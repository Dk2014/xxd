package client_test

import "core/net"
import "core/fail"
import "fmt"
import "game_server/api/protocol"
import (
	"game_server/api/protocol/announcement_api"
	"game_server/api/protocol/arena_api"
	"game_server/api/protocol/awaken_api"
	"game_server/api/protocol/battle_api"
	"game_server/api/protocol/battle_pet_api"
	"game_server/api/protocol/channel_api"
	"game_server/api/protocol/clique_api"
	"game_server/api/protocol/clique_building_api"
	"game_server/api/protocol/clique_escort_api"
	"game_server/api/protocol/clique_quest_api"
	"game_server/api/protocol/daily_sign_in_api"
	"game_server/api/protocol/debug_api"
	"game_server/api/protocol/despair_land_api"
	"game_server/api/protocol/draw_api"
	"game_server/api/protocol/driving_sword_api"
	"game_server/api/protocol/event_api"
	"game_server/api/protocol/fashion_api"
	"game_server/api/protocol/friend_api"
	"game_server/api/protocol/ghost_api"
	"game_server/api/protocol/item_api"
	"game_server/api/protocol/mail_api"
	"game_server/api/protocol/meditation_api"
	"game_server/api/protocol/mission_api"
	"game_server/api/protocol/money_tree_api"
	"game_server/api/protocol/multi_level_api"
	"game_server/api/protocol/notify_api"
	"game_server/api/protocol/pet_virtual_env_api"
	"game_server/api/protocol/player_api"
	"game_server/api/protocol/push_notify_api"
	"game_server/api/protocol/quest_api"
	"game_server/api/protocol/rainbow_api"
	"game_server/api/protocol/role_api"
	"game_server/api/protocol/server_info_api"
	"game_server/api/protocol/skill_api"
	"game_server/api/protocol/sword_soul_api"
	"game_server/api/protocol/taoyuan_api"
	"game_server/api/protocol/team_api"
	"game_server/api/protocol/totem_api"
	"game_server/api/protocol/tower_api"
	"game_server/api/protocol/town_api"
	"game_server/api/protocol/trader_api"
	"game_server/api/protocol/vip_api"
)

type Client struct {
	Name string

	closed bool
	conn   *net.Conn

	OutPlayer_Info                                func(*player_api.Info_Out)
	OutPlayer_Relogin                             func(*player_api.Relogin_Out)
	OutPlayer_BuyPhysical                         func(*player_api.BuyPhysical_Out)
	OutPlayer_SetPlayKey                          func(*player_api.SetPlayKey_Out)
	OutPlayer_GetTime                             func(*player_api.GetTime_Out)
	OutPlayer_FromPlatformLogin                   func(*player_api.FromPlatformLogin_Out)
	OutPlayer_BuyCoins                            func(*player_api.BuyCoins_Out)
	OutPlayer_GetLoginInfo                        func(*player_api.GetLoginInfo_Out)
	OutPlayer_CrossLoginGameServer                func(*player_api.CrossLoginGameServer_Out)
	OutPlayer_NotifyGlobalServerInfo              func(*player_api.NotifyGlobalServerInfo_Out)
	OutPlayer_GlobalLogin                         func(*player_api.GlobalLogin_Out)
	OutPlayer_GetIngot                            func(*player_api.GetIngot_Out)
	OutPlayer_SystemFame                          func(*player_api.SystemFame_Out)
	OutPlayer_GetRanks                            func(*player_api.GetRanks_Out)
	OutTown_Enter                                 func(*town_api.Enter_Out)
	OutTown_Leave                                 func(*town_api.Leave_Out)
	OutTown_Move                                  func(*town_api.Move_Out)
	OutTown_TalkedNpcList                         func(*town_api.TalkedNpcList_Out)
	OutTown_NpcTalk                               func(*town_api.NpcTalk_Out)
	OutTown_NotifyTownPlayers                     func(*town_api.NotifyTownPlayers_Out)
	OutTown_UpdateTownPlayer                      func(*town_api.UpdateTownPlayer_Out)
	OutTown_UpdateTownPlayerMeditationState       func(*town_api.UpdateTownPlayerMeditationState_Out)
	OutTown_ListOpenedTownTreasures               func(*town_api.ListOpenedTownTreasures_Out)
	OutTown_TakeTownTreasures                     func(*town_api.TakeTownTreasures_Out)
	OutTeam_GetFormationInfo                      func(*team_api.GetFormationInfo_Out)
	OutTeam_UpFormation                           func(*team_api.UpFormation_Out)
	OutTeam_DownFormation                         func(*team_api.DownFormation_Out)
	OutTeam_SwapFormation                         func(*team_api.SwapFormation_Out)
	OutTeam_ReplaceFormation                      func(*team_api.ReplaceFormation_Out)
	OutTeam_TrainingTeamship                      func(*team_api.TrainingTeamship_Out)
	OutRole_GetAllRoles                           func(*role_api.GetAllRoles_Out)
	OutRole_GetRoleInfo                           func(*role_api.GetRoleInfo_Out)
	OutRole_GetPlayerInfo                         func(*role_api.GetPlayerInfo_Out)
	OutRole_GetFightNum                           func(*role_api.GetFightNum_Out)
	OutRole_GetPlayerInfoWithOpenid               func(*role_api.GetPlayerInfoWithOpenid_Out)
	OutRole_LevelupRoleFriendship                 func(*role_api.LevelupRoleFriendship_Out)
	OutRole_RecruitBuddy                          func(*role_api.RecruitBuddy_Out)
	OutRole_GetRoleFightNum                       func(*role_api.GetRoleFightNum_Out)
	OutRole_ChangeRoleStatus                      func(*role_api.ChangeRoleStatus_Out)
	OutRole_GetInnRoleList                        func(*role_api.GetInnRoleList_Out)
	OutRole_BuddyCoop                             func(*role_api.BuddyCoop_Out)
	OutMission_Open                               func(*mission_api.Open_Out)
	OutMission_GetMissionLevel                    func(*mission_api.GetMissionLevel_Out)
	OutMission_EnterLevel                         func(*mission_api.EnterLevel_Out)
	OutMission_OpenLevelBox                       func(*mission_api.OpenLevelBox_Out)
	OutMission_UseIngotRelive                     func(*mission_api.UseIngotRelive_Out)
	OutMission_UseItem                            func(*mission_api.UseItem_Out)
	OutMission_Rebuild                            func(*mission_api.Rebuild_Out)
	OutMission_EnterExtendLevel                   func(*mission_api.EnterExtendLevel_Out)
	OutMission_GetExtendLevelInfo                 func(*mission_api.GetExtendLevelInfo_Out)
	OutMission_OpenSmallBox                       func(*mission_api.OpenSmallBox_Out)
	OutMission_NotifyCatchBattlePet               func(*mission_api.NotifyCatchBattlePet_Out)
	OutMission_EnterHardLevel                     func(*mission_api.EnterHardLevel_Out)
	OutMission_GetHardLevel                       func(*mission_api.GetHardLevel_Out)
	OutMission_GetBuddyLevelRoleId                func(*mission_api.GetBuddyLevelRoleId_Out)
	OutMission_SetBuddyLevelTeam                  func(*mission_api.SetBuddyLevelTeam_Out)
	OutMission_AutoFightLevel                     func(*mission_api.AutoFightLevel_Out)
	OutMission_EnterRainbowLevel                  func(*mission_api.EnterRainbowLevel_Out)
	OutMission_OpenMengYao                        func(*mission_api.OpenMengYao_Out)
	OutMission_GetMissionLevelByItemId            func(*mission_api.GetMissionLevelByItemId_Out)
	OutMission_BuyHardLevelTimes                  func(*mission_api.BuyHardLevelTimes_Out)
	OutMission_OpenRandomAwardBox                 func(*mission_api.OpenRandomAwardBox_Out)
	OutMission_EvaluateLevel                      func(*mission_api.EvaluateLevel_Out)
	OutMission_OpenShadedBox                      func(*mission_api.OpenShadedBox_Out)
	OutMission_GetMissionTotalStarInfo            func(*mission_api.GetMissionTotalStarInfo_Out)
	OutMission_GetMissionTotalStarAwards          func(*mission_api.GetMissionTotalStarAwards_Out)
	OutMission_ClearMissionLevelState             func(*mission_api.ClearMissionLevelState_Out)
	OutMission_BuyResourceMissionLevelTimes       func(*mission_api.BuyResourceMissionLevelTimes_Out)
	OutMission_GetDragonBall                      func(*mission_api.GetDragonBall_Out)
	OutMission_GetEventItem                       func(*mission_api.GetEventItem_Out)
	OutSkill_GetAllSkillsInfo                     func(*skill_api.GetAllSkillsInfo_Out)
	OutSkill_EquipSkill                           func(*skill_api.EquipSkill_Out)
	OutSkill_UnequipSkill                         func(*skill_api.UnequipSkill_Out)
	OutSkill_StudySkillByCheat                    func(*skill_api.StudySkillByCheat_Out)
	OutSkill_TrainSkill                           func(*skill_api.TrainSkill_Out)
	OutSkill_FlushSkill                           func(*skill_api.FlushSkill_Out)
	OutBattle_StartBattle                         func(*battle_api.StartBattle_Out)
	OutBattle_NextRound                           func(*battle_api.NextRound_Out)
	OutBattle_End                                 func(*battle_api.End_Out)
	OutBattle_Escape                              func(*battle_api.Escape_Out)
	OutBattle_Fightnum                            func(*battle_api.Fightnum_Out)
	OutBattle_StartReadyTimeout                   func(*battle_api.StartReadyTimeout_Out)
	OutBattle_StartReady                          func(*battle_api.StartReady_Out)
	OutBattle_StateChange                         func(*battle_api.StateChange_Out)
	OutBattle_CallBattlePet                       func(*battle_api.CallBattlePet_Out)
	OutBattle_UseBuddySkill                       func(*battle_api.UseBuddySkill_Out)
	OutBattle_CallNewEnemys                       func(*battle_api.CallNewEnemys_Out)
	OutBattle_NewFighterGroup                     func(*battle_api.NewFighterGroup_Out)
	OutBattle_StartBattleForHijackBoat            func(*battle_api.StartBattleForHijackBoat_Out)
	OutBattle_StartBattleForRecoverBoat           func(*battle_api.StartBattleForRecoverBoat_Out)
	OutBattle_RoundReady                          func(*battle_api.RoundReady_Out)
	OutBattle_InitRound                           func(*battle_api.InitRound_Out)
	OutBattle_SetAuto                             func(*battle_api.SetAuto_Out)
	OutBattle_CancelAuto                          func(*battle_api.CancelAuto_Out)
	OutBattle_SetSkill                            func(*battle_api.SetSkill_Out)
	OutBattle_UseItem                             func(*battle_api.UseItem_Out)
	OutBattle_UseGhost                            func(*battle_api.UseGhost_Out)
	OutBattle_NotifyReady                         func(*battle_api.NotifyReady_Out)
	OutBattle_BattleReconnect                     func(*battle_api.BattleReconnect_Out)
	OutItem_GetAllItems                           func(*item_api.GetAllItems_Out)
	OutItem_DropItem                              func(*item_api.DropItem_Out)
	OutItem_BuyItem                               func(*item_api.BuyItem_Out)
	OutItem_SellItem                              func(*item_api.SellItem_Out)
	OutItem_Dress                                 func(*item_api.Dress_Out)
	OutItem_Undress                               func(*item_api.Undress_Out)
	OutItem_BuyItemBack                           func(*item_api.BuyItemBack_Out)
	OutItem_IsBagFull                             func(*item_api.IsBagFull_Out)
	OutItem_Decompose                             func(*item_api.Decompose_Out)
	OutItem_Refine                                func(*item_api.Refine_Out)
	OutItem_GetRecastInfo                         func(*item_api.GetRecastInfo_Out)
	OutItem_Recast                                func(*item_api.Recast_Out)
	OutItem_UseItem                               func(*item_api.UseItem_Out)
	OutItem_RoleUseCostItem                       func(*item_api.RoleUseCostItem_Out)
	OutItem_BatchUseItem                          func(*item_api.BatchUseItem_Out)
	OutItem_DragonBallExchangeNotify              func(*item_api.DragonBallExchangeNotify_Out)
	OutItem_OpenCornucopia                        func(*item_api.OpenCornucopia_Out)
	OutItem_GetSealedbooks                        func(*item_api.GetSealedbooks_Out)
	OutItem_ActivationSealedbook                  func(*item_api.ActivationSealedbook_Out)
	OutItem_ExchangeGhostCrystal                  func(*item_api.ExchangeGhostCrystal_Out)
	OutItem_ExchangeShopItem                      func(*item_api.ExchangeShopItem_Out)
	OutItem_GetHoildayItemList                    func(*item_api.GetHoildayItemList_Out)
	OutItem_ExchangeHoildayItem                   func(*item_api.ExchangeHoildayItem_Out)
	OutItem_BatchExchangeDragonBall               func(*item_api.BatchExchangeDragonBall_Out)
	OutNotify_PlayerKeyChanged                    func(*notify_api.PlayerKeyChanged_Out)
	OutNotify_MissionLevelLockChanged             func(*notify_api.MissionLevelLockChanged_Out)
	OutNotify_RoleExpChange                       func(*notify_api.RoleExpChange_Out)
	OutNotify_PhysicalChange                      func(*notify_api.PhysicalChange_Out)
	OutNotify_MoneyChange                         func(*notify_api.MoneyChange_Out)
	OutNotify_SkillAdd                            func(*notify_api.SkillAdd_Out)
	OutNotify_ItemChange                          func(*notify_api.ItemChange_Out)
	OutNotify_RoleBattleStatusChange              func(*notify_api.RoleBattleStatusChange_Out)
	OutNotify_NewMail                             func(*notify_api.NewMail_Out)
	OutNotify_HeartChange                         func(*notify_api.HeartChange_Out)
	OutNotify_QuestChange                         func(*notify_api.QuestChange_Out)
	OutNotify_TownLockChange                      func(*notify_api.TownLockChange_Out)
	OutNotify_Chat                                func(*notify_api.Chat_Out)
	OutNotify_FuncKeyChange                       func(*notify_api.FuncKeyChange_Out)
	OutNotify_ItemRecastStateRebuild              func(*notify_api.ItemRecastStateRebuild_Out)
	OutNotify_SendAnnouncement                    func(*notify_api.SendAnnouncement_Out)
	OutNotify_VipLevelChange                      func(*notify_api.VipLevelChange_Out)
	OutNotify_NotifyNewBuddy                      func(*notify_api.NotifyNewBuddy_Out)
	OutNotify_HardLevelLockChanged                func(*notify_api.HardLevelLockChanged_Out)
	OutNotify_SendSwordSoulDrawNumChange          func(*notify_api.SendSwordSoulDrawNumChange_Out)
	OutNotify_SendHaveNewGhost                    func(*notify_api.SendHaveNewGhost_Out)
	OutNotify_SendHeartRecoverTime                func(*notify_api.SendHeartRecoverTime_Out)
	OutNotify_SendGlobalMail                      func(*notify_api.SendGlobalMail_Out)
	OutNotify_SendPhysicalRecoverTime             func(*notify_api.SendPhysicalRecoverTime_Out)
	OutNotify_SendFashionChange                   func(*notify_api.SendFashionChange_Out)
	OutNotify_TransError                          func(*notify_api.TransError_Out)
	OutNotify_SendEventCenterChange               func(*notify_api.SendEventCenterChange_Out)
	OutNotify_MeditationState                     func(*notify_api.MeditationState_Out)
	OutNotify_DeleteAnnouncement                  func(*notify_api.DeleteAnnouncement_Out)
	OutNotify_SendHaveNewPet                      func(*notify_api.SendHaveNewPet_Out)
	OutNotify_SendLogout                          func(*notify_api.SendLogout_Out)
	OutNotify_FameChange                          func(*notify_api.FameChange_Out)
	OutNotify_NotifyMonthCardOpen                 func(*notify_api.NotifyMonthCardOpen_Out)
	OutNotify_NotifyMonthCardRenewal              func(*notify_api.NotifyMonthCardRenewal_Out)
	OutNotify_NotifyNewTotem                      func(*notify_api.NotifyNewTotem_Out)
	OutNotify_NotifyRuneChange                    func(*notify_api.NotifyRuneChange_Out)
	OutNotify_TaoyuanItemChange                   func(*notify_api.TaoyuanItemChange_Out)
	OutNotify_TaoyuanMessageRefresh               func(*notify_api.TaoyuanMessageRefresh_Out)
	OutNotify_TaoyuanQuestCanFinish               func(*notify_api.TaoyuanQuestCanFinish_Out)
	OutNotify_TaoyuanExpRefresh                   func(*notify_api.TaoyuanExpRefresh_Out)
	OutGhost_Info                                 func(*ghost_api.Info_Out)
	OutGhost_Equip                                func(*ghost_api.Equip_Out)
	OutGhost_Unequip                              func(*ghost_api.Unequip_Out)
	OutGhost_Swap                                 func(*ghost_api.Swap_Out)
	OutGhost_EquipPosChange                       func(*ghost_api.EquipPosChange_Out)
	OutGhost_Train                                func(*ghost_api.Train_Out)
	OutGhost_Upgrade                              func(*ghost_api.Upgrade_Out)
	OutGhost_Baptize                              func(*ghost_api.Baptize_Out)
	OutGhost_Compose                              func(*ghost_api.Compose_Out)
	OutGhost_TrainSkill                           func(*ghost_api.TrainSkill_Out)
	OutGhost_FlushAttr                            func(*ghost_api.FlushAttr_Out)
	OutGhost_RelationGhost                        func(*ghost_api.RelationGhost_Out)
	OutSwordSoul_Info                             func(*sword_soul_api.Info_Out)
	OutSwordSoul_Draw                             func(*sword_soul_api.Draw_Out)
	OutSwordSoul_Upgrade                          func(*sword_soul_api.Upgrade_Out)
	OutSwordSoul_Exchange                         func(*sword_soul_api.Exchange_Out)
	OutSwordSoul_Equip                            func(*sword_soul_api.Equip_Out)
	OutSwordSoul_Unequip                          func(*sword_soul_api.Unequip_Out)
	OutSwordSoul_EquipPosChange                   func(*sword_soul_api.EquipPosChange_Out)
	OutSwordSoul_Swap                             func(*sword_soul_api.Swap_Out)
	OutSwordSoul_IsBagFull                        func(*sword_soul_api.IsBagFull_Out)
	OutSwordSoul_EmptyPosNum                      func(*sword_soul_api.EmptyPosNum_Out)
	OutMail_GetList                               func(*mail_api.GetList_Out)
	OutMail_Read                                  func(*mail_api.Read_Out)
	OutMail_TakeAttachment                        func(*mail_api.TakeAttachment_Out)
	OutMail_GetInfos                              func(*mail_api.GetInfos_Out)
	OutMail_RequestGlobalMail                     func(*mail_api.RequestGlobalMail_Out)
	OutQuest_UpdateQuest                          func(*quest_api.UpdateQuest_Out)
	OutQuest_GetDailyInfo                         func(*quest_api.GetDailyInfo_Out)
	OutQuest_AwardDaily                           func(*quest_api.AwardDaily_Out)
	OutQuest_NotifyDailyChange                    func(*quest_api.NotifyDailyChange_Out)
	OutQuest_Guide                                func(*quest_api.Guide_Out)
	OutQuest_GetExtendQuestInfoByNpcId            func(*quest_api.GetExtendQuestInfoByNpcId_Out)
	OutQuest_TakeExtendQuestAward                 func(*quest_api.TakeExtendQuestAward_Out)
	OutQuest_GetPannelQuestInfo                   func(*quest_api.GetPannelQuestInfo_Out)
	OutQuest_GiveUpAdditionQuest                  func(*quest_api.GiveUpAdditionQuest_Out)
	OutQuest_TakeAdditionQuest                    func(*quest_api.TakeAdditionQuest_Out)
	OutQuest_TakeAdditionQuestAward               func(*quest_api.TakeAdditionQuestAward_Out)
	OutQuest_GetAdditionQuest                     func(*quest_api.GetAdditionQuest_Out)
	OutQuest_RefreshAdditionQuest                 func(*quest_api.RefreshAdditionQuest_Out)
	OutQuest_TakeQuestStarsAwaded                 func(*quest_api.TakeQuestStarsAwaded_Out)
	OutFriend_GetFriendList                       func(*friend_api.GetFriendList_Out)
	OutFriend_ListenByNick                        func(*friend_api.ListenByNick_Out)
	OutFriend_CancelListen                        func(*friend_api.CancelListen_Out)
	OutFriend_SendHeart                           func(*friend_api.SendHeart_Out)
	OutFriend_Chat                                func(*friend_api.Chat_Out)
	OutFriend_GetChatHistory                      func(*friend_api.GetChatHistory_Out)
	OutFriend_GetOfflineMessages                  func(*friend_api.GetOfflineMessages_Out)
	OutFriend_Block                               func(*friend_api.Block_Out)
	OutFriend_CancelBlock                         func(*friend_api.CancelBlock_Out)
	OutFriend_CleanChatHistory                    func(*friend_api.CleanChatHistory_Out)
	OutFriend_NotifyListenedState                 func(*friend_api.NotifyListenedState_Out)
	OutFriend_CurrentPlatformFriendNum            func(*friend_api.CurrentPlatformFriendNum_Out)
	OutFriend_GetSendHeartList                    func(*friend_api.GetSendHeartList_Out)
	OutFriend_SendHeartToAllFriends               func(*friend_api.SendHeartToAllFriends_Out)
	OutTower_GetInfo                              func(*tower_api.GetInfo_Out)
	OutTower_UseLadder                            func(*tower_api.UseLadder_Out)
	OutMultiLevel_CreateRoom                      func(*multi_level_api.CreateRoom_Out)
	OutMultiLevel_EnterRoom                       func(*multi_level_api.EnterRoom_Out)
	OutMultiLevel_NotifyRoomInfo                  func(*multi_level_api.NotifyRoomInfo_Out)
	OutMultiLevel_LeaveRoom                       func(*multi_level_api.LeaveRoom_Out)
	OutMultiLevel_NotifyJoinPartner               func(*multi_level_api.NotifyJoinPartner_Out)
	OutMultiLevel_ChangeBuddy                     func(*multi_level_api.ChangeBuddy_Out)
	OutMultiLevel_GetBaseInfo                     func(*multi_level_api.GetBaseInfo_Out)
	OutMultiLevel_GetOnlineFriend                 func(*multi_level_api.GetOnlineFriend_Out)
	OutMultiLevel_InviteIntoRoom                  func(*multi_level_api.InviteIntoRoom_Out)
	OutMultiLevel_NotifyRoomInvite                func(*multi_level_api.NotifyRoomInvite_Out)
	OutMultiLevel_NotifyPlayersInfo               func(*multi_level_api.NotifyPlayersInfo_Out)
	OutMultiLevel_RefuseRoomInvite                func(*multi_level_api.RefuseRoomInvite_Out)
	OutMultiLevel_NotifyRoomInviteRefuse          func(*multi_level_api.NotifyRoomInviteRefuse_Out)
	OutMultiLevel_NotifyUpdatePartner             func(*multi_level_api.NotifyUpdatePartner_Out)
	OutMultiLevel_MatchPlayer                     func(*multi_level_api.MatchPlayer_Out)
	OutMultiLevel_NotifyMatchPlayerSuccess        func(*multi_level_api.NotifyMatchPlayerSuccess_Out)
	OutMultiLevel_MatchRoom                       func(*multi_level_api.MatchRoom_Out)
	OutMultiLevel_CancelMatchRoom                 func(*multi_level_api.CancelMatchRoom_Out)
	OutMultiLevel_GetFormInfo                     func(*multi_level_api.GetFormInfo_Out)
	OutMultiLevel_SetForm                         func(*multi_level_api.SetForm_Out)
	OutMultiLevel_GetBattleRoleInfo               func(*multi_level_api.GetBattleRoleInfo_Out)
	OutMultiLevel_NotifyMatchPlayerInfo           func(*multi_level_api.NotifyMatchPlayerInfo_Out)
	OutMultiLevel_GetMatchInfo                    func(*multi_level_api.GetMatchInfo_Out)
	OutMultiLevel_NotifyFormInfo                  func(*multi_level_api.NotifyFormInfo_Out)
	OutMultiLevel_CancelMatchPlayer               func(*multi_level_api.CancelMatchPlayer_Out)
	OutBattlePet_GetPetInfo                       func(*battle_pet_api.GetPetInfo_Out)
	OutBattlePet_SetPet                           func(*battle_pet_api.SetPet_Out)
	OutBattlePet_SetPetSwap                       func(*battle_pet_api.SetPetSwap_Out)
	OutBattlePet_UpgradePet                       func(*battle_pet_api.UpgradePet_Out)
	OutBattlePet_TrainingPetSkill                 func(*battle_pet_api.TrainingPetSkill_Out)
	OutAnnouncement_GetList                       func(*announcement_api.GetList_Out)
	OutArena_Enter                                func(*arena_api.Enter_Out)
	OutArena_GetRecords                           func(*arena_api.GetRecords_Out)
	OutArena_AwardBox                             func(*arena_api.AwardBox_Out)
	OutArena_NotifyFailedCdTime                   func(*arena_api.NotifyFailedCdTime_Out)
	OutArena_StartBattle                          func(*arena_api.StartBattle_Out)
	OutArena_NotifyLoseRank                       func(*arena_api.NotifyLoseRank_Out)
	OutArena_NotifyArenaInfo                      func(*arena_api.NotifyArenaInfo_Out)
	OutArena_GetTopRank                           func(*arena_api.GetTopRank_Out)
	OutArena_CleanFailedCdTime                    func(*arena_api.CleanFailedCdTime_Out)
	OutVip_Info                                   func(*vip_api.Info_Out)
	OutVip_GetTotal                               func(*vip_api.GetTotal_Out)
	OutVip_VipLevelTotal                          func(*vip_api.VipLevelTotal_Out)
	OutVip_BuyTimes                               func(*vip_api.BuyTimes_Out)
	OutTrader_Info                                func(*trader_api.Info_Out)
	OutTrader_StoreState                          func(*trader_api.StoreState_Out)
	OutTrader_Buy                                 func(*trader_api.Buy_Out)
	OutTrader_Refresh                             func(*trader_api.Refresh_Out)
	OutDailySignIn_Info                           func(*daily_sign_in_api.Info_Out)
	OutDailySignIn_Sign                           func(*daily_sign_in_api.Sign_Out)
	OutDailySignIn_SignPastDay                    func(*daily_sign_in_api.SignPastDay_Out)
	OutRainbow_Info                               func(*rainbow_api.Info_Out)
	OutRainbow_Reset                              func(*rainbow_api.Reset_Out)
	OutRainbow_AwardInfo                          func(*rainbow_api.AwardInfo_Out)
	OutRainbow_TakeAward                          func(*rainbow_api.TakeAward_Out)
	OutRainbow_JumpToSegment                      func(*rainbow_api.JumpToSegment_Out)
	OutRainbow_AutoFight                          func(*rainbow_api.AutoFight_Out)
	OutEvent_LoginAwardInfo                       func(*event_api.LoginAwardInfo_Out)
	OutEvent_TakeLoginAward                       func(*event_api.TakeLoginAward_Out)
	OutEvent_GetEvents                            func(*event_api.GetEvents_Out)
	OutEvent_GetEventAward                        func(*event_api.GetEventAward_Out)
	OutEvent_PlayerEventPhysicalCost              func(*event_api.PlayerEventPhysicalCost_Out)
	OutEvent_PlayerMonthCardInfo                  func(*event_api.PlayerMonthCardInfo_Out)
	OutEvent_GetSevenInfo                         func(*event_api.GetSevenInfo_Out)
	OutEvent_GetSevenAward                        func(*event_api.GetSevenAward_Out)
	OutEvent_GetRichmanClubInfo                   func(*event_api.GetRichmanClubInfo_Out)
	OutEvent_GetRichmanClubAward                  func(*event_api.GetRichmanClubAward_Out)
	OutEvent_InfoShare                            func(*event_api.InfoShare_Out)
	OutEvent_InfoGroupBuy                         func(*event_api.InfoGroupBuy_Out)
	OutEvent_GetIngotChangeTotal                  func(*event_api.GetIngotChangeTotal_Out)
	OutEvent_GetEventTotalAward                   func(*event_api.GetEventTotalAward_Out)
	OutEvent_GetEventArenaRank                    func(*event_api.GetEventArenaRank_Out)
	OutEvent_GetEventTenDrawTimes                 func(*event_api.GetEventTenDrawTimes_Out)
	OutEvent_GetEventRechargeAward                func(*event_api.GetEventRechargeAward_Out)
	OutEvent_GetEventNewYear                      func(*event_api.GetEventNewYear_Out)
	OutEvent_QqVipContinue                        func(*event_api.QqVipContinue_Out)
	OutEvent_DailyOnlineInfo                      func(*event_api.DailyOnlineInfo_Out)
	OutEvent_TakeDailyOnlineAward                 func(*event_api.TakeDailyOnlineAward_Out)
	OutFashion_FashionInfo                        func(*fashion_api.FashionInfo_Out)
	OutFashion_DressFashion                       func(*fashion_api.DressFashion_Out)
	OutPushNotify_PushInfo                        func(*push_notify_api.PushInfo_Out)
	OutPushNotify_PushNotificationSwitch          func(*push_notify_api.PushNotificationSwitch_Out)
	OutMeditation_MeditationInfo                  func(*meditation_api.MeditationInfo_Out)
	OutMeditation_StartMeditation                 func(*meditation_api.StartMeditation_Out)
	OutMeditation_StopMeditation                  func(*meditation_api.StopMeditation_Out)
	OutPetVirtualEnv_Info                         func(*pet_virtual_env_api.Info_Out)
	OutPetVirtualEnv_TakeAward                    func(*pet_virtual_env_api.TakeAward_Out)
	OutPetVirtualEnv_AutoFight                    func(*pet_virtual_env_api.AutoFight_Out)
	OutPetVirtualEnv_PveKills                     func(*pet_virtual_env_api.PveKills_Out)
	OutChannel_GetLatestWorldChannelMessage       func(*channel_api.GetLatestWorldChannelMessage_Out)
	OutChannel_AddWorldChat                       func(*channel_api.AddWorldChat_Out)
	OutChannel_WorldChannelInfo                   func(*channel_api.WorldChannelInfo_Out)
	OutChannel_SendGlobalMessages                 func(*channel_api.SendGlobalMessages_Out)
	OutChannel_AddCliqueChat                      func(*channel_api.AddCliqueChat_Out)
	OutChannel_GetLatestCliqueMessages            func(*channel_api.GetLatestCliqueMessages_Out)
	OutChannel_SendCliqueMessage                  func(*channel_api.SendCliqueMessage_Out)
	OutDrivingSword_CloudMapInfo                  func(*driving_sword_api.CloudMapInfo_Out)
	OutDrivingSword_CloudClimb                    func(*driving_sword_api.CloudClimb_Out)
	OutDrivingSword_CloudTeleport                 func(*driving_sword_api.CloudTeleport_Out)
	OutDrivingSword_AreaTeleport                  func(*driving_sword_api.AreaTeleport_Out)
	OutDrivingSword_AreaMove                      func(*driving_sword_api.AreaMove_Out)
	OutDrivingSword_ExplorerStartBattle           func(*driving_sword_api.ExplorerStartBattle_Out)
	OutDrivingSword_ExplorerAward                 func(*driving_sword_api.ExplorerAward_Out)
	OutDrivingSword_ExplorerGarrison              func(*driving_sword_api.ExplorerGarrison_Out)
	OutDrivingSword_VisitMountain                 func(*driving_sword_api.VisitMountain_Out)
	OutDrivingSword_VisiterStartBattle            func(*driving_sword_api.VisiterStartBattle_Out)
	OutDrivingSword_VisiterAward                  func(*driving_sword_api.VisiterAward_Out)
	OutDrivingSword_MountainTreasureOpen          func(*driving_sword_api.MountainTreasureOpen_Out)
	OutDrivingSword_ListGarrisons                 func(*driving_sword_api.ListGarrisons_Out)
	OutDrivingSword_AwardGarrison                 func(*driving_sword_api.AwardGarrison_Out)
	OutDrivingSword_EndGarrison                   func(*driving_sword_api.EndGarrison_Out)
	OutDrivingSword_BuyAllowedAction              func(*driving_sword_api.BuyAllowedAction_Out)
	OutTotem_Info                                 func(*totem_api.Info_Out)
	OutTotem_CallTotem                            func(*totem_api.CallTotem_Out)
	OutTotem_Upgrade                              func(*totem_api.Upgrade_Out)
	OutTotem_Decompose                            func(*totem_api.Decompose_Out)
	OutTotem_Equip                                func(*totem_api.Equip_Out)
	OutTotem_Unequip                              func(*totem_api.Unequip_Out)
	OutTotem_EquipPosChange                       func(*totem_api.EquipPosChange_Out)
	OutTotem_Swap                                 func(*totem_api.Swap_Out)
	OutTotem_IsBagFull                            func(*totem_api.IsBagFull_Out)
	OutMoneyTree_GetTreeStatus                    func(*money_tree_api.GetTreeStatus_Out)
	OutMoneyTree_GetTreeMoney                     func(*money_tree_api.GetTreeMoney_Out)
	OutMoneyTree_WaveTree                         func(*money_tree_api.WaveTree_Out)
	OutClique_CreateClique                        func(*clique_api.CreateClique_Out)
	OutClique_ApplyJoinClique                     func(*clique_api.ApplyJoinClique_Out)
	OutClique_CancelApplyClique                   func(*clique_api.CancelApplyClique_Out)
	OutClique_ProcessJoinApply                    func(*clique_api.ProcessJoinApply_Out)
	OutClique_ElectOwner                          func(*clique_api.ElectOwner_Out)
	OutClique_MangeMember                         func(*clique_api.MangeMember_Out)
	OutClique_DestoryClique                       func(*clique_api.DestoryClique_Out)
	OutClique_UpdateAnnounce                      func(*clique_api.UpdateAnnounce_Out)
	OutClique_LeaveClique                         func(*clique_api.LeaveClique_Out)
	OutClique_ListClique                          func(*clique_api.ListClique_Out)
	OutClique_CliquePublicInfo                    func(*clique_api.CliquePublicInfo_Out)
	OutClique_CliqueInfo                          func(*clique_api.CliqueInfo_Out)
	OutClique_ListApply                           func(*clique_api.ListApply_Out)
	OutClique_OperaErrorNotify                    func(*clique_api.OperaErrorNotify_Out)
	OutClique_EnterClubhouse                      func(*clique_api.EnterClubhouse_Out)
	OutClique_LeaveClubhouse                      func(*clique_api.LeaveClubhouse_Out)
	OutClique_ClubMove                            func(*clique_api.ClubMove_Out)
	OutClique_NotifyClubhousePlayers              func(*clique_api.NotifyClubhousePlayers_Out)
	OutClique_NotifyNewPlayer                     func(*clique_api.NotifyNewPlayer_Out)
	OutClique_NotifyUpdatePlayer                  func(*clique_api.NotifyUpdatePlayer_Out)
	OutClique_CliquePublicInfoSummary             func(*clique_api.CliquePublicInfoSummary_Out)
	OutClique_CliqueAutoAudit                     func(*clique_api.CliqueAutoAudit_Out)
	OutClique_CliqueBaseDonate                    func(*clique_api.CliqueBaseDonate_Out)
	OutClique_NotifyLeaveClique                   func(*clique_api.NotifyLeaveClique_Out)
	OutClique_NotifyJoincliqueSuccess             func(*clique_api.NotifyJoincliqueSuccess_Out)
	OutClique_NotifyCliqueMangerAction            func(*clique_api.NotifyCliqueMangerAction_Out)
	OutClique_CliqueRecruitment                   func(*clique_api.CliqueRecruitment_Out)
	OutClique_NotifyCliqueAnnounce                func(*clique_api.NotifyCliqueAnnounce_Out)
	OutClique_NotifyCliqueElectowner              func(*clique_api.NotifyCliqueElectowner_Out)
	OutClique_QuickApply                          func(*clique_api.QuickApply_Out)
	OutClique_NotifyContribChange                 func(*clique_api.NotifyContribChange_Out)
	OutClique_OtherClique                         func(*clique_api.OtherClique_Out)
	OutCliqueBuilding_CliqueBaseDonate            func(*clique_building_api.CliqueBaseDonate_Out)
	OutCliqueBuilding_CliqueBuildingStatus        func(*clique_building_api.CliqueBuildingStatus_Out)
	OutCliqueBuilding_CliqueBankDonate            func(*clique_building_api.CliqueBankDonate_Out)
	OutCliqueBuilding_CliqueBankBuy               func(*clique_building_api.CliqueBankBuy_Out)
	OutCliqueBuilding_CliqueBankSold              func(*clique_building_api.CliqueBankSold_Out)
	OutCliqueBuilding_CliqueKongfuDonate          func(*clique_building_api.CliqueKongfuDonate_Out)
	OutCliqueBuilding_CliqueKongfuInfo            func(*clique_building_api.CliqueKongfuInfo_Out)
	OutCliqueBuilding_CliqueKongfuTrain           func(*clique_building_api.CliqueKongfuTrain_Out)
	OutCliqueBuilding_CliqueTempleWorship         func(*clique_building_api.CliqueTempleWorship_Out)
	OutCliqueBuilding_CliqueTempleDonate          func(*clique_building_api.CliqueTempleDonate_Out)
	OutCliqueBuilding_CliqueTempleInfo            func(*clique_building_api.CliqueTempleInfo_Out)
	OutCliqueBuilding_CliqueStoreDonate           func(*clique_building_api.CliqueStoreDonate_Out)
	OutCliqueBuilding_CliqueStoreInfo             func(*clique_building_api.CliqueStoreInfo_Out)
	OutCliqueBuilding_CliqueStoreSendChest        func(*clique_building_api.CliqueStoreSendChest_Out)
	OutCliqueQuest_GetCliqueDailyQuest            func(*clique_quest_api.GetCliqueDailyQuest_Out)
	OutCliqueQuest_AwardCliqueDailyQuest          func(*clique_quest_api.AwardCliqueDailyQuest_Out)
	OutCliqueQuest_NotifyCliqueDailyChange        func(*clique_quest_api.NotifyCliqueDailyChange_Out)
	OutCliqueQuest_GetCliqueBuildingQuest         func(*clique_quest_api.GetCliqueBuildingQuest_Out)
	OutCliqueQuest_AwardCliqueBuildingQuest       func(*clique_quest_api.AwardCliqueBuildingQuest_Out)
	OutCliqueEscort_EscortInfo                    func(*clique_escort_api.EscortInfo_Out)
	OutCliqueEscort_GetIngotBoat                  func(*clique_escort_api.GetIngotBoat_Out)
	OutCliqueEscort_StartEscort                   func(*clique_escort_api.StartEscort_Out)
	OutCliqueEscort_HijackBoat                    func(*clique_escort_api.HijackBoat_Out)
	OutCliqueEscort_RecoverBoat                   func(*clique_escort_api.RecoverBoat_Out)
	OutCliqueEscort_ListBoats                     func(*clique_escort_api.ListBoats_Out)
	OutCliqueEscort_GetRandomBoat                 func(*clique_escort_api.GetRandomBoat_Out)
	OutCliqueEscort_NotifyEscortFinished          func(*clique_escort_api.NotifyEscortFinished_Out)
	OutCliqueEscort_NotifyHijackFinished          func(*clique_escort_api.NotifyHijackFinished_Out)
	OutCliqueEscort_NotifyRecoverBattleWin        func(*clique_escort_api.NotifyRecoverBattleWin_Out)
	OutCliqueEscort_NotifyHijackBattleWin         func(*clique_escort_api.NotifyHijackBattleWin_Out)
	OutCliqueEscort_TakeHijackAward               func(*clique_escort_api.TakeHijackAward_Out)
	OutCliqueEscort_TakeEscortAward               func(*clique_escort_api.TakeEscortAward_Out)
	OutCliqueEscort_GetCliqueBoatMessages         func(*clique_escort_api.GetCliqueBoatMessages_Out)
	OutCliqueEscort_SendCliqueBoatMessage         func(*clique_escort_api.SendCliqueBoatMessage_Out)
	OutCliqueEscort_ReadCliqueBoatMessage         func(*clique_escort_api.ReadCliqueBoatMessage_Out)
	OutCliqueEscort_NotifyBoatStatusChange        func(*clique_escort_api.NotifyBoatStatusChange_Out)
	OutDespairLand_DespairLandInfo                func(*despair_land_api.DespairLandInfo_Out)
	OutDespairLand_DespairLandCampInfo            func(*despair_land_api.DespairLandCampInfo_Out)
	OutDespairLand_DespairLandCampPlayerInfo      func(*despair_land_api.DespairLandCampPlayerInfo_Out)
	OutDespairLand_DespairLandPickBox             func(*despair_land_api.DespairLandPickBox_Out)
	OutDespairLand_DespairLandWatchRecord         func(*despair_land_api.DespairLandWatchRecord_Out)
	OutDespairLand_DespairLandBuyBattleNum        func(*despair_land_api.DespairLandBuyBattleNum_Out)
	OutDespairLand_DespairLandBossInfo            func(*despair_land_api.DespairLandBossInfo_Out)
	OutDespairLand_DespairLandNotifyBossOpen      func(*despair_land_api.DespairLandNotifyBossOpen_Out)
	OutDespairLand_DespairLandNotifyBossDead      func(*despair_land_api.DespairLandNotifyBossDead_Out)
	OutDespairLand_DespairLandBuyBossBattleNum    func(*despair_land_api.DespairLandBuyBossBattleNum_Out)
	OutDespairLand_DespairLandNotifyPass          func(*despair_land_api.DespairLandNotifyPass_Out)
	OutDespairLand_DespairLandPickThreeStarBox    func(*despair_land_api.DespairLandPickThreeStarBox_Out)
	OutDespairLand_DespairLandBattleBossAwardInfo func(*despair_land_api.DespairLandBattleBossAwardInfo_Out)
	OutDespairLand_DespairLandBossOpenInfo        func(*despair_land_api.DespairLandBossOpenInfo_Out)
	OutDespairLand_DespairLandNotifyWeeklyRefresh func(*despair_land_api.DespairLandNotifyWeeklyRefresh_Out)
	OutAwaken_AwakenInfo                          func(*awaken_api.AwakenInfo_Out)
	OutAwaken_LevelupAttr                         func(*awaken_api.LevelupAttr_Out)
	OutTaoyuan_TaoyuanInfo                        func(*taoyuan_api.TaoyuanInfo_Out)
	OutTaoyuan_Bless                              func(*taoyuan_api.Bless_Out)
	OutTaoyuan_ShopBuy                            func(*taoyuan_api.ShopBuy_Out)
	OutTaoyuan_ShopSell                           func(*taoyuan_api.ShopSell_Out)
	OutTaoyuan_GetAllItems                        func(*taoyuan_api.GetAllItems_Out)
	OutTaoyuan_FiledsInfo                         func(*taoyuan_api.FiledsInfo_Out)
	OutTaoyuan_GrowPlant                          func(*taoyuan_api.GrowPlant_Out)
	OutTaoyuan_TakePlant                          func(*taoyuan_api.TakePlant_Out)
	OutTaoyuan_UpgradeFiled                       func(*taoyuan_api.UpgradeFiled_Out)
	OutTaoyuan_OpenFiled                          func(*taoyuan_api.OpenFiled_Out)
	OutTaoyuan_StudySkill                         func(*taoyuan_api.StudySkill_Out)
	OutTaoyuan_MakeProduct                        func(*taoyuan_api.MakeProduct_Out)
	OutTaoyuan_TakeProduct                        func(*taoyuan_api.TakeProduct_Out)
	OutTaoyuan_ProductMakeQueue                   func(*taoyuan_api.ProductMakeQueue_Out)
	OutTaoyuan_QuestInfo                          func(*taoyuan_api.QuestInfo_Out)
	OutTaoyuan_QuestFinish                        func(*taoyuan_api.QuestFinish_Out)
	OutTaoyuan_QuestRefresh                       func(*taoyuan_api.QuestRefresh_Out)
	OutTaoyuan_FriendTaoyuanInfo                  func(*taoyuan_api.FriendTaoyuanInfo_Out)
	OutTaoyuan_SkillInfo                          func(*taoyuan_api.SkillInfo_Out)
	OutTaoyuan_OpenQueue                          func(*taoyuan_api.OpenQueue_Out)
	OutTaoyuan_PlantQuicklyMaturity               func(*taoyuan_api.PlantQuicklyMaturity_Out)
	OutTaoyuan_TaoyuanMessageInfo                 func(*taoyuan_api.TaoyuanMessageInfo_Out)
	OutTaoyuan_TaoyuanMessageRead                 func(*taoyuan_api.TaoyuanMessageRead_Out)
	OutTaoyuan_OpenProductBuilding                func(*taoyuan_api.OpenProductBuilding_Out)
	OutDraw_GetHeartDrawInfo                      func(*draw_api.GetHeartDrawInfo_Out)
	OutDraw_HeartDraw                             func(*draw_api.HeartDraw_Out)
	OutDraw_GetChestInfo                          func(*draw_api.GetChestInfo_Out)
	OutDraw_DrawChest                             func(*draw_api.DrawChest_Out)
	OutDraw_HeartInfo                             func(*draw_api.HeartInfo_Out)
	OutDraw_GetFateBoxInfo                        func(*draw_api.GetFateBoxInfo_Out)
	OutDraw_OpenFateBox                           func(*draw_api.OpenFateBox_Out)
	OutDraw_ExchangeGiftCode                      func(*draw_api.ExchangeGiftCode_Out)
	OutServerInfo_GetVersion                      func(*server_info_api.GetVersion_Out)
	OutServerInfo_GetApiCount                     func(*server_info_api.GetApiCount_Out)
	OutServerInfo_SearchPlayerRole                func(*server_info_api.SearchPlayerRole_Out)
	OutServerInfo_UpdateAccessToken               func(*server_info_api.UpdateAccessToken_Out)
	OutServerInfo_UpdateEventData                 func(*server_info_api.UpdateEventData_Out)
	OutServerInfo_TssData                         func(*server_info_api.TssData_Out)
	OutDebug_AddBuddy                             func(*debug_api.AddBuddy_Out)
	OutDebug_AddItem                              func(*debug_api.AddItem_Out)
	OutDebug_SetRoleLevel                         func(*debug_api.SetRoleLevel_Out)
	OutDebug_SetCoins                             func(*debug_api.SetCoins_Out)
	OutDebug_SetIngot                             func(*debug_api.SetIngot_Out)
	OutDebug_AddGhost                             func(*debug_api.AddGhost_Out)
	OutDebug_SetPlayerPhysical                    func(*debug_api.SetPlayerPhysical_Out)
	OutDebug_ResetLevelEnterCount                 func(*debug_api.ResetLevelEnterCount_Out)
	OutDebug_AddExp                               func(*debug_api.AddExp_Out)
	OutDebug_OpenGhostMission                     func(*debug_api.OpenGhostMission_Out)
	OutDebug_SendMail                             func(*debug_api.SendMail_Out)
	OutDebug_ClearMail                            func(*debug_api.ClearMail_Out)
	OutDebug_OpenMissionLevel                     func(*debug_api.OpenMissionLevel_Out)
	OutDebug_StartBattle                          func(*debug_api.StartBattle_Out)
	OutDebug_ListenByName                         func(*debug_api.ListenByName_Out)
	OutDebug_OpenQuest                            func(*debug_api.OpenQuest_Out)
	OutDebug_OpenFunc                             func(*debug_api.OpenFunc_Out)
	OutDebug_AddSwordSoul                         func(*debug_api.AddSwordSoul_Out)
	OutDebug_AddBattlePet                         func(*debug_api.AddBattlePet_Out)
	OutDebug_ResetMultiLevelEnterCount            func(*debug_api.ResetMultiLevelEnterCount_Out)
	OutDebug_OpenMultiLevel                       func(*debug_api.OpenMultiLevel_Out)
	OutDebug_OpenAllPetGrid                       func(*debug_api.OpenAllPetGrid_Out)
	OutDebug_CreateAnnouncement                   func(*debug_api.CreateAnnouncement_Out)
	OutDebug_AddHeart                             func(*debug_api.AddHeart_Out)
	OutDebug_ResetHardLevelEnterCount             func(*debug_api.ResetHardLevelEnterCount_Out)
	OutDebug_OpenHardLevel                        func(*debug_api.OpenHardLevel_Out)
	OutDebug_SetVipLevel                          func(*debug_api.SetVipLevel_Out)
	OutDebug_SetResourceLevelOpenDay              func(*debug_api.SetResourceLevelOpenDay_Out)
	OutDebug_ResetResourceLevelOpenDay            func(*debug_api.ResetResourceLevelOpenDay_Out)
	OutDebug_ResetArenaDailyCount                 func(*debug_api.ResetArenaDailyCount_Out)
	OutDebug_ResetSwordSoulDrawCd                 func(*debug_api.ResetSwordSoulDrawCd_Out)
	OutDebug_SetFirstLoginTime                    func(*debug_api.SetFirstLoginTime_Out)
	OutDebug_EarlierFirstLoginTime                func(*debug_api.EarlierFirstLoginTime_Out)
	OutDebug_ResetServerOpenTime                  func(*debug_api.ResetServerOpenTime_Out)
	OutDebug_ClearTraderRefreshTime               func(*debug_api.ClearTraderRefreshTime_Out)
	OutDebug_AddTraderRefreshTime                 func(*debug_api.AddTraderRefreshTime_Out)
	OutDebug_ClearTraderSchedule                  func(*debug_api.ClearTraderSchedule_Out)
	OutDebug_AddTraderSchedule                    func(*debug_api.AddTraderSchedule_Out)
	OutDebug_OpenTown                             func(*debug_api.OpenTown_Out)
	OutDebug_AddGlobalMail                        func(*debug_api.AddGlobalMail_Out)
	OutDebug_CreateAnnouncementWithoutTpl         func(*debug_api.CreateAnnouncementWithoutTpl_Out)
	OutDebug_SetLoginDay                          func(*debug_api.SetLoginDay_Out)
	OutDebug_ResetLoginAward                      func(*debug_api.ResetLoginAward_Out)
	OutDebug_RestPlayerAwardLock                  func(*debug_api.RestPlayerAwardLock_Out)
	OutDebug_ResetRainbowLevel                    func(*debug_api.ResetRainbowLevel_Out)
	OutDebug_SetRainbowLevel                      func(*debug_api.SetRainbowLevel_Out)
	OutDebug_SendPushNotification                 func(*debug_api.SendPushNotification_Out)
	OutDebug_ResetPetVirtualEnv                   func(*debug_api.ResetPetVirtualEnv_Out)
	OutDebug_AddFame                              func(*debug_api.AddFame_Out)
	OutDebug_AddWorldChatMessage                  func(*debug_api.AddWorldChatMessage_Out)
	OutDebug_MonthCard                            func(*debug_api.MonthCard_Out)
	OutDebug_EnterSandbox                         func(*debug_api.EnterSandbox_Out)
	OutDebug_SandboxRollback                      func(*debug_api.SandboxRollback_Out)
	OutDebug_ExitSandbox                          func(*debug_api.ExitSandbox_Out)
	OutDebug_ResetShadedMissions                  func(*debug_api.ResetShadedMissions_Out)
	OutDebug_CleanCornucopia                      func(*debug_api.CleanCornucopia_Out)
	OutDebug_AddTotem                             func(*debug_api.AddTotem_Out)
	OutDebug_AddRune                              func(*debug_api.AddRune_Out)
	OutDebug_SendRareItemMessage                  func(*debug_api.SendRareItemMessage_Out)
	OutDebug_AddSwordDrivingAction                func(*debug_api.AddSwordDrivingAction_Out)
	OutDebug_ResetDrivingSwordData                func(*debug_api.ResetDrivingSwordData_Out)
	OutDebug_AddSwordSoulFragment                 func(*debug_api.AddSwordSoulFragment_Out)
	OutDebug_ResetMoneyTreeStatus                 func(*debug_api.ResetMoneyTreeStatus_Out)
	OutDebug_ResetTodayMoneyTree                  func(*debug_api.ResetTodayMoneyTree_Out)
	OutDebug_CleanSwordSoulIngotDrawNums          func(*debug_api.CleanSwordSoulIngotDrawNums_Out)
	OutDebug_PunchDrivingSwordCloud               func(*debug_api.PunchDrivingSwordCloud_Out)
	OutDebug_ClearCliqueDailyDonate               func(*debug_api.ClearCliqueDailyDonate_Out)
	OutDebug_SetCliqueContrib                     func(*debug_api.SetCliqueContrib_Out)
	OutDebug_RefreshCliqueWorship                 func(*debug_api.RefreshCliqueWorship_Out)
	OutDebug_CliqueEscortHijackBattleWin          func(*debug_api.CliqueEscortHijackBattleWin_Out)
	OutDebug_CliqueEscortRecoverBattleWin         func(*debug_api.CliqueEscortRecoverBattleWin_Out)
	OutDebug_CliqueEscortNotifyMessage            func(*debug_api.CliqueEscortNotifyMessage_Out)
	OutDebug_CliqueEscortNotifyDailyQuest         func(*debug_api.CliqueEscortNotifyDailyQuest_Out)
	OutDebug_SetCliqueBuildingLevel               func(*debug_api.SetCliqueBuildingLevel_Out)
	OutDebug_SetCliqueBuildingMoney               func(*debug_api.SetCliqueBuildingMoney_Out)
	OutDebug_EscortBench                          func(*debug_api.EscortBench_Out)
	OutDebug_ResetCliqueEscortDailyNum            func(*debug_api.ResetCliqueEscortDailyNum_Out)
	OutDebug_TakeAdditionQuest                    func(*debug_api.TakeAdditionQuest_Out)
	OutDebug_SetMissionStarMax                    func(*debug_api.SetMissionStarMax_Out)
	OutDebug_CliqueBankCd                         func(*debug_api.CliqueBankCd_Out)
	OutDebug_ResetDespairLandBattleNum            func(*debug_api.ResetDespairLandBattleNum_Out)
	OutDebug_ResetCliqueStoreSendTimes            func(*debug_api.ResetCliqueStoreSendTimes_Out)
	OutDebug_AddCliqueStoreDonate                 func(*debug_api.AddCliqueStoreDonate_Out)
	OutDebug_PassAllDespairLandLevel              func(*debug_api.PassAllDespairLandLevel_Out)
	OutDebug_DespairLandDummyBossKill             func(*debug_api.DespairLandDummyBossKill_Out)
	OutDebug_AddTaoyuanItem                       func(*debug_api.AddTaoyuanItem_Out)
	OutDebug_AddTaoyuanExp                        func(*debug_api.AddTaoyuanExp_Out)
}

