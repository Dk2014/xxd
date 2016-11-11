package mdb

/*
#include "zd_cgo.h"
*/
import "C"
import "sync/atomic"
import "fmt"

type insertOP struct{ db *Database }
type deleteOP struct{ db *Database }
type updateOP struct{ db *Database }
type lookupOP struct{ db *Database }
type selectOP struct{ db *Database }
type allOP    struct{ db *Database }


type Database struct {
	playerId int64
	cId int32
	tables *C.PlayerTables
	Insert insertOP
	Delete deleteOP
	Update updateOP
	Lookup lookupOP
	Select selectOP
	All    allOP
	transController *transController
}

func NewDatabase() *Database {
	return newDatabase(newTransController(g_LockManager.Get(0), g_FileCommitter))
}

func newDatabase(transController *transController) *Database {
	db := new(Database)
	db.Insert.db = db
	db.Delete.db = db
	db.Update.db = db
	db.Lookup.db = db
	db.Select.db = db
	db.All.db    = db
	db.transController = transController
	return db
}

func (db *Database) Init() {
	db.tables = C.NewPlayerTables()
}

func (db *Database) Mount(playerId int64,cId int32) {
	db.mountTables2(playerId,cId)
	db.transController.lock = g_LockManager.Get(playerId)
}

func (db *Database) AgentExecute(pid int64, cb func(*Database)) {
	agentDB := newDatabase(db.transController)
	agentDB.mountTables(pid)
	cb(agentDB)
}

func (db *Database) Rollback() {
	        db.transController.Rollback()
}

func (db *Database) mountTables2(playerId int64,cId int32) {
	db.playerId = playerId
	db.cId = cId
	db.tables = GetPlayerTables(playerId)
	if db.tables == nil {
		panic("tables is nil")
	}
}

func (db *Database) mountTables(playerId int64) {
	db.playerId = playerId
	db.tables = GetPlayerTables(playerId)
	if db.tables == nil {
		panic("tables is nil")
	}
}

func (db *Database) PlayerId() int64 {
	return db.playerId
}

func (db *Database) CId() int32 {
	return db.cId
}

func (db *Database) Transaction(transId uint32, work func()) {
	if db != g_Database {
		g_GlobalLock.RLock()
		defer g_GlobalLock.RUnlock()
	}

	db.transController.Transaction(transId, work)
}

func (db *Database) AddTLog(log TransLog) {
	db.transController.AddTransLog(log)
}

func (db *Database) AddXdLog(log TransLog) {
	db.transController.AddTransLog(log)
}

func (db *Database) addTransLog(log TransLog) {
	db.transController.AddTransLog(log)
}


type rowIds struct {
	GlobalAnnouncement int64
	GlobalClique int64
	GlobalCliqueBoat int64
	GlobalDespairBossHistory int64
	GlobalDespairLand int64
	GlobalDespairLandBattleRecord int64
	GlobalGiftCardRecord int64
	GlobalGroupBuyStatus int64
	GlobalMail int64
	GlobalMailAttachments int64
	GlobalTbXxdOnlinecnt int64
	Player int64
	PlayerAdditionQuest int64
	PlayerAnySdkOrder int64
	PlayerArenaRecord int64
	PlayerAwakenGraphic int64
	PlayerBattlePet int64
	PlayerBattlePetGrid int64
	PlayerDailyQuest int64
	PlayerDailySignInRecord int64
	PlayerDespairLandCampState int64
	PlayerDespairLandCampStateHistory int64
	PlayerDespairLandLevelRecord int64
	PlayerDrivingSwordEvent int64
	PlayerDrivingSwordEventExploring int64
	PlayerDrivingSwordEventTreasure int64
	PlayerDrivingSwordEventVisiting int64
	PlayerDrivingSwordEventmask int64
	PlayerDrivingSwordMap int64
	PlayerEquipment int64
	PlayerEventVnDailyRecharge int64
	PlayerExtendQuest int64
	PlayerFashion int64
	PlayerFirstCharge int64
	PlayerGhost int64
	PlayerGhostEquipment int64
	PlayerGlobalCliqueBuildingQuest int64
	PlayerGlobalCliqueDailyQuest int64
	PlayerGlobalCliqueEscortMessage int64
	PlayerGlobalCliqueKongfu int64
	PlayerGlobalFriend int64
	PlayerGlobalFriendChat int64
	PlayerGlobalGiftCardRecord int64
	PlayerHardLevelRecord int64
	PlayerHeartDraw int64
	PlayerHeartDrawCardRecord int64
	PlayerHeartDrawWheelRecord int64
	PlayerItem int64
	PlayerItemAppendix int64
	PlayerMail int64
	PlayerMailAttachment int64
	PlayerMailAttachmentLg int64
	PlayerMailLg int64
	PlayerMissionLevelRecord int64
	PlayerMissionRecord int64
	PlayerMissionStarAward int64
	PlayerNpcTalkRecord int64
	PlayerOpenedTownTreasure int64
	PlayerPurchaseRecord int64
	PlayerRole int64
	PlayerSendHeartRecord int64
	PlayerSkill int64
	PlayerSwordSoul int64
	PlayerSwordSoulEquipment int64
	PlayerTaoyuanFileds int64
	PlayerTaoyuanItem int64
	PlayerTaoyuanMessage int64
	PlayerTaoyuanPurchaseRecord int64
	PlayerTotem int64
	PlayerTraderRefreshState int64
	PlayerTraderStoreState int64
	PlayerUseSkill int64
}

var g_RowIds rowIds

func initRowIdsWithServerId(serverId int64) {
	g_RowIds.GlobalAnnouncement = serverId
	g_RowIds.GlobalClique = 1000
	g_RowIds.GlobalCliqueBoat = serverId
	g_RowIds.GlobalDespairBossHistory = serverId
	g_RowIds.GlobalDespairLand = serverId
	g_RowIds.GlobalDespairLandBattleRecord = serverId
	g_RowIds.GlobalGiftCardRecord = serverId
	g_RowIds.GlobalGroupBuyStatus = serverId
	g_RowIds.GlobalMail = serverId
	g_RowIds.GlobalMailAttachments = serverId
	g_RowIds.GlobalTbXxdOnlinecnt = serverId
	g_RowIds.Player = serverId
	g_RowIds.PlayerAdditionQuest = serverId
	g_RowIds.PlayerAnySdkOrder = serverId
	g_RowIds.PlayerArenaRecord = serverId
	g_RowIds.PlayerAwakenGraphic = serverId
	g_RowIds.PlayerBattlePet = serverId
	g_RowIds.PlayerBattlePetGrid = serverId
	g_RowIds.PlayerDailyQuest = serverId
	g_RowIds.PlayerDailySignInRecord = serverId
	g_RowIds.PlayerDespairLandCampState = serverId
	g_RowIds.PlayerDespairLandCampStateHistory = serverId
	g_RowIds.PlayerDespairLandLevelRecord = serverId
	g_RowIds.PlayerDrivingSwordEvent = serverId
	g_RowIds.PlayerDrivingSwordEventExploring = serverId
	g_RowIds.PlayerDrivingSwordEventTreasure = serverId
	g_RowIds.PlayerDrivingSwordEventVisiting = serverId
	g_RowIds.PlayerDrivingSwordEventmask = serverId
	g_RowIds.PlayerDrivingSwordMap = serverId
	g_RowIds.PlayerEquipment = serverId
	g_RowIds.PlayerEventVnDailyRecharge = serverId
	g_RowIds.PlayerExtendQuest = serverId
	g_RowIds.PlayerFashion = serverId
	g_RowIds.PlayerFirstCharge = serverId
	g_RowIds.PlayerGhost = serverId
	g_RowIds.PlayerGhostEquipment = serverId
	g_RowIds.PlayerGlobalCliqueBuildingQuest = serverId
	g_RowIds.PlayerGlobalCliqueDailyQuest = serverId
	g_RowIds.PlayerGlobalCliqueEscortMessage = serverId
	g_RowIds.PlayerGlobalCliqueKongfu = serverId
	g_RowIds.PlayerGlobalFriend = serverId
	g_RowIds.PlayerGlobalFriendChat = serverId
	g_RowIds.PlayerGlobalGiftCardRecord = serverId
	g_RowIds.PlayerHardLevelRecord = serverId
	g_RowIds.PlayerHeartDraw = serverId
	g_RowIds.PlayerHeartDrawCardRecord = serverId
	g_RowIds.PlayerHeartDrawWheelRecord = serverId
	g_RowIds.PlayerItem = serverId
	g_RowIds.PlayerItemAppendix = serverId
	g_RowIds.PlayerMail = serverId
	g_RowIds.PlayerMailAttachment = serverId
	g_RowIds.PlayerMailAttachmentLg = serverId
	g_RowIds.PlayerMailLg = serverId
	g_RowIds.PlayerMissionLevelRecord = serverId
	g_RowIds.PlayerMissionRecord = serverId
	g_RowIds.PlayerMissionStarAward = serverId
	g_RowIds.PlayerNpcTalkRecord = serverId
	g_RowIds.PlayerOpenedTownTreasure = serverId
	g_RowIds.PlayerPurchaseRecord = serverId
	g_RowIds.PlayerRole = serverId
	g_RowIds.PlayerSendHeartRecord = serverId
	g_RowIds.PlayerSkill = serverId
	g_RowIds.PlayerSwordSoul = serverId
	g_RowIds.PlayerSwordSoulEquipment = serverId
	g_RowIds.PlayerTaoyuanFileds = serverId
	g_RowIds.PlayerTaoyuanItem = serverId
	g_RowIds.PlayerTaoyuanMessage = serverId
	g_RowIds.PlayerTaoyuanPurchaseRecord = serverId
	g_RowIds.PlayerTotem = serverId
	g_RowIds.PlayerTraderRefreshState = serverId
	g_RowIds.PlayerTraderStoreState = serverId
	g_RowIds.PlayerUseSkill = serverId
}

/* ========== global_announcement ========== */

// 公告列表
type GlobalAnnouncement struct {
	Id int64 // 公告ID
	ExpireTime int64 // 创建时间戳
	TplId int32 // 公告模版ID
	Parameters string // 模版参数
	Content string // 公共内容，有则忽略模版
	SendTime int64 // 发送时间
	SpacingTime int64 // 间隔时间
}

func (this *GlobalAnnouncement) CObject() *C.GlobalAnnouncement {
	object := C.New_GlobalAnnouncement()
	object.Id = C.int64_t(this.Id)
	object.ExpireTime = C.int64_t(this.ExpireTime)
	object.TplId = C.int32_t(this.TplId)
	object.Parameters = C.CString(string(this.Parameters))
	object.Parameters_len = C.int(len(this.Parameters))
	object.Content = C.CString(string(this.Content))
	object.Content_len = C.int(len(this.Content))
	object.SendTime = C.int64_t(this.SendTime)
	object.SpacingTime = C.int64_t(this.SpacingTime)
	return object
}

func (this insertOP) GlobalAnnouncement(row *GlobalAnnouncement) {
	row.Id = atomic.AddInt64(&g_RowIds.GlobalAnnouncement, 1)
	key := C.int64_t(row.Id)
	var old *C.GlobalAnnouncement
	for crow := C.g_GlobalTables.GlobalAnnouncement; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'global_announcement'")
	}
	newRow := row.CObject()
	newRow.next = C.g_GlobalTables.GlobalAnnouncement
	C.g_GlobalTables.GlobalAnnouncement = newRow
	this.db.addTransLog(this.db.newGlobalAnnouncementInsertLog(newRow, row))
}

func (this deleteOP) GlobalAnnouncement(row *GlobalAnnouncement) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalAnnouncement
	for crow := C.g_GlobalTables.GlobalAnnouncement; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'global_announcement'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalAnnouncement = C.g_GlobalTables.GlobalAnnouncement.next
	}
	this.db.addTransLog(this.db.newGlobalAnnouncementDeleteLog(old, row))
}

func (this updateOP) GlobalAnnouncement(row *GlobalAnnouncement) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalAnnouncement
	for crow := C.g_GlobalTables.GlobalAnnouncement; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'global_announcement'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		C.g_GlobalTables.GlobalAnnouncement = newRow
	}
	tmpRow := GlobalAnnouncementRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newGlobalAnnouncementUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) GlobalAnnouncement(key int64) *GlobalAnnouncement {
	for crow := C.g_GlobalTables.GlobalAnnouncement; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := GlobalAnnouncementRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) GlobalAnnouncement(callback func(*GlobalAnnouncementRow)) {
	row := &GlobalAnnouncementRow{}
	for crow := C.g_GlobalTables.GlobalAnnouncement; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) GlobalAnnouncement() (rows []interface{}) {
	row := &GlobalAnnouncementRow{}
	for crow := C.g_GlobalTables.GlobalAnnouncement; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("GlobalAnnouncement", row.GoObject())
	}
	row.row = nil
	fmt.Println("GlobalAnnouncement", len(rows))
	return rows
}

/* ========== global_arena_rank ========== */

// 全局比武场数据
type GlobalArenaRank struct {
	Rank int32 // 排名
	Pid int64 // 玩家ID
}

func (this *GlobalArenaRank) CObject() *C.GlobalArenaRank {
	object := C.New_GlobalArenaRank()
	object.Rank = C.int32_t(this.Rank)
	object.Pid = C.int64_t(this.Pid)
	return object
}

func (this insertOP) GlobalArenaRank(row *GlobalArenaRank) {
	key := C.int32_t(row.Rank)
	var old *C.GlobalArenaRank
	for crow := C.g_GlobalTables.GlobalArenaRank; crow != nil; crow = crow.next {
		if crow.Rank == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'global_arena_rank'")
	}
	newRow := row.CObject()
	newRow.next = C.g_GlobalTables.GlobalArenaRank
	C.g_GlobalTables.GlobalArenaRank = newRow
	this.db.addTransLog(this.db.newGlobalArenaRankInsertLog(newRow, row))
}

func (this deleteOP) GlobalArenaRank(row *GlobalArenaRank) {
	key := C.int32_t(row.Rank)
	var old, prev *C.GlobalArenaRank
	for crow := C.g_GlobalTables.GlobalArenaRank; crow != nil; prev, crow = crow, crow.next {
		if crow.Rank == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'global_arena_rank'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalArenaRank = C.g_GlobalTables.GlobalArenaRank.next
	}
	this.db.addTransLog(this.db.newGlobalArenaRankDeleteLog(old, row))
}

func (this updateOP) GlobalArenaRank(row *GlobalArenaRank) {
	key := C.int32_t(row.Rank)
	var old, prev *C.GlobalArenaRank
	for crow := C.g_GlobalTables.GlobalArenaRank; crow != nil; prev, crow = crow, crow.next {
		if crow.Rank == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'global_arena_rank'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		C.g_GlobalTables.GlobalArenaRank = newRow
	}
	tmpRow := GlobalArenaRankRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newGlobalArenaRankUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) GlobalArenaRank(key int32) *GlobalArenaRank {
	for crow := C.g_GlobalTables.GlobalArenaRank; crow != nil; crow = crow.next {
		if crow.Rank == C.int32_t(key) {
			tmpRow := GlobalArenaRankRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) GlobalArenaRank(callback func(*GlobalArenaRankRow)) {
	row := &GlobalArenaRankRow{}
	for crow := C.g_GlobalTables.GlobalArenaRank; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) GlobalArenaRank() (rows []interface{}) {
	row := &GlobalArenaRankRow{}
	for crow := C.g_GlobalTables.GlobalArenaRank; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("GlobalArenaRank", row.GoObject())
	}
	row.row = nil
	fmt.Println("GlobalArenaRank", len(rows))
	return rows
}

/* ========== global_clique ========== */

// 帮派信息
type GlobalClique struct {
	Id int64 // 主键
	Name string // 名字
	Announce string // 公告
	TotalDonateCoins int64 // 累积贡献铜币
	OwnerPid int64 // 帮主
	OwnerLoginTime int64 // 帮主最近登录时间
	MangerPid1 int64 // 副帮主1
	MangerPid2 int64 // 副帮主2
	CenterBuildingCoins int64 // 总部当前贡献铜币
	TempleBuildingCoins int64 // 宗祠贡献铜币
	HealthBuildingCoins int64 // 回春堂当前贡献铜币
	AttackBuildingCoins int64 // 神兵堂当前贡献铜币
	DefenseBuildingCoins int64 // 金刚堂当前贡献铜币
	StoreBuildingCoins int64 // 战备仓库当前贡献铜币
	BankBuildingCoins int64 // 钱庄当前贡献铜币
	CenterBuildingLevel int16 // 总部当前等级
	TempleBuildingLevel int16 // 宗祠等级
	HealthBuildingLevel int16 // 回春堂当前等级
	AttackBuildingLevel int16 // 神兵堂当前等级
	DefenseBuildingLevel int16 // 金刚堂当前等级
	BankBuildingLevel int16 // 钱庄当前等级
	MemberJson []byte // 成员列表存储 pid 的json 数组
	AutoAudit int8 // 自动审核
	AutoAuditLevel int16 // 自动审核等级下限
	Contrib int64 // 帮派贡献
	ContribClearTime int64 // 帮派贡献清除时间
	RecruitTime int64 // 招募公告时间
	WorshipTime int64 // 上香时间
	WorshipCnt int8 // 当日上香次数
	StoreSendTime int64 // 战备物资发送时间
	StoreSendCnt int8 // 战备物资发送次数
}

func (this *GlobalClique) CObject() *C.GlobalClique {
	object := C.New_GlobalClique()
	object.Id = C.int64_t(this.Id)
	object.Name = C.CString(string(this.Name))
	object.Name_len = C.int(len(this.Name))
	object.Announce = C.CString(string(this.Announce))
	object.Announce_len = C.int(len(this.Announce))
	object.TotalDonateCoins = C.int64_t(this.TotalDonateCoins)
	object.OwnerPid = C.int64_t(this.OwnerPid)
	object.OwnerLoginTime = C.int64_t(this.OwnerLoginTime)
	object.MangerPid1 = C.int64_t(this.MangerPid1)
	object.MangerPid2 = C.int64_t(this.MangerPid2)
	object.CenterBuildingCoins = C.int64_t(this.CenterBuildingCoins)
	object.TempleBuildingCoins = C.int64_t(this.TempleBuildingCoins)
	object.HealthBuildingCoins = C.int64_t(this.HealthBuildingCoins)
	object.AttackBuildingCoins = C.int64_t(this.AttackBuildingCoins)
	object.DefenseBuildingCoins = C.int64_t(this.DefenseBuildingCoins)
	object.StoreBuildingCoins = C.int64_t(this.StoreBuildingCoins)
	object.BankBuildingCoins = C.int64_t(this.BankBuildingCoins)
	object.CenterBuildingLevel = C.int16_t(this.CenterBuildingLevel)
	object.TempleBuildingLevel = C.int16_t(this.TempleBuildingLevel)
	object.HealthBuildingLevel = C.int16_t(this.HealthBuildingLevel)
	object.AttackBuildingLevel = C.int16_t(this.AttackBuildingLevel)
	object.DefenseBuildingLevel = C.int16_t(this.DefenseBuildingLevel)
	object.BankBuildingLevel = C.int16_t(this.BankBuildingLevel)
	object.MemberJson = C.CString(string(this.MemberJson))
	object.MemberJson_len = C.int(len(this.MemberJson))
	object.AutoAudit = C.int8_t(this.AutoAudit)
	object.AutoAuditLevel = C.int16_t(this.AutoAuditLevel)
	object.Contrib = C.int64_t(this.Contrib)
	object.ContribClearTime = C.int64_t(this.ContribClearTime)
	object.RecruitTime = C.int64_t(this.RecruitTime)
	object.WorshipTime = C.int64_t(this.WorshipTime)
	object.WorshipCnt = C.int8_t(this.WorshipCnt)
	object.StoreSendTime = C.int64_t(this.StoreSendTime)
	object.StoreSendCnt = C.int8_t(this.StoreSendCnt)
	return object
}

func (this insertOP) GlobalClique(row *GlobalClique) {
	row.Id = atomic.AddInt64(&g_RowIds.GlobalClique, 1)
	key := C.int64_t(row.Id)
	var old *C.GlobalClique
	for crow := C.g_GlobalTables.GlobalClique; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'global_clique'")
	}
	newRow := row.CObject()
	newRow.next = C.g_GlobalTables.GlobalClique
	C.g_GlobalTables.GlobalClique = newRow
	this.db.addTransLog(this.db.newGlobalCliqueInsertLog(newRow, row))
}

func (this deleteOP) GlobalClique(row *GlobalClique) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalClique
	for crow := C.g_GlobalTables.GlobalClique; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'global_clique'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalClique = C.g_GlobalTables.GlobalClique.next
	}
	this.db.addTransLog(this.db.newGlobalCliqueDeleteLog(old, row))
}

func (this updateOP) GlobalClique(row *GlobalClique) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalClique
	for crow := C.g_GlobalTables.GlobalClique; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'global_clique'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		C.g_GlobalTables.GlobalClique = newRow
	}
	tmpRow := GlobalCliqueRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newGlobalCliqueUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) GlobalClique(key int64) *GlobalClique {
	for crow := C.g_GlobalTables.GlobalClique; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := GlobalCliqueRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) GlobalClique(callback func(*GlobalCliqueRow)) {
	row := &GlobalCliqueRow{}
	for crow := C.g_GlobalTables.GlobalClique; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) GlobalClique() (rows []interface{}) {
	row := &GlobalCliqueRow{}
	for crow := C.g_GlobalTables.GlobalClique; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("GlobalClique", row.GoObject())
	}
	row.row = nil
	fmt.Println("GlobalClique", len(rows))
	return rows
}

/* ========== global_clique_boat ========== */

// 镖船信息
type GlobalCliqueBoat struct {
	Id int64 // ID
	CliqueId int64 // 所属帮派ID
	BoatType int16 // 镖船类型
	OwnerPid int64 // 护送人PID
	EscortStartTimestamp int64 // 本次送镖时间戳
	TotalEscortTime int16 // 累计送镖时间
	RecoverPid int64 // 夺回玩家PID
	HijackerPid int64 // 劫持人PID
	HijackTimestamp int64 // 本次送镖时间戳
	Status int8 // 状态 0--运送中 1--劫持中 2--夺回中 3--劫持完成 
}

func (this *GlobalCliqueBoat) CObject() *C.GlobalCliqueBoat {
	object := C.New_GlobalCliqueBoat()
	object.Id = C.int64_t(this.Id)
	object.CliqueId = C.int64_t(this.CliqueId)
	object.BoatType = C.int16_t(this.BoatType)
	object.OwnerPid = C.int64_t(this.OwnerPid)
	object.EscortStartTimestamp = C.int64_t(this.EscortStartTimestamp)
	object.TotalEscortTime = C.int16_t(this.TotalEscortTime)
	object.RecoverPid = C.int64_t(this.RecoverPid)
	object.HijackerPid = C.int64_t(this.HijackerPid)
	object.HijackTimestamp = C.int64_t(this.HijackTimestamp)
	object.Status = C.int8_t(this.Status)
	return object
}

func (this insertOP) GlobalCliqueBoat(row *GlobalCliqueBoat) {
	row.Id = atomic.AddInt64(&g_RowIds.GlobalCliqueBoat, 1)
	key := C.int64_t(row.Id)
	var old *C.GlobalCliqueBoat
	for crow := C.g_GlobalTables.GlobalCliqueBoat; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'global_clique_boat'")
	}
	newRow := row.CObject()
	newRow.next = C.g_GlobalTables.GlobalCliqueBoat
	C.g_GlobalTables.GlobalCliqueBoat = newRow
	this.db.addTransLog(this.db.newGlobalCliqueBoatInsertLog(newRow, row))
}

func (this deleteOP) GlobalCliqueBoat(row *GlobalCliqueBoat) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalCliqueBoat
	for crow := C.g_GlobalTables.GlobalCliqueBoat; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'global_clique_boat'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalCliqueBoat = C.g_GlobalTables.GlobalCliqueBoat.next
	}
	this.db.addTransLog(this.db.newGlobalCliqueBoatDeleteLog(old, row))
}

func (this updateOP) GlobalCliqueBoat(row *GlobalCliqueBoat) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalCliqueBoat
	for crow := C.g_GlobalTables.GlobalCliqueBoat; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'global_clique_boat'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		C.g_GlobalTables.GlobalCliqueBoat = newRow
	}
	tmpRow := GlobalCliqueBoatRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newGlobalCliqueBoatUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) GlobalCliqueBoat(key int64) *GlobalCliqueBoat {
	for crow := C.g_GlobalTables.GlobalCliqueBoat; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := GlobalCliqueBoatRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) GlobalCliqueBoat(callback func(*GlobalCliqueBoatRow)) {
	row := &GlobalCliqueBoatRow{}
	for crow := C.g_GlobalTables.GlobalCliqueBoat; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) GlobalCliqueBoat() (rows []interface{}) {
	row := &GlobalCliqueBoatRow{}
	for crow := C.g_GlobalTables.GlobalCliqueBoat; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("GlobalCliqueBoat", row.GoObject())
	}
	row.row = nil
	fmt.Println("GlobalCliqueBoat", len(rows))
	return rows
}

/* ========== global_despair_boss ========== */

// 绝望关卡boss
type GlobalDespairBoss struct {
	CampType int8 // 阵营类型
	Level int16 // boss等级
	Timestamp int64 // boss出现时间
	DeadTimestamp int64 // boss死亡时间
	Hp int64 // boss当前血量
	MaxHp int64 // boss最大血量
}

func (this *GlobalDespairBoss) CObject() *C.GlobalDespairBoss {
	object := C.New_GlobalDespairBoss()
	object.CampType = C.int8_t(this.CampType)
	object.Level = C.int16_t(this.Level)
	object.Timestamp = C.int64_t(this.Timestamp)
	object.DeadTimestamp = C.int64_t(this.DeadTimestamp)
	object.Hp = C.int64_t(this.Hp)
	object.MaxHp = C.int64_t(this.MaxHp)
	return object
}

func (this insertOP) GlobalDespairBoss(row *GlobalDespairBoss) {
	key := C.int8_t(row.CampType)
	var old *C.GlobalDespairBoss
	for crow := C.g_GlobalTables.GlobalDespairBoss; crow != nil; crow = crow.next {
		if crow.CampType == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'global_despair_boss'")
	}
	newRow := row.CObject()
	newRow.next = C.g_GlobalTables.GlobalDespairBoss
	C.g_GlobalTables.GlobalDespairBoss = newRow
	this.db.addTransLog(this.db.newGlobalDespairBossInsertLog(newRow, row))
}

func (this deleteOP) GlobalDespairBoss(row *GlobalDespairBoss) {
	key := C.int8_t(row.CampType)
	var old, prev *C.GlobalDespairBoss
	for crow := C.g_GlobalTables.GlobalDespairBoss; crow != nil; prev, crow = crow, crow.next {
		if crow.CampType == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'global_despair_boss'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalDespairBoss = C.g_GlobalTables.GlobalDespairBoss.next
	}
	this.db.addTransLog(this.db.newGlobalDespairBossDeleteLog(old, row))
}

func (this updateOP) GlobalDespairBoss(row *GlobalDespairBoss) {
	key := C.int8_t(row.CampType)
	var old, prev *C.GlobalDespairBoss
	for crow := C.g_GlobalTables.GlobalDespairBoss; crow != nil; prev, crow = crow, crow.next {
		if crow.CampType == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'global_despair_boss'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		C.g_GlobalTables.GlobalDespairBoss = newRow
	}
	tmpRow := GlobalDespairBossRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newGlobalDespairBossUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) GlobalDespairBoss(key int8) *GlobalDespairBoss {
	for crow := C.g_GlobalTables.GlobalDespairBoss; crow != nil; crow = crow.next {
		if crow.CampType == C.int8_t(key) {
			tmpRow := GlobalDespairBossRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) GlobalDespairBoss(callback func(*GlobalDespairBossRow)) {
	row := &GlobalDespairBossRow{}
	for crow := C.g_GlobalTables.GlobalDespairBoss; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) GlobalDespairBoss() (rows []interface{}) {
	row := &GlobalDespairBossRow{}
	for crow := C.g_GlobalTables.GlobalDespairBoss; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("GlobalDespairBoss", row.GoObject())
	}
	row.row = nil
	fmt.Println("GlobalDespairBoss", len(rows))
	return rows
}

/* ========== global_despair_boss_history ========== */

// 绝望关卡boss历史记录
type GlobalDespairBossHistory struct {
	Id int64 // 
	Version int64 // 绝望之地version
	CampType int8 // 阵营类型
	Level int16 // boss等级
	Timestamp int64 // 记录产生时间
	StartTimestamp int64 // boss出现时间
	DeadTimestamp int64 // boss死亡时间
}

func (this *GlobalDespairBossHistory) CObject() *C.GlobalDespairBossHistory {
	object := C.New_GlobalDespairBossHistory()
	object.Id = C.int64_t(this.Id)
	object.Version = C.int64_t(this.Version)
	object.CampType = C.int8_t(this.CampType)
	object.Level = C.int16_t(this.Level)
	object.Timestamp = C.int64_t(this.Timestamp)
	object.StartTimestamp = C.int64_t(this.StartTimestamp)
	object.DeadTimestamp = C.int64_t(this.DeadTimestamp)
	return object
}

func (this insertOP) GlobalDespairBossHistory(row *GlobalDespairBossHistory) {
	row.Id = atomic.AddInt64(&g_RowIds.GlobalDespairBossHistory, 1)
	key := C.int64_t(row.Id)
	var old *C.GlobalDespairBossHistory
	for crow := C.g_GlobalTables.GlobalDespairBossHistory; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'global_despair_boss_history'")
	}
	newRow := row.CObject()
	newRow.next = C.g_GlobalTables.GlobalDespairBossHistory
	C.g_GlobalTables.GlobalDespairBossHistory = newRow
	this.db.addTransLog(this.db.newGlobalDespairBossHistoryInsertLog(newRow, row))
}

func (this deleteOP) GlobalDespairBossHistory(row *GlobalDespairBossHistory) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalDespairBossHistory
	for crow := C.g_GlobalTables.GlobalDespairBossHistory; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'global_despair_boss_history'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalDespairBossHistory = C.g_GlobalTables.GlobalDespairBossHistory.next
	}
	this.db.addTransLog(this.db.newGlobalDespairBossHistoryDeleteLog(old, row))
}

func (this updateOP) GlobalDespairBossHistory(row *GlobalDespairBossHistory) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalDespairBossHistory
	for crow := C.g_GlobalTables.GlobalDespairBossHistory; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'global_despair_boss_history'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		C.g_GlobalTables.GlobalDespairBossHistory = newRow
	}
	tmpRow := GlobalDespairBossHistoryRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newGlobalDespairBossHistoryUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) GlobalDespairBossHistory(key int64) *GlobalDespairBossHistory {
	for crow := C.g_GlobalTables.GlobalDespairBossHistory; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := GlobalDespairBossHistoryRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) GlobalDespairBossHistory(callback func(*GlobalDespairBossHistoryRow)) {
	row := &GlobalDespairBossHistoryRow{}
	for crow := C.g_GlobalTables.GlobalDespairBossHistory; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) GlobalDespairBossHistory() (rows []interface{}) {
	row := &GlobalDespairBossHistoryRow{}
	for crow := C.g_GlobalTables.GlobalDespairBossHistory; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("GlobalDespairBossHistory", row.GoObject())
	}
	row.row = nil
	fmt.Println("GlobalDespairBossHistory", len(rows))
	return rows
}

/* ========== global_despair_land ========== */

// 绝望之地状态
type GlobalDespairLand struct {
	Id int64 // 主键ID
	Version int64 // 绝望之地version
	Timestamp int64 // 绝望之地开始时间戳
}

func (this *GlobalDespairLand) CObject() *C.GlobalDespairLand {
	object := C.New_GlobalDespairLand()
	object.Id = C.int64_t(this.Id)
	object.Version = C.int64_t(this.Version)
	object.Timestamp = C.int64_t(this.Timestamp)
	return object
}

func (this insertOP) GlobalDespairLand(row *GlobalDespairLand) {
	row.Id = atomic.AddInt64(&g_RowIds.GlobalDespairLand, 1)
	key := C.int64_t(row.Id)
	var old *C.GlobalDespairLand
	for crow := C.g_GlobalTables.GlobalDespairLand; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'global_despair_land'")
	}
	newRow := row.CObject()
	newRow.next = C.g_GlobalTables.GlobalDespairLand
	C.g_GlobalTables.GlobalDespairLand = newRow
	this.db.addTransLog(this.db.newGlobalDespairLandInsertLog(newRow, row))
}

func (this deleteOP) GlobalDespairLand(row *GlobalDespairLand) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalDespairLand
	for crow := C.g_GlobalTables.GlobalDespairLand; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'global_despair_land'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalDespairLand = C.g_GlobalTables.GlobalDespairLand.next
	}
	this.db.addTransLog(this.db.newGlobalDespairLandDeleteLog(old, row))
}

func (this updateOP) GlobalDespairLand(row *GlobalDespairLand) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalDespairLand
	for crow := C.g_GlobalTables.GlobalDespairLand; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'global_despair_land'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		C.g_GlobalTables.GlobalDespairLand = newRow
	}
	tmpRow := GlobalDespairLandRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newGlobalDespairLandUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) GlobalDespairLand(key int64) *GlobalDespairLand {
	for crow := C.g_GlobalTables.GlobalDespairLand; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := GlobalDespairLandRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) GlobalDespairLand(callback func(*GlobalDespairLandRow)) {
	row := &GlobalDespairLandRow{}
	for crow := C.g_GlobalTables.GlobalDespairLand; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) GlobalDespairLand() (rows []interface{}) {
	row := &GlobalDespairLandRow{}
	for crow := C.g_GlobalTables.GlobalDespairLand; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("GlobalDespairLand", row.GoObject())
	}
	row.row = nil
	fmt.Println("GlobalDespairLand", len(rows))
	return rows
}

/* ========== global_despair_land_battle_record ========== */

// 玩家战报记录
type GlobalDespairLandBattleRecord struct {
	Id int64 // 主键ID
	Pid int64 // 玩家id
	Fightnum int32 // 玩家战力（当时）
	Timestamp int64 // 战斗时间
	Tag int16 // 1-绝望之地最早通关 2-绝望之地最近通关 4-绝望之地最低战力通关
	BattleVersion int16 // 战场服务端版本
	CampType int8 // 战场类型
	LevelId int32 // 关卡id
	Record []byte // 战报
}

func (this *GlobalDespairLandBattleRecord) CObject() *C.GlobalDespairLandBattleRecord {
	object := C.New_GlobalDespairLandBattleRecord()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.Fightnum = C.int32_t(this.Fightnum)
	object.Timestamp = C.int64_t(this.Timestamp)
	object.Tag = C.int16_t(this.Tag)
	object.BattleVersion = C.int16_t(this.BattleVersion)
	object.CampType = C.int8_t(this.CampType)
	object.LevelId = C.int32_t(this.LevelId)
	object.Record = C.CString(string(this.Record))
	object.Record_len = C.int(len(this.Record))
	return object
}

func (this insertOP) GlobalDespairLandBattleRecord(row *GlobalDespairLandBattleRecord) {
	row.Id = atomic.AddInt64(&g_RowIds.GlobalDespairLandBattleRecord, 1)
	key := C.int64_t(row.Id)
	var old *C.GlobalDespairLandBattleRecord
	for crow := C.g_GlobalTables.GlobalDespairLandBattleRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'global_despair_land_battle_record'")
	}
	newRow := row.CObject()
	newRow.next = C.g_GlobalTables.GlobalDespairLandBattleRecord
	C.g_GlobalTables.GlobalDespairLandBattleRecord = newRow
	this.db.addTransLog(this.db.newGlobalDespairLandBattleRecordInsertLog(newRow, row))
}

func (this deleteOP) GlobalDespairLandBattleRecord(row *GlobalDespairLandBattleRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalDespairLandBattleRecord
	for crow := C.g_GlobalTables.GlobalDespairLandBattleRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'global_despair_land_battle_record'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalDespairLandBattleRecord = C.g_GlobalTables.GlobalDespairLandBattleRecord.next
	}
	this.db.addTransLog(this.db.newGlobalDespairLandBattleRecordDeleteLog(old, row))
}

func (this updateOP) GlobalDespairLandBattleRecord(row *GlobalDespairLandBattleRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalDespairLandBattleRecord
	for crow := C.g_GlobalTables.GlobalDespairLandBattleRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'global_despair_land_battle_record'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		C.g_GlobalTables.GlobalDespairLandBattleRecord = newRow
	}
	tmpRow := GlobalDespairLandBattleRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newGlobalDespairLandBattleRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) GlobalDespairLandBattleRecord(key int64) *GlobalDespairLandBattleRecord {
	for crow := C.g_GlobalTables.GlobalDespairLandBattleRecord; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := GlobalDespairLandBattleRecordRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) GlobalDespairLandBattleRecord(callback func(*GlobalDespairLandBattleRecordRow)) {
	row := &GlobalDespairLandBattleRecordRow{}
	for crow := C.g_GlobalTables.GlobalDespairLandBattleRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) GlobalDespairLandBattleRecord() (rows []interface{}) {
	row := &GlobalDespairLandBattleRecordRow{}
	for crow := C.g_GlobalTables.GlobalDespairLandBattleRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("GlobalDespairLandBattleRecord", row.GoObject())
	}
	row.row = nil
	fmt.Println("GlobalDespairLandBattleRecord", len(rows))
	return rows
}

/* ========== global_despair_land_camp ========== */

// 绝望之地阵营状态
type GlobalDespairLandCamp struct {
	CampType int8 // 阵营
	Timestamp int64 // boss开始时间戳
	DeadTimestamp int64 // boss死亡时间戳
	Level int16 // boss等级也即进度等级
}

func (this *GlobalDespairLandCamp) CObject() *C.GlobalDespairLandCamp {
	object := C.New_GlobalDespairLandCamp()
	object.CampType = C.int8_t(this.CampType)
	object.Timestamp = C.int64_t(this.Timestamp)
	object.DeadTimestamp = C.int64_t(this.DeadTimestamp)
	object.Level = C.int16_t(this.Level)
	return object
}

func (this insertOP) GlobalDespairLandCamp(row *GlobalDespairLandCamp) {
	key := C.int8_t(row.CampType)
	var old *C.GlobalDespairLandCamp
	for crow := C.g_GlobalTables.GlobalDespairLandCamp; crow != nil; crow = crow.next {
		if crow.CampType == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'global_despair_land_camp'")
	}
	newRow := row.CObject()
	newRow.next = C.g_GlobalTables.GlobalDespairLandCamp
	C.g_GlobalTables.GlobalDespairLandCamp = newRow
	this.db.addTransLog(this.db.newGlobalDespairLandCampInsertLog(newRow, row))
}

func (this deleteOP) GlobalDespairLandCamp(row *GlobalDespairLandCamp) {
	key := C.int8_t(row.CampType)
	var old, prev *C.GlobalDespairLandCamp
	for crow := C.g_GlobalTables.GlobalDespairLandCamp; crow != nil; prev, crow = crow, crow.next {
		if crow.CampType == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'global_despair_land_camp'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalDespairLandCamp = C.g_GlobalTables.GlobalDespairLandCamp.next
	}
	this.db.addTransLog(this.db.newGlobalDespairLandCampDeleteLog(old, row))
}

func (this updateOP) GlobalDespairLandCamp(row *GlobalDespairLandCamp) {
	key := C.int8_t(row.CampType)
	var old, prev *C.GlobalDespairLandCamp
	for crow := C.g_GlobalTables.GlobalDespairLandCamp; crow != nil; prev, crow = crow, crow.next {
		if crow.CampType == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'global_despair_land_camp'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		C.g_GlobalTables.GlobalDespairLandCamp = newRow
	}
	tmpRow := GlobalDespairLandCampRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newGlobalDespairLandCampUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) GlobalDespairLandCamp(key int8) *GlobalDespairLandCamp {
	for crow := C.g_GlobalTables.GlobalDespairLandCamp; crow != nil; crow = crow.next {
		if crow.CampType == C.int8_t(key) {
			tmpRow := GlobalDespairLandCampRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) GlobalDespairLandCamp(callback func(*GlobalDespairLandCampRow)) {
	row := &GlobalDespairLandCampRow{}
	for crow := C.g_GlobalTables.GlobalDespairLandCamp; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) GlobalDespairLandCamp() (rows []interface{}) {
	row := &GlobalDespairLandCampRow{}
	for crow := C.g_GlobalTables.GlobalDespairLandCamp; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("GlobalDespairLandCamp", row.GoObject())
	}
	row.row = nil
	fmt.Println("GlobalDespairLandCamp", len(rows))
	return rows
}

/* ========== global_gift_card_record ========== */

// 礼品卡兑换记录
type GlobalGiftCardRecord struct {
	Id int64 // 主键
	Pid int64 // 兑换玩家PID
	Code string // 兑换码
	Timestamp int64 // 兑换时间
}

func (this *GlobalGiftCardRecord) CObject() *C.GlobalGiftCardRecord {
	object := C.New_GlobalGiftCardRecord()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.Code = C.CString(string(this.Code))
	object.Code_len = C.int(len(this.Code))
	object.Timestamp = C.int64_t(this.Timestamp)
	return object
}

func (this insertOP) GlobalGiftCardRecord(row *GlobalGiftCardRecord) {
	row.Id = atomic.AddInt64(&g_RowIds.GlobalGiftCardRecord, 1)
	key := C.int64_t(row.Id)
	var old *C.GlobalGiftCardRecord
	for crow := C.g_GlobalTables.GlobalGiftCardRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'global_gift_card_record'")
	}
	newRow := row.CObject()
	newRow.next = C.g_GlobalTables.GlobalGiftCardRecord
	C.g_GlobalTables.GlobalGiftCardRecord = newRow
	this.db.addTransLog(this.db.newGlobalGiftCardRecordInsertLog(newRow, row))
}

func (this deleteOP) GlobalGiftCardRecord(row *GlobalGiftCardRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalGiftCardRecord
	for crow := C.g_GlobalTables.GlobalGiftCardRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'global_gift_card_record'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalGiftCardRecord = C.g_GlobalTables.GlobalGiftCardRecord.next
	}
	this.db.addTransLog(this.db.newGlobalGiftCardRecordDeleteLog(old, row))
}

func (this updateOP) GlobalGiftCardRecord(row *GlobalGiftCardRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalGiftCardRecord
	for crow := C.g_GlobalTables.GlobalGiftCardRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'global_gift_card_record'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		C.g_GlobalTables.GlobalGiftCardRecord = newRow
	}
	tmpRow := GlobalGiftCardRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newGlobalGiftCardRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) GlobalGiftCardRecord(key int64) *GlobalGiftCardRecord {
	for crow := C.g_GlobalTables.GlobalGiftCardRecord; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := GlobalGiftCardRecordRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) GlobalGiftCardRecord(callback func(*GlobalGiftCardRecordRow)) {
	row := &GlobalGiftCardRecordRow{}
	for crow := C.g_GlobalTables.GlobalGiftCardRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) GlobalGiftCardRecord() (rows []interface{}) {
	row := &GlobalGiftCardRecordRow{}
	for crow := C.g_GlobalTables.GlobalGiftCardRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("GlobalGiftCardRecord", row.GoObject())
	}
	row.row = nil
	fmt.Println("GlobalGiftCardRecord", len(rows))
	return rows
}

/* ========== global_group_buy_status ========== */

// hd服务器记录团购状态信息
type GlobalGroupBuyStatus struct {
	Id int64 // ID标识
	Cid int32 // 外键，指定对应得团购物品记录id
	Status int32 // 当前团购状态，即购买总数
}

func (this *GlobalGroupBuyStatus) CObject() *C.GlobalGroupBuyStatus {
	object := C.New_GlobalGroupBuyStatus()
	object.Id = C.int64_t(this.Id)
	object.Cid = C.int32_t(this.Cid)
	object.Status = C.int32_t(this.Status)
	return object
}

func (this insertOP) GlobalGroupBuyStatus(row *GlobalGroupBuyStatus) {
	row.Id = atomic.AddInt64(&g_RowIds.GlobalGroupBuyStatus, 1)
	key := C.int64_t(row.Id)
	var old *C.GlobalGroupBuyStatus
	for crow := C.g_GlobalTables.GlobalGroupBuyStatus; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'global_group_buy_status'")
	}
	newRow := row.CObject()
	newRow.next = C.g_GlobalTables.GlobalGroupBuyStatus
	C.g_GlobalTables.GlobalGroupBuyStatus = newRow
	this.db.addTransLog(this.db.newGlobalGroupBuyStatusInsertLog(newRow, row))
}

func (this deleteOP) GlobalGroupBuyStatus(row *GlobalGroupBuyStatus) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalGroupBuyStatus
	for crow := C.g_GlobalTables.GlobalGroupBuyStatus; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'global_group_buy_status'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalGroupBuyStatus = C.g_GlobalTables.GlobalGroupBuyStatus.next
	}
	this.db.addTransLog(this.db.newGlobalGroupBuyStatusDeleteLog(old, row))
}

func (this updateOP) GlobalGroupBuyStatus(row *GlobalGroupBuyStatus) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalGroupBuyStatus
	for crow := C.g_GlobalTables.GlobalGroupBuyStatus; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'global_group_buy_status'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		C.g_GlobalTables.GlobalGroupBuyStatus = newRow
	}
	tmpRow := GlobalGroupBuyStatusRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newGlobalGroupBuyStatusUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) GlobalGroupBuyStatus(key int64) *GlobalGroupBuyStatus {
	for crow := C.g_GlobalTables.GlobalGroupBuyStatus; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := GlobalGroupBuyStatusRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) GlobalGroupBuyStatus(callback func(*GlobalGroupBuyStatusRow)) {
	row := &GlobalGroupBuyStatusRow{}
	for crow := C.g_GlobalTables.GlobalGroupBuyStatus; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) GlobalGroupBuyStatus() (rows []interface{}) {
	row := &GlobalGroupBuyStatusRow{}
	for crow := C.g_GlobalTables.GlobalGroupBuyStatus; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("GlobalGroupBuyStatus", row.GoObject())
	}
	row.row = nil
	fmt.Println("GlobalGroupBuyStatus", len(rows))
	return rows
}

/* ========== global_mail ========== */

// 全局邮件
type GlobalMail struct {
	Id int64 // 公告ID
	SendTime int64 // 发送时间戳
	ExpireTime int64 // 创建时间戳
	Title string // 标题
	Content string // 内容
	Priority int8 // 优先级
	MinLevel int16 // 要求最低等级
	MaxLevel int16 // 要求最高等级
	MinVipLevel int16 // 要求VIP最低等级
	MaxVipLevel int16 // 要求VIP最等级
	MinCreateTime int64 // 要求最早注册时间
	MaxCreateTime int64 // 要求最晚注册时间
}

func (this *GlobalMail) CObject() *C.GlobalMail {
	object := C.New_GlobalMail()
	object.Id = C.int64_t(this.Id)
	object.SendTime = C.int64_t(this.SendTime)
	object.ExpireTime = C.int64_t(this.ExpireTime)
	object.Title = C.CString(string(this.Title))
	object.Title_len = C.int(len(this.Title))
	object.Content = C.CString(string(this.Content))
	object.Content_len = C.int(len(this.Content))
	object.Priority = C.int8_t(this.Priority)
	object.MinLevel = C.int16_t(this.MinLevel)
	object.MaxLevel = C.int16_t(this.MaxLevel)
	object.MinVipLevel = C.int16_t(this.MinVipLevel)
	object.MaxVipLevel = C.int16_t(this.MaxVipLevel)
	object.MinCreateTime = C.int64_t(this.MinCreateTime)
	object.MaxCreateTime = C.int64_t(this.MaxCreateTime)
	return object
}

func (this insertOP) GlobalMail(row *GlobalMail) {
	row.Id = atomic.AddInt64(&g_RowIds.GlobalMail, 1)
	key := C.int64_t(row.Id)
	var old *C.GlobalMail
	for crow := C.g_GlobalTables.GlobalMail; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'global_mail'")
	}
	newRow := row.CObject()
	newRow.next = C.g_GlobalTables.GlobalMail
	C.g_GlobalTables.GlobalMail = newRow
	this.db.addTransLog(this.db.newGlobalMailInsertLog(newRow, row))
}

func (this deleteOP) GlobalMail(row *GlobalMail) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalMail
	for crow := C.g_GlobalTables.GlobalMail; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'global_mail'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalMail = C.g_GlobalTables.GlobalMail.next
	}
	this.db.addTransLog(this.db.newGlobalMailDeleteLog(old, row))
}

func (this updateOP) GlobalMail(row *GlobalMail) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalMail
	for crow := C.g_GlobalTables.GlobalMail; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'global_mail'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		C.g_GlobalTables.GlobalMail = newRow
	}
	tmpRow := GlobalMailRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newGlobalMailUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) GlobalMail(key int64) *GlobalMail {
	for crow := C.g_GlobalTables.GlobalMail; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := GlobalMailRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) GlobalMail(callback func(*GlobalMailRow)) {
	row := &GlobalMailRow{}
	for crow := C.g_GlobalTables.GlobalMail; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) GlobalMail() (rows []interface{}) {
	row := &GlobalMailRow{}
	for crow := C.g_GlobalTables.GlobalMail; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("GlobalMail", row.GoObject())
	}
	row.row = nil
	fmt.Println("GlobalMail", len(rows))
	return rows
}

/* ========== global_mail_attachments ========== */

// 全局邮件附件
type GlobalMailAttachments struct {
	Id int64 // 附件邮件ID
	GlobalMailId int64 // 全局邮件表主键
	ItemId int16 // 附件ID
	AttachmentType int8 // 附件类型
	ItemNum int64 // 数量
}

func (this *GlobalMailAttachments) CObject() *C.GlobalMailAttachments {
	object := C.New_GlobalMailAttachments()
	object.Id = C.int64_t(this.Id)
	object.GlobalMailId = C.int64_t(this.GlobalMailId)
	object.ItemId = C.int16_t(this.ItemId)
	object.AttachmentType = C.int8_t(this.AttachmentType)
	object.ItemNum = C.int64_t(this.ItemNum)
	return object
}

func (this insertOP) GlobalMailAttachments(row *GlobalMailAttachments) {
	row.Id = atomic.AddInt64(&g_RowIds.GlobalMailAttachments, 1)
	key := C.int64_t(row.Id)
	var old *C.GlobalMailAttachments
	for crow := C.g_GlobalTables.GlobalMailAttachments; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'global_mail_attachments'")
	}
	newRow := row.CObject()
	newRow.next = C.g_GlobalTables.GlobalMailAttachments
	C.g_GlobalTables.GlobalMailAttachments = newRow
	this.db.addTransLog(this.db.newGlobalMailAttachmentsInsertLog(newRow, row))
}

func (this deleteOP) GlobalMailAttachments(row *GlobalMailAttachments) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalMailAttachments
	for crow := C.g_GlobalTables.GlobalMailAttachments; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'global_mail_attachments'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalMailAttachments = C.g_GlobalTables.GlobalMailAttachments.next
	}
	this.db.addTransLog(this.db.newGlobalMailAttachmentsDeleteLog(old, row))
}

func (this updateOP) GlobalMailAttachments(row *GlobalMailAttachments) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalMailAttachments
	for crow := C.g_GlobalTables.GlobalMailAttachments; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'global_mail_attachments'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		C.g_GlobalTables.GlobalMailAttachments = newRow
	}
	tmpRow := GlobalMailAttachmentsRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newGlobalMailAttachmentsUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) GlobalMailAttachments(key int64) *GlobalMailAttachments {
	for crow := C.g_GlobalTables.GlobalMailAttachments; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := GlobalMailAttachmentsRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) GlobalMailAttachments(callback func(*GlobalMailAttachmentsRow)) {
	row := &GlobalMailAttachmentsRow{}
	for crow := C.g_GlobalTables.GlobalMailAttachments; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) GlobalMailAttachments() (rows []interface{}) {
	row := &GlobalMailAttachmentsRow{}
	for crow := C.g_GlobalTables.GlobalMailAttachments; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("GlobalMailAttachments", row.GoObject())
	}
	row.row = nil
	fmt.Println("GlobalMailAttachments", len(rows))
	return rows
}

/* ========== global_tb_xxd_onlinecnt ========== */

// 腾讯经分在线玩家统计日志
type GlobalTbXxdOnlinecnt struct {
	Id int64 // 主键ID
	Gameappid string // 平台分配的AppID
	Timekey int64 // 当前时间除以60s，下取整
	Gsid int64 // 游戏服务器编号
	Onlinecntios int64 // ios在线人数
	Onlinecntandroid int64 // android在线人数
	Cid int64 // 渠道ID
}

func (this *GlobalTbXxdOnlinecnt) CObject() *C.GlobalTbXxdOnlinecnt {
	object := C.New_GlobalTbXxdOnlinecnt()
	object.Id = C.int64_t(this.Id)
	object.Gameappid = C.CString(string(this.Gameappid))
	object.Gameappid_len = C.int(len(this.Gameappid))
	object.Timekey = C.int64_t(this.Timekey)
	object.Gsid = C.int64_t(this.Gsid)
	object.Onlinecntios = C.int64_t(this.Onlinecntios)
	object.Onlinecntandroid = C.int64_t(this.Onlinecntandroid)
	object.Cid = C.int64_t(this.Cid)
	return object
}

func (this insertOP) GlobalTbXxdOnlinecnt(row *GlobalTbXxdOnlinecnt) {
	row.Id = atomic.AddInt64(&g_RowIds.GlobalTbXxdOnlinecnt, 1)
	key := C.int64_t(row.Id)
	var old *C.GlobalTbXxdOnlinecnt
	for crow := C.g_GlobalTables.GlobalTbXxdOnlinecnt; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'global_tb_xxd_onlinecnt'")
	}
	newRow := row.CObject()
	newRow.next = C.g_GlobalTables.GlobalTbXxdOnlinecnt
	C.g_GlobalTables.GlobalTbXxdOnlinecnt = newRow
	this.db.addTransLog(this.db.newGlobalTbXxdOnlinecntInsertLog(newRow, row))
}

func (this deleteOP) GlobalTbXxdOnlinecnt(row *GlobalTbXxdOnlinecnt) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalTbXxdOnlinecnt
	for crow := C.g_GlobalTables.GlobalTbXxdOnlinecnt; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'global_tb_xxd_onlinecnt'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalTbXxdOnlinecnt = C.g_GlobalTables.GlobalTbXxdOnlinecnt.next
	}
	this.db.addTransLog(this.db.newGlobalTbXxdOnlinecntDeleteLog(old, row))
}

func (this updateOP) GlobalTbXxdOnlinecnt(row *GlobalTbXxdOnlinecnt) {
	key := C.int64_t(row.Id)
	var old, prev *C.GlobalTbXxdOnlinecnt
	for crow := C.g_GlobalTables.GlobalTbXxdOnlinecnt; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'global_tb_xxd_onlinecnt'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		C.g_GlobalTables.GlobalTbXxdOnlinecnt = newRow
	}
	tmpRow := GlobalTbXxdOnlinecntRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newGlobalTbXxdOnlinecntUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) GlobalTbXxdOnlinecnt(key int64) *GlobalTbXxdOnlinecnt {
	for crow := C.g_GlobalTables.GlobalTbXxdOnlinecnt; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := GlobalTbXxdOnlinecntRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) GlobalTbXxdOnlinecnt(callback func(*GlobalTbXxdOnlinecntRow)) {
	row := &GlobalTbXxdOnlinecntRow{}
	for crow := C.g_GlobalTables.GlobalTbXxdOnlinecnt; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) GlobalTbXxdOnlinecnt() (rows []interface{}) {
	row := &GlobalTbXxdOnlinecntRow{}
	for crow := C.g_GlobalTables.GlobalTbXxdOnlinecnt; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("GlobalTbXxdOnlinecnt", row.GoObject())
	}
	row.row = nil
	fmt.Println("GlobalTbXxdOnlinecnt", len(rows))
	return rows
}

/* ========== player ========== */

// 玩家基础信息
type Player struct {
	Id int64 // 玩家ID
	User string // 平台传递过来的用户标识
	Nick string // 玩家昵称
	MainRoleId int64 // 主角ID
	Cid int64 // 渠道ID
}

func (this *Player) CObject() *C.Player {
	object := C.New_Player()
	object.Id = C.int64_t(this.Id)
	object.User = C.CString(string(this.User))
	object.User_len = C.int(len(this.User))
	object.Nick = C.CString(string(this.Nick))
	object.Nick_len = C.int(len(this.Nick))
	object.MainRoleId = C.int64_t(this.MainRoleId)
	object.Cid = C.int64_t(this.Cid)
	return object
}

func (this insertOP) Player(row *Player) {
	if this.db.tables.Player != nil {
		panic("duplicate insert 'player'")
	}
	row.Id = atomic.AddInt64(&g_RowIds.Player, 1)
	this.db.playerId = row.Id
  this.db.transController.lock = g_LockManager.Get(this.db.playerId)
	this.db.tables.Pid = C.int64_t(row.Id)
	SetPlayerTables(this.db.tables)
	newRow := row.CObject()
	this.db.tables.Player = newRow
	this.db.addTransLog(this.db.newPlayerInsertLog(newRow, row))
}

func (this deleteOP) Player(row *Player) {
	if this.db.tables.Player == nil {
		panic("delete not exists 'player'")
	}
	old := this.db.tables.Player
	this.db.tables.Player = nil
	this.db.addTransLog(this.db.newPlayerDeleteLog(old, row))
}

func (this updateOP) Player(row *Player) {
	if this.db.tables.Player == nil {
		panic("update not exists 'player'")
	}
	newRow := row.CObject()
	old := this.db.tables.Player
	this.db.tables.Player = newRow
	tmpRow := PlayerRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) Player(key int64) *Player {
	if this.db.tables.Player == nil {
		return nil
	}
	tmpRow := PlayerRow{row:this.db.tables.Player}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_activity ========== */

// 玩家活跃度
type PlayerActivity struct {
	Pid int64 // 用户ID
	Activity int32 // 用户活跃度
	LastUpdate int64 // 最后一次更新时间戳
}

func (this *PlayerActivity) CObject() *C.PlayerActivity {
	object := C.New_PlayerActivity()
	object.Pid = C.int64_t(this.Pid)
	object.Activity = C.int32_t(this.Activity)
	object.LastUpdate = C.int64_t(this.LastUpdate)
	return object
}

func (this insertOP) PlayerActivity(row *PlayerActivity) {
	if this.db.tables.PlayerActivity != nil {
		panic("duplicate insert 'player_activity'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerActivity = newRow
	this.db.addTransLog(this.db.newPlayerActivityInsertLog(newRow, row))
}

func (this deleteOP) PlayerActivity(row *PlayerActivity) {
	if this.db.tables.PlayerActivity == nil {
		panic("delete not exists 'player_activity'")
	}
	old := this.db.tables.PlayerActivity
	this.db.tables.PlayerActivity = nil
	this.db.addTransLog(this.db.newPlayerActivityDeleteLog(old, row))
}

func (this updateOP) PlayerActivity(row *PlayerActivity) {
	if this.db.tables.PlayerActivity == nil {
		panic("update not exists 'player_activity'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerActivity
	this.db.tables.PlayerActivity = newRow
	tmpRow := PlayerActivityRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerActivityUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerActivity(key int64) *PlayerActivity {
	if this.db.tables.PlayerActivity == nil {
		return nil
	}
	tmpRow := PlayerActivityRow{row:this.db.tables.PlayerActivity}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_addition_quest ========== */

// 玩家支线任务
type PlayerAdditionQuest struct {
	Id int64 // ID
	Pid int64 // 用户ID
	SerialNumber int32 // 任务系列ID
	QuestId int32 // 当前任务ID
	Lock int16 // 任务链权值
	Progress int16 // 任务进度
	State int8 // 0--未完成 1--已完成 2--已奖励 3--已放弃
}

func (this *PlayerAdditionQuest) CObject() *C.PlayerAdditionQuest {
	object := C.New_PlayerAdditionQuest()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.SerialNumber = C.int32_t(this.SerialNumber)
	object.QuestId = C.int32_t(this.QuestId)
	object.Lock = C.int16_t(this.Lock)
	object.Progress = C.int16_t(this.Progress)
	object.State = C.int8_t(this.State)
	return object
}

func (this insertOP) PlayerAdditionQuest(row *PlayerAdditionQuest) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerAdditionQuest, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerAdditionQuest
	for crow := this.db.tables.PlayerAdditionQuest; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_addition_quest'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerAdditionQuest
	this.db.tables.PlayerAdditionQuest = newRow
	this.db.addTransLog(this.db.newPlayerAdditionQuestInsertLog(newRow, row))
}

func (this deleteOP) PlayerAdditionQuest(row *PlayerAdditionQuest) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerAdditionQuest
	for crow := this.db.tables.PlayerAdditionQuest; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_addition_quest'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerAdditionQuest = this.db.tables.PlayerAdditionQuest.next
	}
	this.db.addTransLog(this.db.newPlayerAdditionQuestDeleteLog(old, row))
}

func (this updateOP) PlayerAdditionQuest(row *PlayerAdditionQuest) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerAdditionQuest
	for crow := this.db.tables.PlayerAdditionQuest; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_addition_quest'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerAdditionQuest = newRow
	}
	tmpRow := PlayerAdditionQuestRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerAdditionQuestUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerAdditionQuest(key int64) *PlayerAdditionQuest {
	for crow := this.db.tables.PlayerAdditionQuest; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerAdditionQuestRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerAdditionQuest(callback func(*PlayerAdditionQuestRow)) {
	row := &PlayerAdditionQuestRow{}
	for crow := this.db.tables.PlayerAdditionQuest; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerAdditionQuest() (rows []interface{}) {
	row := &PlayerAdditionQuestRow{}
	for crow := this.db.tables.PlayerAdditionQuest; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerAdditionQuest", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerAdditionQuest", len(rows))
	return rows
}

/* ========== player_any_sdk_order ========== */

// anysdk 订单处理纪录
type PlayerAnySdkOrder struct {
	Id int64 // 玩家ID
	Pid int64 // 玩家ID
	OrderId int64 // 内部订单ID
	Present int64 // 赠送元宝元宝
}

func (this *PlayerAnySdkOrder) CObject() *C.PlayerAnySdkOrder {
	object := C.New_PlayerAnySdkOrder()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.OrderId = C.int64_t(this.OrderId)
	object.Present = C.int64_t(this.Present)
	return object
}

func (this insertOP) PlayerAnySdkOrder(row *PlayerAnySdkOrder) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerAnySdkOrder, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerAnySdkOrder
	for crow := this.db.tables.PlayerAnySdkOrder; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_any_sdk_order'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerAnySdkOrder
	this.db.tables.PlayerAnySdkOrder = newRow
	this.db.addTransLog(this.db.newPlayerAnySdkOrderInsertLog(newRow, row))
}

func (this deleteOP) PlayerAnySdkOrder(row *PlayerAnySdkOrder) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerAnySdkOrder
	for crow := this.db.tables.PlayerAnySdkOrder; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_any_sdk_order'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerAnySdkOrder = this.db.tables.PlayerAnySdkOrder.next
	}
	this.db.addTransLog(this.db.newPlayerAnySdkOrderDeleteLog(old, row))
}

func (this updateOP) PlayerAnySdkOrder(row *PlayerAnySdkOrder) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerAnySdkOrder
	for crow := this.db.tables.PlayerAnySdkOrder; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_any_sdk_order'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerAnySdkOrder = newRow
	}
	tmpRow := PlayerAnySdkOrderRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerAnySdkOrderUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerAnySdkOrder(key int64) *PlayerAnySdkOrder {
	for crow := this.db.tables.PlayerAnySdkOrder; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerAnySdkOrderRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerAnySdkOrder(callback func(*PlayerAnySdkOrderRow)) {
	row := &PlayerAnySdkOrderRow{}
	for crow := this.db.tables.PlayerAnySdkOrder; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerAnySdkOrder() (rows []interface{}) {
	row := &PlayerAnySdkOrderRow{}
	for crow := this.db.tables.PlayerAnySdkOrder; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerAnySdkOrder", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerAnySdkOrder", len(rows))
	return rows
}

/* ========== player_arena ========== */

// 玩家比武场数据
type PlayerArena struct {
	Pid int64 // 玩家ID
	DailyNum int16 // 今日已挑战次数
	FailedCdTime int64 // 战败CD结束时间
	BattleTime int64 // 最近一次挑战时间
	WinTimes int16 // >0 连胜场次; 0 保持不变; -1 下降趋势
	DailyAwardItem int32 // 今日获得暗影果实累计
	NewRecordCount int16 // 新战报计数
	DailyFame int32 // 每日奖励声望
	BuyTimes int16 // 今日购买次数
}

func (this *PlayerArena) CObject() *C.PlayerArena {
	object := C.New_PlayerArena()
	object.Pid = C.int64_t(this.Pid)
	object.DailyNum = C.int16_t(this.DailyNum)
	object.FailedCdTime = C.int64_t(this.FailedCdTime)
	object.BattleTime = C.int64_t(this.BattleTime)
	object.WinTimes = C.int16_t(this.WinTimes)
	object.DailyAwardItem = C.int32_t(this.DailyAwardItem)
	object.NewRecordCount = C.int16_t(this.NewRecordCount)
	object.DailyFame = C.int32_t(this.DailyFame)
	object.BuyTimes = C.int16_t(this.BuyTimes)
	return object
}

func (this insertOP) PlayerArena(row *PlayerArena) {
	if this.db.tables.PlayerArena != nil {
		panic("duplicate insert 'player_arena'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerArena = newRow
	this.db.addTransLog(this.db.newPlayerArenaInsertLog(newRow, row))
}

func (this deleteOP) PlayerArena(row *PlayerArena) {
	if this.db.tables.PlayerArena == nil {
		panic("delete not exists 'player_arena'")
	}
	old := this.db.tables.PlayerArena
	this.db.tables.PlayerArena = nil
	this.db.addTransLog(this.db.newPlayerArenaDeleteLog(old, row))
}

func (this updateOP) PlayerArena(row *PlayerArena) {
	if this.db.tables.PlayerArena == nil {
		panic("update not exists 'player_arena'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerArena
	this.db.tables.PlayerArena = newRow
	tmpRow := PlayerArenaRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerArenaUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerArena(key int64) *PlayerArena {
	if this.db.tables.PlayerArena == nil {
		return nil
	}
	tmpRow := PlayerArenaRow{row:this.db.tables.PlayerArena}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_arena_rank ========== */

// 玩家比武场最近排名记录
type PlayerArenaRank struct {
	Pid int64 // 玩家ID
	Rank int32 // 昨天排名
	Rank1 int32 // 1天前排名
	Rank2 int32 // 2天前排名
	Rank3 int32 // 3天前排名
	Time int64 // 宝箱刷新时间
}

func (this *PlayerArenaRank) CObject() *C.PlayerArenaRank {
	object := C.New_PlayerArenaRank()
	object.Pid = C.int64_t(this.Pid)
	object.Rank = C.int32_t(this.Rank)
	object.Rank1 = C.int32_t(this.Rank1)
	object.Rank2 = C.int32_t(this.Rank2)
	object.Rank3 = C.int32_t(this.Rank3)
	object.Time = C.int64_t(this.Time)
	return object
}

func (this insertOP) PlayerArenaRank(row *PlayerArenaRank) {
	if this.db.tables.PlayerArenaRank != nil {
		panic("duplicate insert 'player_arena_rank'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerArenaRank = newRow
	this.db.addTransLog(this.db.newPlayerArenaRankInsertLog(newRow, row))
}

func (this deleteOP) PlayerArenaRank(row *PlayerArenaRank) {
	if this.db.tables.PlayerArenaRank == nil {
		panic("delete not exists 'player_arena_rank'")
	}
	old := this.db.tables.PlayerArenaRank
	this.db.tables.PlayerArenaRank = nil
	this.db.addTransLog(this.db.newPlayerArenaRankDeleteLog(old, row))
}

func (this updateOP) PlayerArenaRank(row *PlayerArenaRank) {
	if this.db.tables.PlayerArenaRank == nil {
		panic("update not exists 'player_arena_rank'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerArenaRank
	this.db.tables.PlayerArenaRank = newRow
	tmpRow := PlayerArenaRankRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerArenaRankUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerArenaRank(key int64) *PlayerArenaRank {
	if this.db.tables.PlayerArenaRank == nil {
		return nil
	}
	tmpRow := PlayerArenaRankRow{row:this.db.tables.PlayerArenaRank}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_arena_record ========== */

// 玩家比武场记录
type PlayerArenaRecord struct {
	Id int64 // 记录ID
	Pid int64 // 玩家ID
	Mode int8 // 记录类型，0无数据，1挑战成功，2挑战失败，3被挑战且成功，4被挑战且失败
	OldRank int32 // 原排位
	NewRank int32 // 新排位
	FightNum int32 // 战力
	TargetPid int64 // 对手玩家ID
	TargetNick string // 对手昵称
	TargetOldRank int32 // 对手原排位
	TargetNewRank int32 // 对手新排位
	TargetFightNum int32 // 对手战力
	RecordTime int64 // 记录时间
}

func (this *PlayerArenaRecord) CObject() *C.PlayerArenaRecord {
	object := C.New_PlayerArenaRecord()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.Mode = C.int8_t(this.Mode)
	object.OldRank = C.int32_t(this.OldRank)
	object.NewRank = C.int32_t(this.NewRank)
	object.FightNum = C.int32_t(this.FightNum)
	object.TargetPid = C.int64_t(this.TargetPid)
	object.TargetNick = C.CString(string(this.TargetNick))
	object.TargetNick_len = C.int(len(this.TargetNick))
	object.TargetOldRank = C.int32_t(this.TargetOldRank)
	object.TargetNewRank = C.int32_t(this.TargetNewRank)
	object.TargetFightNum = C.int32_t(this.TargetFightNum)
	object.RecordTime = C.int64_t(this.RecordTime)
	return object
}

func (this insertOP) PlayerArenaRecord(row *PlayerArenaRecord) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerArenaRecord, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerArenaRecord
	for crow := this.db.tables.PlayerArenaRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_arena_record'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerArenaRecord
	this.db.tables.PlayerArenaRecord = newRow
	this.db.addTransLog(this.db.newPlayerArenaRecordInsertLog(newRow, row))
}

func (this deleteOP) PlayerArenaRecord(row *PlayerArenaRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerArenaRecord
	for crow := this.db.tables.PlayerArenaRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_arena_record'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerArenaRecord = this.db.tables.PlayerArenaRecord.next
	}
	this.db.addTransLog(this.db.newPlayerArenaRecordDeleteLog(old, row))
}

func (this updateOP) PlayerArenaRecord(row *PlayerArenaRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerArenaRecord
	for crow := this.db.tables.PlayerArenaRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_arena_record'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerArenaRecord = newRow
	}
	tmpRow := PlayerArenaRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerArenaRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerArenaRecord(key int64) *PlayerArenaRecord {
	for crow := this.db.tables.PlayerArenaRecord; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerArenaRecordRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerArenaRecord(callback func(*PlayerArenaRecordRow)) {
	row := &PlayerArenaRecordRow{}
	for crow := this.db.tables.PlayerArenaRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerArenaRecord() (rows []interface{}) {
	row := &PlayerArenaRecordRow{}
	for crow := this.db.tables.PlayerArenaRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerArenaRecord", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerArenaRecord", len(rows))
	return rows
}

/* ========== player_awaken_graphic ========== */

// 玩家觉醒图谱
type PlayerAwakenGraphic struct {
	Id int64 // 主键ID
	Pid int64 // 玩家ID
	RoleId int8 // 角色ID
	AttrImpl int16 // 觉醒属性实例ID
	Level int8 // 觉醒等级
}

func (this *PlayerAwakenGraphic) CObject() *C.PlayerAwakenGraphic {
	object := C.New_PlayerAwakenGraphic()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.RoleId = C.int8_t(this.RoleId)
	object.AttrImpl = C.int16_t(this.AttrImpl)
	object.Level = C.int8_t(this.Level)
	return object
}

func (this insertOP) PlayerAwakenGraphic(row *PlayerAwakenGraphic) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerAwakenGraphic, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerAwakenGraphic
	for crow := this.db.tables.PlayerAwakenGraphic; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_awaken_graphic'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerAwakenGraphic
	this.db.tables.PlayerAwakenGraphic = newRow
	this.db.addTransLog(this.db.newPlayerAwakenGraphicInsertLog(newRow, row))
}

func (this deleteOP) PlayerAwakenGraphic(row *PlayerAwakenGraphic) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerAwakenGraphic
	for crow := this.db.tables.PlayerAwakenGraphic; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_awaken_graphic'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerAwakenGraphic = this.db.tables.PlayerAwakenGraphic.next
	}
	this.db.addTransLog(this.db.newPlayerAwakenGraphicDeleteLog(old, row))
}

func (this updateOP) PlayerAwakenGraphic(row *PlayerAwakenGraphic) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerAwakenGraphic
	for crow := this.db.tables.PlayerAwakenGraphic; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_awaken_graphic'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerAwakenGraphic = newRow
	}
	tmpRow := PlayerAwakenGraphicRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerAwakenGraphicUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerAwakenGraphic(key int64) *PlayerAwakenGraphic {
	for crow := this.db.tables.PlayerAwakenGraphic; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerAwakenGraphicRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerAwakenGraphic(callback func(*PlayerAwakenGraphicRow)) {
	row := &PlayerAwakenGraphicRow{}
	for crow := this.db.tables.PlayerAwakenGraphic; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerAwakenGraphic() (rows []interface{}) {
	row := &PlayerAwakenGraphicRow{}
	for crow := this.db.tables.PlayerAwakenGraphic; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerAwakenGraphic", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerAwakenGraphic", len(rows))
	return rows
}

/* ========== player_battle_pet ========== */

// 玩家灵宠数据
type PlayerBattlePet struct {
	Id int64 // ID
	Pid int64 // 玩家ID
	BattlePetId int32 // 灵宠ID(怪物ID）
	Level int32 // 灵宠等级
	Exp int64 // 当前经验
	SkillLevel int32 // 技能等级
}

func (this *PlayerBattlePet) CObject() *C.PlayerBattlePet {
	object := C.New_PlayerBattlePet()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.BattlePetId = C.int32_t(this.BattlePetId)
	object.Level = C.int32_t(this.Level)
	object.Exp = C.int64_t(this.Exp)
	object.SkillLevel = C.int32_t(this.SkillLevel)
	return object
}

func (this insertOP) PlayerBattlePet(row *PlayerBattlePet) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerBattlePet, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerBattlePet
	for crow := this.db.tables.PlayerBattlePet; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_battle_pet'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerBattlePet
	this.db.tables.PlayerBattlePet = newRow
	this.db.addTransLog(this.db.newPlayerBattlePetInsertLog(newRow, row))
}

func (this deleteOP) PlayerBattlePet(row *PlayerBattlePet) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerBattlePet
	for crow := this.db.tables.PlayerBattlePet; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_battle_pet'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerBattlePet = this.db.tables.PlayerBattlePet.next
	}
	this.db.addTransLog(this.db.newPlayerBattlePetDeleteLog(old, row))
}

func (this updateOP) PlayerBattlePet(row *PlayerBattlePet) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerBattlePet
	for crow := this.db.tables.PlayerBattlePet; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_battle_pet'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerBattlePet = newRow
	}
	tmpRow := PlayerBattlePetRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerBattlePetUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerBattlePet(key int64) *PlayerBattlePet {
	for crow := this.db.tables.PlayerBattlePet; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerBattlePetRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerBattlePet(callback func(*PlayerBattlePetRow)) {
	row := &PlayerBattlePetRow{}
	for crow := this.db.tables.PlayerBattlePet; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerBattlePet() (rows []interface{}) {
	row := &PlayerBattlePetRow{}
	for crow := this.db.tables.PlayerBattlePet; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerBattlePet", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerBattlePet", len(rows))
	return rows
}

/* ========== player_battle_pet_grid ========== */

// 
type PlayerBattlePetGrid struct {
	Id int64 // id
	Pid int64 // 玩家ID
	GridId int8 // 灵宠格子ID
	BattlePetId int32 // 灵宠ID
}

func (this *PlayerBattlePetGrid) CObject() *C.PlayerBattlePetGrid {
	object := C.New_PlayerBattlePetGrid()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.GridId = C.int8_t(this.GridId)
	object.BattlePetId = C.int32_t(this.BattlePetId)
	return object
}

func (this insertOP) PlayerBattlePetGrid(row *PlayerBattlePetGrid) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerBattlePetGrid, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerBattlePetGrid
	for crow := this.db.tables.PlayerBattlePetGrid; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_battle_pet_grid'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerBattlePetGrid
	this.db.tables.PlayerBattlePetGrid = newRow
	this.db.addTransLog(this.db.newPlayerBattlePetGridInsertLog(newRow, row))
}

func (this deleteOP) PlayerBattlePetGrid(row *PlayerBattlePetGrid) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerBattlePetGrid
	for crow := this.db.tables.PlayerBattlePetGrid; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_battle_pet_grid'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerBattlePetGrid = this.db.tables.PlayerBattlePetGrid.next
	}
	this.db.addTransLog(this.db.newPlayerBattlePetGridDeleteLog(old, row))
}

func (this updateOP) PlayerBattlePetGrid(row *PlayerBattlePetGrid) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerBattlePetGrid
	for crow := this.db.tables.PlayerBattlePetGrid; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_battle_pet_grid'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerBattlePetGrid = newRow
	}
	tmpRow := PlayerBattlePetGridRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerBattlePetGridUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerBattlePetGrid(key int64) *PlayerBattlePetGrid {
	for crow := this.db.tables.PlayerBattlePetGrid; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerBattlePetGridRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerBattlePetGrid(callback func(*PlayerBattlePetGridRow)) {
	row := &PlayerBattlePetGridRow{}
	for crow := this.db.tables.PlayerBattlePetGrid; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerBattlePetGrid() (rows []interface{}) {
	row := &PlayerBattlePetGridRow{}
	for crow := this.db.tables.PlayerBattlePetGrid; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerBattlePetGrid", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerBattlePetGrid", len(rows))
	return rows
}

/* ========== player_battle_pet_state ========== */

// 玩家灵宠状态
type PlayerBattlePetState struct {
	Pid int64 // 玩家ID
	UpgradeByIngotNum int32 // 今日使用元宝升级次数
	UpgradeByIngotTime int64 // 最近一次使用元宝升级时间
}

func (this *PlayerBattlePetState) CObject() *C.PlayerBattlePetState {
	object := C.New_PlayerBattlePetState()
	object.Pid = C.int64_t(this.Pid)
	object.UpgradeByIngotNum = C.int32_t(this.UpgradeByIngotNum)
	object.UpgradeByIngotTime = C.int64_t(this.UpgradeByIngotTime)
	return object
}

func (this insertOP) PlayerBattlePetState(row *PlayerBattlePetState) {
	if this.db.tables.PlayerBattlePetState != nil {
		panic("duplicate insert 'player_battle_pet_state'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerBattlePetState = newRow
	this.db.addTransLog(this.db.newPlayerBattlePetStateInsertLog(newRow, row))
}

func (this deleteOP) PlayerBattlePetState(row *PlayerBattlePetState) {
	if this.db.tables.PlayerBattlePetState == nil {
		panic("delete not exists 'player_battle_pet_state'")
	}
	old := this.db.tables.PlayerBattlePetState
	this.db.tables.PlayerBattlePetState = nil
	this.db.addTransLog(this.db.newPlayerBattlePetStateDeleteLog(old, row))
}

func (this updateOP) PlayerBattlePetState(row *PlayerBattlePetState) {
	if this.db.tables.PlayerBattlePetState == nil {
		panic("update not exists 'player_battle_pet_state'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerBattlePetState
	this.db.tables.PlayerBattlePetState = newRow
	tmpRow := PlayerBattlePetStateRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerBattlePetStateUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerBattlePetState(key int64) *PlayerBattlePetState {
	if this.db.tables.PlayerBattlePetState == nil {
		return nil
	}
	tmpRow := PlayerBattlePetStateRow{row:this.db.tables.PlayerBattlePetState}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_chest_state ========== */

// 玩家宝箱状态
type PlayerChestState struct {
	Pid int64 // 玩家id
	FreeCoinChestNum int32 // 每日免费青铜宝箱数
	LastFreeCoinChestAt int64 // 上次开免费青铜宝箱时间
	CoinChestNum int32 // 今天开青铜宝箱次数
	CoinChestTenNum int32 // 今日青铜宝箱十连抽次数
	IsFirstCoinTen int8 // 是否第一次青龙宝箱十连抽
	CoinChestConsume int64 // 今天开青铜宝箱花费铜钱数
	LastCoinChestAt int64 // 上次开消费青铜宝箱时间
	LastFreeIngotChestAt int64 // 上次开免费神龙宝箱时间
	IngotChestNum int32 // 今天开神龙宝箱次数
	IngotChestTenNum int32 // 今日神龙宝箱十连抽次数
	IsFirstIngotTen int8 // 是否第一次神龙宝箱十连抽
	IngotChestConsume int64 // 今天开神龙宝箱花费元宝数
	LastIngotChestAt int64 // 上次开消费神龙宝箱时间
	LastFreePetChestAt int64 // 上次开免费灵兽宝箱时间
}

func (this *PlayerChestState) CObject() *C.PlayerChestState {
	object := C.New_PlayerChestState()
	object.Pid = C.int64_t(this.Pid)
	object.FreeCoinChestNum = C.int32_t(this.FreeCoinChestNum)
	object.LastFreeCoinChestAt = C.int64_t(this.LastFreeCoinChestAt)
	object.CoinChestNum = C.int32_t(this.CoinChestNum)
	object.CoinChestTenNum = C.int32_t(this.CoinChestTenNum)
	object.IsFirstCoinTen = C.int8_t(this.IsFirstCoinTen)
	object.CoinChestConsume = C.int64_t(this.CoinChestConsume)
	object.LastCoinChestAt = C.int64_t(this.LastCoinChestAt)
	object.LastFreeIngotChestAt = C.int64_t(this.LastFreeIngotChestAt)
	object.IngotChestNum = C.int32_t(this.IngotChestNum)
	object.IngotChestTenNum = C.int32_t(this.IngotChestTenNum)
	object.IsFirstIngotTen = C.int8_t(this.IsFirstIngotTen)
	object.IngotChestConsume = C.int64_t(this.IngotChestConsume)
	object.LastIngotChestAt = C.int64_t(this.LastIngotChestAt)
	object.LastFreePetChestAt = C.int64_t(this.LastFreePetChestAt)
	return object
}

func (this insertOP) PlayerChestState(row *PlayerChestState) {
	if this.db.tables.PlayerChestState != nil {
		panic("duplicate insert 'player_chest_state'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerChestState = newRow
	this.db.addTransLog(this.db.newPlayerChestStateInsertLog(newRow, row))
}

func (this deleteOP) PlayerChestState(row *PlayerChestState) {
	if this.db.tables.PlayerChestState == nil {
		panic("delete not exists 'player_chest_state'")
	}
	old := this.db.tables.PlayerChestState
	this.db.tables.PlayerChestState = nil
	this.db.addTransLog(this.db.newPlayerChestStateDeleteLog(old, row))
}

func (this updateOP) PlayerChestState(row *PlayerChestState) {
	if this.db.tables.PlayerChestState == nil {
		panic("update not exists 'player_chest_state'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerChestState
	this.db.tables.PlayerChestState = newRow
	tmpRow := PlayerChestStateRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerChestStateUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerChestState(key int64) *PlayerChestState {
	if this.db.tables.PlayerChestState == nil {
		return nil
	}
	tmpRow := PlayerChestStateRow{row:this.db.tables.PlayerChestState}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_clique_kongfu_attrib ========== */

// 玩家帮派武功属性加成
type PlayerCliqueKongfuAttrib struct {
	Pid int64 // 玩家ID
	Health int32 // 生命
	Attack int32 // 攻击
	Defence int32 // 防御
}

func (this *PlayerCliqueKongfuAttrib) CObject() *C.PlayerCliqueKongfuAttrib {
	object := C.New_PlayerCliqueKongfuAttrib()
	object.Pid = C.int64_t(this.Pid)
	object.Health = C.int32_t(this.Health)
	object.Attack = C.int32_t(this.Attack)
	object.Defence = C.int32_t(this.Defence)
	return object
}

func (this insertOP) PlayerCliqueKongfuAttrib(row *PlayerCliqueKongfuAttrib) {
	if this.db.tables.PlayerCliqueKongfuAttrib != nil {
		panic("duplicate insert 'player_clique_kongfu_attrib'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerCliqueKongfuAttrib = newRow
	this.db.addTransLog(this.db.newPlayerCliqueKongfuAttribInsertLog(newRow, row))
}

func (this deleteOP) PlayerCliqueKongfuAttrib(row *PlayerCliqueKongfuAttrib) {
	if this.db.tables.PlayerCliqueKongfuAttrib == nil {
		panic("delete not exists 'player_clique_kongfu_attrib'")
	}
	old := this.db.tables.PlayerCliqueKongfuAttrib
	this.db.tables.PlayerCliqueKongfuAttrib = nil
	this.db.addTransLog(this.db.newPlayerCliqueKongfuAttribDeleteLog(old, row))
}

func (this updateOP) PlayerCliqueKongfuAttrib(row *PlayerCliqueKongfuAttrib) {
	if this.db.tables.PlayerCliqueKongfuAttrib == nil {
		panic("update not exists 'player_clique_kongfu_attrib'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerCliqueKongfuAttrib
	this.db.tables.PlayerCliqueKongfuAttrib = newRow
	tmpRow := PlayerCliqueKongfuAttribRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerCliqueKongfuAttribUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerCliqueKongfuAttrib(key int64) *PlayerCliqueKongfuAttrib {
	if this.db.tables.PlayerCliqueKongfuAttrib == nil {
		return nil
	}
	tmpRow := PlayerCliqueKongfuAttribRow{row:this.db.tables.PlayerCliqueKongfuAttrib}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_coins ========== */

// 玩家铜币兑换表
type PlayerCoins struct {
	Pid int64 // 玩家ID
	BuyTime int64 // 购买更新时间
	DailyCount int16 // 当天购买次数
	BatchBought int16 // 玩家批量购买铜币次数
}

func (this *PlayerCoins) CObject() *C.PlayerCoins {
	object := C.New_PlayerCoins()
	object.Pid = C.int64_t(this.Pid)
	object.BuyTime = C.int64_t(this.BuyTime)
	object.DailyCount = C.int16_t(this.DailyCount)
	object.BatchBought = C.int16_t(this.BatchBought)
	return object
}

func (this insertOP) PlayerCoins(row *PlayerCoins) {
	if this.db.tables.PlayerCoins != nil {
		panic("duplicate insert 'player_coins'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerCoins = newRow
	this.db.addTransLog(this.db.newPlayerCoinsInsertLog(newRow, row))
}

func (this deleteOP) PlayerCoins(row *PlayerCoins) {
	if this.db.tables.PlayerCoins == nil {
		panic("delete not exists 'player_coins'")
	}
	old := this.db.tables.PlayerCoins
	this.db.tables.PlayerCoins = nil
	this.db.addTransLog(this.db.newPlayerCoinsDeleteLog(old, row))
}

func (this updateOP) PlayerCoins(row *PlayerCoins) {
	if this.db.tables.PlayerCoins == nil {
		panic("update not exists 'player_coins'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerCoins
	this.db.tables.PlayerCoins = newRow
	tmpRow := PlayerCoinsRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerCoinsUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerCoins(key int64) *PlayerCoins {
	if this.db.tables.PlayerCoins == nil {
		return nil
	}
	tmpRow := PlayerCoinsRow{row:this.db.tables.PlayerCoins}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_cornucopia ========== */

// 玩家铜币兑换表
type PlayerCornucopia struct {
	Pid int64 // 玩家ID
	OpenTime int64 // 开启更新时间
	DailyCount int16 // 当天开启次数
}

func (this *PlayerCornucopia) CObject() *C.PlayerCornucopia {
	object := C.New_PlayerCornucopia()
	object.Pid = C.int64_t(this.Pid)
	object.OpenTime = C.int64_t(this.OpenTime)
	object.DailyCount = C.int16_t(this.DailyCount)
	return object
}

func (this insertOP) PlayerCornucopia(row *PlayerCornucopia) {
	if this.db.tables.PlayerCornucopia != nil {
		panic("duplicate insert 'player_cornucopia'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerCornucopia = newRow
	this.db.addTransLog(this.db.newPlayerCornucopiaInsertLog(newRow, row))
}

func (this deleteOP) PlayerCornucopia(row *PlayerCornucopia) {
	if this.db.tables.PlayerCornucopia == nil {
		panic("delete not exists 'player_cornucopia'")
	}
	old := this.db.tables.PlayerCornucopia
	this.db.tables.PlayerCornucopia = nil
	this.db.addTransLog(this.db.newPlayerCornucopiaDeleteLog(old, row))
}

func (this updateOP) PlayerCornucopia(row *PlayerCornucopia) {
	if this.db.tables.PlayerCornucopia == nil {
		panic("update not exists 'player_cornucopia'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerCornucopia
	this.db.tables.PlayerCornucopia = newRow
	tmpRow := PlayerCornucopiaRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerCornucopiaUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerCornucopia(key int64) *PlayerCornucopia {
	if this.db.tables.PlayerCornucopia == nil {
		return nil
	}
	tmpRow := PlayerCornucopiaRow{row:this.db.tables.PlayerCornucopia}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_daily_quest ========== */

// 玩家每日任务
type PlayerDailyQuest struct {
	Id int64 // 记录ID
	Pid int64 // 玩家ID
	QuestId int16 // 任务ID
	FinishCount int16 // 完成数量
	LastUpdateTime int64 // 最近一次更新时间
	AwardStatus int8 // 奖励状态; 0 未奖励； 1已奖励
	Class int16 // 每日任务类别
}

func (this *PlayerDailyQuest) CObject() *C.PlayerDailyQuest {
	object := C.New_PlayerDailyQuest()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.QuestId = C.int16_t(this.QuestId)
	object.FinishCount = C.int16_t(this.FinishCount)
	object.LastUpdateTime = C.int64_t(this.LastUpdateTime)
	object.AwardStatus = C.int8_t(this.AwardStatus)
	object.Class = C.int16_t(this.Class)
	return object
}

func (this insertOP) PlayerDailyQuest(row *PlayerDailyQuest) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerDailyQuest, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerDailyQuest
	for crow := this.db.tables.PlayerDailyQuest; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_daily_quest'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerDailyQuest
	this.db.tables.PlayerDailyQuest = newRow
	this.db.addTransLog(this.db.newPlayerDailyQuestInsertLog(newRow, row))
}

func (this deleteOP) PlayerDailyQuest(row *PlayerDailyQuest) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDailyQuest
	for crow := this.db.tables.PlayerDailyQuest; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_daily_quest'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerDailyQuest = this.db.tables.PlayerDailyQuest.next
	}
	this.db.addTransLog(this.db.newPlayerDailyQuestDeleteLog(old, row))
}

func (this updateOP) PlayerDailyQuest(row *PlayerDailyQuest) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDailyQuest
	for crow := this.db.tables.PlayerDailyQuest; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_daily_quest'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerDailyQuest = newRow
	}
	tmpRow := PlayerDailyQuestRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerDailyQuestUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerDailyQuest(key int64) *PlayerDailyQuest {
	for crow := this.db.tables.PlayerDailyQuest; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerDailyQuestRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerDailyQuest(callback func(*PlayerDailyQuestRow)) {
	row := &PlayerDailyQuestRow{}
	for crow := this.db.tables.PlayerDailyQuest; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerDailyQuest() (rows []interface{}) {
	row := &PlayerDailyQuestRow{}
	for crow := this.db.tables.PlayerDailyQuest; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerDailyQuest", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerDailyQuest", len(rows))
	return rows
}

/* ========== player_daily_quest_star_award_info ========== */

// 玩家每日任务星数奖励状态
type PlayerDailyQuestStarAwardInfo struct {
	Pid int64 // 玩家id
	Stars int32 // 当日星数
	Lastupdatetime int64 // 最后修改时间
	Awarded string // 已领阶段
}

func (this *PlayerDailyQuestStarAwardInfo) CObject() *C.PlayerDailyQuestStarAwardInfo {
	object := C.New_PlayerDailyQuestStarAwardInfo()
	object.Pid = C.int64_t(this.Pid)
	object.Stars = C.int32_t(this.Stars)
	object.Lastupdatetime = C.int64_t(this.Lastupdatetime)
	object.Awarded = C.CString(string(this.Awarded))
	object.Awarded_len = C.int(len(this.Awarded))
	return object
}

func (this insertOP) PlayerDailyQuestStarAwardInfo(row *PlayerDailyQuestStarAwardInfo) {
	if this.db.tables.PlayerDailyQuestStarAwardInfo != nil {
		panic("duplicate insert 'player_daily_quest_star_award_info'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerDailyQuestStarAwardInfo = newRow
	this.db.addTransLog(this.db.newPlayerDailyQuestStarAwardInfoInsertLog(newRow, row))
}

func (this deleteOP) PlayerDailyQuestStarAwardInfo(row *PlayerDailyQuestStarAwardInfo) {
	if this.db.tables.PlayerDailyQuestStarAwardInfo == nil {
		panic("delete not exists 'player_daily_quest_star_award_info'")
	}
	old := this.db.tables.PlayerDailyQuestStarAwardInfo
	this.db.tables.PlayerDailyQuestStarAwardInfo = nil
	this.db.addTransLog(this.db.newPlayerDailyQuestStarAwardInfoDeleteLog(old, row))
}

func (this updateOP) PlayerDailyQuestStarAwardInfo(row *PlayerDailyQuestStarAwardInfo) {
	if this.db.tables.PlayerDailyQuestStarAwardInfo == nil {
		panic("update not exists 'player_daily_quest_star_award_info'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerDailyQuestStarAwardInfo
	this.db.tables.PlayerDailyQuestStarAwardInfo = newRow
	tmpRow := PlayerDailyQuestStarAwardInfoRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerDailyQuestStarAwardInfoUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerDailyQuestStarAwardInfo(key int64) *PlayerDailyQuestStarAwardInfo {
	if this.db.tables.PlayerDailyQuestStarAwardInfo == nil {
		return nil
	}
	tmpRow := PlayerDailyQuestStarAwardInfoRow{row:this.db.tables.PlayerDailyQuestStarAwardInfo}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_daily_sign_in_record ========== */

// 玩家每日签到记录
type PlayerDailySignInRecord struct {
	Id int64 // ID
	Pid int64 // 玩家ID
	SignInTime int64 // 签到时间
}

func (this *PlayerDailySignInRecord) CObject() *C.PlayerDailySignInRecord {
	object := C.New_PlayerDailySignInRecord()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.SignInTime = C.int64_t(this.SignInTime)
	return object
}

func (this insertOP) PlayerDailySignInRecord(row *PlayerDailySignInRecord) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerDailySignInRecord, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerDailySignInRecord
	for crow := this.db.tables.PlayerDailySignInRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_daily_sign_in_record'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerDailySignInRecord
	this.db.tables.PlayerDailySignInRecord = newRow
	this.db.addTransLog(this.db.newPlayerDailySignInRecordInsertLog(newRow, row))
}

func (this deleteOP) PlayerDailySignInRecord(row *PlayerDailySignInRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDailySignInRecord
	for crow := this.db.tables.PlayerDailySignInRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_daily_sign_in_record'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerDailySignInRecord = this.db.tables.PlayerDailySignInRecord.next
	}
	this.db.addTransLog(this.db.newPlayerDailySignInRecordDeleteLog(old, row))
}

func (this updateOP) PlayerDailySignInRecord(row *PlayerDailySignInRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDailySignInRecord
	for crow := this.db.tables.PlayerDailySignInRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_daily_sign_in_record'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerDailySignInRecord = newRow
	}
	tmpRow := PlayerDailySignInRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerDailySignInRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerDailySignInRecord(key int64) *PlayerDailySignInRecord {
	for crow := this.db.tables.PlayerDailySignInRecord; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerDailySignInRecordRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerDailySignInRecord(callback func(*PlayerDailySignInRecordRow)) {
	row := &PlayerDailySignInRecordRow{}
	for crow := this.db.tables.PlayerDailySignInRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerDailySignInRecord() (rows []interface{}) {
	row := &PlayerDailySignInRecordRow{}
	for crow := this.db.tables.PlayerDailySignInRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerDailySignInRecord", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerDailySignInRecord", len(rows))
	return rows
}

/* ========== player_daily_sign_in_state ========== */

// 玩家最近七日每日签到状态
type PlayerDailySignInState struct {
	Pid int64 // 玩家ID
	UpdateTime int64 // 最新签到时间
	Record int16 // 签到记录
	SignedToday int8 // 今天是否已签到
}

func (this *PlayerDailySignInState) CObject() *C.PlayerDailySignInState {
	object := C.New_PlayerDailySignInState()
	object.Pid = C.int64_t(this.Pid)
	object.UpdateTime = C.int64_t(this.UpdateTime)
	object.Record = C.int16_t(this.Record)
	object.SignedToday = C.int8_t(this.SignedToday)
	return object
}

func (this insertOP) PlayerDailySignInState(row *PlayerDailySignInState) {
	if this.db.tables.PlayerDailySignInState != nil {
		panic("duplicate insert 'player_daily_sign_in_state'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerDailySignInState = newRow
	this.db.addTransLog(this.db.newPlayerDailySignInStateInsertLog(newRow, row))
}

func (this deleteOP) PlayerDailySignInState(row *PlayerDailySignInState) {
	if this.db.tables.PlayerDailySignInState == nil {
		panic("delete not exists 'player_daily_sign_in_state'")
	}
	old := this.db.tables.PlayerDailySignInState
	this.db.tables.PlayerDailySignInState = nil
	this.db.addTransLog(this.db.newPlayerDailySignInStateDeleteLog(old, row))
}

func (this updateOP) PlayerDailySignInState(row *PlayerDailySignInState) {
	if this.db.tables.PlayerDailySignInState == nil {
		panic("update not exists 'player_daily_sign_in_state'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerDailySignInState
	this.db.tables.PlayerDailySignInState = newRow
	tmpRow := PlayerDailySignInStateRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerDailySignInStateUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerDailySignInState(key int64) *PlayerDailySignInState {
	if this.db.tables.PlayerDailySignInState == nil {
		return nil
	}
	tmpRow := PlayerDailySignInStateRow{row:this.db.tables.PlayerDailySignInState}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_despair_land_camp_state ========== */

// 玩家绝望之地每周进度
type PlayerDespairLandCampState struct {
	Id int64 // 主键标识
	Pid int64 // 玩家id
	CampType int8 // 阵营
	Timestamp int64 // 记录产生时间戳
	Point int64 // 周讨伐点
	KillNum int64 // 周击杀怪物数
	DeadNum int64 // 周角色死亡次数
	DeadNumBoss int64 // 周boss战总角色死亡次数
	Hurt int64 // 对boss伤害
	BossBattleNum int32 // 参与boss战斗次数
	DailyBossBattleNum int32 // 今日参与boss战斗次数
	BossBattleTimestamp int64 // 最近一次boss战斗结束时间戳
	Awarded int8 // 是否领取奖励
}

func (this *PlayerDespairLandCampState) CObject() *C.PlayerDespairLandCampState {
	object := C.New_PlayerDespairLandCampState()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.CampType = C.int8_t(this.CampType)
	object.Timestamp = C.int64_t(this.Timestamp)
	object.Point = C.int64_t(this.Point)
	object.KillNum = C.int64_t(this.KillNum)
	object.DeadNum = C.int64_t(this.DeadNum)
	object.DeadNumBoss = C.int64_t(this.DeadNumBoss)
	object.Hurt = C.int64_t(this.Hurt)
	object.BossBattleNum = C.int32_t(this.BossBattleNum)
	object.DailyBossBattleNum = C.int32_t(this.DailyBossBattleNum)
	object.BossBattleTimestamp = C.int64_t(this.BossBattleTimestamp)
	object.Awarded = C.int8_t(this.Awarded)
	return object
}

func (this insertOP) PlayerDespairLandCampState(row *PlayerDespairLandCampState) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerDespairLandCampState, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerDespairLandCampState
	for crow := this.db.tables.PlayerDespairLandCampState; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_despair_land_camp_state'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerDespairLandCampState
	this.db.tables.PlayerDespairLandCampState = newRow
	this.db.addTransLog(this.db.newPlayerDespairLandCampStateInsertLog(newRow, row))
}

func (this deleteOP) PlayerDespairLandCampState(row *PlayerDespairLandCampState) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDespairLandCampState
	for crow := this.db.tables.PlayerDespairLandCampState; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_despair_land_camp_state'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerDespairLandCampState = this.db.tables.PlayerDespairLandCampState.next
	}
	this.db.addTransLog(this.db.newPlayerDespairLandCampStateDeleteLog(old, row))
}

func (this updateOP) PlayerDespairLandCampState(row *PlayerDespairLandCampState) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDespairLandCampState
	for crow := this.db.tables.PlayerDespairLandCampState; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_despair_land_camp_state'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerDespairLandCampState = newRow
	}
	tmpRow := PlayerDespairLandCampStateRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerDespairLandCampStateUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerDespairLandCampState(key int64) *PlayerDespairLandCampState {
	for crow := this.db.tables.PlayerDespairLandCampState; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerDespairLandCampStateRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerDespairLandCampState(callback func(*PlayerDespairLandCampStateRow)) {
	row := &PlayerDespairLandCampStateRow{}
	for crow := this.db.tables.PlayerDespairLandCampState; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerDespairLandCampState() (rows []interface{}) {
	row := &PlayerDespairLandCampStateRow{}
	for crow := this.db.tables.PlayerDespairLandCampState; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerDespairLandCampState", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerDespairLandCampState", len(rows))
	return rows
}

/* ========== player_despair_land_camp_state_history ========== */

// 玩家绝望之地每周进度
type PlayerDespairLandCampStateHistory struct {
	Id int64 // 主键标识
	Pid int64 // 玩家id
	CampType int8 // 阵营
	Timestamp int64 // 记录产生时间戳
	Point int64 // 周讨伐点
	KillNum int64 // 周击杀怪物数
	DeadNum int64 // 周角色死亡次数
	DeadNumBoss int64 // 周boss战总角色死亡次数
	Hurt int64 // 对boss伤害
	BossBattleNum int32 // 参与boss战斗次数
	DailyBossBattleNum int32 // 今日参与boss战斗次数
	BossBattleTimestamp int64 // 最近一次boss战斗结束时间戳
	Awarded int8 // 是否领取奖励
}

func (this *PlayerDespairLandCampStateHistory) CObject() *C.PlayerDespairLandCampStateHistory {
	object := C.New_PlayerDespairLandCampStateHistory()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.CampType = C.int8_t(this.CampType)
	object.Timestamp = C.int64_t(this.Timestamp)
	object.Point = C.int64_t(this.Point)
	object.KillNum = C.int64_t(this.KillNum)
	object.DeadNum = C.int64_t(this.DeadNum)
	object.DeadNumBoss = C.int64_t(this.DeadNumBoss)
	object.Hurt = C.int64_t(this.Hurt)
	object.BossBattleNum = C.int32_t(this.BossBattleNum)
	object.DailyBossBattleNum = C.int32_t(this.DailyBossBattleNum)
	object.BossBattleTimestamp = C.int64_t(this.BossBattleTimestamp)
	object.Awarded = C.int8_t(this.Awarded)
	return object
}

func (this insertOP) PlayerDespairLandCampStateHistory(row *PlayerDespairLandCampStateHistory) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerDespairLandCampStateHistory, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerDespairLandCampStateHistory
	for crow := this.db.tables.PlayerDespairLandCampStateHistory; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_despair_land_camp_state_history'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerDespairLandCampStateHistory
	this.db.tables.PlayerDespairLandCampStateHistory = newRow
	this.db.addTransLog(this.db.newPlayerDespairLandCampStateHistoryInsertLog(newRow, row))
}

func (this deleteOP) PlayerDespairLandCampStateHistory(row *PlayerDespairLandCampStateHistory) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDespairLandCampStateHistory
	for crow := this.db.tables.PlayerDespairLandCampStateHistory; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_despair_land_camp_state_history'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerDespairLandCampStateHistory = this.db.tables.PlayerDespairLandCampStateHistory.next
	}
	this.db.addTransLog(this.db.newPlayerDespairLandCampStateHistoryDeleteLog(old, row))
}

func (this updateOP) PlayerDespairLandCampStateHistory(row *PlayerDespairLandCampStateHistory) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDespairLandCampStateHistory
	for crow := this.db.tables.PlayerDespairLandCampStateHistory; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_despair_land_camp_state_history'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerDespairLandCampStateHistory = newRow
	}
	tmpRow := PlayerDespairLandCampStateHistoryRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerDespairLandCampStateHistoryUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerDespairLandCampStateHistory(key int64) *PlayerDespairLandCampStateHistory {
	for crow := this.db.tables.PlayerDespairLandCampStateHistory; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerDespairLandCampStateHistoryRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerDespairLandCampStateHistory(callback func(*PlayerDespairLandCampStateHistoryRow)) {
	row := &PlayerDespairLandCampStateHistoryRow{}
	for crow := this.db.tables.PlayerDespairLandCampStateHistory; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerDespairLandCampStateHistory() (rows []interface{}) {
	row := &PlayerDespairLandCampStateHistoryRow{}
	for crow := this.db.tables.PlayerDespairLandCampStateHistory; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerDespairLandCampStateHistory", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerDespairLandCampStateHistory", len(rows))
	return rows
}

/* ========== player_despair_land_level_record ========== */

// 玩家绝望之地状态
type PlayerDespairLandLevelRecord struct {
	Id int64 // 主键ID
	Pid int64 // 玩家id
	CampType int8 // 阵营
	Timestamp int64 // 记录产生时间戳
	LevelId int32 // 关卡ID
	Round int8 // 最小通关回合数
	PassedTimestamp int64 // 最近通关时间
	FirstPassedTimestamp int64 // 首次通关时间
	FirstPassedFightnum int32 // 首次通关战力
	PassedAward int8 // 首次通奖励
	PerfectAward int8 // 三星通奖励
	MilestoneAward int8 // 阶段通奖励
}

func (this *PlayerDespairLandLevelRecord) CObject() *C.PlayerDespairLandLevelRecord {
	object := C.New_PlayerDespairLandLevelRecord()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.CampType = C.int8_t(this.CampType)
	object.Timestamp = C.int64_t(this.Timestamp)
	object.LevelId = C.int32_t(this.LevelId)
	object.Round = C.int8_t(this.Round)
	object.PassedTimestamp = C.int64_t(this.PassedTimestamp)
	object.FirstPassedTimestamp = C.int64_t(this.FirstPassedTimestamp)
	object.FirstPassedFightnum = C.int32_t(this.FirstPassedFightnum)
	object.PassedAward = C.int8_t(this.PassedAward)
	object.PerfectAward = C.int8_t(this.PerfectAward)
	object.MilestoneAward = C.int8_t(this.MilestoneAward)
	return object
}

func (this insertOP) PlayerDespairLandLevelRecord(row *PlayerDespairLandLevelRecord) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerDespairLandLevelRecord, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerDespairLandLevelRecord
	for crow := this.db.tables.PlayerDespairLandLevelRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_despair_land_level_record'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerDespairLandLevelRecord
	this.db.tables.PlayerDespairLandLevelRecord = newRow
	this.db.addTransLog(this.db.newPlayerDespairLandLevelRecordInsertLog(newRow, row))
}

func (this deleteOP) PlayerDespairLandLevelRecord(row *PlayerDespairLandLevelRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDespairLandLevelRecord
	for crow := this.db.tables.PlayerDespairLandLevelRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_despair_land_level_record'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerDespairLandLevelRecord = this.db.tables.PlayerDespairLandLevelRecord.next
	}
	this.db.addTransLog(this.db.newPlayerDespairLandLevelRecordDeleteLog(old, row))
}

func (this updateOP) PlayerDespairLandLevelRecord(row *PlayerDespairLandLevelRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDespairLandLevelRecord
	for crow := this.db.tables.PlayerDespairLandLevelRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_despair_land_level_record'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerDespairLandLevelRecord = newRow
	}
	tmpRow := PlayerDespairLandLevelRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerDespairLandLevelRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerDespairLandLevelRecord(key int64) *PlayerDespairLandLevelRecord {
	for crow := this.db.tables.PlayerDespairLandLevelRecord; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerDespairLandLevelRecordRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerDespairLandLevelRecord(callback func(*PlayerDespairLandLevelRecordRow)) {
	row := &PlayerDespairLandLevelRecordRow{}
	for crow := this.db.tables.PlayerDespairLandLevelRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerDespairLandLevelRecord() (rows []interface{}) {
	row := &PlayerDespairLandLevelRecordRow{}
	for crow := this.db.tables.PlayerDespairLandLevelRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerDespairLandLevelRecord", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerDespairLandLevelRecord", len(rows))
	return rows
}

/* ========== player_despair_land_state ========== */

// 玩家绝望之地状态
type PlayerDespairLandState struct {
	Pid int64 // 玩家id
	Point int64 // 总讨伐点
	KillNum int64 // 总击杀怪物数
	DeadNum int64 // 总阵亡次数
	DailyBattleNum int16 // 今日讨伐次数
	DailyBattleTimestamp int64 // 今日讨伐时间
	DailyBoughtBattleNum int16 // 今日购买讨伐次数
	DailyBoughtTimestamp int64 // 今日购买讨伐次数时间
	DailyBossBoughtTimestamp int64 // 今日购买boss讨伐次数时间
	DailyBossBeastBoughtBattleNum int16 // 今日购买血兽boss讨伐次数
	DailyBossWalkingDeadBoughtBattleNum int16 // 今日购买尸鬼boss讨伐次数
	DailyBossDevilBoughtBattleNum int16 // 今日购买影魔boss讨伐次数
}

func (this *PlayerDespairLandState) CObject() *C.PlayerDespairLandState {
	object := C.New_PlayerDespairLandState()
	object.Pid = C.int64_t(this.Pid)
	object.Point = C.int64_t(this.Point)
	object.KillNum = C.int64_t(this.KillNum)
	object.DeadNum = C.int64_t(this.DeadNum)
	object.DailyBattleNum = C.int16_t(this.DailyBattleNum)
	object.DailyBattleTimestamp = C.int64_t(this.DailyBattleTimestamp)
	object.DailyBoughtBattleNum = C.int16_t(this.DailyBoughtBattleNum)
	object.DailyBoughtTimestamp = C.int64_t(this.DailyBoughtTimestamp)
	object.DailyBossBoughtTimestamp = C.int64_t(this.DailyBossBoughtTimestamp)
	object.DailyBossBeastBoughtBattleNum = C.int16_t(this.DailyBossBeastBoughtBattleNum)
	object.DailyBossWalkingDeadBoughtBattleNum = C.int16_t(this.DailyBossWalkingDeadBoughtBattleNum)
	object.DailyBossDevilBoughtBattleNum = C.int16_t(this.DailyBossDevilBoughtBattleNum)
	return object
}

func (this insertOP) PlayerDespairLandState(row *PlayerDespairLandState) {
	if this.db.tables.PlayerDespairLandState != nil {
		panic("duplicate insert 'player_despair_land_state'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerDespairLandState = newRow
	this.db.addTransLog(this.db.newPlayerDespairLandStateInsertLog(newRow, row))
}

func (this deleteOP) PlayerDespairLandState(row *PlayerDespairLandState) {
	if this.db.tables.PlayerDespairLandState == nil {
		panic("delete not exists 'player_despair_land_state'")
	}
	old := this.db.tables.PlayerDespairLandState
	this.db.tables.PlayerDespairLandState = nil
	this.db.addTransLog(this.db.newPlayerDespairLandStateDeleteLog(old, row))
}

func (this updateOP) PlayerDespairLandState(row *PlayerDespairLandState) {
	if this.db.tables.PlayerDespairLandState == nil {
		panic("update not exists 'player_despair_land_state'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerDespairLandState
	this.db.tables.PlayerDespairLandState = newRow
	tmpRow := PlayerDespairLandStateRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerDespairLandStateUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerDespairLandState(key int64) *PlayerDespairLandState {
	if this.db.tables.PlayerDespairLandState == nil {
		return nil
	}
	tmpRow := PlayerDespairLandStateRow{row:this.db.tables.PlayerDespairLandState}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_driving_sword_event ========== */

// 玩家云海事件列表
type PlayerDrivingSwordEvent struct {
	Id int64 // 主键ID
	Pid int64 // 玩家ID
	CloudId int16 // 云层id
	X int8 // 事件坐标x
	Y int8 // 事件坐标y
	EventType int8 // 事件类型
	DataId int8 // 事件模版数据id
}

func (this *PlayerDrivingSwordEvent) CObject() *C.PlayerDrivingSwordEvent {
	object := C.New_PlayerDrivingSwordEvent()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.CloudId = C.int16_t(this.CloudId)
	object.X = C.int8_t(this.X)
	object.Y = C.int8_t(this.Y)
	object.EventType = C.int8_t(this.EventType)
	object.DataId = C.int8_t(this.DataId)
	return object
}

func (this insertOP) PlayerDrivingSwordEvent(row *PlayerDrivingSwordEvent) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerDrivingSwordEvent, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerDrivingSwordEvent
	for crow := this.db.tables.PlayerDrivingSwordEvent; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_driving_sword_event'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerDrivingSwordEvent
	this.db.tables.PlayerDrivingSwordEvent = newRow
	this.db.addTransLog(this.db.newPlayerDrivingSwordEventInsertLog(newRow, row))
}

func (this deleteOP) PlayerDrivingSwordEvent(row *PlayerDrivingSwordEvent) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDrivingSwordEvent
	for crow := this.db.tables.PlayerDrivingSwordEvent; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_driving_sword_event'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerDrivingSwordEvent = this.db.tables.PlayerDrivingSwordEvent.next
	}
	this.db.addTransLog(this.db.newPlayerDrivingSwordEventDeleteLog(old, row))
}

func (this updateOP) PlayerDrivingSwordEvent(row *PlayerDrivingSwordEvent) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDrivingSwordEvent
	for crow := this.db.tables.PlayerDrivingSwordEvent; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_driving_sword_event'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerDrivingSwordEvent = newRow
	}
	tmpRow := PlayerDrivingSwordEventRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerDrivingSwordEventUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerDrivingSwordEvent(key int64) *PlayerDrivingSwordEvent {
	for crow := this.db.tables.PlayerDrivingSwordEvent; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerDrivingSwordEventRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerDrivingSwordEvent(callback func(*PlayerDrivingSwordEventRow)) {
	row := &PlayerDrivingSwordEventRow{}
	for crow := this.db.tables.PlayerDrivingSwordEvent; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerDrivingSwordEvent() (rows []interface{}) {
	row := &PlayerDrivingSwordEventRow{}
	for crow := this.db.tables.PlayerDrivingSwordEvent; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerDrivingSwordEvent", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerDrivingSwordEvent", len(rows))
	return rows
}

/* ========== player_driving_sword_event_exploring ========== */

// 玩家云海探险事件信息
type PlayerDrivingSwordEventExploring struct {
	Id int64 // 主键ID
	Pid int64 // 玩家ID
	Status int8 // 事件状态
	GarrisonStart int64 // 驻守开始时间
	GarrisonTime int64 // 已驻守时间
	AwardTime int64 // 已领奖时间
	RoleId int8 // 派驻角色id
	CloudId int16 // 云层id
	X int8 // 事件坐标x
	Y int8 // 事件坐标y
	DataId int8 // 事件模版数据id
}

func (this *PlayerDrivingSwordEventExploring) CObject() *C.PlayerDrivingSwordEventExploring {
	object := C.New_PlayerDrivingSwordEventExploring()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.Status = C.int8_t(this.Status)
	object.GarrisonStart = C.int64_t(this.GarrisonStart)
	object.GarrisonTime = C.int64_t(this.GarrisonTime)
	object.AwardTime = C.int64_t(this.AwardTime)
	object.RoleId = C.int8_t(this.RoleId)
	object.CloudId = C.int16_t(this.CloudId)
	object.X = C.int8_t(this.X)
	object.Y = C.int8_t(this.Y)
	object.DataId = C.int8_t(this.DataId)
	return object
}

func (this insertOP) PlayerDrivingSwordEventExploring(row *PlayerDrivingSwordEventExploring) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerDrivingSwordEventExploring, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerDrivingSwordEventExploring
	for crow := this.db.tables.PlayerDrivingSwordEventExploring; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_driving_sword_event_exploring'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerDrivingSwordEventExploring
	this.db.tables.PlayerDrivingSwordEventExploring = newRow
	this.db.addTransLog(this.db.newPlayerDrivingSwordEventExploringInsertLog(newRow, row))
}

func (this deleteOP) PlayerDrivingSwordEventExploring(row *PlayerDrivingSwordEventExploring) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDrivingSwordEventExploring
	for crow := this.db.tables.PlayerDrivingSwordEventExploring; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_driving_sword_event_exploring'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerDrivingSwordEventExploring = this.db.tables.PlayerDrivingSwordEventExploring.next
	}
	this.db.addTransLog(this.db.newPlayerDrivingSwordEventExploringDeleteLog(old, row))
}

func (this updateOP) PlayerDrivingSwordEventExploring(row *PlayerDrivingSwordEventExploring) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDrivingSwordEventExploring
	for crow := this.db.tables.PlayerDrivingSwordEventExploring; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_driving_sword_event_exploring'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerDrivingSwordEventExploring = newRow
	}
	tmpRow := PlayerDrivingSwordEventExploringRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerDrivingSwordEventExploringUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerDrivingSwordEventExploring(key int64) *PlayerDrivingSwordEventExploring {
	for crow := this.db.tables.PlayerDrivingSwordEventExploring; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerDrivingSwordEventExploringRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerDrivingSwordEventExploring(callback func(*PlayerDrivingSwordEventExploringRow)) {
	row := &PlayerDrivingSwordEventExploringRow{}
	for crow := this.db.tables.PlayerDrivingSwordEventExploring; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerDrivingSwordEventExploring() (rows []interface{}) {
	row := &PlayerDrivingSwordEventExploringRow{}
	for crow := this.db.tables.PlayerDrivingSwordEventExploring; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerDrivingSwordEventExploring", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerDrivingSwordEventExploring", len(rows))
	return rows
}

/* ========== player_driving_sword_event_treasure ========== */

// 玩家云海宝藏事件信息
type PlayerDrivingSwordEventTreasure struct {
	Id int64 // 主键ID
	Pid int64 // 玩家ID
	Progress int8 // 开箱进度
	CloudId int16 // 云层id
	X int8 // 事件坐标x
	Y int8 // 事件坐标y
	DataId int8 // 事件模版数据id
}

func (this *PlayerDrivingSwordEventTreasure) CObject() *C.PlayerDrivingSwordEventTreasure {
	object := C.New_PlayerDrivingSwordEventTreasure()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.Progress = C.int8_t(this.Progress)
	object.CloudId = C.int16_t(this.CloudId)
	object.X = C.int8_t(this.X)
	object.Y = C.int8_t(this.Y)
	object.DataId = C.int8_t(this.DataId)
	return object
}

func (this insertOP) PlayerDrivingSwordEventTreasure(row *PlayerDrivingSwordEventTreasure) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerDrivingSwordEventTreasure, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerDrivingSwordEventTreasure
	for crow := this.db.tables.PlayerDrivingSwordEventTreasure; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_driving_sword_event_treasure'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerDrivingSwordEventTreasure
	this.db.tables.PlayerDrivingSwordEventTreasure = newRow
	this.db.addTransLog(this.db.newPlayerDrivingSwordEventTreasureInsertLog(newRow, row))
}

func (this deleteOP) PlayerDrivingSwordEventTreasure(row *PlayerDrivingSwordEventTreasure) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDrivingSwordEventTreasure
	for crow := this.db.tables.PlayerDrivingSwordEventTreasure; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_driving_sword_event_treasure'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerDrivingSwordEventTreasure = this.db.tables.PlayerDrivingSwordEventTreasure.next
	}
	this.db.addTransLog(this.db.newPlayerDrivingSwordEventTreasureDeleteLog(old, row))
}

func (this updateOP) PlayerDrivingSwordEventTreasure(row *PlayerDrivingSwordEventTreasure) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDrivingSwordEventTreasure
	for crow := this.db.tables.PlayerDrivingSwordEventTreasure; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_driving_sword_event_treasure'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerDrivingSwordEventTreasure = newRow
	}
	tmpRow := PlayerDrivingSwordEventTreasureRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerDrivingSwordEventTreasureUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerDrivingSwordEventTreasure(key int64) *PlayerDrivingSwordEventTreasure {
	for crow := this.db.tables.PlayerDrivingSwordEventTreasure; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerDrivingSwordEventTreasureRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerDrivingSwordEventTreasure(callback func(*PlayerDrivingSwordEventTreasureRow)) {
	row := &PlayerDrivingSwordEventTreasureRow{}
	for crow := this.db.tables.PlayerDrivingSwordEventTreasure; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerDrivingSwordEventTreasure() (rows []interface{}) {
	row := &PlayerDrivingSwordEventTreasureRow{}
	for crow := this.db.tables.PlayerDrivingSwordEventTreasure; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerDrivingSwordEventTreasure", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerDrivingSwordEventTreasure", len(rows))
	return rows
}

/* ========== player_driving_sword_event_visiting ========== */

// 玩家云海拜访事件信息
type PlayerDrivingSwordEventVisiting struct {
	Id int64 // 主键ID
	Pid int64 // 玩家ID
	Status int8 // 事件状态,1代表已经通过还没领奖，2代表已领奖
	TargetPid int64 // 拜访的玩家
	TargetSideState []byte // 拜访玩家战斗状态记录
	CloudId int16 // 云层id
	X int8 // 事件坐标x
	Y int8 // 事件坐标y
	DataId int8 // 事件模版数据id
	TargetStatus string // 拜访玩家相关信息记录
}

func (this *PlayerDrivingSwordEventVisiting) CObject() *C.PlayerDrivingSwordEventVisiting {
	object := C.New_PlayerDrivingSwordEventVisiting()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.Status = C.int8_t(this.Status)
	object.TargetPid = C.int64_t(this.TargetPid)
	object.TargetSideState = C.CString(string(this.TargetSideState))
	object.TargetSideState_len = C.int(len(this.TargetSideState))
	object.CloudId = C.int16_t(this.CloudId)
	object.X = C.int8_t(this.X)
	object.Y = C.int8_t(this.Y)
	object.DataId = C.int8_t(this.DataId)
	object.TargetStatus = C.CString(string(this.TargetStatus))
	object.TargetStatus_len = C.int(len(this.TargetStatus))
	return object
}

func (this insertOP) PlayerDrivingSwordEventVisiting(row *PlayerDrivingSwordEventVisiting) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerDrivingSwordEventVisiting, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerDrivingSwordEventVisiting
	for crow := this.db.tables.PlayerDrivingSwordEventVisiting; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_driving_sword_event_visiting'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerDrivingSwordEventVisiting
	this.db.tables.PlayerDrivingSwordEventVisiting = newRow
	this.db.addTransLog(this.db.newPlayerDrivingSwordEventVisitingInsertLog(newRow, row))
}

func (this deleteOP) PlayerDrivingSwordEventVisiting(row *PlayerDrivingSwordEventVisiting) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDrivingSwordEventVisiting
	for crow := this.db.tables.PlayerDrivingSwordEventVisiting; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_driving_sword_event_visiting'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerDrivingSwordEventVisiting = this.db.tables.PlayerDrivingSwordEventVisiting.next
	}
	this.db.addTransLog(this.db.newPlayerDrivingSwordEventVisitingDeleteLog(old, row))
}

func (this updateOP) PlayerDrivingSwordEventVisiting(row *PlayerDrivingSwordEventVisiting) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDrivingSwordEventVisiting
	for crow := this.db.tables.PlayerDrivingSwordEventVisiting; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_driving_sword_event_visiting'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerDrivingSwordEventVisiting = newRow
	}
	tmpRow := PlayerDrivingSwordEventVisitingRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerDrivingSwordEventVisitingUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerDrivingSwordEventVisiting(key int64) *PlayerDrivingSwordEventVisiting {
	for crow := this.db.tables.PlayerDrivingSwordEventVisiting; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerDrivingSwordEventVisitingRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerDrivingSwordEventVisiting(callback func(*PlayerDrivingSwordEventVisitingRow)) {
	row := &PlayerDrivingSwordEventVisitingRow{}
	for crow := this.db.tables.PlayerDrivingSwordEventVisiting; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerDrivingSwordEventVisiting() (rows []interface{}) {
	row := &PlayerDrivingSwordEventVisitingRow{}
	for crow := this.db.tables.PlayerDrivingSwordEventVisiting; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerDrivingSwordEventVisiting", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerDrivingSwordEventVisiting", len(rows))
	return rows
}

/* ========== player_driving_sword_eventmask ========== */

// 玩家云海事件地图
type PlayerDrivingSwordEventmask struct {
	Id int64 // 主键ID
	Pid int64 // 玩家ID
	CloudId int16 // 云层id
	Events []byte // 事件比特阵列
}

func (this *PlayerDrivingSwordEventmask) CObject() *C.PlayerDrivingSwordEventmask {
	object := C.New_PlayerDrivingSwordEventmask()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.CloudId = C.int16_t(this.CloudId)
	object.Events = C.CString(string(this.Events))
	object.Events_len = C.int(len(this.Events))
	return object
}

func (this insertOP) PlayerDrivingSwordEventmask(row *PlayerDrivingSwordEventmask) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerDrivingSwordEventmask, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerDrivingSwordEventmask
	for crow := this.db.tables.PlayerDrivingSwordEventmask; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_driving_sword_eventmask'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerDrivingSwordEventmask
	this.db.tables.PlayerDrivingSwordEventmask = newRow
	this.db.addTransLog(this.db.newPlayerDrivingSwordEventmaskInsertLog(newRow, row))
}

func (this deleteOP) PlayerDrivingSwordEventmask(row *PlayerDrivingSwordEventmask) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDrivingSwordEventmask
	for crow := this.db.tables.PlayerDrivingSwordEventmask; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_driving_sword_eventmask'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerDrivingSwordEventmask = this.db.tables.PlayerDrivingSwordEventmask.next
	}
	this.db.addTransLog(this.db.newPlayerDrivingSwordEventmaskDeleteLog(old, row))
}

func (this updateOP) PlayerDrivingSwordEventmask(row *PlayerDrivingSwordEventmask) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDrivingSwordEventmask
	for crow := this.db.tables.PlayerDrivingSwordEventmask; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_driving_sword_eventmask'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerDrivingSwordEventmask = newRow
	}
	tmpRow := PlayerDrivingSwordEventmaskRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerDrivingSwordEventmaskUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerDrivingSwordEventmask(key int64) *PlayerDrivingSwordEventmask {
	for crow := this.db.tables.PlayerDrivingSwordEventmask; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerDrivingSwordEventmaskRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerDrivingSwordEventmask(callback func(*PlayerDrivingSwordEventmaskRow)) {
	row := &PlayerDrivingSwordEventmaskRow{}
	for crow := this.db.tables.PlayerDrivingSwordEventmask; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerDrivingSwordEventmask() (rows []interface{}) {
	row := &PlayerDrivingSwordEventmaskRow{}
	for crow := this.db.tables.PlayerDrivingSwordEventmask; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerDrivingSwordEventmask", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerDrivingSwordEventmask", len(rows))
	return rows
}

/* ========== player_driving_sword_info ========== */

// 玩家御剑基本数据
type PlayerDrivingSwordInfo struct {
	Pid int64 // 玩家ID
	CurrentCloud int16 // 当前所在云层
	HighestCloud int16 // 最高开启云层
	CurrentX int8 // 当前坐标x
	CurrentY int8 // 当前坐标y
	AllowedAction int16 // 行动点
	ActionRefreshTime int64 // 行动点刷新时间
	ActionBuyTime int64 // 行动点购买更新时间
	DailyActionBought int8 // 行动点当天购买次数
}

func (this *PlayerDrivingSwordInfo) CObject() *C.PlayerDrivingSwordInfo {
	object := C.New_PlayerDrivingSwordInfo()
	object.Pid = C.int64_t(this.Pid)
	object.CurrentCloud = C.int16_t(this.CurrentCloud)
	object.HighestCloud = C.int16_t(this.HighestCloud)
	object.CurrentX = C.int8_t(this.CurrentX)
	object.CurrentY = C.int8_t(this.CurrentY)
	object.AllowedAction = C.int16_t(this.AllowedAction)
	object.ActionRefreshTime = C.int64_t(this.ActionRefreshTime)
	object.ActionBuyTime = C.int64_t(this.ActionBuyTime)
	object.DailyActionBought = C.int8_t(this.DailyActionBought)
	return object
}

func (this insertOP) PlayerDrivingSwordInfo(row *PlayerDrivingSwordInfo) {
	if this.db.tables.PlayerDrivingSwordInfo != nil {
		panic("duplicate insert 'player_driving_sword_info'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerDrivingSwordInfo = newRow
	this.db.addTransLog(this.db.newPlayerDrivingSwordInfoInsertLog(newRow, row))
}

func (this deleteOP) PlayerDrivingSwordInfo(row *PlayerDrivingSwordInfo) {
	if this.db.tables.PlayerDrivingSwordInfo == nil {
		panic("delete not exists 'player_driving_sword_info'")
	}
	old := this.db.tables.PlayerDrivingSwordInfo
	this.db.tables.PlayerDrivingSwordInfo = nil
	this.db.addTransLog(this.db.newPlayerDrivingSwordInfoDeleteLog(old, row))
}

func (this updateOP) PlayerDrivingSwordInfo(row *PlayerDrivingSwordInfo) {
	if this.db.tables.PlayerDrivingSwordInfo == nil {
		panic("update not exists 'player_driving_sword_info'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerDrivingSwordInfo
	this.db.tables.PlayerDrivingSwordInfo = newRow
	tmpRow := PlayerDrivingSwordInfoRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerDrivingSwordInfoUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerDrivingSwordInfo(key int64) *PlayerDrivingSwordInfo {
	if this.db.tables.PlayerDrivingSwordInfo == nil {
		return nil
	}
	tmpRow := PlayerDrivingSwordInfoRow{row:this.db.tables.PlayerDrivingSwordInfo}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_driving_sword_map ========== */

// 玩家云海地图
type PlayerDrivingSwordMap struct {
	Id int64 // 主键ID
	Pid int64 // 玩家ID
	CloudId int16 // 云层id
	Shadows []byte // 阴影比特阵列
	ObstacleCount int8 // 障碍总数
	ExploringCount int8 // 探险总数
	VisitingCount int8 // 拜访总数
	TreasureCount int8 // 宝藏总数
}

func (this *PlayerDrivingSwordMap) CObject() *C.PlayerDrivingSwordMap {
	object := C.New_PlayerDrivingSwordMap()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.CloudId = C.int16_t(this.CloudId)
	object.Shadows = C.CString(string(this.Shadows))
	object.Shadows_len = C.int(len(this.Shadows))
	object.ObstacleCount = C.int8_t(this.ObstacleCount)
	object.ExploringCount = C.int8_t(this.ExploringCount)
	object.VisitingCount = C.int8_t(this.VisitingCount)
	object.TreasureCount = C.int8_t(this.TreasureCount)
	return object
}

func (this insertOP) PlayerDrivingSwordMap(row *PlayerDrivingSwordMap) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerDrivingSwordMap, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerDrivingSwordMap
	for crow := this.db.tables.PlayerDrivingSwordMap; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_driving_sword_map'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerDrivingSwordMap
	this.db.tables.PlayerDrivingSwordMap = newRow
	this.db.addTransLog(this.db.newPlayerDrivingSwordMapInsertLog(newRow, row))
}

func (this deleteOP) PlayerDrivingSwordMap(row *PlayerDrivingSwordMap) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDrivingSwordMap
	for crow := this.db.tables.PlayerDrivingSwordMap; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_driving_sword_map'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerDrivingSwordMap = this.db.tables.PlayerDrivingSwordMap.next
	}
	this.db.addTransLog(this.db.newPlayerDrivingSwordMapDeleteLog(old, row))
}

func (this updateOP) PlayerDrivingSwordMap(row *PlayerDrivingSwordMap) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerDrivingSwordMap
	for crow := this.db.tables.PlayerDrivingSwordMap; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_driving_sword_map'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerDrivingSwordMap = newRow
	}
	tmpRow := PlayerDrivingSwordMapRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerDrivingSwordMapUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerDrivingSwordMap(key int64) *PlayerDrivingSwordMap {
	for crow := this.db.tables.PlayerDrivingSwordMap; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerDrivingSwordMapRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerDrivingSwordMap(callback func(*PlayerDrivingSwordMapRow)) {
	row := &PlayerDrivingSwordMapRow{}
	for crow := this.db.tables.PlayerDrivingSwordMap; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerDrivingSwordMap() (rows []interface{}) {
	row := &PlayerDrivingSwordMapRow{}
	for crow := this.db.tables.PlayerDrivingSwordMap; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerDrivingSwordMap", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerDrivingSwordMap", len(rows))
	return rows
}

/* ========== player_equipment ========== */

// 玩家装备表
type PlayerEquipment struct {
	Id int64 // 主键ID
	Pid int64 // 玩家ID
	RoleId int8 // 角色ID
	WeaponId int64 // 武器的player_item表主键ID
	ClothesId int64 // 战袍的player_item表主键ID
	AccessoriesId int64 // 饰品的player_item表主键ID
	ShoeId int64 // 靴子的player_item表主键ID
}

func (this *PlayerEquipment) CObject() *C.PlayerEquipment {
	object := C.New_PlayerEquipment()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.RoleId = C.int8_t(this.RoleId)
	object.WeaponId = C.int64_t(this.WeaponId)
	object.ClothesId = C.int64_t(this.ClothesId)
	object.AccessoriesId = C.int64_t(this.AccessoriesId)
	object.ShoeId = C.int64_t(this.ShoeId)
	return object
}

func (this insertOP) PlayerEquipment(row *PlayerEquipment) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerEquipment, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerEquipment
	for crow := this.db.tables.PlayerEquipment; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_equipment'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerEquipment
	this.db.tables.PlayerEquipment = newRow
	this.db.addTransLog(this.db.newPlayerEquipmentInsertLog(newRow, row))
}

func (this deleteOP) PlayerEquipment(row *PlayerEquipment) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerEquipment
	for crow := this.db.tables.PlayerEquipment; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_equipment'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerEquipment = this.db.tables.PlayerEquipment.next
	}
	this.db.addTransLog(this.db.newPlayerEquipmentDeleteLog(old, row))
}

func (this updateOP) PlayerEquipment(row *PlayerEquipment) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerEquipment
	for crow := this.db.tables.PlayerEquipment; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_equipment'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerEquipment = newRow
	}
	tmpRow := PlayerEquipmentRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerEquipmentUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerEquipment(key int64) *PlayerEquipment {
	for crow := this.db.tables.PlayerEquipment; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerEquipmentRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerEquipment(callback func(*PlayerEquipmentRow)) {
	row := &PlayerEquipmentRow{}
	for crow := this.db.tables.PlayerEquipment; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerEquipment() (rows []interface{}) {
	row := &PlayerEquipmentRow{}
	for crow := this.db.tables.PlayerEquipment; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerEquipment", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerEquipment", len(rows))
	return rows
}

/* ========== player_event_award_record ========== */

// 玩家活动奖励领取记录
type PlayerEventAwardRecord struct {
	Pid int64 // 用户ID
	RecordBytes []byte // 奖励领取状态
	JsonEventRecord []byte // json模板配置的玩家活动状态
}

func (this *PlayerEventAwardRecord) CObject() *C.PlayerEventAwardRecord {
	object := C.New_PlayerEventAwardRecord()
	object.Pid = C.int64_t(this.Pid)
	object.RecordBytes = C.CString(string(this.RecordBytes))
	object.RecordBytes_len = C.int(len(this.RecordBytes))
	object.JsonEventRecord = C.CString(string(this.JsonEventRecord))
	object.JsonEventRecord_len = C.int(len(this.JsonEventRecord))
	return object
}

func (this insertOP) PlayerEventAwardRecord(row *PlayerEventAwardRecord) {
	if this.db.tables.PlayerEventAwardRecord != nil {
		panic("duplicate insert 'player_event_award_record'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerEventAwardRecord = newRow
	this.db.addTransLog(this.db.newPlayerEventAwardRecordInsertLog(newRow, row))
}

func (this deleteOP) PlayerEventAwardRecord(row *PlayerEventAwardRecord) {
	if this.db.tables.PlayerEventAwardRecord == nil {
		panic("delete not exists 'player_event_award_record'")
	}
	old := this.db.tables.PlayerEventAwardRecord
	this.db.tables.PlayerEventAwardRecord = nil
	this.db.addTransLog(this.db.newPlayerEventAwardRecordDeleteLog(old, row))
}

func (this updateOP) PlayerEventAwardRecord(row *PlayerEventAwardRecord) {
	if this.db.tables.PlayerEventAwardRecord == nil {
		panic("update not exists 'player_event_award_record'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerEventAwardRecord
	this.db.tables.PlayerEventAwardRecord = newRow
	tmpRow := PlayerEventAwardRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerEventAwardRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerEventAwardRecord(key int64) *PlayerEventAwardRecord {
	if this.db.tables.PlayerEventAwardRecord == nil {
		return nil
	}
	tmpRow := PlayerEventAwardRecordRow{row:this.db.tables.PlayerEventAwardRecord}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_event_daily_online ========== */

// 
type PlayerEventDailyOnline struct {
	Pid int64 // 玩家ID
	Timestamp int64 // 记录创建时间戳
	TotalOnlineTime int32 // 今日累计在线时长
	AwardedOnlinetime int32 // 今日已奖励在线时长
}

func (this *PlayerEventDailyOnline) CObject() *C.PlayerEventDailyOnline {
	object := C.New_PlayerEventDailyOnline()
	object.Pid = C.int64_t(this.Pid)
	object.Timestamp = C.int64_t(this.Timestamp)
	object.TotalOnlineTime = C.int32_t(this.TotalOnlineTime)
	object.AwardedOnlinetime = C.int32_t(this.AwardedOnlinetime)
	return object
}

func (this insertOP) PlayerEventDailyOnline(row *PlayerEventDailyOnline) {
	if this.db.tables.PlayerEventDailyOnline != nil {
		panic("duplicate insert 'player_event_daily_online'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerEventDailyOnline = newRow
	this.db.addTransLog(this.db.newPlayerEventDailyOnlineInsertLog(newRow, row))
}

func (this deleteOP) PlayerEventDailyOnline(row *PlayerEventDailyOnline) {
	if this.db.tables.PlayerEventDailyOnline == nil {
		panic("delete not exists 'player_event_daily_online'")
	}
	old := this.db.tables.PlayerEventDailyOnline
	this.db.tables.PlayerEventDailyOnline = nil
	this.db.addTransLog(this.db.newPlayerEventDailyOnlineDeleteLog(old, row))
}

func (this updateOP) PlayerEventDailyOnline(row *PlayerEventDailyOnline) {
	if this.db.tables.PlayerEventDailyOnline == nil {
		panic("update not exists 'player_event_daily_online'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerEventDailyOnline
	this.db.tables.PlayerEventDailyOnline = newRow
	tmpRow := PlayerEventDailyOnlineRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerEventDailyOnlineUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerEventDailyOnline(key int64) *PlayerEventDailyOnline {
	if this.db.tables.PlayerEventDailyOnline == nil {
		return nil
	}
	tmpRow := PlayerEventDailyOnlineRow{row:this.db.tables.PlayerEventDailyOnline}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_event_vn_daily_recharge ========== */

// 
type PlayerEventVnDailyRecharge struct {
	Id int64 // 
	Pid int64 // 玩家ID
	Page int32 // 玩家ID
	Timestamp int64 // 记录创建时间戳
	Awarded int8 // 今日是否已领取奖励
	DailyRecharge int64 // 今日累计充值
	TotalRecharge int64 // 活动累计充值
}

func (this *PlayerEventVnDailyRecharge) CObject() *C.PlayerEventVnDailyRecharge {
	object := C.New_PlayerEventVnDailyRecharge()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.Page = C.int32_t(this.Page)
	object.Timestamp = C.int64_t(this.Timestamp)
	object.Awarded = C.int8_t(this.Awarded)
	object.DailyRecharge = C.int64_t(this.DailyRecharge)
	object.TotalRecharge = C.int64_t(this.TotalRecharge)
	return object
}

func (this insertOP) PlayerEventVnDailyRecharge(row *PlayerEventVnDailyRecharge) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerEventVnDailyRecharge, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerEventVnDailyRecharge
	for crow := this.db.tables.PlayerEventVnDailyRecharge; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_event_vn_daily_recharge'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerEventVnDailyRecharge
	this.db.tables.PlayerEventVnDailyRecharge = newRow
	this.db.addTransLog(this.db.newPlayerEventVnDailyRechargeInsertLog(newRow, row))
}

func (this deleteOP) PlayerEventVnDailyRecharge(row *PlayerEventVnDailyRecharge) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerEventVnDailyRecharge
	for crow := this.db.tables.PlayerEventVnDailyRecharge; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_event_vn_daily_recharge'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerEventVnDailyRecharge = this.db.tables.PlayerEventVnDailyRecharge.next
	}
	this.db.addTransLog(this.db.newPlayerEventVnDailyRechargeDeleteLog(old, row))
}

func (this updateOP) PlayerEventVnDailyRecharge(row *PlayerEventVnDailyRecharge) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerEventVnDailyRecharge
	for crow := this.db.tables.PlayerEventVnDailyRecharge; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_event_vn_daily_recharge'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerEventVnDailyRecharge = newRow
	}
	tmpRow := PlayerEventVnDailyRechargeRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerEventVnDailyRechargeUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerEventVnDailyRecharge(key int64) *PlayerEventVnDailyRecharge {
	for crow := this.db.tables.PlayerEventVnDailyRecharge; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerEventVnDailyRechargeRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerEventVnDailyRecharge(callback func(*PlayerEventVnDailyRechargeRow)) {
	row := &PlayerEventVnDailyRechargeRow{}
	for crow := this.db.tables.PlayerEventVnDailyRecharge; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerEventVnDailyRecharge() (rows []interface{}) {
	row := &PlayerEventVnDailyRechargeRow{}
	for crow := this.db.tables.PlayerEventVnDailyRecharge; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerEventVnDailyRecharge", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerEventVnDailyRecharge", len(rows))
	return rows
}

/* ========== player_event_vn_dragon_ball_record ========== */

// 越南龙珠活动玩家获取记录
type PlayerEventVnDragonBallRecord struct {
	Pid int64 // 玩家ID
	ItemId int16 // 获取的物品ID
	Createtime int64 // 获取时间戳
}

func (this *PlayerEventVnDragonBallRecord) CObject() *C.PlayerEventVnDragonBallRecord {
	object := C.New_PlayerEventVnDragonBallRecord()
	object.Pid = C.int64_t(this.Pid)
	object.ItemId = C.int16_t(this.ItemId)
	object.Createtime = C.int64_t(this.Createtime)
	return object
}

func (this insertOP) PlayerEventVnDragonBallRecord(row *PlayerEventVnDragonBallRecord) {
	if this.db.tables.PlayerEventVnDragonBallRecord != nil {
		panic("duplicate insert 'player_event_vn_dragon_ball_record'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerEventVnDragonBallRecord = newRow
	this.db.addTransLog(this.db.newPlayerEventVnDragonBallRecordInsertLog(newRow, row))
}

func (this deleteOP) PlayerEventVnDragonBallRecord(row *PlayerEventVnDragonBallRecord) {
	if this.db.tables.PlayerEventVnDragonBallRecord == nil {
		panic("delete not exists 'player_event_vn_dragon_ball_record'")
	}
	old := this.db.tables.PlayerEventVnDragonBallRecord
	this.db.tables.PlayerEventVnDragonBallRecord = nil
	this.db.addTransLog(this.db.newPlayerEventVnDragonBallRecordDeleteLog(old, row))
}

func (this updateOP) PlayerEventVnDragonBallRecord(row *PlayerEventVnDragonBallRecord) {
	if this.db.tables.PlayerEventVnDragonBallRecord == nil {
		panic("update not exists 'player_event_vn_dragon_ball_record'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerEventVnDragonBallRecord
	this.db.tables.PlayerEventVnDragonBallRecord = newRow
	tmpRow := PlayerEventVnDragonBallRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerEventVnDragonBallRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerEventVnDragonBallRecord(key int64) *PlayerEventVnDragonBallRecord {
	if this.db.tables.PlayerEventVnDragonBallRecord == nil {
		return nil
	}
	tmpRow := PlayerEventVnDragonBallRecordRow{row:this.db.tables.PlayerEventVnDragonBallRecord}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_events_fight_power ========== */

// 玩家战力运营活动记录
type PlayerEventsFightPower struct {
	Pid int64 // 玩家ID
	Lock int32 // 当前已奖励的战力档位
}

func (this *PlayerEventsFightPower) CObject() *C.PlayerEventsFightPower {
	object := C.New_PlayerEventsFightPower()
	object.Pid = C.int64_t(this.Pid)
	object.Lock = C.int32_t(this.Lock)
	return object
}

func (this insertOP) PlayerEventsFightPower(row *PlayerEventsFightPower) {
	if this.db.tables.PlayerEventsFightPower != nil {
		panic("duplicate insert 'player_events_fight_power'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerEventsFightPower = newRow
	this.db.addTransLog(this.db.newPlayerEventsFightPowerInsertLog(newRow, row))
}

func (this deleteOP) PlayerEventsFightPower(row *PlayerEventsFightPower) {
	if this.db.tables.PlayerEventsFightPower == nil {
		panic("delete not exists 'player_events_fight_power'")
	}
	old := this.db.tables.PlayerEventsFightPower
	this.db.tables.PlayerEventsFightPower = nil
	this.db.addTransLog(this.db.newPlayerEventsFightPowerDeleteLog(old, row))
}

func (this updateOP) PlayerEventsFightPower(row *PlayerEventsFightPower) {
	if this.db.tables.PlayerEventsFightPower == nil {
		panic("update not exists 'player_events_fight_power'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerEventsFightPower
	this.db.tables.PlayerEventsFightPower = newRow
	tmpRow := PlayerEventsFightPowerRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerEventsFightPowerUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerEventsFightPower(key int64) *PlayerEventsFightPower {
	if this.db.tables.PlayerEventsFightPower == nil {
		return nil
	}
	tmpRow := PlayerEventsFightPowerRow{row:this.db.tables.PlayerEventsFightPower}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_events_ingot_record ========== */

// 玩家元宝充值和消耗活动记录
type PlayerEventsIngotRecord struct {
	Pid int64 // 
	IngotIn int64 // 充值元宝总数
	IngotInEndTime int64 // 累计充值活动结束时间戳
	IngotOut int64 // 消耗元宝总数
	IngotOutEndTime int64 // 消耗元宝活动结束时间戳
}

func (this *PlayerEventsIngotRecord) CObject() *C.PlayerEventsIngotRecord {
	object := C.New_PlayerEventsIngotRecord()
	object.Pid = C.int64_t(this.Pid)
	object.IngotIn = C.int64_t(this.IngotIn)
	object.IngotInEndTime = C.int64_t(this.IngotInEndTime)
	object.IngotOut = C.int64_t(this.IngotOut)
	object.IngotOutEndTime = C.int64_t(this.IngotOutEndTime)
	return object
}

func (this insertOP) PlayerEventsIngotRecord(row *PlayerEventsIngotRecord) {
	if this.db.tables.PlayerEventsIngotRecord != nil {
		panic("duplicate insert 'player_events_ingot_record'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerEventsIngotRecord = newRow
	this.db.addTransLog(this.db.newPlayerEventsIngotRecordInsertLog(newRow, row))
}

func (this deleteOP) PlayerEventsIngotRecord(row *PlayerEventsIngotRecord) {
	if this.db.tables.PlayerEventsIngotRecord == nil {
		panic("delete not exists 'player_events_ingot_record'")
	}
	old := this.db.tables.PlayerEventsIngotRecord
	this.db.tables.PlayerEventsIngotRecord = nil
	this.db.addTransLog(this.db.newPlayerEventsIngotRecordDeleteLog(old, row))
}

func (this updateOP) PlayerEventsIngotRecord(row *PlayerEventsIngotRecord) {
	if this.db.tables.PlayerEventsIngotRecord == nil {
		panic("update not exists 'player_events_ingot_record'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerEventsIngotRecord
	this.db.tables.PlayerEventsIngotRecord = newRow
	tmpRow := PlayerEventsIngotRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerEventsIngotRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerEventsIngotRecord(key int64) *PlayerEventsIngotRecord {
	if this.db.tables.PlayerEventsIngotRecord == nil {
		return nil
	}
	tmpRow := PlayerEventsIngotRecordRow{row:this.db.tables.PlayerEventsIngotRecord}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_extend_level ========== */

// 玩家活动关卡状态
type PlayerExtendLevel struct {
	Pid int64 // 玩家ID
	CoinPassTime int64 // 铜钱关卡通关时间
	ExpPassTime int64 // 经验关卡通关时间
	GhostPassTime int64 // 魂侍关卡通关时间
	PetPassTime int64 // 灵宠关卡通关时间
	BuddyPassTime int64 // 伙伴关卡通关时间
	CoinDailyNum int8 // 经验关卡今日进入次数
	ExpDailyNum int8 // 铜钱关卡今日进入次数
	BuddyDailyNum int8 // 伙伴关卡今日进入次数
	PetDailyNum int8 // 灵宠关卡今日进入次数
	GhostDailyNum int8 // 魂侍关卡今日进入次数
	RandBuddyRoleId int8 // 随机的伙伴角色ID
	BuddyPos int8 // 随机的伙伴角色位置
	BuddyTactical int8 // 伙伴关卡队伍战术
	RolePos int8 // 主角站位
	ExpMaxlevel int16 // 经验关卡通关了的最大等级
	CoinsMaxlevel int16 // 经验关卡通关了的最大等级
	CoinsBuyNum int16 // 元宝购买铜钱关卡次数
	ExpBuyNum int16 // 元宝购买经验关卡次数
	CoinsBuyTime int64 // 元宝购买铜钱关卡时间
	ExpBuyTime int64 // 元宝购买经验关卡时间
}

func (this *PlayerExtendLevel) CObject() *C.PlayerExtendLevel {
	object := C.New_PlayerExtendLevel()
	object.Pid = C.int64_t(this.Pid)
	object.CoinPassTime = C.int64_t(this.CoinPassTime)
	object.ExpPassTime = C.int64_t(this.ExpPassTime)
	object.GhostPassTime = C.int64_t(this.GhostPassTime)
	object.PetPassTime = C.int64_t(this.PetPassTime)
	object.BuddyPassTime = C.int64_t(this.BuddyPassTime)
	object.CoinDailyNum = C.int8_t(this.CoinDailyNum)
	object.ExpDailyNum = C.int8_t(this.ExpDailyNum)
	object.BuddyDailyNum = C.int8_t(this.BuddyDailyNum)
	object.PetDailyNum = C.int8_t(this.PetDailyNum)
	object.GhostDailyNum = C.int8_t(this.GhostDailyNum)
	object.RandBuddyRoleId = C.int8_t(this.RandBuddyRoleId)
	object.BuddyPos = C.int8_t(this.BuddyPos)
	object.BuddyTactical = C.int8_t(this.BuddyTactical)
	object.RolePos = C.int8_t(this.RolePos)
	object.ExpMaxlevel = C.int16_t(this.ExpMaxlevel)
	object.CoinsMaxlevel = C.int16_t(this.CoinsMaxlevel)
	object.CoinsBuyNum = C.int16_t(this.CoinsBuyNum)
	object.ExpBuyNum = C.int16_t(this.ExpBuyNum)
	object.CoinsBuyTime = C.int64_t(this.CoinsBuyTime)
	object.ExpBuyTime = C.int64_t(this.ExpBuyTime)
	return object
}

func (this insertOP) PlayerExtendLevel(row *PlayerExtendLevel) {
	if this.db.tables.PlayerExtendLevel != nil {
		panic("duplicate insert 'player_extend_level'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerExtendLevel = newRow
	this.db.addTransLog(this.db.newPlayerExtendLevelInsertLog(newRow, row))
}

func (this deleteOP) PlayerExtendLevel(row *PlayerExtendLevel) {
	if this.db.tables.PlayerExtendLevel == nil {
		panic("delete not exists 'player_extend_level'")
	}
	old := this.db.tables.PlayerExtendLevel
	this.db.tables.PlayerExtendLevel = nil
	this.db.addTransLog(this.db.newPlayerExtendLevelDeleteLog(old, row))
}

func (this updateOP) PlayerExtendLevel(row *PlayerExtendLevel) {
	if this.db.tables.PlayerExtendLevel == nil {
		panic("update not exists 'player_extend_level'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerExtendLevel
	this.db.tables.PlayerExtendLevel = newRow
	tmpRow := PlayerExtendLevelRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerExtendLevelUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerExtendLevel(key int64) *PlayerExtendLevel {
	if this.db.tables.PlayerExtendLevel == nil {
		return nil
	}
	tmpRow := PlayerExtendLevelRow{row:this.db.tables.PlayerExtendLevel}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_extend_quest ========== */

// 伙伴任务
type PlayerExtendQuest struct {
	Id int64 // ID
	Pid int64 // 用户ID
	QuestId int32 // 任务ID
	Progress int16 // 任务进度
	State int8 // 0--未完成 1--已完成 2--已奖励
}

func (this *PlayerExtendQuest) CObject() *C.PlayerExtendQuest {
	object := C.New_PlayerExtendQuest()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.QuestId = C.int32_t(this.QuestId)
	object.Progress = C.int16_t(this.Progress)
	object.State = C.int8_t(this.State)
	return object
}

func (this insertOP) PlayerExtendQuest(row *PlayerExtendQuest) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerExtendQuest, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerExtendQuest
	for crow := this.db.tables.PlayerExtendQuest; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_extend_quest'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerExtendQuest
	this.db.tables.PlayerExtendQuest = newRow
	this.db.addTransLog(this.db.newPlayerExtendQuestInsertLog(newRow, row))
}

func (this deleteOP) PlayerExtendQuest(row *PlayerExtendQuest) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerExtendQuest
	for crow := this.db.tables.PlayerExtendQuest; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_extend_quest'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerExtendQuest = this.db.tables.PlayerExtendQuest.next
	}
	this.db.addTransLog(this.db.newPlayerExtendQuestDeleteLog(old, row))
}

func (this updateOP) PlayerExtendQuest(row *PlayerExtendQuest) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerExtendQuest
	for crow := this.db.tables.PlayerExtendQuest; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_extend_quest'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerExtendQuest = newRow
	}
	tmpRow := PlayerExtendQuestRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerExtendQuestUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerExtendQuest(key int64) *PlayerExtendQuest {
	for crow := this.db.tables.PlayerExtendQuest; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerExtendQuestRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerExtendQuest(callback func(*PlayerExtendQuestRow)) {
	row := &PlayerExtendQuestRow{}
	for crow := this.db.tables.PlayerExtendQuest; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerExtendQuest() (rows []interface{}) {
	row := &PlayerExtendQuestRow{}
	for crow := this.db.tables.PlayerExtendQuest; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerExtendQuest", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerExtendQuest", len(rows))
	return rows
}

/* ========== player_fame ========== */

// 
type PlayerFame struct {
	Pid int64 // 玩家ID
	Fame int32 // 总声望
	Level int16 // 声望等级
	ArenaFame int32 // 比武场关卡声望
	MultLevelFame int32 // 多人关卡声望
}

func (this *PlayerFame) CObject() *C.PlayerFame {
	object := C.New_PlayerFame()
	object.Pid = C.int64_t(this.Pid)
	object.Fame = C.int32_t(this.Fame)
	object.Level = C.int16_t(this.Level)
	object.ArenaFame = C.int32_t(this.ArenaFame)
	object.MultLevelFame = C.int32_t(this.MultLevelFame)
	return object
}

func (this insertOP) PlayerFame(row *PlayerFame) {
	if this.db.tables.PlayerFame != nil {
		panic("duplicate insert 'player_fame'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerFame = newRow
	this.db.addTransLog(this.db.newPlayerFameInsertLog(newRow, row))
}

func (this deleteOP) PlayerFame(row *PlayerFame) {
	if this.db.tables.PlayerFame == nil {
		panic("delete not exists 'player_fame'")
	}
	old := this.db.tables.PlayerFame
	this.db.tables.PlayerFame = nil
	this.db.addTransLog(this.db.newPlayerFameDeleteLog(old, row))
}

func (this updateOP) PlayerFame(row *PlayerFame) {
	if this.db.tables.PlayerFame == nil {
		panic("update not exists 'player_fame'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerFame
	this.db.tables.PlayerFame = newRow
	tmpRow := PlayerFameRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerFameUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerFame(key int64) *PlayerFame {
	if this.db.tables.PlayerFame == nil {
		return nil
	}
	tmpRow := PlayerFameRow{row:this.db.tables.PlayerFame}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_fashion ========== */

// 玩家时装
type PlayerFashion struct {
	Id int64 // ID
	Pid int64 // 玩家ID
	FashionId int16 // 时装模版ID
	ExpireTime int64 // 过期时间 0--永远有效
}

func (this *PlayerFashion) CObject() *C.PlayerFashion {
	object := C.New_PlayerFashion()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.FashionId = C.int16_t(this.FashionId)
	object.ExpireTime = C.int64_t(this.ExpireTime)
	return object
}

func (this insertOP) PlayerFashion(row *PlayerFashion) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerFashion, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerFashion
	for crow := this.db.tables.PlayerFashion; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_fashion'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerFashion
	this.db.tables.PlayerFashion = newRow
	this.db.addTransLog(this.db.newPlayerFashionInsertLog(newRow, row))
}

func (this deleteOP) PlayerFashion(row *PlayerFashion) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerFashion
	for crow := this.db.tables.PlayerFashion; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_fashion'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerFashion = this.db.tables.PlayerFashion.next
	}
	this.db.addTransLog(this.db.newPlayerFashionDeleteLog(old, row))
}

func (this updateOP) PlayerFashion(row *PlayerFashion) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerFashion
	for crow := this.db.tables.PlayerFashion; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_fashion'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerFashion = newRow
	}
	tmpRow := PlayerFashionRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerFashionUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerFashion(key int64) *PlayerFashion {
	for crow := this.db.tables.PlayerFashion; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerFashionRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerFashion(callback func(*PlayerFashionRow)) {
	row := &PlayerFashionRow{}
	for crow := this.db.tables.PlayerFashion; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerFashion() (rows []interface{}) {
	row := &PlayerFashionRow{}
	for crow := this.db.tables.PlayerFashion; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerFashion", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerFashion", len(rows))
	return rows
}

/* ========== player_fashion_state ========== */

// 玩家时装状态
type PlayerFashionState struct {
	Pid int64 // 玩家ID
	UpdateTime int64 // 状态更新时间
	DressedFashionId int16 // 当前装备的时装
}

func (this *PlayerFashionState) CObject() *C.PlayerFashionState {
	object := C.New_PlayerFashionState()
	object.Pid = C.int64_t(this.Pid)
	object.UpdateTime = C.int64_t(this.UpdateTime)
	object.DressedFashionId = C.int16_t(this.DressedFashionId)
	return object
}

func (this insertOP) PlayerFashionState(row *PlayerFashionState) {
	if this.db.tables.PlayerFashionState != nil {
		panic("duplicate insert 'player_fashion_state'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerFashionState = newRow
	this.db.addTransLog(this.db.newPlayerFashionStateInsertLog(newRow, row))
}

func (this deleteOP) PlayerFashionState(row *PlayerFashionState) {
	if this.db.tables.PlayerFashionState == nil {
		panic("delete not exists 'player_fashion_state'")
	}
	old := this.db.tables.PlayerFashionState
	this.db.tables.PlayerFashionState = nil
	this.db.addTransLog(this.db.newPlayerFashionStateDeleteLog(old, row))
}

func (this updateOP) PlayerFashionState(row *PlayerFashionState) {
	if this.db.tables.PlayerFashionState == nil {
		panic("update not exists 'player_fashion_state'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerFashionState
	this.db.tables.PlayerFashionState = newRow
	tmpRow := PlayerFashionStateRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerFashionStateUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerFashionState(key int64) *PlayerFashionState {
	if this.db.tables.PlayerFashionState == nil {
		return nil
	}
	tmpRow := PlayerFashionStateRow{row:this.db.tables.PlayerFashionState}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_fate_box_state ========== */

// 玩家命锁宝箱状态
type PlayerFateBoxState struct {
	Pid int64 // pid
	Lock int32 // 玩家命锁权值
	StarBoxFreeDrawTimestamp int64 // 星辉命匣免费抽取时间戳
	StarBoxDrawCount int32 // 星辉宝箱抽奖次数
	MoonBoxFreeDrawTimestamp int64 // 月影命匣免费抽取时间戳
	MoonBoxDrawCount int32 // 月影宝箱抽奖次数
	SunBoxFreeDrawTimestamp int64 // 日耀命匣免费抽取时间戳
	SunBoxDrawCount int32 // 日耀宝箱抽奖次数
	HunyuanBoxFreeDrawTimestamp int64 // 混元命匣免费抽取时间戳
	HunyuanBoxDrawCount int32 // 混元宝箱抽奖次数
}

func (this *PlayerFateBoxState) CObject() *C.PlayerFateBoxState {
	object := C.New_PlayerFateBoxState()
	object.Pid = C.int64_t(this.Pid)
	object.Lock = C.int32_t(this.Lock)
	object.StarBoxFreeDrawTimestamp = C.int64_t(this.StarBoxFreeDrawTimestamp)
	object.StarBoxDrawCount = C.int32_t(this.StarBoxDrawCount)
	object.MoonBoxFreeDrawTimestamp = C.int64_t(this.MoonBoxFreeDrawTimestamp)
	object.MoonBoxDrawCount = C.int32_t(this.MoonBoxDrawCount)
	object.SunBoxFreeDrawTimestamp = C.int64_t(this.SunBoxFreeDrawTimestamp)
	object.SunBoxDrawCount = C.int32_t(this.SunBoxDrawCount)
	object.HunyuanBoxFreeDrawTimestamp = C.int64_t(this.HunyuanBoxFreeDrawTimestamp)
	object.HunyuanBoxDrawCount = C.int32_t(this.HunyuanBoxDrawCount)
	return object
}

func (this insertOP) PlayerFateBoxState(row *PlayerFateBoxState) {
	if this.db.tables.PlayerFateBoxState != nil {
		panic("duplicate insert 'player_fate_box_state'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerFateBoxState = newRow
	this.db.addTransLog(this.db.newPlayerFateBoxStateInsertLog(newRow, row))
}

func (this deleteOP) PlayerFateBoxState(row *PlayerFateBoxState) {
	if this.db.tables.PlayerFateBoxState == nil {
		panic("delete not exists 'player_fate_box_state'")
	}
	old := this.db.tables.PlayerFateBoxState
	this.db.tables.PlayerFateBoxState = nil
	this.db.addTransLog(this.db.newPlayerFateBoxStateDeleteLog(old, row))
}

func (this updateOP) PlayerFateBoxState(row *PlayerFateBoxState) {
	if this.db.tables.PlayerFateBoxState == nil {
		panic("update not exists 'player_fate_box_state'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerFateBoxState
	this.db.tables.PlayerFateBoxState = newRow
	tmpRow := PlayerFateBoxStateRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerFateBoxStateUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerFateBoxState(key int64) *PlayerFateBoxState {
	if this.db.tables.PlayerFateBoxState == nil {
		return nil
	}
	tmpRow := PlayerFateBoxStateRow{row:this.db.tables.PlayerFateBoxState}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_fight_num ========== */

// 玩家战斗力
type PlayerFightNum struct {
	Pid int64 // 玩家ID
	FightNum int32 // 战力力
}

func (this *PlayerFightNum) CObject() *C.PlayerFightNum {
	object := C.New_PlayerFightNum()
	object.Pid = C.int64_t(this.Pid)
	object.FightNum = C.int32_t(this.FightNum)
	return object
}

func (this insertOP) PlayerFightNum(row *PlayerFightNum) {
	if this.db.tables.PlayerFightNum != nil {
		panic("duplicate insert 'player_fight_num'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerFightNum = newRow
	this.db.addTransLog(this.db.newPlayerFightNumInsertLog(newRow, row))
}

func (this deleteOP) PlayerFightNum(row *PlayerFightNum) {
	if this.db.tables.PlayerFightNum == nil {
		panic("delete not exists 'player_fight_num'")
	}
	old := this.db.tables.PlayerFightNum
	this.db.tables.PlayerFightNum = nil
	this.db.addTransLog(this.db.newPlayerFightNumDeleteLog(old, row))
}

func (this updateOP) PlayerFightNum(row *PlayerFightNum) {
	if this.db.tables.PlayerFightNum == nil {
		panic("update not exists 'player_fight_num'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerFightNum
	this.db.tables.PlayerFightNum = newRow
	tmpRow := PlayerFightNumRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerFightNumUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerFightNum(key int64) *PlayerFightNum {
	if this.db.tables.PlayerFightNum == nil {
		return nil
	}
	tmpRow := PlayerFightNumRow{row:this.db.tables.PlayerFightNum}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_first_charge ========== */

// 玩家首冲表
type PlayerFirstCharge struct {
	Id int64 // 
	Pid int64 // 玩家ID
	ProductId string // 
}

func (this *PlayerFirstCharge) CObject() *C.PlayerFirstCharge {
	object := C.New_PlayerFirstCharge()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.ProductId = C.CString(string(this.ProductId))
	object.ProductId_len = C.int(len(this.ProductId))
	return object
}

func (this insertOP) PlayerFirstCharge(row *PlayerFirstCharge) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerFirstCharge, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerFirstCharge
	for crow := this.db.tables.PlayerFirstCharge; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_first_charge'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerFirstCharge
	this.db.tables.PlayerFirstCharge = newRow
	this.db.addTransLog(this.db.newPlayerFirstChargeInsertLog(newRow, row))
}

func (this deleteOP) PlayerFirstCharge(row *PlayerFirstCharge) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerFirstCharge
	for crow := this.db.tables.PlayerFirstCharge; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_first_charge'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerFirstCharge = this.db.tables.PlayerFirstCharge.next
	}
	this.db.addTransLog(this.db.newPlayerFirstChargeDeleteLog(old, row))
}

func (this updateOP) PlayerFirstCharge(row *PlayerFirstCharge) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerFirstCharge
	for crow := this.db.tables.PlayerFirstCharge; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_first_charge'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerFirstCharge = newRow
	}
	tmpRow := PlayerFirstChargeRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerFirstChargeUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerFirstCharge(key int64) *PlayerFirstCharge {
	for crow := this.db.tables.PlayerFirstCharge; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerFirstChargeRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerFirstCharge(callback func(*PlayerFirstChargeRow)) {
	row := &PlayerFirstChargeRow{}
	for crow := this.db.tables.PlayerFirstCharge; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerFirstCharge() (rows []interface{}) {
	row := &PlayerFirstChargeRow{}
	for crow := this.db.tables.PlayerFirstCharge; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerFirstCharge", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerFirstCharge", len(rows))
	return rows
}

/* ========== player_formation ========== */

// 玩家阵型站位
type PlayerFormation struct {
	Pid int64 // 玩家ID
	Pos0 int8 // 站位0
	Pos1 int8 // 站位1
	Pos2 int8 // 站位2
	Pos3 int8 // 站位3
	Pos4 int8 // 站位4
	Pos5 int8 // 站位5
	Pos6 int8 // 站位6
	Pos7 int8 // 站位7
	Pos8 int8 // 站位8
}

func (this *PlayerFormation) CObject() *C.PlayerFormation {
	object := C.New_PlayerFormation()
	object.Pid = C.int64_t(this.Pid)
	object.Pos0 = C.int8_t(this.Pos0)
	object.Pos1 = C.int8_t(this.Pos1)
	object.Pos2 = C.int8_t(this.Pos2)
	object.Pos3 = C.int8_t(this.Pos3)
	object.Pos4 = C.int8_t(this.Pos4)
	object.Pos5 = C.int8_t(this.Pos5)
	object.Pos6 = C.int8_t(this.Pos6)
	object.Pos7 = C.int8_t(this.Pos7)
	object.Pos8 = C.int8_t(this.Pos8)
	return object
}

func (this insertOP) PlayerFormation(row *PlayerFormation) {
	if this.db.tables.PlayerFormation != nil {
		panic("duplicate insert 'player_formation'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerFormation = newRow
	this.db.addTransLog(this.db.newPlayerFormationInsertLog(newRow, row))
}

func (this deleteOP) PlayerFormation(row *PlayerFormation) {
	if this.db.tables.PlayerFormation == nil {
		panic("delete not exists 'player_formation'")
	}
	old := this.db.tables.PlayerFormation
	this.db.tables.PlayerFormation = nil
	this.db.addTransLog(this.db.newPlayerFormationDeleteLog(old, row))
}

func (this updateOP) PlayerFormation(row *PlayerFormation) {
	if this.db.tables.PlayerFormation == nil {
		panic("update not exists 'player_formation'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerFormation
	this.db.tables.PlayerFormation = newRow
	tmpRow := PlayerFormationRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerFormationUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerFormation(key int64) *PlayerFormation {
	if this.db.tables.PlayerFormation == nil {
		return nil
	}
	tmpRow := PlayerFormationRow{row:this.db.tables.PlayerFormation}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_func_key ========== */

// 玩家功能开放表
type PlayerFuncKey struct {
	Pid int64 // 玩家ID
	Key int16 // 功能权值
	PlayedKey int16 // 已播放提示的功能
	UniqueKey int64 // 已开启功能的唯一权值
}

func (this *PlayerFuncKey) CObject() *C.PlayerFuncKey {
	object := C.New_PlayerFuncKey()
	object.Pid = C.int64_t(this.Pid)
	object.Key = C.int16_t(this.Key)
	object.PlayedKey = C.int16_t(this.PlayedKey)
	object.UniqueKey = C.int64_t(this.UniqueKey)
	return object
}

func (this insertOP) PlayerFuncKey(row *PlayerFuncKey) {
	if this.db.tables.PlayerFuncKey != nil {
		panic("duplicate insert 'player_func_key'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerFuncKey = newRow
	this.db.addTransLog(this.db.newPlayerFuncKeyInsertLog(newRow, row))
}

func (this deleteOP) PlayerFuncKey(row *PlayerFuncKey) {
	if this.db.tables.PlayerFuncKey == nil {
		panic("delete not exists 'player_func_key'")
	}
	old := this.db.tables.PlayerFuncKey
	this.db.tables.PlayerFuncKey = nil
	this.db.addTransLog(this.db.newPlayerFuncKeyDeleteLog(old, row))
}

func (this updateOP) PlayerFuncKey(row *PlayerFuncKey) {
	if this.db.tables.PlayerFuncKey == nil {
		panic("update not exists 'player_func_key'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerFuncKey
	this.db.tables.PlayerFuncKey = newRow
	tmpRow := PlayerFuncKeyRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerFuncKeyUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerFuncKey(key int64) *PlayerFuncKey {
	if this.db.tables.PlayerFuncKey == nil {
		return nil
	}
	tmpRow := PlayerFuncKeyRow{row:this.db.tables.PlayerFuncKey}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_ghost ========== */

// 玩家魂侍表
type PlayerGhost struct {
	Id int64 // 主键
	Pid int64 // 玩家ID
	GhostId int16 // 魂侍ID
	Star int8 // 星级
	Level int16 // 魂侍等级
	Exp int64 // 魂侍经验
	Pos int16 // 位置
	IsNew int8 // 是否新魂侍
	RoleId int8 // 魂侍装备在那个角色身上 0 未装备
	SkillLevel int16 // 技能等级
	RelationId int64 // 连锁玩家魂侍ID
	AddGrowth int16 // 魂侍洗练添加的成长值
}

func (this *PlayerGhost) CObject() *C.PlayerGhost {
	object := C.New_PlayerGhost()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.GhostId = C.int16_t(this.GhostId)
	object.Star = C.int8_t(this.Star)
	object.Level = C.int16_t(this.Level)
	object.Exp = C.int64_t(this.Exp)
	object.Pos = C.int16_t(this.Pos)
	object.IsNew = C.int8_t(this.IsNew)
	object.RoleId = C.int8_t(this.RoleId)
	object.SkillLevel = C.int16_t(this.SkillLevel)
	object.RelationId = C.int64_t(this.RelationId)
	object.AddGrowth = C.int16_t(this.AddGrowth)
	return object
}

func (this insertOP) PlayerGhost(row *PlayerGhost) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerGhost, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerGhost
	for crow := this.db.tables.PlayerGhost; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_ghost'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerGhost
	this.db.tables.PlayerGhost = newRow
	this.db.addTransLog(this.db.newPlayerGhostInsertLog(newRow, row))
}

func (this deleteOP) PlayerGhost(row *PlayerGhost) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerGhost
	for crow := this.db.tables.PlayerGhost; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_ghost'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerGhost = this.db.tables.PlayerGhost.next
	}
	this.db.addTransLog(this.db.newPlayerGhostDeleteLog(old, row))
}

func (this updateOP) PlayerGhost(row *PlayerGhost) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerGhost
	for crow := this.db.tables.PlayerGhost; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_ghost'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerGhost = newRow
	}
	tmpRow := PlayerGhostRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerGhostUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerGhost(key int64) *PlayerGhost {
	for crow := this.db.tables.PlayerGhost; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerGhostRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerGhost(callback func(*PlayerGhostRow)) {
	row := &PlayerGhostRow{}
	for crow := this.db.tables.PlayerGhost; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerGhost() (rows []interface{}) {
	row := &PlayerGhostRow{}
	for crow := this.db.tables.PlayerGhost; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerGhost", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerGhost", len(rows))
	return rows
}

/* ========== player_ghost_equipment ========== */

// 玩家魂侍装备表
type PlayerGhostEquipment struct {
	Id int64 // 主键
	Pid int64 // 玩家ID
	RoleId int8 // 角色ID
	GhostPower int32 // 魂力
	Pos1 int64 // 装备位置1的魂侍id
	Pos2 int64 // 装备位置2的魂侍id
	Pos3 int64 // 装备位置3的魂侍id
	Pos4 int64 // 装备位置4的魂侍id
}

func (this *PlayerGhostEquipment) CObject() *C.PlayerGhostEquipment {
	object := C.New_PlayerGhostEquipment()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.RoleId = C.int8_t(this.RoleId)
	object.GhostPower = C.int32_t(this.GhostPower)
	object.Pos1 = C.int64_t(this.Pos1)
	object.Pos2 = C.int64_t(this.Pos2)
	object.Pos3 = C.int64_t(this.Pos3)
	object.Pos4 = C.int64_t(this.Pos4)
	return object
}

func (this insertOP) PlayerGhostEquipment(row *PlayerGhostEquipment) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerGhostEquipment, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerGhostEquipment
	for crow := this.db.tables.PlayerGhostEquipment; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_ghost_equipment'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerGhostEquipment
	this.db.tables.PlayerGhostEquipment = newRow
	this.db.addTransLog(this.db.newPlayerGhostEquipmentInsertLog(newRow, row))
}

func (this deleteOP) PlayerGhostEquipment(row *PlayerGhostEquipment) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerGhostEquipment
	for crow := this.db.tables.PlayerGhostEquipment; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_ghost_equipment'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerGhostEquipment = this.db.tables.PlayerGhostEquipment.next
	}
	this.db.addTransLog(this.db.newPlayerGhostEquipmentDeleteLog(old, row))
}

func (this updateOP) PlayerGhostEquipment(row *PlayerGhostEquipment) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerGhostEquipment
	for crow := this.db.tables.PlayerGhostEquipment; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_ghost_equipment'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerGhostEquipment = newRow
	}
	tmpRow := PlayerGhostEquipmentRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerGhostEquipmentUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerGhostEquipment(key int64) *PlayerGhostEquipment {
	for crow := this.db.tables.PlayerGhostEquipment; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerGhostEquipmentRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerGhostEquipment(callback func(*PlayerGhostEquipmentRow)) {
	row := &PlayerGhostEquipmentRow{}
	for crow := this.db.tables.PlayerGhostEquipment; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerGhostEquipment() (rows []interface{}) {
	row := &PlayerGhostEquipmentRow{}
	for crow := this.db.tables.PlayerGhostEquipment; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerGhostEquipment", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerGhostEquipment", len(rows))
	return rows
}

/* ========== player_ghost_state ========== */

// 玩家魂侍相关状态
type PlayerGhostState struct {
	Pid int64 // 玩家ID
	TrainByIngotNum int32 // 今日使用元宝培养次数
	TrainByIngotTime int64 // 最近一次使用元宝培养时间
	LastFlushTime int64 // 最近一次洗点时间
	GhostFightNum int64 // 玩家魂侍总战力
}

func (this *PlayerGhostState) CObject() *C.PlayerGhostState {
	object := C.New_PlayerGhostState()
	object.Pid = C.int64_t(this.Pid)
	object.TrainByIngotNum = C.int32_t(this.TrainByIngotNum)
	object.TrainByIngotTime = C.int64_t(this.TrainByIngotTime)
	object.LastFlushTime = C.int64_t(this.LastFlushTime)
	object.GhostFightNum = C.int64_t(this.GhostFightNum)
	return object
}

func (this insertOP) PlayerGhostState(row *PlayerGhostState) {
	if this.db.tables.PlayerGhostState != nil {
		panic("duplicate insert 'player_ghost_state'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerGhostState = newRow
	this.db.addTransLog(this.db.newPlayerGhostStateInsertLog(newRow, row))
}

func (this deleteOP) PlayerGhostState(row *PlayerGhostState) {
	if this.db.tables.PlayerGhostState == nil {
		panic("delete not exists 'player_ghost_state'")
	}
	old := this.db.tables.PlayerGhostState
	this.db.tables.PlayerGhostState = nil
	this.db.addTransLog(this.db.newPlayerGhostStateDeleteLog(old, row))
}

func (this updateOP) PlayerGhostState(row *PlayerGhostState) {
	if this.db.tables.PlayerGhostState == nil {
		panic("update not exists 'player_ghost_state'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerGhostState
	this.db.tables.PlayerGhostState = newRow
	tmpRow := PlayerGhostStateRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerGhostStateUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerGhostState(key int64) *PlayerGhostState {
	if this.db.tables.PlayerGhostState == nil {
		return nil
	}
	tmpRow := PlayerGhostStateRow{row:this.db.tables.PlayerGhostState}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_global_chat_state ========== */

// 禁言状态
type PlayerGlobalChatState struct {
	Pid int64 // 玩家ID
	BanUntil int64 // <=0 没有禁言 >0禁言
}

func (this *PlayerGlobalChatState) CObject() *C.PlayerGlobalChatState {
	object := C.New_PlayerGlobalChatState()
	object.Pid = C.int64_t(this.Pid)
	object.BanUntil = C.int64_t(this.BanUntil)
	return object
}

func (this insertOP) PlayerGlobalChatState(row *PlayerGlobalChatState) {
	if this.db.tables.PlayerGlobalChatState != nil {
		panic("duplicate insert 'player_global_chat_state'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerGlobalChatState = newRow
	this.db.addTransLog(this.db.newPlayerGlobalChatStateInsertLog(newRow, row))
}

func (this deleteOP) PlayerGlobalChatState(row *PlayerGlobalChatState) {
	if this.db.tables.PlayerGlobalChatState == nil {
		panic("delete not exists 'player_global_chat_state'")
	}
	old := this.db.tables.PlayerGlobalChatState
	this.db.tables.PlayerGlobalChatState = nil
	this.db.addTransLog(this.db.newPlayerGlobalChatStateDeleteLog(old, row))
}

func (this updateOP) PlayerGlobalChatState(row *PlayerGlobalChatState) {
	if this.db.tables.PlayerGlobalChatState == nil {
		panic("update not exists 'player_global_chat_state'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerGlobalChatState
	this.db.tables.PlayerGlobalChatState = newRow
	tmpRow := PlayerGlobalChatStateRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerGlobalChatStateUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerGlobalChatState(key int64) *PlayerGlobalChatState {
	if this.db.tables.PlayerGlobalChatState == nil {
		return nil
	}
	tmpRow := PlayerGlobalChatStateRow{row:this.db.tables.PlayerGlobalChatState}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_global_clique_building ========== */

// 玩家帮派建筑信息
type PlayerGlobalCliqueBuilding struct {
	Pid int64 // 玩家ID
	SilverExchangeTime int64 // 
	GoldExchangeNum int16 // 每日金劵兑换数
	SilverExchangeNum int16 // 每日银劵兑换数
	DonateCoinsCenterBuilding int64 // 总部累积贡献铜币
	DonateCoinsTempleBuilding int64 // 宗祠累积贡献铜币
	DonateCoinsBankBuilding int64 // 钱庄累积贡献铜币
	DonateCoinsHealthBuilding int64 // 回春堂累积贡献铜币
	DonateCoinsAttackBuilding int64 // 神兵堂累积贡献铜币
	DonateCoinsDefenseBuilding int64 // 金刚堂累积贡献铜币
	DonateCoinsStoreBuilding int64 // 战备仓库累积贡献铜币
	HealthLevel int16 // 回春等级
	AttackLevel int16 // 神兵等级
	DefenseLevel int16 // 金刚等级
	WorshipTime int64 // 上香时间
	DonateCoinsCenterBuildingTime int64 // 上次总舵捐钱时间戳
	GlodExchangeTime int64 // 金卷购买时间戳
	WorshipType int8 // 上香类型
}

func (this *PlayerGlobalCliqueBuilding) CObject() *C.PlayerGlobalCliqueBuilding {
	object := C.New_PlayerGlobalCliqueBuilding()
	object.Pid = C.int64_t(this.Pid)
	object.SilverExchangeTime = C.int64_t(this.SilverExchangeTime)
	object.GoldExchangeNum = C.int16_t(this.GoldExchangeNum)
	object.SilverExchangeNum = C.int16_t(this.SilverExchangeNum)
	object.DonateCoinsCenterBuilding = C.int64_t(this.DonateCoinsCenterBuilding)
	object.DonateCoinsTempleBuilding = C.int64_t(this.DonateCoinsTempleBuilding)
	object.DonateCoinsBankBuilding = C.int64_t(this.DonateCoinsBankBuilding)
	object.DonateCoinsHealthBuilding = C.int64_t(this.DonateCoinsHealthBuilding)
	object.DonateCoinsAttackBuilding = C.int64_t(this.DonateCoinsAttackBuilding)
	object.DonateCoinsDefenseBuilding = C.int64_t(this.DonateCoinsDefenseBuilding)
	object.DonateCoinsStoreBuilding = C.int64_t(this.DonateCoinsStoreBuilding)
	object.HealthLevel = C.int16_t(this.HealthLevel)
	object.AttackLevel = C.int16_t(this.AttackLevel)
	object.DefenseLevel = C.int16_t(this.DefenseLevel)
	object.WorshipTime = C.int64_t(this.WorshipTime)
	object.DonateCoinsCenterBuildingTime = C.int64_t(this.DonateCoinsCenterBuildingTime)
	object.GlodExchangeTime = C.int64_t(this.GlodExchangeTime)
	object.WorshipType = C.int8_t(this.WorshipType)
	return object
}

func (this insertOP) PlayerGlobalCliqueBuilding(row *PlayerGlobalCliqueBuilding) {
	if this.db.tables.PlayerGlobalCliqueBuilding != nil {
		panic("duplicate insert 'player_global_clique_building'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerGlobalCliqueBuilding = newRow
	this.db.addTransLog(this.db.newPlayerGlobalCliqueBuildingInsertLog(newRow, row))
}

func (this deleteOP) PlayerGlobalCliqueBuilding(row *PlayerGlobalCliqueBuilding) {
	if this.db.tables.PlayerGlobalCliqueBuilding == nil {
		panic("delete not exists 'player_global_clique_building'")
	}
	old := this.db.tables.PlayerGlobalCliqueBuilding
	this.db.tables.PlayerGlobalCliqueBuilding = nil
	this.db.addTransLog(this.db.newPlayerGlobalCliqueBuildingDeleteLog(old, row))
}

func (this updateOP) PlayerGlobalCliqueBuilding(row *PlayerGlobalCliqueBuilding) {
	if this.db.tables.PlayerGlobalCliqueBuilding == nil {
		panic("update not exists 'player_global_clique_building'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerGlobalCliqueBuilding
	this.db.tables.PlayerGlobalCliqueBuilding = newRow
	tmpRow := PlayerGlobalCliqueBuildingRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerGlobalCliqueBuildingUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerGlobalCliqueBuilding(key int64) *PlayerGlobalCliqueBuilding {
	if this.db.tables.PlayerGlobalCliqueBuilding == nil {
		return nil
	}
	tmpRow := PlayerGlobalCliqueBuildingRow{row:this.db.tables.PlayerGlobalCliqueBuilding}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_global_clique_building_quest ========== */

// 玩家帮派建筑任务
type PlayerGlobalCliqueBuildingQuest struct {
	Id int64 // 记录ID
	Pid int64 // 玩家ID
	QuestId int16 // 任务ID
	AwardStatus int8 // 奖励状态; 0 未奖励； 1已奖励
	BuildingType int16 // 建筑类别
}

func (this *PlayerGlobalCliqueBuildingQuest) CObject() *C.PlayerGlobalCliqueBuildingQuest {
	object := C.New_PlayerGlobalCliqueBuildingQuest()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.QuestId = C.int16_t(this.QuestId)
	object.AwardStatus = C.int8_t(this.AwardStatus)
	object.BuildingType = C.int16_t(this.BuildingType)
	return object
}

func (this insertOP) PlayerGlobalCliqueBuildingQuest(row *PlayerGlobalCliqueBuildingQuest) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerGlobalCliqueBuildingQuest, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerGlobalCliqueBuildingQuest
	for crow := this.db.tables.PlayerGlobalCliqueBuildingQuest; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_global_clique_building_quest'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerGlobalCliqueBuildingQuest
	this.db.tables.PlayerGlobalCliqueBuildingQuest = newRow
	this.db.addTransLog(this.db.newPlayerGlobalCliqueBuildingQuestInsertLog(newRow, row))
}

func (this deleteOP) PlayerGlobalCliqueBuildingQuest(row *PlayerGlobalCliqueBuildingQuest) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerGlobalCliqueBuildingQuest
	for crow := this.db.tables.PlayerGlobalCliqueBuildingQuest; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_global_clique_building_quest'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerGlobalCliqueBuildingQuest = this.db.tables.PlayerGlobalCliqueBuildingQuest.next
	}
	this.db.addTransLog(this.db.newPlayerGlobalCliqueBuildingQuestDeleteLog(old, row))
}

func (this updateOP) PlayerGlobalCliqueBuildingQuest(row *PlayerGlobalCliqueBuildingQuest) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerGlobalCliqueBuildingQuest
	for crow := this.db.tables.PlayerGlobalCliqueBuildingQuest; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_global_clique_building_quest'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerGlobalCliqueBuildingQuest = newRow
	}
	tmpRow := PlayerGlobalCliqueBuildingQuestRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerGlobalCliqueBuildingQuestUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerGlobalCliqueBuildingQuest(key int64) *PlayerGlobalCliqueBuildingQuest {
	for crow := this.db.tables.PlayerGlobalCliqueBuildingQuest; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerGlobalCliqueBuildingQuestRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerGlobalCliqueBuildingQuest(callback func(*PlayerGlobalCliqueBuildingQuestRow)) {
	row := &PlayerGlobalCliqueBuildingQuestRow{}
	for crow := this.db.tables.PlayerGlobalCliqueBuildingQuest; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerGlobalCliqueBuildingQuest() (rows []interface{}) {
	row := &PlayerGlobalCliqueBuildingQuestRow{}
	for crow := this.db.tables.PlayerGlobalCliqueBuildingQuest; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerGlobalCliqueBuildingQuest", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerGlobalCliqueBuildingQuest", len(rows))
	return rows
}

/* ========== player_global_clique_daily_quest ========== */

// 玩家每日帮派任务
type PlayerGlobalCliqueDailyQuest struct {
	Id int64 // 记录ID
	Pid int64 // 玩家ID
	QuestId int16 // 任务ID
	FinishCount int16 // 完成数量
	LastUpdateTime int64 // 最近一次更新时间
	AwardStatus int8 // 奖励状态; 0 未奖励； 1已奖励
	Class int16 // 每日任务类别
}

func (this *PlayerGlobalCliqueDailyQuest) CObject() *C.PlayerGlobalCliqueDailyQuest {
	object := C.New_PlayerGlobalCliqueDailyQuest()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.QuestId = C.int16_t(this.QuestId)
	object.FinishCount = C.int16_t(this.FinishCount)
	object.LastUpdateTime = C.int64_t(this.LastUpdateTime)
	object.AwardStatus = C.int8_t(this.AwardStatus)
	object.Class = C.int16_t(this.Class)
	return object
}

func (this insertOP) PlayerGlobalCliqueDailyQuest(row *PlayerGlobalCliqueDailyQuest) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerGlobalCliqueDailyQuest, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerGlobalCliqueDailyQuest
	for crow := this.db.tables.PlayerGlobalCliqueDailyQuest; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_global_clique_daily_quest'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerGlobalCliqueDailyQuest
	this.db.tables.PlayerGlobalCliqueDailyQuest = newRow
	this.db.addTransLog(this.db.newPlayerGlobalCliqueDailyQuestInsertLog(newRow, row))
}

func (this deleteOP) PlayerGlobalCliqueDailyQuest(row *PlayerGlobalCliqueDailyQuest) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerGlobalCliqueDailyQuest
	for crow := this.db.tables.PlayerGlobalCliqueDailyQuest; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_global_clique_daily_quest'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerGlobalCliqueDailyQuest = this.db.tables.PlayerGlobalCliqueDailyQuest.next
	}
	this.db.addTransLog(this.db.newPlayerGlobalCliqueDailyQuestDeleteLog(old, row))
}

func (this updateOP) PlayerGlobalCliqueDailyQuest(row *PlayerGlobalCliqueDailyQuest) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerGlobalCliqueDailyQuest
	for crow := this.db.tables.PlayerGlobalCliqueDailyQuest; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_global_clique_daily_quest'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerGlobalCliqueDailyQuest = newRow
	}
	tmpRow := PlayerGlobalCliqueDailyQuestRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerGlobalCliqueDailyQuestUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerGlobalCliqueDailyQuest(key int64) *PlayerGlobalCliqueDailyQuest {
	for crow := this.db.tables.PlayerGlobalCliqueDailyQuest; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerGlobalCliqueDailyQuestRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerGlobalCliqueDailyQuest(callback func(*PlayerGlobalCliqueDailyQuestRow)) {
	row := &PlayerGlobalCliqueDailyQuestRow{}
	for crow := this.db.tables.PlayerGlobalCliqueDailyQuest; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerGlobalCliqueDailyQuest() (rows []interface{}) {
	row := &PlayerGlobalCliqueDailyQuestRow{}
	for crow := this.db.tables.PlayerGlobalCliqueDailyQuest; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerGlobalCliqueDailyQuest", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerGlobalCliqueDailyQuest", len(rows))
	return rows
}

/* ========== player_global_clique_escort ========== */

// 玩家帮派运镖信息
type PlayerGlobalCliqueEscort struct {
	Pid int64 // 玩家ID
	DailyEscortTimestamp int64 // 今日首次送镖时间戳
	DailyEscortNum int16 // 送镖次数
	EscortBoatType int16 // 当前标船类型
	Status int8 // 0--无运镖 1--运镖中 2--等待领取奖励
	Hijacked int8 // 0--没有被完成过劫持 1--被完成过劫持 
}

func (this *PlayerGlobalCliqueEscort) CObject() *C.PlayerGlobalCliqueEscort {
	object := C.New_PlayerGlobalCliqueEscort()
	object.Pid = C.int64_t(this.Pid)
	object.DailyEscortTimestamp = C.int64_t(this.DailyEscortTimestamp)
	object.DailyEscortNum = C.int16_t(this.DailyEscortNum)
	object.EscortBoatType = C.int16_t(this.EscortBoatType)
	object.Status = C.int8_t(this.Status)
	object.Hijacked = C.int8_t(this.Hijacked)
	return object
}

func (this insertOP) PlayerGlobalCliqueEscort(row *PlayerGlobalCliqueEscort) {
	if this.db.tables.PlayerGlobalCliqueEscort != nil {
		panic("duplicate insert 'player_global_clique_escort'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerGlobalCliqueEscort = newRow
	this.db.addTransLog(this.db.newPlayerGlobalCliqueEscortInsertLog(newRow, row))
}

func (this deleteOP) PlayerGlobalCliqueEscort(row *PlayerGlobalCliqueEscort) {
	if this.db.tables.PlayerGlobalCliqueEscort == nil {
		panic("delete not exists 'player_global_clique_escort'")
	}
	old := this.db.tables.PlayerGlobalCliqueEscort
	this.db.tables.PlayerGlobalCliqueEscort = nil
	this.db.addTransLog(this.db.newPlayerGlobalCliqueEscortDeleteLog(old, row))
}

func (this updateOP) PlayerGlobalCliqueEscort(row *PlayerGlobalCliqueEscort) {
	if this.db.tables.PlayerGlobalCliqueEscort == nil {
		panic("update not exists 'player_global_clique_escort'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerGlobalCliqueEscort
	this.db.tables.PlayerGlobalCliqueEscort = newRow
	tmpRow := PlayerGlobalCliqueEscortRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerGlobalCliqueEscortUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerGlobalCliqueEscort(key int64) *PlayerGlobalCliqueEscort {
	if this.db.tables.PlayerGlobalCliqueEscort == nil {
		return nil
	}
	tmpRow := PlayerGlobalCliqueEscortRow{row:this.db.tables.PlayerGlobalCliqueEscort}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_global_clique_escort_message ========== */

// 玩家帮派运镖消息
type PlayerGlobalCliqueEscortMessage struct {
	Id int64 // 玩家ID
	Pid int64 // 玩家ID
	TplId int16 // 模版ID
	Parameters string // 模版参数
	Timestamp int64 // 发送时间戳
}

func (this *PlayerGlobalCliqueEscortMessage) CObject() *C.PlayerGlobalCliqueEscortMessage {
	object := C.New_PlayerGlobalCliqueEscortMessage()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.TplId = C.int16_t(this.TplId)
	object.Parameters = C.CString(string(this.Parameters))
	object.Parameters_len = C.int(len(this.Parameters))
	object.Timestamp = C.int64_t(this.Timestamp)
	return object
}

func (this insertOP) PlayerGlobalCliqueEscortMessage(row *PlayerGlobalCliqueEscortMessage) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerGlobalCliqueEscortMessage, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerGlobalCliqueEscortMessage
	for crow := this.db.tables.PlayerGlobalCliqueEscortMessage; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_global_clique_escort_message'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerGlobalCliqueEscortMessage
	this.db.tables.PlayerGlobalCliqueEscortMessage = newRow
	this.db.addTransLog(this.db.newPlayerGlobalCliqueEscortMessageInsertLog(newRow, row))
}

func (this deleteOP) PlayerGlobalCliqueEscortMessage(row *PlayerGlobalCliqueEscortMessage) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerGlobalCliqueEscortMessage
	for crow := this.db.tables.PlayerGlobalCliqueEscortMessage; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_global_clique_escort_message'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerGlobalCliqueEscortMessage = this.db.tables.PlayerGlobalCliqueEscortMessage.next
	}
	this.db.addTransLog(this.db.newPlayerGlobalCliqueEscortMessageDeleteLog(old, row))
}

func (this updateOP) PlayerGlobalCliqueEscortMessage(row *PlayerGlobalCliqueEscortMessage) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerGlobalCliqueEscortMessage
	for crow := this.db.tables.PlayerGlobalCliqueEscortMessage; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_global_clique_escort_message'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerGlobalCliqueEscortMessage = newRow
	}
	tmpRow := PlayerGlobalCliqueEscortMessageRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerGlobalCliqueEscortMessageUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerGlobalCliqueEscortMessage(key int64) *PlayerGlobalCliqueEscortMessage {
	for crow := this.db.tables.PlayerGlobalCliqueEscortMessage; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerGlobalCliqueEscortMessageRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerGlobalCliqueEscortMessage(callback func(*PlayerGlobalCliqueEscortMessageRow)) {
	row := &PlayerGlobalCliqueEscortMessageRow{}
	for crow := this.db.tables.PlayerGlobalCliqueEscortMessage; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerGlobalCliqueEscortMessage() (rows []interface{}) {
	row := &PlayerGlobalCliqueEscortMessageRow{}
	for crow := this.db.tables.PlayerGlobalCliqueEscortMessage; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerGlobalCliqueEscortMessage", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerGlobalCliqueEscortMessage", len(rows))
	return rows
}

/* ========== player_global_clique_hijack ========== */

// 玩家劫镖信息
type PlayerGlobalCliqueHijack struct {
	Pid int64 // 玩家ID
	DailyHijackTimestamp int64 // 劫持时间戳
	DailyHijackNum int16 // 劫持次数
	HijackedBoatType int16 // 待奖励的标船类型
	Status int8 // 0--无劫镖 1--劫镖中 2--等待领取奖励
	BuyTimestamp int64 // 购买时间戳
	BuyNum int16 // 今日购买次数
}

func (this *PlayerGlobalCliqueHijack) CObject() *C.PlayerGlobalCliqueHijack {
	object := C.New_PlayerGlobalCliqueHijack()
	object.Pid = C.int64_t(this.Pid)
	object.DailyHijackTimestamp = C.int64_t(this.DailyHijackTimestamp)
	object.DailyHijackNum = C.int16_t(this.DailyHijackNum)
	object.HijackedBoatType = C.int16_t(this.HijackedBoatType)
	object.Status = C.int8_t(this.Status)
	object.BuyTimestamp = C.int64_t(this.BuyTimestamp)
	object.BuyNum = C.int16_t(this.BuyNum)
	return object
}

func (this insertOP) PlayerGlobalCliqueHijack(row *PlayerGlobalCliqueHijack) {
	if this.db.tables.PlayerGlobalCliqueHijack != nil {
		panic("duplicate insert 'player_global_clique_hijack'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerGlobalCliqueHijack = newRow
	this.db.addTransLog(this.db.newPlayerGlobalCliqueHijackInsertLog(newRow, row))
}

func (this deleteOP) PlayerGlobalCliqueHijack(row *PlayerGlobalCliqueHijack) {
	if this.db.tables.PlayerGlobalCliqueHijack == nil {
		panic("delete not exists 'player_global_clique_hijack'")
	}
	old := this.db.tables.PlayerGlobalCliqueHijack
	this.db.tables.PlayerGlobalCliqueHijack = nil
	this.db.addTransLog(this.db.newPlayerGlobalCliqueHijackDeleteLog(old, row))
}

func (this updateOP) PlayerGlobalCliqueHijack(row *PlayerGlobalCliqueHijack) {
	if this.db.tables.PlayerGlobalCliqueHijack == nil {
		panic("update not exists 'player_global_clique_hijack'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerGlobalCliqueHijack
	this.db.tables.PlayerGlobalCliqueHijack = newRow
	tmpRow := PlayerGlobalCliqueHijackRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerGlobalCliqueHijackUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerGlobalCliqueHijack(key int64) *PlayerGlobalCliqueHijack {
	if this.db.tables.PlayerGlobalCliqueHijack == nil {
		return nil
	}
	tmpRow := PlayerGlobalCliqueHijackRow{row:this.db.tables.PlayerGlobalCliqueHijack}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_global_clique_info ========== */

// 玩家帮派信息
type PlayerGlobalCliqueInfo struct {
	Pid int64 // 玩家ID
	CliqueId int64 // 帮派名称
	JoinTime int64 // 加入帮派时间
	Contrib int64 // 帮派贡献
	ContribClearTime int64 // 帮派贡献结算时间(每周一清理贡献)
	DonateCoinsTime int64 // 贡献铜币时间戳(限制捐献大小)
	DailyDonateCoins int64 // 当日贡献铜币
	TotalContrib int64 // 玩家帮派总贡献
}

func (this *PlayerGlobalCliqueInfo) CObject() *C.PlayerGlobalCliqueInfo {
	object := C.New_PlayerGlobalCliqueInfo()
	object.Pid = C.int64_t(this.Pid)
	object.CliqueId = C.int64_t(this.CliqueId)
	object.JoinTime = C.int64_t(this.JoinTime)
	object.Contrib = C.int64_t(this.Contrib)
	object.ContribClearTime = C.int64_t(this.ContribClearTime)
	object.DonateCoinsTime = C.int64_t(this.DonateCoinsTime)
	object.DailyDonateCoins = C.int64_t(this.DailyDonateCoins)
	object.TotalContrib = C.int64_t(this.TotalContrib)
	return object
}

func (this insertOP) PlayerGlobalCliqueInfo(row *PlayerGlobalCliqueInfo) {
	if this.db.tables.PlayerGlobalCliqueInfo != nil {
		panic("duplicate insert 'player_global_clique_info'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerGlobalCliqueInfo = newRow
	this.db.addTransLog(this.db.newPlayerGlobalCliqueInfoInsertLog(newRow, row))
}

func (this deleteOP) PlayerGlobalCliqueInfo(row *PlayerGlobalCliqueInfo) {
	if this.db.tables.PlayerGlobalCliqueInfo == nil {
		panic("delete not exists 'player_global_clique_info'")
	}
	old := this.db.tables.PlayerGlobalCliqueInfo
	this.db.tables.PlayerGlobalCliqueInfo = nil
	this.db.addTransLog(this.db.newPlayerGlobalCliqueInfoDeleteLog(old, row))
}

func (this updateOP) PlayerGlobalCliqueInfo(row *PlayerGlobalCliqueInfo) {
	if this.db.tables.PlayerGlobalCliqueInfo == nil {
		panic("update not exists 'player_global_clique_info'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerGlobalCliqueInfo
	this.db.tables.PlayerGlobalCliqueInfo = newRow
	tmpRow := PlayerGlobalCliqueInfoRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerGlobalCliqueInfoUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerGlobalCliqueInfo(key int64) *PlayerGlobalCliqueInfo {
	if this.db.tables.PlayerGlobalCliqueInfo == nil {
		return nil
	}
	tmpRow := PlayerGlobalCliqueInfoRow{row:this.db.tables.PlayerGlobalCliqueInfo}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_global_clique_kongfu ========== */

// 玩家帮派武学
type PlayerGlobalCliqueKongfu struct {
	Id int64 // ID
	Pid int64 // 玩家ID
	KongfuId int32 // 武功ID
	Level int16 // 武功等级
}

func (this *PlayerGlobalCliqueKongfu) CObject() *C.PlayerGlobalCliqueKongfu {
	object := C.New_PlayerGlobalCliqueKongfu()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.KongfuId = C.int32_t(this.KongfuId)
	object.Level = C.int16_t(this.Level)
	return object
}

func (this insertOP) PlayerGlobalCliqueKongfu(row *PlayerGlobalCliqueKongfu) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerGlobalCliqueKongfu, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerGlobalCliqueKongfu
	for crow := this.db.tables.PlayerGlobalCliqueKongfu; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_global_clique_kongfu'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerGlobalCliqueKongfu
	this.db.tables.PlayerGlobalCliqueKongfu = newRow
	this.db.addTransLog(this.db.newPlayerGlobalCliqueKongfuInsertLog(newRow, row))
}

func (this deleteOP) PlayerGlobalCliqueKongfu(row *PlayerGlobalCliqueKongfu) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerGlobalCliqueKongfu
	for crow := this.db.tables.PlayerGlobalCliqueKongfu; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_global_clique_kongfu'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerGlobalCliqueKongfu = this.db.tables.PlayerGlobalCliqueKongfu.next
	}
	this.db.addTransLog(this.db.newPlayerGlobalCliqueKongfuDeleteLog(old, row))
}

func (this updateOP) PlayerGlobalCliqueKongfu(row *PlayerGlobalCliqueKongfu) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerGlobalCliqueKongfu
	for crow := this.db.tables.PlayerGlobalCliqueKongfu; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_global_clique_kongfu'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerGlobalCliqueKongfu = newRow
	}
	tmpRow := PlayerGlobalCliqueKongfuRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerGlobalCliqueKongfuUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerGlobalCliqueKongfu(key int64) *PlayerGlobalCliqueKongfu {
	for crow := this.db.tables.PlayerGlobalCliqueKongfu; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerGlobalCliqueKongfuRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerGlobalCliqueKongfu(callback func(*PlayerGlobalCliqueKongfuRow)) {
	row := &PlayerGlobalCliqueKongfuRow{}
	for crow := this.db.tables.PlayerGlobalCliqueKongfu; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerGlobalCliqueKongfu() (rows []interface{}) {
	row := &PlayerGlobalCliqueKongfuRow{}
	for crow := this.db.tables.PlayerGlobalCliqueKongfu; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerGlobalCliqueKongfu", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerGlobalCliqueKongfu", len(rows))
	return rows
}

/* ========== player_global_friend ========== */

// 玩家好友列表
type PlayerGlobalFriend struct {
	Id int64 // 好友关系ID
	Pid int64 // 玩家ID
	FriendPid int64 // 好友ID
	FriendNick string // 玩家昵称
	FriendRoleId int8 // 好友角色ID
	FriendMode int8 // 好友关系:0陌生人,1仅关注,2仅被关注,3互相关注(好友)
	LastChatTime int64 // 最后聊天时间
	FriendTime int64 // 成为好友时间
	SendHeartTime int64 // 上次送爱心时间
	BlockMode int8 // 黑名单状态:0-否,1-是
}

func (this *PlayerGlobalFriend) CObject() *C.PlayerGlobalFriend {
	object := C.New_PlayerGlobalFriend()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.FriendPid = C.int64_t(this.FriendPid)
	object.FriendNick = C.CString(string(this.FriendNick))
	object.FriendNick_len = C.int(len(this.FriendNick))
	object.FriendRoleId = C.int8_t(this.FriendRoleId)
	object.FriendMode = C.int8_t(this.FriendMode)
	object.LastChatTime = C.int64_t(this.LastChatTime)
	object.FriendTime = C.int64_t(this.FriendTime)
	object.SendHeartTime = C.int64_t(this.SendHeartTime)
	object.BlockMode = C.int8_t(this.BlockMode)
	return object
}

func (this insertOP) PlayerGlobalFriend(row *PlayerGlobalFriend) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerGlobalFriend, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerGlobalFriend
	for crow := this.db.tables.PlayerGlobalFriend; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_global_friend'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerGlobalFriend
	this.db.tables.PlayerGlobalFriend = newRow
	this.db.addTransLog(this.db.newPlayerGlobalFriendInsertLog(newRow, row))
}

func (this deleteOP) PlayerGlobalFriend(row *PlayerGlobalFriend) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerGlobalFriend
	for crow := this.db.tables.PlayerGlobalFriend; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_global_friend'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerGlobalFriend = this.db.tables.PlayerGlobalFriend.next
	}
	this.db.addTransLog(this.db.newPlayerGlobalFriendDeleteLog(old, row))
}

func (this updateOP) PlayerGlobalFriend(row *PlayerGlobalFriend) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerGlobalFriend
	for crow := this.db.tables.PlayerGlobalFriend; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_global_friend'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerGlobalFriend = newRow
	}
	tmpRow := PlayerGlobalFriendRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerGlobalFriendUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerGlobalFriend(key int64) *PlayerGlobalFriend {
	for crow := this.db.tables.PlayerGlobalFriend; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerGlobalFriendRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerGlobalFriend(callback func(*PlayerGlobalFriendRow)) {
	row := &PlayerGlobalFriendRow{}
	for crow := this.db.tables.PlayerGlobalFriend; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerGlobalFriend() (rows []interface{}) {
	row := &PlayerGlobalFriendRow{}
	for crow := this.db.tables.PlayerGlobalFriend; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerGlobalFriend", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerGlobalFriend", len(rows))
	return rows
}

/* ========== player_global_friend_chat ========== */

// 玩家聊天记录
type PlayerGlobalFriendChat struct {
	Id int64 // 消息ID
	Pid int64 // 玩家ID
	FriendPid int64 // 对方玩家ID
	Mode int8 // 1发送，2接收
	SendTime int64 // 发送时间戳
	Message string // 消息内容
}

func (this *PlayerGlobalFriendChat) CObject() *C.PlayerGlobalFriendChat {
	object := C.New_PlayerGlobalFriendChat()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.FriendPid = C.int64_t(this.FriendPid)
	object.Mode = C.int8_t(this.Mode)
	object.SendTime = C.int64_t(this.SendTime)
	object.Message = C.CString(string(this.Message))
	object.Message_len = C.int(len(this.Message))
	return object
}

func (this insertOP) PlayerGlobalFriendChat(row *PlayerGlobalFriendChat) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerGlobalFriendChat, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerGlobalFriendChat
	for crow := this.db.tables.PlayerGlobalFriendChat; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_global_friend_chat'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerGlobalFriendChat
	this.db.tables.PlayerGlobalFriendChat = newRow
	this.db.addTransLog(this.db.newPlayerGlobalFriendChatInsertLog(newRow, row))
}

func (this deleteOP) PlayerGlobalFriendChat(row *PlayerGlobalFriendChat) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerGlobalFriendChat
	for crow := this.db.tables.PlayerGlobalFriendChat; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_global_friend_chat'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerGlobalFriendChat = this.db.tables.PlayerGlobalFriendChat.next
	}
	this.db.addTransLog(this.db.newPlayerGlobalFriendChatDeleteLog(old, row))
}

func (this updateOP) PlayerGlobalFriendChat(row *PlayerGlobalFriendChat) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerGlobalFriendChat
	for crow := this.db.tables.PlayerGlobalFriendChat; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_global_friend_chat'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerGlobalFriendChat = newRow
	}
	tmpRow := PlayerGlobalFriendChatRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerGlobalFriendChatUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerGlobalFriendChat(key int64) *PlayerGlobalFriendChat {
	for crow := this.db.tables.PlayerGlobalFriendChat; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerGlobalFriendChatRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerGlobalFriendChat(callback func(*PlayerGlobalFriendChatRow)) {
	row := &PlayerGlobalFriendChatRow{}
	for crow := this.db.tables.PlayerGlobalFriendChat; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerGlobalFriendChat() (rows []interface{}) {
	row := &PlayerGlobalFriendChatRow{}
	for crow := this.db.tables.PlayerGlobalFriendChat; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerGlobalFriendChat", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerGlobalFriendChat", len(rows))
	return rows
}

/* ========== player_global_friend_state ========== */

// 玩家好友功能状态数据
type PlayerGlobalFriendState struct {
	Pid int64 // 玩家ID
	DeleteDayCount int32 // 每日删除好友数量
	DeleteTime int64 // 删除好友时间
	ExistOfflineChat int8 // 0没有离线消息，1有离线消息
	PlatformFriendNum int32 // 平台好友历史最大值
}

func (this *PlayerGlobalFriendState) CObject() *C.PlayerGlobalFriendState {
	object := C.New_PlayerGlobalFriendState()
	object.Pid = C.int64_t(this.Pid)
	object.DeleteDayCount = C.int32_t(this.DeleteDayCount)
	object.DeleteTime = C.int64_t(this.DeleteTime)
	object.ExistOfflineChat = C.int8_t(this.ExistOfflineChat)
	object.PlatformFriendNum = C.int32_t(this.PlatformFriendNum)
	return object
}

func (this insertOP) PlayerGlobalFriendState(row *PlayerGlobalFriendState) {
	if this.db.tables.PlayerGlobalFriendState != nil {
		panic("duplicate insert 'player_global_friend_state'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerGlobalFriendState = newRow
	this.db.addTransLog(this.db.newPlayerGlobalFriendStateInsertLog(newRow, row))
}

func (this deleteOP) PlayerGlobalFriendState(row *PlayerGlobalFriendState) {
	if this.db.tables.PlayerGlobalFriendState == nil {
		panic("delete not exists 'player_global_friend_state'")
	}
	old := this.db.tables.PlayerGlobalFriendState
	this.db.tables.PlayerGlobalFriendState = nil
	this.db.addTransLog(this.db.newPlayerGlobalFriendStateDeleteLog(old, row))
}

func (this updateOP) PlayerGlobalFriendState(row *PlayerGlobalFriendState) {
	if this.db.tables.PlayerGlobalFriendState == nil {
		panic("update not exists 'player_global_friend_state'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerGlobalFriendState
	this.db.tables.PlayerGlobalFriendState = newRow
	tmpRow := PlayerGlobalFriendStateRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerGlobalFriendStateUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerGlobalFriendState(key int64) *PlayerGlobalFriendState {
	if this.db.tables.PlayerGlobalFriendState == nil {
		return nil
	}
	tmpRow := PlayerGlobalFriendStateRow{row:this.db.tables.PlayerGlobalFriendState}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_global_gift_card_record ========== */

// 玩家兑换记录
type PlayerGlobalGiftCardRecord struct {
	Id int64 // 主键
	Pid int64 // 兑换玩家PID
	Code string // 兑换码
}

func (this *PlayerGlobalGiftCardRecord) CObject() *C.PlayerGlobalGiftCardRecord {
	object := C.New_PlayerGlobalGiftCardRecord()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.Code = C.CString(string(this.Code))
	object.Code_len = C.int(len(this.Code))
	return object
}

func (this insertOP) PlayerGlobalGiftCardRecord(row *PlayerGlobalGiftCardRecord) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerGlobalGiftCardRecord, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerGlobalGiftCardRecord
	for crow := this.db.tables.PlayerGlobalGiftCardRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_global_gift_card_record'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerGlobalGiftCardRecord
	this.db.tables.PlayerGlobalGiftCardRecord = newRow
	this.db.addTransLog(this.db.newPlayerGlobalGiftCardRecordInsertLog(newRow, row))
}

func (this deleteOP) PlayerGlobalGiftCardRecord(row *PlayerGlobalGiftCardRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerGlobalGiftCardRecord
	for crow := this.db.tables.PlayerGlobalGiftCardRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_global_gift_card_record'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerGlobalGiftCardRecord = this.db.tables.PlayerGlobalGiftCardRecord.next
	}
	this.db.addTransLog(this.db.newPlayerGlobalGiftCardRecordDeleteLog(old, row))
}

func (this updateOP) PlayerGlobalGiftCardRecord(row *PlayerGlobalGiftCardRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerGlobalGiftCardRecord
	for crow := this.db.tables.PlayerGlobalGiftCardRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_global_gift_card_record'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerGlobalGiftCardRecord = newRow
	}
	tmpRow := PlayerGlobalGiftCardRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerGlobalGiftCardRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerGlobalGiftCardRecord(key int64) *PlayerGlobalGiftCardRecord {
	for crow := this.db.tables.PlayerGlobalGiftCardRecord; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerGlobalGiftCardRecordRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerGlobalGiftCardRecord(callback func(*PlayerGlobalGiftCardRecordRow)) {
	row := &PlayerGlobalGiftCardRecordRow{}
	for crow := this.db.tables.PlayerGlobalGiftCardRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerGlobalGiftCardRecord() (rows []interface{}) {
	row := &PlayerGlobalGiftCardRecordRow{}
	for crow := this.db.tables.PlayerGlobalGiftCardRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerGlobalGiftCardRecord", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerGlobalGiftCardRecord", len(rows))
	return rows
}

/* ========== player_global_mail_state ========== */

// 玩家全局邮件记录
type PlayerGlobalMailState struct {
	Pid int64 // 玩家ID
	MaxTimestamp int64 // 发送时间戳
}

func (this *PlayerGlobalMailState) CObject() *C.PlayerGlobalMailState {
	object := C.New_PlayerGlobalMailState()
	object.Pid = C.int64_t(this.Pid)
	object.MaxTimestamp = C.int64_t(this.MaxTimestamp)
	return object
}

func (this insertOP) PlayerGlobalMailState(row *PlayerGlobalMailState) {
	if this.db.tables.PlayerGlobalMailState != nil {
		panic("duplicate insert 'player_global_mail_state'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerGlobalMailState = newRow
	this.db.addTransLog(this.db.newPlayerGlobalMailStateInsertLog(newRow, row))
}

func (this deleteOP) PlayerGlobalMailState(row *PlayerGlobalMailState) {
	if this.db.tables.PlayerGlobalMailState == nil {
		panic("delete not exists 'player_global_mail_state'")
	}
	old := this.db.tables.PlayerGlobalMailState
	this.db.tables.PlayerGlobalMailState = nil
	this.db.addTransLog(this.db.newPlayerGlobalMailStateDeleteLog(old, row))
}

func (this updateOP) PlayerGlobalMailState(row *PlayerGlobalMailState) {
	if this.db.tables.PlayerGlobalMailState == nil {
		panic("update not exists 'player_global_mail_state'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerGlobalMailState
	this.db.tables.PlayerGlobalMailState = newRow
	tmpRow := PlayerGlobalMailStateRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerGlobalMailStateUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerGlobalMailState(key int64) *PlayerGlobalMailState {
	if this.db.tables.PlayerGlobalMailState == nil {
		return nil
	}
	tmpRow := PlayerGlobalMailStateRow{row:this.db.tables.PlayerGlobalMailState}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_global_world_chat_state ========== */

// 玩家世界聊天状态
type PlayerGlobalWorldChatState struct {
	Pid int64 // 玩家ID
	Timestamp int64 // 最近一次发送聊天时间戳
	DailyNum int16 // 今日次数
}

func (this *PlayerGlobalWorldChatState) CObject() *C.PlayerGlobalWorldChatState {
	object := C.New_PlayerGlobalWorldChatState()
	object.Pid = C.int64_t(this.Pid)
	object.Timestamp = C.int64_t(this.Timestamp)
	object.DailyNum = C.int16_t(this.DailyNum)
	return object
}

func (this insertOP) PlayerGlobalWorldChatState(row *PlayerGlobalWorldChatState) {
	if this.db.tables.PlayerGlobalWorldChatState != nil {
		panic("duplicate insert 'player_global_world_chat_state'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerGlobalWorldChatState = newRow
	this.db.addTransLog(this.db.newPlayerGlobalWorldChatStateInsertLog(newRow, row))
}

func (this deleteOP) PlayerGlobalWorldChatState(row *PlayerGlobalWorldChatState) {
	if this.db.tables.PlayerGlobalWorldChatState == nil {
		panic("delete not exists 'player_global_world_chat_state'")
	}
	old := this.db.tables.PlayerGlobalWorldChatState
	this.db.tables.PlayerGlobalWorldChatState = nil
	this.db.addTransLog(this.db.newPlayerGlobalWorldChatStateDeleteLog(old, row))
}

func (this updateOP) PlayerGlobalWorldChatState(row *PlayerGlobalWorldChatState) {
	if this.db.tables.PlayerGlobalWorldChatState == nil {
		panic("update not exists 'player_global_world_chat_state'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerGlobalWorldChatState
	this.db.tables.PlayerGlobalWorldChatState = newRow
	tmpRow := PlayerGlobalWorldChatStateRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerGlobalWorldChatStateUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerGlobalWorldChatState(key int64) *PlayerGlobalWorldChatState {
	if this.db.tables.PlayerGlobalWorldChatState == nil {
		return nil
	}
	tmpRow := PlayerGlobalWorldChatStateRow{row:this.db.tables.PlayerGlobalWorldChatState}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_hard_level ========== */

// 难度关卡功能权值
type PlayerHardLevel struct {
	Pid int64 // 玩家ID
	Lock int32 // 难度关卡功能权值
	AwardLock int32 // 已获得过奖励关卡的最大lock
}

func (this *PlayerHardLevel) CObject() *C.PlayerHardLevel {
	object := C.New_PlayerHardLevel()
	object.Pid = C.int64_t(this.Pid)
	object.Lock = C.int32_t(this.Lock)
	object.AwardLock = C.int32_t(this.AwardLock)
	return object
}

func (this insertOP) PlayerHardLevel(row *PlayerHardLevel) {
	if this.db.tables.PlayerHardLevel != nil {
		panic("duplicate insert 'player_hard_level'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerHardLevel = newRow
	this.db.addTransLog(this.db.newPlayerHardLevelInsertLog(newRow, row))
}

func (this deleteOP) PlayerHardLevel(row *PlayerHardLevel) {
	if this.db.tables.PlayerHardLevel == nil {
		panic("delete not exists 'player_hard_level'")
	}
	old := this.db.tables.PlayerHardLevel
	this.db.tables.PlayerHardLevel = nil
	this.db.addTransLog(this.db.newPlayerHardLevelDeleteLog(old, row))
}

func (this updateOP) PlayerHardLevel(row *PlayerHardLevel) {
	if this.db.tables.PlayerHardLevel == nil {
		panic("update not exists 'player_hard_level'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerHardLevel
	this.db.tables.PlayerHardLevel = newRow
	tmpRow := PlayerHardLevelRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerHardLevelUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerHardLevel(key int64) *PlayerHardLevel {
	if this.db.tables.PlayerHardLevel == nil {
		return nil
	}
	tmpRow := PlayerHardLevelRow{row:this.db.tables.PlayerHardLevel}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_hard_level_record ========== */

// 难度关卡记录
type PlayerHardLevelRecord struct {
	Id int64 // 记录ID
	Pid int64 // 玩家ID
	LevelId int32 // 开启的关卡ID
	OpenTime int64 // 关卡开启时间
	Score int32 // 得分
	Round int8 // 通关回合数
	DailyNum int8 // 当日已进入关卡的次数
	LastEnterTime int64 // 最后一次进入时间
	BuyTimes int16 // 深渊关卡今日购买次数
	BuyUpdateTime int64 // 深渊关卡上次购买时间戳
}

func (this *PlayerHardLevelRecord) CObject() *C.PlayerHardLevelRecord {
	object := C.New_PlayerHardLevelRecord()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.LevelId = C.int32_t(this.LevelId)
	object.OpenTime = C.int64_t(this.OpenTime)
	object.Score = C.int32_t(this.Score)
	object.Round = C.int8_t(this.Round)
	object.DailyNum = C.int8_t(this.DailyNum)
	object.LastEnterTime = C.int64_t(this.LastEnterTime)
	object.BuyTimes = C.int16_t(this.BuyTimes)
	object.BuyUpdateTime = C.int64_t(this.BuyUpdateTime)
	return object
}

func (this insertOP) PlayerHardLevelRecord(row *PlayerHardLevelRecord) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerHardLevelRecord, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerHardLevelRecord
	for crow := this.db.tables.PlayerHardLevelRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_hard_level_record'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerHardLevelRecord
	this.db.tables.PlayerHardLevelRecord = newRow
	this.db.addTransLog(this.db.newPlayerHardLevelRecordInsertLog(newRow, row))
}

func (this deleteOP) PlayerHardLevelRecord(row *PlayerHardLevelRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerHardLevelRecord
	for crow := this.db.tables.PlayerHardLevelRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_hard_level_record'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerHardLevelRecord = this.db.tables.PlayerHardLevelRecord.next
	}
	this.db.addTransLog(this.db.newPlayerHardLevelRecordDeleteLog(old, row))
}

func (this updateOP) PlayerHardLevelRecord(row *PlayerHardLevelRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerHardLevelRecord
	for crow := this.db.tables.PlayerHardLevelRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_hard_level_record'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerHardLevelRecord = newRow
	}
	tmpRow := PlayerHardLevelRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerHardLevelRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerHardLevelRecord(key int64) *PlayerHardLevelRecord {
	for crow := this.db.tables.PlayerHardLevelRecord; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerHardLevelRecordRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerHardLevelRecord(callback func(*PlayerHardLevelRecordRow)) {
	row := &PlayerHardLevelRecordRow{}
	for crow := this.db.tables.PlayerHardLevelRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerHardLevelRecord() (rows []interface{}) {
	row := &PlayerHardLevelRecordRow{}
	for crow := this.db.tables.PlayerHardLevelRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerHardLevelRecord", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerHardLevelRecord", len(rows))
	return rows
}

/* ========== player_heart ========== */

// 玩家爱心表
type PlayerHeart struct {
	Pid int64 // 玩家ID
	Value int16 // 爱心值
	UpdateTime int64 // 最后更新时间
	AddDayCount int32 // 今日好友赠送数量
	AddTime int64 // 接受好友赠送爱心时间
	RecoverDayCount int16 // 今日恢复数量
}

func (this *PlayerHeart) CObject() *C.PlayerHeart {
	object := C.New_PlayerHeart()
	object.Pid = C.int64_t(this.Pid)
	object.Value = C.int16_t(this.Value)
	object.UpdateTime = C.int64_t(this.UpdateTime)
	object.AddDayCount = C.int32_t(this.AddDayCount)
	object.AddTime = C.int64_t(this.AddTime)
	object.RecoverDayCount = C.int16_t(this.RecoverDayCount)
	return object
}

func (this insertOP) PlayerHeart(row *PlayerHeart) {
	if this.db.tables.PlayerHeart != nil {
		panic("duplicate insert 'player_heart'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerHeart = newRow
	this.db.addTransLog(this.db.newPlayerHeartInsertLog(newRow, row))
}

func (this deleteOP) PlayerHeart(row *PlayerHeart) {
	if this.db.tables.PlayerHeart == nil {
		panic("delete not exists 'player_heart'")
	}
	old := this.db.tables.PlayerHeart
	this.db.tables.PlayerHeart = nil
	this.db.addTransLog(this.db.newPlayerHeartDeleteLog(old, row))
}

func (this updateOP) PlayerHeart(row *PlayerHeart) {
	if this.db.tables.PlayerHeart == nil {
		panic("update not exists 'player_heart'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerHeart
	this.db.tables.PlayerHeart = newRow
	tmpRow := PlayerHeartRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerHeartUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerHeart(key int64) *PlayerHeart {
	if this.db.tables.PlayerHeart == nil {
		return nil
	}
	tmpRow := PlayerHeartRow{row:this.db.tables.PlayerHeart}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_heart_draw ========== */

// 玩家爱心抽奖
type PlayerHeartDraw struct {
	Id int64 // 
	Pid int64 // 玩家ID
	DrawType int8 // 抽奖类型（1-大转盘；2-刮刮卡）
	DailyNum int8 // 当日已抽次数
	DrawTime int64 // 最近一次抽奖时间
}

func (this *PlayerHeartDraw) CObject() *C.PlayerHeartDraw {
	object := C.New_PlayerHeartDraw()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.DrawType = C.int8_t(this.DrawType)
	object.DailyNum = C.int8_t(this.DailyNum)
	object.DrawTime = C.int64_t(this.DrawTime)
	return object
}

func (this insertOP) PlayerHeartDraw(row *PlayerHeartDraw) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerHeartDraw, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerHeartDraw
	for crow := this.db.tables.PlayerHeartDraw; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_heart_draw'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerHeartDraw
	this.db.tables.PlayerHeartDraw = newRow
	this.db.addTransLog(this.db.newPlayerHeartDrawInsertLog(newRow, row))
}

func (this deleteOP) PlayerHeartDraw(row *PlayerHeartDraw) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerHeartDraw
	for crow := this.db.tables.PlayerHeartDraw; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_heart_draw'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerHeartDraw = this.db.tables.PlayerHeartDraw.next
	}
	this.db.addTransLog(this.db.newPlayerHeartDrawDeleteLog(old, row))
}

func (this updateOP) PlayerHeartDraw(row *PlayerHeartDraw) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerHeartDraw
	for crow := this.db.tables.PlayerHeartDraw; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_heart_draw'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerHeartDraw = newRow
	}
	tmpRow := PlayerHeartDrawRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerHeartDrawUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerHeartDraw(key int64) *PlayerHeartDraw {
	for crow := this.db.tables.PlayerHeartDraw; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerHeartDrawRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerHeartDraw(callback func(*PlayerHeartDrawRow)) {
	row := &PlayerHeartDrawRow{}
	for crow := this.db.tables.PlayerHeartDraw; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerHeartDraw() (rows []interface{}) {
	row := &PlayerHeartDrawRow{}
	for crow := this.db.tables.PlayerHeartDraw; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerHeartDraw", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerHeartDraw", len(rows))
	return rows
}

/* ========== player_heart_draw_card_record ========== */

// 玩家爱心刮刮卡抽奖记录
type PlayerHeartDrawCardRecord struct {
	Id int64 // 
	Pid int64 // 玩家ID
	AwardType int8 // 奖品类型（1-铜钱；2-元宝；3-道具）
	AwardNum int16 // 奖品数量
	ItemId int16 // 道具奖品ID
	DrawTime int64 // 抽奖时间
}

func (this *PlayerHeartDrawCardRecord) CObject() *C.PlayerHeartDrawCardRecord {
	object := C.New_PlayerHeartDrawCardRecord()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.AwardType = C.int8_t(this.AwardType)
	object.AwardNum = C.int16_t(this.AwardNum)
	object.ItemId = C.int16_t(this.ItemId)
	object.DrawTime = C.int64_t(this.DrawTime)
	return object
}

func (this insertOP) PlayerHeartDrawCardRecord(row *PlayerHeartDrawCardRecord) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerHeartDrawCardRecord, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerHeartDrawCardRecord
	for crow := this.db.tables.PlayerHeartDrawCardRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_heart_draw_card_record'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerHeartDrawCardRecord
	this.db.tables.PlayerHeartDrawCardRecord = newRow
	this.db.addTransLog(this.db.newPlayerHeartDrawCardRecordInsertLog(newRow, row))
}

func (this deleteOP) PlayerHeartDrawCardRecord(row *PlayerHeartDrawCardRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerHeartDrawCardRecord
	for crow := this.db.tables.PlayerHeartDrawCardRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_heart_draw_card_record'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerHeartDrawCardRecord = this.db.tables.PlayerHeartDrawCardRecord.next
	}
	this.db.addTransLog(this.db.newPlayerHeartDrawCardRecordDeleteLog(old, row))
}

func (this updateOP) PlayerHeartDrawCardRecord(row *PlayerHeartDrawCardRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerHeartDrawCardRecord
	for crow := this.db.tables.PlayerHeartDrawCardRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_heart_draw_card_record'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerHeartDrawCardRecord = newRow
	}
	tmpRow := PlayerHeartDrawCardRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerHeartDrawCardRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerHeartDrawCardRecord(key int64) *PlayerHeartDrawCardRecord {
	for crow := this.db.tables.PlayerHeartDrawCardRecord; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerHeartDrawCardRecordRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerHeartDrawCardRecord(callback func(*PlayerHeartDrawCardRecordRow)) {
	row := &PlayerHeartDrawCardRecordRow{}
	for crow := this.db.tables.PlayerHeartDrawCardRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerHeartDrawCardRecord() (rows []interface{}) {
	row := &PlayerHeartDrawCardRecordRow{}
	for crow := this.db.tables.PlayerHeartDrawCardRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerHeartDrawCardRecord", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerHeartDrawCardRecord", len(rows))
	return rows
}

/* ========== player_heart_draw_wheel_record ========== */

// 玩家爱心大转盘抽奖记录
type PlayerHeartDrawWheelRecord struct {
	Id int64 // 
	Pid int64 // 玩家ID
	AwardType int8 // 奖品类型（1-铜钱；2-元宝；3-道具）
	AwardNum int16 // 奖品数量
	ItemId int16 // 道具奖品ID
	DrawTime int64 // 抽奖时间
}

func (this *PlayerHeartDrawWheelRecord) CObject() *C.PlayerHeartDrawWheelRecord {
	object := C.New_PlayerHeartDrawWheelRecord()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.AwardType = C.int8_t(this.AwardType)
	object.AwardNum = C.int16_t(this.AwardNum)
	object.ItemId = C.int16_t(this.ItemId)
	object.DrawTime = C.int64_t(this.DrawTime)
	return object
}

func (this insertOP) PlayerHeartDrawWheelRecord(row *PlayerHeartDrawWheelRecord) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerHeartDrawWheelRecord, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerHeartDrawWheelRecord
	for crow := this.db.tables.PlayerHeartDrawWheelRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_heart_draw_wheel_record'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerHeartDrawWheelRecord
	this.db.tables.PlayerHeartDrawWheelRecord = newRow
	this.db.addTransLog(this.db.newPlayerHeartDrawWheelRecordInsertLog(newRow, row))
}

func (this deleteOP) PlayerHeartDrawWheelRecord(row *PlayerHeartDrawWheelRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerHeartDrawWheelRecord
	for crow := this.db.tables.PlayerHeartDrawWheelRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_heart_draw_wheel_record'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerHeartDrawWheelRecord = this.db.tables.PlayerHeartDrawWheelRecord.next
	}
	this.db.addTransLog(this.db.newPlayerHeartDrawWheelRecordDeleteLog(old, row))
}

func (this updateOP) PlayerHeartDrawWheelRecord(row *PlayerHeartDrawWheelRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerHeartDrawWheelRecord
	for crow := this.db.tables.PlayerHeartDrawWheelRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_heart_draw_wheel_record'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerHeartDrawWheelRecord = newRow
	}
	tmpRow := PlayerHeartDrawWheelRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerHeartDrawWheelRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerHeartDrawWheelRecord(key int64) *PlayerHeartDrawWheelRecord {
	for crow := this.db.tables.PlayerHeartDrawWheelRecord; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerHeartDrawWheelRecordRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerHeartDrawWheelRecord(callback func(*PlayerHeartDrawWheelRecordRow)) {
	row := &PlayerHeartDrawWheelRecordRow{}
	for crow := this.db.tables.PlayerHeartDrawWheelRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerHeartDrawWheelRecord() (rows []interface{}) {
	row := &PlayerHeartDrawWheelRecordRow{}
	for crow := this.db.tables.PlayerHeartDrawWheelRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerHeartDrawWheelRecord", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerHeartDrawWheelRecord", len(rows))
	return rows
}

/* ========== player_info ========== */

// 玩家信息表
type PlayerInfo struct {
	Pid int64 // 玩家ID
	Ingot int64 // 元宝
	Coins int64 // 铜钱
	NewMailNum int16 // 新邮件数
	LastLoginTime int64 // 上次登录时间
	LastOfflineTime int64 // 上次下线时间
	TotalOnlineTime int64 // 总在线时间
	FirstLoginTime int64 // 玩家注册时间
	NewArenaReportNum int16 // 玩家离线比武场战报数
	LastSkillFlush int64 // 最近一次洗点时间
	SwordSoulFragment int64 // 剑心碎片数量
}

func (this *PlayerInfo) CObject() *C.PlayerInfo {
	object := C.New_PlayerInfo()
	object.Pid = C.int64_t(this.Pid)
	object.Ingot = C.int64_t(this.Ingot)
	object.Coins = C.int64_t(this.Coins)
	object.NewMailNum = C.int16_t(this.NewMailNum)
	object.LastLoginTime = C.int64_t(this.LastLoginTime)
	object.LastOfflineTime = C.int64_t(this.LastOfflineTime)
	object.TotalOnlineTime = C.int64_t(this.TotalOnlineTime)
	object.FirstLoginTime = C.int64_t(this.FirstLoginTime)
	object.NewArenaReportNum = C.int16_t(this.NewArenaReportNum)
	object.LastSkillFlush = C.int64_t(this.LastSkillFlush)
	object.SwordSoulFragment = C.int64_t(this.SwordSoulFragment)
	return object
}

func (this insertOP) PlayerInfo(row *PlayerInfo) {
	if this.db.tables.PlayerInfo != nil {
		panic("duplicate insert 'player_info'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerInfo = newRow
	this.db.addTransLog(this.db.newPlayerInfoInsertLog(newRow, row))
}

func (this deleteOP) PlayerInfo(row *PlayerInfo) {
	if this.db.tables.PlayerInfo == nil {
		panic("delete not exists 'player_info'")
	}
	old := this.db.tables.PlayerInfo
	this.db.tables.PlayerInfo = nil
	this.db.addTransLog(this.db.newPlayerInfoDeleteLog(old, row))
}

func (this updateOP) PlayerInfo(row *PlayerInfo) {
	if this.db.tables.PlayerInfo == nil {
		panic("update not exists 'player_info'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerInfo
	this.db.tables.PlayerInfo = newRow
	tmpRow := PlayerInfoRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerInfoUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerInfo(key int64) *PlayerInfo {
	if this.db.tables.PlayerInfo == nil {
		return nil
	}
	tmpRow := PlayerInfoRow{row:this.db.tables.PlayerInfo}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_is_operated ========== */

// 记录玩家是否第一次操作
type PlayerIsOperated struct {
	Pid int64 // 玩家ID
	OperateValue int64 // 操作值
}

func (this *PlayerIsOperated) CObject() *C.PlayerIsOperated {
	object := C.New_PlayerIsOperated()
	object.Pid = C.int64_t(this.Pid)
	object.OperateValue = C.int64_t(this.OperateValue)
	return object
}

func (this insertOP) PlayerIsOperated(row *PlayerIsOperated) {
	if this.db.tables.PlayerIsOperated != nil {
		panic("duplicate insert 'player_is_operated'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerIsOperated = newRow
	this.db.addTransLog(this.db.newPlayerIsOperatedInsertLog(newRow, row))
}

func (this deleteOP) PlayerIsOperated(row *PlayerIsOperated) {
	if this.db.tables.PlayerIsOperated == nil {
		panic("delete not exists 'player_is_operated'")
	}
	old := this.db.tables.PlayerIsOperated
	this.db.tables.PlayerIsOperated = nil
	this.db.addTransLog(this.db.newPlayerIsOperatedDeleteLog(old, row))
}

func (this updateOP) PlayerIsOperated(row *PlayerIsOperated) {
	if this.db.tables.PlayerIsOperated == nil {
		panic("update not exists 'player_is_operated'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerIsOperated
	this.db.tables.PlayerIsOperated = newRow
	tmpRow := PlayerIsOperatedRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerIsOperatedUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerIsOperated(key int64) *PlayerIsOperated {
	if this.db.tables.PlayerIsOperated == nil {
		return nil
	}
	tmpRow := PlayerIsOperatedRow{row:this.db.tables.PlayerIsOperated}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_item ========== */

// 玩家物品
type PlayerItem struct {
	Id int64 // 主键ID
	Pid int64 // 玩家ID
	ItemId int16 // 物品ID
	Num int16 // 数量
	IsDressed int8 // 是否被装备
	BuyBackState int8 // 记录物品是否在回购栏
	AppendixId int64 // 附加属性ID
	RefineLevelBak int16 // 精炼等级备份
	RefineFailTimes int16 // 精炼失败次数
	RefineLevel int16 // 精炼等级
	Price int32 // 装备精炼价格
}

func (this *PlayerItem) CObject() *C.PlayerItem {
	object := C.New_PlayerItem()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.ItemId = C.int16_t(this.ItemId)
	object.Num = C.int16_t(this.Num)
	object.IsDressed = C.int8_t(this.IsDressed)
	object.BuyBackState = C.int8_t(this.BuyBackState)
	object.AppendixId = C.int64_t(this.AppendixId)
	object.RefineLevelBak = C.int16_t(this.RefineLevelBak)
	object.RefineFailTimes = C.int16_t(this.RefineFailTimes)
	object.RefineLevel = C.int16_t(this.RefineLevel)
	object.Price = C.int32_t(this.Price)
	return object
}

func (this insertOP) PlayerItem(row *PlayerItem) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerItem, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerItem
	for crow := this.db.tables.PlayerItem; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_item'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerItem
	this.db.tables.PlayerItem = newRow
	this.db.addTransLog(this.db.newPlayerItemInsertLog(newRow, row))
}

func (this deleteOP) PlayerItem(row *PlayerItem) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerItem
	for crow := this.db.tables.PlayerItem; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_item'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerItem = this.db.tables.PlayerItem.next
	}
	this.db.addTransLog(this.db.newPlayerItemDeleteLog(old, row))
}

func (this updateOP) PlayerItem(row *PlayerItem) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerItem
	for crow := this.db.tables.PlayerItem; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_item'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerItem = newRow
	}
	tmpRow := PlayerItemRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerItemUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerItem(key int64) *PlayerItem {
	for crow := this.db.tables.PlayerItem; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerItemRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerItem(callback func(*PlayerItemRow)) {
	row := &PlayerItemRow{}
	for crow := this.db.tables.PlayerItem; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerItem() (rows []interface{}) {
	row := &PlayerItemRow{}
	for crow := this.db.tables.PlayerItem; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerItem", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerItem", len(rows))
	return rows
}

/* ========== player_item_appendix ========== */

// 玩家装备追加属性表
type PlayerItemAppendix struct {
	Id int64 // ID
	Pid int64 // 玩家ID
	Health int32 // 生命
	Cultivation int32 // 内力
	Speed int32 // 速度
	Attack int32 // 攻击
	Defence int32 // 防御
	DodgeLevel int32 // 闪避
	HitLevel int32 // 命中
	BlockLevel int32 // 格挡
	CriticalLevel int32 // 暴击
	TenacityLevel int32 // 韧性
	DestroyLevel int32 // 破击
	RecastAttr int8 // 重铸属性
}

func (this *PlayerItemAppendix) CObject() *C.PlayerItemAppendix {
	object := C.New_PlayerItemAppendix()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.Health = C.int32_t(this.Health)
	object.Cultivation = C.int32_t(this.Cultivation)
	object.Speed = C.int32_t(this.Speed)
	object.Attack = C.int32_t(this.Attack)
	object.Defence = C.int32_t(this.Defence)
	object.DodgeLevel = C.int32_t(this.DodgeLevel)
	object.HitLevel = C.int32_t(this.HitLevel)
	object.BlockLevel = C.int32_t(this.BlockLevel)
	object.CriticalLevel = C.int32_t(this.CriticalLevel)
	object.TenacityLevel = C.int32_t(this.TenacityLevel)
	object.DestroyLevel = C.int32_t(this.DestroyLevel)
	object.RecastAttr = C.int8_t(this.RecastAttr)
	return object
}

func (this insertOP) PlayerItemAppendix(row *PlayerItemAppendix) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerItemAppendix, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerItemAppendix
	for crow := this.db.tables.PlayerItemAppendix; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_item_appendix'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerItemAppendix
	this.db.tables.PlayerItemAppendix = newRow
	this.db.addTransLog(this.db.newPlayerItemAppendixInsertLog(newRow, row))
}

func (this deleteOP) PlayerItemAppendix(row *PlayerItemAppendix) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerItemAppendix
	for crow := this.db.tables.PlayerItemAppendix; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_item_appendix'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerItemAppendix = this.db.tables.PlayerItemAppendix.next
	}
	this.db.addTransLog(this.db.newPlayerItemAppendixDeleteLog(old, row))
}

func (this updateOP) PlayerItemAppendix(row *PlayerItemAppendix) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerItemAppendix
	for crow := this.db.tables.PlayerItemAppendix; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_item_appendix'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerItemAppendix = newRow
	}
	tmpRow := PlayerItemAppendixRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerItemAppendixUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerItemAppendix(key int64) *PlayerItemAppendix {
	for crow := this.db.tables.PlayerItemAppendix; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerItemAppendixRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerItemAppendix(callback func(*PlayerItemAppendixRow)) {
	row := &PlayerItemAppendixRow{}
	for crow := this.db.tables.PlayerItemAppendix; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerItemAppendix() (rows []interface{}) {
	row := &PlayerItemAppendixRow{}
	for crow := this.db.tables.PlayerItemAppendix; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerItemAppendix", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerItemAppendix", len(rows))
	return rows
}

/* ========== player_item_buyback ========== */

// 玩家物品回购表
type PlayerItemBuyback struct {
	Pid int64 // 玩家ID
	BackId1 int64 // 回购格子1,player_item表主键ID
	BackId2 int64 // 回购格子2,player_item表主键ID
	BackId3 int64 // 回购格子3,player_item表主键ID
	BackId4 int64 // 回购格子4,player_item表主键ID
	BackId5 int64 // 回购格子5,player_item表主键ID
	BackId6 int64 // 回购格子6,player_item表主键ID
}

func (this *PlayerItemBuyback) CObject() *C.PlayerItemBuyback {
	object := C.New_PlayerItemBuyback()
	object.Pid = C.int64_t(this.Pid)
	object.BackId1 = C.int64_t(this.BackId1)
	object.BackId2 = C.int64_t(this.BackId2)
	object.BackId3 = C.int64_t(this.BackId3)
	object.BackId4 = C.int64_t(this.BackId4)
	object.BackId5 = C.int64_t(this.BackId5)
	object.BackId6 = C.int64_t(this.BackId6)
	return object
}

func (this insertOP) PlayerItemBuyback(row *PlayerItemBuyback) {
	if this.db.tables.PlayerItemBuyback != nil {
		panic("duplicate insert 'player_item_buyback'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerItemBuyback = newRow
	this.db.addTransLog(this.db.newPlayerItemBuybackInsertLog(newRow, row))
}

func (this deleteOP) PlayerItemBuyback(row *PlayerItemBuyback) {
	if this.db.tables.PlayerItemBuyback == nil {
		panic("delete not exists 'player_item_buyback'")
	}
	old := this.db.tables.PlayerItemBuyback
	this.db.tables.PlayerItemBuyback = nil
	this.db.addTransLog(this.db.newPlayerItemBuybackDeleteLog(old, row))
}

func (this updateOP) PlayerItemBuyback(row *PlayerItemBuyback) {
	if this.db.tables.PlayerItemBuyback == nil {
		panic("update not exists 'player_item_buyback'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerItemBuyback
	this.db.tables.PlayerItemBuyback = newRow
	tmpRow := PlayerItemBuybackRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerItemBuybackUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerItemBuyback(key int64) *PlayerItemBuyback {
	if this.db.tables.PlayerItemBuyback == nil {
		return nil
	}
	tmpRow := PlayerItemBuybackRow{row:this.db.tables.PlayerItemBuyback}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_item_recast_state ========== */

// 玩家装备重铸状态
type PlayerItemRecastState struct {
	Pid int64 // 玩家ID
	PlayerItemId int64 // 玩家装备ID
	SelectedAttr int8 // 选中的属性
	Attr1Type int8 // 重铸属性1类型
	Attr1Value int32 // 重铸属性1数值
	Attr2Type int8 // 重铸属性2类型
	Attr2Value int32 // 重铸属性2数值
	Attr3Type int8 // 重铸属性3类型
	Attr3Value int32 // 重铸属性3数值
}

func (this *PlayerItemRecastState) CObject() *C.PlayerItemRecastState {
	object := C.New_PlayerItemRecastState()
	object.Pid = C.int64_t(this.Pid)
	object.PlayerItemId = C.int64_t(this.PlayerItemId)
	object.SelectedAttr = C.int8_t(this.SelectedAttr)
	object.Attr1Type = C.int8_t(this.Attr1Type)
	object.Attr1Value = C.int32_t(this.Attr1Value)
	object.Attr2Type = C.int8_t(this.Attr2Type)
	object.Attr2Value = C.int32_t(this.Attr2Value)
	object.Attr3Type = C.int8_t(this.Attr3Type)
	object.Attr3Value = C.int32_t(this.Attr3Value)
	return object
}

func (this insertOP) PlayerItemRecastState(row *PlayerItemRecastState) {
	if this.db.tables.PlayerItemRecastState != nil {
		panic("duplicate insert 'player_item_recast_state'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerItemRecastState = newRow
	this.db.addTransLog(this.db.newPlayerItemRecastStateInsertLog(newRow, row))
}

func (this deleteOP) PlayerItemRecastState(row *PlayerItemRecastState) {
	if this.db.tables.PlayerItemRecastState == nil {
		panic("delete not exists 'player_item_recast_state'")
	}
	old := this.db.tables.PlayerItemRecastState
	this.db.tables.PlayerItemRecastState = nil
	this.db.addTransLog(this.db.newPlayerItemRecastStateDeleteLog(old, row))
}

func (this updateOP) PlayerItemRecastState(row *PlayerItemRecastState) {
	if this.db.tables.PlayerItemRecastState == nil {
		panic("update not exists 'player_item_recast_state'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerItemRecastState
	this.db.tables.PlayerItemRecastState = newRow
	tmpRow := PlayerItemRecastStateRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerItemRecastStateUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerItemRecastState(key int64) *PlayerItemRecastState {
	if this.db.tables.PlayerItemRecastState == nil {
		return nil
	}
	tmpRow := PlayerItemRecastStateRow{row:this.db.tables.PlayerItemRecastState}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_level_award_info ========== */

// 玩家等级奖励领取状态表
type PlayerLevelAwardInfo struct {
	Pid int64 // 玩家id
	Awarded string // 玩家已领取状态
}

func (this *PlayerLevelAwardInfo) CObject() *C.PlayerLevelAwardInfo {
	object := C.New_PlayerLevelAwardInfo()
	object.Pid = C.int64_t(this.Pid)
	object.Awarded = C.CString(string(this.Awarded))
	object.Awarded_len = C.int(len(this.Awarded))
	return object
}

func (this insertOP) PlayerLevelAwardInfo(row *PlayerLevelAwardInfo) {
	if this.db.tables.PlayerLevelAwardInfo != nil {
		panic("duplicate insert 'player_level_award_info'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerLevelAwardInfo = newRow
	this.db.addTransLog(this.db.newPlayerLevelAwardInfoInsertLog(newRow, row))
}

func (this deleteOP) PlayerLevelAwardInfo(row *PlayerLevelAwardInfo) {
	if this.db.tables.PlayerLevelAwardInfo == nil {
		panic("delete not exists 'player_level_award_info'")
	}
	old := this.db.tables.PlayerLevelAwardInfo
	this.db.tables.PlayerLevelAwardInfo = nil
	this.db.addTransLog(this.db.newPlayerLevelAwardInfoDeleteLog(old, row))
}

func (this updateOP) PlayerLevelAwardInfo(row *PlayerLevelAwardInfo) {
	if this.db.tables.PlayerLevelAwardInfo == nil {
		panic("update not exists 'player_level_award_info'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerLevelAwardInfo
	this.db.tables.PlayerLevelAwardInfo = newRow
	tmpRow := PlayerLevelAwardInfoRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerLevelAwardInfoUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerLevelAwardInfo(key int64) *PlayerLevelAwardInfo {
	if this.db.tables.PlayerLevelAwardInfo == nil {
		return nil
	}
	tmpRow := PlayerLevelAwardInfoRow{row:this.db.tables.PlayerLevelAwardInfo}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_login_award_record ========== */

// 玩家奖励领取情况
type PlayerLoginAwardRecord struct {
	Pid int64 // 玩家ID
	ActiveDays int32 // 累计活跃天数
	Record int32 // 七天奖励领取记录
	UpdateTimestamp int64 // 更新时间戳
}

func (this *PlayerLoginAwardRecord) CObject() *C.PlayerLoginAwardRecord {
	object := C.New_PlayerLoginAwardRecord()
	object.Pid = C.int64_t(this.Pid)
	object.ActiveDays = C.int32_t(this.ActiveDays)
	object.Record = C.int32_t(this.Record)
	object.UpdateTimestamp = C.int64_t(this.UpdateTimestamp)
	return object
}

func (this insertOP) PlayerLoginAwardRecord(row *PlayerLoginAwardRecord) {
	if this.db.tables.PlayerLoginAwardRecord != nil {
		panic("duplicate insert 'player_login_award_record'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerLoginAwardRecord = newRow
	this.db.addTransLog(this.db.newPlayerLoginAwardRecordInsertLog(newRow, row))
}

func (this deleteOP) PlayerLoginAwardRecord(row *PlayerLoginAwardRecord) {
	if this.db.tables.PlayerLoginAwardRecord == nil {
		panic("delete not exists 'player_login_award_record'")
	}
	old := this.db.tables.PlayerLoginAwardRecord
	this.db.tables.PlayerLoginAwardRecord = nil
	this.db.addTransLog(this.db.newPlayerLoginAwardRecordDeleteLog(old, row))
}

func (this updateOP) PlayerLoginAwardRecord(row *PlayerLoginAwardRecord) {
	if this.db.tables.PlayerLoginAwardRecord == nil {
		panic("update not exists 'player_login_award_record'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerLoginAwardRecord
	this.db.tables.PlayerLoginAwardRecord = newRow
	tmpRow := PlayerLoginAwardRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerLoginAwardRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerLoginAwardRecord(key int64) *PlayerLoginAwardRecord {
	if this.db.tables.PlayerLoginAwardRecord == nil {
		return nil
	}
	tmpRow := PlayerLoginAwardRecordRow{row:this.db.tables.PlayerLoginAwardRecord}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_mail ========== */

// 玩家邮件表
type PlayerMail struct {
	Id int64 // 玩家邮件ID
	Pid int64 // 玩家ID
	MailId int32 // 邮件模版ID
	State int8 // 0未读，1已读
	SendTime int64 // 发送时间戳
	Parameters string // 模版参数
	HaveAttachment int8 // 是否有附件
	Title string // 标题
	Content string // 内容
	ExpireTime int64 // 邮件删除时机
	Priority int8 // 优先级
}

func (this *PlayerMail) CObject() *C.PlayerMail {
	object := C.New_PlayerMail()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.MailId = C.int32_t(this.MailId)
	object.State = C.int8_t(this.State)
	object.SendTime = C.int64_t(this.SendTime)
	object.Parameters = C.CString(string(this.Parameters))
	object.Parameters_len = C.int(len(this.Parameters))
	object.HaveAttachment = C.int8_t(this.HaveAttachment)
	object.Title = C.CString(string(this.Title))
	object.Title_len = C.int(len(this.Title))
	object.Content = C.CString(string(this.Content))
	object.Content_len = C.int(len(this.Content))
	object.ExpireTime = C.int64_t(this.ExpireTime)
	object.Priority = C.int8_t(this.Priority)
	return object
}

func (this insertOP) PlayerMail(row *PlayerMail) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerMail, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerMail
	for crow := this.db.tables.PlayerMail; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_mail'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerMail
	this.db.tables.PlayerMail = newRow
	this.db.addTransLog(this.db.newPlayerMailInsertLog(newRow, row))
}

func (this deleteOP) PlayerMail(row *PlayerMail) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerMail
	for crow := this.db.tables.PlayerMail; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_mail'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerMail = this.db.tables.PlayerMail.next
	}
	this.db.addTransLog(this.db.newPlayerMailDeleteLog(old, row))
}

func (this updateOP) PlayerMail(row *PlayerMail) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerMail
	for crow := this.db.tables.PlayerMail; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_mail'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerMail = newRow
	}
	tmpRow := PlayerMailRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerMailUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerMail(key int64) *PlayerMail {
	for crow := this.db.tables.PlayerMail; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerMailRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerMail(callback func(*PlayerMailRow)) {
	row := &PlayerMailRow{}
	for crow := this.db.tables.PlayerMail; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerMail() (rows []interface{}) {
	row := &PlayerMailRow{}
	for crow := this.db.tables.PlayerMail; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerMail", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerMail", len(rows))
	return rows
}

/* ========== player_mail_attachment ========== */

// 玩家邮件附件表
type PlayerMailAttachment struct {
	Id int64 // 玩家邮件附件ID
	Pid int64 // 玩家ID
	PlayerMailId int64 // player_mail 主键ID
	AttachmentType int8 // 附件类型
	ItemId int16 // 物品
	ItemNum int64 // 数量
}

func (this *PlayerMailAttachment) CObject() *C.PlayerMailAttachment {
	object := C.New_PlayerMailAttachment()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.PlayerMailId = C.int64_t(this.PlayerMailId)
	object.AttachmentType = C.int8_t(this.AttachmentType)
	object.ItemId = C.int16_t(this.ItemId)
	object.ItemNum = C.int64_t(this.ItemNum)
	return object
}

func (this insertOP) PlayerMailAttachment(row *PlayerMailAttachment) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerMailAttachment, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerMailAttachment
	for crow := this.db.tables.PlayerMailAttachment; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_mail_attachment'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerMailAttachment
	this.db.tables.PlayerMailAttachment = newRow
	this.db.addTransLog(this.db.newPlayerMailAttachmentInsertLog(newRow, row))
}

func (this deleteOP) PlayerMailAttachment(row *PlayerMailAttachment) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerMailAttachment
	for crow := this.db.tables.PlayerMailAttachment; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_mail_attachment'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerMailAttachment = this.db.tables.PlayerMailAttachment.next
	}
	this.db.addTransLog(this.db.newPlayerMailAttachmentDeleteLog(old, row))
}

func (this updateOP) PlayerMailAttachment(row *PlayerMailAttachment) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerMailAttachment
	for crow := this.db.tables.PlayerMailAttachment; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_mail_attachment'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerMailAttachment = newRow
	}
	tmpRow := PlayerMailAttachmentRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerMailAttachmentUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerMailAttachment(key int64) *PlayerMailAttachment {
	for crow := this.db.tables.PlayerMailAttachment; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerMailAttachmentRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerMailAttachment(callback func(*PlayerMailAttachmentRow)) {
	row := &PlayerMailAttachmentRow{}
	for crow := this.db.tables.PlayerMailAttachment; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerMailAttachment() (rows []interface{}) {
	row := &PlayerMailAttachmentRow{}
	for crow := this.db.tables.PlayerMailAttachment; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerMailAttachment", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerMailAttachment", len(rows))
	return rows
}

/* ========== player_mail_attachment_lg ========== */

// 玩家领取邮件附件表
type PlayerMailAttachmentLg struct {
	Id int64 // 玩家邮件附件ID
	Pid int64 // 玩家ID
	PlayerMailId int64 // player_mail 主键ID
	AttachmentType int8 // 附件类型
	ItemId int16 // 物品
	ItemNum int64 // 数量
	TakeTimestamp int64 // 附件领取时间
}

func (this *PlayerMailAttachmentLg) CObject() *C.PlayerMailAttachmentLg {
	object := C.New_PlayerMailAttachmentLg()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.PlayerMailId = C.int64_t(this.PlayerMailId)
	object.AttachmentType = C.int8_t(this.AttachmentType)
	object.ItemId = C.int16_t(this.ItemId)
	object.ItemNum = C.int64_t(this.ItemNum)
	object.TakeTimestamp = C.int64_t(this.TakeTimestamp)
	return object
}

func (this insertOP) PlayerMailAttachmentLg(row *PlayerMailAttachmentLg) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerMailAttachmentLg, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerMailAttachmentLg
	for crow := this.db.tables.PlayerMailAttachmentLg; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_mail_attachment_lg'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerMailAttachmentLg
	this.db.tables.PlayerMailAttachmentLg = newRow
	this.db.addTransLog(this.db.newPlayerMailAttachmentLgInsertLog(newRow, row))
}

func (this deleteOP) PlayerMailAttachmentLg(row *PlayerMailAttachmentLg) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerMailAttachmentLg
	for crow := this.db.tables.PlayerMailAttachmentLg; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_mail_attachment_lg'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerMailAttachmentLg = this.db.tables.PlayerMailAttachmentLg.next
	}
	this.db.addTransLog(this.db.newPlayerMailAttachmentLgDeleteLog(old, row))
}

func (this updateOP) PlayerMailAttachmentLg(row *PlayerMailAttachmentLg) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerMailAttachmentLg
	for crow := this.db.tables.PlayerMailAttachmentLg; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_mail_attachment_lg'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerMailAttachmentLg = newRow
	}
	tmpRow := PlayerMailAttachmentLgRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerMailAttachmentLgUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerMailAttachmentLg(key int64) *PlayerMailAttachmentLg {
	for crow := this.db.tables.PlayerMailAttachmentLg; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerMailAttachmentLgRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerMailAttachmentLg(callback func(*PlayerMailAttachmentLgRow)) {
	row := &PlayerMailAttachmentLgRow{}
	for crow := this.db.tables.PlayerMailAttachmentLg; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerMailAttachmentLg() (rows []interface{}) {
	row := &PlayerMailAttachmentLgRow{}
	for crow := this.db.tables.PlayerMailAttachmentLg; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerMailAttachmentLg", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerMailAttachmentLg", len(rows))
	return rows
}

/* ========== player_mail_lg ========== */

// 玩家已删除邮件表
type PlayerMailLg struct {
	Id int64 // ID
	Pmid int64 // 玩家邮件ID
	Pid int64 // 玩家ID
	MailId int32 // 邮件模版ID
	State int8 // 0未读，1已读
	SendTime int64 // 发送时间戳
	Parameters string // 模版参数
	HaveAttachment int8 // 是否有附件
	Title string // 标题
	Content string // 内容
}

func (this *PlayerMailLg) CObject() *C.PlayerMailLg {
	object := C.New_PlayerMailLg()
	object.Id = C.int64_t(this.Id)
	object.Pmid = C.int64_t(this.Pmid)
	object.Pid = C.int64_t(this.Pid)
	object.MailId = C.int32_t(this.MailId)
	object.State = C.int8_t(this.State)
	object.SendTime = C.int64_t(this.SendTime)
	object.Parameters = C.CString(string(this.Parameters))
	object.Parameters_len = C.int(len(this.Parameters))
	object.HaveAttachment = C.int8_t(this.HaveAttachment)
	object.Title = C.CString(string(this.Title))
	object.Title_len = C.int(len(this.Title))
	object.Content = C.CString(string(this.Content))
	object.Content_len = C.int(len(this.Content))
	return object
}

func (this insertOP) PlayerMailLg(row *PlayerMailLg) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerMailLg, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerMailLg
	for crow := this.db.tables.PlayerMailLg; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_mail_lg'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerMailLg
	this.db.tables.PlayerMailLg = newRow
	this.db.addTransLog(this.db.newPlayerMailLgInsertLog(newRow, row))
}

func (this deleteOP) PlayerMailLg(row *PlayerMailLg) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerMailLg
	for crow := this.db.tables.PlayerMailLg; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_mail_lg'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerMailLg = this.db.tables.PlayerMailLg.next
	}
	this.db.addTransLog(this.db.newPlayerMailLgDeleteLog(old, row))
}

func (this updateOP) PlayerMailLg(row *PlayerMailLg) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerMailLg
	for crow := this.db.tables.PlayerMailLg; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_mail_lg'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerMailLg = newRow
	}
	tmpRow := PlayerMailLgRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerMailLgUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerMailLg(key int64) *PlayerMailLg {
	for crow := this.db.tables.PlayerMailLg; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerMailLgRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerMailLg(callback func(*PlayerMailLgRow)) {
	row := &PlayerMailLgRow{}
	for crow := this.db.tables.PlayerMailLg; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerMailLg() (rows []interface{}) {
	row := &PlayerMailLgRow{}
	for crow := this.db.tables.PlayerMailLg; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerMailLg", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerMailLg", len(rows))
	return rows
}

/* ========== player_meditation_state ========== */

// 玩家打坐状态
type PlayerMeditationState struct {
	Pid int64 // player id
	AccumulateTime int32 // 光明钥匙奖励累积时间
	StartTimestamp int64 // 打坐开始时间 0-未未打坐状态
}

func (this *PlayerMeditationState) CObject() *C.PlayerMeditationState {
	object := C.New_PlayerMeditationState()
	object.Pid = C.int64_t(this.Pid)
	object.AccumulateTime = C.int32_t(this.AccumulateTime)
	object.StartTimestamp = C.int64_t(this.StartTimestamp)
	return object
}

func (this insertOP) PlayerMeditationState(row *PlayerMeditationState) {
	if this.db.tables.PlayerMeditationState != nil {
		panic("duplicate insert 'player_meditation_state'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerMeditationState = newRow
	this.db.addTransLog(this.db.newPlayerMeditationStateInsertLog(newRow, row))
}

func (this deleteOP) PlayerMeditationState(row *PlayerMeditationState) {
	if this.db.tables.PlayerMeditationState == nil {
		panic("delete not exists 'player_meditation_state'")
	}
	old := this.db.tables.PlayerMeditationState
	this.db.tables.PlayerMeditationState = nil
	this.db.addTransLog(this.db.newPlayerMeditationStateDeleteLog(old, row))
}

func (this updateOP) PlayerMeditationState(row *PlayerMeditationState) {
	if this.db.tables.PlayerMeditationState == nil {
		panic("update not exists 'player_meditation_state'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerMeditationState
	this.db.tables.PlayerMeditationState = newRow
	tmpRow := PlayerMeditationStateRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerMeditationStateUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerMeditationState(key int64) *PlayerMeditationState {
	if this.db.tables.PlayerMeditationState == nil {
		return nil
	}
	tmpRow := PlayerMeditationStateRow{row:this.db.tables.PlayerMeditationState}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_mission ========== */

// 玩家区域数据
type PlayerMission struct {
	Pid int64 // 玩家ID
	Key int32 // 拥有的区域钥匙数
	MaxOrder int8 // 已开启区域的最大序号
}

func (this *PlayerMission) CObject() *C.PlayerMission {
	object := C.New_PlayerMission()
	object.Pid = C.int64_t(this.Pid)
	object.Key = C.int32_t(this.Key)
	object.MaxOrder = C.int8_t(this.MaxOrder)
	return object
}

func (this insertOP) PlayerMission(row *PlayerMission) {
	if this.db.tables.PlayerMission != nil {
		panic("duplicate insert 'player_mission'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerMission = newRow
	this.db.addTransLog(this.db.newPlayerMissionInsertLog(newRow, row))
}

func (this deleteOP) PlayerMission(row *PlayerMission) {
	if this.db.tables.PlayerMission == nil {
		panic("delete not exists 'player_mission'")
	}
	old := this.db.tables.PlayerMission
	this.db.tables.PlayerMission = nil
	this.db.addTransLog(this.db.newPlayerMissionDeleteLog(old, row))
}

func (this updateOP) PlayerMission(row *PlayerMission) {
	if this.db.tables.PlayerMission == nil {
		panic("update not exists 'player_mission'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerMission
	this.db.tables.PlayerMission = newRow
	tmpRow := PlayerMissionRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerMissionUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerMission(key int64) *PlayerMission {
	if this.db.tables.PlayerMission == nil {
		return nil
	}
	tmpRow := PlayerMissionRow{row:this.db.tables.PlayerMission}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_mission_level ========== */

// 玩家区域关卡数据
type PlayerMissionLevel struct {
	Pid int64 // 玩家ID
	Lock int32 // 当前的关卡权值
	MaxLock int32 // 已开启的关卡最大权值
	AwardLock int32 // 已获得过奖励关卡的最大lock
}

func (this *PlayerMissionLevel) CObject() *C.PlayerMissionLevel {
	object := C.New_PlayerMissionLevel()
	object.Pid = C.int64_t(this.Pid)
	object.Lock = C.int32_t(this.Lock)
	object.MaxLock = C.int32_t(this.MaxLock)
	object.AwardLock = C.int32_t(this.AwardLock)
	return object
}

func (this insertOP) PlayerMissionLevel(row *PlayerMissionLevel) {
	if this.db.tables.PlayerMissionLevel != nil {
		panic("duplicate insert 'player_mission_level'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerMissionLevel = newRow
	this.db.addTransLog(this.db.newPlayerMissionLevelInsertLog(newRow, row))
}

func (this deleteOP) PlayerMissionLevel(row *PlayerMissionLevel) {
	if this.db.tables.PlayerMissionLevel == nil {
		panic("delete not exists 'player_mission_level'")
	}
	old := this.db.tables.PlayerMissionLevel
	this.db.tables.PlayerMissionLevel = nil
	this.db.addTransLog(this.db.newPlayerMissionLevelDeleteLog(old, row))
}

func (this updateOP) PlayerMissionLevel(row *PlayerMissionLevel) {
	if this.db.tables.PlayerMissionLevel == nil {
		panic("update not exists 'player_mission_level'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerMissionLevel
	this.db.tables.PlayerMissionLevel = newRow
	tmpRow := PlayerMissionLevelRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerMissionLevelUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerMissionLevel(key int64) *PlayerMissionLevel {
	if this.db.tables.PlayerMissionLevel == nil {
		return nil
	}
	tmpRow := PlayerMissionLevelRow{row:this.db.tables.PlayerMissionLevel}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_mission_level_record ========== */

// 关卡记录
type PlayerMissionLevelRecord struct {
	Id int64 // 记录ID
	Pid int64 // 玩家ID
	MissionId int16 // 区域ID
	MissionLevelId int32 // 开启的关卡ID
	OpenTime int64 // 关卡开启时间
	Score int32 // boss战得分
	Round int8 // 通关回合数
	DailyNum int8 // 当日已进入关卡的次数
	LastEnterTime int64 // 最后一次进入时间
	EmptyShadowBits int16 // 清剿过的影之间隙
}

func (this *PlayerMissionLevelRecord) CObject() *C.PlayerMissionLevelRecord {
	object := C.New_PlayerMissionLevelRecord()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.MissionId = C.int16_t(this.MissionId)
	object.MissionLevelId = C.int32_t(this.MissionLevelId)
	object.OpenTime = C.int64_t(this.OpenTime)
	object.Score = C.int32_t(this.Score)
	object.Round = C.int8_t(this.Round)
	object.DailyNum = C.int8_t(this.DailyNum)
	object.LastEnterTime = C.int64_t(this.LastEnterTime)
	object.EmptyShadowBits = C.int16_t(this.EmptyShadowBits)
	return object
}

func (this insertOP) PlayerMissionLevelRecord(row *PlayerMissionLevelRecord) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerMissionLevelRecord, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerMissionLevelRecord
	for crow := this.db.tables.PlayerMissionLevelRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_mission_level_record'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerMissionLevelRecord
	this.db.tables.PlayerMissionLevelRecord = newRow
	this.db.addTransLog(this.db.newPlayerMissionLevelRecordInsertLog(newRow, row))
}

func (this deleteOP) PlayerMissionLevelRecord(row *PlayerMissionLevelRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerMissionLevelRecord
	for crow := this.db.tables.PlayerMissionLevelRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_mission_level_record'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerMissionLevelRecord = this.db.tables.PlayerMissionLevelRecord.next
	}
	this.db.addTransLog(this.db.newPlayerMissionLevelRecordDeleteLog(old, row))
}

func (this updateOP) PlayerMissionLevelRecord(row *PlayerMissionLevelRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerMissionLevelRecord
	for crow := this.db.tables.PlayerMissionLevelRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_mission_level_record'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerMissionLevelRecord = newRow
	}
	tmpRow := PlayerMissionLevelRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerMissionLevelRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerMissionLevelRecord(key int64) *PlayerMissionLevelRecord {
	for crow := this.db.tables.PlayerMissionLevelRecord; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerMissionLevelRecordRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerMissionLevelRecord(callback func(*PlayerMissionLevelRecordRow)) {
	row := &PlayerMissionLevelRecordRow{}
	for crow := this.db.tables.PlayerMissionLevelRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerMissionLevelRecord() (rows []interface{}) {
	row := &PlayerMissionLevelRecordRow{}
	for crow := this.db.tables.PlayerMissionLevelRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerMissionLevelRecord", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerMissionLevelRecord", len(rows))
	return rows
}

/* ========== player_mission_level_state_bin ========== */

// 玩家区域关卡状态保存
type PlayerMissionLevelStateBin struct {
	Pid int64 // 玩家ID
	Bin []byte // 状态MissionLevelState
}

func (this *PlayerMissionLevelStateBin) CObject() *C.PlayerMissionLevelStateBin {
	object := C.New_PlayerMissionLevelStateBin()
	object.Pid = C.int64_t(this.Pid)
	object.Bin = C.CString(string(this.Bin))
	object.Bin_len = C.int(len(this.Bin))
	return object
}

func (this insertOP) PlayerMissionLevelStateBin(row *PlayerMissionLevelStateBin) {
	if this.db.tables.PlayerMissionLevelStateBin != nil {
		panic("duplicate insert 'player_mission_level_state_bin'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerMissionLevelStateBin = newRow
	this.db.addTransLog(this.db.newPlayerMissionLevelStateBinInsertLog(newRow, row))
}

func (this deleteOP) PlayerMissionLevelStateBin(row *PlayerMissionLevelStateBin) {
	if this.db.tables.PlayerMissionLevelStateBin == nil {
		panic("delete not exists 'player_mission_level_state_bin'")
	}
	old := this.db.tables.PlayerMissionLevelStateBin
	this.db.tables.PlayerMissionLevelStateBin = nil
	this.db.addTransLog(this.db.newPlayerMissionLevelStateBinDeleteLog(old, row))
}

func (this updateOP) PlayerMissionLevelStateBin(row *PlayerMissionLevelStateBin) {
	if this.db.tables.PlayerMissionLevelStateBin == nil {
		panic("update not exists 'player_mission_level_state_bin'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerMissionLevelStateBin
	this.db.tables.PlayerMissionLevelStateBin = newRow
	tmpRow := PlayerMissionLevelStateBinRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerMissionLevelStateBinUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerMissionLevelStateBin(key int64) *PlayerMissionLevelStateBin {
	if this.db.tables.PlayerMissionLevelStateBin == nil {
		return nil
	}
	tmpRow := PlayerMissionLevelStateBinRow{row:this.db.tables.PlayerMissionLevelStateBin}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_mission_record ========== */

// 区域记录
type PlayerMissionRecord struct {
	Id int64 // 记录ID
	Pid int64 // 玩家ID
	TownId int16 // 城镇ID
	MissionId int16 // 开启的区域ID
	OpenTime int64 // 开启的区域时间
}

func (this *PlayerMissionRecord) CObject() *C.PlayerMissionRecord {
	object := C.New_PlayerMissionRecord()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.TownId = C.int16_t(this.TownId)
	object.MissionId = C.int16_t(this.MissionId)
	object.OpenTime = C.int64_t(this.OpenTime)
	return object
}

func (this insertOP) PlayerMissionRecord(row *PlayerMissionRecord) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerMissionRecord, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerMissionRecord
	for crow := this.db.tables.PlayerMissionRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_mission_record'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerMissionRecord
	this.db.tables.PlayerMissionRecord = newRow
	this.db.addTransLog(this.db.newPlayerMissionRecordInsertLog(newRow, row))
}

func (this deleteOP) PlayerMissionRecord(row *PlayerMissionRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerMissionRecord
	for crow := this.db.tables.PlayerMissionRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_mission_record'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerMissionRecord = this.db.tables.PlayerMissionRecord.next
	}
	this.db.addTransLog(this.db.newPlayerMissionRecordDeleteLog(old, row))
}

func (this updateOP) PlayerMissionRecord(row *PlayerMissionRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerMissionRecord
	for crow := this.db.tables.PlayerMissionRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_mission_record'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerMissionRecord = newRow
	}
	tmpRow := PlayerMissionRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerMissionRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerMissionRecord(key int64) *PlayerMissionRecord {
	for crow := this.db.tables.PlayerMissionRecord; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerMissionRecordRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerMissionRecord(callback func(*PlayerMissionRecordRow)) {
	row := &PlayerMissionRecordRow{}
	for crow := this.db.tables.PlayerMissionRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerMissionRecord() (rows []interface{}) {
	row := &PlayerMissionRecordRow{}
	for crow := this.db.tables.PlayerMissionRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerMissionRecord", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerMissionRecord", len(rows))
	return rows
}

/* ========== player_mission_star_award ========== */

// 玩家区域评星领奖记录
type PlayerMissionStarAward struct {
	Id int64 // 主键
	Pid int64 // 玩家ID
	TownId int16 // 阵印ID
	BoxType int8 // 宝箱类型 1:铜 2:银 3:金
}

func (this *PlayerMissionStarAward) CObject() *C.PlayerMissionStarAward {
	object := C.New_PlayerMissionStarAward()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.TownId = C.int16_t(this.TownId)
	object.BoxType = C.int8_t(this.BoxType)
	return object
}

func (this insertOP) PlayerMissionStarAward(row *PlayerMissionStarAward) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerMissionStarAward, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerMissionStarAward
	for crow := this.db.tables.PlayerMissionStarAward; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_mission_star_award'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerMissionStarAward
	this.db.tables.PlayerMissionStarAward = newRow
	this.db.addTransLog(this.db.newPlayerMissionStarAwardInsertLog(newRow, row))
}

func (this deleteOP) PlayerMissionStarAward(row *PlayerMissionStarAward) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerMissionStarAward
	for crow := this.db.tables.PlayerMissionStarAward; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_mission_star_award'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerMissionStarAward = this.db.tables.PlayerMissionStarAward.next
	}
	this.db.addTransLog(this.db.newPlayerMissionStarAwardDeleteLog(old, row))
}

func (this updateOP) PlayerMissionStarAward(row *PlayerMissionStarAward) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerMissionStarAward
	for crow := this.db.tables.PlayerMissionStarAward; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_mission_star_award'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerMissionStarAward = newRow
	}
	tmpRow := PlayerMissionStarAwardRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerMissionStarAwardUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerMissionStarAward(key int64) *PlayerMissionStarAward {
	for crow := this.db.tables.PlayerMissionStarAward; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerMissionStarAwardRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerMissionStarAward(callback func(*PlayerMissionStarAwardRow)) {
	row := &PlayerMissionStarAwardRow{}
	for crow := this.db.tables.PlayerMissionStarAward; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerMissionStarAward() (rows []interface{}) {
	row := &PlayerMissionStarAwardRow{}
	for crow := this.db.tables.PlayerMissionStarAward; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerMissionStarAward", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerMissionStarAward", len(rows))
	return rows
}

/* ========== player_money_tree ========== */

// 玩家摇钱树记录
type PlayerMoneyTree struct {
	Pid int64 // 玩家id
	Total int32 // 摇钱树铜钱总额
	WavedTimesTotal int8 // 总的摇树次数
	WavedTimes int8 // 当天已经摇过的次数
	LastWavedTime int64 // 上次摇树的时间戳
}

func (this *PlayerMoneyTree) CObject() *C.PlayerMoneyTree {
	object := C.New_PlayerMoneyTree()
	object.Pid = C.int64_t(this.Pid)
	object.Total = C.int32_t(this.Total)
	object.WavedTimesTotal = C.int8_t(this.WavedTimesTotal)
	object.WavedTimes = C.int8_t(this.WavedTimes)
	object.LastWavedTime = C.int64_t(this.LastWavedTime)
	return object
}

func (this insertOP) PlayerMoneyTree(row *PlayerMoneyTree) {
	if this.db.tables.PlayerMoneyTree != nil {
		panic("duplicate insert 'player_money_tree'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerMoneyTree = newRow
	this.db.addTransLog(this.db.newPlayerMoneyTreeInsertLog(newRow, row))
}

func (this deleteOP) PlayerMoneyTree(row *PlayerMoneyTree) {
	if this.db.tables.PlayerMoneyTree == nil {
		panic("delete not exists 'player_money_tree'")
	}
	old := this.db.tables.PlayerMoneyTree
	this.db.tables.PlayerMoneyTree = nil
	this.db.addTransLog(this.db.newPlayerMoneyTreeDeleteLog(old, row))
}

func (this updateOP) PlayerMoneyTree(row *PlayerMoneyTree) {
	if this.db.tables.PlayerMoneyTree == nil {
		panic("update not exists 'player_money_tree'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerMoneyTree
	this.db.tables.PlayerMoneyTree = newRow
	tmpRow := PlayerMoneyTreeRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerMoneyTreeUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerMoneyTree(key int64) *PlayerMoneyTree {
	if this.db.tables.PlayerMoneyTree == nil {
		return nil
	}
	tmpRow := PlayerMoneyTreeRow{row:this.db.tables.PlayerMoneyTree}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_month_card_award_record ========== */

// 玩家最后领取月卡时间
type PlayerMonthCardAwardRecord struct {
	Pid int64 // 用户ID
	LastUpdate int64 // 最后一次更新时间戳
}

func (this *PlayerMonthCardAwardRecord) CObject() *C.PlayerMonthCardAwardRecord {
	object := C.New_PlayerMonthCardAwardRecord()
	object.Pid = C.int64_t(this.Pid)
	object.LastUpdate = C.int64_t(this.LastUpdate)
	return object
}

func (this insertOP) PlayerMonthCardAwardRecord(row *PlayerMonthCardAwardRecord) {
	if this.db.tables.PlayerMonthCardAwardRecord != nil {
		panic("duplicate insert 'player_month_card_award_record'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerMonthCardAwardRecord = newRow
	this.db.addTransLog(this.db.newPlayerMonthCardAwardRecordInsertLog(newRow, row))
}

func (this deleteOP) PlayerMonthCardAwardRecord(row *PlayerMonthCardAwardRecord) {
	if this.db.tables.PlayerMonthCardAwardRecord == nil {
		panic("delete not exists 'player_month_card_award_record'")
	}
	old := this.db.tables.PlayerMonthCardAwardRecord
	this.db.tables.PlayerMonthCardAwardRecord = nil
	this.db.addTransLog(this.db.newPlayerMonthCardAwardRecordDeleteLog(old, row))
}

func (this updateOP) PlayerMonthCardAwardRecord(row *PlayerMonthCardAwardRecord) {
	if this.db.tables.PlayerMonthCardAwardRecord == nil {
		panic("update not exists 'player_month_card_award_record'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerMonthCardAwardRecord
	this.db.tables.PlayerMonthCardAwardRecord = newRow
	tmpRow := PlayerMonthCardAwardRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerMonthCardAwardRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerMonthCardAwardRecord(key int64) *PlayerMonthCardAwardRecord {
	if this.db.tables.PlayerMonthCardAwardRecord == nil {
		return nil
	}
	tmpRow := PlayerMonthCardAwardRecordRow{row:this.db.tables.PlayerMonthCardAwardRecord}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_month_card_info ========== */

// 玩家月卡信息
type PlayerMonthCardInfo struct {
	Pid int64 // 用户ID
	Starttime int64 // 月卡开始时间
	Endtime int64 // 月卡结束时间
	Buytimes int64 // 购买次数
	Presenttotal int64 // 赠送总金额
}

func (this *PlayerMonthCardInfo) CObject() *C.PlayerMonthCardInfo {
	object := C.New_PlayerMonthCardInfo()
	object.Pid = C.int64_t(this.Pid)
	object.Starttime = C.int64_t(this.Starttime)
	object.Endtime = C.int64_t(this.Endtime)
	object.Buytimes = C.int64_t(this.Buytimes)
	object.Presenttotal = C.int64_t(this.Presenttotal)
	return object
}

func (this insertOP) PlayerMonthCardInfo(row *PlayerMonthCardInfo) {
	if this.db.tables.PlayerMonthCardInfo != nil {
		panic("duplicate insert 'player_month_card_info'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerMonthCardInfo = newRow
	this.db.addTransLog(this.db.newPlayerMonthCardInfoInsertLog(newRow, row))
}

func (this deleteOP) PlayerMonthCardInfo(row *PlayerMonthCardInfo) {
	if this.db.tables.PlayerMonthCardInfo == nil {
		panic("delete not exists 'player_month_card_info'")
	}
	old := this.db.tables.PlayerMonthCardInfo
	this.db.tables.PlayerMonthCardInfo = nil
	this.db.addTransLog(this.db.newPlayerMonthCardInfoDeleteLog(old, row))
}

func (this updateOP) PlayerMonthCardInfo(row *PlayerMonthCardInfo) {
	if this.db.tables.PlayerMonthCardInfo == nil {
		panic("update not exists 'player_month_card_info'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerMonthCardInfo
	this.db.tables.PlayerMonthCardInfo = newRow
	tmpRow := PlayerMonthCardInfoRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerMonthCardInfoUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerMonthCardInfo(key int64) *PlayerMonthCardInfo {
	if this.db.tables.PlayerMonthCardInfo == nil {
		return nil
	}
	tmpRow := PlayerMonthCardInfoRow{row:this.db.tables.PlayerMonthCardInfo}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_multi_level_info ========== */

// 玩家多人关卡信息
type PlayerMultiLevelInfo struct {
	Pid int64 // 玩家ID
	BuddyRoleId int8 // 上阵伙伴角色模板ID
	BuddyRow int8 // 上阵伙伴所在行（1或2)
	TacticalGrid int8 // 战术
	DailyNum int8 // 今日已战斗次数
	BattleTime int64 // 最近一次战斗时间
	Lock int32 // 关卡开启权值
}

func (this *PlayerMultiLevelInfo) CObject() *C.PlayerMultiLevelInfo {
	object := C.New_PlayerMultiLevelInfo()
	object.Pid = C.int64_t(this.Pid)
	object.BuddyRoleId = C.int8_t(this.BuddyRoleId)
	object.BuddyRow = C.int8_t(this.BuddyRow)
	object.TacticalGrid = C.int8_t(this.TacticalGrid)
	object.DailyNum = C.int8_t(this.DailyNum)
	object.BattleTime = C.int64_t(this.BattleTime)
	object.Lock = C.int32_t(this.Lock)
	return object
}

func (this insertOP) PlayerMultiLevelInfo(row *PlayerMultiLevelInfo) {
	if this.db.tables.PlayerMultiLevelInfo != nil {
		panic("duplicate insert 'player_multi_level_info'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerMultiLevelInfo = newRow
	this.db.addTransLog(this.db.newPlayerMultiLevelInfoInsertLog(newRow, row))
}

func (this deleteOP) PlayerMultiLevelInfo(row *PlayerMultiLevelInfo) {
	if this.db.tables.PlayerMultiLevelInfo == nil {
		panic("delete not exists 'player_multi_level_info'")
	}
	old := this.db.tables.PlayerMultiLevelInfo
	this.db.tables.PlayerMultiLevelInfo = nil
	this.db.addTransLog(this.db.newPlayerMultiLevelInfoDeleteLog(old, row))
}

func (this updateOP) PlayerMultiLevelInfo(row *PlayerMultiLevelInfo) {
	if this.db.tables.PlayerMultiLevelInfo == nil {
		panic("update not exists 'player_multi_level_info'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerMultiLevelInfo
	this.db.tables.PlayerMultiLevelInfo = newRow
	tmpRow := PlayerMultiLevelInfoRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerMultiLevelInfoUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerMultiLevelInfo(key int64) *PlayerMultiLevelInfo {
	if this.db.tables.PlayerMultiLevelInfo == nil {
		return nil
	}
	tmpRow := PlayerMultiLevelInfoRow{row:this.db.tables.PlayerMultiLevelInfo}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_new_year_consume_record ========== */

// 春节玩家消费记录
type PlayerNewYearConsumeRecord struct {
	Pid int64 // 玩家id
	ConsumeRecord string // 玩家消费情况
}

func (this *PlayerNewYearConsumeRecord) CObject() *C.PlayerNewYearConsumeRecord {
	object := C.New_PlayerNewYearConsumeRecord()
	object.Pid = C.int64_t(this.Pid)
	object.ConsumeRecord = C.CString(string(this.ConsumeRecord))
	object.ConsumeRecord_len = C.int(len(this.ConsumeRecord))
	return object
}

func (this insertOP) PlayerNewYearConsumeRecord(row *PlayerNewYearConsumeRecord) {
	if this.db.tables.PlayerNewYearConsumeRecord != nil {
		panic("duplicate insert 'player_new_year_consume_record'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerNewYearConsumeRecord = newRow
	this.db.addTransLog(this.db.newPlayerNewYearConsumeRecordInsertLog(newRow, row))
}

func (this deleteOP) PlayerNewYearConsumeRecord(row *PlayerNewYearConsumeRecord) {
	if this.db.tables.PlayerNewYearConsumeRecord == nil {
		panic("delete not exists 'player_new_year_consume_record'")
	}
	old := this.db.tables.PlayerNewYearConsumeRecord
	this.db.tables.PlayerNewYearConsumeRecord = nil
	this.db.addTransLog(this.db.newPlayerNewYearConsumeRecordDeleteLog(old, row))
}

func (this updateOP) PlayerNewYearConsumeRecord(row *PlayerNewYearConsumeRecord) {
	if this.db.tables.PlayerNewYearConsumeRecord == nil {
		panic("update not exists 'player_new_year_consume_record'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerNewYearConsumeRecord
	this.db.tables.PlayerNewYearConsumeRecord = newRow
	tmpRow := PlayerNewYearConsumeRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerNewYearConsumeRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerNewYearConsumeRecord(key int64) *PlayerNewYearConsumeRecord {
	if this.db.tables.PlayerNewYearConsumeRecord == nil {
		return nil
	}
	tmpRow := PlayerNewYearConsumeRecordRow{row:this.db.tables.PlayerNewYearConsumeRecord}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_npc_talk_record ========== */

// 玩家与NPC对话奖励记录
type PlayerNpcTalkRecord struct {
	Id int64 // 记录ID
	Pid int64 // 玩家ID
	NpcId int32 // NPC ID
	TownId int16 // 相关城镇
	QuestId int16 // 任务ID  首次对话作为特殊任务ID -1
}

func (this *PlayerNpcTalkRecord) CObject() *C.PlayerNpcTalkRecord {
	object := C.New_PlayerNpcTalkRecord()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.NpcId = C.int32_t(this.NpcId)
	object.TownId = C.int16_t(this.TownId)
	object.QuestId = C.int16_t(this.QuestId)
	return object
}

func (this insertOP) PlayerNpcTalkRecord(row *PlayerNpcTalkRecord) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerNpcTalkRecord, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerNpcTalkRecord
	for crow := this.db.tables.PlayerNpcTalkRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_npc_talk_record'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerNpcTalkRecord
	this.db.tables.PlayerNpcTalkRecord = newRow
	this.db.addTransLog(this.db.newPlayerNpcTalkRecordInsertLog(newRow, row))
}

func (this deleteOP) PlayerNpcTalkRecord(row *PlayerNpcTalkRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerNpcTalkRecord
	for crow := this.db.tables.PlayerNpcTalkRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_npc_talk_record'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerNpcTalkRecord = this.db.tables.PlayerNpcTalkRecord.next
	}
	this.db.addTransLog(this.db.newPlayerNpcTalkRecordDeleteLog(old, row))
}

func (this updateOP) PlayerNpcTalkRecord(row *PlayerNpcTalkRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerNpcTalkRecord
	for crow := this.db.tables.PlayerNpcTalkRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_npc_talk_record'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerNpcTalkRecord = newRow
	}
	tmpRow := PlayerNpcTalkRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerNpcTalkRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerNpcTalkRecord(key int64) *PlayerNpcTalkRecord {
	for crow := this.db.tables.PlayerNpcTalkRecord; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerNpcTalkRecordRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerNpcTalkRecord(callback func(*PlayerNpcTalkRecordRow)) {
	row := &PlayerNpcTalkRecordRow{}
	for crow := this.db.tables.PlayerNpcTalkRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerNpcTalkRecord() (rows []interface{}) {
	row := &PlayerNpcTalkRecordRow{}
	for crow := this.db.tables.PlayerNpcTalkRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerNpcTalkRecord", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerNpcTalkRecord", len(rows))
	return rows
}

/* ========== player_opened_town_treasure ========== */

// 玩家已开启的城镇宝箱
type PlayerOpenedTownTreasure struct {
	Id int64 // 主键
	Pid int64 // 玩家ID
	TownId int16 // 城镇ID
}

func (this *PlayerOpenedTownTreasure) CObject() *C.PlayerOpenedTownTreasure {
	object := C.New_PlayerOpenedTownTreasure()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.TownId = C.int16_t(this.TownId)
	return object
}

func (this insertOP) PlayerOpenedTownTreasure(row *PlayerOpenedTownTreasure) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerOpenedTownTreasure, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerOpenedTownTreasure
	for crow := this.db.tables.PlayerOpenedTownTreasure; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_opened_town_treasure'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerOpenedTownTreasure
	this.db.tables.PlayerOpenedTownTreasure = newRow
	this.db.addTransLog(this.db.newPlayerOpenedTownTreasureInsertLog(newRow, row))
}

func (this deleteOP) PlayerOpenedTownTreasure(row *PlayerOpenedTownTreasure) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerOpenedTownTreasure
	for crow := this.db.tables.PlayerOpenedTownTreasure; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_opened_town_treasure'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerOpenedTownTreasure = this.db.tables.PlayerOpenedTownTreasure.next
	}
	this.db.addTransLog(this.db.newPlayerOpenedTownTreasureDeleteLog(old, row))
}

func (this updateOP) PlayerOpenedTownTreasure(row *PlayerOpenedTownTreasure) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerOpenedTownTreasure
	for crow := this.db.tables.PlayerOpenedTownTreasure; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_opened_town_treasure'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerOpenedTownTreasure = newRow
	}
	tmpRow := PlayerOpenedTownTreasureRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerOpenedTownTreasureUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerOpenedTownTreasure(key int64) *PlayerOpenedTownTreasure {
	for crow := this.db.tables.PlayerOpenedTownTreasure; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerOpenedTownTreasureRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerOpenedTownTreasure(callback func(*PlayerOpenedTownTreasureRow)) {
	row := &PlayerOpenedTownTreasureRow{}
	for crow := this.db.tables.PlayerOpenedTownTreasure; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerOpenedTownTreasure() (rows []interface{}) {
	row := &PlayerOpenedTownTreasureRow{}
	for crow := this.db.tables.PlayerOpenedTownTreasure; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerOpenedTownTreasure", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerOpenedTownTreasure", len(rows))
	return rows
}

/* ========== player_physical ========== */

// 玩家体力表
type PlayerPhysical struct {
	Pid int64 // 玩家ID
	Value int16 // 体力值
	UpdateTime int64 // 最后更新时间
	BuyCount int64 // 购买次数
	BuyUpdateTime int64 // 购买次数更新时间
	DailyCount int16 // 当天购买次数
}

func (this *PlayerPhysical) CObject() *C.PlayerPhysical {
	object := C.New_PlayerPhysical()
	object.Pid = C.int64_t(this.Pid)
	object.Value = C.int16_t(this.Value)
	object.UpdateTime = C.int64_t(this.UpdateTime)
	object.BuyCount = C.int64_t(this.BuyCount)
	object.BuyUpdateTime = C.int64_t(this.BuyUpdateTime)
	object.DailyCount = C.int16_t(this.DailyCount)
	return object
}

func (this insertOP) PlayerPhysical(row *PlayerPhysical) {
	if this.db.tables.PlayerPhysical != nil {
		panic("duplicate insert 'player_physical'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerPhysical = newRow
	this.db.addTransLog(this.db.newPlayerPhysicalInsertLog(newRow, row))
}

func (this deleteOP) PlayerPhysical(row *PlayerPhysical) {
	if this.db.tables.PlayerPhysical == nil {
		panic("delete not exists 'player_physical'")
	}
	old := this.db.tables.PlayerPhysical
	this.db.tables.PlayerPhysical = nil
	this.db.addTransLog(this.db.newPlayerPhysicalDeleteLog(old, row))
}

func (this updateOP) PlayerPhysical(row *PlayerPhysical) {
	if this.db.tables.PlayerPhysical == nil {
		panic("update not exists 'player_physical'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerPhysical
	this.db.tables.PlayerPhysical = newRow
	tmpRow := PlayerPhysicalRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerPhysicalUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerPhysical(key int64) *PlayerPhysical {
	if this.db.tables.PlayerPhysical == nil {
		return nil
	}
	tmpRow := PlayerPhysicalRow{row:this.db.tables.PlayerPhysical}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_purchase_record ========== */

// 玩家物品购买记录 仅记录限购商品
type PlayerPurchaseRecord struct {
	Id int64 // 主键ID
	Pid int64 // 玩家ID
	ItemId int16 // 物品ID
	Num int16 // 已购买数量
	Timestamp int64 // 上次购买时间
}

func (this *PlayerPurchaseRecord) CObject() *C.PlayerPurchaseRecord {
	object := C.New_PlayerPurchaseRecord()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.ItemId = C.int16_t(this.ItemId)
	object.Num = C.int16_t(this.Num)
	object.Timestamp = C.int64_t(this.Timestamp)
	return object
}

func (this insertOP) PlayerPurchaseRecord(row *PlayerPurchaseRecord) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerPurchaseRecord, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerPurchaseRecord
	for crow := this.db.tables.PlayerPurchaseRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_purchase_record'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerPurchaseRecord
	this.db.tables.PlayerPurchaseRecord = newRow
	this.db.addTransLog(this.db.newPlayerPurchaseRecordInsertLog(newRow, row))
}

func (this deleteOP) PlayerPurchaseRecord(row *PlayerPurchaseRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerPurchaseRecord
	for crow := this.db.tables.PlayerPurchaseRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_purchase_record'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerPurchaseRecord = this.db.tables.PlayerPurchaseRecord.next
	}
	this.db.addTransLog(this.db.newPlayerPurchaseRecordDeleteLog(old, row))
}

func (this updateOP) PlayerPurchaseRecord(row *PlayerPurchaseRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerPurchaseRecord
	for crow := this.db.tables.PlayerPurchaseRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_purchase_record'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerPurchaseRecord = newRow
	}
	tmpRow := PlayerPurchaseRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerPurchaseRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerPurchaseRecord(key int64) *PlayerPurchaseRecord {
	for crow := this.db.tables.PlayerPurchaseRecord; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerPurchaseRecordRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerPurchaseRecord(callback func(*PlayerPurchaseRecordRow)) {
	row := &PlayerPurchaseRecordRow{}
	for crow := this.db.tables.PlayerPurchaseRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerPurchaseRecord() (rows []interface{}) {
	row := &PlayerPurchaseRecordRow{}
	for crow := this.db.tables.PlayerPurchaseRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerPurchaseRecord", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerPurchaseRecord", len(rows))
	return rows
}

/* ========== player_push_notify_switch ========== */

// 玩家推送通知开关列表
type PlayerPushNotifySwitch struct {
	Pid int64 // 玩家ID
	Options int64 // 推送通知开关
}

func (this *PlayerPushNotifySwitch) CObject() *C.PlayerPushNotifySwitch {
	object := C.New_PlayerPushNotifySwitch()
	object.Pid = C.int64_t(this.Pid)
	object.Options = C.int64_t(this.Options)
	return object
}

func (this insertOP) PlayerPushNotifySwitch(row *PlayerPushNotifySwitch) {
	if this.db.tables.PlayerPushNotifySwitch != nil {
		panic("duplicate insert 'player_push_notify_switch'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerPushNotifySwitch = newRow
	this.db.addTransLog(this.db.newPlayerPushNotifySwitchInsertLog(newRow, row))
}

func (this deleteOP) PlayerPushNotifySwitch(row *PlayerPushNotifySwitch) {
	if this.db.tables.PlayerPushNotifySwitch == nil {
		panic("delete not exists 'player_push_notify_switch'")
	}
	old := this.db.tables.PlayerPushNotifySwitch
	this.db.tables.PlayerPushNotifySwitch = nil
	this.db.addTransLog(this.db.newPlayerPushNotifySwitchDeleteLog(old, row))
}

func (this updateOP) PlayerPushNotifySwitch(row *PlayerPushNotifySwitch) {
	if this.db.tables.PlayerPushNotifySwitch == nil {
		panic("update not exists 'player_push_notify_switch'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerPushNotifySwitch
	this.db.tables.PlayerPushNotifySwitch = newRow
	tmpRow := PlayerPushNotifySwitchRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerPushNotifySwitchUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerPushNotifySwitch(key int64) *PlayerPushNotifySwitch {
	if this.db.tables.PlayerPushNotifySwitch == nil {
		return nil
	}
	tmpRow := PlayerPushNotifySwitchRow{row:this.db.tables.PlayerPushNotifySwitch}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_pve_state ========== */

// 玩家灵宠状态
type PlayerPveState struct {
	Pid int64 // 玩家ID
	MaxPassedFloor int16 // 已通关最大层数
	MaxAwardedFloor int16 // 已奖励最大层数
	UnpassedFloorEnemyNum int16 // 未通关关卡杀敌数
	EnterTime int64 // 最后一次进入关卡次数
	DailyNum int8 // 今日进入次数
}

func (this *PlayerPveState) CObject() *C.PlayerPveState {
	object := C.New_PlayerPveState()
	object.Pid = C.int64_t(this.Pid)
	object.MaxPassedFloor = C.int16_t(this.MaxPassedFloor)
	object.MaxAwardedFloor = C.int16_t(this.MaxAwardedFloor)
	object.UnpassedFloorEnemyNum = C.int16_t(this.UnpassedFloorEnemyNum)
	object.EnterTime = C.int64_t(this.EnterTime)
	object.DailyNum = C.int8_t(this.DailyNum)
	return object
}

func (this insertOP) PlayerPveState(row *PlayerPveState) {
	if this.db.tables.PlayerPveState != nil {
		panic("duplicate insert 'player_pve_state'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerPveState = newRow
	this.db.addTransLog(this.db.newPlayerPveStateInsertLog(newRow, row))
}

func (this deleteOP) PlayerPveState(row *PlayerPveState) {
	if this.db.tables.PlayerPveState == nil {
		panic("delete not exists 'player_pve_state'")
	}
	old := this.db.tables.PlayerPveState
	this.db.tables.PlayerPveState = nil
	this.db.addTransLog(this.db.newPlayerPveStateDeleteLog(old, row))
}

func (this updateOP) PlayerPveState(row *PlayerPveState) {
	if this.db.tables.PlayerPveState == nil {
		panic("update not exists 'player_pve_state'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerPveState
	this.db.tables.PlayerPveState = newRow
	tmpRow := PlayerPveStateRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerPveStateUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerPveState(key int64) *PlayerPveState {
	if this.db.tables.PlayerPveState == nil {
		return nil
	}
	tmpRow := PlayerPveStateRow{row:this.db.tables.PlayerPveState}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_qq_vip_gift ========== */

// 玩家qq服务礼包领取记录
type PlayerQqVipGift struct {
	Pid int64 // 
	Qqvip int16 // qq会员的礼包领取记录，1代表开通礼包已领取，2代表续费礼包已领取
	Surper int16 // qq超级会员礼包领取记录，值同上
}

func (this *PlayerQqVipGift) CObject() *C.PlayerQqVipGift {
	object := C.New_PlayerQqVipGift()
	object.Pid = C.int64_t(this.Pid)
	object.Qqvip = C.int16_t(this.Qqvip)
	object.Surper = C.int16_t(this.Surper)
	return object
}

func (this insertOP) PlayerQqVipGift(row *PlayerQqVipGift) {
	if this.db.tables.PlayerQqVipGift != nil {
		panic("duplicate insert 'player_qq_vip_gift'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerQqVipGift = newRow
	this.db.addTransLog(this.db.newPlayerQqVipGiftInsertLog(newRow, row))
}

func (this deleteOP) PlayerQqVipGift(row *PlayerQqVipGift) {
	if this.db.tables.PlayerQqVipGift == nil {
		panic("delete not exists 'player_qq_vip_gift'")
	}
	old := this.db.tables.PlayerQqVipGift
	this.db.tables.PlayerQqVipGift = nil
	this.db.addTransLog(this.db.newPlayerQqVipGiftDeleteLog(old, row))
}

func (this updateOP) PlayerQqVipGift(row *PlayerQqVipGift) {
	if this.db.tables.PlayerQqVipGift == nil {
		panic("update not exists 'player_qq_vip_gift'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerQqVipGift
	this.db.tables.PlayerQqVipGift = newRow
	tmpRow := PlayerQqVipGiftRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerQqVipGiftUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerQqVipGift(key int64) *PlayerQqVipGift {
	if this.db.tables.PlayerQqVipGift == nil {
		return nil
	}
	tmpRow := PlayerQqVipGiftRow{row:this.db.tables.PlayerQqVipGift}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_quest ========== */

// 玩家任务
type PlayerQuest struct {
	Pid int64 // 玩家ID
	QuestId int16 // 当前任务ID
	State int8 // 
}

func (this *PlayerQuest) CObject() *C.PlayerQuest {
	object := C.New_PlayerQuest()
	object.Pid = C.int64_t(this.Pid)
	object.QuestId = C.int16_t(this.QuestId)
	object.State = C.int8_t(this.State)
	return object
}

func (this insertOP) PlayerQuest(row *PlayerQuest) {
	if this.db.tables.PlayerQuest != nil {
		panic("duplicate insert 'player_quest'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerQuest = newRow
	this.db.addTransLog(this.db.newPlayerQuestInsertLog(newRow, row))
}

func (this deleteOP) PlayerQuest(row *PlayerQuest) {
	if this.db.tables.PlayerQuest == nil {
		panic("delete not exists 'player_quest'")
	}
	old := this.db.tables.PlayerQuest
	this.db.tables.PlayerQuest = nil
	this.db.addTransLog(this.db.newPlayerQuestDeleteLog(old, row))
}

func (this updateOP) PlayerQuest(row *PlayerQuest) {
	if this.db.tables.PlayerQuest == nil {
		panic("update not exists 'player_quest'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerQuest
	this.db.tables.PlayerQuest = newRow
	tmpRow := PlayerQuestRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerQuestUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerQuest(key int64) *PlayerQuest {
	if this.db.tables.PlayerQuest == nil {
		return nil
	}
	tmpRow := PlayerQuestRow{row:this.db.tables.PlayerQuest}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_rainbow_level ========== */

// 玩家彩虹关卡状态
type PlayerRainbowLevel struct {
	Pid int64 // 玩家ID
	ResetTimestamp int64 // 重置彩虹关卡时间戳
	ResetNum int32 // 已重置次数
	Segment int16 // 段数
	Order int8 // 彩虹关段内第X关顺序[1,7]
	Status int8 // 状态 0--打败Boss  1--未领取奖励 2--已奖励未进入下一关卡
	BestSegment int16 // 最好记录段数
	BestOrder int8 // 最好记录关卡顺序
	BestRecordTimestamp int64 // 最好记录时间戳
	MaxOpenSegment int16 // 可跳转的最大段数
	MaxPassSegment int16 // 打通的最高段数
	AutoFightNum int8 // 今日扫荡次数
	BuyTimes int16 // 今日购买次数
	AutoFightTime int64 // 扫荡时间
	BuyTimestamp int64 // 购买彩虹关卡次数时间戳
}

func (this *PlayerRainbowLevel) CObject() *C.PlayerRainbowLevel {
	object := C.New_PlayerRainbowLevel()
	object.Pid = C.int64_t(this.Pid)
	object.ResetTimestamp = C.int64_t(this.ResetTimestamp)
	object.ResetNum = C.int32_t(this.ResetNum)
	object.Segment = C.int16_t(this.Segment)
	object.Order = C.int8_t(this.Order)
	object.Status = C.int8_t(this.Status)
	object.BestSegment = C.int16_t(this.BestSegment)
	object.BestOrder = C.int8_t(this.BestOrder)
	object.BestRecordTimestamp = C.int64_t(this.BestRecordTimestamp)
	object.MaxOpenSegment = C.int16_t(this.MaxOpenSegment)
	object.MaxPassSegment = C.int16_t(this.MaxPassSegment)
	object.AutoFightNum = C.int8_t(this.AutoFightNum)
	object.BuyTimes = C.int16_t(this.BuyTimes)
	object.AutoFightTime = C.int64_t(this.AutoFightTime)
	object.BuyTimestamp = C.int64_t(this.BuyTimestamp)
	return object
}

func (this insertOP) PlayerRainbowLevel(row *PlayerRainbowLevel) {
	if this.db.tables.PlayerRainbowLevel != nil {
		panic("duplicate insert 'player_rainbow_level'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerRainbowLevel = newRow
	this.db.addTransLog(this.db.newPlayerRainbowLevelInsertLog(newRow, row))
}

func (this deleteOP) PlayerRainbowLevel(row *PlayerRainbowLevel) {
	if this.db.tables.PlayerRainbowLevel == nil {
		panic("delete not exists 'player_rainbow_level'")
	}
	old := this.db.tables.PlayerRainbowLevel
	this.db.tables.PlayerRainbowLevel = nil
	this.db.addTransLog(this.db.newPlayerRainbowLevelDeleteLog(old, row))
}

func (this updateOP) PlayerRainbowLevel(row *PlayerRainbowLevel) {
	if this.db.tables.PlayerRainbowLevel == nil {
		panic("update not exists 'player_rainbow_level'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerRainbowLevel
	this.db.tables.PlayerRainbowLevel = newRow
	tmpRow := PlayerRainbowLevelRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerRainbowLevelUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerRainbowLevel(key int64) *PlayerRainbowLevel {
	if this.db.tables.PlayerRainbowLevel == nil {
		return nil
	}
	tmpRow := PlayerRainbowLevelRow{row:this.db.tables.PlayerRainbowLevel}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_rainbow_level_state_bin ========== */

// 玩家彩虹关卡状态
type PlayerRainbowLevelStateBin struct {
	Pid int64 // 玩家ID
	Bin []byte // 彩虹状态
}

func (this *PlayerRainbowLevelStateBin) CObject() *C.PlayerRainbowLevelStateBin {
	object := C.New_PlayerRainbowLevelStateBin()
	object.Pid = C.int64_t(this.Pid)
	object.Bin = C.CString(string(this.Bin))
	object.Bin_len = C.int(len(this.Bin))
	return object
}

func (this insertOP) PlayerRainbowLevelStateBin(row *PlayerRainbowLevelStateBin) {
	if this.db.tables.PlayerRainbowLevelStateBin != nil {
		panic("duplicate insert 'player_rainbow_level_state_bin'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerRainbowLevelStateBin = newRow
	this.db.addTransLog(this.db.newPlayerRainbowLevelStateBinInsertLog(newRow, row))
}

func (this deleteOP) PlayerRainbowLevelStateBin(row *PlayerRainbowLevelStateBin) {
	if this.db.tables.PlayerRainbowLevelStateBin == nil {
		panic("delete not exists 'player_rainbow_level_state_bin'")
	}
	old := this.db.tables.PlayerRainbowLevelStateBin
	this.db.tables.PlayerRainbowLevelStateBin = nil
	this.db.addTransLog(this.db.newPlayerRainbowLevelStateBinDeleteLog(old, row))
}

func (this updateOP) PlayerRainbowLevelStateBin(row *PlayerRainbowLevelStateBin) {
	if this.db.tables.PlayerRainbowLevelStateBin == nil {
		panic("update not exists 'player_rainbow_level_state_bin'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerRainbowLevelStateBin
	this.db.tables.PlayerRainbowLevelStateBin = newRow
	tmpRow := PlayerRainbowLevelStateBinRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerRainbowLevelStateBinUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerRainbowLevelStateBin(key int64) *PlayerRainbowLevelStateBin {
	if this.db.tables.PlayerRainbowLevelStateBin == nil {
		return nil
	}
	tmpRow := PlayerRainbowLevelStateBinRow{row:this.db.tables.PlayerRainbowLevelStateBin}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_role ========== */

// 玩家角色数据
type PlayerRole struct {
	Id int64 // 玩家角色ID
	Pid int64 // 玩家ID
	RoleId int8 // 角色模板ID
	Level int16 // 等级
	Exp int64 // 经验
	FriendshipLevel int32 // 角色的羁绊等级
	Status int16 // 伙伴状态，0表示正常，1表示在客栈
	CoopId int16 // 协力组合
	Timestamp int64 // 等级变化时间
	FightNum int32 // 角色战力
}

func (this *PlayerRole) CObject() *C.PlayerRole {
	object := C.New_PlayerRole()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.RoleId = C.int8_t(this.RoleId)
	object.Level = C.int16_t(this.Level)
	object.Exp = C.int64_t(this.Exp)
	object.FriendshipLevel = C.int32_t(this.FriendshipLevel)
	object.Status = C.int16_t(this.Status)
	object.CoopId = C.int16_t(this.CoopId)
	object.Timestamp = C.int64_t(this.Timestamp)
	object.FightNum = C.int32_t(this.FightNum)
	return object
}

func (this insertOP) PlayerRole(row *PlayerRole) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerRole, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerRole
	for crow := this.db.tables.PlayerRole; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_role'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerRole
	this.db.tables.PlayerRole = newRow
	this.db.addTransLog(this.db.newPlayerRoleInsertLog(newRow, row))
}

func (this deleteOP) PlayerRole(row *PlayerRole) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerRole
	for crow := this.db.tables.PlayerRole; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_role'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerRole = this.db.tables.PlayerRole.next
	}
	this.db.addTransLog(this.db.newPlayerRoleDeleteLog(old, row))
}

func (this updateOP) PlayerRole(row *PlayerRole) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerRole
	for crow := this.db.tables.PlayerRole; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_role'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerRole = newRow
	}
	tmpRow := PlayerRoleRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerRoleUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerRole(key int64) *PlayerRole {
	for crow := this.db.tables.PlayerRole; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerRoleRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerRole(callback func(*PlayerRoleRow)) {
	row := &PlayerRoleRow{}
	for crow := this.db.tables.PlayerRole; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerRole() (rows []interface{}) {
	row := &PlayerRoleRow{}
	for crow := this.db.tables.PlayerRole; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerRole", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerRole", len(rows))
	return rows
}

/* ========== player_sealedbook ========== */

// 玩家天书记录
type PlayerSealedbook struct {
	Pid int64 // 用户ID
	SealedRecord []byte // 玩家天书记录
}

func (this *PlayerSealedbook) CObject() *C.PlayerSealedbook {
	object := C.New_PlayerSealedbook()
	object.Pid = C.int64_t(this.Pid)
	object.SealedRecord = C.CString(string(this.SealedRecord))
	object.SealedRecord_len = C.int(len(this.SealedRecord))
	return object
}

func (this insertOP) PlayerSealedbook(row *PlayerSealedbook) {
	if this.db.tables.PlayerSealedbook != nil {
		panic("duplicate insert 'player_sealedbook'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerSealedbook = newRow
	this.db.addTransLog(this.db.newPlayerSealedbookInsertLog(newRow, row))
}

func (this deleteOP) PlayerSealedbook(row *PlayerSealedbook) {
	if this.db.tables.PlayerSealedbook == nil {
		panic("delete not exists 'player_sealedbook'")
	}
	old := this.db.tables.PlayerSealedbook
	this.db.tables.PlayerSealedbook = nil
	this.db.addTransLog(this.db.newPlayerSealedbookDeleteLog(old, row))
}

func (this updateOP) PlayerSealedbook(row *PlayerSealedbook) {
	if this.db.tables.PlayerSealedbook == nil {
		panic("update not exists 'player_sealedbook'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerSealedbook
	this.db.tables.PlayerSealedbook = newRow
	tmpRow := PlayerSealedbookRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerSealedbookUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerSealedbook(key int64) *PlayerSealedbook {
	if this.db.tables.PlayerSealedbook == nil {
		return nil
	}
	tmpRow := PlayerSealedbookRow{row:this.db.tables.PlayerSealedbook}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_send_heart_record ========== */

// 玩家好友列表
type PlayerSendHeartRecord struct {
	Id int64 // ID
	Pid int64 // 玩家ID
	FriendPid int64 // 好友ID
	SendHeartTime int64 // 上次送爱心时间
}

func (this *PlayerSendHeartRecord) CObject() *C.PlayerSendHeartRecord {
	object := C.New_PlayerSendHeartRecord()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.FriendPid = C.int64_t(this.FriendPid)
	object.SendHeartTime = C.int64_t(this.SendHeartTime)
	return object
}

func (this insertOP) PlayerSendHeartRecord(row *PlayerSendHeartRecord) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerSendHeartRecord, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerSendHeartRecord
	for crow := this.db.tables.PlayerSendHeartRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_send_heart_record'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerSendHeartRecord
	this.db.tables.PlayerSendHeartRecord = newRow
	this.db.addTransLog(this.db.newPlayerSendHeartRecordInsertLog(newRow, row))
}

func (this deleteOP) PlayerSendHeartRecord(row *PlayerSendHeartRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerSendHeartRecord
	for crow := this.db.tables.PlayerSendHeartRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_send_heart_record'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerSendHeartRecord = this.db.tables.PlayerSendHeartRecord.next
	}
	this.db.addTransLog(this.db.newPlayerSendHeartRecordDeleteLog(old, row))
}

func (this updateOP) PlayerSendHeartRecord(row *PlayerSendHeartRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerSendHeartRecord
	for crow := this.db.tables.PlayerSendHeartRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_send_heart_record'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerSendHeartRecord = newRow
	}
	tmpRow := PlayerSendHeartRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerSendHeartRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerSendHeartRecord(key int64) *PlayerSendHeartRecord {
	for crow := this.db.tables.PlayerSendHeartRecord; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerSendHeartRecordRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerSendHeartRecord(callback func(*PlayerSendHeartRecordRow)) {
	row := &PlayerSendHeartRecordRow{}
	for crow := this.db.tables.PlayerSendHeartRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerSendHeartRecord() (rows []interface{}) {
	row := &PlayerSendHeartRecordRow{}
	for crow := this.db.tables.PlayerSendHeartRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerSendHeartRecord", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerSendHeartRecord", len(rows))
	return rows
}

/* ========== player_skill ========== */

// 玩家角色绝招表
type PlayerSkill struct {
	Id int64 // 主键ID
	Pid int64 // 玩家ID
	RoleId int8 // 角色ID
	SkillId int16 // 绝招ID
	SkillTrnlv int32 // 技能等级
}

func (this *PlayerSkill) CObject() *C.PlayerSkill {
	object := C.New_PlayerSkill()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.RoleId = C.int8_t(this.RoleId)
	object.SkillId = C.int16_t(this.SkillId)
	object.SkillTrnlv = C.int32_t(this.SkillTrnlv)
	return object
}

func (this insertOP) PlayerSkill(row *PlayerSkill) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerSkill, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerSkill
	for crow := this.db.tables.PlayerSkill; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_skill'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerSkill
	this.db.tables.PlayerSkill = newRow
	this.db.addTransLog(this.db.newPlayerSkillInsertLog(newRow, row))
}

func (this deleteOP) PlayerSkill(row *PlayerSkill) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerSkill
	for crow := this.db.tables.PlayerSkill; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_skill'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerSkill = this.db.tables.PlayerSkill.next
	}
	this.db.addTransLog(this.db.newPlayerSkillDeleteLog(old, row))
}

func (this updateOP) PlayerSkill(row *PlayerSkill) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerSkill
	for crow := this.db.tables.PlayerSkill; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_skill'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerSkill = newRow
	}
	tmpRow := PlayerSkillRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerSkillUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerSkill(key int64) *PlayerSkill {
	for crow := this.db.tables.PlayerSkill; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerSkillRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerSkill(callback func(*PlayerSkillRow)) {
	row := &PlayerSkillRow{}
	for crow := this.db.tables.PlayerSkill; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerSkill() (rows []interface{}) {
	row := &PlayerSkillRow{}
	for crow := this.db.tables.PlayerSkill; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerSkill", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerSkill", len(rows))
	return rows
}

/* ========== player_state ========== */

// 玩家状态
type PlayerState struct {
	Pid int64 // 玩家ID
	BanStartTime int64 // 被冻结的时间
	BanEndTime int64 // 被冻结时长, -1 无冻结；0 永久
}

func (this *PlayerState) CObject() *C.PlayerState {
	object := C.New_PlayerState()
	object.Pid = C.int64_t(this.Pid)
	object.BanStartTime = C.int64_t(this.BanStartTime)
	object.BanEndTime = C.int64_t(this.BanEndTime)
	return object
}

func (this insertOP) PlayerState(row *PlayerState) {
	if this.db.tables.PlayerState != nil {
		panic("duplicate insert 'player_state'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerState = newRow
	this.db.addTransLog(this.db.newPlayerStateInsertLog(newRow, row))
}

func (this deleteOP) PlayerState(row *PlayerState) {
	if this.db.tables.PlayerState == nil {
		panic("delete not exists 'player_state'")
	}
	old := this.db.tables.PlayerState
	this.db.tables.PlayerState = nil
	this.db.addTransLog(this.db.newPlayerStateDeleteLog(old, row))
}

func (this updateOP) PlayerState(row *PlayerState) {
	if this.db.tables.PlayerState == nil {
		panic("update not exists 'player_state'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerState
	this.db.tables.PlayerState = newRow
	tmpRow := PlayerStateRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerStateUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerState(key int64) *PlayerState {
	if this.db.tables.PlayerState == nil {
		return nil
	}
	tmpRow := PlayerStateRow{row:this.db.tables.PlayerState}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_sword_soul ========== */

// 玩家剑心数据
type PlayerSwordSoul struct {
	Id int64 // 玩家物品ID
	Pid int64 // 玩家ID
	Pos int8 // 是否已装备  1-已装备 0-未装备
	SwordSoulId int16 // 剑心ID
	Exp int32 // 当前经验
	Level int8 // 等级
}

func (this *PlayerSwordSoul) CObject() *C.PlayerSwordSoul {
	object := C.New_PlayerSwordSoul()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.Pos = C.int8_t(this.Pos)
	object.SwordSoulId = C.int16_t(this.SwordSoulId)
	object.Exp = C.int32_t(this.Exp)
	object.Level = C.int8_t(this.Level)
	return object
}

func (this insertOP) PlayerSwordSoul(row *PlayerSwordSoul) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerSwordSoul, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerSwordSoul
	for crow := this.db.tables.PlayerSwordSoul; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_sword_soul'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerSwordSoul
	this.db.tables.PlayerSwordSoul = newRow
	this.db.addTransLog(this.db.newPlayerSwordSoulInsertLog(newRow, row))
}

func (this deleteOP) PlayerSwordSoul(row *PlayerSwordSoul) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerSwordSoul
	for crow := this.db.tables.PlayerSwordSoul; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_sword_soul'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerSwordSoul = this.db.tables.PlayerSwordSoul.next
	}
	this.db.addTransLog(this.db.newPlayerSwordSoulDeleteLog(old, row))
}

func (this updateOP) PlayerSwordSoul(row *PlayerSwordSoul) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerSwordSoul
	for crow := this.db.tables.PlayerSwordSoul; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_sword_soul'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerSwordSoul = newRow
	}
	tmpRow := PlayerSwordSoulRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerSwordSoulUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerSwordSoul(key int64) *PlayerSwordSoul {
	for crow := this.db.tables.PlayerSwordSoul; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerSwordSoulRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerSwordSoul(callback func(*PlayerSwordSoulRow)) {
	row := &PlayerSwordSoulRow{}
	for crow := this.db.tables.PlayerSwordSoul; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerSwordSoul() (rows []interface{}) {
	row := &PlayerSwordSoulRow{}
	for crow := this.db.tables.PlayerSwordSoul; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerSwordSoul", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerSwordSoul", len(rows))
	return rows
}

/* ========== player_sword_soul_equipment ========== */

// 玩家剑心装备表
type PlayerSwordSoulEquipment struct {
	Id int64 // 主键ID
	Pid int64 // 玩家ID
	RoleId int8 // 角色ID
	IsEquippedXuanyuan int8 // 
	TypeBits int64 // 所有装备剑心类型的位标记
	Pos1 int64 // 装备位置1的剑心
	Pos2 int64 // 装备位置2的剑心
	Pos3 int64 // 装备位置3的剑心
	Pos4 int64 // 装备位置4的剑心
	Pos5 int64 // 装备位置5的剑心
	Pos6 int64 // 装备位置6的剑心
	Pos7 int64 // 装备位置7的剑心
	Pos8 int64 // 装备位置8的剑心
	Pos9 int64 // 装备位置9的剑心
}

func (this *PlayerSwordSoulEquipment) CObject() *C.PlayerSwordSoulEquipment {
	object := C.New_PlayerSwordSoulEquipment()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.RoleId = C.int8_t(this.RoleId)
	object.IsEquippedXuanyuan = C.int8_t(this.IsEquippedXuanyuan)
	object.TypeBits = C.int64_t(this.TypeBits)
	object.Pos1 = C.int64_t(this.Pos1)
	object.Pos2 = C.int64_t(this.Pos2)
	object.Pos3 = C.int64_t(this.Pos3)
	object.Pos4 = C.int64_t(this.Pos4)
	object.Pos5 = C.int64_t(this.Pos5)
	object.Pos6 = C.int64_t(this.Pos6)
	object.Pos7 = C.int64_t(this.Pos7)
	object.Pos8 = C.int64_t(this.Pos8)
	object.Pos9 = C.int64_t(this.Pos9)
	return object
}

func (this insertOP) PlayerSwordSoulEquipment(row *PlayerSwordSoulEquipment) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerSwordSoulEquipment, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerSwordSoulEquipment
	for crow := this.db.tables.PlayerSwordSoulEquipment; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_sword_soul_equipment'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerSwordSoulEquipment
	this.db.tables.PlayerSwordSoulEquipment = newRow
	this.db.addTransLog(this.db.newPlayerSwordSoulEquipmentInsertLog(newRow, row))
}

func (this deleteOP) PlayerSwordSoulEquipment(row *PlayerSwordSoulEquipment) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerSwordSoulEquipment
	for crow := this.db.tables.PlayerSwordSoulEquipment; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_sword_soul_equipment'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerSwordSoulEquipment = this.db.tables.PlayerSwordSoulEquipment.next
	}
	this.db.addTransLog(this.db.newPlayerSwordSoulEquipmentDeleteLog(old, row))
}

func (this updateOP) PlayerSwordSoulEquipment(row *PlayerSwordSoulEquipment) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerSwordSoulEquipment
	for crow := this.db.tables.PlayerSwordSoulEquipment; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_sword_soul_equipment'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerSwordSoulEquipment = newRow
	}
	tmpRow := PlayerSwordSoulEquipmentRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerSwordSoulEquipmentUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerSwordSoulEquipment(key int64) *PlayerSwordSoulEquipment {
	for crow := this.db.tables.PlayerSwordSoulEquipment; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerSwordSoulEquipmentRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerSwordSoulEquipment(callback func(*PlayerSwordSoulEquipmentRow)) {
	row := &PlayerSwordSoulEquipmentRow{}
	for crow := this.db.tables.PlayerSwordSoulEquipment; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerSwordSoulEquipment() (rows []interface{}) {
	row := &PlayerSwordSoulEquipmentRow{}
	for crow := this.db.tables.PlayerSwordSoulEquipment; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerSwordSoulEquipment", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerSwordSoulEquipment", len(rows))
	return rows
}

/* ========== player_sword_soul_state ========== */

// 玩家拔剑状态
type PlayerSwordSoulState struct {
	Pid int64 // 玩家ID
	BoxState int8 // 开箱子的状态(位操作)
	Num int16 // 当前可拔剑次数
	UpdateTime int64 // 更新时间
	IngotNum int64 // 
	SupersoulAdditionalPossible int8 // 
	IngotYelloOne int16 // 
	LastIngotDrawTime int64 // 
	SwordSoulNum int32 // 玩家剑心值
}

func (this *PlayerSwordSoulState) CObject() *C.PlayerSwordSoulState {
	object := C.New_PlayerSwordSoulState()
	object.Pid = C.int64_t(this.Pid)
	object.BoxState = C.int8_t(this.BoxState)
	object.Num = C.int16_t(this.Num)
	object.UpdateTime = C.int64_t(this.UpdateTime)
	object.IngotNum = C.int64_t(this.IngotNum)
	object.SupersoulAdditionalPossible = C.int8_t(this.SupersoulAdditionalPossible)
	object.IngotYelloOne = C.int16_t(this.IngotYelloOne)
	object.LastIngotDrawTime = C.int64_t(this.LastIngotDrawTime)
	object.SwordSoulNum = C.int32_t(this.SwordSoulNum)
	return object
}

func (this insertOP) PlayerSwordSoulState(row *PlayerSwordSoulState) {
	if this.db.tables.PlayerSwordSoulState != nil {
		panic("duplicate insert 'player_sword_soul_state'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerSwordSoulState = newRow
	this.db.addTransLog(this.db.newPlayerSwordSoulStateInsertLog(newRow, row))
}

func (this deleteOP) PlayerSwordSoulState(row *PlayerSwordSoulState) {
	if this.db.tables.PlayerSwordSoulState == nil {
		panic("delete not exists 'player_sword_soul_state'")
	}
	old := this.db.tables.PlayerSwordSoulState
	this.db.tables.PlayerSwordSoulState = nil
	this.db.addTransLog(this.db.newPlayerSwordSoulStateDeleteLog(old, row))
}

func (this updateOP) PlayerSwordSoulState(row *PlayerSwordSoulState) {
	if this.db.tables.PlayerSwordSoulState == nil {
		panic("update not exists 'player_sword_soul_state'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerSwordSoulState
	this.db.tables.PlayerSwordSoulState = newRow
	tmpRow := PlayerSwordSoulStateRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerSwordSoulStateUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerSwordSoulState(key int64) *PlayerSwordSoulState {
	if this.db.tables.PlayerSwordSoulState == nil {
		return nil
	}
	tmpRow := PlayerSwordSoulStateRow{row:this.db.tables.PlayerSwordSoulState}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_taoyuan_bless ========== */

// 玩家桃源祈福
type PlayerTaoyuanBless struct {
	Pid int64 // 玩家ID
	DailyBlessTimes int32 // 每日祈福次数
	LastBlessTime int64 // 最后一次祈福时间
	DailyBeBlessTimes int32 // 每日被祈福次数
	LastBeBlessTime int64 // 最后一次被祈福时间
	BlessPid1 int64 // 祈福玩家1
	BlessPid2 int64 // 祈福玩家2
	BlessPid3 int64 // 祈福玩家3
	BlessPid4 int64 // 祈福玩家4
	BlessPid5 int64 // 祈福玩家5
}

func (this *PlayerTaoyuanBless) CObject() *C.PlayerTaoyuanBless {
	object := C.New_PlayerTaoyuanBless()
	object.Pid = C.int64_t(this.Pid)
	object.DailyBlessTimes = C.int32_t(this.DailyBlessTimes)
	object.LastBlessTime = C.int64_t(this.LastBlessTime)
	object.DailyBeBlessTimes = C.int32_t(this.DailyBeBlessTimes)
	object.LastBeBlessTime = C.int64_t(this.LastBeBlessTime)
	object.BlessPid1 = C.int64_t(this.BlessPid1)
	object.BlessPid2 = C.int64_t(this.BlessPid2)
	object.BlessPid3 = C.int64_t(this.BlessPid3)
	object.BlessPid4 = C.int64_t(this.BlessPid4)
	object.BlessPid5 = C.int64_t(this.BlessPid5)
	return object
}

func (this insertOP) PlayerTaoyuanBless(row *PlayerTaoyuanBless) {
	if this.db.tables.PlayerTaoyuanBless != nil {
		panic("duplicate insert 'player_taoyuan_bless'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerTaoyuanBless = newRow
	this.db.addTransLog(this.db.newPlayerTaoyuanBlessInsertLog(newRow, row))
}

func (this deleteOP) PlayerTaoyuanBless(row *PlayerTaoyuanBless) {
	if this.db.tables.PlayerTaoyuanBless == nil {
		panic("delete not exists 'player_taoyuan_bless'")
	}
	old := this.db.tables.PlayerTaoyuanBless
	this.db.tables.PlayerTaoyuanBless = nil
	this.db.addTransLog(this.db.newPlayerTaoyuanBlessDeleteLog(old, row))
}

func (this updateOP) PlayerTaoyuanBless(row *PlayerTaoyuanBless) {
	if this.db.tables.PlayerTaoyuanBless == nil {
		panic("update not exists 'player_taoyuan_bless'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerTaoyuanBless
	this.db.tables.PlayerTaoyuanBless = newRow
	tmpRow := PlayerTaoyuanBlessRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerTaoyuanBlessUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerTaoyuanBless(key int64) *PlayerTaoyuanBless {
	if this.db.tables.PlayerTaoyuanBless == nil {
		return nil
	}
	tmpRow := PlayerTaoyuanBlessRow{row:this.db.tables.PlayerTaoyuanBless}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_taoyuan_fileds ========== */

// 玩家桃源田地作物
type PlayerTaoyuanFileds struct {
	Id int64 // 
	Pid int64 // 玩家ID
	FiledId int16 // 对应田地编号
	FiledStatus int16 // 对应田地状态0-普通,1-黑土地
	PlantId int16 // 种植植物ID
	GrowTime int64 // 种植时间
}

func (this *PlayerTaoyuanFileds) CObject() *C.PlayerTaoyuanFileds {
	object := C.New_PlayerTaoyuanFileds()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.FiledId = C.int16_t(this.FiledId)
	object.FiledStatus = C.int16_t(this.FiledStatus)
	object.PlantId = C.int16_t(this.PlantId)
	object.GrowTime = C.int64_t(this.GrowTime)
	return object
}

func (this insertOP) PlayerTaoyuanFileds(row *PlayerTaoyuanFileds) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerTaoyuanFileds, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerTaoyuanFileds
	for crow := this.db.tables.PlayerTaoyuanFileds; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_taoyuan_fileds'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerTaoyuanFileds
	this.db.tables.PlayerTaoyuanFileds = newRow
	this.db.addTransLog(this.db.newPlayerTaoyuanFiledsInsertLog(newRow, row))
}

func (this deleteOP) PlayerTaoyuanFileds(row *PlayerTaoyuanFileds) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerTaoyuanFileds
	for crow := this.db.tables.PlayerTaoyuanFileds; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_taoyuan_fileds'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerTaoyuanFileds = this.db.tables.PlayerTaoyuanFileds.next
	}
	this.db.addTransLog(this.db.newPlayerTaoyuanFiledsDeleteLog(old, row))
}

func (this updateOP) PlayerTaoyuanFileds(row *PlayerTaoyuanFileds) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerTaoyuanFileds
	for crow := this.db.tables.PlayerTaoyuanFileds; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_taoyuan_fileds'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerTaoyuanFileds = newRow
	}
	tmpRow := PlayerTaoyuanFiledsRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerTaoyuanFiledsUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerTaoyuanFileds(key int64) *PlayerTaoyuanFileds {
	for crow := this.db.tables.PlayerTaoyuanFileds; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerTaoyuanFiledsRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerTaoyuanFileds(callback func(*PlayerTaoyuanFiledsRow)) {
	row := &PlayerTaoyuanFiledsRow{}
	for crow := this.db.tables.PlayerTaoyuanFileds; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerTaoyuanFileds() (rows []interface{}) {
	row := &PlayerTaoyuanFiledsRow{}
	for crow := this.db.tables.PlayerTaoyuanFileds; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerTaoyuanFileds", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerTaoyuanFileds", len(rows))
	return rows
}

/* ========== player_taoyuan_heart ========== */

// 玩家桃源之心
type PlayerTaoyuanHeart struct {
	Pid int64 // 玩家ID
	Level int16 // 桃源之心等级
	Exp int64 // 桃源之心经验
}

func (this *PlayerTaoyuanHeart) CObject() *C.PlayerTaoyuanHeart {
	object := C.New_PlayerTaoyuanHeart()
	object.Pid = C.int64_t(this.Pid)
	object.Level = C.int16_t(this.Level)
	object.Exp = C.int64_t(this.Exp)
	return object
}

func (this insertOP) PlayerTaoyuanHeart(row *PlayerTaoyuanHeart) {
	if this.db.tables.PlayerTaoyuanHeart != nil {
		panic("duplicate insert 'player_taoyuan_heart'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerTaoyuanHeart = newRow
	this.db.addTransLog(this.db.newPlayerTaoyuanHeartInsertLog(newRow, row))
}

func (this deleteOP) PlayerTaoyuanHeart(row *PlayerTaoyuanHeart) {
	if this.db.tables.PlayerTaoyuanHeart == nil {
		panic("delete not exists 'player_taoyuan_heart'")
	}
	old := this.db.tables.PlayerTaoyuanHeart
	this.db.tables.PlayerTaoyuanHeart = nil
	this.db.addTransLog(this.db.newPlayerTaoyuanHeartDeleteLog(old, row))
}

func (this updateOP) PlayerTaoyuanHeart(row *PlayerTaoyuanHeart) {
	if this.db.tables.PlayerTaoyuanHeart == nil {
		panic("update not exists 'player_taoyuan_heart'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerTaoyuanHeart
	this.db.tables.PlayerTaoyuanHeart = newRow
	tmpRow := PlayerTaoyuanHeartRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerTaoyuanHeartUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerTaoyuanHeart(key int64) *PlayerTaoyuanHeart {
	if this.db.tables.PlayerTaoyuanHeart == nil {
		return nil
	}
	tmpRow := PlayerTaoyuanHeartRow{row:this.db.tables.PlayerTaoyuanHeart}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_taoyuan_item ========== */

// 玩家桃源物品
type PlayerTaoyuanItem struct {
	Id int64 // 主键ID
	Pid int64 // 玩家ID
	ItemId int16 // 物品ID
	Num int16 // 数量
}

func (this *PlayerTaoyuanItem) CObject() *C.PlayerTaoyuanItem {
	object := C.New_PlayerTaoyuanItem()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.ItemId = C.int16_t(this.ItemId)
	object.Num = C.int16_t(this.Num)
	return object
}

func (this insertOP) PlayerTaoyuanItem(row *PlayerTaoyuanItem) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerTaoyuanItem, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerTaoyuanItem
	for crow := this.db.tables.PlayerTaoyuanItem; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_taoyuan_item'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerTaoyuanItem
	this.db.tables.PlayerTaoyuanItem = newRow
	this.db.addTransLog(this.db.newPlayerTaoyuanItemInsertLog(newRow, row))
}

func (this deleteOP) PlayerTaoyuanItem(row *PlayerTaoyuanItem) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerTaoyuanItem
	for crow := this.db.tables.PlayerTaoyuanItem; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_taoyuan_item'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerTaoyuanItem = this.db.tables.PlayerTaoyuanItem.next
	}
	this.db.addTransLog(this.db.newPlayerTaoyuanItemDeleteLog(old, row))
}

func (this updateOP) PlayerTaoyuanItem(row *PlayerTaoyuanItem) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerTaoyuanItem
	for crow := this.db.tables.PlayerTaoyuanItem; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_taoyuan_item'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerTaoyuanItem = newRow
	}
	tmpRow := PlayerTaoyuanItemRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerTaoyuanItemUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerTaoyuanItem(key int64) *PlayerTaoyuanItem {
	for crow := this.db.tables.PlayerTaoyuanItem; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerTaoyuanItemRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerTaoyuanItem(callback func(*PlayerTaoyuanItemRow)) {
	row := &PlayerTaoyuanItemRow{}
	for crow := this.db.tables.PlayerTaoyuanItem; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerTaoyuanItem() (rows []interface{}) {
	row := &PlayerTaoyuanItemRow{}
	for crow := this.db.tables.PlayerTaoyuanItem; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerTaoyuanItem", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerTaoyuanItem", len(rows))
	return rows
}

/* ========== player_taoyuan_message ========== */

// 祈福消息列表
type PlayerTaoyuanMessage struct {
	Id int64 // 主键ID
	Pid int64 // 玩家ID
	Nick string // 祈福玩家名
	Exp int32 // 纳福获得桃源之心经验
}

func (this *PlayerTaoyuanMessage) CObject() *C.PlayerTaoyuanMessage {
	object := C.New_PlayerTaoyuanMessage()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.Nick = C.CString(string(this.Nick))
	object.Nick_len = C.int(len(this.Nick))
	object.Exp = C.int32_t(this.Exp)
	return object
}

func (this insertOP) PlayerTaoyuanMessage(row *PlayerTaoyuanMessage) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerTaoyuanMessage, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerTaoyuanMessage
	for crow := this.db.tables.PlayerTaoyuanMessage; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_taoyuan_message'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerTaoyuanMessage
	this.db.tables.PlayerTaoyuanMessage = newRow
	this.db.addTransLog(this.db.newPlayerTaoyuanMessageInsertLog(newRow, row))
}

func (this deleteOP) PlayerTaoyuanMessage(row *PlayerTaoyuanMessage) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerTaoyuanMessage
	for crow := this.db.tables.PlayerTaoyuanMessage; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_taoyuan_message'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerTaoyuanMessage = this.db.tables.PlayerTaoyuanMessage.next
	}
	this.db.addTransLog(this.db.newPlayerTaoyuanMessageDeleteLog(old, row))
}

func (this updateOP) PlayerTaoyuanMessage(row *PlayerTaoyuanMessage) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerTaoyuanMessage
	for crow := this.db.tables.PlayerTaoyuanMessage; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_taoyuan_message'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerTaoyuanMessage = newRow
	}
	tmpRow := PlayerTaoyuanMessageRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerTaoyuanMessageUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerTaoyuanMessage(key int64) *PlayerTaoyuanMessage {
	for crow := this.db.tables.PlayerTaoyuanMessage; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerTaoyuanMessageRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerTaoyuanMessage(callback func(*PlayerTaoyuanMessageRow)) {
	row := &PlayerTaoyuanMessageRow{}
	for crow := this.db.tables.PlayerTaoyuanMessage; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerTaoyuanMessage() (rows []interface{}) {
	row := &PlayerTaoyuanMessageRow{}
	for crow := this.db.tables.PlayerTaoyuanMessage; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerTaoyuanMessage", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerTaoyuanMessage", len(rows))
	return rows
}

/* ========== player_taoyuan_product ========== */

// 玩家桃源食坊、酒坊
type PlayerTaoyuanProduct struct {
	Pid int64 // 玩家ID
	SkillId1 int16 // 桃源酒坊技能
	SkillId2 int16 // 桃源食坊技能
	SameTimeWineQueue int16 // 酿酒队列
	MakeWineTimes int64 // 酿酒次数
	SameTimeFoodQueue int16 // 烹饪队列
	MakeFoodTimes int64 // 烹饪次数
	FoodQueue1 int16 // 队列1美食id
	FoodQueue1StartTimestamp int64 // 队列1美食开始时间
	FoodQueue1EndTimestamp int64 // 队列1美食结束时间
	FoodQueue2 int16 // 队列2美食id
	FoodQueue2StartTimestamp int64 // 队列2美食开始时间
	FoodQueue2EndTimestamp int64 // 队列2美食结束时间
	FoodQueue3 int16 // 队列3美食id
	FoodQueue3StartTimestamp int64 // 队列3美食开始时间
	FoodQueue3EndTimestamp int64 // 队列3美食结束时间
	FoodQueue4 int16 // 队列4美食id
	FoodQueue4StartTimestamp int64 // 队列4美食开始时间
	FoodQueue4EndTimestamp int64 // 队列4美食结束时间
	WineQueue1 int16 // 队列1美酒id
	WineQueue1StartTimestamp int64 // 队列1美酒开始时间
	WineQueue1EndTimestamp int64 // 队列1美酒结束时间
	WineQueue2 int16 // 队列2美酒id
	WineQueue2StartTimestamp int64 // 队列2美酒开始时间
	WineQueue2EndTimestamp int64 // 队列2美酒结束时间
	WineQueue3 int16 // 队列3美酒id
	WineQueue3StartTimestamp int64 // 队列3美酒开始时间
	WineQueue3EndTimestamp int64 // 队列3美酒结束时间
	WineQueue4 int16 // 队列4美酒id
	WineQueue4StartTimestamp int64 // 队列4美酒开始时间
	WineQueue4EndTimestamp int64 // 队列4美酒结束时间
	IsOpenWine int8 // 酒坊是否开启
	IsOpenFood int8 // 食坊是否开启
}

func (this *PlayerTaoyuanProduct) CObject() *C.PlayerTaoyuanProduct {
	object := C.New_PlayerTaoyuanProduct()
	object.Pid = C.int64_t(this.Pid)
	object.SkillId1 = C.int16_t(this.SkillId1)
	object.SkillId2 = C.int16_t(this.SkillId2)
	object.SameTimeWineQueue = C.int16_t(this.SameTimeWineQueue)
	object.MakeWineTimes = C.int64_t(this.MakeWineTimes)
	object.SameTimeFoodQueue = C.int16_t(this.SameTimeFoodQueue)
	object.MakeFoodTimes = C.int64_t(this.MakeFoodTimes)
	object.FoodQueue1 = C.int16_t(this.FoodQueue1)
	object.FoodQueue1StartTimestamp = C.int64_t(this.FoodQueue1StartTimestamp)
	object.FoodQueue1EndTimestamp = C.int64_t(this.FoodQueue1EndTimestamp)
	object.FoodQueue2 = C.int16_t(this.FoodQueue2)
	object.FoodQueue2StartTimestamp = C.int64_t(this.FoodQueue2StartTimestamp)
	object.FoodQueue2EndTimestamp = C.int64_t(this.FoodQueue2EndTimestamp)
	object.FoodQueue3 = C.int16_t(this.FoodQueue3)
	object.FoodQueue3StartTimestamp = C.int64_t(this.FoodQueue3StartTimestamp)
	object.FoodQueue3EndTimestamp = C.int64_t(this.FoodQueue3EndTimestamp)
	object.FoodQueue4 = C.int16_t(this.FoodQueue4)
	object.FoodQueue4StartTimestamp = C.int64_t(this.FoodQueue4StartTimestamp)
	object.FoodQueue4EndTimestamp = C.int64_t(this.FoodQueue4EndTimestamp)
	object.WineQueue1 = C.int16_t(this.WineQueue1)
	object.WineQueue1StartTimestamp = C.int64_t(this.WineQueue1StartTimestamp)
	object.WineQueue1EndTimestamp = C.int64_t(this.WineQueue1EndTimestamp)
	object.WineQueue2 = C.int16_t(this.WineQueue2)
	object.WineQueue2StartTimestamp = C.int64_t(this.WineQueue2StartTimestamp)
	object.WineQueue2EndTimestamp = C.int64_t(this.WineQueue2EndTimestamp)
	object.WineQueue3 = C.int16_t(this.WineQueue3)
	object.WineQueue3StartTimestamp = C.int64_t(this.WineQueue3StartTimestamp)
	object.WineQueue3EndTimestamp = C.int64_t(this.WineQueue3EndTimestamp)
	object.WineQueue4 = C.int16_t(this.WineQueue4)
	object.WineQueue4StartTimestamp = C.int64_t(this.WineQueue4StartTimestamp)
	object.WineQueue4EndTimestamp = C.int64_t(this.WineQueue4EndTimestamp)
	object.IsOpenWine = C.int8_t(this.IsOpenWine)
	object.IsOpenFood = C.int8_t(this.IsOpenFood)
	return object
}

func (this insertOP) PlayerTaoyuanProduct(row *PlayerTaoyuanProduct) {
	if this.db.tables.PlayerTaoyuanProduct != nil {
		panic("duplicate insert 'player_taoyuan_product'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerTaoyuanProduct = newRow
	this.db.addTransLog(this.db.newPlayerTaoyuanProductInsertLog(newRow, row))
}

func (this deleteOP) PlayerTaoyuanProduct(row *PlayerTaoyuanProduct) {
	if this.db.tables.PlayerTaoyuanProduct == nil {
		panic("delete not exists 'player_taoyuan_product'")
	}
	old := this.db.tables.PlayerTaoyuanProduct
	this.db.tables.PlayerTaoyuanProduct = nil
	this.db.addTransLog(this.db.newPlayerTaoyuanProductDeleteLog(old, row))
}

func (this updateOP) PlayerTaoyuanProduct(row *PlayerTaoyuanProduct) {
	if this.db.tables.PlayerTaoyuanProduct == nil {
		panic("update not exists 'player_taoyuan_product'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerTaoyuanProduct
	this.db.tables.PlayerTaoyuanProduct = newRow
	tmpRow := PlayerTaoyuanProductRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerTaoyuanProductUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerTaoyuanProduct(key int64) *PlayerTaoyuanProduct {
	if this.db.tables.PlayerTaoyuanProduct == nil {
		return nil
	}
	tmpRow := PlayerTaoyuanProductRow{row:this.db.tables.PlayerTaoyuanProduct}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_taoyuan_purchase_record ========== */

// 玩家桃源物品购买记录 仅记录限购商品
type PlayerTaoyuanPurchaseRecord struct {
	Id int64 // 主键ID
	Pid int64 // 玩家ID
	ItemId int16 // 物品ID
	Num int16 // 已购买数量
	Timestamp int64 // 上次购买时间
}

func (this *PlayerTaoyuanPurchaseRecord) CObject() *C.PlayerTaoyuanPurchaseRecord {
	object := C.New_PlayerTaoyuanPurchaseRecord()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.ItemId = C.int16_t(this.ItemId)
	object.Num = C.int16_t(this.Num)
	object.Timestamp = C.int64_t(this.Timestamp)
	return object
}

func (this insertOP) PlayerTaoyuanPurchaseRecord(row *PlayerTaoyuanPurchaseRecord) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerTaoyuanPurchaseRecord, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerTaoyuanPurchaseRecord
	for crow := this.db.tables.PlayerTaoyuanPurchaseRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_taoyuan_purchase_record'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerTaoyuanPurchaseRecord
	this.db.tables.PlayerTaoyuanPurchaseRecord = newRow
	this.db.addTransLog(this.db.newPlayerTaoyuanPurchaseRecordInsertLog(newRow, row))
}

func (this deleteOP) PlayerTaoyuanPurchaseRecord(row *PlayerTaoyuanPurchaseRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerTaoyuanPurchaseRecord
	for crow := this.db.tables.PlayerTaoyuanPurchaseRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_taoyuan_purchase_record'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerTaoyuanPurchaseRecord = this.db.tables.PlayerTaoyuanPurchaseRecord.next
	}
	this.db.addTransLog(this.db.newPlayerTaoyuanPurchaseRecordDeleteLog(old, row))
}

func (this updateOP) PlayerTaoyuanPurchaseRecord(row *PlayerTaoyuanPurchaseRecord) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerTaoyuanPurchaseRecord
	for crow := this.db.tables.PlayerTaoyuanPurchaseRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_taoyuan_purchase_record'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerTaoyuanPurchaseRecord = newRow
	}
	tmpRow := PlayerTaoyuanPurchaseRecordRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerTaoyuanPurchaseRecordUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerTaoyuanPurchaseRecord(key int64) *PlayerTaoyuanPurchaseRecord {
	for crow := this.db.tables.PlayerTaoyuanPurchaseRecord; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerTaoyuanPurchaseRecordRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerTaoyuanPurchaseRecord(callback func(*PlayerTaoyuanPurchaseRecordRow)) {
	row := &PlayerTaoyuanPurchaseRecordRow{}
	for crow := this.db.tables.PlayerTaoyuanPurchaseRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerTaoyuanPurchaseRecord() (rows []interface{}) {
	row := &PlayerTaoyuanPurchaseRecordRow{}
	for crow := this.db.tables.PlayerTaoyuanPurchaseRecord; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerTaoyuanPurchaseRecord", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerTaoyuanPurchaseRecord", len(rows))
	return rows
}

/* ========== player_taoyuan_quest ========== */

// 玩家桃源愿望
type PlayerTaoyuanQuest struct {
	Pid int64 // 玩家ID
	LastQuestFlashTime int64 // 最后一次桃源愿望刷新时间
	Quest1ItemId int16 // 桃源愿望1物品ID
	Quest1ItemNum int16 // 桃源愿望1物品需求数量
	Quest1TotalExp int16 // 桃源愿望1获得总经验
	Quest1TotalCoins int64 // 桃源愿望1获得总铜钱
	Quest1FinishTime int64 // 桃源愿望1完成时间
	Quest2ItemId int16 // 桃源愿望2物品ID
	Quest2ItemNum int16 // 桃源愿望2物品需求数量
	Quest2TotalExp int16 // 桃源愿望2获得总经验
	Quest2TotalCoins int64 // 桃源愿望2获得总铜钱
	Quest2FinishTime int64 // 桃源愿望2完成时间
	Quest3ItemId int16 // 桃源愿望3物品ID
	Quest3ItemNum int16 // 桃源愿望3物品需求数量
	Quest3TotalExp int16 // 桃源愿望3获得总经验
	Quest3TotalCoins int64 // 桃源愿望3获得总铜钱
	Quest3FinishTime int64 // 桃源愿望3完成时间
	Quest4ItemId int16 // 桃源愿望4物品ID
	Quest4ItemNum int16 // 桃源愿望4物品需求数量
	Quest4TotalExp int16 // 桃源愿望4获得总经验
	Quest4TotalCoins int64 // 桃源愿望4获得总铜钱
	Quest4FinishTime int64 // 桃源愿望4完成时间
	Quest5ItemId int16 // 桃源愿望5物品ID
	Quest5ItemNum int16 // 桃源愿望5物品需求数量
	Quest5TotalExp int16 // 桃源愿望5获得总经验
	Quest5TotalCoins int64 // 桃源愿望5获得总铜钱
	Quest5FinishTime int64 // 桃源愿望5完成时间
}

func (this *PlayerTaoyuanQuest) CObject() *C.PlayerTaoyuanQuest {
	object := C.New_PlayerTaoyuanQuest()
	object.Pid = C.int64_t(this.Pid)
	object.LastQuestFlashTime = C.int64_t(this.LastQuestFlashTime)
	object.Quest1ItemId = C.int16_t(this.Quest1ItemId)
	object.Quest1ItemNum = C.int16_t(this.Quest1ItemNum)
	object.Quest1TotalExp = C.int16_t(this.Quest1TotalExp)
	object.Quest1TotalCoins = C.int64_t(this.Quest1TotalCoins)
	object.Quest1FinishTime = C.int64_t(this.Quest1FinishTime)
	object.Quest2ItemId = C.int16_t(this.Quest2ItemId)
	object.Quest2ItemNum = C.int16_t(this.Quest2ItemNum)
	object.Quest2TotalExp = C.int16_t(this.Quest2TotalExp)
	object.Quest2TotalCoins = C.int64_t(this.Quest2TotalCoins)
	object.Quest2FinishTime = C.int64_t(this.Quest2FinishTime)
	object.Quest3ItemId = C.int16_t(this.Quest3ItemId)
	object.Quest3ItemNum = C.int16_t(this.Quest3ItemNum)
	object.Quest3TotalExp = C.int16_t(this.Quest3TotalExp)
	object.Quest3TotalCoins = C.int64_t(this.Quest3TotalCoins)
	object.Quest3FinishTime = C.int64_t(this.Quest3FinishTime)
	object.Quest4ItemId = C.int16_t(this.Quest4ItemId)
	object.Quest4ItemNum = C.int16_t(this.Quest4ItemNum)
	object.Quest4TotalExp = C.int16_t(this.Quest4TotalExp)
	object.Quest4TotalCoins = C.int64_t(this.Quest4TotalCoins)
	object.Quest4FinishTime = C.int64_t(this.Quest4FinishTime)
	object.Quest5ItemId = C.int16_t(this.Quest5ItemId)
	object.Quest5ItemNum = C.int16_t(this.Quest5ItemNum)
	object.Quest5TotalExp = C.int16_t(this.Quest5TotalExp)
	object.Quest5TotalCoins = C.int64_t(this.Quest5TotalCoins)
	object.Quest5FinishTime = C.int64_t(this.Quest5FinishTime)
	return object
}

func (this insertOP) PlayerTaoyuanQuest(row *PlayerTaoyuanQuest) {
	if this.db.tables.PlayerTaoyuanQuest != nil {
		panic("duplicate insert 'player_taoyuan_quest'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerTaoyuanQuest = newRow
	this.db.addTransLog(this.db.newPlayerTaoyuanQuestInsertLog(newRow, row))
}

func (this deleteOP) PlayerTaoyuanQuest(row *PlayerTaoyuanQuest) {
	if this.db.tables.PlayerTaoyuanQuest == nil {
		panic("delete not exists 'player_taoyuan_quest'")
	}
	old := this.db.tables.PlayerTaoyuanQuest
	this.db.tables.PlayerTaoyuanQuest = nil
	this.db.addTransLog(this.db.newPlayerTaoyuanQuestDeleteLog(old, row))
}

func (this updateOP) PlayerTaoyuanQuest(row *PlayerTaoyuanQuest) {
	if this.db.tables.PlayerTaoyuanQuest == nil {
		panic("update not exists 'player_taoyuan_quest'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerTaoyuanQuest
	this.db.tables.PlayerTaoyuanQuest = newRow
	tmpRow := PlayerTaoyuanQuestRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerTaoyuanQuestUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerTaoyuanQuest(key int64) *PlayerTaoyuanQuest {
	if this.db.tables.PlayerTaoyuanQuest == nil {
		return nil
	}
	tmpRow := PlayerTaoyuanQuestRow{row:this.db.tables.PlayerTaoyuanQuest}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_tb_xxd_roleinfo ========== */

// 腾讯经分用户信息表
type PlayerTbXxdRoleinfo struct {
	Pid int64 // 玩家id
	Gameappid string // 平台分配的AppID
	Openid string // 玩家平台唯一标识
	Regtime int64 // 注册时间
	Level int16 // 玩家等级
	IFriends int16 // 玩家好友数
	Moneyios int64 // ios金钱存量
	Moneyandroid int64 // android金钱存量
	Diamondios int64 // ios钻石存量
	Diamondandroid int64 // android钻石存量
}

func (this *PlayerTbXxdRoleinfo) CObject() *C.PlayerTbXxdRoleinfo {
	object := C.New_PlayerTbXxdRoleinfo()
	object.Pid = C.int64_t(this.Pid)
	object.Gameappid = C.CString(string(this.Gameappid))
	object.Gameappid_len = C.int(len(this.Gameappid))
	object.Openid = C.CString(string(this.Openid))
	object.Openid_len = C.int(len(this.Openid))
	object.Regtime = C.int64_t(this.Regtime)
	object.Level = C.int16_t(this.Level)
	object.IFriends = C.int16_t(this.IFriends)
	object.Moneyios = C.int64_t(this.Moneyios)
	object.Moneyandroid = C.int64_t(this.Moneyandroid)
	object.Diamondios = C.int64_t(this.Diamondios)
	object.Diamondandroid = C.int64_t(this.Diamondandroid)
	return object
}

func (this insertOP) PlayerTbXxdRoleinfo(row *PlayerTbXxdRoleinfo) {
	if this.db.tables.PlayerTbXxdRoleinfo != nil {
		panic("duplicate insert 'player_tb_xxd_roleinfo'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerTbXxdRoleinfo = newRow
	this.db.addTransLog(this.db.newPlayerTbXxdRoleinfoInsertLog(newRow, row))
}

func (this deleteOP) PlayerTbXxdRoleinfo(row *PlayerTbXxdRoleinfo) {
	if this.db.tables.PlayerTbXxdRoleinfo == nil {
		panic("delete not exists 'player_tb_xxd_roleinfo'")
	}
	old := this.db.tables.PlayerTbXxdRoleinfo
	this.db.tables.PlayerTbXxdRoleinfo = nil
	this.db.addTransLog(this.db.newPlayerTbXxdRoleinfoDeleteLog(old, row))
}

func (this updateOP) PlayerTbXxdRoleinfo(row *PlayerTbXxdRoleinfo) {
	if this.db.tables.PlayerTbXxdRoleinfo == nil {
		panic("update not exists 'player_tb_xxd_roleinfo'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerTbXxdRoleinfo
	this.db.tables.PlayerTbXxdRoleinfo = newRow
	tmpRow := PlayerTbXxdRoleinfoRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerTbXxdRoleinfoUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerTbXxdRoleinfo(key int64) *PlayerTbXxdRoleinfo {
	if this.db.tables.PlayerTbXxdRoleinfo == nil {
		return nil
	}
	tmpRow := PlayerTbXxdRoleinfoRow{row:this.db.tables.PlayerTbXxdRoleinfo}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_team_info ========== */

// 玩家队伍相关信息
type PlayerTeamInfo struct {
	Pid int64 // 玩家ID
	Relationship int32 // 友情点数
	HealthLv int16 // 生命项等级
	AttackLv int16 // 攻击项等级
	DefenceLv int16 // 防御项等级
}

func (this *PlayerTeamInfo) CObject() *C.PlayerTeamInfo {
	object := C.New_PlayerTeamInfo()
	object.Pid = C.int64_t(this.Pid)
	object.Relationship = C.int32_t(this.Relationship)
	object.HealthLv = C.int16_t(this.HealthLv)
	object.AttackLv = C.int16_t(this.AttackLv)
	object.DefenceLv = C.int16_t(this.DefenceLv)
	return object
}

func (this insertOP) PlayerTeamInfo(row *PlayerTeamInfo) {
	if this.db.tables.PlayerTeamInfo != nil {
		panic("duplicate insert 'player_team_info'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerTeamInfo = newRow
	this.db.addTransLog(this.db.newPlayerTeamInfoInsertLog(newRow, row))
}

func (this deleteOP) PlayerTeamInfo(row *PlayerTeamInfo) {
	if this.db.tables.PlayerTeamInfo == nil {
		panic("delete not exists 'player_team_info'")
	}
	old := this.db.tables.PlayerTeamInfo
	this.db.tables.PlayerTeamInfo = nil
	this.db.addTransLog(this.db.newPlayerTeamInfoDeleteLog(old, row))
}

func (this updateOP) PlayerTeamInfo(row *PlayerTeamInfo) {
	if this.db.tables.PlayerTeamInfo == nil {
		panic("update not exists 'player_team_info'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerTeamInfo
	this.db.tables.PlayerTeamInfo = newRow
	tmpRow := PlayerTeamInfoRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerTeamInfoUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerTeamInfo(key int64) *PlayerTeamInfo {
	if this.db.tables.PlayerTeamInfo == nil {
		return nil
	}
	tmpRow := PlayerTeamInfoRow{row:this.db.tables.PlayerTeamInfo}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_totem ========== */

// 玩家阵印表
type PlayerTotem struct {
	Id int64 // 主键
	Pid int64 // 玩家ID
	TotemId int16 // 阵印ID
	Level int8 // 等级
	SkillId int16 // 技能
}

func (this *PlayerTotem) CObject() *C.PlayerTotem {
	object := C.New_PlayerTotem()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.TotemId = C.int16_t(this.TotemId)
	object.Level = C.int8_t(this.Level)
	object.SkillId = C.int16_t(this.SkillId)
	return object
}

func (this insertOP) PlayerTotem(row *PlayerTotem) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerTotem, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerTotem
	for crow := this.db.tables.PlayerTotem; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_totem'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerTotem
	this.db.tables.PlayerTotem = newRow
	this.db.addTransLog(this.db.newPlayerTotemInsertLog(newRow, row))
}

func (this deleteOP) PlayerTotem(row *PlayerTotem) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerTotem
	for crow := this.db.tables.PlayerTotem; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_totem'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerTotem = this.db.tables.PlayerTotem.next
	}
	this.db.addTransLog(this.db.newPlayerTotemDeleteLog(old, row))
}

func (this updateOP) PlayerTotem(row *PlayerTotem) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerTotem
	for crow := this.db.tables.PlayerTotem; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_totem'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerTotem = newRow
	}
	tmpRow := PlayerTotemRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerTotemUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerTotem(key int64) *PlayerTotem {
	for crow := this.db.tables.PlayerTotem; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerTotemRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerTotem(callback func(*PlayerTotemRow)) {
	row := &PlayerTotemRow{}
	for crow := this.db.tables.PlayerTotem; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerTotem() (rows []interface{}) {
	row := &PlayerTotemRow{}
	for crow := this.db.tables.PlayerTotem; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerTotem", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerTotem", len(rows))
	return rows
}

/* ========== player_totem_info ========== */

// 玩家阵印装备表
type PlayerTotemInfo struct {
	Pid int64 // 玩家ID
	IngotCallDailyNum int16 // 每日元宝召唤次数
	IngotCallTimestamp int64 // 元宝召唤时间戳
	RockRuneNum int32 // 石附文数量
	JadeRuneNum int32 // 玉附文数量
	Pos1 int64 // 装备位置1的阵印id
	Pos2 int64 // 装备位置2的阵印id
	Pos3 int64 // 装备位置3的阵印id
	Pos4 int64 // 装备位置4的阵印id
	Pos5 int64 // 装备位置4的阵印id
	IngotDrawTimes int16 // 玩家元宝抽取次数
}

func (this *PlayerTotemInfo) CObject() *C.PlayerTotemInfo {
	object := C.New_PlayerTotemInfo()
	object.Pid = C.int64_t(this.Pid)
	object.IngotCallDailyNum = C.int16_t(this.IngotCallDailyNum)
	object.IngotCallTimestamp = C.int64_t(this.IngotCallTimestamp)
	object.RockRuneNum = C.int32_t(this.RockRuneNum)
	object.JadeRuneNum = C.int32_t(this.JadeRuneNum)
	object.Pos1 = C.int64_t(this.Pos1)
	object.Pos2 = C.int64_t(this.Pos2)
	object.Pos3 = C.int64_t(this.Pos3)
	object.Pos4 = C.int64_t(this.Pos4)
	object.Pos5 = C.int64_t(this.Pos5)
	object.IngotDrawTimes = C.int16_t(this.IngotDrawTimes)
	return object
}

func (this insertOP) PlayerTotemInfo(row *PlayerTotemInfo) {
	if this.db.tables.PlayerTotemInfo != nil {
		panic("duplicate insert 'player_totem_info'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerTotemInfo = newRow
	this.db.addTransLog(this.db.newPlayerTotemInfoInsertLog(newRow, row))
}

func (this deleteOP) PlayerTotemInfo(row *PlayerTotemInfo) {
	if this.db.tables.PlayerTotemInfo == nil {
		panic("delete not exists 'player_totem_info'")
	}
	old := this.db.tables.PlayerTotemInfo
	this.db.tables.PlayerTotemInfo = nil
	this.db.addTransLog(this.db.newPlayerTotemInfoDeleteLog(old, row))
}

func (this updateOP) PlayerTotemInfo(row *PlayerTotemInfo) {
	if this.db.tables.PlayerTotemInfo == nil {
		panic("update not exists 'player_totem_info'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerTotemInfo
	this.db.tables.PlayerTotemInfo = newRow
	tmpRow := PlayerTotemInfoRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerTotemInfoUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerTotemInfo(key int64) *PlayerTotemInfo {
	if this.db.tables.PlayerTotemInfo == nil {
		return nil
	}
	tmpRow := PlayerTotemInfoRow{row:this.db.tables.PlayerTotemInfo}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_town ========== */

// 玩家城镇数据
type PlayerTown struct {
	Pid int64 // 玩家ID
	TownId int16 // 当前玩家所处的城镇ID
	Lock int32 // 当前拥有的城镇权值
	AtPosX int16 // 当前城镇的X轴位置
	AtPosY int16 // 当前城镇的y轴位置
}

func (this *PlayerTown) CObject() *C.PlayerTown {
	object := C.New_PlayerTown()
	object.Pid = C.int64_t(this.Pid)
	object.TownId = C.int16_t(this.TownId)
	object.Lock = C.int32_t(this.Lock)
	object.AtPosX = C.int16_t(this.AtPosX)
	object.AtPosY = C.int16_t(this.AtPosY)
	return object
}

func (this insertOP) PlayerTown(row *PlayerTown) {
	if this.db.tables.PlayerTown != nil {
		panic("duplicate insert 'player_town'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerTown = newRow
	this.db.addTransLog(this.db.newPlayerTownInsertLog(newRow, row))
}

func (this deleteOP) PlayerTown(row *PlayerTown) {
	if this.db.tables.PlayerTown == nil {
		panic("delete not exists 'player_town'")
	}
	old := this.db.tables.PlayerTown
	this.db.tables.PlayerTown = nil
	this.db.addTransLog(this.db.newPlayerTownDeleteLog(old, row))
}

func (this updateOP) PlayerTown(row *PlayerTown) {
	if this.db.tables.PlayerTown == nil {
		panic("update not exists 'player_town'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerTown
	this.db.tables.PlayerTown = newRow
	tmpRow := PlayerTownRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerTownUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerTown(key int64) *PlayerTown {
	if this.db.tables.PlayerTown == nil {
		return nil
	}
	tmpRow := PlayerTownRow{row:this.db.tables.PlayerTown}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

/* ========== player_trader_refresh_state ========== */

// 玩家随机商店手动刷新次数状态
type PlayerTraderRefreshState struct {
	Id int64 // ID
	Pid int64 // 玩家ID
	LastUpdateTime int64 // 最近一次*手动*刷新时间
	AutoUpdateTime int64 // 最近一次*自动*刷新时间
	TraderId int16 // 随机商人ID
	RefreshNum int16 // 已刷新次数
}

func (this *PlayerTraderRefreshState) CObject() *C.PlayerTraderRefreshState {
	object := C.New_PlayerTraderRefreshState()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.LastUpdateTime = C.int64_t(this.LastUpdateTime)
	object.AutoUpdateTime = C.int64_t(this.AutoUpdateTime)
	object.TraderId = C.int16_t(this.TraderId)
	object.RefreshNum = C.int16_t(this.RefreshNum)
	return object
}

func (this insertOP) PlayerTraderRefreshState(row *PlayerTraderRefreshState) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerTraderRefreshState, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerTraderRefreshState
	for crow := this.db.tables.PlayerTraderRefreshState; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_trader_refresh_state'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerTraderRefreshState
	this.db.tables.PlayerTraderRefreshState = newRow
	this.db.addTransLog(this.db.newPlayerTraderRefreshStateInsertLog(newRow, row))
}

func (this deleteOP) PlayerTraderRefreshState(row *PlayerTraderRefreshState) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerTraderRefreshState
	for crow := this.db.tables.PlayerTraderRefreshState; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_trader_refresh_state'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerTraderRefreshState = this.db.tables.PlayerTraderRefreshState.next
	}
	this.db.addTransLog(this.db.newPlayerTraderRefreshStateDeleteLog(old, row))
}

func (this updateOP) PlayerTraderRefreshState(row *PlayerTraderRefreshState) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerTraderRefreshState
	for crow := this.db.tables.PlayerTraderRefreshState; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_trader_refresh_state'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerTraderRefreshState = newRow
	}
	tmpRow := PlayerTraderRefreshStateRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerTraderRefreshStateUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerTraderRefreshState(key int64) *PlayerTraderRefreshState {
	for crow := this.db.tables.PlayerTraderRefreshState; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerTraderRefreshStateRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerTraderRefreshState(callback func(*PlayerTraderRefreshStateRow)) {
	row := &PlayerTraderRefreshStateRow{}
	for crow := this.db.tables.PlayerTraderRefreshState; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerTraderRefreshState() (rows []interface{}) {
	row := &PlayerTraderRefreshStateRow{}
	for crow := this.db.tables.PlayerTraderRefreshState; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerTraderRefreshState", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerTraderRefreshState", len(rows))
	return rows
}

/* ========== player_trader_store_state ========== */

// 玩家随机商店状态
type PlayerTraderStoreState struct {
	Id int64 // ID
	Pid int64 // 玩家ID
	TraderId int16 // 随机商人ID
	GridId int32 // 格子ID
	ItemId int16 // 物品ID
	Num int16 // 物品数量
	Cost int64 // 价格
	Stock int8 // 剩余可购买次数
	GoodsType int8 // 货物类型0-物品 1-爱心 2-剑心 3-魂侍
}

func (this *PlayerTraderStoreState) CObject() *C.PlayerTraderStoreState {
	object := C.New_PlayerTraderStoreState()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.TraderId = C.int16_t(this.TraderId)
	object.GridId = C.int32_t(this.GridId)
	object.ItemId = C.int16_t(this.ItemId)
	object.Num = C.int16_t(this.Num)
	object.Cost = C.int64_t(this.Cost)
	object.Stock = C.int8_t(this.Stock)
	object.GoodsType = C.int8_t(this.GoodsType)
	return object
}

func (this insertOP) PlayerTraderStoreState(row *PlayerTraderStoreState) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerTraderStoreState, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerTraderStoreState
	for crow := this.db.tables.PlayerTraderStoreState; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_trader_store_state'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerTraderStoreState
	this.db.tables.PlayerTraderStoreState = newRow
	this.db.addTransLog(this.db.newPlayerTraderStoreStateInsertLog(newRow, row))
}

func (this deleteOP) PlayerTraderStoreState(row *PlayerTraderStoreState) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerTraderStoreState
	for crow := this.db.tables.PlayerTraderStoreState; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_trader_store_state'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerTraderStoreState = this.db.tables.PlayerTraderStoreState.next
	}
	this.db.addTransLog(this.db.newPlayerTraderStoreStateDeleteLog(old, row))
}

func (this updateOP) PlayerTraderStoreState(row *PlayerTraderStoreState) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerTraderStoreState
	for crow := this.db.tables.PlayerTraderStoreState; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_trader_store_state'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerTraderStoreState = newRow
	}
	tmpRow := PlayerTraderStoreStateRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerTraderStoreStateUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerTraderStoreState(key int64) *PlayerTraderStoreState {
	for crow := this.db.tables.PlayerTraderStoreState; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerTraderStoreStateRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerTraderStoreState(callback func(*PlayerTraderStoreStateRow)) {
	row := &PlayerTraderStoreStateRow{}
	for crow := this.db.tables.PlayerTraderStoreState; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerTraderStoreState() (rows []interface{}) {
	row := &PlayerTraderStoreStateRow{}
	for crow := this.db.tables.PlayerTraderStoreState; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerTraderStoreState", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerTraderStoreState", len(rows))
	return rows
}

/* ========== player_use_skill ========== */

// 玩家角色当前使用的绝招表
type PlayerUseSkill struct {
	Id int64 // 主键ID
	Pid int64 // 玩家ID
	RoleId int8 // 角色ID
	SkillId1 int16 // 招式1
	SkillId4 int16 // 招式4
	SkillId2 int16 // 招式2
	SkillId3 int16 // 招式3
}

func (this *PlayerUseSkill) CObject() *C.PlayerUseSkill {
	object := C.New_PlayerUseSkill()
	object.Id = C.int64_t(this.Id)
	object.Pid = C.int64_t(this.Pid)
	object.RoleId = C.int8_t(this.RoleId)
	object.SkillId1 = C.int16_t(this.SkillId1)
	object.SkillId4 = C.int16_t(this.SkillId4)
	object.SkillId2 = C.int16_t(this.SkillId2)
	object.SkillId3 = C.int16_t(this.SkillId3)
	return object
}

func (this insertOP) PlayerUseSkill(row *PlayerUseSkill) {
	row.Id = atomic.AddInt64(&g_RowIds.PlayerUseSkill, 1)
	key := C.int64_t(row.Id)
	var old *C.PlayerUseSkill
	for crow := this.db.tables.PlayerUseSkill; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("duplicate insert 'player_use_skill'")
	}
	newRow := row.CObject()
	newRow.next = this.db.tables.PlayerUseSkill
	this.db.tables.PlayerUseSkill = newRow
	this.db.addTransLog(this.db.newPlayerUseSkillInsertLog(newRow, row))
}

func (this deleteOP) PlayerUseSkill(row *PlayerUseSkill) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerUseSkill
	for crow := this.db.tables.PlayerUseSkill; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("delete not exists 'player_use_skill'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		this.db.tables.PlayerUseSkill = this.db.tables.PlayerUseSkill.next
	}
	this.db.addTransLog(this.db.newPlayerUseSkillDeleteLog(old, row))
}

func (this updateOP) PlayerUseSkill(row *PlayerUseSkill) {
	key := C.int64_t(row.Id)
	var old, prev *C.PlayerUseSkill
	for crow := this.db.tables.PlayerUseSkill; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old == nil {
		panic("update not exists 'player_use_skill'")
	}
	newRow := row.CObject()
	newRow.next = old.next
	if prev != nil {
		prev.next = newRow
	} else {
		this.db.tables.PlayerUseSkill = newRow
	}
	tmpRow := PlayerUseSkillRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerUseSkillUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerUseSkill(key int64) *PlayerUseSkill {
	for crow := this.db.tables.PlayerUseSkill; crow != nil; crow = crow.next {
		if crow.Id == C.int64_t(key) {
			tmpRow := PlayerUseSkillRow{row:crow}
			goRow := tmpRow.GoObject()
			tmpRow.row = nil
			return goRow
		}
	}
	return nil
}

func (this selectOP) PlayerUseSkill(callback func(*PlayerUseSkillRow)) {
	row := &PlayerUseSkillRow{}
	for crow := this.db.tables.PlayerUseSkill; crow != nil; crow = crow.next {
		row.reset(crow)
		callback(row)
		if row.isBreak {
			break
		}
	}
	row.row = nil
}

func (this allOP) PlayerUseSkill() (rows []interface{}) {
	row := &PlayerUseSkillRow{}
	for crow := this.db.tables.PlayerUseSkill; crow != nil; crow = crow.next {
		row.reset(crow)
		rows = append(rows, row.GoObject())
		fmt.Println("PlayerUseSkill", row.GoObject())
	}
	row.row = nil
	fmt.Println("PlayerUseSkill", len(rows))
	return rows
}

/* ========== player_vip ========== */

// 玩家VIP卡信息
type PlayerVip struct {
	Pid int64 // 玩家ID
	Ingot int64 // 累计充值元宝数
	Level int16 // VIP等级
	CardId string // VIP卡编号
	PresentExp int64 // 赠送vip经验
}

func (this *PlayerVip) CObject() *C.PlayerVip {
	object := C.New_PlayerVip()
	object.Pid = C.int64_t(this.Pid)
	object.Ingot = C.int64_t(this.Ingot)
	object.Level = C.int16_t(this.Level)
	object.CardId = C.CString(string(this.CardId))
	object.CardId_len = C.int(len(this.CardId))
	object.PresentExp = C.int64_t(this.PresentExp)
	return object
}

func (this insertOP) PlayerVip(row *PlayerVip) {
	if this.db.tables.PlayerVip != nil {
		panic("duplicate insert 'player_vip'")
	}
	newRow := row.CObject()
	this.db.tables.PlayerVip = newRow
	this.db.addTransLog(this.db.newPlayerVipInsertLog(newRow, row))
}

func (this deleteOP) PlayerVip(row *PlayerVip) {
	if this.db.tables.PlayerVip == nil {
		panic("delete not exists 'player_vip'")
	}
	old := this.db.tables.PlayerVip
	this.db.tables.PlayerVip = nil
	this.db.addTransLog(this.db.newPlayerVipDeleteLog(old, row))
}

func (this updateOP) PlayerVip(row *PlayerVip) {
	if this.db.tables.PlayerVip == nil {
		panic("update not exists 'player_vip'")
	}
	newRow := row.CObject()
	old := this.db.tables.PlayerVip
	this.db.tables.PlayerVip = newRow
	tmpRow := PlayerVipRow{row:old}
	goOld := tmpRow.GoObject()
	tmpRow.row = nil
	this.db.addTransLog(this.db.newPlayerVipUpdateLog(newRow, old, row, goOld))
}

func (this lookupOP) PlayerVip(key int64) *PlayerVip {
	if this.db.tables.PlayerVip == nil {
		return nil
	}
	tmpRow := PlayerVipRow{row:this.db.tables.PlayerVip}
	goRow := tmpRow.GoObject()
	tmpRow.row = nil
	return goRow
}

