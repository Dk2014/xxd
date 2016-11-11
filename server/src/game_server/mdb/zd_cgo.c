#include "zd_cgo.h"

GlobalTables g_GlobalTables;
PlayerTables* g_PlayerTables;

void Init() {
	bzero(&g_GlobalTables, sizeof(GlobalTables));
}

// 创建并返回新的玩家数据库切片
PlayerTables* NewPlayerTables() {
	return (PlayerTables*)calloc(1, sizeof(PlayerTables));
}

/* ========== global_announcement ========== */

GlobalAnnouncement* New_GlobalAnnouncement() {
	return (GlobalAnnouncement*)calloc(1, sizeof(GlobalAnnouncement));
}

void Free_GlobalAnnouncement(GlobalAnnouncement* row) {
	if (row->Parameters != NULL) {
		free(row->Parameters);
	}
	if (row->Content != NULL) {
		free(row->Content);
	}
	free(row);
}

/* ========== global_arena_rank ========== */

GlobalArenaRank* New_GlobalArenaRank() {
	return (GlobalArenaRank*)calloc(1, sizeof(GlobalArenaRank));
}

void Free_GlobalArenaRank(GlobalArenaRank* row) {
	free(row);
}

/* ========== global_clique ========== */

GlobalClique* New_GlobalClique() {
	return (GlobalClique*)calloc(1, sizeof(GlobalClique));
}

void Free_GlobalClique(GlobalClique* row) {
	if (row->Name != NULL) {
		free(row->Name);
	}
	if (row->Announce != NULL) {
		free(row->Announce);
	}
	if (row->MemberJson != NULL) {
		free(row->MemberJson);
	}
	free(row);
}

/* ========== global_clique_boat ========== */

GlobalCliqueBoat* New_GlobalCliqueBoat() {
	return (GlobalCliqueBoat*)calloc(1, sizeof(GlobalCliqueBoat));
}

void Free_GlobalCliqueBoat(GlobalCliqueBoat* row) {
	free(row);
}

/* ========== global_despair_boss ========== */

GlobalDespairBoss* New_GlobalDespairBoss() {
	return (GlobalDespairBoss*)calloc(1, sizeof(GlobalDespairBoss));
}

void Free_GlobalDespairBoss(GlobalDespairBoss* row) {
	free(row);
}

/* ========== global_despair_boss_history ========== */

GlobalDespairBossHistory* New_GlobalDespairBossHistory() {
	return (GlobalDespairBossHistory*)calloc(1, sizeof(GlobalDespairBossHistory));
}

void Free_GlobalDespairBossHistory(GlobalDespairBossHistory* row) {
	free(row);
}

/* ========== global_despair_land ========== */

GlobalDespairLand* New_GlobalDespairLand() {
	return (GlobalDespairLand*)calloc(1, sizeof(GlobalDespairLand));
}

void Free_GlobalDespairLand(GlobalDespairLand* row) {
	free(row);
}

/* ========== global_despair_land_battle_record ========== */

GlobalDespairLandBattleRecord* New_GlobalDespairLandBattleRecord() {
	return (GlobalDespairLandBattleRecord*)calloc(1, sizeof(GlobalDespairLandBattleRecord));
}

void Free_GlobalDespairLandBattleRecord(GlobalDespairLandBattleRecord* row) {
	if (row->Record != NULL) {
		free(row->Record);
	}
	free(row);
}

/* ========== global_despair_land_camp ========== */

GlobalDespairLandCamp* New_GlobalDespairLandCamp() {
	return (GlobalDespairLandCamp*)calloc(1, sizeof(GlobalDespairLandCamp));
}

void Free_GlobalDespairLandCamp(GlobalDespairLandCamp* row) {
	free(row);
}

/* ========== global_gift_card_record ========== */

GlobalGiftCardRecord* New_GlobalGiftCardRecord() {
	return (GlobalGiftCardRecord*)calloc(1, sizeof(GlobalGiftCardRecord));
}

void Free_GlobalGiftCardRecord(GlobalGiftCardRecord* row) {
	if (row->Code != NULL) {
		free(row->Code);
	}
	free(row);
}

/* ========== global_group_buy_status ========== */

GlobalGroupBuyStatus* New_GlobalGroupBuyStatus() {
	return (GlobalGroupBuyStatus*)calloc(1, sizeof(GlobalGroupBuyStatus));
}

void Free_GlobalGroupBuyStatus(GlobalGroupBuyStatus* row) {
	free(row);
}

/* ========== global_mail ========== */

GlobalMail* New_GlobalMail() {
	return (GlobalMail*)calloc(1, sizeof(GlobalMail));
}

void Free_GlobalMail(GlobalMail* row) {
	if (row->Title != NULL) {
		free(row->Title);
	}
	if (row->Content != NULL) {
		free(row->Content);
	}
	free(row);
}

/* ========== global_mail_attachments ========== */

GlobalMailAttachments* New_GlobalMailAttachments() {
	return (GlobalMailAttachments*)calloc(1, sizeof(GlobalMailAttachments));
}

void Free_GlobalMailAttachments(GlobalMailAttachments* row) {
	free(row);
}

/* ========== global_tb_xxd_onlinecnt ========== */

GlobalTbXxdOnlinecnt* New_GlobalTbXxdOnlinecnt() {
	return (GlobalTbXxdOnlinecnt*)calloc(1, sizeof(GlobalTbXxdOnlinecnt));
}

void Free_GlobalTbXxdOnlinecnt(GlobalTbXxdOnlinecnt* row) {
	if (row->Gameappid != NULL) {
		free(row->Gameappid);
	}
	free(row);
}

/* ========== player ========== */

Player* New_Player() {
	return (Player*)calloc(1, sizeof(Player));
}

void Free_Player(Player* row) {
	if (row->User != NULL) {
		free(row->User);
	}
	if (row->Nick != NULL) {
		free(row->Nick);
	}
	free(row);
}

/* ========== player_activity ========== */

PlayerActivity* New_PlayerActivity() {
	return (PlayerActivity*)calloc(1, sizeof(PlayerActivity));
}

void Free_PlayerActivity(PlayerActivity* row) {
	free(row);
}

/* ========== player_addition_quest ========== */

