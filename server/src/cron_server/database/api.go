package database

import (
	"encoding/json"
	"fmt"
	"time"
)

type Player struct {
	Tag          string `json:"tag"`
	Cid          int32  `json:"cid"`          // (必填)平台渠道ID
	Sid          int32  `json:"sid"`          // (必填)区唯一ID,如果不分区则为0
	Account      string `json:"account"`      // (必填)平台帐号
	Pid          int64  `json:"pid"`          // (必填)玩家角色ID
	Name         string `json:"name"`         // (必填)玩家名
	Coins        int64  `json:"coins"`        // (必填)当前剩余的全部游戏币数量
	Charge_coins int64  `json:"charge_coins"` // (必填)剩余的充值的游戏币数量
	Login_time   int64  `json:"login_time"`   // (必填)最后登录时间
	Item         string `json:"item"`         // (必填)对应物品id的json数据
	Level        string `json:"level"`        // (必填)对应角色等级的json数据
	Vip_level    int32  `json:"vip_level"`    // (必填)VIP等级
}

type Item struct {
	ItemId int32
	Num    int32
}

type Role struct {
	RoleId int32
	Level  int32
}

func FetchPlayer(sid int32) (results []*Player) {
	defer func() {
		if g_db != nil {
			g_db.Close()
		}
	}()
	rows, err := g_db.Query(fmt.Sprintf(`
	select p.cid as cid,p.user as user,p.id as pid,p.nick as nick,i.coins as coins,i.ingot as ingot,i.last_login_time as last_login_time,v.level as vip_level from player as p 
	left join player_info as i on (p.id = i.pid)
	left join player_vip as v on (p.id = v.pid)`))
	if err != nil {
		panic(err)
	}
	defer func() {
		rows.Close()
	}()
	for rows.Next() {
		var player Player
		rows.Scan(&player.Cid, &player.Account, &player.Pid, &player.Name, &player.Coins, &player.Charge_coins, &player.Login_time, &player.Vip_level)
		player.Tag = "snapshot"
		player.Sid = sid
		results = append(results, &player)
	}
	for _, v := range results {
		rows1, err := g_db.Query(fmt.Sprintf(`select item_id,sum(num) as num from player_item where pid = %d group by item_id`, v.Pid))
		if err != nil {
			panic(err)
		}
		defer func() {
			rows1.Close()
		}()
		item_list := make(map[string]int32)
		for rows1.Next() {
			var item Item
			rows1.Scan(&item.ItemId, &item.Num)
			item_list[fmt.Sprintf("%d", item.ItemId)] = item.Num
		}
		item_str, err := json.Marshal(item_list)
		if err != nil {
			panic(err)
		}
		v.Item = string(item_str)

		rows2, err := g_db.Query(fmt.Sprintf(`select role_id,level from player_role where pid = %d`, v.Pid))
		if err != nil {
			panic(err)
		}
		defer func() {
			rows2.Close()
		}()
		role_list := make(map[string]int32)
		for rows2.Next() {
			var role Role
			rows2.Scan(&role.RoleId, &role.Level)
			role_list[fmt.Sprintf("%d", role.RoleId)] = role.Level
		}
		level_str, err := json.Marshal(role_list)
		if err != nil {
			panic(err)
		}
		v.Level = string(level_str)
	}
	return results
}

type Online struct {
	Tag          string `json:"tag"`
	Cid          int32  `json:"cid"`           // (必填)平台渠道ID
	Sid          int32  `json:"sid"`           // (必填)区唯一ID,如果不分区则为0
	Time         int64  `json:"time"`          // (必填)平台帐号
	CountIos     int64  `json:"count_ios"`     // (必填)ios在线
	CountAndroid int64  `json:"count_android"` // (必填)ios在线
	CountAll     int64  `json:"count_all"`     // (必填)玩家角色ID
}

func FetchOnline(sid int32, tablename string) (results []*Online) {
	rows, err := o_db.Query(fmt.Sprintf(`select cid,onlinecntios,onlinecntandroid from %s where floor(gsid/10) = %d order by id desc limit 1`, tablename, sid))
	if err != nil {
		panic(err)
	}
	defer func() {
		rows.Close()
	}()
	for rows.Next() {
		var online Online
		rows.Scan(&online.Cid, &online.CountIos, &online.CountAndroid)
		online.Tag = "online"
		online.Sid = sid
		nowtime := time.Now()
		online.Time = time.Date(nowtime.Year(), nowtime.Month(), nowtime.Day(), nowtime.Hour(), nowtime.Minute(), 0, 0, time.Local).Unix()
		online.CountAll = online.CountIos + online.CountAndroid
		results = append(results, &online)
	}
	return results
}
