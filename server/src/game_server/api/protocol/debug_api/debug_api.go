package debug_api

import "core/net"

type Request interface {
	Process(*net.Session)
	TypeName() string
	GetModuleIdAndActionId() (int8, int8)
}

var (
	g_InHandler  InHandler
	g_OutHandler OutHandler
)

func SetInHandler(handler InHandler) {
	g_InHandler = handler
}

func SetOutHandler(handler OutHandler) {
	g_OutHandler = handler
}

type InHandler interface {
	AddBuddy(*net.Session, *AddBuddy_In)
	AddItem(*net.Session, *AddItem_In)
	SetRoleLevel(*net.Session, *SetRoleLevel_In)
	SetCoins(*net.Session, *SetCoins_In)
	SetIngot(*net.Session, *SetIngot_In)
	AddGhost(*net.Session, *AddGhost_In)
	SetPlayerPhysical(*net.Session, *SetPlayerPhysical_In)
	ResetLevelEnterCount(*net.Session, *ResetLevelEnterCount_In)
	AddExp(*net.Session, *AddExp_In)
	OpenGhostMission(*net.Session, *OpenGhostMission_In)
	SendMail(*net.Session, *SendMail_In)
	ClearMail(*net.Session, *ClearMail_In)
	OpenMissionLevel(*net.Session, *OpenMissionLevel_In)
	StartBattle(*net.Session, *StartBattle_In)
	ListenByName(*net.Session, *ListenByName_In)
	OpenQuest(*net.Session, *OpenQuest_In)
	OpenFunc(*net.Session, *OpenFunc_In)
	AddSwordSoul(*net.Session, *AddSwordSoul_In)
	AddBattlePet(*net.Session, *AddBattlePet_In)
	ResetMultiLevelEnterCount(*net.Session, *ResetMultiLevelEnterCount_In)
	OpenMultiLevel(*net.Session, *OpenMultiLevel_In)
	OpenAllPetGrid(*net.Session, *OpenAllPetGrid_In)
	CreateAnnouncement(*net.Session, *CreateAnnouncement_In)
	AddHeart(*net.Session, *AddHeart_In)
	ResetHardLevelEnterCount(*net.Session, *ResetHardLevelEnterCount_In)
	OpenHardLevel(*net.Session, *OpenHardLevel_In)
	SetVipLevel(*net.Session, *SetVipLevel_In)
	SetResourceLevelOpenDay(*net.Session, *SetResourceLevelOpenDay_In)
	ResetResourceLevelOpenDay(*net.Session, *ResetResourceLevelOpenDay_In)
	ResetArenaDailyCount(*net.Session, *ResetArenaDailyCount_In)
	ResetSwordSoulDrawCd(*net.Session, *ResetSwordSoulDrawCd_In)
	SetFirstLoginTime(*net.Session, *SetFirstLoginTime_In)
	EarlierFirstLoginTime(*net.Session, *EarlierFirstLoginTime_In)
	ResetServerOpenTime(*net.Session, *ResetServerOpenTime_In)
	ClearTraderRefreshTime(*net.Session, *ClearTraderRefreshTime_In)
	AddTraderRefreshTime(*net.Session, *AddTraderRefreshTime_In)
	ClearTraderSchedule(*net.Session, *ClearTraderSchedule_In)
	AddTraderSchedule(*net.Session, *AddTraderSchedule_In)
	OpenTown(*net.Session, *OpenTown_In)
	AddGlobalMail(*net.Session, *AddGlobalMail_In)
	CreateAnnouncementWithoutTpl(*net.Session, *CreateAnnouncementWithoutTpl_In)
	SetLoginDay(*net.Session, *SetLoginDay_In)
	ResetLoginAward(*net.Session, *ResetLoginAward_In)
	RestPlayerAwardLock(*net.Session, *RestPlayerAwardLock_In)
	ResetRainbowLevel(*net.Session, *ResetRainbowLevel_In)
	SetRainbowLevel(*net.Session, *SetRainbowLevel_In)
	SendPushNotification(*net.Session, *SendPushNotification_In)
	ResetPetVirtualEnv(*net.Session, *ResetPetVirtualEnv_In)
	AddFame(*net.Session, *AddFame_In)
	AddWorldChatMessage(*net.Session, *AddWorldChatMessage_In)
	MonthCard(*net.Session, *MonthCard_In)
	EnterSandbox(*net.Session, *EnterSandbox_In)
	SandboxRollback(*net.Session, *SandboxRollback_In)
	ExitSandbox(*net.Session, *ExitSandbox_In)
	ResetShadedMissions(*net.Session, *ResetShadedMissions_In)
	CleanCornucopia(*net.Session, *CleanCornucopia_In)
	AddTotem(*net.Session, *AddTotem_In)
	AddRune(*net.Session, *AddRune_In)
	SendRareItemMessage(*net.Session, *SendRareItemMessage_In)
	AddSwordDrivingAction(*net.Session, *AddSwordDrivingAction_In)
	ResetDrivingSwordData(*net.Session, *ResetDrivingSwordData_In)
	AddSwordSoulFragment(*net.Session, *AddSwordSoulFragment_In)
	ResetMoneyTreeStatus(*net.Session, *ResetMoneyTreeStatus_In)
	ResetTodayMoneyTree(*net.Session, *ResetTodayMoneyTree_In)
	CleanSwordSoulIngotDrawNums(*net.Session, *CleanSwordSoulIngotDrawNums_In)
	PunchDrivingSwordCloud(*net.Session, *PunchDrivingSwordCloud_In)
	ClearCliqueDailyDonate(*net.Session, *ClearCliqueDailyDonate_In)
	SetCliqueContrib(*net.Session, *SetCliqueContrib_In)
	RefreshCliqueWorship(*net.Session, *RefreshCliqueWorship_In)
	CliqueEscortHijackBattleWin(*net.Session, *CliqueEscortHijackBattleWin_In)
	CliqueEscortRecoverBattleWin(*net.Session, *CliqueEscortRecoverBattleWin_In)
	CliqueEscortNotifyMessage(*net.Session, *CliqueEscortNotifyMessage_In)
	CliqueEscortNotifyDailyQuest(*net.Session, *CliqueEscortNotifyDailyQuest_In)
	SetCliqueBuildingLevel(*net.Session, *SetCliqueBuildingLevel_In)
	SetCliqueBuildingMoney(*net.Session, *SetCliqueBuildingMoney_In)
	EscortBench(*net.Session, *EscortBench_In)
	ResetCliqueEscortDailyNum(*net.Session, *ResetCliqueEscortDailyNum_In)
	TakeAdditionQuest(*net.Session, *TakeAdditionQuest_In)
	SetMissionStarMax(*net.Session, *SetMissionStarMax_In)
	CliqueBankCd(*net.Session, *CliqueBankCd_In)
	ResetDespairLandBattleNum(*net.Session, *ResetDespairLandBattleNum_In)
	ResetCliqueStoreSendTimes(*net.Session, *ResetCliqueStoreSendTimes_In)
	AddCliqueStoreDonate(*net.Session, *AddCliqueStoreDonate_In)
	PassAllDespairLandLevel(*net.Session, *PassAllDespairLandLevel_In)
	DespairLandDummyBossKill(*net.Session, *DespairLandDummyBossKill_In)
	AddTaoyuanItem(*net.Session, *AddTaoyuanItem_In)
	AddTaoyuanExp(*net.Session, *AddTaoyuanExp_In)
}

