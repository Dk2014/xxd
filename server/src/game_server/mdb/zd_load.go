package mdb

/*
#include "zd_cgo.h"
*/
import "C"
import "strconv"
import "core/log"
import "core/mysql"

func itoa(i int) string {
	return strconv.Itoa(i)
}

func playerIdToStr(playerId int64) string {
	return strconv.FormatInt(playerId, 10)
}

func initRowId(db *mysql.Connection, id *int64, sql string) bool {
	qr, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	row := qr.Rows[0]
	if row[0] == nil {
		return false
	}
	*id = row.Int64(0)
	return true
}

func newPlayerTables(db *mysql.Connection) {
	sql := "SELECT `id` FROM `player`"

	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")

	for _, row := range res.Rows {
		playerId := C.int64_t(row.Int64(iId))

		playerDb := C.NewPlayerTables()
		playerDb.Pid = C.int64_t(playerId)

		SetPlayerTables(playerDb)
	}
}

func initGlobalTables(db *mysql.Connection) {
	loadGlobalAnnouncement(db)
	loadGlobalArenaRank(db)
	loadGlobalClique(db)
	loadGlobalCliqueBoat(db)
	loadGlobalDespairBoss(db)
	loadGlobalDespairBossHistory(db)
	loadGlobalDespairLand(db)
	loadGlobalDespairLandBattleRecord(db)
	loadGlobalDespairLandCamp(db)
	loadGlobalGiftCardRecord(db)
	loadGlobalGroupBuyStatus(db)
	loadGlobalMail(db)
	loadGlobalMailAttachments(db)
	loadGlobalTbXxdOnlinecnt(db)
}

func initPlayerGlobalTables(db *mysql.Connection, sid, wgn, procId int) {
	loadPlayerGlobalChatState(db, 0, sid, wgn, procId)
	loadPlayerGlobalCliqueBuilding(db, 0, sid, wgn, procId)
	loadPlayerGlobalCliqueBuildingQuest(db, 0, sid, wgn, procId)
	loadPlayerGlobalCliqueDailyQuest(db, 0, sid, wgn, procId)
	loadPlayerGlobalCliqueEscort(db, 0, sid, wgn, procId)
	loadPlayerGlobalCliqueEscortMessage(db, 0, sid, wgn, procId)
	loadPlayerGlobalCliqueHijack(db, 0, sid, wgn, procId)
	loadPlayerGlobalCliqueInfo(db, 0, sid, wgn, procId)
	loadPlayerGlobalCliqueKongfu(db, 0, sid, wgn, procId)
	loadPlayerGlobalFriend(db, 0, sid, wgn, procId)
	loadPlayerGlobalFriendChat(db, 0, sid, wgn, procId)
	loadPlayerGlobalFriendState(db, 0, sid, wgn, procId)
	loadPlayerGlobalGiftCardRecord(db, 0, sid, wgn, procId)
	loadPlayerGlobalMailState(db, 0, sid, wgn, procId)
	loadPlayerGlobalWorldChatState(db, 0, sid, wgn, procId)
}

func initPlayerTables(db *mysql.Connection, sid, wgn, procId int) {
	loadPlayer(db, 0, sid, wgn, procId)
	loadPlayerActivity(db, 0, sid, wgn, procId)
	loadPlayerAdditionQuest(db, 0, sid, wgn, procId)
	loadPlayerAnySdkOrder(db, 0, sid, wgn, procId)
	loadPlayerArena(db, 0, sid, wgn, procId)
	loadPlayerArenaRank(db, 0, sid, wgn, procId)
	loadPlayerArenaRecord(db, 0, sid, wgn, procId)
	loadPlayerAwakenGraphic(db, 0, sid, wgn, procId)
	loadPlayerBattlePet(db, 0, sid, wgn, procId)
	loadPlayerBattlePetGrid(db, 0, sid, wgn, procId)
	loadPlayerBattlePetState(db, 0, sid, wgn, procId)
	loadPlayerChestState(db, 0, sid, wgn, procId)
	loadPlayerCliqueKongfuAttrib(db, 0, sid, wgn, procId)
	loadPlayerCoins(db, 0, sid, wgn, procId)
	loadPlayerCornucopia(db, 0, sid, wgn, procId)
	loadPlayerDailyQuest(db, 0, sid, wgn, procId)
	loadPlayerDailyQuestStarAwardInfo(db, 0, sid, wgn, procId)
	loadPlayerDailySignInRecord(db, 0, sid, wgn, procId)
	loadPlayerDailySignInState(db, 0, sid, wgn, procId)
	loadPlayerDespairLandCampState(db, 0, sid, wgn, procId)
	loadPlayerDespairLandCampStateHistory(db, 0, sid, wgn, procId)
	loadPlayerDespairLandLevelRecord(db, 0, sid, wgn, procId)
	loadPlayerDespairLandState(db, 0, sid, wgn, procId)
	loadPlayerDrivingSwordEvent(db, 0, sid, wgn, procId)
	loadPlayerDrivingSwordEventExploring(db, 0, sid, wgn, procId)
	loadPlayerDrivingSwordEventTreasure(db, 0, sid, wgn, procId)
	loadPlayerDrivingSwordEventVisiting(db, 0, sid, wgn, procId)
	loadPlayerDrivingSwordEventmask(db, 0, sid, wgn, procId)
	loadPlayerDrivingSwordInfo(db, 0, sid, wgn, procId)
	loadPlayerDrivingSwordMap(db, 0, sid, wgn, procId)
	loadPlayerEquipment(db, 0, sid, wgn, procId)
	loadPlayerEventAwardRecord(db, 0, sid, wgn, procId)
	loadPlayerEventDailyOnline(db, 0, sid, wgn, procId)
	loadPlayerEventVnDailyRecharge(db, 0, sid, wgn, procId)
	loadPlayerEventVnDragonBallRecord(db, 0, sid, wgn, procId)
	loadPlayerEventsFightPower(db, 0, sid, wgn, procId)
	loadPlayerEventsIngotRecord(db, 0, sid, wgn, procId)
	loadPlayerExtendLevel(db, 0, sid, wgn, procId)
	loadPlayerExtendQuest(db, 0, sid, wgn, procId)
	loadPlayerFame(db, 0, sid, wgn, procId)
	loadPlayerFashion(db, 0, sid, wgn, procId)
	loadPlayerFashionState(db, 0, sid, wgn, procId)
	loadPlayerFateBoxState(db, 0, sid, wgn, procId)
	loadPlayerFightNum(db, 0, sid, wgn, procId)
	loadPlayerFirstCharge(db, 0, sid, wgn, procId)
	loadPlayerFormation(db, 0, sid, wgn, procId)
	loadPlayerFuncKey(db, 0, sid, wgn, procId)
	loadPlayerGhost(db, 0, sid, wgn, procId)
	loadPlayerGhostEquipment(db, 0, sid, wgn, procId)
	loadPlayerGhostState(db, 0, sid, wgn, procId)
	loadPlayerHardLevel(db, 0, sid, wgn, procId)
	loadPlayerHardLevelRecord(db, 0, sid, wgn, procId)
	loadPlayerHeart(db, 0, sid, wgn, procId)
	loadPlayerHeartDraw(db, 0, sid, wgn, procId)
	loadPlayerHeartDrawCardRecord(db, 0, sid, wgn, procId)
	loadPlayerHeartDrawWheelRecord(db, 0, sid, wgn, procId)
	loadPlayerInfo(db, 0, sid, wgn, procId)
	loadPlayerIsOperated(db, 0, sid, wgn, procId)
	loadPlayerItem(db, 0, sid, wgn, procId)
	loadPlayerItemAppendix(db, 0, sid, wgn, procId)
	loadPlayerItemBuyback(db, 0, sid, wgn, procId)
	loadPlayerItemRecastState(db, 0, sid, wgn, procId)
	loadPlayerLevelAwardInfo(db, 0, sid, wgn, procId)
	loadPlayerLoginAwardRecord(db, 0, sid, wgn, procId)
	loadPlayerMail(db, 0, sid, wgn, procId)
	loadPlayerMailAttachment(db, 0, sid, wgn, procId)
	loadPlayerMailAttachmentLg(db, 0, sid, wgn, procId)
	loadPlayerMailLg(db, 0, sid, wgn, procId)
	loadPlayerMeditationState(db, 0, sid, wgn, procId)
	loadPlayerMission(db, 0, sid, wgn, procId)
	loadPlayerMissionLevel(db, 0, sid, wgn, procId)
	loadPlayerMissionLevelRecord(db, 0, sid, wgn, procId)
	loadPlayerMissionLevelStateBin(db, 0, sid, wgn, procId)
	loadPlayerMissionRecord(db, 0, sid, wgn, procId)
	loadPlayerMissionStarAward(db, 0, sid, wgn, procId)
	loadPlayerMoneyTree(db, 0, sid, wgn, procId)
	loadPlayerMonthCardAwardRecord(db, 0, sid, wgn, procId)
	loadPlayerMonthCardInfo(db, 0, sid, wgn, procId)
	loadPlayerMultiLevelInfo(db, 0, sid, wgn, procId)
	loadPlayerNewYearConsumeRecord(db, 0, sid, wgn, procId)
	loadPlayerNpcTalkRecord(db, 0, sid, wgn, procId)
	loadPlayerOpenedTownTreasure(db, 0, sid, wgn, procId)
	loadPlayerPhysical(db, 0, sid, wgn, procId)
	loadPlayerPurchaseRecord(db, 0, sid, wgn, procId)
	loadPlayerPushNotifySwitch(db, 0, sid, wgn, procId)
	loadPlayerPveState(db, 0, sid, wgn, procId)
	loadPlayerQqVipGift(db, 0, sid, wgn, procId)
	loadPlayerQuest(db, 0, sid, wgn, procId)
	loadPlayerRainbowLevel(db, 0, sid, wgn, procId)
	loadPlayerRainbowLevelStateBin(db, 0, sid, wgn, procId)
	loadPlayerRole(db, 0, sid, wgn, procId)
	loadPlayerSealedbook(db, 0, sid, wgn, procId)
	loadPlayerSendHeartRecord(db, 0, sid, wgn, procId)
	loadPlayerSkill(db, 0, sid, wgn, procId)
	loadPlayerState(db, 0, sid, wgn, procId)
	loadPlayerSwordSoul(db, 0, sid, wgn, procId)
	loadPlayerSwordSoulEquipment(db, 0, sid, wgn, procId)
	loadPlayerSwordSoulState(db, 0, sid, wgn, procId)
	loadPlayerTaoyuanBless(db, 0, sid, wgn, procId)
	loadPlayerTaoyuanFileds(db, 0, sid, wgn, procId)
	loadPlayerTaoyuanHeart(db, 0, sid, wgn, procId)
	loadPlayerTaoyuanItem(db, 0, sid, wgn, procId)
	loadPlayerTaoyuanMessage(db, 0, sid, wgn, procId)
	loadPlayerTaoyuanProduct(db, 0, sid, wgn, procId)
	loadPlayerTaoyuanPurchaseRecord(db, 0, sid, wgn, procId)
	loadPlayerTaoyuanQuest(db, 0, sid, wgn, procId)
	loadPlayerTbXxdRoleinfo(db, 0, sid, wgn, procId)
	loadPlayerTeamInfo(db, 0, sid, wgn, procId)
	loadPlayerTotem(db, 0, sid, wgn, procId)
	loadPlayerTotemInfo(db, 0, sid, wgn, procId)
	loadPlayerTown(db, 0, sid, wgn, procId)
	loadPlayerTraderRefreshState(db, 0, sid, wgn, procId)
	loadPlayerTraderStoreState(db, 0, sid, wgn, procId)
	loadPlayerUseSkill(db, 0, sid, wgn, procId)
	loadPlayerVip(db, 0, sid, wgn, procId)
}