PlayerAdditionQuest* New_PlayerAdditionQuest() {
	return (PlayerAdditionQuest*)calloc(1, sizeof(PlayerAdditionQuest));
}

void Free_PlayerAdditionQuest(PlayerAdditionQuest* row) {
	free(row);
}

/* ========== player_any_sdk_order ========== */

PlayerAnySdkOrder* New_PlayerAnySdkOrder() {
	return (PlayerAnySdkOrder*)calloc(1, sizeof(PlayerAnySdkOrder));
}

void Free_PlayerAnySdkOrder(PlayerAnySdkOrder* row) {
	free(row);
}

/* ========== player_arena ========== */

PlayerArena* New_PlayerArena() {
	return (PlayerArena*)calloc(1, sizeof(PlayerArena));
}

void Free_PlayerArena(PlayerArena* row) {
	free(row);
}

/* ========== player_arena_rank ========== */

PlayerArenaRank* New_PlayerArenaRank() {
	return (PlayerArenaRank*)calloc(1, sizeof(PlayerArenaRank));
}

void Free_PlayerArenaRank(PlayerArenaRank* row) {
	free(row);
}

/* ========== player_arena_record ========== */

PlayerArenaRecord* New_PlayerArenaRecord() {
	return (PlayerArenaRecord*)calloc(1, sizeof(PlayerArenaRecord));
}

void Free_PlayerArenaRecord(PlayerArenaRecord* row) {
	if (row->TargetNick != NULL) {
		free(row->TargetNick);
	}
	free(row);
}

/* ========== player_awaken_graphic ========== */

PlayerAwakenGraphic* New_PlayerAwakenGraphic() {
	return (PlayerAwakenGraphic*)calloc(1, sizeof(PlayerAwakenGraphic));
}

void Free_PlayerAwakenGraphic(PlayerAwakenGraphic* row) {
	free(row);
}

/* ========== player_battle_pet ========== */

PlayerBattlePet* New_PlayerBattlePet() {
	return (PlayerBattlePet*)calloc(1, sizeof(PlayerBattlePet));
}

void Free_PlayerBattlePet(PlayerBattlePet* row) {
	free(row);
}

/* ========== player_battle_pet_grid ========== */

PlayerBattlePetGrid* New_PlayerBattlePetGrid() {
	return (PlayerBattlePetGrid*)calloc(1, sizeof(PlayerBattlePetGrid));
}

void Free_PlayerBattlePetGrid(PlayerBattlePetGrid* row) {
	free(row);
}

/* ========== player_battle_pet_state ========== */

PlayerBattlePetState* New_PlayerBattlePetState() {
	return (PlayerBattlePetState*)calloc(1, sizeof(PlayerBattlePetState));
}

void Free_PlayerBattlePetState(PlayerBattlePetState* row) {
	free(row);
}

/* ========== player_chest_state ========== */

PlayerChestState* New_PlayerChestState() {
	return (PlayerChestState*)calloc(1, sizeof(PlayerChestState));
}

void Free_PlayerChestState(PlayerChestState* row) {
	free(row);
}

/* ========== player_clique_kongfu_attrib ========== */

PlayerCliqueKongfuAttrib* New_PlayerCliqueKongfuAttrib() {
	return (PlayerCliqueKongfuAttrib*)calloc(1, sizeof(PlayerCliqueKongfuAttrib));
}

void Free_PlayerCliqueKongfuAttrib(PlayerCliqueKongfuAttrib* row) {
	free(row);
}

/* ========== player_coins ========== */

PlayerCoins* New_PlayerCoins() {
	return (PlayerCoins*)calloc(1, sizeof(PlayerCoins));
}

void Free_PlayerCoins(PlayerCoins* row) {
	free(row);
}

/* ========== player_cornucopia ========== */

PlayerCornucopia* New_PlayerCornucopia() {
	return (PlayerCornucopia*)calloc(1, sizeof(PlayerCornucopia));
}

void Free_PlayerCornucopia(PlayerCornucopia* row) {
	free(row);
}

/* ========== player_daily_quest ========== */

PlayerDailyQuest* New_PlayerDailyQuest() {
	return (PlayerDailyQuest*)calloc(1, sizeof(PlayerDailyQuest));
}

void Free_PlayerDailyQuest(PlayerDailyQuest* row) {
	free(row);
}

/* ========== player_daily_quest_star_award_info ========== */

PlayerDailyQuestStarAwardInfo* New_PlayerDailyQuestStarAwardInfo() {
	return (PlayerDailyQuestStarAwardInfo*)calloc(1, sizeof(PlayerDailyQuestStarAwardInfo));
}

void Free_PlayerDailyQuestStarAwardInfo(PlayerDailyQuestStarAwardInfo* row) {
	if (row->Awarded != NULL) {
		free(row->Awarded);
	}
	free(row);
}

/* ========== player_daily_sign_in_record ========== */

PlayerDailySignInRecord* New_PlayerDailySignInRecord() {
	return (PlayerDailySignInRecord*)calloc(1, sizeof(PlayerDailySignInRecord));
}

void Free_PlayerDailySignInRecord(PlayerDailySignInRecord* row) {
	free(row);
}

/* ========== player_daily_sign_in_state ========== */

PlayerDailySignInState* New_PlayerDailySignInState() {
	return (PlayerDailySignInState*)calloc(1, sizeof(PlayerDailySignInState));
}

void Free_PlayerDailySignInState(PlayerDailySignInState* row) {
	free(row);
}

/* ========== player_despair_land_camp_state ========== */

PlayerDespairLandCampState* New_PlayerDespairLandCampState() {
	return (PlayerDespairLandCampState*)calloc(1, sizeof(PlayerDespairLandCampState));
}

void Free_PlayerDespairLandCampState(PlayerDespairLandCampState* row) {
	free(row);
}

/* ========== player_despair_land_camp_state_history ========== */

PlayerDespairLandCampStateHistory* New_PlayerDespairLandCampStateHistory() {
	return (PlayerDespairLandCampStateHistory*)calloc(1, sizeof(PlayerDespairLandCampStateHistory));
}

void Free_PlayerDespairLandCampStateHistory(PlayerDespairLandCampStateHistory* row) {
	free(row);
}

/* ========== player_despair_land_level_record ========== */

