package mdb

import "core/mysql"

func prepare(db *mysql.Connection, sql string) *mysql.Stmt {
	stmt, err := db.Prepare([]byte(sql))
	if err != nil {
		panic(err)
	}
	return stmt
}
var g_SQL tableSQL

type tableSQL struct {
	GlobalAnnouncement globalAnnouncementSQL
	GlobalArenaRank globalArenaRankSQL
	GlobalClique globalCliqueSQL
	GlobalCliqueBoat globalCliqueBoatSQL
	GlobalDespairBoss globalDespairBossSQL
	GlobalDespairBossHistory globalDespairBossHistorySQL
	GlobalDespairLand globalDespairLandSQL
	GlobalDespairLandBattleRecord globalDespairLandBattleRecordSQL
	GlobalDespairLandCamp globalDespairLandCampSQL
	GlobalGiftCardRecord globalGiftCardRecordSQL
	GlobalGroupBuyStatus globalGroupBuyStatusSQL
	GlobalMail globalMailSQL
	GlobalMailAttachments globalMailAttachmentsSQL
	GlobalTbXxdOnlinecnt globalTbXxdOnlinecntSQL
	Player playerSQL
	PlayerActivity playerActivitySQL
	PlayerAdditionQuest playerAdditionQuestSQL
	PlayerAnySdkOrder playerAnySdkOrderSQL
	PlayerArena playerArenaSQL
	PlayerArenaRank playerArenaRankSQL
	PlayerArenaRecord playerArenaRecordSQL
	PlayerAwakenGraphic playerAwakenGraphicSQL
	PlayerBattlePet playerBattlePetSQL
	PlayerBattlePetGrid playerBattlePetGridSQL
	PlayerBattlePetState playerBattlePetStateSQL
	PlayerChestState playerChestStateSQL
	PlayerCliqueKongfuAttrib playerCliqueKongfuAttribSQL
	PlayerCoins playerCoinsSQL
	PlayerCornucopia playerCornucopiaSQL
	PlayerDailyQuest playerDailyQuestSQL
	PlayerDailyQuestStarAwardInfo playerDailyQuestStarAwardInfoSQL
	PlayerDailySignInRecord playerDailySignInRecordSQL
	PlayerDailySignInState playerDailySignInStateSQL
	PlayerDespairLandCampState playerDespairLandCampStateSQL
	PlayerDespairLandCampStateHistory playerDespairLandCampStateHistorySQL
	PlayerDespairLandLevelRecord playerDespairLandLevelRecordSQL
	PlayerDespairLandState playerDespairLandStateSQL
	PlayerDrivingSwordEvent playerDrivingSwordEventSQL
	PlayerDrivingSwordEventExploring playerDrivingSwordEventExploringSQL
	PlayerDrivingSwordEventTreasure playerDrivingSwordEventTreasureSQL
	PlayerDrivingSwordEventVisiting playerDrivingSwordEventVisitingSQL
	PlayerDrivingSwordEventmask playerDrivingSwordEventmaskSQL
	PlayerDrivingSwordInfo playerDrivingSwordInfoSQL
	PlayerDrivingSwordMap playerDrivingSwordMapSQL
	PlayerEquipment playerEquipmentSQL
	PlayerEventAwardRecord playerEventAwardRecordSQL
	PlayerEventDailyOnline playerEventDailyOnlineSQL
	PlayerEventVnDailyRecharge playerEventVnDailyRechargeSQL
	PlayerEventVnDragonBallRecord playerEventVnDragonBallRecordSQL
	PlayerEventsFightPower playerEventsFightPowerSQL
	PlayerEventsIngotRecord playerEventsIngotRecordSQL
	PlayerExtendLevel playerExtendLevelSQL
	PlayerExtendQuest playerExtendQuestSQL
	PlayerFame playerFameSQL
	PlayerFashion playerFashionSQL
	PlayerFashionState playerFashionStateSQL
	PlayerFateBoxState playerFateBoxStateSQL
	PlayerFightNum playerFightNumSQL
	PlayerFirstCharge playerFirstChargeSQL
	PlayerFormation playerFormationSQL
	PlayerFuncKey playerFuncKeySQL
	PlayerGhost playerGhostSQL
	PlayerGhostEquipment playerGhostEquipmentSQL
	PlayerGhostState playerGhostStateSQL
	PlayerGlobalChatState playerGlobalChatStateSQL
	PlayerGlobalCliqueBuilding playerGlobalCliqueBuildingSQL
	PlayerGlobalCliqueBuildingQuest playerGlobalCliqueBuildingQuestSQL
	PlayerGlobalCliqueDailyQuest playerGlobalCliqueDailyQuestSQL
	PlayerGlobalCliqueEscort playerGlobalCliqueEscortSQL
	PlayerGlobalCliqueEscortMessage playerGlobalCliqueEscortMessageSQL
	PlayerGlobalCliqueHijack playerGlobalCliqueHijackSQL
	PlayerGlobalCliqueInfo playerGlobalCliqueInfoSQL
	PlayerGlobalCliqueKongfu playerGlobalCliqueKongfuSQL
	PlayerGlobalFriend playerGlobalFriendSQL
	PlayerGlobalFriendChat playerGlobalFriendChatSQL
	PlayerGlobalFriendState playerGlobalFriendStateSQL
	PlayerGlobalGiftCardRecord playerGlobalGiftCardRecordSQL
	PlayerGlobalMailState playerGlobalMailStateSQL
	PlayerGlobalWorldChatState playerGlobalWorldChatStateSQL
	PlayerHardLevel playerHardLevelSQL
	PlayerHardLevelRecord playerHardLevelRecordSQL
	PlayerHeart playerHeartSQL
	PlayerHeartDraw playerHeartDrawSQL
	PlayerHeartDrawCardRecord playerHeartDrawCardRecordSQL
	PlayerHeartDrawWheelRecord playerHeartDrawWheelRecordSQL
	PlayerInfo playerInfoSQL
	PlayerIsOperated playerIsOperatedSQL
	PlayerItem playerItemSQL
	PlayerItemAppendix playerItemAppendixSQL
	PlayerItemBuyback playerItemBuybackSQL
	PlayerItemRecastState playerItemRecastStateSQL
	PlayerLevelAwardInfo playerLevelAwardInfoSQL
	PlayerLoginAwardRecord playerLoginAwardRecordSQL
	PlayerMail playerMailSQL
	PlayerMailAttachment playerMailAttachmentSQL
	PlayerMailAttachmentLg playerMailAttachmentLgSQL
	PlayerMailLg playerMailLgSQL
	PlayerMeditationState playerMeditationStateSQL
	PlayerMission playerMissionSQL
	PlayerMissionLevel playerMissionLevelSQL
	PlayerMissionLevelRecord playerMissionLevelRecordSQL
	PlayerMissionLevelStateBin playerMissionLevelStateBinSQL
	PlayerMissionRecord playerMissionRecordSQL
	PlayerMissionStarAward playerMissionStarAwardSQL
	PlayerMoneyTree playerMoneyTreeSQL
	PlayerMonthCardAwardRecord playerMonthCardAwardRecordSQL
	PlayerMonthCardInfo playerMonthCardInfoSQL
	PlayerMultiLevelInfo playerMultiLevelInfoSQL
	PlayerNewYearConsumeRecord playerNewYearConsumeRecordSQL
	PlayerNpcTalkRecord playerNpcTalkRecordSQL
	PlayerOpenedTownTreasure playerOpenedTownTreasureSQL
	PlayerPhysical playerPhysicalSQL
	PlayerPurchaseRecord playerPurchaseRecordSQL
	PlayerPushNotifySwitch playerPushNotifySwitchSQL
	PlayerPveState playerPveStateSQL
	PlayerQqVipGift playerQqVipGiftSQL
	PlayerQuest playerQuestSQL
	PlayerRainbowLevel playerRainbowLevelSQL
	PlayerRainbowLevelStateBin playerRainbowLevelStateBinSQL
	PlayerRole playerRoleSQL
	PlayerSealedbook playerSealedbookSQL
	PlayerSendHeartRecord playerSendHeartRecordSQL
	PlayerSkill playerSkillSQL
	PlayerState playerStateSQL
	PlayerSwordSoul playerSwordSoulSQL
	PlayerSwordSoulEquipment playerSwordSoulEquipmentSQL
	PlayerSwordSoulState playerSwordSoulStateSQL
	PlayerTaoyuanBless playerTaoyuanBlessSQL
	PlayerTaoyuanFileds playerTaoyuanFiledsSQL
	PlayerTaoyuanHeart playerTaoyuanHeartSQL
	PlayerTaoyuanItem playerTaoyuanItemSQL
	PlayerTaoyuanMessage playerTaoyuanMessageSQL
	PlayerTaoyuanProduct playerTaoyuanProductSQL
	PlayerTaoyuanPurchaseRecord playerTaoyuanPurchaseRecordSQL
	PlayerTaoyuanQuest playerTaoyuanQuestSQL
	PlayerTbXxdRoleinfo playerTbXxdRoleinfoSQL
	PlayerTeamInfo playerTeamInfoSQL
	PlayerTotem playerTotemSQL
	PlayerTotemInfo playerTotemInfoSQL
	PlayerTown playerTownSQL
	PlayerTraderRefreshState playerTraderRefreshStateSQL
	PlayerTraderStoreState playerTraderStoreStateSQL
	PlayerUseSkill playerUseSkillSQL
	PlayerVip playerVipSQL
}

type globalAnnouncementSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type globalArenaRankSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type globalCliqueSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type globalCliqueBoatSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type globalDespairBossSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type globalDespairBossHistorySQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type globalDespairLandSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type globalDespairLandBattleRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type globalDespairLandCampSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type globalGiftCardRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type globalGroupBuyStatusSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type globalMailSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type globalMailAttachmentsSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type globalTbXxdOnlinecntSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerActivitySQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerAdditionQuestSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerAnySdkOrderSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerArenaSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerArenaRankSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerArenaRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerAwakenGraphicSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerBattlePetSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerBattlePetGridSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerBattlePetStateSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerChestStateSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerCliqueKongfuAttribSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerCoinsSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerCornucopiaSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerDailyQuestSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerDailyQuestStarAwardInfoSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerDailySignInRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerDailySignInStateSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerDespairLandCampStateSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerDespairLandCampStateHistorySQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerDespairLandLevelRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerDespairLandStateSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerDrivingSwordEventSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerDrivingSwordEventExploringSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerDrivingSwordEventTreasureSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerDrivingSwordEventVisitingSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerDrivingSwordEventmaskSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerDrivingSwordInfoSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerDrivingSwordMapSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerEquipmentSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerEventAwardRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerEventDailyOnlineSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerEventVnDailyRechargeSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerEventVnDragonBallRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerEventsFightPowerSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerEventsIngotRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerExtendLevelSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerExtendQuestSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerFameSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerFashionSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerFashionStateSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerFateBoxStateSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerFightNumSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerFirstChargeSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerFormationSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerFuncKeySQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerGhostSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerGhostEquipmentSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerGhostStateSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerGlobalChatStateSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerGlobalCliqueBuildingSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerGlobalCliqueBuildingQuestSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerGlobalCliqueDailyQuestSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerGlobalCliqueEscortSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerGlobalCliqueEscortMessageSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerGlobalCliqueHijackSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerGlobalCliqueInfoSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerGlobalCliqueKongfuSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerGlobalFriendSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerGlobalFriendChatSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerGlobalFriendStateSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerGlobalGiftCardRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerGlobalMailStateSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerGlobalWorldChatStateSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerHardLevelSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerHardLevelRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerHeartSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerHeartDrawSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerHeartDrawCardRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerHeartDrawWheelRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerInfoSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerIsOperatedSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerItemSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerItemAppendixSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerItemBuybackSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerItemRecastStateSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerLevelAwardInfoSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerLoginAwardRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerMailSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerMailAttachmentSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerMailAttachmentLgSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerMailLgSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerMeditationStateSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerMissionSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerMissionLevelSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerMissionLevelRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerMissionLevelStateBinSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerMissionRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerMissionStarAwardSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerMoneyTreeSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerMonthCardAwardRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerMonthCardInfoSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerMultiLevelInfoSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerNewYearConsumeRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerNpcTalkRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerOpenedTownTreasureSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerPhysicalSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerPurchaseRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerPushNotifySwitchSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerPveStateSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerQqVipGiftSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerQuestSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerRainbowLevelSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerRainbowLevelStateBinSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerRoleSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerSealedbookSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerSendHeartRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerSkillSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerStateSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerSwordSoulSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerSwordSoulEquipmentSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerSwordSoulStateSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerTaoyuanBlessSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerTaoyuanFiledsSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerTaoyuanHeartSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerTaoyuanItemSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerTaoyuanMessageSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerTaoyuanProductSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerTaoyuanPurchaseRecordSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerTaoyuanQuestSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerTbXxdRoleinfoSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerTeamInfoSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerTotemSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerTotemInfoSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerTownSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerTraderRefreshStateSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerTraderStoreStateSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerUseSkillSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

type playerVipSQL struct {
	Insert *mysql.Stmt
	Delete *mysql.Stmt
	Update *mysql.Stmt
}

