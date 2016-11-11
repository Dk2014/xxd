package tlog

import (
	"core/fail"
	"game_server/api/protocol/draw_api"
	. "game_server/config"
	"game_server/dat/battle_pet_dat"
	"game_server/dat/mission_dat"
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
	"strconv"
	"strings"
)

type ItemList struct {
	ChestItem []TChestItem
}

type TChestItem struct {
	Id       int16
	Itemtype int8
	Num      int32
}

//副本流水表
func PlayerMissionFlowLog(db *mdb.Database, missionID, action int32) {
	levelInfo := mission_dat.GetMissionLevelById(missionID)
	player_info := module.Player.GetPlayer(db)
	main_role_info := module.Role.GetMainRole(db)
	role_vip := module.VIP.VIPInfo(db)
	mission_level_lock := db.Lookup.PlayerMissionLevel(db.PlayerId())
	db.AddTLog(NewPlayerMissionFlow(player_info.User, int32(main_role_info.Level), int32(levelInfo.ParentType), missionID, levelInfo.Lock, mission_level_lock.MaxLock, int32(role_vip.Level), action))
}

//体力流水表
func PlayerPhysicalFlowLog(db *mdb.Database, count, aftercount, addorreduce, reason int32) {
	player_info := module.Player.GetPlayer(db)
	main_role_info := module.Role.GetMainRole(db)
	role_vip := module.VIP.VIPInfo(db)
	db.AddTLog(NewPlayerPhysicalFlow(player_info.User, int32(main_role_info.Level), aftercount, count, int32(role_vip.Level), reason, addorreduce))
}

//PVP流水表
func PlayerPvpFlowLog(db *mdb.Database, beforerank, afterrank, action int32) {
	player_info := module.Player.GetPlayer(db)
	main_role_info := module.Role.GetMainRole(db)
	role_vip := module.VIP.VIPInfo(db)
	db.AddTLog(NewPlayerPvpFlow(player_info.User, int32(main_role_info.Level), int32(beforerank), int32(afterrank), int32(role_vip.Level), action))
}

//新手引导流水表
func PlayerGuideFlowLog(db *mdb.Database, guidetype, action int32) {
	player_info := module.Player.GetPlayer(db)
	main_role_info := module.Role.GetMainRole(db)
	role_vip := module.VIP.VIPInfo(db)
	db.AddTLog(NewPlayerGuideFlow(player_info.User, int32(main_role_info.Level), guidetype, int32(role_vip.Level), action))
}

//道具流水表
func PlayerItemFlowLog(db *mdb.Database, itemid int16, itemtype int32, num int16, AddOrReduce, itemFlowReason int32) {
	player_info := module.Player.GetPlayer(db)
	var count, aftercount int32
	count = int32(num)
	if itemtype > 0 {
		aftercount = int32(module.Item.GetItemNum(db, itemid))
	} else if itemtype == IT_SWORDSOULFRAGMENT {
		db.Select.PlayerSwordSoul(func(row *mdb.PlayerSwordSoulRow) {
			if row.Pid() == db.PlayerId() && row.SwordSoulId() == itemid {
				aftercount++
			}
		})
	} else {
		aftercount = int32(num)
	}
	role_vip := module.VIP.VIPInfo(db)
	db.AddTLog(NewItemFlow(0, player_info.User, int32(itemtype), int32(itemid), aftercount, count, itemFlowReason, 0, AddOrReduce, int32(role_vip.Level)))
}

//系统模块参与流水表
func PlayerSystemModuleFlowLog(db *mdb.Database, moduletype int32) {
	if ServerCfg.EnableGlobalServer {
		main_role_info := global.GetPlayerInfo(db.PlayerId())
		db.AddTLog(NewPlayerSystemModuleFlow(string(main_role_info.Openid), int32(main_role_info.RoleLevel), moduletype, int32(main_role_info.VIPLevel)))
	} else {
		player_info := module.Player.GetPlayer(db)
		main_role_info := module.Role.GetMainRole(db)
		role_vip := module.VIP.VIPInfo(db)
		db.AddTLog(NewPlayerSystemModuleFlow(player_info.User, int32(main_role_info.Level), moduletype, int32(role_vip.Level)))
	}
}