func init() {

	player_api.SetOutHandler(OutPlayer{})
	town_api.SetOutHandler(OutTown{})
	team_api.SetOutHandler(OutTeam{})
	role_api.SetOutHandler(OutRole{})
	mission_api.SetOutHandler(OutMission{})
	skill_api.SetOutHandler(OutSkill{})
	battle_api.SetOutHandler(OutBattle{})
	item_api.SetOutHandler(OutItem{})
	notify_api.SetOutHandler(OutNotify{})
	ghost_api.SetOutHandler(OutGhost{})
	sword_soul_api.SetOutHandler(OutSwordSoul{})
	mail_api.SetOutHandler(OutMail{})
	quest_api.SetOutHandler(OutQuest{})
	friend_api.SetOutHandler(OutFriend{})
	tower_api.SetOutHandler(OutTower{})
	multi_level_api.SetOutHandler(OutMultiLevel{})
	battle_pet_api.SetOutHandler(OutBattlePet{})
	announcement_api.SetOutHandler(OutAnnouncement{})
	arena_api.SetOutHandler(OutArena{})
	vip_api.SetOutHandler(OutVip{})
	trader_api.SetOutHandler(OutTrader{})
	daily_sign_in_api.SetOutHandler(OutDailySignIn{})
	rainbow_api.SetOutHandler(OutRainbow{})
	event_api.SetOutHandler(OutEvent{})
	fashion_api.SetOutHandler(OutFashion{})
	push_notify_api.SetOutHandler(OutPushNotify{})
	meditation_api.SetOutHandler(OutMeditation{})
	pet_virtual_env_api.SetOutHandler(OutPetVirtualEnv{})
	channel_api.SetOutHandler(OutChannel{})
	driving_sword_api.SetOutHandler(OutDrivingSword{})
	totem_api.SetOutHandler(OutTotem{})
	money_tree_api.SetOutHandler(OutMoneyTree{})
	clique_api.SetOutHandler(OutClique{})
	clique_building_api.SetOutHandler(OutCliqueBuilding{})
	clique_quest_api.SetOutHandler(OutCliqueQuest{})
	clique_escort_api.SetOutHandler(OutCliqueEscort{})
	despair_land_api.SetOutHandler(OutDespairLand{})
	awaken_api.SetOutHandler(OutAwaken{})
	taoyuan_api.SetOutHandler(OutTaoyuan{})
	draw_api.SetOutHandler(OutDraw{})
	server_info_api.SetOutHandler(OutServerInfo{})
	debug_api.SetOutHandler(OutDebug{})

}

func NewClient(ip string) *Client {

	c := &Client{}

	var err error
	c.conn, err = net.Dial("tcp", ip, net.PacketN(8, net.LittleEndian))

	fail.When(err != nil, err)

	c.recv()

	return c

}

func (c *Client) Close() {

	c.closed = true
	c.conn.Close()
}

func (c *Client) RecordResponse(rsp interface{}) {}

func (c *Client) recv() {

	go func() {

		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Client %s recv() panic. %v\n", c.Name, err)
			}
		}()
		s := &net.Session{State: c}

		for {

			raw, err := c.conn.Read()

			if err != nil {
				if c.closed {
					return
				} else {
					panic(err)
				}
			}
			protocol.DecodeOut(raw).Process(s)

		}

	}()

}

type OutPlayer struct{}
type OutTown struct{}
type OutTeam struct{}
type OutRole struct{}
type OutMission struct{}
type OutSkill struct{}
type OutBattle struct{}
type OutItem struct{}
type OutNotify struct{}
type OutGhost struct{}
type OutSwordSoul struct{}
type OutMail struct{}
type OutQuest struct{}
type OutFriend struct{}
type OutTower struct{}
type OutMultiLevel struct{}
type OutBattlePet struct{}
type OutAnnouncement struct{}
type OutArena struct{}
type OutVip struct{}
type OutTrader struct{}
type OutDailySignIn struct{}
type OutRainbow struct{}
type OutEvent struct{}
type OutFashion struct{}
type OutPushNotify struct{}
type OutMeditation struct{}
type OutPetVirtualEnv struct{}
type OutChannel struct{}
type OutDrivingSword struct{}
type OutTotem struct{}
type OutMoneyTree struct{}
type OutClique struct{}
type OutCliqueBuilding struct{}
type OutCliqueQuest struct{}
type OutCliqueEscort struct{}
type OutDespairLand struct{}
type OutAwaken struct{}
type OutTaoyuan struct{}
type OutDraw struct{}
type OutServerInfo struct{}
type OutDebug struct{}

func (o OutPlayer) Info(session *net.Session, out *player_api.Info_Out) {

	if session.State.(*Client).OutPlayer_Info != nil {

		session.State.(*Client).OutPlayer_Info(out)

	}

}

func (o OutPlayer) Relogin(session *net.Session, out *player_api.Relogin_Out) {

	if session.State.(*Client).OutPlayer_Relogin != nil {

		session.State.(*Client).OutPlayer_Relogin(out)

	}

}

func (o OutPlayer) BuyPhysical(session *net.Session, out *player_api.BuyPhysical_Out) {

	if session.State.(*Client).OutPlayer_BuyPhysical != nil {

		session.State.(*Client).OutPlayer_BuyPhysical(out)

	}

}

func (o OutPlayer) SetPlayKey(session *net.Session, out *player_api.SetPlayKey_Out) {

	if session.State.(*Client).OutPlayer_SetPlayKey != nil {

		session.State.(*Client).OutPlayer_SetPlayKey(out)

	}

}

func (o OutPlayer) GetTime(session *net.Session, out *player_api.GetTime_Out) {

	if session.State.(*Client).OutPlayer_GetTime != nil {

		session.State.(*Client).OutPlayer_GetTime(out)

	}

}

func (o OutPlayer) FromPlatformLogin(session *net.Session, out *player_api.FromPlatformLogin_Out) {

	if session.State.(*Client).OutPlayer_FromPlatformLogin != nil {

		session.State.(*Client).OutPlayer_FromPlatformLogin(out)

	}

}

func (o OutPlayer) BuyCoins(session *net.Session, out *player_api.BuyCoins_Out) {

	if session.State.(*Client).OutPlayer_BuyCoins != nil {

		session.State.(*Client).OutPlayer_BuyCoins(out)

	}

}

func (o OutPlayer) GetLoginInfo(session *net.Session, out *player_api.GetLoginInfo_Out) {

	if session.State.(*Client).OutPlayer_GetLoginInfo != nil {

		session.State.(*Client).OutPlayer_GetLoginInfo(out)

	}

}

func (o OutPlayer) CrossLoginGameServer(session *net.Session, out *player_api.CrossLoginGameServer_Out) {

	if session.State.(*Client).OutPlayer_CrossLoginGameServer != nil {

		session.State.(*Client).OutPlayer_CrossLoginGameServer(out)

	}

}

func (o OutPlayer) NotifyGlobalServerInfo(session *net.Session, out *player_api.NotifyGlobalServerInfo_Out) {

	if session.State.(*Client).OutPlayer_NotifyGlobalServerInfo != nil {

		session.State.(*Client).OutPlayer_NotifyGlobalServerInfo(out)

	}

}

func (o OutPlayer) GlobalLogin(session *net.Session, out *player_api.GlobalLogin_Out) {

	if session.State.(*Client).OutPlayer_GlobalLogin != nil {

		session.State.(*Client).OutPlayer_GlobalLogin(out)

	}

}

func (o OutPlayer) GetIngot(session *net.Session, out *player_api.GetIngot_Out) {

	if session.State.(*Client).OutPlayer_GetIngot != nil {

		session.State.(*Client).OutPlayer_GetIngot(out)

	}

}

func (o OutPlayer) SystemFame(session *net.Session, out *player_api.SystemFame_Out) {

	if session.State.(*Client).OutPlayer_SystemFame != nil {

		session.State.(*Client).OutPlayer_SystemFame(out)

	}

}

func (o OutPlayer) GetRanks(session *net.Session, out *player_api.GetRanks_Out) {

	if session.State.(*Client).OutPlayer_GetRanks != nil {

		session.State.(*Client).OutPlayer_GetRanks(out)

	}

}

func (o OutTown) Enter(session *net.Session, out *town_api.Enter_Out) {

	if session.State.(*Client).OutTown_Enter != nil {

		session.State.(*Client).OutTown_Enter(out)

	}

}

func (o OutTown) Leave(session *net.Session, out *town_api.Leave_Out) {

	if session.State.(*Client).OutTown_Leave != nil {

		session.State.(*Client).OutTown_Leave(out)

	}

}

func (o OutTown) Move(session *net.Session, out *town_api.Move_Out) {

	if session.State.(*Client).OutTown_Move != nil {

		session.State.(*Client).OutTown_Move(out)

	}

}

func (o OutTown) TalkedNpcList(session *net.Session, out *town_api.TalkedNpcList_Out) {

	if session.State.(*Client).OutTown_TalkedNpcList != nil {

		session.State.(*Client).OutTown_TalkedNpcList(out)

	}

}

func (o OutTown) NpcTalk(session *net.Session, out *town_api.NpcTalk_Out) {

	if session.State.(*Client).OutTown_NpcTalk != nil {

		session.State.(*Client).OutTown_NpcTalk(out)

	}

}

func (o OutTown) NotifyTownPlayers(session *net.Session, out *town_api.NotifyTownPlayers_Out) {

	if session.State.(*Client).OutTown_NotifyTownPlayers != nil {

		session.State.(*Client).OutTown_NotifyTownPlayers(out)

	}

}

func (o OutTown) UpdateTownPlayer(session *net.Session, out *town_api.UpdateTownPlayer_Out) {

	if session.State.(*Client).OutTown_UpdateTownPlayer != nil {

		session.State.(*Client).OutTown_UpdateTownPlayer(out)

	}

}

func (o OutTown) UpdateTownPlayerMeditationState(session *net.Session, out *town_api.UpdateTownPlayerMeditationState_Out) {

	if session.State.(*Client).OutTown_UpdateTownPlayerMeditationState != nil {

		session.State.(*Client).OutTown_UpdateTownPlayerMeditationState(out)

	}

}

func (o OutTown) ListOpenedTownTreasures(session *net.Session, out *town_api.ListOpenedTownTreasures_Out) {

	if session.State.(*Client).OutTown_ListOpenedTownTreasures != nil {

		session.State.(*Client).OutTown_ListOpenedTownTreasures(out)

	}

}

func (o OutTown) TakeTownTreasures(session *net.Session, out *town_api.TakeTownTreasures_Out) {

	if session.State.(*Client).OutTown_TakeTownTreasures != nil {

		session.State.(*Client).OutTown_TakeTownTreasures(out)

	}

}

func (o OutTeam) GetFormationInfo(session *net.Session, out *team_api.GetFormationInfo_Out) {

	if session.State.(*Client).OutTeam_GetFormationInfo != nil {

		session.State.(*Client).OutTeam_GetFormationInfo(out)

	}

}

func (o OutTeam) UpFormation(session *net.Session, out *team_api.UpFormation_Out) {

	if session.State.(*Client).OutTeam_UpFormation != nil {

		session.State.(*Client).OutTeam_UpFormation(out)

	}

}

func (o OutTeam) DownFormation(session *net.Session, out *team_api.DownFormation_Out) {

	if session.State.(*Client).OutTeam_DownFormation != nil {

		session.State.(*Client).OutTeam_DownFormation(out)

	}

}

func (o OutTeam) SwapFormation(session *net.Session, out *team_api.SwapFormation_Out) {

	if session.State.(*Client).OutTeam_SwapFormation != nil {

		session.State.(*Client).OutTeam_SwapFormation(out)

	}

}

func (o OutTeam) ReplaceFormation(session *net.Session, out *team_api.ReplaceFormation_Out) {

	if session.State.(*Client).OutTeam_ReplaceFormation != nil {

		session.State.(*Client).OutTeam_ReplaceFormation(out)

	}

}

func (o OutTeam) TrainingTeamship(session *net.Session, out *team_api.TrainingTeamship_Out) {

	if session.State.(*Client).OutTeam_TrainingTeamship != nil {

		session.State.(*Client).OutTeam_TrainingTeamship(out)

	}

}

func (o OutRole) GetAllRoles(session *net.Session, out *role_api.GetAllRoles_Out) {

	if session.State.(*Client).OutRole_GetAllRoles != nil {

		session.State.(*Client).OutRole_GetAllRoles(out)

	}

}