func loadGlobalAnnouncement(db *mysql.Connection) {
	log.Infof("load global_announcement")
	sql := "SELECT * FROM `global_announcement`"
	if !initRowId(db, &(g_RowIds.GlobalAnnouncement), "SELECT MAX(`id`) FROM `global_announcement`") {
		return
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iExpireTime := res.Map("expire_time")
	iTplId := res.Map("tpl_id")
	iParameters := res.Map("parameters")
	iContent := res.Map("content")
	iSendTime := res.Map("send_time")
	iSpacingTime := res.Map("spacing_time")
	for _, row := range res.Rows {
		crow := C.New_GlobalAnnouncement()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.ExpireTime = C.int64_t(row.Int64(iExpireTime))
		crow.TplId = C.int32_t(row.Int32(iTplId))
		row_Parameters := row.Str(iParameters)
		crow.Parameters = C.CString(string(row_Parameters))
		crow.Parameters_len = C.int(len(row_Parameters))
		row_Content := row.Str(iContent)
		crow.Content = C.CString(string(row_Content))
		crow.Content_len = C.int(len(row_Content))
		crow.SendTime = C.int64_t(row.Int64(iSendTime))
		crow.SpacingTime = C.int64_t(row.Int64(iSpacingTime))
		crow.next = C.g_GlobalTables.GlobalAnnouncement
		C.g_GlobalTables.GlobalAnnouncement = crow
		if g_Hooks.GlobalAnnouncement != nil {
			g_Hooks.GlobalAnnouncement.Load(&GlobalAnnouncementRow{row: crow})
		}
	}
}

func loadGlobalArenaRank(db *mysql.Connection) {
	log.Infof("load global_arena_rank")
	sql := "SELECT * FROM `global_arena_rank`"
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iRank := res.Map("rank")
	iPid := res.Map("pid")
	for _, row := range res.Rows {
		crow := C.New_GlobalArenaRank()
		crow.Rank = C.int32_t(row.Int32(iRank))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.next = C.g_GlobalTables.GlobalArenaRank
		C.g_GlobalTables.GlobalArenaRank = crow
		if g_Hooks.GlobalArenaRank != nil {
			g_Hooks.GlobalArenaRank.Load(&GlobalArenaRankRow{row: crow})
		}
	}
}

func loadGlobalClique(db *mysql.Connection) {
	log.Infof("load global_clique")
	sql := "SELECT * FROM `global_clique`"
	if !initRowId(db, &(g_RowIds.GlobalClique), "SELECT MAX(`id`) FROM `global_clique`") {
		return
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iName := res.Map("name")
	iAnnounce := res.Map("announce")
	iTotalDonateCoins := res.Map("total_donate_coins")
	iOwnerPid := res.Map("owner_pid")
	iOwnerLoginTime := res.Map("owner_login_time")
	iMangerPid1 := res.Map("manger_pid1")
	iMangerPid2 := res.Map("manger_pid2")
	iCenterBuildingCoins := res.Map("center_building_coins")
	iTempleBuildingCoins := res.Map("temple_building_coins")
	iHealthBuildingCoins := res.Map("health_building_coins")
	iAttackBuildingCoins := res.Map("attack_building_coins")
	iDefenseBuildingCoins := res.Map("defense_building_coins")
	iStoreBuildingCoins := res.Map("store_building_coins")
	iBankBuildingCoins := res.Map("bank_building_coins")
	iCenterBuildingLevel := res.Map("center_building_level")
	iTempleBuildingLevel := res.Map("temple_building_level")
	iHealthBuildingLevel := res.Map("health_building_level")
	iAttackBuildingLevel := res.Map("attack_building_level")
	iDefenseBuildingLevel := res.Map("defense_building_level")
	iBankBuildingLevel := res.Map("bank_building_level")
	iMemberJson := res.Map("member_json")
	iAutoAudit := res.Map("auto_audit")
	iAutoAuditLevel := res.Map("auto_audit_level")
	iContrib := res.Map("contrib")
	iContribClearTime := res.Map("contrib_clear_time")
	iRecruitTime := res.Map("recruit_time")
	iWorshipTime := res.Map("worship_time")
	iWorshipCnt := res.Map("worship_cnt")
	iStoreSendTime := res.Map("store_send_time")
	iStoreSendCnt := res.Map("store_send_cnt")
	for _, row := range res.Rows {
		crow := C.New_GlobalClique()
		crow.Id = C.int64_t(row.Int64(iId))
		row_Name := row.Str(iName)
		crow.Name = C.CString(string(row_Name))
		crow.Name_len = C.int(len(row_Name))
		row_Announce := row.Str(iAnnounce)
		crow.Announce = C.CString(string(row_Announce))
		crow.Announce_len = C.int(len(row_Announce))
		crow.TotalDonateCoins = C.int64_t(row.Int64(iTotalDonateCoins))
		crow.OwnerPid = C.int64_t(row.Int64(iOwnerPid))
		crow.OwnerLoginTime = C.int64_t(row.Int64(iOwnerLoginTime))
		crow.MangerPid1 = C.int64_t(row.Int64(iMangerPid1))
		crow.MangerPid2 = C.int64_t(row.Int64(iMangerPid2))
		crow.CenterBuildingCoins = C.int64_t(row.Int64(iCenterBuildingCoins))
		crow.TempleBuildingCoins = C.int64_t(row.Int64(iTempleBuildingCoins))
		crow.HealthBuildingCoins = C.int64_t(row.Int64(iHealthBuildingCoins))
		crow.AttackBuildingCoins = C.int64_t(row.Int64(iAttackBuildingCoins))
		crow.DefenseBuildingCoins = C.int64_t(row.Int64(iDefenseBuildingCoins))
		crow.StoreBuildingCoins = C.int64_t(row.Int64(iStoreBuildingCoins))
		crow.BankBuildingCoins = C.int64_t(row.Int64(iBankBuildingCoins))
		crow.CenterBuildingLevel = C.int16_t(row.Int16(iCenterBuildingLevel))
		crow.TempleBuildingLevel = C.int16_t(row.Int16(iTempleBuildingLevel))
		crow.HealthBuildingLevel = C.int16_t(row.Int16(iHealthBuildingLevel))
		crow.AttackBuildingLevel = C.int16_t(row.Int16(iAttackBuildingLevel))
		crow.DefenseBuildingLevel = C.int16_t(row.Int16(iDefenseBuildingLevel))
		crow.BankBuildingLevel = C.int16_t(row.Int16(iBankBuildingLevel))
		row_MemberJson := row.Bin(iMemberJson)
		crow.MemberJson = C.CString(string(row_MemberJson))
		crow.MemberJson_len = C.int(len(row_MemberJson))
		crow.AutoAudit = C.int8_t(row.Int8(iAutoAudit))
		crow.AutoAuditLevel = C.int16_t(row.Int16(iAutoAuditLevel))
		crow.Contrib = C.int64_t(row.Int64(iContrib))
		crow.ContribClearTime = C.int64_t(row.Int64(iContribClearTime))
		crow.RecruitTime = C.int64_t(row.Int64(iRecruitTime))
		crow.WorshipTime = C.int64_t(row.Int64(iWorshipTime))
		crow.WorshipCnt = C.int8_t(row.Int8(iWorshipCnt))
		crow.StoreSendTime = C.int64_t(row.Int64(iStoreSendTime))
		crow.StoreSendCnt = C.int8_t(row.Int8(iStoreSendCnt))
		crow.next = C.g_GlobalTables.GlobalClique
		C.g_GlobalTables.GlobalClique = crow
		if g_Hooks.GlobalClique != nil {
			g_Hooks.GlobalClique.Load(&GlobalCliqueRow{row: crow})
		}
	}
}

func loadGlobalCliqueBoat(db *mysql.Connection) {
	log.Infof("load global_clique_boat")
	sql := "SELECT * FROM `global_clique_boat`"
	if !initRowId(db, &(g_RowIds.GlobalCliqueBoat), "SELECT MAX(`id`) FROM `global_clique_boat`") {
		return
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iCliqueId := res.Map("clique_id")
	iBoatType := res.Map("boat_type")
	iOwnerPid := res.Map("owner_pid")
	iEscortStartTimestamp := res.Map("escort_start_timestamp")
	iTotalEscortTime := res.Map("total_escort_time")
	iRecoverPid := res.Map("recover_pid")
	iHijackerPid := res.Map("hijacker_pid")
	iHijackTimestamp := res.Map("hijack_timestamp")
	iStatus := res.Map("status")
	for _, row := range res.Rows {
		crow := C.New_GlobalCliqueBoat()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.CliqueId = C.int64_t(row.Int64(iCliqueId))
		crow.BoatType = C.int16_t(row.Int16(iBoatType))
		crow.OwnerPid = C.int64_t(row.Int64(iOwnerPid))
		crow.EscortStartTimestamp = C.int64_t(row.Int64(iEscortStartTimestamp))
		crow.TotalEscortTime = C.int16_t(row.Int16(iTotalEscortTime))
		crow.RecoverPid = C.int64_t(row.Int64(iRecoverPid))
		crow.HijackerPid = C.int64_t(row.Int64(iHijackerPid))
		crow.HijackTimestamp = C.int64_t(row.Int64(iHijackTimestamp))
		crow.Status = C.int8_t(row.Int8(iStatus))
		crow.next = C.g_GlobalTables.GlobalCliqueBoat
		C.g_GlobalTables.GlobalCliqueBoat = crow
		if g_Hooks.GlobalCliqueBoat != nil {
			g_Hooks.GlobalCliqueBoat.Load(&GlobalCliqueBoatRow{row: crow})
		}
	}
}

func loadGlobalDespairBoss(db *mysql.Connection) {
	log.Infof("load global_despair_boss")
	sql := "SELECT * FROM `global_despair_boss`"
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iCampType := res.Map("camp_type")
	iLevel := res.Map("level")
	iTimestamp := res.Map("timestamp")
	iDeadTimestamp := res.Map("dead_timestamp")
	iHp := res.Map("hp")
	iMaxHp := res.Map("max_hp")
	for _, row := range res.Rows {
		crow := C.New_GlobalDespairBoss()
		crow.CampType = C.int8_t(row.Int8(iCampType))
		crow.Level = C.int16_t(row.Int16(iLevel))
		crow.Timestamp = C.int64_t(row.Int64(iTimestamp))
		crow.DeadTimestamp = C.int64_t(row.Int64(iDeadTimestamp))
		crow.Hp = C.int64_t(row.Int64(iHp))
		crow.MaxHp = C.int64_t(row.Int64(iMaxHp))
		crow.next = C.g_GlobalTables.GlobalDespairBoss
		C.g_GlobalTables.GlobalDespairBoss = crow
		if g_Hooks.GlobalDespairBoss != nil {
			g_Hooks.GlobalDespairBoss.Load(&GlobalDespairBossRow{row: crow})
		}
	}
}

func loadGlobalDespairBossHistory(db *mysql.Connection) {
	log.Infof("load global_despair_boss_history")
	sql := "SELECT * FROM `global_despair_boss_history`"
	if !initRowId(db, &(g_RowIds.GlobalDespairBossHistory), "SELECT MAX(`id`) FROM `global_despair_boss_history`") {
		return
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iVersion := res.Map("version")
	iCampType := res.Map("camp_type")
	iLevel := res.Map("level")
	iTimestamp := res.Map("timestamp")
	iStartTimestamp := res.Map("start_timestamp")
	iDeadTimestamp := res.Map("dead_timestamp")
	for _, row := range res.Rows {
		crow := C.New_GlobalDespairBossHistory()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Version = C.int64_t(row.Int64(iVersion))
		crow.CampType = C.int8_t(row.Int8(iCampType))
		crow.Level = C.int16_t(row.Int16(iLevel))
		crow.Timestamp = C.int64_t(row.Int64(iTimestamp))
		crow.StartTimestamp = C.int64_t(row.Int64(iStartTimestamp))
		crow.DeadTimestamp = C.int64_t(row.Int64(iDeadTimestamp))
		crow.next = C.g_GlobalTables.GlobalDespairBossHistory
		C.g_GlobalTables.GlobalDespairBossHistory = crow
		if g_Hooks.GlobalDespairBossHistory != nil {
			g_Hooks.GlobalDespairBossHistory.Load(&GlobalDespairBossHistoryRow{row: crow})
		}
	}
}

func loadGlobalDespairLand(db *mysql.Connection) {
	log.Infof("load global_despair_land")
	sql := "SELECT * FROM `global_despair_land`"
	if !initRowId(db, &(g_RowIds.GlobalDespairLand), "SELECT MAX(`id`) FROM `global_despair_land`") {
		return
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iVersion := res.Map("version")
	iTimestamp := res.Map("timestamp")
	for _, row := range res.Rows {
		crow := C.New_GlobalDespairLand()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Version = C.int64_t(row.Int64(iVersion))
		crow.Timestamp = C.int64_t(row.Int64(iTimestamp))
		crow.next = C.g_GlobalTables.GlobalDespairLand
		C.g_GlobalTables.GlobalDespairLand = crow
		if g_Hooks.GlobalDespairLand != nil {
			g_Hooks.GlobalDespairLand.Load(&GlobalDespairLandRow{row: crow})
		}
	}
}

func loadGlobalDespairLandBattleRecord(db *mysql.Connection) {
	log.Infof("load global_despair_land_battle_record")
	sql := "SELECT * FROM `global_despair_land_battle_record`"
	if !initRowId(db, &(g_RowIds.GlobalDespairLandBattleRecord), "SELECT MAX(`id`) FROM `global_despair_land_battle_record`") {
		return
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iFightnum := res.Map("fightnum")
	iTimestamp := res.Map("timestamp")
	iTag := res.Map("tag")
	iBattleVersion := res.Map("battle_version")
	iCampType := res.Map("camp_type")
	iLevelId := res.Map("level_id")
	iRecord := res.Map("record")
	for _, row := range res.Rows {
		crow := C.New_GlobalDespairLandBattleRecord()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Fightnum = C.int32_t(row.Int32(iFightnum))
		crow.Timestamp = C.int64_t(row.Int64(iTimestamp))
		crow.Tag = C.int16_t(row.Int16(iTag))
		crow.BattleVersion = C.int16_t(row.Int16(iBattleVersion))
		crow.CampType = C.int8_t(row.Int8(iCampType))
		crow.LevelId = C.int32_t(row.Int32(iLevelId))
		row_Record := row.Bin(iRecord)
		crow.Record = C.CString(string(row_Record))
		crow.Record_len = C.int(len(row_Record))
		crow.next = C.g_GlobalTables.GlobalDespairLandBattleRecord
		C.g_GlobalTables.GlobalDespairLandBattleRecord = crow
		if g_Hooks.GlobalDespairLandBattleRecord != nil {
			g_Hooks.GlobalDespairLandBattleRecord.Load(&GlobalDespairLandBattleRecordRow{row: crow})
		}
	}
}

func loadGlobalDespairLandCamp(db *mysql.Connection) {
	log.Infof("load global_despair_land_camp")
	sql := "SELECT * FROM `global_despair_land_camp`"
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iCampType := res.Map("camp_type")
	iTimestamp := res.Map("timestamp")
	iDeadTimestamp := res.Map("dead_timestamp")
	iLevel := res.Map("level")
	for _, row := range res.Rows {
		crow := C.New_GlobalDespairLandCamp()
		crow.CampType = C.int8_t(row.Int8(iCampType))
		crow.Timestamp = C.int64_t(row.Int64(iTimestamp))
		crow.DeadTimestamp = C.int64_t(row.Int64(iDeadTimestamp))
		crow.Level = C.int16_t(row.Int16(iLevel))
		crow.next = C.g_GlobalTables.GlobalDespairLandCamp
		C.g_GlobalTables.GlobalDespairLandCamp = crow
		if g_Hooks.GlobalDespairLandCamp != nil {
			g_Hooks.GlobalDespairLandCamp.Load(&GlobalDespairLandCampRow{row: crow})
		}
	}
}

func loadGlobalGiftCardRecord(db *mysql.Connection) {
	log.Infof("load global_gift_card_record")
	sql := "SELECT * FROM `global_gift_card_record`"
	if !initRowId(db, &(g_RowIds.GlobalGiftCardRecord), "SELECT MAX(`id`) FROM `global_gift_card_record`") {
		return
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iCode := res.Map("code")
	iTimestamp := res.Map("timestamp")
	for _, row := range res.Rows {
		crow := C.New_GlobalGiftCardRecord()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		row_Code := row.Str(iCode)
		crow.Code = C.CString(string(row_Code))
		crow.Code_len = C.int(len(row_Code))
		crow.Timestamp = C.int64_t(row.Int64(iTimestamp))
		crow.next = C.g_GlobalTables.GlobalGiftCardRecord
		C.g_GlobalTables.GlobalGiftCardRecord = crow
		if g_Hooks.GlobalGiftCardRecord != nil {
			g_Hooks.GlobalGiftCardRecord.Load(&GlobalGiftCardRecordRow{row: crow})
		}
	}
}

func loadGlobalGroupBuyStatus(db *mysql.Connection) {
	log.Infof("load global_group_buy_status")
	sql := "SELECT * FROM `global_group_buy_status`"
	if !initRowId(db, &(g_RowIds.GlobalGroupBuyStatus), "SELECT MAX(`id`) FROM `global_group_buy_status`") {
		return
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iCid := res.Map("cid")
	iStatus := res.Map("status")
	for _, row := range res.Rows {
		crow := C.New_GlobalGroupBuyStatus()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Cid = C.int32_t(row.Int32(iCid))
		crow.Status = C.int32_t(row.Int32(iStatus))
		crow.next = C.g_GlobalTables.GlobalGroupBuyStatus
		C.g_GlobalTables.GlobalGroupBuyStatus = crow
		if g_Hooks.GlobalGroupBuyStatus != nil {
			g_Hooks.GlobalGroupBuyStatus.Load(&GlobalGroupBuyStatusRow{row: crow})
		}
	}
}

func loadGlobalMail(db *mysql.Connection) {
	log.Infof("load global_mail")
	sql := "SELECT * FROM `global_mail`"
	if !initRowId(db, &(g_RowIds.GlobalMail), "SELECT MAX(`id`) FROM `global_mail`") {
		return
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iSendTime := res.Map("send_time")
	iExpireTime := res.Map("expire_time")
	iTitle := res.Map("title")
	iContent := res.Map("content")
	iPriority := res.Map("priority")
	iMinLevel := res.Map("min_level")
	iMaxLevel := res.Map("max_level")
	iMinVipLevel := res.Map("min_vip_level")
	iMaxVipLevel := res.Map("max_vip_level")
	iMinCreateTime := res.Map("min_create_time")
	iMaxCreateTime := res.Map("max_create_time")
	for _, row := range res.Rows {
		crow := C.New_GlobalMail()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.SendTime = C.int64_t(row.Int64(iSendTime))
		crow.ExpireTime = C.int64_t(row.Int64(iExpireTime))
		row_Title := row.Str(iTitle)
		crow.Title = C.CString(string(row_Title))
		crow.Title_len = C.int(len(row_Title))
		row_Content := row.Str(iContent)
		crow.Content = C.CString(string(row_Content))
		crow.Content_len = C.int(len(row_Content))
		crow.Priority = C.int8_t(row.Int8(iPriority))
		crow.MinLevel = C.int16_t(row.Int16(iMinLevel))
		crow.MaxLevel = C.int16_t(row.Int16(iMaxLevel))
		crow.MinVipLevel = C.int16_t(row.Int16(iMinVipLevel))
		crow.MaxVipLevel = C.int16_t(row.Int16(iMaxVipLevel))
		crow.MinCreateTime = C.int64_t(row.Int64(iMinCreateTime))
		crow.MaxCreateTime = C.int64_t(row.Int64(iMaxCreateTime))
		crow.next = C.g_GlobalTables.GlobalMail
		C.g_GlobalTables.GlobalMail = crow
		if g_Hooks.GlobalMail != nil {
			g_Hooks.GlobalMail.Load(&GlobalMailRow{row: crow})
		}
	}
}

func loadGlobalMailAttachments(db *mysql.Connection) {
	log.Infof("load global_mail_attachments")
	sql := "SELECT * FROM `global_mail_attachments`"
	if !initRowId(db, &(g_RowIds.GlobalMailAttachments), "SELECT MAX(`id`) FROM `global_mail_attachments`") {
		return
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iGlobalMailId := res.Map("global_mail_id")
	iItemId := res.Map("item_id")
	iAttachmentType := res.Map("attachment_type")
	iItemNum := res.Map("item_num")
	for _, row := range res.Rows {
		crow := C.New_GlobalMailAttachments()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.GlobalMailId = C.int64_t(row.Int64(iGlobalMailId))
		crow.ItemId = C.int16_t(row.Int16(iItemId))
		crow.AttachmentType = C.int8_t(row.Int8(iAttachmentType))
		crow.ItemNum = C.int64_t(row.Int64(iItemNum))
		crow.next = C.g_GlobalTables.GlobalMailAttachments
		C.g_GlobalTables.GlobalMailAttachments = crow
		if g_Hooks.GlobalMailAttachments != nil {
			g_Hooks.GlobalMailAttachments.Load(&GlobalMailAttachmentsRow{row: crow})
		}
	}
}

func loadGlobalTbXxdOnlinecnt(db *mysql.Connection) {
	log.Infof("load global_tb_xxd_onlinecnt")
	sql := "SELECT * FROM `global_tb_xxd_onlinecnt`"
	if !initRowId(db, &(g_RowIds.GlobalTbXxdOnlinecnt), "SELECT MAX(`id`) FROM `global_tb_xxd_onlinecnt`") {
		return
	}
	func(s string){}(sql)
}

func loadPlayer(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player")
	}
	sql := "SELECT * FROM `player`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.Player), "SELECT MAX(`id`) FROM `player` WHERE (id >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (id >> 32) = " + itoa(sid) + " AND  id % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE id = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iUser := res.Map("user")
	iNick := res.Map("nick")
	iMainRoleId := res.Map("main_role_id")
	iCid := res.Map("cid")
	for _, row := range res.Rows {
		crow := C.New_Player()
		crow.Id = C.int64_t(row.Int64(iId))
		row_User := row.Str(iUser)
		crow.User = C.CString(string(row_User))
		crow.User_len = C.int(len(row_User))
		row_Nick := row.Str(iNick)
		crow.Nick = C.CString(string(row_Nick))
		crow.Nick_len = C.int(len(row_Nick))
		crow.MainRoleId = C.int64_t(row.Int64(iMainRoleId))
		crow.Cid = C.int64_t(row.Int64(iCid))
		playerDb := GetPlayerTables(int64(crow.Id))
		playerDb.Player = crow
		if g_Hooks.Player != nil {
			g_Hooks.Player.Load(&PlayerRow{row: crow})
		}
	}
}

func loadPlayerActivity(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_activity")
	}
	sql := "SELECT * FROM `player_activity`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iActivity := res.Map("activity")
	iLastUpdate := res.Map("last_update")
	for _, row := range res.Rows {
		crow := C.New_PlayerActivity()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Activity = C.int32_t(row.Int32(iActivity))
		crow.LastUpdate = C.int64_t(row.Int64(iLastUpdate))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerActivity = crow
		if g_Hooks.PlayerActivity != nil {
			g_Hooks.PlayerActivity.Load(&PlayerActivityRow{row: crow})
		}
	}
}

func loadPlayerAdditionQuest(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_addition_quest")
	}
	sql := "SELECT * FROM `player_addition_quest`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerAdditionQuest), "SELECT MAX(`id`) FROM `player_addition_quest` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iSerialNumber := res.Map("serial_number")
	iQuestId := res.Map("quest_id")
	iLock := res.Map("lock")
	iProgress := res.Map("progress")
	iState := res.Map("state")
	for _, row := range res.Rows {
		crow := C.New_PlayerAdditionQuest()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.SerialNumber = C.int32_t(row.Int32(iSerialNumber))
		crow.QuestId = C.int32_t(row.Int32(iQuestId))
		crow.Lock = C.int16_t(row.Int16(iLock))
		crow.Progress = C.int16_t(row.Int16(iProgress))
		crow.State = C.int8_t(row.Int8(iState))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerAdditionQuest
		playerDb.PlayerAdditionQuest = crow
		if g_Hooks.PlayerAdditionQuest != nil {
			g_Hooks.PlayerAdditionQuest.Load(&PlayerAdditionQuestRow{row: crow})
		}
	}
}

func loadPlayerAnySdkOrder(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_any_sdk_order")
	}
	sql := "SELECT * FROM `player_any_sdk_order`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerAnySdkOrder), "SELECT MAX(`id`) FROM `player_any_sdk_order` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iOrderId := res.Map("order_id")
	iPresent := res.Map("present")
	for _, row := range res.Rows {
		crow := C.New_PlayerAnySdkOrder()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.OrderId = C.int64_t(row.Int64(iOrderId))
		crow.Present = C.int64_t(row.Int64(iPresent))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerAnySdkOrder
		playerDb.PlayerAnySdkOrder = crow
		if g_Hooks.PlayerAnySdkOrder != nil {
			g_Hooks.PlayerAnySdkOrder.Load(&PlayerAnySdkOrderRow{row: crow})
		}
	}
}

func loadPlayerArena(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_arena")
	}
	sql := "SELECT * FROM `player_arena`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iDailyNum := res.Map("daily_num")
	iFailedCdTime := res.Map("failed_cd_time")
	iBattleTime := res.Map("battle_time")
	iWinTimes := res.Map("win_times")
	iDailyAwardItem := res.Map("daily_award_item")
	iNewRecordCount := res.Map("new_record_count")
	iDailyFame := res.Map("daily_fame")
	iBuyTimes := res.Map("buy_times")
	for _, row := range res.Rows {
		crow := C.New_PlayerArena()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.DailyNum = C.int16_t(row.Int16(iDailyNum))
		crow.FailedCdTime = C.int64_t(row.Int64(iFailedCdTime))
		crow.BattleTime = C.int64_t(row.Int64(iBattleTime))
		crow.WinTimes = C.int16_t(row.Int16(iWinTimes))
		crow.DailyAwardItem = C.int32_t(row.Int32(iDailyAwardItem))
		crow.NewRecordCount = C.int16_t(row.Int16(iNewRecordCount))
		crow.DailyFame = C.int32_t(row.Int32(iDailyFame))
		crow.BuyTimes = C.int16_t(row.Int16(iBuyTimes))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerArena = crow
		if g_Hooks.PlayerArena != nil {
			g_Hooks.PlayerArena.Load(&PlayerArenaRow{row: crow})
		}
	}
}

func loadPlayerArenaRank(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_arena_rank")
	}
	sql := "SELECT * FROM `player_arena_rank`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iRank := res.Map("rank")
	iRank1 := res.Map("rank1")
	iRank2 := res.Map("rank2")
	iRank3 := res.Map("rank3")
	iTime := res.Map("time")
	for _, row := range res.Rows {
		crow := C.New_PlayerArenaRank()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Rank = C.int32_t(row.Int32(iRank))
		crow.Rank1 = C.int32_t(row.Int32(iRank1))
		crow.Rank2 = C.int32_t(row.Int32(iRank2))
		crow.Rank3 = C.int32_t(row.Int32(iRank3))
		crow.Time = C.int64_t(row.Int64(iTime))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerArenaRank = crow
		if g_Hooks.PlayerArenaRank != nil {
			g_Hooks.PlayerArenaRank.Load(&PlayerArenaRankRow{row: crow})
		}
	}
}

func loadPlayerArenaRecord(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_arena_record")
	}
	sql := "SELECT * FROM `player_arena_record`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerArenaRecord), "SELECT MAX(`id`) FROM `player_arena_record` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iMode := res.Map("mode")
	iOldRank := res.Map("old_rank")
	iNewRank := res.Map("new_rank")
	iFightNum := res.Map("fight_num")
	iTargetPid := res.Map("target_pid")
	iTargetNick := res.Map("target_nick")
	iTargetOldRank := res.Map("target_old_rank")
	iTargetNewRank := res.Map("target_new_rank")
	iTargetFightNum := res.Map("target_fight_num")
	iRecordTime := res.Map("record_time")
	for _, row := range res.Rows {
		crow := C.New_PlayerArenaRecord()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Mode = C.int8_t(row.Int8(iMode))
		crow.OldRank = C.int32_t(row.Int32(iOldRank))
		crow.NewRank = C.int32_t(row.Int32(iNewRank))
		crow.FightNum = C.int32_t(row.Int32(iFightNum))
		crow.TargetPid = C.int64_t(row.Int64(iTargetPid))
		row_TargetNick := row.Str(iTargetNick)
		crow.TargetNick = C.CString(string(row_TargetNick))
		crow.TargetNick_len = C.int(len(row_TargetNick))
		crow.TargetOldRank = C.int32_t(row.Int32(iTargetOldRank))
		crow.TargetNewRank = C.int32_t(row.Int32(iTargetNewRank))
		crow.TargetFightNum = C.int32_t(row.Int32(iTargetFightNum))
		crow.RecordTime = C.int64_t(row.Int64(iRecordTime))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerArenaRecord
		playerDb.PlayerArenaRecord = crow
		if g_Hooks.PlayerArenaRecord != nil {
			g_Hooks.PlayerArenaRecord.Load(&PlayerArenaRecordRow{row: crow})
		}
	}
}

func loadPlayerAwakenGraphic(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_awaken_graphic")
	}
	sql := "SELECT * FROM `player_awaken_graphic`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerAwakenGraphic), "SELECT MAX(`id`) FROM `player_awaken_graphic` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iRoleId := res.Map("role_id")
	iAttrImpl := res.Map("attr_impl")
	iLevel := res.Map("level")
	for _, row := range res.Rows {
		crow := C.New_PlayerAwakenGraphic()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.RoleId = C.int8_t(row.Int8(iRoleId))
		crow.AttrImpl = C.int16_t(row.Int16(iAttrImpl))
		crow.Level = C.int8_t(row.Int8(iLevel))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerAwakenGraphic
		playerDb.PlayerAwakenGraphic = crow
		if g_Hooks.PlayerAwakenGraphic != nil {
			g_Hooks.PlayerAwakenGraphic.Load(&PlayerAwakenGraphicRow{row: crow})
		}
	}
}

func loadPlayerBattlePet(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_battle_pet")
	}
	sql := "SELECT * FROM `player_battle_pet`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerBattlePet), "SELECT MAX(`id`) FROM `player_battle_pet` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iBattlePetId := res.Map("battle_pet_id")
	iLevel := res.Map("level")
	iExp := res.Map("exp")
	iSkillLevel := res.Map("skill_level")
	for _, row := range res.Rows {
		crow := C.New_PlayerBattlePet()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.BattlePetId = C.int32_t(row.Int32(iBattlePetId))
		crow.Level = C.int32_t(row.Int32(iLevel))
		crow.Exp = C.int64_t(row.Int64(iExp))
		crow.SkillLevel = C.int32_t(row.Int32(iSkillLevel))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerBattlePet
		playerDb.PlayerBattlePet = crow
		if g_Hooks.PlayerBattlePet != nil {
			g_Hooks.PlayerBattlePet.Load(&PlayerBattlePetRow{row: crow})
		}
	}
}

func loadPlayerBattlePetGrid(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_battle_pet_grid")
	}
	sql := "SELECT * FROM `player_battle_pet_grid`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerBattlePetGrid), "SELECT MAX(`id`) FROM `player_battle_pet_grid` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iGridId := res.Map("grid_id")
	iBattlePetId := res.Map("battle_pet_id")
	for _, row := range res.Rows {
		crow := C.New_PlayerBattlePetGrid()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.GridId = C.int8_t(row.Int8(iGridId))
		crow.BattlePetId = C.int32_t(row.Int32(iBattlePetId))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerBattlePetGrid
		playerDb.PlayerBattlePetGrid = crow
		if g_Hooks.PlayerBattlePetGrid != nil {
			g_Hooks.PlayerBattlePetGrid.Load(&PlayerBattlePetGridRow{row: crow})
		}
	}
}

func loadPlayerBattlePetState(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_battle_pet_state")
	}
	sql := "SELECT * FROM `player_battle_pet_state`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iUpgradeByIngotNum := res.Map("upgrade_by_ingot_num")
	iUpgradeByIngotTime := res.Map("upgrade_by_ingot_time")
	for _, row := range res.Rows {
		crow := C.New_PlayerBattlePetState()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.UpgradeByIngotNum = C.int32_t(row.Int32(iUpgradeByIngotNum))
		crow.UpgradeByIngotTime = C.int64_t(row.Int64(iUpgradeByIngotTime))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerBattlePetState = crow
		if g_Hooks.PlayerBattlePetState != nil {
			g_Hooks.PlayerBattlePetState.Load(&PlayerBattlePetStateRow{row: crow})
		}
	}
}

func loadPlayerChestState(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_chest_state")
	}
	sql := "SELECT * FROM `player_chest_state`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iFreeCoinChestNum := res.Map("free_coin_chest_num")
	iLastFreeCoinChestAt := res.Map("last_free_coin_chest_at")
	iCoinChestNum := res.Map("coin_chest_num")
	iCoinChestTenNum := res.Map("coin_chest_ten_num")
	iIsFirstCoinTen := res.Map("is_first_coin_ten")
	iCoinChestConsume := res.Map("coin_chest_consume")
	iLastCoinChestAt := res.Map("last_coin_chest_at")
	iLastFreeIngotChestAt := res.Map("last_free_ingot_chest_at")
	iIngotChestNum := res.Map("ingot_chest_num")
	iIngotChestTenNum := res.Map("ingot_chest_ten_num")
	iIsFirstIngotTen := res.Map("is_first_ingot_ten")
	iIngotChestConsume := res.Map("ingot_chest_consume")
	iLastIngotChestAt := res.Map("last_ingot_chest_at")
	iLastFreePetChestAt := res.Map("last_free_pet_chest_at")
	for _, row := range res.Rows {
		crow := C.New_PlayerChestState()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.FreeCoinChestNum = C.int32_t(row.Int32(iFreeCoinChestNum))
		crow.LastFreeCoinChestAt = C.int64_t(row.Int64(iLastFreeCoinChestAt))
		crow.CoinChestNum = C.int32_t(row.Int32(iCoinChestNum))
		crow.CoinChestTenNum = C.int32_t(row.Int32(iCoinChestTenNum))
		crow.IsFirstCoinTen = C.int8_t(row.Int8(iIsFirstCoinTen))
		crow.CoinChestConsume = C.int64_t(row.Int64(iCoinChestConsume))
		crow.LastCoinChestAt = C.int64_t(row.Int64(iLastCoinChestAt))
		crow.LastFreeIngotChestAt = C.int64_t(row.Int64(iLastFreeIngotChestAt))
		crow.IngotChestNum = C.int32_t(row.Int32(iIngotChestNum))
		crow.IngotChestTenNum = C.int32_t(row.Int32(iIngotChestTenNum))
		crow.IsFirstIngotTen = C.int8_t(row.Int8(iIsFirstIngotTen))
		crow.IngotChestConsume = C.int64_t(row.Int64(iIngotChestConsume))
		crow.LastIngotChestAt = C.int64_t(row.Int64(iLastIngotChestAt))
		crow.LastFreePetChestAt = C.int64_t(row.Int64(iLastFreePetChestAt))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerChestState = crow
		if g_Hooks.PlayerChestState != nil {
			g_Hooks.PlayerChestState.Load(&PlayerChestStateRow{row: crow})
		}
	}
}

func loadPlayerCliqueKongfuAttrib(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_clique_kongfu_attrib")
	}
	sql := "SELECT * FROM `player_clique_kongfu_attrib`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iHealth := res.Map("health")
	iAttack := res.Map("attack")
	iDefence := res.Map("defence")
	for _, row := range res.Rows {
		crow := C.New_PlayerCliqueKongfuAttrib()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Health = C.int32_t(row.Int32(iHealth))
		crow.Attack = C.int32_t(row.Int32(iAttack))
		crow.Defence = C.int32_t(row.Int32(iDefence))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerCliqueKongfuAttrib = crow
		if g_Hooks.PlayerCliqueKongfuAttrib != nil {
			g_Hooks.PlayerCliqueKongfuAttrib.Load(&PlayerCliqueKongfuAttribRow{row: crow})
		}
	}
}