//剑心抽取流水
func PlayerSwordDrawFlowLog(db *mdb.Database, moneytype, money, swordid int32) {
	player_info := module.Player.GetPlayer(db)
	main_role_info := module.Role.GetMainRole(db)
	role_vip := module.VIP.VIPInfo(db)
	db.AddTLog(NewPlayerSwordDrawFlow(player_info.User, int32(main_role_info.Level), int32(role_vip.Level), money, moneytype, swordid))
}

//神龙宝箱
func PlayerChestDrawFlowLog(db *mdb.Database, moneytype, money, chesttype int32, itemlist ItemList) {
	player_info := module.Player.GetPlayer(db)
	main_role_info := module.Role.GetMainRole(db)
	role_vip := module.VIP.VIPInfo(db)
	var itemlistarr []string
	var itemliststr string
	for _, v := range itemlist.ChestItem {
		switch draw_api.ItemType(v.Itemtype) {
		case draw_api.ITEM_TYPE_COIN:
		case draw_api.ITEM_TYPE_INGOT:
		case draw_api.ITEM_TYPE_EQUIPMENT:
			itemlistarr = append(itemlistarr, strconv.FormatInt(int64(v.Id), 10)+" "+strconv.FormatInt(int64(IT_PROPS), 10)+" "+strconv.FormatInt(int64(v.Num), 10))
		case draw_api.ITEM_TYPE_PREFERENCE:
			itemlistarr = append(itemlistarr, strconv.FormatInt(int64(v.Id), 10)+" "+strconv.FormatInt(int64(IT_PREFERENCE), 10)+" "+strconv.FormatInt(int64(v.Num), 10))
		case draw_api.ITEM_TYPE_ITEM:
			itemlistarr = append(itemlistarr, strconv.FormatInt(int64(v.Id), 10)+" "+strconv.FormatInt(int64(IT_PROPS), 10)+" "+strconv.FormatInt(int64(v.Num), 10))
		case draw_api.ITEM_TYPE_GHOST:
			itemlistarr = append(itemlistarr, strconv.FormatInt(int64(v.Id), 10)+" "+strconv.FormatInt(int64(IT_GHOST), 10)+" "+strconv.FormatInt(int64(v.Num), 10))
		case draw_api.ITEM_TYPE_SWORD_SOUL:
			itemlistarr = append(itemlistarr, strconv.FormatInt(int64(v.Id), 10)+" "+strconv.FormatInt(int64(IT_SWORD), 10)+" "+strconv.FormatInt(int64(v.Num), 10))
		case draw_api.ITEM_TYPE_PET:
			itemlistarr = append(itemlistarr, strconv.FormatInt(int64(v.Id), 10)+" "+strconv.FormatInt(int64(IT_PET), 10)+" "+strconv.FormatInt(int64(v.Num), 10))
		case draw_api.ITEM_TYPE_GHOST_FRAGMENT:
			itemlistarr = append(itemlistarr, strconv.FormatInt(int64(v.Id), 10)+" "+strconv.FormatInt(int64(IT_PROPS), 10)+" "+strconv.FormatInt(int64(v.Num), 10))
		default:
			fail.When(true, "wrong item type")
		}
	}
	if len(itemlistarr) > 0 {
		itemliststr = strings.Join(itemlistarr, " ")
	} else {
		return
	}
	db.AddTLog(NewPlayerChestDrawFlow(player_info.User, int32(main_role_info.Level), int32(role_vip.Level), money, moneytype, itemliststr, chesttype))
}

//人物战力
func PlayerFightNumFlowLog(db *mdb.Database, fightNum int32) {
	player_info := module.Player.GetPlayer(db)
	main_role_info := module.Role.GetMainRole(db)
	role_vip := module.VIP.VIPInfo(db)
	db.AddTLog(NewPlayerFightNumFlow(player_info.User, int32(main_role_info.Level), int32(role_vip.Level), fightNum))
}