PlayerDespairLandLevelRecord* New_PlayerDespairLandLevelRecord() {
	return (PlayerDespairLandLevelRecord*)calloc(1, sizeof(PlayerDespairLandLevelRecord));
}

void Free_PlayerDespairLandLevelRecord(PlayerDespairLandLevelRecord* row) {
	free(row);
}

/* ========== player_despair_land_state ========== */

PlayerDespairLandState* New_PlayerDespairLandState() {
	return (PlayerDespairLandState*)calloc(1, sizeof(PlayerDespairLandState));
}

void Free_PlayerDespairLandState(PlayerDespairLandState* row) {
	free(row);
}

/* ========== player_driving_sword_event ========== */

PlayerDrivingSwordEvent* New_PlayerDrivingSwordEvent() {
	return (PlayerDrivingSwordEvent*)calloc(1, sizeof(PlayerDrivingSwordEvent));
}

void Free_PlayerDrivingSwordEvent(PlayerDrivingSwordEvent* row) {
	free(row);
}

/* ========== player_driving_sword_event_exploring ========== */

PlayerDrivingSwordEventExploring* New_PlayerDrivingSwordEventExploring() {
	return (PlayerDrivingSwordEventExploring*)calloc(1, sizeof(PlayerDrivingSwordEventExploring));
}

void Free_PlayerDrivingSwordEventExploring(PlayerDrivingSwordEventExploring* row) {
	free(row);
}

/* ========== player_driving_sword_event_treasure ========== */

PlayerDrivingSwordEventTreasure* New_PlayerDrivingSwordEventTreasure() {
	return (PlayerDrivingSwordEventTreasure*)calloc(1, sizeof(PlayerDrivingSwordEventTreasure));
}

void Free_PlayerDrivingSwordEventTreasure(PlayerDrivingSwordEventTreasure* row) {
	free(row);
}

/* ========== player_driving_sword_event_visiting ========== */

PlayerDrivingSwordEventVisiting* New_PlayerDrivingSwordEventVisiting() {
	return (PlayerDrivingSwordEventVisiting*)calloc(1, sizeof(PlayerDrivingSwordEventVisiting));
}

void Free_PlayerDrivingSwordEventVisiting(PlayerDrivingSwordEventVisiting* row) {
	if (row->TargetSideState != NULL) {
		free(row->TargetSideState);
	}
	if (row->TargetStatus != NULL) {
		free(row->TargetStatus);
	}
	free(row);
}

/* ========== player_driving_sword_eventmask ========== */

PlayerDrivingSwordEventmask* New_PlayerDrivingSwordEventmask() {
	return (PlayerDrivingSwordEventmask*)calloc(1, sizeof(PlayerDrivingSwordEventmask));
}

void Free_PlayerDrivingSwordEventmask(PlayerDrivingSwordEventmask* row) {
	if (row->Events != NULL) {
		free(row->Events);
	}
	free(row);
}

/* ========== player_driving_sword_info ========== */

PlayerDrivingSwordInfo* New_PlayerDrivingSwordInfo() {
	return (PlayerDrivingSwordInfo*)calloc(1, sizeof(PlayerDrivingSwordInfo));
}

void Free_PlayerDrivingSwordInfo(PlayerDrivingSwordInfo* row) {
	free(row);
}

/* ========== player_driving_sword_map ========== */

PlayerDrivingSwordMap* New_PlayerDrivingSwordMap() {
	return (PlayerDrivingSwordMap*)calloc(1, sizeof(PlayerDrivingSwordMap));
}

void Free_PlayerDrivingSwordMap(PlayerDrivingSwordMap* row) {
	if (row->Shadows != NULL) {
		free(row->Shadows);
	}
	free(row);
}

/* ========== player_equipment ========== */

PlayerEquipment* New_PlayerEquipment() {
	return (PlayerEquipment*)calloc(1, sizeof(PlayerEquipment));
}

void Free_PlayerEquipment(PlayerEquipment* row) {
	free(row);
}

/* ========== player_event_award_record ========== */

PlayerEventAwardRecord* New_PlayerEventAwardRecord() {
	return (PlayerEventAwardRecord*)calloc(1, sizeof(PlayerEventAwardRecord));
}

void Free_PlayerEventAwardRecord(PlayerEventAwardRecord* row) {
	if (row->RecordBytes != NULL) {
		free(row->RecordBytes);
	}
	if (row->JsonEventRecord != NULL) {
		free(row->JsonEventRecord);
	}
	free(row);
}

/* ========== player_event_daily_online ========== */

PlayerEventDailyOnline* New_PlayerEventDailyOnline() {
	return (PlayerEventDailyOnline*)calloc(1, sizeof(PlayerEventDailyOnline));
}

void Free_PlayerEventDailyOnline(PlayerEventDailyOnline* row) {
	free(row);
}

/* ========== player_event_vn_daily_recharge ========== */

PlayerEventVnDailyRecharge* New_PlayerEventVnDailyRecharge() {
	return (PlayerEventVnDailyRecharge*)calloc(1, sizeof(PlayerEventVnDailyRecharge));
}

void Free_PlayerEventVnDailyRecharge(PlayerEventVnDailyRecharge* row) {
	free(row);
}

/* ========== player_event_vn_dragon_ball_record ========== */

PlayerEventVnDragonBallRecord* New_PlayerEventVnDragonBallRecord() {
	return (PlayerEventVnDragonBallRecord*)calloc(1, sizeof(PlayerEventVnDragonBallRecord));
}

void Free_PlayerEventVnDragonBallRecord(PlayerEventVnDragonBallRecord* row) {
	free(row);
}

/* ========== player_events_fight_power ========== */

PlayerEventsFightPower* New_PlayerEventsFightPower() {
	return (PlayerEventsFightPower*)calloc(1, sizeof(PlayerEventsFightPower));
}

void Free_PlayerEventsFightPower(PlayerEventsFightPower* row) {
	free(row);
}

/* ========== player_events_ingot_record ========== */

PlayerEventsIngotRecord* New_PlayerEventsIngotRecord() {
	return (PlayerEventsIngotRecord*)calloc(1, sizeof(PlayerEventsIngotRecord));
}

void Free_PlayerEventsIngotRecord(PlayerEventsIngotRecord* row) {
	free(row);
}