func (o OutRole) GetRoleInfo(session *net.Session, out *role_api.GetRoleInfo_Out) {

	if session.State.(*Client).OutRole_GetRoleInfo != nil {

		session.State.(*Client).OutRole_GetRoleInfo(out)

	}

}

func (o OutRole) GetPlayerInfo(session *net.Session, out *role_api.GetPlayerInfo_Out) {

	if session.State.(*Client).OutRole_GetPlayerInfo != nil {

		session.State.(*Client).OutRole_GetPlayerInfo(out)

	}

}

func (o OutRole) GetFightNum(session *net.Session, out *role_api.GetFightNum_Out) {

	if session.State.(*Client).OutRole_GetFightNum != nil {

		session.State.(*Client).OutRole_GetFightNum(out)

	}

}

func (o OutRole) GetPlayerInfoWithOpenid(session *net.Session, out *role_api.GetPlayerInfoWithOpenid_Out) {

	if session.State.(*Client).OutRole_GetPlayerInfoWithOpenid != nil {

		session.State.(*Client).OutRole_GetPlayerInfoWithOpenid(out)

	}

}

func (o OutRole) LevelupRoleFriendship(session *net.Session, out *role_api.LevelupRoleFriendship_Out) {

	if session.State.(*Client).OutRole_LevelupRoleFriendship != nil {

		session.State.(*Client).OutRole_LevelupRoleFriendship(out)

	}

}

func (o OutRole) RecruitBuddy(session *net.Session, out *role_api.RecruitBuddy_Out) {

	if session.State.(*Client).OutRole_RecruitBuddy != nil {

		session.State.(*Client).OutRole_RecruitBuddy(out)

	}

}

func (o OutRole) GetRoleFightNum(session *net.Session, out *role_api.GetRoleFightNum_Out) {

	if session.State.(*Client).OutRole_GetRoleFightNum != nil {

		session.State.(*Client).OutRole_GetRoleFightNum(out)

	}

}

func (o OutRole) ChangeRoleStatus(session *net.Session, out *role_api.ChangeRoleStatus_Out) {

	if session.State.(*Client).OutRole_ChangeRoleStatus != nil {

		session.State.(*Client).OutRole_ChangeRoleStatus(out)

	}

}

func (o OutRole) GetInnRoleList(session *net.Session, out *role_api.GetInnRoleList_Out) {

	if session.State.(*Client).OutRole_GetInnRoleList != nil {

		session.State.(*Client).OutRole_GetInnRoleList(out)

	}

}

func (o OutRole) BuddyCoop(session *net.Session, out *role_api.BuddyCoop_Out) {

	if session.State.(*Client).OutRole_BuddyCoop != nil {

		session.State.(*Client).OutRole_BuddyCoop(out)

	}

}

func (o OutMission) Open(session *net.Session, out *mission_api.Open_Out) {

	if session.State.(*Client).OutMission_Open != nil {

		session.State.(*Client).OutMission_Open(out)

	}

}

func (o OutMission) GetMissionLevel(session *net.Session, out *mission_api.GetMissionLevel_Out) {

	if session.State.(*Client).OutMission_GetMissionLevel != nil {

		session.State.(*Client).OutMission_GetMissionLevel(out)

	}

}

func (o OutMission) EnterLevel(session *net.Session, out *mission_api.EnterLevel_Out) {

	if session.State.(*Client).OutMission_EnterLevel != nil {

		session.State.(*Client).OutMission_EnterLevel(out)

	}

}

func (o OutMission) OpenLevelBox(session *net.Session, out *mission_api.OpenLevelBox_Out) {

	if session.State.(*Client).OutMission_OpenLevelBox != nil {

		session.State.(*Client).OutMission_OpenLevelBox(out)

	}

}

func (o OutMission) UseIngotRelive(session *net.Session, out *mission_api.UseIngotRelive_Out) {

	if session.State.(*Client).OutMission_UseIngotRelive != nil {

		session.State.(*Client).OutMission_UseIngotRelive(out)

	}

}

func (o OutMission) UseItem(session *net.Session, out *mission_api.UseItem_Out) {

	if session.State.(*Client).OutMission_UseItem != nil {

		session.State.(*Client).OutMission_UseItem(out)

	}

}

func (o OutMission) Rebuild(session *net.Session, out *mission_api.Rebuild_Out) {

	if session.State.(*Client).OutMission_Rebuild != nil {

		session.State.(*Client).OutMission_Rebuild(out)

	}

}

func (o OutMission) EnterExtendLevel(session *net.Session, out *mission_api.EnterExtendLevel_Out) {

	if session.State.(*Client).OutMission_EnterExtendLevel != nil {

		session.State.(*Client).OutMission_EnterExtendLevel(out)

	}

}

func (o OutMission) GetExtendLevelInfo(session *net.Session, out *mission_api.GetExtendLevelInfo_Out) {

	if session.State.(*Client).OutMission_GetExtendLevelInfo != nil {

		session.State.(*Client).OutMission_GetExtendLevelInfo(out)

	}

}

func (o OutMission) OpenSmallBox(session *net.Session, out *mission_api.OpenSmallBox_Out) {

	if session.State.(*Client).OutMission_OpenSmallBox != nil {

		session.State.(*Client).OutMission_OpenSmallBox(out)

	}

}

func (o OutMission) NotifyCatchBattlePet(session *net.Session, out *mission_api.NotifyCatchBattlePet_Out) {

	if session.State.(*Client).OutMission_NotifyCatchBattlePet != nil {

		session.State.(*Client).OutMission_NotifyCatchBattlePet(out)

	}

}

func (o OutMission) EnterHardLevel(session *net.Session, out *mission_api.EnterHardLevel_Out) {

	if session.State.(*Client).OutMission_EnterHardLevel != nil {

		session.State.(*Client).OutMission_EnterHardLevel(out)

	}

}

func (o OutMission) GetHardLevel(session *net.Session, out *mission_api.GetHardLevel_Out) {

	if session.State.(*Client).OutMission_GetHardLevel != nil {

		session.State.(*Client).OutMission_GetHardLevel(out)

	}

}

func (o OutMission) GetBuddyLevelRoleId(session *net.Session, out *mission_api.GetBuddyLevelRoleId_Out) {

	if session.State.(*Client).OutMission_GetBuddyLevelRoleId != nil {

		session.State.(*Client).OutMission_GetBuddyLevelRoleId(out)

	}

}

func (o OutMission) SetBuddyLevelTeam(session *net.Session, out *mission_api.SetBuddyLevelTeam_Out) {

	if session.State.(*Client).OutMission_SetBuddyLevelTeam != nil {

		session.State.(*Client).OutMission_SetBuddyLevelTeam(out)

	}

}

func (o OutMission) AutoFightLevel(session *net.Session, out *mission_api.AutoFightLevel_Out) {

	if session.State.(*Client).OutMission_AutoFightLevel != nil {

		session.State.(*Client).OutMission_AutoFightLevel(out)

	}

}

func (o OutMission) EnterRainbowLevel(session *net.Session, out *mission_api.EnterRainbowLevel_Out) {

	if session.State.(*Client).OutMission_EnterRainbowLevel != nil {

		session.State.(*Client).OutMission_EnterRainbowLevel(out)

	}

}

func (o OutMission) OpenMengYao(session *net.Session, out *mission_api.OpenMengYao_Out) {

	if session.State.(*Client).OutMission_OpenMengYao != nil {

		session.State.(*Client).OutMission_OpenMengYao(out)

	}

}

func (o OutMission) GetMissionLevelByItemId(session *net.Session, out *mission_api.GetMissionLevelByItemId_Out) {

	if session.State.(*Client).OutMission_GetMissionLevelByItemId != nil {

		session.State.(*Client).OutMission_GetMissionLevelByItemId(out)

	}

}

func (o OutMission) BuyHardLevelTimes(session *net.Session, out *mission_api.BuyHardLevelTimes_Out) {

	if session.State.(*Client).OutMission_BuyHardLevelTimes != nil {

		session.State.(*Client).OutMission_BuyHardLevelTimes(out)

	}

}

func (o OutMission) OpenRandomAwardBox(session *net.Session, out *mission_api.OpenRandomAwardBox_Out) {

	if session.State.(*Client).OutMission_OpenRandomAwardBox != nil {

		session.State.(*Client).OutMission_OpenRandomAwardBox(out)

	}

}

func (o OutMission) EvaluateLevel(session *net.Session, out *mission_api.EvaluateLevel_Out) {

	if session.State.(*Client).OutMission_EvaluateLevel != nil {

		session.State.(*Client).OutMission_EvaluateLevel(out)

	}

}

func (o OutMission) OpenShadedBox(session *net.Session, out *mission_api.OpenShadedBox_Out) {

	if session.State.(*Client).OutMission_OpenShadedBox != nil {

		session.State.(*Client).OutMission_OpenShadedBox(out)

	}

}

func (o OutMission) GetMissionTotalStarInfo(session *net.Session, out *mission_api.GetMissionTotalStarInfo_Out) {

	if session.State.(*Client).OutMission_GetMissionTotalStarInfo != nil {

		session.State.(*Client).OutMission_GetMissionTotalStarInfo(out)

	}

}

func (o OutMission) GetMissionTotalStarAwards(session *net.Session, out *mission_api.GetMissionTotalStarAwards_Out) {

	if session.State.(*Client).OutMission_GetMissionTotalStarAwards != nil {

		session.State.(*Client).OutMission_GetMissionTotalStarAwards(out)

	}

}

func (o OutMission) ClearMissionLevelState(session *net.Session, out *mission_api.ClearMissionLevelState_Out) {

	if session.State.(*Client).OutMission_ClearMissionLevelState != nil {

		session.State.(*Client).OutMission_ClearMissionLevelState(out)

	}

}

func (o OutMission) BuyResourceMissionLevelTimes(session *net.Session, out *mission_api.BuyResourceMissionLevelTimes_Out) {

	if session.State.(*Client).OutMission_BuyResourceMissionLevelTimes != nil {

		session.State.(*Client).OutMission_BuyResourceMissionLevelTimes(out)

	}

}

func (o OutMission) GetDragonBall(session *net.Session, out *mission_api.GetDragonBall_Out) {

	if session.State.(*Client).OutMission_GetDragonBall != nil {

		session.State.(*Client).OutMission_GetDragonBall(out)

	}

}

func (o OutMission) GetEventItem(session *net.Session, out *mission_api.GetEventItem_Out) {

	if session.State.(*Client).OutMission_GetEventItem != nil {

		session.State.(*Client).OutMission_GetEventItem(out)

	}

}

func (o OutSkill) GetAllSkillsInfo(session *net.Session, out *skill_api.GetAllSkillsInfo_Out) {

	if session.State.(*Client).OutSkill_GetAllSkillsInfo != nil {

		session.State.(*Client).OutSkill_GetAllSkillsInfo(out)

	}

}

func (o OutSkill) EquipSkill(session *net.Session, out *skill_api.EquipSkill_Out) {

	if session.State.(*Client).OutSkill_EquipSkill != nil {

		session.State.(*Client).OutSkill_EquipSkill(out)

	}

}

func (o OutSkill) UnequipSkill(session *net.Session, out *skill_api.UnequipSkill_Out) {

	if session.State.(*Client).OutSkill_UnequipSkill != nil {

		session.State.(*Client).OutSkill_UnequipSkill(out)

	}

}

func (o OutSkill) StudySkillByCheat(session *net.Session, out *skill_api.StudySkillByCheat_Out) {

	if session.State.(*Client).OutSkill_StudySkillByCheat != nil {

		session.State.(*Client).OutSkill_StudySkillByCheat(out)

	}

}

func (o OutSkill) TrainSkill(session *net.Session, out *skill_api.TrainSkill_Out) {

	if session.State.(*Client).OutSkill_TrainSkill != nil {

		session.State.(*Client).OutSkill_TrainSkill(out)

	}

}

func (o OutSkill) FlushSkill(session *net.Session, out *skill_api.FlushSkill_Out) {

	if session.State.(*Client).OutSkill_FlushSkill != nil {

		session.State.(*Client).OutSkill_FlushSkill(out)

	}

}

func (o OutBattle) StartBattle(session *net.Session, out *battle_api.StartBattle_Out) {

	if session.State.(*Client).OutBattle_StartBattle != nil {

		session.State.(*Client).OutBattle_StartBattle(out)

	}

}

func (o OutBattle) NextRound(session *net.Session, out *battle_api.NextRound_Out) {

	if session.State.(*Client).OutBattle_NextRound != nil {

		session.State.(*Client).OutBattle_NextRound(out)

	}

}

func (o OutBattle) End(session *net.Session, out *battle_api.End_Out) {

	if session.State.(*Client).OutBattle_End != nil {

		session.State.(*Client).OutBattle_End(out)

	}

}

func (o OutBattle) Escape(session *net.Session, out *battle_api.Escape_Out) {

	if session.State.(*Client).OutBattle_Escape != nil {

		session.State.(*Client).OutBattle_Escape(out)

	}

}

func (o OutBattle) Fightnum(session *net.Session, out *battle_api.Fightnum_Out) {

	if session.State.(*Client).OutBattle_Fightnum != nil {

		session.State.(*Client).OutBattle_Fightnum(out)

	}

}

func (o OutBattle) StartReadyTimeout(session *net.Session, out *battle_api.StartReadyTimeout_Out) {

	if session.State.(*Client).OutBattle_StartReadyTimeout != nil {

		session.State.(*Client).OutBattle_StartReadyTimeout(out)

	}

}

func (o OutBattle) StartReady(session *net.Session, out *battle_api.StartReady_Out) {

	if session.State.(*Client).OutBattle_StartReady != nil {

		session.State.(*Client).OutBattle_StartReady(out)

	}

}

func (o OutBattle) StateChange(session *net.Session, out *battle_api.StateChange_Out) {

	if session.State.(*Client).OutBattle_StateChange != nil {

		session.State.(*Client).OutBattle_StateChange(out)

	}

}

func (o OutBattle) CallBattlePet(session *net.Session, out *battle_api.CallBattlePet_Out) {

	if session.State.(*Client).OutBattle_CallBattlePet != nil {

		session.State.(*Client).OutBattle_CallBattlePet(out)

	}

}

func (o OutBattle) UseBuddySkill(session *net.Session, out *battle_api.UseBuddySkill_Out) {

	if session.State.(*Client).OutBattle_UseBuddySkill != nil {

		session.State.(*Client).OutBattle_UseBuddySkill(out)

	}

}

func (o OutBattle) CallNewEnemys(session *net.Session, out *battle_api.CallNewEnemys_Out) {

	if session.State.(*Client).OutBattle_CallNewEnemys != nil {

		session.State.(*Client).OutBattle_CallNewEnemys(out)

	}

}

func (o OutBattle) NewFighterGroup(session *net.Session, out *battle_api.NewFighterGroup_Out) {

	if session.State.(*Client).OutBattle_NewFighterGroup != nil {

		session.State.(*Client).OutBattle_NewFighterGroup(out)

	}

}

func (o OutBattle) StartBattleForHijackBoat(session *net.Session, out *battle_api.StartBattleForHijackBoat_Out) {

	if session.State.(*Client).OutBattle_StartBattleForHijackBoat != nil {

		session.State.(*Client).OutBattle_StartBattleForHijackBoat(out)

	}

}

func (o OutBattle) StartBattleForRecoverBoat(session *net.Session, out *battle_api.StartBattleForRecoverBoat_Out) {

	if session.State.(*Client).OutBattle_StartBattleForRecoverBoat != nil {

		session.State.(*Client).OutBattle_StartBattleForRecoverBoat(out)

	}

}

func (o OutBattle) RoundReady(session *net.Session, out *battle_api.RoundReady_Out) {

	if session.State.(*Client).OutBattle_RoundReady != nil {

		session.State.(*Client).OutBattle_RoundReady(out)

	}

}

func (o OutBattle) InitRound(session *net.Session, out *battle_api.InitRound_Out) {

	if session.State.(*Client).OutBattle_InitRound != nil {

		session.State.(*Client).OutBattle_InitRound(out)

	}

}

func (o OutBattle) SetAuto(session *net.Session, out *battle_api.SetAuto_Out) {

	if session.State.(*Client).OutBattle_SetAuto != nil {

		session.State.(*Client).OutBattle_SetAuto(out)

	}

}

func (o OutBattle) CancelAuto(session *net.Session, out *battle_api.CancelAuto_Out) {

	if session.State.(*Client).OutBattle_CancelAuto != nil {

		session.State.(*Client).OutBattle_CancelAuto(out)

	}

}

func (o OutBattle) SetSkill(session *net.Session, out *battle_api.SetSkill_Out) {

	if session.State.(*Client).OutBattle_SetSkill != nil {

		session.State.(*Client).OutBattle_SetSkill(out)

	}

}

func (o OutBattle) UseItem(session *net.Session, out *battle_api.UseItem_Out) {

	if session.State.(*Client).OutBattle_UseItem != nil {

		session.State.(*Client).OutBattle_UseItem(out)

	}

}

func (o OutBattle) UseGhost(session *net.Session, out *battle_api.UseGhost_Out) {

	if session.State.(*Client).OutBattle_UseGhost != nil {

		session.State.(*Client).OutBattle_UseGhost(out)

	}

}

func (o OutBattle) NotifyReady(session *net.Session, out *battle_api.NotifyReady_Out) {

	if session.State.(*Client).OutBattle_NotifyReady != nil {

		session.State.(*Client).OutBattle_NotifyReady(out)

	}

}

func (o OutBattle) BattleReconnect(session *net.Session, out *battle_api.BattleReconnect_Out) {

	if session.State.(*Client).OutBattle_BattleReconnect != nil {

		session.State.(*Client).OutBattle_BattleReconnect(out)

	}

}

func (o OutItem) GetAllItems(session *net.Session, out *item_api.GetAllItems_Out) {

	if session.State.(*Client).OutItem_GetAllItems != nil {

		session.State.(*Client).OutItem_GetAllItems(out)

	}

}

func (o OutItem) DropItem(session *net.Session, out *item_api.DropItem_Out) {

	if session.State.(*Client).OutItem_DropItem != nil {

		session.State.(*Client).OutItem_DropItem(out)

	}

}

func (o OutItem) BuyItem(session *net.Session, out *item_api.BuyItem_Out) {

	if session.State.(*Client).OutItem_BuyItem != nil {

		session.State.(*Client).OutItem_BuyItem(out)

	}

}

func (o OutItem) SellItem(session *net.Session, out *item_api.SellItem_Out) {

	if session.State.(*Client).OutItem_SellItem != nil {

		session.State.(*Client).OutItem_SellItem(out)

	}

}

func (o OutItem) Dress(session *net.Session, out *item_api.Dress_Out) {

	if session.State.(*Client).OutItem_Dress != nil {

		session.State.(*Client).OutItem_Dress(out)

	}

}

func (o OutItem) Undress(session *net.Session, out *item_api.Undress_Out) {

	if session.State.(*Client).OutItem_Undress != nil {

		session.State.(*Client).OutItem_Undress(out)

	}

}

func (o OutItem) BuyItemBack(session *net.Session, out *item_api.BuyItemBack_Out) {

	if session.State.(*Client).OutItem_BuyItemBack != nil {

		session.State.(*Client).OutItem_BuyItemBack(out)

	}

}

func (o OutItem) IsBagFull(session *net.Session, out *item_api.IsBagFull_Out) {

	if session.State.(*Client).OutItem_IsBagFull != nil {

		session.State.(*Client).OutItem_IsBagFull(out)

	}

}

func (o OutItem) Decompose(session *net.Session, out *item_api.Decompose_Out) {

	if session.State.(*Client).OutItem_Decompose != nil {

		session.State.(*Client).OutItem_Decompose(out)

	}

}

func (o OutItem) Refine(session *net.Session, out *item_api.Refine_Out) {

	if session.State.(*Client).OutItem_Refine != nil {

		session.State.(*Client).OutItem_Refine(out)

	}

}

func (o OutItem) GetRecastInfo(session *net.Session, out *item_api.GetRecastInfo_Out) {

	if session.State.(*Client).OutItem_GetRecastInfo != nil {

		session.State.(*Client).OutItem_GetRecastInfo(out)

	}

}

func (o OutItem) Recast(session *net.Session, out *item_api.Recast_Out) {

	if session.State.(*Client).OutItem_Recast != nil {

		session.State.(*Client).OutItem_Recast(out)

	}

}

func (o OutItem) UseItem(session *net.Session, out *item_api.UseItem_Out) {

	if session.State.(*Client).OutItem_UseItem != nil {

		session.State.(*Client).OutItem_UseItem(out)

	}

}

func (o OutItem) RoleUseCostItem(session *net.Session, out *item_api.RoleUseCostItem_Out) {

	if session.State.(*Client).OutItem_RoleUseCostItem != nil {

		session.State.(*Client).OutItem_RoleUseCostItem(out)

	}

}

func (o OutItem) BatchUseItem(session *net.Session, out *item_api.BatchUseItem_Out) {

	if session.State.(*Client).OutItem_BatchUseItem != nil {

		session.State.(*Client).OutItem_BatchUseItem(out)

	}

}

func (o OutItem) DragonBallExchangeNotify(session *net.Session, out *item_api.DragonBallExchangeNotify_Out) {

	if session.State.(*Client).OutItem_DragonBallExchangeNotify != nil {

		session.State.(*Client).OutItem_DragonBallExchangeNotify(out)

	}

}

func (o OutItem) OpenCornucopia(session *net.Session, out *item_api.OpenCornucopia_Out) {

	if session.State.(*Client).OutItem_OpenCornucopia != nil {

		session.State.(*Client).OutItem_OpenCornucopia(out)

	}

}

func (o OutItem) GetSealedbooks(session *net.Session, out *item_api.GetSealedbooks_Out) {

	if session.State.(*Client).OutItem_GetSealedbooks != nil {

		session.State.(*Client).OutItem_GetSealedbooks(out)

	}

}

func (o OutItem) ActivationSealedbook(session *net.Session, out *item_api.ActivationSealedbook_Out) {

	if session.State.(*Client).OutItem_ActivationSealedbook != nil {

		session.State.(*Client).OutItem_ActivationSealedbook(out)

	}

}

func (o OutItem) ExchangeGhostCrystal(session *net.Session, out *item_api.ExchangeGhostCrystal_Out) {

	if session.State.(*Client).OutItem_ExchangeGhostCrystal != nil {

		session.State.(*Client).OutItem_ExchangeGhostCrystal(out)

	}

}

func (o OutItem) ExchangeShopItem(session *net.Session, out *item_api.ExchangeShopItem_Out) {

	if session.State.(*Client).OutItem_ExchangeShopItem != nil {

		session.State.(*Client).OutItem_ExchangeShopItem(out)

	}

}

func (o OutItem) GetHoildayItemList(session *net.Session, out *item_api.GetHoildayItemList_Out) {

	if session.State.(*Client).OutItem_GetHoildayItemList != nil {

		session.State.(*Client).OutItem_GetHoildayItemList(out)

	}

}

func (o OutItem) ExchangeHoildayItem(session *net.Session, out *item_api.ExchangeHoildayItem_Out) {

	if session.State.(*Client).OutItem_ExchangeHoildayItem != nil {

		session.State.(*Client).OutItem_ExchangeHoildayItem(out)

	}

}

func (o OutItem) BatchExchangeDragonBall(session *net.Session, out *item_api.BatchExchangeDragonBall_Out) {

	if session.State.(*Client).OutItem_BatchExchangeDragonBall != nil {

		session.State.(*Client).OutItem_BatchExchangeDragonBall(out)

	}

}

func (o OutNotify) PlayerKeyChanged(session *net.Session, out *notify_api.PlayerKeyChanged_Out) {

	if session.State.(*Client).OutNotify_PlayerKeyChanged != nil {

		session.State.(*Client).OutNotify_PlayerKeyChanged(out)

	}

}

func (o OutNotify) MissionLevelLockChanged(session *net.Session, out *notify_api.MissionLevelLockChanged_Out) {

	if session.State.(*Client).OutNotify_MissionLevelLockChanged != nil {

		session.State.(*Client).OutNotify_MissionLevelLockChanged(out)

	}

}

func (o OutNotify) RoleExpChange(session *net.Session, out *notify_api.RoleExpChange_Out) {

	if session.State.(*Client).OutNotify_RoleExpChange != nil {

		session.State.(*Client).OutNotify_RoleExpChange(out)

	}

}

func (o OutNotify) PhysicalChange(session *net.Session, out *notify_api.PhysicalChange_Out) {

	if session.State.(*Client).OutNotify_PhysicalChange != nil {

		session.State.(*Client).OutNotify_PhysicalChange(out)

	}

}

func (o OutNotify) MoneyChange(session *net.Session, out *notify_api.MoneyChange_Out) {

	if session.State.(*Client).OutNotify_MoneyChange != nil {

		session.State.(*Client).OutNotify_MoneyChange(out)

	}

}

func (o OutNotify) SkillAdd(session *net.Session, out *notify_api.SkillAdd_Out) {

	if session.State.(*Client).OutNotify_SkillAdd != nil {

		session.State.(*Client).OutNotify_SkillAdd(out)

	}

}

func (o OutNotify) ItemChange(session *net.Session, out *notify_api.ItemChange_Out) {

	if session.State.(*Client).OutNotify_ItemChange != nil {

		session.State.(*Client).OutNotify_ItemChange(out)

	}

}

func (o OutNotify) RoleBattleStatusChange(session *net.Session, out *notify_api.RoleBattleStatusChange_Out) {

	if session.State.(*Client).OutNotify_RoleBattleStatusChange != nil {

		session.State.(*Client).OutNotify_RoleBattleStatusChange(out)

	}

}

func (o OutNotify) NewMail(session *net.Session, out *notify_api.NewMail_Out) {

	if session.State.(*Client).OutNotify_NewMail != nil {

		session.State.(*Client).OutNotify_NewMail(out)

	}

}

func (o OutNotify) HeartChange(session *net.Session, out *notify_api.HeartChange_Out) {

	if session.State.(*Client).OutNotify_HeartChange != nil {

		session.State.(*Client).OutNotify_HeartChange(out)

	}

}

func (o OutNotify) QuestChange(session *net.Session, out *notify_api.QuestChange_Out) {

	if session.State.(*Client).OutNotify_QuestChange != nil {

		session.State.(*Client).OutNotify_QuestChange(out)

	}

}

func (o OutNotify) TownLockChange(session *net.Session, out *notify_api.TownLockChange_Out) {

	if session.State.(*Client).OutNotify_TownLockChange != nil {

		session.State.(*Client).OutNotify_TownLockChange(out)

	}

}

func (o OutNotify) Chat(session *net.Session, out *notify_api.Chat_Out) {

	if session.State.(*Client).OutNotify_Chat != nil {

		session.State.(*Client).OutNotify_Chat(out)

	}

}

func (o OutNotify) FuncKeyChange(session *net.Session, out *notify_api.FuncKeyChange_Out) {

	if session.State.(*Client).OutNotify_FuncKeyChange != nil {

		session.State.(*Client).OutNotify_FuncKeyChange(out)

	}

}

func (o OutNotify) ItemRecastStateRebuild(session *net.Session, out *notify_api.ItemRecastStateRebuild_Out) {

	if session.State.(*Client).OutNotify_ItemRecastStateRebuild != nil {

		session.State.(*Client).OutNotify_ItemRecastStateRebuild(out)

	}

}

func (o OutNotify) SendAnnouncement(session *net.Session, out *notify_api.SendAnnouncement_Out) {

	if session.State.(*Client).OutNotify_SendAnnouncement != nil {

		session.State.(*Client).OutNotify_SendAnnouncement(out)

	}

}

func (o OutNotify) VipLevelChange(session *net.Session, out *notify_api.VipLevelChange_Out) {

	if session.State.(*Client).OutNotify_VipLevelChange != nil {

		session.State.(*Client).OutNotify_VipLevelChange(out)

	}

}

func (o OutNotify) NotifyNewBuddy(session *net.Session, out *notify_api.NotifyNewBuddy_Out) {

	if session.State.(*Client).OutNotify_NotifyNewBuddy != nil {

		session.State.(*Client).OutNotify_NotifyNewBuddy(out)

	}

}

func (o OutNotify) HardLevelLockChanged(session *net.Session, out *notify_api.HardLevelLockChanged_Out) {

	if session.State.(*Client).OutNotify_HardLevelLockChanged != nil {

		session.State.(*Client).OutNotify_HardLevelLockChanged(out)

	}

}

func (o OutNotify) SendSwordSoulDrawNumChange(session *net.Session, out *notify_api.SendSwordSoulDrawNumChange_Out) {

	if session.State.(*Client).OutNotify_SendSwordSoulDrawNumChange != nil {

		session.State.(*Client).OutNotify_SendSwordSoulDrawNumChange(out)

	}

}

func (o OutNotify) SendHaveNewGhost(session *net.Session, out *notify_api.SendHaveNewGhost_Out) {

	if session.State.(*Client).OutNotify_SendHaveNewGhost != nil {

		session.State.(*Client).OutNotify_SendHaveNewGhost(out)

	}

}

func (o OutNotify) SendHeartRecoverTime(session *net.Session, out *notify_api.SendHeartRecoverTime_Out) {

	if session.State.(*Client).OutNotify_SendHeartRecoverTime != nil {

		session.State.(*Client).OutNotify_SendHeartRecoverTime(out)

	}

}

func (o OutNotify) SendGlobalMail(session *net.Session, out *notify_api.SendGlobalMail_Out) {

	if session.State.(*Client).OutNotify_SendGlobalMail != nil {

		session.State.(*Client).OutNotify_SendGlobalMail(out)

	}

}

func (o OutNotify) SendPhysicalRecoverTime(session *net.Session, out *notify_api.SendPhysicalRecoverTime_Out) {

	if session.State.(*Client).OutNotify_SendPhysicalRecoverTime != nil {

		session.State.(*Client).OutNotify_SendPhysicalRecoverTime(out)

	}

}

func (o OutNotify) SendFashionChange(session *net.Session, out *notify_api.SendFashionChange_Out) {

	if session.State.(*Client).OutNotify_SendFashionChange != nil {

		session.State.(*Client).OutNotify_SendFashionChange(out)

	}

}

func (o OutNotify) TransError(session *net.Session, out *notify_api.TransError_Out) {

	if session.State.(*Client).OutNotify_TransError != nil {

		session.State.(*Client).OutNotify_TransError(out)

	}

}

func (o OutNotify) SendEventCenterChange(session *net.Session, out *notify_api.SendEventCenterChange_Out) {

	if session.State.(*Client).OutNotify_SendEventCenterChange != nil {

		session.State.(*Client).OutNotify_SendEventCenterChange(out)

	}

}

func (o OutNotify) MeditationState(session *net.Session, out *notify_api.MeditationState_Out) {

	if session.State.(*Client).OutNotify_MeditationState != nil {

		session.State.(*Client).OutNotify_MeditationState(out)

	}

}

func (o OutNotify) DeleteAnnouncement(session *net.Session, out *notify_api.DeleteAnnouncement_Out) {

	if session.State.(*Client).OutNotify_DeleteAnnouncement != nil {

		session.State.(*Client).OutNotify_DeleteAnnouncement(out)

	}

}

func (o OutNotify) SendHaveNewPet(session *net.Session, out *notify_api.SendHaveNewPet_Out) {

	if session.State.(*Client).OutNotify_SendHaveNewPet != nil {

		session.State.(*Client).OutNotify_SendHaveNewPet(out)

	}

}

func (o OutNotify) SendLogout(session *net.Session, out *notify_api.SendLogout_Out) {

	if session.State.(*Client).OutNotify_SendLogout != nil {

		session.State.(*Client).OutNotify_SendLogout(out)

	}

}

func (o OutNotify) FameChange(session *net.Session, out *notify_api.FameChange_Out) {

	if session.State.(*Client).OutNotify_FameChange != nil {

		session.State.(*Client).OutNotify_FameChange(out)

	}

}

func (o OutNotify) NotifyMonthCardOpen(session *net.Session, out *notify_api.NotifyMonthCardOpen_Out) {

	if session.State.(*Client).OutNotify_NotifyMonthCardOpen != nil {

		session.State.(*Client).OutNotify_NotifyMonthCardOpen(out)

	}

}

func (o OutNotify) NotifyMonthCardRenewal(session *net.Session, out *notify_api.NotifyMonthCardRenewal_Out) {

	if session.State.(*Client).OutNotify_NotifyMonthCardRenewal != nil {

		session.State.(*Client).OutNotify_NotifyMonthCardRenewal(out)

	}

}

func (o OutNotify) NotifyNewTotem(session *net.Session, out *notify_api.NotifyNewTotem_Out) {

	if session.State.(*Client).OutNotify_NotifyNewTotem != nil {

		session.State.(*Client).OutNotify_NotifyNewTotem(out)

	}

}

func (o OutNotify) NotifyRuneChange(session *net.Session, out *notify_api.NotifyRuneChange_Out) {

	if session.State.(*Client).OutNotify_NotifyRuneChange != nil {

		session.State.(*Client).OutNotify_NotifyRuneChange(out)

	}

}

func (o OutNotify) TaoyuanItemChange(session *net.Session, out *notify_api.TaoyuanItemChange_Out) {

	if session.State.(*Client).OutNotify_TaoyuanItemChange != nil {

		session.State.(*Client).OutNotify_TaoyuanItemChange(out)

	}

}

