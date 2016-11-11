package totem

import (
	"core/fail"
	"core/net"
	"core/time"
	"fmt"
	"game_server/api/protocol/notify_api"
	"game_server/api/protocol/totem_api"
	"game_server/dat/channel_dat"
	"game_server/dat/item_dat"
	"game_server/dat/mail_dat"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/dat/totem_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
	"game_server/xdlog"
	"math/rand"
)

const (
	//兽骨召唤概率
	BONE_CALL_BLUE_RATE   = 95
	BONE_CALL_PURPLE_RATE = 5
	BONE_CALL_GOLDEN_RATE = 0

	//圣器以及元宝召唤概率
	HALLOW_CALL_BLUE_RATE   = 40
	HALLOW_CALL_PURPLE_RATE = 50
	HALLOW_CALL_GOLDEN_RATE = 10
)

//获取阵印信息
func getTotemInfo(db *mdb.Database, out *totem_api.Info_Out) {
	totemFuncCheck(db)
	db.Select.PlayerTotem(func(row *mdb.PlayerTotemRow) {
		out.Totem = append(out.Totem, totem_api.Info_Out_Totem{
			Id:      row.Id(),
			TotemId: row.TotemId(),
			Skill:   row.SkillId(),
			Level:   row.Level(),
		})
	})
	playerTotemInfo := db.Lookup.PlayerTotemInfo(db.PlayerId())
	//刷新元宝购买次数
	if !time.IsInPointHour(player_dat.RESET_BUY_TOTEM_INGOT_CALL_TIMES_IN_HOUR, playerTotemInfo.IngotCallTimestamp) {
		playerTotemInfo.IngotCallDailyNum = 0
	}

	out.CallNum = playerTotemInfo.IngotCallDailyNum
	out.RockRuneNum = playerTotemInfo.RockRuneNum
	out.JadeRuneNum = playerTotemInfo.JadeRuneNum
	out.Pos1Id = playerTotemInfo.Pos1
	out.Pos2Id = playerTotemInfo.Pos2
	out.Pos3Id = playerTotemInfo.Pos3
	out.Pos4Id = playerTotemInfo.Pos4
	out.Pos5Id = playerTotemInfo.Pos5
}

func totemEquipPosChange(db *mdb.Database, fromPos, toPos int8) {
	totemFuncCheck(db)
	totemEquipPosCheck(db, toPos)
	playerTotemInfo := db.Lookup.PlayerTotemInfo(db.PlayerId())
	equipmentPosDataPtr := []*int64{
		&playerTotemInfo.Pos1,
		&playerTotemInfo.Pos2,
		&playerTotemInfo.Pos3,
		&playerTotemInfo.Pos4,
		&playerTotemInfo.Pos5,
	}
	*equipmentPosDataPtr[fromPos], *equipmentPosDataPtr[toPos] = *equipmentPosDataPtr[toPos], *equipmentPosDataPtr[fromPos]
	db.Update.PlayerTotemInfo(playerTotemInfo)
}

func swapTotem(db *mdb.Database, equipId, inbagId int64) {
	totemFuncCheck(db)
	playerTotemInfo := db.Lookup.PlayerTotemInfo(db.PlayerId())
	equipmentPosDataPtr := []*int64{
		&playerTotemInfo.Pos1,
		&playerTotemInfo.Pos2,
		&playerTotemInfo.Pos3,
		&playerTotemInfo.Pos4,
		&playerTotemInfo.Pos5,
	}
	var findEquipedTotem bool = false
	var targetPosPtr *int64
	inBagTotem := db.Lookup.PlayerTotem(inbagId)
	fail.When(inBagTotem == nil, "没有拥有阵印")
	equipedTotem := db.Lookup.PlayerTotem(equipId)
	fail.When(equipedTotem == nil, "没有拥有阵印")

	for _, ptr := range equipmentPosDataPtr {
		if *ptr == equipId {
			findEquipedTotem = true
			targetPosPtr = ptr
		} else if *ptr != totem_dat.TOTEM_POS_EMPTY {
			equipedTotem := db.Lookup.PlayerTotem(*ptr)
			fail.When(equipedTotem.TotemId == inBagTotem.TotemId, "不能装备相同的阵印")
		}
		//if *ptr == inbagId {
		//	fail.When(true, "阵印已装备")
		//}
	}
	fail.When(!findEquipedTotem, "阵印没有装备")
	*targetPosPtr = inbagId
	db.Update.PlayerTotemInfo(playerTotemInfo)
}

