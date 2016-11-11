package rpc

import (
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
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
	BankBuildingLevel    int16           `json:"bank_building_level"`    //钱庄等级
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

func (this *RemoteServe) XdgmCliqueList(args *Args_XdgmCliqueList, reply *Reply_XdgmCliqueList) error {
	return Remote.Serve(mdb.RPC_Remote_XdgmCliqueList, args, mdb.TRANS_TAG_RPC_Serve_XdgmCliqueList, func() error {
		var total int16

		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.Select.GlobalClique(func(row *mdb.GlobalCliqueRow) {
				total++
				if total-1 < args.Offset || total > args.Offset+args.Limit {
					return
				}
				clique := &CliqueInfo{
					CliqueId:             row.Id(),
					Name:                 row.Name(),
					TotalDonateCoins:     row.TotalDonateCoins(),
					OwnerPid:             row.OwnerPid(),
					CenterBuildingLevel:  row.CenterBuildingLevel(),
					TempleBuildingLevel:  row.TempleBuildingLevel(),
					BankBuildingLevel:    row.BankBuildingLevel(),
					HealthBuildingLevel:  row.HealthBuildingLevel(),
					AttackBuildingLevel:  row.AttackBuildingLevel(),
					DefenseBuildingLevel: row.DefenseBuildingLevel(),
				}

				ownrInfo := global.GetPlayerInfo(row.OwnerPid())
				clique.OwnerNick = string(ownrInfo.PlayerNick)
				memberPids := module.CliqueRPC.CliqueInfoListPid(row.Id())
				for _, pid := range memberPids {
					playerInfo := global.GetPlayerInfo(pid)
					clique.Members = append(clique.Members, &CliqueMember{
						Pid:   playerInfo.PlayerId,
						Nick:  string(playerInfo.PlayerNick),
						Level: playerInfo.RoleLevel,
					})
				}
				reply.CliqueList = append(reply.CliqueList, clique)
			})
			reply.Total = total
			reply.Offset = args.Offset
		})
		return nil
	})
}