//任务流水
func PlayerQuestFlowLog(db *mdb.Database, questid, questtype, action int32) {
	player_info := module.Player.GetPlayer(db)
	main_role_info := module.Role.GetMainRole(db)
	role_vip := module.VIP.VIPInfo(db)
	db.AddTLog(NewPlayerQuestFlow(player_info.User, int32(main_role_info.Level), int32(role_vip.Level), questid, questtype, action))
}

//商人消费流水
func PlayerBusinessManFlowLog(db *mdb.Database, itemid, count, cost, addorreduce int32) {
	player_info := module.Player.GetPlayer(db)
	main_role_info := module.Role.GetMainRole(db)
	role_vip := module.VIP.VIPInfo(db)
	db.AddTLog(NewBusinessManFlow(player_info.User, itemid, cost, count, int32(main_role_info.Level), addorreduce, int32(role_vip.Level)))
}

//瀛海商人消费流水
func PlayerTradeBuyFlowLog(db *mdb.Database, itemid, count, cost, moneytype int32) {
	player_info := module.Player.GetPlayer(db)
	main_role_info := module.Role.GetMainRole(db)
	role_vip := module.VIP.VIPInfo(db)
	db.AddTLog(NewTradeBuyFlow(player_info.User, itemid, count, cost, moneytype, int32(main_role_info.Level), int32(role_vip.Level)))

}

type EquipRefineList struct {
	Bluecount          int32
	Blueaftercount     int32
	Purplecount        int32
	Purpleaftercount   int32
	Goldcount          int32
	Goldaftercount     int32
	Orangecount        int32
	Orangeaftercount   int32
	Fragmentid         int32
	Fragmentcount      int32
	Fragmentaftercount int32
}

//装备精炼流水
func PlayerEquipRefineFlowLog(db *mdb.Database, itemid, imoneytype, imoney, afterimoney, refinelevel, afterrefinelevel int32, costinfo EquipRefineList) {
	player_info := module.Player.GetPlayer(db)
	main_role_info := module.Role.GetMainRole(db)
	role_vip := module.VIP.VIPInfo(db)
	db.AddTLog(NewEquipRefineFlow(player_info.User, itemid, costinfo.Bluecount, costinfo.Blueaftercount, costinfo.Purplecount, costinfo.Purpleaftercount, costinfo.Goldcount, costinfo.Goldaftercount, costinfo.Orangecount, costinfo.Orangeaftercount, costinfo.Fragmentid, costinfo.Fragmentcount, costinfo.Fragmentaftercount, imoneytype, imoney, afterimoney, refinelevel, afterrefinelevel, int32(main_role_info.Level), int32(role_vip.Level)))
}

//装备分解流水 (装备分解不给结晶了)
func PlayerEquipDecomposeFlowLog(db *mdb.Database, itemid, crystaid, crystacount, crystaaftercount, fragmentid, fragmentcount, fragmentaftercount int32) {
	player_info := module.Player.GetPlayer(db)
	main_role_info := module.Role.GetMainRole(db)
	role_vip := module.VIP.VIPInfo(db)
	db.AddTLog(NewEquipDecomposeFlow(player_info.User, itemid, crystaid, crystacount, crystaaftercount, fragmentid, fragmentcount, fragmentaftercount, int32(main_role_info.Level), int32(role_vip.Level)))
}

//魂侍强化流水
func PlayerGhostTrainFlowLog(db *mdb.Database, ghostid, fruitcount, fruitaftercount, imoneytype, imoney, ghostlevel, afterghostlevel int32) {
	player_info := module.Player.GetPlayer(db)
	main_role_info := module.Role.GetMainRole(db)
	role_vip := module.VIP.VIPInfo(db)
	db.AddTLog(NewGhostTrainFlow(player_info.User, ghostid, fruitcount, fruitaftercount, imoneytype, imoney, ghostlevel, afterghostlevel, int32(main_role_info.Level), int32(role_vip.Level)))
}

