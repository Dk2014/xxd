package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"
)

import (
	"core/debug"
	"core/i18l"
	"core/log"
	"core/mysql"
	"core/syncutil"
	util "core/time"
	"game_server/api"
	. "game_server/config"
	"game_server/dat"
	"game_server/mdb"
	"game_server/tlog"
	"game_server/web"
	"game_server/xdlog"
	golog "log"
	godebug "runtime/debug"
)

import (
	_ "game_server/module/announcement_rpc"
	_ "game_server/module/arena"
	_ "game_server/module/arena_rpc"
	_ "game_server/module/battle"
	_ "game_server/module/battle_pet"
	_ "game_server/module/channel_rpc"
	_ "game_server/module/clique_building_rpc"
	_ "game_server/module/clique_rpc"
	_ "game_server/module/daily_sign_in"
	_ "game_server/module/debug"
	_ "game_server/module/fashion"
	_ "game_server/module/friend_rpc"
	_ "game_server/module/ghost"
	_ "game_server/module/heart"
	_ "game_server/module/item"
	_ "game_server/module/mail"
	_ "game_server/module/mail_rpc"
	_ "game_server/module/meditation"
	_ "game_server/module/mission"
	_ "game_server/module/multi_level"
	_ "game_server/module/notify"
	_ "game_server/module/pet_virtual_env"
	_ "game_server/module/physical"
	_ "game_server/module/player"
	_ "game_server/module/push_notify"
	_ "game_server/module/quest"
	_ "game_server/module/rainbow"
	_ "game_server/module/role"
	_ "game_server/module/skill"
	_ "game_server/module/sword_soul"
	_ "game_server/module/team"
	_ "game_server/module/totem"
	//_ "game_server/module/tower"
	_ "game_server/module/awaken"
	_ "game_server/module/clique_escort_rpc"
	_ "game_server/module/clique_quest_rpc"
	_ "game_server/module/draw"
	_ "game_server/module/driving_sword"
	_ "game_server/module/money_tree"
	_ "game_server/module/town"
	_ "game_server/module/trader"
	_ "game_server/module/vip"
)

import (
	"game_server/api/protocol/server_info_api"
	"game_server/global"
	"game_server/module"
	"game_server/module/event"
	"game_server/module/rpc"
	"game_server/module/server_info"
)

import (
	"io/ioutil"
	"math"
	"regexp"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
)

var (
	exitChan                           = make(chan bool)
	iosTotalOnline, androidTotalOnline int64
)

func main() {
	rand.Seed(time.Now().UnixNano())

	runtime.GOMAXPROCS(runtime.NumCPU())

	logPrefix := "gs"
	if ServerCfg.EnableGlobalServer {
		logPrefix = "hd"
	}
	log.SetLogFilePrefix(logPrefix)
	log.Setup(ServerCfg.LogDir, true)
	if !ServerCfg.EnableDebugAPI { //verbose debug info warn error EnableDebugAPI == false 意味正式环境，正式环境不打印 verbose
		log.SetLogLevel(log.LEVEL_DEBUG)
	}

	dbconfig := GetDBConfig()
	mdb.Start(ServerCfg.EnableGlobalServer, ServerCfg.ServerId, dbconfig, ServerCfg.DBFileDir, ServerCfg.TlogFileDir, ServerCfg.XdlogFileDir, ServerCfg.EnableXDlog, ServerCfg.EnableTlog)

	// 语言包加载
	i18l.T = i18l.NewLocale(ServerCfg.LocalConf, "zh")
	if ServerCfg.CurrLocal != "" {
		i18l.T.SetLocale(ServerCfg.CurrLocal)
	}

	dat.Load(dbconfig)

	// 启动时指定了运营配置文件（该操作不能在dat.Load之前进行）
	//if Eventconf != "" {
	event.LoadEventExtend(Eventconf)
	//}

	// 开启互动服务
	if ServerCfg.EnableGlobalServer {
		global.Start(dbconfig)
		module.PostManStart()
		module.BroadcastChannelStart()
		onlineCount()
		module.Cliquesalary()
		module.PlayerRPC.LoadGlobal()
		if ServerCfg.EnableGiftCode {
			mdb.Transaction(mdb.TRANS_TAG_CMD, func() {
				mdb.GlobalExecute(func(db *mdb.Database) {
					//对 LoadGiftCodeFromDB 作修改请看函数实现处的注释
					global.LoadGiftCodeFromDB(db, int(ServerCfg.ServerId/10), GetGiftCodeDBConfig())
				})
			})
		}
	}

	serverId := fmt.Sprintf("%d", ServerCfg.ServerId)
	if ServerCfg.EnableTlog {
		if err := tlog.Start(ServerCfg.TlogServerAddr, ServerCfg.GameAppID, serverId, ServerCfg.TlogPlatformId); err != nil {
			panic(err)
		}
	}

	api.Start("tcp4", fmt.Sprintf("0.0.0.0:%d", ServerCfg.GamePort))
	rpc.Remote.Start(fmt.Sprintf("0.0.0.0:%d", ServerCfg.GamePort+1))

	if ServerCfg.DebugWebServerAddr != "" {
		go web.StartDebugServer(fmt.Sprintf("%s:%d", ServerCfg.DebugWebServerAddr, ServerCfg.GamePort+2))
	}

	//日志中的游戏服ID 等于游戏服务器ID/10
	xdlog.Init(int32(ServerCfg.ServerId / 10))

	go sighandler()

	<-exitChan

	api.Stop()
	mdb.Stop(ServerCfg.EnableXDlog, ServerCfg.EnableTlog)
	tlog.Stop()
}

