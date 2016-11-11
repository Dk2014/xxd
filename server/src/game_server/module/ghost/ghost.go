package ghost

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/ghost_api"
	"game_server/battle"
	"game_server/dat/channel_dat"
	"game_server/dat/ghost_dat"
	"game_server/dat/item_dat"
	"game_server/dat/mail_dat"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/dat/role_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
	"game_server/xdlog"
	"math"
	"math/rand"
)

func Info(session *net.Session, out *ghost_api.Info_Out) {
	state := module.State(session)
	db := state.Database

	playerGhostState := getPlayerGhostState(db)
	out.TrainTimes = playerGhostState.TrainByIngotNum
	out.FlushTime = playerGhostState.LastFlushTime

	var ghostUsed bool
	// 魂侍信息
	db.Select.PlayerGhost(func(row *mdb.PlayerGhostRow) {

		playerGhost := row.GoObject()
		ghostAddData := getGhostAddData(playerGhost)
		if state.MissionLevelState != nil && state.MissionLevelState.UsedGhost != nil {
			_, ghostUsed = state.MissionLevelState.UsedGhost[row.GhostId()]
		} else {
			ghostUsed = false
		}

		out.Ghosts = append(out.Ghosts, ghost_api.Info_Out_Ghosts{
			Id:         row.Id(),
			GhostId:    row.GhostId(),
			Level:      row.Level(),
			SkillLevel: row.SkillLevel(),
			Star:       row.Star(),
			Exp:        row.Exp(),
			Pos:        row.Pos(),
			Health:     ghostAddData.Health,
			Attack:     ghostAddData.Attack,
			Defence:    ghostAddData.Defence,
			AddGrowth:  ghostAddData.Growth,
			RelationId: row.RelationId(),
			Used:       ghostUsed,
		})
	})

	var usedGhostSkill bool
	// 装备信息
	db.Select.PlayerGhostEquipment(func(row *mdb.PlayerGhostEquipmentRow) {
		usedGhostSkill = false
		if state.MissionLevelState != nil {
			if fighterAttr, ok := state.MissionLevelState.AttackerInfo.Fighters[int(row.RoleId())]; ok {
				usedGhostSkill = fighterAttr.UsedGhostSkill
			}
		}
		out.RoleEquip = append(out.RoleEquip, ghost_api.Info_Out_RoleEquip{
			RoleId:          row.RoleId(),
			GhostPower:      row.GhostPower(),
			Pos1Id:          row.Pos1(),
			Pos2Id:          row.Pos2(),
			Pos3Id:          row.Pos3(),
			Pos4Id:          row.Pos4(),
			AlreadyUseGhost: usedGhostSkill,
		})
	})
}

// 获取魂侍加成数据
func getGhostAddData(playerGhost *mdb.PlayerGhost) (ghostAddData *ghost_dat.GhostAddData) {

	ghost := ghost_dat.GetGhost(playerGhost.GhostId)
	ghostStarDat := ghost_dat.GetGhostStar(ghost.Quality, playerGhost.Star)

	//基础数据
	ghostAddData = &ghost_dat.GhostAddData{
		Health:  ghost.Health + ghostStarDat.Health,
		Attack:  ghost.Attack + ghostStarDat.Attack,
		Defence: ghost.Defence + ghostStarDat.Defence,
		Growth:  playerGhost.AddGrowth,
	}

	levelAddData := int32((ghostStarDat.Growth + ghostAddData.Growth) * playerGhost.Level)
	ghostAddData.Health += levelAddData * 2
	ghostAddData.Attack += int32(math.Ceil(float64(levelAddData) * 0.6))
	ghostAddData.Defence += int32(math.Ceil(float64(levelAddData) * 0.25))

	return ghostAddData
}

// 装备魂侍和魂侍替换位置
func Swap(session *net.Session, in *ghost_api.Swap_In) {
	state := module.State(session)
	db := state.Database
	roleId := in.RoleId
	idInBag := in.IdBag
	idInEquip := in.IdEquip

	swap(db, roleId, idInBag, idInEquip)
}

