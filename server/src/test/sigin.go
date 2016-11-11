package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"

	"client_test"
	"encoding/base64"
	"game_server/api/protocol/daily_sign_in_api"
	"game_server/api/protocol/item_api"
	"game_server/api/protocol/notify_api"
	"game_server/api/protocol/player_api"
	"game_server/api/protocol/town_api"
	"game_server/module/player"
)

var (
	GSAddr   string
	exitChan chan bool
)

const (
	MAIN_ROLE = 1
)

type Robot struct {
	id   int
	pid  int64
	name string

	gs *client_test.Client
	hd *client_test.Client

	callback func()
	callOnce func()

	Player PlayerData
}

type PlayerData struct {
	FuncKey int16 //功能权值
	RoleLv  int16 //主角等级
}

func (this *Robot) Action(action func(*client_test.Client, *Robot), gameserver bool, sleep int32) *Robot {
	oldCallBack := this.callback
	this.callback = func() {
		if oldCallBack != nil {
			oldCallBack()
		}
		if gameserver {
			action(this.gs, this)
		} else {
			action(this.hd, this)
		}
		time.Sleep(time.Duration(sleep) * time.Second)
	}
	return this
}

func (this *Robot) ActOnce(action func(*client_test.Client, *Robot), gameserver bool, sleep int32) *Robot {
	oldCallBack := this.callOnce
	this.callOnce = func() {
		if oldCallBack != nil {
			oldCallBack()
		}
		if gameserver {
			action(this.gs, this)
		} else {
			action(this.hd, this)
		}
	}
	return this
}

func NewRobot() *Robot {
	robot := &Robot{gs: client_test.NewClient(GSAddr)}

	robot.gs.OutPlayer_FromPlatformLogin = func(out *player_api.FromPlatformLogin_Out) {
		robot.pid = out.PlayerId
		robot.gs.Player_Info()
	}

	robot.gs.OutPlayer_Info = func(out *player_api.Info_Out) {
		robot.Player.FuncKey = out.FuncKey
		robot.Player.RoleLv = out.RoleLv
	}

	robot.gs.OutPlayer_NotifyGlobalServerInfo = func(out *player_api.NotifyGlobalServerInfo_Out) {
		// robot.hd = client_test.NewClient(string(out.Addr))
		robot.gs.Town_Enter(1)
	}

	robot.gs.OutTown_Enter = func(out *town_api.Enter_Out) {

		// robot.hd.Player_GlobalLogin(robot.pid, robot.name, robot.name, 1, 1, 920, time.Now().Unix(), "")
		if out.Player.PlayerId != robot.pid {
			return
		}

		go func() {
			if robot.callOnce != nil {
				robot.callOnce()
			}

			if robot.callback == nil {
				return
			}

			for {
				robot.callback()
			}
		}()
	}

	return robot
}

func (this *Robot) Login(id int) {
	this.name = fmt.Sprintf("robot_%d", id)
	this.gs.Name = this.name
	nowTime := time.Now().Unix()

	hash := base64.StdEncoding.EncodeToString(player.GetPlatformLoginHash(1, MAIN_ROLE, []byte(this.name), []byte(this.name), nowTime))
	this.gs.Player_FromPlatformLogin(1, this.name, this.name, MAIN_ROLE, hash, nowTime, false /*restore*/, 0 /*recvCount*/, "" /*openkey*/, "" /*pay_token*/, "" /*pfkey*/, 0 /*zoneid*/, "" /*pf*/, 0 /*channel*/, "" /*telecom_oper*/, "" /*client_version*/, "" /*system_hardware*/, "" /*network*/)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 2)
	exitChan = make(chan bool, 2)

	if len(os.Args) == 1 {
		fmt.Println("usage:\n./robot \"127.0.0.1:8080\" 3000")
		fmt.Println(" - 127.0.0.1:8080 (gs addr)")
		fmt.Println(" - 3000 (robot number)")
		return
	}

	GSAddr = os.Args[1]
	robotNum, _ := strconv.ParseInt(os.Args[2], 10, 32)

	for i := 1; i <= int(robotNum); i++ {
		func(id int, robot *Robot) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("robot %s panic. %v\n", robot.name, err)
				}
			}()
			robot.ActOnce(ActOnce_SigninInfo, true, 1)
			robot.ActOnce(ActOnce_DailySign, true, 1)

			robot.Login(i)

			// robot.ActOnce(ActOnce_OpenAllFunc, true, 1)
			// robot.ActOnce(ActOnce_SetMainRoleLevelTo25, true, 1)
			// robot.ActOnce(ActOnce_AddSword, true, 1)
			// robot.Action(Action_Equiement, true, 1)
			//robot.Action(Action_Town_Move, true, 1)
		}(i, NewRobot())
	}

	<-exitChan
	<-exitChan
}