func loadPlayerCoins(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_coins")
	}
	sql := "SELECT * FROM `player_coins`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iBuyTime := res.Map("buy_time")
	iDailyCount := res.Map("daily_count")
	iBatchBought := res.Map("batch_bought")
	for _, row := range res.Rows {
		crow := C.New_PlayerCoins()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.BuyTime = C.int64_t(row.Int64(iBuyTime))
		crow.DailyCount = C.int16_t(row.Int16(iDailyCount))
		crow.BatchBought = C.int16_t(row.Int16(iBatchBought))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerCoins = crow
		if g_Hooks.PlayerCoins != nil {
			g_Hooks.PlayerCoins.Load(&PlayerCoinsRow{row: crow})
		}
	}
}

func loadPlayerCornucopia(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_cornucopia")
	}
	sql := "SELECT * FROM `player_cornucopia`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iOpenTime := res.Map("open_time")
	iDailyCount := res.Map("daily_count")
	for _, row := range res.Rows {
		crow := C.New_PlayerCornucopia()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.OpenTime = C.int64_t(row.Int64(iOpenTime))
		crow.DailyCount = C.int16_t(row.Int16(iDailyCount))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerCornucopia = crow
		if g_Hooks.PlayerCornucopia != nil {
			g_Hooks.PlayerCornucopia.Load(&PlayerCornucopiaRow{row: crow})
		}
	}
}

func loadPlayerDailyQuest(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_daily_quest")
	}
	sql := "SELECT * FROM `player_daily_quest`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerDailyQuest), "SELECT MAX(`id`) FROM `player_daily_quest` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iQuestId := res.Map("quest_id")
	iFinishCount := res.Map("finish_count")
	iLastUpdateTime := res.Map("last_update_time")
	iAwardStatus := res.Map("award_status")
	iClass := res.Map("class")
	for _, row := range res.Rows {
		crow := C.New_PlayerDailyQuest()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.QuestId = C.int16_t(row.Int16(iQuestId))
		crow.FinishCount = C.int16_t(row.Int16(iFinishCount))
		crow.LastUpdateTime = C.int64_t(row.Int64(iLastUpdateTime))
		crow.AwardStatus = C.int8_t(row.Int8(iAwardStatus))
		crow.Class = C.int16_t(row.Int16(iClass))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerDailyQuest
		playerDb.PlayerDailyQuest = crow
		if g_Hooks.PlayerDailyQuest != nil {
			g_Hooks.PlayerDailyQuest.Load(&PlayerDailyQuestRow{row: crow})
		}
	}
}

func loadPlayerDailyQuestStarAwardInfo(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_daily_quest_star_award_info")
	}
	sql := "SELECT * FROM `player_daily_quest_star_award_info`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iStars := res.Map("stars")
	iLastupdatetime := res.Map("lastupdatetime")
	iAwarded := res.Map("awarded")
	for _, row := range res.Rows {
		crow := C.New_PlayerDailyQuestStarAwardInfo()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Stars = C.int32_t(row.Int32(iStars))
		crow.Lastupdatetime = C.int64_t(row.Int64(iLastupdatetime))
		row_Awarded := row.Str(iAwarded)
		crow.Awarded = C.CString(string(row_Awarded))
		crow.Awarded_len = C.int(len(row_Awarded))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerDailyQuestStarAwardInfo = crow
		if g_Hooks.PlayerDailyQuestStarAwardInfo != nil {
			g_Hooks.PlayerDailyQuestStarAwardInfo.Load(&PlayerDailyQuestStarAwardInfoRow{row: crow})
		}
	}
}

func loadPlayerDailySignInRecord(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_daily_sign_in_record")
	}
	sql := "SELECT * FROM `player_daily_sign_in_record`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerDailySignInRecord), "SELECT MAX(`id`) FROM `player_daily_sign_in_record` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iSignInTime := res.Map("sign_in_time")
	for _, row := range res.Rows {
		crow := C.New_PlayerDailySignInRecord()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.SignInTime = C.int64_t(row.Int64(iSignInTime))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerDailySignInRecord
		playerDb.PlayerDailySignInRecord = crow
		if g_Hooks.PlayerDailySignInRecord != nil {
			g_Hooks.PlayerDailySignInRecord.Load(&PlayerDailySignInRecordRow{row: crow})
		}
	}
}

func loadPlayerDailySignInState(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_daily_sign_in_state")
	}
	sql := "SELECT * FROM `player_daily_sign_in_state`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iUpdateTime := res.Map("update_time")
	iRecord := res.Map("record")
	iSignedToday := res.Map("signed_today")
	for _, row := range res.Rows {
		crow := C.New_PlayerDailySignInState()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.UpdateTime = C.int64_t(row.Int64(iUpdateTime))
		crow.Record = C.int16_t(row.Int16(iRecord))
		crow.SignedToday = C.int8_t(row.Int8(iSignedToday))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerDailySignInState = crow
		if g_Hooks.PlayerDailySignInState != nil {
			g_Hooks.PlayerDailySignInState.Load(&PlayerDailySignInStateRow{row: crow})
		}
	}
}

func loadPlayerDespairLandCampState(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_despair_land_camp_state")
	}
	sql := "SELECT * FROM `player_despair_land_camp_state`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerDespairLandCampState), "SELECT MAX(`id`) FROM `player_despair_land_camp_state` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iCampType := res.Map("camp_type")
	iTimestamp := res.Map("timestamp")
	iPoint := res.Map("point")
	iKillNum := res.Map("kill_num")
	iDeadNum := res.Map("dead_num")
	iDeadNumBoss := res.Map("dead_num_boss")
	iHurt := res.Map("hurt")
	iBossBattleNum := res.Map("boss_battle_num")
	iDailyBossBattleNum := res.Map("daily_boss_battle_num")
	iBossBattleTimestamp := res.Map("boss_battle_timestamp")
	iAwarded := res.Map("awarded")
	for _, row := range res.Rows {
		crow := C.New_PlayerDespairLandCampState()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.CampType = C.int8_t(row.Int8(iCampType))
		crow.Timestamp = C.int64_t(row.Int64(iTimestamp))
		crow.Point = C.int64_t(row.Int64(iPoint))
		crow.KillNum = C.int64_t(row.Int64(iKillNum))
		crow.DeadNum = C.int64_t(row.Int64(iDeadNum))
		crow.DeadNumBoss = C.int64_t(row.Int64(iDeadNumBoss))
		crow.Hurt = C.int64_t(row.Int64(iHurt))
		crow.BossBattleNum = C.int32_t(row.Int32(iBossBattleNum))
		crow.DailyBossBattleNum = C.int32_t(row.Int32(iDailyBossBattleNum))
		crow.BossBattleTimestamp = C.int64_t(row.Int64(iBossBattleTimestamp))
		crow.Awarded = C.int8_t(row.Int8(iAwarded))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerDespairLandCampState
		playerDb.PlayerDespairLandCampState = crow
		if g_Hooks.PlayerDespairLandCampState != nil {
			g_Hooks.PlayerDespairLandCampState.Load(&PlayerDespairLandCampStateRow{row: crow})
		}
	}
}

func loadPlayerDespairLandCampStateHistory(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_despair_land_camp_state_history")
	}
	sql := "SELECT * FROM `player_despair_land_camp_state_history`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerDespairLandCampStateHistory), "SELECT MAX(`id`) FROM `player_despair_land_camp_state_history` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iCampType := res.Map("camp_type")
	iTimestamp := res.Map("timestamp")
	iPoint := res.Map("point")
	iKillNum := res.Map("kill_num")
	iDeadNum := res.Map("dead_num")
	iDeadNumBoss := res.Map("dead_num_boss")
	iHurt := res.Map("hurt")
	iBossBattleNum := res.Map("boss_battle_num")
	iDailyBossBattleNum := res.Map("daily_boss_battle_num")
	iBossBattleTimestamp := res.Map("boss_battle_timestamp")
	iAwarded := res.Map("awarded")
	for _, row := range res.Rows {
		crow := C.New_PlayerDespairLandCampStateHistory()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.CampType = C.int8_t(row.Int8(iCampType))
		crow.Timestamp = C.int64_t(row.Int64(iTimestamp))
		crow.Point = C.int64_t(row.Int64(iPoint))
		crow.KillNum = C.int64_t(row.Int64(iKillNum))
		crow.DeadNum = C.int64_t(row.Int64(iDeadNum))
		crow.DeadNumBoss = C.int64_t(row.Int64(iDeadNumBoss))
		crow.Hurt = C.int64_t(row.Int64(iHurt))
		crow.BossBattleNum = C.int32_t(row.Int32(iBossBattleNum))
		crow.DailyBossBattleNum = C.int32_t(row.Int32(iDailyBossBattleNum))
		crow.BossBattleTimestamp = C.int64_t(row.Int64(iBossBattleTimestamp))
		crow.Awarded = C.int8_t(row.Int8(iAwarded))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerDespairLandCampStateHistory
		playerDb.PlayerDespairLandCampStateHistory = crow
		if g_Hooks.PlayerDespairLandCampStateHistory != nil {
			g_Hooks.PlayerDespairLandCampStateHistory.Load(&PlayerDespairLandCampStateHistoryRow{row: crow})
		}
	}
}

func loadPlayerDespairLandLevelRecord(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_despair_land_level_record")
	}
	sql := "SELECT * FROM `player_despair_land_level_record`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerDespairLandLevelRecord), "SELECT MAX(`id`) FROM `player_despair_land_level_record` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iCampType := res.Map("camp_type")
	iTimestamp := res.Map("timestamp")
	iLevelId := res.Map("level_id")
	iRound := res.Map("round")
	iPassedTimestamp := res.Map("passed_timestamp")
	iFirstPassedTimestamp := res.Map("first_passed_timestamp")
	iFirstPassedFightnum := res.Map("first_passed_fightnum")
	iPassedAward := res.Map("passed_award")
	iPerfectAward := res.Map("perfect_award")
	iMilestoneAward := res.Map("milestone_award")
	for _, row := range res.Rows {
		crow := C.New_PlayerDespairLandLevelRecord()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.CampType = C.int8_t(row.Int8(iCampType))
		crow.Timestamp = C.int64_t(row.Int64(iTimestamp))
		crow.LevelId = C.int32_t(row.Int32(iLevelId))
		crow.Round = C.int8_t(row.Int8(iRound))
		crow.PassedTimestamp = C.int64_t(row.Int64(iPassedTimestamp))
		crow.FirstPassedTimestamp = C.int64_t(row.Int64(iFirstPassedTimestamp))
		crow.FirstPassedFightnum = C.int32_t(row.Int32(iFirstPassedFightnum))
		crow.PassedAward = C.int8_t(row.Int8(iPassedAward))
		crow.PerfectAward = C.int8_t(row.Int8(iPerfectAward))
		crow.MilestoneAward = C.int8_t(row.Int8(iMilestoneAward))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerDespairLandLevelRecord
		playerDb.PlayerDespairLandLevelRecord = crow
		if g_Hooks.PlayerDespairLandLevelRecord != nil {
			g_Hooks.PlayerDespairLandLevelRecord.Load(&PlayerDespairLandLevelRecordRow{row: crow})
		}
	}
}

func loadPlayerDespairLandState(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_despair_land_state")
	}
	sql := "SELECT * FROM `player_despair_land_state`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iPoint := res.Map("point")
	iKillNum := res.Map("kill_num")
	iDeadNum := res.Map("dead_num")
	iDailyBattleNum := res.Map("daily_battle_num")
	iDailyBattleTimestamp := res.Map("daily_battle_timestamp")
	iDailyBoughtBattleNum := res.Map("daily_bought_battle_num")
	iDailyBoughtTimestamp := res.Map("daily_bought_timestamp")
	iDailyBossBoughtTimestamp := res.Map("daily_boss_bought_timestamp")
	iDailyBossBeastBoughtBattleNum := res.Map("daily_boss_beast_bought_battle_num")
	iDailyBossWalkingDeadBoughtBattleNum := res.Map("daily_boss_walking_dead_bought_battle_num")
	iDailyBossDevilBoughtBattleNum := res.Map("daily_boss_devil_bought_battle_num")
	for _, row := range res.Rows {
		crow := C.New_PlayerDespairLandState()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Point = C.int64_t(row.Int64(iPoint))
		crow.KillNum = C.int64_t(row.Int64(iKillNum))
		crow.DeadNum = C.int64_t(row.Int64(iDeadNum))
		crow.DailyBattleNum = C.int16_t(row.Int16(iDailyBattleNum))
		crow.DailyBattleTimestamp = C.int64_t(row.Int64(iDailyBattleTimestamp))
		crow.DailyBoughtBattleNum = C.int16_t(row.Int16(iDailyBoughtBattleNum))
		crow.DailyBoughtTimestamp = C.int64_t(row.Int64(iDailyBoughtTimestamp))
		crow.DailyBossBoughtTimestamp = C.int64_t(row.Int64(iDailyBossBoughtTimestamp))
		crow.DailyBossBeastBoughtBattleNum = C.int16_t(row.Int16(iDailyBossBeastBoughtBattleNum))
		crow.DailyBossWalkingDeadBoughtBattleNum = C.int16_t(row.Int16(iDailyBossWalkingDeadBoughtBattleNum))
		crow.DailyBossDevilBoughtBattleNum = C.int16_t(row.Int16(iDailyBossDevilBoughtBattleNum))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerDespairLandState = crow
		if g_Hooks.PlayerDespairLandState != nil {
			g_Hooks.PlayerDespairLandState.Load(&PlayerDespairLandStateRow{row: crow})
		}
	}
}

func loadPlayerDrivingSwordEvent(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_driving_sword_event")
	}
	sql := "SELECT * FROM `player_driving_sword_event`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerDrivingSwordEvent), "SELECT MAX(`id`) FROM `player_driving_sword_event` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iCloudId := res.Map("cloud_id")
	iX := res.Map("x")
	iY := res.Map("y")
	iEventType := res.Map("event_type")
	iDataId := res.Map("data_id")
	for _, row := range res.Rows {
		crow := C.New_PlayerDrivingSwordEvent()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.CloudId = C.int16_t(row.Int16(iCloudId))
		crow.X = C.int8_t(row.Int8(iX))
		crow.Y = C.int8_t(row.Int8(iY))
		crow.EventType = C.int8_t(row.Int8(iEventType))
		crow.DataId = C.int8_t(row.Int8(iDataId))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerDrivingSwordEvent
		playerDb.PlayerDrivingSwordEvent = crow
		if g_Hooks.PlayerDrivingSwordEvent != nil {
			g_Hooks.PlayerDrivingSwordEvent.Load(&PlayerDrivingSwordEventRow{row: crow})
		}
	}
}

func loadPlayerDrivingSwordEventExploring(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_driving_sword_event_exploring")
	}
	sql := "SELECT * FROM `player_driving_sword_event_exploring`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerDrivingSwordEventExploring), "SELECT MAX(`id`) FROM `player_driving_sword_event_exploring` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iStatus := res.Map("status")
	iGarrisonStart := res.Map("garrison_start")
	iGarrisonTime := res.Map("garrison_time")
	iAwardTime := res.Map("award_time")
	iRoleId := res.Map("role_id")
	iCloudId := res.Map("cloud_id")
	iX := res.Map("x")
	iY := res.Map("y")
	iDataId := res.Map("data_id")
	for _, row := range res.Rows {
		crow := C.New_PlayerDrivingSwordEventExploring()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Status = C.int8_t(row.Int8(iStatus))
		crow.GarrisonStart = C.int64_t(row.Int64(iGarrisonStart))
		crow.GarrisonTime = C.int64_t(row.Int64(iGarrisonTime))
		crow.AwardTime = C.int64_t(row.Int64(iAwardTime))
		crow.RoleId = C.int8_t(row.Int8(iRoleId))
		crow.CloudId = C.int16_t(row.Int16(iCloudId))
		crow.X = C.int8_t(row.Int8(iX))
		crow.Y = C.int8_t(row.Int8(iY))
		crow.DataId = C.int8_t(row.Int8(iDataId))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerDrivingSwordEventExploring
		playerDb.PlayerDrivingSwordEventExploring = crow
		if g_Hooks.PlayerDrivingSwordEventExploring != nil {
			g_Hooks.PlayerDrivingSwordEventExploring.Load(&PlayerDrivingSwordEventExploringRow{row: crow})
		}
	}
}

func loadPlayerDrivingSwordEventTreasure(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_driving_sword_event_treasure")
	}
	sql := "SELECT * FROM `player_driving_sword_event_treasure`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerDrivingSwordEventTreasure), "SELECT MAX(`id`) FROM `player_driving_sword_event_treasure` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iProgress := res.Map("progress")
	iCloudId := res.Map("cloud_id")
	iX := res.Map("x")
	iY := res.Map("y")
	iDataId := res.Map("data_id")
	for _, row := range res.Rows {
		crow := C.New_PlayerDrivingSwordEventTreasure()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Progress = C.int8_t(row.Int8(iProgress))
		crow.CloudId = C.int16_t(row.Int16(iCloudId))
		crow.X = C.int8_t(row.Int8(iX))
		crow.Y = C.int8_t(row.Int8(iY))
		crow.DataId = C.int8_t(row.Int8(iDataId))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerDrivingSwordEventTreasure
		playerDb.PlayerDrivingSwordEventTreasure = crow
		if g_Hooks.PlayerDrivingSwordEventTreasure != nil {
			g_Hooks.PlayerDrivingSwordEventTreasure.Load(&PlayerDrivingSwordEventTreasureRow{row: crow})
		}
	}
}

func loadPlayerDrivingSwordEventVisiting(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_driving_sword_event_visiting")
	}
	sql := "SELECT * FROM `player_driving_sword_event_visiting`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerDrivingSwordEventVisiting), "SELECT MAX(`id`) FROM `player_driving_sword_event_visiting` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iStatus := res.Map("status")
	iTargetPid := res.Map("target_pid")
	iTargetSideState := res.Map("target_side_state")
	iCloudId := res.Map("cloud_id")
	iX := res.Map("x")
	iY := res.Map("y")
	iDataId := res.Map("data_id")
	iTargetStatus := res.Map("target_status")
	for _, row := range res.Rows {
		crow := C.New_PlayerDrivingSwordEventVisiting()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Status = C.int8_t(row.Int8(iStatus))
		crow.TargetPid = C.int64_t(row.Int64(iTargetPid))
		row_TargetSideState := row.Bin(iTargetSideState)
		crow.TargetSideState = C.CString(string(row_TargetSideState))
		crow.TargetSideState_len = C.int(len(row_TargetSideState))
		crow.CloudId = C.int16_t(row.Int16(iCloudId))
		crow.X = C.int8_t(row.Int8(iX))
		crow.Y = C.int8_t(row.Int8(iY))
		crow.DataId = C.int8_t(row.Int8(iDataId))
		row_TargetStatus := row.Str(iTargetStatus)
		crow.TargetStatus = C.CString(string(row_TargetStatus))
		crow.TargetStatus_len = C.int(len(row_TargetStatus))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerDrivingSwordEventVisiting
		playerDb.PlayerDrivingSwordEventVisiting = crow
		if g_Hooks.PlayerDrivingSwordEventVisiting != nil {
			g_Hooks.PlayerDrivingSwordEventVisiting.Load(&PlayerDrivingSwordEventVisitingRow{row: crow})
		}
	}
}

func loadPlayerDrivingSwordEventmask(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_driving_sword_eventmask")
	}
	sql := "SELECT * FROM `player_driving_sword_eventmask`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerDrivingSwordEventmask), "SELECT MAX(`id`) FROM `player_driving_sword_eventmask` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iCloudId := res.Map("cloud_id")
	iEvents := res.Map("events")
	for _, row := range res.Rows {
		crow := C.New_PlayerDrivingSwordEventmask()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.CloudId = C.int16_t(row.Int16(iCloudId))
		row_Events := row.Bin(iEvents)
		crow.Events = C.CString(string(row_Events))
		crow.Events_len = C.int(len(row_Events))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerDrivingSwordEventmask
		playerDb.PlayerDrivingSwordEventmask = crow
		if g_Hooks.PlayerDrivingSwordEventmask != nil {
			g_Hooks.PlayerDrivingSwordEventmask.Load(&PlayerDrivingSwordEventmaskRow{row: crow})
		}
	}
}

func loadPlayerDrivingSwordInfo(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_driving_sword_info")
	}
	sql := "SELECT * FROM `player_driving_sword_info`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iCurrentCloud := res.Map("current_cloud")
	iHighestCloud := res.Map("highest_cloud")
	iCurrentX := res.Map("current_x")
	iCurrentY := res.Map("current_y")
	iAllowedAction := res.Map("allowed_action")
	iActionRefreshTime := res.Map("action_refresh_time")
	iActionBuyTime := res.Map("action_buy_time")
	iDailyActionBought := res.Map("daily_action_bought")
	for _, row := range res.Rows {
		crow := C.New_PlayerDrivingSwordInfo()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.CurrentCloud = C.int16_t(row.Int16(iCurrentCloud))
		crow.HighestCloud = C.int16_t(row.Int16(iHighestCloud))
		crow.CurrentX = C.int8_t(row.Int8(iCurrentX))
		crow.CurrentY = C.int8_t(row.Int8(iCurrentY))
		crow.AllowedAction = C.int16_t(row.Int16(iAllowedAction))
		crow.ActionRefreshTime = C.int64_t(row.Int64(iActionRefreshTime))
		crow.ActionBuyTime = C.int64_t(row.Int64(iActionBuyTime))
		crow.DailyActionBought = C.int8_t(row.Int8(iDailyActionBought))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerDrivingSwordInfo = crow
		if g_Hooks.PlayerDrivingSwordInfo != nil {
			g_Hooks.PlayerDrivingSwordInfo.Load(&PlayerDrivingSwordInfoRow{row: crow})
		}
	}
}

func loadPlayerDrivingSwordMap(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_driving_sword_map")
	}
	sql := "SELECT * FROM `player_driving_sword_map`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerDrivingSwordMap), "SELECT MAX(`id`) FROM `player_driving_sword_map` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iCloudId := res.Map("cloud_id")
	iShadows := res.Map("shadows")
	iObstacleCount := res.Map("obstacle_count")
	iExploringCount := res.Map("exploring_count")
	iVisitingCount := res.Map("visiting_count")
	iTreasureCount := res.Map("treasure_count")
	for _, row := range res.Rows {
		crow := C.New_PlayerDrivingSwordMap()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.CloudId = C.int16_t(row.Int16(iCloudId))
		row_Shadows := row.Bin(iShadows)
		crow.Shadows = C.CString(string(row_Shadows))
		crow.Shadows_len = C.int(len(row_Shadows))
		crow.ObstacleCount = C.int8_t(row.Int8(iObstacleCount))
		crow.ExploringCount = C.int8_t(row.Int8(iExploringCount))
		crow.VisitingCount = C.int8_t(row.Int8(iVisitingCount))
		crow.TreasureCount = C.int8_t(row.Int8(iTreasureCount))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerDrivingSwordMap
		playerDb.PlayerDrivingSwordMap = crow
		if g_Hooks.PlayerDrivingSwordMap != nil {
			g_Hooks.PlayerDrivingSwordMap.Load(&PlayerDrivingSwordMapRow{row: crow})
		}
	}
}