// 装备魂侍和魂侍替换位置
func swap(db *mdb.Database, roleId int8, idInBag, idInEquip int64) {
	ghostBag := db.Lookup.PlayerGhost(idInBag)
	ghostEquip := db.Lookup.PlayerGhost(idInEquip)

	playerGhostEquipment := getPlayerGhostEquipment(db, roleId)

	ghostEquipmentPosList := []*int64{
		&playerGhostEquipment.Pos1,
		&playerGhostEquipment.Pos2,
		&playerGhostEquipment.Pos3,
		&playerGhostEquipment.Pos4,
	}
	for _, posPtr := range ghostEquipmentPosList {
		fail.When(idInBag == *posPtr, "魂侍已装备")
	}

	// 设置装备位id
	switch idInEquip {
	case playerGhostEquipment.Pos1:
		playerGhostEquipment.Pos1 = idInBag
	case playerGhostEquipment.Pos2:
		playerGhostEquipment.Pos2 = idInBag
	case playerGhostEquipment.Pos3:
		playerGhostEquipment.Pos3 = idInBag
	case playerGhostEquipment.Pos4:
		playerGhostEquipment.Pos4 = idInBag
	}

	ghostEquip.Pos = ghost_dat.UNEQUIPPED
	ghostEquip.RoleId = ghost_dat.NO_ROLE
	ghostBag.Pos = ghost_dat.EQUIPPED
	ghostBag.RoleId = roleId

	// 更新数据
	db.Update.PlayerGhost(ghostBag)
	db.Update.PlayerGhost(ghostEquip)
	db.Update.PlayerGhostEquipment(playerGhostEquipment)
}

// 获取角色已经装备的魂侍列表
func getPlayerGhostEquipment(db *mdb.Database, roleId int8) (playerGhostEquipment *mdb.PlayerGhostEquipment) {

	db.Select.PlayerGhostEquipment(func(row *mdb.PlayerGhostEquipmentRow) {
		if row.RoleId() == roleId {
			playerGhostEquipment = row.GoObject()
			row.Break()
		}
	})

	fail.When(playerGhostEquipment == nil, "no this playerGhostEquipment")
	return playerGhostEquipment
}

func Equip(session *net.Session, in *ghost_api.Equip_In) {
	state := module.State(session)
	playerGhostId := in.FromId
	roleId := in.RoleId
	equipPos := in.ToPos

	equip(state, playerGhostId, roleId, equipPos)
}

//装备魂侍
func equip(state *module.SessionState, playerGhostId int64, roleId int8, equipPos ghost_api.EquipPos) {

	db := state.Database

	playerGhost := db.Lookup.PlayerGhost(playerGhostId)
	playerGhostEquipment := getPlayerGhostEquipment(db, roleId)
	role := module.Role.GetBuddyRoleInTeam(state.Database, roleId)

	fail.When(role.Level < GetGhostGridRequireLevel(equipPos), "等级未足够开启魂侍装备格子")

	fail.When(playerGhost.Pos == ghost_dat.EQUIPPED, "魂侍已装备")
	playerGhost.Pos = ghost_dat.EQUIPPED
	playerGhost.RoleId = roleId

	ghostEquipmentPosList := []*int64{
		&playerGhostEquipment.Pos1,
		&playerGhostEquipment.Pos2,
		&playerGhostEquipment.Pos3,
		&playerGhostEquipment.Pos4,
	}

	fail.When(*ghostEquipmentPosList[int(equipPos)] > 0, "该位置已装备魂侍")
	*ghostEquipmentPosList[int(equipPos)] = playerGhostId

	db.Update.PlayerGhost(playerGhost)
	db.Update.PlayerGhostEquipment(playerGhostEquipment)
	tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_GHOST)
}

//卸载一个魂侍
func Unequip(session *net.Session, in *ghost_api.Unequip_In) {
	state := module.State(session)

	roleId := in.RoleId
	playerGhostId := in.FromId

	unequip(state, roleId, playerGhostId)
}

