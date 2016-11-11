#ifndef _cgo_types_h_
#define _cgo_types_h_

#include <stdint.h>
#include <stdlib.h>
#include <string.h>
extern void Init();

typedef struct GlobalTables GlobalTables;
typedef struct PlayerTables PlayerTables;

extern GlobalTables g_GlobalTables;
extern PlayerTables* NewPlayerTables();
/* ========== global_announcement ========== */

typedef struct global_announcement {
	int64_t Id;
	int64_t ExpireTime;
	int32_t TplId;
	char* Parameters;
	int Parameters_len;
	char* Content;
	int Content_len;
	int64_t SendTime;
	int64_t SpacingTime;
	struct global_announcement* next;
} GlobalAnnouncement;

extern GlobalAnnouncement* New_GlobalAnnouncement();
extern void Free_GlobalAnnouncement(GlobalAnnouncement*);

/* ========== global_arena_rank ========== */

typedef struct global_arena_rank {
	int32_t Rank;
	int64_t Pid;
	struct global_arena_rank* next;
} GlobalArenaRank;

extern GlobalArenaRank* New_GlobalArenaRank();
extern void Free_GlobalArenaRank(GlobalArenaRank*);

/* ========== global_clique ========== */

typedef struct global_clique {
	int64_t Id;
	char* Name;
	int Name_len;
	char* Announce;
	int Announce_len;
	int64_t TotalDonateCoins;
	int64_t OwnerPid;
	int64_t OwnerLoginTime;
	int64_t MangerPid1;
	int64_t MangerPid2;
	int64_t CenterBuildingCoins;
	int64_t TempleBuildingCoins;
	int64_t HealthBuildingCoins;
	int64_t AttackBuildingCoins;
	int64_t DefenseBuildingCoins;
	int64_t StoreBuildingCoins;
	int64_t BankBuildingCoins;
	int16_t CenterBuildingLevel;
	int16_t TempleBuildingLevel;
	int16_t HealthBuildingLevel;
	int16_t AttackBuildingLevel;
	int16_t DefenseBuildingLevel;
	int16_t BankBuildingLevel;
	char* MemberJson;
	int MemberJson_len;
	int8_t AutoAudit;
	int16_t AutoAuditLevel;
	int64_t Contrib;
	int64_t ContribClearTime;
	int64_t RecruitTime;
	int64_t WorshipTime;
	int8_t WorshipCnt;
	int64_t StoreSendTime;
	int8_t StoreSendCnt;
	struct global_clique* next;
} GlobalClique;

extern GlobalClique* New_GlobalClique();
extern void Free_GlobalClique(GlobalClique*);

/* ========== global_clique_boat ========== */

typedef struct global_clique_boat {
	int64_t Id;
	int64_t CliqueId;
	int16_t BoatType;
	int64_t OwnerPid;
	int64_t EscortStartTimestamp;
	int16_t TotalEscortTime;
	int64_t RecoverPid;
	int64_t HijackerPid;
	int64_t HijackTimestamp;
	int8_t Status;
	struct global_clique_boat* next;
} GlobalCliqueBoat;

extern GlobalCliqueBoat* New_GlobalCliqueBoat();
extern void Free_GlobalCliqueBoat(GlobalCliqueBoat*);

/* ========== global_despair_boss ========== */

typedef struct global_despair_boss {
	int8_t CampType;
	int16_t Level;
	int64_t Timestamp;
	int64_t DeadTimestamp;
	int64_t Hp;
	int64_t MaxHp;
	struct global_despair_boss* next;
} GlobalDespairBoss;

extern GlobalDespairBoss* New_GlobalDespairBoss();
extern void Free_GlobalDespairBoss(GlobalDespairBoss*);

/* ========== global_despair_boss_history ========== */

typedef struct global_despair_boss_history {
	int64_t Id;
	int64_t Version;
	int8_t CampType;
	int16_t Level;
	int64_t Timestamp;
	int64_t StartTimestamp;
	int64_t DeadTimestamp;
	struct global_despair_boss_history* next;
} GlobalDespairBossHistory;

extern GlobalDespairBossHistory* New_GlobalDespairBossHistory();
extern void Free_GlobalDespairBossHistory(GlobalDespairBossHistory*);

/* ========== global_despair_land ========== */

typedef struct global_despair_land {
	int64_t Id;
	int64_t Version;
	int64_t Timestamp;
	struct global_despair_land* next;
} GlobalDespairLand;

extern GlobalDespairLand* New_GlobalDespairLand();
extern void Free_GlobalDespairLand(GlobalDespairLand*);

/* ========== global_despair_land_battle_record ========== */

typedef struct global_despair_land_battle_record {
	int64_t Id;
	int64_t Pid;
	int32_t Fightnum;
	int64_t Timestamp;
	int16_t Tag;
	int16_t BattleVersion;
	int8_t CampType;
	int32_t LevelId;
	char* Record;
	int Record_len;
	struct global_despair_land_battle_record* next;
} GlobalDespairLandBattleRecord;

extern GlobalDespairLandBattleRecord* New_GlobalDespairLandBattleRecord();
extern void Free_GlobalDespairLandBattleRecord(GlobalDespairLandBattleRecord*);

/* ========== global_despair_land_camp ========== */

typedef struct global_despair_land_camp {
	int8_t CampType;
	int64_t Timestamp;
	int64_t DeadTimestamp;
	int16_t Level;
	struct global_despair_land_camp* next;
} GlobalDespairLandCamp;

extern GlobalDespairLandCamp* New_GlobalDespairLandCamp();
extern void Free_GlobalDespairLandCamp(GlobalDespairLandCamp*);

/* ========== global_gift_card_record ========== */

typedef struct global_gift_card_record {
	int64_t Id;
	int64_t Pid;
	char* Code;
	int Code_len;
	int64_t Timestamp;
	struct global_gift_card_record* next;
} GlobalGiftCardRecord;

extern GlobalGiftCardRecord* New_GlobalGiftCardRecord();
extern void Free_GlobalGiftCardRecord(GlobalGiftCardRecord*);

/* ========== global_group_buy_status ========== */

typedef struct global_group_buy_status {
	int64_t Id;
	int32_t Cid;
	int32_t Status;
	struct global_group_buy_status* next;
} GlobalGroupBuyStatus;

extern GlobalGroupBuyStatus* New_GlobalGroupBuyStatus();
extern void Free_GlobalGroupBuyStatus(GlobalGroupBuyStatus*);

/* ========== global_mail ========== */

typedef struct global_mail {
	int64_t Id;
	int64_t SendTime;
	int64_t ExpireTime;
	char* Title;
	int Title_len;
	char* Content;
	int Content_len;
	int8_t Priority;
	int16_t MinLevel;
	int16_t MaxLevel;
	int16_t MinVipLevel;
	int16_t MaxVipLevel;
	int64_t MinCreateTime;
	int64_t MaxCreateTime;
	struct global_mail* next;
} GlobalMail;

extern GlobalMail* New_GlobalMail();
extern void Free_GlobalMail(GlobalMail*);

/* ========== global_mail_attachments ========== */

typedef struct global_mail_attachments {
	int64_t Id;
	int64_t GlobalMailId;
	int16_t ItemId;
	int8_t AttachmentType;
	int64_t ItemNum;
	struct global_mail_attachments* next;
} GlobalMailAttachments;

extern GlobalMailAttachments* New_GlobalMailAttachments();
extern void Free_GlobalMailAttachments(GlobalMailAttachments*);

/* ========== global_tb_xxd_onlinecnt ========== */

typedef struct global_tb_xxd_onlinecnt {
	int64_t Id;
	char* Gameappid;
	int Gameappid_len;
	int64_t Timekey;
	int64_t Gsid;
	int64_t Onlinecntios;
	int64_t Onlinecntandroid;
	int64_t Cid;
	struct global_tb_xxd_onlinecnt* next;
} GlobalTbXxdOnlinecnt;

extern GlobalTbXxdOnlinecnt* New_GlobalTbXxdOnlinecnt();
extern void Free_GlobalTbXxdOnlinecnt(GlobalTbXxdOnlinecnt*);

/* ========== player ========== */

typedef struct player {
	int64_t Id;
	char* User;
	int User_len;
	char* Nick;
	int Nick_len;
	int64_t MainRoleId;
	int64_t Cid;
} Player;

extern Player* New_Player();
extern void Free_Player(Player*);

/* ========== player_activity ========== */

typedef struct player_activity {
	int64_t Pid;
	int32_t Activity;
	int64_t LastUpdate;
} PlayerActivity;

extern PlayerActivity* New_PlayerActivity();
extern void Free_PlayerActivity(PlayerActivity*);

/* ========== player_addition_quest ========== */

typedef struct player_addition_quest {
	int64_t Id;
	int64_t Pid;
	int32_t SerialNumber;
	int32_t QuestId;
	int16_t Lock;
	int16_t Progress;
	int8_t State;
	struct player_addition_quest* next;
} PlayerAdditionQuest;

extern PlayerAdditionQuest* New_PlayerAdditionQuest();
extern void Free_PlayerAdditionQuest(PlayerAdditionQuest*);

/* ========== player_any_sdk_order ========== */

typedef struct player_any_sdk_order {
	int64_t Id;
	int64_t Pid;
	int64_t OrderId;
	int64_t Present;
	struct player_any_sdk_order* next;
} PlayerAnySdkOrder;

extern PlayerAnySdkOrder* New_PlayerAnySdkOrder();
extern void Free_PlayerAnySdkOrder(PlayerAnySdkOrder*);

/* ========== player_arena ========== */

typedef struct player_arena {
	int64_t Pid;
	int16_t DailyNum;
	int64_t FailedCdTime;
	int64_t BattleTime;
	int16_t WinTimes;
	int32_t DailyAwardItem;
	int16_t NewRecordCount;
	int32_t DailyFame;
	int16_t BuyTimes;
} PlayerArena;

extern PlayerArena* New_PlayerArena();
extern void Free_PlayerArena(PlayerArena*);

/* ========== player_arena_rank ========== */

typedef struct player_arena_rank {
	int64_t Pid;
	int32_t Rank;
	int32_t Rank1;
	int32_t Rank2;
	int32_t Rank3;
	int64_t Time;
} PlayerArenaRank;