func (o OutNotify) TaoyuanMessageRefresh(session *net.Session, out *notify_api.TaoyuanMessageRefresh_Out) {

	if session.State.(*Client).OutNotify_TaoyuanMessageRefresh != nil {

		session.State.(*Client).OutNotify_TaoyuanMessageRefresh(out)

	}

}

func (o OutNotify) TaoyuanQuestCanFinish(session *net.Session, out *notify_api.TaoyuanQuestCanFinish_Out) {

	if session.State.(*Client).OutNotify_TaoyuanQuestCanFinish != nil {

		session.State.(*Client).OutNotify_TaoyuanQuestCanFinish(out)

	}

}

func (o OutNotify) TaoyuanExpRefresh(session *net.Session, out *notify_api.TaoyuanExpRefresh_Out) {

	if session.State.(*Client).OutNotify_TaoyuanExpRefresh != nil {

		session.State.(*Client).OutNotify_TaoyuanExpRefresh(out)

	}

}

func (o OutGhost) Info(session *net.Session, out *ghost_api.Info_Out) {

	if session.State.(*Client).OutGhost_Info != nil {

		session.State.(*Client).OutGhost_Info(out)

	}

}

func (o OutGhost) Equip(session *net.Session, out *ghost_api.Equip_Out) {

	if session.State.(*Client).OutGhost_Equip != nil {

		session.State.(*Client).OutGhost_Equip(out)

	}

}

func (o OutGhost) Unequip(session *net.Session, out *ghost_api.Unequip_Out) {

	if session.State.(*Client).OutGhost_Unequip != nil {

		session.State.(*Client).OutGhost_Unequip(out)

	}

}

func (o OutGhost) Swap(session *net.Session, out *ghost_api.Swap_Out) {

	if session.State.(*Client).OutGhost_Swap != nil {

		session.State.(*Client).OutGhost_Swap(out)

	}

}

func (o OutGhost) EquipPosChange(session *net.Session, out *ghost_api.EquipPosChange_Out) {

	if session.State.(*Client).OutGhost_EquipPosChange != nil {

		session.State.(*Client).OutGhost_EquipPosChange(out)

	}

}

func (o OutGhost) Train(session *net.Session, out *ghost_api.Train_Out) {

	if session.State.(*Client).OutGhost_Train != nil {

		session.State.(*Client).OutGhost_Train(out)

	}

}

func (o OutGhost) Upgrade(session *net.Session, out *ghost_api.Upgrade_Out) {

	if session.State.(*Client).OutGhost_Upgrade != nil {

		session.State.(*Client).OutGhost_Upgrade(out)

	}

}

func (o OutGhost) Baptize(session *net.Session, out *ghost_api.Baptize_Out) {

	if session.State.(*Client).OutGhost_Baptize != nil {

		session.State.(*Client).OutGhost_Baptize(out)

	}

}

func (o OutGhost) Compose(session *net.Session, out *ghost_api.Compose_Out) {

	if session.State.(*Client).OutGhost_Compose != nil {

		session.State.(*Client).OutGhost_Compose(out)

	}

}

func (o OutGhost) TrainSkill(session *net.Session, out *ghost_api.TrainSkill_Out) {

	if session.State.(*Client).OutGhost_TrainSkill != nil {

		session.State.(*Client).OutGhost_TrainSkill(out)

	}

}

func (o OutGhost) FlushAttr(session *net.Session, out *ghost_api.FlushAttr_Out) {

	if session.State.(*Client).OutGhost_FlushAttr != nil {

		session.State.(*Client).OutGhost_FlushAttr(out)

	}

}

func (o OutGhost) RelationGhost(session *net.Session, out *ghost_api.RelationGhost_Out) {

	if session.State.(*Client).OutGhost_RelationGhost != nil {

		session.State.(*Client).OutGhost_RelationGhost(out)

	}

}

func (o OutSwordSoul) Info(session *net.Session, out *sword_soul_api.Info_Out) {

	if session.State.(*Client).OutSwordSoul_Info != nil {

		session.State.(*Client).OutSwordSoul_Info(out)

	}

}

func (o OutSwordSoul) Draw(session *net.Session, out *sword_soul_api.Draw_Out) {

	if session.State.(*Client).OutSwordSoul_Draw != nil {

		session.State.(*Client).OutSwordSoul_Draw(out)

	}

}

func (o OutSwordSoul) Upgrade(session *net.Session, out *sword_soul_api.Upgrade_Out) {

	if session.State.(*Client).OutSwordSoul_Upgrade != nil {

		session.State.(*Client).OutSwordSoul_Upgrade(out)

	}

}

func (o OutSwordSoul) Exchange(session *net.Session, out *sword_soul_api.Exchange_Out) {

	if session.State.(*Client).OutSwordSoul_Exchange != nil {

		session.State.(*Client).OutSwordSoul_Exchange(out)

	}

}

func (o OutSwordSoul) Equip(session *net.Session, out *sword_soul_api.Equip_Out) {

	if session.State.(*Client).OutSwordSoul_Equip != nil {

		session.State.(*Client).OutSwordSoul_Equip(out)

	}

}

func (o OutSwordSoul) Unequip(session *net.Session, out *sword_soul_api.Unequip_Out) {

	if session.State.(*Client).OutSwordSoul_Unequip != nil {

		session.State.(*Client).OutSwordSoul_Unequip(out)

	}

}

func (o OutSwordSoul) EquipPosChange(session *net.Session, out *sword_soul_api.EquipPosChange_Out) {

	if session.State.(*Client).OutSwordSoul_EquipPosChange != nil {

		session.State.(*Client).OutSwordSoul_EquipPosChange(out)

	}

}

func (o OutSwordSoul) Swap(session *net.Session, out *sword_soul_api.Swap_Out) {

	if session.State.(*Client).OutSwordSoul_Swap != nil {

		session.State.(*Client).OutSwordSoul_Swap(out)

	}

}

func (o OutSwordSoul) IsBagFull(session *net.Session, out *sword_soul_api.IsBagFull_Out) {

	if session.State.(*Client).OutSwordSoul_IsBagFull != nil {

		session.State.(*Client).OutSwordSoul_IsBagFull(out)

	}

}

func (o OutSwordSoul) EmptyPosNum(session *net.Session, out *sword_soul_api.EmptyPosNum_Out) {

	if session.State.(*Client).OutSwordSoul_EmptyPosNum != nil {

		session.State.(*Client).OutSwordSoul_EmptyPosNum(out)

	}

}

func (o OutMail) GetList(session *net.Session, out *mail_api.GetList_Out) {

	if session.State.(*Client).OutMail_GetList != nil {

		session.State.(*Client).OutMail_GetList(out)

	}

}

func (o OutMail) Read(session *net.Session, out *mail_api.Read_Out) {

	if session.State.(*Client).OutMail_Read != nil {

		session.State.(*Client).OutMail_Read(out)

	}

}

func (o OutMail) TakeAttachment(session *net.Session, out *mail_api.TakeAttachment_Out) {

	if session.State.(*Client).OutMail_TakeAttachment != nil {

		session.State.(*Client).OutMail_TakeAttachment(out)

	}

}

func (o OutMail) GetInfos(session *net.Session, out *mail_api.GetInfos_Out) {

	if session.State.(*Client).OutMail_GetInfos != nil {

		session.State.(*Client).OutMail_GetInfos(out)

	}

}

func (o OutMail) RequestGlobalMail(session *net.Session, out *mail_api.RequestGlobalMail_Out) {

	if session.State.(*Client).OutMail_RequestGlobalMail != nil {

		session.State.(*Client).OutMail_RequestGlobalMail(out)

	}

}

func (o OutQuest) UpdateQuest(session *net.Session, out *quest_api.UpdateQuest_Out) {

	if session.State.(*Client).OutQuest_UpdateQuest != nil {

		session.State.(*Client).OutQuest_UpdateQuest(out)

	}

}

func (o OutQuest) GetDailyInfo(session *net.Session, out *quest_api.GetDailyInfo_Out) {

	if session.State.(*Client).OutQuest_GetDailyInfo != nil {

		session.State.(*Client).OutQuest_GetDailyInfo(out)

	}

}

func (o OutQuest) AwardDaily(session *net.Session, out *quest_api.AwardDaily_Out) {

	if session.State.(*Client).OutQuest_AwardDaily != nil {

		session.State.(*Client).OutQuest_AwardDaily(out)

	}

}

func (o OutQuest) NotifyDailyChange(session *net.Session, out *quest_api.NotifyDailyChange_Out) {

	if session.State.(*Client).OutQuest_NotifyDailyChange != nil {

		session.State.(*Client).OutQuest_NotifyDailyChange(out)

	}

}

func (o OutQuest) Guide(session *net.Session, out *quest_api.Guide_Out) {

	if session.State.(*Client).OutQuest_Guide != nil {

		session.State.(*Client).OutQuest_Guide(out)

	}

}

func (o OutQuest) GetExtendQuestInfoByNpcId(session *net.Session, out *quest_api.GetExtendQuestInfoByNpcId_Out) {

	if session.State.(*Client).OutQuest_GetExtendQuestInfoByNpcId != nil {

		session.State.(*Client).OutQuest_GetExtendQuestInfoByNpcId(out)

	}

}

func (o OutQuest) TakeExtendQuestAward(session *net.Session, out *quest_api.TakeExtendQuestAward_Out) {

	if session.State.(*Client).OutQuest_TakeExtendQuestAward != nil {

		session.State.(*Client).OutQuest_TakeExtendQuestAward(out)

	}

}

func (o OutQuest) GetPannelQuestInfo(session *net.Session, out *quest_api.GetPannelQuestInfo_Out) {

	if session.State.(*Client).OutQuest_GetPannelQuestInfo != nil {

		session.State.(*Client).OutQuest_GetPannelQuestInfo(out)

	}

}

func (o OutQuest) GiveUpAdditionQuest(session *net.Session, out *quest_api.GiveUpAdditionQuest_Out) {

	if session.State.(*Client).OutQuest_GiveUpAdditionQuest != nil {

		session.State.(*Client).OutQuest_GiveUpAdditionQuest(out)

	}

}

func (o OutQuest) TakeAdditionQuest(session *net.Session, out *quest_api.TakeAdditionQuest_Out) {

	if session.State.(*Client).OutQuest_TakeAdditionQuest != nil {

		session.State.(*Client).OutQuest_TakeAdditionQuest(out)

	}

}

func (o OutQuest) TakeAdditionQuestAward(session *net.Session, out *quest_api.TakeAdditionQuestAward_Out) {

	if session.State.(*Client).OutQuest_TakeAdditionQuestAward != nil {

		session.State.(*Client).OutQuest_TakeAdditionQuestAward(out)

	}

}

func (o OutQuest) GetAdditionQuest(session *net.Session, out *quest_api.GetAdditionQuest_Out) {

	if session.State.(*Client).OutQuest_GetAdditionQuest != nil {

		session.State.(*Client).OutQuest_GetAdditionQuest(out)

	}

}

func (o OutQuest) RefreshAdditionQuest(session *net.Session, out *quest_api.RefreshAdditionQuest_Out) {

	if session.State.(*Client).OutQuest_RefreshAdditionQuest != nil {

		session.State.(*Client).OutQuest_RefreshAdditionQuest(out)

	}

}

func (o OutQuest) TakeQuestStarsAwaded(session *net.Session, out *quest_api.TakeQuestStarsAwaded_Out) {

	if session.State.(*Client).OutQuest_TakeQuestStarsAwaded != nil {

		session.State.(*Client).OutQuest_TakeQuestStarsAwaded(out)

	}

}

func (o OutFriend) GetFriendList(session *net.Session, out *friend_api.GetFriendList_Out) {

	if session.State.(*Client).OutFriend_GetFriendList != nil {

		session.State.(*Client).OutFriend_GetFriendList(out)

	}

}

func (o OutFriend) ListenByNick(session *net.Session, out *friend_api.ListenByNick_Out) {

	if session.State.(*Client).OutFriend_ListenByNick != nil {

		session.State.(*Client).OutFriend_ListenByNick(out)

	}

}

func (o OutFriend) CancelListen(session *net.Session, out *friend_api.CancelListen_Out) {

	if session.State.(*Client).OutFriend_CancelListen != nil {

		session.State.(*Client).OutFriend_CancelListen(out)

	}

}

func (o OutFriend) SendHeart(session *net.Session, out *friend_api.SendHeart_Out) {

	if session.State.(*Client).OutFriend_SendHeart != nil {

		session.State.(*Client).OutFriend_SendHeart(out)

	}

}

func (o OutFriend) Chat(session *net.Session, out *friend_api.Chat_Out) {

	if session.State.(*Client).OutFriend_Chat != nil {

		session.State.(*Client).OutFriend_Chat(out)

	}

}

func (o OutFriend) GetChatHistory(session *net.Session, out *friend_api.GetChatHistory_Out) {

	if session.State.(*Client).OutFriend_GetChatHistory != nil {

		session.State.(*Client).OutFriend_GetChatHistory(out)

	}

}

func (o OutFriend) GetOfflineMessages(session *net.Session, out *friend_api.GetOfflineMessages_Out) {

	if session.State.(*Client).OutFriend_GetOfflineMessages != nil {

		session.State.(*Client).OutFriend_GetOfflineMessages(out)

	}

}

func (o OutFriend) Block(session *net.Session, out *friend_api.Block_Out) {

	if session.State.(*Client).OutFriend_Block != nil {

		session.State.(*Client).OutFriend_Block(out)

	}

}

func (o OutFriend) CancelBlock(session *net.Session, out *friend_api.CancelBlock_Out) {

	if session.State.(*Client).OutFriend_CancelBlock != nil {

		session.State.(*Client).OutFriend_CancelBlock(out)

	}

}

func (o OutFriend) CleanChatHistory(session *net.Session, out *friend_api.CleanChatHistory_Out) {

	if session.State.(*Client).OutFriend_CleanChatHistory != nil {

		session.State.(*Client).OutFriend_CleanChatHistory(out)

	}

}

func (o OutFriend) NotifyListenedState(session *net.Session, out *friend_api.NotifyListenedState_Out) {

	if session.State.(*Client).OutFriend_NotifyListenedState != nil {

		session.State.(*Client).OutFriend_NotifyListenedState(out)

	}

}

func (o OutFriend) CurrentPlatformFriendNum(session *net.Session, out *friend_api.CurrentPlatformFriendNum_Out) {

	if session.State.(*Client).OutFriend_CurrentPlatformFriendNum != nil {

		session.State.(*Client).OutFriend_CurrentPlatformFriendNum(out)

	}

}

func (o OutFriend) GetSendHeartList(session *net.Session, out *friend_api.GetSendHeartList_Out) {

	if session.State.(*Client).OutFriend_GetSendHeartList != nil {

		session.State.(*Client).OutFriend_GetSendHeartList(out)

	}

}

func (o OutFriend) SendHeartToAllFriends(session *net.Session, out *friend_api.SendHeartToAllFriends_Out) {

	if session.State.(*Client).OutFriend_SendHeartToAllFriends != nil {

		session.State.(*Client).OutFriend_SendHeartToAllFriends(out)

	}

}

func (o OutTower) GetInfo(session *net.Session, out *tower_api.GetInfo_Out) {

	if session.State.(*Client).OutTower_GetInfo != nil {

		session.State.(*Client).OutTower_GetInfo(out)

	}

}

func (o OutTower) UseLadder(session *net.Session, out *tower_api.UseLadder_Out) {

	if session.State.(*Client).OutTower_UseLadder != nil {

		session.State.(*Client).OutTower_UseLadder(out)

	}

}

func (o OutMultiLevel) CreateRoom(session *net.Session, out *multi_level_api.CreateRoom_Out) {

	if session.State.(*Client).OutMultiLevel_CreateRoom != nil {

		session.State.(*Client).OutMultiLevel_CreateRoom(out)

	}

}

func (o OutMultiLevel) EnterRoom(session *net.Session, out *multi_level_api.EnterRoom_Out) {

	if session.State.(*Client).OutMultiLevel_EnterRoom != nil {

		session.State.(*Client).OutMultiLevel_EnterRoom(out)

	}

}

func (o OutMultiLevel) NotifyRoomInfo(session *net.Session, out *multi_level_api.NotifyRoomInfo_Out) {

	if session.State.(*Client).OutMultiLevel_NotifyRoomInfo != nil {

		session.State.(*Client).OutMultiLevel_NotifyRoomInfo(out)

	}

}

func (o OutMultiLevel) LeaveRoom(session *net.Session, out *multi_level_api.LeaveRoom_Out) {

	if session.State.(*Client).OutMultiLevel_LeaveRoom != nil {

		session.State.(*Client).OutMultiLevel_LeaveRoom(out)

	}

}

func (o OutMultiLevel) NotifyJoinPartner(session *net.Session, out *multi_level_api.NotifyJoinPartner_Out) {

	if session.State.(*Client).OutMultiLevel_NotifyJoinPartner != nil {

		session.State.(*Client).OutMultiLevel_NotifyJoinPartner(out)

	}

}

func (o OutMultiLevel) ChangeBuddy(session *net.Session, out *multi_level_api.ChangeBuddy_Out) {

	if session.State.(*Client).OutMultiLevel_ChangeBuddy != nil {

		session.State.(*Client).OutMultiLevel_ChangeBuddy(out)

	}

}

func (o OutMultiLevel) GetBaseInfo(session *net.Session, out *multi_level_api.GetBaseInfo_Out) {

	if session.State.(*Client).OutMultiLevel_GetBaseInfo != nil {

		session.State.(*Client).OutMultiLevel_GetBaseInfo(out)

	}

}

func (o OutMultiLevel) GetOnlineFriend(session *net.Session, out *multi_level_api.GetOnlineFriend_Out) {

	if session.State.(*Client).OutMultiLevel_GetOnlineFriend != nil {

		session.State.(*Client).OutMultiLevel_GetOnlineFriend(out)

	}

}

func (o OutMultiLevel) InviteIntoRoom(session *net.Session, out *multi_level_api.InviteIntoRoom_Out) {

	if session.State.(*Client).OutMultiLevel_InviteIntoRoom != nil {

		session.State.(*Client).OutMultiLevel_InviteIntoRoom(out)

	}

}

func (o OutMultiLevel) NotifyRoomInvite(session *net.Session, out *multi_level_api.NotifyRoomInvite_Out) {

	if session.State.(*Client).OutMultiLevel_NotifyRoomInvite != nil {

		session.State.(*Client).OutMultiLevel_NotifyRoomInvite(out)

	}

}

func (o OutMultiLevel) NotifyPlayersInfo(session *net.Session, out *multi_level_api.NotifyPlayersInfo_Out) {

	if session.State.(*Client).OutMultiLevel_NotifyPlayersInfo != nil {

		session.State.(*Client).OutMultiLevel_NotifyPlayersInfo(out)

	}

}

func (o OutMultiLevel) RefuseRoomInvite(session *net.Session, out *multi_level_api.RefuseRoomInvite_Out) {

	if session.State.(*Client).OutMultiLevel_RefuseRoomInvite != nil {

		session.State.(*Client).OutMultiLevel_RefuseRoomInvite(out)

	}

}

func (o OutMultiLevel) NotifyRoomInviteRefuse(session *net.Session, out *multi_level_api.NotifyRoomInviteRefuse_Out) {

	if session.State.(*Client).OutMultiLevel_NotifyRoomInviteRefuse != nil {

		session.State.(*Client).OutMultiLevel_NotifyRoomInviteRefuse(out)

	}

}

func (o OutMultiLevel) NotifyUpdatePartner(session *net.Session, out *multi_level_api.NotifyUpdatePartner_Out) {

	if session.State.(*Client).OutMultiLevel_NotifyUpdatePartner != nil {

		session.State.(*Client).OutMultiLevel_NotifyUpdatePartner(out)

	}

}

func (o OutMultiLevel) MatchPlayer(session *net.Session, out *multi_level_api.MatchPlayer_Out) {

	if session.State.(*Client).OutMultiLevel_MatchPlayer != nil {

		session.State.(*Client).OutMultiLevel_MatchPlayer(out)

	}

}

func (o OutMultiLevel) NotifyMatchPlayerSuccess(session *net.Session, out *multi_level_api.NotifyMatchPlayerSuccess_Out) {

	if session.State.(*Client).OutMultiLevel_NotifyMatchPlayerSuccess != nil {

		session.State.(*Client).OutMultiLevel_NotifyMatchPlayerSuccess(out)

	}

}

func (o OutMultiLevel) MatchRoom(session *net.Session, out *multi_level_api.MatchRoom_Out) {

	if session.State.(*Client).OutMultiLevel_MatchRoom != nil {

		session.State.(*Client).OutMultiLevel_MatchRoom(out)

	}

}

func (o OutMultiLevel) CancelMatchRoom(session *net.Session, out *multi_level_api.CancelMatchRoom_Out) {

	if session.State.(*Client).OutMultiLevel_CancelMatchRoom != nil {

		session.State.(*Client).OutMultiLevel_CancelMatchRoom(out)

	}

}

func (o OutMultiLevel) GetFormInfo(session *net.Session, out *multi_level_api.GetFormInfo_Out) {

	if session.State.(*Client).OutMultiLevel_GetFormInfo != nil {

		session.State.(*Client).OutMultiLevel_GetFormInfo(out)

	}

}

func (o OutMultiLevel) SetForm(session *net.Session, out *multi_level_api.SetForm_Out) {

	if session.State.(*Client).OutMultiLevel_SetForm != nil {

		session.State.(*Client).OutMultiLevel_SetForm(out)

	}

}

func (o OutMultiLevel) GetBattleRoleInfo(session *net.Session, out *multi_level_api.GetBattleRoleInfo_Out) {

	if session.State.(*Client).OutMultiLevel_GetBattleRoleInfo != nil {

		session.State.(*Client).OutMultiLevel_GetBattleRoleInfo(out)

	}

}

func (o OutMultiLevel) NotifyMatchPlayerInfo(session *net.Session, out *multi_level_api.NotifyMatchPlayerInfo_Out) {

	if session.State.(*Client).OutMultiLevel_NotifyMatchPlayerInfo != nil {

		session.State.(*Client).OutMultiLevel_NotifyMatchPlayerInfo(out)

	}

}

func (o OutMultiLevel) GetMatchInfo(session *net.Session, out *multi_level_api.GetMatchInfo_Out) {

	if session.State.(*Client).OutMultiLevel_GetMatchInfo != nil {

		session.State.(*Client).OutMultiLevel_GetMatchInfo(out)

	}

}

func (o OutMultiLevel) NotifyFormInfo(session *net.Session, out *multi_level_api.NotifyFormInfo_Out) {

	if session.State.(*Client).OutMultiLevel_NotifyFormInfo != nil {

		session.State.(*Client).OutMultiLevel_NotifyFormInfo(out)

	}

}

func (o OutMultiLevel) CancelMatchPlayer(session *net.Session, out *multi_level_api.CancelMatchPlayer_Out) {

	if session.State.(*Client).OutMultiLevel_CancelMatchPlayer != nil {

		session.State.(*Client).OutMultiLevel_CancelMatchPlayer(out)

	}

}

func (o OutBattlePet) GetPetInfo(session *net.Session, out *battle_pet_api.GetPetInfo_Out) {

	if session.State.(*Client).OutBattlePet_GetPetInfo != nil {

		session.State.(*Client).OutBattlePet_GetPetInfo(out)

	}

}

func (o OutBattlePet) SetPet(session *net.Session, out *battle_pet_api.SetPet_Out) {

	if session.State.(*Client).OutBattlePet_SetPet != nil {

		session.State.(*Client).OutBattlePet_SetPet(out)

	}

}

func (o OutBattlePet) SetPetSwap(session *net.Session, out *battle_pet_api.SetPetSwap_Out) {

	if session.State.(*Client).OutBattlePet_SetPetSwap != nil {

		session.State.(*Client).OutBattlePet_SetPetSwap(out)

	}

}

func (o OutBattlePet) UpgradePet(session *net.Session, out *battle_pet_api.UpgradePet_Out) {

	if session.State.(*Client).OutBattlePet_UpgradePet != nil {

		session.State.(*Client).OutBattlePet_UpgradePet(out)

	}

}

func (o OutBattlePet) TrainingPetSkill(session *net.Session, out *battle_pet_api.TrainingPetSkill_Out) {

	if session.State.(*Client).OutBattlePet_TrainingPetSkill != nil {

		session.State.(*Client).OutBattlePet_TrainingPetSkill(out)

	}

}

func (o OutAnnouncement) GetList(session *net.Session, out *announcement_api.GetList_Out) {

	if session.State.(*Client).OutAnnouncement_GetList != nil {

		session.State.(*Client).OutAnnouncement_GetList(out)

	}

}

func (o OutArena) Enter(session *net.Session, out *arena_api.Enter_Out) {

	if session.State.(*Client).OutArena_Enter != nil {

		session.State.(*Client).OutArena_Enter(out)

	}

}

func (o OutArena) GetRecords(session *net.Session, out *arena_api.GetRecords_Out) {

	if session.State.(*Client).OutArena_GetRecords != nil {

		session.State.(*Client).OutArena_GetRecords(out)

	}

}

func (o OutArena) AwardBox(session *net.Session, out *arena_api.AwardBox_Out) {

	if session.State.(*Client).OutArena_AwardBox != nil {

		session.State.(*Client).OutArena_AwardBox(out)

	}

}

func (o OutArena) NotifyFailedCdTime(session *net.Session, out *arena_api.NotifyFailedCdTime_Out) {

	if session.State.(*Client).OutArena_NotifyFailedCdTime != nil {

		session.State.(*Client).OutArena_NotifyFailedCdTime(out)

	}

}

func (o OutArena) StartBattle(session *net.Session, out *arena_api.StartBattle_Out) {

	if session.State.(*Client).OutArena_StartBattle != nil {

		session.State.(*Client).OutArena_StartBattle(out)

	}

}

func (o OutArena) NotifyLoseRank(session *net.Session, out *arena_api.NotifyLoseRank_Out) {

	if session.State.(*Client).OutArena_NotifyLoseRank != nil {

		session.State.(*Client).OutArena_NotifyLoseRank(out)

	}

}

func (o OutArena) NotifyArenaInfo(session *net.Session, out *arena_api.NotifyArenaInfo_Out) {

	if session.State.(*Client).OutArena_NotifyArenaInfo != nil {

		session.State.(*Client).OutArena_NotifyArenaInfo(out)

	}

}

func (o OutArena) GetTopRank(session *net.Session, out *arena_api.GetTopRank_Out) {

	if session.State.(*Client).OutArena_GetTopRank != nil {

		session.State.(*Client).OutArena_GetTopRank(out)

	}

}

func (o OutArena) CleanFailedCdTime(session *net.Session, out *arena_api.CleanFailedCdTime_Out) {

	if session.State.(*Client).OutArena_CleanFailedCdTime != nil {

		session.State.(*Client).OutArena_CleanFailedCdTime(out)

	}

}

func (o OutVip) Info(session *net.Session, out *vip_api.Info_Out) {

	if session.State.(*Client).OutVip_Info != nil {

		session.State.(*Client).OutVip_Info(out)

	}

}

func (o OutVip) GetTotal(session *net.Session, out *vip_api.GetTotal_Out) {

	if session.State.(*Client).OutVip_GetTotal != nil {

		session.State.(*Client).OutVip_GetTotal(out)

	}

}

func (o OutVip) VipLevelTotal(session *net.Session, out *vip_api.VipLevelTotal_Out) {

	if session.State.(*Client).OutVip_VipLevelTotal != nil {

		session.State.(*Client).OutVip_VipLevelTotal(out)

	}

}

func (o OutVip) BuyTimes(session *net.Session, out *vip_api.BuyTimes_Out) {

	if session.State.(*Client).OutVip_BuyTimes != nil {

		session.State.(*Client).OutVip_BuyTimes(out)

	}

}

func (o OutTrader) Info(session *net.Session, out *trader_api.Info_Out) {

	if session.State.(*Client).OutTrader_Info != nil {

		session.State.(*Client).OutTrader_Info(out)

	}

}

func (o OutTrader) StoreState(session *net.Session, out *trader_api.StoreState_Out) {

	if session.State.(*Client).OutTrader_StoreState != nil {

		session.State.(*Client).OutTrader_StoreState(out)

	}

}

func (o OutTrader) Buy(session *net.Session, out *trader_api.Buy_Out) {

	if session.State.(*Client).OutTrader_Buy != nil {

		session.State.(*Client).OutTrader_Buy(out)

	}

}

func (o OutTrader) Refresh(session *net.Session, out *trader_api.Refresh_Out) {

	if session.State.(*Client).OutTrader_Refresh != nil {

		session.State.(*Client).OutTrader_Refresh(out)

	}

}

func (o OutDailySignIn) Info(session *net.Session, out *daily_sign_in_api.Info_Out) {

	if session.State.(*Client).OutDailySignIn_Info != nil {

		session.State.(*Client).OutDailySignIn_Info(out)

	}

}

func (o OutDailySignIn) Sign(session *net.Session, out *daily_sign_in_api.Sign_Out) {

	if session.State.(*Client).OutDailySignIn_Sign != nil {

		session.State.(*Client).OutDailySignIn_Sign(out)

	}

}

func (o OutDailySignIn) SignPastDay(session *net.Session, out *daily_sign_in_api.SignPastDay_Out) {

	if session.State.(*Client).OutDailySignIn_SignPastDay != nil {

		session.State.(*Client).OutDailySignIn_SignPastDay(out)

	}

}

func (o OutRainbow) Info(session *net.Session, out *rainbow_api.Info_Out) {

	if session.State.(*Client).OutRainbow_Info != nil {

		session.State.(*Client).OutRainbow_Info(out)

	}

}

func (o OutRainbow) Reset(session *net.Session, out *rainbow_api.Reset_Out) {

	if session.State.(*Client).OutRainbow_Reset != nil {

		session.State.(*Client).OutRainbow_Reset(out)

	}

}

func (o OutRainbow) AwardInfo(session *net.Session, out *rainbow_api.AwardInfo_Out) {

	if session.State.(*Client).OutRainbow_AwardInfo != nil {

		session.State.(*Client).OutRainbow_AwardInfo(out)

	}

}

func (o OutRainbow) TakeAward(session *net.Session, out *rainbow_api.TakeAward_Out) {

	if session.State.(*Client).OutRainbow_TakeAward != nil {

		session.State.(*Client).OutRainbow_TakeAward(out)

	}

}

func (o OutRainbow) JumpToSegment(session *net.Session, out *rainbow_api.JumpToSegment_Out) {

	if session.State.(*Client).OutRainbow_JumpToSegment != nil {

		session.State.(*Client).OutRainbow_JumpToSegment(out)

	}

}

func (o OutRainbow) AutoFight(session *net.Session, out *rainbow_api.AutoFight_Out) {

	if session.State.(*Client).OutRainbow_AutoFight != nil {

		session.State.(*Client).OutRainbow_AutoFight(out)

	}

}

func (o OutEvent) LoginAwardInfo(session *net.Session, out *event_api.LoginAwardInfo_Out) {

	if session.State.(*Client).OutEvent_LoginAwardInfo != nil {

		session.State.(*Client).OutEvent_LoginAwardInfo(out)

	}

}

func (o OutEvent) TakeLoginAward(session *net.Session, out *event_api.TakeLoginAward_Out) {

	if session.State.(*Client).OutEvent_TakeLoginAward != nil {

		session.State.(*Client).OutEvent_TakeLoginAward(out)

	}

}

func (o OutEvent) GetEvents(session *net.Session, out *event_api.GetEvents_Out) {

	if session.State.(*Client).OutEvent_GetEvents != nil {

		session.State.(*Client).OutEvent_GetEvents(out)

	}

}

func (o OutEvent) GetEventAward(session *net.Session, out *event_api.GetEventAward_Out) {

	if session.State.(*Client).OutEvent_GetEventAward != nil {

		session.State.(*Client).OutEvent_GetEventAward(out)

	}

}

func (o OutEvent) PlayerEventPhysicalCost(session *net.Session, out *event_api.PlayerEventPhysicalCost_Out) {

	if session.State.(*Client).OutEvent_PlayerEventPhysicalCost != nil {

		session.State.(*Client).OutEvent_PlayerEventPhysicalCost(out)

	}

}

func (o OutEvent) PlayerMonthCardInfo(session *net.Session, out *event_api.PlayerMonthCardInfo_Out) {

	if session.State.(*Client).OutEvent_PlayerMonthCardInfo != nil {

		session.State.(*Client).OutEvent_PlayerMonthCardInfo(out)

	}

}

func (o OutEvent) GetSevenInfo(session *net.Session, out *event_api.GetSevenInfo_Out) {

	if session.State.(*Client).OutEvent_GetSevenInfo != nil {

		session.State.(*Client).OutEvent_GetSevenInfo(out)

	}

}

func (o OutEvent) GetSevenAward(session *net.Session, out *event_api.GetSevenAward_Out) {

	if session.State.(*Client).OutEvent_GetSevenAward != nil {

		session.State.(*Client).OutEvent_GetSevenAward(out)

	}

}

func (o OutEvent) GetRichmanClubInfo(session *net.Session, out *event_api.GetRichmanClubInfo_Out) {

	if session.State.(*Client).OutEvent_GetRichmanClubInfo != nil {

		session.State.(*Client).OutEvent_GetRichmanClubInfo(out)

	}

}

func (o OutEvent) GetRichmanClubAward(session *net.Session, out *event_api.GetRichmanClubAward_Out) {

	if session.State.(*Client).OutEvent_GetRichmanClubAward != nil {

		session.State.(*Client).OutEvent_GetRichmanClubAward(out)

	}

}

func (o OutEvent) InfoShare(session *net.Session, out *event_api.InfoShare_Out) {

	if session.State.(*Client).OutEvent_InfoShare != nil {

		session.State.(*Client).OutEvent_InfoShare(out)

	}

}