/* ========== player_extend_level ========== */

PlayerExtendLevel* New_PlayerExtendLevel() {
	return (PlayerExtendLevel*)calloc(1, sizeof(PlayerExtendLevel));
}

void Free_PlayerExtendLevel(PlayerExtendLevel* row) {
	free(row);
}

/* ========== player_extend_quest ========== */

PlayerExtendQuest* New_PlayerExtendQuest() {
	return (PlayerExtendQuest*)calloc(1, sizeof(PlayerExtendQuest));
}

void Free_PlayerExtendQuest(PlayerExtendQuest* row) {
	free(row);
}

/* ========== player_fame ========== */

PlayerFame* New_PlayerFame() {
	return (PlayerFame*)calloc(1, sizeof(PlayerFame));
}

void Free_PlayerFame(PlayerFame* row) {
	free(row);
}

/* ========== player_fashion ========== */

PlayerFashion* New_PlayerFashion() {
	return (PlayerFashion*)calloc(1, sizeof(PlayerFashion));
}

void Free_PlayerFashion(PlayerFashion* row) {
	free(row);
}

/* ========== player_fashion_state ========== */

PlayerFashionState* New_PlayerFashionState() {
	return (PlayerFashionState*)calloc(1, sizeof(PlayerFashionState));
}

void Free_PlayerFashionState(PlayerFashionState* row) {
	free(row);
}

/* ========== player_fate_box_state ========== */

PlayerFateBoxState* New_PlayerFateBoxState() {
	return (PlayerFateBoxState*)calloc(1, sizeof(PlayerFateBoxState));
}

void Free_PlayerFateBoxState(PlayerFateBoxState* row) {
	free(row);
}

/* ========== player_fight_num ========== */

PlayerFightNum* New_PlayerFightNum() {
	return (PlayerFightNum*)calloc(1, sizeof(PlayerFightNum));
}

void Free_PlayerFightNum(PlayerFightNum* row) {
	free(row);
}

/* ========== player_first_charge ========== */

PlayerFirstCharge* New_PlayerFirstCharge() {
	return (PlayerFirstCharge*)calloc(1, sizeof(PlayerFirstCharge));
}

void Free_PlayerFirstCharge(PlayerFirstCharge* row) {
	if (row->ProductId != NULL) {
		free(row->ProductId);
	}
	free(row);
}

/* ========== player_formation ========== */

PlayerFormation* New_PlayerFormation() {
	return (PlayerFormation*)calloc(1, sizeof(PlayerFormation));
}

void Free_PlayerFormation(PlayerFormation* row) {
	free(row);
}

/* ========== player_func_key ========== */

PlayerFuncKey* New_PlayerFuncKey() {
	return (PlayerFuncKey*)calloc(1, sizeof(PlayerFuncKey));
}

void Free_PlayerFuncKey(PlayerFuncKey* row) {
	free(row);
}

/* ========== player_ghost ========== */

PlayerGhost* New_PlayerGhost() {
	return (PlayerGhost*)calloc(1, sizeof(PlayerGhost));
}

void Free_PlayerGhost(PlayerGhost* row) {
	free(row);
}

/* ========== player_ghost_equipment ========== */

PlayerGhostEquipment* New_PlayerGhostEquipment() {
	return (PlayerGhostEquipment*)calloc(1, sizeof(PlayerGhostEquipment));
}

void Free_PlayerGhostEquipment(PlayerGhostEquipment* row) {
	free(row);
}

/* ========== player_ghost_state ========== */

PlayerGhostState* New_PlayerGhostState() {
	return (PlayerGhostState*)calloc(1, sizeof(PlayerGhostState));
}

void Free_PlayerGhostState(PlayerGhostState* row) {
	free(row);
}

/* ========== player_global_chat_state ========== */

PlayerGlobalChatState* New_PlayerGlobalChatState() {
	return (PlayerGlobalChatState*)calloc(1, sizeof(PlayerGlobalChatState));
}

void Free_PlayerGlobalChatState(PlayerGlobalChatState* row) {
	free(row);
}

/* ========== player_global_clique_building ========== */

PlayerGlobalCliqueBuilding* New_PlayerGlobalCliqueBuilding() {
	return (PlayerGlobalCliqueBuilding*)calloc(1, sizeof(PlayerGlobalCliqueBuilding));
}

void Free_PlayerGlobalCliqueBuilding(PlayerGlobalCliqueBuilding* row) {
	free(row);
}

/* ========== player_global_clique_building_quest ========== */

PlayerGlobalCliqueBuildingQuest* New_PlayerGlobalCliqueBuildingQuest() {
	return (PlayerGlobalCliqueBuildingQuest*)calloc(1, sizeof(PlayerGlobalCliqueBuildingQuest));
}

void Free_PlayerGlobalCliqueBuildingQuest(PlayerGlobalCliqueBuildingQuest* row) {
	free(row);
}

/* ========== player_global_clique_daily_quest ========== */

PlayerGlobalCliqueDailyQuest* New_PlayerGlobalCliqueDailyQuest() {
	return (PlayerGlobalCliqueDailyQuest*)calloc(1, sizeof(PlayerGlobalCliqueDailyQuest));
}

void Free_PlayerGlobalCliqueDailyQuest(PlayerGlobalCliqueDailyQuest* row) {
	free(row);
}

/* ========== player_global_clique_escort ========== */

PlayerGlobalCliqueEscort* New_PlayerGlobalCliqueEscort() {
	return (PlayerGlobalCliqueEscort*)calloc(1, sizeof(PlayerGlobalCliqueEscort));
}

void Free_PlayerGlobalCliqueEscort(PlayerGlobalCliqueEscort* row) {
	free(row);
}

/* ========== player_global_clique_escort_message ========== */

PlayerGlobalCliqueEscortMessage* New_PlayerGlobalCliqueEscortMessage() {
	return (PlayerGlobalCliqueEscortMessage*)calloc(1, sizeof(PlayerGlobalCliqueEscortMessage));
}

void Free_PlayerGlobalCliqueEscortMessage(PlayerGlobalCliqueEscortMessage* row) {
	if (row->Parameters != NULL) {
		free(row->Parameters);
	}
	free(row);
}

/* ========== player_global_clique_hijack ========== */