extern PlayerArenaRank* New_PlayerArenaRank();
extern void Free_PlayerArenaRank(PlayerArenaRank*);

/* ========== player_arena_record ========== */

typedef struct player_arena_record {
	int64_t Id;
	int64_t Pid;
	int8_t Mode;
	int32_t OldRank;
	int32_t NewRank;
	int32_t FightNum;
	int64_t TargetPid;
	char* TargetNick;
	int TargetNick_len;
	int32_t TargetOldRank;
	int32_t TargetNewRank;
	int32_t TargetFightNum;
	int64_t RecordTime;
	struct player_arena_record* next;
} PlayerArenaRecord;

extern PlayerArenaRecord* New_PlayerArenaRecord();
extern void Free_PlayerArenaRecord(PlayerArenaRecord*);

/* ========== player_awaken_graphic ========== */

typedef struct player_awaken_graphic {
	int64_t Id;
	int64_t Pid;
	int8_t RoleId;
	int16_t AttrImpl;
	int8_t Level;
	struct player_awaken_graphic* next;
} PlayerAwakenGraphic;

extern PlayerAwakenGraphic* New_PlayerAwakenGraphic();
extern void Free_PlayerAwakenGraphic(PlayerAwakenGraphic*);

/* ========== player_battle_pet ========== */

typedef struct player_battle_pet {
	int64_t Id;
	int64_t Pid;
	int32_t BattlePetId;
	int32_t Level;
	int64_t Exp;
	int32_t SkillLevel;
	struct player_battle_pet* next;
} PlayerBattlePet;

extern PlayerBattlePet* New_PlayerBattlePet();
extern void Free_PlayerBattlePet(PlayerBattlePet*);

/* ========== player_battle_pet_grid ========== */

typedef struct player_battle_pet_grid {
	int64_t Id;
	int64_t Pid;
	int8_t GridId;
	int32_t BattlePetId;
	struct player_battle_pet_grid* next;
} PlayerBattlePetGrid;

extern PlayerBattlePetGrid* New_PlayerBattlePetGrid();
extern void Free_PlayerBattlePetGrid(PlayerBattlePetGrid*);

/* ========== player_battle_pet_state ========== */

typedef struct player_battle_pet_state {
	int64_t Pid;
	int32_t UpgradeByIngotNum;
	int64_t UpgradeByIngotTime;
} PlayerBattlePetState;

extern PlayerBattlePetState* New_PlayerBattlePetState();
extern void Free_PlayerBattlePetState(PlayerBattlePetState*);

/* ========== player_chest_state ========== */

typedef struct player_chest_state {
	int64_t Pid;
	int32_t FreeCoinChestNum;
	int64_t LastFreeCoinChestAt;
	int32_t CoinChestNum;
	int32_t CoinChestTenNum;
	int8_t IsFirstCoinTen;
	int64_t CoinChestConsume;
	int64_t LastCoinChestAt;
	int64_t LastFreeIngotChestAt;
	int32_t IngotChestNum;
	int32_t IngotChestTenNum;
	int8_t IsFirstIngotTen;
	int64_t IngotChestConsume;
	int64_t LastIngotChestAt;
	int64_t LastFreePetChestAt;
} PlayerChestState;

extern PlayerChestState* New_PlayerChestState();
extern void Free_PlayerChestState(PlayerChestState*);

/* ========== player_clique_kongfu_attrib ========== */

typedef struct player_clique_kongfu_attrib {
	int64_t Pid;
	int32_t Health;
	int32_t Attack;
	int32_t Defence;
} PlayerCliqueKongfuAttrib;

extern PlayerCliqueKongfuAttrib* New_PlayerCliqueKongfuAttrib();
extern void Free_PlayerCliqueKongfuAttrib(PlayerCliqueKongfuAttrib*);

/* ========== player_coins ========== */

typedef struct player_coins {
	int64_t Pid;
	int64_t BuyTime;
	int16_t DailyCount;
	int16_t BatchBought;
} PlayerCoins;

extern PlayerCoins* New_PlayerCoins();
extern void Free_PlayerCoins(PlayerCoins*);

/* ========== player_cornucopia ========== */

typedef struct player_cornucopia {
	int64_t Pid;
	int64_t OpenTime;
	int16_t DailyCount;
} PlayerCornucopia;

extern PlayerCornucopia* New_PlayerCornucopia();
extern void Free_PlayerCornucopia(PlayerCornucopia*);

/* ========== player_daily_quest ========== */

typedef struct player_daily_quest {
	int64_t Id;
	int64_t Pid;
	int16_t QuestId;
	int16_t FinishCount;
	int64_t LastUpdateTime;
	int8_t AwardStatus;
	int16_t Class;
	struct player_daily_quest* next;
} PlayerDailyQuest;

extern PlayerDailyQuest* New_PlayerDailyQuest();
extern void Free_PlayerDailyQuest(PlayerDailyQuest*);

/* ========== player_daily_quest_star_award_info ========== */

typedef struct player_daily_quest_star_award_info {
	int64_t Pid;
	int32_t Stars;
	int64_t Lastupdatetime;
	char* Awarded;
	int Awarded_len;
} PlayerDailyQuestStarAwardInfo;

extern PlayerDailyQuestStarAwardInfo* New_PlayerDailyQuestStarAwardInfo();
extern void Free_PlayerDailyQuestStarAwardInfo(PlayerDailyQuestStarAwardInfo*);

/* ========== player_daily_sign_in_record ========== */

typedef struct player_daily_sign_in_record {
	int64_t Id;
	int64_t Pid;
	int64_t SignInTime;
	struct player_daily_sign_in_record* next;
} PlayerDailySignInRecord;

extern PlayerDailySignInRecord* New_PlayerDailySignInRecord();
extern void Free_PlayerDailySignInRecord(PlayerDailySignInRecord*);

/* ========== player_daily_sign_in_state ========== */

typedef struct player_daily_sign_in_state {
	int64_t Pid;
	int64_t UpdateTime;
	int16_t Record;
	int8_t SignedToday;
} PlayerDailySignInState;

extern PlayerDailySignInState* New_PlayerDailySignInState();
extern void Free_PlayerDailySignInState(PlayerDailySignInState*);

/* ========== player_despair_land_camp_state ========== */

typedef struct player_despair_land_camp_state {
	int64_t Id;
	int64_t Pid;
	int8_t CampType;
	int64_t Timestamp;
	int64_t Point;
	int64_t KillNum;
	int64_t DeadNum;
	int64_t DeadNumBoss;
	int64_t Hurt;
	int32_t BossBattleNum;
	int32_t DailyBossBattleNum;
	int64_t BossBattleTimestamp;
	int8_t Awarded;
	struct player_despair_land_camp_state* next;
} PlayerDespairLandCampState;

extern PlayerDespairLandCampState* New_PlayerDespairLandCampState();
extern void Free_PlayerDespairLandCampState(PlayerDespairLandCampState*);

/* ========== player_despair_land_camp_state_history ========== */

typedef struct player_despair_land_camp_state_history {
	int64_t Id;
	int64_t Pid;
	int8_t CampType;
	int64_t Timestamp;
	int64_t Point;
	int64_t KillNum;
	int64_t DeadNum;
	int64_t DeadNumBoss;
	int64_t Hurt;
	int32_t BossBattleNum;
	int32_t DailyBossBattleNum;
	int64_t BossBattleTimestamp;
	int8_t Awarded;
	struct player_despair_land_camp_state_history* next;
} PlayerDespairLandCampStateHistory;

extern PlayerDespairLandCampStateHistory* New_PlayerDespairLandCampStateHistory();
extern void Free_PlayerDespairLandCampStateHistory(PlayerDespairLandCampStateHistory*);

/* ========== player_despair_land_level_record ========== */

typedef struct player_despair_land_level_record {
	int64_t Id;
	int64_t Pid;
	int8_t CampType;
	int64_t Timestamp;
	int32_t LevelId;
	int8_t Round;
	int64_t PassedTimestamp;
	int64_t FirstPassedTimestamp;
	int32_t FirstPassedFightnum;
	int8_t PassedAward;
	int8_t PerfectAward;
	int8_t MilestoneAward;
	struct player_despair_land_level_record* next;
} PlayerDespairLandLevelRecord;

extern PlayerDespairLandLevelRecord* New_PlayerDespairLandLevelRecord();
extern void Free_PlayerDespairLandLevelRecord(PlayerDespairLandLevelRecord*);

/* ========== player_despair_land_state ========== */

typedef struct player_despair_land_state {
	int64_t Pid;
	int64_t Point;
	int64_t KillNum;
	int64_t DeadNum;
	int16_t DailyBattleNum;
	int64_t DailyBattleTimestamp;
	int16_t DailyBoughtBattleNum;
	int64_t DailyBoughtTimestamp;
	int64_t DailyBossBoughtTimestamp;
	int16_t DailyBossBeastBoughtBattleNum;
	int16_t DailyBossWalkingDeadBoughtBattleNum;
	int16_t DailyBossDevilBoughtBattleNum;
} PlayerDespairLandState;

extern PlayerDespairLandState* New_PlayerDespairLandState();
extern void Free_PlayerDespairLandState(PlayerDespairLandState*);

/* ========== player_driving_sword_event ========== */

typedef struct player_driving_sword_event {
	int64_t Id;
	int64_t Pid;
	int16_t CloudId;
	int8_t X;
	int8_t Y;
	int8_t EventType;
	int8_t DataId;
	struct player_driving_sword_event* next;
} PlayerDrivingSwordEvent;

extern PlayerDrivingSwordEvent* New_PlayerDrivingSwordEvent();
extern void Free_PlayerDrivingSwordEvent(PlayerDrivingSwordEvent*);

/* ========== player_driving_sword_event_exploring ========== */

typedef struct player_driving_sword_event_exploring {
	int64_t Id;
	int64_t Pid;
	int8_t Status;
	int64_t GarrisonStart;
	int64_t GarrisonTime;
	int64_t AwardTime;
	int8_t RoleId;
	int16_t CloudId;
	int8_t X;
	int8_t Y;
	int8_t DataId;
	struct player_driving_sword_event_exploring* next;
} PlayerDrivingSwordEventExploring;