func sighandler() {
	defer func() { recover() }()
	signalTermChan := make(chan os.Signal, 1)
	signalTermUSR1 := make(chan os.Signal, 1)

	signal.Notify(signalTermChan, syscall.SIGTERM)
	signal.Notify(signalTermUSR1, syscall.SIGUSR1)

	for {
		select {
		case <-signalTermChan:
			exitChan <- true
		case <-signalTermUSR1:
			go readAndExecuteCmdFromFile("cmd.in")
		}
	}
}

var onlineDB *mysql.Connection
var onlineStmt *mysql.Stmt

func onlineCount() {
	time.AfterFunc(time.Minute*1, func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("onlineCount() error: ", err)
			}
		}()

		var countmap map[int32]int64

		defer onlineCount()

		rpc.RemoteGetOnlineNumber(func(back map[int32]int64) {
			countmap = back
			var err error
			if onlineDB == nil {
				if onlineDB, err = mysql.Connect(GetOnlineDBConfig()); err != nil {
					log.Errorf("onlineCount connect mysql error: ", err)
					return
				}
				onlineStmt = mdb.PrepareOnlineInfoToOnlineTable(onlineDB, ServerCfg.OnlineDB.TableName)
			}

			timeKey := math.Floor(float64(util.GetNowTime()) / 60)
			for k, v := range countmap {
				//腾讯
				if k == 0 {
					if ServerCfg.TlogPlatformId == 0 { // ios
						iosTotalOnline = v
						androidTotalOnline = 0
					} else {
						iosTotalOnline = 0
						androidTotalOnline = v
					}
				} else if k == 1 || k == 3 {
					iosTotalOnline = v
					androidTotalOnline = 0
				} else if k == 2 || k == 4 {
					iosTotalOnline = 0
					androidTotalOnline = v
				}
				log.Infof("Cid: %d Total Online: %d(ios), %d(android)\n", k, iosTotalOnline, androidTotalOnline)
				err = mdb.DoInsertOnline(onlineStmt, ServerCfg.GameAppID, int64(timeKey), int64(ServerCfg.ServerId), iosTotalOnline, androidTotalOnline, int64(k))
				if err != nil {
					log.Errorf("onlineCount::DoInsertOnline error: ", err)
					onlineStmt.Close()
					onlineDB.Close()
					onlineDB = nil
				}
			}
		})
	})
}

