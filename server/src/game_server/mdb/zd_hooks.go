package mdb

var g_Hooks tableHooks

type tableHooks struct {
	GlobalAnnouncement GlobalAnnouncementHook
	GlobalArenaRank GlobalArenaRankHook
	GlobalClique GlobalCliqueHook
	GlobalCliqueBoat GlobalCliqueBoatHook
	GlobalDespairBoss GlobalDespairBossHook
	GlobalDespairBossHistory GlobalDespairBossHistoryHook
	GlobalDespairLand GlobalDespairLandHook
	GlobalDespairLandBattleRecord GlobalDespairLandBattleRecordHook
	GlobalDespairLandCamp GlobalDespairLandCampHook
	GlobalGiftCardRecord GlobalGiftCardRecordHook
	GlobalGroupBuyStatus GlobalGroupBuyStatusHook
	GlobalMail GlobalMailHook
	GlobalMailAttachments GlobalMailAttachmentsHook
	GlobalTbXxdOnlinecnt GlobalTbXxdOnlinecntHook
	Player PlayerHook
	PlayerActivity PlayerActivityHook
	PlayerAdditionQuest PlayerAdditionQuestHook
	PlayerAnySdkOrder PlayerAnySdkOrderHook
	PlayerArena PlayerArenaHook
	PlayerArenaRank PlayerArenaRankHook
	PlayerArenaRecord PlayerArenaRecordHook
	PlayerAwakenGraphic PlayerAwakenGraphicHook
	PlayerBattlePet PlayerBattlePetHook
	PlayerBattlePetGrid PlayerBattlePetGridHook
	PlayerBattlePetState PlayerBattlePetStateHook
	PlayerChestState PlayerChestStateHook
	PlayerCliqueKongfuAttrib PlayerCliqueKongfuAttribHook
	PlayerCoins PlayerCoinsHook
	PlayerCornucopia PlayerCornucopiaHook
	PlayerDailyQuest PlayerDailyQuestHook
	PlayerDailyQuestStarAwardInfo PlayerDailyQuestStarAwardInfoHook
	PlayerDailySignInRecord PlayerDailySignInRecordHook
	PlayerDailySignInState PlayerDailySignInStateHook
	PlayerDespairLandCampState PlayerDespairLandCampStateHook
	PlayerDespairLandCampStateHistory PlayerDespairLandCampStateHistoryHook
	PlayerDespairLandLevelRecord PlayerDespairLandLevelRecordHook
	PlayerDespairLandState PlayerDespairLandStateHook
	PlayerDrivingSwordEvent PlayerDrivingSwordEventHook
	PlayerDrivingSwordEventExploring PlayerDrivingSwordEventExploringHook
	PlayerDrivingSwordEventTreasure PlayerDrivingSwordEventTreasureHook
	PlayerDrivingSwordEventVisiting PlayerDrivingSwordEventVisitingHook
	PlayerDrivingSwordEventmask PlayerDrivingSwordEventmaskHook
	PlayerDrivingSwordInfo PlayerDrivingSwordInfoHook
	PlayerDrivingSwordMap PlayerDrivingSwordMapHook
	PlayerEquipment PlayerEquipmentHook
	PlayerEventAwardRecord PlayerEventAwardRecordHook
	PlayerEventDailyOnline PlayerEventDailyOnlineHook
	PlayerEventVnDailyRecharge PlayerEventVnDailyRechargeHook
	PlayerEventVnDragonBallRecord PlayerEventVnDragonBallRecordHook
	PlayerEventsFightPower PlayerEventsFightPowerHook
	PlayerEventsIngotRecord PlayerEventsIngotRecordHook
	PlayerExtendLevel PlayerExtendLevelHook
	PlayerExtendQuest PlayerExtendQuestHook
	PlayerFame PlayerFameHook
	PlayerFashion PlayerFashionHook
	PlayerFashionState PlayerFashionStateHook
	PlayerFateBoxState PlayerFateBoxStateHook
	PlayerFightNum PlayerFightNumHook
	PlayerFirstCharge PlayerFirstChargeHook
	PlayerFormation PlayerFormationHook
	PlayerFuncKey PlayerFuncKeyHook
	PlayerGhost PlayerGhostHook
	PlayerGhostEquipment PlayerGhostEquipmentHook
	PlayerGhostState PlayerGhostStateHook
	PlayerGlobalChatState PlayerGlobalChatStateHook
	PlayerGlobalCliqueBuilding PlayerGlobalCliqueBuildingHook
	PlayerGlobalCliqueBuildingQuest PlayerGlobalCliqueBuildingQuestHook
	PlayerGlobalCliqueDailyQuest PlayerGlobalCliqueDailyQuestHook
	PlayerGlobalCliqueEscort PlayerGlobalCliqueEscortHook
	PlayerGlobalCliqueEscortMessage PlayerGlobalCliqueEscortMessageHook
	PlayerGlobalCliqueHijack PlayerGlobalCliqueHijackHook
	PlayerGlobalCliqueInfo PlayerGlobalCliqueInfoHook
	PlayerGlobalCliqueKongfu PlayerGlobalCliqueKongfuHook
	PlayerGlobalFriend PlayerGlobalFriendHook
	PlayerGlobalFriendChat PlayerGlobalFriendChatHook
	PlayerGlobalFriendState PlayerGlobalFriendStateHook
	PlayerGlobalGiftCardRecord PlayerGlobalGiftCardRecordHook
	PlayerGlobalMailState PlayerGlobalMailStateHook
	PlayerGlobalWorldChatState PlayerGlobalWorldChatStateHook
	PlayerHardLevel PlayerHardLevelHook
	PlayerHardLevelRecord PlayerHardLevelRecordHook
	PlayerHeart PlayerHeartHook
	PlayerHeartDraw PlayerHeartDrawHook
	PlayerHeartDrawCardRecord PlayerHeartDrawCardRecordHook
	PlayerHeartDrawWheelRecord PlayerHeartDrawWheelRecordHook
	PlayerInfo PlayerInfoHook
	PlayerIsOperated PlayerIsOperatedHook
	PlayerItem PlayerItemHook
	PlayerItemAppendix PlayerItemAppendixHook
	PlayerItemBuyback PlayerItemBuybackHook
	PlayerItemRecastState PlayerItemRecastStateHook
	PlayerLevelAwardInfo PlayerLevelAwardInfoHook
	PlayerLoginAwardRecord PlayerLoginAwardRecordHook
	PlayerMail PlayerMailHook
	PlayerMailAttachment PlayerMailAttachmentHook
	PlayerMailAttachmentLg PlayerMailAttachmentLgHook
	PlayerMailLg PlayerMailLgHook
	PlayerMeditationState PlayerMeditationStateHook
	PlayerMission PlayerMissionHook
	PlayerMissionLevel PlayerMissionLevelHook
	PlayerMissionLevelRecord PlayerMissionLevelRecordHook
	PlayerMissionLevelStateBin PlayerMissionLevelStateBinHook
	PlayerMissionRecord PlayerMissionRecordHook
	PlayerMissionStarAward PlayerMissionStarAwardHook
	PlayerMoneyTree PlayerMoneyTreeHook
	PlayerMonthCardAwardRecord PlayerMonthCardAwardRecordHook
	PlayerMonthCardInfo PlayerMonthCardInfoHook
	PlayerMultiLevelInfo PlayerMultiLevelInfoHook
	PlayerNewYearConsumeRecord PlayerNewYearConsumeRecordHook
	PlayerNpcTalkRecord PlayerNpcTalkRecordHook
	PlayerOpenedTownTreasure PlayerOpenedTownTreasureHook
	PlayerPhysical PlayerPhysicalHook
	PlayerPurchaseRecord PlayerPurchaseRecordHook
	PlayerPushNotifySwitch PlayerPushNotifySwitchHook
	PlayerPveState PlayerPveStateHook
	PlayerQqVipGift PlayerQqVipGiftHook
	PlayerQuest PlayerQuestHook
	PlayerRainbowLevel PlayerRainbowLevelHook
	PlayerRainbowLevelStateBin PlayerRainbowLevelStateBinHook
	PlayerRole PlayerRoleHook
	PlayerSealedbook PlayerSealedbookHook
	PlayerSendHeartRecord PlayerSendHeartRecordHook
	PlayerSkill PlayerSkillHook
	PlayerState PlayerStateHook
	PlayerSwordSoul PlayerSwordSoulHook
	PlayerSwordSoulEquipment PlayerSwordSoulEquipmentHook
	PlayerSwordSoulState PlayerSwordSoulStateHook
	PlayerTaoyuanBless PlayerTaoyuanBlessHook
	PlayerTaoyuanFileds PlayerTaoyuanFiledsHook
	PlayerTaoyuanHeart PlayerTaoyuanHeartHook
	PlayerTaoyuanItem PlayerTaoyuanItemHook
	PlayerTaoyuanMessage PlayerTaoyuanMessageHook
	PlayerTaoyuanProduct PlayerTaoyuanProductHook
	PlayerTaoyuanPurchaseRecord PlayerTaoyuanPurchaseRecordHook
	PlayerTaoyuanQuest PlayerTaoyuanQuestHook
	PlayerTbXxdRoleinfo PlayerTbXxdRoleinfoHook
	PlayerTeamInfo PlayerTeamInfoHook
	PlayerTotem PlayerTotemHook
	PlayerTotemInfo PlayerTotemInfoHook
	PlayerTown PlayerTownHook
	PlayerTraderRefreshState PlayerTraderRefreshStateHook
	PlayerTraderStoreState PlayerTraderStoreStateHook
	PlayerUseSkill PlayerUseSkillHook
	PlayerVip PlayerVipHook
}