func loadPlayerEquipment(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_equipment")
	}
	sql := "SELECT * FROM `player_equipment`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerEquipment), "SELECT MAX(`id`) FROM `player_equipment` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iRoleId := res.Map("role_id")
	iWeaponId := res.Map("weapon_id")
	iClothesId := res.Map("clothes_id")
	iAccessoriesId := res.Map("accessories_id")
	iShoeId := res.Map("shoe_id")
	for _, row := range res.Rows {
		crow := C.New_PlayerEquipment()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.RoleId = C.int8_t(row.Int8(iRoleId))
		crow.WeaponId = C.int64_t(row.Int64(iWeaponId))
		crow.ClothesId = C.int64_t(row.Int64(iClothesId))
		crow.AccessoriesId = C.int64_t(row.Int64(iAccessoriesId))
		crow.ShoeId = C.int64_t(row.Int64(iShoeId))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerEquipment
		playerDb.PlayerEquipment = crow
		if g_Hooks.PlayerEquipment != nil {
			g_Hooks.PlayerEquipment.Load(&PlayerEquipmentRow{row: crow})
		}
	}
}

func loadPlayerEventAwardRecord(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_event_award_record")
	}
	sql := "SELECT * FROM `player_event_award_record`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iRecordBytes := res.Map("record_bytes")
	iJsonEventRecord := res.Map("json_event_record")
	for _, row := range res.Rows {
		crow := C.New_PlayerEventAwardRecord()
		crow.Pid = C.int64_t(row.Int64(iPid))
		row_RecordBytes := row.Bin(iRecordBytes)
		crow.RecordBytes = C.CString(string(row_RecordBytes))
		crow.RecordBytes_len = C.int(len(row_RecordBytes))
		row_JsonEventRecord := row.Bin(iJsonEventRecord)
		crow.JsonEventRecord = C.CString(string(row_JsonEventRecord))
		crow.JsonEventRecord_len = C.int(len(row_JsonEventRecord))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerEventAwardRecord = crow
		if g_Hooks.PlayerEventAwardRecord != nil {
			g_Hooks.PlayerEventAwardRecord.Load(&PlayerEventAwardRecordRow{row: crow})
		}
	}
}

func loadPlayerEventDailyOnline(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_event_daily_online")
	}
	sql := "SELECT * FROM `player_event_daily_online`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iTimestamp := res.Map("timestamp")
	iTotalOnlineTime := res.Map("total_online_time")
	iAwardedOnlinetime := res.Map("awarded_onlinetime")
	for _, row := range res.Rows {
		crow := C.New_PlayerEventDailyOnline()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Timestamp = C.int64_t(row.Int64(iTimestamp))
		crow.TotalOnlineTime = C.int32_t(row.Int32(iTotalOnlineTime))
		crow.AwardedOnlinetime = C.int32_t(row.Int32(iAwardedOnlinetime))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerEventDailyOnline = crow
		if g_Hooks.PlayerEventDailyOnline != nil {
			g_Hooks.PlayerEventDailyOnline.Load(&PlayerEventDailyOnlineRow{row: crow})
		}
	}
}

func loadPlayerEventVnDailyRecharge(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_event_vn_daily_recharge")
	}
	sql := "SELECT * FROM `player_event_vn_daily_recharge`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerEventVnDailyRecharge), "SELECT MAX(`id`) FROM `player_event_vn_daily_recharge` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iPage := res.Map("page")
	iTimestamp := res.Map("timestamp")
	iAwarded := res.Map("awarded")
	iDailyRecharge := res.Map("daily_recharge")
	iTotalRecharge := res.Map("total_recharge")
	for _, row := range res.Rows {
		crow := C.New_PlayerEventVnDailyRecharge()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Page = C.int32_t(row.Int32(iPage))
		crow.Timestamp = C.int64_t(row.Int64(iTimestamp))
		crow.Awarded = C.int8_t(row.Int8(iAwarded))
		crow.DailyRecharge = C.int64_t(row.Int64(iDailyRecharge))
		crow.TotalRecharge = C.int64_t(row.Int64(iTotalRecharge))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerEventVnDailyRecharge
		playerDb.PlayerEventVnDailyRecharge = crow
		if g_Hooks.PlayerEventVnDailyRecharge != nil {
			g_Hooks.PlayerEventVnDailyRecharge.Load(&PlayerEventVnDailyRechargeRow{row: crow})
		}
	}
}

func loadPlayerEventVnDragonBallRecord(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_event_vn_dragon_ball_record")
	}
	sql := "SELECT * FROM `player_event_vn_dragon_ball_record`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iItemId := res.Map("item_id")
	iCreatetime := res.Map("createtime")
	for _, row := range res.Rows {
		crow := C.New_PlayerEventVnDragonBallRecord()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.ItemId = C.int16_t(row.Int16(iItemId))
		crow.Createtime = C.int64_t(row.Int64(iCreatetime))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerEventVnDragonBallRecord = crow
		if g_Hooks.PlayerEventVnDragonBallRecord != nil {
			g_Hooks.PlayerEventVnDragonBallRecord.Load(&PlayerEventVnDragonBallRecordRow{row: crow})
		}
	}
}

func loadPlayerEventsFightPower(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_events_fight_power")
	}
	sql := "SELECT * FROM `player_events_fight_power`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iLock := res.Map("lock")
	for _, row := range res.Rows {
		crow := C.New_PlayerEventsFightPower()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Lock = C.int32_t(row.Int32(iLock))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerEventsFightPower = crow
		if g_Hooks.PlayerEventsFightPower != nil {
			g_Hooks.PlayerEventsFightPower.Load(&PlayerEventsFightPowerRow{row: crow})
		}
	}
}

func loadPlayerEventsIngotRecord(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_events_ingot_record")
	}
	sql := "SELECT * FROM `player_events_ingot_record`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iIngotIn := res.Map("ingot_in")
	iIngotInEndTime := res.Map("ingot_in_end_time")
	iIngotOut := res.Map("ingot_out")
	iIngotOutEndTime := res.Map("ingot_out_end_time")
	for _, row := range res.Rows {
		crow := C.New_PlayerEventsIngotRecord()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.IngotIn = C.int64_t(row.Int64(iIngotIn))
		crow.IngotInEndTime = C.int64_t(row.Int64(iIngotInEndTime))
		crow.IngotOut = C.int64_t(row.Int64(iIngotOut))
		crow.IngotOutEndTime = C.int64_t(row.Int64(iIngotOutEndTime))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerEventsIngotRecord = crow
		if g_Hooks.PlayerEventsIngotRecord != nil {
			g_Hooks.PlayerEventsIngotRecord.Load(&PlayerEventsIngotRecordRow{row: crow})
		}
	}
}

func loadPlayerExtendLevel(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_extend_level")
	}
	sql := "SELECT * FROM `player_extend_level`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iCoinPassTime := res.Map("coin_pass_time")
	iExpPassTime := res.Map("exp_pass_time")
	iGhostPassTime := res.Map("ghost_pass_time")
	iPetPassTime := res.Map("pet_pass_time")
	iBuddyPassTime := res.Map("buddy_pass_time")
	iCoinDailyNum := res.Map("coin_daily_num")
	iExpDailyNum := res.Map("exp_daily_num")
	iBuddyDailyNum := res.Map("buddy_daily_num")
	iPetDailyNum := res.Map("pet_daily_num")
	iGhostDailyNum := res.Map("ghost_daily_num")
	iRandBuddyRoleId := res.Map("rand_buddy_role_id")
	iBuddyPos := res.Map("buddy_pos")
	iBuddyTactical := res.Map("buddy_tactical")
	iRolePos := res.Map("role_pos")
	iExpMaxlevel := res.Map("exp_maxlevel")
	iCoinsMaxlevel := res.Map("coins_maxlevel")
	iCoinsBuyNum := res.Map("coins_buy_num")
	iExpBuyNum := res.Map("exp_buy_num")
	iCoinsBuyTime := res.Map("coins_buy_time")
	iExpBuyTime := res.Map("exp_buy_time")
	for _, row := range res.Rows {
		crow := C.New_PlayerExtendLevel()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.CoinPassTime = C.int64_t(row.Int64(iCoinPassTime))
		crow.ExpPassTime = C.int64_t(row.Int64(iExpPassTime))
		crow.GhostPassTime = C.int64_t(row.Int64(iGhostPassTime))
		crow.PetPassTime = C.int64_t(row.Int64(iPetPassTime))
		crow.BuddyPassTime = C.int64_t(row.Int64(iBuddyPassTime))
		crow.CoinDailyNum = C.int8_t(row.Int8(iCoinDailyNum))
		crow.ExpDailyNum = C.int8_t(row.Int8(iExpDailyNum))
		crow.BuddyDailyNum = C.int8_t(row.Int8(iBuddyDailyNum))
		crow.PetDailyNum = C.int8_t(row.Int8(iPetDailyNum))
		crow.GhostDailyNum = C.int8_t(row.Int8(iGhostDailyNum))
		crow.RandBuddyRoleId = C.int8_t(row.Int8(iRandBuddyRoleId))
		crow.BuddyPos = C.int8_t(row.Int8(iBuddyPos))
		crow.BuddyTactical = C.int8_t(row.Int8(iBuddyTactical))
		crow.RolePos = C.int8_t(row.Int8(iRolePos))
		crow.ExpMaxlevel = C.int16_t(row.Int16(iExpMaxlevel))
		crow.CoinsMaxlevel = C.int16_t(row.Int16(iCoinsMaxlevel))
		crow.CoinsBuyNum = C.int16_t(row.Int16(iCoinsBuyNum))
		crow.ExpBuyNum = C.int16_t(row.Int16(iExpBuyNum))
		crow.CoinsBuyTime = C.int64_t(row.Int64(iCoinsBuyTime))
		crow.ExpBuyTime = C.int64_t(row.Int64(iExpBuyTime))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerExtendLevel = crow
		if g_Hooks.PlayerExtendLevel != nil {
			g_Hooks.PlayerExtendLevel.Load(&PlayerExtendLevelRow{row: crow})
		}
	}
}

func loadPlayerExtendQuest(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_extend_quest")
	}
	sql := "SELECT * FROM `player_extend_quest`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerExtendQuest), "SELECT MAX(`id`) FROM `player_extend_quest` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iQuestId := res.Map("quest_id")
	iProgress := res.Map("progress")
	iState := res.Map("state")
	for _, row := range res.Rows {
		crow := C.New_PlayerExtendQuest()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.QuestId = C.int32_t(row.Int32(iQuestId))
		crow.Progress = C.int16_t(row.Int16(iProgress))
		crow.State = C.int8_t(row.Int8(iState))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerExtendQuest
		playerDb.PlayerExtendQuest = crow
		if g_Hooks.PlayerExtendQuest != nil {
			g_Hooks.PlayerExtendQuest.Load(&PlayerExtendQuestRow{row: crow})
		}
	}
}

func loadPlayerFame(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_fame")
	}
	sql := "SELECT * FROM `player_fame`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iFame := res.Map("fame")
	iLevel := res.Map("level")
	iArenaFame := res.Map("arena_fame")
	iMultLevelFame := res.Map("mult_level_fame")
	for _, row := range res.Rows {
		crow := C.New_PlayerFame()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Fame = C.int32_t(row.Int32(iFame))
		crow.Level = C.int16_t(row.Int16(iLevel))
		crow.ArenaFame = C.int32_t(row.Int32(iArenaFame))
		crow.MultLevelFame = C.int32_t(row.Int32(iMultLevelFame))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerFame = crow
		if g_Hooks.PlayerFame != nil {
			g_Hooks.PlayerFame.Load(&PlayerFameRow{row: crow})
		}
	}
}

func loadPlayerFashion(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_fashion")
	}
	sql := "SELECT * FROM `player_fashion`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerFashion), "SELECT MAX(`id`) FROM `player_fashion` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iFashionId := res.Map("fashion_id")
	iExpireTime := res.Map("expire_time")
	for _, row := range res.Rows {
		crow := C.New_PlayerFashion()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.FashionId = C.int16_t(row.Int16(iFashionId))
		crow.ExpireTime = C.int64_t(row.Int64(iExpireTime))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerFashion
		playerDb.PlayerFashion = crow
		if g_Hooks.PlayerFashion != nil {
			g_Hooks.PlayerFashion.Load(&PlayerFashionRow{row: crow})
		}
	}
}

func loadPlayerFashionState(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_fashion_state")
	}
	sql := "SELECT * FROM `player_fashion_state`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iUpdateTime := res.Map("update_time")
	iDressedFashionId := res.Map("dressed_fashion_id")
	for _, row := range res.Rows {
		crow := C.New_PlayerFashionState()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.UpdateTime = C.int64_t(row.Int64(iUpdateTime))
		crow.DressedFashionId = C.int16_t(row.Int16(iDressedFashionId))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerFashionState = crow
		if g_Hooks.PlayerFashionState != nil {
			g_Hooks.PlayerFashionState.Load(&PlayerFashionStateRow{row: crow})
		}
	}
}

func loadPlayerFateBoxState(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_fate_box_state")
	}
	sql := "SELECT * FROM `player_fate_box_state`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iLock := res.Map("lock")
	iStarBoxFreeDrawTimestamp := res.Map("star_box_free_draw_timestamp")
	iStarBoxDrawCount := res.Map("star_box_draw_count")
	iMoonBoxFreeDrawTimestamp := res.Map("moon_box_free_draw_timestamp")
	iMoonBoxDrawCount := res.Map("moon_box_draw_count")
	iSunBoxFreeDrawTimestamp := res.Map("sun_box_free_draw_timestamp")
	iSunBoxDrawCount := res.Map("sun_box_draw_count")
	iHunyuanBoxFreeDrawTimestamp := res.Map("hunyuan_box_free_draw_timestamp")
	iHunyuanBoxDrawCount := res.Map("hunyuan_box_draw_count")
	for _, row := range res.Rows {
		crow := C.New_PlayerFateBoxState()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Lock = C.int32_t(row.Int32(iLock))
		crow.StarBoxFreeDrawTimestamp = C.int64_t(row.Int64(iStarBoxFreeDrawTimestamp))
		crow.StarBoxDrawCount = C.int32_t(row.Int32(iStarBoxDrawCount))
		crow.MoonBoxFreeDrawTimestamp = C.int64_t(row.Int64(iMoonBoxFreeDrawTimestamp))
		crow.MoonBoxDrawCount = C.int32_t(row.Int32(iMoonBoxDrawCount))
		crow.SunBoxFreeDrawTimestamp = C.int64_t(row.Int64(iSunBoxFreeDrawTimestamp))
		crow.SunBoxDrawCount = C.int32_t(row.Int32(iSunBoxDrawCount))
		crow.HunyuanBoxFreeDrawTimestamp = C.int64_t(row.Int64(iHunyuanBoxFreeDrawTimestamp))
		crow.HunyuanBoxDrawCount = C.int32_t(row.Int32(iHunyuanBoxDrawCount))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerFateBoxState = crow
		if g_Hooks.PlayerFateBoxState != nil {
			g_Hooks.PlayerFateBoxState.Load(&PlayerFateBoxStateRow{row: crow})
		}
	}
}

func loadPlayerFightNum(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_fight_num")
	}
	sql := "SELECT * FROM `player_fight_num`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iFightNum := res.Map("fight_num")
	for _, row := range res.Rows {
		crow := C.New_PlayerFightNum()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.FightNum = C.int32_t(row.Int32(iFightNum))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerFightNum = crow
		if g_Hooks.PlayerFightNum != nil {
			g_Hooks.PlayerFightNum.Load(&PlayerFightNumRow{row: crow})
		}
	}
}

func loadPlayerFirstCharge(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_first_charge")
	}
	sql := "SELECT * FROM `player_first_charge`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerFirstCharge), "SELECT MAX(`id`) FROM `player_first_charge` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iProductId := res.Map("product_id")
	for _, row := range res.Rows {
		crow := C.New_PlayerFirstCharge()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		row_ProductId := row.Str(iProductId)
		crow.ProductId = C.CString(string(row_ProductId))
		crow.ProductId_len = C.int(len(row_ProductId))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerFirstCharge
		playerDb.PlayerFirstCharge = crow
		if g_Hooks.PlayerFirstCharge != nil {
			g_Hooks.PlayerFirstCharge.Load(&PlayerFirstChargeRow{row: crow})
		}
	}
}

func loadPlayerFormation(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_formation")
	}
	sql := "SELECT * FROM `player_formation`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iPos0 := res.Map("pos0")
	iPos1 := res.Map("pos1")
	iPos2 := res.Map("pos2")
	iPos3 := res.Map("pos3")
	iPos4 := res.Map("pos4")
	iPos5 := res.Map("pos5")
	iPos6 := res.Map("pos6")
	iPos7 := res.Map("pos7")
	iPos8 := res.Map("pos8")
	for _, row := range res.Rows {
		crow := C.New_PlayerFormation()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Pos0 = C.int8_t(row.Int8(iPos0))
		crow.Pos1 = C.int8_t(row.Int8(iPos1))
		crow.Pos2 = C.int8_t(row.Int8(iPos2))
		crow.Pos3 = C.int8_t(row.Int8(iPos3))
		crow.Pos4 = C.int8_t(row.Int8(iPos4))
		crow.Pos5 = C.int8_t(row.Int8(iPos5))
		crow.Pos6 = C.int8_t(row.Int8(iPos6))
		crow.Pos7 = C.int8_t(row.Int8(iPos7))
		crow.Pos8 = C.int8_t(row.Int8(iPos8))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerFormation = crow
		if g_Hooks.PlayerFormation != nil {
			g_Hooks.PlayerFormation.Load(&PlayerFormationRow{row: crow})
		}
	}
}

func loadPlayerFuncKey(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_func_key")
	}
	sql := "SELECT * FROM `player_func_key`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iKey := res.Map("key")
	iPlayedKey := res.Map("played_key")
	iUniqueKey := res.Map("unique_key")
	for _, row := range res.Rows {
		crow := C.New_PlayerFuncKey()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Key = C.int16_t(row.Int16(iKey))
		crow.PlayedKey = C.int16_t(row.Int16(iPlayedKey))
		crow.UniqueKey = C.int64_t(row.Int64(iUniqueKey))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerFuncKey = crow
		if g_Hooks.PlayerFuncKey != nil {
			g_Hooks.PlayerFuncKey.Load(&PlayerFuncKeyRow{row: crow})
		}
	}
}

func loadPlayerGhost(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_ghost")
	}
	sql := "SELECT * FROM `player_ghost`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerGhost), "SELECT MAX(`id`) FROM `player_ghost` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iGhostId := res.Map("ghost_id")
	iStar := res.Map("star")
	iLevel := res.Map("level")
	iExp := res.Map("exp")
	iPos := res.Map("pos")
	iIsNew := res.Map("is_new")
	iRoleId := res.Map("role_id")
	iSkillLevel := res.Map("skill_level")
	iRelationId := res.Map("relation_id")
	iAddGrowth := res.Map("add_growth")
	for _, row := range res.Rows {
		crow := C.New_PlayerGhost()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.GhostId = C.int16_t(row.Int16(iGhostId))
		crow.Star = C.int8_t(row.Int8(iStar))
		crow.Level = C.int16_t(row.Int16(iLevel))
		crow.Exp = C.int64_t(row.Int64(iExp))
		crow.Pos = C.int16_t(row.Int16(iPos))
		crow.IsNew = C.int8_t(row.Int8(iIsNew))
		crow.RoleId = C.int8_t(row.Int8(iRoleId))
		crow.SkillLevel = C.int16_t(row.Int16(iSkillLevel))
		crow.RelationId = C.int64_t(row.Int64(iRelationId))
		crow.AddGrowth = C.int16_t(row.Int16(iAddGrowth))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerGhost
		playerDb.PlayerGhost = crow
		if g_Hooks.PlayerGhost != nil {
			g_Hooks.PlayerGhost.Load(&PlayerGhostRow{row: crow})
		}
	}
}

func loadPlayerGhostEquipment(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_ghost_equipment")
	}
	sql := "SELECT * FROM `player_ghost_equipment`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerGhostEquipment), "SELECT MAX(`id`) FROM `player_ghost_equipment` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iRoleId := res.Map("role_id")
	iGhostPower := res.Map("ghost_power")
	iPos1 := res.Map("pos1")
	iPos2 := res.Map("pos2")
	iPos3 := res.Map("pos3")
	iPos4 := res.Map("pos4")
	for _, row := range res.Rows {
		crow := C.New_PlayerGhostEquipment()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.RoleId = C.int8_t(row.Int8(iRoleId))
		crow.GhostPower = C.int32_t(row.Int32(iGhostPower))
		crow.Pos1 = C.int64_t(row.Int64(iPos1))
		crow.Pos2 = C.int64_t(row.Int64(iPos2))
		crow.Pos3 = C.int64_t(row.Int64(iPos3))
		crow.Pos4 = C.int64_t(row.Int64(iPos4))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerGhostEquipment
		playerDb.PlayerGhostEquipment = crow
		if g_Hooks.PlayerGhostEquipment != nil {
			g_Hooks.PlayerGhostEquipment.Load(&PlayerGhostEquipmentRow{row: crow})
		}
	}
}

func loadPlayerGhostState(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_ghost_state")
	}
	sql := "SELECT * FROM `player_ghost_state`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iTrainByIngotNum := res.Map("train_by_ingot_num")
	iTrainByIngotTime := res.Map("train_by_ingot_time")
	iLastFlushTime := res.Map("last_flush_time")
	iGhostFightNum := res.Map("ghost_fight_num")
	for _, row := range res.Rows {
		crow := C.New_PlayerGhostState()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.TrainByIngotNum = C.int32_t(row.Int32(iTrainByIngotNum))
		crow.TrainByIngotTime = C.int64_t(row.Int64(iTrainByIngotTime))
		crow.LastFlushTime = C.int64_t(row.Int64(iLastFlushTime))
		crow.GhostFightNum = C.int64_t(row.Int64(iGhostFightNum))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerGhostState = crow
		if g_Hooks.PlayerGhostState != nil {
			g_Hooks.PlayerGhostState.Load(&PlayerGhostStateRow{row: crow})
		}
	}
}

func loadPlayerGlobalChatState(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_global_chat_state")
	}
	sql := "SELECT * FROM `player_global_chat_state`"
	if playerId == 0 {
		sql += " WHERE  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iBanUntil := res.Map("ban_until")
	for _, row := range res.Rows {
		crow := C.New_PlayerGlobalChatState()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.BanUntil = C.int64_t(row.Int64(iBanUntil))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerGlobalChatState = crow
		if g_Hooks.PlayerGlobalChatState != nil {
			g_Hooks.PlayerGlobalChatState.Load(&PlayerGlobalChatStateRow{row: crow})
		}
	}
}