func readAndExecuteCmdFromFile(file string) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf(`readAndExecuteCmdFromFile
Error = %v
Stack = 
%s`,
				err,
				debug.Stack(1, "    "),
			)
		}
	}()

	cmd, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
		return
	}

	var cpuProfile *os.File

	cmdStr := strings.Trim(string(cmd), "\t\r\n ")

	switch cmdStr {
	case "online":
		golog.Printf("在线人数 %d\n", module.Player.OnlinePlayerNum())

	case "exit":
		exitChan <- true

	case "reload":
		mdb.Transaction(mdb.TRANS_TAG_CMD, func() {
			dat.Load(GetDBConfig())
		})

	case "trace stop":
		module.Tracer.SetTarget(-1)

	case "print go num":
		log.Infof("Num Goroutine: %v", runtime.NumGoroutine())

	case "print go":
		p := pprof.Lookup("goroutine")
		if file, err := os.Create(fmt.Sprintf("gs_%d.goroutine", ServerCfg.ServerId)); err == nil {
			p.WriteTo(file, 2)
			file.Close()
		} else {
			log.Errorf("[print go] error. %v", err)
		}

	case "print block":
		p := pprof.Lookup("block")
		if file, err := os.Create(fmt.Sprintf("gs_%d.block", ServerCfg.ServerId)); err == nil {
			p.WriteTo(file, 2)
			file.Close()
		} else {
			log.Errorf("[print block] error. %v", err)
		}

	case "start cpuprof":
		if cpuProfile == nil {
			if f, err := os.Create(fmt.Sprintf("gs_%d.cpuprof", ServerCfg.ServerId)); err != nil {
				log.Errorf("start cpu profile failed: %v", err)
			} else {
				log.Info("start cpu profile")
				pprof.StartCPUProfile(f)
				cpuProfile = f
			}
		}

	case "stop cpuprof":
		if cpuProfile != nil {
			pprof.StopCPUProfile()
			cpuProfile.Close()
			cpuProfile = nil
			log.Info("stop cpu profile")
		}

	case "get memprof":
		if f, err := os.Create(fmt.Sprintf("gs_%d.memprof", ServerCfg.ServerId)); err != nil {
			golog.Printf("record memory profile failed: %v\n", err)
		} else {
			pprof.WriteHeapProfile(f)
			f.Close()
			golog.Println("record memory profile")
		}

	case "lookup heap":
		p := pprof.Lookup("heap")
		if file, err := os.Create(fmt.Sprintf("gs_%d.heap", ServerCfg.ServerId)); err == nil {
			p.WriteTo(file, 2)
			file.Close()
		} else {
			golog.Println(err)
		}

	case "gc summary":
		golog.Println("excute gc summary:")
		PrintGCSummary()

	case "get apiprof":
		api.GetAPIProfile()

	case "debug lock on":
		syncutil.EnableLockDebug()

	case "debug lock off":
		syncutil.DisableLockDebug()

	case "lock dump":
		if file, err := os.Create(fmt.Sprintf("gs_%d.lockdump", ServerCfg.ServerId)); err == nil {
			file.Write(syncutil.Dump())
			file.Close()
		} else {
			golog.Println(err)
		}

	case "srvprofile":
		onlineStr := fmt.Sprintf("Total Online: %d(gs-%d), %d(ios), %d(android)\n",
			module.Player.OnlinePlayerNum(), ServerCfg.ServerId, iosTotalOnline, androidTotalOnline)

		api.GetAPIProfile()

		if file, err := os.Create(fmt.Sprintf("gs_%d_%d.srvprofile", ServerCfg.ServerId, util.GetNowTime())); err == nil {
			file.WriteString(onlineStr)
			mdb.ProfileFileCommitToFile(file, ServerCfg.EnableXDlog, ServerCfg.EnableTlog)
			file.Close()
		} else {
			log.Info(onlineStr)
			mdb.ProfileFileCommitToLog(ServerCfg.EnableXDlog, ServerCfg.EnableTlog)
		}

	case "reload_srv_version": // 刷新服务端版本
		server_info.RefreshServerVersion()

	case "reload_conf": // 部分conf配置文件支持热更新
		var localConf ServerConfig
		LoadJSONConfig(ServerConf, &localConf)

		ServerCfg.TssServerId = localConf.TssServerId
		ServerCfg.GlobalServerAddr = localConf.GlobalServerAddr

		ServerCfg.MSDKApiAddr = localConf.MSDKApiAddr
		ServerCfg.TlogPlatformId = localConf.TlogPlatformId

		ServerCfg.XGSdk = localConf.XGSdk
		ServerCfg.MoneySdk = localConf.MoneySdk
		ServerCfg.OnlineDB = localConf.OnlineDB
		ServerCfg.RPCServerList = localConf.RPCServerList

		tlog.UpdateTlogPlatId(ServerCfg.TlogPlatformId)

		mdb.Transaction(mdb.TRANS_TAG_CMD, func() {
			rpc.Remote.RefreshRPCServerList()
		})

	default:
		parseCmd(`give\s([0-9]+)\s([0-9]+)\s([0-9]+)`, cmdStr, func(matches []string) {
			pid, _ := strconv.ParseInt(matches[1], 10, 64)
			itemId, _ := strconv.ParseInt(matches[2], 10, 64)
			num, _ := strconv.ParseInt(matches[3], 10, 64)
			fmt.Printf("give pid(%d) item_id(%d) num(%d)\n", pid, itemId, num)
			mdb.Transaction(mdb.TRANS_TAG_CMD, func() {
				mdb.GlobalExecute(func(globalDB *mdb.Database) {
					globalDB.AgentExecute(pid, func(agentDB *mdb.Database) {
						module.Item.AddItem(agentDB, int16(itemId), int16(num), 0, 0, "")
					})
				})
			})
		})
		parseCmd(`^trace\s(up|down)\s([0-9]+)\s(\-?[0-9]+)$`, cmdStr, func(matches []string) {
			mid, err := strconv.ParseInt(matches[2], 10, 64)
			if err != nil {
				panic(err)
			}
			aid, err := strconv.ParseInt(matches[3], 10, 64)
			if err != nil {
				panic(err)
			}
			if matches[1] == "up" {
				module.SetTraceUp(int8(mid), int8(aid))
			} else if matches[1] == "down" {
				module.SetTraceDown(int8(mid), int8(aid))
			}

		})
		// trace带参数，如 trace 123456
		parseCmd(`^trace\s([0-9]+)$`, cmdStr, func(matches []string) {
			playerId, err := strconv.ParseInt(matches[1], 10, 64)
			if err != nil {
				panic(err)
			}
			module.Tracer.SetTarget(playerId)
		})

		// loadevent带参数, 如 loadevent /root/xx/xxx/ev_001.json
		// 加载运营数据配置
		parseCmd(`^loadevent\s([/\w.-]+\.json$)`, cmdStr, func(matches []string) {
			mdb.Transaction(mdb.TRANS_TAG_CMD, func() {
				event.LoadEventExtend(matches[1])
				// 通知在线玩家
				module.API.Broadcast(module.Player, &server_info_api.UpdateEventData_Out{
					Json: event.EventConfigRawData,
				})
			})
		})

		// query player mdb xxxx xxxxxx, 后跟要查得表名 pid key
		// 查询内存中的玩家数据
		parseCmd(`^query\splayer\smdb\s([0-9\w\_]+)\s([0-9]+)\s([0-9]+)`, cmdStr, func(matches []string) {
			mdb.Transaction(mdb.TRANS_TAG_CMD, func() {
				tableName := matches[1]
				pid, err := strconv.ParseInt(matches[2], 10, 64)
				if err != nil {
					panic(err)
				}
				key, err := strconv.ParseInt(matches[3], 10, 64)
				if err != nil {
					panic(err)
				}
				mdb.GlobalExecute(func(globalDB *mdb.Database) {
					globalDB.AgentExecute(pid, func(agenDB *mdb.Database) {
						lookupOperate := agenDB.Lookup
						// reflect
						in := make([]reflect.Value, 0)
						in = append(in, reflect.ValueOf(key))
						result := reflect.ValueOf(lookupOperate).MethodByName(tableName).Call(in)
						if len(result) > 0 {
							structType := reflect.TypeOf(result[0].Elem().Interface())
							for i := 0; i < result[0].Elem().NumField(); i++ {
								val := result[0].Elem().Field(i)
								log.Infof("%s : %v", structType.Field(i).Name, val.Interface())
							}
						}
					})
				})
			})
		})

		log.Infof("######### success execute cmd: %s #########", cmdStr)
	}
}