PlayerGlobalCliqueHijack* New_PlayerGlobalCliqueHijack() {
	return (PlayerGlobalCliqueHijack*)calloc(1, sizeof(PlayerGlobalCliqueHijack));
}

void Free_PlayerGlobalCliqueHijack(PlayerGlobalCliqueHijack* row) {
	free(row);
}

/* ========== player_global_clique_info ========== */

PlayerGlobalCliqueInfo* New_PlayerGlobalCliqueInfo() {
	return (PlayerGlobalCliqueInfo*)calloc(1, sizeof(PlayerGlobalCliqueInfo));
}

void Free_PlayerGlobalCliqueInfo(PlayerGlobalCliqueInfo* row) {
	free(row);
}

/* ========== player_global_clique_kongfu ========== */

PlayerGlobalCliqueKongfu* New_PlayerGlobalCliqueKongfu() {
	return (PlayerGlobalCliqueKongfu*)calloc(1, sizeof(PlayerGlobalCliqueKongfu));
}

void Free_PlayerGlobalCliqueKongfu(PlayerGlobalCliqueKongfu* row) {
	free(row);
}

/* ========== player_global_friend ========== */

PlayerGlobalFriend* New_PlayerGlobalFriend() {
	return (PlayerGlobalFriend*)calloc(1, sizeof(PlayerGlobalFriend));
}

void Free_PlayerGlobalFriend(PlayerGlobalFriend* row) {
	if (row->FriendNick != NULL) {
		free(row->FriendNick);
	}
	free(row);
}

/* ========== player_global_friend_chat ========== */

PlayerGlobalFriendChat* New_PlayerGlobalFriendChat() {
	return (PlayerGlobalFriendChat*)calloc(1, sizeof(PlayerGlobalFriendChat));
}

void Free_PlayerGlobalFriendChat(PlayerGlobalFriendChat* row) {
	if (row->Message != NULL) {
		free(row->Message);
	}
	free(row);
}

/* ========== player_global_friend_state ========== */

PlayerGlobalFriendState* New_PlayerGlobalFriendState() {
	return (PlayerGlobalFriendState*)calloc(1, sizeof(PlayerGlobalFriendState));
}

void Free_PlayerGlobalFriendState(PlayerGlobalFriendState* row) {
	free(row);
}

/* ========== player_global_gift_card_record ========== */

PlayerGlobalGiftCardRecord* New_PlayerGlobalGiftCardRecord() {
	return (PlayerGlobalGiftCardRecord*)calloc(1, sizeof(PlayerGlobalGiftCardRecord));
}

void Free_PlayerGlobalGiftCardRecord(PlayerGlobalGiftCardRecord* row) {
	if (row->Code != NULL) {
		free(row->Code);
	}
	free(row);
}

/* ========== player_global_mail_state ========== */

PlayerGlobalMailState* New_PlayerGlobalMailState() {
	return (PlayerGlobalMailState*)calloc(1, sizeof(PlayerGlobalMailState));
}

void Free_PlayerGlobalMailState(PlayerGlobalMailState* row) {
	free(row);
}

/* ========== player_global_world_chat_state ========== */

PlayerGlobalWorldChatState* New_PlayerGlobalWorldChatState() {
	return (PlayerGlobalWorldChatState*)calloc(1, sizeof(PlayerGlobalWorldChatState));
}

void Free_PlayerGlobalWorldChatState(PlayerGlobalWorldChatState* row) {
	free(row);
}

/* ========== player_hard_level ========== */

PlayerHardLevel* New_PlayerHardLevel() {
	return (PlayerHardLevel*)calloc(1, sizeof(PlayerHardLevel));
}

void Free_PlayerHardLevel(PlayerHardLevel* row) {
	free(row);
}

/* ========== player_hard_level_record ========== */

PlayerHardLevelRecord* New_PlayerHardLevelRecord() {
	return (PlayerHardLevelRecord*)calloc(1, sizeof(PlayerHardLevelRecord));
}

void Free_PlayerHardLevelRecord(PlayerHardLevelRecord* row) {
	free(row);
}

/* ========== player_heart ========== */

PlayerHeart* New_PlayerHeart() {
	return (PlayerHeart*)calloc(1, sizeof(PlayerHeart));
}

void Free_PlayerHeart(PlayerHeart* row) {
	free(row);
}

/* ========== player_heart_draw ========== */

PlayerHeartDraw* New_PlayerHeartDraw() {
	return (PlayerHeartDraw*)calloc(1, sizeof(PlayerHeartDraw));
}

void Free_PlayerHeartDraw(PlayerHeartDraw* row) {
	free(row);
}

/* ========== player_heart_draw_card_record ========== */

PlayerHeartDrawCardRecord* New_PlayerHeartDrawCardRecord() {
	return (PlayerHeartDrawCardRecord*)calloc(1, sizeof(PlayerHeartDrawCardRecord));
}

void Free_PlayerHeartDrawCardRecord(PlayerHeartDrawCardRecord* row) {
	free(row);
}

/* ========== player_heart_draw_wheel_record ========== */

PlayerHeartDrawWheelRecord* New_PlayerHeartDrawWheelRecord() {
	return (PlayerHeartDrawWheelRecord*)calloc(1, sizeof(PlayerHeartDrawWheelRecord));
}

void Free_PlayerHeartDrawWheelRecord(PlayerHeartDrawWheelRecord* row) {
	free(row);
}

/* ========== player_info ========== */

PlayerInfo* New_PlayerInfo() {
	return (PlayerInfo*)calloc(1, sizeof(PlayerInfo));
}

void Free_PlayerInfo(PlayerInfo* row) {
	free(row);
}

/* ========== player_is_operated ========== */

PlayerIsOperated* New_PlayerIsOperated() {
	return (PlayerIsOperated*)calloc(1, sizeof(PlayerIsOperated));
}

void Free_PlayerIsOperated(PlayerIsOperated* row) {
	free(row);
}

/* ========== player_item ========== */

PlayerItem* New_PlayerItem() {
	return (PlayerItem*)calloc(1, sizeof(PlayerItem));
}

void Free_PlayerItem(PlayerItem* row) {
	free(row);
}

/* ========== player_item_appendix ========== */