//魂侍升星流水
func PlayerGhostUpgradeFlowLog(db *mdb.Database, ghostid, fragmentid, fragmentcount, fragmentaftercount, imoneytype, imoney, ghoststart, afterghoststar int32) {
	player_info := module.Player.GetPlayer(db)
	main_role_info := module.Role.GetMainRole(db)
	role_vip := module.VIP.VIPInfo(db)
	db.AddTLog(NewGhostUpgradeFlow(player_info.User, ghostid, fragmentid, fragmentcount, fragmentaftercount, imoneytype, imoney, ghoststart, afterghoststar, int32(main_role_info.Level), int32(role_vip.Level)))
}

//剑心升星流水
func PlayerSwordUpgradeFlowLog(db *mdb.Database, swordid, swordlevel, afterswordlevel int32, swordlist string) {
	player_info := module.Player.GetPlayer(db)
	main_role_info := module.Role.GetMainRole(db)
	role_vip := module.VIP.VIPInfo(db)
	db.AddTLog(NewGhostSwordUpgradeFlow(player_info.User, swordid, swordlist, swordlevel, afterswordlevel, int32(main_role_info.Level), int32(role_vip.Level)))
}

//赠送爱心流水
func PlayerSendHeartFlowLog(db *mdb.Database, friendid int64) {
	if ServerCfg.EnableGlobalServer {
		main_role_info := global.GetPlayerInfo(db.PlayerId())
		db.AddTLog(NewSendHeartFlow(string(main_role_info.Openid), uint64(friendid), int32(main_role_info.RoleLevel), int32(main_role_info.VIPLevel)))
	} else {
		player_info := module.Player.GetPlayer(db)
		main_role_info := module.Role.GetMainRole(db)
		role_vip := module.VIP.VIPInfo(db)
		db.AddTLog(NewSendHeartFlow(player_info.User, uint64(friendid), int32(main_role_info.Level), int32(role_vip.Level)))
	}
}

//灵宠信息流水
func PlayerAddPetFlowLog(db *mdb.Database, petid int32) {
	if ServerCfg.EnableGlobalServer {
		main_role_info := global.GetPlayerInfo(db.PlayerId())
		db.AddTLog(NewAddPetFlow(string(main_role_info.Openid), petid, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, int32(main_role_info.RoleLevel), int32(main_role_info.VIPLevel)))
	} else {
		player_info := module.Player.GetPlayer(db)
		main_role_info := module.Role.GetMainRole(db)
		role_vip := module.VIP.VIPInfo(db)
		petinfo := battle_pet_dat.GetBattlePetWithEnemyId(petid)
		db.AddTLog(NewAddPetFlow(player_info.User, petid, int32(petinfo.RoundAttack), int32(petinfo.CostPower), int32(petinfo.LiveRound), int32(petinfo.LivePos), int32(petinfo.Quality), 0, 0, 0, 0, 0, 0, 0, int32(main_role_info.Level), int32(role_vip.Level)))
	}
}

//灵宠格子升级流水
func PlayerUpgradePetGridFlowLog(db *mdb.Database, gridnum int8, prematrial int16, postmatrial int16) {
	player_info := module.Player.GetPlayer(db)
	main_role_info := module.Role.GetMainRole(db)
	role_vip := module.VIP.VIPInfo(db)
	db.AddTLog(NewUpgradePetGridFlow(player_info.User, int32(gridnum), int32(prematrial), int32(postmatrial), int32(main_role_info.Level), int32(role_vip.Level)))
}

//IDIP接口日志
func IdipFlowLog(db *mdb.Database, openid, serial string, areaid, platid, partition, itemid, itemnum, itemtype, source, cmd uint32) {
	db.AddTLog(NewIDIPFLOW(areaid, platid, partition, openid, itemid, itemnum, itemtype, serial, source, cmd))
}