func parseCmd(cmdExpr string, cmdStr string, execute func(matches []string)) {
	exp := regexp.MustCompile(cmdExpr)
	matches := exp.FindStringSubmatch(cmdStr)
	if len(matches) > 0 {
		execute(matches)
	}
}

var startTime = time.Now()

func PrintGCSummary() {
	memStats := &runtime.MemStats{}
	runtime.ReadMemStats(memStats)
	gcstats := &godebug.GCStats{PauseQuantiles: make([]time.Duration, 100)}
	godebug.ReadGCStats(gcstats)

	printGC(memStats, gcstats)
}

func printGC(memStats *runtime.MemStats, gcstats *godebug.GCStats) {
	if gcstats.NumGC > 0 {
		lastPause := gcstats.Pause[0]
		elapsed := time.Now().Sub(startTime)
		overhead := float64(gcstats.PauseTotal) / float64(elapsed) * 100
		allocatedRate := float64(memStats.TotalAlloc) / elapsed.Seconds()

		golog.Printf("NumGC:%d Pause:%s Pause(Avg):%s Overhead:%3.2f%% Alloc:%s Sys:%s Alloc(Rate):%s/s Histogram:%s %s %s \n",
			gcstats.NumGC,
			util.Time(lastPause),
			util.Time(avg(gcstats.Pause)),
			overhead,
			util.Size(memStats.Alloc),
			util.Size(memStats.Sys),
			util.Size(uint64(allocatedRate)),
			util.Time(gcstats.PauseQuantiles[94]),
			util.Time(gcstats.PauseQuantiles[98]),
			util.Time(gcstats.PauseQuantiles[99]))
	} else {
		// while GC has disabled
		elapsed := time.Now().Sub(startTime)
		allocatedRate := float64(memStats.TotalAlloc) / elapsed.Seconds()

		golog.Printf("Alloc:%s Sys:%s Alloc(Rate):%s/s\n",
			util.Size(memStats.Alloc),
			util.Size(memStats.Sys),
			util.Size(uint64(allocatedRate)))
	}
}

func avg(items []time.Duration) time.Duration {
	var sum time.Duration
	for _, item := range items {
		sum += item
	}
	return time.Duration(int64(sum) / int64(len(items)))
}