type hooksOP int
const Hook hooksOP = 0
func (h hooksOP) GlobalAnnouncement(hook GlobalAnnouncementHook) {
	if g_Hooks.GlobalAnnouncement != nil {
		panic("duplicate hook GlobalAnnouncement")
	}
	g_Hooks.GlobalAnnouncement = hook
}
func (h hooksOP) GlobalArenaRank(hook GlobalArenaRankHook) {
	if g_Hooks.GlobalArenaRank != nil {
		panic("duplicate hook GlobalArenaRank")
	}
	g_Hooks.GlobalArenaRank = hook
}
func (h hooksOP) GlobalClique(hook GlobalCliqueHook) {
	if g_Hooks.GlobalClique != nil {
		panic("duplicate hook GlobalClique")
	}
	g_Hooks.GlobalClique = hook
}
func (h hooksOP) GlobalCliqueBoat(hook GlobalCliqueBoatHook) {
	if g_Hooks.GlobalCliqueBoat != nil {
		panic("duplicate hook GlobalCliqueBoat")
	}
	g_Hooks.GlobalCliqueBoat = hook
}
func (h hooksOP) GlobalDespairBoss(hook GlobalDespairBossHook) {
	if g_Hooks.GlobalDespairBoss != nil {
		panic("duplicate hook GlobalDespairBoss")
	}
	g_Hooks.GlobalDespairBoss = hook
}
func (h hooksOP) GlobalDespairBossHistory(hook GlobalDespairBossHistoryHook) {
	if g_Hooks.GlobalDespairBossHistory != nil {
		panic("duplicate hook GlobalDespairBossHistory")
	}
	g_Hooks.GlobalDespairBossHistory = hook
}
func (h hooksOP) GlobalDespairLand(hook GlobalDespairLandHook) {
	if g_Hooks.GlobalDespairLand != nil {
		panic("duplicate hook GlobalDespairLand")
	}
	g_Hooks.GlobalDespairLand = hook
}
func (h hooksOP) GlobalDespairLandBattleRecord(hook GlobalDespairLandBattleRecordHook) {
	if g_Hooks.GlobalDespairLandBattleRecord != nil {
		panic("duplicate hook GlobalDespairLandBattleRecord")
	}
	g_Hooks.GlobalDespairLandBattleRecord = hook
}
func (h hooksOP) GlobalDespairLandCamp(hook GlobalDespairLandCampHook) {
	if g_Hooks.GlobalDespairLandCamp != nil {
		panic("duplicate hook GlobalDespairLandCamp")
	}
	g_Hooks.GlobalDespairLandCamp = hook
}
func (h hooksOP) GlobalGiftCardRecord(hook GlobalGiftCardRecordHook) {
	if g_Hooks.GlobalGiftCardRecord != nil {
		panic("duplicate hook GlobalGiftCardRecord")
	}
	g_Hooks.GlobalGiftCardRecord = hook
}
func (h hooksOP) GlobalGroupBuyStatus(hook GlobalGroupBuyStatusHook) {
	if g_Hooks.GlobalGroupBuyStatus != nil {
		panic("duplicate hook GlobalGroupBuyStatus")
	}
	g_Hooks.GlobalGroupBuyStatus = hook
}
func (h hooksOP) GlobalMail(hook GlobalMailHook) {
	if g_Hooks.GlobalMail != nil {
		panic("duplicate hook GlobalMail")
	}
	g_Hooks.GlobalMail = hook
}
func (h hooksOP) GlobalMailAttachments(hook GlobalMailAttachmentsHook) {
	if g_Hooks.GlobalMailAttachments != nil {
		panic("duplicate hook GlobalMailAttachments")
	}
	g_Hooks.GlobalMailAttachments = hook
}
func (h hooksOP) GlobalTbXxdOnlinecnt(hook GlobalTbXxdOnlinecntHook) {
	if g_Hooks.GlobalTbXxdOnlinecnt != nil {
		panic("duplicate hook GlobalTbXxdOnlinecnt")
	}
	g_Hooks.GlobalTbXxdOnlinecnt = hook
}
func (h hooksOP) Player(hook PlayerHook) {
	if g_Hooks.Player != nil {
		panic("duplicate hook Player")
	}
	g_Hooks.Player = hook
}
func (h hooksOP) PlayerActivity(hook PlayerActivityHook) {
	if g_Hooks.PlayerActivity != nil {
		panic("duplicate hook PlayerActivity")
	}
	g_Hooks.PlayerActivity = hook
}
func (h hooksOP) PlayerAdditionQuest(hook PlayerAdditionQuestHook) {
	if g_Hooks.PlayerAdditionQuest != nil {
		panic("duplicate hook PlayerAdditionQuest")
	}
	g_Hooks.PlayerAdditionQuest = hook
}
func (h hooksOP) PlayerAnySdkOrder(hook PlayerAnySdkOrderHook) {
	if g_Hooks.PlayerAnySdkOrder != nil {
		panic("duplicate hook PlayerAnySdkOrder")
	}
	g_Hooks.PlayerAnySdkOrder = hook
}
func (h hooksOP) PlayerArena(hook PlayerArenaHook) {
	if g_Hooks.PlayerArena != nil {
		panic("duplicate hook PlayerArena")
	}
	g_Hooks.PlayerArena = hook
}
func (h hooksOP) PlayerArenaRank(hook PlayerArenaRankHook) {
	if g_Hooks.PlayerArenaRank != nil {
		panic("duplicate hook PlayerArenaRank")
	}
	g_Hooks.PlayerArenaRank = hook
}
func (h hooksOP) PlayerArenaRecord(hook PlayerArenaRecordHook) {
	if g_Hooks.PlayerArenaRecord != nil {
		panic("duplicate hook PlayerArenaRecord")
	}
	g_Hooks.PlayerArenaRecord = hook
}
func (h hooksOP) PlayerAwakenGraphic(hook PlayerAwakenGraphicHook) {
	if g_Hooks.PlayerAwakenGraphic != nil {
		panic("duplicate hook PlayerAwakenGraphic")
	}
	g_Hooks.PlayerAwakenGraphic = hook
}
func (h hooksOP) PlayerBattlePet(hook PlayerBattlePetHook) {
	if g_Hooks.PlayerBattlePet != nil {
		panic("duplicate hook PlayerBattlePet")
	}
	g_Hooks.PlayerBattlePet = hook
}
func (h hooksOP) PlayerBattlePetGrid(hook PlayerBattlePetGridHook) {
	if g_Hooks.PlayerBattlePetGrid != nil {
		panic("duplicate hook PlayerBattlePetGrid")
	}
	g_Hooks.PlayerBattlePetGrid = hook
}
func (h hooksOP) PlayerBattlePetState(hook PlayerBattlePetStateHook) {
	if g_Hooks.PlayerBattlePetState != nil {
		panic("duplicate hook PlayerBattlePetState")
	}
	g_Hooks.PlayerBattlePetState = hook
}
func (h hooksOP) PlayerChestState(hook PlayerChestStateHook) {
	if g_Hooks.PlayerChestState != nil {
		panic("duplicate hook PlayerChestState")
	}
	g_Hooks.PlayerChestState = hook
}
func (h hooksOP) PlayerCliqueKongfuAttrib(hook PlayerCliqueKongfuAttribHook) {
	if g_Hooks.PlayerCliqueKongfuAttrib != nil {
		panic("duplicate hook PlayerCliqueKongfuAttrib")
	}
	g_Hooks.PlayerCliqueKongfuAttrib = hook
}
func (h hooksOP) PlayerCoins(hook PlayerCoinsHook) {
	if g_Hooks.PlayerCoins != nil {
		panic("duplicate hook PlayerCoins")
	}
	g_Hooks.PlayerCoins = hook
}
func (h hooksOP) PlayerCornucopia(hook PlayerCornucopiaHook) {
	if g_Hooks.PlayerCornucopia != nil {
		panic("duplicate hook PlayerCornucopia")
	}
	g_Hooks.PlayerCornucopia = hook
}
func (h hooksOP) PlayerDailyQuest(hook PlayerDailyQuestHook) {
	if g_Hooks.PlayerDailyQuest != nil {
		panic("duplicate hook PlayerDailyQuest")
	}
	g_Hooks.PlayerDailyQuest = hook
}
func (h hooksOP) PlayerDailyQuestStarAwardInfo(hook PlayerDailyQuestStarAwardInfoHook) {
	if g_Hooks.PlayerDailyQuestStarAwardInfo != nil {
		panic("duplicate hook PlayerDailyQuestStarAwardInfo")
	}
	g_Hooks.PlayerDailyQuestStarAwardInfo = hook
}
func (h hooksOP) PlayerDailySignInRecord(hook PlayerDailySignInRecordHook) {
	if g_Hooks.PlayerDailySignInRecord != nil {
		panic("duplicate hook PlayerDailySignInRecord")
	}
	g_Hooks.PlayerDailySignInRecord = hook
}
func (h hooksOP) PlayerDailySignInState(hook PlayerDailySignInStateHook) {
	if g_Hooks.PlayerDailySignInState != nil {
		panic("duplicate hook PlayerDailySignInState")
	}
	g_Hooks.PlayerDailySignInState = hook
}
func (h hooksOP) PlayerDespairLandCampState(hook PlayerDespairLandCampStateHook) {
	if g_Hooks.PlayerDespairLandCampState != nil {
		panic("duplicate hook PlayerDespairLandCampState")
	}
	g_Hooks.PlayerDespairLandCampState = hook
}
func (h hooksOP) PlayerDespairLandCampStateHistory(hook PlayerDespairLandCampStateHistoryHook) {
	if g_Hooks.PlayerDespairLandCampStateHistory != nil {
		panic("duplicate hook PlayerDespairLandCampStateHistory")
	}
	g_Hooks.PlayerDespairLandCampStateHistory = hook
}
func (h hooksOP) PlayerDespairLandLevelRecord(hook PlayerDespairLandLevelRecordHook) {
	if g_Hooks.PlayerDespairLandLevelRecord != nil {
		panic("duplicate hook PlayerDespairLandLevelRecord")
	}
	g_Hooks.PlayerDespairLandLevelRecord = hook
}
func (h hooksOP) PlayerDespairLandState(hook PlayerDespairLandStateHook) {
	if g_Hooks.PlayerDespairLandState != nil {
		panic("duplicate hook PlayerDespairLandState")
	}
	g_Hooks.PlayerDespairLandState = hook
}
func (h hooksOP) PlayerDrivingSwordEvent(hook PlayerDrivingSwordEventHook) {
	if g_Hooks.PlayerDrivingSwordEvent != nil {
		panic("duplicate hook PlayerDrivingSwordEvent")
	}
	g_Hooks.PlayerDrivingSwordEvent = hook
}
func (h hooksOP) PlayerDrivingSwordEventExploring(hook PlayerDrivingSwordEventExploringHook) {
	if g_Hooks.PlayerDrivingSwordEventExploring != nil {
		panic("duplicate hook PlayerDrivingSwordEventExploring")
	}
	g_Hooks.PlayerDrivingSwordEventExploring = hook
}
func (h hooksOP) PlayerDrivingSwordEventTreasure(hook PlayerDrivingSwordEventTreasureHook) {
	if g_Hooks.PlayerDrivingSwordEventTreasure != nil {
		panic("duplicate hook PlayerDrivingSwordEventTreasure")
	}
	g_Hooks.PlayerDrivingSwordEventTreasure = hook
}
func (h hooksOP) PlayerDrivingSwordEventVisiting(hook PlayerDrivingSwordEventVisitingHook) {
	if g_Hooks.PlayerDrivingSwordEventVisiting != nil {
		panic("duplicate hook PlayerDrivingSwordEventVisiting")
	}
	g_Hooks.PlayerDrivingSwordEventVisiting = hook
}
func (h hooksOP) PlayerDrivingSwordEventmask(hook PlayerDrivingSwordEventmaskHook) {
	if g_Hooks.PlayerDrivingSwordEventmask != nil {
		panic("duplicate hook PlayerDrivingSwordEventmask")
	}
	g_Hooks.PlayerDrivingSwordEventmask = hook
}
func (h hooksOP) PlayerDrivingSwordInfo(hook PlayerDrivingSwordInfoHook) {
	if g_Hooks.PlayerDrivingSwordInfo != nil {
		panic("duplicate hook PlayerDrivingSwordInfo")
	}
	g_Hooks.PlayerDrivingSwordInfo = hook
}
func (h hooksOP) PlayerDrivingSwordMap(hook PlayerDrivingSwordMapHook) {
	if g_Hooks.PlayerDrivingSwordMap != nil {
		panic("duplicate hook PlayerDrivingSwordMap")
	}
	g_Hooks.PlayerDrivingSwordMap = hook
}
func (h hooksOP) PlayerEquipment(hook PlayerEquipmentHook) {
	if g_Hooks.PlayerEquipment != nil {
		panic("duplicate hook PlayerEquipment")
	}
	g_Hooks.PlayerEquipment = hook
}
func (h hooksOP) PlayerEventAwardRecord(hook PlayerEventAwardRecordHook) {
	if g_Hooks.PlayerEventAwardRecord != nil {
		panic("duplicate hook PlayerEventAwardRecord")
	}
	g_Hooks.PlayerEventAwardRecord = hook
}
func (h hooksOP) PlayerEventDailyOnline(hook PlayerEventDailyOnlineHook) {
	if g_Hooks.PlayerEventDailyOnline != nil {
		panic("duplicate hook PlayerEventDailyOnline")
	}
	g_Hooks.PlayerEventDailyOnline = hook
}
func (h hooksOP) PlayerEventVnDailyRecharge(hook PlayerEventVnDailyRechargeHook) {
	if g_Hooks.PlayerEventVnDailyRecharge != nil {
		panic("duplicate hook PlayerEventVnDailyRecharge")
	}
	g_Hooks.PlayerEventVnDailyRecharge = hook
}
func (h hooksOP) PlayerEventVnDragonBallRecord(hook PlayerEventVnDragonBallRecordHook) {
	if g_Hooks.PlayerEventVnDragonBallRecord != nil {
		panic("duplicate hook PlayerEventVnDragonBallRecord")
	}
	g_Hooks.PlayerEventVnDragonBallRecord = hook
}
func (h hooksOP) PlayerEventsFightPower(hook PlayerEventsFightPowerHook) {
	if g_Hooks.PlayerEventsFightPower != nil {
		panic("duplicate hook PlayerEventsFightPower")
	}
	g_Hooks.PlayerEventsFightPower = hook
}
func (h hooksOP) PlayerEventsIngotRecord(hook PlayerEventsIngotRecordHook) {
	if g_Hooks.PlayerEventsIngotRecord != nil {
		panic("duplicate hook PlayerEventsIngotRecord")
	}
	g_Hooks.PlayerEventsIngotRecord = hook
}
func (h hooksOP) PlayerExtendLevel(hook PlayerExtendLevelHook) {
	if g_Hooks.PlayerExtendLevel != nil {
		panic("duplicate hook PlayerExtendLevel")
	}
	g_Hooks.PlayerExtendLevel = hook
}
func (h hooksOP) PlayerExtendQuest(hook PlayerExtendQuestHook) {
	if g_Hooks.PlayerExtendQuest != nil {
		panic("duplicate hook PlayerExtendQuest")
	}
	g_Hooks.PlayerExtendQuest = hook
}
func (h hooksOP) PlayerFame(hook PlayerFameHook) {
	if g_Hooks.PlayerFame != nil {
		panic("duplicate hook PlayerFame")
	}
	g_Hooks.PlayerFame = hook
}
func (h hooksOP) PlayerFashion(hook PlayerFashionHook) {
	if g_Hooks.PlayerFashion != nil {
		panic("duplicate hook PlayerFashion")
	}
	g_Hooks.PlayerFashion = hook
}
func (h hooksOP) PlayerFashionState(hook PlayerFashionStateHook) {
	if g_Hooks.PlayerFashionState != nil {
		panic("duplicate hook PlayerFashionState")
	}
	g_Hooks.PlayerFashionState = hook
}
func (h hooksOP) PlayerFateBoxState(hook PlayerFateBoxStateHook) {
	if g_Hooks.PlayerFateBoxState != nil {
		panic("duplicate hook PlayerFateBoxState")
	}
	g_Hooks.PlayerFateBoxState = hook
}
func (h hooksOP) PlayerFightNum(hook PlayerFightNumHook) {
	if g_Hooks.PlayerFightNum != nil {
		panic("duplicate hook PlayerFightNum")
	}
	g_Hooks.PlayerFightNum = hook
}
func (h hooksOP) PlayerFirstCharge(hook PlayerFirstChargeHook) {
	if g_Hooks.PlayerFirstCharge != nil {
		panic("duplicate hook PlayerFirstCharge")
	}
	g_Hooks.PlayerFirstCharge = hook
}
func (h hooksOP) PlayerFormation(hook PlayerFormationHook) {
	if g_Hooks.PlayerFormation != nil {
		panic("duplicate hook PlayerFormation")
	}
	g_Hooks.PlayerFormation = hook
}
func (h hooksOP) PlayerFuncKey(hook PlayerFuncKeyHook) {
	if g_Hooks.PlayerFuncKey != nil {
		panic("duplicate hook PlayerFuncKey")
	}
	g_Hooks.PlayerFuncKey = hook
}
func (h hooksOP) PlayerGhost(hook PlayerGhostHook) {
	if g_Hooks.PlayerGhost != nil {
		panic("duplicate hook PlayerGhost")
	}
	g_Hooks.PlayerGhost = hook
}
func (h hooksOP) PlayerGhostEquipment(hook PlayerGhostEquipmentHook) {
	if g_Hooks.PlayerGhostEquipment != nil {
		panic("duplicate hook PlayerGhostEquipment")
	}
	g_Hooks.PlayerGhostEquipment = hook
}
func (h hooksOP) PlayerGhostState(hook PlayerGhostStateHook) {
	if g_Hooks.PlayerGhostState != nil {
		panic("duplicate hook PlayerGhostState")
	}
	g_Hooks.PlayerGhostState = hook
}
func (h hooksOP) PlayerGlobalChatState(hook PlayerGlobalChatStateHook) {
	if g_Hooks.PlayerGlobalChatState != nil {
		panic("duplicate hook PlayerGlobalChatState")
	}
	g_Hooks.PlayerGlobalChatState = hook
}
func (h hooksOP) PlayerGlobalCliqueBuilding(hook PlayerGlobalCliqueBuildingHook) {
	if g_Hooks.PlayerGlobalCliqueBuilding != nil {
		panic("duplicate hook PlayerGlobalCliqueBuilding")
	}
	g_Hooks.PlayerGlobalCliqueBuilding = hook
}
func (h hooksOP) PlayerGlobalCliqueBuildingQuest(hook PlayerGlobalCliqueBuildingQuestHook) {
	if g_Hooks.PlayerGlobalCliqueBuildingQuest != nil {
		panic("duplicate hook PlayerGlobalCliqueBuildingQuest")
	}
	g_Hooks.PlayerGlobalCliqueBuildingQuest = hook
}
func (h hooksOP) PlayerGlobalCliqueDailyQuest(hook PlayerGlobalCliqueDailyQuestHook) {
	if g_Hooks.PlayerGlobalCliqueDailyQuest != nil {
		panic("duplicate hook PlayerGlobalCliqueDailyQuest")
	}
	g_Hooks.PlayerGlobalCliqueDailyQuest = hook
}
func (h hooksOP) PlayerGlobalCliqueEscort(hook PlayerGlobalCliqueEscortHook) {
	if g_Hooks.PlayerGlobalCliqueEscort != nil {
		panic("duplicate hook PlayerGlobalCliqueEscort")
	}
	g_Hooks.PlayerGlobalCliqueEscort = hook
}
func (h hooksOP) PlayerGlobalCliqueEscortMessage(hook PlayerGlobalCliqueEscortMessageHook) {
	if g_Hooks.PlayerGlobalCliqueEscortMessage != nil {
		panic("duplicate hook PlayerGlobalCliqueEscortMessage")
	}
	g_Hooks.PlayerGlobalCliqueEscortMessage = hook
}
func (h hooksOP) PlayerGlobalCliqueHijack(hook PlayerGlobalCliqueHijackHook) {
	if g_Hooks.PlayerGlobalCliqueHijack != nil {
		panic("duplicate hook PlayerGlobalCliqueHijack")
	}
	g_Hooks.PlayerGlobalCliqueHijack = hook
}
func (h hooksOP) PlayerGlobalCliqueInfo(hook PlayerGlobalCliqueInfoHook) {
	if g_Hooks.PlayerGlobalCliqueInfo != nil {
		panic("duplicate hook PlayerGlobalCliqueInfo")
	}
	g_Hooks.PlayerGlobalCliqueInfo = hook
}
func (h hooksOP) PlayerGlobalCliqueKongfu(hook PlayerGlobalCliqueKongfuHook) {
	if g_Hooks.PlayerGlobalCliqueKongfu != nil {
		panic("duplicate hook PlayerGlobalCliqueKongfu")
	}
	g_Hooks.PlayerGlobalCliqueKongfu = hook
}
func (h hooksOP) PlayerGlobalFriend(hook PlayerGlobalFriendHook) {
	if g_Hooks.PlayerGlobalFriend != nil {
		panic("duplicate hook PlayerGlobalFriend")
	}
	g_Hooks.PlayerGlobalFriend = hook
}
func (h hooksOP) PlayerGlobalFriendChat(hook PlayerGlobalFriendChatHook) {
	if g_Hooks.PlayerGlobalFriendChat != nil {
		panic("duplicate hook PlayerGlobalFriendChat")
	}
	g_Hooks.PlayerGlobalFriendChat = hook
}
func (h hooksOP) PlayerGlobalFriendState(hook PlayerGlobalFriendStateHook) {
	if g_Hooks.PlayerGlobalFriendState != nil {
		panic("duplicate hook PlayerGlobalFriendState")
	}
	g_Hooks.PlayerGlobalFriendState = hook
}
func (h hooksOP) PlayerGlobalGiftCardRecord(hook PlayerGlobalGiftCardRecordHook) {
	if g_Hooks.PlayerGlobalGiftCardRecord != nil {
		panic("duplicate hook PlayerGlobalGiftCardRecord")
	}
	g_Hooks.PlayerGlobalGiftCardRecord = hook
}
func (h hooksOP) PlayerGlobalMailState(hook PlayerGlobalMailStateHook) {
	if g_Hooks.PlayerGlobalMailState != nil {
		panic("duplicate hook PlayerGlobalMailState")
	}
	g_Hooks.PlayerGlobalMailState = hook
}
func (h hooksOP) PlayerGlobalWorldChatState(hook PlayerGlobalWorldChatStateHook) {
	if g_Hooks.PlayerGlobalWorldChatState != nil {
		panic("duplicate hook PlayerGlobalWorldChatState")
	}
	g_Hooks.PlayerGlobalWorldChatState = hook
}
func (h hooksOP) PlayerHardLevel(hook PlayerHardLevelHook) {
	if g_Hooks.PlayerHardLevel != nil {
		panic("duplicate hook PlayerHardLevel")
	}
	g_Hooks.PlayerHardLevel = hook
}
func (h hooksOP) PlayerHardLevelRecord(hook PlayerHardLevelRecordHook) {
	if g_Hooks.PlayerHardLevelRecord != nil {
		panic("duplicate hook PlayerHardLevelRecord")
	}
	g_Hooks.PlayerHardLevelRecord = hook
}
func (h hooksOP) PlayerHeart(hook PlayerHeartHook) {
	if g_Hooks.PlayerHeart != nil {
		panic("duplicate hook PlayerHeart")
	}
	g_Hooks.PlayerHeart = hook
}
func (h hooksOP) PlayerHeartDraw(hook PlayerHeartDrawHook) {
	if g_Hooks.PlayerHeartDraw != nil {
		panic("duplicate hook PlayerHeartDraw")
	}
	g_Hooks.PlayerHeartDraw = hook
}
func (h hooksOP) PlayerHeartDrawCardRecord(hook PlayerHeartDrawCardRecordHook) {
	if g_Hooks.PlayerHeartDrawCardRecord != nil {
		panic("duplicate hook PlayerHeartDrawCardRecord")
	}
	g_Hooks.PlayerHeartDrawCardRecord = hook
}
func (h hooksOP) PlayerHeartDrawWheelRecord(hook PlayerHeartDrawWheelRecordHook) {
	if g_Hooks.PlayerHeartDrawWheelRecord != nil {
		panic("duplicate hook PlayerHeartDrawWheelRecord")
	}
	g_Hooks.PlayerHeartDrawWheelRecord = hook
}
func (h hooksOP) PlayerInfo(hook PlayerInfoHook) {
	if g_Hooks.PlayerInfo != nil {
		panic("duplicate hook PlayerInfo")
	}
	g_Hooks.PlayerInfo = hook
}
func (h hooksOP) PlayerIsOperated(hook PlayerIsOperatedHook) {
	if g_Hooks.PlayerIsOperated != nil {
		panic("duplicate hook PlayerIsOperated")
	}
	g_Hooks.PlayerIsOperated = hook
}
func (h hooksOP) PlayerItem(hook PlayerItemHook) {
	if g_Hooks.PlayerItem != nil {
		panic("duplicate hook PlayerItem")
	}
	g_Hooks.PlayerItem = hook
}
func (h hooksOP) PlayerItemAppendix(hook PlayerItemAppendixHook) {
	if g_Hooks.PlayerItemAppendix != nil {
		panic("duplicate hook PlayerItemAppendix")
	}
	g_Hooks.PlayerItemAppendix = hook
}
func (h hooksOP) PlayerItemBuyback(hook PlayerItemBuybackHook) {
	if g_Hooks.PlayerItemBuyback != nil {
		panic("duplicate hook PlayerItemBuyback")
	}
	g_Hooks.PlayerItemBuyback = hook
}
func (h hooksOP) PlayerItemRecastState(hook PlayerItemRecastStateHook) {
	if g_Hooks.PlayerItemRecastState != nil {
		panic("duplicate hook PlayerItemRecastState")
	}
	g_Hooks.PlayerItemRecastState = hook
}
func (h hooksOP) PlayerLevelAwardInfo(hook PlayerLevelAwardInfoHook) {
	if g_Hooks.PlayerLevelAwardInfo != nil {
		panic("duplicate hook PlayerLevelAwardInfo")
	}
	g_Hooks.PlayerLevelAwardInfo = hook
}
func (h hooksOP) PlayerLoginAwardRecord(hook PlayerLoginAwardRecordHook) {
	if g_Hooks.PlayerLoginAwardRecord != nil {
		panic("duplicate hook PlayerLoginAwardRecord")
	}
	g_Hooks.PlayerLoginAwardRecord = hook
}
func (h hooksOP) PlayerMail(hook PlayerMailHook) {
	if g_Hooks.PlayerMail != nil {
		panic("duplicate hook PlayerMail")
	}
	g_Hooks.PlayerMail = hook
}
func (h hooksOP) PlayerMailAttachment(hook PlayerMailAttachmentHook) {
	if g_Hooks.PlayerMailAttachment != nil {
		panic("duplicate hook PlayerMailAttachment")
	}
	g_Hooks.PlayerMailAttachment = hook
}
func (h hooksOP) PlayerMailAttachmentLg(hook PlayerMailAttachmentLgHook) {
	if g_Hooks.PlayerMailAttachmentLg != nil {
		panic("duplicate hook PlayerMailAttachmentLg")
	}
	g_Hooks.PlayerMailAttachmentLg = hook
}
func (h hooksOP) PlayerMailLg(hook PlayerMailLgHook) {
	if g_Hooks.PlayerMailLg != nil {
		panic("duplicate hook PlayerMailLg")
	}
	g_Hooks.PlayerMailLg = hook
}
func (h hooksOP) PlayerMeditationState(hook PlayerMeditationStateHook) {
	if g_Hooks.PlayerMeditationState != nil {
		panic("duplicate hook PlayerMeditationState")
	}
	g_Hooks.PlayerMeditationState = hook
}
func (h hooksOP) PlayerMission(hook PlayerMissionHook) {
	if g_Hooks.PlayerMission != nil {
		panic("duplicate hook PlayerMission")
	}
	g_Hooks.PlayerMission = hook
}
func (h hooksOP) PlayerMissionLevel(hook PlayerMissionLevelHook) {
	if g_Hooks.PlayerMissionLevel != nil {
		panic("duplicate hook PlayerMissionLevel")
	}
	g_Hooks.PlayerMissionLevel = hook
}
func (h hooksOP) PlayerMissionLevelRecord(hook PlayerMissionLevelRecordHook) {
	if g_Hooks.PlayerMissionLevelRecord != nil {
		panic("duplicate hook PlayerMissionLevelRecord")
	}
	g_Hooks.PlayerMissionLevelRecord = hook
}
func (h hooksOP) PlayerMissionLevelStateBin(hook PlayerMissionLevelStateBinHook) {
	if g_Hooks.PlayerMissionLevelStateBin != nil {
		panic("duplicate hook PlayerMissionLevelStateBin")
	}
	g_Hooks.PlayerMissionLevelStateBin = hook
}
func (h hooksOP) PlayerMissionRecord(hook PlayerMissionRecordHook) {
	if g_Hooks.PlayerMissionRecord != nil {
		panic("duplicate hook PlayerMissionRecord")
	}
	g_Hooks.PlayerMissionRecord = hook
}
func (h hooksOP) PlayerMissionStarAward(hook PlayerMissionStarAwardHook) {
	if g_Hooks.PlayerMissionStarAward != nil {
		panic("duplicate hook PlayerMissionStarAward")
	}
	g_Hooks.PlayerMissionStarAward = hook
}
func (h hooksOP) PlayerMoneyTree(hook PlayerMoneyTreeHook) {
	if g_Hooks.PlayerMoneyTree != nil {
		panic("duplicate hook PlayerMoneyTree")
	}
	g_Hooks.PlayerMoneyTree = hook
}
func (h hooksOP) PlayerMonthCardAwardRecord(hook PlayerMonthCardAwardRecordHook) {
	if g_Hooks.PlayerMonthCardAwardRecord != nil {
		panic("duplicate hook PlayerMonthCardAwardRecord")
	}
	g_Hooks.PlayerMonthCardAwardRecord = hook
}
func (h hooksOP) PlayerMonthCardInfo(hook PlayerMonthCardInfoHook) {
	if g_Hooks.PlayerMonthCardInfo != nil {
		panic("duplicate hook PlayerMonthCardInfo")
	}
	g_Hooks.PlayerMonthCardInfo = hook
}
func (h hooksOP) PlayerMultiLevelInfo(hook PlayerMultiLevelInfoHook) {
	if g_Hooks.PlayerMultiLevelInfo != nil {
		panic("duplicate hook PlayerMultiLevelInfo")
	}
	g_Hooks.PlayerMultiLevelInfo = hook
}
func (h hooksOP) PlayerNewYearConsumeRecord(hook PlayerNewYearConsumeRecordHook) {
	if g_Hooks.PlayerNewYearConsumeRecord != nil {
		panic("duplicate hook PlayerNewYearConsumeRecord")
	}
	g_Hooks.PlayerNewYearConsumeRecord = hook
}
func (h hooksOP) PlayerNpcTalkRecord(hook PlayerNpcTalkRecordHook) {
	if g_Hooks.PlayerNpcTalkRecord != nil {
		panic("duplicate hook PlayerNpcTalkRecord")
	}
	g_Hooks.PlayerNpcTalkRecord = hook
}
func (h hooksOP) PlayerOpenedTownTreasure(hook PlayerOpenedTownTreasureHook) {
	if g_Hooks.PlayerOpenedTownTreasure != nil {
		panic("duplicate hook PlayerOpenedTownTreasure")
	}
	g_Hooks.PlayerOpenedTownTreasure = hook
}
func (h hooksOP) PlayerPhysical(hook PlayerPhysicalHook) {
	if g_Hooks.PlayerPhysical != nil {
		panic("duplicate hook PlayerPhysical")
	}
	g_Hooks.PlayerPhysical = hook
}
func (h hooksOP) PlayerPurchaseRecord(hook PlayerPurchaseRecordHook) {
	if g_Hooks.PlayerPurchaseRecord != nil {
		panic("duplicate hook PlayerPurchaseRecord")
	}
	g_Hooks.PlayerPurchaseRecord = hook
}
func (h hooksOP) PlayerPushNotifySwitch(hook PlayerPushNotifySwitchHook) {
	if g_Hooks.PlayerPushNotifySwitch != nil {
		panic("duplicate hook PlayerPushNotifySwitch")
	}
	g_Hooks.PlayerPushNotifySwitch = hook
}
func (h hooksOP) PlayerPveState(hook PlayerPveStateHook) {
	if g_Hooks.PlayerPveState != nil {
		panic("duplicate hook PlayerPveState")
	}
	g_Hooks.PlayerPveState = hook
}
func (h hooksOP) PlayerQqVipGift(hook PlayerQqVipGiftHook) {
	if g_Hooks.PlayerQqVipGift != nil {
		panic("duplicate hook PlayerQqVipGift")
	}
	g_Hooks.PlayerQqVipGift = hook
}
func (h hooksOP) PlayerQuest(hook PlayerQuestHook) {
	if g_Hooks.PlayerQuest != nil {
		panic("duplicate hook PlayerQuest")
	}
	g_Hooks.PlayerQuest = hook
}
func (h hooksOP) PlayerRainbowLevel(hook PlayerRainbowLevelHook) {
	if g_Hooks.PlayerRainbowLevel != nil {
		panic("duplicate hook PlayerRainbowLevel")
	}
	g_Hooks.PlayerRainbowLevel = hook
}
func (h hooksOP) PlayerRainbowLevelStateBin(hook PlayerRainbowLevelStateBinHook) {
	if g_Hooks.PlayerRainbowLevelStateBin != nil {
		panic("duplicate hook PlayerRainbowLevelStateBin")
	}
	g_Hooks.PlayerRainbowLevelStateBin = hook
}
func (h hooksOP) PlayerRole(hook PlayerRoleHook) {
	if g_Hooks.PlayerRole != nil {
		panic("duplicate hook PlayerRole")
	}
	g_Hooks.PlayerRole = hook
}
func (h hooksOP) PlayerSealedbook(hook PlayerSealedbookHook) {
	if g_Hooks.PlayerSealedbook != nil {
		panic("duplicate hook PlayerSealedbook")
	}
	g_Hooks.PlayerSealedbook = hook
}
func (h hooksOP) PlayerSendHeartRecord(hook PlayerSendHeartRecordHook) {
	if g_Hooks.PlayerSendHeartRecord != nil {
		panic("duplicate hook PlayerSendHeartRecord")
	}
	g_Hooks.PlayerSendHeartRecord = hook
}
func (h hooksOP) PlayerSkill(hook PlayerSkillHook) {
	if g_Hooks.PlayerSkill != nil {
		panic("duplicate hook PlayerSkill")
	}
	g_Hooks.PlayerSkill = hook
}
func (h hooksOP) PlayerState(hook PlayerStateHook) {
	if g_Hooks.PlayerState != nil {
		panic("duplicate hook PlayerState")
	}
	g_Hooks.PlayerState = hook
}
func (h hooksOP) PlayerSwordSoul(hook PlayerSwordSoulHook) {
	if g_Hooks.PlayerSwordSoul != nil {
		panic("duplicate hook PlayerSwordSoul")
	}
	g_Hooks.PlayerSwordSoul = hook
}
func (h hooksOP) PlayerSwordSoulEquipment(hook PlayerSwordSoulEquipmentHook) {
	if g_Hooks.PlayerSwordSoulEquipment != nil {
		panic("duplicate hook PlayerSwordSoulEquipment")
	}
	g_Hooks.PlayerSwordSoulEquipment = hook
}
func (h hooksOP) PlayerSwordSoulState(hook PlayerSwordSoulStateHook) {
	if g_Hooks.PlayerSwordSoulState != nil {
		panic("duplicate hook PlayerSwordSoulState")
	}
	g_Hooks.PlayerSwordSoulState = hook
}
func (h hooksOP) PlayerTaoyuanBless(hook PlayerTaoyuanBlessHook) {
	if g_Hooks.PlayerTaoyuanBless != nil {
		panic("duplicate hook PlayerTaoyuanBless")
	}
	g_Hooks.PlayerTaoyuanBless = hook
}
func (h hooksOP) PlayerTaoyuanFileds(hook PlayerTaoyuanFiledsHook) {
	if g_Hooks.PlayerTaoyuanFileds != nil {
		panic("duplicate hook PlayerTaoyuanFileds")
	}
	g_Hooks.PlayerTaoyuanFileds = hook
}
func (h hooksOP) PlayerTaoyuanHeart(hook PlayerTaoyuanHeartHook) {
	if g_Hooks.PlayerTaoyuanHeart != nil {
		panic("duplicate hook PlayerTaoyuanHeart")
	}
	g_Hooks.PlayerTaoyuanHeart = hook
}
func (h hooksOP) PlayerTaoyuanItem(hook PlayerTaoyuanItemHook) {
	if g_Hooks.PlayerTaoyuanItem != nil {
		panic("duplicate hook PlayerTaoyuanItem")
	}
	g_Hooks.PlayerTaoyuanItem = hook
}
func (h hooksOP) PlayerTaoyuanMessage(hook PlayerTaoyuanMessageHook) {
	if g_Hooks.PlayerTaoyuanMessage != nil {
		panic("duplicate hook PlayerTaoyuanMessage")
	}
	g_Hooks.PlayerTaoyuanMessage = hook
}
func (h hooksOP) PlayerTaoyuanProduct(hook PlayerTaoyuanProductHook) {
	if g_Hooks.PlayerTaoyuanProduct != nil {
		panic("duplicate hook PlayerTaoyuanProduct")
	}
	g_Hooks.PlayerTaoyuanProduct = hook
}
func (h hooksOP) PlayerTaoyuanPurchaseRecord(hook PlayerTaoyuanPurchaseRecordHook) {
	if g_Hooks.PlayerTaoyuanPurchaseRecord != nil {
		panic("duplicate hook PlayerTaoyuanPurchaseRecord")
	}
	g_Hooks.PlayerTaoyuanPurchaseRecord = hook
}
func (h hooksOP) PlayerTaoyuanQuest(hook PlayerTaoyuanQuestHook) {
	if g_Hooks.PlayerTaoyuanQuest != nil {
		panic("duplicate hook PlayerTaoyuanQuest")
	}
	g_Hooks.PlayerTaoyuanQuest = hook
}
func (h hooksOP) PlayerTbXxdRoleinfo(hook PlayerTbXxdRoleinfoHook) {
	if g_Hooks.PlayerTbXxdRoleinfo != nil {
		panic("duplicate hook PlayerTbXxdRoleinfo")
	}
	g_Hooks.PlayerTbXxdRoleinfo = hook
}
func (h hooksOP) PlayerTeamInfo(hook PlayerTeamInfoHook) {
	if g_Hooks.PlayerTeamInfo != nil {
		panic("duplicate hook PlayerTeamInfo")
	}
	g_Hooks.PlayerTeamInfo = hook
}
func (h hooksOP) PlayerTotem(hook PlayerTotemHook) {
	if g_Hooks.PlayerTotem != nil {
		panic("duplicate hook PlayerTotem")
	}
	g_Hooks.PlayerTotem = hook
}
func (h hooksOP) PlayerTotemInfo(hook PlayerTotemInfoHook) {
	if g_Hooks.PlayerTotemInfo != nil {
		panic("duplicate hook PlayerTotemInfo")
	}
	g_Hooks.PlayerTotemInfo = hook
}
func (h hooksOP) PlayerTown(hook PlayerTownHook) {
	if g_Hooks.PlayerTown != nil {
		panic("duplicate hook PlayerTown")
	}
	g_Hooks.PlayerTown = hook
}
func (h hooksOP) PlayerTraderRefreshState(hook PlayerTraderRefreshStateHook) {
	if g_Hooks.PlayerTraderRefreshState != nil {
		panic("duplicate hook PlayerTraderRefreshState")
	}
	g_Hooks.PlayerTraderRefreshState = hook
}
func (h hooksOP) PlayerTraderStoreState(hook PlayerTraderStoreStateHook) {
	if g_Hooks.PlayerTraderStoreState != nil {
		panic("duplicate hook PlayerTraderStoreState")
	}
	g_Hooks.PlayerTraderStoreState = hook
}
func (h hooksOP) PlayerUseSkill(hook PlayerUseSkillHook) {
	if g_Hooks.PlayerUseSkill != nil {
		panic("duplicate hook PlayerUseSkill")
	}
	g_Hooks.PlayerUseSkill = hook
}
func (h hooksOP) PlayerVip(hook PlayerVipHook) {
	if g_Hooks.PlayerVip != nil {
		panic("duplicate hook PlayerVip")
	}
	g_Hooks.PlayerVip = hook
}
/* ========== global_announcement ========== */