type OutHandler interface {
	AddBuddy(*net.Session, *AddBuddy_Out)
	AddItem(*net.Session, *AddItem_Out)
	SetRoleLevel(*net.Session, *SetRoleLevel_Out)
	SetCoins(*net.Session, *SetCoins_Out)
	SetIngot(*net.Session, *SetIngot_Out)
	AddGhost(*net.Session, *AddGhost_Out)
	SetPlayerPhysical(*net.Session, *SetPlayerPhysical_Out)
	ResetLevelEnterCount(*net.Session, *ResetLevelEnterCount_Out)
	AddExp(*net.Session, *AddExp_Out)
	OpenGhostMission(*net.Session, *OpenGhostMission_Out)
	SendMail(*net.Session, *SendMail_Out)
	ClearMail(*net.Session, *ClearMail_Out)
	OpenMissionLevel(*net.Session, *OpenMissionLevel_Out)
	StartBattle(*net.Session, *StartBattle_Out)
	ListenByName(*net.Session, *ListenByName_Out)
	OpenQuest(*net.Session, *OpenQuest_Out)
	OpenFunc(*net.Session, *OpenFunc_Out)
	AddSwordSoul(*net.Session, *AddSwordSoul_Out)
	AddBattlePet(*net.Session, *AddBattlePet_Out)
	ResetMultiLevelEnterCount(*net.Session, *ResetMultiLevelEnterCount_Out)
	OpenMultiLevel(*net.Session, *OpenMultiLevel_Out)
	OpenAllPetGrid(*net.Session, *OpenAllPetGrid_Out)
	CreateAnnouncement(*net.Session, *CreateAnnouncement_Out)
	AddHeart(*net.Session, *AddHeart_Out)
	ResetHardLevelEnterCount(*net.Session, *ResetHardLevelEnterCount_Out)
	OpenHardLevel(*net.Session, *OpenHardLevel_Out)
	SetVipLevel(*net.Session, *SetVipLevel_Out)
	SetResourceLevelOpenDay(*net.Session, *SetResourceLevelOpenDay_Out)
	ResetResourceLevelOpenDay(*net.Session, *ResetResourceLevelOpenDay_Out)
	ResetArenaDailyCount(*net.Session, *ResetArenaDailyCount_Out)
	ResetSwordSoulDrawCd(*net.Session, *ResetSwordSoulDrawCd_Out)
	SetFirstLoginTime(*net.Session, *SetFirstLoginTime_Out)
	EarlierFirstLoginTime(*net.Session, *EarlierFirstLoginTime_Out)
	ResetServerOpenTime(*net.Session, *ResetServerOpenTime_Out)
	ClearTraderRefreshTime(*net.Session, *ClearTraderRefreshTime_Out)
	AddTraderRefreshTime(*net.Session, *AddTraderRefreshTime_Out)
	ClearTraderSchedule(*net.Session, *ClearTraderSchedule_Out)
	AddTraderSchedule(*net.Session, *AddTraderSchedule_Out)
	OpenTown(*net.Session, *OpenTown_Out)
	AddGlobalMail(*net.Session, *AddGlobalMail_Out)
	CreateAnnouncementWithoutTpl(*net.Session, *CreateAnnouncementWithoutTpl_Out)
	SetLoginDay(*net.Session, *SetLoginDay_Out)
	ResetLoginAward(*net.Session, *ResetLoginAward_Out)
	RestPlayerAwardLock(*net.Session, *RestPlayerAwardLock_Out)
	ResetRainbowLevel(*net.Session, *ResetRainbowLevel_Out)
	SetRainbowLevel(*net.Session, *SetRainbowLevel_Out)
	SendPushNotification(*net.Session, *SendPushNotification_Out)
	ResetPetVirtualEnv(*net.Session, *ResetPetVirtualEnv_Out)
	AddFame(*net.Session, *AddFame_Out)
	AddWorldChatMessage(*net.Session, *AddWorldChatMessage_Out)
	MonthCard(*net.Session, *MonthCard_Out)
	EnterSandbox(*net.Session, *EnterSandbox_Out)
	SandboxRollback(*net.Session, *SandboxRollback_Out)
	ExitSandbox(*net.Session, *ExitSandbox_Out)
	ResetShadedMissions(*net.Session, *ResetShadedMissions_Out)
	CleanCornucopia(*net.Session, *CleanCornucopia_Out)
	AddTotem(*net.Session, *AddTotem_Out)
	AddRune(*net.Session, *AddRune_Out)
	SendRareItemMessage(*net.Session, *SendRareItemMessage_Out)
	AddSwordDrivingAction(*net.Session, *AddSwordDrivingAction_Out)
	ResetDrivingSwordData(*net.Session, *ResetDrivingSwordData_Out)
	AddSwordSoulFragment(*net.Session, *AddSwordSoulFragment_Out)
	ResetMoneyTreeStatus(*net.Session, *ResetMoneyTreeStatus_Out)
	ResetTodayMoneyTree(*net.Session, *ResetTodayMoneyTree_Out)
	CleanSwordSoulIngotDrawNums(*net.Session, *CleanSwordSoulIngotDrawNums_Out)
	PunchDrivingSwordCloud(*net.Session, *PunchDrivingSwordCloud_Out)
	ClearCliqueDailyDonate(*net.Session, *ClearCliqueDailyDonate_Out)
	SetCliqueContrib(*net.Session, *SetCliqueContrib_Out)
	RefreshCliqueWorship(*net.Session, *RefreshCliqueWorship_Out)
	CliqueEscortHijackBattleWin(*net.Session, *CliqueEscortHijackBattleWin_Out)
	CliqueEscortRecoverBattleWin(*net.Session, *CliqueEscortRecoverBattleWin_Out)
	CliqueEscortNotifyMessage(*net.Session, *CliqueEscortNotifyMessage_Out)
	CliqueEscortNotifyDailyQuest(*net.Session, *CliqueEscortNotifyDailyQuest_Out)
	SetCliqueBuildingLevel(*net.Session, *SetCliqueBuildingLevel_Out)
	SetCliqueBuildingMoney(*net.Session, *SetCliqueBuildingMoney_Out)
	EscortBench(*net.Session, *EscortBench_Out)
	ResetCliqueEscortDailyNum(*net.Session, *ResetCliqueEscortDailyNum_Out)
	TakeAdditionQuest(*net.Session, *TakeAdditionQuest_Out)
	SetMissionStarMax(*net.Session, *SetMissionStarMax_Out)
	CliqueBankCd(*net.Session, *CliqueBankCd_Out)
	ResetDespairLandBattleNum(*net.Session, *ResetDespairLandBattleNum_Out)
	ResetCliqueStoreSendTimes(*net.Session, *ResetCliqueStoreSendTimes_Out)
	AddCliqueStoreDonate(*net.Session, *AddCliqueStoreDonate_Out)
	PassAllDespairLandLevel(*net.Session, *PassAllDespairLandLevel_Out)
	DespairLandDummyBossKill(*net.Session, *DespairLandDummyBossKill_Out)
	AddTaoyuanItem(*net.Session, *AddTaoyuanItem_Out)
	AddTaoyuanExp(*net.Session, *AddTaoyuanExp_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(AddBuddy_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(AddItem_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(SetRoleLevel_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(SetCoins_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(SetIngot_In)
		request.Decode(buffer)
		return request
	case 11:
		request := new(AddGhost_In)
		request.Decode(buffer)
		return request
	case 12:
		request := new(SetPlayerPhysical_In)
		request.Decode(buffer)
		return request
	case 13:
		request := new(ResetLevelEnterCount_In)
		request.Decode(buffer)
		return request
	case 14:
		request := new(AddExp_In)
		request.Decode(buffer)
		return request
	case 15:
		request := new(OpenGhostMission_In)
		request.Decode(buffer)
		return request
	case 16:
		request := new(SendMail_In)
		request.Decode(buffer)
		return request
	case 17:
		request := new(ClearMail_In)
		request.Decode(buffer)
		return request
	case 18:
		request := new(OpenMissionLevel_In)
		request.Decode(buffer)
		return request
	case 19:
		request := new(StartBattle_In)
		request.Decode(buffer)
		return request
	case 20:
		request := new(ListenByName_In)
		request.Decode(buffer)
		return request
	case 21:
		request := new(OpenQuest_In)
		request.Decode(buffer)
		return request
	case 22:
		request := new(OpenFunc_In)
		request.Decode(buffer)
		return request
	case 23:
		request := new(AddSwordSoul_In)
		request.Decode(buffer)
		return request
	case 25:
		request := new(AddBattlePet_In)
		request.Decode(buffer)
		return request
	case 26:
		request := new(ResetMultiLevelEnterCount_In)
		request.Decode(buffer)
		return request
	case 27:
		request := new(OpenMultiLevel_In)
		request.Decode(buffer)
		return request
	case 28:
		request := new(OpenAllPetGrid_In)
		request.Decode(buffer)
		return request
	case 29:
		request := new(CreateAnnouncement_In)
		request.Decode(buffer)
		return request
	case 30:
		request := new(AddHeart_In)
		request.Decode(buffer)
		return request
	case 31:
		request := new(ResetHardLevelEnterCount_In)
		request.Decode(buffer)
		return request
	case 32:
		request := new(OpenHardLevel_In)
		request.Decode(buffer)
		return request
	case 33:
		request := new(SetVipLevel_In)
		request.Decode(buffer)
		return request
	case 34:
		request := new(SetResourceLevelOpenDay_In)
		request.Decode(buffer)
		return request
	case 35:
		request := new(ResetResourceLevelOpenDay_In)
		request.Decode(buffer)
		return request
	case 36:
		request := new(ResetArenaDailyCount_In)
		request.Decode(buffer)
		return request
	case 37:
		request := new(ResetSwordSoulDrawCd_In)
		request.Decode(buffer)
		return request
	case 38:
		request := new(SetFirstLoginTime_In)
		request.Decode(buffer)
		return request
	case 39:
		request := new(EarlierFirstLoginTime_In)
		request.Decode(buffer)
		return request
	case 40:
		request := new(ResetServerOpenTime_In)
		request.Decode(buffer)
		return request
	case 41:
		request := new(ClearTraderRefreshTime_In)
		request.Decode(buffer)
		return request
	case 42:
		request := new(AddTraderRefreshTime_In)
		request.Decode(buffer)
		return request
	case 43:
		request := new(ClearTraderSchedule_In)
		request.Decode(buffer)
		return request
	case 44:
		request := new(AddTraderSchedule_In)
		request.Decode(buffer)
		return request
	case 45:
		request := new(OpenTown_In)
		request.Decode(buffer)
		return request
	case 46:
		request := new(AddGlobalMail_In)
		request.Decode(buffer)
		return request
	case 47:
		request := new(CreateAnnouncementWithoutTpl_In)
		request.Decode(buffer)
		return request
	case 48:
		request := new(SetLoginDay_In)
		request.Decode(buffer)
		return request
	case 49:
		request := new(ResetLoginAward_In)
		request.Decode(buffer)
		return request
	case 50:
		request := new(RestPlayerAwardLock_In)
		request.Decode(buffer)
		return request
	case 51:
		request := new(ResetRainbowLevel_In)
		request.Decode(buffer)
		return request
	case 52:
		request := new(SetRainbowLevel_In)
		request.Decode(buffer)
		return request
	case 53:
		request := new(SendPushNotification_In)
		request.Decode(buffer)
		return request
	case 54:
		request := new(ResetPetVirtualEnv_In)
		request.Decode(buffer)
		return request
	case 55:
		request := new(AddFame_In)
		request.Decode(buffer)
		return request
	case 56:
		request := new(AddWorldChatMessage_In)
		request.Decode(buffer)
		return request
	case 57:
		request := new(MonthCard_In)
		request.Decode(buffer)
		return request
	case 58:
		request := new(EnterSandbox_In)
		request.Decode(buffer)
		return request
	case 59:
		request := new(SandboxRollback_In)
		request.Decode(buffer)
		return request
	case 60:
		request := new(ExitSandbox_In)
		request.Decode(buffer)
		return request
	case 61:
		request := new(ResetShadedMissions_In)
		request.Decode(buffer)
		return request
	case 62:
		request := new(CleanCornucopia_In)
		request.Decode(buffer)
		return request
	case 63:
		request := new(AddTotem_In)
		request.Decode(buffer)
		return request
	case 64:
		request := new(AddRune_In)
		request.Decode(buffer)
		return request
	case 65:
		request := new(SendRareItemMessage_In)
		request.Decode(buffer)
		return request
	case 66:
		request := new(AddSwordDrivingAction_In)
		request.Decode(buffer)
		return request
	case 67:
		request := new(ResetDrivingSwordData_In)
		request.Decode(buffer)
		return request
	case 68:
		request := new(AddSwordSoulFragment_In)
		request.Decode(buffer)
		return request
	case 69:
		request := new(ResetMoneyTreeStatus_In)
		request.Decode(buffer)
		return request
	case 70:
		request := new(ResetTodayMoneyTree_In)
		request.Decode(buffer)
		return request
	case 71:
		request := new(CleanSwordSoulIngotDrawNums_In)
		request.Decode(buffer)
		return request
	case 72:
		request := new(PunchDrivingSwordCloud_In)
		request.Decode(buffer)
		return request
	case 73:
		request := new(ClearCliqueDailyDonate_In)
		request.Decode(buffer)
		return request
	case 74:
		request := new(SetCliqueContrib_In)
		request.Decode(buffer)
		return request
	case 75:
		request := new(RefreshCliqueWorship_In)
		request.Decode(buffer)
		return request
	case 76:
		request := new(CliqueEscortHijackBattleWin_In)
		request.Decode(buffer)
		return request
	case 77:
		request := new(CliqueEscortRecoverBattleWin_In)
		request.Decode(buffer)
		return request
	case 87:
		request := new(CliqueEscortNotifyMessage_In)
		request.Decode(buffer)
		return request
	case 88:
		request := new(CliqueEscortNotifyDailyQuest_In)
		request.Decode(buffer)
		return request
	case 89:
		request := new(SetCliqueBuildingLevel_In)
		request.Decode(buffer)
		return request
	case 90:
		request := new(SetCliqueBuildingMoney_In)
		request.Decode(buffer)
		return request
	case 91:
		request := new(EscortBench_In)
		request.Decode(buffer)
		return request
	case 92:
		request := new(ResetCliqueEscortDailyNum_In)
		request.Decode(buffer)
		return request
	case 93:
		request := new(TakeAdditionQuest_In)
		request.Decode(buffer)
		return request
	case 94:
		request := new(SetMissionStarMax_In)
		request.Decode(buffer)
		return request
	case 95:
		request := new(CliqueBankCd_In)
		request.Decode(buffer)
		return request
	case 96:
		request := new(ResetDespairLandBattleNum_In)
		request.Decode(buffer)
		return request
	case 97:
		request := new(ResetCliqueStoreSendTimes_In)
		request.Decode(buffer)
		return request
	case 98:
		request := new(AddCliqueStoreDonate_In)
		request.Decode(buffer)
		return request
	case 99:
		request := new(PassAllDespairLandLevel_In)
		request.Decode(buffer)
		return request
	case 100:
		request := new(DespairLandDummyBossKill_In)
		request.Decode(buffer)
		return request
	case 101:
		request := new(AddTaoyuanItem_In)
		request.Decode(buffer)
		return request
	case 102:
		request := new(AddTaoyuanExp_In)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported request")
}

func DecodeOut(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(AddBuddy_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(AddItem_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(SetRoleLevel_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(SetCoins_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(SetIngot_Out)
		request.Decode(buffer)
		return request
	case 11:
		request := new(AddGhost_Out)
		request.Decode(buffer)
		return request
	case 12:
		request := new(SetPlayerPhysical_Out)
		request.Decode(buffer)
		return request
	case 13:
		request := new(ResetLevelEnterCount_Out)
		request.Decode(buffer)
		return request
	case 14:
		request := new(AddExp_Out)
		request.Decode(buffer)
		return request
	case 15:
		request := new(OpenGhostMission_Out)
		request.Decode(buffer)
		return request
	case 16:
		request := new(SendMail_Out)
		request.Decode(buffer)
		return request
	case 17:
		request := new(ClearMail_Out)
		request.Decode(buffer)
		return request
	case 18:
		request := new(OpenMissionLevel_Out)
		request.Decode(buffer)
		return request
	case 19:
		request := new(StartBattle_Out)
		request.Decode(buffer)
		return request
	case 20:
		request := new(ListenByName_Out)
		request.Decode(buffer)
		return request
	case 21:
		request := new(OpenQuest_Out)
		request.Decode(buffer)
		return request
	case 22:
		request := new(OpenFunc_Out)
		request.Decode(buffer)
		return request
	case 23:
		request := new(AddSwordSoul_Out)
		request.Decode(buffer)
		return request
	case 25:
		request := new(AddBattlePet_Out)
		request.Decode(buffer)
		return request
	case 26:
		request := new(ResetMultiLevelEnterCount_Out)
		request.Decode(buffer)
		return request
	case 27:
		request := new(OpenMultiLevel_Out)
		request.Decode(buffer)
		return request
	case 28:
		request := new(OpenAllPetGrid_Out)
		request.Decode(buffer)
		return request
	case 29:
		request := new(CreateAnnouncement_Out)
		request.Decode(buffer)
		return request
	case 30:
		request := new(AddHeart_Out)
		request.Decode(buffer)
		return request
	case 31:
		request := new(ResetHardLevelEnterCount_Out)
		request.Decode(buffer)
		return request
	case 32:
		request := new(OpenHardLevel_Out)
		request.Decode(buffer)
		return request
	case 33:
		request := new(SetVipLevel_Out)
		request.Decode(buffer)
		return request
	case 34:
		request := new(SetResourceLevelOpenDay_Out)
		request.Decode(buffer)
		return request
	case 35:
		request := new(ResetResourceLevelOpenDay_Out)
		request.Decode(buffer)
		return request
	case 36:
		request := new(ResetArenaDailyCount_Out)
		request.Decode(buffer)
		return request
	case 37:
		request := new(ResetSwordSoulDrawCd_Out)
		request.Decode(buffer)
		return request
	case 38:
		request := new(SetFirstLoginTime_Out)
		request.Decode(buffer)
		return request
	case 39:
		request := new(EarlierFirstLoginTime_Out)
		request.Decode(buffer)
		return request
	case 40:
		request := new(ResetServerOpenTime_Out)
		request.Decode(buffer)
		return request
	case 41:
		request := new(ClearTraderRefreshTime_Out)
		request.Decode(buffer)
		return request
	case 42:
		request := new(AddTraderRefreshTime_Out)
		request.Decode(buffer)
		return request
	case 43:
		request := new(ClearTraderSchedule_Out)
		request.Decode(buffer)
		return request
	case 44:
		request := new(AddTraderSchedule_Out)
		request.Decode(buffer)
		return request
	case 45:
		request := new(OpenTown_Out)
		request.Decode(buffer)
		return request
	case 46:
		request := new(AddGlobalMail_Out)
		request.Decode(buffer)
		return request
	case 47:
		request := new(CreateAnnouncementWithoutTpl_Out)
		request.Decode(buffer)
		return request
	case 48:
		request := new(SetLoginDay_Out)
		request.Decode(buffer)
		return request
	case 49:
		request := new(ResetLoginAward_Out)
		request.Decode(buffer)
		return request
	case 50:
		request := new(RestPlayerAwardLock_Out)
		request.Decode(buffer)
		return request
	case 51:
		request := new(ResetRainbowLevel_Out)
		request.Decode(buffer)
		return request
	case 52:
		request := new(SetRainbowLevel_Out)
		request.Decode(buffer)
		return request
	case 53:
		request := new(SendPushNotification_Out)
		request.Decode(buffer)
		return request
	case 54:
		request := new(ResetPetVirtualEnv_Out)
		request.Decode(buffer)
		return request
	case 55:
		request := new(AddFame_Out)
		request.Decode(buffer)
		return request
	case 56:
		request := new(AddWorldChatMessage_Out)
		request.Decode(buffer)
		return request
	case 57:
		request := new(MonthCard_Out)
		request.Decode(buffer)
		return request
	case 58:
		request := new(EnterSandbox_Out)
		request.Decode(buffer)
		return request
	case 59:
		request := new(SandboxRollback_Out)
		request.Decode(buffer)
		return request
	case 60:
		request := new(ExitSandbox_Out)
		request.Decode(buffer)
		return request
	case 61:
		request := new(ResetShadedMissions_Out)
		request.Decode(buffer)
		return request
	case 62:
		request := new(CleanCornucopia_Out)
		request.Decode(buffer)
		return request
	case 63:
		request := new(AddTotem_Out)
		request.Decode(buffer)
		return request
	case 64:
		request := new(AddRune_Out)
		request.Decode(buffer)
		return request
	case 65:
		request := new(SendRareItemMessage_Out)
		request.Decode(buffer)
		return request
	case 66:
		request := new(AddSwordDrivingAction_Out)
		request.Decode(buffer)
		return request
	case 67:
		request := new(ResetDrivingSwordData_Out)
		request.Decode(buffer)
		return request
	case 68:
		request := new(AddSwordSoulFragment_Out)
		request.Decode(buffer)
		return request
	case 69:
		request := new(ResetMoneyTreeStatus_Out)
		request.Decode(buffer)
		return request
	case 70:
		request := new(ResetTodayMoneyTree_Out)
		request.Decode(buffer)
		return request
	case 71:
		request := new(CleanSwordSoulIngotDrawNums_Out)
		request.Decode(buffer)
		return request
	case 72:
		request := new(PunchDrivingSwordCloud_Out)
		request.Decode(buffer)
		return request
	case 73:
		request := new(ClearCliqueDailyDonate_Out)
		request.Decode(buffer)
		return request
	case 74:
		request := new(SetCliqueContrib_Out)
		request.Decode(buffer)
		return request
	case 75:
		request := new(RefreshCliqueWorship_Out)
		request.Decode(buffer)
		return request
	case 76:
		request := new(CliqueEscortHijackBattleWin_Out)
		request.Decode(buffer)
		return request
	case 77:
		request := new(CliqueEscortRecoverBattleWin_Out)
		request.Decode(buffer)
		return request
	case 87:
		request := new(CliqueEscortNotifyMessage_Out)
		request.Decode(buffer)
		return request
	case 88:
		request := new(CliqueEscortNotifyDailyQuest_Out)
		request.Decode(buffer)
		return request
	case 89:
		request := new(SetCliqueBuildingLevel_Out)
		request.Decode(buffer)
		return request
	case 90:
		request := new(SetCliqueBuildingMoney_Out)
		request.Decode(buffer)
		return request
	case 91:
		request := new(EscortBench_Out)
		request.Decode(buffer)
		return request
	case 92:
		request := new(ResetCliqueEscortDailyNum_Out)
		request.Decode(buffer)
		return request
	case 93:
		request := new(TakeAdditionQuest_Out)
		request.Decode(buffer)
		return request
	case 94:
		request := new(SetMissionStarMax_Out)
		request.Decode(buffer)
		return request
	case 95:
		request := new(CliqueBankCd_Out)
		request.Decode(buffer)
		return request
	case 96:
		request := new(ResetDespairLandBattleNum_Out)
		request.Decode(buffer)
		return request
	case 97:
		request := new(ResetCliqueStoreSendTimes_Out)
		request.Decode(buffer)
		return request
	case 98:
		request := new(AddCliqueStoreDonate_Out)
		request.Decode(buffer)
		return request
	case 99:
		request := new(PassAllDespairLandLevel_Out)
		request.Decode(buffer)
		return request
	case 100:
		request := new(DespairLandDummyBossKill_Out)
		request.Decode(buffer)
		return request
	case 101:
		request := new(AddTaoyuanItem_Out)
		request.Decode(buffer)
		return request
	case 102:
		request := new(AddTaoyuanExp_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type AddBuddy_In struct {
	RoleId int8 `json:"role_id"`
}

func (this *AddBuddy_In) Process(session *net.Session) {
	g_InHandler.AddBuddy(session, this)
}

func (this *AddBuddy_In) TypeName() string {
	return "debug.add_buddy.in"
}

func (this *AddBuddy_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 0
}

type AddBuddy_Out struct {
}

func (this *AddBuddy_Out) Process(session *net.Session) {
	g_OutHandler.AddBuddy(session, this)
}

func (this *AddBuddy_Out) TypeName() string {
	return "debug.add_buddy.out"
}

func (this *AddBuddy_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 0
}

func (this *AddBuddy_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AddItem_In struct {
	ItemId int16 `json:"item_id"`
	Number int16 `json:"number"`
}

func (this *AddItem_In) Process(session *net.Session) {
	g_InHandler.AddItem(session, this)
}

func (this *AddItem_In) TypeName() string {
	return "debug.add_item.in"
}

func (this *AddItem_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 2
}

type AddItem_Out struct {
}

func (this *AddItem_Out) Process(session *net.Session) {
	g_OutHandler.AddItem(session, this)
}

func (this *AddItem_Out) TypeName() string {
	return "debug.add_item.out"
}

func (this *AddItem_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 2
}

func (this *AddItem_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SetRoleLevel_In struct {
	RoleId int8  `json:"role_id"`
	Level  int16 `json:"level"`
}

func (this *SetRoleLevel_In) Process(session *net.Session) {
	g_InHandler.SetRoleLevel(session, this)
}

func (this *SetRoleLevel_In) TypeName() string {
	return "debug.set_role_level.in"
}

func (this *SetRoleLevel_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 3
}

type SetRoleLevel_Out struct {
}

func (this *SetRoleLevel_Out) Process(session *net.Session) {
	g_OutHandler.SetRoleLevel(session, this)
}

func (this *SetRoleLevel_Out) TypeName() string {
	return "debug.set_role_level.out"
}

func (this *SetRoleLevel_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 3
}

func (this *SetRoleLevel_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SetCoins_In struct {
	Number int64 `json:"number"`
}

func (this *SetCoins_In) Process(session *net.Session) {
	g_InHandler.SetCoins(session, this)
}

func (this *SetCoins_In) TypeName() string {
	return "debug.set_coins.in"
}

func (this *SetCoins_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 4
}

type SetCoins_Out struct {
}

func (this *SetCoins_Out) Process(session *net.Session) {
	g_OutHandler.SetCoins(session, this)
}

func (this *SetCoins_Out) TypeName() string {
	return "debug.set_coins.out"
}

func (this *SetCoins_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 4
}

func (this *SetCoins_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SetIngot_In struct {
	Number int64 `json:"number"`
}

func (this *SetIngot_In) Process(session *net.Session) {
	g_InHandler.SetIngot(session, this)
}

func (this *SetIngot_In) TypeName() string {
	return "debug.set_ingot.in"
}

func (this *SetIngot_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 5
}

type SetIngot_Out struct {
}

func (this *SetIngot_Out) Process(session *net.Session) {
	g_OutHandler.SetIngot(session, this)
}

func (this *SetIngot_Out) TypeName() string {
	return "debug.set_ingot.out"
}

func (this *SetIngot_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 5
}

func (this *SetIngot_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AddGhost_In struct {
	GhostId int16 `json:"ghost_id"`
}

func (this *AddGhost_In) Process(session *net.Session) {
	g_InHandler.AddGhost(session, this)
}

func (this *AddGhost_In) TypeName() string {
	return "debug.add_ghost.in"
}

func (this *AddGhost_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 11
}

type AddGhost_Out struct {
}

func (this *AddGhost_Out) Process(session *net.Session) {
	g_OutHandler.AddGhost(session, this)
}

func (this *AddGhost_Out) TypeName() string {
	return "debug.add_ghost.out"
}

func (this *AddGhost_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 11
}

func (this *AddGhost_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SetPlayerPhysical_In struct {
	Physical int16 `json:"physical"`
}

func (this *SetPlayerPhysical_In) Process(session *net.Session) {
	g_InHandler.SetPlayerPhysical(session, this)
}

func (this *SetPlayerPhysical_In) TypeName() string {
	return "debug.set_player_physical.in"
}

func (this *SetPlayerPhysical_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 12
}

type SetPlayerPhysical_Out struct {
}

func (this *SetPlayerPhysical_Out) Process(session *net.Session) {
	g_OutHandler.SetPlayerPhysical(session, this)
}

func (this *SetPlayerPhysical_Out) TypeName() string {
	return "debug.set_player_physical.out"
}

func (this *SetPlayerPhysical_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 12
}

func (this *SetPlayerPhysical_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ResetLevelEnterCount_In struct {
	LevelId int32 `json:"level_id"`
}

func (this *ResetLevelEnterCount_In) Process(session *net.Session) {
	g_InHandler.ResetLevelEnterCount(session, this)
}

func (this *ResetLevelEnterCount_In) TypeName() string {
	return "debug.reset_level_enter_count.in"
}

func (this *ResetLevelEnterCount_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 13
}

type ResetLevelEnterCount_Out struct {
}

func (this *ResetLevelEnterCount_Out) Process(session *net.Session) {
	g_OutHandler.ResetLevelEnterCount(session, this)
}

func (this *ResetLevelEnterCount_Out) TypeName() string {
	return "debug.reset_level_enter_count.out"
}

func (this *ResetLevelEnterCount_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 13
}

func (this *ResetLevelEnterCount_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AddExp_In struct {
	RoleId int8  `json:"role_id"`
	AddExp int64 `json:"add_exp"`
}

func (this *AddExp_In) Process(session *net.Session) {
	g_InHandler.AddExp(session, this)
}

func (this *AddExp_In) TypeName() string {
	return "debug.add_exp.in"
}

func (this *AddExp_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 14
}

type AddExp_Out struct {
}

func (this *AddExp_Out) Process(session *net.Session) {
	g_OutHandler.AddExp(session, this)
}

func (this *AddExp_Out) TypeName() string {
	return "debug.add_exp.out"
}

func (this *AddExp_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 14
}

func (this *AddExp_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type OpenGhostMission_In struct {
	MissionId int16 `json:"mission_id"`
}

func (this *OpenGhostMission_In) Process(session *net.Session) {
	g_InHandler.OpenGhostMission(session, this)
}

func (this *OpenGhostMission_In) TypeName() string {
	return "debug.open_ghost_mission.in"
}

func (this *OpenGhostMission_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 15
}

type OpenGhostMission_Out struct {
}

func (this *OpenGhostMission_Out) Process(session *net.Session) {
	g_OutHandler.OpenGhostMission(session, this)
}

func (this *OpenGhostMission_Out) TypeName() string {
	return "debug.open_ghost_mission.out"
}

func (this *OpenGhostMission_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 15
}

func (this *OpenGhostMission_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SendMail_In struct {
	MailId int32 `json:"mail_id"`
}

func (this *SendMail_In) Process(session *net.Session) {
	g_InHandler.SendMail(session, this)
}

func (this *SendMail_In) TypeName() string {
	return "debug.send_mail.in"
}

func (this *SendMail_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 16
}

type SendMail_Out struct {
}

func (this *SendMail_Out) Process(session *net.Session) {
	g_OutHandler.SendMail(session, this)
}

func (this *SendMail_Out) TypeName() string {
	return "debug.send_mail.out"
}

func (this *SendMail_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 16
}

func (this *SendMail_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ClearMail_In struct {
}

func (this *ClearMail_In) Process(session *net.Session) {
	g_InHandler.ClearMail(session, this)
}

func (this *ClearMail_In) TypeName() string {
	return "debug.clear_mail.in"
}

func (this *ClearMail_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 17
}

type ClearMail_Out struct {
}

func (this *ClearMail_Out) Process(session *net.Session) {
	g_OutHandler.ClearMail(session, this)
}

func (this *ClearMail_Out) TypeName() string {
	return "debug.clear_mail.out"
}

func (this *ClearMail_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 17
}

func (this *ClearMail_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type OpenMissionLevel_In struct {
	LevelId int32 `json:"level_id"`
}

func (this *OpenMissionLevel_In) Process(session *net.Session) {
	g_InHandler.OpenMissionLevel(session, this)
}

func (this *OpenMissionLevel_In) TypeName() string {
	return "debug.open_mission_level.in"
}

func (this *OpenMissionLevel_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 18
}

type OpenMissionLevel_Out struct {
}

func (this *OpenMissionLevel_Out) Process(session *net.Session) {
	g_OutHandler.OpenMissionLevel(session, this)
}

func (this *OpenMissionLevel_Out) TypeName() string {
	return "debug.open_mission_level.out"
}

func (this *OpenMissionLevel_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 18
}

func (this *OpenMissionLevel_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type StartBattle_In struct {
	BattleType int8  `json:"battle_type"`
	EnemyId    int32 `json:"enemy_id"`
}

func (this *StartBattle_In) Process(session *net.Session) {
	g_InHandler.StartBattle(session, this)
}

func (this *StartBattle_In) TypeName() string {
	return "debug.start_battle.in"
}

func (this *StartBattle_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 19
}

type StartBattle_Out struct {
}

func (this *StartBattle_Out) Process(session *net.Session) {
	g_OutHandler.StartBattle(session, this)
}

func (this *StartBattle_Out) TypeName() string {
	return "debug.start_battle.out"
}

func (this *StartBattle_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 19
}

func (this *StartBattle_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ListenByName_In struct {
	Name []byte `json:"name"`
}

func (this *ListenByName_In) Process(session *net.Session) {
	g_InHandler.ListenByName(session, this)
}

func (this *ListenByName_In) TypeName() string {
	return "debug.listen_by_name.in"
}

func (this *ListenByName_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 20
}

type ListenByName_Out struct {
}

func (this *ListenByName_Out) Process(session *net.Session) {
	g_OutHandler.ListenByName(session, this)
}

func (this *ListenByName_Out) TypeName() string {
	return "debug.listen_by_name.out"
}

func (this *ListenByName_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 20
}

func (this *ListenByName_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type OpenQuest_In struct {
	QuestId int16 `json:"quest_id"`
}

func (this *OpenQuest_In) Process(session *net.Session) {
	g_InHandler.OpenQuest(session, this)
}

func (this *OpenQuest_In) TypeName() string {
	return "debug.open_quest.in"
}

func (this *OpenQuest_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 21
}

type OpenQuest_Out struct {
}

func (this *OpenQuest_Out) Process(session *net.Session) {
	g_OutHandler.OpenQuest(session, this)
}

func (this *OpenQuest_Out) TypeName() string {
	return "debug.open_quest.out"
}

func (this *OpenQuest_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 21
}

func (this *OpenQuest_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type OpenFunc_In struct {
	Lock int16 `json:"lock"`
}

func (this *OpenFunc_In) Process(session *net.Session) {
	g_InHandler.OpenFunc(session, this)
}

func (this *OpenFunc_In) TypeName() string {
	return "debug.open_func.in"
}

func (this *OpenFunc_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 22
}

type OpenFunc_Out struct {
}

func (this *OpenFunc_Out) Process(session *net.Session) {
	g_OutHandler.OpenFunc(session, this)
}

func (this *OpenFunc_Out) TypeName() string {
	return "debug.open_func.out"
}

func (this *OpenFunc_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 22
}

func (this *OpenFunc_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AddSwordSoul_In struct {
	SwordSoulId int16 `json:"sword_soul_id"`
}

func (this *AddSwordSoul_In) Process(session *net.Session) {
	g_InHandler.AddSwordSoul(session, this)
}

func (this *AddSwordSoul_In) TypeName() string {
	return "debug.add_sword_soul.in"
}

func (this *AddSwordSoul_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 23
}

type AddSwordSoul_Out struct {
}

func (this *AddSwordSoul_Out) Process(session *net.Session) {
	g_OutHandler.AddSwordSoul(session, this)
}

func (this *AddSwordSoul_Out) TypeName() string {
	return "debug.add_sword_soul.out"
}

func (this *AddSwordSoul_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 23
}

func (this *AddSwordSoul_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AddBattlePet_In struct {
	PetId int16 `json:"petId"`
}

func (this *AddBattlePet_In) Process(session *net.Session) {
	g_InHandler.AddBattlePet(session, this)
}

func (this *AddBattlePet_In) TypeName() string {
	return "debug.add_battle_pet.in"
}

func (this *AddBattlePet_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 25
}

type AddBattlePet_Out struct {
}

func (this *AddBattlePet_Out) Process(session *net.Session) {
	g_OutHandler.AddBattlePet(session, this)
}

func (this *AddBattlePet_Out) TypeName() string {
	return "debug.add_battle_pet.out"
}

func (this *AddBattlePet_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 25
}

func (this *AddBattlePet_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ResetMultiLevelEnterCount_In struct {
}

func (this *ResetMultiLevelEnterCount_In) Process(session *net.Session) {
	g_InHandler.ResetMultiLevelEnterCount(session, this)
}

func (this *ResetMultiLevelEnterCount_In) TypeName() string {
	return "debug.reset_multi_level_enter_count.in"
}

func (this *ResetMultiLevelEnterCount_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 26
}

type ResetMultiLevelEnterCount_Out struct {
}

func (this *ResetMultiLevelEnterCount_Out) Process(session *net.Session) {
	g_OutHandler.ResetMultiLevelEnterCount(session, this)
}

func (this *ResetMultiLevelEnterCount_Out) TypeName() string {
	return "debug.reset_multi_level_enter_count.out"
}

func (this *ResetMultiLevelEnterCount_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 26
}

func (this *ResetMultiLevelEnterCount_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type OpenMultiLevel_In struct {
	LevelId int16 `json:"level_id"`
}

func (this *OpenMultiLevel_In) Process(session *net.Session) {
	g_InHandler.OpenMultiLevel(session, this)
}

func (this *OpenMultiLevel_In) TypeName() string {
	return "debug.open_multi_level.in"
}

func (this *OpenMultiLevel_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 27
}

type OpenMultiLevel_Out struct {
}

func (this *OpenMultiLevel_Out) Process(session *net.Session) {
	g_OutHandler.OpenMultiLevel(session, this)
}

func (this *OpenMultiLevel_Out) TypeName() string {
	return "debug.open_multi_level.out"
}

func (this *OpenMultiLevel_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 27
}

func (this *OpenMultiLevel_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type OpenAllPetGrid_In struct {
}

func (this *OpenAllPetGrid_In) Process(session *net.Session) {
	g_InHandler.OpenAllPetGrid(session, this)
}

func (this *OpenAllPetGrid_In) TypeName() string {
	return "debug.open_all_pet_grid.in"
}

func (this *OpenAllPetGrid_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 28
}

type OpenAllPetGrid_Out struct {
}

func (this *OpenAllPetGrid_Out) Process(session *net.Session) {
	g_OutHandler.OpenAllPetGrid(session, this)
}

func (this *OpenAllPetGrid_Out) TypeName() string {
	return "debug.open_all_pet_grid.out"
}

func (this *OpenAllPetGrid_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 28
}

func (this *OpenAllPetGrid_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CreateAnnouncement_In struct {
}

func (this *CreateAnnouncement_In) Process(session *net.Session) {
	g_InHandler.CreateAnnouncement(session, this)
}

func (this *CreateAnnouncement_In) TypeName() string {
	return "debug.create_announcement.in"
}

func (this *CreateAnnouncement_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 29
}

type CreateAnnouncement_Out struct {
}

func (this *CreateAnnouncement_Out) Process(session *net.Session) {
	g_OutHandler.CreateAnnouncement(session, this)
}

func (this *CreateAnnouncement_Out) TypeName() string {
	return "debug.create_announcement.out"
}

func (this *CreateAnnouncement_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 29
}

func (this *CreateAnnouncement_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AddHeart_In struct {
	Number int16 `json:"number"`
}

func (this *AddHeart_In) Process(session *net.Session) {
	g_InHandler.AddHeart(session, this)
}

func (this *AddHeart_In) TypeName() string {
	return "debug.add_heart.in"
}

func (this *AddHeart_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 30
}

type AddHeart_Out struct {
}

func (this *AddHeart_Out) Process(session *net.Session) {
	g_OutHandler.AddHeart(session, this)
}

func (this *AddHeart_Out) TypeName() string {
	return "debug.add_heart.out"
}

func (this *AddHeart_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 30
}

func (this *AddHeart_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ResetHardLevelEnterCount_In struct {
}

func (this *ResetHardLevelEnterCount_In) Process(session *net.Session) {
	g_InHandler.ResetHardLevelEnterCount(session, this)
}

func (this *ResetHardLevelEnterCount_In) TypeName() string {
	return "debug.reset_hard_level_enter_count.in"
}

func (this *ResetHardLevelEnterCount_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 31
}

type ResetHardLevelEnterCount_Out struct {
}

func (this *ResetHardLevelEnterCount_Out) Process(session *net.Session) {
	g_OutHandler.ResetHardLevelEnterCount(session, this)
}

func (this *ResetHardLevelEnterCount_Out) TypeName() string {
	return "debug.reset_hard_level_enter_count.out"
}

func (this *ResetHardLevelEnterCount_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 31
}

func (this *ResetHardLevelEnterCount_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type OpenHardLevel_In struct {
	LevelId int32 `json:"level_id"`
}

func (this *OpenHardLevel_In) Process(session *net.Session) {
	g_InHandler.OpenHardLevel(session, this)
}

func (this *OpenHardLevel_In) TypeName() string {
	return "debug.open_hard_level.in"
}

func (this *OpenHardLevel_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 32
}

type OpenHardLevel_Out struct {
}

func (this *OpenHardLevel_Out) Process(session *net.Session) {
	g_OutHandler.OpenHardLevel(session, this)
}

func (this *OpenHardLevel_Out) TypeName() string {
	return "debug.open_hard_level.out"
}

func (this *OpenHardLevel_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 32
}

func (this *OpenHardLevel_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SetVipLevel_In struct {
	Level int16 `json:"level"`
}

func (this *SetVipLevel_In) Process(session *net.Session) {
	g_InHandler.SetVipLevel(session, this)
}

func (this *SetVipLevel_In) TypeName() string {
	return "debug.set_vip_level.in"
}

func (this *SetVipLevel_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 33
}

type SetVipLevel_Out struct {
}

func (this *SetVipLevel_Out) Process(session *net.Session) {
	g_OutHandler.SetVipLevel(session, this)
}

func (this *SetVipLevel_Out) TypeName() string {
	return "debug.set_vip_level.out"
}

func (this *SetVipLevel_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 33
}

func (this *SetVipLevel_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SetResourceLevelOpenDay_In struct {
	LevelType int8 `json:"level_type"`
	OpenDay   int8 `json:"open_day"`
}

func (this *SetResourceLevelOpenDay_In) Process(session *net.Session) {
	g_InHandler.SetResourceLevelOpenDay(session, this)
}

func (this *SetResourceLevelOpenDay_In) TypeName() string {
	return "debug.set_resource_level_open_day.in"
}

func (this *SetResourceLevelOpenDay_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 34
}

type SetResourceLevelOpenDay_Out struct {
}

func (this *SetResourceLevelOpenDay_Out) Process(session *net.Session) {
	g_OutHandler.SetResourceLevelOpenDay(session, this)
}

func (this *SetResourceLevelOpenDay_Out) TypeName() string {
	return "debug.set_resource_level_open_day.out"
}

func (this *SetResourceLevelOpenDay_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 34
}

func (this *SetResourceLevelOpenDay_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ResetResourceLevelOpenDay_In struct {
}

func (this *ResetResourceLevelOpenDay_In) Process(session *net.Session) {
	g_InHandler.ResetResourceLevelOpenDay(session, this)
}

func (this *ResetResourceLevelOpenDay_In) TypeName() string {
	return "debug.reset_resource_level_open_day.in"
}

func (this *ResetResourceLevelOpenDay_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 35
}

type ResetResourceLevelOpenDay_Out struct {
}

func (this *ResetResourceLevelOpenDay_Out) Process(session *net.Session) {
	g_OutHandler.ResetResourceLevelOpenDay(session, this)
}

func (this *ResetResourceLevelOpenDay_Out) TypeName() string {
	return "debug.reset_resource_level_open_day.out"
}

func (this *ResetResourceLevelOpenDay_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 35
}

func (this *ResetResourceLevelOpenDay_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ResetArenaDailyCount_In struct {
}

func (this *ResetArenaDailyCount_In) Process(session *net.Session) {
	g_InHandler.ResetArenaDailyCount(session, this)
}

func (this *ResetArenaDailyCount_In) TypeName() string {
	return "debug.reset_arena_daily_count.in"
}

func (this *ResetArenaDailyCount_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 36
}

type ResetArenaDailyCount_Out struct {
}

func (this *ResetArenaDailyCount_Out) Process(session *net.Session) {
	g_OutHandler.ResetArenaDailyCount(session, this)
}

func (this *ResetArenaDailyCount_Out) TypeName() string {
	return "debug.reset_arena_daily_count.out"
}

func (this *ResetArenaDailyCount_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 36
}

func (this *ResetArenaDailyCount_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ResetSwordSoulDrawCd_In struct {
}

func (this *ResetSwordSoulDrawCd_In) Process(session *net.Session) {
	g_InHandler.ResetSwordSoulDrawCd(session, this)
}

func (this *ResetSwordSoulDrawCd_In) TypeName() string {
	return "debug.reset_sword_soul_draw_cd.in"
}

func (this *ResetSwordSoulDrawCd_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 37
}

type ResetSwordSoulDrawCd_Out struct {
}

func (this *ResetSwordSoulDrawCd_Out) Process(session *net.Session) {
	g_OutHandler.ResetSwordSoulDrawCd(session, this)
}

func (this *ResetSwordSoulDrawCd_Out) TypeName() string {
	return "debug.reset_sword_soul_draw_cd.out"
}

func (this *ResetSwordSoulDrawCd_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 37
}

func (this *ResetSwordSoulDrawCd_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SetFirstLoginTime_In struct {
	Timestamp int64 `json:"timestamp"`
}

func (this *SetFirstLoginTime_In) Process(session *net.Session) {
	g_InHandler.SetFirstLoginTime(session, this)
}

func (this *SetFirstLoginTime_In) TypeName() string {
	return "debug.set_first_login_time.in"
}

func (this *SetFirstLoginTime_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 38
}

type SetFirstLoginTime_Out struct {
}

func (this *SetFirstLoginTime_Out) Process(session *net.Session) {
	g_OutHandler.SetFirstLoginTime(session, this)
}

func (this *SetFirstLoginTime_Out) TypeName() string {
	return "debug.set_first_login_time.out"
}

func (this *SetFirstLoginTime_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 38
}

func (this *SetFirstLoginTime_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type EarlierFirstLoginTime_In struct {
}

func (this *EarlierFirstLoginTime_In) Process(session *net.Session) {
	g_InHandler.EarlierFirstLoginTime(session, this)
}

func (this *EarlierFirstLoginTime_In) TypeName() string {
	return "debug.earlier_first_login_time.in"
}

func (this *EarlierFirstLoginTime_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 39
}

type EarlierFirstLoginTime_Out struct {
}

func (this *EarlierFirstLoginTime_Out) Process(session *net.Session) {
	g_OutHandler.EarlierFirstLoginTime(session, this)
}

func (this *EarlierFirstLoginTime_Out) TypeName() string {
	return "debug.earlier_first_login_time.out"
}

func (this *EarlierFirstLoginTime_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 39
}

func (this *EarlierFirstLoginTime_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ResetServerOpenTime_In struct {
}

func (this *ResetServerOpenTime_In) Process(session *net.Session) {
	g_InHandler.ResetServerOpenTime(session, this)
}

func (this *ResetServerOpenTime_In) TypeName() string {
	return "debug.reset_server_open_time.in"
}

func (this *ResetServerOpenTime_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 40
}

type ResetServerOpenTime_Out struct {
}

func (this *ResetServerOpenTime_Out) Process(session *net.Session) {
	g_OutHandler.ResetServerOpenTime(session, this)
}

func (this *ResetServerOpenTime_Out) TypeName() string {
	return "debug.reset_server_open_time.out"
}

func (this *ResetServerOpenTime_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 40
}

func (this *ResetServerOpenTime_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ClearTraderRefreshTime_In struct {
	TraderId int16 `json:"trader_id"`
}

func (this *ClearTraderRefreshTime_In) Process(session *net.Session) {
	g_InHandler.ClearTraderRefreshTime(session, this)
}

func (this *ClearTraderRefreshTime_In) TypeName() string {
	return "debug.clear_trader_refresh_time.in"
}

func (this *ClearTraderRefreshTime_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 41
}

type ClearTraderRefreshTime_Out struct {
}

func (this *ClearTraderRefreshTime_Out) Process(session *net.Session) {
	g_OutHandler.ClearTraderRefreshTime(session, this)
}

func (this *ClearTraderRefreshTime_Out) TypeName() string {
	return "debug.clear_trader_refresh_time.out"
}

func (this *ClearTraderRefreshTime_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 41
}

func (this *ClearTraderRefreshTime_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AddTraderRefreshTime_In struct {
	TraderId int16 `json:"trader_id"`
	Timing   int64 `json:"timing"`
}

func (this *AddTraderRefreshTime_In) Process(session *net.Session) {
	g_InHandler.AddTraderRefreshTime(session, this)
}

func (this *AddTraderRefreshTime_In) TypeName() string {
	return "debug.add_trader_refresh_time.in"
}

func (this *AddTraderRefreshTime_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 42
}

type AddTraderRefreshTime_Out struct {
}

func (this *AddTraderRefreshTime_Out) Process(session *net.Session) {
	g_OutHandler.AddTraderRefreshTime(session, this)
}

func (this *AddTraderRefreshTime_Out) TypeName() string {
	return "debug.add_trader_refresh_time.out"
}

func (this *AddTraderRefreshTime_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 42
}

func (this *AddTraderRefreshTime_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ClearTraderSchedule_In struct {
	TraderId int16 `json:"trader_id"`
}

func (this *ClearTraderSchedule_In) Process(session *net.Session) {
	g_InHandler.ClearTraderSchedule(session, this)
}

func (this *ClearTraderSchedule_In) TypeName() string {
	return "debug.clear_trader_schedule.in"
}

func (this *ClearTraderSchedule_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 43
}

type ClearTraderSchedule_Out struct {
}

func (this *ClearTraderSchedule_Out) Process(session *net.Session) {
	g_OutHandler.ClearTraderSchedule(session, this)
}

func (this *ClearTraderSchedule_Out) TypeName() string {
	return "debug.clear_trader_schedule.out"
}

func (this *ClearTraderSchedule_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 43
}

func (this *ClearTraderSchedule_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AddTraderSchedule_In struct {
	TraderId  int16 `json:"trader_id"`
	Expire    int64 `json:"expire"`
	Showup    int64 `json:"showup"`
	Disappear int64 `json:"disappear"`
}

func (this *AddTraderSchedule_In) Process(session *net.Session) {
	g_InHandler.AddTraderSchedule(session, this)
}

func (this *AddTraderSchedule_In) TypeName() string {
	return "debug.add_trader_schedule.in"
}

func (this *AddTraderSchedule_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 44
}

type AddTraderSchedule_Out struct {
}

func (this *AddTraderSchedule_Out) Process(session *net.Session) {
	g_OutHandler.AddTraderSchedule(session, this)
}

func (this *AddTraderSchedule_Out) TypeName() string {
	return "debug.add_trader_schedule.out"
}

func (this *AddTraderSchedule_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 44
}

func (this *AddTraderSchedule_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type OpenTown_In struct {
	TownId int16 `json:"town_id"`
}

func (this *OpenTown_In) Process(session *net.Session) {
	g_InHandler.OpenTown(session, this)
}

func (this *OpenTown_In) TypeName() string {
	return "debug.open_town.in"
}

func (this *OpenTown_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 45
}

type OpenTown_Out struct {
}

func (this *OpenTown_Out) Process(session *net.Session) {
	g_OutHandler.OpenTown(session, this)
}

func (this *OpenTown_Out) TypeName() string {
	return "debug.open_town.out"
}

func (this *OpenTown_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 45
}

func (this *OpenTown_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AddGlobalMail_In struct {
	SendDelay   int64 `json:"send_delay"`
	ExpireDelay int64 `json:"expire_delay"`
}

func (this *AddGlobalMail_In) Process(session *net.Session) {
	g_InHandler.AddGlobalMail(session, this)
}

func (this *AddGlobalMail_In) TypeName() string {
	return "debug.add_global_mail.in"
}

func (this *AddGlobalMail_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 46
}

type AddGlobalMail_Out struct {
}

func (this *AddGlobalMail_Out) Process(session *net.Session) {
	g_OutHandler.AddGlobalMail(session, this)
}

func (this *AddGlobalMail_Out) TypeName() string {
	return "debug.add_global_mail.out"
}

func (this *AddGlobalMail_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 46
}

func (this *AddGlobalMail_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CreateAnnouncementWithoutTpl_In struct {
}

func (this *CreateAnnouncementWithoutTpl_In) Process(session *net.Session) {
	g_InHandler.CreateAnnouncementWithoutTpl(session, this)
}

func (this *CreateAnnouncementWithoutTpl_In) TypeName() string {
	return "debug.create_announcement_without_tpl.in"
}

func (this *CreateAnnouncementWithoutTpl_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 47
}

type CreateAnnouncementWithoutTpl_Out struct {
}

func (this *CreateAnnouncementWithoutTpl_Out) Process(session *net.Session) {
	g_OutHandler.CreateAnnouncementWithoutTpl(session, this)
}

func (this *CreateAnnouncementWithoutTpl_Out) TypeName() string {
	return "debug.create_announcement_without_tpl.out"
}

func (this *CreateAnnouncementWithoutTpl_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 47
}

func (this *CreateAnnouncementWithoutTpl_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SetLoginDay_In struct {
	Days int32 `json:"days"`
}

func (this *SetLoginDay_In) Process(session *net.Session) {
	g_InHandler.SetLoginDay(session, this)
}

func (this *SetLoginDay_In) TypeName() string {
	return "debug.set_login_day.in"
}

func (this *SetLoginDay_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 48
}

type SetLoginDay_Out struct {
}

func (this *SetLoginDay_Out) Process(session *net.Session) {
	g_OutHandler.SetLoginDay(session, this)
}

func (this *SetLoginDay_Out) TypeName() string {
	return "debug.set_login_day.out"
}

func (this *SetLoginDay_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 48
}

func (this *SetLoginDay_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ResetLoginAward_In struct {
}

func (this *ResetLoginAward_In) Process(session *net.Session) {
	g_InHandler.ResetLoginAward(session, this)
}

func (this *ResetLoginAward_In) TypeName() string {
	return "debug.reset_login_award.in"
}

func (this *ResetLoginAward_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 49
}

type ResetLoginAward_Out struct {
}

func (this *ResetLoginAward_Out) Process(session *net.Session) {
	g_OutHandler.ResetLoginAward(session, this)
}

func (this *ResetLoginAward_Out) TypeName() string {
	return "debug.reset_login_award.out"
}

func (this *ResetLoginAward_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 49
}

func (this *ResetLoginAward_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type RestPlayerAwardLock_In struct {
}

func (this *RestPlayerAwardLock_In) Process(session *net.Session) {
	g_InHandler.RestPlayerAwardLock(session, this)
}

func (this *RestPlayerAwardLock_In) TypeName() string {
	return "debug.rest_player_award_lock.in"
}

func (this *RestPlayerAwardLock_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 50
}

type RestPlayerAwardLock_Out struct {
}

func (this *RestPlayerAwardLock_Out) Process(session *net.Session) {
	g_OutHandler.RestPlayerAwardLock(session, this)
}

func (this *RestPlayerAwardLock_Out) TypeName() string {
	return "debug.rest_player_award_lock.out"
}

func (this *RestPlayerAwardLock_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 50
}

func (this *RestPlayerAwardLock_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ResetRainbowLevel_In struct {
}

func (this *ResetRainbowLevel_In) Process(session *net.Session) {
	g_InHandler.ResetRainbowLevel(session, this)
}

func (this *ResetRainbowLevel_In) TypeName() string {
	return "debug.reset_rainbow_level.in"
}

func (this *ResetRainbowLevel_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 51
}

type ResetRainbowLevel_Out struct {
}

func (this *ResetRainbowLevel_Out) Process(session *net.Session) {
	g_OutHandler.ResetRainbowLevel(session, this)
}

func (this *ResetRainbowLevel_Out) TypeName() string {
	return "debug.reset_rainbow_level.out"
}

func (this *ResetRainbowLevel_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 51
}

func (this *ResetRainbowLevel_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SetRainbowLevel_In struct {
	Segment int16 `json:"segment"`
	Order   int8  `json:"order"`
}

func (this *SetRainbowLevel_In) Process(session *net.Session) {
	g_InHandler.SetRainbowLevel(session, this)
}

func (this *SetRainbowLevel_In) TypeName() string {
	return "debug.set_rainbow_level.in"
}

func (this *SetRainbowLevel_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 52
}

type SetRainbowLevel_Out struct {
}

func (this *SetRainbowLevel_Out) Process(session *net.Session) {
	g_OutHandler.SetRainbowLevel(session, this)
}

func (this *SetRainbowLevel_Out) TypeName() string {
	return "debug.set_rainbow_level.out"
}

func (this *SetRainbowLevel_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 52
}

func (this *SetRainbowLevel_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SendPushNotification_In struct {
}

func (this *SendPushNotification_In) Process(session *net.Session) {
	g_InHandler.SendPushNotification(session, this)
}

func (this *SendPushNotification_In) TypeName() string {
	return "debug.send_push_notification.in"
}

func (this *SendPushNotification_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 53
}

type SendPushNotification_Out struct {
}

func (this *SendPushNotification_Out) Process(session *net.Session) {
	g_OutHandler.SendPushNotification(session, this)
}

func (this *SendPushNotification_Out) TypeName() string {
	return "debug.send_push_notification.out"
}

func (this *SendPushNotification_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 53
}

func (this *SendPushNotification_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ResetPetVirtualEnv_In struct {
}

func (this *ResetPetVirtualEnv_In) Process(session *net.Session) {
	g_InHandler.ResetPetVirtualEnv(session, this)
}

func (this *ResetPetVirtualEnv_In) TypeName() string {
	return "debug.reset_pet_virtual_env.in"
}

func (this *ResetPetVirtualEnv_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 54
}

type ResetPetVirtualEnv_Out struct {
}

func (this *ResetPetVirtualEnv_Out) Process(session *net.Session) {
	g_OutHandler.ResetPetVirtualEnv(session, this)
}

func (this *ResetPetVirtualEnv_Out) TypeName() string {
	return "debug.reset_pet_virtual_env.out"
}

func (this *ResetPetVirtualEnv_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 54
}

func (this *ResetPetVirtualEnv_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AddFame_In struct {
	System int16 `json:"system"`
	Val    int32 `json:"val"`
}

func (this *AddFame_In) Process(session *net.Session) {
	g_InHandler.AddFame(session, this)
}

func (this *AddFame_In) TypeName() string {
	return "debug.add_fame.in"
}

func (this *AddFame_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 55
}

type AddFame_Out struct {
}

func (this *AddFame_Out) Process(session *net.Session) {
	g_OutHandler.AddFame(session, this)
}

func (this *AddFame_Out) TypeName() string {
	return "debug.add_fame.out"
}

func (this *AddFame_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 55
}

func (this *AddFame_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AddWorldChatMessage_In struct {
	Num int16 `json:"num"`
}

func (this *AddWorldChatMessage_In) Process(session *net.Session) {
	g_InHandler.AddWorldChatMessage(session, this)
}

func (this *AddWorldChatMessage_In) TypeName() string {
	return "debug.add_world_chat_message.in"
}

func (this *AddWorldChatMessage_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 56
}

type AddWorldChatMessage_Out struct {
}

func (this *AddWorldChatMessage_Out) Process(session *net.Session) {
	g_OutHandler.AddWorldChatMessage(session, this)
}

func (this *AddWorldChatMessage_Out) TypeName() string {
	return "debug.add_world_chat_message.out"
}

func (this *AddWorldChatMessage_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 56
}

func (this *AddWorldChatMessage_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type MonthCard_In struct {
}

func (this *MonthCard_In) Process(session *net.Session) {
	g_InHandler.MonthCard(session, this)
}

func (this *MonthCard_In) TypeName() string {
	return "debug.month_card.in"
}

func (this *MonthCard_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 57
}

type MonthCard_Out struct {
}

func (this *MonthCard_Out) Process(session *net.Session) {
	g_OutHandler.MonthCard(session, this)
}

func (this *MonthCard_Out) TypeName() string {
	return "debug.month_card.out"
}

func (this *MonthCard_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 57
}

func (this *MonthCard_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type EnterSandbox_In struct {
}

func (this *EnterSandbox_In) Process(session *net.Session) {
	g_InHandler.EnterSandbox(session, this)
}

func (this *EnterSandbox_In) TypeName() string {
	return "debug.enter_sandbox.in"
}

func (this *EnterSandbox_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 58
}

type EnterSandbox_Out struct {
}

func (this *EnterSandbox_Out) Process(session *net.Session) {
	g_OutHandler.EnterSandbox(session, this)
}

func (this *EnterSandbox_Out) TypeName() string {
	return "debug.enter_sandbox.out"
}

func (this *EnterSandbox_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 58
}

func (this *EnterSandbox_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SandboxRollback_In struct {
}

func (this *SandboxRollback_In) Process(session *net.Session) {
	g_InHandler.SandboxRollback(session, this)
}

func (this *SandboxRollback_In) TypeName() string {
	return "debug.sandbox_rollback.in"
}

func (this *SandboxRollback_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 59
}

type SandboxRollback_Out struct {
}

func (this *SandboxRollback_Out) Process(session *net.Session) {
	g_OutHandler.SandboxRollback(session, this)
}

func (this *SandboxRollback_Out) TypeName() string {
	return "debug.sandbox_rollback.out"
}

func (this *SandboxRollback_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 59
}

func (this *SandboxRollback_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ExitSandbox_In struct {
}

func (this *ExitSandbox_In) Process(session *net.Session) {
	g_InHandler.ExitSandbox(session, this)
}

func (this *ExitSandbox_In) TypeName() string {
	return "debug.exit_sandbox.in"
}

func (this *ExitSandbox_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 60
}

type ExitSandbox_Out struct {
}

func (this *ExitSandbox_Out) Process(session *net.Session) {
	g_OutHandler.ExitSandbox(session, this)
}

func (this *ExitSandbox_Out) TypeName() string {
	return "debug.exit_sandbox.out"
}

func (this *ExitSandbox_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 60
}

func (this *ExitSandbox_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ResetShadedMissions_In struct {
	LevelId int32 `json:"level_id"`
}

func (this *ResetShadedMissions_In) Process(session *net.Session) {
	g_InHandler.ResetShadedMissions(session, this)
}

func (this *ResetShadedMissions_In) TypeName() string {
	return "debug.reset_shaded_missions.in"
}

func (this *ResetShadedMissions_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 61
}

type ResetShadedMissions_Out struct {
}

func (this *ResetShadedMissions_Out) Process(session *net.Session) {
	g_OutHandler.ResetShadedMissions(session, this)
}

func (this *ResetShadedMissions_Out) TypeName() string {
	return "debug.reset_shaded_missions.out"
}

func (this *ResetShadedMissions_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 61
}

func (this *ResetShadedMissions_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CleanCornucopia_In struct {
}

func (this *CleanCornucopia_In) Process(session *net.Session) {
	g_InHandler.CleanCornucopia(session, this)
}

func (this *CleanCornucopia_In) TypeName() string {
	return "debug.clean_cornucopia.in"
}

func (this *CleanCornucopia_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 62
}

type CleanCornucopia_Out struct {
}

func (this *CleanCornucopia_Out) Process(session *net.Session) {
	g_OutHandler.CleanCornucopia(session, this)
}

func (this *CleanCornucopia_Out) TypeName() string {
	return "debug.clean_cornucopia.out"
}

func (this *CleanCornucopia_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 62
}

func (this *CleanCornucopia_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AddTotem_In struct {
	TotemId int16 `json:"totem_id"`
}

func (this *AddTotem_In) Process(session *net.Session) {
	g_InHandler.AddTotem(session, this)
}

func (this *AddTotem_In) TypeName() string {
	return "debug.add_totem.in"
}

func (this *AddTotem_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 63
}

type AddTotem_Out struct {
}

func (this *AddTotem_Out) Process(session *net.Session) {
	g_OutHandler.AddTotem(session, this)
}

func (this *AddTotem_Out) TypeName() string {
	return "debug.add_totem.out"
}

func (this *AddTotem_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 63
}

func (this *AddTotem_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AddRune_In struct {
	JadeNum int32 `json:"jade_num"`
	RockNum int32 `json:"rock_num"`
}

func (this *AddRune_In) Process(session *net.Session) {
	g_InHandler.AddRune(session, this)
}

func (this *AddRune_In) TypeName() string {
	return "debug.add_rune.in"
}

func (this *AddRune_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 64
}

type AddRune_Out struct {
}

func (this *AddRune_Out) Process(session *net.Session) {
	g_OutHandler.AddRune(session, this)
}

func (this *AddRune_Out) TypeName() string {
	return "debug.add_rune.out"
}

func (this *AddRune_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 64
}

func (this *AddRune_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SendRareItemMessage_In struct {
}

func (this *SendRareItemMessage_In) Process(session *net.Session) {
	g_InHandler.SendRareItemMessage(session, this)
}

func (this *SendRareItemMessage_In) TypeName() string {
	return "debug.send_rare_item_message.in"
}

func (this *SendRareItemMessage_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 65
}

type SendRareItemMessage_Out struct {
}

func (this *SendRareItemMessage_Out) Process(session *net.Session) {
	g_OutHandler.SendRareItemMessage(session, this)
}

func (this *SendRareItemMessage_Out) TypeName() string {
	return "debug.send_rare_item_message.out"
}

func (this *SendRareItemMessage_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 65
}

func (this *SendRareItemMessage_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AddSwordDrivingAction_In struct {
	Point int16 `json:"point"`
}

func (this *AddSwordDrivingAction_In) Process(session *net.Session) {
	g_InHandler.AddSwordDrivingAction(session, this)
}

func (this *AddSwordDrivingAction_In) TypeName() string {
	return "debug.add_sword_driving_action.in"
}

func (this *AddSwordDrivingAction_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 66
}

type AddSwordDrivingAction_Out struct {
}

func (this *AddSwordDrivingAction_Out) Process(session *net.Session) {
	g_OutHandler.AddSwordDrivingAction(session, this)
}

func (this *AddSwordDrivingAction_Out) TypeName() string {
	return "debug.add_sword_driving_action.out"
}

func (this *AddSwordDrivingAction_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 66
}

func (this *AddSwordDrivingAction_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ResetDrivingSwordData_In struct {
	Cloud int16 `json:"cloud"`
}

func (this *ResetDrivingSwordData_In) Process(session *net.Session) {
	g_InHandler.ResetDrivingSwordData(session, this)
}

func (this *ResetDrivingSwordData_In) TypeName() string {
	return "debug.reset_driving_sword_data.in"
}

func (this *ResetDrivingSwordData_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 67
}

type ResetDrivingSwordData_Out struct {
}

func (this *ResetDrivingSwordData_Out) Process(session *net.Session) {
	g_OutHandler.ResetDrivingSwordData(session, this)
}

func (this *ResetDrivingSwordData_Out) TypeName() string {
	return "debug.reset_driving_sword_data.out"
}

func (this *ResetDrivingSwordData_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 67
}

func (this *ResetDrivingSwordData_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AddSwordSoulFragment_In struct {
	Number int64 `json:"number"`
}

func (this *AddSwordSoulFragment_In) Process(session *net.Session) {
	g_InHandler.AddSwordSoulFragment(session, this)
}

func (this *AddSwordSoulFragment_In) TypeName() string {
	return "debug.add_sword_soul_fragment.in"
}

func (this *AddSwordSoulFragment_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 68
}

type AddSwordSoulFragment_Out struct {
}

func (this *AddSwordSoulFragment_Out) Process(session *net.Session) {
	g_OutHandler.AddSwordSoulFragment(session, this)
}

func (this *AddSwordSoulFragment_Out) TypeName() string {
	return "debug.add_sword_soul_fragment.out"
}

func (this *AddSwordSoulFragment_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 68
}

func (this *AddSwordSoulFragment_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ResetMoneyTreeStatus_In struct {
}

func (this *ResetMoneyTreeStatus_In) Process(session *net.Session) {
	g_InHandler.ResetMoneyTreeStatus(session, this)
}

func (this *ResetMoneyTreeStatus_In) TypeName() string {
	return "debug.reset_money_tree_status.in"
}

func (this *ResetMoneyTreeStatus_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 69
}

type ResetMoneyTreeStatus_Out struct {
}

func (this *ResetMoneyTreeStatus_Out) Process(session *net.Session) {
	g_OutHandler.ResetMoneyTreeStatus(session, this)
}

func (this *ResetMoneyTreeStatus_Out) TypeName() string {
	return "debug.reset_money_tree_status.out"
}

func (this *ResetMoneyTreeStatus_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 69
}

func (this *ResetMoneyTreeStatus_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ResetTodayMoneyTree_In struct {
}

func (this *ResetTodayMoneyTree_In) Process(session *net.Session) {
	g_InHandler.ResetTodayMoneyTree(session, this)
}

func (this *ResetTodayMoneyTree_In) TypeName() string {
	return "debug.reset_today_money_tree.in"
}

func (this *ResetTodayMoneyTree_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 70
}

type ResetTodayMoneyTree_Out struct {
}

func (this *ResetTodayMoneyTree_Out) Process(session *net.Session) {
	g_OutHandler.ResetTodayMoneyTree(session, this)
}

func (this *ResetTodayMoneyTree_Out) TypeName() string {
	return "debug.reset_today_money_tree.out"
}

func (this *ResetTodayMoneyTree_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 70
}

func (this *ResetTodayMoneyTree_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CleanSwordSoulIngotDrawNums_In struct {
}

func (this *CleanSwordSoulIngotDrawNums_In) Process(session *net.Session) {
	g_InHandler.CleanSwordSoulIngotDrawNums(session, this)
}

func (this *CleanSwordSoulIngotDrawNums_In) TypeName() string {
	return "debug.clean_sword_soul_ingot_draw_nums.in"
}

func (this *CleanSwordSoulIngotDrawNums_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 71
}

type CleanSwordSoulIngotDrawNums_Out struct {
}

func (this *CleanSwordSoulIngotDrawNums_Out) Process(session *net.Session) {
	g_OutHandler.CleanSwordSoulIngotDrawNums(session, this)
}

func (this *CleanSwordSoulIngotDrawNums_Out) TypeName() string {
	return "debug.clean_sword_soul_ingot_draw_nums.out"
}

func (this *CleanSwordSoulIngotDrawNums_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 71
}

func (this *CleanSwordSoulIngotDrawNums_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type PunchDrivingSwordCloud_In struct {
}

func (this *PunchDrivingSwordCloud_In) Process(session *net.Session) {
	g_InHandler.PunchDrivingSwordCloud(session, this)
}

func (this *PunchDrivingSwordCloud_In) TypeName() string {
	return "debug.punch_driving_sword_cloud.in"
}

func (this *PunchDrivingSwordCloud_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 72
}

type PunchDrivingSwordCloud_Out struct {
}

func (this *PunchDrivingSwordCloud_Out) Process(session *net.Session) {
	g_OutHandler.PunchDrivingSwordCloud(session, this)
}

func (this *PunchDrivingSwordCloud_Out) TypeName() string {
	return "debug.punch_driving_sword_cloud.out"
}

func (this *PunchDrivingSwordCloud_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 72
}

func (this *PunchDrivingSwordCloud_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ClearCliqueDailyDonate_In struct {
}

func (this *ClearCliqueDailyDonate_In) Process(session *net.Session) {
	g_InHandler.ClearCliqueDailyDonate(session, this)
}

func (this *ClearCliqueDailyDonate_In) TypeName() string {
	return "debug.clear_clique_daily_donate.in"
}

func (this *ClearCliqueDailyDonate_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 73
}

type ClearCliqueDailyDonate_Out struct {
}

func (this *ClearCliqueDailyDonate_Out) Process(session *net.Session) {
	g_OutHandler.ClearCliqueDailyDonate(session, this)
}

func (this *ClearCliqueDailyDonate_Out) TypeName() string {
	return "debug.clear_clique_daily_donate.out"
}

func (this *ClearCliqueDailyDonate_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 73
}

func (this *ClearCliqueDailyDonate_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SetCliqueContrib_In struct {
	Contrib int64 `json:"contrib"`
}

func (this *SetCliqueContrib_In) Process(session *net.Session) {
	g_InHandler.SetCliqueContrib(session, this)
}

func (this *SetCliqueContrib_In) TypeName() string {
	return "debug.set_clique_contrib.in"
}

func (this *SetCliqueContrib_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 74
}

type SetCliqueContrib_Out struct {
}

func (this *SetCliqueContrib_Out) Process(session *net.Session) {
	g_OutHandler.SetCliqueContrib(session, this)
}

func (this *SetCliqueContrib_Out) TypeName() string {
	return "debug.set_clique_contrib.out"
}

func (this *SetCliqueContrib_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 74
}

func (this *SetCliqueContrib_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type RefreshCliqueWorship_In struct {
}

func (this *RefreshCliqueWorship_In) Process(session *net.Session) {
	g_InHandler.RefreshCliqueWorship(session, this)
}

func (this *RefreshCliqueWorship_In) TypeName() string {
	return "debug.refresh_clique_worship.in"
}

func (this *RefreshCliqueWorship_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 75
}

type RefreshCliqueWorship_Out struct {
}

func (this *RefreshCliqueWorship_Out) Process(session *net.Session) {
	g_OutHandler.RefreshCliqueWorship(session, this)
}

func (this *RefreshCliqueWorship_Out) TypeName() string {
	return "debug.refresh_clique_worship.out"
}

func (this *RefreshCliqueWorship_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 75
}

func (this *RefreshCliqueWorship_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueEscortHijackBattleWin_In struct {
	BoatId int64 `json:"boat_id"`
}

func (this *CliqueEscortHijackBattleWin_In) Process(session *net.Session) {
	g_InHandler.CliqueEscortHijackBattleWin(session, this)
}

func (this *CliqueEscortHijackBattleWin_In) TypeName() string {
	return "debug.clique_escort_hijack_battle_win.in"
}

func (this *CliqueEscortHijackBattleWin_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 76
}

type CliqueEscortHijackBattleWin_Out struct {
}

func (this *CliqueEscortHijackBattleWin_Out) Process(session *net.Session) {
	g_OutHandler.CliqueEscortHijackBattleWin(session, this)
}

func (this *CliqueEscortHijackBattleWin_Out) TypeName() string {
	return "debug.clique_escort_hijack_battle_win.out"
}

func (this *CliqueEscortHijackBattleWin_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 76
}

func (this *CliqueEscortHijackBattleWin_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueEscortRecoverBattleWin_In struct {
	BoatId int64 `json:"boat_id"`
}

func (this *CliqueEscortRecoverBattleWin_In) Process(session *net.Session) {
	g_InHandler.CliqueEscortRecoverBattleWin(session, this)
}

func (this *CliqueEscortRecoverBattleWin_In) TypeName() string {
	return "debug.clique_escort_recover_battle_win.in"
}

func (this *CliqueEscortRecoverBattleWin_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 77
}

type CliqueEscortRecoverBattleWin_Out struct {
}

func (this *CliqueEscortRecoverBattleWin_Out) Process(session *net.Session) {
	g_OutHandler.CliqueEscortRecoverBattleWin(session, this)
}

func (this *CliqueEscortRecoverBattleWin_Out) TypeName() string {
	return "debug.clique_escort_recover_battle_win.out"
}

func (this *CliqueEscortRecoverBattleWin_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 77
}

func (this *CliqueEscortRecoverBattleWin_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueEscortNotifyMessage_In struct {
}

func (this *CliqueEscortNotifyMessage_In) Process(session *net.Session) {
	g_InHandler.CliqueEscortNotifyMessage(session, this)
}

func (this *CliqueEscortNotifyMessage_In) TypeName() string {
	return "debug.clique_escort_notify_message.in"
}

func (this *CliqueEscortNotifyMessage_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 87
}

type CliqueEscortNotifyMessage_Out struct {
}

func (this *CliqueEscortNotifyMessage_Out) Process(session *net.Session) {
	g_OutHandler.CliqueEscortNotifyMessage(session, this)
}

func (this *CliqueEscortNotifyMessage_Out) TypeName() string {
	return "debug.clique_escort_notify_message.out"
}

func (this *CliqueEscortNotifyMessage_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 87
}

func (this *CliqueEscortNotifyMessage_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueEscortNotifyDailyQuest_In struct {
}

func (this *CliqueEscortNotifyDailyQuest_In) Process(session *net.Session) {
	g_InHandler.CliqueEscortNotifyDailyQuest(session, this)
}

func (this *CliqueEscortNotifyDailyQuest_In) TypeName() string {
	return "debug.clique_escort_notify_daily_quest.in"
}

func (this *CliqueEscortNotifyDailyQuest_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 88
}

type CliqueEscortNotifyDailyQuest_Out struct {
}

func (this *CliqueEscortNotifyDailyQuest_Out) Process(session *net.Session) {
	g_OutHandler.CliqueEscortNotifyDailyQuest(session, this)
}

func (this *CliqueEscortNotifyDailyQuest_Out) TypeName() string {
	return "debug.clique_escort_notify_daily_quest.out"
}

func (this *CliqueEscortNotifyDailyQuest_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 88
}

func (this *CliqueEscortNotifyDailyQuest_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SetCliqueBuildingLevel_In struct {
	Building int32 `json:"building"`
	Level    int16 `json:"level"`
}

func (this *SetCliqueBuildingLevel_In) Process(session *net.Session) {
	g_InHandler.SetCliqueBuildingLevel(session, this)
}

func (this *SetCliqueBuildingLevel_In) TypeName() string {
	return "debug.set_clique_building_level.in"
}

func (this *SetCliqueBuildingLevel_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 89
}

type SetCliqueBuildingLevel_Out struct {
}

func (this *SetCliqueBuildingLevel_Out) Process(session *net.Session) {
	g_OutHandler.SetCliqueBuildingLevel(session, this)
}

func (this *SetCliqueBuildingLevel_Out) TypeName() string {
	return "debug.set_clique_building_level.out"
}

func (this *SetCliqueBuildingLevel_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 89
}

func (this *SetCliqueBuildingLevel_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SetCliqueBuildingMoney_In struct {
	Building int32 `json:"building"`
	Money    int64 `json:"money"`
}

func (this *SetCliqueBuildingMoney_In) Process(session *net.Session) {
	g_InHandler.SetCliqueBuildingMoney(session, this)
}

func (this *SetCliqueBuildingMoney_In) TypeName() string {
	return "debug.set_clique_building_money.in"
}

func (this *SetCliqueBuildingMoney_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 90
}

type SetCliqueBuildingMoney_Out struct {
}

func (this *SetCliqueBuildingMoney_Out) Process(session *net.Session) {
	g_OutHandler.SetCliqueBuildingMoney(session, this)
}

func (this *SetCliqueBuildingMoney_Out) TypeName() string {
	return "debug.set_clique_building_money.out"
}

func (this *SetCliqueBuildingMoney_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 90
}

func (this *SetCliqueBuildingMoney_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type EscortBench_In struct {
}

func (this *EscortBench_In) Process(session *net.Session) {
	g_InHandler.EscortBench(session, this)
}

func (this *EscortBench_In) TypeName() string {
	return "debug.escort_bench.in"
}

func (this *EscortBench_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 91
}

type EscortBench_Out struct {
}

func (this *EscortBench_Out) Process(session *net.Session) {
	g_OutHandler.EscortBench(session, this)
}

func (this *EscortBench_Out) TypeName() string {
	return "debug.escort_bench.out"
}

func (this *EscortBench_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 91
}

func (this *EscortBench_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ResetCliqueEscortDailyNum_In struct {
}

func (this *ResetCliqueEscortDailyNum_In) Process(session *net.Session) {
	g_InHandler.ResetCliqueEscortDailyNum(session, this)
}

func (this *ResetCliqueEscortDailyNum_In) TypeName() string {
	return "debug.reset_clique_escort_daily_num.in"
}

func (this *ResetCliqueEscortDailyNum_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 92
}

type ResetCliqueEscortDailyNum_Out struct {
}

func (this *ResetCliqueEscortDailyNum_Out) Process(session *net.Session) {
	g_OutHandler.ResetCliqueEscortDailyNum(session, this)
}

func (this *ResetCliqueEscortDailyNum_Out) TypeName() string {
	return "debug.reset_clique_escort_daily_num.out"
}

func (this *ResetCliqueEscortDailyNum_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 92
}

func (this *ResetCliqueEscortDailyNum_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TakeAdditionQuest_In struct {
	FirstQuestId int32 `json:"first_quest_id"`
}

func (this *TakeAdditionQuest_In) Process(session *net.Session) {
	g_InHandler.TakeAdditionQuest(session, this)
}

func (this *TakeAdditionQuest_In) TypeName() string {
	return "debug.take_addition_quest.in"
}

func (this *TakeAdditionQuest_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 93
}

type TakeAdditionQuest_Out struct {
	Msg []byte `json:"msg"`
}

func (this *TakeAdditionQuest_Out) Process(session *net.Session) {
	g_OutHandler.TakeAdditionQuest(session, this)
}

func (this *TakeAdditionQuest_Out) TypeName() string {
	return "debug.take_addition_quest.out"
}

func (this *TakeAdditionQuest_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 93
}

func (this *TakeAdditionQuest_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SetMissionStarMax_In struct {
}

func (this *SetMissionStarMax_In) Process(session *net.Session) {
	g_InHandler.SetMissionStarMax(session, this)
}

func (this *SetMissionStarMax_In) TypeName() string {
	return "debug.set_mission_star_max.in"
}

func (this *SetMissionStarMax_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 94
}

type SetMissionStarMax_Out struct {
}

func (this *SetMissionStarMax_Out) Process(session *net.Session) {
	g_OutHandler.SetMissionStarMax(session, this)
}

func (this *SetMissionStarMax_Out) TypeName() string {
	return "debug.set_mission_star_max.out"
}

func (this *SetMissionStarMax_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 94
}

func (this *SetMissionStarMax_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueBankCd_In struct {
}

func (this *CliqueBankCd_In) Process(session *net.Session) {
	g_InHandler.CliqueBankCd(session, this)
}

func (this *CliqueBankCd_In) TypeName() string {
	return "debug.clique_bank_cd.in"
}

func (this *CliqueBankCd_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 95
}

type CliqueBankCd_Out struct {
}

func (this *CliqueBankCd_Out) Process(session *net.Session) {
	g_OutHandler.CliqueBankCd(session, this)
}

func (this *CliqueBankCd_Out) TypeName() string {
	return "debug.clique_bank_cd.out"
}

func (this *CliqueBankCd_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 95
}

func (this *CliqueBankCd_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ResetDespairLandBattleNum_In struct {
}

func (this *ResetDespairLandBattleNum_In) Process(session *net.Session) {
	g_InHandler.ResetDespairLandBattleNum(session, this)
}

func (this *ResetDespairLandBattleNum_In) TypeName() string {
	return "debug.reset_despair_land_battle_num.in"
}

func (this *ResetDespairLandBattleNum_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 96
}

type ResetDespairLandBattleNum_Out struct {
}

func (this *ResetDespairLandBattleNum_Out) Process(session *net.Session) {
	g_OutHandler.ResetDespairLandBattleNum(session, this)
}

func (this *ResetDespairLandBattleNum_Out) TypeName() string {
	return "debug.reset_despair_land_battle_num.out"
}

func (this *ResetDespairLandBattleNum_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 96
}

func (this *ResetDespairLandBattleNum_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ResetCliqueStoreSendTimes_In struct {
}

func (this *ResetCliqueStoreSendTimes_In) Process(session *net.Session) {
	g_InHandler.ResetCliqueStoreSendTimes(session, this)
}

func (this *ResetCliqueStoreSendTimes_In) TypeName() string {
	return "debug.reset_clique_store_send_times.in"
}

func (this *ResetCliqueStoreSendTimes_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 97
}

type ResetCliqueStoreSendTimes_Out struct {
}

func (this *ResetCliqueStoreSendTimes_Out) Process(session *net.Session) {
	g_OutHandler.ResetCliqueStoreSendTimes(session, this)
}

func (this *ResetCliqueStoreSendTimes_Out) TypeName() string {
	return "debug.reset_clique_store_send_times.out"
}

func (this *ResetCliqueStoreSendTimes_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 97
}

func (this *ResetCliqueStoreSendTimes_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AddCliqueStoreDonate_In struct {
}

func (this *AddCliqueStoreDonate_In) Process(session *net.Session) {
	g_InHandler.AddCliqueStoreDonate(session, this)
}

func (this *AddCliqueStoreDonate_In) TypeName() string {
	return "debug.add_clique_store_donate.in"
}

func (this *AddCliqueStoreDonate_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 98
}

type AddCliqueStoreDonate_Out struct {
}

func (this *AddCliqueStoreDonate_Out) Process(session *net.Session) {
	g_OutHandler.AddCliqueStoreDonate(session, this)
}

func (this *AddCliqueStoreDonate_Out) TypeName() string {
	return "debug.add_clique_store_donate.out"
}

func (this *AddCliqueStoreDonate_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 98
}

func (this *AddCliqueStoreDonate_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type PassAllDespairLandLevel_In struct {
	Star     int8 `json:"star"`
	CampType int8 `json:"camp_type"`
}

func (this *PassAllDespairLandLevel_In) Process(session *net.Session) {
	g_InHandler.PassAllDespairLandLevel(session, this)
}

func (this *PassAllDespairLandLevel_In) TypeName() string {
	return "debug.pass_all_despair_land_level.in"
}

func (this *PassAllDespairLandLevel_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 99
}

type PassAllDespairLandLevel_Out struct {
}

func (this *PassAllDespairLandLevel_Out) Process(session *net.Session) {
	g_OutHandler.PassAllDespairLandLevel(session, this)
}

func (this *PassAllDespairLandLevel_Out) TypeName() string {
	return "debug.pass_all_despair_land_level.out"
}

func (this *PassAllDespairLandLevel_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 99
}

func (this *PassAllDespairLandLevel_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DespairLandDummyBossKill_In struct {
}

func (this *DespairLandDummyBossKill_In) Process(session *net.Session) {
	g_InHandler.DespairLandDummyBossKill(session, this)
}

func (this *DespairLandDummyBossKill_In) TypeName() string {
	return "debug.despair_land_dummy_boss_kill.in"
}

func (this *DespairLandDummyBossKill_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 100
}

type DespairLandDummyBossKill_Out struct {
}

func (this *DespairLandDummyBossKill_Out) Process(session *net.Session) {
	g_OutHandler.DespairLandDummyBossKill(session, this)
}

func (this *DespairLandDummyBossKill_Out) TypeName() string {
	return "debug.despair_land_dummy_boss_kill.out"
}

func (this *DespairLandDummyBossKill_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 100
}

func (this *DespairLandDummyBossKill_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AddTaoyuanItem_In struct {
	ItemId int16 `json:"item_id"`
	Number int16 `json:"number"`
}

func (this *AddTaoyuanItem_In) Process(session *net.Session) {
	g_InHandler.AddTaoyuanItem(session, this)
}

func (this *AddTaoyuanItem_In) TypeName() string {
	return "debug.add_taoyuan_item.in"
}

func (this *AddTaoyuanItem_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 101
}

type AddTaoyuanItem_Out struct {
}

func (this *AddTaoyuanItem_Out) Process(session *net.Session) {
	g_OutHandler.AddTaoyuanItem(session, this)
}

func (this *AddTaoyuanItem_Out) TypeName() string {
	return "debug.add_taoyuan_item.out"
}

func (this *AddTaoyuanItem_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 101
}

func (this *AddTaoyuanItem_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AddTaoyuanExp_In struct {
	AddExp int64 `json:"add_exp"`
}

func (this *AddTaoyuanExp_In) Process(session *net.Session) {
	g_InHandler.AddTaoyuanExp(session, this)
}

func (this *AddTaoyuanExp_In) TypeName() string {
	return "debug.add_taoyuan_exp.in"
}

func (this *AddTaoyuanExp_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 102
}

type AddTaoyuanExp_Out struct {
}

func (this *AddTaoyuanExp_Out) Process(session *net.Session) {
	g_OutHandler.AddTaoyuanExp(session, this)
}

func (this *AddTaoyuanExp_Out) TypeName() string {
	return "debug.add_taoyuan_exp.out"
}

func (this *AddTaoyuanExp_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 99, 102
}

func (this *AddTaoyuanExp_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *AddBuddy_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
}

func (this *AddBuddy_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(0)
	buffer.WriteUint8(uint8(this.RoleId))
}

func (this *AddBuddy_In) ByteSize() int {
	size := 3
	return size
}

func (this *AddBuddy_Out) Decode(buffer *net.Buffer) {
}

func (this *AddBuddy_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(0)
}

func (this *AddBuddy_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AddItem_In) Decode(buffer *net.Buffer) {
	this.ItemId = int16(buffer.ReadUint16LE())
	this.Number = int16(buffer.ReadUint16LE())
}

func (this *AddItem_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(2)
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint16LE(uint16(this.Number))
}

func (this *AddItem_In) ByteSize() int {
	size := 6
	return size
}

func (this *AddItem_Out) Decode(buffer *net.Buffer) {
}

func (this *AddItem_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(2)
}

func (this *AddItem_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SetRoleLevel_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.Level = int16(buffer.ReadUint16LE())
}

func (this *SetRoleLevel_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(3)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.Level))
}

func (this *SetRoleLevel_In) ByteSize() int {
	size := 5
	return size
}

func (this *SetRoleLevel_Out) Decode(buffer *net.Buffer) {
}

func (this *SetRoleLevel_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(3)
}

func (this *SetRoleLevel_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SetCoins_In) Decode(buffer *net.Buffer) {
	this.Number = int64(buffer.ReadUint64LE())
}

func (this *SetCoins_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(4)
	buffer.WriteUint64LE(uint64(this.Number))
}

func (this *SetCoins_In) ByteSize() int {
	size := 10
	return size
}

func (this *SetCoins_Out) Decode(buffer *net.Buffer) {
}

func (this *SetCoins_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(4)
}

func (this *SetCoins_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SetIngot_In) Decode(buffer *net.Buffer) {
	this.Number = int64(buffer.ReadUint64LE())
}

func (this *SetIngot_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(5)
	buffer.WriteUint64LE(uint64(this.Number))
}

func (this *SetIngot_In) ByteSize() int {
	size := 10
	return size
}

func (this *SetIngot_Out) Decode(buffer *net.Buffer) {
}

func (this *SetIngot_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(5)
}

func (this *SetIngot_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AddGhost_In) Decode(buffer *net.Buffer) {
	this.GhostId = int16(buffer.ReadUint16LE())
}

func (this *AddGhost_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(11)
	buffer.WriteUint16LE(uint16(this.GhostId))
}

func (this *AddGhost_In) ByteSize() int {
	size := 4
	return size
}

func (this *AddGhost_Out) Decode(buffer *net.Buffer) {
}

func (this *AddGhost_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(11)
}

func (this *AddGhost_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SetPlayerPhysical_In) Decode(buffer *net.Buffer) {
	this.Physical = int16(buffer.ReadUint16LE())
}

func (this *SetPlayerPhysical_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(12)
	buffer.WriteUint16LE(uint16(this.Physical))
}

func (this *SetPlayerPhysical_In) ByteSize() int {
	size := 4
	return size
}

func (this *SetPlayerPhysical_Out) Decode(buffer *net.Buffer) {
}

func (this *SetPlayerPhysical_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(12)
}

func (this *SetPlayerPhysical_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ResetLevelEnterCount_In) Decode(buffer *net.Buffer) {
	this.LevelId = int32(buffer.ReadUint32LE())
}

func (this *ResetLevelEnterCount_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(13)
	buffer.WriteUint32LE(uint32(this.LevelId))
}

func (this *ResetLevelEnterCount_In) ByteSize() int {
	size := 6
	return size
}

func (this *ResetLevelEnterCount_Out) Decode(buffer *net.Buffer) {
}

func (this *ResetLevelEnterCount_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(13)
}

func (this *ResetLevelEnterCount_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AddExp_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.AddExp = int64(buffer.ReadUint64LE())
}

func (this *AddExp_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(14)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint64LE(uint64(this.AddExp))
}

func (this *AddExp_In) ByteSize() int {
	size := 11
	return size
}

func (this *AddExp_Out) Decode(buffer *net.Buffer) {
}

func (this *AddExp_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(14)
}

func (this *AddExp_Out) ByteSize() int {
	size := 2
	return size
}

func (this *OpenGhostMission_In) Decode(buffer *net.Buffer) {
	this.MissionId = int16(buffer.ReadUint16LE())
}

func (this *OpenGhostMission_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(15)
	buffer.WriteUint16LE(uint16(this.MissionId))
}

func (this *OpenGhostMission_In) ByteSize() int {
	size := 4
	return size
}

func (this *OpenGhostMission_Out) Decode(buffer *net.Buffer) {
}

func (this *OpenGhostMission_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(15)
}

func (this *OpenGhostMission_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SendMail_In) Decode(buffer *net.Buffer) {
	this.MailId = int32(buffer.ReadUint32LE())
}

func (this *SendMail_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(16)
	buffer.WriteUint32LE(uint32(this.MailId))
}

func (this *SendMail_In) ByteSize() int {
	size := 6
	return size
}

func (this *SendMail_Out) Decode(buffer *net.Buffer) {
}

func (this *SendMail_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(16)
}

func (this *SendMail_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ClearMail_In) Decode(buffer *net.Buffer) {
}

func (this *ClearMail_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(17)
}

func (this *ClearMail_In) ByteSize() int {
	size := 2
	return size
}

func (this *ClearMail_Out) Decode(buffer *net.Buffer) {
}

func (this *ClearMail_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(17)
}

func (this *ClearMail_Out) ByteSize() int {
	size := 2
	return size
}

func (this *OpenMissionLevel_In) Decode(buffer *net.Buffer) {
	this.LevelId = int32(buffer.ReadUint32LE())
}

func (this *OpenMissionLevel_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(18)
	buffer.WriteUint32LE(uint32(this.LevelId))
}

func (this *OpenMissionLevel_In) ByteSize() int {
	size := 6
	return size
}

func (this *OpenMissionLevel_Out) Decode(buffer *net.Buffer) {
}

func (this *OpenMissionLevel_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(18)
}

func (this *OpenMissionLevel_Out) ByteSize() int {
	size := 2
	return size
}

func (this *StartBattle_In) Decode(buffer *net.Buffer) {
	this.BattleType = int8(buffer.ReadUint8())
	this.EnemyId = int32(buffer.ReadUint32LE())
}

func (this *StartBattle_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(19)
	buffer.WriteUint8(uint8(this.BattleType))
	buffer.WriteUint32LE(uint32(this.EnemyId))
}

func (this *StartBattle_In) ByteSize() int {
	size := 7
	return size
}

func (this *StartBattle_Out) Decode(buffer *net.Buffer) {
}

func (this *StartBattle_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(19)
}

func (this *StartBattle_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ListenByName_In) Decode(buffer *net.Buffer) {
	this.Name = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *ListenByName_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(20)
	buffer.WriteUint16LE(uint16(len(this.Name)))
	buffer.WriteBytes(this.Name)
}

func (this *ListenByName_In) ByteSize() int {
	size := 4
	size += len(this.Name)
	return size
}

func (this *ListenByName_Out) Decode(buffer *net.Buffer) {
}

func (this *ListenByName_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(20)
}

func (this *ListenByName_Out) ByteSize() int {
	size := 2
	return size
}

func (this *OpenQuest_In) Decode(buffer *net.Buffer) {
	this.QuestId = int16(buffer.ReadUint16LE())
}

func (this *OpenQuest_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(21)
	buffer.WriteUint16LE(uint16(this.QuestId))
}

func (this *OpenQuest_In) ByteSize() int {
	size := 4
	return size
}

func (this *OpenQuest_Out) Decode(buffer *net.Buffer) {
}

func (this *OpenQuest_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(21)
}

func (this *OpenQuest_Out) ByteSize() int {
	size := 2
	return size
}

func (this *OpenFunc_In) Decode(buffer *net.Buffer) {
	this.Lock = int16(buffer.ReadUint16LE())
}

func (this *OpenFunc_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(22)
	buffer.WriteUint16LE(uint16(this.Lock))
}

func (this *OpenFunc_In) ByteSize() int {
	size := 4
	return size
}

func (this *OpenFunc_Out) Decode(buffer *net.Buffer) {
}

func (this *OpenFunc_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(22)
}

func (this *OpenFunc_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AddSwordSoul_In) Decode(buffer *net.Buffer) {
	this.SwordSoulId = int16(buffer.ReadUint16LE())
}

func (this *AddSwordSoul_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(23)
	buffer.WriteUint16LE(uint16(this.SwordSoulId))
}

func (this *AddSwordSoul_In) ByteSize() int {
	size := 4
	return size
}

func (this *AddSwordSoul_Out) Decode(buffer *net.Buffer) {
}

func (this *AddSwordSoul_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(23)
}

func (this *AddSwordSoul_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AddBattlePet_In) Decode(buffer *net.Buffer) {
	this.PetId = int16(buffer.ReadUint16LE())
}

func (this *AddBattlePet_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(25)
	buffer.WriteUint16LE(uint16(this.PetId))
}

func (this *AddBattlePet_In) ByteSize() int {
	size := 4
	return size
}

func (this *AddBattlePet_Out) Decode(buffer *net.Buffer) {
}

func (this *AddBattlePet_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(25)
}

func (this *AddBattlePet_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ResetMultiLevelEnterCount_In) Decode(buffer *net.Buffer) {
}

func (this *ResetMultiLevelEnterCount_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(26)
}

func (this *ResetMultiLevelEnterCount_In) ByteSize() int {
	size := 2
	return size
}

func (this *ResetMultiLevelEnterCount_Out) Decode(buffer *net.Buffer) {
}

func (this *ResetMultiLevelEnterCount_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(26)
}

func (this *ResetMultiLevelEnterCount_Out) ByteSize() int {
	size := 2
	return size
}

func (this *OpenMultiLevel_In) Decode(buffer *net.Buffer) {
	this.LevelId = int16(buffer.ReadUint16LE())
}

func (this *OpenMultiLevel_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(27)
	buffer.WriteUint16LE(uint16(this.LevelId))
}

func (this *OpenMultiLevel_In) ByteSize() int {
	size := 4
	return size
}

func (this *OpenMultiLevel_Out) Decode(buffer *net.Buffer) {
}

func (this *OpenMultiLevel_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(27)
}

func (this *OpenMultiLevel_Out) ByteSize() int {
	size := 2
	return size
}

func (this *OpenAllPetGrid_In) Decode(buffer *net.Buffer) {
}

func (this *OpenAllPetGrid_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(28)
}

func (this *OpenAllPetGrid_In) ByteSize() int {
	size := 2
	return size
}

func (this *OpenAllPetGrid_Out) Decode(buffer *net.Buffer) {
}

func (this *OpenAllPetGrid_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(28)
}

func (this *OpenAllPetGrid_Out) ByteSize() int {
	size := 2
	return size
}

func (this *CreateAnnouncement_In) Decode(buffer *net.Buffer) {
}

func (this *CreateAnnouncement_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(29)
}

func (this *CreateAnnouncement_In) ByteSize() int {
	size := 2
	return size
}

func (this *CreateAnnouncement_Out) Decode(buffer *net.Buffer) {
}

func (this *CreateAnnouncement_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(29)
}

func (this *CreateAnnouncement_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AddHeart_In) Decode(buffer *net.Buffer) {
	this.Number = int16(buffer.ReadUint16LE())
}

func (this *AddHeart_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(30)
	buffer.WriteUint16LE(uint16(this.Number))
}

func (this *AddHeart_In) ByteSize() int {
	size := 4
	return size
}

func (this *AddHeart_Out) Decode(buffer *net.Buffer) {
}

func (this *AddHeart_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(30)
}

func (this *AddHeart_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ResetHardLevelEnterCount_In) Decode(buffer *net.Buffer) {
}

func (this *ResetHardLevelEnterCount_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(31)
}

func (this *ResetHardLevelEnterCount_In) ByteSize() int {
	size := 2
	return size
}

func (this *ResetHardLevelEnterCount_Out) Decode(buffer *net.Buffer) {
}

func (this *ResetHardLevelEnterCount_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(31)
}

func (this *ResetHardLevelEnterCount_Out) ByteSize() int {
	size := 2
	return size
}

func (this *OpenHardLevel_In) Decode(buffer *net.Buffer) {
	this.LevelId = int32(buffer.ReadUint32LE())
}

func (this *OpenHardLevel_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(32)
	buffer.WriteUint32LE(uint32(this.LevelId))
}

func (this *OpenHardLevel_In) ByteSize() int {
	size := 6
	return size
}

func (this *OpenHardLevel_Out) Decode(buffer *net.Buffer) {
}

func (this *OpenHardLevel_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(32)
}

func (this *OpenHardLevel_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SetVipLevel_In) Decode(buffer *net.Buffer) {
	this.Level = int16(buffer.ReadUint16LE())
}

func (this *SetVipLevel_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(33)
	buffer.WriteUint16LE(uint16(this.Level))
}

func (this *SetVipLevel_In) ByteSize() int {
	size := 4
	return size
}

func (this *SetVipLevel_Out) Decode(buffer *net.Buffer) {
}

func (this *SetVipLevel_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(33)
}

func (this *SetVipLevel_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SetResourceLevelOpenDay_In) Decode(buffer *net.Buffer) {
	this.LevelType = int8(buffer.ReadUint8())
	this.OpenDay = int8(buffer.ReadUint8())
}

func (this *SetResourceLevelOpenDay_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(34)
	buffer.WriteUint8(uint8(this.LevelType))
	buffer.WriteUint8(uint8(this.OpenDay))
}

func (this *SetResourceLevelOpenDay_In) ByteSize() int {
	size := 4
	return size
}

func (this *SetResourceLevelOpenDay_Out) Decode(buffer *net.Buffer) {
}

func (this *SetResourceLevelOpenDay_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(34)
}

func (this *SetResourceLevelOpenDay_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ResetResourceLevelOpenDay_In) Decode(buffer *net.Buffer) {
}

func (this *ResetResourceLevelOpenDay_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(35)
}

func (this *ResetResourceLevelOpenDay_In) ByteSize() int {
	size := 2
	return size
}

func (this *ResetResourceLevelOpenDay_Out) Decode(buffer *net.Buffer) {
}

func (this *ResetResourceLevelOpenDay_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(35)
}

func (this *ResetResourceLevelOpenDay_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ResetArenaDailyCount_In) Decode(buffer *net.Buffer) {
}

func (this *ResetArenaDailyCount_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(36)
}

func (this *ResetArenaDailyCount_In) ByteSize() int {
	size := 2
	return size
}

func (this *ResetArenaDailyCount_Out) Decode(buffer *net.Buffer) {
}

func (this *ResetArenaDailyCount_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(36)
}

func (this *ResetArenaDailyCount_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ResetSwordSoulDrawCd_In) Decode(buffer *net.Buffer) {
}

func (this *ResetSwordSoulDrawCd_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(37)
}

func (this *ResetSwordSoulDrawCd_In) ByteSize() int {
	size := 2
	return size
}

func (this *ResetSwordSoulDrawCd_Out) Decode(buffer *net.Buffer) {
}

func (this *ResetSwordSoulDrawCd_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(37)
}

func (this *ResetSwordSoulDrawCd_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SetFirstLoginTime_In) Decode(buffer *net.Buffer) {
	this.Timestamp = int64(buffer.ReadUint64LE())
}

func (this *SetFirstLoginTime_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(38)
	buffer.WriteUint64LE(uint64(this.Timestamp))
}

func (this *SetFirstLoginTime_In) ByteSize() int {
	size := 10
	return size
}

func (this *SetFirstLoginTime_Out) Decode(buffer *net.Buffer) {
}

func (this *SetFirstLoginTime_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(38)
}

func (this *SetFirstLoginTime_Out) ByteSize() int {
	size := 2
	return size
}

func (this *EarlierFirstLoginTime_In) Decode(buffer *net.Buffer) {
}

func (this *EarlierFirstLoginTime_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(39)
}

func (this *EarlierFirstLoginTime_In) ByteSize() int {
	size := 2
	return size
}

func (this *EarlierFirstLoginTime_Out) Decode(buffer *net.Buffer) {
}

func (this *EarlierFirstLoginTime_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(39)
}

func (this *EarlierFirstLoginTime_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ResetServerOpenTime_In) Decode(buffer *net.Buffer) {
}

func (this *ResetServerOpenTime_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(40)
}

func (this *ResetServerOpenTime_In) ByteSize() int {
	size := 2
	return size
}

func (this *ResetServerOpenTime_Out) Decode(buffer *net.Buffer) {
}

func (this *ResetServerOpenTime_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(40)
}

func (this *ResetServerOpenTime_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ClearTraderRefreshTime_In) Decode(buffer *net.Buffer) {
	this.TraderId = int16(buffer.ReadUint16LE())
}

func (this *ClearTraderRefreshTime_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(41)
	buffer.WriteUint16LE(uint16(this.TraderId))
}

func (this *ClearTraderRefreshTime_In) ByteSize() int {
	size := 4
	return size
}

func (this *ClearTraderRefreshTime_Out) Decode(buffer *net.Buffer) {
}

func (this *ClearTraderRefreshTime_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(41)
}

func (this *ClearTraderRefreshTime_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AddTraderRefreshTime_In) Decode(buffer *net.Buffer) {
	this.TraderId = int16(buffer.ReadUint16LE())
	this.Timing = int64(buffer.ReadUint64LE())
}

func (this *AddTraderRefreshTime_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(42)
	buffer.WriteUint16LE(uint16(this.TraderId))
	buffer.WriteUint64LE(uint64(this.Timing))
}

func (this *AddTraderRefreshTime_In) ByteSize() int {
	size := 12
	return size
}

func (this *AddTraderRefreshTime_Out) Decode(buffer *net.Buffer) {
}

func (this *AddTraderRefreshTime_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(42)
}

func (this *AddTraderRefreshTime_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ClearTraderSchedule_In) Decode(buffer *net.Buffer) {
	this.TraderId = int16(buffer.ReadUint16LE())
}

func (this *ClearTraderSchedule_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(43)
	buffer.WriteUint16LE(uint16(this.TraderId))
}

func (this *ClearTraderSchedule_In) ByteSize() int {
	size := 4
	return size
}

func (this *ClearTraderSchedule_Out) Decode(buffer *net.Buffer) {
}

func (this *ClearTraderSchedule_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(43)
}

func (this *ClearTraderSchedule_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AddTraderSchedule_In) Decode(buffer *net.Buffer) {
	this.TraderId = int16(buffer.ReadUint16LE())
	this.Expire = int64(buffer.ReadUint64LE())
	this.Showup = int64(buffer.ReadUint64LE())
	this.Disappear = int64(buffer.ReadUint64LE())
}

func (this *AddTraderSchedule_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(44)
	buffer.WriteUint16LE(uint16(this.TraderId))
	buffer.WriteUint64LE(uint64(this.Expire))
	buffer.WriteUint64LE(uint64(this.Showup))
	buffer.WriteUint64LE(uint64(this.Disappear))
}

func (this *AddTraderSchedule_In) ByteSize() int {
	size := 28
	return size
}

func (this *AddTraderSchedule_Out) Decode(buffer *net.Buffer) {
}

func (this *AddTraderSchedule_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(44)
}

func (this *AddTraderSchedule_Out) ByteSize() int {
	size := 2
	return size
}

func (this *OpenTown_In) Decode(buffer *net.Buffer) {
	this.TownId = int16(buffer.ReadUint16LE())
}

func (this *OpenTown_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(45)
	buffer.WriteUint16LE(uint16(this.TownId))
}

func (this *OpenTown_In) ByteSize() int {
	size := 4
	return size
}

func (this *OpenTown_Out) Decode(buffer *net.Buffer) {
}

func (this *OpenTown_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(45)
}

func (this *OpenTown_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AddGlobalMail_In) Decode(buffer *net.Buffer) {
	this.SendDelay = int64(buffer.ReadUint64LE())
	this.ExpireDelay = int64(buffer.ReadUint64LE())
}

func (this *AddGlobalMail_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(46)
	buffer.WriteUint64LE(uint64(this.SendDelay))
	buffer.WriteUint64LE(uint64(this.ExpireDelay))
}

func (this *AddGlobalMail_In) ByteSize() int {
	size := 18
	return size
}

func (this *AddGlobalMail_Out) Decode(buffer *net.Buffer) {
}

func (this *AddGlobalMail_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(46)
}

func (this *AddGlobalMail_Out) ByteSize() int {
	size := 2
	return size
}

func (this *CreateAnnouncementWithoutTpl_In) Decode(buffer *net.Buffer) {
}

func (this *CreateAnnouncementWithoutTpl_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(47)
}

func (this *CreateAnnouncementWithoutTpl_In) ByteSize() int {
	size := 2
	return size
}

func (this *CreateAnnouncementWithoutTpl_Out) Decode(buffer *net.Buffer) {
}

func (this *CreateAnnouncementWithoutTpl_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(47)
}

func (this *CreateAnnouncementWithoutTpl_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SetLoginDay_In) Decode(buffer *net.Buffer) {
	this.Days = int32(buffer.ReadUint32LE())
}

func (this *SetLoginDay_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(48)
	buffer.WriteUint32LE(uint32(this.Days))
}

func (this *SetLoginDay_In) ByteSize() int {
	size := 6
	return size
}

func (this *SetLoginDay_Out) Decode(buffer *net.Buffer) {
}

func (this *SetLoginDay_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(48)
}

func (this *SetLoginDay_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ResetLoginAward_In) Decode(buffer *net.Buffer) {
}

func (this *ResetLoginAward_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(49)
}

func (this *ResetLoginAward_In) ByteSize() int {
	size := 2
	return size
}

func (this *ResetLoginAward_Out) Decode(buffer *net.Buffer) {
}

func (this *ResetLoginAward_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(49)
}

func (this *ResetLoginAward_Out) ByteSize() int {
	size := 2
	return size
}

func (this *RestPlayerAwardLock_In) Decode(buffer *net.Buffer) {
}

func (this *RestPlayerAwardLock_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(50)
}

func (this *RestPlayerAwardLock_In) ByteSize() int {
	size := 2
	return size
}

func (this *RestPlayerAwardLock_Out) Decode(buffer *net.Buffer) {
}

func (this *RestPlayerAwardLock_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(50)
}

func (this *RestPlayerAwardLock_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ResetRainbowLevel_In) Decode(buffer *net.Buffer) {
}

func (this *ResetRainbowLevel_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(51)
}

func (this *ResetRainbowLevel_In) ByteSize() int {
	size := 2
	return size
}

func (this *ResetRainbowLevel_Out) Decode(buffer *net.Buffer) {
}

func (this *ResetRainbowLevel_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(51)
}

func (this *ResetRainbowLevel_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SetRainbowLevel_In) Decode(buffer *net.Buffer) {
	this.Segment = int16(buffer.ReadUint16LE())
	this.Order = int8(buffer.ReadUint8())
}

func (this *SetRainbowLevel_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(52)
	buffer.WriteUint16LE(uint16(this.Segment))
	buffer.WriteUint8(uint8(this.Order))
}

func (this *SetRainbowLevel_In) ByteSize() int {
	size := 5
	return size
}

func (this *SetRainbowLevel_Out) Decode(buffer *net.Buffer) {
}

func (this *SetRainbowLevel_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(52)
}

func (this *SetRainbowLevel_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SendPushNotification_In) Decode(buffer *net.Buffer) {
}

func (this *SendPushNotification_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(53)
}

func (this *SendPushNotification_In) ByteSize() int {
	size := 2
	return size
}

func (this *SendPushNotification_Out) Decode(buffer *net.Buffer) {
}

func (this *SendPushNotification_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(53)
}

func (this *SendPushNotification_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ResetPetVirtualEnv_In) Decode(buffer *net.Buffer) {
}

func (this *ResetPetVirtualEnv_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(54)
}

func (this *ResetPetVirtualEnv_In) ByteSize() int {
	size := 2
	return size
}

func (this *ResetPetVirtualEnv_Out) Decode(buffer *net.Buffer) {
}

func (this *ResetPetVirtualEnv_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(54)
}

func (this *ResetPetVirtualEnv_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AddFame_In) Decode(buffer *net.Buffer) {
	this.System = int16(buffer.ReadUint16LE())
	this.Val = int32(buffer.ReadUint32LE())
}

func (this *AddFame_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(55)
	buffer.WriteUint16LE(uint16(this.System))
	buffer.WriteUint32LE(uint32(this.Val))
}

func (this *AddFame_In) ByteSize() int {
	size := 8
	return size
}

func (this *AddFame_Out) Decode(buffer *net.Buffer) {
}

func (this *AddFame_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(55)
}

func (this *AddFame_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AddWorldChatMessage_In) Decode(buffer *net.Buffer) {
	this.Num = int16(buffer.ReadUint16LE())
}

func (this *AddWorldChatMessage_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(56)
	buffer.WriteUint16LE(uint16(this.Num))
}

func (this *AddWorldChatMessage_In) ByteSize() int {
	size := 4
	return size
}

func (this *AddWorldChatMessage_Out) Decode(buffer *net.Buffer) {
}

func (this *AddWorldChatMessage_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(56)
}

func (this *AddWorldChatMessage_Out) ByteSize() int {
	size := 2
	return size
}

func (this *MonthCard_In) Decode(buffer *net.Buffer) {
}

func (this *MonthCard_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(57)
}

func (this *MonthCard_In) ByteSize() int {
	size := 2
	return size
}

func (this *MonthCard_Out) Decode(buffer *net.Buffer) {
}

func (this *MonthCard_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(57)
}

func (this *MonthCard_Out) ByteSize() int {
	size := 2
	return size
}

func (this *EnterSandbox_In) Decode(buffer *net.Buffer) {
}

func (this *EnterSandbox_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(58)
}

func (this *EnterSandbox_In) ByteSize() int {
	size := 2
	return size
}

func (this *EnterSandbox_Out) Decode(buffer *net.Buffer) {
}

func (this *EnterSandbox_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(58)
}

func (this *EnterSandbox_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SandboxRollback_In) Decode(buffer *net.Buffer) {
}

func (this *SandboxRollback_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(59)
}

func (this *SandboxRollback_In) ByteSize() int {
	size := 2
	return size
}

func (this *SandboxRollback_Out) Decode(buffer *net.Buffer) {
}

func (this *SandboxRollback_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(59)
}

func (this *SandboxRollback_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ExitSandbox_In) Decode(buffer *net.Buffer) {
}

func (this *ExitSandbox_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(60)
}

func (this *ExitSandbox_In) ByteSize() int {
	size := 2
	return size
}

func (this *ExitSandbox_Out) Decode(buffer *net.Buffer) {
}

func (this *ExitSandbox_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(60)
}

func (this *ExitSandbox_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ResetShadedMissions_In) Decode(buffer *net.Buffer) {
	this.LevelId = int32(buffer.ReadUint32LE())
}

func (this *ResetShadedMissions_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(61)
	buffer.WriteUint32LE(uint32(this.LevelId))
}

func (this *ResetShadedMissions_In) ByteSize() int {
	size := 6
	return size
}

func (this *ResetShadedMissions_Out) Decode(buffer *net.Buffer) {
}

func (this *ResetShadedMissions_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(61)
}

func (this *ResetShadedMissions_Out) ByteSize() int {
	size := 2
	return size
}

func (this *CleanCornucopia_In) Decode(buffer *net.Buffer) {
}

func (this *CleanCornucopia_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(62)
}

func (this *CleanCornucopia_In) ByteSize() int {
	size := 2
	return size
}

func (this *CleanCornucopia_Out) Decode(buffer *net.Buffer) {
}

func (this *CleanCornucopia_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(62)
}

func (this *CleanCornucopia_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AddTotem_In) Decode(buffer *net.Buffer) {
	this.TotemId = int16(buffer.ReadUint16LE())
}

func (this *AddTotem_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(63)
	buffer.WriteUint16LE(uint16(this.TotemId))
}

func (this *AddTotem_In) ByteSize() int {
	size := 4
	return size
}

func (this *AddTotem_Out) Decode(buffer *net.Buffer) {
}

func (this *AddTotem_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(63)
}

func (this *AddTotem_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AddRune_In) Decode(buffer *net.Buffer) {
	this.JadeNum = int32(buffer.ReadUint32LE())
	this.RockNum = int32(buffer.ReadUint32LE())
}

func (this *AddRune_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(64)
	buffer.WriteUint32LE(uint32(this.JadeNum))
	buffer.WriteUint32LE(uint32(this.RockNum))
}

func (this *AddRune_In) ByteSize() int {
	size := 10
	return size
}

func (this *AddRune_Out) Decode(buffer *net.Buffer) {
}

func (this *AddRune_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(64)
}

func (this *AddRune_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SendRareItemMessage_In) Decode(buffer *net.Buffer) {
}

func (this *SendRareItemMessage_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(65)
}

func (this *SendRareItemMessage_In) ByteSize() int {
	size := 2
	return size
}

func (this *SendRareItemMessage_Out) Decode(buffer *net.Buffer) {
}

func (this *SendRareItemMessage_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(65)
}

func (this *SendRareItemMessage_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AddSwordDrivingAction_In) Decode(buffer *net.Buffer) {
	this.Point = int16(buffer.ReadUint16LE())
}

func (this *AddSwordDrivingAction_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(66)
	buffer.WriteUint16LE(uint16(this.Point))
}

func (this *AddSwordDrivingAction_In) ByteSize() int {
	size := 4
	return size
}

func (this *AddSwordDrivingAction_Out) Decode(buffer *net.Buffer) {
}

func (this *AddSwordDrivingAction_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(66)
}

func (this *AddSwordDrivingAction_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ResetDrivingSwordData_In) Decode(buffer *net.Buffer) {
	this.Cloud = int16(buffer.ReadUint16LE())
}

func (this *ResetDrivingSwordData_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(67)
	buffer.WriteUint16LE(uint16(this.Cloud))
}

func (this *ResetDrivingSwordData_In) ByteSize() int {
	size := 4
	return size
}

func (this *ResetDrivingSwordData_Out) Decode(buffer *net.Buffer) {
}

func (this *ResetDrivingSwordData_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(67)
}

func (this *ResetDrivingSwordData_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AddSwordSoulFragment_In) Decode(buffer *net.Buffer) {
	this.Number = int64(buffer.ReadUint64LE())
}

func (this *AddSwordSoulFragment_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(68)
	buffer.WriteUint64LE(uint64(this.Number))
}

func (this *AddSwordSoulFragment_In) ByteSize() int {
	size := 10
	return size
}

func (this *AddSwordSoulFragment_Out) Decode(buffer *net.Buffer) {
}

func (this *AddSwordSoulFragment_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(68)
}

func (this *AddSwordSoulFragment_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ResetMoneyTreeStatus_In) Decode(buffer *net.Buffer) {
}

func (this *ResetMoneyTreeStatus_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(69)
}

func (this *ResetMoneyTreeStatus_In) ByteSize() int {
	size := 2
	return size
}

func (this *ResetMoneyTreeStatus_Out) Decode(buffer *net.Buffer) {
}

func (this *ResetMoneyTreeStatus_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(69)
}

func (this *ResetMoneyTreeStatus_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ResetTodayMoneyTree_In) Decode(buffer *net.Buffer) {
}

func (this *ResetTodayMoneyTree_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(70)
}

func (this *ResetTodayMoneyTree_In) ByteSize() int {
	size := 2
	return size
}

func (this *ResetTodayMoneyTree_Out) Decode(buffer *net.Buffer) {
}

func (this *ResetTodayMoneyTree_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(70)
}

func (this *ResetTodayMoneyTree_Out) ByteSize() int {
	size := 2
	return size
}

func (this *CleanSwordSoulIngotDrawNums_In) Decode(buffer *net.Buffer) {
}

func (this *CleanSwordSoulIngotDrawNums_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(71)
}

func (this *CleanSwordSoulIngotDrawNums_In) ByteSize() int {
	size := 2
	return size
}

func (this *CleanSwordSoulIngotDrawNums_Out) Decode(buffer *net.Buffer) {
}

func (this *CleanSwordSoulIngotDrawNums_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(71)
}

func (this *CleanSwordSoulIngotDrawNums_Out) ByteSize() int {
	size := 2
	return size
}

func (this *PunchDrivingSwordCloud_In) Decode(buffer *net.Buffer) {
}

func (this *PunchDrivingSwordCloud_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(72)
}

func (this *PunchDrivingSwordCloud_In) ByteSize() int {
	size := 2
	return size
}

func (this *PunchDrivingSwordCloud_Out) Decode(buffer *net.Buffer) {
}

func (this *PunchDrivingSwordCloud_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(72)
}

func (this *PunchDrivingSwordCloud_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ClearCliqueDailyDonate_In) Decode(buffer *net.Buffer) {
}

func (this *ClearCliqueDailyDonate_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(73)
}

func (this *ClearCliqueDailyDonate_In) ByteSize() int {
	size := 2
	return size
}

func (this *ClearCliqueDailyDonate_Out) Decode(buffer *net.Buffer) {
}

func (this *ClearCliqueDailyDonate_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(73)
}

func (this *ClearCliqueDailyDonate_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SetCliqueContrib_In) Decode(buffer *net.Buffer) {
	this.Contrib = int64(buffer.ReadUint64LE())
}

func (this *SetCliqueContrib_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(74)
	buffer.WriteUint64LE(uint64(this.Contrib))
}

func (this *SetCliqueContrib_In) ByteSize() int {
	size := 10
	return size
}

func (this *SetCliqueContrib_Out) Decode(buffer *net.Buffer) {
}

func (this *SetCliqueContrib_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(74)
}

func (this *SetCliqueContrib_Out) ByteSize() int {
	size := 2
	return size
}

func (this *RefreshCliqueWorship_In) Decode(buffer *net.Buffer) {
}

func (this *RefreshCliqueWorship_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(75)
}

func (this *RefreshCliqueWorship_In) ByteSize() int {
	size := 2
	return size
}

func (this *RefreshCliqueWorship_Out) Decode(buffer *net.Buffer) {
}

func (this *RefreshCliqueWorship_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(75)
}

func (this *RefreshCliqueWorship_Out) ByteSize() int {
	size := 2
	return size
}

func (this *CliqueEscortHijackBattleWin_In) Decode(buffer *net.Buffer) {
	this.BoatId = int64(buffer.ReadUint64LE())
}

func (this *CliqueEscortHijackBattleWin_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(76)
	buffer.WriteUint64LE(uint64(this.BoatId))
}

func (this *CliqueEscortHijackBattleWin_In) ByteSize() int {
	size := 10
	return size
}

func (this *CliqueEscortHijackBattleWin_Out) Decode(buffer *net.Buffer) {
}

func (this *CliqueEscortHijackBattleWin_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(76)
}

func (this *CliqueEscortHijackBattleWin_Out) ByteSize() int {
	size := 2
	return size
}

func (this *CliqueEscortRecoverBattleWin_In) Decode(buffer *net.Buffer) {
	this.BoatId = int64(buffer.ReadUint64LE())
}

func (this *CliqueEscortRecoverBattleWin_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(77)
	buffer.WriteUint64LE(uint64(this.BoatId))
}

func (this *CliqueEscortRecoverBattleWin_In) ByteSize() int {
	size := 10
	return size
}

func (this *CliqueEscortRecoverBattleWin_Out) Decode(buffer *net.Buffer) {
}

func (this *CliqueEscortRecoverBattleWin_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(77)
}

func (this *CliqueEscortRecoverBattleWin_Out) ByteSize() int {
	size := 2
	return size
}

func (this *CliqueEscortNotifyMessage_In) Decode(buffer *net.Buffer) {
}

func (this *CliqueEscortNotifyMessage_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(87)
}

func (this *CliqueEscortNotifyMessage_In) ByteSize() int {
	size := 2
	return size
}

func (this *CliqueEscortNotifyMessage_Out) Decode(buffer *net.Buffer) {
}

func (this *CliqueEscortNotifyMessage_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(87)
}

func (this *CliqueEscortNotifyMessage_Out) ByteSize() int {
	size := 2
	return size
}

func (this *CliqueEscortNotifyDailyQuest_In) Decode(buffer *net.Buffer) {
}

func (this *CliqueEscortNotifyDailyQuest_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(88)
}

func (this *CliqueEscortNotifyDailyQuest_In) ByteSize() int {
	size := 2
	return size
}

func (this *CliqueEscortNotifyDailyQuest_Out) Decode(buffer *net.Buffer) {
}

func (this *CliqueEscortNotifyDailyQuest_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(88)
}

func (this *CliqueEscortNotifyDailyQuest_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SetCliqueBuildingLevel_In) Decode(buffer *net.Buffer) {
	this.Building = int32(buffer.ReadUint32LE())
	this.Level = int16(buffer.ReadUint16LE())
}

func (this *SetCliqueBuildingLevel_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(89)
	buffer.WriteUint32LE(uint32(this.Building))
	buffer.WriteUint16LE(uint16(this.Level))
}

func (this *SetCliqueBuildingLevel_In) ByteSize() int {
	size := 8
	return size
}

func (this *SetCliqueBuildingLevel_Out) Decode(buffer *net.Buffer) {
}

func (this *SetCliqueBuildingLevel_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(89)
}

func (this *SetCliqueBuildingLevel_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SetCliqueBuildingMoney_In) Decode(buffer *net.Buffer) {
	this.Building = int32(buffer.ReadUint32LE())
	this.Money = int64(buffer.ReadUint64LE())
}

func (this *SetCliqueBuildingMoney_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(90)
	buffer.WriteUint32LE(uint32(this.Building))
	buffer.WriteUint64LE(uint64(this.Money))
}

func (this *SetCliqueBuildingMoney_In) ByteSize() int {
	size := 14
	return size
}

func (this *SetCliqueBuildingMoney_Out) Decode(buffer *net.Buffer) {
}

func (this *SetCliqueBuildingMoney_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(90)
}

func (this *SetCliqueBuildingMoney_Out) ByteSize() int {
	size := 2
	return size
}

func (this *EscortBench_In) Decode(buffer *net.Buffer) {
}

func (this *EscortBench_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(91)
}

func (this *EscortBench_In) ByteSize() int {
	size := 2
	return size
}

func (this *EscortBench_Out) Decode(buffer *net.Buffer) {
}

func (this *EscortBench_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(91)
}

func (this *EscortBench_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ResetCliqueEscortDailyNum_In) Decode(buffer *net.Buffer) {
}

func (this *ResetCliqueEscortDailyNum_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(92)
}

func (this *ResetCliqueEscortDailyNum_In) ByteSize() int {
	size := 2
	return size
}

func (this *ResetCliqueEscortDailyNum_Out) Decode(buffer *net.Buffer) {
}

func (this *ResetCliqueEscortDailyNum_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(92)
}

func (this *ResetCliqueEscortDailyNum_Out) ByteSize() int {
	size := 2
	return size
}

func (this *TakeAdditionQuest_In) Decode(buffer *net.Buffer) {
	this.FirstQuestId = int32(buffer.ReadUint32LE())
}

func (this *TakeAdditionQuest_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(93)
	buffer.WriteUint32LE(uint32(this.FirstQuestId))
}

func (this *TakeAdditionQuest_In) ByteSize() int {
	size := 6
	return size
}

func (this *TakeAdditionQuest_Out) Decode(buffer *net.Buffer) {
	this.Msg = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *TakeAdditionQuest_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(93)
	buffer.WriteUint16LE(uint16(len(this.Msg)))
	buffer.WriteBytes(this.Msg)
}

func (this *TakeAdditionQuest_Out) ByteSize() int {
	size := 4
	size += len(this.Msg)
	return size
}

func (this *SetMissionStarMax_In) Decode(buffer *net.Buffer) {
}

func (this *SetMissionStarMax_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(94)
}

func (this *SetMissionStarMax_In) ByteSize() int {
	size := 2
	return size
}

func (this *SetMissionStarMax_Out) Decode(buffer *net.Buffer) {
}

func (this *SetMissionStarMax_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(94)
}

func (this *SetMissionStarMax_Out) ByteSize() int {
	size := 2
	return size
}

func (this *CliqueBankCd_In) Decode(buffer *net.Buffer) {
}

func (this *CliqueBankCd_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(95)
}

func (this *CliqueBankCd_In) ByteSize() int {
	size := 2
	return size
}

func (this *CliqueBankCd_Out) Decode(buffer *net.Buffer) {
}

func (this *CliqueBankCd_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(95)
}

func (this *CliqueBankCd_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ResetDespairLandBattleNum_In) Decode(buffer *net.Buffer) {
}

func (this *ResetDespairLandBattleNum_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(96)
}

func (this *ResetDespairLandBattleNum_In) ByteSize() int {
	size := 2
	return size
}

func (this *ResetDespairLandBattleNum_Out) Decode(buffer *net.Buffer) {
}

func (this *ResetDespairLandBattleNum_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(96)
}

func (this *ResetDespairLandBattleNum_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ResetCliqueStoreSendTimes_In) Decode(buffer *net.Buffer) {
}

func (this *ResetCliqueStoreSendTimes_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(97)
}

func (this *ResetCliqueStoreSendTimes_In) ByteSize() int {
	size := 2
	return size
}

func (this *ResetCliqueStoreSendTimes_Out) Decode(buffer *net.Buffer) {
}

func (this *ResetCliqueStoreSendTimes_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(97)
}

func (this *ResetCliqueStoreSendTimes_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AddCliqueStoreDonate_In) Decode(buffer *net.Buffer) {
}

func (this *AddCliqueStoreDonate_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(98)
}

func (this *AddCliqueStoreDonate_In) ByteSize() int {
	size := 2
	return size
}

func (this *AddCliqueStoreDonate_Out) Decode(buffer *net.Buffer) {
}

func (this *AddCliqueStoreDonate_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(98)
}

func (this *AddCliqueStoreDonate_Out) ByteSize() int {
	size := 2
	return size
}

func (this *PassAllDespairLandLevel_In) Decode(buffer *net.Buffer) {
	this.Star = int8(buffer.ReadUint8())
	this.CampType = int8(buffer.ReadUint8())
}

func (this *PassAllDespairLandLevel_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(99)
	buffer.WriteUint8(uint8(this.Star))
	buffer.WriteUint8(uint8(this.CampType))
}

func (this *PassAllDespairLandLevel_In) ByteSize() int {
	size := 4
	return size
}

func (this *PassAllDespairLandLevel_Out) Decode(buffer *net.Buffer) {
}

func (this *PassAllDespairLandLevel_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(99)
}

func (this *PassAllDespairLandLevel_Out) ByteSize() int {
	size := 2
	return size
}

func (this *DespairLandDummyBossKill_In) Decode(buffer *net.Buffer) {
}

func (this *DespairLandDummyBossKill_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(100)
}

func (this *DespairLandDummyBossKill_In) ByteSize() int {
	size := 2
	return size
}

func (this *DespairLandDummyBossKill_Out) Decode(buffer *net.Buffer) {
}

func (this *DespairLandDummyBossKill_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(100)
}

func (this *DespairLandDummyBossKill_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AddTaoyuanItem_In) Decode(buffer *net.Buffer) {
	this.ItemId = int16(buffer.ReadUint16LE())
	this.Number = int16(buffer.ReadUint16LE())
}

func (this *AddTaoyuanItem_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(101)
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint16LE(uint16(this.Number))
}

func (this *AddTaoyuanItem_In) ByteSize() int {
	size := 6
	return size
}

func (this *AddTaoyuanItem_Out) Decode(buffer *net.Buffer) {
}

func (this *AddTaoyuanItem_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(101)
}

func (this *AddTaoyuanItem_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AddTaoyuanExp_In) Decode(buffer *net.Buffer) {
	this.AddExp = int64(buffer.ReadUint64LE())
}

func (this *AddTaoyuanExp_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(102)
	buffer.WriteUint64LE(uint64(this.AddExp))
}

func (this *AddTaoyuanExp_In) ByteSize() int {
	size := 10
	return size
}

func (this *AddTaoyuanExp_Out) Decode(buffer *net.Buffer) {
}

func (this *AddTaoyuanExp_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(99)
	buffer.WriteUint8(102)
}

func (this *AddTaoyuanExp_Out) ByteSize() int {
	size := 2
	return size
}