func (o OutEvent) InfoGroupBuy(session *net.Session, out *event_api.InfoGroupBuy_Out) {

	if session.State.(*Client).OutEvent_InfoGroupBuy != nil {

		session.State.(*Client).OutEvent_InfoGroupBuy(out)

	}

}

func (o OutEvent) GetIngotChangeTotal(session *net.Session, out *event_api.GetIngotChangeTotal_Out) {

	if session.State.(*Client).OutEvent_GetIngotChangeTotal != nil {

		session.State.(*Client).OutEvent_GetIngotChangeTotal(out)

	}

}

func (o OutEvent) GetEventTotalAward(session *net.Session, out *event_api.GetEventTotalAward_Out) {

	if session.State.(*Client).OutEvent_GetEventTotalAward != nil {

		session.State.(*Client).OutEvent_GetEventTotalAward(out)

	}

}

func (o OutEvent) GetEventArenaRank(session *net.Session, out *event_api.GetEventArenaRank_Out) {

	if session.State.(*Client).OutEvent_GetEventArenaRank != nil {

		session.State.(*Client).OutEvent_GetEventArenaRank(out)

	}

}

func (o OutEvent) GetEventTenDrawTimes(session *net.Session, out *event_api.GetEventTenDrawTimes_Out) {

	if session.State.(*Client).OutEvent_GetEventTenDrawTimes != nil {

		session.State.(*Client).OutEvent_GetEventTenDrawTimes(out)

	}

}

func (o OutEvent) GetEventRechargeAward(session *net.Session, out *event_api.GetEventRechargeAward_Out) {

	if session.State.(*Client).OutEvent_GetEventRechargeAward != nil {

		session.State.(*Client).OutEvent_GetEventRechargeAward(out)

	}

}

func (o OutEvent) GetEventNewYear(session *net.Session, out *event_api.GetEventNewYear_Out) {

	if session.State.(*Client).OutEvent_GetEventNewYear != nil {

		session.State.(*Client).OutEvent_GetEventNewYear(out)

	}

}

func (o OutEvent) QqVipContinue(session *net.Session, out *event_api.QqVipContinue_Out) {

	if session.State.(*Client).OutEvent_QqVipContinue != nil {

		session.State.(*Client).OutEvent_QqVipContinue(out)

	}

}

func (o OutEvent) DailyOnlineInfo(session *net.Session, out *event_api.DailyOnlineInfo_Out) {

	if session.State.(*Client).OutEvent_DailyOnlineInfo != nil {

		session.State.(*Client).OutEvent_DailyOnlineInfo(out)

	}

}

func (o OutEvent) TakeDailyOnlineAward(session *net.Session, out *event_api.TakeDailyOnlineAward_Out) {

	if session.State.(*Client).OutEvent_TakeDailyOnlineAward != nil {

		session.State.(*Client).OutEvent_TakeDailyOnlineAward(out)

	}

}

func (o OutFashion) FashionInfo(session *net.Session, out *fashion_api.FashionInfo_Out) {

	if session.State.(*Client).OutFashion_FashionInfo != nil {

		session.State.(*Client).OutFashion_FashionInfo(out)

	}

}

func (o OutFashion) DressFashion(session *net.Session, out *fashion_api.DressFashion_Out) {

	if session.State.(*Client).OutFashion_DressFashion != nil {

		session.State.(*Client).OutFashion_DressFashion(out)

	}

}

func (o OutPushNotify) PushInfo(session *net.Session, out *push_notify_api.PushInfo_Out) {

	if session.State.(*Client).OutPushNotify_PushInfo != nil {

		session.State.(*Client).OutPushNotify_PushInfo(out)

	}

}

func (o OutPushNotify) PushNotificationSwitch(session *net.Session, out *push_notify_api.PushNotificationSwitch_Out) {

	if session.State.(*Client).OutPushNotify_PushNotificationSwitch != nil {

		session.State.(*Client).OutPushNotify_PushNotificationSwitch(out)

	}

}

func (o OutMeditation) MeditationInfo(session *net.Session, out *meditation_api.MeditationInfo_Out) {

	if session.State.(*Client).OutMeditation_MeditationInfo != nil {

		session.State.(*Client).OutMeditation_MeditationInfo(out)

	}

}

func (o OutMeditation) StartMeditation(session *net.Session, out *meditation_api.StartMeditation_Out) {

	if session.State.(*Client).OutMeditation_StartMeditation != nil {

		session.State.(*Client).OutMeditation_StartMeditation(out)

	}

}

func (o OutMeditation) StopMeditation(session *net.Session, out *meditation_api.StopMeditation_Out) {

	if session.State.(*Client).OutMeditation_StopMeditation != nil {

		session.State.(*Client).OutMeditation_StopMeditation(out)

	}

}

func (o OutPetVirtualEnv) Info(session *net.Session, out *pet_virtual_env_api.Info_Out) {

	if session.State.(*Client).OutPetVirtualEnv_Info != nil {

		session.State.(*Client).OutPetVirtualEnv_Info(out)

	}

}

func (o OutPetVirtualEnv) TakeAward(session *net.Session, out *pet_virtual_env_api.TakeAward_Out) {

	if session.State.(*Client).OutPetVirtualEnv_TakeAward != nil {

		session.State.(*Client).OutPetVirtualEnv_TakeAward(out)

	}

}

func (o OutPetVirtualEnv) AutoFight(session *net.Session, out *pet_virtual_env_api.AutoFight_Out) {

	if session.State.(*Client).OutPetVirtualEnv_AutoFight != nil {

		session.State.(*Client).OutPetVirtualEnv_AutoFight(out)

	}

}

func (o OutPetVirtualEnv) PveKills(session *net.Session, out *pet_virtual_env_api.PveKills_Out) {

	if session.State.(*Client).OutPetVirtualEnv_PveKills != nil {

		session.State.(*Client).OutPetVirtualEnv_PveKills(out)

	}

}

func (o OutChannel) GetLatestWorldChannelMessage(session *net.Session, out *channel_api.GetLatestWorldChannelMessage_Out) {

	if session.State.(*Client).OutChannel_GetLatestWorldChannelMessage != nil {

		session.State.(*Client).OutChannel_GetLatestWorldChannelMessage(out)

	}

}

func (o OutChannel) AddWorldChat(session *net.Session, out *channel_api.AddWorldChat_Out) {

	if session.State.(*Client).OutChannel_AddWorldChat != nil {

		session.State.(*Client).OutChannel_AddWorldChat(out)

	}

}

func (o OutChannel) WorldChannelInfo(session *net.Session, out *channel_api.WorldChannelInfo_Out) {

	if session.State.(*Client).OutChannel_WorldChannelInfo != nil {

		session.State.(*Client).OutChannel_WorldChannelInfo(out)

	}

}

func (o OutChannel) SendGlobalMessages(session *net.Session, out *channel_api.SendGlobalMessages_Out) {

	if session.State.(*Client).OutChannel_SendGlobalMessages != nil {

		session.State.(*Client).OutChannel_SendGlobalMessages(out)

	}

}

func (o OutChannel) AddCliqueChat(session *net.Session, out *channel_api.AddCliqueChat_Out) {

	if session.State.(*Client).OutChannel_AddCliqueChat != nil {

		session.State.(*Client).OutChannel_AddCliqueChat(out)

	}

}

func (o OutChannel) GetLatestCliqueMessages(session *net.Session, out *channel_api.GetLatestCliqueMessages_Out) {

	if session.State.(*Client).OutChannel_GetLatestCliqueMessages != nil {

		session.State.(*Client).OutChannel_GetLatestCliqueMessages(out)

	}

}

func (o OutChannel) SendCliqueMessage(session *net.Session, out *channel_api.SendCliqueMessage_Out) {

	if session.State.(*Client).OutChannel_SendCliqueMessage != nil {

		session.State.(*Client).OutChannel_SendCliqueMessage(out)

	}

}

func (o OutDrivingSword) CloudMapInfo(session *net.Session, out *driving_sword_api.CloudMapInfo_Out) {

	if session.State.(*Client).OutDrivingSword_CloudMapInfo != nil {

		session.State.(*Client).OutDrivingSword_CloudMapInfo(out)

	}

}

func (o OutDrivingSword) CloudClimb(session *net.Session, out *driving_sword_api.CloudClimb_Out) {

	if session.State.(*Client).OutDrivingSword_CloudClimb != nil {

		session.State.(*Client).OutDrivingSword_CloudClimb(out)

	}

}

func (o OutDrivingSword) CloudTeleport(session *net.Session, out *driving_sword_api.CloudTeleport_Out) {

	if session.State.(*Client).OutDrivingSword_CloudTeleport != nil {

		session.State.(*Client).OutDrivingSword_CloudTeleport(out)

	}

}

func (o OutDrivingSword) AreaTeleport(session *net.Session, out *driving_sword_api.AreaTeleport_Out) {

	if session.State.(*Client).OutDrivingSword_AreaTeleport != nil {

		session.State.(*Client).OutDrivingSword_AreaTeleport(out)

	}

}

func (o OutDrivingSword) AreaMove(session *net.Session, out *driving_sword_api.AreaMove_Out) {

	if session.State.(*Client).OutDrivingSword_AreaMove != nil {

		session.State.(*Client).OutDrivingSword_AreaMove(out)

	}

}

func (o OutDrivingSword) ExplorerStartBattle(session *net.Session, out *driving_sword_api.ExplorerStartBattle_Out) {

	if session.State.(*Client).OutDrivingSword_ExplorerStartBattle != nil {

		session.State.(*Client).OutDrivingSword_ExplorerStartBattle(out)

	}

}

func (o OutDrivingSword) ExplorerAward(session *net.Session, out *driving_sword_api.ExplorerAward_Out) {

	if session.State.(*Client).OutDrivingSword_ExplorerAward != nil {

		session.State.(*Client).OutDrivingSword_ExplorerAward(out)

	}

}

func (o OutDrivingSword) ExplorerGarrison(session *net.Session, out *driving_sword_api.ExplorerGarrison_Out) {

	if session.State.(*Client).OutDrivingSword_ExplorerGarrison != nil {

		session.State.(*Client).OutDrivingSword_ExplorerGarrison(out)

	}

}

func (o OutDrivingSword) VisitMountain(session *net.Session, out *driving_sword_api.VisitMountain_Out) {

	if session.State.(*Client).OutDrivingSword_VisitMountain != nil {

		session.State.(*Client).OutDrivingSword_VisitMountain(out)

	}

}

func (o OutDrivingSword) VisiterStartBattle(session *net.Session, out *driving_sword_api.VisiterStartBattle_Out) {

	if session.State.(*Client).OutDrivingSword_VisiterStartBattle != nil {

		session.State.(*Client).OutDrivingSword_VisiterStartBattle(out)

	}

}

func (o OutDrivingSword) VisiterAward(session *net.Session, out *driving_sword_api.VisiterAward_Out) {

	if session.State.(*Client).OutDrivingSword_VisiterAward != nil {

		session.State.(*Client).OutDrivingSword_VisiterAward(out)

	}

}

func (o OutDrivingSword) MountainTreasureOpen(session *net.Session, out *driving_sword_api.MountainTreasureOpen_Out) {

	if session.State.(*Client).OutDrivingSword_MountainTreasureOpen != nil {

		session.State.(*Client).OutDrivingSword_MountainTreasureOpen(out)

	}

}

func (o OutDrivingSword) ListGarrisons(session *net.Session, out *driving_sword_api.ListGarrisons_Out) {

	if session.State.(*Client).OutDrivingSword_ListGarrisons != nil {

		session.State.(*Client).OutDrivingSword_ListGarrisons(out)

	}

}

func (o OutDrivingSword) AwardGarrison(session *net.Session, out *driving_sword_api.AwardGarrison_Out) {

	if session.State.(*Client).OutDrivingSword_AwardGarrison != nil {

		session.State.(*Client).OutDrivingSword_AwardGarrison(out)

	}

}

func (o OutDrivingSword) EndGarrison(session *net.Session, out *driving_sword_api.EndGarrison_Out) {

	if session.State.(*Client).OutDrivingSword_EndGarrison != nil {

		session.State.(*Client).OutDrivingSword_EndGarrison(out)

	}

}

func (o OutDrivingSword) BuyAllowedAction(session *net.Session, out *driving_sword_api.BuyAllowedAction_Out) {

	if session.State.(*Client).OutDrivingSword_BuyAllowedAction != nil {

		session.State.(*Client).OutDrivingSword_BuyAllowedAction(out)

	}

}

func (o OutTotem) Info(session *net.Session, out *totem_api.Info_Out) {

	if session.State.(*Client).OutTotem_Info != nil {

		session.State.(*Client).OutTotem_Info(out)

	}

}

func (o OutTotem) CallTotem(session *net.Session, out *totem_api.CallTotem_Out) {

	if session.State.(*Client).OutTotem_CallTotem != nil {

		session.State.(*Client).OutTotem_CallTotem(out)

	}

}

func (o OutTotem) Upgrade(session *net.Session, out *totem_api.Upgrade_Out) {

	if session.State.(*Client).OutTotem_Upgrade != nil {

		session.State.(*Client).OutTotem_Upgrade(out)

	}

}

func (o OutTotem) Decompose(session *net.Session, out *totem_api.Decompose_Out) {

	if session.State.(*Client).OutTotem_Decompose != nil {

		session.State.(*Client).OutTotem_Decompose(out)

	}

}

func (o OutTotem) Equip(session *net.Session, out *totem_api.Equip_Out) {

	if session.State.(*Client).OutTotem_Equip != nil {

		session.State.(*Client).OutTotem_Equip(out)

	}

}

func (o OutTotem) Unequip(session *net.Session, out *totem_api.Unequip_Out) {

	if session.State.(*Client).OutTotem_Unequip != nil {

		session.State.(*Client).OutTotem_Unequip(out)

	}

}

func (o OutTotem) EquipPosChange(session *net.Session, out *totem_api.EquipPosChange_Out) {

	if session.State.(*Client).OutTotem_EquipPosChange != nil {

		session.State.(*Client).OutTotem_EquipPosChange(out)

	}

}

func (o OutTotem) Swap(session *net.Session, out *totem_api.Swap_Out) {

	if session.State.(*Client).OutTotem_Swap != nil {

		session.State.(*Client).OutTotem_Swap(out)

	}

}

func (o OutTotem) IsBagFull(session *net.Session, out *totem_api.IsBagFull_Out) {

	if session.State.(*Client).OutTotem_IsBagFull != nil {

		session.State.(*Client).OutTotem_IsBagFull(out)

	}

}

func (o OutMoneyTree) GetTreeStatus(session *net.Session, out *money_tree_api.GetTreeStatus_Out) {

	if session.State.(*Client).OutMoneyTree_GetTreeStatus != nil {

		session.State.(*Client).OutMoneyTree_GetTreeStatus(out)

	}

}

func (o OutMoneyTree) GetTreeMoney(session *net.Session, out *money_tree_api.GetTreeMoney_Out) {

	if session.State.(*Client).OutMoneyTree_GetTreeMoney != nil {

		session.State.(*Client).OutMoneyTree_GetTreeMoney(out)

	}

}

func (o OutMoneyTree) WaveTree(session *net.Session, out *money_tree_api.WaveTree_Out) {

	if session.State.(*Client).OutMoneyTree_WaveTree != nil {

		session.State.(*Client).OutMoneyTree_WaveTree(out)

	}

}

func (o OutClique) CreateClique(session *net.Session, out *clique_api.CreateClique_Out) {

	if session.State.(*Client).OutClique_CreateClique != nil {

		session.State.(*Client).OutClique_CreateClique(out)

	}

}

func (o OutClique) ApplyJoinClique(session *net.Session, out *clique_api.ApplyJoinClique_Out) {

	if session.State.(*Client).OutClique_ApplyJoinClique != nil {

		session.State.(*Client).OutClique_ApplyJoinClique(out)

	}

}

func (o OutClique) CancelApplyClique(session *net.Session, out *clique_api.CancelApplyClique_Out) {

	if session.State.(*Client).OutClique_CancelApplyClique != nil {

		session.State.(*Client).OutClique_CancelApplyClique(out)

	}

}

func (o OutClique) ProcessJoinApply(session *net.Session, out *clique_api.ProcessJoinApply_Out) {

	if session.State.(*Client).OutClique_ProcessJoinApply != nil {

		session.State.(*Client).OutClique_ProcessJoinApply(out)

	}

}

func (o OutClique) ElectOwner(session *net.Session, out *clique_api.ElectOwner_Out) {

	if session.State.(*Client).OutClique_ElectOwner != nil {

		session.State.(*Client).OutClique_ElectOwner(out)

	}

}

func (o OutClique) MangeMember(session *net.Session, out *clique_api.MangeMember_Out) {

	if session.State.(*Client).OutClique_MangeMember != nil {

		session.State.(*Client).OutClique_MangeMember(out)

	}

}

func (o OutClique) DestoryClique(session *net.Session, out *clique_api.DestoryClique_Out) {

	if session.State.(*Client).OutClique_DestoryClique != nil {

		session.State.(*Client).OutClique_DestoryClique(out)

	}

}

func (o OutClique) UpdateAnnounce(session *net.Session, out *clique_api.UpdateAnnounce_Out) {

	if session.State.(*Client).OutClique_UpdateAnnounce != nil {

		session.State.(*Client).OutClique_UpdateAnnounce(out)

	}

}

func (o OutClique) LeaveClique(session *net.Session, out *clique_api.LeaveClique_Out) {

	if session.State.(*Client).OutClique_LeaveClique != nil {

		session.State.(*Client).OutClique_LeaveClique(out)

	}

}

func (o OutClique) ListClique(session *net.Session, out *clique_api.ListClique_Out) {

	if session.State.(*Client).OutClique_ListClique != nil {

		session.State.(*Client).OutClique_ListClique(out)

	}

}

func (o OutClique) CliquePublicInfo(session *net.Session, out *clique_api.CliquePublicInfo_Out) {

	if session.State.(*Client).OutClique_CliquePublicInfo != nil {

		session.State.(*Client).OutClique_CliquePublicInfo(out)

	}

}

func (o OutClique) CliqueInfo(session *net.Session, out *clique_api.CliqueInfo_Out) {

	if session.State.(*Client).OutClique_CliqueInfo != nil {

		session.State.(*Client).OutClique_CliqueInfo(out)

	}

}

func (o OutClique) ListApply(session *net.Session, out *clique_api.ListApply_Out) {

	if session.State.(*Client).OutClique_ListApply != nil {

		session.State.(*Client).OutClique_ListApply(out)

	}

}

func (o OutClique) OperaErrorNotify(session *net.Session, out *clique_api.OperaErrorNotify_Out) {

	if session.State.(*Client).OutClique_OperaErrorNotify != nil {

		session.State.(*Client).OutClique_OperaErrorNotify(out)

	}

}

func (o OutClique) EnterClubhouse(session *net.Session, out *clique_api.EnterClubhouse_Out) {

	if session.State.(*Client).OutClique_EnterClubhouse != nil {

		session.State.(*Client).OutClique_EnterClubhouse(out)

	}

}

func (o OutClique) LeaveClubhouse(session *net.Session, out *clique_api.LeaveClubhouse_Out) {

	if session.State.(*Client).OutClique_LeaveClubhouse != nil {

		session.State.(*Client).OutClique_LeaveClubhouse(out)

	}

}

func (o OutClique) ClubMove(session *net.Session, out *clique_api.ClubMove_Out) {

	if session.State.(*Client).OutClique_ClubMove != nil {

		session.State.(*Client).OutClique_ClubMove(out)

	}

}

func (o OutClique) NotifyClubhousePlayers(session *net.Session, out *clique_api.NotifyClubhousePlayers_Out) {

	if session.State.(*Client).OutClique_NotifyClubhousePlayers != nil {

		session.State.(*Client).OutClique_NotifyClubhousePlayers(out)

	}

}

func (o OutClique) NotifyNewPlayer(session *net.Session, out *clique_api.NotifyNewPlayer_Out) {

	if session.State.(*Client).OutClique_NotifyNewPlayer != nil {

		session.State.(*Client).OutClique_NotifyNewPlayer(out)

	}

}

func (o OutClique) NotifyUpdatePlayer(session *net.Session, out *clique_api.NotifyUpdatePlayer_Out) {

	if session.State.(*Client).OutClique_NotifyUpdatePlayer != nil {

		session.State.(*Client).OutClique_NotifyUpdatePlayer(out)

	}

}

func (o OutClique) CliquePublicInfoSummary(session *net.Session, out *clique_api.CliquePublicInfoSummary_Out) {

	if session.State.(*Client).OutClique_CliquePublicInfoSummary != nil {

		session.State.(*Client).OutClique_CliquePublicInfoSummary(out)

	}

}

func (o OutClique) CliqueAutoAudit(session *net.Session, out *clique_api.CliqueAutoAudit_Out) {

	if session.State.(*Client).OutClique_CliqueAutoAudit != nil {

		session.State.(*Client).OutClique_CliqueAutoAudit(out)

	}

}

func (o OutClique) CliqueBaseDonate(session *net.Session, out *clique_api.CliqueBaseDonate_Out) {

	if session.State.(*Client).OutClique_CliqueBaseDonate != nil {

		session.State.(*Client).OutClique_CliqueBaseDonate(out)

	}

}

func (o OutClique) NotifyLeaveClique(session *net.Session, out *clique_api.NotifyLeaveClique_Out) {

	if session.State.(*Client).OutClique_NotifyLeaveClique != nil {

		session.State.(*Client).OutClique_NotifyLeaveClique(out)

	}

}

func (o OutClique) NotifyJoincliqueSuccess(session *net.Session, out *clique_api.NotifyJoincliqueSuccess_Out) {

	if session.State.(*Client).OutClique_NotifyJoincliqueSuccess != nil {

		session.State.(*Client).OutClique_NotifyJoincliqueSuccess(out)

	}

}

func (o OutClique) NotifyCliqueMangerAction(session *net.Session, out *clique_api.NotifyCliqueMangerAction_Out) {

	if session.State.(*Client).OutClique_NotifyCliqueMangerAction != nil {

		session.State.(*Client).OutClique_NotifyCliqueMangerAction(out)

	}

}

func (o OutClique) CliqueRecruitment(session *net.Session, out *clique_api.CliqueRecruitment_Out) {

	if session.State.(*Client).OutClique_CliqueRecruitment != nil {

		session.State.(*Client).OutClique_CliqueRecruitment(out)

	}

}

func (o OutClique) NotifyCliqueAnnounce(session *net.Session, out *clique_api.NotifyCliqueAnnounce_Out) {

	if session.State.(*Client).OutClique_NotifyCliqueAnnounce != nil {

		session.State.(*Client).OutClique_NotifyCliqueAnnounce(out)

	}

}

func (o OutClique) NotifyCliqueElectowner(session *net.Session, out *clique_api.NotifyCliqueElectowner_Out) {

	if session.State.(*Client).OutClique_NotifyCliqueElectowner != nil {

		session.State.(*Client).OutClique_NotifyCliqueElectowner(out)

	}

}

func (o OutClique) QuickApply(session *net.Session, out *clique_api.QuickApply_Out) {

	if session.State.(*Client).OutClique_QuickApply != nil {

		session.State.(*Client).OutClique_QuickApply(out)

	}

}

func (o OutClique) NotifyContribChange(session *net.Session, out *clique_api.NotifyContribChange_Out) {

	if session.State.(*Client).OutClique_NotifyContribChange != nil {

		session.State.(*Client).OutClique_NotifyContribChange(out)

	}

}

func (o OutClique) OtherClique(session *net.Session, out *clique_api.OtherClique_Out) {

	if session.State.(*Client).OutClique_OtherClique != nil {

		session.State.(*Client).OutClique_OtherClique(out)

	}

}

func (o OutCliqueBuilding) CliqueBaseDonate(session *net.Session, out *clique_building_api.CliqueBaseDonate_Out) {

	if session.State.(*Client).OutCliqueBuilding_CliqueBaseDonate != nil {

		session.State.(*Client).OutCliqueBuilding_CliqueBaseDonate(out)

	}

}

func (o OutCliqueBuilding) CliqueBuildingStatus(session *net.Session, out *clique_building_api.CliqueBuildingStatus_Out) {

	if session.State.(*Client).OutCliqueBuilding_CliqueBuildingStatus != nil {

		session.State.(*Client).OutCliqueBuilding_CliqueBuildingStatus(out)

	}

}

func (o OutCliqueBuilding) CliqueBankDonate(session *net.Session, out *clique_building_api.CliqueBankDonate_Out) {

	if session.State.(*Client).OutCliqueBuilding_CliqueBankDonate != nil {

		session.State.(*Client).OutCliqueBuilding_CliqueBankDonate(out)

	}

}

func (o OutCliqueBuilding) CliqueBankBuy(session *net.Session, out *clique_building_api.CliqueBankBuy_Out) {

	if session.State.(*Client).OutCliqueBuilding_CliqueBankBuy != nil {

		session.State.(*Client).OutCliqueBuilding_CliqueBankBuy(out)

	}

}

func (o OutCliqueBuilding) CliqueBankSold(session *net.Session, out *clique_building_api.CliqueBankSold_Out) {

	if session.State.(*Client).OutCliqueBuilding_CliqueBankSold != nil {

		session.State.(*Client).OutCliqueBuilding_CliqueBankSold(out)

	}

}

func (o OutCliqueBuilding) CliqueKongfuDonate(session *net.Session, out *clique_building_api.CliqueKongfuDonate_Out) {

	if session.State.(*Client).OutCliqueBuilding_CliqueKongfuDonate != nil {

		session.State.(*Client).OutCliqueBuilding_CliqueKongfuDonate(out)

	}

}

func (o OutCliqueBuilding) CliqueKongfuInfo(session *net.Session, out *clique_building_api.CliqueKongfuInfo_Out) {

	if session.State.(*Client).OutCliqueBuilding_CliqueKongfuInfo != nil {

		session.State.(*Client).OutCliqueBuilding_CliqueKongfuInfo(out)

	}

}

func (o OutCliqueBuilding) CliqueKongfuTrain(session *net.Session, out *clique_building_api.CliqueKongfuTrain_Out) {

	if session.State.(*Client).OutCliqueBuilding_CliqueKongfuTrain != nil {

		session.State.(*Client).OutCliqueBuilding_CliqueKongfuTrain(out)

	}

}

func (o OutCliqueBuilding) CliqueTempleWorship(session *net.Session, out *clique_building_api.CliqueTempleWorship_Out) {

	if session.State.(*Client).OutCliqueBuilding_CliqueTempleWorship != nil {

		session.State.(*Client).OutCliqueBuilding_CliqueTempleWorship(out)

	}

}

func (o OutCliqueBuilding) CliqueTempleDonate(session *net.Session, out *clique_building_api.CliqueTempleDonate_Out) {

	if session.State.(*Client).OutCliqueBuilding_CliqueTempleDonate != nil {

		session.State.(*Client).OutCliqueBuilding_CliqueTempleDonate(out)

	}

}

func (o OutCliqueBuilding) CliqueTempleInfo(session *net.Session, out *clique_building_api.CliqueTempleInfo_Out) {

	if session.State.(*Client).OutCliqueBuilding_CliqueTempleInfo != nil {

		session.State.(*Client).OutCliqueBuilding_CliqueTempleInfo(out)

	}

}

func (o OutCliqueBuilding) CliqueStoreDonate(session *net.Session, out *clique_building_api.CliqueStoreDonate_Out) {

	if session.State.(*Client).OutCliqueBuilding_CliqueStoreDonate != nil {

		session.State.(*Client).OutCliqueBuilding_CliqueStoreDonate(out)

	}

}

func (o OutCliqueBuilding) CliqueStoreInfo(session *net.Session, out *clique_building_api.CliqueStoreInfo_Out) {

	if session.State.(*Client).OutCliqueBuilding_CliqueStoreInfo != nil {

		session.State.(*Client).OutCliqueBuilding_CliqueStoreInfo(out)

	}

}

func (o OutCliqueBuilding) CliqueStoreSendChest(session *net.Session, out *clique_building_api.CliqueStoreSendChest_Out) {

	if session.State.(*Client).OutCliqueBuilding_CliqueStoreSendChest != nil {

		session.State.(*Client).OutCliqueBuilding_CliqueStoreSendChest(out)

	}

}

func (o OutCliqueQuest) GetCliqueDailyQuest(session *net.Session, out *clique_quest_api.GetCliqueDailyQuest_Out) {

	if session.State.(*Client).OutCliqueQuest_GetCliqueDailyQuest != nil {

		session.State.(*Client).OutCliqueQuest_GetCliqueDailyQuest(out)

	}

}

func (o OutCliqueQuest) AwardCliqueDailyQuest(session *net.Session, out *clique_quest_api.AwardCliqueDailyQuest_Out) {

	if session.State.(*Client).OutCliqueQuest_AwardCliqueDailyQuest != nil {

		session.State.(*Client).OutCliqueQuest_AwardCliqueDailyQuest(out)

	}

}

func (o OutCliqueQuest) NotifyCliqueDailyChange(session *net.Session, out *clique_quest_api.NotifyCliqueDailyChange_Out) {

	if session.State.(*Client).OutCliqueQuest_NotifyCliqueDailyChange != nil {

		session.State.(*Client).OutCliqueQuest_NotifyCliqueDailyChange(out)

	}

}

func (o OutCliqueQuest) GetCliqueBuildingQuest(session *net.Session, out *clique_quest_api.GetCliqueBuildingQuest_Out) {

	if session.State.(*Client).OutCliqueQuest_GetCliqueBuildingQuest != nil {

		session.State.(*Client).OutCliqueQuest_GetCliqueBuildingQuest(out)

	}

}

func (o OutCliqueQuest) AwardCliqueBuildingQuest(session *net.Session, out *clique_quest_api.AwardCliqueBuildingQuest_Out) {

	if session.State.(*Client).OutCliqueQuest_AwardCliqueBuildingQuest != nil {

		session.State.(*Client).OutCliqueQuest_AwardCliqueBuildingQuest(out)

	}

}

func (o OutCliqueEscort) EscortInfo(session *net.Session, out *clique_escort_api.EscortInfo_Out) {

	if session.State.(*Client).OutCliqueEscort_EscortInfo != nil {

		session.State.(*Client).OutCliqueEscort_EscortInfo(out)

	}

}

func (o OutCliqueEscort) GetIngotBoat(session *net.Session, out *clique_escort_api.GetIngotBoat_Out) {

	if session.State.(*Client).OutCliqueEscort_GetIngotBoat != nil {

		session.State.(*Client).OutCliqueEscort_GetIngotBoat(out)

	}

}

func (o OutCliqueEscort) StartEscort(session *net.Session, out *clique_escort_api.StartEscort_Out) {

	if session.State.(*Client).OutCliqueEscort_StartEscort != nil {

		session.State.(*Client).OutCliqueEscort_StartEscort(out)

	}

}

func (o OutCliqueEscort) HijackBoat(session *net.Session, out *clique_escort_api.HijackBoat_Out) {

	if session.State.(*Client).OutCliqueEscort_HijackBoat != nil {

		session.State.(*Client).OutCliqueEscort_HijackBoat(out)

	}

}

func (o OutCliqueEscort) RecoverBoat(session *net.Session, out *clique_escort_api.RecoverBoat_Out) {

	if session.State.(*Client).OutCliqueEscort_RecoverBoat != nil {

		session.State.(*Client).OutCliqueEscort_RecoverBoat(out)

	}

}

func (o OutCliqueEscort) ListBoats(session *net.Session, out *clique_escort_api.ListBoats_Out) {

	if session.State.(*Client).OutCliqueEscort_ListBoats != nil {

		session.State.(*Client).OutCliqueEscort_ListBoats(out)

	}

}

func (o OutCliqueEscort) GetRandomBoat(session *net.Session, out *clique_escort_api.GetRandomBoat_Out) {

	if session.State.(*Client).OutCliqueEscort_GetRandomBoat != nil {

		session.State.(*Client).OutCliqueEscort_GetRandomBoat(out)

	}

}

func (o OutCliqueEscort) NotifyEscortFinished(session *net.Session, out *clique_escort_api.NotifyEscortFinished_Out) {

	if session.State.(*Client).OutCliqueEscort_NotifyEscortFinished != nil {

		session.State.(*Client).OutCliqueEscort_NotifyEscortFinished(out)

	}

}

func (o OutCliqueEscort) NotifyHijackFinished(session *net.Session, out *clique_escort_api.NotifyHijackFinished_Out) {

	if session.State.(*Client).OutCliqueEscort_NotifyHijackFinished != nil {

		session.State.(*Client).OutCliqueEscort_NotifyHijackFinished(out)

	}

}

func (o OutCliqueEscort) NotifyRecoverBattleWin(session *net.Session, out *clique_escort_api.NotifyRecoverBattleWin_Out) {

	if session.State.(*Client).OutCliqueEscort_NotifyRecoverBattleWin != nil {

		session.State.(*Client).OutCliqueEscort_NotifyRecoverBattleWin(out)

	}

}

func (o OutCliqueEscort) NotifyHijackBattleWin(session *net.Session, out *clique_escort_api.NotifyHijackBattleWin_Out) {

	if session.State.(*Client).OutCliqueEscort_NotifyHijackBattleWin != nil {

		session.State.(*Client).OutCliqueEscort_NotifyHijackBattleWin(out)

	}

}

func (o OutCliqueEscort) TakeHijackAward(session *net.Session, out *clique_escort_api.TakeHijackAward_Out) {

	if session.State.(*Client).OutCliqueEscort_TakeHijackAward != nil {

		session.State.(*Client).OutCliqueEscort_TakeHijackAward(out)

	}

}

