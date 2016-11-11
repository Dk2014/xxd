package rpc

import (
	"game_server/mdb"
)

type Args_XdgmCliqueList struct {
	RPCArgTag
	Limit  int16
	Offset int16
}

type CliqueMember struct {
	Pid   int64  `json:"pid"`   //pid
	Nick  string `json:"nick"`  //nick
	Level int16  `json:"level"` //level
}

type CliqueInfo struct {
	CliqueId             int64           `json:"clique_id"`              //帮派ID
	Name                 string          `json:"name"`                   //帮派名称
	TotalDonateCoins     int64           `json:"total_donate_coins"`     //帮派累计捐献
	OwnerLoginTime       int64           `json:"owner_login_time"`       //帮主最近登录时间
	OwnerPid             int64           `json:"owner_pid"`              //帮主PID
	OwnerNick            string          `json:"owner_nick"`             //帮主昵称
	CenterBuildingLevel  int16           `json:"center_building_level"`  //总舵等级
	TempleBuildingLevel  int16           `json:"temple_building_level"`  //宗祠等级
	BankBuildingLevel    int16           `json:"bank_building_levej"`    //钱庄等级
	HealthBuildingLevel  int16           `json:"health_building_level"`  //回春堂等级
	AttackBuildingLevel  int16           `json:"attack_building_level"`  //神兵堂等级
	DefenseBuildingLevel int16           `json:"defense_building_level"` //金刚堂等级
	Members              []*CliqueMember `json:"members"`                //成员列表
}

type Reply_XdgmCliqueList struct {
	Total      int16         `json:"total"`
	Offset     int16         `json:"offset"`
	CliqueList []*CliqueInfo `json:"clique_list"`
}

func RemoteXdgmCliqueList(sid int, offset, limit int16, callback func(*Reply_XdgmCliqueList, error)) {
	reply := &Reply_XdgmCliqueList{}
	args := &Args_XdgmCliqueList{
		Offset: offset,
		Limit:  limit,
	}
	Remote.Call(sid, mdb.RPC_Remote_XdgmCliqueList, args, reply, func(err error) {
		callback(reply, err)
	})
}