// 卸载一个魂侍
func unequip(state *module.SessionState, roleId int8, playerGhostId int64) {

	playerGhost := state.Database.Lookup.PlayerGhost(playerGhostId)

	fail.When(playerGhost.Pos == ghost_dat.UNEQUIPPED, "不能卸载未装备的魂侍")
	playerGhost.Pos = ghost_dat.UNEQUIPPED
	playerGhost.RoleId = ghost_dat.NO_ROLE

	playerGhostEquipment := getPlayerGhostEquipment(state.Database, roleId)

	ghostEquipmentPosList := []*int64{
		&playerGhostEquipment.Pos1,
		&playerGhostEquipment.Pos2,
		&playerGhostEquipment.Pos3,
		&playerGhostEquipment.Pos4,
	}

	var foundTargetGhost bool

	//卸载的魂侍必须侍最后一个有效魂侍
	var index int
	for index = len(ghostEquipmentPosList) - 1; index >= 0; index-- {
		ghostId := *ghostEquipmentPosList[index]
		if ghostId == playerGhostId {
			foundTargetGhost = true
			*ghostEquipmentPosList[index] = 0
			break
		}
	}

	fail.When(!foundTargetGhost, "找不到需要卸载的魂侍")

	// 魂侍关卡限制；主角必须要有一个魂侍装配好
	if state.MissionLevelState != nil && state.MissionLevelState.LevelType == battle.BT_GHOST_LEVEL && role_dat.IsMainRole(roleId) {
		fail.When(playerGhostEquipment.Pos1 <= 0, "魂侍关卡主角必须装备至少一个魂侍")
	}

	state.Database.Update.PlayerGhostEquipment(playerGhostEquipment)
	state.Database.Update.PlayerGhost(playerGhost)

	// 清理ghost的连锁关系
	clearGhostRelation(state.Database, playerGhost)
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_GHOST)
}

func EquipPosChange(session *net.Session, in *ghost_api.EquipPosChange_In) {
	state := module.State(session)
	db := state.Database

	roleId := in.RoleId
	playerGhostId := in.FromId
	toPos := in.ToPos

	equipPosChange(db, roleId, playerGhostId, toPos)
}

// 已装备魂侍交换位置或者把已装备魂侍放到空的格子
func equipPosChange(db *mdb.Database, roleId int8, id int64, toPos ghost_api.EquipPos) {
	role := module.Role.GetBuddyRoleInTeam(db, roleId)
	fail.When(role.Level < GetGhostGridRequireLevel(toPos), "等级未足够开启魂侍装备格子")

	playerGhostEquipment := getPlayerGhostEquipment(db, roleId)

	var fromPos ghost_api.EquipPos
	switch id {
	case playerGhostEquipment.Pos1:
		fromPos = ghost_api.EQUIP_POS_POS1
	case playerGhostEquipment.Pos2:
		fromPos = ghost_api.EQUIP_POS_POS2
	case playerGhostEquipment.Pos3:
		fromPos = ghost_api.EQUIP_POS_POS3
	case playerGhostEquipment.Pos4:
		fromPos = ghost_api.EQUIP_POS_POS4
	}

	playerGhostIds := []*int64{
		&playerGhostEquipment.Pos1,
		&playerGhostEquipment.Pos2,
		&playerGhostEquipment.Pos3,
		&playerGhostEquipment.Pos4,
	}

	srcGhost := db.Lookup.PlayerGhost(*playerGhostIds[fromPos])
	if *playerGhostIds[toPos] > 0 {
		targetGhost := db.Lookup.PlayerGhost(*playerGhostIds[toPos])
		fail.When(srcGhost.RoleId != targetGhost.RoleId, "装备在不同角色身上的魂侍不能交换")
	}

	*playerGhostIds[fromPos], *playerGhostIds[toPos] = *playerGhostIds[toPos], *playerGhostIds[fromPos]

	db.Update.PlayerGhostEquipment(playerGhostEquipment)
	tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_GHOST)
}

func Train(session *net.Session, in *ghost_api.Train_In, out *ghost_api.Train_Out) {
	state := module.State(session)

	playerGhostId := in.Id

	out.AddExp = train(state, playerGhostId)
}