func (o OutCliqueEscort) TakeEscortAward(session *net.Session, out *clique_escort_api.TakeEscortAward_Out) {

	if session.State.(*Client).OutCliqueEscort_TakeEscortAward != nil {

		session.State.(*Client).OutCliqueEscort_TakeEscortAward(out)

	}

}

func (o OutCliqueEscort) GetCliqueBoatMessages(session *net.Session, out *clique_escort_api.GetCliqueBoatMessages_Out) {

	if session.State.(*Client).OutCliqueEscort_GetCliqueBoatMessages != nil {

		session.State.(*Client).OutCliqueEscort_GetCliqueBoatMessages(out)

	}

}

func (o OutCliqueEscort) SendCliqueBoatMessage(session *net.Session, out *clique_escort_api.SendCliqueBoatMessage_Out) {

	if session.State.(*Client).OutCliqueEscort_SendCliqueBoatMessage != nil {

		session.State.(*Client).OutCliqueEscort_SendCliqueBoatMessage(out)

	}

}

func (o OutCliqueEscort) ReadCliqueBoatMessage(session *net.Session, out *clique_escort_api.ReadCliqueBoatMessage_Out) {

	if session.State.(*Client).OutCliqueEscort_ReadCliqueBoatMessage != nil {

		session.State.(*Client).OutCliqueEscort_ReadCliqueBoatMessage(out)

	}

}

func (o OutCliqueEscort) NotifyBoatStatusChange(session *net.Session, out *clique_escort_api.NotifyBoatStatusChange_Out) {

	if session.State.(*Client).OutCliqueEscort_NotifyBoatStatusChange != nil {

		session.State.(*Client).OutCliqueEscort_NotifyBoatStatusChange(out)

	}

}

func (o OutDespairLand) DespairLandInfo(session *net.Session, out *despair_land_api.DespairLandInfo_Out) {

	if session.State.(*Client).OutDespairLand_DespairLandInfo != nil {

		session.State.(*Client).OutDespairLand_DespairLandInfo(out)

	}

}

func (o OutDespairLand) DespairLandCampInfo(session *net.Session, out *despair_land_api.DespairLandCampInfo_Out) {

	if session.State.(*Client).OutDespairLand_DespairLandCampInfo != nil {

		session.State.(*Client).OutDespairLand_DespairLandCampInfo(out)

	}

}

func (o OutDespairLand) DespairLandCampPlayerInfo(session *net.Session, out *despair_land_api.DespairLandCampPlayerInfo_Out) {

	if session.State.(*Client).OutDespairLand_DespairLandCampPlayerInfo != nil {

		session.State.(*Client).OutDespairLand_DespairLandCampPlayerInfo(out)

	}

}

func (o OutDespairLand) DespairLandPickBox(session *net.Session, out *despair_land_api.DespairLandPickBox_Out) {

	if session.State.(*Client).OutDespairLand_DespairLandPickBox != nil {

		session.State.(*Client).OutDespairLand_DespairLandPickBox(out)

	}

}

func (o OutDespairLand) DespairLandWatchRecord(session *net.Session, out *despair_land_api.DespairLandWatchRecord_Out) {

	if session.State.(*Client).OutDespairLand_DespairLandWatchRecord != nil {

		session.State.(*Client).OutDespairLand_DespairLandWatchRecord(out)

	}

}

func (o OutDespairLand) DespairLandBuyBattleNum(session *net.Session, out *despair_land_api.DespairLandBuyBattleNum_Out) {

	if session.State.(*Client).OutDespairLand_DespairLandBuyBattleNum != nil {

		session.State.(*Client).OutDespairLand_DespairLandBuyBattleNum(out)

	}

}

func (o OutDespairLand) DespairLandBossInfo(session *net.Session, out *despair_land_api.DespairLandBossInfo_Out) {

	if session.State.(*Client).OutDespairLand_DespairLandBossInfo != nil {

		session.State.(*Client).OutDespairLand_DespairLandBossInfo(out)

	}

}

func (o OutDespairLand) DespairLandNotifyBossOpen(session *net.Session, out *despair_land_api.DespairLandNotifyBossOpen_Out) {

	if session.State.(*Client).OutDespairLand_DespairLandNotifyBossOpen != nil {

		session.State.(*Client).OutDespairLand_DespairLandNotifyBossOpen(out)

	}

}

func (o OutDespairLand) DespairLandNotifyBossDead(session *net.Session, out *despair_land_api.DespairLandNotifyBossDead_Out) {

	if session.State.(*Client).OutDespairLand_DespairLandNotifyBossDead != nil {

		session.State.(*Client).OutDespairLand_DespairLandNotifyBossDead(out)

	}

}

func (o OutDespairLand) DespairLandBuyBossBattleNum(session *net.Session, out *despair_land_api.DespairLandBuyBossBattleNum_Out) {

	if session.State.(*Client).OutDespairLand_DespairLandBuyBossBattleNum != nil {

		session.State.(*Client).OutDespairLand_DespairLandBuyBossBattleNum(out)

	}

}

func (o OutDespairLand) DespairLandNotifyPass(session *net.Session, out *despair_land_api.DespairLandNotifyPass_Out) {

	if session.State.(*Client).OutDespairLand_DespairLandNotifyPass != nil {

		session.State.(*Client).OutDespairLand_DespairLandNotifyPass(out)

	}

}

func (o OutDespairLand) DespairLandPickThreeStarBox(session *net.Session, out *despair_land_api.DespairLandPickThreeStarBox_Out) {

	if session.State.(*Client).OutDespairLand_DespairLandPickThreeStarBox != nil {

		session.State.(*Client).OutDespairLand_DespairLandPickThreeStarBox(out)

	}

}

func (o OutDespairLand) DespairLandBattleBossAwardInfo(session *net.Session, out *despair_land_api.DespairLandBattleBossAwardInfo_Out) {

	if session.State.(*Client).OutDespairLand_DespairLandBattleBossAwardInfo != nil {

		session.State.(*Client).OutDespairLand_DespairLandBattleBossAwardInfo(out)

	}

}

func (o OutDespairLand) DespairLandBossOpenInfo(session *net.Session, out *despair_land_api.DespairLandBossOpenInfo_Out) {

	if session.State.(*Client).OutDespairLand_DespairLandBossOpenInfo != nil {

		session.State.(*Client).OutDespairLand_DespairLandBossOpenInfo(out)

	}

}

func (o OutDespairLand) DespairLandNotifyWeeklyRefresh(session *net.Session, out *despair_land_api.DespairLandNotifyWeeklyRefresh_Out) {

	if session.State.(*Client).OutDespairLand_DespairLandNotifyWeeklyRefresh != nil {

		session.State.(*Client).OutDespairLand_DespairLandNotifyWeeklyRefresh(out)

	}

}

func (o OutAwaken) AwakenInfo(session *net.Session, out *awaken_api.AwakenInfo_Out) {

	if session.State.(*Client).OutAwaken_AwakenInfo != nil {

		session.State.(*Client).OutAwaken_AwakenInfo(out)

	}

}

func (o OutAwaken) LevelupAttr(session *net.Session, out *awaken_api.LevelupAttr_Out) {

	if session.State.(*Client).OutAwaken_LevelupAttr != nil {

		session.State.(*Client).OutAwaken_LevelupAttr(out)

	}

}

func (o OutTaoyuan) TaoyuanInfo(session *net.Session, out *taoyuan_api.TaoyuanInfo_Out) {

	if session.State.(*Client).OutTaoyuan_TaoyuanInfo != nil {

		session.State.(*Client).OutTaoyuan_TaoyuanInfo(out)

	}

}

func (o OutTaoyuan) Bless(session *net.Session, out *taoyuan_api.Bless_Out) {

	if session.State.(*Client).OutTaoyuan_Bless != nil {

		session.State.(*Client).OutTaoyuan_Bless(out)

	}

}

func (o OutTaoyuan) ShopBuy(session *net.Session, out *taoyuan_api.ShopBuy_Out) {

	if session.State.(*Client).OutTaoyuan_ShopBuy != nil {

		session.State.(*Client).OutTaoyuan_ShopBuy(out)

	}

}

func (o OutTaoyuan) ShopSell(session *net.Session, out *taoyuan_api.ShopSell_Out) {

	if session.State.(*Client).OutTaoyuan_ShopSell != nil {

		session.State.(*Client).OutTaoyuan_ShopSell(out)

	}

}

func (o OutTaoyuan) GetAllItems(session *net.Session, out *taoyuan_api.GetAllItems_Out) {

	if session.State.(*Client).OutTaoyuan_GetAllItems != nil {

		session.State.(*Client).OutTaoyuan_GetAllItems(out)

	}

}

func (o OutTaoyuan) FiledsInfo(session *net.Session, out *taoyuan_api.FiledsInfo_Out) {

	if session.State.(*Client).OutTaoyuan_FiledsInfo != nil {

		session.State.(*Client).OutTaoyuan_FiledsInfo(out)

	}

}

func (o OutTaoyuan) GrowPlant(session *net.Session, out *taoyuan_api.GrowPlant_Out) {

	if session.State.(*Client).OutTaoyuan_GrowPlant != nil {

		session.State.(*Client).OutTaoyuan_GrowPlant(out)

	}

}

func (o OutTaoyuan) TakePlant(session *net.Session, out *taoyuan_api.TakePlant_Out) {

	if session.State.(*Client).OutTaoyuan_TakePlant != nil {

		session.State.(*Client).OutTaoyuan_TakePlant(out)

	}

}

func (o OutTaoyuan) UpgradeFiled(session *net.Session, out *taoyuan_api.UpgradeFiled_Out) {

	if session.State.(*Client).OutTaoyuan_UpgradeFiled != nil {

		session.State.(*Client).OutTaoyuan_UpgradeFiled(out)

	}

}

func (o OutTaoyuan) OpenFiled(session *net.Session, out *taoyuan_api.OpenFiled_Out) {

	if session.State.(*Client).OutTaoyuan_OpenFiled != nil {

		session.State.(*Client).OutTaoyuan_OpenFiled(out)

	}

}

func (o OutTaoyuan) StudySkill(session *net.Session, out *taoyuan_api.StudySkill_Out) {

	if session.State.(*Client).OutTaoyuan_StudySkill != nil {

		session.State.(*Client).OutTaoyuan_StudySkill(out)

	}

}

func (o OutTaoyuan) MakeProduct(session *net.Session, out *taoyuan_api.MakeProduct_Out) {

	if session.State.(*Client).OutTaoyuan_MakeProduct != nil {

		session.State.(*Client).OutTaoyuan_MakeProduct(out)

	}

}

func (o OutTaoyuan) TakeProduct(session *net.Session, out *taoyuan_api.TakeProduct_Out) {

	if session.State.(*Client).OutTaoyuan_TakeProduct != nil {

		session.State.(*Client).OutTaoyuan_TakeProduct(out)

	}

}

func (o OutTaoyuan) ProductMakeQueue(session *net.Session, out *taoyuan_api.ProductMakeQueue_Out) {

	if session.State.(*Client).OutTaoyuan_ProductMakeQueue != nil {

		session.State.(*Client).OutTaoyuan_ProductMakeQueue(out)

	}

}

func (o OutTaoyuan) QuestInfo(session *net.Session, out *taoyuan_api.QuestInfo_Out) {

	if session.State.(*Client).OutTaoyuan_QuestInfo != nil {

		session.State.(*Client).OutTaoyuan_QuestInfo(out)

	}

}

func (o OutTaoyuan) QuestFinish(session *net.Session, out *taoyuan_api.QuestFinish_Out) {

	if session.State.(*Client).OutTaoyuan_QuestFinish != nil {

		session.State.(*Client).OutTaoyuan_QuestFinish(out)

	}

}

func (o OutTaoyuan) QuestRefresh(session *net.Session, out *taoyuan_api.QuestRefresh_Out) {

	if session.State.(*Client).OutTaoyuan_QuestRefresh != nil {

		session.State.(*Client).OutTaoyuan_QuestRefresh(out)

	}

}

func (o OutTaoyuan) FriendTaoyuanInfo(session *net.Session, out *taoyuan_api.FriendTaoyuanInfo_Out) {

	if session.State.(*Client).OutTaoyuan_FriendTaoyuanInfo != nil {

		session.State.(*Client).OutTaoyuan_FriendTaoyuanInfo(out)

	}

}

func (o OutTaoyuan) SkillInfo(session *net.Session, out *taoyuan_api.SkillInfo_Out) {

	if session.State.(*Client).OutTaoyuan_SkillInfo != nil {

		session.State.(*Client).OutTaoyuan_SkillInfo(out)

	}

}

func (o OutTaoyuan) OpenQueue(session *net.Session, out *taoyuan_api.OpenQueue_Out) {

	if session.State.(*Client).OutTaoyuan_OpenQueue != nil {

		session.State.(*Client).OutTaoyuan_OpenQueue(out)

	}

}

func (o OutTaoyuan) PlantQuicklyMaturity(session *net.Session, out *taoyuan_api.PlantQuicklyMaturity_Out) {

	if session.State.(*Client).OutTaoyuan_PlantQuicklyMaturity != nil {

		session.State.(*Client).OutTaoyuan_PlantQuicklyMaturity(out)

	}

}

func (o OutTaoyuan) TaoyuanMessageInfo(session *net.Session, out *taoyuan_api.TaoyuanMessageInfo_Out) {

	if session.State.(*Client).OutTaoyuan_TaoyuanMessageInfo != nil {

		session.State.(*Client).OutTaoyuan_TaoyuanMessageInfo(out)

	}

}

func (o OutTaoyuan) TaoyuanMessageRead(session *net.Session, out *taoyuan_api.TaoyuanMessageRead_Out) {

	if session.State.(*Client).OutTaoyuan_TaoyuanMessageRead != nil {

		session.State.(*Client).OutTaoyuan_TaoyuanMessageRead(out)

	}

}

func (o OutTaoyuan) OpenProductBuilding(session *net.Session, out *taoyuan_api.OpenProductBuilding_Out) {

	if session.State.(*Client).OutTaoyuan_OpenProductBuilding != nil {

		session.State.(*Client).OutTaoyuan_OpenProductBuilding(out)

	}

}

func (o OutDraw) GetHeartDrawInfo(session *net.Session, out *draw_api.GetHeartDrawInfo_Out) {

	if session.State.(*Client).OutDraw_GetHeartDrawInfo != nil {

		session.State.(*Client).OutDraw_GetHeartDrawInfo(out)

	}

}

func (o OutDraw) HeartDraw(session *net.Session, out *draw_api.HeartDraw_Out) {

	if session.State.(*Client).OutDraw_HeartDraw != nil {

		session.State.(*Client).OutDraw_HeartDraw(out)

	}

}

func (o OutDraw) GetChestInfo(session *net.Session, out *draw_api.GetChestInfo_Out) {

	if session.State.(*Client).OutDraw_GetChestInfo != nil {

		session.State.(*Client).OutDraw_GetChestInfo(out)

	}

}

func (o OutDraw) DrawChest(session *net.Session, out *draw_api.DrawChest_Out) {

	if session.State.(*Client).OutDraw_DrawChest != nil {

		session.State.(*Client).OutDraw_DrawChest(out)

	}

}

func (o OutDraw) HeartInfo(session *net.Session, out *draw_api.HeartInfo_Out) {

	if session.State.(*Client).OutDraw_HeartInfo != nil {

		session.State.(*Client).OutDraw_HeartInfo(out)

	}

}

func (o OutDraw) GetFateBoxInfo(session *net.Session, out *draw_api.GetFateBoxInfo_Out) {

	if session.State.(*Client).OutDraw_GetFateBoxInfo != nil {

		session.State.(*Client).OutDraw_GetFateBoxInfo(out)

	}

}

func (o OutDraw) OpenFateBox(session *net.Session, out *draw_api.OpenFateBox_Out) {

	if session.State.(*Client).OutDraw_OpenFateBox != nil {

		session.State.(*Client).OutDraw_OpenFateBox(out)

	}

}

func (o OutDraw) ExchangeGiftCode(session *net.Session, out *draw_api.ExchangeGiftCode_Out) {

	if session.State.(*Client).OutDraw_ExchangeGiftCode != nil {

		session.State.(*Client).OutDraw_ExchangeGiftCode(out)

	}

}

func (o OutServerInfo) GetVersion(session *net.Session, out *server_info_api.GetVersion_Out) {

	if session.State.(*Client).OutServerInfo_GetVersion != nil {

		session.State.(*Client).OutServerInfo_GetVersion(out)

	}

}

func (o OutServerInfo) GetApiCount(session *net.Session, out *server_info_api.GetApiCount_Out) {

	if session.State.(*Client).OutServerInfo_GetApiCount != nil {

		session.State.(*Client).OutServerInfo_GetApiCount(out)

	}

}

func (o OutServerInfo) SearchPlayerRole(session *net.Session, out *server_info_api.SearchPlayerRole_Out) {

	if session.State.(*Client).OutServerInfo_SearchPlayerRole != nil {

		session.State.(*Client).OutServerInfo_SearchPlayerRole(out)

	}

}

func (o OutServerInfo) UpdateAccessToken(session *net.Session, out *server_info_api.UpdateAccessToken_Out) {

	if session.State.(*Client).OutServerInfo_UpdateAccessToken != nil {

		session.State.(*Client).OutServerInfo_UpdateAccessToken(out)

	}

}

func (o OutServerInfo) UpdateEventData(session *net.Session, out *server_info_api.UpdateEventData_Out) {

	if session.State.(*Client).OutServerInfo_UpdateEventData != nil {

		session.State.(*Client).OutServerInfo_UpdateEventData(out)

	}

}

func (o OutServerInfo) TssData(session *net.Session, out *server_info_api.TssData_Out) {

	if session.State.(*Client).OutServerInfo_TssData != nil {

		session.State.(*Client).OutServerInfo_TssData(out)

	}

}

func (o OutDebug) AddBuddy(session *net.Session, out *debug_api.AddBuddy_Out) {

	if session.State.(*Client).OutDebug_AddBuddy != nil {

		session.State.(*Client).OutDebug_AddBuddy(out)

	}

}

func (o OutDebug) AddItem(session *net.Session, out *debug_api.AddItem_Out) {

	if session.State.(*Client).OutDebug_AddItem != nil {

		session.State.(*Client).OutDebug_AddItem(out)

	}

}

func (o OutDebug) SetRoleLevel(session *net.Session, out *debug_api.SetRoleLevel_Out) {

	if session.State.(*Client).OutDebug_SetRoleLevel != nil {

		session.State.(*Client).OutDebug_SetRoleLevel(out)

	}

}

func (o OutDebug) SetCoins(session *net.Session, out *debug_api.SetCoins_Out) {

	if session.State.(*Client).OutDebug_SetCoins != nil {

		session.State.(*Client).OutDebug_SetCoins(out)

	}

}

func (o OutDebug) SetIngot(session *net.Session, out *debug_api.SetIngot_Out) {

	if session.State.(*Client).OutDebug_SetIngot != nil {

		session.State.(*Client).OutDebug_SetIngot(out)

	}

}

func (o OutDebug) AddGhost(session *net.Session, out *debug_api.AddGhost_Out) {

	if session.State.(*Client).OutDebug_AddGhost != nil {

		session.State.(*Client).OutDebug_AddGhost(out)

	}

}

func (o OutDebug) SetPlayerPhysical(session *net.Session, out *debug_api.SetPlayerPhysical_Out) {

	if session.State.(*Client).OutDebug_SetPlayerPhysical != nil {

		session.State.(*Client).OutDebug_SetPlayerPhysical(out)

	}

}

func (o OutDebug) ResetLevelEnterCount(session *net.Session, out *debug_api.ResetLevelEnterCount_Out) {

	if session.State.(*Client).OutDebug_ResetLevelEnterCount != nil {

		session.State.(*Client).OutDebug_ResetLevelEnterCount(out)

	}

}

func (o OutDebug) AddExp(session *net.Session, out *debug_api.AddExp_Out) {

	if session.State.(*Client).OutDebug_AddExp != nil {

		session.State.(*Client).OutDebug_AddExp(out)

	}

}

func (o OutDebug) OpenGhostMission(session *net.Session, out *debug_api.OpenGhostMission_Out) {

	if session.State.(*Client).OutDebug_OpenGhostMission != nil {

		session.State.(*Client).OutDebug_OpenGhostMission(out)

	}

}

func (o OutDebug) SendMail(session *net.Session, out *debug_api.SendMail_Out) {

	if session.State.(*Client).OutDebug_SendMail != nil {

		session.State.(*Client).OutDebug_SendMail(out)

	}

}

func (o OutDebug) ClearMail(session *net.Session, out *debug_api.ClearMail_Out) {

	if session.State.(*Client).OutDebug_ClearMail != nil {

		session.State.(*Client).OutDebug_ClearMail(out)

	}

}

func (o OutDebug) OpenMissionLevel(session *net.Session, out *debug_api.OpenMissionLevel_Out) {

	if session.State.(*Client).OutDebug_OpenMissionLevel != nil {

		session.State.(*Client).OutDebug_OpenMissionLevel(out)

	}

}

func (o OutDebug) StartBattle(session *net.Session, out *debug_api.StartBattle_Out) {

	if session.State.(*Client).OutDebug_StartBattle != nil {

		session.State.(*Client).OutDebug_StartBattle(out)

	}

}

func (o OutDebug) ListenByName(session *net.Session, out *debug_api.ListenByName_Out) {

	if session.State.(*Client).OutDebug_ListenByName != nil {

		session.State.(*Client).OutDebug_ListenByName(out)

	}

}

func (o OutDebug) OpenQuest(session *net.Session, out *debug_api.OpenQuest_Out) {

	if session.State.(*Client).OutDebug_OpenQuest != nil {

		session.State.(*Client).OutDebug_OpenQuest(out)

	}

}

func (o OutDebug) OpenFunc(session *net.Session, out *debug_api.OpenFunc_Out) {

	if session.State.(*Client).OutDebug_OpenFunc != nil {

		session.State.(*Client).OutDebug_OpenFunc(out)

	}

}

func (o OutDebug) AddSwordSoul(session *net.Session, out *debug_api.AddSwordSoul_Out) {

	if session.State.(*Client).OutDebug_AddSwordSoul != nil {

		session.State.(*Client).OutDebug_AddSwordSoul(out)

	}

}

func (o OutDebug) AddBattlePet(session *net.Session, out *debug_api.AddBattlePet_Out) {

	if session.State.(*Client).OutDebug_AddBattlePet != nil {

		session.State.(*Client).OutDebug_AddBattlePet(out)

	}

}

func (o OutDebug) ResetMultiLevelEnterCount(session *net.Session, out *debug_api.ResetMultiLevelEnterCount_Out) {

	if session.State.(*Client).OutDebug_ResetMultiLevelEnterCount != nil {

		session.State.(*Client).OutDebug_ResetMultiLevelEnterCount(out)

	}

}

func (o OutDebug) OpenMultiLevel(session *net.Session, out *debug_api.OpenMultiLevel_Out) {

	if session.State.(*Client).OutDebug_OpenMultiLevel != nil {

		session.State.(*Client).OutDebug_OpenMultiLevel(out)

	}

}

func (o OutDebug) OpenAllPetGrid(session *net.Session, out *debug_api.OpenAllPetGrid_Out) {

	if session.State.(*Client).OutDebug_OpenAllPetGrid != nil {

		session.State.(*Client).OutDebug_OpenAllPetGrid(out)

	}

}

func (o OutDebug) CreateAnnouncement(session *net.Session, out *debug_api.CreateAnnouncement_Out) {

	if session.State.(*Client).OutDebug_CreateAnnouncement != nil {

		session.State.(*Client).OutDebug_CreateAnnouncement(out)

	}

}

func (o OutDebug) AddHeart(session *net.Session, out *debug_api.AddHeart_Out) {

	if session.State.(*Client).OutDebug_AddHeart != nil {

		session.State.(*Client).OutDebug_AddHeart(out)

	}

}

func (o OutDebug) ResetHardLevelEnterCount(session *net.Session, out *debug_api.ResetHardLevelEnterCount_Out) {

	if session.State.(*Client).OutDebug_ResetHardLevelEnterCount != nil {

		session.State.(*Client).OutDebug_ResetHardLevelEnterCount(out)

	}

}

func (o OutDebug) OpenHardLevel(session *net.Session, out *debug_api.OpenHardLevel_Out) {

	if session.State.(*Client).OutDebug_OpenHardLevel != nil {

		session.State.(*Client).OutDebug_OpenHardLevel(out)

	}

}

func (o OutDebug) SetVipLevel(session *net.Session, out *debug_api.SetVipLevel_Out) {

	if session.State.(*Client).OutDebug_SetVipLevel != nil {

		session.State.(*Client).OutDebug_SetVipLevel(out)

	}

}

func (o OutDebug) SetResourceLevelOpenDay(session *net.Session, out *debug_api.SetResourceLevelOpenDay_Out) {

	if session.State.(*Client).OutDebug_SetResourceLevelOpenDay != nil {

		session.State.(*Client).OutDebug_SetResourceLevelOpenDay(out)

	}

}

func (o OutDebug) ResetResourceLevelOpenDay(session *net.Session, out *debug_api.ResetResourceLevelOpenDay_Out) {

	if session.State.(*Client).OutDebug_ResetResourceLevelOpenDay != nil {

		session.State.(*Client).OutDebug_ResetResourceLevelOpenDay(out)

	}

}

func (o OutDebug) ResetArenaDailyCount(session *net.Session, out *debug_api.ResetArenaDailyCount_Out) {

	if session.State.(*Client).OutDebug_ResetArenaDailyCount != nil {

		session.State.(*Client).OutDebug_ResetArenaDailyCount(out)

	}

}

func (o OutDebug) ResetSwordSoulDrawCd(session *net.Session, out *debug_api.ResetSwordSoulDrawCd_Out) {

	if session.State.(*Client).OutDebug_ResetSwordSoulDrawCd != nil {

		session.State.(*Client).OutDebug_ResetSwordSoulDrawCd(out)

	}

}

func (o OutDebug) SetFirstLoginTime(session *net.Session, out *debug_api.SetFirstLoginTime_Out) {

	if session.State.(*Client).OutDebug_SetFirstLoginTime != nil {

		session.State.(*Client).OutDebug_SetFirstLoginTime(out)

	}

}

func (o OutDebug) EarlierFirstLoginTime(session *net.Session, out *debug_api.EarlierFirstLoginTime_Out) {

	if session.State.(*Client).OutDebug_EarlierFirstLoginTime != nil {

		session.State.(*Client).OutDebug_EarlierFirstLoginTime(out)

	}

}

func (o OutDebug) ResetServerOpenTime(session *net.Session, out *debug_api.ResetServerOpenTime_Out) {

	if session.State.(*Client).OutDebug_ResetServerOpenTime != nil {

		session.State.(*Client).OutDebug_ResetServerOpenTime(out)

	}

}

func (o OutDebug) ClearTraderRefreshTime(session *net.Session, out *debug_api.ClearTraderRefreshTime_Out) {

	if session.State.(*Client).OutDebug_ClearTraderRefreshTime != nil {

		session.State.(*Client).OutDebug_ClearTraderRefreshTime(out)

	}

}

func (o OutDebug) AddTraderRefreshTime(session *net.Session, out *debug_api.AddTraderRefreshTime_Out) {

	if session.State.(*Client).OutDebug_AddTraderRefreshTime != nil {

		session.State.(*Client).OutDebug_AddTraderRefreshTime(out)

	}

}

func (o OutDebug) ClearTraderSchedule(session *net.Session, out *debug_api.ClearTraderSchedule_Out) {

	if session.State.(*Client).OutDebug_ClearTraderSchedule != nil {

		session.State.(*Client).OutDebug_ClearTraderSchedule(out)

	}

}

func (o OutDebug) AddTraderSchedule(session *net.Session, out *debug_api.AddTraderSchedule_Out) {

	if session.State.(*Client).OutDebug_AddTraderSchedule != nil {

		session.State.(*Client).OutDebug_AddTraderSchedule(out)

	}

}

func (o OutDebug) OpenTown(session *net.Session, out *debug_api.OpenTown_Out) {

	if session.State.(*Client).OutDebug_OpenTown != nil {

		session.State.(*Client).OutDebug_OpenTown(out)

	}

}

func (o OutDebug) AddGlobalMail(session *net.Session, out *debug_api.AddGlobalMail_Out) {

	if session.State.(*Client).OutDebug_AddGlobalMail != nil {

		session.State.(*Client).OutDebug_AddGlobalMail(out)

	}

}

func (o OutDebug) CreateAnnouncementWithoutTpl(session *net.Session, out *debug_api.CreateAnnouncementWithoutTpl_Out) {

	if session.State.(*Client).OutDebug_CreateAnnouncementWithoutTpl != nil {

		session.State.(*Client).OutDebug_CreateAnnouncementWithoutTpl(out)

	}

}

func (o OutDebug) SetLoginDay(session *net.Session, out *debug_api.SetLoginDay_Out) {

	if session.State.(*Client).OutDebug_SetLoginDay != nil {

		session.State.(*Client).OutDebug_SetLoginDay(out)

	}

}

func (o OutDebug) ResetLoginAward(session *net.Session, out *debug_api.ResetLoginAward_Out) {

	if session.State.(*Client).OutDebug_ResetLoginAward != nil {

		session.State.(*Client).OutDebug_ResetLoginAward(out)

	}

}

func (o OutDebug) RestPlayerAwardLock(session *net.Session, out *debug_api.RestPlayerAwardLock_Out) {

	if session.State.(*Client).OutDebug_RestPlayerAwardLock != nil {

		session.State.(*Client).OutDebug_RestPlayerAwardLock(out)

	}

}

func (o OutDebug) ResetRainbowLevel(session *net.Session, out *debug_api.ResetRainbowLevel_Out) {

	if session.State.(*Client).OutDebug_ResetRainbowLevel != nil {

		session.State.(*Client).OutDebug_ResetRainbowLevel(out)

	}

}

func (o OutDebug) SetRainbowLevel(session *net.Session, out *debug_api.SetRainbowLevel_Out) {

	if session.State.(*Client).OutDebug_SetRainbowLevel != nil {

		session.State.(*Client).OutDebug_SetRainbowLevel(out)

	}

}

func (o OutDebug) SendPushNotification(session *net.Session, out *debug_api.SendPushNotification_Out) {

	if session.State.(*Client).OutDebug_SendPushNotification != nil {

		session.State.(*Client).OutDebug_SendPushNotification(out)

	}

}

func (o OutDebug) ResetPetVirtualEnv(session *net.Session, out *debug_api.ResetPetVirtualEnv_Out) {

	if session.State.(*Client).OutDebug_ResetPetVirtualEnv != nil {

		session.State.(*Client).OutDebug_ResetPetVirtualEnv(out)

	}

}

func (o OutDebug) AddFame(session *net.Session, out *debug_api.AddFame_Out) {

	if session.State.(*Client).OutDebug_AddFame != nil {

		session.State.(*Client).OutDebug_AddFame(out)

	}

}

func (o OutDebug) AddWorldChatMessage(session *net.Session, out *debug_api.AddWorldChatMessage_Out) {

	if session.State.(*Client).OutDebug_AddWorldChatMessage != nil {

		session.State.(*Client).OutDebug_AddWorldChatMessage(out)

	}

}

func (o OutDebug) MonthCard(session *net.Session, out *debug_api.MonthCard_Out) {

	if session.State.(*Client).OutDebug_MonthCard != nil {

		session.State.(*Client).OutDebug_MonthCard(out)

	}

}

func (o OutDebug) EnterSandbox(session *net.Session, out *debug_api.EnterSandbox_Out) {

	if session.State.(*Client).OutDebug_EnterSandbox != nil {

		session.State.(*Client).OutDebug_EnterSandbox(out)

	}

}

func (o OutDebug) SandboxRollback(session *net.Session, out *debug_api.SandboxRollback_Out) {

	if session.State.(*Client).OutDebug_SandboxRollback != nil {

		session.State.(*Client).OutDebug_SandboxRollback(out)

	}

}

func (o OutDebug) ExitSandbox(session *net.Session, out *debug_api.ExitSandbox_Out) {

	if session.State.(*Client).OutDebug_ExitSandbox != nil {

		session.State.(*Client).OutDebug_ExitSandbox(out)

	}

}

func (o OutDebug) ResetShadedMissions(session *net.Session, out *debug_api.ResetShadedMissions_Out) {

	if session.State.(*Client).OutDebug_ResetShadedMissions != nil {

		session.State.(*Client).OutDebug_ResetShadedMissions(out)

	}

}

func (o OutDebug) CleanCornucopia(session *net.Session, out *debug_api.CleanCornucopia_Out) {

	if session.State.(*Client).OutDebug_CleanCornucopia != nil {

		session.State.(*Client).OutDebug_CleanCornucopia(out)

	}

}

func (o OutDebug) AddTotem(session *net.Session, out *debug_api.AddTotem_Out) {

	if session.State.(*Client).OutDebug_AddTotem != nil {

		session.State.(*Client).OutDebug_AddTotem(out)

	}

}

func (o OutDebug) AddRune(session *net.Session, out *debug_api.AddRune_Out) {

	if session.State.(*Client).OutDebug_AddRune != nil {

		session.State.(*Client).OutDebug_AddRune(out)

	}

}

func (o OutDebug) SendRareItemMessage(session *net.Session, out *debug_api.SendRareItemMessage_Out) {

	if session.State.(*Client).OutDebug_SendRareItemMessage != nil {

		session.State.(*Client).OutDebug_SendRareItemMessage(out)

	}

}