// 公告列表
type GlobalAnnouncementHook interface {
	Load(row *GlobalAnnouncementRow)
	Insert(row *GlobalAnnouncementRow)
	Delete(old *GlobalAnnouncementRow)
	Update(row, old *GlobalAnnouncementRow)
}

/* ========== global_arena_rank ========== */

// 全局比武场数据
type GlobalArenaRankHook interface {
	Load(row *GlobalArenaRankRow)
	Insert(row *GlobalArenaRankRow)
	Delete(old *GlobalArenaRankRow)
	Update(row, old *GlobalArenaRankRow)
}

/* ========== global_clique ========== */

// 帮派信息
type GlobalCliqueHook interface {
	Load(row *GlobalCliqueRow)
	Insert(row *GlobalCliqueRow)
	Delete(old *GlobalCliqueRow)
	Update(row, old *GlobalCliqueRow)
}

/* ========== global_clique_boat ========== */

// 镖船信息
type GlobalCliqueBoatHook interface {
	Load(row *GlobalCliqueBoatRow)
	Insert(row *GlobalCliqueBoatRow)
	Delete(old *GlobalCliqueBoatRow)
	Update(row, old *GlobalCliqueBoatRow)
}

/* ========== global_despair_boss ========== */

// 绝望关卡boss
type GlobalDespairBossHook interface {
	Load(row *GlobalDespairBossRow)
	Insert(row *GlobalDespairBossRow)
	Delete(old *GlobalDespairBossRow)
	Update(row, old *GlobalDespairBossRow)
}