extern PlayerDrivingSwordEventExploring* New_PlayerDrivingSwordEventExploring();
extern void Free_PlayerDrivingSwordEventExploring(PlayerDrivingSwordEventExploring*);

/* ========== player_driving_sword_event_treasure ========== */

typedef struct player_driving_sword_event_treasure {
	int64_t Id;
	int64_t Pid;
	int8_t Progress;
	int16_t CloudId;
	int8_t X;
	int8_t Y;
	int8_t DataId;
	struct player_driving_sword_event_treasure* next;
} PlayerDrivingSwordEventTreasure;

extern PlayerDrivingSwordEventTreasure* New_PlayerDrivingSwordEventTreasure();
extern void Free_PlayerDrivingSwordEventTreasure(PlayerDrivingSwordEventTreasure*);

/* ========== player_driving_sword_event_visiting ========== */

typedef struct player_driving_sword_event_visiting {
	int64_t Id;
	int64_t Pid;
	int8_t Status;
	int64_t TargetPid;
	char* TargetSideState;
	int TargetSideState_len;
	int16_t CloudId;
	int8_t X;
	int8_t Y;
	int8_t DataId;
	char* TargetStatus;
	int TargetStatus_len;
	struct player_driving_sword_event_visiting* next;
} PlayerDrivingSwordEventVisiting;

extern PlayerDrivingSwordEventVisiting* New_PlayerDrivingSwordEventVisiting();
extern void Free_PlayerDrivingSwordEventVisiting(PlayerDrivingSwordEventVisiting*);

/* ========== player_driving_sword_eventmask ========== */

typedef struct player_driving_sword_eventmask {
	int64_t Id;
	int64_t Pid;
	int16_t CloudId;
	char* Events;
	int Events_len;
	struct player_driving_sword_eventmask* next;
} PlayerDrivingSwordEventmask;

extern PlayerDrivingSwordEventmask* New_PlayerDrivingSwordEventmask();
extern void Free_PlayerDrivingSwordEventmask(PlayerDrivingSwordEventmask*);

/* ========== player_driving_sword_info ========== */

typedef struct player_driving_sword_info {
	int64_t Pid;
	int16_t CurrentCloud;
	int16_t HighestCloud;
	int8_t CurrentX;
	int8_t CurrentY;
	int16_t AllowedAction;
	int64_t ActionRefreshTime;
	int64_t ActionBuyTime;
	int8_t DailyActionBought;
} PlayerDrivingSwordInfo;

extern PlayerDrivingSwordInfo* New_PlayerDrivingSwordInfo();
extern void Free_PlayerDrivingSwordInfo(PlayerDrivingSwordInfo*);

/* ========== player_driving_sword_map ========== */

typedef struct player_driving_sword_map {
	int64_t Id;
	int64_t Pid;
	int16_t CloudId;
	char* Shadows;
	int Shadows_len;
	int8_t ObstacleCount;
	int8_t ExploringCount;
	int8_t VisitingCount;
	int8_t TreasureCount;
	struct player_driving_sword_map* next;
} PlayerDrivingSwordMap;

extern PlayerDrivingSwordMap* New_PlayerDrivingSwordMap();
extern void Free_PlayerDrivingSwordMap(PlayerDrivingSwordMap*);

/* ========== player_equipment ========== */

typedef struct player_equipment {
	int64_t Id;
	int64_t Pid;
	int8_t RoleId;
	int64_t WeaponId;
	int64_t ClothesId;
	int64_t AccessoriesId;
	int64_t ShoeId;
	struct player_equipment* next;
} PlayerEquipment;

extern PlayerEquipment* New_PlayerEquipment();
extern void Free_PlayerEquipment(PlayerEquipment*);

/* ========== player_event_award_record ========== */

typedef struct player_event_award_record {
	int64_t Pid;
	char* RecordBytes;
	int RecordBytes_len;
	char* JsonEventRecord;
	int JsonEventRecord_len;
} PlayerEventAwardRecord;

extern PlayerEventAwardRecord* New_PlayerEventAwardRecord();
extern void Free_PlayerEventAwardRecord(PlayerEventAwardRecord*);

/* ========== player_event_daily_online ========== */

typedef struct player_event_daily_online {
	int64_t Pid;
	int64_t Timestamp;
	int32_t TotalOnlineTime;
	int32_t AwardedOnlinetime;
} PlayerEventDailyOnline;

extern PlayerEventDailyOnline* New_PlayerEventDailyOnline();
extern void Free_PlayerEventDailyOnline(PlayerEventDailyOnline*);

/* ========== player_event_vn_daily_recharge ========== */

typedef struct player_event_vn_daily_recharge {
	int64_t Id;
	int64_t Pid;
	int32_t Page;
	int64_t Timestamp;
	int8_t Awarded;
	int64_t DailyRecharge;
	int64_t TotalRecharge;
	struct player_event_vn_daily_recharge* next;
} PlayerEventVnDailyRecharge;

extern PlayerEventVnDailyRecharge* New_PlayerEventVnDailyRecharge();
extern void Free_PlayerEventVnDailyRecharge(PlayerEventVnDailyRecharge*);

/* ========== player_event_vn_dragon_ball_record ========== */

typedef struct player_event_vn_dragon_ball_record {
	int64_t Pid;
	int16_t ItemId;
	int64_t Createtime;
} PlayerEventVnDragonBallRecord;

extern PlayerEventVnDragonBallRecord* New_PlayerEventVnDragonBallRecord();
extern void Free_PlayerEventVnDragonBallRecord(PlayerEventVnDragonBallRecord*);

/* ========== player_events_fight_power ========== */

typedef struct player_events_fight_power {
	int64_t Pid;
	int32_t Lock;
} PlayerEventsFightPower;

extern PlayerEventsFightPower* New_PlayerEventsFightPower();
extern void Free_PlayerEventsFightPower(PlayerEventsFightPower*);

/* ========== player_events_ingot_record ========== */

typedef struct player_events_ingot_record {
	int64_t Pid;
	int64_t IngotIn;
	int64_t IngotInEndTime;
	int64_t IngotOut;
	int64_t IngotOutEndTime;
} PlayerEventsIngotRecord;

extern PlayerEventsIngotRecord* New_PlayerEventsIngotRecord();
extern void Free_PlayerEventsIngotRecord(PlayerEventsIngotRecord*);

/* ========== player_extend_level ========== */

typedef struct player_extend_level {
	int64_t Pid;
	int64_t CoinPassTime;
	int64_t ExpPassTime;
	int64_t GhostPassTime;
	int64_t PetPassTime;
	int64_t BuddyPassTime;
	int8_t CoinDailyNum;
	int8_t ExpDailyNum;
	int8_t BuddyDailyNum;
	int8_t PetDailyNum;
	int8_t GhostDailyNum;
	int8_t RandBuddyRoleId;
	int8_t BuddyPos;
	int8_t BuddyTactical;
	int8_t RolePos;
	int16_t ExpMaxlevel;
	int16_t CoinsMaxlevel;
	int16_t CoinsBuyNum;
	int16_t ExpBuyNum;
	int64_t CoinsBuyTime;
	int64_t ExpBuyTime;
} PlayerExtendLevel;

extern PlayerExtendLevel* New_PlayerExtendLevel();
extern void Free_PlayerExtendLevel(PlayerExtendLevel*);

/* ========== player_extend_quest ========== */

typedef struct player_extend_quest {
	int64_t Id;
	int64_t Pid;
	int32_t QuestId;
	int16_t Progress;
	int8_t State;
	struct player_extend_quest* next;
} PlayerExtendQuest;

extern PlayerExtendQuest* New_PlayerExtendQuest();
extern void Free_PlayerExtendQuest(PlayerExtendQuest*);

/* ========== player_fame ========== */

typedef struct player_fame {
	int64_t Pid;
	int32_t Fame;
	int16_t Level;
	int32_t ArenaFame;
	int32_t MultLevelFame;
} PlayerFame;

extern PlayerFame* New_PlayerFame();
extern void Free_PlayerFame(PlayerFame*);

/* ========== player_fashion ========== */

typedef struct player_fashion {
	int64_t Id;
	int64_t Pid;
	int16_t FashionId;
	int64_t ExpireTime;
	struct player_fashion* next;
} PlayerFashion;

extern PlayerFashion* New_PlayerFashion();
extern void Free_PlayerFashion(PlayerFashion*);

/* ========== player_fashion_state ========== */

typedef struct player_fashion_state {
	int64_t Pid;
	int64_t UpdateTime;
	int16_t DressedFashionId;
} PlayerFashionState;

extern PlayerFashionState* New_PlayerFashionState();
extern void Free_PlayerFashionState(PlayerFashionState*);

/* ========== player_fate_box_state ========== */

typedef struct player_fate_box_state {
	int64_t Pid;
	int32_t Lock;
	int64_t StarBoxFreeDrawTimestamp;
	int32_t StarBoxDrawCount;
	int64_t MoonBoxFreeDrawTimestamp;
	int32_t MoonBoxDrawCount;
	int64_t SunBoxFreeDrawTimestamp;
	int32_t SunBoxDrawCount;
	int64_t HunyuanBoxFreeDrawTimestamp;
	int32_t HunyuanBoxDrawCount;
} PlayerFateBoxState;

extern PlayerFateBoxState* New_PlayerFateBoxState();
extern void Free_PlayerFateBoxState(PlayerFateBoxState*);

/* ========== player_fight_num ========== */

typedef struct player_fight_num {
	int64_t Pid;
	int32_t FightNum;
} PlayerFightNum;

extern PlayerFightNum* New_PlayerFightNum();
extern void Free_PlayerFightNum(PlayerFightNum*);

/* ========== player_first_charge ========== */

typedef struct player_first_charge {
	int64_t Id;
	int64_t Pid;
	char* ProductId;
	int ProductId_len;
	struct player_first_charge* next;
} PlayerFirstCharge;

extern PlayerFirstCharge* New_PlayerFirstCharge();
extern void Free_PlayerFirstCharge(PlayerFirstCharge*);

/* ========== player_formation ========== */

typedef struct player_formation {
	int64_t Pid;
	int8_t Pos0;
	int8_t Pos1;
	int8_t Pos2;
	int8_t Pos3;
	int8_t Pos4;
	int8_t Pos5;
	int8_t Pos6;
	int8_t Pos7;
	int8_t Pos8;
} PlayerFormation;