func loadPlayerGlobalCliqueBuilding(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_global_clique_building")
	}
	sql := "SELECT * FROM `player_global_clique_building`"
	if playerId == 0 {
		sql += " WHERE  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iSilverExchangeTime := res.Map("silver_exchange_time")
	iGoldExchangeNum := res.Map("gold_exchange_num")
	iSilverExchangeNum := res.Map("silver_exchange_num")
	iDonateCoinsCenterBuilding := res.Map("donate_coins_center_building")
	iDonateCoinsTempleBuilding := res.Map("donate_coins_temple_building")
	iDonateCoinsBankBuilding := res.Map("donate_coins_bank_building")
	iDonateCoinsHealthBuilding := res.Map("donate_coins_health_building")
	iDonateCoinsAttackBuilding := res.Map("donate_coins_attack_building")
	iDonateCoinsDefenseBuilding := res.Map("donate_coins_defense_building")
	iDonateCoinsStoreBuilding := res.Map("donate_coins_store_building")
	iHealthLevel := res.Map("health_level")
	iAttackLevel := res.Map("attack_level")
	iDefenseLevel := res.Map("defense_level")
	iWorshipTime := res.Map("worship_time")
	iDonateCoinsCenterBuildingTime := res.Map("donate_coins_center_building_time")
	iGlodExchangeTime := res.Map("glod_exchange_time")
	iWorshipType := res.Map("worship_type")
	for _, row := range res.Rows {
		crow := C.New_PlayerGlobalCliqueBuilding()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.SilverExchangeTime = C.int64_t(row.Int64(iSilverExchangeTime))
		crow.GoldExchangeNum = C.int16_t(row.Int16(iGoldExchangeNum))
		crow.SilverExchangeNum = C.int16_t(row.Int16(iSilverExchangeNum))
		crow.DonateCoinsCenterBuilding = C.int64_t(row.Int64(iDonateCoinsCenterBuilding))
		crow.DonateCoinsTempleBuilding = C.int64_t(row.Int64(iDonateCoinsTempleBuilding))
		crow.DonateCoinsBankBuilding = C.int64_t(row.Int64(iDonateCoinsBankBuilding))
		crow.DonateCoinsHealthBuilding = C.int64_t(row.Int64(iDonateCoinsHealthBuilding))
		crow.DonateCoinsAttackBuilding = C.int64_t(row.Int64(iDonateCoinsAttackBuilding))
		crow.DonateCoinsDefenseBuilding = C.int64_t(row.Int64(iDonateCoinsDefenseBuilding))
		crow.DonateCoinsStoreBuilding = C.int64_t(row.Int64(iDonateCoinsStoreBuilding))
		crow.HealthLevel = C.int16_t(row.Int16(iHealthLevel))
		crow.AttackLevel = C.int16_t(row.Int16(iAttackLevel))
		crow.DefenseLevel = C.int16_t(row.Int16(iDefenseLevel))
		crow.WorshipTime = C.int64_t(row.Int64(iWorshipTime))
		crow.DonateCoinsCenterBuildingTime = C.int64_t(row.Int64(iDonateCoinsCenterBuildingTime))
		crow.GlodExchangeTime = C.int64_t(row.Int64(iGlodExchangeTime))
		crow.WorshipType = C.int8_t(row.Int8(iWorshipType))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerGlobalCliqueBuilding = crow
		if g_Hooks.PlayerGlobalCliqueBuilding != nil {
			g_Hooks.PlayerGlobalCliqueBuilding.Load(&PlayerGlobalCliqueBuildingRow{row: crow})
		}
	}
}

func loadPlayerGlobalCliqueBuildingQuest(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_global_clique_building_quest")
	}
	sql := "SELECT * FROM `player_global_clique_building_quest`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerGlobalCliqueBuildingQuest), "SELECT MAX(`id`) FROM `player_global_clique_building_quest` ") {
			return
		}
		sql += " WHERE  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iQuestId := res.Map("quest_id")
	iAwardStatus := res.Map("award_status")
	iBuildingType := res.Map("building_type")
	for _, row := range res.Rows {
		crow := C.New_PlayerGlobalCliqueBuildingQuest()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.QuestId = C.int16_t(row.Int16(iQuestId))
		crow.AwardStatus = C.int8_t(row.Int8(iAwardStatus))
		crow.BuildingType = C.int16_t(row.Int16(iBuildingType))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerGlobalCliqueBuildingQuest
		playerDb.PlayerGlobalCliqueBuildingQuest = crow
		if g_Hooks.PlayerGlobalCliqueBuildingQuest != nil {
			g_Hooks.PlayerGlobalCliqueBuildingQuest.Load(&PlayerGlobalCliqueBuildingQuestRow{row: crow})
		}
	}
}

func loadPlayerGlobalCliqueDailyQuest(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_global_clique_daily_quest")
	}
	sql := "SELECT * FROM `player_global_clique_daily_quest`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerGlobalCliqueDailyQuest), "SELECT MAX(`id`) FROM `player_global_clique_daily_quest` ") {
			return
		}
		sql += " WHERE  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iQuestId := res.Map("quest_id")
	iFinishCount := res.Map("finish_count")
	iLastUpdateTime := res.Map("last_update_time")
	iAwardStatus := res.Map("award_status")
	iClass := res.Map("class")
	for _, row := range res.Rows {
		crow := C.New_PlayerGlobalCliqueDailyQuest()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.QuestId = C.int16_t(row.Int16(iQuestId))
		crow.FinishCount = C.int16_t(row.Int16(iFinishCount))
		crow.LastUpdateTime = C.int64_t(row.Int64(iLastUpdateTime))
		crow.AwardStatus = C.int8_t(row.Int8(iAwardStatus))
		crow.Class = C.int16_t(row.Int16(iClass))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerGlobalCliqueDailyQuest
		playerDb.PlayerGlobalCliqueDailyQuest = crow
		if g_Hooks.PlayerGlobalCliqueDailyQuest != nil {
			g_Hooks.PlayerGlobalCliqueDailyQuest.Load(&PlayerGlobalCliqueDailyQuestRow{row: crow})
		}
	}
}

func loadPlayerGlobalCliqueEscort(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_global_clique_escort")
	}
	sql := "SELECT * FROM `player_global_clique_escort`"
	if playerId == 0 {
		sql += " WHERE  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iDailyEscortTimestamp := res.Map("daily_escort_timestamp")
	iDailyEscortNum := res.Map("daily_escort_num")
	iEscortBoatType := res.Map("escort_boat_type")
	iStatus := res.Map("status")
	iHijacked := res.Map("hijacked")
	for _, row := range res.Rows {
		crow := C.New_PlayerGlobalCliqueEscort()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.DailyEscortTimestamp = C.int64_t(row.Int64(iDailyEscortTimestamp))
		crow.DailyEscortNum = C.int16_t(row.Int16(iDailyEscortNum))
		crow.EscortBoatType = C.int16_t(row.Int16(iEscortBoatType))
		crow.Status = C.int8_t(row.Int8(iStatus))
		crow.Hijacked = C.int8_t(row.Int8(iHijacked))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerGlobalCliqueEscort = crow
		if g_Hooks.PlayerGlobalCliqueEscort != nil {
			g_Hooks.PlayerGlobalCliqueEscort.Load(&PlayerGlobalCliqueEscortRow{row: crow})
		}
	}
}

func loadPlayerGlobalCliqueEscortMessage(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_global_clique_escort_message")
	}
	sql := "SELECT * FROM `player_global_clique_escort_message`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerGlobalCliqueEscortMessage), "SELECT MAX(`id`) FROM `player_global_clique_escort_message` ") {
			return
		}
		sql += " WHERE  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iTplId := res.Map("tpl_id")
	iParameters := res.Map("parameters")
	iTimestamp := res.Map("timestamp")
	for _, row := range res.Rows {
		crow := C.New_PlayerGlobalCliqueEscortMessage()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.TplId = C.int16_t(row.Int16(iTplId))
		row_Parameters := row.Str(iParameters)
		crow.Parameters = C.CString(string(row_Parameters))
		crow.Parameters_len = C.int(len(row_Parameters))
		crow.Timestamp = C.int64_t(row.Int64(iTimestamp))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerGlobalCliqueEscortMessage
		playerDb.PlayerGlobalCliqueEscortMessage = crow
		if g_Hooks.PlayerGlobalCliqueEscortMessage != nil {
			g_Hooks.PlayerGlobalCliqueEscortMessage.Load(&PlayerGlobalCliqueEscortMessageRow{row: crow})
		}
	}
}

func loadPlayerGlobalCliqueHijack(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_global_clique_hijack")
	}
	sql := "SELECT * FROM `player_global_clique_hijack`"
	if playerId == 0 {
		sql += " WHERE  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iDailyHijackTimestamp := res.Map("daily_hijack_timestamp")
	iDailyHijackNum := res.Map("daily_hijack_num")
	iHijackedBoatType := res.Map("hijacked_boat_type")
	iStatus := res.Map("status")
	iBuyTimestamp := res.Map("buy_timestamp")
	iBuyNum := res.Map("buy_num")
	for _, row := range res.Rows {
		crow := C.New_PlayerGlobalCliqueHijack()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.DailyHijackTimestamp = C.int64_t(row.Int64(iDailyHijackTimestamp))
		crow.DailyHijackNum = C.int16_t(row.Int16(iDailyHijackNum))
		crow.HijackedBoatType = C.int16_t(row.Int16(iHijackedBoatType))
		crow.Status = C.int8_t(row.Int8(iStatus))
		crow.BuyTimestamp = C.int64_t(row.Int64(iBuyTimestamp))
		crow.BuyNum = C.int16_t(row.Int16(iBuyNum))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerGlobalCliqueHijack = crow
		if g_Hooks.PlayerGlobalCliqueHijack != nil {
			g_Hooks.PlayerGlobalCliqueHijack.Load(&PlayerGlobalCliqueHijackRow{row: crow})
		}
	}
}

func loadPlayerGlobalCliqueInfo(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_global_clique_info")
	}
	sql := "SELECT * FROM `player_global_clique_info`"
	if playerId == 0 {
		sql += " WHERE  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iCliqueId := res.Map("clique_id")
	iJoinTime := res.Map("join_time")
	iContrib := res.Map("contrib")
	iContribClearTime := res.Map("contrib_clear_time")
	iDonateCoinsTime := res.Map("donate_coins_time")
	iDailyDonateCoins := res.Map("daily_donate_coins")
	iTotalContrib := res.Map("total_contrib")
	for _, row := range res.Rows {
		crow := C.New_PlayerGlobalCliqueInfo()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.CliqueId = C.int64_t(row.Int64(iCliqueId))
		crow.JoinTime = C.int64_t(row.Int64(iJoinTime))
		crow.Contrib = C.int64_t(row.Int64(iContrib))
		crow.ContribClearTime = C.int64_t(row.Int64(iContribClearTime))
		crow.DonateCoinsTime = C.int64_t(row.Int64(iDonateCoinsTime))
		crow.DailyDonateCoins = C.int64_t(row.Int64(iDailyDonateCoins))
		crow.TotalContrib = C.int64_t(row.Int64(iTotalContrib))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerGlobalCliqueInfo = crow
		if g_Hooks.PlayerGlobalCliqueInfo != nil {
			g_Hooks.PlayerGlobalCliqueInfo.Load(&PlayerGlobalCliqueInfoRow{row: crow})
		}
	}
}

func loadPlayerGlobalCliqueKongfu(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_global_clique_kongfu")
	}
	sql := "SELECT * FROM `player_global_clique_kongfu`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerGlobalCliqueKongfu), "SELECT MAX(`id`) FROM `player_global_clique_kongfu` ") {
			return
		}
		sql += " WHERE  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iKongfuId := res.Map("kongfu_id")
	iLevel := res.Map("level")
	for _, row := range res.Rows {
		crow := C.New_PlayerGlobalCliqueKongfu()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.KongfuId = C.int32_t(row.Int32(iKongfuId))
		crow.Level = C.int16_t(row.Int16(iLevel))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerGlobalCliqueKongfu
		playerDb.PlayerGlobalCliqueKongfu = crow
		if g_Hooks.PlayerGlobalCliqueKongfu != nil {
			g_Hooks.PlayerGlobalCliqueKongfu.Load(&PlayerGlobalCliqueKongfuRow{row: crow})
		}
	}
}

func loadPlayerGlobalFriend(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_global_friend")
	}
	sql := "SELECT * FROM `player_global_friend`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerGlobalFriend), "SELECT MAX(`id`) FROM `player_global_friend` ") {
			return
		}
		sql += " WHERE  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iFriendPid := res.Map("friend_pid")
	iFriendNick := res.Map("friend_nick")
	iFriendRoleId := res.Map("friend_role_id")
	iFriendMode := res.Map("friend_mode")
	iLastChatTime := res.Map("last_chat_time")
	iFriendTime := res.Map("friend_time")
	iSendHeartTime := res.Map("send_heart_time")
	iBlockMode := res.Map("block_mode")
	for _, row := range res.Rows {
		crow := C.New_PlayerGlobalFriend()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.FriendPid = C.int64_t(row.Int64(iFriendPid))
		row_FriendNick := row.Str(iFriendNick)
		crow.FriendNick = C.CString(string(row_FriendNick))
		crow.FriendNick_len = C.int(len(row_FriendNick))
		crow.FriendRoleId = C.int8_t(row.Int8(iFriendRoleId))
		crow.FriendMode = C.int8_t(row.Int8(iFriendMode))
		crow.LastChatTime = C.int64_t(row.Int64(iLastChatTime))
		crow.FriendTime = C.int64_t(row.Int64(iFriendTime))
		crow.SendHeartTime = C.int64_t(row.Int64(iSendHeartTime))
		crow.BlockMode = C.int8_t(row.Int8(iBlockMode))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerGlobalFriend
		playerDb.PlayerGlobalFriend = crow
		if g_Hooks.PlayerGlobalFriend != nil {
			g_Hooks.PlayerGlobalFriend.Load(&PlayerGlobalFriendRow{row: crow})
		}
	}
}

func loadPlayerGlobalFriendChat(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_global_friend_chat")
	}
	sql := "SELECT * FROM `player_global_friend_chat`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerGlobalFriendChat), "SELECT MAX(`id`) FROM `player_global_friend_chat` ") {
			return
		}
		sql += " WHERE  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iFriendPid := res.Map("friend_pid")
	iMode := res.Map("mode")
	iSendTime := res.Map("send_time")
	iMessage := res.Map("message")
	for _, row := range res.Rows {
		crow := C.New_PlayerGlobalFriendChat()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.FriendPid = C.int64_t(row.Int64(iFriendPid))
		crow.Mode = C.int8_t(row.Int8(iMode))
		crow.SendTime = C.int64_t(row.Int64(iSendTime))
		row_Message := row.Str(iMessage)
		crow.Message = C.CString(string(row_Message))
		crow.Message_len = C.int(len(row_Message))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerGlobalFriendChat
		playerDb.PlayerGlobalFriendChat = crow
		if g_Hooks.PlayerGlobalFriendChat != nil {
			g_Hooks.PlayerGlobalFriendChat.Load(&PlayerGlobalFriendChatRow{row: crow})
		}
	}
}

func loadPlayerGlobalFriendState(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_global_friend_state")
	}
	sql := "SELECT * FROM `player_global_friend_state`"
	if playerId == 0 {
		sql += " WHERE  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iDeleteDayCount := res.Map("delete_day_count")
	iDeleteTime := res.Map("delete_time")
	iExistOfflineChat := res.Map("exist_offline_chat")
	iPlatformFriendNum := res.Map("platform_friend_num")
	for _, row := range res.Rows {
		crow := C.New_PlayerGlobalFriendState()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.DeleteDayCount = C.int32_t(row.Int32(iDeleteDayCount))
		crow.DeleteTime = C.int64_t(row.Int64(iDeleteTime))
		crow.ExistOfflineChat = C.int8_t(row.Int8(iExistOfflineChat))
		crow.PlatformFriendNum = C.int32_t(row.Int32(iPlatformFriendNum))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerGlobalFriendState = crow
		if g_Hooks.PlayerGlobalFriendState != nil {
			g_Hooks.PlayerGlobalFriendState.Load(&PlayerGlobalFriendStateRow{row: crow})
		}
	}
}

func loadPlayerGlobalGiftCardRecord(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_global_gift_card_record")
	}
	sql := "SELECT * FROM `player_global_gift_card_record`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerGlobalGiftCardRecord), "SELECT MAX(`id`) FROM `player_global_gift_card_record` ") {
			return
		}
		sql += " WHERE  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iCode := res.Map("code")
	for _, row := range res.Rows {
		crow := C.New_PlayerGlobalGiftCardRecord()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		row_Code := row.Str(iCode)
		crow.Code = C.CString(string(row_Code))
		crow.Code_len = C.int(len(row_Code))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerGlobalGiftCardRecord
		playerDb.PlayerGlobalGiftCardRecord = crow
		if g_Hooks.PlayerGlobalGiftCardRecord != nil {
			g_Hooks.PlayerGlobalGiftCardRecord.Load(&PlayerGlobalGiftCardRecordRow{row: crow})
		}
	}
}

func loadPlayerGlobalMailState(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_global_mail_state")
	}
	sql := "SELECT * FROM `player_global_mail_state`"
	if playerId == 0 {
		sql += " WHERE  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iMaxTimestamp := res.Map("max_timestamp")
	for _, row := range res.Rows {
		crow := C.New_PlayerGlobalMailState()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.MaxTimestamp = C.int64_t(row.Int64(iMaxTimestamp))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerGlobalMailState = crow
		if g_Hooks.PlayerGlobalMailState != nil {
			g_Hooks.PlayerGlobalMailState.Load(&PlayerGlobalMailStateRow{row: crow})
		}
	}
}

func loadPlayerGlobalWorldChatState(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_global_world_chat_state")
	}
	sql := "SELECT * FROM `player_global_world_chat_state`"
	if playerId == 0 {
		sql += " WHERE  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iTimestamp := res.Map("timestamp")
	iDailyNum := res.Map("daily_num")
	for _, row := range res.Rows {
		crow := C.New_PlayerGlobalWorldChatState()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Timestamp = C.int64_t(row.Int64(iTimestamp))
		crow.DailyNum = C.int16_t(row.Int16(iDailyNum))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerGlobalWorldChatState = crow
		if g_Hooks.PlayerGlobalWorldChatState != nil {
			g_Hooks.PlayerGlobalWorldChatState.Load(&PlayerGlobalWorldChatStateRow{row: crow})
		}
	}
}

func loadPlayerHardLevel(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_hard_level")
	}
	sql := "SELECT * FROM `player_hard_level`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iLock := res.Map("lock")
	iAwardLock := res.Map("award_lock")
	for _, row := range res.Rows {
		crow := C.New_PlayerHardLevel()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Lock = C.int32_t(row.Int32(iLock))
		crow.AwardLock = C.int32_t(row.Int32(iAwardLock))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerHardLevel = crow
		if g_Hooks.PlayerHardLevel != nil {
			g_Hooks.PlayerHardLevel.Load(&PlayerHardLevelRow{row: crow})
		}
	}
}

func loadPlayerHardLevelRecord(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_hard_level_record")
	}
	sql := "SELECT * FROM `player_hard_level_record`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerHardLevelRecord), "SELECT MAX(`id`) FROM `player_hard_level_record` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iLevelId := res.Map("level_id")
	iOpenTime := res.Map("open_time")
	iScore := res.Map("score")
	iRound := res.Map("round")
	iDailyNum := res.Map("daily_num")
	iLastEnterTime := res.Map("last_enter_time")
	iBuyTimes := res.Map("buy_times")
	iBuyUpdateTime := res.Map("buy_update_time")
	for _, row := range res.Rows {
		crow := C.New_PlayerHardLevelRecord()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.LevelId = C.int32_t(row.Int32(iLevelId))
		crow.OpenTime = C.int64_t(row.Int64(iOpenTime))
		crow.Score = C.int32_t(row.Int32(iScore))
		crow.Round = C.int8_t(row.Int8(iRound))
		crow.DailyNum = C.int8_t(row.Int8(iDailyNum))
		crow.LastEnterTime = C.int64_t(row.Int64(iLastEnterTime))
		crow.BuyTimes = C.int16_t(row.Int16(iBuyTimes))
		crow.BuyUpdateTime = C.int64_t(row.Int64(iBuyUpdateTime))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerHardLevelRecord
		playerDb.PlayerHardLevelRecord = crow
		if g_Hooks.PlayerHardLevelRecord != nil {
			g_Hooks.PlayerHardLevelRecord.Load(&PlayerHardLevelRecordRow{row: crow})
		}
	}
}

func loadPlayerHeart(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_heart")
	}
	sql := "SELECT * FROM `player_heart`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iValue := res.Map("value")
	iUpdateTime := res.Map("update_time")
	iAddDayCount := res.Map("add_day_count")
	iAddTime := res.Map("add_time")
	iRecoverDayCount := res.Map("recover_day_count")
	for _, row := range res.Rows {
		crow := C.New_PlayerHeart()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Value = C.int16_t(row.Int16(iValue))
		crow.UpdateTime = C.int64_t(row.Int64(iUpdateTime))
		crow.AddDayCount = C.int32_t(row.Int32(iAddDayCount))
		crow.AddTime = C.int64_t(row.Int64(iAddTime))
		crow.RecoverDayCount = C.int16_t(row.Int16(iRecoverDayCount))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerHeart = crow
		if g_Hooks.PlayerHeart != nil {
			g_Hooks.PlayerHeart.Load(&PlayerHeartRow{row: crow})
		}
	}
}

func loadPlayerHeartDraw(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_heart_draw")
	}
	sql := "SELECT * FROM `player_heart_draw`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerHeartDraw), "SELECT MAX(`id`) FROM `player_heart_draw` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iDrawType := res.Map("draw_type")
	iDailyNum := res.Map("daily_num")
	iDrawTime := res.Map("draw_time")
	for _, row := range res.Rows {
		crow := C.New_PlayerHeartDraw()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.DrawType = C.int8_t(row.Int8(iDrawType))
		crow.DailyNum = C.int8_t(row.Int8(iDailyNum))
		crow.DrawTime = C.int64_t(row.Int64(iDrawTime))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerHeartDraw
		playerDb.PlayerHeartDraw = crow
		if g_Hooks.PlayerHeartDraw != nil {
			g_Hooks.PlayerHeartDraw.Load(&PlayerHeartDrawRow{row: crow})
		}
	}
}

func loadPlayerHeartDrawCardRecord(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_heart_draw_card_record")
	}
	sql := "SELECT * FROM `player_heart_draw_card_record`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerHeartDrawCardRecord), "SELECT MAX(`id`) FROM `player_heart_draw_card_record` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iAwardType := res.Map("award_type")
	iAwardNum := res.Map("award_num")
	iItemId := res.Map("item_id")
	iDrawTime := res.Map("draw_time")
	for _, row := range res.Rows {
		crow := C.New_PlayerHeartDrawCardRecord()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.AwardType = C.int8_t(row.Int8(iAwardType))
		crow.AwardNum = C.int16_t(row.Int16(iAwardNum))
		crow.ItemId = C.int16_t(row.Int16(iItemId))
		crow.DrawTime = C.int64_t(row.Int64(iDrawTime))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerHeartDrawCardRecord
		playerDb.PlayerHeartDrawCardRecord = crow
		if g_Hooks.PlayerHeartDrawCardRecord != nil {
			g_Hooks.PlayerHeartDrawCardRecord.Load(&PlayerHeartDrawCardRecordRow{row: crow})
		}
	}
}

func loadPlayerHeartDrawWheelRecord(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_heart_draw_wheel_record")
	}
	sql := "SELECT * FROM `player_heart_draw_wheel_record`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerHeartDrawWheelRecord), "SELECT MAX(`id`) FROM `player_heart_draw_wheel_record` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iAwardType := res.Map("award_type")
	iAwardNum := res.Map("award_num")
	iItemId := res.Map("item_id")
	iDrawTime := res.Map("draw_time")
	for _, row := range res.Rows {
		crow := C.New_PlayerHeartDrawWheelRecord()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.AwardType = C.int8_t(row.Int8(iAwardType))
		crow.AwardNum = C.int16_t(row.Int16(iAwardNum))
		crow.ItemId = C.int16_t(row.Int16(iItemId))
		crow.DrawTime = C.int64_t(row.Int64(iDrawTime))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerHeartDrawWheelRecord
		playerDb.PlayerHeartDrawWheelRecord = crow
		if g_Hooks.PlayerHeartDrawWheelRecord != nil {
			g_Hooks.PlayerHeartDrawWheelRecord.Load(&PlayerHeartDrawWheelRecordRow{row: crow})
		}
	}
}

