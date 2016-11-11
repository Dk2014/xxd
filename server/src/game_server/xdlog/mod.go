package xdlog

import (
	// "core/fail"
	// "encoding/json"
	// "fmt"
	. "game_server/config"
	"game_server/dat/mission_dat"
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
)

func PlayerLoginLog(db *mdb.Database, ip string, logtype int32) {
	if !ServerCfg.EnableXDlog {
		return
	}
	playerinfo := module.Player.GetPlayer(db)
	main_role_info := module.Role.GetMainRole(db)
	role_vip := module.VIP.VIPInfo(db)
	ispay := ispayed(role_vip.Ingot)

	db.AddXdLog(NewLoginFlow(db.CId(), playerinfo.User, db.PlayerId(), playerinfo.Nick, ip, logtype, ispay, int32(main_role_info.Level), int32(role_vip.Level)))
}

func PlayerCreateLog(db *mdb.Database, ip string, cid int32) {
	if !ServerCfg.EnableXDlog {
		return
	}
	playerinfo := module.Player.GetPlayer(db)
	db.AddXdLog(NewCreateFlow(cid, playerinfo.User, db.PlayerId(), playerinfo.Nick, ip, 1))
}

func PlayerChargeLog(db *mdb.Database, ip, oid, poid, paytype string, amount float64, coins int64, currency string) {
	if !ServerCfg.EnableXDlog {
		return
	}
	playerinfo := module.Player.GetPlayer(db)
	main_role_info := module.Role.GetMainRole(db)
	db.AddXdLog(NewChargeFlow(db.CId(), playerinfo.User, db.PlayerId(), playerinfo.Nick, ip, oid, poid, paytype, int32(main_role_info.Level), amount, coins, currency))
}

func EventLog(db *mdb.Database, eventType int32) {
	if !ServerCfg.EnableXDlog {
		return
	}
	playerinfo := module.Player.GetPlayer(db)
	role_vip := module.VIP.VIPInfo(db)
	ispay := ispayed(role_vip.Ingot)
	db.AddXdLog(NewEventFlow(db.CId(), playerinfo.User, db.PlayerId(), playerinfo.Nick, ispay, eventType, 1))
}

func IncomeLog(db *mdb.Database, value, aftervalue, cointype, eventType int32, eventParam string) {
	if !ServerCfg.EnableXDlog {
		return
	}
	EventLog(db, eventType)
	playerinfo := module.Player.GetPlayer(db)
	role_vip := module.VIP.VIPInfo(db)
	ispay := ispayed(role_vip.Ingot)
	db.AddXdLog(NewIncomeFlow(db.CId(), playerinfo.User, db.PlayerId(), playerinfo.Nick, ispay, value, cointype, eventType, aftervalue, eventParam))
}

func ConsumeLog(db *mdb.Database, value, aftervalue, cointype, eventType int32) {
	if !ServerCfg.EnableXDlog {
		return
	}
	EventLog(db, eventType)
	playerinfo := module.Player.GetPlayer(db)
	role_vip := module.VIP.VIPInfo(db)
	ispay := ispayed(role_vip.Ingot)
	db.AddXdLog(NewConsumeFlow(db.CId(), playerinfo.User, db.PlayerId(), playerinfo.Nick, ispay, value, cointype, eventType, 1, aftervalue))
}

//
func ItemLog(db *mdb.Database, itemid, value int16, eventType int32, eventParam string) {
	if !ServerCfg.EnableXDlog {
		return
	}
	EventLog(db, eventType)
	var aftervalue int32
	playerinfo := module.Player.GetPlayer(db)
	role_vip := module.VIP.VIPInfo(db)
	ispay := ispayed(role_vip.Ingot)
	aftervalue = int32(module.Item.GetItemNum(db, itemid))
	db.AddXdLog(NewItemFlow(db.CId(), playerinfo.User, db.PlayerId(), playerinfo.Nick, ispay, int32(itemid), int32(value), 0, eventType, aftervalue, eventParam))
}

func PropsLog(db *mdb.Database, itemid, value int16, eventType int32) {
	if !ServerCfg.EnableXDlog {
		return
	}
	EventLog(db, eventType)
	var aftervalue int32
	playerinfo := module.Player.GetPlayer(db)
	role_vip := module.VIP.VIPInfo(db)
	ispay := ispayed(role_vip.Ingot)
	aftervalue = int32(module.Item.GetItemNum(db, itemid))
	db.AddXdLog(NewPropsFlow(db.CId(), playerinfo.User, db.PlayerId(), playerinfo.Nick, ispay, int32(itemid), int32(value), eventType, aftervalue))
}