func equipTotem(db *mdb.Database, playerTotemId int64, emptyPos totem_api.EquipPos) {
	totemFuncCheck(db)
	totemEquipPosCheck(db, int8(emptyPos))
	playerTotemInfo := db.Lookup.PlayerTotemInfo(db.PlayerId())
	equipmentPosDataPtr := []*int64{
		&playerTotemInfo.Pos1,
		&playerTotemInfo.Pos2,
		&playerTotemInfo.Pos3,
		&playerTotemInfo.Pos4,
		&playerTotemInfo.Pos5,
	}
	fail.When(*equipmentPosDataPtr[int(emptyPos)] != totem_dat.TOTEM_POS_EMPTY, "目标位置已装备阵印")
	inBagTotem := db.Lookup.PlayerTotem(playerTotemId)
	fail.When(inBagTotem == nil, "没有拥有阵印")
	for _, ptr := range equipmentPosDataPtr {
		if *ptr != totem_dat.TOTEM_POS_EMPTY {
			equipedTotem := db.Lookup.PlayerTotem(*ptr)
			fail.When(equipedTotem.TotemId == inBagTotem.TotemId, "不能装备相同的阵印")
		}
	}
	*equipmentPosDataPtr[int(emptyPos)] = playerTotemId
	db.Update.PlayerTotemInfo(playerTotemInfo)
}

func unequipTotem(db *mdb.Database, pos totem_api.EquipPos) {
	totemFuncCheck(db)
	fail.When(isTotemBagFull(db), "背包已满无法召唤")
	playerTotemInfo := db.Lookup.PlayerTotemInfo(db.PlayerId())
	switch pos {
	case totem_api.EQUIP_POS_POS1:
		playerTotemInfo.Pos1 = totem_dat.TOTEM_POS_EMPTY
	case totem_api.EQUIP_POS_POS2:
		playerTotemInfo.Pos2 = totem_dat.TOTEM_POS_EMPTY
	case totem_api.EQUIP_POS_POS3:
		playerTotemInfo.Pos3 = totem_dat.TOTEM_POS_EMPTY
	case totem_api.EQUIP_POS_POS4:
		playerTotemInfo.Pos4 = totem_dat.TOTEM_POS_EMPTY
	case totem_api.EQUIP_POS_POS5:
		playerTotemInfo.Pos5 = totem_dat.TOTEM_POS_EMPTY
	default:
		fail.When(true, fmt.Sprintf("undefine totem pos %v", pos))
	}
	db.Update.PlayerTotemInfo(playerTotemInfo)
}

func decomposeTotem(db *mdb.Database, toremId int64) (rockRuneNum, jadeRuneNum int32) {
	totemFuncCheck(db)
	playerTotem := db.Lookup.PlayerTotem(toremId)
	fail.When(playerTotem == nil, "未拥有阵印")
	playerTotemInfo := db.Lookup.PlayerTotemInfo(db.PlayerId())
	switch toremId {
	case playerTotemInfo.Pos1:
		playerTotemInfo.Pos1 = totem_dat.TOTEM_POS_EMPTY
	case playerTotemInfo.Pos2:
		playerTotemInfo.Pos2 = totem_dat.TOTEM_POS_EMPTY
	case playerTotemInfo.Pos3:
		playerTotemInfo.Pos3 = totem_dat.TOTEM_POS_EMPTY
	case playerTotemInfo.Pos4:
		playerTotemInfo.Pos4 = totem_dat.TOTEM_POS_EMPTY
	case playerTotemInfo.Pos5:
		playerTotemInfo.Pos5 = totem_dat.TOTEM_POS_EMPTY
	}
	//1. 获得物品
	quality := totem_dat.GetTotemQualiyById(playerTotem.TotemId)
	totemLevelInfo := totem_dat.GetTotemLevelInfo(quality, playerTotem.Level)
	chance := totemLevelInfo.RockRuneRate + totemLevelInfo.JadeRuneRate
	randNum := rand.Intn(int(chance)) + 1
	if randNum <= int(totemLevelInfo.RockRuneRate) {
		playerTotemInfo.RockRuneNum += int32(totemLevelInfo.RockRuneNum)
	} else {
		playerTotemInfo.JadeRuneNum += int32(totemLevelInfo.JadeRuneNum)
	}

	db.Update.PlayerTotemInfo(playerTotemInfo)
	//2. 删除图腾
	db.Delete.PlayerTotem(playerTotem)
	return playerTotemInfo.RockRuneNum, playerTotemInfo.JadeRuneNum
}

