package main

import (
	"client_test"
	"core/tprof"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

import (
	"game_server/api/protocol/player_api"
)

type SlowLog []slowLogItem

type slowLogItem struct {
	Threshold time.Duration
	Count     int64
}

func NewSlowLog(thresholds []time.Duration) SlowLog {
	log := make(SlowLog, len(thresholds))
	for i := 0; i < len(log); i++ {
		log[i].Threshold = thresholds[i]
	}
	return log
}

func (log SlowLog) Record(d time.Duration) {
	for i := 0; i < len(log); i++ {
		if log[i].Threshold <= d {
			atomic.AddInt64(&log[i].Count, 1)
			break
		}
	}
}

func (log SlowLog) Save(w io.Writer) {
	fmt.Fprintln(w, "Threshold,Count")
	for i := 0; i < len(log); i++ {
		fmt.Fprintf(w, "%s,%d\n", log[i].Threshold.String(), atomic.LoadInt64(&log[i].Count))
	}
}

var prof = tprof.New()

var slow = NewSlowLog([]time.Duration{
	time.Hour,
	time.Millisecond * 2000,
	time.Millisecond * 1000,
	time.Millisecond * 800,
	time.Millisecond * 600,
	time.Millisecond * 400,
	time.Millisecond * 200,
	time.Millisecond * 80,
	time.Millisecond * 60,
	time.Millisecond * 40,
	time.Millisecond * 20,
	time.Nanosecond * 800,
	time.Nanosecond * 600,
	time.Nanosecond * 400,
	time.Nanosecond * 200,
	time.Nanosecond * 80,
	time.Nanosecond * 60,
	time.Nanosecond * 40,
	time.Nanosecond * 20,
	time.Nanosecond,
})

func main() {
	rand.Seed(time.Now().UnixNano())

	var (
		serverAddr   = os.Getenv("BM_ADDR")                 // 服务器地址
		clientNum, _ = strconv.Atoi(os.Getenv("BM_NUM"))    // 客户端数量
		clientPrefix = os.Getenv("BM_PREFIX")               // 账号前缀
		threadNum, _ = strconv.Atoi(os.Getenv("BM_THREAD")) // 线程数量，当type == loop的时候对应在线玩家数量
		testMode     = os.Getenv("BM_MODE")                 // 测试模式，如果是 loop则不会持续创建角色，保持连接持续请求
		testType     = os.Getenv("BM_TYPE")                 // 测试类型，1-获取用户信息，2-心跳包
		thinkTime, _ = strconv.Atoi(os.Getenv("BM_THINK"))  // 每次请求间的思考时间
	)

	if clientPrefix == "" {
		clientPrefix = strconv.Itoa(int(time.Now().UnixNano()))
	}

	if threadNum == 0 {
		threadNum = 1
	}

	for m := 0; m < threadNum; m++ {
		go func(m int) {
			for i := m * (clientNum / threadNum); i < (m+1)*(clientNum/threadNum); i++ {
				println(i)
				client := newClient(i, serverAddr, clientPrefix)

				for testMode == "loop" {
					switch testType {
					case "1":
						doRequest1(client)
					case "2":
						doRequest2(client)
					}

					// 模拟思考时间
					t1 := time.Now()
					time.Sleep(time.Millisecond * time.Duration(rand.Intn(thinkTime)))
					prof.Record("think", time.Since(t1))
				}

				client.Close()
			}
		}(m)
	}

	fmt.Scanln()

	if file, e := os.Create("benchmark_tprof.csv"); e == nil {
		prof.Save(file)
		file.Close()
	} else {
		println(e.Error())
	}

	if file, e := os.Create("benchmark_slow.csv"); e == nil {
		slow.Save(file)
		file.Close()
	} else {
		println(e.Error())
	}
}

func newClient(i int, serverAddr string, clientPrefix string) *client_test.Client {
	username := clientPrefix + strconv.Itoa(i)

	client := client_test.NewClient(serverAddr)

	var wg sync.WaitGroup
	wg.Add(1)
	t1 := time.Now()
	client.OutPlayer_FromPlatformLogin = func(out *player_api.FromPlatformLogin_Out) {
		d1 := time.Since(t1)
		prof.Record("login", d1)
		slow.Record(d1)
		if out.Status != player_api.LOGIN_STATUS_SUCCEED && out.Status != player_api.LOGIN_STATUS_FIRST_TIME {
			log.Printf("login failed (%d)", out.Status)
			client = nil
		}
		wg.Done()
	}
	client.Player_FromPlatformLogin(
		player_api.PLATFORM_TYPE_MOBILE_IOS_GUEST, // platform_type,
		player_api.CHANNEL_ID_MOBILE,
		username, // username,
		username, // nickname,
		1,        // role_id,
		"",       // hashcode,
		0,        // unixtime,
		false,    // restore,
		0,        // recvCount,
		0,        // gsid,
		"",       // openkey,
		"",       // pay_token,
		"",       // pfkey,
		0,        // zoneid,
		"",       // pf,
		0,        // channel,
		"",       // telecom_oper,
		"",       // client_version,
		"",       // system_hardware,
		"",       // network,
		"",
	)
	wg.Wait()

	return client
}

func doRequest1(client *client_test.Client) {
	var wg sync.WaitGroup
	wg.Add(1)
	t1 := time.Now()
	client.OutPlayer_Info = func(out *player_api.Info_Out) {
		d1 := time.Since(t1)
		prof.Record("info", d1)
		slow.Record(d1)
		wg.Done()
	}
	client.Player_Info()
	wg.Wait()
}

func doRequest2(client *client_test.Client) {
	var wg sync.WaitGroup
	wg.Add(1)
	t1 := time.Now()
	client.OutPlayer_GetTime = func(out *player_api.GetTime_Out) {
		d1 := time.Since(t1)
		prof.Record("gettime", d1)
		slow.Record(d1)
		wg.Done()
	}
	client.Player_GetTime()
	wg.Wait()
}