PlayerItemAppendix* New_PlayerItemAppendix() {
	return (PlayerItemAppendix*)calloc(1, sizeof(PlayerItemAppendix));
}

void Free_PlayerItemAppendix(PlayerItemAppendix* row) {
	free(row);
}

/* ========== player_item_buyback ========== */

PlayerItemBuyback* New_PlayerItemBuyback() {
	return (PlayerItemBuyback*)calloc(1, sizeof(PlayerItemBuyback));
}

void Free_PlayerItemBuyback(PlayerItemBuyback* row) {
	free(row);
}

/* ========== player_item_recast_state ========== */

PlayerItemRecastState* New_PlayerItemRecastState() {
	return (PlayerItemRecastState*)calloc(1, sizeof(PlayerItemRecastState));
}

void Free_PlayerItemRecastState(PlayerItemRecastState* row) {
	free(row);
}

/* ========== player_level_award_info ========== */

PlayerLevelAwardInfo* New_PlayerLevelAwardInfo() {
	return (PlayerLevelAwardInfo*)calloc(1, sizeof(PlayerLevelAwardInfo));
}

void Free_PlayerLevelAwardInfo(PlayerLevelAwardInfo* row) {
	if (row->Awarded != NULL) {
		free(row->Awarded);
	}
	free(row);
}

/* ========== player_login_award_record ========== */

PlayerLoginAwardRecord* New_PlayerLoginAwardRecord() {
	return (PlayerLoginAwardRecord*)calloc(1, sizeof(PlayerLoginAwardRecord));
}

void Free_PlayerLoginAwardRecord(PlayerLoginAwardRecord* row) {
	free(row);
}

/* ========== player_mail ========== */

PlayerMail* New_PlayerMail() {
	return (PlayerMail*)calloc(1, sizeof(PlayerMail));
}

void Free_PlayerMail(PlayerMail* row) {
	if (row->Parameters != NULL) {
		free(row->Parameters);
	}
	if (row->Title != NULL) {
		free(row->Title);
	}
	if (row->Content != NULL) {
		free(row->Content);
	}
	free(row);
}

/* ========== player_mail_attachment ========== */

PlayerMailAttachment* New_PlayerMailAttachment() {
	return (PlayerMailAttachment*)calloc(1, sizeof(PlayerMailAttachment));
}

void Free_PlayerMailAttachment(PlayerMailAttachment* row) {
	free(row);
}

/* ========== player_mail_attachment_lg ========== */

PlayerMailAttachmentLg* New_PlayerMailAttachmentLg() {
	return (PlayerMailAttachmentLg*)calloc(1, sizeof(PlayerMailAttachmentLg));
}

void Free_PlayerMailAttachmentLg(PlayerMailAttachmentLg* row) {
	free(row);
}

/* ========== player_mail_lg ========== */

PlayerMailLg* New_PlayerMailLg() {
	return (PlayerMailLg*)calloc(1, sizeof(PlayerMailLg));
}

void Free_PlayerMailLg(PlayerMailLg* row) {
	if (row->Parameters != NULL) {
		free(row->Parameters);
	}
	if (row->Title != NULL) {
		free(row->Title);
	}
	if (row->Content != NULL) {
		free(row->Content);
	}
	free(row);
}

/* ========== player_meditation_state ========== */

PlayerMeditationState* New_PlayerMeditationState() {
	return (PlayerMeditationState*)calloc(1, sizeof(PlayerMeditationState));
}

void Free_PlayerMeditationState(PlayerMeditationState* row) {
	free(row);
}

/* ========== player_mission ========== */

PlayerMission* New_PlayerMission() {
	return (PlayerMission*)calloc(1, sizeof(PlayerMission));
}

void Free_PlayerMission(PlayerMission* row) {
	free(row);
}

/* ========== player_mission_level ========== */

PlayerMissionLevel* New_PlayerMissionLevel() {
	return (PlayerMissionLevel*)calloc(1, sizeof(PlayerMissionLevel));
}

void Free_PlayerMissionLevel(PlayerMissionLevel* row) {
	free(row);
}

/* ========== player_mission_level_record ========== */

PlayerMissionLevelRecord* New_PlayerMissionLevelRecord() {
	return (PlayerMissionLevelRecord*)calloc(1, sizeof(PlayerMissionLevelRecord));
}

void Free_PlayerMissionLevelRecord(PlayerMissionLevelRecord* row) {
	free(row);
}

/* ========== player_mission_level_state_bin ========== */

PlayerMissionLevelStateBin* New_PlayerMissionLevelStateBin() {
	return (PlayerMissionLevelStateBin*)calloc(1, sizeof(PlayerMissionLevelStateBin));
}

void Free_PlayerMissionLevelStateBin(PlayerMissionLevelStateBin* row) {
	if (row->Bin != NULL) {
		free(row->Bin);
	}
	free(row);
}

/* ========== player_mission_record ========== */

PlayerMissionRecord* New_PlayerMissionRecord() {
	return (PlayerMissionRecord*)calloc(1, sizeof(PlayerMissionRecord));
}

void Free_PlayerMissionRecord(PlayerMissionRecord* row) {
	free(row);
}

/* ========== player_mission_star_award ========== */

PlayerMissionStarAward* New_PlayerMissionStarAward() {
	return (PlayerMissionStarAward*)calloc(1, sizeof(PlayerMissionStarAward));
}

void Free_PlayerMissionStarAward(PlayerMissionStarAward* row) {
	free(row);
}

/* ========== player_money_tree ========== */

PlayerMoneyTree* New_PlayerMoneyTree() {
	return (PlayerMoneyTree*)calloc(1, sizeof(PlayerMoneyTree));
}

void Free_PlayerMoneyTree(PlayerMoneyTree* row) {
	free(row);
}

/* ========== player_month_card_award_record ========== */

PlayerMonthCardAwardRecord* New_PlayerMonthCardAwardRecord() {
	return (PlayerMonthCardAwardRecord*)calloc(1, sizeof(PlayerMonthCardAwardRecord));
}

void Free_PlayerMonthCardAwardRecord(PlayerMonthCardAwardRecord* row) {
	free(row);
}

/* ========== player_month_card_info ========== */

PlayerMonthCardInfo* New_PlayerMonthCardInfo() {
	return (PlayerMonthCardInfo*)calloc(1, sizeof(PlayerMonthCardInfo));
}