func upgradeTotem(db *mdb.Database, totemId int64, out *totem_api.Upgrade_Out) (rockRuneNum, jadeRuneNum int32, ok bool) {
	totemFuncCheck(db)

	playerTotem := db.Lookup.PlayerTotem(totemId)
	fail.When(playerTotem == nil, "未拥有阵印")
	quality := totem_dat.GetTotemQualiyById(playerTotem.TotemId)
	//扣符文
	playerTotemInfo := db.Lookup.PlayerTotemInfo(db.PlayerId())
	rockRuneNum = playerTotemInfo.RockRuneNum
	jadeRuneNum = playerTotemInfo.JadeRuneNum
	totemLevel := playerTotem.Level
	if !descRune(playerTotemInfo, quality, totemLevel) {
		return rockRuneNum, jadeRuneNum, false
	}
	db.Update.PlayerTotemInfo(playerTotemInfo)

	skills := totem_dat.GetSkillByTotemId(playerTotem.TotemId)
	playerTotem.SkillId = skills[rand.Intn(len(skills))]
	if playerTotem.Level < totem_dat.TOTEM_MAX_LEVEL {
		playerTotem.Level++
	}
	db.Update.PlayerTotem(playerTotem)
	out.Id = totemId
	out.Level = playerTotem.Level
	out.Skill = playerTotem.SkillId
	return playerTotemInfo.RockRuneNum, playerTotemInfo.JadeRuneNum, true

}

//升级扣除符文，扣除规则写死，如果需改需要与客户端同步
func descRune(playerTotemInfo *mdb.PlayerTotemInfo, quality, totemLevel int8) (ok bool) {
	totemLevelInfo := totem_dat.GetTotemLevelInfo(quality, totemLevel)
	playerTotemInfo.RockRuneNum -= int32(totemLevelInfo.UpgradeNeedRock)
	playerTotemInfo.JadeRuneNum -= int32(totemLevelInfo.UpgradeNeedJade)

	return playerTotemInfo.RockRuneNum >= 0 && playerTotemInfo.JadeRuneNum >= 0
	//fail.When(playerTotemInfo.RockRuneNum < 0 || playerTotemInfo.JadeRuneNum < 0, "符文不足")

}

func callTotem(session *net.Session, callType totem_api.CallType) {
	state := module.State(session)
	db := state.Database
	totemFuncCheck(db)
	fail.When(isTotemBagFull(db), "背包已满无法召唤")
	playerTotemInfo := db.Lookup.PlayerTotemInfo(db.PlayerId())

	var blueRate, purpleRate, goldenRate int
	switch callType {
	case totem_api.CALL_TYPE_BONE:
		module.Item.DelItemByItemId(db, item_dat.ITEM_BONE, 1, tlog.IFR_TOTEM_CALL, xdlog.ET_TOTEM)
		blueRate, purpleRate, goldenRate = BONE_CALL_BLUE_RATE, BONE_CALL_PURPLE_RATE, BONE_CALL_GOLDEN_RATE
	case totem_api.CALL_TYPE_HALLOW:
		module.Item.DelItemByItemId(db, item_dat.ITEM_HALLOW, 1, tlog.IFR_TOTEM_CALL, xdlog.ET_TOTEM)
		blueRate, purpleRate, goldenRate = HALLOW_CALL_BLUE_RATE, HALLOW_CALL_PURPLE_RATE, HALLOW_CALL_GOLDEN_RATE
	case totem_api.CALL_TYPE_INGOT:
		if !time.IsInPointHour(player_dat.RESET_BUY_TOTEM_INGOT_CALL_TIMES_IN_HOUR, playerTotemInfo.IngotCallTimestamp) {
			playerTotemInfo.IngotCallDailyNum = 0
		}
		playerTotemInfo.RockRuneNum = playerTotemInfo.RockRuneNum + 5
		db.Update.PlayerTotemInfo(playerTotemInfo)
		session.Send(&notify_api.NotifyRuneChange_Out{
			RockRuneNum: playerTotemInfo.RockRuneNum,
			JadeRuneNum: playerTotemInfo.JadeRuneNum,
		})
		playerTotemInfo.IngotCallDailyNum++
		cost := totem_dat.GetTotemCallCost(playerTotemInfo.IngotCallDailyNum)
		module.Player.DecMoney(db, state.MoneyState, int64(cost), player_dat.INGOT, tlog.MFR_TOTEM_CALL, xdlog.ET_TOTEM)
		playerTotemInfo.IngotCallTimestamp = time.GetNowTime()
		blueRate, purpleRate, goldenRate = HALLOW_CALL_BLUE_RATE, HALLOW_CALL_PURPLE_RATE, HALLOW_CALL_GOLDEN_RATE
		db.Update.PlayerTotemInfo(playerTotemInfo)
	default:
		fail.When(true, "undefine totem call type")
	}
	var totems []int16
	randNum := rand.Intn(blueRate+purpleRate+goldenRate) + 1
	if randNum <= blueRate {
		totems = totem_dat.GetTotemsByQuality(totem_dat.TOTEM_QUALITY_BLUE)
	} else if randNum <= (blueRate + purpleRate) {
		totems = totem_dat.GetTotemsByQuality(totem_dat.TOTEM_QUALITY_PURPLE)
	} else {
		totems = totem_dat.GetTotemsByQuality(totem_dat.TOTEM_QUALITY_GOLDEN)
	}
	totemId := totems[rand.Intn(len(totems))]
	playerTotemId, skillId := addTotem(db, totemId)
	module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_CALL_TOTEM)
	session.Send(&notify_api.NotifyNewTotem_Out{
		Id:      playerTotemId,
		TotemId: totemId,
		Skill:   skillId,
	})
}