func loadPlayerInfo(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_info")
	}
	sql := "SELECT * FROM `player_info`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iIngot := res.Map("ingot")
	iCoins := res.Map("coins")
	iNewMailNum := res.Map("new_mail_num")
	iLastLoginTime := res.Map("last_login_time")
	iLastOfflineTime := res.Map("last_offline_time")
	iTotalOnlineTime := res.Map("total_online_time")
	iFirstLoginTime := res.Map("first_login_time")
	iNewArenaReportNum := res.Map("new_arena_report_num")
	iLastSkillFlush := res.Map("last_skill_flush")
	iSwordSoulFragment := res.Map("sword_soul_fragment")
	for _, row := range res.Rows {
		crow := C.New_PlayerInfo()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Ingot = C.int64_t(row.Int64(iIngot))
		crow.Coins = C.int64_t(row.Int64(iCoins))
		crow.NewMailNum = C.int16_t(row.Int16(iNewMailNum))
		crow.LastLoginTime = C.int64_t(row.Int64(iLastLoginTime))
		crow.LastOfflineTime = C.int64_t(row.Int64(iLastOfflineTime))
		crow.TotalOnlineTime = C.int64_t(row.Int64(iTotalOnlineTime))
		crow.FirstLoginTime = C.int64_t(row.Int64(iFirstLoginTime))
		crow.NewArenaReportNum = C.int16_t(row.Int16(iNewArenaReportNum))
		crow.LastSkillFlush = C.int64_t(row.Int64(iLastSkillFlush))
		crow.SwordSoulFragment = C.int64_t(row.Int64(iSwordSoulFragment))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerInfo = crow
		if g_Hooks.PlayerInfo != nil {
			g_Hooks.PlayerInfo.Load(&PlayerInfoRow{row: crow})
		}
	}
}

func loadPlayerIsOperated(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_is_operated")
	}
	sql := "SELECT * FROM `player_is_operated`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iOperateValue := res.Map("operate_value")
	for _, row := range res.Rows {
		crow := C.New_PlayerIsOperated()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.OperateValue = C.int64_t(row.Int64(iOperateValue))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerIsOperated = crow
		if g_Hooks.PlayerIsOperated != nil {
			g_Hooks.PlayerIsOperated.Load(&PlayerIsOperatedRow{row: crow})
		}
	}
}

func loadPlayerItem(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_item")
	}
	sql := "SELECT * FROM `player_item`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerItem), "SELECT MAX(`id`) FROM `player_item` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iItemId := res.Map("item_id")
	iNum := res.Map("num")
	iIsDressed := res.Map("is_dressed")
	iBuyBackState := res.Map("buy_back_state")
	iAppendixId := res.Map("appendix_id")
	iRefineLevelBak := res.Map("refine_level_bak")
	iRefineFailTimes := res.Map("refine_fail_times")
	iRefineLevel := res.Map("refine_level")
	iPrice := res.Map("price")
	for _, row := range res.Rows {
		crow := C.New_PlayerItem()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.ItemId = C.int16_t(row.Int16(iItemId))
		crow.Num = C.int16_t(row.Int16(iNum))
		crow.IsDressed = C.int8_t(row.Int8(iIsDressed))
		crow.BuyBackState = C.int8_t(row.Int8(iBuyBackState))
		crow.AppendixId = C.int64_t(row.Int64(iAppendixId))
		crow.RefineLevelBak = C.int16_t(row.Int16(iRefineLevelBak))
		crow.RefineFailTimes = C.int16_t(row.Int16(iRefineFailTimes))
		crow.RefineLevel = C.int16_t(row.Int16(iRefineLevel))
		crow.Price = C.int32_t(row.Int32(iPrice))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerItem
		playerDb.PlayerItem = crow
		if g_Hooks.PlayerItem != nil {
			g_Hooks.PlayerItem.Load(&PlayerItemRow{row: crow})
		}
	}
}

func loadPlayerItemAppendix(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_item_appendix")
	}
	sql := "SELECT * FROM `player_item_appendix`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerItemAppendix), "SELECT MAX(`id`) FROM `player_item_appendix` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iHealth := res.Map("health")
	iCultivation := res.Map("cultivation")
	iSpeed := res.Map("speed")
	iAttack := res.Map("attack")
	iDefence := res.Map("defence")
	iDodgeLevel := res.Map("dodge_level")
	iHitLevel := res.Map("hit_level")
	iBlockLevel := res.Map("block_level")
	iCriticalLevel := res.Map("critical_level")
	iTenacityLevel := res.Map("tenacity_level")
	iDestroyLevel := res.Map("destroy_level")
	iRecastAttr := res.Map("recast_attr")
	for _, row := range res.Rows {
		crow := C.New_PlayerItemAppendix()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Health = C.int32_t(row.Int32(iHealth))
		crow.Cultivation = C.int32_t(row.Int32(iCultivation))
		crow.Speed = C.int32_t(row.Int32(iSpeed))
		crow.Attack = C.int32_t(row.Int32(iAttack))
		crow.Defence = C.int32_t(row.Int32(iDefence))
		crow.DodgeLevel = C.int32_t(row.Int32(iDodgeLevel))
		crow.HitLevel = C.int32_t(row.Int32(iHitLevel))
		crow.BlockLevel = C.int32_t(row.Int32(iBlockLevel))
		crow.CriticalLevel = C.int32_t(row.Int32(iCriticalLevel))
		crow.TenacityLevel = C.int32_t(row.Int32(iTenacityLevel))
		crow.DestroyLevel = C.int32_t(row.Int32(iDestroyLevel))
		crow.RecastAttr = C.int8_t(row.Int8(iRecastAttr))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerItemAppendix
		playerDb.PlayerItemAppendix = crow
		if g_Hooks.PlayerItemAppendix != nil {
			g_Hooks.PlayerItemAppendix.Load(&PlayerItemAppendixRow{row: crow})
		}
	}
}

func loadPlayerItemBuyback(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_item_buyback")
	}
	sql := "SELECT * FROM `player_item_buyback`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iBackId1 := res.Map("back_id1")
	iBackId2 := res.Map("back_id2")
	iBackId3 := res.Map("back_id3")
	iBackId4 := res.Map("back_id4")
	iBackId5 := res.Map("back_id5")
	iBackId6 := res.Map("back_id6")
	for _, row := range res.Rows {
		crow := C.New_PlayerItemBuyback()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.BackId1 = C.int64_t(row.Int64(iBackId1))
		crow.BackId2 = C.int64_t(row.Int64(iBackId2))
		crow.BackId3 = C.int64_t(row.Int64(iBackId3))
		crow.BackId4 = C.int64_t(row.Int64(iBackId4))
		crow.BackId5 = C.int64_t(row.Int64(iBackId5))
		crow.BackId6 = C.int64_t(row.Int64(iBackId6))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerItemBuyback = crow
		if g_Hooks.PlayerItemBuyback != nil {
			g_Hooks.PlayerItemBuyback.Load(&PlayerItemBuybackRow{row: crow})
		}
	}
}

func loadPlayerItemRecastState(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_item_recast_state")
	}
	sql := "SELECT * FROM `player_item_recast_state`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iPlayerItemId := res.Map("player_item_id")
	iSelectedAttr := res.Map("selected_attr")
	iAttr1Type := res.Map("attr1_type")
	iAttr1Value := res.Map("attr1_value")
	iAttr2Type := res.Map("attr2_type")
	iAttr2Value := res.Map("attr2_value")
	iAttr3Type := res.Map("attr3_type")
	iAttr3Value := res.Map("attr3_value")
	for _, row := range res.Rows {
		crow := C.New_PlayerItemRecastState()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.PlayerItemId = C.int64_t(row.Int64(iPlayerItemId))
		crow.SelectedAttr = C.int8_t(row.Int8(iSelectedAttr))
		crow.Attr1Type = C.int8_t(row.Int8(iAttr1Type))
		crow.Attr1Value = C.int32_t(row.Int32(iAttr1Value))
		crow.Attr2Type = C.int8_t(row.Int8(iAttr2Type))
		crow.Attr2Value = C.int32_t(row.Int32(iAttr2Value))
		crow.Attr3Type = C.int8_t(row.Int8(iAttr3Type))
		crow.Attr3Value = C.int32_t(row.Int32(iAttr3Value))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerItemRecastState = crow
		if g_Hooks.PlayerItemRecastState != nil {
			g_Hooks.PlayerItemRecastState.Load(&PlayerItemRecastStateRow{row: crow})
		}
	}
}

func loadPlayerLevelAwardInfo(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_level_award_info")
	}
	sql := "SELECT * FROM `player_level_award_info`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iAwarded := res.Map("awarded")
	for _, row := range res.Rows {
		crow := C.New_PlayerLevelAwardInfo()
		crow.Pid = C.int64_t(row.Int64(iPid))
		row_Awarded := row.Str(iAwarded)
		crow.Awarded = C.CString(string(row_Awarded))
		crow.Awarded_len = C.int(len(row_Awarded))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerLevelAwardInfo = crow
		if g_Hooks.PlayerLevelAwardInfo != nil {
			g_Hooks.PlayerLevelAwardInfo.Load(&PlayerLevelAwardInfoRow{row: crow})
		}
	}
}

func loadPlayerLoginAwardRecord(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_login_award_record")
	}
	sql := "SELECT * FROM `player_login_award_record`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iActiveDays := res.Map("active_days")
	iRecord := res.Map("record")
	iUpdateTimestamp := res.Map("update_timestamp")
	for _, row := range res.Rows {
		crow := C.New_PlayerLoginAwardRecord()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.ActiveDays = C.int32_t(row.Int32(iActiveDays))
		crow.Record = C.int32_t(row.Int32(iRecord))
		crow.UpdateTimestamp = C.int64_t(row.Int64(iUpdateTimestamp))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerLoginAwardRecord = crow
		if g_Hooks.PlayerLoginAwardRecord != nil {
			g_Hooks.PlayerLoginAwardRecord.Load(&PlayerLoginAwardRecordRow{row: crow})
		}
	}
}

func loadPlayerMail(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_mail")
	}
	sql := "SELECT * FROM `player_mail`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerMail), "SELECT MAX(`id`) FROM `player_mail` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iMailId := res.Map("mail_id")
	iState := res.Map("state")
	iSendTime := res.Map("send_time")
	iParameters := res.Map("parameters")
	iHaveAttachment := res.Map("have_attachment")
	iTitle := res.Map("title")
	iContent := res.Map("content")
	iExpireTime := res.Map("expire_time")
	iPriority := res.Map("priority")
	for _, row := range res.Rows {
		crow := C.New_PlayerMail()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.MailId = C.int32_t(row.Int32(iMailId))
		crow.State = C.int8_t(row.Int8(iState))
		crow.SendTime = C.int64_t(row.Int64(iSendTime))
		row_Parameters := row.Str(iParameters)
		crow.Parameters = C.CString(string(row_Parameters))
		crow.Parameters_len = C.int(len(row_Parameters))
		crow.HaveAttachment = C.int8_t(row.Int8(iHaveAttachment))
		row_Title := row.Str(iTitle)
		crow.Title = C.CString(string(row_Title))
		crow.Title_len = C.int(len(row_Title))
		row_Content := row.Str(iContent)
		crow.Content = C.CString(string(row_Content))
		crow.Content_len = C.int(len(row_Content))
		crow.ExpireTime = C.int64_t(row.Int64(iExpireTime))
		crow.Priority = C.int8_t(row.Int8(iPriority))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerMail
		playerDb.PlayerMail = crow
		if g_Hooks.PlayerMail != nil {
			g_Hooks.PlayerMail.Load(&PlayerMailRow{row: crow})
		}
	}
}

func loadPlayerMailAttachment(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_mail_attachment")
	}
	sql := "SELECT * FROM `player_mail_attachment`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerMailAttachment), "SELECT MAX(`id`) FROM `player_mail_attachment` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iPlayerMailId := res.Map("player_mail_id")
	iAttachmentType := res.Map("attachment_type")
	iItemId := res.Map("item_id")
	iItemNum := res.Map("item_num")
	for _, row := range res.Rows {
		crow := C.New_PlayerMailAttachment()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.PlayerMailId = C.int64_t(row.Int64(iPlayerMailId))
		crow.AttachmentType = C.int8_t(row.Int8(iAttachmentType))
		crow.ItemId = C.int16_t(row.Int16(iItemId))
		crow.ItemNum = C.int64_t(row.Int64(iItemNum))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerMailAttachment
		playerDb.PlayerMailAttachment = crow
		if g_Hooks.PlayerMailAttachment != nil {
			g_Hooks.PlayerMailAttachment.Load(&PlayerMailAttachmentRow{row: crow})
		}
	}
}

func loadPlayerMailAttachmentLg(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_mail_attachment_lg")
	}
	sql := "SELECT * FROM `player_mail_attachment_lg`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerMailAttachmentLg), "SELECT MAX(`id`) FROM `player_mail_attachment_lg` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iPlayerMailId := res.Map("player_mail_id")
	iAttachmentType := res.Map("attachment_type")
	iItemId := res.Map("item_id")
	iItemNum := res.Map("item_num")
	iTakeTimestamp := res.Map("take_timestamp")
	for _, row := range res.Rows {
		crow := C.New_PlayerMailAttachmentLg()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.PlayerMailId = C.int64_t(row.Int64(iPlayerMailId))
		crow.AttachmentType = C.int8_t(row.Int8(iAttachmentType))
		crow.ItemId = C.int16_t(row.Int16(iItemId))
		crow.ItemNum = C.int64_t(row.Int64(iItemNum))
		crow.TakeTimestamp = C.int64_t(row.Int64(iTakeTimestamp))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerMailAttachmentLg
		playerDb.PlayerMailAttachmentLg = crow
		if g_Hooks.PlayerMailAttachmentLg != nil {
			g_Hooks.PlayerMailAttachmentLg.Load(&PlayerMailAttachmentLgRow{row: crow})
		}
	}
}

func loadPlayerMailLg(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_mail_lg")
	}
	sql := "SELECT * FROM `player_mail_lg`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerMailLg), "SELECT MAX(`id`) FROM `player_mail_lg` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPmid := res.Map("pmid")
	iPid := res.Map("pid")
	iMailId := res.Map("mail_id")
	iState := res.Map("state")
	iSendTime := res.Map("send_time")
	iParameters := res.Map("parameters")
	iHaveAttachment := res.Map("have_attachment")
	iTitle := res.Map("title")
	iContent := res.Map("content")
	for _, row := range res.Rows {
		crow := C.New_PlayerMailLg()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pmid = C.int64_t(row.Int64(iPmid))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.MailId = C.int32_t(row.Int32(iMailId))
		crow.State = C.int8_t(row.Int8(iState))
		crow.SendTime = C.int64_t(row.Int64(iSendTime))
		row_Parameters := row.Str(iParameters)
		crow.Parameters = C.CString(string(row_Parameters))
		crow.Parameters_len = C.int(len(row_Parameters))
		crow.HaveAttachment = C.int8_t(row.Int8(iHaveAttachment))
		row_Title := row.Str(iTitle)
		crow.Title = C.CString(string(row_Title))
		crow.Title_len = C.int(len(row_Title))
		row_Content := row.Str(iContent)
		crow.Content = C.CString(string(row_Content))
		crow.Content_len = C.int(len(row_Content))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerMailLg
		playerDb.PlayerMailLg = crow
		if g_Hooks.PlayerMailLg != nil {
			g_Hooks.PlayerMailLg.Load(&PlayerMailLgRow{row: crow})
		}
	}
}

func loadPlayerMeditationState(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_meditation_state")
	}
	sql := "SELECT * FROM `player_meditation_state`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iAccumulateTime := res.Map("accumulate_time")
	iStartTimestamp := res.Map("start_timestamp")
	for _, row := range res.Rows {
		crow := C.New_PlayerMeditationState()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.AccumulateTime = C.int32_t(row.Int32(iAccumulateTime))
		crow.StartTimestamp = C.int64_t(row.Int64(iStartTimestamp))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerMeditationState = crow
		if g_Hooks.PlayerMeditationState != nil {
			g_Hooks.PlayerMeditationState.Load(&PlayerMeditationStateRow{row: crow})
		}
	}
}

func loadPlayerMission(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_mission")
	}
	sql := "SELECT * FROM `player_mission`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iKey := res.Map("key")
	iMaxOrder := res.Map("max_order")
	for _, row := range res.Rows {
		crow := C.New_PlayerMission()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Key = C.int32_t(row.Int32(iKey))
		crow.MaxOrder = C.int8_t(row.Int8(iMaxOrder))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerMission = crow
		if g_Hooks.PlayerMission != nil {
			g_Hooks.PlayerMission.Load(&PlayerMissionRow{row: crow})
		}
	}
}

func loadPlayerMissionLevel(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_mission_level")
	}
	sql := "SELECT * FROM `player_mission_level`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iLock := res.Map("lock")
	iMaxLock := res.Map("max_lock")
	iAwardLock := res.Map("award_lock")
	for _, row := range res.Rows {
		crow := C.New_PlayerMissionLevel()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Lock = C.int32_t(row.Int32(iLock))
		crow.MaxLock = C.int32_t(row.Int32(iMaxLock))
		crow.AwardLock = C.int32_t(row.Int32(iAwardLock))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerMissionLevel = crow
		if g_Hooks.PlayerMissionLevel != nil {
			g_Hooks.PlayerMissionLevel.Load(&PlayerMissionLevelRow{row: crow})
		}
	}
}

func loadPlayerMissionLevelRecord(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_mission_level_record")
	}
	sql := "SELECT * FROM `player_mission_level_record`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerMissionLevelRecord), "SELECT MAX(`id`) FROM `player_mission_level_record` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iMissionId := res.Map("mission_id")
	iMissionLevelId := res.Map("mission_level_id")
	iOpenTime := res.Map("open_time")
	iScore := res.Map("score")
	iRound := res.Map("round")
	iDailyNum := res.Map("daily_num")
	iLastEnterTime := res.Map("last_enter_time")
	iEmptyShadowBits := res.Map("empty_shadow_bits")
	for _, row := range res.Rows {
		crow := C.New_PlayerMissionLevelRecord()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.MissionId = C.int16_t(row.Int16(iMissionId))
		crow.MissionLevelId = C.int32_t(row.Int32(iMissionLevelId))
		crow.OpenTime = C.int64_t(row.Int64(iOpenTime))
		crow.Score = C.int32_t(row.Int32(iScore))
		crow.Round = C.int8_t(row.Int8(iRound))
		crow.DailyNum = C.int8_t(row.Int8(iDailyNum))
		crow.LastEnterTime = C.int64_t(row.Int64(iLastEnterTime))
		crow.EmptyShadowBits = C.int16_t(row.Int16(iEmptyShadowBits))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerMissionLevelRecord
		playerDb.PlayerMissionLevelRecord = crow
		if g_Hooks.PlayerMissionLevelRecord != nil {
			g_Hooks.PlayerMissionLevelRecord.Load(&PlayerMissionLevelRecordRow{row: crow})
		}
	}
}

func loadPlayerMissionLevelStateBin(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_mission_level_state_bin")
	}
	sql := "SELECT * FROM `player_mission_level_state_bin`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iBin := res.Map("bin")
	for _, row := range res.Rows {
		crow := C.New_PlayerMissionLevelStateBin()
		crow.Pid = C.int64_t(row.Int64(iPid))
		row_Bin := row.Bin(iBin)
		crow.Bin = C.CString(string(row_Bin))
		crow.Bin_len = C.int(len(row_Bin))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerMissionLevelStateBin = crow
		if g_Hooks.PlayerMissionLevelStateBin != nil {
			g_Hooks.PlayerMissionLevelStateBin.Load(&PlayerMissionLevelStateBinRow{row: crow})
		}
	}
}

func loadPlayerMissionRecord(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_mission_record")
	}
	sql := "SELECT * FROM `player_mission_record`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerMissionRecord), "SELECT MAX(`id`) FROM `player_mission_record` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iTownId := res.Map("town_id")
	iMissionId := res.Map("mission_id")
	iOpenTime := res.Map("open_time")
	for _, row := range res.Rows {
		crow := C.New_PlayerMissionRecord()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.TownId = C.int16_t(row.Int16(iTownId))
		crow.MissionId = C.int16_t(row.Int16(iMissionId))
		crow.OpenTime = C.int64_t(row.Int64(iOpenTime))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerMissionRecord
		playerDb.PlayerMissionRecord = crow
		if g_Hooks.PlayerMissionRecord != nil {
			g_Hooks.PlayerMissionRecord.Load(&PlayerMissionRecordRow{row: crow})
		}
	}
}

func loadPlayerMissionStarAward(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_mission_star_award")
	}
	sql := "SELECT * FROM `player_mission_star_award`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerMissionStarAward), "SELECT MAX(`id`) FROM `player_mission_star_award` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iTownId := res.Map("town_id")
	iBoxType := res.Map("box_type")
	for _, row := range res.Rows {
		crow := C.New_PlayerMissionStarAward()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.TownId = C.int16_t(row.Int16(iTownId))
		crow.BoxType = C.int8_t(row.Int8(iBoxType))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerMissionStarAward
		playerDb.PlayerMissionStarAward = crow
		if g_Hooks.PlayerMissionStarAward != nil {
			g_Hooks.PlayerMissionStarAward.Load(&PlayerMissionStarAwardRow{row: crow})
		}
	}
}

func loadPlayerMoneyTree(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_money_tree")
	}
	sql := "SELECT * FROM `player_money_tree`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iTotal := res.Map("total")
	iWavedTimesTotal := res.Map("waved_times_total")
	iWavedTimes := res.Map("waved_times")
	iLastWavedTime := res.Map("last_waved_time")
	for _, row := range res.Rows {
		crow := C.New_PlayerMoneyTree()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Total = C.int32_t(row.Int32(iTotal))
		crow.WavedTimesTotal = C.int8_t(row.Int8(iWavedTimesTotal))
		crow.WavedTimes = C.int8_t(row.Int8(iWavedTimes))
		crow.LastWavedTime = C.int64_t(row.Int64(iLastWavedTime))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerMoneyTree = crow
		if g_Hooks.PlayerMoneyTree != nil {
			g_Hooks.PlayerMoneyTree.Load(&PlayerMoneyTreeRow{row: crow})
		}
	}
}

func loadPlayerMonthCardAwardRecord(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_month_card_award_record")
	}
	sql := "SELECT * FROM `player_month_card_award_record`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iLastUpdate := res.Map("last_update")
	for _, row := range res.Rows {
		crow := C.New_PlayerMonthCardAwardRecord()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.LastUpdate = C.int64_t(row.Int64(iLastUpdate))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerMonthCardAwardRecord = crow
		if g_Hooks.PlayerMonthCardAwardRecord != nil {
			g_Hooks.PlayerMonthCardAwardRecord.Load(&PlayerMonthCardAwardRecordRow{row: crow})
		}
	}
}

func loadPlayerMonthCardInfo(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_month_card_info")
	}
	sql := "SELECT * FROM `player_month_card_info`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iStarttime := res.Map("starttime")
	iEndtime := res.Map("endtime")
	iBuytimes := res.Map("buytimes")
	iPresenttotal := res.Map("presenttotal")
	for _, row := range res.Rows {
		crow := C.New_PlayerMonthCardInfo()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Starttime = C.int64_t(row.Int64(iStarttime))
		crow.Endtime = C.int64_t(row.Int64(iEndtime))
		crow.Buytimes = C.int64_t(row.Int64(iBuytimes))
		crow.Presenttotal = C.int64_t(row.Int64(iPresenttotal))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerMonthCardInfo = crow
		if g_Hooks.PlayerMonthCardInfo != nil {
			g_Hooks.PlayerMonthCardInfo.Load(&PlayerMonthCardInfoRow{row: crow})
		}
	}
}

func loadPlayerMultiLevelInfo(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_multi_level_info")
	}
	sql := "SELECT * FROM `player_multi_level_info`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iBuddyRoleId := res.Map("buddy_role_id")
	iBuddyRow := res.Map("buddy_row")
	iTacticalGrid := res.Map("tactical_grid")
	iDailyNum := res.Map("daily_num")
	iBattleTime := res.Map("battle_time")
	iLock := res.Map("lock")
	for _, row := range res.Rows {
		crow := C.New_PlayerMultiLevelInfo()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.BuddyRoleId = C.int8_t(row.Int8(iBuddyRoleId))
		crow.BuddyRow = C.int8_t(row.Int8(iBuddyRow))
		crow.TacticalGrid = C.int8_t(row.Int8(iTacticalGrid))
		crow.DailyNum = C.int8_t(row.Int8(iDailyNum))
		crow.BattleTime = C.int64_t(row.Int64(iBattleTime))
		crow.Lock = C.int32_t(row.Int32(iLock))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerMultiLevelInfo = crow
		if g_Hooks.PlayerMultiLevelInfo != nil {
			g_Hooks.PlayerMultiLevelInfo.Load(&PlayerMultiLevelInfoRow{row: crow})
		}
	}
}