func train(state *module.SessionState, playerGhostId int64) (exp int64) {
	db := state.Database

	playerGhost := db.Lookup.PlayerGhost(playerGhostId)
	ghostLevel := ghost_dat.GetGhostLevel(playerGhost.Level)

	roleLevel := module.Role.GetMainRole(db).Level

	// 混屎等级不能高于主角等级，不能高于最大等级
	if playerGhost.Level >= roleLevel || playerGhost.Level >= ghost_dat.MAX_GHOST_LEVEL {
		return
	}

	itemNum := module.Item.GetItemNum(db, item_dat.ITEM_YINGJIEGUOSHI)
	var afternum, ghostlevel int32
	afternum = int32(itemNum) - int32(ghostLevel.NeedFruitNum)
	module.Item.DelItemByItemId(db, item_dat.ITEM_YINGJIEGUOSHI, int16(ghostLevel.NeedFruitNum), tlog.IFR_GHOST_TRAIN, xdlog.ET_GHOST_TRAIN)

	ghostlevel = int32(playerGhost.Level)
	exp = int64(rand.Int63n(ghostLevel.MaxAddExp-ghostLevel.MinAddExp+1) + ghostLevel.MinAddExp)
	addExp(db, playerGhost, exp)

	tlog.PlayerGhostTrainFlowLog(db, int32(playerGhost.GhostId), int32(itemNum), afternum, tlog.MT_INGOT, 0 /*废弃*/, ghostlevel, int32(playerGhost.Level))

	module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_TRAIN_GHOST)
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_GHOST)

	return exp
}

//设置魂侍等级
func setlevel(db *mdb.Database, playerGhostId int64, ghostLevel int16) {
	var ghostexists = false
	roleinfo := module.Role.GetMainRole(db)
	db.Select.PlayerGhost(func(row *mdb.PlayerGhostRow) {
		if row.GhostId() == int16(playerGhostId) {
			ghostexists = true
			playerGhost := row.GoObject()
			if ghostLevel >= roleinfo.Level {
				ghostLevel = roleinfo.Level
			}
			fail.When(ghostLevel <= playerGhost.Level, "set level must great than player ghost current level")
			if ghostLevel > ghost_dat.MAX_GHOST_LEVEL {
				playerGhost.Level = ghost_dat.MAX_GHOST_LEVEL
			} else {
				playerGhost.Level = ghostLevel
			}
			playerGhost.Exp = 0
			db.Update.PlayerGhost(playerGhost)
			row.Break()
		}
	})
	fail.When(!ghostexists, "playerGhostId not exists")
}

// 为魂侍加经验
func addExp(db *mdb.Database, playerGhost *mdb.PlayerGhost, exp int64) {

	roleLevel := module.Role.GetMainRole(db).Level

	ghostLevel := playerGhost.Level
	ghostExp := playerGhost.Exp + exp

	for {
		//魂侍的等级不能大于主角等级
		if ghostLevel == roleLevel {
			ghostExp = 0
			break
		}

		needExp := ghost_dat.GetGhostLevel(ghostLevel).Exp

		if ghostExp < needExp {
			break
		}

		ghostLevel = ghostLevel + 1
		ghostExp = ghostExp - needExp
	}

	playerGhost.Level = ghostLevel
	playerGhost.Exp = ghostExp
	db.Update.PlayerGhost(playerGhost)
}

func Upgrade(session *net.Session, in *ghost_api.Upgrade_In) {
	state := module.State(session)
	playerGhostId := in.Id

	upgrade(state, playerGhostId)
}

func Baptize(session *net.Session, in *ghost_api.Baptize_In, out *ghost_api.Baptize_Out) {
	state := module.State(session)
	playerGhostId := in.Id
	out.AddGrowth = baptize(state, playerGhostId)
}

func baptize(state *module.SessionState, playerGhostId int64) (addGrowth int8) {
	db := state.Database

	mainRoleLevel := int32(module.Role.GetMainRole(db).Level)
	fail.When(mainRoleLevel < ghost_dat.BAPTIZE_PLAYER_MIN_LEVEL, "Cant baptize cause player level not enough")

	playerGhost := db.Lookup.PlayerGhost(playerGhostId)
	ghostBaptize := ghost_dat.GetGhostBaptize(playerGhost.Star)

	fail.When(int32(playerGhost.AddGrowth) >= int32(ghostBaptize.MaxAddGrowth), "Cant baptize cause current add growth is max")

	probability1 := int32(ghostBaptize.Probablity1)
	minAddGrowth1 := int32(ghostBaptize.MinAddGrowth1)
	maxAddGrowth1 := int32(ghostBaptize.MaxAddGrowth1)

	probability2 := int32(ghostBaptize.Probablity2)
	minAddGrowth2 := int32(ghostBaptize.MinAddGrowth2)
	maxAddGrowth2 := int32(ghostBaptize.MaxAddGrowth2)

	probAll := probability1 + probability2
	prob := rand.Int31n(probAll)

	if prob <= probability1 {
		addGrowth = int8(rand.Int31n(maxAddGrowth1-minAddGrowth1+1) + minAddGrowth1)
	} else if prob <= probability1+probability2 {
		addGrowth = int8(rand.Int31n(maxAddGrowth2-minAddGrowth2+1) + minAddGrowth2)
	}

	if int8(playerGhost.AddGrowth)+addGrowth > ghostBaptize.MaxAddGrowth {
		addGrowth = ghostBaptize.MaxAddGrowth - int8(playerGhost.AddGrowth)
	}

	module.Item.DelItemByItemId(db, item_dat.ITEM_GHOST_CRYSTAL_ID, ghost_dat.BAPTIZE_FRAME_NUM, tlog.IFR_GHOST_BAPTIZE, xdlog.ET_GHOST_BAPTIZE)

	playerGhost.AddGrowth = int16(addGrowth)
	db.Update.PlayerGhost(playerGhost)
	return addGrowth
}