void Free_PlayerMonthCardInfo(PlayerMonthCardInfo* row) {
	free(row);
}

/* ========== player_multi_level_info ========== */

PlayerMultiLevelInfo* New_PlayerMultiLevelInfo() {
	return (PlayerMultiLevelInfo*)calloc(1, sizeof(PlayerMultiLevelInfo));
}

void Free_PlayerMultiLevelInfo(PlayerMultiLevelInfo* row) {
	free(row);
}

/* ========== player_new_year_consume_record ========== */

PlayerNewYearConsumeRecord* New_PlayerNewYearConsumeRecord() {
	return (PlayerNewYearConsumeRecord*)calloc(1, sizeof(PlayerNewYearConsumeRecord));
}

void Free_PlayerNewYearConsumeRecord(PlayerNewYearConsumeRecord* row) {
	if (row->ConsumeRecord != NULL) {
		free(row->ConsumeRecord);
	}
	free(row);
}

/* ========== player_npc_talk_record ========== */

PlayerNpcTalkRecord* New_PlayerNpcTalkRecord() {
	return (PlayerNpcTalkRecord*)calloc(1, sizeof(PlayerNpcTalkRecord));
}

void Free_PlayerNpcTalkRecord(PlayerNpcTalkRecord* row) {
	free(row);
}

/* ========== player_opened_town_treasure ========== */

PlayerOpenedTownTreasure* New_PlayerOpenedTownTreasure() {
	return (PlayerOpenedTownTreasure*)calloc(1, sizeof(PlayerOpenedTownTreasure));
}

void Free_PlayerOpenedTownTreasure(PlayerOpenedTownTreasure* row) {
	free(row);
}

/* ========== player_physical ========== */

PlayerPhysical* New_PlayerPhysical() {
	return (PlayerPhysical*)calloc(1, sizeof(PlayerPhysical));
}

void Free_PlayerPhysical(PlayerPhysical* row) {
	free(row);
}

/* ========== player_purchase_record ========== */

PlayerPurchaseRecord* New_PlayerPurchaseRecord() {
	return (PlayerPurchaseRecord*)calloc(1, sizeof(PlayerPurchaseRecord));
}

void Free_PlayerPurchaseRecord(PlayerPurchaseRecord* row) {
	free(row);
}

/* ========== player_push_notify_switch ========== */

PlayerPushNotifySwitch* New_PlayerPushNotifySwitch() {
	return (PlayerPushNotifySwitch*)calloc(1, sizeof(PlayerPushNotifySwitch));
}

void Free_PlayerPushNotifySwitch(PlayerPushNotifySwitch* row) {
	free(row);
}

/* ========== player_pve_state ========== */

PlayerPveState* New_PlayerPveState() {
	return (PlayerPveState*)calloc(1, sizeof(PlayerPveState));
}

void Free_PlayerPveState(PlayerPveState* row) {
	free(row);
}

/* ========== player_qq_vip_gift ========== */

PlayerQqVipGift* New_PlayerQqVipGift() {
	return (PlayerQqVipGift*)calloc(1, sizeof(PlayerQqVipGift));
}

void Free_PlayerQqVipGift(PlayerQqVipGift* row) {
	free(row);
}

/* ========== player_quest ========== */

PlayerQuest* New_PlayerQuest() {
	return (PlayerQuest*)calloc(1, sizeof(PlayerQuest));
}

void Free_PlayerQuest(PlayerQuest* row) {
	free(row);
}

/* ========== player_rainbow_level ========== */

PlayerRainbowLevel* New_PlayerRainbowLevel() {
	return (PlayerRainbowLevel*)calloc(1, sizeof(PlayerRainbowLevel));
}

void Free_PlayerRainbowLevel(PlayerRainbowLevel* row) {
	free(row);
}

/* ========== player_rainbow_level_state_bin ========== */

PlayerRainbowLevelStateBin* New_PlayerRainbowLevelStateBin() {
	return (PlayerRainbowLevelStateBin*)calloc(1, sizeof(PlayerRainbowLevelStateBin));
}

void Free_PlayerRainbowLevelStateBin(PlayerRainbowLevelStateBin* row) {
	if (row->Bin != NULL) {
		free(row->Bin);
	}
	free(row);
}

/* ========== player_role ========== */

PlayerRole* New_PlayerRole() {
	return (PlayerRole*)calloc(1, sizeof(PlayerRole));
}

void Free_PlayerRole(PlayerRole* row) {
	free(row);
}

/* ========== player_sealedbook ========== */

PlayerSealedbook* New_PlayerSealedbook() {
	return (PlayerSealedbook*)calloc(1, sizeof(PlayerSealedbook));
}

void Free_PlayerSealedbook(PlayerSealedbook* row) {
	if (row->SealedRecord != NULL) {
		free(row->SealedRecord);
	}
	free(row);
}

/* ========== player_send_heart_record ========== */

PlayerSendHeartRecord* New_PlayerSendHeartRecord() {
	return (PlayerSendHeartRecord*)calloc(1, sizeof(PlayerSendHeartRecord));
}

void Free_PlayerSendHeartRecord(PlayerSendHeartRecord* row) {
	free(row);
}

/* ========== player_skill ========== */

PlayerSkill* New_PlayerSkill() {
	return (PlayerSkill*)calloc(1, sizeof(PlayerSkill));
}

void Free_PlayerSkill(PlayerSkill* row) {
	free(row);
}

/* ========== player_state ========== */

PlayerState* New_PlayerState() {
	return (PlayerState*)calloc(1, sizeof(PlayerState));
}

void Free_PlayerState(PlayerState* row) {
	free(row);
}

/* ========== player_sword_soul ========== */

PlayerSwordSoul* New_PlayerSwordSoul() {
	return (PlayerSwordSoul*)calloc(1, sizeof(PlayerSwordSoul));
}

void Free_PlayerSwordSoul(PlayerSwordSoul* row) {
	free(row);
}

/* ========== player_sword_soul_equipment ========== */

PlayerSwordSoulEquipment* New_PlayerSwordSoulEquipment() {
	return (PlayerSwordSoulEquipment*)calloc(1, sizeof(PlayerSwordSoulEquipment));
}