/* ========== global_despair_boss_history ========== */

// 绝望关卡boss历史记录
type GlobalDespairBossHistoryHook interface {
	Load(row *GlobalDespairBossHistoryRow)
	Insert(row *GlobalDespairBossHistoryRow)
	Delete(old *GlobalDespairBossHistoryRow)
	Update(row, old *GlobalDespairBossHistoryRow)
}

/* ========== global_despair_land ========== */

// 绝望之地状态
type GlobalDespairLandHook interface {
	Load(row *GlobalDespairLandRow)
	Insert(row *GlobalDespairLandRow)
	Delete(old *GlobalDespairLandRow)
	Update(row, old *GlobalDespairLandRow)
}

/* ========== global_despair_land_battle_record ========== */

// 玩家战报记录
type GlobalDespairLandBattleRecordHook interface {
	Load(row *GlobalDespairLandBattleRecordRow)
	Insert(row *GlobalDespairLandBattleRecordRow)
	Delete(old *GlobalDespairLandBattleRecordRow)
	Update(row, old *GlobalDespairLandBattleRecordRow)
}

/* ========== global_despair_land_camp ========== */

// 绝望之地阵营状态
type GlobalDespairLandCampHook interface {
	Load(row *GlobalDespairLandCampRow)
	Insert(row *GlobalDespairLandCampRow)
	Delete(old *GlobalDespairLandCampRow)
	Update(row, old *GlobalDespairLandCampRow)
}