func upgrade(state *module.SessionState, playerGhostId int64) {
	db := state.Database

	playerGhost := db.Lookup.PlayerGhost(playerGhostId)
	ghost := ghost_dat.GetGhost(playerGhost.GhostId)

	fail.When(playerGhost.Star == 5, "cant Upgrade")

	dstStar := playerGhost.Star + 1
	ghostStar := ghost_dat.GetGhostStar(ghost.Quality, dstStar)

	itemNum := module.Item.GetItemNum(db, ghost.FragmentId)
	var afterfragmentcount int32
	module.Item.DelItemByItemId(db, ghost.FragmentId, ghostStar.NeedFragmentNum, tlog.IFR_GHOST_UPGRADE, xdlog.ET_GHOST_UPGRADE)
	afterfragmentcount = int32(itemNum) - int32(ghostStar.NeedFragmentNum)

	module.Player.DecMoney(db, state.MoneyState, ghostStar.Costs, player_dat.COINS, tlog.MFR_GHOST_UPGRADE, xdlog.ET_GHOST_UPGRADE)
	tlog.PlayerGhostUpgradeFlowLog(db, int32(playerGhost.GhostId), int32(ghost.FragmentId), int32(itemNum), afterfragmentcount, tlog.MT_INGOT, 0, int32(playerGhost.Star), int32(dstStar))

	playerGhost.Star = dstStar

	db.Update.PlayerGhost(playerGhost)
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_GHOST)
}

// 合成
func Compose(session *net.Session, in *ghost_api.Compose_In, out *ghost_api.Compose_Out) {
	state := module.State(session)

	db := state.Database
	ghostId := in.GhostId

	fail.When(isHasGhost(db, ghostId), "already have this ghost")

	ghost := ghost_dat.GetGhost(ghostId)

	ghostStar := ghost_dat.GetGhostStar(ghost.Quality, ghost.InitStar)

	// 消费
	module.Item.DelItemByItemId(db, ghost.FragmentId, ghostStar.NeedFragmentNum, tlog.IFR_GHOST_COMPOSE, xdlog.ET_GHOST_COMPOSE)
	module.Player.DecMoney(db, state.MoneyState, ghostStar.Costs, player_dat.COINS, tlog.MFR_GHOST_COMPOSE, xdlog.ET_GHOST_COMPOSE)

	if ghost.Quality == ghost_dat.COLOR_GOLD {
		rpc.RemoteAddWorldChannelTplMessage(state.PlayerId, []channel_dat.MessageTpl{
			channel_dat.MessageComposeGhost{
				Player: channel_dat.ParamPlayer{state.PlayerNick, state.PlayerId},
				Item:   channel_dat.ParamItem{mail_dat.ATTACHMENT_GHOST, ghost.Id},
			},
		})
	}
	// 增加魂侍
	out.Id = addGhost(state.Database, ghostId, tlog.IFR_GHOST_COMPOSE, xdlog.ET_GHOST_COMPOSE)
	out.GhostId = ghostId
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_GHOST)
}

func getPlayerGhostState(db *mdb.Database) *mdb.PlayerGhostState {
	playerGhostState := db.Lookup.PlayerGhostState(db.PlayerId())
	if !time.IsInPointHour(player_dat.RESET_GHOST_TRAIN_INFO_IN_HOUR, playerGhostState.TrainByIngotTime) {
		playerGhostState.TrainByIngotNum = 0
		db.Update.PlayerGhostState(playerGhostState)
	}
	return playerGhostState
}