extern PlayerFormation* New_PlayerFormation();
extern void Free_PlayerFormation(PlayerFormation*);

/* ========== player_func_key ========== */

typedef struct player_func_key {
	int64_t Pid;
	int16_t Key;
	int16_t PlayedKey;
	int64_t UniqueKey;
} PlayerFuncKey;

extern PlayerFuncKey* New_PlayerFuncKey();
extern void Free_PlayerFuncKey(PlayerFuncKey*);

/* ========== player_ghost ========== */

typedef struct player_ghost {
	int64_t Id;
	int64_t Pid;
	int16_t GhostId;
	int8_t Star;
	int16_t Level;
	int64_t Exp;
	int16_t Pos;
	int8_t IsNew;
	int8_t RoleId;
	int16_t SkillLevel;
	int64_t RelationId;
	int16_t AddGrowth;
	struct player_ghost* next;
} PlayerGhost;

extern PlayerGhost* New_PlayerGhost();
extern void Free_PlayerGhost(PlayerGhost*);

/* ========== player_ghost_equipment ========== */

typedef struct player_ghost_equipment {
	int64_t Id;
	int64_t Pid;
	int8_t RoleId;
	int32_t GhostPower;
	int64_t Pos1;
	int64_t Pos2;
	int64_t Pos3;
	int64_t Pos4;
	struct player_ghost_equipment* next;
} PlayerGhostEquipment;

extern PlayerGhostEquipment* New_PlayerGhostEquipment();
extern void Free_PlayerGhostEquipment(PlayerGhostEquipment*);

/* ========== player_ghost_state ========== */

typedef struct player_ghost_state {
	int64_t Pid;
	int32_t TrainByIngotNum;
	int64_t TrainByIngotTime;
	int64_t LastFlushTime;
	int64_t GhostFightNum;
} PlayerGhostState;

extern PlayerGhostState* New_PlayerGhostState();
extern void Free_PlayerGhostState(PlayerGhostState*);

/* ========== player_global_chat_state ========== */

typedef struct player_global_chat_state {
	int64_t Pid;
	int64_t BanUntil;
} PlayerGlobalChatState;

extern PlayerGlobalChatState* New_PlayerGlobalChatState();
extern void Free_PlayerGlobalChatState(PlayerGlobalChatState*);

/* ========== player_global_clique_building ========== */

typedef struct player_global_clique_building {
	int64_t Pid;
	int64_t SilverExchangeTime;
	int16_t GoldExchangeNum;
	int16_t SilverExchangeNum;
	int64_t DonateCoinsCenterBuilding;
	int64_t DonateCoinsTempleBuilding;
	int64_t DonateCoinsBankBuilding;
	int64_t DonateCoinsHealthBuilding;
	int64_t DonateCoinsAttackBuilding;
	int64_t DonateCoinsDefenseBuilding;
	int64_t DonateCoinsStoreBuilding;
	int16_t HealthLevel;
	int16_t AttackLevel;
	int16_t DefenseLevel;
	int64_t WorshipTime;
	int64_t DonateCoinsCenterBuildingTime;
	int64_t GlodExchangeTime;
	int8_t WorshipType;
} PlayerGlobalCliqueBuilding;

extern PlayerGlobalCliqueBuilding* New_PlayerGlobalCliqueBuilding();
extern void Free_PlayerGlobalCliqueBuilding(PlayerGlobalCliqueBuilding*);

/* ========== player_global_clique_building_quest ========== */

typedef struct player_global_clique_building_quest {
	int64_t Id;
	int64_t Pid;
	int16_t QuestId;
	int8_t AwardStatus;
	int16_t BuildingType;
	struct player_global_clique_building_quest* next;
} PlayerGlobalCliqueBuildingQuest;

extern PlayerGlobalCliqueBuildingQuest* New_PlayerGlobalCliqueBuildingQuest();
extern void Free_PlayerGlobalCliqueBuildingQuest(PlayerGlobalCliqueBuildingQuest*);

/* ========== player_global_clique_daily_quest ========== */

typedef struct player_global_clique_daily_quest {
	int64_t Id;
	int64_t Pid;
	int16_t QuestId;
	int16_t FinishCount;
	int64_t LastUpdateTime;
	int8_t AwardStatus;
	int16_t Class;
	struct player_global_clique_daily_quest* next;
} PlayerGlobalCliqueDailyQuest;

extern PlayerGlobalCliqueDailyQuest* New_PlayerGlobalCliqueDailyQuest();
extern void Free_PlayerGlobalCliqueDailyQuest(PlayerGlobalCliqueDailyQuest*);

/* ========== player_global_clique_escort ========== */

typedef struct player_global_clique_escort {
	int64_t Pid;
	int64_t DailyEscortTimestamp;
	int16_t DailyEscortNum;
	int16_t EscortBoatType;
	int8_t Status;
	int8_t Hijacked;
} PlayerGlobalCliqueEscort;

extern PlayerGlobalCliqueEscort* New_PlayerGlobalCliqueEscort();
extern void Free_PlayerGlobalCliqueEscort(PlayerGlobalCliqueEscort*);

/* ========== player_global_clique_escort_message ========== */

typedef struct player_global_clique_escort_message {
	int64_t Id;
	int64_t Pid;
	int16_t TplId;
	char* Parameters;
	int Parameters_len;
	int64_t Timestamp;
	struct player_global_clique_escort_message* next;
} PlayerGlobalCliqueEscortMessage;

extern PlayerGlobalCliqueEscortMessage* New_PlayerGlobalCliqueEscortMessage();
extern void Free_PlayerGlobalCliqueEscortMessage(PlayerGlobalCliqueEscortMessage*);

/* ========== player_global_clique_hijack ========== */

typedef struct player_global_clique_hijack {
	int64_t Pid;
	int64_t DailyHijackTimestamp;
	int16_t DailyHijackNum;
	int16_t HijackedBoatType;
	int8_t Status;
	int64_t BuyTimestamp;
	int16_t BuyNum;
} PlayerGlobalCliqueHijack;

extern PlayerGlobalCliqueHijack* New_PlayerGlobalCliqueHijack();
extern void Free_PlayerGlobalCliqueHijack(PlayerGlobalCliqueHijack*);

/* ========== player_global_clique_info ========== */

typedef struct player_global_clique_info {
	int64_t Pid;
	int64_t CliqueId;
	int64_t JoinTime;
	int64_t Contrib;
	int64_t ContribClearTime;
	int64_t DonateCoinsTime;
	int64_t DailyDonateCoins;
	int64_t TotalContrib;
} PlayerGlobalCliqueInfo;

extern PlayerGlobalCliqueInfo* New_PlayerGlobalCliqueInfo();
extern void Free_PlayerGlobalCliqueInfo(PlayerGlobalCliqueInfo*);

/* ========== player_global_clique_kongfu ========== */

typedef struct player_global_clique_kongfu {
	int64_t Id;
	int64_t Pid;
	int32_t KongfuId;
	int16_t Level;
	struct player_global_clique_kongfu* next;
} PlayerGlobalCliqueKongfu;

extern PlayerGlobalCliqueKongfu* New_PlayerGlobalCliqueKongfu();
extern void Free_PlayerGlobalCliqueKongfu(PlayerGlobalCliqueKongfu*);

/* ========== player_global_friend ========== */

typedef struct player_global_friend {
	int64_t Id;
	int64_t Pid;
	int64_t FriendPid;
	char* FriendNick;
	int FriendNick_len;
	int8_t FriendRoleId;
	int8_t FriendMode;
	int64_t LastChatTime;
	int64_t FriendTime;
	int64_t SendHeartTime;
	int8_t BlockMode;
	struct player_global_friend* next;
} PlayerGlobalFriend;

extern PlayerGlobalFriend* New_PlayerGlobalFriend();
extern void Free_PlayerGlobalFriend(PlayerGlobalFriend*);

/* ========== player_global_friend_chat ========== */

typedef struct player_global_friend_chat {
	int64_t Id;
	int64_t Pid;
	int64_t FriendPid;
	int8_t Mode;
	int64_t SendTime;
	char* Message;
	int Message_len;
	struct player_global_friend_chat* next;
} PlayerGlobalFriendChat;

extern PlayerGlobalFriendChat* New_PlayerGlobalFriendChat();
extern void Free_PlayerGlobalFriendChat(PlayerGlobalFriendChat*);

/* ========== player_global_friend_state ========== */

typedef struct player_global_friend_state {
	int64_t Pid;
	int32_t DeleteDayCount;
	int64_t DeleteTime;
	int8_t ExistOfflineChat;
	int32_t PlatformFriendNum;
} PlayerGlobalFriendState;

extern PlayerGlobalFriendState* New_PlayerGlobalFriendState();
extern void Free_PlayerGlobalFriendState(PlayerGlobalFriendState*);

/* ========== player_global_gift_card_record ========== */

typedef struct player_global_gift_card_record {
	int64_t Id;
	int64_t Pid;
	char* Code;
	int Code_len;
	struct player_global_gift_card_record* next;
} PlayerGlobalGiftCardRecord;

extern PlayerGlobalGiftCardRecord* New_PlayerGlobalGiftCardRecord();
extern void Free_PlayerGlobalGiftCardRecord(PlayerGlobalGiftCardRecord*);

/* ========== player_global_mail_state ========== */

typedef struct player_global_mail_state {
	int64_t Pid;
	int64_t MaxTimestamp;
} PlayerGlobalMailState;

extern PlayerGlobalMailState* New_PlayerGlobalMailState();
extern void Free_PlayerGlobalMailState(PlayerGlobalMailState*);

/* ========== player_global_world_chat_state ========== */

typedef struct player_global_world_chat_state {
	int64_t Pid;
	int64_t Timestamp;
	int16_t DailyNum;
} PlayerGlobalWorldChatState;

extern PlayerGlobalWorldChatState* New_PlayerGlobalWorldChatState();
extern void Free_PlayerGlobalWorldChatState(PlayerGlobalWorldChatState*);

/* ========== player_hard_level ========== */

typedef struct player_hard_level {
	int64_t Pid;
	int32_t Lock;
	int32_t AwardLock;
} PlayerHardLevel;