/* ========== global_gift_card_record ========== */

// 礼品卡兑换记录
type GlobalGiftCardRecordHook interface {
	Load(row *GlobalGiftCardRecordRow)
	Insert(row *GlobalGiftCardRecordRow)
	Delete(old *GlobalGiftCardRecordRow)
	Update(row, old *GlobalGiftCardRecordRow)
}

/* ========== global_group_buy_status ========== */

// hd服务器记录团购状态信息
type GlobalGroupBuyStatusHook interface {
	Load(row *GlobalGroupBuyStatusRow)
	Insert(row *GlobalGroupBuyStatusRow)
	Delete(old *GlobalGroupBuyStatusRow)
	Update(row, old *GlobalGroupBuyStatusRow)
}

/* ========== global_mail ========== */

// 全局邮件
type GlobalMailHook interface {
	Load(row *GlobalMailRow)
	Insert(row *GlobalMailRow)
	Delete(old *GlobalMailRow)
	Update(row, old *GlobalMailRow)
}

/* ========== global_mail_attachments ========== */

// 全局邮件附件
type GlobalMailAttachmentsHook interface {
	Load(row *GlobalMailAttachmentsRow)
	Insert(row *GlobalMailAttachmentsRow)
	Delete(old *GlobalMailAttachmentsRow)
	Update(row, old *GlobalMailAttachmentsRow)
}

/* ========== global_tb_xxd_onlinecnt ========== */

// 腾讯经分在线玩家统计日志
type GlobalTbXxdOnlinecntHook interface {
	Load(row *GlobalTbXxdOnlinecntRow)
	Insert(row *GlobalTbXxdOnlinecntRow)
	Delete(old *GlobalTbXxdOnlinecntRow)
	Update(row, old *GlobalTbXxdOnlinecntRow)
}

/* ========== player ========== */

// 玩家基础信息
type PlayerHook interface {
	Load(row *PlayerRow)
	Insert(row *PlayerRow)
	Delete(old *PlayerRow)
	Update(row, old *PlayerRow)
}

/* ========== player_activity ========== */

// 玩家活跃度
type PlayerActivityHook interface {
	Load(row *PlayerActivityRow)
	Insert(row *PlayerActivityRow)
	Delete(old *PlayerActivityRow)
	Update(row, old *PlayerActivityRow)
}

/* ========== player_addition_quest ========== */

// 玩家支线任务
type PlayerAdditionQuestHook interface {
	Load(row *PlayerAdditionQuestRow)
	Insert(row *PlayerAdditionQuestRow)
	Delete(old *PlayerAdditionQuestRow)
	Update(row, old *PlayerAdditionQuestRow)
}

/* ========== player_any_sdk_order ========== */

// anysdk 订单处理纪录
type PlayerAnySdkOrderHook interface {
	Load(row *PlayerAnySdkOrderRow)
	Insert(row *PlayerAnySdkOrderRow)
	Delete(old *PlayerAnySdkOrderRow)
	Update(row, old *PlayerAnySdkOrderRow)
}

/* ========== player_arena ========== */

// 玩家比武场数据
type PlayerArenaHook interface {
	Load(row *PlayerArenaRow)
	Insert(row *PlayerArenaRow)
	Delete(old *PlayerArenaRow)
	Update(row, old *PlayerArenaRow)
}

/* ========== player_arena_rank ========== */

// 玩家比武场最近排名记录
type PlayerArenaRankHook interface {
	Load(row *PlayerArenaRankRow)
	Insert(row *PlayerArenaRankRow)
	Delete(old *PlayerArenaRankRow)
	Update(row, old *PlayerArenaRankRow)
}

/* ========== player_arena_record ========== */

// 玩家比武场记录
type PlayerArenaRecordHook interface {
	Load(row *PlayerArenaRecordRow)
	Insert(row *PlayerArenaRecordRow)
	Delete(old *PlayerArenaRecordRow)
	Update(row, old *PlayerArenaRecordRow)
}

/* ========== player_awaken_graphic ========== */

// 玩家觉醒图谱
type PlayerAwakenGraphicHook interface {
	Load(row *PlayerAwakenGraphicRow)
	Insert(row *PlayerAwakenGraphicRow)
	Delete(old *PlayerAwakenGraphicRow)
	Update(row, old *PlayerAwakenGraphicRow)
}

/* ========== player_battle_pet ========== */

// 玩家灵宠数据
type PlayerBattlePetHook interface {
	Load(row *PlayerBattlePetRow)
	Insert(row *PlayerBattlePetRow)
	Delete(old *PlayerBattlePetRow)
	Update(row, old *PlayerBattlePetRow)
}

/* ========== player_battle_pet_grid ========== */

// 
type PlayerBattlePetGridHook interface {
	Load(row *PlayerBattlePetGridRow)
	Insert(row *PlayerBattlePetGridRow)
	Delete(old *PlayerBattlePetGridRow)
	Update(row, old *PlayerBattlePetGridRow)
}

/* ========== player_battle_pet_state ========== */

// 玩家灵宠状态
type PlayerBattlePetStateHook interface {
	Load(row *PlayerBattlePetStateRow)
	Insert(row *PlayerBattlePetStateRow)
	Delete(old *PlayerBattlePetStateRow)
	Update(row, old *PlayerBattlePetStateRow)
}

/* ========== player_chest_state ========== */

// 玩家宝箱状态
type PlayerChestStateHook interface {
	Load(row *PlayerChestStateRow)
	Insert(row *PlayerChestStateRow)
	Delete(old *PlayerChestStateRow)
	Update(row, old *PlayerChestStateRow)
}

/* ========== player_clique_kongfu_attrib ========== */

// 玩家帮派武功属性加成
type PlayerCliqueKongfuAttribHook interface {
	Load(row *PlayerCliqueKongfuAttribRow)
	Insert(row *PlayerCliqueKongfuAttribRow)
	Delete(old *PlayerCliqueKongfuAttribRow)
	Update(row, old *PlayerCliqueKongfuAttribRow)
}

/* ========== player_coins ========== */

// 玩家铜币兑换表
type PlayerCoinsHook interface {
	Load(row *PlayerCoinsRow)
	Insert(row *PlayerCoinsRow)
	Delete(old *PlayerCoinsRow)
	Update(row, old *PlayerCoinsRow)
}

/* ========== player_cornucopia ========== */

// 玩家铜币兑换表
type PlayerCornucopiaHook interface {
	Load(row *PlayerCornucopiaRow)
	Insert(row *PlayerCornucopiaRow)
	Delete(old *PlayerCornucopiaRow)
	Update(row, old *PlayerCornucopiaRow)
}

/* ========== player_daily_quest ========== */

// 玩家每日任务
type PlayerDailyQuestHook interface {
	Load(row *PlayerDailyQuestRow)
	Insert(row *PlayerDailyQuestRow)
	Delete(old *PlayerDailyQuestRow)
	Update(row, old *PlayerDailyQuestRow)
}

/* ========== player_daily_quest_star_award_info ========== */

// 玩家每日任务星数奖励状态
type PlayerDailyQuestStarAwardInfoHook interface {
	Load(row *PlayerDailyQuestStarAwardInfoRow)
	Insert(row *PlayerDailyQuestStarAwardInfoRow)
	Delete(old *PlayerDailyQuestStarAwardInfoRow)
	Update(row, old *PlayerDailyQuestStarAwardInfoRow)
}

/* ========== player_daily_sign_in_record ========== */

// 玩家每日签到记录
type PlayerDailySignInRecordHook interface {
	Load(row *PlayerDailySignInRecordRow)
	Insert(row *PlayerDailySignInRecordRow)
	Delete(old *PlayerDailySignInRecordRow)
	Update(row, old *PlayerDailySignInRecordRow)
}

/* ========== player_daily_sign_in_state ========== */

// 玩家最近七日每日签到状态
type PlayerDailySignInStateHook interface {
	Load(row *PlayerDailySignInStateRow)
	Insert(row *PlayerDailySignInStateRow)
	Delete(old *PlayerDailySignInStateRow)
	Update(row, old *PlayerDailySignInStateRow)
}

/* ========== player_despair_land_camp_state ========== */

// 玩家绝望之地每周进度
type PlayerDespairLandCampStateHook interface {
	Load(row *PlayerDespairLandCampStateRow)
	Insert(row *PlayerDespairLandCampStateRow)
	Delete(old *PlayerDespairLandCampStateRow)
	Update(row, old *PlayerDespairLandCampStateRow)
}

/* ========== player_despair_land_camp_state_history ========== */

// 玩家绝望之地每周进度
type PlayerDespairLandCampStateHistoryHook interface {
	Load(row *PlayerDespairLandCampStateHistoryRow)
	Insert(row *PlayerDespairLandCampStateHistoryRow)
	Delete(old *PlayerDespairLandCampStateHistoryRow)
	Update(row, old *PlayerDespairLandCampStateHistoryRow)
}

/* ========== player_despair_land_level_record ========== */

// 玩家绝望之地状态
type PlayerDespairLandLevelRecordHook interface {
	Load(row *PlayerDespairLandLevelRecordRow)
	Insert(row *PlayerDespairLandLevelRecordRow)
	Delete(old *PlayerDespairLandLevelRecordRow)
	Update(row, old *PlayerDespairLandLevelRecordRow)
}