void Free_PlayerSwordSoulEquipment(PlayerSwordSoulEquipment* row) {
	free(row);
}

/* ========== player_sword_soul_state ========== */

PlayerSwordSoulState* New_PlayerSwordSoulState() {
	return (PlayerSwordSoulState*)calloc(1, sizeof(PlayerSwordSoulState));
}

void Free_PlayerSwordSoulState(PlayerSwordSoulState* row) {
	free(row);
}

/* ========== player_taoyuan_bless ========== */

PlayerTaoyuanBless* New_PlayerTaoyuanBless() {
	return (PlayerTaoyuanBless*)calloc(1, sizeof(PlayerTaoyuanBless));
}

void Free_PlayerTaoyuanBless(PlayerTaoyuanBless* row) {
	free(row);
}

/* ========== player_taoyuan_fileds ========== */

PlayerTaoyuanFileds* New_PlayerTaoyuanFileds() {
	return (PlayerTaoyuanFileds*)calloc(1, sizeof(PlayerTaoyuanFileds));
}

void Free_PlayerTaoyuanFileds(PlayerTaoyuanFileds* row) {
	free(row);
}

/* ========== player_taoyuan_heart ========== */

PlayerTaoyuanHeart* New_PlayerTaoyuanHeart() {
	return (PlayerTaoyuanHeart*)calloc(1, sizeof(PlayerTaoyuanHeart));
}

void Free_PlayerTaoyuanHeart(PlayerTaoyuanHeart* row) {
	free(row);
}

/* ========== player_taoyuan_item ========== */

PlayerTaoyuanItem* New_PlayerTaoyuanItem() {
	return (PlayerTaoyuanItem*)calloc(1, sizeof(PlayerTaoyuanItem));
}

void Free_PlayerTaoyuanItem(PlayerTaoyuanItem* row) {
	free(row);
}

/* ========== player_taoyuan_message ========== */

PlayerTaoyuanMessage* New_PlayerTaoyuanMessage() {
	return (PlayerTaoyuanMessage*)calloc(1, sizeof(PlayerTaoyuanMessage));
}

void Free_PlayerTaoyuanMessage(PlayerTaoyuanMessage* row) {
	if (row->Nick != NULL) {
		free(row->Nick);
	}
	free(row);
}

/* ========== player_taoyuan_product ========== */

PlayerTaoyuanProduct* New_PlayerTaoyuanProduct() {
	return (PlayerTaoyuanProduct*)calloc(1, sizeof(PlayerTaoyuanProduct));
}

void Free_PlayerTaoyuanProduct(PlayerTaoyuanProduct* row) {
	free(row);
}

/* ========== player_taoyuan_purchase_record ========== */

PlayerTaoyuanPurchaseRecord* New_PlayerTaoyuanPurchaseRecord() {
	return (PlayerTaoyuanPurchaseRecord*)calloc(1, sizeof(PlayerTaoyuanPurchaseRecord));
}

void Free_PlayerTaoyuanPurchaseRecord(PlayerTaoyuanPurchaseRecord* row) {
	free(row);
}

/* ========== player_taoyuan_quest ========== */

PlayerTaoyuanQuest* New_PlayerTaoyuanQuest() {
	return (PlayerTaoyuanQuest*)calloc(1, sizeof(PlayerTaoyuanQuest));
}

void Free_PlayerTaoyuanQuest(PlayerTaoyuanQuest* row) {
	free(row);
}

/* ========== player_tb_xxd_roleinfo ========== */

PlayerTbXxdRoleinfo* New_PlayerTbXxdRoleinfo() {
	return (PlayerTbXxdRoleinfo*)calloc(1, sizeof(PlayerTbXxdRoleinfo));
}

void Free_PlayerTbXxdRoleinfo(PlayerTbXxdRoleinfo* row) {
	if (row->Gameappid != NULL) {
		free(row->Gameappid);
	}
	if (row->Openid != NULL) {
		free(row->Openid);
	}
	free(row);
}

/* ========== player_team_info ========== */

PlayerTeamInfo* New_PlayerTeamInfo() {
	return (PlayerTeamInfo*)calloc(1, sizeof(PlayerTeamInfo));
}

void Free_PlayerTeamInfo(PlayerTeamInfo* row) {
	free(row);
}

/* ========== player_totem ========== */

PlayerTotem* New_PlayerTotem() {
	return (PlayerTotem*)calloc(1, sizeof(PlayerTotem));
}

void Free_PlayerTotem(PlayerTotem* row) {
	free(row);
}

/* ========== player_totem_info ========== */

PlayerTotemInfo* New_PlayerTotemInfo() {
	return (PlayerTotemInfo*)calloc(1, sizeof(PlayerTotemInfo));
}

void Free_PlayerTotemInfo(PlayerTotemInfo* row) {
	free(row);
}

/* ========== player_town ========== */

PlayerTown* New_PlayerTown() {
	return (PlayerTown*)calloc(1, sizeof(PlayerTown));
}

void Free_PlayerTown(PlayerTown* row) {
	free(row);
}

/* ========== player_trader_refresh_state ========== */

PlayerTraderRefreshState* New_PlayerTraderRefreshState() {
	return (PlayerTraderRefreshState*)calloc(1, sizeof(PlayerTraderRefreshState));
}

void Free_PlayerTraderRefreshState(PlayerTraderRefreshState* row) {
	free(row);
}

/* ========== player_trader_store_state ========== */

PlayerTraderStoreState* New_PlayerTraderStoreState() {
	return (PlayerTraderStoreState*)calloc(1, sizeof(PlayerTraderStoreState));
}

void Free_PlayerTraderStoreState(PlayerTraderStoreState* row) {
	free(row);
}

/* ========== player_use_skill ========== */

PlayerUseSkill* New_PlayerUseSkill() {
	return (PlayerUseSkill*)calloc(1, sizeof(PlayerUseSkill));
}

void Free_PlayerUseSkill(PlayerUseSkill* row) {
	free(row);
}

/* ========== player_vip ========== */

PlayerVip* New_PlayerVip() {
	return (PlayerVip*)calloc(1, sizeof(PlayerVip));
}

void Free_PlayerVip(PlayerVip* row) {
	if (row->CardId != NULL) {
		free(row->CardId);
	}
	free(row);
}