func LevelLog(db *mdb.Database, exp, fromlevel, tolevel int32) {
	if !ServerCfg.EnableXDlog {
		return
	}
	playerinfo := module.Player.GetPlayer(db)
	role_vip := module.VIP.VIPInfo(db)
	ispay := ispayed(role_vip.Ingot)
	db.AddXdLog(NewLevelFlow(db.CId(), playerinfo.User, db.PlayerId(), playerinfo.Nick, exp, fromlevel, tolevel, ispay))
}

func VipLog(db *mdb.Database, ingot, ispay, fromlevel, tolevel int32) {
	if !ServerCfg.EnableXDlog {
		return
	}
	playerinfo := module.Player.GetPlayer(db)
	db.AddXdLog(NewVipFlow(db.CId(), playerinfo.User, db.PlayerId(), playerinfo.Nick, ingot, fromlevel, tolevel, ispay))
}

// func OnlineLog(db *mdb.Database) {
// 	if !ServerCfg.EnableXDlog {
// 		return
// 	}
// 	db.AddXdLog(NewOnlineFlow(1, 1, 1, 1))
// }

// func SnapshotLog(db *mdb.Database) {
// 	if !ServerCfg.EnableXDlog {
// 		return
// 	}
// 	itemmap := make(map[string]int16)
// 	rolemap := make(map[string]int16)
// 	playerinfo := module.Player.GetPlayer(db)
// 	player := db.Lookup.PlayerInfo(playerinfo.Id)
// 	role_vip := module.VIP.VIPInfo(db)
// 	db.Select.PlayerItem(func(row *mdb.PlayerItemRow) {
// 		if row.BuyBackState() == 0 {
// 			itemmap[fmt.Sprintf("%s", row.ItemId())] += row.Num()
// 		}
// 	})
// 	db.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
// 		rolemap[fmt.Sprintf("%s", row.RoleId())] = row.Level()
// 	})
// 	itemjs, err := json.Marshal(itemmap)
// 	fail.When(err != nil, err)
// 	rolejs, err := json.Marshal(rolemap)
// 	fail.When(err != nil, err)
// 	db.AddXdLog(NewSnapshotFlow(playerinfo.User, db.PlayerId(), player.Coins, player.Ingot, player.LastLoginTime, string(itemjs), string(rolejs), int32(role_vip.Level)))
// }

func MissionLog(db *mdb.Database, missionid, action int32) {
	if !ServerCfg.EnableXDlog {
		return
	}
	playerinfo := module.Player.GetPlayer(db)
	role_vip := module.VIP.VIPInfo(db)
	ispay := ispayed(role_vip.Ingot)
	levelInfo := mission_dat.GetMissionLevelById(missionid)
	mission_level_lock := db.Lookup.PlayerMissionLevel(db.PlayerId())
	main_role_info := module.Role.GetMainRole(db)
	db.AddXdLog(NewMissionFlow(db.CId(), playerinfo.User, db.PlayerId(), playerinfo.Nick, missionid, int32(levelInfo.ParentType), levelInfo.Lock, mission_level_lock.MaxLock, ispay, int32(main_role_info.Level), int32(role_vip.Level), action))
}

func GiftCodeLog(db *mdb.Database, giftcode string, codetype int32) {
	if !ServerCfg.EnableXDlog {
		return
	}
	ispay := NOTPAY
	playerinfoGlobal := global.GetPlayerInfo(db.PlayerId())
	if playerinfoGlobal.VIPLevel > 0 {
		ispay = ISPAY
	}
	db.AddXdLog(NewGiftcodeFlow(int32(playerinfoGlobal.CId), string(playerinfoGlobal.Openid), db.PlayerId(), string(playerinfoGlobal.PlayerNick), giftcode, codetype, int32(ispay), int32(playerinfoGlobal.RoleLevel)))
}

func FightNumLog(db *mdb.Database, fightNum int64) {
	if !ServerCfg.EnableXDlog {
		return
	}
	playerinfo := module.Player.GetPlayer(db)
	role_vip := module.VIP.VIPInfo(db)
	ispay := ispayed(role_vip.Ingot)
	main_role_info := module.Role.GetMainRole(db)
	db.AddXdLog(NewFightnumFlow(db.CId(), playerinfo.User, db.PlayerId(), playerinfo.Nick, fightNum, ispay, int32(main_role_info.Level)))
}

func ispayed(ingot int64) int32 {
	if ingot > 0 {
		return ISPAY
	}
	return NOTPAY
}