func trainGhostSkill(state *module.SessionState, playerGhostId int64) {
	playerGhost := state.Database.Lookup.PlayerGhost(playerGhostId)
	fail.When(playerGhost == nil, "找不到目标魂侍")

	// 魂侍技能已满
	if playerGhost.SkillLevel >= playerGhost.Level {
		return
	}

	trainPrice := ghost_dat.GetGhostSkillTrainPrice(playerGhost.SkillLevel)
	if !module.Player.CheckMoney(state, trainPrice, player_dat.COINS) {
		return
	}

	module.Player.DecMoney(state.Database, state.MoneyState, trainPrice, player_dat.COINS, tlog.MFR_GHOST_SKILL_LEVELUP, xdlog.ET_GHOST_SKILL)
	playerGhost.SkillLevel += 1
	state.Database.Update.PlayerGhost(playerGhost)
}

func flushGhostAttr(state *module.SessionState, playerGhostId int64) int64 {
	playerGhostState := getPlayerGhostState(state.Database)

	nowTime := time.GetNowTime()
	fail.When(nowTime-playerGhostState.LastFlushTime < ghost_dat.FLUSH_ATTR_CD, "skill flushing is in CD")

	playerGhost := state.Database.Lookup.PlayerGhost(playerGhostId)
	fail.When(playerGhost == nil, "找不到目标魂侍")

	has_flushed_sth := false

	//清空魂侍技能等级
	price := float64(ghost_dat.GetGhostSkillTrainTotalPrice(playerGhost.SkillLevel - 1))
	price *= ghost_dat.FLUSH_ATTR_BACK_PERCENT
	totalPrice := int64(price)

	if totalPrice > 0 {
		module.Player.IncMoney(state.Database, state.MoneyState, totalPrice, player_dat.COINS, 0 /*TODO*/, xdlog.ET_FLUSH_GHOST_ATTR, "")
		has_flushed_sth = true
	}
	playerGhost.SkillLevel = 1

	//清空魂侍等级
	totalFruit := float64(ghost_dat.GetGhostFruitCost(playerGhost.Level-1, playerGhost.Exp))
	totalFruit *= ghost_dat.FLUSH_ATTR_BACK_PERCENT
	totalFruit = math.Ceil(totalFruit)

	if totalFruit > 0 {
		module.Item.AddItem(state.Database, item_dat.ITEM_YINGJIEGUOSHI, int16(totalFruit), 0 /*TODO*/, xdlog.ET_FLUSH_GHOST_ATTR, "")
		has_flushed_sth = true
	}
	playerGhost.Level = 1
	playerGhost.Exp = 0

	state.Database.Update.PlayerGhost(playerGhost)

	if has_flushed_sth {
		playerGhostState.LastFlushTime = nowTime
		state.Database.Update.PlayerGhostState(playerGhostState)
	}

	return playerGhostState.LastFlushTime
}

func relation(state *module.SessionState, master, slave int64) {
	masterGhost := state.Database.Lookup.PlayerGhost(master)
	slaveGhost := state.Database.Lookup.PlayerGhost(slave)
	fail.When(masterGhost == nil || slaveGhost == nil, "can not find the ghosts ")
	fail.When(masterGhost.RoleId != slaveGhost.RoleId, "this two ghosts must be equiped by the same role")
	fail.When(masterGhost.Pos == 0 || slaveGhost.Pos == 0, "all ghosts must be equiped")

	relationship := ghost_dat.GetGhostRelationship(masterGhost.GhostId, slaveGhost.GhostId)
	fail.When(relationship == nil, "does not exist relationship between this two ghosts")
	fail.When(!ghost_dat.CheckGhostsRelation(relationship, masterGhost.GhostId, masterGhost.Star, slaveGhost.GhostId, slaveGhost.Star), "the star level is not satisfied between this two ghosts")

	// 之前存在关系的,需要先清除
	clearGhostRelation(state.Database, masterGhost)
	clearGhostRelation(state.Database, slaveGhost)

	masterGhost.RelationId = slave
	slaveGhost.RelationId = master
	state.Database.Update.PlayerGhost(masterGhost)
	state.Database.Update.PlayerGhost(slaveGhost)
}