func loadPlayerNewYearConsumeRecord(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_new_year_consume_record")
	}
	sql := "SELECT * FROM `player_new_year_consume_record`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iConsumeRecord := res.Map("consume_record")
	for _, row := range res.Rows {
		crow := C.New_PlayerNewYearConsumeRecord()
		crow.Pid = C.int64_t(row.Int64(iPid))
		row_ConsumeRecord := row.Str(iConsumeRecord)
		crow.ConsumeRecord = C.CString(string(row_ConsumeRecord))
		crow.ConsumeRecord_len = C.int(len(row_ConsumeRecord))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerNewYearConsumeRecord = crow
		if g_Hooks.PlayerNewYearConsumeRecord != nil {
			g_Hooks.PlayerNewYearConsumeRecord.Load(&PlayerNewYearConsumeRecordRow{row: crow})
		}
	}
}

func loadPlayerNpcTalkRecord(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_npc_talk_record")
	}
	sql := "SELECT * FROM `player_npc_talk_record`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerNpcTalkRecord), "SELECT MAX(`id`) FROM `player_npc_talk_record` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iNpcId := res.Map("npc_id")
	iTownId := res.Map("town_id")
	iQuestId := res.Map("quest_id")
	for _, row := range res.Rows {
		crow := C.New_PlayerNpcTalkRecord()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.NpcId = C.int32_t(row.Int32(iNpcId))
		crow.TownId = C.int16_t(row.Int16(iTownId))
		crow.QuestId = C.int16_t(row.Int16(iQuestId))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerNpcTalkRecord
		playerDb.PlayerNpcTalkRecord = crow
		if g_Hooks.PlayerNpcTalkRecord != nil {
			g_Hooks.PlayerNpcTalkRecord.Load(&PlayerNpcTalkRecordRow{row: crow})
		}
	}
}

func loadPlayerOpenedTownTreasure(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_opened_town_treasure")
	}
	sql := "SELECT * FROM `player_opened_town_treasure`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerOpenedTownTreasure), "SELECT MAX(`id`) FROM `player_opened_town_treasure` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iTownId := res.Map("town_id")
	for _, row := range res.Rows {
		crow := C.New_PlayerOpenedTownTreasure()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.TownId = C.int16_t(row.Int16(iTownId))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerOpenedTownTreasure
		playerDb.PlayerOpenedTownTreasure = crow
		if g_Hooks.PlayerOpenedTownTreasure != nil {
			g_Hooks.PlayerOpenedTownTreasure.Load(&PlayerOpenedTownTreasureRow{row: crow})
		}
	}
}

func loadPlayerPhysical(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_physical")
	}
	sql := "SELECT * FROM `player_physical`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iValue := res.Map("value")
	iUpdateTime := res.Map("update_time")
	iBuyCount := res.Map("buy_count")
	iBuyUpdateTime := res.Map("buy_update_time")
	iDailyCount := res.Map("daily_count")
	for _, row := range res.Rows {
		crow := C.New_PlayerPhysical()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Value = C.int16_t(row.Int16(iValue))
		crow.UpdateTime = C.int64_t(row.Int64(iUpdateTime))
		crow.BuyCount = C.int64_t(row.Int64(iBuyCount))
		crow.BuyUpdateTime = C.int64_t(row.Int64(iBuyUpdateTime))
		crow.DailyCount = C.int16_t(row.Int16(iDailyCount))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerPhysical = crow
		if g_Hooks.PlayerPhysical != nil {
			g_Hooks.PlayerPhysical.Load(&PlayerPhysicalRow{row: crow})
		}
	}
}

func loadPlayerPurchaseRecord(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_purchase_record")
	}
	sql := "SELECT * FROM `player_purchase_record`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerPurchaseRecord), "SELECT MAX(`id`) FROM `player_purchase_record` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iItemId := res.Map("item_id")
	iNum := res.Map("num")
	iTimestamp := res.Map("timestamp")
	for _, row := range res.Rows {
		crow := C.New_PlayerPurchaseRecord()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.ItemId = C.int16_t(row.Int16(iItemId))
		crow.Num = C.int16_t(row.Int16(iNum))
		crow.Timestamp = C.int64_t(row.Int64(iTimestamp))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerPurchaseRecord
		playerDb.PlayerPurchaseRecord = crow
		if g_Hooks.PlayerPurchaseRecord != nil {
			g_Hooks.PlayerPurchaseRecord.Load(&PlayerPurchaseRecordRow{row: crow})
		}
	}
}

func loadPlayerPushNotifySwitch(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_push_notify_switch")
	}
	sql := "SELECT * FROM `player_push_notify_switch`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iOptions := res.Map("options")
	for _, row := range res.Rows {
		crow := C.New_PlayerPushNotifySwitch()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Options = C.int64_t(row.Int64(iOptions))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerPushNotifySwitch = crow
		if g_Hooks.PlayerPushNotifySwitch != nil {
			g_Hooks.PlayerPushNotifySwitch.Load(&PlayerPushNotifySwitchRow{row: crow})
		}
	}
}

func loadPlayerPveState(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_pve_state")
	}
	sql := "SELECT * FROM `player_pve_state`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iMaxPassedFloor := res.Map("max_passed_floor")
	iMaxAwardedFloor := res.Map("max_awarded_floor")
	iUnpassedFloorEnemyNum := res.Map("unpassed_floor_enemy_num")
	iEnterTime := res.Map("enter_time")
	iDailyNum := res.Map("daily_num")
	for _, row := range res.Rows {
		crow := C.New_PlayerPveState()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.MaxPassedFloor = C.int16_t(row.Int16(iMaxPassedFloor))
		crow.MaxAwardedFloor = C.int16_t(row.Int16(iMaxAwardedFloor))
		crow.UnpassedFloorEnemyNum = C.int16_t(row.Int16(iUnpassedFloorEnemyNum))
		crow.EnterTime = C.int64_t(row.Int64(iEnterTime))
		crow.DailyNum = C.int8_t(row.Int8(iDailyNum))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerPveState = crow
		if g_Hooks.PlayerPveState != nil {
			g_Hooks.PlayerPveState.Load(&PlayerPveStateRow{row: crow})
		}
	}
}

func loadPlayerQqVipGift(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_qq_vip_gift")
	}
	sql := "SELECT * FROM `player_qq_vip_gift`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iQqvip := res.Map("qqvip")
	iSurper := res.Map("surper")
	for _, row := range res.Rows {
		crow := C.New_PlayerQqVipGift()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Qqvip = C.int16_t(row.Int16(iQqvip))
		crow.Surper = C.int16_t(row.Int16(iSurper))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerQqVipGift = crow
		if g_Hooks.PlayerQqVipGift != nil {
			g_Hooks.PlayerQqVipGift.Load(&PlayerQqVipGiftRow{row: crow})
		}
	}
}

func loadPlayerQuest(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_quest")
	}
	sql := "SELECT * FROM `player_quest`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iQuestId := res.Map("quest_id")
	iState := res.Map("state")
	for _, row := range res.Rows {
		crow := C.New_PlayerQuest()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.QuestId = C.int16_t(row.Int16(iQuestId))
		crow.State = C.int8_t(row.Int8(iState))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerQuest = crow
		if g_Hooks.PlayerQuest != nil {
			g_Hooks.PlayerQuest.Load(&PlayerQuestRow{row: crow})
		}
	}
}

func loadPlayerRainbowLevel(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_rainbow_level")
	}
	sql := "SELECT * FROM `player_rainbow_level`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iResetTimestamp := res.Map("reset_timestamp")
	iResetNum := res.Map("reset_num")
	iSegment := res.Map("segment")
	iOrder := res.Map("order")
	iStatus := res.Map("status")
	iBestSegment := res.Map("best_segment")
	iBestOrder := res.Map("best_order")
	iBestRecordTimestamp := res.Map("best_record_timestamp")
	iMaxOpenSegment := res.Map("max_open_segment")
	iMaxPassSegment := res.Map("max_pass_segment")
	iAutoFightNum := res.Map("auto_fight_num")
	iBuyTimes := res.Map("buy_times")
	iAutoFightTime := res.Map("auto_fight_time")
	iBuyTimestamp := res.Map("buy_timestamp")
	for _, row := range res.Rows {
		crow := C.New_PlayerRainbowLevel()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.ResetTimestamp = C.int64_t(row.Int64(iResetTimestamp))
		crow.ResetNum = C.int32_t(row.Int32(iResetNum))
		crow.Segment = C.int16_t(row.Int16(iSegment))
		crow.Order = C.int8_t(row.Int8(iOrder))
		crow.Status = C.int8_t(row.Int8(iStatus))
		crow.BestSegment = C.int16_t(row.Int16(iBestSegment))
		crow.BestOrder = C.int8_t(row.Int8(iBestOrder))
		crow.BestRecordTimestamp = C.int64_t(row.Int64(iBestRecordTimestamp))
		crow.MaxOpenSegment = C.int16_t(row.Int16(iMaxOpenSegment))
		crow.MaxPassSegment = C.int16_t(row.Int16(iMaxPassSegment))
		crow.AutoFightNum = C.int8_t(row.Int8(iAutoFightNum))
		crow.BuyTimes = C.int16_t(row.Int16(iBuyTimes))
		crow.AutoFightTime = C.int64_t(row.Int64(iAutoFightTime))
		crow.BuyTimestamp = C.int64_t(row.Int64(iBuyTimestamp))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerRainbowLevel = crow
		if g_Hooks.PlayerRainbowLevel != nil {
			g_Hooks.PlayerRainbowLevel.Load(&PlayerRainbowLevelRow{row: crow})
		}
	}
}

func loadPlayerRainbowLevelStateBin(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_rainbow_level_state_bin")
	}
	sql := "SELECT * FROM `player_rainbow_level_state_bin`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iBin := res.Map("bin")
	for _, row := range res.Rows {
		crow := C.New_PlayerRainbowLevelStateBin()
		crow.Pid = C.int64_t(row.Int64(iPid))
		row_Bin := row.Bin(iBin)
		crow.Bin = C.CString(string(row_Bin))
		crow.Bin_len = C.int(len(row_Bin))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerRainbowLevelStateBin = crow
		if g_Hooks.PlayerRainbowLevelStateBin != nil {
			g_Hooks.PlayerRainbowLevelStateBin.Load(&PlayerRainbowLevelStateBinRow{row: crow})
		}
	}
}

func loadPlayerRole(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_role")
	}
	sql := "SELECT * FROM `player_role`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerRole), "SELECT MAX(`id`) FROM `player_role` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iRoleId := res.Map("role_id")
	iLevel := res.Map("level")
	iExp := res.Map("exp")
	iFriendshipLevel := res.Map("friendship_level")
	iStatus := res.Map("status")
	iCoopId := res.Map("coop_id")
	iTimestamp := res.Map("timestamp")
	iFightNum := res.Map("fight_num")
	for _, row := range res.Rows {
		crow := C.New_PlayerRole()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.RoleId = C.int8_t(row.Int8(iRoleId))
		crow.Level = C.int16_t(row.Int16(iLevel))
		crow.Exp = C.int64_t(row.Int64(iExp))
		crow.FriendshipLevel = C.int32_t(row.Int32(iFriendshipLevel))
		crow.Status = C.int16_t(row.Int16(iStatus))
		crow.CoopId = C.int16_t(row.Int16(iCoopId))
		crow.Timestamp = C.int64_t(row.Int64(iTimestamp))
		crow.FightNum = C.int32_t(row.Int32(iFightNum))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerRole
		playerDb.PlayerRole = crow
		if g_Hooks.PlayerRole != nil {
			g_Hooks.PlayerRole.Load(&PlayerRoleRow{row: crow})
		}
	}
}

func loadPlayerSealedbook(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_sealedbook")
	}
	sql := "SELECT * FROM `player_sealedbook`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iSealedRecord := res.Map("sealed_record")
	for _, row := range res.Rows {
		crow := C.New_PlayerSealedbook()
		crow.Pid = C.int64_t(row.Int64(iPid))
		row_SealedRecord := row.Bin(iSealedRecord)
		crow.SealedRecord = C.CString(string(row_SealedRecord))
		crow.SealedRecord_len = C.int(len(row_SealedRecord))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerSealedbook = crow
		if g_Hooks.PlayerSealedbook != nil {
			g_Hooks.PlayerSealedbook.Load(&PlayerSealedbookRow{row: crow})
		}
	}
}

func loadPlayerSendHeartRecord(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_send_heart_record")
	}
	sql := "SELECT * FROM `player_send_heart_record`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerSendHeartRecord), "SELECT MAX(`id`) FROM `player_send_heart_record` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iFriendPid := res.Map("friend_pid")
	iSendHeartTime := res.Map("send_heart_time")
	for _, row := range res.Rows {
		crow := C.New_PlayerSendHeartRecord()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.FriendPid = C.int64_t(row.Int64(iFriendPid))
		crow.SendHeartTime = C.int64_t(row.Int64(iSendHeartTime))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerSendHeartRecord
		playerDb.PlayerSendHeartRecord = crow
		if g_Hooks.PlayerSendHeartRecord != nil {
			g_Hooks.PlayerSendHeartRecord.Load(&PlayerSendHeartRecordRow{row: crow})
		}
	}
}

func loadPlayerSkill(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_skill")
	}
	sql := "SELECT * FROM `player_skill`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerSkill), "SELECT MAX(`id`) FROM `player_skill` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iRoleId := res.Map("role_id")
	iSkillId := res.Map("skill_id")
	iSkillTrnlv := res.Map("skill_trnlv")
	for _, row := range res.Rows {
		crow := C.New_PlayerSkill()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.RoleId = C.int8_t(row.Int8(iRoleId))
		crow.SkillId = C.int16_t(row.Int16(iSkillId))
		crow.SkillTrnlv = C.int32_t(row.Int32(iSkillTrnlv))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerSkill
		playerDb.PlayerSkill = crow
		if g_Hooks.PlayerSkill != nil {
			g_Hooks.PlayerSkill.Load(&PlayerSkillRow{row: crow})
		}
	}
}

func loadPlayerState(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_state")
	}
	sql := "SELECT * FROM `player_state`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iBanStartTime := res.Map("ban_start_time")
	iBanEndTime := res.Map("ban_end_time")
	for _, row := range res.Rows {
		crow := C.New_PlayerState()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.BanStartTime = C.int64_t(row.Int64(iBanStartTime))
		crow.BanEndTime = C.int64_t(row.Int64(iBanEndTime))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerState = crow
		if g_Hooks.PlayerState != nil {
			g_Hooks.PlayerState.Load(&PlayerStateRow{row: crow})
		}
	}
}

func loadPlayerSwordSoul(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_sword_soul")
	}
	sql := "SELECT * FROM `player_sword_soul`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerSwordSoul), "SELECT MAX(`id`) FROM `player_sword_soul` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iPos := res.Map("pos")
	iSwordSoulId := res.Map("sword_soul_id")
	iExp := res.Map("exp")
	iLevel := res.Map("level")
	for _, row := range res.Rows {
		crow := C.New_PlayerSwordSoul()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Pos = C.int8_t(row.Int8(iPos))
		crow.SwordSoulId = C.int16_t(row.Int16(iSwordSoulId))
		crow.Exp = C.int32_t(row.Int32(iExp))
		crow.Level = C.int8_t(row.Int8(iLevel))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerSwordSoul
		playerDb.PlayerSwordSoul = crow
		if g_Hooks.PlayerSwordSoul != nil {
			g_Hooks.PlayerSwordSoul.Load(&PlayerSwordSoulRow{row: crow})
		}
	}
}

func loadPlayerSwordSoulEquipment(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_sword_soul_equipment")
	}
	sql := "SELECT * FROM `player_sword_soul_equipment`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerSwordSoulEquipment), "SELECT MAX(`id`) FROM `player_sword_soul_equipment` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iRoleId := res.Map("role_id")
	iIsEquippedXuanyuan := res.Map("is_equipped_xuanyuan")
	iTypeBits := res.Map("type_bits")
	iPos1 := res.Map("pos1")
	iPos2 := res.Map("pos2")
	iPos3 := res.Map("pos3")
	iPos4 := res.Map("pos4")
	iPos5 := res.Map("pos5")
	iPos6 := res.Map("pos6")
	iPos7 := res.Map("pos7")
	iPos8 := res.Map("pos8")
	iPos9 := res.Map("pos9")
	for _, row := range res.Rows {
		crow := C.New_PlayerSwordSoulEquipment()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.RoleId = C.int8_t(row.Int8(iRoleId))
		crow.IsEquippedXuanyuan = C.int8_t(row.Int8(iIsEquippedXuanyuan))
		crow.TypeBits = C.int64_t(row.Int64(iTypeBits))
		crow.Pos1 = C.int64_t(row.Int64(iPos1))
		crow.Pos2 = C.int64_t(row.Int64(iPos2))
		crow.Pos3 = C.int64_t(row.Int64(iPos3))
		crow.Pos4 = C.int64_t(row.Int64(iPos4))
		crow.Pos5 = C.int64_t(row.Int64(iPos5))
		crow.Pos6 = C.int64_t(row.Int64(iPos6))
		crow.Pos7 = C.int64_t(row.Int64(iPos7))
		crow.Pos8 = C.int64_t(row.Int64(iPos8))
		crow.Pos9 = C.int64_t(row.Int64(iPos9))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerSwordSoulEquipment
		playerDb.PlayerSwordSoulEquipment = crow
		if g_Hooks.PlayerSwordSoulEquipment != nil {
			g_Hooks.PlayerSwordSoulEquipment.Load(&PlayerSwordSoulEquipmentRow{row: crow})
		}
	}
}

func loadPlayerSwordSoulState(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_sword_soul_state")
	}
	sql := "SELECT * FROM `player_sword_soul_state`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iBoxState := res.Map("box_state")
	iNum := res.Map("num")
	iUpdateTime := res.Map("update_time")
	iIngotNum := res.Map("ingot_num")
	iSupersoulAdditionalPossible := res.Map("supersoul_additional_possible")
	iIngotYelloOne := res.Map("ingot_yello_one")
	iLastIngotDrawTime := res.Map("last_ingot_draw_time")
	iSwordSoulNum := res.Map("sword_soul_num")
	for _, row := range res.Rows {
		crow := C.New_PlayerSwordSoulState()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.BoxState = C.int8_t(row.Int8(iBoxState))
		crow.Num = C.int16_t(row.Int16(iNum))
		crow.UpdateTime = C.int64_t(row.Int64(iUpdateTime))
		crow.IngotNum = C.int64_t(row.Int64(iIngotNum))
		crow.SupersoulAdditionalPossible = C.int8_t(row.Int8(iSupersoulAdditionalPossible))
		crow.IngotYelloOne = C.int16_t(row.Int16(iIngotYelloOne))
		crow.LastIngotDrawTime = C.int64_t(row.Int64(iLastIngotDrawTime))
		crow.SwordSoulNum = C.int32_t(row.Int32(iSwordSoulNum))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerSwordSoulState = crow
		if g_Hooks.PlayerSwordSoulState != nil {
			g_Hooks.PlayerSwordSoulState.Load(&PlayerSwordSoulStateRow{row: crow})
		}
	}
}

func loadPlayerTaoyuanBless(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_taoyuan_bless")
	}
	sql := "SELECT * FROM `player_taoyuan_bless`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iDailyBlessTimes := res.Map("daily_bless_times")
	iLastBlessTime := res.Map("last_bless_time")
	iDailyBeBlessTimes := res.Map("daily_be_bless_times")
	iLastBeBlessTime := res.Map("last_be_bless_time")
	iBlessPid1 := res.Map("bless_pid1")
	iBlessPid2 := res.Map("bless_pid2")
	iBlessPid3 := res.Map("bless_pid3")
	iBlessPid4 := res.Map("bless_pid4")
	iBlessPid5 := res.Map("bless_pid5")
	for _, row := range res.Rows {
		crow := C.New_PlayerTaoyuanBless()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.DailyBlessTimes = C.int32_t(row.Int32(iDailyBlessTimes))
		crow.LastBlessTime = C.int64_t(row.Int64(iLastBlessTime))
		crow.DailyBeBlessTimes = C.int32_t(row.Int32(iDailyBeBlessTimes))
		crow.LastBeBlessTime = C.int64_t(row.Int64(iLastBeBlessTime))
		crow.BlessPid1 = C.int64_t(row.Int64(iBlessPid1))
		crow.BlessPid2 = C.int64_t(row.Int64(iBlessPid2))
		crow.BlessPid3 = C.int64_t(row.Int64(iBlessPid3))
		crow.BlessPid4 = C.int64_t(row.Int64(iBlessPid4))
		crow.BlessPid5 = C.int64_t(row.Int64(iBlessPid5))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerTaoyuanBless = crow
		if g_Hooks.PlayerTaoyuanBless != nil {
			g_Hooks.PlayerTaoyuanBless.Load(&PlayerTaoyuanBlessRow{row: crow})
		}
	}
}

func loadPlayerTaoyuanFileds(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_taoyuan_fileds")
	}
	sql := "SELECT * FROM `player_taoyuan_fileds`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerTaoyuanFileds), "SELECT MAX(`id`) FROM `player_taoyuan_fileds` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iFiledId := res.Map("filed_id")
	iFiledStatus := res.Map("filed_status")
	iPlantId := res.Map("plant_id")
	iGrowTime := res.Map("grow_time")
	for _, row := range res.Rows {
		crow := C.New_PlayerTaoyuanFileds()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.FiledId = C.int16_t(row.Int16(iFiledId))
		crow.FiledStatus = C.int16_t(row.Int16(iFiledStatus))
		crow.PlantId = C.int16_t(row.Int16(iPlantId))
		crow.GrowTime = C.int64_t(row.Int64(iGrowTime))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerTaoyuanFileds
		playerDb.PlayerTaoyuanFileds = crow
		if g_Hooks.PlayerTaoyuanFileds != nil {
			g_Hooks.PlayerTaoyuanFileds.Load(&PlayerTaoyuanFiledsRow{row: crow})
		}
	}
}

func loadPlayerTaoyuanHeart(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_taoyuan_heart")
	}
	sql := "SELECT * FROM `player_taoyuan_heart`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iLevel := res.Map("level")
	iExp := res.Map("exp")
	for _, row := range res.Rows {
		crow := C.New_PlayerTaoyuanHeart()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Level = C.int16_t(row.Int16(iLevel))
		crow.Exp = C.int64_t(row.Int64(iExp))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerTaoyuanHeart = crow
		if g_Hooks.PlayerTaoyuanHeart != nil {
			g_Hooks.PlayerTaoyuanHeart.Load(&PlayerTaoyuanHeartRow{row: crow})
		}
	}
}

func loadPlayerTaoyuanItem(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_taoyuan_item")
	}
	sql := "SELECT * FROM `player_taoyuan_item`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerTaoyuanItem), "SELECT MAX(`id`) FROM `player_taoyuan_item` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iItemId := res.Map("item_id")
	iNum := res.Map("num")
	for _, row := range res.Rows {
		crow := C.New_PlayerTaoyuanItem()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.ItemId = C.int16_t(row.Int16(iItemId))
		crow.Num = C.int16_t(row.Int16(iNum))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerTaoyuanItem
		playerDb.PlayerTaoyuanItem = crow
		if g_Hooks.PlayerTaoyuanItem != nil {
			g_Hooks.PlayerTaoyuanItem.Load(&PlayerTaoyuanItemRow{row: crow})
		}
	}
}