func addTotem(db *mdb.Database, totemId int16) (int64, int16) {
	//for sealedbook
	if session, exist := module.Player.GetPlayerOnline(db.PlayerId()); exist {
		state := module.State(session)
		if _, result := state.GetSealedBookRecord().FindRecord(item_dat.STEALDBOOK_TYPE_TOTEMS, totemId); !result {
			sealedbook := item_dat.GetSealedBookInfo(item_dat.STEALDBOOK_TYPE_TOTEMS, totemId)
			if sealedbook != nil {
				state.GetSealedBookRecord().AddRecord(item_dat.STEALDBOOK_TYPE_TOTEMS, totemId, item_dat.STEALDBOOK_HAVING, db)
			}
		}
	}

	skills := totem_dat.GetSkillByTotemId(totemId)
	skill := skills[rand.Intn(len(skills))]
	playerTotem := &mdb.PlayerTotem{
		Pid:     db.PlayerId(),
		TotemId: totemId,
		SkillId: skill,
		Level:   1,
	}
	db.Insert.PlayerTotem(playerTotem)
	totem := totem_dat.GetTotemById(totemId)
	if totem.IsRareItem() {
		player := module.Player.GetPlayer(db)
		if player != nil {
			rpc.RemoteAddWorldChannelTplMessage(db.PlayerId(), []channel_dat.MessageTpl{
				channel_dat.MessageCallTotem{
					Player: channel_dat.ParamPlayer{[]byte(player.Nick), player.Id},
					Item:   channel_dat.ParamItem{mail_dat.ATTACHMENT_TOTEM, totem.Id},
				},
			})
		}
	}
	return playerTotem.Id, skill
}
func isTotemBagFull(db *mdb.Database) bool {
	var num int
	db.Select.PlayerTotem(func(row *mdb.PlayerTotemRow) {
		num++
	})
	playerTotemInfo := db.Lookup.PlayerTotemInfo(db.PlayerId())

	if playerTotemInfo.Pos1 != totem_dat.TOTEM_POS_EMPTY {
		num--
	}
	if playerTotemInfo.Pos2 != totem_dat.TOTEM_POS_EMPTY {
		num--
	}
	if playerTotemInfo.Pos3 != totem_dat.TOTEM_POS_EMPTY {
		num--
	}
	if playerTotemInfo.Pos4 != totem_dat.TOTEM_POS_EMPTY {
		num--
	}
	if playerTotemInfo.Pos5 != totem_dat.TOTEM_POS_EMPTY {
		num--
	}
	return num >= totem_dat.TOTEM_BAG_SIZE
}

func totemFuncCheck(db *mdb.Database) {
	fail.When(!module.Player.IsOpenFunc(db, player_dat.FUNC_TOTEM), "阵印功能未开启")
}

//每个阵印开启有等级要求，
func totemEquipPosCheck(db *mdb.Database, pos int8) {
	mainRoleLv := module.Role.GetMainRole(db).Level
	var requireLv int16
	switch pos {
	//case totem_api.EQUIP_POS_POS1:
	case int8(totem_api.EQUIP_POS_POS2):
		requireLv = 55
	case int8(totem_api.EQUIP_POS_POS3):
		requireLv = 65
	case int8(totem_api.EQUIP_POS_POS4):
		requireLv = 75
	case int8(totem_api.EQUIP_POS_POS5):
		requireLv = 85
	}
	fail.When(mainRoleLv < requireLv, "等级不够")
}

func sendMailWithTotem(db *mdb.Database, totemIds []int16) {

	totemAttachments := []*mail_dat.Attachment{}

	for _, totemId := range totemIds {
		totemAttachments = append(totemAttachments, &mail_dat.Attachment{mail_dat.ATTACHMENT_TOTEM, totemId, 1})
	}

	rpc.RemoteMailSend(db.PlayerId(), mail_dat.MailTotemBagFull{
		Num:         fmt.Sprintf("%d", len(totemIds)),
		Attachments: totemAttachments,
	})
}