/* ========== player_despair_land_state ========== */

// 玩家绝望之地状态
type PlayerDespairLandStateHook interface {
	Load(row *PlayerDespairLandStateRow)
	Insert(row *PlayerDespairLandStateRow)
	Delete(old *PlayerDespairLandStateRow)
	Update(row, old *PlayerDespairLandStateRow)
}

/* ========== player_driving_sword_event ========== */

// 玩家云海事件列表
type PlayerDrivingSwordEventHook interface {
	Load(row *PlayerDrivingSwordEventRow)
	Insert(row *PlayerDrivingSwordEventRow)
	Delete(old *PlayerDrivingSwordEventRow)
	Update(row, old *PlayerDrivingSwordEventRow)
}

/* ========== player_driving_sword_event_exploring ========== */

// 玩家云海探险事件信息
type PlayerDrivingSwordEventExploringHook interface {
	Load(row *PlayerDrivingSwordEventExploringRow)
	Insert(row *PlayerDrivingSwordEventExploringRow)
	Delete(old *PlayerDrivingSwordEventExploringRow)
	Update(row, old *PlayerDrivingSwordEventExploringRow)
}

/* ========== player_driving_sword_event_treasure ========== */

// 玩家云海宝藏事件信息
type PlayerDrivingSwordEventTreasureHook interface {
	Load(row *PlayerDrivingSwordEventTreasureRow)
	Insert(row *PlayerDrivingSwordEventTreasureRow)
	Delete(old *PlayerDrivingSwordEventTreasureRow)
	Update(row, old *PlayerDrivingSwordEventTreasureRow)
}

/* ========== player_driving_sword_event_visiting ========== */

// 玩家云海拜访事件信息
type PlayerDrivingSwordEventVisitingHook interface {
	Load(row *PlayerDrivingSwordEventVisitingRow)
	Insert(row *PlayerDrivingSwordEventVisitingRow)
	Delete(old *PlayerDrivingSwordEventVisitingRow)
	Update(row, old *PlayerDrivingSwordEventVisitingRow)
}

/* ========== player_driving_sword_eventmask ========== */

// 玩家云海事件地图
type PlayerDrivingSwordEventmaskHook interface {
	Load(row *PlayerDrivingSwordEventmaskRow)
	Insert(row *PlayerDrivingSwordEventmaskRow)
	Delete(old *PlayerDrivingSwordEventmaskRow)
	Update(row, old *PlayerDrivingSwordEventmaskRow)
}

/* ========== player_driving_sword_info ========== */

// 玩家御剑基本数据
type PlayerDrivingSwordInfoHook interface {
	Load(row *PlayerDrivingSwordInfoRow)
	Insert(row *PlayerDrivingSwordInfoRow)
	Delete(old *PlayerDrivingSwordInfoRow)
	Update(row, old *PlayerDrivingSwordInfoRow)
}

/* ========== player_driving_sword_map ========== */

// 玩家云海地图
type PlayerDrivingSwordMapHook interface {
	Load(row *PlayerDrivingSwordMapRow)
	Insert(row *PlayerDrivingSwordMapRow)
	Delete(old *PlayerDrivingSwordMapRow)
	Update(row, old *PlayerDrivingSwordMapRow)
}

/* ========== player_equipment ========== */

// 玩家装备表
type PlayerEquipmentHook interface {
	Load(row *PlayerEquipmentRow)
	Insert(row *PlayerEquipmentRow)
	Delete(old *PlayerEquipmentRow)
	Update(row, old *PlayerEquipmentRow)
}

/* ========== player_event_award_record ========== */

// 玩家活动奖励领取记录
type PlayerEventAwardRecordHook interface {
	Load(row *PlayerEventAwardRecordRow)
	Insert(row *PlayerEventAwardRecordRow)
	Delete(old *PlayerEventAwardRecordRow)
	Update(row, old *PlayerEventAwardRecordRow)
}

/* ========== player_event_daily_online ========== */

// 
type PlayerEventDailyOnlineHook interface {
	Load(row *PlayerEventDailyOnlineRow)
	Insert(row *PlayerEventDailyOnlineRow)
	Delete(old *PlayerEventDailyOnlineRow)
	Update(row, old *PlayerEventDailyOnlineRow)
}

/* ========== player_event_vn_daily_recharge ========== */

// 
type PlayerEventVnDailyRechargeHook interface {
	Load(row *PlayerEventVnDailyRechargeRow)
	Insert(row *PlayerEventVnDailyRechargeRow)
	Delete(old *PlayerEventVnDailyRechargeRow)
	Update(row, old *PlayerEventVnDailyRechargeRow)
}

/* ========== player_event_vn_dragon_ball_record ========== */

// 越南龙珠活动玩家获取记录
type PlayerEventVnDragonBallRecordHook interface {
	Load(row *PlayerEventVnDragonBallRecordRow)
	Insert(row *PlayerEventVnDragonBallRecordRow)
	Delete(old *PlayerEventVnDragonBallRecordRow)
	Update(row, old *PlayerEventVnDragonBallRecordRow)
}

/* ========== player_events_fight_power ========== */

// 玩家战力运营活动记录
type PlayerEventsFightPowerHook interface {
	Load(row *PlayerEventsFightPowerRow)
	Insert(row *PlayerEventsFightPowerRow)
	Delete(old *PlayerEventsFightPowerRow)
	Update(row, old *PlayerEventsFightPowerRow)
}

/* ========== player_events_ingot_record ========== */

// 玩家元宝充值和消耗活动记录
type PlayerEventsIngotRecordHook interface {
	Load(row *PlayerEventsIngotRecordRow)
	Insert(row *PlayerEventsIngotRecordRow)
	Delete(old *PlayerEventsIngotRecordRow)
	Update(row, old *PlayerEventsIngotRecordRow)
}

/* ========== player_extend_level ========== */

// 玩家活动关卡状态
type PlayerExtendLevelHook interface {
	Load(row *PlayerExtendLevelRow)
	Insert(row *PlayerExtendLevelRow)
	Delete(old *PlayerExtendLevelRow)
	Update(row, old *PlayerExtendLevelRow)
}

/* ========== player_extend_quest ========== */

// 伙伴任务
type PlayerExtendQuestHook interface {
	Load(row *PlayerExtendQuestRow)
	Insert(row *PlayerExtendQuestRow)
	Delete(old *PlayerExtendQuestRow)
	Update(row, old *PlayerExtendQuestRow)
}

/* ========== player_fame ========== */

// 
type PlayerFameHook interface {
	Load(row *PlayerFameRow)
	Insert(row *PlayerFameRow)
	Delete(old *PlayerFameRow)
	Update(row, old *PlayerFameRow)
}

/* ========== player_fashion ========== */

// 玩家时装
type PlayerFashionHook interface {
	Load(row *PlayerFashionRow)
	Insert(row *PlayerFashionRow)
	Delete(old *PlayerFashionRow)
	Update(row, old *PlayerFashionRow)
}

/* ========== player_fashion_state ========== */

// 玩家时装状态
type PlayerFashionStateHook interface {
	Load(row *PlayerFashionStateRow)
	Insert(row *PlayerFashionStateRow)
	Delete(old *PlayerFashionStateRow)
	Update(row, old *PlayerFashionStateRow)
}

/* ========== player_fate_box_state ========== */

// 玩家命锁宝箱状态
type PlayerFateBoxStateHook interface {
	Load(row *PlayerFateBoxStateRow)
	Insert(row *PlayerFateBoxStateRow)
	Delete(old *PlayerFateBoxStateRow)
	Update(row, old *PlayerFateBoxStateRow)
}

/* ========== player_fight_num ========== */

// 玩家战斗力
type PlayerFightNumHook interface {
	Load(row *PlayerFightNumRow)
	Insert(row *PlayerFightNumRow)
	Delete(old *PlayerFightNumRow)
	Update(row, old *PlayerFightNumRow)
}

/* ========== player_first_charge ========== */

// 玩家首冲表
type PlayerFirstChargeHook interface {
	Load(row *PlayerFirstChargeRow)
	Insert(row *PlayerFirstChargeRow)
	Delete(old *PlayerFirstChargeRow)
	Update(row, old *PlayerFirstChargeRow)
}

/* ========== player_formation ========== */

// 玩家阵型站位
type PlayerFormationHook interface {
	Load(row *PlayerFormationRow)
	Insert(row *PlayerFormationRow)
	Delete(old *PlayerFormationRow)
	Update(row, old *PlayerFormationRow)
}

/* ========== player_func_key ========== */

// 玩家功能开放表
type PlayerFuncKeyHook interface {
	Load(row *PlayerFuncKeyRow)
	Insert(row *PlayerFuncKeyRow)
	Delete(old *PlayerFuncKeyRow)
	Update(row, old *PlayerFuncKeyRow)
}

/* ========== player_ghost ========== */

// 玩家魂侍表
type PlayerGhostHook interface {
	Load(row *PlayerGhostRow)
	Insert(row *PlayerGhostRow)
	Delete(old *PlayerGhostRow)
	Update(row, old *PlayerGhostRow)
}

/* ========== player_ghost_equipment ========== */

// 玩家魂侍装备表
type PlayerGhostEquipmentHook interface {
	Load(row *PlayerGhostEquipmentRow)
	Insert(row *PlayerGhostEquipmentRow)
	Delete(old *PlayerGhostEquipmentRow)
	Update(row, old *PlayerGhostEquipmentRow)
}

/* ========== player_ghost_state ========== */

// 玩家魂侍相关状态
type PlayerGhostStateHook interface {
	Load(row *PlayerGhostStateRow)
	Insert(row *PlayerGhostStateRow)
	Delete(old *PlayerGhostStateRow)
	Update(row, old *PlayerGhostStateRow)
}

/* ========== player_global_chat_state ========== */

// 禁言状态
type PlayerGlobalChatStateHook interface {
	Load(row *PlayerGlobalChatStateRow)
	Insert(row *PlayerGlobalChatStateRow)
	Delete(old *PlayerGlobalChatStateRow)
	Update(row, old *PlayerGlobalChatStateRow)
}

/* ========== player_global_clique_building ========== */

// 玩家帮派建筑信息
type PlayerGlobalCliqueBuildingHook interface {
	Load(row *PlayerGlobalCliqueBuildingRow)
	Insert(row *PlayerGlobalCliqueBuildingRow)
	Delete(old *PlayerGlobalCliqueBuildingRow)
	Update(row, old *PlayerGlobalCliqueBuildingRow)
}

/* ========== player_global_clique_building_quest ========== */

// 玩家帮派建筑任务
type PlayerGlobalCliqueBuildingQuestHook interface {
	Load(row *PlayerGlobalCliqueBuildingQuestRow)
	Insert(row *PlayerGlobalCliqueBuildingQuestRow)
	Delete(old *PlayerGlobalCliqueBuildingQuestRow)
	Update(row, old *PlayerGlobalCliqueBuildingQuestRow)
}

/* ========== player_global_clique_daily_quest ========== */

// 玩家每日帮派任务
type PlayerGlobalCliqueDailyQuestHook interface {
	Load(row *PlayerGlobalCliqueDailyQuestRow)
	Insert(row *PlayerGlobalCliqueDailyQuestRow)
	Delete(old *PlayerGlobalCliqueDailyQuestRow)
	Update(row, old *PlayerGlobalCliqueDailyQuestRow)
}

/* ========== player_global_clique_escort ========== */

// 玩家帮派运镖信息
type PlayerGlobalCliqueEscortHook interface {
	Load(row *PlayerGlobalCliqueEscortRow)
	Insert(row *PlayerGlobalCliqueEscortRow)
	Delete(old *PlayerGlobalCliqueEscortRow)
	Update(row, old *PlayerGlobalCliqueEscortRow)
}

/* ========== player_global_clique_escort_message ========== */

// 玩家帮派运镖消息
type PlayerGlobalCliqueEscortMessageHook interface {
	Load(row *PlayerGlobalCliqueEscortMessageRow)
	Insert(row *PlayerGlobalCliqueEscortMessageRow)
	Delete(old *PlayerGlobalCliqueEscortMessageRow)
	Update(row, old *PlayerGlobalCliqueEscortMessageRow)
}

/* ========== player_global_clique_hijack ========== */

// 玩家劫镖信息
type PlayerGlobalCliqueHijackHook interface {
	Load(row *PlayerGlobalCliqueHijackRow)
	Insert(row *PlayerGlobalCliqueHijackRow)
	Delete(old *PlayerGlobalCliqueHijackRow)
	Update(row, old *PlayerGlobalCliqueHijackRow)
}