func loadPlayerTaoyuanMessage(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_taoyuan_message")
	}
	sql := "SELECT * FROM `player_taoyuan_message`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerTaoyuanMessage), "SELECT MAX(`id`) FROM `player_taoyuan_message` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iNick := res.Map("nick")
	iExp := res.Map("exp")
	for _, row := range res.Rows {
		crow := C.New_PlayerTaoyuanMessage()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		row_Nick := row.Str(iNick)
		crow.Nick = C.CString(string(row_Nick))
		crow.Nick_len = C.int(len(row_Nick))
		crow.Exp = C.int32_t(row.Int32(iExp))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerTaoyuanMessage
		playerDb.PlayerTaoyuanMessage = crow
		if g_Hooks.PlayerTaoyuanMessage != nil {
			g_Hooks.PlayerTaoyuanMessage.Load(&PlayerTaoyuanMessageRow{row: crow})
		}
	}
}

func loadPlayerTaoyuanProduct(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_taoyuan_product")
	}
	sql := "SELECT * FROM `player_taoyuan_product`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iSkillId1 := res.Map("skill_id1")
	iSkillId2 := res.Map("skill_id2")
	iSameTimeWineQueue := res.Map("same_time_wine_queue")
	iMakeWineTimes := res.Map("make_wine_times")
	iSameTimeFoodQueue := res.Map("same_time_food_queue")
	iMakeFoodTimes := res.Map("make_food_times")
	iFoodQueue1 := res.Map("food_queue1")
	iFoodQueue1StartTimestamp := res.Map("food_queue1_start_timestamp")
	iFoodQueue1EndTimestamp := res.Map("food_queue1_end_timestamp")
	iFoodQueue2 := res.Map("food_queue2")
	iFoodQueue2StartTimestamp := res.Map("food_queue2_start_timestamp")
	iFoodQueue2EndTimestamp := res.Map("food_queue2_end_timestamp")
	iFoodQueue3 := res.Map("food_queue3")
	iFoodQueue3StartTimestamp := res.Map("food_queue3_start_timestamp")
	iFoodQueue3EndTimestamp := res.Map("food_queue3_end_timestamp")
	iFoodQueue4 := res.Map("food_queue4")
	iFoodQueue4StartTimestamp := res.Map("food_queue4_start_timestamp")
	iFoodQueue4EndTimestamp := res.Map("food_queue4_end_timestamp")
	iWineQueue1 := res.Map("wine_queue1")
	iWineQueue1StartTimestamp := res.Map("wine_queue1_start_timestamp")
	iWineQueue1EndTimestamp := res.Map("wine_queue1_end_timestamp")
	iWineQueue2 := res.Map("wine_queue2")
	iWineQueue2StartTimestamp := res.Map("wine_queue2_start_timestamp")
	iWineQueue2EndTimestamp := res.Map("wine_queue2_end_timestamp")
	iWineQueue3 := res.Map("wine_queue3")
	iWineQueue3StartTimestamp := res.Map("wine_queue3_start_timestamp")
	iWineQueue3EndTimestamp := res.Map("wine_queue3_end_timestamp")
	iWineQueue4 := res.Map("wine_queue4")
	iWineQueue4StartTimestamp := res.Map("wine_queue4_start_timestamp")
	iWineQueue4EndTimestamp := res.Map("wine_queue4_end_timestamp")
	iIsOpenWine := res.Map("is_open_wine")
	iIsOpenFood := res.Map("is_open_food")
	for _, row := range res.Rows {
		crow := C.New_PlayerTaoyuanProduct()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.SkillId1 = C.int16_t(row.Int16(iSkillId1))
		crow.SkillId2 = C.int16_t(row.Int16(iSkillId2))
		crow.SameTimeWineQueue = C.int16_t(row.Int16(iSameTimeWineQueue))
		crow.MakeWineTimes = C.int64_t(row.Int64(iMakeWineTimes))
		crow.SameTimeFoodQueue = C.int16_t(row.Int16(iSameTimeFoodQueue))
		crow.MakeFoodTimes = C.int64_t(row.Int64(iMakeFoodTimes))
		crow.FoodQueue1 = C.int16_t(row.Int16(iFoodQueue1))
		crow.FoodQueue1StartTimestamp = C.int64_t(row.Int64(iFoodQueue1StartTimestamp))
		crow.FoodQueue1EndTimestamp = C.int64_t(row.Int64(iFoodQueue1EndTimestamp))
		crow.FoodQueue2 = C.int16_t(row.Int16(iFoodQueue2))
		crow.FoodQueue2StartTimestamp = C.int64_t(row.Int64(iFoodQueue2StartTimestamp))
		crow.FoodQueue2EndTimestamp = C.int64_t(row.Int64(iFoodQueue2EndTimestamp))
		crow.FoodQueue3 = C.int16_t(row.Int16(iFoodQueue3))
		crow.FoodQueue3StartTimestamp = C.int64_t(row.Int64(iFoodQueue3StartTimestamp))
		crow.FoodQueue3EndTimestamp = C.int64_t(row.Int64(iFoodQueue3EndTimestamp))
		crow.FoodQueue4 = C.int16_t(row.Int16(iFoodQueue4))
		crow.FoodQueue4StartTimestamp = C.int64_t(row.Int64(iFoodQueue4StartTimestamp))
		crow.FoodQueue4EndTimestamp = C.int64_t(row.Int64(iFoodQueue4EndTimestamp))
		crow.WineQueue1 = C.int16_t(row.Int16(iWineQueue1))
		crow.WineQueue1StartTimestamp = C.int64_t(row.Int64(iWineQueue1StartTimestamp))
		crow.WineQueue1EndTimestamp = C.int64_t(row.Int64(iWineQueue1EndTimestamp))
		crow.WineQueue2 = C.int16_t(row.Int16(iWineQueue2))
		crow.WineQueue2StartTimestamp = C.int64_t(row.Int64(iWineQueue2StartTimestamp))
		crow.WineQueue2EndTimestamp = C.int64_t(row.Int64(iWineQueue2EndTimestamp))
		crow.WineQueue3 = C.int16_t(row.Int16(iWineQueue3))
		crow.WineQueue3StartTimestamp = C.int64_t(row.Int64(iWineQueue3StartTimestamp))
		crow.WineQueue3EndTimestamp = C.int64_t(row.Int64(iWineQueue3EndTimestamp))
		crow.WineQueue4 = C.int16_t(row.Int16(iWineQueue4))
		crow.WineQueue4StartTimestamp = C.int64_t(row.Int64(iWineQueue4StartTimestamp))
		crow.WineQueue4EndTimestamp = C.int64_t(row.Int64(iWineQueue4EndTimestamp))
		crow.IsOpenWine = C.int8_t(row.Int8(iIsOpenWine))
		crow.IsOpenFood = C.int8_t(row.Int8(iIsOpenFood))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerTaoyuanProduct = crow
		if g_Hooks.PlayerTaoyuanProduct != nil {
			g_Hooks.PlayerTaoyuanProduct.Load(&PlayerTaoyuanProductRow{row: crow})
		}
	}
}

func loadPlayerTaoyuanPurchaseRecord(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_taoyuan_purchase_record")
	}
	sql := "SELECT * FROM `player_taoyuan_purchase_record`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerTaoyuanPurchaseRecord), "SELECT MAX(`id`) FROM `player_taoyuan_purchase_record` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iItemId := res.Map("item_id")
	iNum := res.Map("num")
	iTimestamp := res.Map("timestamp")
	for _, row := range res.Rows {
		crow := C.New_PlayerTaoyuanPurchaseRecord()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.ItemId = C.int16_t(row.Int16(iItemId))
		crow.Num = C.int16_t(row.Int16(iNum))
		crow.Timestamp = C.int64_t(row.Int64(iTimestamp))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerTaoyuanPurchaseRecord
		playerDb.PlayerTaoyuanPurchaseRecord = crow
		if g_Hooks.PlayerTaoyuanPurchaseRecord != nil {
			g_Hooks.PlayerTaoyuanPurchaseRecord.Load(&PlayerTaoyuanPurchaseRecordRow{row: crow})
		}
	}
}

func loadPlayerTaoyuanQuest(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_taoyuan_quest")
	}
	sql := "SELECT * FROM `player_taoyuan_quest`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iLastQuestFlashTime := res.Map("last_quest_flash_time")
	iQuest1ItemId := res.Map("quest1_item_id")
	iQuest1ItemNum := res.Map("quest1_item_num")
	iQuest1TotalExp := res.Map("quest1_total_exp")
	iQuest1TotalCoins := res.Map("quest1_total_coins")
	iQuest1FinishTime := res.Map("quest1_finish_time")
	iQuest2ItemId := res.Map("quest2_item_id")
	iQuest2ItemNum := res.Map("quest2_item_num")
	iQuest2TotalExp := res.Map("quest2_total_exp")
	iQuest2TotalCoins := res.Map("quest2_total_coins")
	iQuest2FinishTime := res.Map("quest2_finish_time")
	iQuest3ItemId := res.Map("quest3_item_id")
	iQuest3ItemNum := res.Map("quest3_item_num")
	iQuest3TotalExp := res.Map("quest3_total_exp")
	iQuest3TotalCoins := res.Map("quest3_total_coins")
	iQuest3FinishTime := res.Map("quest3_finish_time")
	iQuest4ItemId := res.Map("quest4_item_id")
	iQuest4ItemNum := res.Map("quest4_item_num")
	iQuest4TotalExp := res.Map("quest4_total_exp")
	iQuest4TotalCoins := res.Map("quest4_total_coins")
	iQuest4FinishTime := res.Map("quest4_finish_time")
	iQuest5ItemId := res.Map("quest5_item_id")
	iQuest5ItemNum := res.Map("quest5_item_num")
	iQuest5TotalExp := res.Map("quest5_total_exp")
	iQuest5TotalCoins := res.Map("quest5_total_coins")
	iQuest5FinishTime := res.Map("quest5_finish_time")
	for _, row := range res.Rows {
		crow := C.New_PlayerTaoyuanQuest()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.LastQuestFlashTime = C.int64_t(row.Int64(iLastQuestFlashTime))
		crow.Quest1ItemId = C.int16_t(row.Int16(iQuest1ItemId))
		crow.Quest1ItemNum = C.int16_t(row.Int16(iQuest1ItemNum))
		crow.Quest1TotalExp = C.int16_t(row.Int16(iQuest1TotalExp))
		crow.Quest1TotalCoins = C.int64_t(row.Int64(iQuest1TotalCoins))
		crow.Quest1FinishTime = C.int64_t(row.Int64(iQuest1FinishTime))
		crow.Quest2ItemId = C.int16_t(row.Int16(iQuest2ItemId))
		crow.Quest2ItemNum = C.int16_t(row.Int16(iQuest2ItemNum))
		crow.Quest2TotalExp = C.int16_t(row.Int16(iQuest2TotalExp))
		crow.Quest2TotalCoins = C.int64_t(row.Int64(iQuest2TotalCoins))
		crow.Quest2FinishTime = C.int64_t(row.Int64(iQuest2FinishTime))
		crow.Quest3ItemId = C.int16_t(row.Int16(iQuest3ItemId))
		crow.Quest3ItemNum = C.int16_t(row.Int16(iQuest3ItemNum))
		crow.Quest3TotalExp = C.int16_t(row.Int16(iQuest3TotalExp))
		crow.Quest3TotalCoins = C.int64_t(row.Int64(iQuest3TotalCoins))
		crow.Quest3FinishTime = C.int64_t(row.Int64(iQuest3FinishTime))
		crow.Quest4ItemId = C.int16_t(row.Int16(iQuest4ItemId))
		crow.Quest4ItemNum = C.int16_t(row.Int16(iQuest4ItemNum))
		crow.Quest4TotalExp = C.int16_t(row.Int16(iQuest4TotalExp))
		crow.Quest4TotalCoins = C.int64_t(row.Int64(iQuest4TotalCoins))
		crow.Quest4FinishTime = C.int64_t(row.Int64(iQuest4FinishTime))
		crow.Quest5ItemId = C.int16_t(row.Int16(iQuest5ItemId))
		crow.Quest5ItemNum = C.int16_t(row.Int16(iQuest5ItemNum))
		crow.Quest5TotalExp = C.int16_t(row.Int16(iQuest5TotalExp))
		crow.Quest5TotalCoins = C.int64_t(row.Int64(iQuest5TotalCoins))
		crow.Quest5FinishTime = C.int64_t(row.Int64(iQuest5FinishTime))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerTaoyuanQuest = crow
		if g_Hooks.PlayerTaoyuanQuest != nil {
			g_Hooks.PlayerTaoyuanQuest.Load(&PlayerTaoyuanQuestRow{row: crow})
		}
	}
}

func loadPlayerTbXxdRoleinfo(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_tb_xxd_roleinfo")
	}
	sql := "SELECT * FROM `player_tb_xxd_roleinfo`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iGameappid := res.Map("gameappid")
	iOpenid := res.Map("openid")
	iRegtime := res.Map("regtime")
	iLevel := res.Map("level")
	iIFriends := res.Map("iFriends")
	iMoneyios := res.Map("moneyios")
	iMoneyandroid := res.Map("moneyandroid")
	iDiamondios := res.Map("diamondios")
	iDiamondandroid := res.Map("diamondandroid")
	for _, row := range res.Rows {
		crow := C.New_PlayerTbXxdRoleinfo()
		crow.Pid = C.int64_t(row.Int64(iPid))
		row_Gameappid := row.Str(iGameappid)
		crow.Gameappid = C.CString(string(row_Gameappid))
		crow.Gameappid_len = C.int(len(row_Gameappid))
		row_Openid := row.Str(iOpenid)
		crow.Openid = C.CString(string(row_Openid))
		crow.Openid_len = C.int(len(row_Openid))
		crow.Regtime = C.int64_t(row.Int64(iRegtime))
		crow.Level = C.int16_t(row.Int16(iLevel))
		crow.IFriends = C.int16_t(row.Int16(iIFriends))
		crow.Moneyios = C.int64_t(row.Int64(iMoneyios))
		crow.Moneyandroid = C.int64_t(row.Int64(iMoneyandroid))
		crow.Diamondios = C.int64_t(row.Int64(iDiamondios))
		crow.Diamondandroid = C.int64_t(row.Int64(iDiamondandroid))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerTbXxdRoleinfo = crow
		if g_Hooks.PlayerTbXxdRoleinfo != nil {
			g_Hooks.PlayerTbXxdRoleinfo.Load(&PlayerTbXxdRoleinfoRow{row: crow})
		}
	}
}

func loadPlayerTeamInfo(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_team_info")
	}
	sql := "SELECT * FROM `player_team_info`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iRelationship := res.Map("relationship")
	iHealthLv := res.Map("health_lv")
	iAttackLv := res.Map("attack_lv")
	iDefenceLv := res.Map("defence_lv")
	for _, row := range res.Rows {
		crow := C.New_PlayerTeamInfo()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Relationship = C.int32_t(row.Int32(iRelationship))
		crow.HealthLv = C.int16_t(row.Int16(iHealthLv))
		crow.AttackLv = C.int16_t(row.Int16(iAttackLv))
		crow.DefenceLv = C.int16_t(row.Int16(iDefenceLv))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerTeamInfo = crow
		if g_Hooks.PlayerTeamInfo != nil {
			g_Hooks.PlayerTeamInfo.Load(&PlayerTeamInfoRow{row: crow})
		}
	}
}

func loadPlayerTotem(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_totem")
	}
	sql := "SELECT * FROM `player_totem`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerTotem), "SELECT MAX(`id`) FROM `player_totem` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iTotemId := res.Map("totem_id")
	iLevel := res.Map("level")
	iSkillId := res.Map("skill_id")
	for _, row := range res.Rows {
		crow := C.New_PlayerTotem()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.TotemId = C.int16_t(row.Int16(iTotemId))
		crow.Level = C.int8_t(row.Int8(iLevel))
		crow.SkillId = C.int16_t(row.Int16(iSkillId))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerTotem
		playerDb.PlayerTotem = crow
		if g_Hooks.PlayerTotem != nil {
			g_Hooks.PlayerTotem.Load(&PlayerTotemRow{row: crow})
		}
	}
}

func loadPlayerTotemInfo(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_totem_info")
	}
	sql := "SELECT * FROM `player_totem_info`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iIngotCallDailyNum := res.Map("ingot_call_daily_num")
	iIngotCallTimestamp := res.Map("ingot_call_timestamp")
	iRockRuneNum := res.Map("rock_rune_num")
	iJadeRuneNum := res.Map("jade_rune_num")
	iPos1 := res.Map("pos1")
	iPos2 := res.Map("pos2")
	iPos3 := res.Map("pos3")
	iPos4 := res.Map("pos4")
	iPos5 := res.Map("pos5")
	iIngotDrawTimes := res.Map("ingot_draw_times")
	for _, row := range res.Rows {
		crow := C.New_PlayerTotemInfo()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.IngotCallDailyNum = C.int16_t(row.Int16(iIngotCallDailyNum))
		crow.IngotCallTimestamp = C.int64_t(row.Int64(iIngotCallTimestamp))
		crow.RockRuneNum = C.int32_t(row.Int32(iRockRuneNum))
		crow.JadeRuneNum = C.int32_t(row.Int32(iJadeRuneNum))
		crow.Pos1 = C.int64_t(row.Int64(iPos1))
		crow.Pos2 = C.int64_t(row.Int64(iPos2))
		crow.Pos3 = C.int64_t(row.Int64(iPos3))
		crow.Pos4 = C.int64_t(row.Int64(iPos4))
		crow.Pos5 = C.int64_t(row.Int64(iPos5))
		crow.IngotDrawTimes = C.int16_t(row.Int16(iIngotDrawTimes))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerTotemInfo = crow
		if g_Hooks.PlayerTotemInfo != nil {
			g_Hooks.PlayerTotemInfo.Load(&PlayerTotemInfoRow{row: crow})
		}
	}
}

func loadPlayerTown(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_town")
	}
	sql := "SELECT * FROM `player_town`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iTownId := res.Map("town_id")
	iLock := res.Map("lock")
	iAtPosX := res.Map("at_pos_x")
	iAtPosY := res.Map("at_pos_y")
	for _, row := range res.Rows {
		crow := C.New_PlayerTown()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.TownId = C.int16_t(row.Int16(iTownId))
		crow.Lock = C.int32_t(row.Int32(iLock))
		crow.AtPosX = C.int16_t(row.Int16(iAtPosX))
		crow.AtPosY = C.int16_t(row.Int16(iAtPosY))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerTown = crow
		if g_Hooks.PlayerTown != nil {
			g_Hooks.PlayerTown.Load(&PlayerTownRow{row: crow})
		}
	}
}

func loadPlayerTraderRefreshState(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_trader_refresh_state")
	}
	sql := "SELECT * FROM `player_trader_refresh_state`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerTraderRefreshState), "SELECT MAX(`id`) FROM `player_trader_refresh_state` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iLastUpdateTime := res.Map("last_update_time")
	iAutoUpdateTime := res.Map("auto_update_time")
	iTraderId := res.Map("trader_id")
	iRefreshNum := res.Map("refresh_num")
	for _, row := range res.Rows {
		crow := C.New_PlayerTraderRefreshState()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.LastUpdateTime = C.int64_t(row.Int64(iLastUpdateTime))
		crow.AutoUpdateTime = C.int64_t(row.Int64(iAutoUpdateTime))
		crow.TraderId = C.int16_t(row.Int16(iTraderId))
		crow.RefreshNum = C.int16_t(row.Int16(iRefreshNum))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerTraderRefreshState
		playerDb.PlayerTraderRefreshState = crow
		if g_Hooks.PlayerTraderRefreshState != nil {
			g_Hooks.PlayerTraderRefreshState.Load(&PlayerTraderRefreshStateRow{row: crow})
		}
	}
}

func loadPlayerTraderStoreState(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_trader_store_state")
	}
	sql := "SELECT * FROM `player_trader_store_state`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerTraderStoreState), "SELECT MAX(`id`) FROM `player_trader_store_state` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iTraderId := res.Map("trader_id")
	iGridId := res.Map("grid_id")
	iItemId := res.Map("item_id")
	iNum := res.Map("num")
	iCost := res.Map("cost")
	iStock := res.Map("stock")
	iGoodsType := res.Map("goods_type")
	for _, row := range res.Rows {
		crow := C.New_PlayerTraderStoreState()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.TraderId = C.int16_t(row.Int16(iTraderId))
		crow.GridId = C.int32_t(row.Int32(iGridId))
		crow.ItemId = C.int16_t(row.Int16(iItemId))
		crow.Num = C.int16_t(row.Int16(iNum))
		crow.Cost = C.int64_t(row.Int64(iCost))
		crow.Stock = C.int8_t(row.Int8(iStock))
		crow.GoodsType = C.int8_t(row.Int8(iGoodsType))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerTraderStoreState
		playerDb.PlayerTraderStoreState = crow
		if g_Hooks.PlayerTraderStoreState != nil {
			g_Hooks.PlayerTraderStoreState.Load(&PlayerTraderStoreStateRow{row: crow})
		}
	}
}

func loadPlayerUseSkill(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_use_skill")
	}
	sql := "SELECT * FROM `player_use_skill`"
	if playerId == 0 {
		if !initRowId(db, &(g_RowIds.PlayerUseSkill), "SELECT MAX(`id`) FROM `player_use_skill` WHERE (pid >> 32) = " + itoa(sid) + "") {
			return
		}
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iPid := res.Map("pid")
	iRoleId := res.Map("role_id")
	iSkillId1 := res.Map("skill_id1")
	iSkillId4 := res.Map("skill_id4")
	iSkillId2 := res.Map("skill_id2")
	iSkillId3 := res.Map("skill_id3")
	for _, row := range res.Rows {
		crow := C.New_PlayerUseSkill()
		crow.Id = C.int64_t(row.Int64(iId))
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.RoleId = C.int8_t(row.Int8(iRoleId))
		crow.SkillId1 = C.int16_t(row.Int16(iSkillId1))
		crow.SkillId4 = C.int16_t(row.Int16(iSkillId4))
		crow.SkillId2 = C.int16_t(row.Int16(iSkillId2))
		crow.SkillId3 = C.int16_t(row.Int16(iSkillId3))
		playerDb := GetPlayerTables(int64(crow.Pid))
		crow.next =  playerDb.PlayerUseSkill
		playerDb.PlayerUseSkill = crow
		if g_Hooks.PlayerUseSkill != nil {
			g_Hooks.PlayerUseSkill.Load(&PlayerUseSkillRow{row: crow})
		}
	}
}

func loadPlayerVip(db *mysql.Connection, playerId int64, sid, wgn, procId int) {
	if playerId == 0 {
		log.Infof("load player_vip")
	}
	sql := "SELECT * FROM `player_vip`"
	if playerId == 0 {
		sql += " WHERE (pid >> 32) = " + itoa(sid) + " AND  pid % " + itoa(wgn) + " = " + itoa(procId)
	} else {
		sql += " WHERE pid = " + playerIdToStr(playerId)
	}
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	iPid := res.Map("pid")
	iIngot := res.Map("ingot")
	iLevel := res.Map("level")
	iCardId := res.Map("card_id")
	iPresentExp := res.Map("present_exp")
	for _, row := range res.Rows {
		crow := C.New_PlayerVip()
		crow.Pid = C.int64_t(row.Int64(iPid))
		crow.Ingot = C.int64_t(row.Int64(iIngot))
		crow.Level = C.int16_t(row.Int16(iLevel))
		row_CardId := row.Str(iCardId)
		crow.CardId = C.CString(string(row_CardId))
		crow.CardId_len = C.int(len(row_CardId))
		crow.PresentExp = C.int64_t(row.Int64(iPresentExp))
		playerDb := GetPlayerTables(int64(crow.Pid))
		playerDb.PlayerVip = crow
		if g_Hooks.PlayerVip != nil {
			g_Hooks.PlayerVip.Load(&PlayerVipRow{row: crow})
		}
	}
}