func (o OutDebug) AddSwordDrivingAction(session *net.Session, out *debug_api.AddSwordDrivingAction_Out) {

	if session.State.(*Client).OutDebug_AddSwordDrivingAction != nil {

		session.State.(*Client).OutDebug_AddSwordDrivingAction(out)

	}

}

func (o OutDebug) ResetDrivingSwordData(session *net.Session, out *debug_api.ResetDrivingSwordData_Out) {

	if session.State.(*Client).OutDebug_ResetDrivingSwordData != nil {

		session.State.(*Client).OutDebug_ResetDrivingSwordData(out)

	}

}

func (o OutDebug) AddSwordSoulFragment(session *net.Session, out *debug_api.AddSwordSoulFragment_Out) {

	if session.State.(*Client).OutDebug_AddSwordSoulFragment != nil {

		session.State.(*Client).OutDebug_AddSwordSoulFragment(out)

	}

}

func (o OutDebug) ResetMoneyTreeStatus(session *net.Session, out *debug_api.ResetMoneyTreeStatus_Out) {

	if session.State.(*Client).OutDebug_ResetMoneyTreeStatus != nil {

		session.State.(*Client).OutDebug_ResetMoneyTreeStatus(out)

	}

}

func (o OutDebug) ResetTodayMoneyTree(session *net.Session, out *debug_api.ResetTodayMoneyTree_Out) {

	if session.State.(*Client).OutDebug_ResetTodayMoneyTree != nil {

		session.State.(*Client).OutDebug_ResetTodayMoneyTree(out)

	}

}

func (o OutDebug) CleanSwordSoulIngotDrawNums(session *net.Session, out *debug_api.CleanSwordSoulIngotDrawNums_Out) {

	if session.State.(*Client).OutDebug_CleanSwordSoulIngotDrawNums != nil {

		session.State.(*Client).OutDebug_CleanSwordSoulIngotDrawNums(out)

	}

}

func (o OutDebug) PunchDrivingSwordCloud(session *net.Session, out *debug_api.PunchDrivingSwordCloud_Out) {

	if session.State.(*Client).OutDebug_PunchDrivingSwordCloud != nil {

		session.State.(*Client).OutDebug_PunchDrivingSwordCloud(out)

	}

}

func (o OutDebug) ClearCliqueDailyDonate(session *net.Session, out *debug_api.ClearCliqueDailyDonate_Out) {

	if session.State.(*Client).OutDebug_ClearCliqueDailyDonate != nil {

		session.State.(*Client).OutDebug_ClearCliqueDailyDonate(out)

	}

}

func (o OutDebug) SetCliqueContrib(session *net.Session, out *debug_api.SetCliqueContrib_Out) {

	if session.State.(*Client).OutDebug_SetCliqueContrib != nil {

		session.State.(*Client).OutDebug_SetCliqueContrib(out)

	}

}

func (o OutDebug) RefreshCliqueWorship(session *net.Session, out *debug_api.RefreshCliqueWorship_Out) {

	if session.State.(*Client).OutDebug_RefreshCliqueWorship != nil {

		session.State.(*Client).OutDebug_RefreshCliqueWorship(out)

	}

}

func (o OutDebug) CliqueEscortHijackBattleWin(session *net.Session, out *debug_api.CliqueEscortHijackBattleWin_Out) {

	if session.State.(*Client).OutDebug_CliqueEscortHijackBattleWin != nil {

		session.State.(*Client).OutDebug_CliqueEscortHijackBattleWin(out)

	}

}

func (o OutDebug) CliqueEscortRecoverBattleWin(session *net.Session, out *debug_api.CliqueEscortRecoverBattleWin_Out) {

	if session.State.(*Client).OutDebug_CliqueEscortRecoverBattleWin != nil {

		session.State.(*Client).OutDebug_CliqueEscortRecoverBattleWin(out)

	}

}

func (o OutDebug) CliqueEscortNotifyMessage(session *net.Session, out *debug_api.CliqueEscortNotifyMessage_Out) {

	if session.State.(*Client).OutDebug_CliqueEscortNotifyMessage != nil {

		session.State.(*Client).OutDebug_CliqueEscortNotifyMessage(out)

	}

}

func (o OutDebug) CliqueEscortNotifyDailyQuest(session *net.Session, out *debug_api.CliqueEscortNotifyDailyQuest_Out) {

	if session.State.(*Client).OutDebug_CliqueEscortNotifyDailyQuest != nil {

		session.State.(*Client).OutDebug_CliqueEscortNotifyDailyQuest(out)

	}

}

func (o OutDebug) SetCliqueBuildingLevel(session *net.Session, out *debug_api.SetCliqueBuildingLevel_Out) {

	if session.State.(*Client).OutDebug_SetCliqueBuildingLevel != nil {

		session.State.(*Client).OutDebug_SetCliqueBuildingLevel(out)

	}

}

func (o OutDebug) SetCliqueBuildingMoney(session *net.Session, out *debug_api.SetCliqueBuildingMoney_Out) {

	if session.State.(*Client).OutDebug_SetCliqueBuildingMoney != nil {

		session.State.(*Client).OutDebug_SetCliqueBuildingMoney(out)

	}

}

func (o OutDebug) EscortBench(session *net.Session, out *debug_api.EscortBench_Out) {

	if session.State.(*Client).OutDebug_EscortBench != nil {

		session.State.(*Client).OutDebug_EscortBench(out)

	}

}

func (o OutDebug) ResetCliqueEscortDailyNum(session *net.Session, out *debug_api.ResetCliqueEscortDailyNum_Out) {

	if session.State.(*Client).OutDebug_ResetCliqueEscortDailyNum != nil {

		session.State.(*Client).OutDebug_ResetCliqueEscortDailyNum(out)

	}

}

func (o OutDebug) TakeAdditionQuest(session *net.Session, out *debug_api.TakeAdditionQuest_Out) {

	if session.State.(*Client).OutDebug_TakeAdditionQuest != nil {

		session.State.(*Client).OutDebug_TakeAdditionQuest(out)

	}

}

func (o OutDebug) SetMissionStarMax(session *net.Session, out *debug_api.SetMissionStarMax_Out) {

	if session.State.(*Client).OutDebug_SetMissionStarMax != nil {

		session.State.(*Client).OutDebug_SetMissionStarMax(out)

	}

}

func (o OutDebug) CliqueBankCd(session *net.Session, out *debug_api.CliqueBankCd_Out) {

	if session.State.(*Client).OutDebug_CliqueBankCd != nil {

		session.State.(*Client).OutDebug_CliqueBankCd(out)

	}

}

func (o OutDebug) ResetDespairLandBattleNum(session *net.Session, out *debug_api.ResetDespairLandBattleNum_Out) {

	if session.State.(*Client).OutDebug_ResetDespairLandBattleNum != nil {

		session.State.(*Client).OutDebug_ResetDespairLandBattleNum(out)

	}

}

func (o OutDebug) ResetCliqueStoreSendTimes(session *net.Session, out *debug_api.ResetCliqueStoreSendTimes_Out) {

	if session.State.(*Client).OutDebug_ResetCliqueStoreSendTimes != nil {

		session.State.(*Client).OutDebug_ResetCliqueStoreSendTimes(out)

	}

}

func (o OutDebug) AddCliqueStoreDonate(session *net.Session, out *debug_api.AddCliqueStoreDonate_Out) {

	if session.State.(*Client).OutDebug_AddCliqueStoreDonate != nil {

		session.State.(*Client).OutDebug_AddCliqueStoreDonate(out)

	}

}

func (o OutDebug) PassAllDespairLandLevel(session *net.Session, out *debug_api.PassAllDespairLandLevel_Out) {

	if session.State.(*Client).OutDebug_PassAllDespairLandLevel != nil {

		session.State.(*Client).OutDebug_PassAllDespairLandLevel(out)

	}

}

func (o OutDebug) DespairLandDummyBossKill(session *net.Session, out *debug_api.DespairLandDummyBossKill_Out) {

	if session.State.(*Client).OutDebug_DespairLandDummyBossKill != nil {

		session.State.(*Client).OutDebug_DespairLandDummyBossKill(out)

	}

}

func (o OutDebug) AddTaoyuanItem(session *net.Session, out *debug_api.AddTaoyuanItem_Out) {

	if session.State.(*Client).OutDebug_AddTaoyuanItem != nil {

		session.State.(*Client).OutDebug_AddTaoyuanItem(out)

	}

}

func (o OutDebug) AddTaoyuanExp(session *net.Session, out *debug_api.AddTaoyuanExp_Out) {

	if session.State.(*Client).OutDebug_AddTaoyuanExp != nil {

		session.State.(*Client).OutDebug_AddTaoyuanExp(out)

	}

}