/* ========== player_global_clique_info ========== */

// 玩家帮派信息
type PlayerGlobalCliqueInfoHook interface {
	Load(row *PlayerGlobalCliqueInfoRow)
	Insert(row *PlayerGlobalCliqueInfoRow)
	Delete(old *PlayerGlobalCliqueInfoRow)
	Update(row, old *PlayerGlobalCliqueInfoRow)
}

/* ========== player_global_clique_kongfu ========== */

// 玩家帮派武学
type PlayerGlobalCliqueKongfuHook interface {
	Load(row *PlayerGlobalCliqueKongfuRow)
	Insert(row *PlayerGlobalCliqueKongfuRow)
	Delete(old *PlayerGlobalCliqueKongfuRow)
	Update(row, old *PlayerGlobalCliqueKongfuRow)
}

/* ========== player_global_friend ========== */

// 玩家好友列表
type PlayerGlobalFriendHook interface {
	Load(row *PlayerGlobalFriendRow)
	Insert(row *PlayerGlobalFriendRow)
	Delete(old *PlayerGlobalFriendRow)
	Update(row, old *PlayerGlobalFriendRow)
}

/* ========== player_global_friend_chat ========== */

// 玩家聊天记录
type PlayerGlobalFriendChatHook interface {
	Load(row *PlayerGlobalFriendChatRow)
	Insert(row *PlayerGlobalFriendChatRow)
	Delete(old *PlayerGlobalFriendChatRow)
	Update(row, old *PlayerGlobalFriendChatRow)
}

/* ========== player_global_friend_state ========== */

// 玩家好友功能状态数据
type PlayerGlobalFriendStateHook interface {
	Load(row *PlayerGlobalFriendStateRow)
	Insert(row *PlayerGlobalFriendStateRow)
	Delete(old *PlayerGlobalFriendStateRow)
	Update(row, old *PlayerGlobalFriendStateRow)
}

/* ========== player_global_gift_card_record ========== */

// 玩家兑换记录
type PlayerGlobalGiftCardRecordHook interface {
	Load(row *PlayerGlobalGiftCardRecordRow)
	Insert(row *PlayerGlobalGiftCardRecordRow)
	Delete(old *PlayerGlobalGiftCardRecordRow)
	Update(row, old *PlayerGlobalGiftCardRecordRow)
}

/* ========== player_global_mail_state ========== */

// 玩家全局邮件记录
type PlayerGlobalMailStateHook interface {
	Load(row *PlayerGlobalMailStateRow)
	Insert(row *PlayerGlobalMailStateRow)
	Delete(old *PlayerGlobalMailStateRow)
	Update(row, old *PlayerGlobalMailStateRow)
}

/* ========== player_global_world_chat_state ========== */

// 玩家世界聊天状态
type PlayerGlobalWorldChatStateHook interface {
	Load(row *PlayerGlobalWorldChatStateRow)
	Insert(row *PlayerGlobalWorldChatStateRow)
	Delete(old *PlayerGlobalWorldChatStateRow)
	Update(row, old *PlayerGlobalWorldChatStateRow)
}

/* ========== player_hard_level ========== */

// 难度关卡功能权值
type PlayerHardLevelHook interface {
	Load(row *PlayerHardLevelRow)
	Insert(row *PlayerHardLevelRow)
	Delete(old *PlayerHardLevelRow)
	Update(row, old *PlayerHardLevelRow)
}

/* ========== player_hard_level_record ========== */

// 难度关卡记录
type PlayerHardLevelRecordHook interface {
	Load(row *PlayerHardLevelRecordRow)
	Insert(row *PlayerHardLevelRecordRow)
	Delete(old *PlayerHardLevelRecordRow)
	Update(row, old *PlayerHardLevelRecordRow)
}

/* ========== player_heart ========== */

// 玩家爱心表
type PlayerHeartHook interface {
	Load(row *PlayerHeartRow)
	Insert(row *PlayerHeartRow)
	Delete(old *PlayerHeartRow)
	Update(row, old *PlayerHeartRow)
}

/* ========== player_heart_draw ========== */

// 玩家爱心抽奖
type PlayerHeartDrawHook interface {
	Load(row *PlayerHeartDrawRow)
	Insert(row *PlayerHeartDrawRow)
	Delete(old *PlayerHeartDrawRow)
	Update(row, old *PlayerHeartDrawRow)
}

/* ========== player_heart_draw_card_record ========== */

// 玩家爱心刮刮卡抽奖记录
type PlayerHeartDrawCardRecordHook interface {
	Load(row *PlayerHeartDrawCardRecordRow)
	Insert(row *PlayerHeartDrawCardRecordRow)
	Delete(old *PlayerHeartDrawCardRecordRow)
	Update(row, old *PlayerHeartDrawCardRecordRow)
}

/* ========== player_heart_draw_wheel_record ========== */

// 玩家爱心大转盘抽奖记录
type PlayerHeartDrawWheelRecordHook interface {
	Load(row *PlayerHeartDrawWheelRecordRow)
	Insert(row *PlayerHeartDrawWheelRecordRow)
	Delete(old *PlayerHeartDrawWheelRecordRow)
	Update(row, old *PlayerHeartDrawWheelRecordRow)
}

/* ========== player_info ========== */

// 玩家信息表
type PlayerInfoHook interface {
	Load(row *PlayerInfoRow)
	Insert(row *PlayerInfoRow)
	Delete(old *PlayerInfoRow)
	Update(row, old *PlayerInfoRow)
}

/* ========== player_is_operated ========== */

// 记录玩家是否第一次操作
type PlayerIsOperatedHook interface {
	Load(row *PlayerIsOperatedRow)
	Insert(row *PlayerIsOperatedRow)
	Delete(old *PlayerIsOperatedRow)
	Update(row, old *PlayerIsOperatedRow)
}

/* ========== player_item ========== */

// 玩家物品
type PlayerItemHook interface {
	Load(row *PlayerItemRow)
	Insert(row *PlayerItemRow)
	Delete(old *PlayerItemRow)
	Update(row, old *PlayerItemRow)
}

/* ========== player_item_appendix ========== */

// 玩家装备追加属性表
type PlayerItemAppendixHook interface {
	Load(row *PlayerItemAppendixRow)
	Insert(row *PlayerItemAppendixRow)
	Delete(old *PlayerItemAppendixRow)
	Update(row, old *PlayerItemAppendixRow)
}

/* ========== player_item_buyback ========== */

// 玩家物品回购表
type PlayerItemBuybackHook interface {
	Load(row *PlayerItemBuybackRow)
	Insert(row *PlayerItemBuybackRow)
	Delete(old *PlayerItemBuybackRow)
	Update(row, old *PlayerItemBuybackRow)
}

/* ========== player_item_recast_state ========== */

// 玩家装备重铸状态
type PlayerItemRecastStateHook interface {
	Load(row *PlayerItemRecastStateRow)
	Insert(row *PlayerItemRecastStateRow)
	Delete(old *PlayerItemRecastStateRow)
	Update(row, old *PlayerItemRecastStateRow)
}

/* ========== player_level_award_info ========== */

// 玩家等级奖励领取状态表
type PlayerLevelAwardInfoHook interface {
	Load(row *PlayerLevelAwardInfoRow)
	Insert(row *PlayerLevelAwardInfoRow)
	Delete(old *PlayerLevelAwardInfoRow)
	Update(row, old *PlayerLevelAwardInfoRow)
}

/* ========== player_login_award_record ========== */

// 玩家奖励领取情况
type PlayerLoginAwardRecordHook interface {
	Load(row *PlayerLoginAwardRecordRow)
	Insert(row *PlayerLoginAwardRecordRow)
	Delete(old *PlayerLoginAwardRecordRow)
	Update(row, old *PlayerLoginAwardRecordRow)
}

/* ========== player_mail ========== */

// 玩家邮件表
type PlayerMailHook interface {
	Load(row *PlayerMailRow)
	Insert(row *PlayerMailRow)
	Delete(old *PlayerMailRow)
	Update(row, old *PlayerMailRow)
}

/* ========== player_mail_attachment ========== */

// 玩家邮件附件表
type PlayerMailAttachmentHook interface {
	Load(row *PlayerMailAttachmentRow)
	Insert(row *PlayerMailAttachmentRow)
	Delete(old *PlayerMailAttachmentRow)
	Update(row, old *PlayerMailAttachmentRow)
}

/* ========== player_mail_attachment_lg ========== */

// 玩家领取邮件附件表
type PlayerMailAttachmentLgHook interface {
	Load(row *PlayerMailAttachmentLgRow)
	Insert(row *PlayerMailAttachmentLgRow)
	Delete(old *PlayerMailAttachmentLgRow)
	Update(row, old *PlayerMailAttachmentLgRow)
}

/* ========== player_mail_lg ========== */

// 玩家已删除邮件表
type PlayerMailLgHook interface {
	Load(row *PlayerMailLgRow)
	Insert(row *PlayerMailLgRow)
	Delete(old *PlayerMailLgRow)
	Update(row, old *PlayerMailLgRow)
}

/* ========== player_meditation_state ========== */

// 玩家打坐状态
type PlayerMeditationStateHook interface {
	Load(row *PlayerMeditationStateRow)
	Insert(row *PlayerMeditationStateRow)
	Delete(old *PlayerMeditationStateRow)
	Update(row, old *PlayerMeditationStateRow)
}

/* ========== player_mission ========== */

// 玩家区域数据
type PlayerMissionHook interface {
	Load(row *PlayerMissionRow)
	Insert(row *PlayerMissionRow)
	Delete(old *PlayerMissionRow)
	Update(row, old *PlayerMissionRow)
}

/* ========== player_mission_level ========== */

// 玩家区域关卡数据
type PlayerMissionLevelHook interface {
	Load(row *PlayerMissionLevelRow)
	Insert(row *PlayerMissionLevelRow)
	Delete(old *PlayerMissionLevelRow)
	Update(row, old *PlayerMissionLevelRow)
}

/* ========== player_mission_level_record ========== */

// 关卡记录
type PlayerMissionLevelRecordHook interface {
	Load(row *PlayerMissionLevelRecordRow)
	Insert(row *PlayerMissionLevelRecordRow)
	Delete(old *PlayerMissionLevelRecordRow)
	Update(row, old *PlayerMissionLevelRecordRow)
}

/* ========== player_mission_level_state_bin ========== */

// 玩家区域关卡状态保存
type PlayerMissionLevelStateBinHook interface {
	Load(row *PlayerMissionLevelStateBinRow)
	Insert(row *PlayerMissionLevelStateBinRow)
	Delete(old *PlayerMissionLevelStateBinRow)
	Update(row, old *PlayerMissionLevelStateBinRow)
}

/* ========== player_mission_record ========== */

// 区域记录
type PlayerMissionRecordHook interface {
	Load(row *PlayerMissionRecordRow)
	Insert(row *PlayerMissionRecordRow)
	Delete(old *PlayerMissionRecordRow)
	Update(row, old *PlayerMissionRecordRow)
}

/* ========== player_mission_star_award ========== */

// 玩家区域评星领奖记录
type PlayerMissionStarAwardHook interface {
	Load(row *PlayerMissionStarAwardRow)
	Insert(row *PlayerMissionStarAwardRow)
	Delete(old *PlayerMissionStarAwardRow)
	Update(row, old *PlayerMissionStarAwardRow)
}

/* ========== player_money_tree ========== */

// 玩家摇钱树记录
type PlayerMoneyTreeHook interface {
	Load(row *PlayerMoneyTreeRow)
	Insert(row *PlayerMoneyTreeRow)
	Delete(old *PlayerMoneyTreeRow)
	Update(row, old *PlayerMoneyTreeRow)
}

/* ========== player_month_card_award_record ========== */

// 玩家最后领取月卡时间
type PlayerMonthCardAwardRecordHook interface {
	Load(row *PlayerMonthCardAwardRecordRow)
	Insert(row *PlayerMonthCardAwardRecordRow)
	Delete(old *PlayerMonthCardAwardRecordRow)
	Update(row, old *PlayerMonthCardAwardRecordRow)
}

/* ========== player_month_card_info ========== */

// 玩家月卡信息
type PlayerMonthCardInfoHook interface {
	Load(row *PlayerMonthCardInfoRow)
	Insert(row *PlayerMonthCardInfoRow)
	Delete(old *PlayerMonthCardInfoRow)
	Update(row, old *PlayerMonthCardInfoRow)
}

/* ========== player_multi_level_info ========== */