extern PlayerHardLevel* New_PlayerHardLevel();
extern void Free_PlayerHardLevel(PlayerHardLevel*);

/* ========== player_hard_level_record ========== */

typedef struct player_hard_level_record {
	int64_t Id;
	int64_t Pid;
	int32_t LevelId;
	int64_t OpenTime;
	int32_t Score;
	int8_t Round;
	int8_t DailyNum;
	int64_t LastEnterTime;
	int16_t BuyTimes;
	int64_t BuyUpdateTime;
	struct player_hard_level_record* next;
} PlayerHardLevelRecord;

extern PlayerHardLevelRecord* New_PlayerHardLevelRecord();
extern void Free_PlayerHardLevelRecord(PlayerHardLevelRecord*);

/* ========== player_heart ========== */

typedef struct player_heart {
	int64_t Pid;
	int16_t Value;
	int64_t UpdateTime;
	int32_t AddDayCount;
	int64_t AddTime;
	int16_t RecoverDayCount;
} PlayerHeart;

extern PlayerHeart* New_PlayerHeart();
extern void Free_PlayerHeart(PlayerHeart*);

/* ========== player_heart_draw ========== */

typedef struct player_heart_draw {
	int64_t Id;
	int64_t Pid;
	int8_t DrawType;
	int8_t DailyNum;
	int64_t DrawTime;
	struct player_heart_draw* next;
} PlayerHeartDraw;

extern PlayerHeartDraw* New_PlayerHeartDraw();
extern void Free_PlayerHeartDraw(PlayerHeartDraw*);

/* ========== player_heart_draw_card_record ========== */

typedef struct player_heart_draw_card_record {
	int64_t Id;
	int64_t Pid;
	int8_t AwardType;
	int16_t AwardNum;
	int16_t ItemId;
	int64_t DrawTime;
	struct player_heart_draw_card_record* next;
} PlayerHeartDrawCardRecord;

extern PlayerHeartDrawCardRecord* New_PlayerHeartDrawCardRecord();
extern void Free_PlayerHeartDrawCardRecord(PlayerHeartDrawCardRecord*);

/* ========== player_heart_draw_wheel_record ========== */

typedef struct player_heart_draw_wheel_record {
	int64_t Id;
	int64_t Pid;
	int8_t AwardType;
	int16_t AwardNum;
	int16_t ItemId;
	int64_t DrawTime;
	struct player_heart_draw_wheel_record* next;
} PlayerHeartDrawWheelRecord;

extern PlayerHeartDrawWheelRecord* New_PlayerHeartDrawWheelRecord();
extern void Free_PlayerHeartDrawWheelRecord(PlayerHeartDrawWheelRecord*);

/* ========== player_info ========== */

typedef struct player_info {
	int64_t Pid;
	int64_t Ingot;
	int64_t Coins;
	int16_t NewMailNum;
	int64_t LastLoginTime;
	int64_t LastOfflineTime;
	int64_t TotalOnlineTime;
	int64_t FirstLoginTime;
	int16_t NewArenaReportNum;
	int64_t LastSkillFlush;
	int64_t SwordSoulFragment;
} PlayerInfo;

extern PlayerInfo* New_PlayerInfo();
extern void Free_PlayerInfo(PlayerInfo*);

/* ========== player_is_operated ========== */

typedef struct player_is_operated {
	int64_t Pid;
	int64_t OperateValue;
} PlayerIsOperated;

extern PlayerIsOperated* New_PlayerIsOperated();
extern void Free_PlayerIsOperated(PlayerIsOperated*);

/* ========== player_item ========== */

typedef struct player_item {
	int64_t Id;
	int64_t Pid;
	int16_t ItemId;
	int16_t Num;
	int8_t IsDressed;
	int8_t BuyBackState;
	int64_t AppendixId;
	int16_t RefineLevelBak;
	int16_t RefineFailTimes;
	int16_t RefineLevel;
	int32_t Price;
	struct player_item* next;
} PlayerItem;

extern PlayerItem* New_PlayerItem();
extern void Free_PlayerItem(PlayerItem*);

/* ========== player_item_appendix ========== */

typedef struct player_item_appendix {
	int64_t Id;
	int64_t Pid;
	int32_t Health;
	int32_t Cultivation;
	int32_t Speed;
	int32_t Attack;
	int32_t Defence;
	int32_t DodgeLevel;
	int32_t HitLevel;
	int32_t BlockLevel;
	int32_t CriticalLevel;
	int32_t TenacityLevel;
	int32_t DestroyLevel;
	int8_t RecastAttr;
	struct player_item_appendix* next;
} PlayerItemAppendix;

extern PlayerItemAppendix* New_PlayerItemAppendix();
extern void Free_PlayerItemAppendix(PlayerItemAppendix*);

/* ========== player_item_buyback ========== */

typedef struct player_item_buyback {
	int64_t Pid;
	int64_t BackId1;
	int64_t BackId2;
	int64_t BackId3;
	int64_t BackId4;
	int64_t BackId5;
	int64_t BackId6;
} PlayerItemBuyback;

extern PlayerItemBuyback* New_PlayerItemBuyback();
extern void Free_PlayerItemBuyback(PlayerItemBuyback*);

/* ========== player_item_recast_state ========== */

typedef struct player_item_recast_state {
	int64_t Pid;
	int64_t PlayerItemId;
	int8_t SelectedAttr;
	int8_t Attr1Type;
	int32_t Attr1Value;
	int8_t Attr2Type;
	int32_t Attr2Value;
	int8_t Attr3Type;
	int32_t Attr3Value;
} PlayerItemRecastState;

extern PlayerItemRecastState* New_PlayerItemRecastState();
extern void Free_PlayerItemRecastState(PlayerItemRecastState*);

/* ========== player_level_award_info ========== */

typedef struct player_level_award_info {
	int64_t Pid;
	char* Awarded;
	int Awarded_len;
} PlayerLevelAwardInfo;

extern PlayerLevelAwardInfo* New_PlayerLevelAwardInfo();
extern void Free_PlayerLevelAwardInfo(PlayerLevelAwardInfo*);

/* ========== player_login_award_record ========== */

typedef struct player_login_award_record {
	int64_t Pid;
	int32_t ActiveDays;
	int32_t Record;
	int64_t UpdateTimestamp;
} PlayerLoginAwardRecord;

extern PlayerLoginAwardRecord* New_PlayerLoginAwardRecord();
extern void Free_PlayerLoginAwardRecord(PlayerLoginAwardRecord*);

/* ========== player_mail ========== */

typedef struct player_mail {
	int64_t Id;
	int64_t Pid;
	int32_t MailId;
	int8_t State;
	int64_t SendTime;
	char* Parameters;
	int Parameters_len;
	int8_t HaveAttachment;
	char* Title;
	int Title_len;
	char* Content;
	int Content_len;
	int64_t ExpireTime;
	int8_t Priority;
	struct player_mail* next;
} PlayerMail;

extern PlayerMail* New_PlayerMail();
extern void Free_PlayerMail(PlayerMail*);

/* ========== player_mail_attachment ========== */

typedef struct player_mail_attachment {
	int64_t Id;
	int64_t Pid;
	int64_t PlayerMailId;
	int8_t AttachmentType;
	int16_t ItemId;
	int64_t ItemNum;
	struct player_mail_attachment* next;
} PlayerMailAttachment;

extern PlayerMailAttachment* New_PlayerMailAttachment();
extern void Free_PlayerMailAttachment(PlayerMailAttachment*);

/* ========== player_mail_attachment_lg ========== */

typedef struct player_mail_attachment_lg {
	int64_t Id;
	int64_t Pid;
	int64_t PlayerMailId;
	int8_t AttachmentType;
	int16_t ItemId;
	int64_t ItemNum;
	int64_t TakeTimestamp;
	struct player_mail_attachment_lg* next;
} PlayerMailAttachmentLg;

extern PlayerMailAttachmentLg* New_PlayerMailAttachmentLg();
extern void Free_PlayerMailAttachmentLg(PlayerMailAttachmentLg*);

/* ========== player_mail_lg ========== */

typedef struct player_mail_lg {
	int64_t Id;
	int64_t Pmid;
	int64_t Pid;
	int32_t MailId;
	int8_t State;
	int64_t SendTime;
	char* Parameters;
	int Parameters_len;
	int8_t HaveAttachment;
	char* Title;
	int Title_len;
	char* Content;
	int Content_len;
	struct player_mail_lg* next;
} PlayerMailLg;

extern PlayerMailLg* New_PlayerMailLg();
extern void Free_PlayerMailLg(PlayerMailLg*);

/* ========== player_meditation_state ========== */

typedef struct player_meditation_state {
	int64_t Pid;
	int32_t AccumulateTime;
	int64_t StartTimestamp;
} PlayerMeditationState;

extern PlayerMeditationState* New_PlayerMeditationState();
extern void Free_PlayerMeditationState(PlayerMeditationState*);

/* ========== player_mission ========== */

typedef struct player_mission {
	int64_t Pid;
	int32_t Key;
	int8_t MaxOrder;
} PlayerMission;

extern PlayerMission* New_PlayerMission();
extern void Free_PlayerMission(PlayerMission*);

/* ========== player_mission_level ========== */

typedef struct player_mission_level {
	int64_t Pid;
	int32_t Lock;
	int32_t MaxLock;
	int32_t AwardLock;
} PlayerMissionLevel;

extern PlayerMissionLevel* New_PlayerMissionLevel();
extern void Free_PlayerMissionLevel(PlayerMissionLevel*);

/* ========== player_mission_level_record ========== */

typedef struct player_mission_level_record {
	int64_t Id;
	int64_t Pid;
	int16_t MissionId;
	int32_t MissionLevelId;
	int64_t OpenTime;
	int32_t Score;
	int8_t Round;
	int8_t DailyNum;
	int64_t LastEnterTime;
	int16_t EmptyShadowBits;
	struct player_mission_level_record* next;
} PlayerMissionLevelRecord;

extern PlayerMissionLevelRecord* New_PlayerMissionLevelRecord();
extern void Free_PlayerMissionLevelRecord(PlayerMissionLevelRecord*);

/* ========== player_mission_level_state_bin ========== */

typedef struct player_mission_level_state_bin {
	int64_t Pid;
	char* Bin;
	int Bin_len;
} PlayerMissionLevelStateBin;