func prepareSQL(db *mysql.Connection) {
	g_SQL.GlobalAnnouncement.Insert = prepare(db, "INSERT INTO `global_announcement` SET `id`=?, `expire_time`=?, `tpl_id`=?, `parameters`=?, `content`=?, `send_time`=?, `spacing_time`=?")
	g_SQL.GlobalAnnouncement.Update = prepare(db, "UPDATE `global_announcement` SET `expire_time`=?, `tpl_id`=?, `parameters`=?, `content`=?, `send_time`=?, `spacing_time`=? WHERE `id`=?")
	g_SQL.GlobalAnnouncement.Delete = prepare(db, "DELETE FROM `global_announcement` WHERE `id`=?")
	g_SQL.GlobalArenaRank.Insert = prepare(db, "INSERT INTO `global_arena_rank` SET `rank`=?, `pid`=?")
	g_SQL.GlobalArenaRank.Update = prepare(db, "UPDATE `global_arena_rank` SET `pid`=? WHERE `rank`=?")
	g_SQL.GlobalArenaRank.Delete = prepare(db, "DELETE FROM `global_arena_rank` WHERE `rank`=?")
	g_SQL.GlobalClique.Insert = prepare(db, "INSERT INTO `global_clique` SET `id`=?, `name`=?, `announce`=?, `total_donate_coins`=?, `owner_pid`=?, `owner_login_time`=?, `manger_pid1`=?, `manger_pid2`=?, `center_building_coins`=?, `temple_building_coins`=?, `health_building_coins`=?, `attack_building_coins`=?, `defense_building_coins`=?, `store_building_coins`=?, `bank_building_coins`=?, `center_building_level`=?, `temple_building_level`=?, `health_building_level`=?, `attack_building_level`=?, `defense_building_level`=?, `bank_building_level`=?, `member_json`=?, `auto_audit`=?, `auto_audit_level`=?, `contrib`=?, `contrib_clear_time`=?, `recruit_time`=?, `worship_time`=?, `worship_cnt`=?, `store_send_time`=?, `store_send_cnt`=?")
	g_SQL.GlobalClique.Update = prepare(db, "UPDATE `global_clique` SET `name`=?, `announce`=?, `total_donate_coins`=?, `owner_pid`=?, `owner_login_time`=?, `manger_pid1`=?, `manger_pid2`=?, `center_building_coins`=?, `temple_building_coins`=?, `health_building_coins`=?, `attack_building_coins`=?, `defense_building_coins`=?, `store_building_coins`=?, `bank_building_coins`=?, `center_building_level`=?, `temple_building_level`=?, `health_building_level`=?, `attack_building_level`=?, `defense_building_level`=?, `bank_building_level`=?, `member_json`=?, `auto_audit`=?, `auto_audit_level`=?, `contrib`=?, `contrib_clear_time`=?, `recruit_time`=?, `worship_time`=?, `worship_cnt`=?, `store_send_time`=?, `store_send_cnt`=? WHERE `id`=?")
	g_SQL.GlobalClique.Delete = prepare(db, "DELETE FROM `global_clique` WHERE `id`=?")
	g_SQL.GlobalCliqueBoat.Insert = prepare(db, "INSERT INTO `global_clique_boat` SET `id`=?, `clique_id`=?, `boat_type`=?, `owner_pid`=?, `escort_start_timestamp`=?, `total_escort_time`=?, `recover_pid`=?, `hijacker_pid`=?, `hijack_timestamp`=?, `status`=?")
	g_SQL.GlobalCliqueBoat.Update = prepare(db, "UPDATE `global_clique_boat` SET `clique_id`=?, `boat_type`=?, `owner_pid`=?, `escort_start_timestamp`=?, `total_escort_time`=?, `recover_pid`=?, `hijacker_pid`=?, `hijack_timestamp`=?, `status`=? WHERE `id`=?")
	g_SQL.GlobalCliqueBoat.Delete = prepare(db, "DELETE FROM `global_clique_boat` WHERE `id`=?")
	g_SQL.GlobalDespairBoss.Insert = prepare(db, "INSERT INTO `global_despair_boss` SET `camp_type`=?, `level`=?, `timestamp`=?, `dead_timestamp`=?, `hp`=?, `max_hp`=?")
	g_SQL.GlobalDespairBoss.Update = prepare(db, "UPDATE `global_despair_boss` SET `level`=?, `timestamp`=?, `dead_timestamp`=?, `hp`=?, `max_hp`=? WHERE `camp_type`=?")
	g_SQL.GlobalDespairBoss.Delete = prepare(db, "DELETE FROM `global_despair_boss` WHERE `camp_type`=?")
	g_SQL.GlobalDespairBossHistory.Insert = prepare(db, "INSERT INTO `global_despair_boss_history` SET `id`=?, `version`=?, `camp_type`=?, `level`=?, `timestamp`=?, `start_timestamp`=?, `dead_timestamp`=?")
	g_SQL.GlobalDespairBossHistory.Update = prepare(db, "UPDATE `global_despair_boss_history` SET `version`=?, `camp_type`=?, `level`=?, `timestamp`=?, `start_timestamp`=?, `dead_timestamp`=? WHERE `id`=?")
	g_SQL.GlobalDespairBossHistory.Delete = prepare(db, "DELETE FROM `global_despair_boss_history` WHERE `id`=?")
	g_SQL.GlobalDespairLand.Insert = prepare(db, "INSERT INTO `global_despair_land` SET `id`=?, `version`=?, `timestamp`=?")
	g_SQL.GlobalDespairLand.Update = prepare(db, "UPDATE `global_despair_land` SET `version`=?, `timestamp`=? WHERE `id`=?")
	g_SQL.GlobalDespairLand.Delete = prepare(db, "DELETE FROM `global_despair_land` WHERE `id`=?")
	g_SQL.GlobalDespairLandBattleRecord.Insert = prepare(db, "INSERT INTO `global_despair_land_battle_record` SET `id`=?, `pid`=?, `fightnum`=?, `timestamp`=?, `tag`=?, `battle_version`=?, `camp_type`=?, `level_id`=?, `record`=?")
	g_SQL.GlobalDespairLandBattleRecord.Update = prepare(db, "UPDATE `global_despair_land_battle_record` SET `pid`=?, `fightnum`=?, `timestamp`=?, `tag`=?, `battle_version`=?, `camp_type`=?, `level_id`=?, `record`=? WHERE `id`=?")
	g_SQL.GlobalDespairLandBattleRecord.Delete = prepare(db, "DELETE FROM `global_despair_land_battle_record` WHERE `id`=?")
	g_SQL.GlobalDespairLandCamp.Insert = prepare(db, "INSERT INTO `global_despair_land_camp` SET `camp_type`=?, `timestamp`=?, `dead_timestamp`=?, `level`=?")
	g_SQL.GlobalDespairLandCamp.Update = prepare(db, "UPDATE `global_despair_land_camp` SET `timestamp`=?, `dead_timestamp`=?, `level`=? WHERE `camp_type`=?")
	g_SQL.GlobalDespairLandCamp.Delete = prepare(db, "DELETE FROM `global_despair_land_camp` WHERE `camp_type`=?")
	g_SQL.GlobalGiftCardRecord.Insert = prepare(db, "INSERT INTO `global_gift_card_record` SET `id`=?, `pid`=?, `code`=?, `timestamp`=?")
	g_SQL.GlobalGiftCardRecord.Update = prepare(db, "UPDATE `global_gift_card_record` SET `pid`=?, `code`=?, `timestamp`=? WHERE `id`=?")
	g_SQL.GlobalGiftCardRecord.Delete = prepare(db, "DELETE FROM `global_gift_card_record` WHERE `id`=?")
	g_SQL.GlobalGroupBuyStatus.Insert = prepare(db, "INSERT INTO `global_group_buy_status` SET `id`=?, `cid`=?, `status`=?")
	g_SQL.GlobalGroupBuyStatus.Update = prepare(db, "UPDATE `global_group_buy_status` SET `cid`=?, `status`=? WHERE `id`=?")
	g_SQL.GlobalGroupBuyStatus.Delete = prepare(db, "DELETE FROM `global_group_buy_status` WHERE `id`=?")
	g_SQL.GlobalMail.Insert = prepare(db, "INSERT INTO `global_mail` SET `id`=?, `send_time`=?, `expire_time`=?, `title`=?, `content`=?, `priority`=?, `min_level`=?, `max_level`=?, `min_vip_level`=?, `max_vip_level`=?, `min_create_time`=?, `max_create_time`=?")
	g_SQL.GlobalMail.Update = prepare(db, "UPDATE `global_mail` SET `send_time`=?, `expire_time`=?, `title`=?, `content`=?, `priority`=?, `min_level`=?, `max_level`=?, `min_vip_level`=?, `max_vip_level`=?, `min_create_time`=?, `max_create_time`=? WHERE `id`=?")
	g_SQL.GlobalMail.Delete = prepare(db, "DELETE FROM `global_mail` WHERE `id`=?")
	g_SQL.GlobalMailAttachments.Insert = prepare(db, "INSERT INTO `global_mail_attachments` SET `id`=?, `global_mail_id`=?, `item_id`=?, `attachment_type`=?, `item_num`=?")
	g_SQL.GlobalMailAttachments.Update = prepare(db, "UPDATE `global_mail_attachments` SET `global_mail_id`=?, `item_id`=?, `attachment_type`=?, `item_num`=? WHERE `id`=?")
	g_SQL.GlobalMailAttachments.Delete = prepare(db, "DELETE FROM `global_mail_attachments` WHERE `id`=?")
	g_SQL.GlobalTbXxdOnlinecnt.Insert = prepare(db, "INSERT INTO `global_tb_xxd_onlinecnt` SET `id`=?, `gameappid`=?, `timekey`=?, `gsid`=?, `onlinecntios`=?, `onlinecntandroid`=?, `cid`=?")
	g_SQL.GlobalTbXxdOnlinecnt.Update = prepare(db, "UPDATE `global_tb_xxd_onlinecnt` SET `gameappid`=?, `timekey`=?, `gsid`=?, `onlinecntios`=?, `onlinecntandroid`=?, `cid`=? WHERE `id`=?")
	g_SQL.GlobalTbXxdOnlinecnt.Delete = prepare(db, "DELETE FROM `global_tb_xxd_onlinecnt` WHERE `id`=?")
	g_SQL.Player.Insert = prepare(db, "INSERT INTO `player` SET `id`=?, `user`=?, `nick`=?, `main_role_id`=?, `cid`=?")
	g_SQL.Player.Update = prepare(db, "UPDATE `player` SET `user`=?, `nick`=?, `main_role_id`=?, `cid`=? WHERE `id`=?")
	g_SQL.Player.Delete = prepare(db, "DELETE FROM `player` WHERE `id`=?")
	g_SQL.PlayerActivity.Insert = prepare(db, "INSERT INTO `player_activity` SET `pid`=?, `activity`=?, `last_update`=?")
	g_SQL.PlayerActivity.Update = prepare(db, "UPDATE `player_activity` SET `activity`=?, `last_update`=? WHERE `pid`=?")
	g_SQL.PlayerActivity.Delete = prepare(db, "DELETE FROM `player_activity` WHERE `pid`=?")
	g_SQL.PlayerAdditionQuest.Insert = prepare(db, "INSERT INTO `player_addition_quest` SET `id`=?, `pid`=?, `serial_number`=?, `quest_id`=?, `lock`=?, `progress`=?, `state`=?")
	g_SQL.PlayerAdditionQuest.Update = prepare(db, "UPDATE `player_addition_quest` SET `pid`=?, `serial_number`=?, `quest_id`=?, `lock`=?, `progress`=?, `state`=? WHERE `id`=?")
	g_SQL.PlayerAdditionQuest.Delete = prepare(db, "DELETE FROM `player_addition_quest` WHERE `id`=?")
	g_SQL.PlayerAnySdkOrder.Insert = prepare(db, "INSERT INTO `player_any_sdk_order` SET `id`=?, `pid`=?, `order_id`=?, `present`=?")
	g_SQL.PlayerAnySdkOrder.Update = prepare(db, "UPDATE `player_any_sdk_order` SET `pid`=?, `order_id`=?, `present`=? WHERE `id`=?")
	g_SQL.PlayerAnySdkOrder.Delete = prepare(db, "DELETE FROM `player_any_sdk_order` WHERE `id`=?")
	g_SQL.PlayerArena.Insert = prepare(db, "INSERT INTO `player_arena` SET `pid`=?, `daily_num`=?, `failed_cd_time`=?, `battle_time`=?, `win_times`=?, `daily_award_item`=?, `new_record_count`=?, `daily_fame`=?, `buy_times`=?")
	g_SQL.PlayerArena.Update = prepare(db, "UPDATE `player_arena` SET `daily_num`=?, `failed_cd_time`=?, `battle_time`=?, `win_times`=?, `daily_award_item`=?, `new_record_count`=?, `daily_fame`=?, `buy_times`=? WHERE `pid`=?")
	g_SQL.PlayerArena.Delete = prepare(db, "DELETE FROM `player_arena` WHERE `pid`=?")
	g_SQL.PlayerArenaRank.Insert = prepare(db, "INSERT INTO `player_arena_rank` SET `pid`=?, `rank`=?, `rank1`=?, `rank2`=?, `rank3`=?, `time`=?")
	g_SQL.PlayerArenaRank.Update = prepare(db, "UPDATE `player_arena_rank` SET `rank`=?, `rank1`=?, `rank2`=?, `rank3`=?, `time`=? WHERE `pid`=?")
	g_SQL.PlayerArenaRank.Delete = prepare(db, "DELETE FROM `player_arena_rank` WHERE `pid`=?")
	g_SQL.PlayerArenaRecord.Insert = prepare(db, "INSERT INTO `player_arena_record` SET `id`=?, `pid`=?, `mode`=?, `old_rank`=?, `new_rank`=?, `fight_num`=?, `target_pid`=?, `target_nick`=?, `target_old_rank`=?, `target_new_rank`=?, `target_fight_num`=?, `record_time`=?")
	g_SQL.PlayerArenaRecord.Update = prepare(db, "UPDATE `player_arena_record` SET `pid`=?, `mode`=?, `old_rank`=?, `new_rank`=?, `fight_num`=?, `target_pid`=?, `target_nick`=?, `target_old_rank`=?, `target_new_rank`=?, `target_fight_num`=?, `record_time`=? WHERE `id`=?")
	g_SQL.PlayerArenaRecord.Delete = prepare(db, "DELETE FROM `player_arena_record` WHERE `id`=?")
	g_SQL.PlayerAwakenGraphic.Insert = prepare(db, "INSERT INTO `player_awaken_graphic` SET `id`=?, `pid`=?, `role_id`=?, `attr_impl`=?, `level`=?")
	g_SQL.PlayerAwakenGraphic.Update = prepare(db, "UPDATE `player_awaken_graphic` SET `pid`=?, `role_id`=?, `attr_impl`=?, `level`=? WHERE `id`=?")
	g_SQL.PlayerAwakenGraphic.Delete = prepare(db, "DELETE FROM `player_awaken_graphic` WHERE `id`=?")
	g_SQL.PlayerBattlePet.Insert = prepare(db, "INSERT INTO `player_battle_pet` SET `id`=?, `pid`=?, `battle_pet_id`=?, `level`=?, `exp`=?, `skill_level`=?")
	g_SQL.PlayerBattlePet.Update = prepare(db, "UPDATE `player_battle_pet` SET `pid`=?, `battle_pet_id`=?, `level`=?, `exp`=?, `skill_level`=? WHERE `id`=?")
	g_SQL.PlayerBattlePet.Delete = prepare(db, "DELETE FROM `player_battle_pet` WHERE `id`=?")
	g_SQL.PlayerBattlePetGrid.Insert = prepare(db, "INSERT INTO `player_battle_pet_grid` SET `id`=?, `pid`=?, `grid_id`=?, `battle_pet_id`=?")
	g_SQL.PlayerBattlePetGrid.Update = prepare(db, "UPDATE `player_battle_pet_grid` SET `pid`=?, `grid_id`=?, `battle_pet_id`=? WHERE `id`=?")
	g_SQL.PlayerBattlePetGrid.Delete = prepare(db, "DELETE FROM `player_battle_pet_grid` WHERE `id`=?")
	g_SQL.PlayerBattlePetState.Insert = prepare(db, "INSERT INTO `player_battle_pet_state` SET `pid`=?, `upgrade_by_ingot_num`=?, `upgrade_by_ingot_time`=?")
	g_SQL.PlayerBattlePetState.Update = prepare(db, "UPDATE `player_battle_pet_state` SET `upgrade_by_ingot_num`=?, `upgrade_by_ingot_time`=? WHERE `pid`=?")
	g_SQL.PlayerBattlePetState.Delete = prepare(db, "DELETE FROM `player_battle_pet_state` WHERE `pid`=?")
	g_SQL.PlayerChestState.Insert = prepare(db, "INSERT INTO `player_chest_state` SET `pid`=?, `free_coin_chest_num`=?, `last_free_coin_chest_at`=?, `coin_chest_num`=?, `coin_chest_ten_num`=?, `is_first_coin_ten`=?, `coin_chest_consume`=?, `last_coin_chest_at`=?, `last_free_ingot_chest_at`=?, `ingot_chest_num`=?, `ingot_chest_ten_num`=?, `is_first_ingot_ten`=?, `ingot_chest_consume`=?, `last_ingot_chest_at`=?, `last_free_pet_chest_at`=?")
	g_SQL.PlayerChestState.Update = prepare(db, "UPDATE `player_chest_state` SET `free_coin_chest_num`=?, `last_free_coin_chest_at`=?, `coin_chest_num`=?, `coin_chest_ten_num`=?, `is_first_coin_ten`=?, `coin_chest_consume`=?, `last_coin_chest_at`=?, `last_free_ingot_chest_at`=?, `ingot_chest_num`=?, `ingot_chest_ten_num`=?, `is_first_ingot_ten`=?, `ingot_chest_consume`=?, `last_ingot_chest_at`=?, `last_free_pet_chest_at`=? WHERE `pid`=?")
	g_SQL.PlayerChestState.Delete = prepare(db, "DELETE FROM `player_chest_state` WHERE `pid`=?")
	g_SQL.PlayerCliqueKongfuAttrib.Insert = prepare(db, "INSERT INTO `player_clique_kongfu_attrib` SET `pid`=?, `health`=?, `attack`=?, `defence`=?")
	g_SQL.PlayerCliqueKongfuAttrib.Update = prepare(db, "UPDATE `player_clique_kongfu_attrib` SET `health`=?, `attack`=?, `defence`=? WHERE `pid`=?")
	g_SQL.PlayerCliqueKongfuAttrib.Delete = prepare(db, "DELETE FROM `player_clique_kongfu_attrib` WHERE `pid`=?")
	g_SQL.PlayerCoins.Insert = prepare(db, "INSERT INTO `player_coins` SET `pid`=?, `buy_time`=?, `daily_count`=?, `batch_bought`=?")
	g_SQL.PlayerCoins.Update = prepare(db, "UPDATE `player_coins` SET `buy_time`=?, `daily_count`=?, `batch_bought`=? WHERE `pid`=?")
	g_SQL.PlayerCoins.Delete = prepare(db, "DELETE FROM `player_coins` WHERE `pid`=?")
	g_SQL.PlayerCornucopia.Insert = prepare(db, "INSERT INTO `player_cornucopia` SET `pid`=?, `open_time`=?, `daily_count`=?")
	g_SQL.PlayerCornucopia.Update = prepare(db, "UPDATE `player_cornucopia` SET `open_time`=?, `daily_count`=? WHERE `pid`=?")
	g_SQL.PlayerCornucopia.Delete = prepare(db, "DELETE FROM `player_cornucopia` WHERE `pid`=?")
	g_SQL.PlayerDailyQuest.Insert = prepare(db, "INSERT INTO `player_daily_quest` SET `id`=?, `pid`=?, `quest_id`=?, `finish_count`=?, `last_update_time`=?, `award_status`=?, `class`=?")
	g_SQL.PlayerDailyQuest.Update = prepare(db, "UPDATE `player_daily_quest` SET `pid`=?, `quest_id`=?, `finish_count`=?, `last_update_time`=?, `award_status`=?, `class`=? WHERE `id`=?")
	g_SQL.PlayerDailyQuest.Delete = prepare(db, "DELETE FROM `player_daily_quest` WHERE `id`=?")
	g_SQL.PlayerDailyQuestStarAwardInfo.Insert = prepare(db, "INSERT INTO `player_daily_quest_star_award_info` SET `pid`=?, `stars`=?, `lastupdatetime`=?, `awarded`=?")
	g_SQL.PlayerDailyQuestStarAwardInfo.Update = prepare(db, "UPDATE `player_daily_quest_star_award_info` SET `stars`=?, `lastupdatetime`=?, `awarded`=? WHERE `pid`=?")
	g_SQL.PlayerDailyQuestStarAwardInfo.Delete = prepare(db, "DELETE FROM `player_daily_quest_star_award_info` WHERE `pid`=?")
	g_SQL.PlayerDailySignInRecord.Insert = prepare(db, "INSERT INTO `player_daily_sign_in_record` SET `id`=?, `pid`=?, `sign_in_time`=?")
	g_SQL.PlayerDailySignInRecord.Update = prepare(db, "UPDATE `player_daily_sign_in_record` SET `pid`=?, `sign_in_time`=? WHERE `id`=?")
	g_SQL.PlayerDailySignInRecord.Delete = prepare(db, "DELETE FROM `player_daily_sign_in_record` WHERE `id`=?")
	g_SQL.PlayerDailySignInState.Insert = prepare(db, "INSERT INTO `player_daily_sign_in_state` SET `pid`=?, `update_time`=?, `record`=?, `signed_today`=?")
	g_SQL.PlayerDailySignInState.Update = prepare(db, "UPDATE `player_daily_sign_in_state` SET `update_time`=?, `record`=?, `signed_today`=? WHERE `pid`=?")
	g_SQL.PlayerDailySignInState.Delete = prepare(db, "DELETE FROM `player_daily_sign_in_state` WHERE `pid`=?")
	g_SQL.PlayerDespairLandCampState.Insert = prepare(db, "INSERT INTO `player_despair_land_camp_state` SET `id`=?, `pid`=?, `camp_type`=?, `timestamp`=?, `point`=?, `kill_num`=?, `dead_num`=?, `dead_num_boss`=?, `hurt`=?, `boss_battle_num`=?, `daily_boss_battle_num`=?, `boss_battle_timestamp`=?, `awarded`=?")
	g_SQL.PlayerDespairLandCampState.Update = prepare(db, "UPDATE `player_despair_land_camp_state` SET `pid`=?, `camp_type`=?, `timestamp`=?, `point`=?, `kill_num`=?, `dead_num`=?, `dead_num_boss`=?, `hurt`=?, `boss_battle_num`=?, `daily_boss_battle_num`=?, `boss_battle_timestamp`=?, `awarded`=? WHERE `id`=?")
	g_SQL.PlayerDespairLandCampState.Delete = prepare(db, "DELETE FROM `player_despair_land_camp_state` WHERE `id`=?")
	g_SQL.PlayerDespairLandCampStateHistory.Insert = prepare(db, "INSERT INTO `player_despair_land_camp_state_history` SET `id`=?, `pid`=?, `camp_type`=?, `timestamp`=?, `point`=?, `kill_num`=?, `dead_num`=?, `dead_num_boss`=?, `hurt`=?, `boss_battle_num`=?, `daily_boss_battle_num`=?, `boss_battle_timestamp`=?, `awarded`=?")
	g_SQL.PlayerDespairLandCampStateHistory.Update = prepare(db, "UPDATE `player_despair_land_camp_state_history` SET `pid`=?, `camp_type`=?, `timestamp`=?, `point`=?, `kill_num`=?, `dead_num`=?, `dead_num_boss`=?, `hurt`=?, `boss_battle_num`=?, `daily_boss_battle_num`=?, `boss_battle_timestamp`=?, `awarded`=? WHERE `id`=?")
	g_SQL.PlayerDespairLandCampStateHistory.Delete = prepare(db, "DELETE FROM `player_despair_land_camp_state_history` WHERE `id`=?")
	g_SQL.PlayerDespairLandLevelRecord.Insert = prepare(db, "INSERT INTO `player_despair_land_level_record` SET `id`=?, `pid`=?, `camp_type`=?, `timestamp`=?, `level_id`=?, `round`=?, `passed_timestamp`=?, `first_passed_timestamp`=?, `first_passed_fightnum`=?, `passed_award`=?, `perfect_award`=?, `milestone_award`=?")
	g_SQL.PlayerDespairLandLevelRecord.Update = prepare(db, "UPDATE `player_despair_land_level_record` SET `pid`=?, `camp_type`=?, `timestamp`=?, `level_id`=?, `round`=?, `passed_timestamp`=?, `first_passed_timestamp`=?, `first_passed_fightnum`=?, `passed_award`=?, `perfect_award`=?, `milestone_award`=? WHERE `id`=?")
	g_SQL.PlayerDespairLandLevelRecord.Delete = prepare(db, "DELETE FROM `player_despair_land_level_record` WHERE `id`=?")
	g_SQL.PlayerDespairLandState.Insert = prepare(db, "INSERT INTO `player_despair_land_state` SET `pid`=?, `point`=?, `kill_num`=?, `dead_num`=?, `daily_battle_num`=?, `daily_battle_timestamp`=?, `daily_bought_battle_num`=?, `daily_bought_timestamp`=?, `daily_boss_bought_timestamp`=?, `daily_boss_beast_bought_battle_num`=?, `daily_boss_walking_dead_bought_battle_num`=?, `daily_boss_devil_bought_battle_num`=?")
	g_SQL.PlayerDespairLandState.Update = prepare(db, "UPDATE `player_despair_land_state` SET `point`=?, `kill_num`=?, `dead_num`=?, `daily_battle_num`=?, `daily_battle_timestamp`=?, `daily_bought_battle_num`=?, `daily_bought_timestamp`=?, `daily_boss_bought_timestamp`=?, `daily_boss_beast_bought_battle_num`=?, `daily_boss_walking_dead_bought_battle_num`=?, `daily_boss_devil_bought_battle_num`=? WHERE `pid`=?")
	g_SQL.PlayerDespairLandState.Delete = prepare(db, "DELETE FROM `player_despair_land_state` WHERE `pid`=?")
	g_SQL.PlayerDrivingSwordEvent.Insert = prepare(db, "INSERT INTO `player_driving_sword_event` SET `id`=?, `pid`=?, `cloud_id`=?, `x`=?, `y`=?, `event_type`=?, `data_id`=?")
	g_SQL.PlayerDrivingSwordEvent.Update = prepare(db, "UPDATE `player_driving_sword_event` SET `pid`=?, `cloud_id`=?, `x`=?, `y`=?, `event_type`=?, `data_id`=? WHERE `id`=?")
	g_SQL.PlayerDrivingSwordEvent.Delete = prepare(db, "DELETE FROM `player_driving_sword_event` WHERE `id`=?")
	g_SQL.PlayerDrivingSwordEventExploring.Insert = prepare(db, "INSERT INTO `player_driving_sword_event_exploring` SET `id`=?, `pid`=?, `status`=?, `garrison_start`=?, `garrison_time`=?, `award_time`=?, `role_id`=?, `cloud_id`=?, `x`=?, `y`=?, `data_id`=?")
	g_SQL.PlayerDrivingSwordEventExploring.Update = prepare(db, "UPDATE `player_driving_sword_event_exploring` SET `pid`=?, `status`=?, `garrison_start`=?, `garrison_time`=?, `award_time`=?, `role_id`=?, `cloud_id`=?, `x`=?, `y`=?, `data_id`=? WHERE `id`=?")
	g_SQL.PlayerDrivingSwordEventExploring.Delete = prepare(db, "DELETE FROM `player_driving_sword_event_exploring` WHERE `id`=?")
	g_SQL.PlayerDrivingSwordEventTreasure.Insert = prepare(db, "INSERT INTO `player_driving_sword_event_treasure` SET `id`=?, `pid`=?, `progress`=?, `cloud_id`=?, `x`=?, `y`=?, `data_id`=?")
	g_SQL.PlayerDrivingSwordEventTreasure.Update = prepare(db, "UPDATE `player_driving_sword_event_treasure` SET `pid`=?, `progress`=?, `cloud_id`=?, `x`=?, `y`=?, `data_id`=? WHERE `id`=?")
	g_SQL.PlayerDrivingSwordEventTreasure.Delete = prepare(db, "DELETE FROM `player_driving_sword_event_treasure` WHERE `id`=?")
	g_SQL.PlayerDrivingSwordEventVisiting.Insert = prepare(db, "INSERT INTO `player_driving_sword_event_visiting` SET `id`=?, `pid`=?, `status`=?, `target_pid`=?, `target_side_state`=?, `cloud_id`=?, `x`=?, `y`=?, `data_id`=?, `target_status`=?")
	g_SQL.PlayerDrivingSwordEventVisiting.Update = prepare(db, "UPDATE `player_driving_sword_event_visiting` SET `pid`=?, `status`=?, `target_pid`=?, `target_side_state`=?, `cloud_id`=?, `x`=?, `y`=?, `data_id`=?, `target_status`=? WHERE `id`=?")
	g_SQL.PlayerDrivingSwordEventVisiting.Delete = prepare(db, "DELETE FROM `player_driving_sword_event_visiting` WHERE `id`=?")
	g_SQL.PlayerDrivingSwordEventmask.Insert = prepare(db, "INSERT INTO `player_driving_sword_eventmask` SET `id`=?, `pid`=?, `cloud_id`=?, `events`=?")
	g_SQL.PlayerDrivingSwordEventmask.Update = prepare(db, "UPDATE `player_driving_sword_eventmask` SET `pid`=?, `cloud_id`=?, `events`=? WHERE `id`=?")
	g_SQL.PlayerDrivingSwordEventmask.Delete = prepare(db, "DELETE FROM `player_driving_sword_eventmask` WHERE `id`=?")
	g_SQL.PlayerDrivingSwordInfo.Insert = prepare(db, "INSERT INTO `player_driving_sword_info` SET `pid`=?, `current_cloud`=?, `highest_cloud`=?, `current_x`=?, `current_y`=?, `allowed_action`=?, `action_refresh_time`=?, `action_buy_time`=?, `daily_action_bought`=?")
	g_SQL.PlayerDrivingSwordInfo.Update = prepare(db, "UPDATE `player_driving_sword_info` SET `current_cloud`=?, `highest_cloud`=?, `current_x`=?, `current_y`=?, `allowed_action`=?, `action_refresh_time`=?, `action_buy_time`=?, `daily_action_bought`=? WHERE `pid`=?")
	g_SQL.PlayerDrivingSwordInfo.Delete = prepare(db, "DELETE FROM `player_driving_sword_info` WHERE `pid`=?")
	g_SQL.PlayerDrivingSwordMap.Insert = prepare(db, "INSERT INTO `player_driving_sword_map` SET `id`=?, `pid`=?, `cloud_id`=?, `shadows`=?, `obstacle_count`=?, `exploring_count`=?, `visiting_count`=?, `treasure_count`=?")
	g_SQL.PlayerDrivingSwordMap.Update = prepare(db, "UPDATE `player_driving_sword_map` SET `pid`=?, `cloud_id`=?, `shadows`=?, `obstacle_count`=?, `exploring_count`=?, `visiting_count`=?, `treasure_count`=? WHERE `id`=?")
	g_SQL.PlayerDrivingSwordMap.Delete = prepare(db, "DELETE FROM `player_driving_sword_map` WHERE `id`=?")
	g_SQL.PlayerEquipment.Insert = prepare(db, "INSERT INTO `player_equipment` SET `id`=?, `pid`=?, `role_id`=?, `weapon_id`=?, `clothes_id`=?, `accessories_id`=?, `shoe_id`=?")
	g_SQL.PlayerEquipment.Update = prepare(db, "UPDATE `player_equipment` SET `pid`=?, `role_id`=?, `weapon_id`=?, `clothes_id`=?, `accessories_id`=?, `shoe_id`=? WHERE `id`=?")
	g_SQL.PlayerEquipment.Delete = prepare(db, "DELETE FROM `player_equipment` WHERE `id`=?")
	g_SQL.PlayerEventAwardRecord.Insert = prepare(db, "INSERT INTO `player_event_award_record` SET `pid`=?, `record_bytes`=?, `json_event_record`=?")
	g_SQL.PlayerEventAwardRecord.Update = prepare(db, "UPDATE `player_event_award_record` SET `record_bytes`=?, `json_event_record`=? WHERE `pid`=?")
	g_SQL.PlayerEventAwardRecord.Delete = prepare(db, "DELETE FROM `player_event_award_record` WHERE `pid`=?")
	g_SQL.PlayerEventDailyOnline.Insert = prepare(db, "INSERT INTO `player_event_daily_online` SET `pid`=?, `timestamp`=?, `total_online_time`=?, `awarded_onlinetime`=?")
	g_SQL.PlayerEventDailyOnline.Update = prepare(db, "UPDATE `player_event_daily_online` SET `timestamp`=?, `total_online_time`=?, `awarded_onlinetime`=? WHERE `pid`=?")
	g_SQL.PlayerEventDailyOnline.Delete = prepare(db, "DELETE FROM `player_event_daily_online` WHERE `pid`=?")
	g_SQL.PlayerEventVnDailyRecharge.Insert = prepare(db, "INSERT INTO `player_event_vn_daily_recharge` SET `id`=?, `pid`=?, `page`=?, `timestamp`=?, `awarded`=?, `daily_recharge`=?, `total_recharge`=?")
	g_SQL.PlayerEventVnDailyRecharge.Update = prepare(db, "UPDATE `player_event_vn_daily_recharge` SET `pid`=?, `page`=?, `timestamp`=?, `awarded`=?, `daily_recharge`=?, `total_recharge`=? WHERE `id`=?")
	g_SQL.PlayerEventVnDailyRecharge.Delete = prepare(db, "DELETE FROM `player_event_vn_daily_recharge` WHERE `id`=?")
	g_SQL.PlayerEventVnDragonBallRecord.Insert = prepare(db, "INSERT INTO `player_event_vn_dragon_ball_record` SET `pid`=?, `item_id`=?, `createtime`=?")
	g_SQL.PlayerEventVnDragonBallRecord.Update = prepare(db, "UPDATE `player_event_vn_dragon_ball_record` SET `item_id`=?, `createtime`=? WHERE `pid`=?")
	g_SQL.PlayerEventVnDragonBallRecord.Delete = prepare(db, "DELETE FROM `player_event_vn_dragon_ball_record` WHERE `pid`=?")
	g_SQL.PlayerEventsFightPower.Insert = prepare(db, "INSERT INTO `player_events_fight_power` SET `pid`=?, `lock`=?")
	g_SQL.PlayerEventsFightPower.Update = prepare(db, "UPDATE `player_events_fight_power` SET `lock`=? WHERE `pid`=?")
	g_SQL.PlayerEventsFightPower.Delete = prepare(db, "DELETE FROM `player_events_fight_power` WHERE `pid`=?")
	g_SQL.PlayerEventsIngotRecord.Insert = prepare(db, "INSERT INTO `player_events_ingot_record` SET `pid`=?, `ingot_in`=?, `ingot_in_end_time`=?, `ingot_out`=?, `ingot_out_end_time`=?")
	g_SQL.PlayerEventsIngotRecord.Update = prepare(db, "UPDATE `player_events_ingot_record` SET `ingot_in`=?, `ingot_in_end_time`=?, `ingot_out`=?, `ingot_out_end_time`=? WHERE `pid`=?")
	g_SQL.PlayerEventsIngotRecord.Delete = prepare(db, "DELETE FROM `player_events_ingot_record` WHERE `pid`=?")
	g_SQL.PlayerExtendLevel.Insert = prepare(db, "INSERT INTO `player_extend_level` SET `pid`=?, `coin_pass_time`=?, `exp_pass_time`=?, `ghost_pass_time`=?, `pet_pass_time`=?, `buddy_pass_time`=?, `coin_daily_num`=?, `exp_daily_num`=?, `buddy_daily_num`=?, `pet_daily_num`=?, `ghost_daily_num`=?, `rand_buddy_role_id`=?, `buddy_pos`=?, `buddy_tactical`=?, `role_pos`=?, `exp_maxlevel`=?, `coins_maxlevel`=?, `coins_buy_num`=?, `exp_buy_num`=?, `coins_buy_time`=?, `exp_buy_time`=?")
	g_SQL.PlayerExtendLevel.Update = prepare(db, "UPDATE `player_extend_level` SET `coin_pass_time`=?, `exp_pass_time`=?, `ghost_pass_time`=?, `pet_pass_time`=?, `buddy_pass_time`=?, `coin_daily_num`=?, `exp_daily_num`=?, `buddy_daily_num`=?, `pet_daily_num`=?, `ghost_daily_num`=?, `rand_buddy_role_id`=?, `buddy_pos`=?, `buddy_tactical`=?, `role_pos`=?, `exp_maxlevel`=?, `coins_maxlevel`=?, `coins_buy_num`=?, `exp_buy_num`=?, `coins_buy_time`=?, `exp_buy_time`=? WHERE `pid`=?")
	g_SQL.PlayerExtendLevel.Delete = prepare(db, "DELETE FROM `player_extend_level` WHERE `pid`=?")
	g_SQL.PlayerExtendQuest.Insert = prepare(db, "INSERT INTO `player_extend_quest` SET `id`=?, `pid`=?, `quest_id`=?, `progress`=?, `state`=?")
	g_SQL.PlayerExtendQuest.Update = prepare(db, "UPDATE `player_extend_quest` SET `pid`=?, `quest_id`=?, `progress`=?, `state`=? WHERE `id`=?")
	g_SQL.PlayerExtendQuest.Delete = prepare(db, "DELETE FROM `player_extend_quest` WHERE `id`=?")
	g_SQL.PlayerFame.Insert = prepare(db, "INSERT INTO `player_fame` SET `pid`=?, `fame`=?, `level`=?, `arena_fame`=?, `mult_level_fame`=?")
	g_SQL.PlayerFame.Update = prepare(db, "UPDATE `player_fame` SET `fame`=?, `level`=?, `arena_fame`=?, `mult_level_fame`=? WHERE `pid`=?")
	g_SQL.PlayerFame.Delete = prepare(db, "DELETE FROM `player_fame` WHERE `pid`=?")
	g_SQL.PlayerFashion.Insert = prepare(db, "INSERT INTO `player_fashion` SET `id`=?, `pid`=?, `fashion_id`=?, `expire_time`=?")
	g_SQL.PlayerFashion.Update = prepare(db, "UPDATE `player_fashion` SET `pid`=?, `fashion_id`=?, `expire_time`=? WHERE `id`=?")
	g_SQL.PlayerFashion.Delete = prepare(db, "DELETE FROM `player_fashion` WHERE `id`=?")
	g_SQL.PlayerFashionState.Insert = prepare(db, "INSERT INTO `player_fashion_state` SET `pid`=?, `update_time`=?, `dressed_fashion_id`=?")
	g_SQL.PlayerFashionState.Update = prepare(db, "UPDATE `player_fashion_state` SET `update_time`=?, `dressed_fashion_id`=? WHERE `pid`=?")
	g_SQL.PlayerFashionState.Delete = prepare(db, "DELETE FROM `player_fashion_state` WHERE `pid`=?")
	g_SQL.PlayerFateBoxState.Insert = prepare(db, "INSERT INTO `player_fate_box_state` SET `pid`=?, `lock`=?, `star_box_free_draw_timestamp`=?, `star_box_draw_count`=?, `moon_box_free_draw_timestamp`=?, `moon_box_draw_count`=?, `sun_box_free_draw_timestamp`=?, `sun_box_draw_count`=?, `hunyuan_box_free_draw_timestamp`=?, `hunyuan_box_draw_count`=?")
	g_SQL.PlayerFateBoxState.Update = prepare(db, "UPDATE `player_fate_box_state` SET `lock`=?, `star_box_free_draw_timestamp`=?, `star_box_draw_count`=?, `moon_box_free_draw_timestamp`=?, `moon_box_draw_count`=?, `sun_box_free_draw_timestamp`=?, `sun_box_draw_count`=?, `hunyuan_box_free_draw_timestamp`=?, `hunyuan_box_draw_count`=? WHERE `pid`=?")
	g_SQL.PlayerFateBoxState.Delete = prepare(db, "DELETE FROM `player_fate_box_state` WHERE `pid`=?")
	g_SQL.PlayerFightNum.Insert = prepare(db, "INSERT INTO `player_fight_num` SET `pid`=?, `fight_num`=?")
	g_SQL.PlayerFightNum.Update = prepare(db, "UPDATE `player_fight_num` SET `fight_num`=? WHERE `pid`=?")
	g_SQL.PlayerFightNum.Delete = prepare(db, "DELETE FROM `player_fight_num` WHERE `pid`=?")
	g_SQL.PlayerFirstCharge.Insert = prepare(db, "INSERT INTO `player_first_charge` SET `id`=?, `pid`=?, `product_id`=?")
	g_SQL.PlayerFirstCharge.Update = prepare(db, "UPDATE `player_first_charge` SET `pid`=?, `product_id`=? WHERE `id`=?")
	g_SQL.PlayerFirstCharge.Delete = prepare(db, "DELETE FROM `player_first_charge` WHERE `id`=?")
	g_SQL.PlayerFormation.Insert = prepare(db, "INSERT INTO `player_formation` SET `pid`=?, `pos0`=?, `pos1`=?, `pos2`=?, `pos3`=?, `pos4`=?, `pos5`=?, `pos6`=?, `pos7`=?, `pos8`=?")
	g_SQL.PlayerFormation.Update = prepare(db, "UPDATE `player_formation` SET `pos0`=?, `pos1`=?, `pos2`=?, `pos3`=?, `pos4`=?, `pos5`=?, `pos6`=?, `pos7`=?, `pos8`=? WHERE `pid`=?")
	g_SQL.PlayerFormation.Delete = prepare(db, "DELETE FROM `player_formation` WHERE `pid`=?")
	g_SQL.PlayerFuncKey.Insert = prepare(db, "INSERT INTO `player_func_key` SET `pid`=?, `key`=?, `played_key`=?, `unique_key`=?")
	g_SQL.PlayerFuncKey.Update = prepare(db, "UPDATE `player_func_key` SET `key`=?, `played_key`=?, `unique_key`=? WHERE `pid`=?")
	g_SQL.PlayerFuncKey.Delete = prepare(db, "DELETE FROM `player_func_key` WHERE `pid`=?")
	g_SQL.PlayerGhost.Insert = prepare(db, "INSERT INTO `player_ghost` SET `id`=?, `pid`=?, `ghost_id`=?, `star`=?, `level`=?, `exp`=?, `pos`=?, `is_new`=?, `role_id`=?, `skill_level`=?, `relation_id`=?, `add_growth`=?")
	g_SQL.PlayerGhost.Update = prepare(db, "UPDATE `player_ghost` SET `pid`=?, `ghost_id`=?, `star`=?, `level`=?, `exp`=?, `pos`=?, `is_new`=?, `role_id`=?, `skill_level`=?, `relation_id`=?, `add_growth`=? WHERE `id`=?")
	g_SQL.PlayerGhost.Delete = prepare(db, "DELETE FROM `player_ghost` WHERE `id`=?")
	g_SQL.PlayerGhostEquipment.Insert = prepare(db, "INSERT INTO `player_ghost_equipment` SET `id`=?, `pid`=?, `role_id`=?, `ghost_power`=?, `pos1`=?, `pos2`=?, `pos3`=?, `pos4`=?")
	g_SQL.PlayerGhostEquipment.Update = prepare(db, "UPDATE `player_ghost_equipment` SET `pid`=?, `role_id`=?, `ghost_power`=?, `pos1`=?, `pos2`=?, `pos3`=?, `pos4`=? WHERE `id`=?")
	g_SQL.PlayerGhostEquipment.Delete = prepare(db, "DELETE FROM `player_ghost_equipment` WHERE `id`=?")
	g_SQL.PlayerGhostState.Insert = prepare(db, "INSERT INTO `player_ghost_state` SET `pid`=?, `train_by_ingot_num`=?, `train_by_ingot_time`=?, `last_flush_time`=?, `ghost_fight_num`=?")
	g_SQL.PlayerGhostState.Update = prepare(db, "UPDATE `player_ghost_state` SET `train_by_ingot_num`=?, `train_by_ingot_time`=?, `last_flush_time`=?, `ghost_fight_num`=? WHERE `pid`=?")
	g_SQL.PlayerGhostState.Delete = prepare(db, "DELETE FROM `player_ghost_state` WHERE `pid`=?")
	g_SQL.PlayerGlobalChatState.Insert = prepare(db, "INSERT INTO `player_global_chat_state` SET `pid`=?, `ban_until`=?")
	g_SQL.PlayerGlobalChatState.Update = prepare(db, "UPDATE `player_global_chat_state` SET `ban_until`=? WHERE `pid`=?")
	g_SQL.PlayerGlobalChatState.Delete = prepare(db, "DELETE FROM `player_global_chat_state` WHERE `pid`=?")
	g_SQL.PlayerGlobalCliqueBuilding.Insert = prepare(db, "INSERT INTO `player_global_clique_building` SET `pid`=?, `silver_exchange_time`=?, `gold_exchange_num`=?, `silver_exchange_num`=?, `donate_coins_center_building`=?, `donate_coins_temple_building`=?, `donate_coins_bank_building`=?, `donate_coins_health_building`=?, `donate_coins_attack_building`=?, `donate_coins_defense_building`=?, `donate_coins_store_building`=?, `health_level`=?, `attack_level`=?, `defense_level`=?, `worship_time`=?, `donate_coins_center_building_time`=?, `glod_exchange_time`=?, `worship_type`=?")
	g_SQL.PlayerGlobalCliqueBuilding.Update = prepare(db, "UPDATE `player_global_clique_building` SET `silver_exchange_time`=?, `gold_exchange_num`=?, `silver_exchange_num`=?, `donate_coins_center_building`=?, `donate_coins_temple_building`=?, `donate_coins_bank_building`=?, `donate_coins_health_building`=?, `donate_coins_attack_building`=?, `donate_coins_defense_building`=?, `donate_coins_store_building`=?, `health_level`=?, `attack_level`=?, `defense_level`=?, `worship_time`=?, `donate_coins_center_building_time`=?, `glod_exchange_time`=?, `worship_type`=? WHERE `pid`=?")
	g_SQL.PlayerGlobalCliqueBuilding.Delete = prepare(db, "DELETE FROM `player_global_clique_building` WHERE `pid`=?")
	g_SQL.PlayerGlobalCliqueBuildingQuest.Insert = prepare(db, "INSERT INTO `player_global_clique_building_quest` SET `id`=?, `pid`=?, `quest_id`=?, `award_status`=?, `building_type`=?")
	g_SQL.PlayerGlobalCliqueBuildingQuest.Update = prepare(db, "UPDATE `player_global_clique_building_quest` SET `pid`=?, `quest_id`=?, `award_status`=?, `building_type`=? WHERE `id`=?")
	g_SQL.PlayerGlobalCliqueBuildingQuest.Delete = prepare(db, "DELETE FROM `player_global_clique_building_quest` WHERE `id`=?")
	g_SQL.PlayerGlobalCliqueDailyQuest.Insert = prepare(db, "INSERT INTO `player_global_clique_daily_quest` SET `id`=?, `pid`=?, `quest_id`=?, `finish_count`=?, `last_update_time`=?, `award_status`=?, `class`=?")
	g_SQL.PlayerGlobalCliqueDailyQuest.Update = prepare(db, "UPDATE `player_global_clique_daily_quest` SET `pid`=?, `quest_id`=?, `finish_count`=?, `last_update_time`=?, `award_status`=?, `class`=? WHERE `id`=?")
	g_SQL.PlayerGlobalCliqueDailyQuest.Delete = prepare(db, "DELETE FROM `player_global_clique_daily_quest` WHERE `id`=?")
	g_SQL.PlayerGlobalCliqueEscort.Insert = prepare(db, "INSERT INTO `player_global_clique_escort` SET `pid`=?, `daily_escort_timestamp`=?, `daily_escort_num`=?, `escort_boat_type`=?, `status`=?, `hijacked`=?")
	g_SQL.PlayerGlobalCliqueEscort.Update = prepare(db, "UPDATE `player_global_clique_escort` SET `daily_escort_timestamp`=?, `daily_escort_num`=?, `escort_boat_type`=?, `status`=?, `hijacked`=? WHERE `pid`=?")
	g_SQL.PlayerGlobalCliqueEscort.Delete = prepare(db, "DELETE FROM `player_global_clique_escort` WHERE `pid`=?")
	g_SQL.PlayerGlobalCliqueEscortMessage.Insert = prepare(db, "INSERT INTO `player_global_clique_escort_message` SET `id`=?, `pid`=?, `tpl_id`=?, `parameters`=?, `timestamp`=?")
	g_SQL.PlayerGlobalCliqueEscortMessage.Update = prepare(db, "UPDATE `player_global_clique_escort_message` SET `pid`=?, `tpl_id`=?, `parameters`=?, `timestamp`=? WHERE `id`=?")
	g_SQL.PlayerGlobalCliqueEscortMessage.Delete = prepare(db, "DELETE FROM `player_global_clique_escort_message` WHERE `id`=?")
	g_SQL.PlayerGlobalCliqueHijack.Insert = prepare(db, "INSERT INTO `player_global_clique_hijack` SET `pid`=?, `daily_hijack_timestamp`=?, `daily_hijack_num`=?, `hijacked_boat_type`=?, `status`=?, `buy_timestamp`=?, `buy_num`=?")
	g_SQL.PlayerGlobalCliqueHijack.Update = prepare(db, "UPDATE `player_global_clique_hijack` SET `daily_hijack_timestamp`=?, `daily_hijack_num`=?, `hijacked_boat_type`=?, `status`=?, `buy_timestamp`=?, `buy_num`=? WHERE `pid`=?")
	g_SQL.PlayerGlobalCliqueHijack.Delete = prepare(db, "DELETE FROM `player_global_clique_hijack` WHERE `pid`=?")
	g_SQL.PlayerGlobalCliqueInfo.Insert = prepare(db, "INSERT INTO `player_global_clique_info` SET `pid`=?, `clique_id`=?, `join_time`=?, `contrib`=?, `contrib_clear_time`=?, `donate_coins_time`=?, `daily_donate_coins`=?, `total_contrib`=?")
	g_SQL.PlayerGlobalCliqueInfo.Update = prepare(db, "UPDATE `player_global_clique_info` SET `clique_id`=?, `join_time`=?, `contrib`=?, `contrib_clear_time`=?, `donate_coins_time`=?, `daily_donate_coins`=?, `total_contrib`=? WHERE `pid`=?")
	g_SQL.PlayerGlobalCliqueInfo.Delete = prepare(db, "DELETE FROM `player_global_clique_info` WHERE `pid`=?")
	g_SQL.PlayerGlobalCliqueKongfu.Insert = prepare(db, "INSERT INTO `player_global_clique_kongfu` SET `id`=?, `pid`=?, `kongfu_id`=?, `level`=?")
	g_SQL.PlayerGlobalCliqueKongfu.Update = prepare(db, "UPDATE `player_global_clique_kongfu` SET `pid`=?, `kongfu_id`=?, `level`=? WHERE `id`=?")
	g_SQL.PlayerGlobalCliqueKongfu.Delete = prepare(db, "DELETE FROM `player_global_clique_kongfu` WHERE `id`=?")
	g_SQL.PlayerGlobalFriend.Insert = prepare(db, "INSERT INTO `player_global_friend` SET `id`=?, `pid`=?, `friend_pid`=?, `friend_nick`=?, `friend_role_id`=?, `friend_mode`=?, `last_chat_time`=?, `friend_time`=?, `send_heart_time`=?, `block_mode`=?")
	g_SQL.PlayerGlobalFriend.Update = prepare(db, "UPDATE `player_global_friend` SET `pid`=?, `friend_pid`=?, `friend_nick`=?, `friend_role_id`=?, `friend_mode`=?, `last_chat_time`=?, `friend_time`=?, `send_heart_time`=?, `block_mode`=? WHERE `id`=?")
	g_SQL.PlayerGlobalFriend.Delete = prepare(db, "DELETE FROM `player_global_friend` WHERE `id`=?")
	g_SQL.PlayerGlobalFriendChat.Insert = prepare(db, "INSERT INTO `player_global_friend_chat` SET `id`=?, `pid`=?, `friend_pid`=?, `mode`=?, `send_time`=?, `message`=?")
	g_SQL.PlayerGlobalFriendChat.Update = prepare(db, "UPDATE `player_global_friend_chat` SET `pid`=?, `friend_pid`=?, `mode`=?, `send_time`=?, `message`=? WHERE `id`=?")
	g_SQL.PlayerGlobalFriendChat.Delete = prepare(db, "DELETE FROM `player_global_friend_chat` WHERE `id`=?")
	g_SQL.PlayerGlobalFriendState.Insert = prepare(db, "INSERT INTO `player_global_friend_state` SET `pid`=?, `delete_day_count`=?, `delete_time`=?, `exist_offline_chat`=?, `platform_friend_num`=?")
	g_SQL.PlayerGlobalFriendState.Update = prepare(db, "UPDATE `player_global_friend_state` SET `delete_day_count`=?, `delete_time`=?, `exist_offline_chat`=?, `platform_friend_num`=? WHERE `pid`=?")
	g_SQL.PlayerGlobalFriendState.Delete = prepare(db, "DELETE FROM `player_global_friend_state` WHERE `pid`=?")
	g_SQL.PlayerGlobalGiftCardRecord.Insert = prepare(db, "INSERT INTO `player_global_gift_card_record` SET `id`=?, `pid`=?, `code`=?")
	g_SQL.PlayerGlobalGiftCardRecord.Update = prepare(db, "UPDATE `player_global_gift_card_record` SET `pid`=?, `code`=? WHERE `id`=?")
	g_SQL.PlayerGlobalGiftCardRecord.Delete = prepare(db, "DELETE FROM `player_global_gift_card_record` WHERE `id`=?")
	g_SQL.PlayerGlobalMailState.Insert = prepare(db, "INSERT INTO `player_global_mail_state` SET `pid`=?, `max_timestamp`=?")
	g_SQL.PlayerGlobalMailState.Update = prepare(db, "UPDATE `player_global_mail_state` SET `max_timestamp`=? WHERE `pid`=?")
	g_SQL.PlayerGlobalMailState.Delete = prepare(db, "DELETE FROM `player_global_mail_state` WHERE `pid`=?")
	g_SQL.PlayerGlobalWorldChatState.Insert = prepare(db, "INSERT INTO `player_global_world_chat_state` SET `pid`=?, `timestamp`=?, `daily_num`=?")
	g_SQL.PlayerGlobalWorldChatState.Update = prepare(db, "UPDATE `player_global_world_chat_state` SET `timestamp`=?, `daily_num`=? WHERE `pid`=?")
	g_SQL.PlayerGlobalWorldChatState.Delete = prepare(db, "DELETE FROM `player_global_world_chat_state` WHERE `pid`=?")
	g_SQL.PlayerHardLevel.Insert = prepare(db, "INSERT INTO `player_hard_level` SET `pid`=?, `lock`=?, `award_lock`=?")
	g_SQL.PlayerHardLevel.Update = prepare(db, "UPDATE `player_hard_level` SET `lock`=?, `award_lock`=? WHERE `pid`=?")
	g_SQL.PlayerHardLevel.Delete = prepare(db, "DELETE FROM `player_hard_level` WHERE `pid`=?")
	g_SQL.PlayerHardLevelRecord.Insert = prepare(db, "INSERT INTO `player_hard_level_record` SET `id`=?, `pid`=?, `level_id`=?, `open_time`=?, `score`=?, `round`=?, `daily_num`=?, `last_enter_time`=?, `buy_times`=?, `buy_update_time`=?")
	g_SQL.PlayerHardLevelRecord.Update = prepare(db, "UPDATE `player_hard_level_record` SET `pid`=?, `level_id`=?, `open_time`=?, `score`=?, `round`=?, `daily_num`=?, `last_enter_time`=?, `buy_times`=?, `buy_update_time`=? WHERE `id`=?")
	g_SQL.PlayerHardLevelRecord.Delete = prepare(db, "DELETE FROM `player_hard_level_record` WHERE `id`=?")
	g_SQL.PlayerHeart.Insert = prepare(db, "INSERT INTO `player_heart` SET `pid`=?, `value`=?, `update_time`=?, `add_day_count`=?, `add_time`=?, `recover_day_count`=?")
	g_SQL.PlayerHeart.Update = prepare(db, "UPDATE `player_heart` SET `value`=?, `update_time`=?, `add_day_count`=?, `add_time`=?, `recover_day_count`=? WHERE `pid`=?")
	g_SQL.PlayerHeart.Delete = prepare(db, "DELETE FROM `player_heart` WHERE `pid`=?")
	g_SQL.PlayerHeartDraw.Insert = prepare(db, "INSERT INTO `player_heart_draw` SET `id`=?, `pid`=?, `draw_type`=?, `daily_num`=?, `draw_time`=?")
	g_SQL.PlayerHeartDraw.Update = prepare(db, "UPDATE `player_heart_draw` SET `pid`=?, `draw_type`=?, `daily_num`=?, `draw_time`=? WHERE `id`=?")
	g_SQL.PlayerHeartDraw.Delete = prepare(db, "DELETE FROM `player_heart_draw` WHERE `id`=?")
	g_SQL.PlayerHeartDrawCardRecord.Insert = prepare(db, "INSERT INTO `player_heart_draw_card_record` SET `id`=?, `pid`=?, `award_type`=?, `award_num`=?, `item_id`=?, `draw_time`=?")
	g_SQL.PlayerHeartDrawCardRecord.Update = prepare(db, "UPDATE `player_heart_draw_card_record` SET `pid`=?, `award_type`=?, `award_num`=?, `item_id`=?, `draw_time`=? WHERE `id`=?")
	g_SQL.PlayerHeartDrawCardRecord.Delete = prepare(db, "DELETE FROM `player_heart_draw_card_record` WHERE `id`=?")
	g_SQL.PlayerHeartDrawWheelRecord.Insert = prepare(db, "INSERT INTO `player_heart_draw_wheel_record` SET `id`=?, `pid`=?, `award_type`=?, `award_num`=?, `item_id`=?, `draw_time`=?")
	g_SQL.PlayerHeartDrawWheelRecord.Update = prepare(db, "UPDATE `player_heart_draw_wheel_record` SET `pid`=?, `award_type`=?, `award_num`=?, `item_id`=?, `draw_time`=? WHERE `id`=?")
	g_SQL.PlayerHeartDrawWheelRecord.Delete = prepare(db, "DELETE FROM `player_heart_draw_wheel_record` WHERE `id`=?")
	g_SQL.PlayerInfo.Insert = prepare(db, "INSERT INTO `player_info` SET `pid`=?, `ingot`=?, `coins`=?, `new_mail_num`=?, `last_login_time`=?, `last_offline_time`=?, `total_online_time`=?, `first_login_time`=?, `new_arena_report_num`=?, `last_skill_flush`=?, `sword_soul_fragment`=?")
	g_SQL.PlayerInfo.Update = prepare(db, "UPDATE `player_info` SET `ingot`=?, `coins`=?, `new_mail_num`=?, `last_login_time`=?, `last_offline_time`=?, `total_online_time`=?, `first_login_time`=?, `new_arena_report_num`=?, `last_skill_flush`=?, `sword_soul_fragment`=? WHERE `pid`=?")
	g_SQL.PlayerInfo.Delete = prepare(db, "DELETE FROM `player_info` WHERE `pid`=?")
	g_SQL.PlayerIsOperated.Insert = prepare(db, "INSERT INTO `player_is_operated` SET `pid`=?, `operate_value`=?")
	g_SQL.PlayerIsOperated.Update = prepare(db, "UPDATE `player_is_operated` SET `operate_value`=? WHERE `pid`=?")
	g_SQL.PlayerIsOperated.Delete = prepare(db, "DELETE FROM `player_is_operated` WHERE `pid`=?")
	g_SQL.PlayerItem.Insert = prepare(db, "INSERT INTO `player_item` SET `id`=?, `pid`=?, `item_id`=?, `num`=?, `is_dressed`=?, `buy_back_state`=?, `appendix_id`=?, `refine_level_bak`=?, `refine_fail_times`=?, `refine_level`=?, `price`=?")
	g_SQL.PlayerItem.Update = prepare(db, "UPDATE `player_item` SET `pid`=?, `item_id`=?, `num`=?, `is_dressed`=?, `buy_back_state`=?, `appendix_id`=?, `refine_level_bak`=?, `refine_fail_times`=?, `refine_level`=?, `price`=? WHERE `id`=?")
	g_SQL.PlayerItem.Delete = prepare(db, "DELETE FROM `player_item` WHERE `id`=?")
	g_SQL.PlayerItemAppendix.Insert = prepare(db, "INSERT INTO `player_item_appendix` SET `id`=?, `pid`=?, `health`=?, `cultivation`=?, `speed`=?, `attack`=?, `defence`=?, `dodge_level`=?, `hit_level`=?, `block_level`=?, `critical_level`=?, `tenacity_level`=?, `destroy_level`=?, `recast_attr`=?")
	g_SQL.PlayerItemAppendix.Update = prepare(db, "UPDATE `player_item_appendix` SET `pid`=?, `health`=?, `cultivation`=?, `speed`=?, `attack`=?, `defence`=?, `dodge_level`=?, `hit_level`=?, `block_level`=?, `critical_level`=?, `tenacity_level`=?, `destroy_level`=?, `recast_attr`=? WHERE `id`=?")
	g_SQL.PlayerItemAppendix.Delete = prepare(db, "DELETE FROM `player_item_appendix` WHERE `id`=?")
	g_SQL.PlayerItemBuyback.Insert = prepare(db, "INSERT INTO `player_item_buyback` SET `pid`=?, `back_id1`=?, `back_id2`=?, `back_id3`=?, `back_id4`=?, `back_id5`=?, `back_id6`=?")
	g_SQL.PlayerItemBuyback.Update = prepare(db, "UPDATE `player_item_buyback` SET `back_id1`=?, `back_id2`=?, `back_id3`=?, `back_id4`=?, `back_id5`=?, `back_id6`=? WHERE `pid`=?")
	g_SQL.PlayerItemBuyback.Delete = prepare(db, "DELETE FROM `player_item_buyback` WHERE `pid`=?")
	g_SQL.PlayerItemRecastState.Insert = prepare(db, "INSERT INTO `player_item_recast_state` SET `pid`=?, `player_item_id`=?, `selected_attr`=?, `attr1_type`=?, `attr1_value`=?, `attr2_type`=?, `attr2_value`=?, `attr3_type`=?, `attr3_value`=?")
	g_SQL.PlayerItemRecastState.Update = prepare(db, "UPDATE `player_item_recast_state` SET `player_item_id`=?, `selected_attr`=?, `attr1_type`=?, `attr1_value`=?, `attr2_type`=?, `attr2_value`=?, `attr3_type`=?, `attr3_value`=? WHERE `pid`=?")
	g_SQL.PlayerItemRecastState.Delete = prepare(db, "DELETE FROM `player_item_recast_state` WHERE `pid`=?")
	g_SQL.PlayerLevelAwardInfo.Insert = prepare(db, "INSERT INTO `player_level_award_info` SET `pid`=?, `awarded`=?")
	g_SQL.PlayerLevelAwardInfo.Update = prepare(db, "UPDATE `player_level_award_info` SET `awarded`=? WHERE `pid`=?")
	g_SQL.PlayerLevelAwardInfo.Delete = prepare(db, "DELETE FROM `player_level_award_info` WHERE `pid`=?")
	g_SQL.PlayerLoginAwardRecord.Insert = prepare(db, "INSERT INTO `player_login_award_record` SET `pid`=?, `active_days`=?, `record`=?, `update_timestamp`=?")
	g_SQL.PlayerLoginAwardRecord.Update = prepare(db, "UPDATE `player_login_award_record` SET `active_days`=?, `record`=?, `update_timestamp`=? WHERE `pid`=?")
	g_SQL.PlayerLoginAwardRecord.Delete = prepare(db, "DELETE FROM `player_login_award_record` WHERE `pid`=?")
	g_SQL.PlayerMail.Insert = prepare(db, "INSERT INTO `player_mail` SET `id`=?, `pid`=?, `mail_id`=?, `state`=?, `send_time`=?, `parameters`=?, `have_attachment`=?, `title`=?, `content`=?, `expire_time`=?, `priority`=?")
	g_SQL.PlayerMail.Update = prepare(db, "UPDATE `player_mail` SET `pid`=?, `mail_id`=?, `state`=?, `send_time`=?, `parameters`=?, `have_attachment`=?, `title`=?, `content`=?, `expire_time`=?, `priority`=? WHERE `id`=?")
	g_SQL.PlayerMail.Delete = prepare(db, "DELETE FROM `player_mail` WHERE `id`=?")
	g_SQL.PlayerMailAttachment.Insert = prepare(db, "INSERT INTO `player_mail_attachment` SET `id`=?, `pid`=?, `player_mail_id`=?, `attachment_type`=?, `item_id`=?, `item_num`=?")
	g_SQL.PlayerMailAttachment.Update = prepare(db, "UPDATE `player_mail_attachment` SET `pid`=?, `player_mail_id`=?, `attachment_type`=?, `item_id`=?, `item_num`=? WHERE `id`=?")
	g_SQL.PlayerMailAttachment.Delete = prepare(db, "DELETE FROM `player_mail_attachment` WHERE `id`=?")
	g_SQL.PlayerMailAttachmentLg.Insert = prepare(db, "INSERT INTO `player_mail_attachment_lg` SET `id`=?, `pid`=?, `player_mail_id`=?, `attachment_type`=?, `item_id`=?, `item_num`=?, `take_timestamp`=?")
	g_SQL.PlayerMailAttachmentLg.Update = prepare(db, "UPDATE `player_mail_attachment_lg` SET `pid`=?, `player_mail_id`=?, `attachment_type`=?, `item_id`=?, `item_num`=?, `take_timestamp`=? WHERE `id`=?")
	g_SQL.PlayerMailAttachmentLg.Delete = prepare(db, "DELETE FROM `player_mail_attachment_lg` WHERE `id`=?")
	g_SQL.PlayerMailLg.Insert = prepare(db, "INSERT INTO `player_mail_lg` SET `id`=?, `pmid`=?, `pid`=?, `mail_id`=?, `state`=?, `send_time`=?, `parameters`=?, `have_attachment`=?, `title`=?, `content`=?")
	g_SQL.PlayerMailLg.Update = prepare(db, "UPDATE `player_mail_lg` SET `pmid`=?, `pid`=?, `mail_id`=?, `state`=?, `send_time`=?, `parameters`=?, `have_attachment`=?, `title`=?, `content`=? WHERE `id`=?")
	g_SQL.PlayerMailLg.Delete = prepare(db, "DELETE FROM `player_mail_lg` WHERE `id`=?")
	g_SQL.PlayerMeditationState.Insert = prepare(db, "INSERT INTO `player_meditation_state` SET `pid`=?, `accumulate_time`=?, `start_timestamp`=?")
	g_SQL.PlayerMeditationState.Update = prepare(db, "UPDATE `player_meditation_state` SET `accumulate_time`=?, `start_timestamp`=? WHERE `pid`=?")
	g_SQL.PlayerMeditationState.Delete = prepare(db, "DELETE FROM `player_meditation_state` WHERE `pid`=?")
	g_SQL.PlayerMission.Insert = prepare(db, "INSERT INTO `player_mission` SET `pid`=?, `key`=?, `max_order`=?")
	g_SQL.PlayerMission.Update = prepare(db, "UPDATE `player_mission` SET `key`=?, `max_order`=? WHERE `pid`=?")
	g_SQL.PlayerMission.Delete = prepare(db, "DELETE FROM `player_mission` WHERE `pid`=?")
	g_SQL.PlayerMissionLevel.Insert = prepare(db, "INSERT INTO `player_mission_level` SET `pid`=?, `lock`=?, `max_lock`=?, `award_lock`=?")
	g_SQL.PlayerMissionLevel.Update = prepare(db, "UPDATE `player_mission_level` SET `lock`=?, `max_lock`=?, `award_lock`=? WHERE `pid`=?")
	g_SQL.PlayerMissionLevel.Delete = prepare(db, "DELETE FROM `player_mission_level` WHERE `pid`=?")
	g_SQL.PlayerMissionLevelRecord.Insert = prepare(db, "INSERT INTO `player_mission_level_record` SET `id`=?, `pid`=?, `mission_id`=?, `mission_level_id`=?, `open_time`=?, `score`=?, `round`=?, `daily_num`=?, `last_enter_time`=?, `empty_shadow_bits`=?")
	g_SQL.PlayerMissionLevelRecord.Update = prepare(db, "UPDATE `player_mission_level_record` SET `pid`=?, `mission_id`=?, `mission_level_id`=?, `open_time`=?, `score`=?, `round`=?, `daily_num`=?, `last_enter_time`=?, `empty_shadow_bits`=? WHERE `id`=?")
	g_SQL.PlayerMissionLevelRecord.Delete = prepare(db, "DELETE FROM `player_mission_level_record` WHERE `id`=?")
	g_SQL.PlayerMissionLevelStateBin.Insert = prepare(db, "INSERT INTO `player_mission_level_state_bin` SET `pid`=?, `bin`=?")
	g_SQL.PlayerMissionLevelStateBin.Update = prepare(db, "UPDATE `player_mission_level_state_bin` SET `bin`=? WHERE `pid`=?")
	g_SQL.PlayerMissionLevelStateBin.Delete = prepare(db, "DELETE FROM `player_mission_level_state_bin` WHERE `pid`=?")
	g_SQL.PlayerMissionRecord.Insert = prepare(db, "INSERT INTO `player_mission_record` SET `id`=?, `pid`=?, `town_id`=?, `mission_id`=?, `open_time`=?")
	g_SQL.PlayerMissionRecord.Update = prepare(db, "UPDATE `player_mission_record` SET `pid`=?, `town_id`=?, `mission_id`=?, `open_time`=? WHERE `id`=?")
	g_SQL.PlayerMissionRecord.Delete = prepare(db, "DELETE FROM `player_mission_record` WHERE `id`=?")
	g_SQL.PlayerMissionStarAward.Insert = prepare(db, "INSERT INTO `player_mission_star_award` SET `id`=?, `pid`=?, `town_id`=?, `box_type`=?")
	g_SQL.PlayerMissionStarAward.Update = prepare(db, "UPDATE `player_mission_star_award` SET `pid`=?, `town_id`=?, `box_type`=? WHERE `id`=?")
	g_SQL.PlayerMissionStarAward.Delete = prepare(db, "DELETE FROM `player_mission_star_award` WHERE `id`=?")
	g_SQL.PlayerMoneyTree.Insert = prepare(db, "INSERT INTO `player_money_tree` SET `pid`=?, `total`=?, `waved_times_total`=?, `waved_times`=?, `last_waved_time`=?")
	g_SQL.PlayerMoneyTree.Update = prepare(db, "UPDATE `player_money_tree` SET `total`=?, `waved_times_total`=?, `waved_times`=?, `last_waved_time`=? WHERE `pid`=?")
	g_SQL.PlayerMoneyTree.Delete = prepare(db, "DELETE FROM `player_money_tree` WHERE `pid`=?")
	g_SQL.PlayerMonthCardAwardRecord.Insert = prepare(db, "INSERT INTO `player_month_card_award_record` SET `pid`=?, `last_update`=?")
	g_SQL.PlayerMonthCardAwardRecord.Update = prepare(db, "UPDATE `player_month_card_award_record` SET `last_update`=? WHERE `pid`=?")
	g_SQL.PlayerMonthCardAwardRecord.Delete = prepare(db, "DELETE FROM `player_month_card_award_record` WHERE `pid`=?")
	g_SQL.PlayerMonthCardInfo.Insert = prepare(db, "INSERT INTO `player_month_card_info` SET `pid`=?, `starttime`=?, `endtime`=?, `buytimes`=?, `presenttotal`=?")
	g_SQL.PlayerMonthCardInfo.Update = prepare(db, "UPDATE `player_month_card_info` SET `starttime`=?, `endtime`=?, `buytimes`=?, `presenttotal`=? WHERE `pid`=?")
	g_SQL.PlayerMonthCardInfo.Delete = prepare(db, "DELETE FROM `player_month_card_info` WHERE `pid`=?")
	g_SQL.PlayerMultiLevelInfo.Insert = prepare(db, "INSERT INTO `player_multi_level_info` SET `pid`=?, `buddy_role_id`=?, `buddy_row`=?, `tactical_grid`=?, `daily_num`=?, `battle_time`=?, `lock`=?")
	g_SQL.PlayerMultiLevelInfo.Update = prepare(db, "UPDATE `player_multi_level_info` SET `buddy_role_id`=?, `buddy_row`=?, `tactical_grid`=?, `daily_num`=?, `battle_time`=?, `lock`=? WHERE `pid`=?")
	g_SQL.PlayerMultiLevelInfo.Delete = prepare(db, "DELETE FROM `player_multi_level_info` WHERE `pid`=?")
	g_SQL.PlayerNewYearConsumeRecord.Insert = prepare(db, "INSERT INTO `player_new_year_consume_record` SET `pid`=?, `consume_record`=?")
	g_SQL.PlayerNewYearConsumeRecord.Update = prepare(db, "UPDATE `player_new_year_consume_record` SET `consume_record`=? WHERE `pid`=?")
	g_SQL.PlayerNewYearConsumeRecord.Delete = prepare(db, "DELETE FROM `player_new_year_consume_record` WHERE `pid`=?")
	g_SQL.PlayerNpcTalkRecord.Insert = prepare(db, "INSERT INTO `player_npc_talk_record` SET `id`=?, `pid`=?, `npc_id`=?, `town_id`=?, `quest_id`=?")
	g_SQL.PlayerNpcTalkRecord.Update = prepare(db, "UPDATE `player_npc_talk_record` SET `pid`=?, `npc_id`=?, `town_id`=?, `quest_id`=? WHERE `id`=?")
	g_SQL.PlayerNpcTalkRecord.Delete = prepare(db, "DELETE FROM `player_npc_talk_record` WHERE `id`=?")
	g_SQL.PlayerOpenedTownTreasure.Insert = prepare(db, "INSERT INTO `player_opened_town_treasure` SET `id`=?, `pid`=?, `town_id`=?")
	g_SQL.PlayerOpenedTownTreasure.Update = prepare(db, "UPDATE `player_opened_town_treasure` SET `pid`=?, `town_id`=? WHERE `id`=?")
	g_SQL.PlayerOpenedTownTreasure.Delete = prepare(db, "DELETE FROM `player_opened_town_treasure` WHERE `id`=?")
	g_SQL.PlayerPhysical.Insert = prepare(db, "INSERT INTO `player_physical` SET `pid`=?, `value`=?, `update_time`=?, `buy_count`=?, `buy_update_time`=?, `daily_count`=?")
	g_SQL.PlayerPhysical.Update = prepare(db, "UPDATE `player_physical` SET `value`=?, `update_time`=?, `buy_count`=?, `buy_update_time`=?, `daily_count`=? WHERE `pid`=?")
	g_SQL.PlayerPhysical.Delete = prepare(db, "DELETE FROM `player_physical` WHERE `pid`=?")
	g_SQL.PlayerPurchaseRecord.Insert = prepare(db, "INSERT INTO `player_purchase_record` SET `id`=?, `pid`=?, `item_id`=?, `num`=?, `timestamp`=?")
	g_SQL.PlayerPurchaseRecord.Update = prepare(db, "UPDATE `player_purchase_record` SET `pid`=?, `item_id`=?, `num`=?, `timestamp`=? WHERE `id`=?")
	g_SQL.PlayerPurchaseRecord.Delete = prepare(db, "DELETE FROM `player_purchase_record` WHERE `id`=?")
	g_SQL.PlayerPushNotifySwitch.Insert = prepare(db, "INSERT INTO `player_push_notify_switch` SET `pid`=?, `options`=?")
	g_SQL.PlayerPushNotifySwitch.Update = prepare(db, "UPDATE `player_push_notify_switch` SET `options`=? WHERE `pid`=?")
	g_SQL.PlayerPushNotifySwitch.Delete = prepare(db, "DELETE FROM `player_push_notify_switch` WHERE `pid`=?")
	g_SQL.PlayerPveState.Insert = prepare(db, "INSERT INTO `player_pve_state` SET `pid`=?, `max_passed_floor`=?, `max_awarded_floor`=?, `unpassed_floor_enemy_num`=?, `enter_time`=?, `daily_num`=?")
	g_SQL.PlayerPveState.Update = prepare(db, "UPDATE `player_pve_state` SET `max_passed_floor`=?, `max_awarded_floor`=?, `unpassed_floor_enemy_num`=?, `enter_time`=?, `daily_num`=? WHERE `pid`=?")
	g_SQL.PlayerPveState.Delete = prepare(db, "DELETE FROM `player_pve_state` WHERE `pid`=?")
	g_SQL.PlayerQqVipGift.Insert = prepare(db, "INSERT INTO `player_qq_vip_gift` SET `pid`=?, `qqvip`=?, `surper`=?")
	g_SQL.PlayerQqVipGift.Update = prepare(db, "UPDATE `player_qq_vip_gift` SET `qqvip`=?, `surper`=? WHERE `pid`=?")
	g_SQL.PlayerQqVipGift.Delete = prepare(db, "DELETE FROM `player_qq_vip_gift` WHERE `pid`=?")
	g_SQL.PlayerQuest.Insert = prepare(db, "INSERT INTO `player_quest` SET `pid`=?, `quest_id`=?, `state`=?")
	g_SQL.PlayerQuest.Update = prepare(db, "UPDATE `player_quest` SET `quest_id`=?, `state`=? WHERE `pid`=?")
	g_SQL.PlayerQuest.Delete = prepare(db, "DELETE FROM `player_quest` WHERE `pid`=?")
	g_SQL.PlayerRainbowLevel.Insert = prepare(db, "INSERT INTO `player_rainbow_level` SET `pid`=?, `reset_timestamp`=?, `reset_num`=?, `segment`=?, `order`=?, `status`=?, `best_segment`=?, `best_order`=?, `best_record_timestamp`=?, `max_open_segment`=?, `max_pass_segment`=?, `auto_fight_num`=?, `buy_times`=?, `auto_fight_time`=?, `buy_timestamp`=?")
	g_SQL.PlayerRainbowLevel.Update = prepare(db, "UPDATE `player_rainbow_level` SET `reset_timestamp`=?, `reset_num`=?, `segment`=?, `order`=?, `status`=?, `best_segment`=?, `best_order`=?, `best_record_timestamp`=?, `max_open_segment`=?, `max_pass_segment`=?, `auto_fight_num`=?, `buy_times`=?, `auto_fight_time`=?, `buy_timestamp`=? WHERE `pid`=?")
	g_SQL.PlayerRainbowLevel.Delete = prepare(db, "DELETE FROM `player_rainbow_level` WHERE `pid`=?")
	g_SQL.PlayerRainbowLevelStateBin.Insert = prepare(db, "INSERT INTO `player_rainbow_level_state_bin` SET `pid`=?, `bin`=?")
	g_SQL.PlayerRainbowLevelStateBin.Update = prepare(db, "UPDATE `player_rainbow_level_state_bin` SET `bin`=? WHERE `pid`=?")
	g_SQL.PlayerRainbowLevelStateBin.Delete = prepare(db, "DELETE FROM `player_rainbow_level_state_bin` WHERE `pid`=?")
	g_SQL.PlayerRole.Insert = prepare(db, "INSERT INTO `player_role` SET `id`=?, `pid`=?, `role_id`=?, `level`=?, `exp`=?, `friendship_level`=?, `status`=?, `coop_id`=?, `timestamp`=?, `fight_num`=?")
	g_SQL.PlayerRole.Update = prepare(db, "UPDATE `player_role` SET `pid`=?, `role_id`=?, `level`=?, `exp`=?, `friendship_level`=?, `status`=?, `coop_id`=?, `timestamp`=?, `fight_num`=? WHERE `id`=?")
	g_SQL.PlayerRole.Delete = prepare(db, "DELETE FROM `player_role` WHERE `id`=?")
	g_SQL.PlayerSealedbook.Insert = prepare(db, "INSERT INTO `player_sealedbook` SET `pid`=?, `sealed_record`=?")
	g_SQL.PlayerSealedbook.Update = prepare(db, "UPDATE `player_sealedbook` SET `sealed_record`=? WHERE `pid`=?")
	g_SQL.PlayerSealedbook.Delete = prepare(db, "DELETE FROM `player_sealedbook` WHERE `pid`=?")
	g_SQL.PlayerSendHeartRecord.Insert = prepare(db, "INSERT INTO `player_send_heart_record` SET `id`=?, `pid`=?, `friend_pid`=?, `send_heart_time`=?")
	g_SQL.PlayerSendHeartRecord.Update = prepare(db, "UPDATE `player_send_heart_record` SET `pid`=?, `friend_pid`=?, `send_heart_time`=? WHERE `id`=?")
	g_SQL.PlayerSendHeartRecord.Delete = prepare(db, "DELETE FROM `player_send_heart_record` WHERE `id`=?")
	g_SQL.PlayerSkill.Insert = prepare(db, "INSERT INTO `player_skill` SET `id`=?, `pid`=?, `role_id`=?, `skill_id`=?, `skill_trnlv`=?")
	g_SQL.PlayerSkill.Update = prepare(db, "UPDATE `player_skill` SET `pid`=?, `role_id`=?, `skill_id`=?, `skill_trnlv`=? WHERE `id`=?")
	g_SQL.PlayerSkill.Delete = prepare(db, "DELETE FROM `player_skill` WHERE `id`=?")
	g_SQL.PlayerState.Insert = prepare(db, "INSERT INTO `player_state` SET `pid`=?, `ban_start_time`=?, `ban_end_time`=?")
	g_SQL.PlayerState.Update = prepare(db, "UPDATE `player_state` SET `ban_start_time`=?, `ban_end_time`=? WHERE `pid`=?")
	g_SQL.PlayerState.Delete = prepare(db, "DELETE FROM `player_state` WHERE `pid`=?")
	g_SQL.PlayerSwordSoul.Insert = prepare(db, "INSERT INTO `player_sword_soul` SET `id`=?, `pid`=?, `pos`=?, `sword_soul_id`=?, `exp`=?, `level`=?")
	g_SQL.PlayerSwordSoul.Update = prepare(db, "UPDATE `player_sword_soul` SET `pid`=?, `pos`=?, `sword_soul_id`=?, `exp`=?, `level`=? WHERE `id`=?")
	g_SQL.PlayerSwordSoul.Delete = prepare(db, "DELETE FROM `player_sword_soul` WHERE `id`=?")
	g_SQL.PlayerSwordSoulEquipment.Insert = prepare(db, "INSERT INTO `player_sword_soul_equipment` SET `id`=?, `pid`=?, `role_id`=?, `is_equipped_xuanyuan`=?, `type_bits`=?, `pos1`=?, `pos2`=?, `pos3`=?, `pos4`=?, `pos5`=?, `pos6`=?, `pos7`=?, `pos8`=?, `pos9`=?")
	g_SQL.PlayerSwordSoulEquipment.Update = prepare(db, "UPDATE `player_sword_soul_equipment` SET `pid`=?, `role_id`=?, `is_equipped_xuanyuan`=?, `type_bits`=?, `pos1`=?, `pos2`=?, `pos3`=?, `pos4`=?, `pos5`=?, `pos6`=?, `pos7`=?, `pos8`=?, `pos9`=? WHERE `id`=?")
	g_SQL.PlayerSwordSoulEquipment.Delete = prepare(db, "DELETE FROM `player_sword_soul_equipment` WHERE `id`=?")
	g_SQL.PlayerSwordSoulState.Insert = prepare(db, "INSERT INTO `player_sword_soul_state` SET `pid`=?, `box_state`=?, `num`=?, `update_time`=?, `ingot_num`=?, `supersoul_additional_possible`=?, `ingot_yello_one`=?, `last_ingot_draw_time`=?, `sword_soul_num`=?")
	g_SQL.PlayerSwordSoulState.Update = prepare(db, "UPDATE `player_sword_soul_state` SET `box_state`=?, `num`=?, `update_time`=?, `ingot_num`=?, `supersoul_additional_possible`=?, `ingot_yello_one`=?, `last_ingot_draw_time`=?, `sword_soul_num`=? WHERE `pid`=?")
	g_SQL.PlayerSwordSoulState.Delete = prepare(db, "DELETE FROM `player_sword_soul_state` WHERE `pid`=?")
	g_SQL.PlayerTaoyuanBless.Insert = prepare(db, "INSERT INTO `player_taoyuan_bless` SET `pid`=?, `daily_bless_times`=?, `last_bless_time`=?, `daily_be_bless_times`=?, `last_be_bless_time`=?, `bless_pid1`=?, `bless_pid2`=?, `bless_pid3`=?, `bless_pid4`=?, `bless_pid5`=?")
	g_SQL.PlayerTaoyuanBless.Update = prepare(db, "UPDATE `player_taoyuan_bless` SET `daily_bless_times`=?, `last_bless_time`=?, `daily_be_bless_times`=?, `last_be_bless_time`=?, `bless_pid1`=?, `bless_pid2`=?, `bless_pid3`=?, `bless_pid4`=?, `bless_pid5`=? WHERE `pid`=?")
	g_SQL.PlayerTaoyuanBless.Delete = prepare(db, "DELETE FROM `player_taoyuan_bless` WHERE `pid`=?")
	g_SQL.PlayerTaoyuanFileds.Insert = prepare(db, "INSERT INTO `player_taoyuan_fileds` SET `id`=?, `pid`=?, `filed_id`=?, `filed_status`=?, `plant_id`=?, `grow_time`=?")
	g_SQL.PlayerTaoyuanFileds.Update = prepare(db, "UPDATE `player_taoyuan_fileds` SET `pid`=?, `filed_id`=?, `filed_status`=?, `plant_id`=?, `grow_time`=? WHERE `id`=?")
	g_SQL.PlayerTaoyuanFileds.Delete = prepare(db, "DELETE FROM `player_taoyuan_fileds` WHERE `id`=?")
	g_SQL.PlayerTaoyuanHeart.Insert = prepare(db, "INSERT INTO `player_taoyuan_heart` SET `pid`=?, `level`=?, `exp`=?")
	g_SQL.PlayerTaoyuanHeart.Update = prepare(db, "UPDATE `player_taoyuan_heart` SET `level`=?, `exp`=? WHERE `pid`=?")
	g_SQL.PlayerTaoyuanHeart.Delete = prepare(db, "DELETE FROM `player_taoyuan_heart` WHERE `pid`=?")
	g_SQL.PlayerTaoyuanItem.Insert = prepare(db, "INSERT INTO `player_taoyuan_item` SET `id`=?, `pid`=?, `item_id`=?, `num`=?")
	g_SQL.PlayerTaoyuanItem.Update = prepare(db, "UPDATE `player_taoyuan_item` SET `pid`=?, `item_id`=?, `num`=? WHERE `id`=?")
	g_SQL.PlayerTaoyuanItem.Delete = prepare(db, "DELETE FROM `player_taoyuan_item` WHERE `id`=?")
	g_SQL.PlayerTaoyuanMessage.Insert = prepare(db, "INSERT INTO `player_taoyuan_message` SET `id`=?, `pid`=?, `nick`=?, `exp`=?")
	g_SQL.PlayerTaoyuanMessage.Update = prepare(db, "UPDATE `player_taoyuan_message` SET `pid`=?, `nick`=?, `exp`=? WHERE `id`=?")
	g_SQL.PlayerTaoyuanMessage.Delete = prepare(db, "DELETE FROM `player_taoyuan_message` WHERE `id`=?")
	g_SQL.PlayerTaoyuanProduct.Insert = prepare(db, "INSERT INTO `player_taoyuan_product` SET `pid`=?, `skill_id1`=?, `skill_id2`=?, `same_time_wine_queue`=?, `make_wine_times`=?, `same_time_food_queue`=?, `make_food_times`=?, `food_queue1`=?, `food_queue1_start_timestamp`=?, `food_queue1_end_timestamp`=?, `food_queue2`=?, `food_queue2_start_timestamp`=?, `food_queue2_end_timestamp`=?, `food_queue3`=?, `food_queue3_start_timestamp`=?, `food_queue3_end_timestamp`=?, `food_queue4`=?, `food_queue4_start_timestamp`=?, `food_queue4_end_timestamp`=?, `wine_queue1`=?, `wine_queue1_start_timestamp`=?, `wine_queue1_end_timestamp`=?, `wine_queue2`=?, `wine_queue2_start_timestamp`=?, `wine_queue2_end_timestamp`=?, `wine_queue3`=?, `wine_queue3_start_timestamp`=?, `wine_queue3_end_timestamp`=?, `wine_queue4`=?, `wine_queue4_start_timestamp`=?, `wine_queue4_end_timestamp`=?, `is_open_wine`=?, `is_open_food`=?")
	g_SQL.PlayerTaoyuanProduct.Update = prepare(db, "UPDATE `player_taoyuan_product` SET `skill_id1`=?, `skill_id2`=?, `same_time_wine_queue`=?, `make_wine_times`=?, `same_time_food_queue`=?, `make_food_times`=?, `food_queue1`=?, `food_queue1_start_timestamp`=?, `food_queue1_end_timestamp`=?, `food_queue2`=?, `food_queue2_start_timestamp`=?, `food_queue2_end_timestamp`=?, `food_queue3`=?, `food_queue3_start_timestamp`=?, `food_queue3_end_timestamp`=?, `food_queue4`=?, `food_queue4_start_timestamp`=?, `food_queue4_end_timestamp`=?, `wine_queue1`=?, `wine_queue1_start_timestamp`=?, `wine_queue1_end_timestamp`=?, `wine_queue2`=?, `wine_queue2_start_timestamp`=?, `wine_queue2_end_timestamp`=?, `wine_queue3`=?, `wine_queue3_start_timestamp`=?, `wine_queue3_end_timestamp`=?, `wine_queue4`=?, `wine_queue4_start_timestamp`=?, `wine_queue4_end_timestamp`=?, `is_open_wine`=?, `is_open_food`=? WHERE `pid`=?")
	g_SQL.PlayerTaoyuanProduct.Delete = prepare(db, "DELETE FROM `player_taoyuan_product` WHERE `pid`=?")
	g_SQL.PlayerTaoyuanPurchaseRecord.Insert = prepare(db, "INSERT INTO `player_taoyuan_purchase_record` SET `id`=?, `pid`=?, `item_id`=?, `num`=?, `timestamp`=?")
	g_SQL.PlayerTaoyuanPurchaseRecord.Update = prepare(db, "UPDATE `player_taoyuan_purchase_record` SET `pid`=?, `item_id`=?, `num`=?, `timestamp`=? WHERE `id`=?")
	g_SQL.PlayerTaoyuanPurchaseRecord.Delete = prepare(db, "DELETE FROM `player_taoyuan_purchase_record` WHERE `id`=?")
	g_SQL.PlayerTaoyuanQuest.Insert = prepare(db, "INSERT INTO `player_taoyuan_quest` SET `pid`=?, `last_quest_flash_time`=?, `quest1_item_id`=?, `quest1_item_num`=?, `quest1_total_exp`=?, `quest1_total_coins`=?, `quest1_finish_time`=?, `quest2_item_id`=?, `quest2_item_num`=?, `quest2_total_exp`=?, `quest2_total_coins`=?, `quest2_finish_time`=?, `quest3_item_id`=?, `quest3_item_num`=?, `quest3_total_exp`=?, `quest3_total_coins`=?, `quest3_finish_time`=?, `quest4_item_id`=?, `quest4_item_num`=?, `quest4_total_exp`=?, `quest4_total_coins`=?, `quest4_finish_time`=?, `quest5_item_id`=?, `quest5_item_num`=?, `quest5_total_exp`=?, `quest5_total_coins`=?, `quest5_finish_time`=?")
	g_SQL.PlayerTaoyuanQuest.Update = prepare(db, "UPDATE `player_taoyuan_quest` SET `last_quest_flash_time`=?, `quest1_item_id`=?, `quest1_item_num`=?, `quest1_total_exp`=?, `quest1_total_coins`=?, `quest1_finish_time`=?, `quest2_item_id`=?, `quest2_item_num`=?, `quest2_total_exp`=?, `quest2_total_coins`=?, `quest2_finish_time`=?, `quest3_item_id`=?, `quest3_item_num`=?, `quest3_total_exp`=?, `quest3_total_coins`=?, `quest3_finish_time`=?, `quest4_item_id`=?, `quest4_item_num`=?, `quest4_total_exp`=?, `quest4_total_coins`=?, `quest4_finish_time`=?, `quest5_item_id`=?, `quest5_item_num`=?, `quest5_total_exp`=?, `quest5_total_coins`=?, `quest5_finish_time`=? WHERE `pid`=?")
	g_SQL.PlayerTaoyuanQuest.Delete = prepare(db, "DELETE FROM `player_taoyuan_quest` WHERE `pid`=?")
	g_SQL.PlayerTbXxdRoleinfo.Insert = prepare(db, "INSERT INTO `player_tb_xxd_roleinfo` SET `pid`=?, `gameappid`=?, `openid`=?, `regtime`=?, `level`=?, `iFriends`=?, `moneyios`=?, `moneyandroid`=?, `diamondios`=?, `diamondandroid`=?")
	g_SQL.PlayerTbXxdRoleinfo.Update = prepare(db, "UPDATE `player_tb_xxd_roleinfo` SET `gameappid`=?, `openid`=?, `regtime`=?, `level`=?, `iFriends`=?, `moneyios`=?, `moneyandroid`=?, `diamondios`=?, `diamondandroid`=? WHERE `pid`=?")
	g_SQL.PlayerTbXxdRoleinfo.Delete = prepare(db, "DELETE FROM `player_tb_xxd_roleinfo` WHERE `pid`=?")
	g_SQL.PlayerTeamInfo.Insert = prepare(db, "INSERT INTO `player_team_info` SET `pid`=?, `relationship`=?, `health_lv`=?, `attack_lv`=?, `defence_lv`=?")
	g_SQL.PlayerTeamInfo.Update = prepare(db, "UPDATE `player_team_info` SET `relationship`=?, `health_lv`=?, `attack_lv`=?, `defence_lv`=? WHERE `pid`=?")
	g_SQL.PlayerTeamInfo.Delete = prepare(db, "DELETE FROM `player_team_info` WHERE `pid`=?")
	g_SQL.PlayerTotem.Insert = prepare(db, "INSERT INTO `player_totem` SET `id`=?, `pid`=?, `totem_id`=?, `level`=?, `skill_id`=?")
	g_SQL.PlayerTotem.Update = prepare(db, "UPDATE `player_totem` SET `pid`=?, `totem_id`=?, `level`=?, `skill_id`=? WHERE `id`=?")
	g_SQL.PlayerTotem.Delete = prepare(db, "DELETE FROM `player_totem` WHERE `id`=?")
	g_SQL.PlayerTotemInfo.Insert = prepare(db, "INSERT INTO `player_totem_info` SET `pid`=?, `ingot_call_daily_num`=?, `ingot_call_timestamp`=?, `rock_rune_num`=?, `jade_rune_num`=?, `pos1`=?, `pos2`=?, `pos3`=?, `pos4`=?, `pos5`=?, `ingot_draw_times`=?")
	g_SQL.PlayerTotemInfo.Update = prepare(db, "UPDATE `player_totem_info` SET `ingot_call_daily_num`=?, `ingot_call_timestamp`=?, `rock_rune_num`=?, `jade_rune_num`=?, `pos1`=?, `pos2`=?, `pos3`=?, `pos4`=?, `pos5`=?, `ingot_draw_times`=? WHERE `pid`=?")
	g_SQL.PlayerTotemInfo.Delete = prepare(db, "DELETE FROM `player_totem_info` WHERE `pid`=?")
	g_SQL.PlayerTown.Insert = prepare(db, "INSERT INTO `player_town` SET `pid`=?, `town_id`=?, `lock`=?, `at_pos_x`=?, `at_pos_y`=?")
	g_SQL.PlayerTown.Update = prepare(db, "UPDATE `player_town` SET `town_id`=?, `lock`=?, `at_pos_x`=?, `at_pos_y`=? WHERE `pid`=?")
	g_SQL.PlayerTown.Delete = prepare(db, "DELETE FROM `player_town` WHERE `pid`=?")
	g_SQL.PlayerTraderRefreshState.Insert = prepare(db, "INSERT INTO `player_trader_refresh_state` SET `id`=?, `pid`=?, `last_update_time`=?, `auto_update_time`=?, `trader_id`=?, `refresh_num`=?")
	g_SQL.PlayerTraderRefreshState.Update = prepare(db, "UPDATE `player_trader_refresh_state` SET `pid`=?, `last_update_time`=?, `auto_update_time`=?, `trader_id`=?, `refresh_num`=? WHERE `id`=?")
	g_SQL.PlayerTraderRefreshState.Delete = prepare(db, "DELETE FROM `player_trader_refresh_state` WHERE `id`=?")
	g_SQL.PlayerTraderStoreState.Insert = prepare(db, "INSERT INTO `player_trader_store_state` SET `id`=?, `pid`=?, `trader_id`=?, `grid_id`=?, `item_id`=?, `num`=?, `cost`=?, `stock`=?, `goods_type`=?")
	g_SQL.PlayerTraderStoreState.Update = prepare(db, "UPDATE `player_trader_store_state` SET `pid`=?, `trader_id`=?, `grid_id`=?, `item_id`=?, `num`=?, `cost`=?, `stock`=?, `goods_type`=? WHERE `id`=?")
	g_SQL.PlayerTraderStoreState.Delete = prepare(db, "DELETE FROM `player_trader_store_state` WHERE `id`=?")
	g_SQL.PlayerUseSkill.Insert = prepare(db, "INSERT INTO `player_use_skill` SET `id`=?, `pid`=?, `role_id`=?, `skill_id1`=?, `skill_id4`=?, `skill_id2`=?, `skill_id3`=?")
	g_SQL.PlayerUseSkill.Update = prepare(db, "UPDATE `player_use_skill` SET `pid`=?, `role_id`=?, `skill_id1`=?, `skill_id4`=?, `skill_id2`=?, `skill_id3`=? WHERE `id`=?")
	g_SQL.PlayerUseSkill.Delete = prepare(db, "DELETE FROM `player_use_skill` WHERE `id`=?")
	g_SQL.PlayerVip.Insert = prepare(db, "INSERT INTO `player_vip` SET `pid`=?, `ingot`=?, `level`=?, `card_id`=?, `present_exp`=?")
	g_SQL.PlayerVip.Update = prepare(db, "UPDATE `player_vip` SET `ingot`=?, `level`=?, `card_id`=?, `present_exp`=? WHERE `pid`=?")
	g_SQL.PlayerVip.Delete = prepare(db, "DELETE FROM `player_vip` WHERE `pid`=?")
}