// 玩家多人关卡信息
type PlayerMultiLevelInfoHook interface {
	Load(row *PlayerMultiLevelInfoRow)
	Insert(row *PlayerMultiLevelInfoRow)
	Delete(old *PlayerMultiLevelInfoRow)
	Update(row, old *PlayerMultiLevelInfoRow)
}

/* ========== player_new_year_consume_record ========== */

// 春节玩家消费记录
type PlayerNewYearConsumeRecordHook interface {
	Load(row *PlayerNewYearConsumeRecordRow)
	Insert(row *PlayerNewYearConsumeRecordRow)
	Delete(old *PlayerNewYearConsumeRecordRow)
	Update(row, old *PlayerNewYearConsumeRecordRow)
}

/* ========== player_npc_talk_record ========== */

// 玩家与NPC对话奖励记录
type PlayerNpcTalkRecordHook interface {
	Load(row *PlayerNpcTalkRecordRow)
	Insert(row *PlayerNpcTalkRecordRow)
	Delete(old *PlayerNpcTalkRecordRow)
	Update(row, old *PlayerNpcTalkRecordRow)
}

/* ========== player_opened_town_treasure ========== */

// 玩家已开启的城镇宝箱
type PlayerOpenedTownTreasureHook interface {
	Load(row *PlayerOpenedTownTreasureRow)
	Insert(row *PlayerOpenedTownTreasureRow)
	Delete(old *PlayerOpenedTownTreasureRow)
	Update(row, old *PlayerOpenedTownTreasureRow)
}

/* ========== player_physical ========== */

// 玩家体力表
type PlayerPhysicalHook interface {
	Load(row *PlayerPhysicalRow)
	Insert(row *PlayerPhysicalRow)
	Delete(old *PlayerPhysicalRow)
	Update(row, old *PlayerPhysicalRow)
}

/* ========== player_purchase_record ========== */

// 玩家物品购买记录 仅记录限购商品
type PlayerPurchaseRecordHook interface {
	Load(row *PlayerPurchaseRecordRow)
	Insert(row *PlayerPurchaseRecordRow)
	Delete(old *PlayerPurchaseRecordRow)
	Update(row, old *PlayerPurchaseRecordRow)
}

/* ========== player_push_notify_switch ========== */

// 玩家推送通知开关列表
type PlayerPushNotifySwitchHook interface {
	Load(row *PlayerPushNotifySwitchRow)
	Insert(row *PlayerPushNotifySwitchRow)
	Delete(old *PlayerPushNotifySwitchRow)
	Update(row, old *PlayerPushNotifySwitchRow)
}

/* ========== player_pve_state ========== */

// 玩家灵宠状态
type PlayerPveStateHook interface {
	Load(row *PlayerPveStateRow)
	Insert(row *PlayerPveStateRow)
	Delete(old *PlayerPveStateRow)
	Update(row, old *PlayerPveStateRow)
}

/* ========== player_qq_vip_gift ========== */

// 玩家qq服务礼包领取记录
type PlayerQqVipGiftHook interface {
	Load(row *PlayerQqVipGiftRow)
	Insert(row *PlayerQqVipGiftRow)
	Delete(old *PlayerQqVipGiftRow)
	Update(row, old *PlayerQqVipGiftRow)
}

/* ========== player_quest ========== */

// 玩家任务
type PlayerQuestHook interface {
	Load(row *PlayerQuestRow)
	Insert(row *PlayerQuestRow)
	Delete(old *PlayerQuestRow)
	Update(row, old *PlayerQuestRow)
}

/* ========== player_rainbow_level ========== */

// 玩家彩虹关卡状态
type PlayerRainbowLevelHook interface {
	Load(row *PlayerRainbowLevelRow)
	Insert(row *PlayerRainbowLevelRow)
	Delete(old *PlayerRainbowLevelRow)
	Update(row, old *PlayerRainbowLevelRow)
}

/* ========== player_rainbow_level_state_bin ========== */

// 玩家彩虹关卡状态
type PlayerRainbowLevelStateBinHook interface {
	Load(row *PlayerRainbowLevelStateBinRow)
	Insert(row *PlayerRainbowLevelStateBinRow)
	Delete(old *PlayerRainbowLevelStateBinRow)
	Update(row, old *PlayerRainbowLevelStateBinRow)
}

/* ========== player_role ========== */

// 玩家角色数据
type PlayerRoleHook interface {
	Load(row *PlayerRoleRow)
	Insert(row *PlayerRoleRow)
	Delete(old *PlayerRoleRow)
	Update(row, old *PlayerRoleRow)
}

/* ========== player_sealedbook ========== */

// 玩家天书记录
type PlayerSealedbookHook interface {
	Load(row *PlayerSealedbookRow)
	Insert(row *PlayerSealedbookRow)
	Delete(old *PlayerSealedbookRow)
	Update(row, old *PlayerSealedbookRow)
}

/* ========== player_send_heart_record ========== */

// 玩家好友列表
type PlayerSendHeartRecordHook interface {
	Load(row *PlayerSendHeartRecordRow)
	Insert(row *PlayerSendHeartRecordRow)
	Delete(old *PlayerSendHeartRecordRow)
	Update(row, old *PlayerSendHeartRecordRow)
}

/* ========== player_skill ========== */

// 玩家角色绝招表
type PlayerSkillHook interface {
	Load(row *PlayerSkillRow)
	Insert(row *PlayerSkillRow)
	Delete(old *PlayerSkillRow)
	Update(row, old *PlayerSkillRow)
}

/* ========== player_state ========== */

// 玩家状态
type PlayerStateHook interface {
	Load(row *PlayerStateRow)
	Insert(row *PlayerStateRow)
	Delete(old *PlayerStateRow)
	Update(row, old *PlayerStateRow)
}

/* ========== player_sword_soul ========== */

// 玩家剑心数据
type PlayerSwordSoulHook interface {
	Load(row *PlayerSwordSoulRow)
	Insert(row *PlayerSwordSoulRow)
	Delete(old *PlayerSwordSoulRow)
	Update(row, old *PlayerSwordSoulRow)
}

/* ========== player_sword_soul_equipment ========== */

// 玩家剑心装备表
type PlayerSwordSoulEquipmentHook interface {
	Load(row *PlayerSwordSoulEquipmentRow)
	Insert(row *PlayerSwordSoulEquipmentRow)
	Delete(old *PlayerSwordSoulEquipmentRow)
	Update(row, old *PlayerSwordSoulEquipmentRow)
}

/* ========== player_sword_soul_state ========== */

// 玩家拔剑状态
type PlayerSwordSoulStateHook interface {
	Load(row *PlayerSwordSoulStateRow)
	Insert(row *PlayerSwordSoulStateRow)
	Delete(old *PlayerSwordSoulStateRow)
	Update(row, old *PlayerSwordSoulStateRow)
}

/* ========== player_taoyuan_bless ========== */

// 玩家桃源祈福
type PlayerTaoyuanBlessHook interface {
	Load(row *PlayerTaoyuanBlessRow)
	Insert(row *PlayerTaoyuanBlessRow)
	Delete(old *PlayerTaoyuanBlessRow)
	Update(row, old *PlayerTaoyuanBlessRow)
}

/* ========== player_taoyuan_fileds ========== */

// 玩家桃源田地作物
type PlayerTaoyuanFiledsHook interface {
	Load(row *PlayerTaoyuanFiledsRow)
	Insert(row *PlayerTaoyuanFiledsRow)
	Delete(old *PlayerTaoyuanFiledsRow)
	Update(row, old *PlayerTaoyuanFiledsRow)
}

/* ========== player_taoyuan_heart ========== */

// 玩家桃源之心
type PlayerTaoyuanHeartHook interface {
	Load(row *PlayerTaoyuanHeartRow)
	Insert(row *PlayerTaoyuanHeartRow)
	Delete(old *PlayerTaoyuanHeartRow)
	Update(row, old *PlayerTaoyuanHeartRow)
}

/* ========== player_taoyuan_item ========== */

// 玩家桃源物品
type PlayerTaoyuanItemHook interface {
	Load(row *PlayerTaoyuanItemRow)
	Insert(row *PlayerTaoyuanItemRow)
	Delete(old *PlayerTaoyuanItemRow)
	Update(row, old *PlayerTaoyuanItemRow)
}

/* ========== player_taoyuan_message ========== */

// 祈福消息列表
type PlayerTaoyuanMessageHook interface {
	Load(row *PlayerTaoyuanMessageRow)
	Insert(row *PlayerTaoyuanMessageRow)
	Delete(old *PlayerTaoyuanMessageRow)
	Update(row, old *PlayerTaoyuanMessageRow)
}

/* ========== player_taoyuan_product ========== */

// 玩家桃源食坊、酒坊
type PlayerTaoyuanProductHook interface {
	Load(row *PlayerTaoyuanProductRow)
	Insert(row *PlayerTaoyuanProductRow)
	Delete(old *PlayerTaoyuanProductRow)
	Update(row, old *PlayerTaoyuanProductRow)
}

/* ========== player_taoyuan_purchase_record ========== */

// 玩家桃源物品购买记录 仅记录限购商品
type PlayerTaoyuanPurchaseRecordHook interface {
	Load(row *PlayerTaoyuanPurchaseRecordRow)
	Insert(row *PlayerTaoyuanPurchaseRecordRow)
	Delete(old *PlayerTaoyuanPurchaseRecordRow)
	Update(row, old *PlayerTaoyuanPurchaseRecordRow)
}

/* ========== player_taoyuan_quest ========== */

// 玩家桃源愿望
type PlayerTaoyuanQuestHook interface {
	Load(row *PlayerTaoyuanQuestRow)
	Insert(row *PlayerTaoyuanQuestRow)
	Delete(old *PlayerTaoyuanQuestRow)
	Update(row, old *PlayerTaoyuanQuestRow)
}

/* ========== player_tb_xxd_roleinfo ========== */

// 腾讯经分用户信息表
type PlayerTbXxdRoleinfoHook interface {
	Load(row *PlayerTbXxdRoleinfoRow)
	Insert(row *PlayerTbXxdRoleinfoRow)
	Delete(old *PlayerTbXxdRoleinfoRow)
	Update(row, old *PlayerTbXxdRoleinfoRow)
}

/* ========== player_team_info ========== */

// 玩家队伍相关信息
type PlayerTeamInfoHook interface {
	Load(row *PlayerTeamInfoRow)
	Insert(row *PlayerTeamInfoRow)
	Delete(old *PlayerTeamInfoRow)
	Update(row, old *PlayerTeamInfoRow)
}

/* ========== player_totem ========== */

// 玩家阵印表
type PlayerTotemHook interface {
	Load(row *PlayerTotemRow)
	Insert(row *PlayerTotemRow)
	Delete(old *PlayerTotemRow)
	Update(row, old *PlayerTotemRow)
}

/* ========== player_totem_info ========== */

// 玩家阵印装备表
type PlayerTotemInfoHook interface {
	Load(row *PlayerTotemInfoRow)
	Insert(row *PlayerTotemInfoRow)
	Delete(old *PlayerTotemInfoRow)
	Update(row, old *PlayerTotemInfoRow)
}

/* ========== player_town ========== */

// 玩家城镇数据
type PlayerTownHook interface {
	Load(row *PlayerTownRow)
	Insert(row *PlayerTownRow)
	Delete(old *PlayerTownRow)
	Update(row, old *PlayerTownRow)
}

/* ========== player_trader_refresh_state ========== */

// 玩家随机商店手动刷新次数状态
type PlayerTraderRefreshStateHook interface {
	Load(row *PlayerTraderRefreshStateRow)
	Insert(row *PlayerTraderRefreshStateRow)
	Delete(old *PlayerTraderRefreshStateRow)
	Update(row, old *PlayerTraderRefreshStateRow)
}

/* ========== player_trader_store_state ========== */

// 玩家随机商店状态
type PlayerTraderStoreStateHook interface {
	Load(row *PlayerTraderStoreStateRow)
	Insert(row *PlayerTraderStoreStateRow)
	Delete(old *PlayerTraderStoreStateRow)
	Update(row, old *PlayerTraderStoreStateRow)
}

/* ========== player_use_skill ========== */

// 玩家角色当前使用的绝招表
type PlayerUseSkillHook interface {
	Load(row *PlayerUseSkillRow)
	Insert(row *PlayerUseSkillRow)
	Delete(old *PlayerUseSkillRow)
	Update(row, old *PlayerUseSkillRow)
}

/* ========== player_vip ========== */

// 玩家VIP卡信息
type PlayerVipHook interface {
	Load(row *PlayerVipRow)
	Insert(row *PlayerVipRow)
	Delete(old *PlayerVipRow)
	Update(row, old *PlayerVipRow)
}