extern PlayerMissionLevelStateBin* New_PlayerMissionLevelStateBin();
extern void Free_PlayerMissionLevelStateBin(PlayerMissionLevelStateBin*);

/* ========== player_mission_record ========== */

typedef struct player_mission_record {
	int64_t Id;
	int64_t Pid;
	int16_t TownId;
	int16_t MissionId;
	int64_t OpenTime;
	struct player_mission_record* next;
} PlayerMissionRecord;

extern PlayerMissionRecord* New_PlayerMissionRecord();
extern void Free_PlayerMissionRecord(PlayerMissionRecord*);

/* ========== player_mission_star_award ========== */

typedef struct player_mission_star_award {
	int64_t Id;
	int64_t Pid;
	int16_t TownId;
	int8_t BoxType;
	struct player_mission_star_award* next;
} PlayerMissionStarAward;

extern PlayerMissionStarAward* New_PlayerMissionStarAward();
extern void Free_PlayerMissionStarAward(PlayerMissionStarAward*);

/* ========== player_money_tree ========== */

typedef struct player_money_tree {
	int64_t Pid;
	int32_t Total;
	int8_t WavedTimesTotal;
	int8_t WavedTimes;
	int64_t LastWavedTime;
} PlayerMoneyTree;

extern PlayerMoneyTree* New_PlayerMoneyTree();
extern void Free_PlayerMoneyTree(PlayerMoneyTree*);

/* ========== player_month_card_award_record ========== */

typedef struct player_month_card_award_record {
	int64_t Pid;
	int64_t LastUpdate;
} PlayerMonthCardAwardRecord;

extern PlayerMonthCardAwardRecord* New_PlayerMonthCardAwardRecord();
extern void Free_PlayerMonthCardAwardRecord(PlayerMonthCardAwardRecord*);

/* ========== player_month_card_info ========== */

typedef struct player_month_card_info {
	int64_t Pid;
	int64_t Starttime;
	int64_t Endtime;
	int64_t Buytimes;
	int64_t Presenttotal;
} PlayerMonthCardInfo;

extern PlayerMonthCardInfo* New_PlayerMonthCardInfo();
extern void Free_PlayerMonthCardInfo(PlayerMonthCardInfo*);

/* ========== player_multi_level_info ========== */

typedef struct player_multi_level_info {
	int64_t Pid;
	int8_t BuddyRoleId;
	int8_t BuddyRow;
	int8_t TacticalGrid;
	int8_t DailyNum;
	int64_t BattleTime;
	int32_t Lock;
} PlayerMultiLevelInfo;

extern PlayerMultiLevelInfo* New_PlayerMultiLevelInfo();
extern void Free_PlayerMultiLevelInfo(PlayerMultiLevelInfo*);

/* ========== player_new_year_consume_record ========== */

typedef struct player_new_year_consume_record {
	int64_t Pid;
	char* ConsumeRecord;
	int ConsumeRecord_len;
} PlayerNewYearConsumeRecord;

extern PlayerNewYearConsumeRecord* New_PlayerNewYearConsumeRecord();
extern void Free_PlayerNewYearConsumeRecord(PlayerNewYearConsumeRecord*);

/* ========== player_npc_talk_record ========== */

typedef struct player_npc_talk_record {
	int64_t Id;
	int64_t Pid;
	int32_t NpcId;
	int16_t TownId;
	int16_t QuestId;
	struct player_npc_talk_record* next;
} PlayerNpcTalkRecord;

extern PlayerNpcTalkRecord* New_PlayerNpcTalkRecord();
extern void Free_PlayerNpcTalkRecord(PlayerNpcTalkRecord*);

/* ========== player_opened_town_treasure ========== */

typedef struct player_opened_town_treasure {
	int64_t Id;
	int64_t Pid;
	int16_t TownId;
	struct player_opened_town_treasure* next;
} PlayerOpenedTownTreasure;

extern PlayerOpenedTownTreasure* New_PlayerOpenedTownTreasure();
extern void Free_PlayerOpenedTownTreasure(PlayerOpenedTownTreasure*);

/* ========== player_physical ========== */

typedef struct player_physical {
	int64_t Pid;
	int16_t Value;
	int64_t UpdateTime;
	int64_t BuyCount;
	int64_t BuyUpdateTime;
	int16_t DailyCount;
} PlayerPhysical;

extern PlayerPhysical* New_PlayerPhysical();
extern void Free_PlayerPhysical(PlayerPhysical*);

/* ========== player_purchase_record ========== */

typedef struct player_purchase_record {
	int64_t Id;
	int64_t Pid;
	int16_t ItemId;
	int16_t Num;
	int64_t Timestamp;
	struct player_purchase_record* next;
} PlayerPurchaseRecord;

extern PlayerPurchaseRecord* New_PlayerPurchaseRecord();
extern void Free_PlayerPurchaseRecord(PlayerPurchaseRecord*);

/* ========== player_push_notify_switch ========== */

typedef struct player_push_notify_switch {
	int64_t Pid;
	int64_t Options;
} PlayerPushNotifySwitch;

extern PlayerPushNotifySwitch* New_PlayerPushNotifySwitch();
extern void Free_PlayerPushNotifySwitch(PlayerPushNotifySwitch*);

/* ========== player_pve_state ========== */

typedef struct player_pve_state {
	int64_t Pid;
	int16_t MaxPassedFloor;
	int16_t MaxAwardedFloor;
	int16_t UnpassedFloorEnemyNum;
	int64_t EnterTime;
	int8_t DailyNum;
} PlayerPveState;

extern PlayerPveState* New_PlayerPveState();
extern void Free_PlayerPveState(PlayerPveState*);

/* ========== player_qq_vip_gift ========== */

typedef struct player_qq_vip_gift {
	int64_t Pid;
	int16_t Qqvip;
	int16_t Surper;
} PlayerQqVipGift;

extern PlayerQqVipGift* New_PlayerQqVipGift();
extern void Free_PlayerQqVipGift(PlayerQqVipGift*);

/* ========== player_quest ========== */

typedef struct player_quest {
	int64_t Pid;
	int16_t QuestId;
	int8_t State;
} PlayerQuest;

extern PlayerQuest* New_PlayerQuest();
extern void Free_PlayerQuest(PlayerQuest*);

/* ========== player_rainbow_level ========== */

typedef struct player_rainbow_level {
	int64_t Pid;
	int64_t ResetTimestamp;
	int32_t ResetNum;
	int16_t Segment;
	int8_t Order;
	int8_t Status;
	int16_t BestSegment;
	int8_t BestOrder;
	int64_t BestRecordTimestamp;
	int16_t MaxOpenSegment;
	int16_t MaxPassSegment;
	int8_t AutoFightNum;
	int16_t BuyTimes;
	int64_t AutoFightTime;
	int64_t BuyTimestamp;
} PlayerRainbowLevel;

extern PlayerRainbowLevel* New_PlayerRainbowLevel();
extern void Free_PlayerRainbowLevel(PlayerRainbowLevel*);

/* ========== player_rainbow_level_state_bin ========== */

typedef struct player_rainbow_level_state_bin {
	int64_t Pid;
	char* Bin;
	int Bin_len;
} PlayerRainbowLevelStateBin;

extern PlayerRainbowLevelStateBin* New_PlayerRainbowLevelStateBin();
extern void Free_PlayerRainbowLevelStateBin(PlayerRainbowLevelStateBin*);

/* ========== player_role ========== */

typedef struct player_role {
	int64_t Id;
	int64_t Pid;
	int8_t RoleId;
	int16_t Level;
	int64_t Exp;
	int32_t FriendshipLevel;
	int16_t Status;
	int16_t CoopId;
	int64_t Timestamp;
	int32_t FightNum;
	struct player_role* next;
} PlayerRole;

extern PlayerRole* New_PlayerRole();
extern void Free_PlayerRole(PlayerRole*);

/* ========== player_sealedbook ========== */

typedef struct player_sealedbook {
	int64_t Pid;
	char* SealedRecord;
	int SealedRecord_len;
} PlayerSealedbook;

extern PlayerSealedbook* New_PlayerSealedbook();
extern void Free_PlayerSealedbook(PlayerSealedbook*);

/* ========== player_send_heart_record ========== */

typedef struct player_send_heart_record {
	int64_t Id;
	int64_t Pid;
	int64_t FriendPid;
	int64_t SendHeartTime;
	struct player_send_heart_record* next;
} PlayerSendHeartRecord;

extern PlayerSendHeartRecord* New_PlayerSendHeartRecord();
extern void Free_PlayerSendHeartRecord(PlayerSendHeartRecord*);

/* ========== player_skill ========== */

typedef struct player_skill {
	int64_t Id;
	int64_t Pid;
	int8_t RoleId;
	int16_t SkillId;
	int32_t SkillTrnlv;
	struct player_skill* next;
} PlayerSkill;

extern PlayerSkill* New_PlayerSkill();
extern void Free_PlayerSkill(PlayerSkill*);

/* ========== player_state ========== */

typedef struct player_state {
	int64_t Pid;
	int64_t BanStartTime;
	int64_t BanEndTime;
} PlayerState;

extern PlayerState* New_PlayerState();
extern void Free_PlayerState(PlayerState*);

/* ========== player_sword_soul ========== */

typedef struct player_sword_soul {
	int64_t Id;
	int64_t Pid;
	int8_t Pos;
	int16_t SwordSoulId;
	int32_t Exp;
	int8_t Level;
	struct player_sword_soul* next;
} PlayerSwordSoul;

extern PlayerSwordSoul* New_PlayerSwordSoul();
extern void Free_PlayerSwordSoul(PlayerSwordSoul*);

/* ========== player_sword_soul_equipment ========== */

typedef struct player_sword_soul_equipment {
	int64_t Id;
	int64_t Pid;
	int8_t RoleId;
	int8_t IsEquippedXuanyuan;
	int64_t TypeBits;
	int64_t Pos1;
	int64_t Pos2;
	int64_t Pos3;
	int64_t Pos4;
	int64_t Pos5;
	int64_t Pos6;
	int64_t Pos7;
	int64_t Pos8;
	int64_t Pos9;
	struct player_sword_soul_equipment* next;
} PlayerSwordSoulEquipment;