func prepareSQLClean() {
	g_SQL.GlobalAnnouncement.Insert.Close()
	g_SQL.GlobalAnnouncement.Update.Close()
	g_SQL.GlobalAnnouncement.Delete.Close()
	g_SQL.GlobalArenaRank.Insert.Close()
	g_SQL.GlobalArenaRank.Update.Close()
	g_SQL.GlobalArenaRank.Delete.Close()
	g_SQL.GlobalClique.Insert.Close()
	g_SQL.GlobalClique.Update.Close()
	g_SQL.GlobalClique.Delete.Close()
	g_SQL.GlobalCliqueBoat.Insert.Close()
	g_SQL.GlobalCliqueBoat.Update.Close()
	g_SQL.GlobalCliqueBoat.Delete.Close()
	g_SQL.GlobalDespairBoss.Insert.Close()
	g_SQL.GlobalDespairBoss.Update.Close()
	g_SQL.GlobalDespairBoss.Delete.Close()
	g_SQL.GlobalDespairBossHistory.Insert.Close()
	g_SQL.GlobalDespairBossHistory.Update.Close()
	g_SQL.GlobalDespairBossHistory.Delete.Close()
	g_SQL.GlobalDespairLand.Insert.Close()
	g_SQL.GlobalDespairLand.Update.Close()
	g_SQL.GlobalDespairLand.Delete.Close()
	g_SQL.GlobalDespairLandBattleRecord.Insert.Close()
	g_SQL.GlobalDespairLandBattleRecord.Update.Close()
	g_SQL.GlobalDespairLandBattleRecord.Delete.Close()
	g_SQL.GlobalDespairLandCamp.Insert.Close()
	g_SQL.GlobalDespairLandCamp.Update.Close()
	g_SQL.GlobalDespairLandCamp.Delete.Close()
	g_SQL.GlobalGiftCardRecord.Insert.Close()
	g_SQL.GlobalGiftCardRecord.Update.Close()
	g_SQL.GlobalGiftCardRecord.Delete.Close()
	g_SQL.GlobalGroupBuyStatus.Insert.Close()
	g_SQL.GlobalGroupBuyStatus.Update.Close()
	g_SQL.GlobalGroupBuyStatus.Delete.Close()
	g_SQL.GlobalMail.Insert.Close()
	g_SQL.GlobalMail.Update.Close()
	g_SQL.GlobalMail.Delete.Close()
	g_SQL.GlobalMailAttachments.Insert.Close()
	g_SQL.GlobalMailAttachments.Update.Close()
	g_SQL.GlobalMailAttachments.Delete.Close()
	g_SQL.GlobalTbXxdOnlinecnt.Insert.Close()
	g_SQL.GlobalTbXxdOnlinecnt.Update.Close()
	g_SQL.GlobalTbXxdOnlinecnt.Delete.Close()
	g_SQL.Player.Insert.Close()
	g_SQL.Player.Update.Close()
	g_SQL.Player.Delete.Close()
	g_SQL.PlayerActivity.Insert.Close()
	g_SQL.PlayerActivity.Update.Close()
	g_SQL.PlayerActivity.Delete.Close()
	g_SQL.PlayerAdditionQuest.Insert.Close()
	g_SQL.PlayerAdditionQuest.Update.Close()
	g_SQL.PlayerAdditionQuest.Delete.Close()
	g_SQL.PlayerAnySdkOrder.Insert.Close()
	g_SQL.PlayerAnySdkOrder.Update.Close()
	g_SQL.PlayerAnySdkOrder.Delete.Close()
	g_SQL.PlayerArena.Insert.Close()
	g_SQL.PlayerArena.Update.Close()
	g_SQL.PlayerArena.Delete.Close()
	g_SQL.PlayerArenaRank.Insert.Close()
	g_SQL.PlayerArenaRank.Update.Close()
	g_SQL.PlayerArenaRank.Delete.Close()
	g_SQL.PlayerArenaRecord.Insert.Close()
	g_SQL.PlayerArenaRecord.Update.Close()
	g_SQL.PlayerArenaRecord.Delete.Close()
	g_SQL.PlayerAwakenGraphic.Insert.Close()
	g_SQL.PlayerAwakenGraphic.Update.Close()
	g_SQL.PlayerAwakenGraphic.Delete.Close()
	g_SQL.PlayerBattlePet.Insert.Close()
	g_SQL.PlayerBattlePet.Update.Close()
	g_SQL.PlayerBattlePet.Delete.Close()
	g_SQL.PlayerBattlePetGrid.Insert.Close()
	g_SQL.PlayerBattlePetGrid.Update.Close()
	g_SQL.PlayerBattlePetGrid.Delete.Close()
	g_SQL.PlayerBattlePetState.Insert.Close()
	g_SQL.PlayerBattlePetState.Update.Close()
	g_SQL.PlayerBattlePetState.Delete.Close()
	g_SQL.PlayerChestState.Insert.Close()
	g_SQL.PlayerChestState.Update.Close()
	g_SQL.PlayerChestState.Delete.Close()
	g_SQL.PlayerCliqueKongfuAttrib.Insert.Close()
	g_SQL.PlayerCliqueKongfuAttrib.Update.Close()
	g_SQL.PlayerCliqueKongfuAttrib.Delete.Close()
	g_SQL.PlayerCoins.Insert.Close()
	g_SQL.PlayerCoins.Update.Close()
	g_SQL.PlayerCoins.Delete.Close()
	g_SQL.PlayerCornucopia.Insert.Close()
	g_SQL.PlayerCornucopia.Update.Close()
	g_SQL.PlayerCornucopia.Delete.Close()
	g_SQL.PlayerDailyQuest.Insert.Close()
	g_SQL.PlayerDailyQuest.Update.Close()
	g_SQL.PlayerDailyQuest.Delete.Close()
	g_SQL.PlayerDailyQuestStarAwardInfo.Insert.Close()
	g_SQL.PlayerDailyQuestStarAwardInfo.Update.Close()
	g_SQL.PlayerDailyQuestStarAwardInfo.Delete.Close()
	g_SQL.PlayerDailySignInRecord.Insert.Close()
	g_SQL.PlayerDailySignInRecord.Update.Close()
	g_SQL.PlayerDailySignInRecord.Delete.Close()
	g_SQL.PlayerDailySignInState.Insert.Close()
	g_SQL.PlayerDailySignInState.Update.Close()
	g_SQL.PlayerDailySignInState.Delete.Close()
	g_SQL.PlayerDespairLandCampState.Insert.Close()
	g_SQL.PlayerDespairLandCampState.Update.Close()
	g_SQL.PlayerDespairLandCampState.Delete.Close()
	g_SQL.PlayerDespairLandCampStateHistory.Insert.Close()
	g_SQL.PlayerDespairLandCampStateHistory.Update.Close()
	g_SQL.PlayerDespairLandCampStateHistory.Delete.Close()
	g_SQL.PlayerDespairLandLevelRecord.Insert.Close()
	g_SQL.PlayerDespairLandLevelRecord.Update.Close()
	g_SQL.PlayerDespairLandLevelRecord.Delete.Close()
	g_SQL.PlayerDespairLandState.Insert.Close()
	g_SQL.PlayerDespairLandState.Update.Close()
	g_SQL.PlayerDespairLandState.Delete.Close()
	g_SQL.PlayerDrivingSwordEvent.Insert.Close()
	g_SQL.PlayerDrivingSwordEvent.Update.Close()
	g_SQL.PlayerDrivingSwordEvent.Delete.Close()
	g_SQL.PlayerDrivingSwordEventExploring.Insert.Close()
	g_SQL.PlayerDrivingSwordEventExploring.Update.Close()
	g_SQL.PlayerDrivingSwordEventExploring.Delete.Close()
	g_SQL.PlayerDrivingSwordEventTreasure.Insert.Close()
	g_SQL.PlayerDrivingSwordEventTreasure.Update.Close()
	g_SQL.PlayerDrivingSwordEventTreasure.Delete.Close()
	g_SQL.PlayerDrivingSwordEventVisiting.Insert.Close()
	g_SQL.PlayerDrivingSwordEventVisiting.Update.Close()
	g_SQL.PlayerDrivingSwordEventVisiting.Delete.Close()
	g_SQL.PlayerDrivingSwordEventmask.Insert.Close()
	g_SQL.PlayerDrivingSwordEventmask.Update.Close()
	g_SQL.PlayerDrivingSwordEventmask.Delete.Close()
	g_SQL.PlayerDrivingSwordInfo.Insert.Close()
	g_SQL.PlayerDrivingSwordInfo.Update.Close()
	g_SQL.PlayerDrivingSwordInfo.Delete.Close()
	g_SQL.PlayerDrivingSwordMap.Insert.Close()
	g_SQL.PlayerDrivingSwordMap.Update.Close()
	g_SQL.PlayerDrivingSwordMap.Delete.Close()
	g_SQL.PlayerEquipment.Insert.Close()
	g_SQL.PlayerEquipment.Update.Close()
	g_SQL.PlayerEquipment.Delete.Close()
	g_SQL.PlayerEventAwardRecord.Insert.Close()
	g_SQL.PlayerEventAwardRecord.Update.Close()
	g_SQL.PlayerEventAwardRecord.Delete.Close()
	g_SQL.PlayerEventDailyOnline.Insert.Close()
	g_SQL.PlayerEventDailyOnline.Update.Close()
	g_SQL.PlayerEventDailyOnline.Delete.Close()
	g_SQL.PlayerEventVnDailyRecharge.Insert.Close()
	g_SQL.PlayerEventVnDailyRecharge.Update.Close()
	g_SQL.PlayerEventVnDailyRecharge.Delete.Close()
	g_SQL.PlayerEventVnDragonBallRecord.Insert.Close()
	g_SQL.PlayerEventVnDragonBallRecord.Update.Close()
	g_SQL.PlayerEventVnDragonBallRecord.Delete.Close()
	g_SQL.PlayerEventsFightPower.Insert.Close()
	g_SQL.PlayerEventsFightPower.Update.Close()
	g_SQL.PlayerEventsFightPower.Delete.Close()
	g_SQL.PlayerEventsIngotRecord.Insert.Close()
	g_SQL.PlayerEventsIngotRecord.Update.Close()
	g_SQL.PlayerEventsIngotRecord.Delete.Close()
	g_SQL.PlayerExtendLevel.Insert.Close()
	g_SQL.PlayerExtendLevel.Update.Close()
	g_SQL.PlayerExtendLevel.Delete.Close()
	g_SQL.PlayerExtendQuest.Insert.Close()
	g_SQL.PlayerExtendQuest.Update.Close()
	g_SQL.PlayerExtendQuest.Delete.Close()
	g_SQL.PlayerFame.Insert.Close()
	g_SQL.PlayerFame.Update.Close()
	g_SQL.PlayerFame.Delete.Close()
	g_SQL.PlayerFashion.Insert.Close()
	g_SQL.PlayerFashion.Update.Close()
	g_SQL.PlayerFashion.Delete.Close()
	g_SQL.PlayerFashionState.Insert.Close()
	g_SQL.PlayerFashionState.Update.Close()
	g_SQL.PlayerFashionState.Delete.Close()
	g_SQL.PlayerFateBoxState.Insert.Close()
	g_SQL.PlayerFateBoxState.Update.Close()
	g_SQL.PlayerFateBoxState.Delete.Close()
	g_SQL.PlayerFightNum.Insert.Close()
	g_SQL.PlayerFightNum.Update.Close()
	g_SQL.PlayerFightNum.Delete.Close()
	g_SQL.PlayerFirstCharge.Insert.Close()
	g_SQL.PlayerFirstCharge.Update.Close()
	g_SQL.PlayerFirstCharge.Delete.Close()
	g_SQL.PlayerFormation.Insert.Close()
	g_SQL.PlayerFormation.Update.Close()
	g_SQL.PlayerFormation.Delete.Close()
	g_SQL.PlayerFuncKey.Insert.Close()
	g_SQL.PlayerFuncKey.Update.Close()
	g_SQL.PlayerFuncKey.Delete.Close()
	g_SQL.PlayerGhost.Insert.Close()
	g_SQL.PlayerGhost.Update.Close()
	g_SQL.PlayerGhost.Delete.Close()
	g_SQL.PlayerGhostEquipment.Insert.Close()
	g_SQL.PlayerGhostEquipment.Update.Close()
	g_SQL.PlayerGhostEquipment.Delete.Close()
	g_SQL.PlayerGhostState.Insert.Close()
	g_SQL.PlayerGhostState.Update.Close()
	g_SQL.PlayerGhostState.Delete.Close()
	g_SQL.PlayerGlobalChatState.Insert.Close()
	g_SQL.PlayerGlobalChatState.Update.Close()
	g_SQL.PlayerGlobalChatState.Delete.Close()
	g_SQL.PlayerGlobalCliqueBuilding.Insert.Close()
	g_SQL.PlayerGlobalCliqueBuilding.Update.Close()
	g_SQL.PlayerGlobalCliqueBuilding.Delete.Close()
	g_SQL.PlayerGlobalCliqueBuildingQuest.Insert.Close()
	g_SQL.PlayerGlobalCliqueBuildingQuest.Update.Close()
	g_SQL.PlayerGlobalCliqueBuildingQuest.Delete.Close()
	g_SQL.PlayerGlobalCliqueDailyQuest.Insert.Close()
	g_SQL.PlayerGlobalCliqueDailyQuest.Update.Close()
	g_SQL.PlayerGlobalCliqueDailyQuest.Delete.Close()
	g_SQL.PlayerGlobalCliqueEscort.Insert.Close()
	g_SQL.PlayerGlobalCliqueEscort.Update.Close()
	g_SQL.PlayerGlobalCliqueEscort.Delete.Close()
	g_SQL.PlayerGlobalCliqueEscortMessage.Insert.Close()
	g_SQL.PlayerGlobalCliqueEscortMessage.Update.Close()
	g_SQL.PlayerGlobalCliqueEscortMessage.Delete.Close()
	g_SQL.PlayerGlobalCliqueHijack.Insert.Close()
	g_SQL.PlayerGlobalCliqueHijack.Update.Close()
	g_SQL.PlayerGlobalCliqueHijack.Delete.Close()
	g_SQL.PlayerGlobalCliqueInfo.Insert.Close()
	g_SQL.PlayerGlobalCliqueInfo.Update.Close()
	g_SQL.PlayerGlobalCliqueInfo.Delete.Close()
	g_SQL.PlayerGlobalCliqueKongfu.Insert.Close()
	g_SQL.PlayerGlobalCliqueKongfu.Update.Close()
	g_SQL.PlayerGlobalCliqueKongfu.Delete.Close()
	g_SQL.PlayerGlobalFriend.Insert.Close()
	g_SQL.PlayerGlobalFriend.Update.Close()
	g_SQL.PlayerGlobalFriend.Delete.Close()
	g_SQL.PlayerGlobalFriendChat.Insert.Close()
	g_SQL.PlayerGlobalFriendChat.Update.Close()
	g_SQL.PlayerGlobalFriendChat.Delete.Close()
	g_SQL.PlayerGlobalFriendState.Insert.Close()
	g_SQL.PlayerGlobalFriendState.Update.Close()
	g_SQL.PlayerGlobalFriendState.Delete.Close()
	g_SQL.PlayerGlobalGiftCardRecord.Insert.Close()
	g_SQL.PlayerGlobalGiftCardRecord.Update.Close()
	g_SQL.PlayerGlobalGiftCardRecord.Delete.Close()
	g_SQL.PlayerGlobalMailState.Insert.Close()
	g_SQL.PlayerGlobalMailState.Update.Close()
	g_SQL.PlayerGlobalMailState.Delete.Close()
	g_SQL.PlayerGlobalWorldChatState.Insert.Close()
	g_SQL.PlayerGlobalWorldChatState.Update.Close()
	g_SQL.PlayerGlobalWorldChatState.Delete.Close()
	g_SQL.PlayerHardLevel.Insert.Close()
	g_SQL.PlayerHardLevel.Update.Close()
	g_SQL.PlayerHardLevel.Delete.Close()
	g_SQL.PlayerHardLevelRecord.Insert.Close()
	g_SQL.PlayerHardLevelRecord.Update.Close()
	g_SQL.PlayerHardLevelRecord.Delete.Close()
	g_SQL.PlayerHeart.Insert.Close()
	g_SQL.PlayerHeart.Update.Close()
	g_SQL.PlayerHeart.Delete.Close()
	g_SQL.PlayerHeartDraw.Insert.Close()
	g_SQL.PlayerHeartDraw.Update.Close()
	g_SQL.PlayerHeartDraw.Delete.Close()
	g_SQL.PlayerHeartDrawCardRecord.Insert.Close()
	g_SQL.PlayerHeartDrawCardRecord.Update.Close()
	g_SQL.PlayerHeartDrawCardRecord.Delete.Close()
	g_SQL.PlayerHeartDrawWheelRecord.Insert.Close()
	g_SQL.PlayerHeartDrawWheelRecord.Update.Close()
	g_SQL.PlayerHeartDrawWheelRecord.Delete.Close()
	g_SQL.PlayerInfo.Insert.Close()
	g_SQL.PlayerInfo.Update.Close()
	g_SQL.PlayerInfo.Delete.Close()
	g_SQL.PlayerIsOperated.Insert.Close()
	g_SQL.PlayerIsOperated.Update.Close()
	g_SQL.PlayerIsOperated.Delete.Close()
	g_SQL.PlayerItem.Insert.Close()
	g_SQL.PlayerItem.Update.Close()
	g_SQL.PlayerItem.Delete.Close()
	g_SQL.PlayerItemAppendix.Insert.Close()
	g_SQL.PlayerItemAppendix.Update.Close()
	g_SQL.PlayerItemAppendix.Delete.Close()
	g_SQL.PlayerItemBuyback.Insert.Close()
	g_SQL.PlayerItemBuyback.Update.Close()
	g_SQL.PlayerItemBuyback.Delete.Close()
	g_SQL.PlayerItemRecastState.Insert.Close()
	g_SQL.PlayerItemRecastState.Update.Close()
	g_SQL.PlayerItemRecastState.Delete.Close()
	g_SQL.PlayerLevelAwardInfo.Insert.Close()
	g_SQL.PlayerLevelAwardInfo.Update.Close()
	g_SQL.PlayerLevelAwardInfo.Delete.Close()
	g_SQL.PlayerLoginAwardRecord.Insert.Close()
	g_SQL.PlayerLoginAwardRecord.Update.Close()
	g_SQL.PlayerLoginAwardRecord.Delete.Close()
	g_SQL.PlayerMail.Insert.Close()
	g_SQL.PlayerMail.Update.Close()
	g_SQL.PlayerMail.Delete.Close()
	g_SQL.PlayerMailAttachment.Insert.Close()
	g_SQL.PlayerMailAttachment.Update.Close()
	g_SQL.PlayerMailAttachment.Delete.Close()
	g_SQL.PlayerMailAttachmentLg.Insert.Close()
	g_SQL.PlayerMailAttachmentLg.Update.Close()
	g_SQL.PlayerMailAttachmentLg.Delete.Close()
	g_SQL.PlayerMailLg.Insert.Close()
	g_SQL.PlayerMailLg.Update.Close()
	g_SQL.PlayerMailLg.Delete.Close()
	g_SQL.PlayerMeditationState.Insert.Close()
	g_SQL.PlayerMeditationState.Update.Close()
	g_SQL.PlayerMeditationState.Delete.Close()
	g_SQL.PlayerMission.Insert.Close()
	g_SQL.PlayerMission.Update.Close()
	g_SQL.PlayerMission.Delete.Close()
	g_SQL.PlayerMissionLevel.Insert.Close()
	g_SQL.PlayerMissionLevel.Update.Close()
	g_SQL.PlayerMissionLevel.Delete.Close()
	g_SQL.PlayerMissionLevelRecord.Insert.Close()
	g_SQL.PlayerMissionLevelRecord.Update.Close()
	g_SQL.PlayerMissionLevelRecord.Delete.Close()
	g_SQL.PlayerMissionLevelStateBin.Insert.Close()
	g_SQL.PlayerMissionLevelStateBin.Update.Close()
	g_SQL.PlayerMissionLevelStateBin.Delete.Close()
	g_SQL.PlayerMissionRecord.Insert.Close()
	g_SQL.PlayerMissionRecord.Update.Close()
	g_SQL.PlayerMissionRecord.Delete.Close()
	g_SQL.PlayerMissionStarAward.Insert.Close()
	g_SQL.PlayerMissionStarAward.Update.Close()
	g_SQL.PlayerMissionStarAward.Delete.Close()
	g_SQL.PlayerMoneyTree.Insert.Close()
	g_SQL.PlayerMoneyTree.Update.Close()
	g_SQL.PlayerMoneyTree.Delete.Close()
	g_SQL.PlayerMonthCardAwardRecord.Insert.Close()
	g_SQL.PlayerMonthCardAwardRecord.Update.Close()
	g_SQL.PlayerMonthCardAwardRecord.Delete.Close()
	g_SQL.PlayerMonthCardInfo.Insert.Close()
	g_SQL.PlayerMonthCardInfo.Update.Close()
	g_SQL.PlayerMonthCardInfo.Delete.Close()
	g_SQL.PlayerMultiLevelInfo.Insert.Close()
	g_SQL.PlayerMultiLevelInfo.Update.Close()
	g_SQL.PlayerMultiLevelInfo.Delete.Close()
	g_SQL.PlayerNewYearConsumeRecord.Insert.Close()
	g_SQL.PlayerNewYearConsumeRecord.Update.Close()
	g_SQL.PlayerNewYearConsumeRecord.Delete.Close()
	g_SQL.PlayerNpcTalkRecord.Insert.Close()
	g_SQL.PlayerNpcTalkRecord.Update.Close()
	g_SQL.PlayerNpcTalkRecord.Delete.Close()
	g_SQL.PlayerOpenedTownTreasure.Insert.Close()
	g_SQL.PlayerOpenedTownTreasure.Update.Close()
	g_SQL.PlayerOpenedTownTreasure.Delete.Close()
	g_SQL.PlayerPhysical.Insert.Close()
	g_SQL.PlayerPhysical.Update.Close()
	g_SQL.PlayerPhysical.Delete.Close()
	g_SQL.PlayerPurchaseRecord.Insert.Close()
	g_SQL.PlayerPurchaseRecord.Update.Close()
	g_SQL.PlayerPurchaseRecord.Delete.Close()
	g_SQL.PlayerPushNotifySwitch.Insert.Close()
	g_SQL.PlayerPushNotifySwitch.Update.Close()
	g_SQL.PlayerPushNotifySwitch.Delete.Close()
	g_SQL.PlayerPveState.Insert.Close()
	g_SQL.PlayerPveState.Update.Close()
	g_SQL.PlayerPveState.Delete.Close()
	g_SQL.PlayerQqVipGift.Insert.Close()
	g_SQL.PlayerQqVipGift.Update.Close()
	g_SQL.PlayerQqVipGift.Delete.Close()
	g_SQL.PlayerQuest.Insert.Close()
	g_SQL.PlayerQuest.Update.Close()
	g_SQL.PlayerQuest.Delete.Close()
	g_SQL.PlayerRainbowLevel.Insert.Close()
	g_SQL.PlayerRainbowLevel.Update.Close()
	g_SQL.PlayerRainbowLevel.Delete.Close()
	g_SQL.PlayerRainbowLevelStateBin.Insert.Close()
	g_SQL.PlayerRainbowLevelStateBin.Update.Close()
	g_SQL.PlayerRainbowLevelStateBin.Delete.Close()
	g_SQL.PlayerRole.Insert.Close()
	g_SQL.PlayerRole.Update.Close()
	g_SQL.PlayerRole.Delete.Close()
	g_SQL.PlayerSealedbook.Insert.Close()
	g_SQL.PlayerSealedbook.Update.Close()
	g_SQL.PlayerSealedbook.Delete.Close()
	g_SQL.PlayerSendHeartRecord.Insert.Close()
	g_SQL.PlayerSendHeartRecord.Update.Close()
	g_SQL.PlayerSendHeartRecord.Delete.Close()
	g_SQL.PlayerSkill.Insert.Close()
	g_SQL.PlayerSkill.Update.Close()
	g_SQL.PlayerSkill.Delete.Close()
	g_SQL.PlayerState.Insert.Close()
	g_SQL.PlayerState.Update.Close()
	g_SQL.PlayerState.Delete.Close()
	g_SQL.PlayerSwordSoul.Insert.Close()
	g_SQL.PlayerSwordSoul.Update.Close()
	g_SQL.PlayerSwordSoul.Delete.Close()
	g_SQL.PlayerSwordSoulEquipment.Insert.Close()
	g_SQL.PlayerSwordSoulEquipment.Update.Close()
	g_SQL.PlayerSwordSoulEquipment.Delete.Close()
	g_SQL.PlayerSwordSoulState.Insert.Close()
	g_SQL.PlayerSwordSoulState.Update.Close()
	g_SQL.PlayerSwordSoulState.Delete.Close()
	g_SQL.PlayerTaoyuanBless.Insert.Close()
	g_SQL.PlayerTaoyuanBless.Update.Close()
	g_SQL.PlayerTaoyuanBless.Delete.Close()
	g_SQL.PlayerTaoyuanFileds.Insert.Close()
	g_SQL.PlayerTaoyuanFileds.Update.Close()
	g_SQL.PlayerTaoyuanFileds.Delete.Close()
	g_SQL.PlayerTaoyuanHeart.Insert.Close()
	g_SQL.PlayerTaoyuanHeart.Update.Close()
	g_SQL.PlayerTaoyuanHeart.Delete.Close()
	g_SQL.PlayerTaoyuanItem.Insert.Close()
	g_SQL.PlayerTaoyuanItem.Update.Close()
	g_SQL.PlayerTaoyuanItem.Delete.Close()
	g_SQL.PlayerTaoyuanMessage.Insert.Close()
	g_SQL.PlayerTaoyuanMessage.Update.Close()
	g_SQL.PlayerTaoyuanMessage.Delete.Close()
	g_SQL.PlayerTaoyuanProduct.Insert.Close()
	g_SQL.PlayerTaoyuanProduct.Update.Close()
	g_SQL.PlayerTaoyuanProduct.Delete.Close()
	g_SQL.PlayerTaoyuanPurchaseRecord.Insert.Close()
	g_SQL.PlayerTaoyuanPurchaseRecord.Update.Close()
	g_SQL.PlayerTaoyuanPurchaseRecord.Delete.Close()
	g_SQL.PlayerTaoyuanQuest.Insert.Close()
	g_SQL.PlayerTaoyuanQuest.Update.Close()
	g_SQL.PlayerTaoyuanQuest.Delete.Close()
	g_SQL.PlayerTbXxdRoleinfo.Insert.Close()
	g_SQL.PlayerTbXxdRoleinfo.Update.Close()
	g_SQL.PlayerTbXxdRoleinfo.Delete.Close()
	g_SQL.PlayerTeamInfo.Insert.Close()
	g_SQL.PlayerTeamInfo.Update.Close()
	g_SQL.PlayerTeamInfo.Delete.Close()
	g_SQL.PlayerTotem.Insert.Close()
	g_SQL.PlayerTotem.Update.Close()
	g_SQL.PlayerTotem.Delete.Close()
	g_SQL.PlayerTotemInfo.Insert.Close()
	g_SQL.PlayerTotemInfo.Update.Close()
	g_SQL.PlayerTotemInfo.Delete.Close()
	g_SQL.PlayerTown.Insert.Close()
	g_SQL.PlayerTown.Update.Close()
	g_SQL.PlayerTown.Delete.Close()
	g_SQL.PlayerTraderRefreshState.Insert.Close()
	g_SQL.PlayerTraderRefreshState.Update.Close()
	g_SQL.PlayerTraderRefreshState.Delete.Close()
	g_SQL.PlayerTraderStoreState.Insert.Close()
	g_SQL.PlayerTraderStoreState.Update.Close()
	g_SQL.PlayerTraderStoreState.Delete.Close()
	g_SQL.PlayerUseSkill.Insert.Close()
	g_SQL.PlayerUseSkill.Update.Close()
	g_SQL.PlayerUseSkill.Delete.Close()
	g_SQL.PlayerVip.Insert.Close()
	g_SQL.PlayerVip.Update.Close()
	g_SQL.PlayerVip.Delete.Close()
}