//在城镇移动
func Action_Town_Move(client *client_test.Client, robot *Robot) {
	x := int16(rand.Int() % 700)
	y := int16(rand.Int() % 400)
	client.Town_Move(1200+x, 700+y)
}

//开启所有功能
func ActOnce_OpenAllFunc(client *client_test.Client, robot *Robot) {
	if robot.Player.FuncKey < 1500 {
		client.Debug_OpenFunc(1500)
		//TODO 收到通知需要更新 robot.Player.FuncKey
	}
	client.OutNotify_FuncKeyChange = func(out *notify_api.FuncKeyChange_Out) {
		robot.Player.FuncKey = out.FuncKey
	}
}

//角色升到25级
func ActOnce_SetMainRoleLevelTo25(client *client_test.Client, robot *Robot) {
	if robot.Player.RoleLv < 24 && robot.Player.RoleLv > 0 {
		client.Debug_SetRoleLevel(MAIN_ROLE, 24)
	}
	client.OutNotify_RoleExpChange = func(out *notify_api.RoleExpChange_Out) {
		robot.Player.RoleLv = out.Level
	}
}

//获得装备
func ActOnce_AddSword(client *client_test.Client, robot *Robot) {
	client.OutItem_GetAllItems = func(out *item_api.GetAllItems_Out) {
		var hasSword bool = false //拥有鹰扬剑
		var hasDress bool = false //拥有鹰扬袍
		var hasShoe bool = false  //拥有鹰扬鞋
		var hasRing bool = false  //拥有鹰扬戒
		for _, item := range out.Items {
			if item.ItemId == 31 {
				hasSword = true
			}
			if item.ItemId == 36 {
				hasDress = true
			}
			if item.ItemId == 37 {
				hasShoe = true
			}
			if item.ItemId == 38 {
				hasRing = true
			}
		}
		if !hasSword {
			client.Debug_AddItem(31, 5)
		}
		if !hasDress {
			client.Debug_AddItem(36, 5)
		}
		if !hasShoe {
			client.Debug_AddItem(37, 5)
		}
		if !hasRing {
			client.Debug_AddItem(38, 5)
		}
	}
	client.Item_GetAllItems()
}

//装备武器 然后卸载
func Action_Equiement(client *client_test.Client, robot *Robot) {
	client.OutItem_GetAllItems = func(out *item_api.GetAllItems_Out) {
		var swordUID int64
		var dressUID int64
		var shoeUID int64
		var ringUID int64
		for _, item := range out.Items {
			if item.ItemId == 31 {
				swordUID = item.Id
			}
			if item.ItemId == 36 {
				dressUID = item.Id
			}
			if item.ItemId == 37 {
				shoeUID = item.Id
			}
			if item.ItemId == 38 {
				ringUID = item.Id
			}
		}
		fmt.Printf("%d %d %d %d\n", swordUID, dressUID, shoeUID, ringUID)
		if swordUID > 0 {
			client.Item_Dress(MAIN_ROLE, swordUID)
		}
		if shoeUID > 0 {
			client.Item_Dress(MAIN_ROLE, shoeUID)
		}
		if dressUID > 0 {
			client.Item_Dress(MAIN_ROLE, dressUID)
		}
		if ringUID > 0 {
			client.Item_Dress(MAIN_ROLE, ringUID)
		}
	}
	client.Item_GetAllItems()
}

func ActOnce_SigninInfo(client *client_test.Client, robot *Robot) {
	client.OutDailySignIn_Info = func(out *daily_sign_in_api.Info_Out) {
		PrettyPrint(out)
		exitChan <- true
	}
	client.DailySignIn_Info()
}

func ActOnce_DailySign(client *client_test.Client, robot *Robot) {
	client.OutDailySignIn_Sign = func(out *daily_sign_in_api.Sign_Out) {
		PrettyPrint(out)
		exitChan <- true
	}
	client.DailySignIn_Sign()
}

/*
func ActOnce_GetGhost(client *client_test.Client, robot *Robot) {
	client.OutGhost_Info = func(out *ghost_api.Info_Out) {
	}
}

func Action_EquipGhost(client *client_test.Client, robot *Robot) {
}
*/

func PrettyPrint(data interface{}) {
	fmt.Println("\n==============")
	bytes, _ := json.MarshalIndent(data, "", "\t")
	fmt.Printf(string(bytes))
	fmt.Println("==============\n")
}