extern PlayerSwordSoulEquipment* New_PlayerSwordSoulEquipment();
extern void Free_PlayerSwordSoulEquipment(PlayerSwordSoulEquipment*);

/* ========== player_sword_soul_state ========== */

typedef struct player_sword_soul_state {
	int64_t Pid;
	int8_t BoxState;
	int16_t Num;
	int64_t UpdateTime;
	int64_t IngotNum;
	int8_t SupersoulAdditionalPossible;
	int16_t IngotYelloOne;
	int64_t LastIngotDrawTime;
	int32_t SwordSoulNum;
} PlayerSwordSoulState;

extern PlayerSwordSoulState* New_PlayerSwordSoulState();
extern void Free_PlayerSwordSoulState(PlayerSwordSoulState*);

/* ========== player_taoyuan_bless ========== */

typedef struct player_taoyuan_bless {
	int64_t Pid;
	int32_t DailyBlessTimes;
	int64_t LastBlessTime;
	int32_t DailyBeBlessTimes;
	int64_t LastBeBlessTime;
	int64_t BlessPid1;
	int64_t BlessPid2;
	int64_t BlessPid3;
	int64_t BlessPid4;
	int64_t BlessPid5;
} PlayerTaoyuanBless;

extern PlayerTaoyuanBless* New_PlayerTaoyuanBless();
extern void Free_PlayerTaoyuanBless(PlayerTaoyuanBless*);

/* ========== player_taoyuan_fileds ========== */

typedef struct player_taoyuan_fileds {
	int64_t Id;
	int64_t Pid;
	int16_t FiledId;
	int16_t FiledStatus;
	int16_t PlantId;
	int64_t GrowTime;
	struct player_taoyuan_fileds* next;
} PlayerTaoyuanFileds;

extern PlayerTaoyuanFileds* New_PlayerTaoyuanFileds();
extern void Free_PlayerTaoyuanFileds(PlayerTaoyuanFileds*);

/* ========== player_taoyuan_heart ========== */

typedef struct player_taoyuan_heart {
	int64_t Pid;
	int16_t Level;
	int64_t Exp;
} PlayerTaoyuanHeart;

extern PlayerTaoyuanHeart* New_PlayerTaoyuanHeart();
extern void Free_PlayerTaoyuanHeart(PlayerTaoyuanHeart*);

/* ========== player_taoyuan_item ========== */

typedef struct player_taoyuan_item {
	int64_t Id;
	int64_t Pid;
	int16_t ItemId;
	int16_t Num;
	struct player_taoyuan_item* next;
} PlayerTaoyuanItem;

extern PlayerTaoyuanItem* New_PlayerTaoyuanItem();
extern void Free_PlayerTaoyuanItem(PlayerTaoyuanItem*);

/* ========== player_taoyuan_message ========== */

typedef struct player_taoyuan_message {
	int64_t Id;
	int64_t Pid;
	char* Nick;
	int Nick_len;
	int32_t Exp;
	struct player_taoyuan_message* next;
} PlayerTaoyuanMessage;

extern PlayerTaoyuanMessage* New_PlayerTaoyuanMessage();
extern void Free_PlayerTaoyuanMessage(PlayerTaoyuanMessage*);

/* ========== player_taoyuan_product ========== */

typedef struct player_taoyuan_product {
	int64_t Pid;
	int16_t SkillId1;
	int16_t SkillId2;
	int16_t SameTimeWineQueue;
	int64_t MakeWineTimes;
	int16_t SameTimeFoodQueue;
	int64_t MakeFoodTimes;
	int16_t FoodQueue1;
	int64_t FoodQueue1StartTimestamp;
	int64_t FoodQueue1EndTimestamp;
	int16_t FoodQueue2;
	int64_t FoodQueue2StartTimestamp;
	int64_t FoodQueue2EndTimestamp;
	int16_t FoodQueue3;
	int64_t FoodQueue3StartTimestamp;
	int64_t FoodQueue3EndTimestamp;
	int16_t FoodQueue4;
	int64_t FoodQueue4StartTimestamp;
	int64_t FoodQueue4EndTimestamp;
	int16_t WineQueue1;
	int64_t WineQueue1StartTimestamp;
	int64_t WineQueue1EndTimestamp;
	int16_t WineQueue2;
	int64_t WineQueue2StartTimestamp;
	int64_t WineQueue2EndTimestamp;
	int16_t WineQueue3;
	int64_t WineQueue3StartTimestamp;
	int64_t WineQueue3EndTimestamp;
	int16_t WineQueue4;
	int64_t WineQueue4StartTimestamp;
	int64_t WineQueue4EndTimestamp;
	int8_t IsOpenWine;
	int8_t IsOpenFood;
} PlayerTaoyuanProduct;

extern PlayerTaoyuanProduct* New_PlayerTaoyuanProduct();
extern void Free_PlayerTaoyuanProduct(PlayerTaoyuanProduct*);

/* ========== player_taoyuan_purchase_record ========== */

typedef struct player_taoyuan_purchase_record {
	int64_t Id;
	int64_t Pid;
	int16_t ItemId;
	int16_t Num;
	int64_t Timestamp;
	struct player_taoyuan_purchase_record* next;
} PlayerTaoyuanPurchaseRecord;

extern PlayerTaoyuanPurchaseRecord* New_PlayerTaoyuanPurchaseRecord();
extern void Free_PlayerTaoyuanPurchaseRecord(PlayerTaoyuanPurchaseRecord*);

/* ========== player_taoyuan_quest ========== */

typedef struct player_taoyuan_quest {
	int64_t Pid;
	int64_t LastQuestFlashTime;
	int16_t Quest1ItemId;
	int16_t Quest1ItemNum;
	int16_t Quest1TotalExp;
	int64_t Quest1TotalCoins;
	int64_t Quest1FinishTime;
	int16_t Quest2ItemId;
	int16_t Quest2ItemNum;
	int16_t Quest2TotalExp;
	int64_t Quest2TotalCoins;
	int64_t Quest2FinishTime;
	int16_t Quest3ItemId;
	int16_t Quest3ItemNum;
	int16_t Quest3TotalExp;
	int64_t Quest3TotalCoins;
	int64_t Quest3FinishTime;
	int16_t Quest4ItemId;
	int16_t Quest4ItemNum;
	int16_t Quest4TotalExp;
	int64_t Quest4TotalCoins;
	int64_t Quest4FinishTime;
	int16_t Quest5ItemId;
	int16_t Quest5ItemNum;
	int16_t Quest5TotalExp;
	int64_t Quest5TotalCoins;
	int64_t Quest5FinishTime;
} PlayerTaoyuanQuest;

extern PlayerTaoyuanQuest* New_PlayerTaoyuanQuest();
extern void Free_PlayerTaoyuanQuest(PlayerTaoyuanQuest*);

/* ========== player_tb_xxd_roleinfo ========== */

typedef struct player_tb_xxd_roleinfo {
	int64_t Pid;
	char* Gameappid;
	int Gameappid_len;
	char* Openid;
	int Openid_len;
	int64_t Regtime;
	int16_t Level;
	int16_t IFriends;
	int64_t Moneyios;
	int64_t Moneyandroid;
	int64_t Diamondios;
	int64_t Diamondandroid;
} PlayerTbXxdRoleinfo;

extern PlayerTbXxdRoleinfo* New_PlayerTbXxdRoleinfo();
extern void Free_PlayerTbXxdRoleinfo(PlayerTbXxdRoleinfo*);

/* ========== player_team_info ========== */

typedef struct player_team_info {
	int64_t Pid;
	int32_t Relationship;
	int16_t HealthLv;
	int16_t AttackLv;
	int16_t DefenceLv;
} PlayerTeamInfo;

extern PlayerTeamInfo* New_PlayerTeamInfo();
extern void Free_PlayerTeamInfo(PlayerTeamInfo*);

/* ========== player_totem ========== */

typedef struct player_totem {
	int64_t Id;
	int64_t Pid;
	int16_t TotemId;
	int8_t Level;
	int16_t SkillId;
	struct player_totem* next;
} PlayerTotem;

extern PlayerTotem* New_PlayerTotem();
extern void Free_PlayerTotem(PlayerTotem*);

/* ========== player_totem_info ========== */

typedef struct player_totem_info {
	int64_t Pid;
	int16_t IngotCallDailyNum;
	int64_t IngotCallTimestamp;
	int32_t RockRuneNum;
	int32_t JadeRuneNum;
	int64_t Pos1;
	int64_t Pos2;
	int64_t Pos3;
	int64_t Pos4;
	int64_t Pos5;
	int16_t IngotDrawTimes;
} PlayerTotemInfo;

extern PlayerTotemInfo* New_PlayerTotemInfo();
extern void Free_PlayerTotemInfo(PlayerTotemInfo*);

/* ========== player_town ========== */

typedef struct player_town {
	int64_t Pid;
	int16_t TownId;
	int32_t Lock;
	int16_t AtPosX;
	int16_t AtPosY;
} PlayerTown;

extern PlayerTown* New_PlayerTown();
extern void Free_PlayerTown(PlayerTown*);

/* ========== player_trader_refresh_state ========== */

typedef struct player_trader_refresh_state {
	int64_t Id;
	int64_t Pid;
	int64_t LastUpdateTime;
	int64_t AutoUpdateTime;
	int16_t TraderId;
	int16_t RefreshNum;
	struct player_trader_refresh_state* next;
} PlayerTraderRefreshState;

extern PlayerTraderRefreshState* New_PlayerTraderRefreshState();
extern void Free_PlayerTraderRefreshState(PlayerTraderRefreshState*);

/* ========== player_trader_store_state ========== */

typedef struct player_trader_store_state {
	int64_t Id;
	int64_t Pid;
	int16_t TraderId;
	int32_t GridId;
	int16_t ItemId;
	int16_t Num;
	int64_t Cost;
	int8_t Stock;
	int8_t GoodsType;
	struct player_trader_store_state* next;
} PlayerTraderStoreState;

extern PlayerTraderStoreState* New_PlayerTraderStoreState();
extern void Free_PlayerTraderStoreState(PlayerTraderStoreState*);

/* ========== player_use_skill ========== */