func (c *Client) Player_Info() {

	in := &player_api.Info_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Player_Relogin() {

	in := &player_api.Relogin_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Player_BuyPhysical() {

	in := &player_api.BuyPhysical_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Player_SetPlayKey(lock int16) {

	in := &player_api.SetPlayKey_In{

		Lock: lock,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Player_GetTime() {

	in := &player_api.GetTime_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Player_FromPlatformLogin(platform_type player_api.PlatformType, channel_id player_api.ChannelId, username string, nickname string, role_id int8, hashcode string, unixtime int64, restore bool, recvCount int64, gsid int32, openkey string, pay_token string, pfkey string, zoneid int32, pf string, channel int32, telecom_oper string, client_version string, system_hardware string, network string) {

	in := &player_api.FromPlatformLogin_In{

		PlatformType:   platform_type,
		ChannelId:      channel_id,
		Username:       []byte(username),
		Nickname:       []byte(nickname),
		RoleId:         role_id,
		Hashcode:       []byte(hashcode),
		Unixtime:       unixtime,
		Restore:        restore,
		RecvCount:      recvCount,
		Gsid:           gsid,
		Openkey:        []byte(openkey),
		PayToken:       []byte(pay_token),
		Pfkey:          []byte(pfkey),
		Zoneid:         zoneid,
		Pf:             []byte(pf),
		Channel:        channel,
		TelecomOper:    []byte(telecom_oper),
		ClientVersion:  []byte(client_version),
		SystemHardware: []byte(system_hardware),
		Network:        []byte(network),
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Player_BuyCoins(number int16) {

	in := &player_api.BuyCoins_In{

		Number: number,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Player_GetLoginInfo() {

	in := &player_api.GetLoginInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Player_CrossLoginGameServer(pid int64, openid string, nick string, role_id int8, role_level int16, time int64, hash string) {

	in := &player_api.CrossLoginGameServer_In{

		Pid:       pid,
		Openid:    []byte(openid),
		Nick:      []byte(nick),
		RoleId:    role_id,
		RoleLevel: role_level,
		Time:      time,
		Hash:      []byte(hash),
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Player_GlobalLogin(pid int64, openid string, nick string, role_id int8, role_level int16, fight_num int32, time int64, hash string, gsid int32, platform_type player_api.PlatformType, channel_id player_api.ChannelId, username string, openkey string, pay_token string, pfkey string, zoneid int32, pf string) {

	in := &player_api.GlobalLogin_In{

		Pid:          pid,
		Openid:       []byte(openid),
		Nick:         []byte(nick),
		RoleId:       role_id,
		RoleLevel:    role_level,
		FightNum:     fight_num,
		Time:         time,
		Hash:         []byte(hash),
		Gsid:         gsid,
		PlatformType: platform_type,
		ChannelId:    channel_id,
		Username:     []byte(username),
		Openkey:      []byte(openkey),
		PayToken:     []byte(pay_token),
		Pfkey:        []byte(pfkey),
		Zoneid:       zoneid,
		Pf:           []byte(pf),
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Player_GetIngot() {

	in := &player_api.GetIngot_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Player_SystemFame(system int16) {

	in := &player_api.SystemFame_In{

		System: system,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Player_GetRanks(flag player_api.RankType, page_index int32) {

	in := &player_api.GetRanks_In{

		Flag:      flag,
		PageIndex: page_index,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Town_Enter(town_id int16) {

	in := &town_api.Enter_In{

		TownId: town_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Town_Leave() {

	in := &town_api.Leave_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Town_Move(to_x int16, to_y int16) {

	in := &town_api.Move_In{

		ToX: to_x,
		ToY: to_y,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Town_TalkedNpcList(town_id int16) {

	in := &town_api.TalkedNpcList_In{

		TownId: town_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Town_NpcTalk(npc_id int32) {

	in := &town_api.NpcTalk_In{

		NpcId: npc_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Town_ListOpenedTownTreasures() {

	in := &town_api.ListOpenedTownTreasures_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Town_TakeTownTreasures(town_id int16) {

	in := &town_api.TakeTownTreasures_In{

		TownId: town_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Team_GetFormationInfo() {

	in := &team_api.GetFormationInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Team_UpFormation(role_id int8, pos int8) {

	in := &team_api.UpFormation_In{

		RoleId: role_id,
		Pos:    pos,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Team_DownFormation(pos int8) {

	in := &team_api.DownFormation_In{

		Pos: pos,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Team_SwapFormation(pos_from int8, pos_to int8) {

	in := &team_api.SwapFormation_In{

		PosFrom: pos_from,
		PosTo:   pos_to,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Team_ReplaceFormation(role_id int8, pos int8) {

	in := &team_api.ReplaceFormation_In{

		RoleId: role_id,
		Pos:    pos,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Team_TrainingTeamship(attr_ind int8) {

	in := &team_api.TrainingTeamship_In{

		AttrInd: attr_ind,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Role_GetAllRoles() {

	in := &role_api.GetAllRoles_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Role_GetRoleInfo(role_id int8) {

	in := &role_api.GetRoleInfo_In{

		RoleId: role_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Role_GetPlayerInfo(pid int64) {

	in := &role_api.GetPlayerInfo_In{

		Pid: pid,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Role_GetFightNum(fight_type role_api.FightnumType) {

	in := &role_api.GetFightNum_In{

		FightType: fight_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Role_GetPlayerInfoWithOpenid(openid string, game_server_id int32) {

	in := &role_api.GetPlayerInfoWithOpenid_In{

		Openid:       []byte(openid),
		GameServerId: game_server_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Role_LevelupRoleFriendship(role_id int8) {

	in := &role_api.LevelupRoleFriendship_In{

		RoleId: role_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Role_RecruitBuddy(role_id int8) {

	in := &role_api.RecruitBuddy_In{

		RoleId: role_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Role_GetRoleFightNum(role_id int8) {

	in := &role_api.GetRoleFightNum_In{

		RoleId: role_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Role_ChangeRoleStatus(role_id int8, status int8) {

	in := &role_api.ChangeRoleStatus_In{

		RoleId: role_id,
		Status: status,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Role_GetInnRoleList() {

	in := &role_api.GetInnRoleList_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Role_BuddyCoop(coop_id int16) {

	in := &role_api.BuddyCoop_In{

		CoopId: coop_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_Open(mission_id int16) {

	in := &mission_api.Open_In{

		MissionId: mission_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_GetMissionLevel(mission_id int16) {

	in := &mission_api.GetMissionLevel_In{

		MissionId: mission_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_EnterLevel(mission_level_id int32) {

	in := &mission_api.EnterLevel_In{

		MissionLevelId: mission_level_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_OpenLevelBox(times int32) {

	in := &mission_api.OpenLevelBox_In{

		Times: times,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_UseIngotRelive() {

	in := &mission_api.UseIngotRelive_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_UseItem(role_id int8, item_id int16) {

	in := &mission_api.UseItem_In{

		RoleId: role_id,
		ItemId: item_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_EnterExtendLevel(level_type mission_api.ExtendLevelType, level_id int32) {

	in := &mission_api.EnterExtendLevel_In{

		LevelType: level_type,
		LevelId:   level_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_GetExtendLevelInfo(level_type mission_api.ExtendType) {

	in := &mission_api.GetExtendLevelInfo_In{

		LevelType: level_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_OpenSmallBox(box_id int32) {

	in := &mission_api.OpenSmallBox_In{

		BoxId: box_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_EnterHardLevel(level_id int32) {

	in := &mission_api.EnterHardLevel_In{

		LevelId: level_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_GetHardLevel() {

	in := &mission_api.GetHardLevel_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_GetBuddyLevelRoleId() {

	in := &mission_api.GetBuddyLevelRoleId_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_SetBuddyLevelTeam(role_pos int8, buddy_pos int8, tactical int8) {

	in := &mission_api.SetBuddyLevelTeam_In{

		RolePos:  role_pos,
		BuddyPos: buddy_pos,
		Tactical: tactical,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_AutoFightLevel(level_type mission_api.BattleType, level_id int32, times int8) {

	in := &mission_api.AutoFightLevel_In{

		LevelType: level_type,
		LevelId:   level_id,
		Times:     times,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_EnterRainbowLevel(mission_level_id int32) {

	in := &mission_api.EnterRainbowLevel_In{

		MissionLevelId: mission_level_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_OpenMengYao(meng_yao_id int32) {

	in := &mission_api.OpenMengYao_In{

		MengYaoId: meng_yao_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_GetMissionLevelByItemId(item_id int16) {

	in := &mission_api.GetMissionLevelByItemId_In{

		ItemId: item_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_BuyHardLevelTimes(level_id int32) {

	in := &mission_api.BuyHardLevelTimes_In{

		LevelId: level_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_OpenRandomAwardBox(level_id int32) {

	in := &mission_api.OpenRandomAwardBox_In{

		LevelId: level_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_EvaluateLevel() {

	in := &mission_api.EvaluateLevel_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_OpenShadedBox(shaded_id int32) {

	in := &mission_api.OpenShadedBox_In{

		ShadedId: shaded_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_GetMissionTotalStarInfo(town_id int16) {

	in := &mission_api.GetMissionTotalStarInfo_In{

		TownId: town_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_GetMissionTotalStarAwards(town_id int16, box_type int8) {

	in := &mission_api.GetMissionTotalStarAwards_In{

		TownId:  town_id,
		BoxType: box_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_ClearMissionLevelState() {

	in := &mission_api.ClearMissionLevelState_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_BuyResourceMissionLevelTimes(sub_type int8) {

	in := &mission_api.BuyResourceMissionLevelTimes_In{

		SubType: sub_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_GetDragonBall() {

	in := &mission_api.GetDragonBall_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mission_GetEventItem() {

	in := &mission_api.GetEventItem_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Skill_GetAllSkillsInfo() {

	in := &skill_api.GetAllSkillsInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Skill_EquipSkill(role_id int8, order_number int8, skill_id int16) {

	in := &skill_api.EquipSkill_In{

		RoleId:      role_id,
		OrderNumber: order_number,
		SkillId:     skill_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Skill_UnequipSkill(role_id int8, order_number int8) {

	in := &skill_api.UnequipSkill_In{

		RoleId:      role_id,
		OrderNumber: order_number,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Skill_StudySkillByCheat(role_id int8, item_id int16) {

	in := &skill_api.StudySkillByCheat_In{

		RoleId: role_id,
		ItemId: item_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Skill_TrainSkill(role_id int8, skill_id int16) {

	in := &skill_api.TrainSkill_In{

		RoleId:  role_id,
		SkillId: skill_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Skill_FlushSkill(role_id int8) {

	in := &skill_api.FlushSkill_In{

		RoleId: role_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Battle_StartBattle(battle_type battle_api.BattleType, battle_id int64) {

	in := &battle_api.StartBattle_In{

		BattleType: battle_type,
		BattleId:   battle_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Battle_NextRound(use_skill int8, use_item int16, auto_fight bool, is_attacker bool, position int8, job_index int8, send_num int16, use_sword_soul bool, use_ghost_skill_position int8, use_ghost_skill_id int32, use_totem bool) {

	in := &battle_api.NextRound_In{

		UseSkill:              use_skill,
		UseItem:               use_item,
		AutoFight:             auto_fight,
		IsAttacker:            is_attacker,
		Position:              position,
		JobIndex:              job_index,
		SendNum:               send_num,
		UseSwordSoul:          use_sword_soul,
		UseGhostSkillPosition: use_ghost_skill_position,
		UseGhostSkillId:       use_ghost_skill_id,
		UseTotem:              use_totem,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Battle_Escape() {

	in := &battle_api.Escape_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Battle_StartReady(ok bool) {

	in := &battle_api.StartReady_In{

		Ok: ok,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Battle_CallBattlePet(grid_num int8) {

	in := &battle_api.CallBattlePet_In{

		GridNum: grid_num,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Battle_UseBuddySkill(pos int8, use_skill int8) {

	in := &battle_api.UseBuddySkill_In{

		Pos:      pos,
		UseSkill: use_skill,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Battle_StartBattleForHijackBoat(pid int64, boat_id int64) {

	in := &battle_api.StartBattleForHijackBoat_In{

		Pid:    pid,
		BoatId: boat_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Battle_StartBattleForRecoverBoat(pid int64, boat_id int64) {

	in := &battle_api.StartBattleForRecoverBoat_In{

		Pid:    pid,
		BoatId: boat_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Battle_RoundReady(is_auto bool) {

	in := &battle_api.RoundReady_In{

		IsAuto: is_auto,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Battle_InitRound() {

	in := &battle_api.InitRound_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Battle_SetAuto() {

	in := &battle_api.SetAuto_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Battle_CancelAuto() {

	in := &battle_api.CancelAuto_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Battle_SetSkill(is_attacker bool, pos_idx int8, skill_idx int8) {

	in := &battle_api.SetSkill_In{

		IsAttacker: is_attacker,
		PosIdx:     pos_idx,
		SkillIdx:   skill_idx,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Battle_UseItem(is_attacker bool, position int8, item_id int16) {

	in := &battle_api.UseItem_In{

		IsAttacker: is_attacker,
		Position:   position,
		ItemId:     item_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Battle_UseGhost(is_attacker bool, position int8) {

	in := &battle_api.UseGhost_In{

		IsAttacker: is_attacker,
		Position:   position,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Battle_BattleReconnect() {

	in := &battle_api.BattleReconnect_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_GetAllItems() {

	in := &item_api.GetAllItems_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_DropItem(id int64) {

	in := &item_api.DropItem_In{

		Id: id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_BuyItem(item_id int16) {

	in := &item_api.BuyItem_In{

		ItemId: item_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_SellItem(id int64) {

	in := &item_api.SellItem_In{

		Id: id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_Dress(role_id int8, id int64) {

	in := &item_api.Dress_In{

		RoleId: role_id,
		Id:     id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_Undress(role_id int8, pos item_api.EquipmentPos) {

	in := &item_api.Undress_In{

		RoleId: role_id,
		Pos:    pos,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_BuyItemBack(id int64) {

	in := &item_api.BuyItemBack_In{

		Id: id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_IsBagFull() {

	in := &item_api.IsBagFull_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_Decompose(id int64) {

	in := &item_api.Decompose_In{

		Id: id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_Refine(id int64, is_batch bool) {

	in := &item_api.Refine_In{

		Id:      id,
		IsBatch: is_batch,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_GetRecastInfo(id int64, attr item_api.Attribute) {

	in := &item_api.GetRecastInfo_In{

		Id:   id,
		Attr: attr,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_Recast(attr item_api.Attribute) {

	in := &item_api.Recast_In{

		Attr: attr,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_UseItem(id int64) {

	in := &item_api.UseItem_In{

		Id: id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_RoleUseCostItem(role_id int8, item_id int64) {

	in := &item_api.RoleUseCostItem_In{

		RoleId: role_id,
		ItemId: item_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_BatchUseItem(role_id int8, id int64, num int32) {

	in := &item_api.BatchUseItem_In{

		RoleId: role_id,
		Id:     id,
		Num:    num,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_OpenCornucopia(id int64) {

	in := &item_api.OpenCornucopia_In{

		Id: id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_GetSealedbooks(item_type int8) {

	in := &item_api.GetSealedbooks_In{

		ItemType: item_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_ActivationSealedbook(item_type int8, item_id int64) {

	in := &item_api.ActivationSealedbook_In{

		ItemType: item_type,
		ItemId:   item_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_ExchangeGhostCrystal(item_id int16, exchange_type int8) {

	in := &item_api.ExchangeGhostCrystal_In{

		ItemId:       item_id,
		ExchangeType: exchange_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_ExchangeShopItem(kind int8) {

	in := &item_api.ExchangeShopItem_In{

		Kind: kind,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_GetHoildayItemList() {

	in := &item_api.GetHoildayItemList_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_ExchangeHoildayItem(id int32) {

	in := &item_api.ExchangeHoildayItem_In{

		Id: id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Item_BatchExchangeDragonBall(ball_id int16) {

	in := &item_api.BatchExchangeDragonBall_In{

		BallId: ball_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Notify_PlayerKeyChanged() {

	in := &notify_api.PlayerKeyChanged_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Ghost_Info() {

	in := &ghost_api.Info_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Ghost_Equip(from_id int64, role_id int8, to_pos ghost_api.EquipPos) {

	in := &ghost_api.Equip_In{

		FromId: from_id,
		RoleId: role_id,
		ToPos:  to_pos,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Ghost_Unequip(role_id int8, from_id int64) {

	in := &ghost_api.Unequip_In{

		RoleId: role_id,
		FromId: from_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Ghost_Swap(role_id int8, id_bag int64, id_equip int64) {

	in := &ghost_api.Swap_In{

		RoleId:  role_id,
		IdBag:   id_bag,
		IdEquip: id_equip,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Ghost_EquipPosChange(role_id int8, from_id int64, to_pos ghost_api.EquipPos) {

	in := &ghost_api.EquipPosChange_In{

		RoleId: role_id,
		FromId: from_id,
		ToPos:  to_pos,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Ghost_Train(id int64) {

	in := &ghost_api.Train_In{

		Id: id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Ghost_Upgrade(id int64) {

	in := &ghost_api.Upgrade_In{

		Id: id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Ghost_Baptize(id int64) {

	in := &ghost_api.Baptize_In{

		Id: id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Ghost_Compose(ghost_id int16) {

	in := &ghost_api.Compose_In{

		GhostId: ghost_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Ghost_TrainSkill(id int64) {

	in := &ghost_api.TrainSkill_In{

		Id: id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Ghost_FlushAttr(id int64) {

	in := &ghost_api.FlushAttr_In{

		Id: id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Ghost_RelationGhost(master int64, slave int64) {

	in := &ghost_api.RelationGhost_In{

		Master: master,
		Slave:  slave,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) SwordSoul_Info() {

	in := &sword_soul_api.Info_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) SwordSoul_Draw(box sword_soul_api.Box, draw_type sword_soul_api.DrawType) {

	in := &sword_soul_api.Draw_In{

		Box:      box,
		DrawType: draw_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) SwordSoul_Upgrade(target_id int64, sword_souls []sword_soul_api.Upgrade_In_SwordSouls) {

	in := &sword_soul_api.Upgrade_In{

		TargetId:   target_id,
		SwordSouls: sword_souls,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) SwordSoul_Exchange(sword_soul_id int16) {

	in := &sword_soul_api.Exchange_In{

		SwordSoulId: sword_soul_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) SwordSoul_Equip(from_id int64, role_id int8, equip_pos sword_soul_api.EquipPos) {

	in := &sword_soul_api.Equip_In{

		FromId:   from_id,
		RoleId:   role_id,
		EquipPos: equip_pos,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) SwordSoul_Unequip(role_id int8, from_id int64) {

	in := &sword_soul_api.Unequip_In{

		RoleId: role_id,
		FromId: from_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) SwordSoul_EquipPosChange(role_id int8, from_id int64, to_pos sword_soul_api.EquipPos) {

	in := &sword_soul_api.EquipPosChange_In{

		RoleId: role_id,
		FromId: from_id,
		ToPos:  to_pos,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) SwordSoul_Swap(role_id int8, from_id int64, to_id int64) {

	in := &sword_soul_api.Swap_In{

		RoleId: role_id,
		FromId: from_id,
		ToId:   to_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) SwordSoul_IsBagFull() {

	in := &sword_soul_api.IsBagFull_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) SwordSoul_EmptyPosNum() {

	in := &sword_soul_api.EmptyPosNum_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mail_GetList() {

	in := &mail_api.GetList_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mail_Read(id int64) {

	in := &mail_api.Read_In{

		Id: id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mail_TakeAttachment(attachment_id int64) {

	in := &mail_api.TakeAttachment_In{

		AttachmentId: attachment_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mail_GetInfos() {

	in := &mail_api.GetInfos_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Mail_RequestGlobalMail() {

	in := &mail_api.RequestGlobalMail_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Quest_UpdateQuest() {

	in := &quest_api.UpdateQuest_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Quest_GetDailyInfo() {

	in := &quest_api.GetDailyInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Quest_AwardDaily(id int16) {

	in := &quest_api.AwardDaily_In{

		Id: id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Quest_Guide(guide_type quest_api.GuideType, action quest_api.GuideAction) {

	in := &quest_api.Guide_In{

		GuideType: guide_type,
		Action:    action,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Quest_GetExtendQuestInfoByNpcId(npc_id int32) {

	in := &quest_api.GetExtendQuestInfoByNpcId_In{

		NpcId: npc_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Quest_TakeExtendQuestAward(quest_id int32) {

	in := &quest_api.TakeExtendQuestAward_In{

		QuestId: quest_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Quest_GetPannelQuestInfo() {

	in := &quest_api.GetPannelQuestInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Quest_GiveUpAdditionQuest(quest_id int32) {

	in := &quest_api.GiveUpAdditionQuest_In{

		QuestId: quest_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Quest_TakeAdditionQuest(quest_id int32) {

	in := &quest_api.TakeAdditionQuest_In{

		QuestId: quest_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Quest_TakeAdditionQuestAward(quest_id int32) {

	in := &quest_api.TakeAdditionQuestAward_In{

		QuestId: quest_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Quest_GetAdditionQuest() {

	in := &quest_api.GetAdditionQuest_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Quest_RefreshAdditionQuest(quest_id int32) {

	in := &quest_api.RefreshAdditionQuest_In{

		QuestId: quest_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Quest_TakeQuestStarsAwaded(stars_level int32) {

	in := &quest_api.TakeQuestStarsAwaded_In{

		StarsLevel: stars_level,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Friend_GetFriendList() {

	in := &friend_api.GetFriendList_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Friend_ListenByNick(nick string) {

	in := &friend_api.ListenByNick_In{

		Nick: []byte(nick),
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Friend_CancelListen(pid int64) {

	in := &friend_api.CancelListen_In{

		Pid: pid,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Friend_SendHeart(nickname string, friend_type friend_api.FriendType, pid int64) {

	in := &friend_api.SendHeart_In{

		Nickname:   []byte(nickname),
		FriendType: friend_type,
		Pid:        pid,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Friend_Chat(pid int64, message string) {

	in := &friend_api.Chat_In{

		Pid:     pid,
		Message: []byte(message),
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Friend_GetChatHistory(pid int64) {

	in := &friend_api.GetChatHistory_In{

		Pid: pid,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Friend_Block(pid int64) {

	in := &friend_api.Block_In{

		Pid: pid,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Friend_CancelBlock(pid int64) {

	in := &friend_api.CancelBlock_In{

		Pid: pid,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Friend_CleanChatHistory(pid int64) {

	in := &friend_api.CleanChatHistory_In{

		Pid: pid,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Friend_CurrentPlatformFriendNum(num int32) {

	in := &friend_api.CurrentPlatformFriendNum_In{

		Num: num,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Friend_GetSendHeartList() {

	in := &friend_api.GetSendHeartList_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Friend_SendHeartToAllFriends(friend_type friend_api.FriendType, friends []friend_api.SendHeartToAllFriends_In_Friends) {

	in := &friend_api.SendHeartToAllFriends_In{

		FriendType: friend_type,
		Friends:    friends,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Tower_GetInfo() {

	in := &tower_api.GetInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Tower_UseLadder() {

	in := &tower_api.UseLadder_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) MultiLevel_CreateRoom(level_id int32) {

	in := &multi_level_api.CreateRoom_In{

		LevelId: level_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) MultiLevel_EnterRoom(room_id int64, room_token uint64) {

	in := &multi_level_api.EnterRoom_In{

		RoomId:    room_id,
		RoomToken: room_token,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) MultiLevel_LeaveRoom() {

	in := &multi_level_api.LeaveRoom_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) MultiLevel_ChangeBuddy(buddy_role_id int8) {

	in := &multi_level_api.ChangeBuddy_In{

		BuddyRoleId: buddy_role_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) MultiLevel_GetBaseInfo() {

	in := &multi_level_api.GetBaseInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) MultiLevel_GetOnlineFriend() {

	in := &multi_level_api.GetOnlineFriend_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) MultiLevel_InviteIntoRoom(pid int64) {

	in := &multi_level_api.InviteIntoRoom_In{

		Pid: pid,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) MultiLevel_RefuseRoomInvite(room_id int64, inviter_id int64) {

	in := &multi_level_api.RefuseRoomInvite_In{

		RoomId:    room_id,
		InviterId: inviter_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) MultiLevel_MatchPlayer(level int32) {

	in := &multi_level_api.MatchPlayer_In{

		Level: level,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) MultiLevel_MatchRoom(pos int8) {

	in := &multi_level_api.MatchRoom_In{

		Pos: pos,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) MultiLevel_CancelMatchRoom(match_token uint64) {

	in := &multi_level_api.CancelMatchRoom_In{

		MatchToken: match_token,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) MultiLevel_GetFormInfo() {

	in := &multi_level_api.GetFormInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) MultiLevel_SetForm(formation []multi_level_api.SetForm_In_Formation) {

	in := &multi_level_api.SetForm_In{

		Formation: formation,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) MultiLevel_GetBattleRoleInfo() {

	in := &multi_level_api.GetBattleRoleInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) MultiLevel_GetMatchInfo() {

	in := &multi_level_api.GetMatchInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) MultiLevel_CancelMatchPlayer() {

	in := &multi_level_api.CancelMatchPlayer_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) BattlePet_GetPetInfo() {

	in := &battle_pet_api.GetPetInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) BattlePet_SetPet(grid_num int8, pet_id int32) {

	in := &battle_pet_api.SetPet_In{

		GridNum: grid_num,
		PetId:   pet_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) BattlePet_SetPetSwap(from_grid_num int8, to_grid_num int8) {

	in := &battle_pet_api.SetPetSwap_In{

		FromGridNum: from_grid_num,
		ToGridNum:   to_grid_num,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) BattlePet_UpgradePet(pet_id int32) {

	in := &battle_pet_api.UpgradePet_In{

		PetId: pet_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) BattlePet_TrainingPetSkill(pet_id int32) {

	in := &battle_pet_api.TrainingPetSkill_In{

		PetId: pet_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Announcement_GetList() {

	in := &announcement_api.GetList_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Arena_Enter() {

	in := &arena_api.Enter_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Arena_GetRecords() {

	in := &arena_api.GetRecords_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Arena_AwardBox(num int8) {

	in := &arena_api.AwardBox_In{

		Num: num,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Arena_StartBattle(player_id int64, player_rank int32) {

	in := &arena_api.StartBattle_In{

		PlayerId:   player_id,
		PlayerRank: player_rank,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Arena_GetTopRank() {

	in := &arena_api.GetTopRank_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Arena_CleanFailedCdTime() {

	in := &arena_api.CleanFailedCdTime_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Vip_Info(player_id int64) {

	in := &vip_api.Info_In{

		PlayerId: player_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Vip_GetTotal() {

	in := &vip_api.GetTotal_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Vip_VipLevelTotal() {

	in := &vip_api.VipLevelTotal_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Vip_BuyTimes(buytype vip_api.BuyTimesType) {

	in := &vip_api.BuyTimes_In{

		Buytype: buytype,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Trader_Info() {

	in := &trader_api.Info_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Trader_StoreState(trader_id int16) {

	in := &trader_api.StoreState_In{

		TraderId: trader_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Trader_Buy(grid_id int32) {

	in := &trader_api.Buy_In{

		GridId: grid_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Trader_Refresh(trader_id int16) {

	in := &trader_api.Refresh_In{

		TraderId: trader_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DailySignIn_Info() {

	in := &daily_sign_in_api.Info_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DailySignIn_Sign() {

	in := &daily_sign_in_api.Sign_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DailySignIn_SignPastDay(index int8) {

	in := &daily_sign_in_api.SignPastDay_In{

		Index: index,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Rainbow_Info() {

	in := &rainbow_api.Info_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Rainbow_Reset() {

	in := &rainbow_api.Reset_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Rainbow_AwardInfo() {

	in := &rainbow_api.AwardInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Rainbow_TakeAward(pos1 int8, pos2 int8) {

	in := &rainbow_api.TakeAward_In{

		Pos1: pos1,
		Pos2: pos2,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Rainbow_JumpToSegment(segment int16) {

	in := &rainbow_api.JumpToSegment_In{

		Segment: segment,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Rainbow_AutoFight(segment int16) {

	in := &rainbow_api.AutoFight_In{

		Segment: segment,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_LoginAwardInfo() {

	in := &event_api.LoginAwardInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_TakeLoginAward(day int32) {

	in := &event_api.TakeLoginAward_In{

		Day: day,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_GetEvents() {

	in := &event_api.GetEvents_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_GetEventAward(event_id int16, page int32, param1 int32, param2 int32, param3 int32) {

	in := &event_api.GetEventAward_In{

		EventId: event_id,
		Page:    page,
		Param1:  param1,
		Param2:  param2,
		Param3:  param3,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_PlayerEventPhysicalCost() {

	in := &event_api.PlayerEventPhysicalCost_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_PlayerMonthCardInfo() {

	in := &event_api.PlayerMonthCardInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_GetSevenInfo() {

	in := &event_api.GetSevenInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_GetSevenAward() {

	in := &event_api.GetSevenAward_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_GetRichmanClubInfo() {

	in := &event_api.GetRichmanClubInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_GetRichmanClubAward(column int8, sequence int8) {

	in := &event_api.GetRichmanClubAward_In{

		Column:   column,
		Sequence: sequence,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_InfoShare(is_share bool) {

	in := &event_api.InfoShare_In{

		IsShare: is_share,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_InfoGroupBuy() {

	in := &event_api.InfoGroupBuy_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_GetIngotChangeTotal(is_in bool) {

	in := &event_api.GetIngotChangeTotal_In{

		IsIn: is_in,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_GetEventTotalAward(order int16) {

	in := &event_api.GetEventTotalAward_In{

		Order: order,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_GetEventArenaRank() {

	in := &event_api.GetEventArenaRank_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_GetEventTenDrawTimes() {

	in := &event_api.GetEventTenDrawTimes_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_GetEventRechargeAward(page int32, requireid int32) {

	in := &event_api.GetEventRechargeAward_In{

		Page:      page,
		Requireid: requireid,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_GetEventNewYear() {

	in := &event_api.GetEventNewYear_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_QqVipContinue(kind int8) {

	in := &event_api.QqVipContinue_In{

		Kind: kind,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_DailyOnlineInfo() {

	in := &event_api.DailyOnlineInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Event_TakeDailyOnlineAward() {

	in := &event_api.TakeDailyOnlineAward_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Fashion_FashionInfo() {

	in := &fashion_api.FashionInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Fashion_DressFashion(fashion_id int16, in_clubhouse bool) {

	in := &fashion_api.DressFashion_In{

		FashionId:   fashion_id,
		InClubhouse: in_clubhouse,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) PushNotify_PushInfo() {

	in := &push_notify_api.PushInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) PushNotify_PushNotificationSwitch(notification_id int32, turn_on bool) {

	in := &push_notify_api.PushNotificationSwitch_In{

		NotificationId: notification_id,
		TurnOn:         turn_on,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Meditation_MeditationInfo() {

	in := &meditation_api.MeditationInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Meditation_StartMeditation(in_clubhouse bool) {

	in := &meditation_api.StartMeditation_In{

		InClubhouse: in_clubhouse,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Meditation_StopMeditation(in_clubhouse bool) {

	in := &meditation_api.StopMeditation_In{

		InClubhouse: in_clubhouse,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) PetVirtualEnv_Info() {

	in := &pet_virtual_env_api.Info_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) PetVirtualEnv_TakeAward() {

	in := &pet_virtual_env_api.TakeAward_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) PetVirtualEnv_AutoFight(floor int16) {

	in := &pet_virtual_env_api.AutoFight_In{

		Floor: floor,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Channel_GetLatestWorldChannelMessage() {

	in := &channel_api.GetLatestWorldChannelMessage_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Channel_AddWorldChat(content string) {

	in := &channel_api.AddWorldChat_In{

		Content: []byte(content),
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Channel_WorldChannelInfo() {

	in := &channel_api.WorldChannelInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Channel_AddCliqueChat(content string) {

	in := &channel_api.AddCliqueChat_In{

		Content: []byte(content),
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Channel_GetLatestCliqueMessages() {

	in := &channel_api.GetLatestCliqueMessages_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DrivingSword_CloudMapInfo() {

	in := &driving_sword_api.CloudMapInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DrivingSword_CloudClimb() {

	in := &driving_sword_api.CloudClimb_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DrivingSword_CloudTeleport(cloud int16) {

	in := &driving_sword_api.CloudTeleport_In{

		Cloud: cloud,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DrivingSword_AreaTeleport() {

	in := &driving_sword_api.AreaTeleport_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DrivingSword_AreaMove(direction driving_sword_api.MovingDirection) {

	in := &driving_sword_api.AreaMove_In{

		Direction: direction,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DrivingSword_ExplorerStartBattle() {

	in := &driving_sword_api.ExplorerStartBattle_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DrivingSword_ExplorerAward() {

	in := &driving_sword_api.ExplorerAward_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DrivingSword_ExplorerGarrison(role_id int8) {

	in := &driving_sword_api.ExplorerGarrison_In{

		RoleId: role_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DrivingSword_VisitMountain() {

	in := &driving_sword_api.VisitMountain_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DrivingSword_VisiterStartBattle() {

	in := &driving_sword_api.VisiterStartBattle_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DrivingSword_VisiterAward() {

	in := &driving_sword_api.VisiterAward_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DrivingSword_MountainTreasureOpen() {

	in := &driving_sword_api.MountainTreasureOpen_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DrivingSword_ListGarrisons() {

	in := &driving_sword_api.ListGarrisons_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DrivingSword_AwardGarrison(role_id int8) {

	in := &driving_sword_api.AwardGarrison_In{

		RoleId: role_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DrivingSword_EndGarrison(role_id int8) {

	in := &driving_sword_api.EndGarrison_In{

		RoleId: role_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DrivingSword_BuyAllowedAction() {

	in := &driving_sword_api.BuyAllowedAction_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Totem_Info() {

	in := &totem_api.Info_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Totem_CallTotem(call_type totem_api.CallType) {

	in := &totem_api.CallTotem_In{

		CallType: call_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Totem_Upgrade(target_id int64) {

	in := &totem_api.Upgrade_In{

		TargetId: target_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Totem_Decompose(totem_id int64) {

	in := &totem_api.Decompose_In{

		TotemId: totem_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Totem_Equip(totem_id int64, equip_pos totem_api.EquipPos) {

	in := &totem_api.Equip_In{

		TotemId:  totem_id,
		EquipPos: equip_pos,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Totem_Unequip(equip_pos totem_api.EquipPos) {

	in := &totem_api.Unequip_In{

		EquipPos: equip_pos,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Totem_EquipPosChange(from_pos totem_api.EquipPos, to_pos totem_api.EquipPos) {

	in := &totem_api.EquipPosChange_In{

		FromPos: from_pos,
		ToPos:   to_pos,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Totem_Swap(equiped_id int64, inbag_id int64) {

	in := &totem_api.Swap_In{

		EquipedId: equiped_id,
		InbagId:   inbag_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Totem_IsBagFull() {

	in := &totem_api.IsBagFull_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) MoneyTree_GetTreeStatus() {

	in := &money_tree_api.GetTreeStatus_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) MoneyTree_GetTreeMoney() {

	in := &money_tree_api.GetTreeMoney_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) MoneyTree_WaveTree() {

	in := &money_tree_api.WaveTree_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_CreateClique(name string, announce string) {

	in := &clique_api.CreateClique_In{

		Name:     []byte(name),
		Announce: []byte(announce),
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_ApplyJoinClique(clique_id int64) {

	in := &clique_api.ApplyJoinClique_In{

		CliqueId: clique_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_CancelApplyClique(clique_id int64) {

	in := &clique_api.CancelApplyClique_In{

		CliqueId: clique_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_ProcessJoinApply(agree bool, pidlist []clique_api.ProcessJoinApply_In_Pidlist) {

	in := &clique_api.ProcessJoinApply_In{

		Agree:   agree,
		Pidlist: pidlist,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_ElectOwner() {

	in := &clique_api.ElectOwner_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_MangeMember(pid int64, action clique_api.MangeMemberAction) {

	in := &clique_api.MangeMember_In{

		Pid:    pid,
		Action: action,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_DestoryClique() {

	in := &clique_api.DestoryClique_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_UpdateAnnounce(content string) {

	in := &clique_api.UpdateAnnounce_In{

		Content: []byte(content),
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_LeaveClique() {

	in := &clique_api.LeaveClique_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_ListClique(offset int16, limit int16) {

	in := &clique_api.ListClique_In{

		Offset: offset,
		Limit:  limit,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_CliquePublicInfo(clique_id int64) {

	in := &clique_api.CliquePublicInfo_In{

		CliqueId: clique_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_CliqueInfo() {

	in := &clique_api.CliqueInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_ListApply(offset int16, limit int16) {

	in := &clique_api.ListApply_In{

		Offset: offset,
		Limit:  limit,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_EnterClubhouse() {

	in := &clique_api.EnterClubhouse_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_LeaveClubhouse() {

	in := &clique_api.LeaveClubhouse_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_ClubMove(to_x int16, to_y int16) {

	in := &clique_api.ClubMove_In{

		ToX: to_x,
		ToY: to_y,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_CliquePublicInfoSummary(clique_id int64) {

	in := &clique_api.CliquePublicInfoSummary_In{

		CliqueId: clique_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_CliqueAutoAudit(level int16, enable bool) {

	in := &clique_api.CliqueAutoAudit_In{

		Level:  level,
		Enable: enable,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_CliqueBaseDonate() {

	in := &clique_api.CliqueBaseDonate_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_CliqueRecruitment() {

	in := &clique_api.CliqueRecruitment_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_QuickApply() {

	in := &clique_api.QuickApply_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Clique_OtherClique(page int16) {

	in := &clique_api.OtherClique_In{

		Page: page,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueBuilding_CliqueBaseDonate(money int64) {

	in := &clique_building_api.CliqueBaseDonate_In{

		Money: money,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueBuilding_CliqueBuildingStatus() {

	in := &clique_building_api.CliqueBuildingStatus_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueBuilding_CliqueBankDonate(money int64) {

	in := &clique_building_api.CliqueBankDonate_In{

		Money: money,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueBuilding_CliqueBankBuy(kind int8, num int8) {

	in := &clique_building_api.CliqueBankBuy_In{

		Kind: kind,
		Num:  num,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueBuilding_CliqueBankSold(kind int8) {

	in := &clique_building_api.CliqueBankSold_In{

		Kind: kind,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueBuilding_CliqueKongfuDonate(building int32, money int64) {

	in := &clique_building_api.CliqueKongfuDonate_In{

		Building: building,
		Money:    money,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueBuilding_CliqueKongfuInfo(building int32) {

	in := &clique_building_api.CliqueKongfuInfo_In{

		Building: building,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueBuilding_CliqueKongfuTrain(kongfu_id int32) {

	in := &clique_building_api.CliqueKongfuTrain_In{

		KongfuId: kongfu_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueBuilding_CliqueTempleWorship(worship_type clique_building_api.AncestralHallWorship) {

	in := &clique_building_api.CliqueTempleWorship_In{

		WorshipType: worship_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueBuilding_CliqueTempleDonate(money int64) {

	in := &clique_building_api.CliqueTempleDonate_In{

		Money: money,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueBuilding_CliqueTempleInfo() {

	in := &clique_building_api.CliqueTempleInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueBuilding_CliqueStoreDonate(money int64) {

	in := &clique_building_api.CliqueStoreDonate_In{

		Money: money,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueBuilding_CliqueStoreInfo() {

	in := &clique_building_api.CliqueStoreInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueBuilding_CliqueStoreSendChest(store_chest_type clique_building_api.CliqueStoreChest) {

	in := &clique_building_api.CliqueStoreSendChest_In{

		StoreChestType: store_chest_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueQuest_GetCliqueDailyQuest() {

	in := &clique_quest_api.GetCliqueDailyQuest_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueQuest_AwardCliqueDailyQuest(id int16) {

	in := &clique_quest_api.AwardCliqueDailyQuest_In{

		Id: id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueQuest_GetCliqueBuildingQuest() {

	in := &clique_quest_api.GetCliqueBuildingQuest_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueQuest_AwardCliqueBuildingQuest(id int16) {

	in := &clique_quest_api.AwardCliqueBuildingQuest_In{

		Id: id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueEscort_EscortInfo() {

	in := &clique_escort_api.EscortInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueEscort_GetIngotBoat() {

	in := &clique_escort_api.GetIngotBoat_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueEscort_StartEscort() {

	in := &clique_escort_api.StartEscort_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueEscort_HijackBoat(boat_id int64, pay bool) {

	in := &clique_escort_api.HijackBoat_In{

		BoatId: boat_id,
		Pay:    pay,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueEscort_RecoverBoat(boat_id int64) {

	in := &clique_escort_api.RecoverBoat_In{

		BoatId: boat_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueEscort_ListBoats() {

	in := &clique_escort_api.ListBoats_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueEscort_GetRandomBoat() {

	in := &clique_escort_api.GetRandomBoat_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueEscort_TakeHijackAward() {

	in := &clique_escort_api.TakeHijackAward_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueEscort_TakeEscortAward() {

	in := &clique_escort_api.TakeEscortAward_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueEscort_GetCliqueBoatMessages() {

	in := &clique_escort_api.GetCliqueBoatMessages_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) CliqueEscort_ReadCliqueBoatMessage(id int64) {

	in := &clique_escort_api.ReadCliqueBoatMessage_In{

		Id: id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DespairLand_DespairLandInfo() {

	in := &despair_land_api.DespairLandInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DespairLand_DespairLandCampInfo(camp_type despair_land_api.DespairLandCamp) {

	in := &despair_land_api.DespairLandCampInfo_In{

		CampType: camp_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DespairLand_DespairLandCampPlayerInfo(level_id int32) {

	in := &despair_land_api.DespairLandCampPlayerInfo_In{

		LevelId: level_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DespairLand_DespairLandPickBox(level_id int32, camp_type despair_land_api.DespairLandCamp) {

	in := &despair_land_api.DespairLandPickBox_In{

		LevelId:  level_id,
		CampType: camp_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DespairLand_DespairLandWatchRecord(record_type despair_land_api.DespairLandBattleRecordType, camp_type despair_land_api.DespairLandCamp, level_id int32) {

	in := &despair_land_api.DespairLandWatchRecord_In{

		RecordType: record_type,
		CampType:   camp_type,
		LevelId:    level_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DespairLand_DespairLandBuyBattleNum() {

	in := &despair_land_api.DespairLandBuyBattleNum_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DespairLand_DespairLandBossInfo(camp_type despair_land_api.DespairLandCamp) {

	in := &despair_land_api.DespairLandBossInfo_In{

		CampType: camp_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DespairLand_DespairLandBuyBossBattleNum(camp_type despair_land_api.DespairLandCamp) {

	in := &despair_land_api.DespairLandBuyBossBattleNum_In{

		CampType: camp_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DespairLand_DespairLandPickThreeStarBox(level_id int32, camp_type despair_land_api.DespairLandCamp) {

	in := &despair_land_api.DespairLandPickThreeStarBox_In{

		LevelId:  level_id,
		CampType: camp_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DespairLand_DespairLandBattleBossAwardInfo(camp_type despair_land_api.DespairLandCamp) {

	in := &despair_land_api.DespairLandBattleBossAwardInfo_In{

		CampType: camp_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) DespairLand_DespairLandBossOpenInfo() {

	in := &despair_land_api.DespairLandBossOpenInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Awaken_AwakenInfo(role_id int8) {

	in := &awaken_api.AwakenInfo_In{

		RoleId: role_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Awaken_LevelupAttr(role_id int8, attr_impl int16) {

	in := &awaken_api.LevelupAttr_In{

		RoleId:   role_id,
		AttrImpl: attr_impl,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_TaoyuanInfo() {

	in := &taoyuan_api.TaoyuanInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_Bless(other_pid int64) {

	in := &taoyuan_api.Bless_In{

		OtherPid: other_pid,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_ShopBuy(item_id int16, num int16) {

	in := &taoyuan_api.ShopBuy_In{

		ItemId: item_id,
		Num:    num,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_ShopSell(id int64, num int16) {

	in := &taoyuan_api.ShopSell_In{

		Id:  id,
		Num: num,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_GetAllItems() {

	in := &taoyuan_api.GetAllItems_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_FiledsInfo() {

	in := &taoyuan_api.FiledsInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_GrowPlant(filed_id int16, plant_id int16) {

	in := &taoyuan_api.GrowPlant_In{

		FiledId: filed_id,
		PlantId: plant_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_TakePlant(filed_id int16) {

	in := &taoyuan_api.TakePlant_In{

		FiledId: filed_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_UpgradeFiled(filed_id int16) {

	in := &taoyuan_api.UpgradeFiled_In{

		FiledId: filed_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_OpenFiled(filed_id int16) {

	in := &taoyuan_api.OpenFiled_In{

		FiledId: filed_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_StudySkill(skill_id int16) {

	in := &taoyuan_api.StudySkill_In{

		SkillId: skill_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_MakeProduct(item_id int16, product_type int8) {

	in := &taoyuan_api.MakeProduct_In{

		ItemId:      item_id,
		ProductType: product_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_TakeProduct(product_type int8, is_ingot bool) {

	in := &taoyuan_api.TakeProduct_In{

		ProductType: product_type,
		IsIngot:     is_ingot,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_ProductMakeQueue(product_type int8) {

	in := &taoyuan_api.ProductMakeQueue_In{

		ProductType: product_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_QuestInfo() {

	in := &taoyuan_api.QuestInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_QuestFinish(quset_id int16) {

	in := &taoyuan_api.QuestFinish_In{

		QusetId: quset_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_QuestRefresh() {

	in := &taoyuan_api.QuestRefresh_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_FriendTaoyuanInfo(pid int64) {

	in := &taoyuan_api.FriendTaoyuanInfo_In{

		Pid: pid,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_SkillInfo() {

	in := &taoyuan_api.SkillInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_OpenQueue(product_type int8) {

	in := &taoyuan_api.OpenQueue_In{

		ProductType: product_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_PlantQuicklyMaturity(filed_id int16) {

	in := &taoyuan_api.PlantQuicklyMaturity_In{

		FiledId: filed_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_TaoyuanMessageInfo() {

	in := &taoyuan_api.TaoyuanMessageInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_TaoyuanMessageRead(id int64) {

	in := &taoyuan_api.TaoyuanMessageRead_In{

		Id: id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Taoyuan_OpenProductBuilding(product_type int8) {

	in := &taoyuan_api.OpenProductBuilding_In{

		ProductType: product_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Draw_GetHeartDrawInfo(draw_type int8, award_record bool) {

	in := &draw_api.GetHeartDrawInfo_In{

		DrawType:    draw_type,
		AwardRecord: award_record,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Draw_HeartDraw(draw_type int8) {

	in := &draw_api.HeartDraw_In{

		DrawType: draw_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Draw_GetChestInfo() {

	in := &draw_api.GetChestInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Draw_DrawChest(chest_type draw_api.ChestType) {

	in := &draw_api.DrawChest_In{

		ChestType: chest_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Draw_HeartInfo() {

	in := &draw_api.HeartInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Draw_GetFateBoxInfo() {

	in := &draw_api.GetFateBoxInfo_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Draw_OpenFateBox(box_type int32, times int8) {

	in := &draw_api.OpenFateBox_In{

		BoxType: box_type,
		Times:   times,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Draw_ExchangeGiftCode(code string) {

	in := &draw_api.ExchangeGiftCode_In{

		Code: []byte(code),
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) ServerInfo_GetVersion() {

	in := &server_info_api.GetVersion_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) ServerInfo_GetApiCount() {

	in := &server_info_api.GetApiCount_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) ServerInfo_SearchPlayerRole(openid string) {

	in := &server_info_api.SearchPlayerRole_In{

		Openid: []byte(openid),
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) ServerInfo_UpdateAccessToken(token string, pfkey string) {

	in := &server_info_api.UpdateAccessToken_In{

		Token: []byte(token),
		Pfkey: []byte(pfkey),
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) ServerInfo_UpdateEventData(version int32) {

	in := &server_info_api.UpdateEventData_In{

		Version: version,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) ServerInfo_TssData(data string) {

	in := &server_info_api.TssData_In{

		Data: []byte(data),
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_AddBuddy(role_id int8) {

	in := &debug_api.AddBuddy_In{

		RoleId: role_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_AddItem(item_id int16, number int16) {

	in := &debug_api.AddItem_In{

		ItemId: item_id,
		Number: number,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_SetRoleLevel(role_id int8, level int16) {

	in := &debug_api.SetRoleLevel_In{

		RoleId: role_id,
		Level:  level,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_SetCoins(number int64) {

	in := &debug_api.SetCoins_In{

		Number: number,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_SetIngot(number int64) {

	in := &debug_api.SetIngot_In{

		Number: number,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_AddGhost(ghost_id int16) {

	in := &debug_api.AddGhost_In{

		GhostId: ghost_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_SetPlayerPhysical(physical int16) {

	in := &debug_api.SetPlayerPhysical_In{

		Physical: physical,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ResetLevelEnterCount(level_id int32) {

	in := &debug_api.ResetLevelEnterCount_In{

		LevelId: level_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_AddExp(role_id int8, add_exp int64) {

	in := &debug_api.AddExp_In{

		RoleId: role_id,
		AddExp: add_exp,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_OpenGhostMission(mission_id int16) {

	in := &debug_api.OpenGhostMission_In{

		MissionId: mission_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_SendMail(mail_id int32) {

	in := &debug_api.SendMail_In{

		MailId: mail_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ClearMail() {

	in := &debug_api.ClearMail_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_OpenMissionLevel(level_id int32) {

	in := &debug_api.OpenMissionLevel_In{

		LevelId: level_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_StartBattle(battle_type int8, enemy_id int32) {

	in := &debug_api.StartBattle_In{

		BattleType: battle_type,
		EnemyId:    enemy_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ListenByName(name string) {

	in := &debug_api.ListenByName_In{

		Name: []byte(name),
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_OpenQuest(quest_id int16) {

	in := &debug_api.OpenQuest_In{

		QuestId: quest_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_OpenFunc(lock int16) {

	in := &debug_api.OpenFunc_In{

		Lock: lock,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_AddSwordSoul(sword_soul_id int16) {

	in := &debug_api.AddSwordSoul_In{

		SwordSoulId: sword_soul_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_AddBattlePet(petId int16) {

	in := &debug_api.AddBattlePet_In{

		PetId: petId,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ResetMultiLevelEnterCount() {

	in := &debug_api.ResetMultiLevelEnterCount_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_OpenMultiLevel(level_id int16) {

	in := &debug_api.OpenMultiLevel_In{

		LevelId: level_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_OpenAllPetGrid() {

	in := &debug_api.OpenAllPetGrid_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_CreateAnnouncement() {

	in := &debug_api.CreateAnnouncement_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_AddHeart(number int16) {

	in := &debug_api.AddHeart_In{

		Number: number,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ResetHardLevelEnterCount() {

	in := &debug_api.ResetHardLevelEnterCount_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_OpenHardLevel(level_id int32) {

	in := &debug_api.OpenHardLevel_In{

		LevelId: level_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_SetVipLevel(level int16) {

	in := &debug_api.SetVipLevel_In{

		Level: level,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_SetResourceLevelOpenDay(level_type int8, open_day int8) {

	in := &debug_api.SetResourceLevelOpenDay_In{

		LevelType: level_type,
		OpenDay:   open_day,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ResetResourceLevelOpenDay() {

	in := &debug_api.ResetResourceLevelOpenDay_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ResetArenaDailyCount() {

	in := &debug_api.ResetArenaDailyCount_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ResetSwordSoulDrawCd() {

	in := &debug_api.ResetSwordSoulDrawCd_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_SetFirstLoginTime(timestamp int64) {

	in := &debug_api.SetFirstLoginTime_In{

		Timestamp: timestamp,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_EarlierFirstLoginTime() {

	in := &debug_api.EarlierFirstLoginTime_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ResetServerOpenTime() {

	in := &debug_api.ResetServerOpenTime_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ClearTraderRefreshTime(trader_id int16) {

	in := &debug_api.ClearTraderRefreshTime_In{

		TraderId: trader_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_AddTraderRefreshTime(trader_id int16, timing int64) {

	in := &debug_api.AddTraderRefreshTime_In{

		TraderId: trader_id,
		Timing:   timing,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ClearTraderSchedule(trader_id int16) {

	in := &debug_api.ClearTraderSchedule_In{

		TraderId: trader_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_AddTraderSchedule(trader_id int16, expire int64, showup int64, disappear int64) {

	in := &debug_api.AddTraderSchedule_In{

		TraderId:  trader_id,
		Expire:    expire,
		Showup:    showup,
		Disappear: disappear,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_OpenTown(town_id int16) {

	in := &debug_api.OpenTown_In{

		TownId: town_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_AddGlobalMail(send_delay int64, expire_delay int64) {

	in := &debug_api.AddGlobalMail_In{

		SendDelay:   send_delay,
		ExpireDelay: expire_delay,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_CreateAnnouncementWithoutTpl() {

	in := &debug_api.CreateAnnouncementWithoutTpl_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_SetLoginDay(days int32) {

	in := &debug_api.SetLoginDay_In{

		Days: days,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ResetLoginAward() {

	in := &debug_api.ResetLoginAward_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_RestPlayerAwardLock() {

	in := &debug_api.RestPlayerAwardLock_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ResetRainbowLevel() {

	in := &debug_api.ResetRainbowLevel_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_SetRainbowLevel(segment int16, order int8) {

	in := &debug_api.SetRainbowLevel_In{

		Segment: segment,
		Order:   order,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_SendPushNotification() {

	in := &debug_api.SendPushNotification_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ResetPetVirtualEnv() {

	in := &debug_api.ResetPetVirtualEnv_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_AddFame(system int16, val int32) {

	in := &debug_api.AddFame_In{

		System: system,
		Val:    val,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_AddWorldChatMessage(num int16) {

	in := &debug_api.AddWorldChatMessage_In{

		Num: num,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_MonthCard() {

	in := &debug_api.MonthCard_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_EnterSandbox() {

	in := &debug_api.EnterSandbox_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_SandboxRollback() {

	in := &debug_api.SandboxRollback_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ExitSandbox() {

	in := &debug_api.ExitSandbox_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ResetShadedMissions(level_id int32) {

	in := &debug_api.ResetShadedMissions_In{

		LevelId: level_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_CleanCornucopia() {

	in := &debug_api.CleanCornucopia_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_AddTotem(totem_id int16) {

	in := &debug_api.AddTotem_In{

		TotemId: totem_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_AddRune(jade_num int32, rock_num int32) {

	in := &debug_api.AddRune_In{

		JadeNum: jade_num,
		RockNum: rock_num,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_SendRareItemMessage() {

	in := &debug_api.SendRareItemMessage_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_AddSwordDrivingAction(point int16) {

	in := &debug_api.AddSwordDrivingAction_In{

		Point: point,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ResetDrivingSwordData(cloud int16) {

	in := &debug_api.ResetDrivingSwordData_In{

		Cloud: cloud,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_AddSwordSoulFragment(number int64) {

	in := &debug_api.AddSwordSoulFragment_In{

		Number: number,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ResetMoneyTreeStatus() {

	in := &debug_api.ResetMoneyTreeStatus_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ResetTodayMoneyTree() {

	in := &debug_api.ResetTodayMoneyTree_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_CleanSwordSoulIngotDrawNums() {

	in := &debug_api.CleanSwordSoulIngotDrawNums_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_PunchDrivingSwordCloud() {

	in := &debug_api.PunchDrivingSwordCloud_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ClearCliqueDailyDonate() {

	in := &debug_api.ClearCliqueDailyDonate_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_SetCliqueContrib(contrib int64) {

	in := &debug_api.SetCliqueContrib_In{

		Contrib: contrib,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_RefreshCliqueWorship() {

	in := &debug_api.RefreshCliqueWorship_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_CliqueEscortHijackBattleWin(boat_id int64) {

	in := &debug_api.CliqueEscortHijackBattleWin_In{

		BoatId: boat_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_CliqueEscortRecoverBattleWin(boat_id int64) {

	in := &debug_api.CliqueEscortRecoverBattleWin_In{

		BoatId: boat_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_CliqueEscortNotifyMessage() {

	in := &debug_api.CliqueEscortNotifyMessage_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_CliqueEscortNotifyDailyQuest() {

	in := &debug_api.CliqueEscortNotifyDailyQuest_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_SetCliqueBuildingLevel(building int32, level int16) {

	in := &debug_api.SetCliqueBuildingLevel_In{

		Building: building,
		Level:    level,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_SetCliqueBuildingMoney(building int32, money int64) {

	in := &debug_api.SetCliqueBuildingMoney_In{

		Building: building,
		Money:    money,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_EscortBench() {

	in := &debug_api.EscortBench_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ResetCliqueEscortDailyNum() {

	in := &debug_api.ResetCliqueEscortDailyNum_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_TakeAdditionQuest(first_quest_id int32) {

	in := &debug_api.TakeAdditionQuest_In{

		FirstQuestId: first_quest_id,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_SetMissionStarMax() {

	in := &debug_api.SetMissionStarMax_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_CliqueBankCd() {

	in := &debug_api.CliqueBankCd_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ResetDespairLandBattleNum() {

	in := &debug_api.ResetDespairLandBattleNum_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_ResetCliqueStoreSendTimes() {

	in := &debug_api.ResetCliqueStoreSendTimes_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_AddCliqueStoreDonate() {

	in := &debug_api.AddCliqueStoreDonate_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_PassAllDespairLandLevel(star int8, camp_type int8) {

	in := &debug_api.PassAllDespairLandLevel_In{

		Star:     star,
		CampType: camp_type,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_DespairLandDummyBossKill() {

	in := &debug_api.DespairLandDummyBossKill_In{}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_AddTaoyuanItem(item_id int16, number int16) {

	in := &debug_api.AddTaoyuanItem_In{

		ItemId: item_id,
		Number: number,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}

func (c *Client) Debug_AddTaoyuanExp(add_exp int64) {

	in := &debug_api.AddTaoyuanExp_In{

		AddExp: add_exp,
	}

	buff := net.NewBuffer(make([]byte, in.ByteSize()))

	in.Encode(buff)

	c.conn.Write(buff.Get())

}