typedef struct player_use_skill {
	int64_t Id;
	int64_t Pid;
	int8_t RoleId;
	int16_t SkillId1;
	int16_t SkillId4;
	int16_t SkillId2;
	int16_t SkillId3;
	struct player_use_skill* next;
} PlayerUseSkill;

extern PlayerUseSkill* New_PlayerUseSkill();
extern void Free_PlayerUseSkill(PlayerUseSkill*);

/* ========== player_vip ========== */

typedef struct player_vip {
	int64_t Pid;
	int64_t Ingot;
	int16_t Level;
	char* CardId;
	int CardId_len;
	int64_t PresentExp;
} PlayerVip;

extern PlayerVip* New_PlayerVip();
extern void Free_PlayerVip(PlayerVip*);

/* ================================== */

struct GlobalTables {
	GlobalAnnouncement* GlobalAnnouncement;
	GlobalArenaRank* GlobalArenaRank;
	GlobalClique* GlobalClique;
	GlobalCliqueBoat* GlobalCliqueBoat;
	GlobalDespairBoss* GlobalDespairBoss;
	GlobalDespairBossHistory* GlobalDespairBossHistory;
	GlobalDespairLand* GlobalDespairLand;
	GlobalDespairLandBattleRecord* GlobalDespairLandBattleRecord;
	GlobalDespairLandCamp* GlobalDespairLandCamp;
	GlobalGiftCardRecord* GlobalGiftCardRecord;
	GlobalGroupBuyStatus* GlobalGroupBuyStatus;
	GlobalMail* GlobalMail;
	GlobalMailAttachments* GlobalMailAttachments;
	GlobalTbXxdOnlinecnt* GlobalTbXxdOnlinecnt;
};

struct PlayerTables {
	int64_t Pid;
	Player* Player;
	PlayerActivity* PlayerActivity;
	PlayerAdditionQuest* PlayerAdditionQuest;
	PlayerAnySdkOrder* PlayerAnySdkOrder;
	PlayerArena* PlayerArena;
	PlayerArenaRank* PlayerArenaRank;
	PlayerArenaRecord* PlayerArenaRecord;
	PlayerAwakenGraphic* PlayerAwakenGraphic;
	PlayerBattlePet* PlayerBattlePet;
	PlayerBattlePetGrid* PlayerBattlePetGrid;
	PlayerBattlePetState* PlayerBattlePetState;
	PlayerChestState* PlayerChestState;
	PlayerCliqueKongfuAttrib* PlayerCliqueKongfuAttrib;
	PlayerCoins* PlayerCoins;
	PlayerCornucopia* PlayerCornucopia;
	PlayerDailyQuest* PlayerDailyQuest;
	PlayerDailyQuestStarAwardInfo* PlayerDailyQuestStarAwardInfo;
	PlayerDailySignInRecord* PlayerDailySignInRecord;
	PlayerDailySignInState* PlayerDailySignInState;
	PlayerDespairLandCampState* PlayerDespairLandCampState;
	PlayerDespairLandCampStateHistory* PlayerDespairLandCampStateHistory;
	PlayerDespairLandLevelRecord* PlayerDespairLandLevelRecord;
	PlayerDespairLandState* PlayerDespairLandState;
	PlayerDrivingSwordEvent* PlayerDrivingSwordEvent;
	PlayerDrivingSwordEventExploring* PlayerDrivingSwordEventExploring;
	PlayerDrivingSwordEventTreasure* PlayerDrivingSwordEventTreasure;
	PlayerDrivingSwordEventVisiting* PlayerDrivingSwordEventVisiting;
	PlayerDrivingSwordEventmask* PlayerDrivingSwordEventmask;
	PlayerDrivingSwordInfo* PlayerDrivingSwordInfo;
	PlayerDrivingSwordMap* PlayerDrivingSwordMap;
	PlayerEquipment* PlayerEquipment;
	PlayerEventAwardRecord* PlayerEventAwardRecord;
	PlayerEventDailyOnline* PlayerEventDailyOnline;
	PlayerEventVnDailyRecharge* PlayerEventVnDailyRecharge;
	PlayerEventVnDragonBallRecord* PlayerEventVnDragonBallRecord;
	PlayerEventsFightPower* PlayerEventsFightPower;
	PlayerEventsIngotRecord* PlayerEventsIngotRecord;
	PlayerExtendLevel* PlayerExtendLevel;
	PlayerExtendQuest* PlayerExtendQuest;
	PlayerFame* PlayerFame;
	PlayerFashion* PlayerFashion;
	PlayerFashionState* PlayerFashionState;
	PlayerFateBoxState* PlayerFateBoxState;
	PlayerFightNum* PlayerFightNum;
	PlayerFirstCharge* PlayerFirstCharge;
	PlayerFormation* PlayerFormation;
	PlayerFuncKey* PlayerFuncKey;
	PlayerGhost* PlayerGhost;
	PlayerGhostEquipment* PlayerGhostEquipment;
	PlayerGhostState* PlayerGhostState;
	PlayerGlobalChatState* PlayerGlobalChatState;
	PlayerGlobalCliqueBuilding* PlayerGlobalCliqueBuilding;
	PlayerGlobalCliqueBuildingQuest* PlayerGlobalCliqueBuildingQuest;
	PlayerGlobalCliqueDailyQuest* PlayerGlobalCliqueDailyQuest;
	PlayerGlobalCliqueEscort* PlayerGlobalCliqueEscort;
	PlayerGlobalCliqueEscortMessage* PlayerGlobalCliqueEscortMessage;
	PlayerGlobalCliqueHijack* PlayerGlobalCliqueHijack;
	PlayerGlobalCliqueInfo* PlayerGlobalCliqueInfo;
	PlayerGlobalCliqueKongfu* PlayerGlobalCliqueKongfu;
	PlayerGlobalFriend* PlayerGlobalFriend;
	PlayerGlobalFriendChat* PlayerGlobalFriendChat;
	PlayerGlobalFriendState* PlayerGlobalFriendState;
	PlayerGlobalGiftCardRecord* PlayerGlobalGiftCardRecord;
	PlayerGlobalMailState* PlayerGlobalMailState;
	PlayerGlobalWorldChatState* PlayerGlobalWorldChatState;
	PlayerHardLevel* PlayerHardLevel;
	PlayerHardLevelRecord* PlayerHardLevelRecord;
	PlayerHeart* PlayerHeart;
	PlayerHeartDraw* PlayerHeartDraw;
	PlayerHeartDrawCardRecord* PlayerHeartDrawCardRecord;
	PlayerHeartDrawWheelRecord* PlayerHeartDrawWheelRecord;
	PlayerInfo* PlayerInfo;
	PlayerIsOperated* PlayerIsOperated;
	PlayerItem* PlayerItem;
	PlayerItemAppendix* PlayerItemAppendix;
	PlayerItemBuyback* PlayerItemBuyback;
	PlayerItemRecastState* PlayerItemRecastState;
	PlayerLevelAwardInfo* PlayerLevelAwardInfo;
	PlayerLoginAwardRecord* PlayerLoginAwardRecord;
	PlayerMail* PlayerMail;
	PlayerMailAttachment* PlayerMailAttachment;
	PlayerMailAttachmentLg* PlayerMailAttachmentLg;
	PlayerMailLg* PlayerMailLg;
	PlayerMeditationState* PlayerMeditationState;
	PlayerMission* PlayerMission;
	PlayerMissionLevel* PlayerMissionLevel;
	PlayerMissionLevelRecord* PlayerMissionLevelRecord;
	PlayerMissionLevelStateBin* PlayerMissionLevelStateBin;
	PlayerMissionRecord* PlayerMissionRecord;
	PlayerMissionStarAward* PlayerMissionStarAward;
	PlayerMoneyTree* PlayerMoneyTree;
	PlayerMonthCardAwardRecord* PlayerMonthCardAwardRecord;
	PlayerMonthCardInfo* PlayerMonthCardInfo;
	PlayerMultiLevelInfo* PlayerMultiLevelInfo;
	PlayerNewYearConsumeRecord* PlayerNewYearConsumeRecord;
	PlayerNpcTalkRecord* PlayerNpcTalkRecord;
	PlayerOpenedTownTreasure* PlayerOpenedTownTreasure;
	PlayerPhysical* PlayerPhysical;
	PlayerPurchaseRecord* PlayerPurchaseRecord;
	PlayerPushNotifySwitch* PlayerPushNotifySwitch;
	PlayerPveState* PlayerPveState;
	PlayerQqVipGift* PlayerQqVipGift;
	PlayerQuest* PlayerQuest;
	PlayerRainbowLevel* PlayerRainbowLevel;
	PlayerRainbowLevelStateBin* PlayerRainbowLevelStateBin;
	PlayerRole* PlayerRole;
	PlayerSealedbook* PlayerSealedbook;
	PlayerSendHeartRecord* PlayerSendHeartRecord;
	PlayerSkill* PlayerSkill;
	PlayerState* PlayerState;
	PlayerSwordSoul* PlayerSwordSoul;
	PlayerSwordSoulEquipment* PlayerSwordSoulEquipment;
	PlayerSwordSoulState* PlayerSwordSoulState;
	PlayerTaoyuanBless* PlayerTaoyuanBless;
	PlayerTaoyuanFileds* PlayerTaoyuanFileds;
	PlayerTaoyuanHeart* PlayerTaoyuanHeart;
	PlayerTaoyuanItem* PlayerTaoyuanItem;
	PlayerTaoyuanMessage* PlayerTaoyuanMessage;
	PlayerTaoyuanProduct* PlayerTaoyuanProduct;
	PlayerTaoyuanPurchaseRecord* PlayerTaoyuanPurchaseRecord;
	PlayerTaoyuanQuest* PlayerTaoyuanQuest;
	PlayerTbXxdRoleinfo* PlayerTbXxdRoleinfo;
	PlayerTeamInfo* PlayerTeamInfo;
	PlayerTotem* PlayerTotem;
	PlayerTotemInfo* PlayerTotemInfo;
	PlayerTown* PlayerTown;
	PlayerTraderRefreshState* PlayerTraderRefreshState;
	PlayerTraderStoreState* PlayerTraderStoreState;
	PlayerUseSkill* PlayerUseSkill;
	PlayerVip* PlayerVip;
};

#endif
