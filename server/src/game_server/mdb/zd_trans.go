package mdb

/*
#include "zd_cgo.h"
*/
import "C"
import "unsafe"

import "time"

/* ========== global_announcement ========== */

type GlobalAnnouncementInsertLog struct {
	db *Database 
	Row *C.GlobalAnnouncement
	GoRow *GlobalAnnouncement
}

func (db *Database) newGlobalAnnouncementInsertLog(row *C.GlobalAnnouncement, goRow *GlobalAnnouncement) *GlobalAnnouncementInsertLog {
	return &GlobalAnnouncementInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *GlobalAnnouncementInsertLog) Free() {
}

func (log *GlobalAnnouncementInsertLog) InvokeHook() {
	if g_Hooks.GlobalAnnouncement != nil {
		g_Hooks.GlobalAnnouncement.Insert(&GlobalAnnouncementRow{row: log.Row})
	}
}

func (log *GlobalAnnouncementInsertLog) SetEventId(time.Time) {
}

func (log *GlobalAnnouncementInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(0)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.ExpireTime)
	file.WriteInt32(log.GoRow.TplId)
	file.WriteString(log.GoRow.Parameters)
	file.WriteString(log.GoRow.Content)
	file.WriteInt64(log.GoRow.SendTime)
	file.WriteInt64(log.GoRow.SpacingTime)
}

func (log *GlobalAnnouncementInsertLog) CommitToMySql() error {
	stmt := g_SQL.GlobalAnnouncement.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ExpireTime)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.TplId)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Parameters), int(log.Row.Parameters_len))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Content), int(log.Row.Content_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SendTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SpacingTime)))
	return stmt.Execute()
}

func (log *GlobalAnnouncementInsertLog) CommitToTLog() {
}

func (log *GlobalAnnouncementInsertLog) CommitToXdLog() {
}

func (log *GlobalAnnouncementInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalAnnouncementDeleteLog struct {
	db *Database
	Old *C.GlobalAnnouncement
	GoOld *GlobalAnnouncement
}

func (db *Database) newGlobalAnnouncementDeleteLog(old *C.GlobalAnnouncement, goOld *GlobalAnnouncement) *GlobalAnnouncementDeleteLog {
	return &GlobalAnnouncementDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *GlobalAnnouncementDeleteLog) Free() {
	C.Free_GlobalAnnouncement(log.Old)
}

func (log *GlobalAnnouncementDeleteLog) InvokeHook() {
	if g_Hooks.GlobalAnnouncement != nil {
		g_Hooks.GlobalAnnouncement.Delete(&GlobalAnnouncementRow{row: log.Old})
	}
}

func (log *GlobalAnnouncementDeleteLog) SetEventId(time.Time) {
}

func (log *GlobalAnnouncementDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(0)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.ExpireTime)
	file.WriteInt32(log.GoOld.TplId)
	file.WriteString(log.GoOld.Parameters)
	file.WriteString(log.GoOld.Content)
	file.WriteInt64(log.GoOld.SendTime)
	file.WriteInt64(log.GoOld.SpacingTime)
}

func (log *GlobalAnnouncementDeleteLog) CommitToMySql() error {
	stmt := g_SQL.GlobalAnnouncement.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *GlobalAnnouncementDeleteLog) CommitToTLog() {
}

func (log *GlobalAnnouncementDeleteLog) CommitToXdLog() {
}

func (log *GlobalAnnouncementDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalAnnouncementUpdateLog struct {
	db *Database 
	Row *C.GlobalAnnouncement
	Old *C.GlobalAnnouncement
	GoNew *GlobalAnnouncement
	GoOld *GlobalAnnouncement
}

func (db *Database) newGlobalAnnouncementUpdateLog(row, old *C.GlobalAnnouncement, goNew, goOld *GlobalAnnouncement) *GlobalAnnouncementUpdateLog {
	return &GlobalAnnouncementUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *GlobalAnnouncementUpdateLog) Free() {
	C.Free_GlobalAnnouncement(log.Old)
}

func (log *GlobalAnnouncementUpdateLog) InvokeHook() {
	if g_Hooks.GlobalAnnouncement != nil {
		g_Hooks.GlobalAnnouncement.Update(&GlobalAnnouncementRow{row: log.Row}, &GlobalAnnouncementRow{row: log.Old})
	}
}

func (log *GlobalAnnouncementUpdateLog) SetEventId(time.Time) {
}

func (log *GlobalAnnouncementUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(0)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.ExpireTime)
	file.WriteInt32(log.GoNew.TplId)
	file.WriteString(log.GoNew.Parameters)
	file.WriteString(log.GoNew.Content)
	file.WriteInt64(log.GoNew.SendTime)
	file.WriteInt64(log.GoNew.SpacingTime)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.ExpireTime)
	file.WriteInt32(log.GoOld.TplId)
	file.WriteString(log.GoOld.Parameters)
	file.WriteString(log.GoOld.Content)
	file.WriteInt64(log.GoOld.SendTime)
	file.WriteInt64(log.GoOld.SpacingTime)
}

func (log *GlobalAnnouncementUpdateLog) CommitToMySql() error {
	stmt := g_SQL.GlobalAnnouncement.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ExpireTime)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.TplId)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Parameters), int(log.Row.Parameters_len))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Content), int(log.Row.Content_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SendTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SpacingTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *GlobalAnnouncementUpdateLog) CommitToTLog() {
}

func (log *GlobalAnnouncementUpdateLog) CommitToXdLog() {
}

func (log *GlobalAnnouncementUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *GlobalAnnouncementInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.GlobalAnnouncement
	for crow := C.g_GlobalTables.GlobalAnnouncement; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'GlobalAnnouncement'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalAnnouncement = C.g_GlobalTables.GlobalAnnouncement.next
	}
	C.Free_GlobalAnnouncement(log.Row)
}

func (log *GlobalAnnouncementDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.GlobalAnnouncement
	for crow := C.g_GlobalTables.GlobalAnnouncement; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'global_announcement'")
	}
	log.Old.next = C.g_GlobalTables.GlobalAnnouncement
	C.g_GlobalTables.GlobalAnnouncement = log.Old
}

func (log *GlobalAnnouncementUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.GlobalAnnouncement
	for crow := C.g_GlobalTables.GlobalAnnouncement; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'global_announcement'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		C.g_GlobalTables.GlobalAnnouncement = log.Old
	}
	C.Free_GlobalAnnouncement(log.Row)
}

/* ========== global_arena_rank ========== */

type GlobalArenaRankInsertLog struct {
	db *Database 
	Row *C.GlobalArenaRank
	GoRow *GlobalArenaRank
}

func (db *Database) newGlobalArenaRankInsertLog(row *C.GlobalArenaRank, goRow *GlobalArenaRank) *GlobalArenaRankInsertLog {
	return &GlobalArenaRankInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *GlobalArenaRankInsertLog) Free() {
}

func (log *GlobalArenaRankInsertLog) InvokeHook() {
	if g_Hooks.GlobalArenaRank != nil {
		g_Hooks.GlobalArenaRank.Insert(&GlobalArenaRankRow{row: log.Row})
	}
}

func (log *GlobalArenaRankInsertLog) SetEventId(time.Time) {
}

func (log *GlobalArenaRankInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(1)
	file.WriteInt32(log.GoRow.Rank)
	file.WriteInt64(log.GoRow.Pid)
}

func (log *GlobalArenaRankInsertLog) CommitToMySql() error {
	stmt := g_SQL.GlobalArenaRank.Insert
	stmt.CleanBind()
	stmt.BindInt(unsafe.Pointer(&(log.Row.Rank)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	return stmt.Execute()
}

func (log *GlobalArenaRankInsertLog) CommitToTLog() {
}

func (log *GlobalArenaRankInsertLog) CommitToXdLog() {
}

func (log *GlobalArenaRankInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalArenaRankDeleteLog struct {
	db *Database
	Old *C.GlobalArenaRank
	GoOld *GlobalArenaRank
}

func (db *Database) newGlobalArenaRankDeleteLog(old *C.GlobalArenaRank, goOld *GlobalArenaRank) *GlobalArenaRankDeleteLog {
	return &GlobalArenaRankDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *GlobalArenaRankDeleteLog) Free() {
	C.Free_GlobalArenaRank(log.Old)
}

func (log *GlobalArenaRankDeleteLog) InvokeHook() {
	if g_Hooks.GlobalArenaRank != nil {
		g_Hooks.GlobalArenaRank.Delete(&GlobalArenaRankRow{row: log.Old})
	}
}

func (log *GlobalArenaRankDeleteLog) SetEventId(time.Time) {
}

func (log *GlobalArenaRankDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(1)
	file.WriteInt32(log.GoOld.Rank)
	file.WriteInt64(log.GoOld.Pid)
}

func (log *GlobalArenaRankDeleteLog) CommitToMySql() error {
	stmt := g_SQL.GlobalArenaRank.Delete
	stmt.CleanBind()
	stmt.BindInt(unsafe.Pointer(&(log.Old.Rank)))
	return stmt.Execute()
}

func (log *GlobalArenaRankDeleteLog) CommitToTLog() {
}

func (log *GlobalArenaRankDeleteLog) CommitToXdLog() {
}

func (log *GlobalArenaRankDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalArenaRankUpdateLog struct {
	db *Database 
	Row *C.GlobalArenaRank
	Old *C.GlobalArenaRank
	GoNew *GlobalArenaRank
	GoOld *GlobalArenaRank
}

func (db *Database) newGlobalArenaRankUpdateLog(row, old *C.GlobalArenaRank, goNew, goOld *GlobalArenaRank) *GlobalArenaRankUpdateLog {
	return &GlobalArenaRankUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *GlobalArenaRankUpdateLog) Free() {
	C.Free_GlobalArenaRank(log.Old)
}

func (log *GlobalArenaRankUpdateLog) InvokeHook() {
	if g_Hooks.GlobalArenaRank != nil {
		g_Hooks.GlobalArenaRank.Update(&GlobalArenaRankRow{row: log.Row}, &GlobalArenaRankRow{row: log.Old})
	}
}

func (log *GlobalArenaRankUpdateLog) SetEventId(time.Time) {
}

func (log *GlobalArenaRankUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(1)
	file.WriteInt32(log.GoNew.Rank)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoOld.Rank)
	file.WriteInt64(log.GoOld.Pid)
}

func (log *GlobalArenaRankUpdateLog) CommitToMySql() error {
	stmt := g_SQL.GlobalArenaRank.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Rank)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *GlobalArenaRankUpdateLog) CommitToTLog() {
}

func (log *GlobalArenaRankUpdateLog) CommitToXdLog() {
}

func (log *GlobalArenaRankUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *GlobalArenaRankInsertLog) Rollback() {
	key := log.Row.Rank
	var old, prev *C.GlobalArenaRank
	for crow := C.g_GlobalTables.GlobalArenaRank; crow != nil; prev, crow = crow, crow.next {
		if crow.Rank == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'GlobalArenaRank'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalArenaRank = C.g_GlobalTables.GlobalArenaRank.next
	}
	C.Free_GlobalArenaRank(log.Row)
}

func (log *GlobalArenaRankDeleteLog) Rollback() {
	key := log.Old.Rank
	var old *C.GlobalArenaRank
	for crow := C.g_GlobalTables.GlobalArenaRank; crow != nil; crow = crow.next {
		if crow.Rank == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'global_arena_rank'")
	}
	log.Old.next = C.g_GlobalTables.GlobalArenaRank
	C.g_GlobalTables.GlobalArenaRank = log.Old
}

func (log *GlobalArenaRankUpdateLog) Rollback() {
	key := log.Old.Rank
	var old, prev *C.GlobalArenaRank
	for crow := C.g_GlobalTables.GlobalArenaRank; crow != nil; prev, crow = crow, crow.next {
		if crow.Rank == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'global_arena_rank'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		C.g_GlobalTables.GlobalArenaRank = log.Old
	}
	C.Free_GlobalArenaRank(log.Row)
}

/* ========== global_clique ========== */

type GlobalCliqueInsertLog struct {
	db *Database 
	Row *C.GlobalClique
	GoRow *GlobalClique
}

func (db *Database) newGlobalCliqueInsertLog(row *C.GlobalClique, goRow *GlobalClique) *GlobalCliqueInsertLog {
	return &GlobalCliqueInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *GlobalCliqueInsertLog) Free() {
}

func (log *GlobalCliqueInsertLog) InvokeHook() {
	if g_Hooks.GlobalClique != nil {
		g_Hooks.GlobalClique.Insert(&GlobalCliqueRow{row: log.Row})
	}
}

func (log *GlobalCliqueInsertLog) SetEventId(time.Time) {
}

func (log *GlobalCliqueInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(2)
	file.WriteInt64(log.GoRow.Id)
	file.WriteString(log.GoRow.Name)
	file.WriteString(log.GoRow.Announce)
	file.WriteInt64(log.GoRow.TotalDonateCoins)
	file.WriteInt64(log.GoRow.OwnerPid)
	file.WriteInt64(log.GoRow.OwnerLoginTime)
	file.WriteInt64(log.GoRow.MangerPid1)
	file.WriteInt64(log.GoRow.MangerPid2)
	file.WriteInt64(log.GoRow.CenterBuildingCoins)
	file.WriteInt64(log.GoRow.TempleBuildingCoins)
	file.WriteInt64(log.GoRow.HealthBuildingCoins)
	file.WriteInt64(log.GoRow.AttackBuildingCoins)
	file.WriteInt64(log.GoRow.DefenseBuildingCoins)
	file.WriteInt64(log.GoRow.StoreBuildingCoins)
	file.WriteInt64(log.GoRow.BankBuildingCoins)
	file.WriteInt16(log.GoRow.CenterBuildingLevel)
	file.WriteInt16(log.GoRow.TempleBuildingLevel)
	file.WriteInt16(log.GoRow.HealthBuildingLevel)
	file.WriteInt16(log.GoRow.AttackBuildingLevel)
	file.WriteInt16(log.GoRow.DefenseBuildingLevel)
	file.WriteInt16(log.GoRow.BankBuildingLevel)
	file.WriteBytes(log.GoRow.MemberJson)
	file.WriteInt8(log.GoRow.AutoAudit)
	file.WriteInt16(log.GoRow.AutoAuditLevel)
	file.WriteInt64(log.GoRow.Contrib)
	file.WriteInt64(log.GoRow.ContribClearTime)
	file.WriteInt64(log.GoRow.RecruitTime)
	file.WriteInt64(log.GoRow.WorshipTime)
	file.WriteInt8(log.GoRow.WorshipCnt)
	file.WriteInt64(log.GoRow.StoreSendTime)
	file.WriteInt8(log.GoRow.StoreSendCnt)
}

func (log *GlobalCliqueInsertLog) CommitToMySql() error {
	stmt := g_SQL.GlobalClique.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Name), int(log.Row.Name_len))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Announce), int(log.Row.Announce_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.TotalDonateCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.OwnerPid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.OwnerLoginTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.MangerPid1)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.MangerPid2)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.CenterBuildingCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.TempleBuildingCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.HealthBuildingCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.AttackBuildingCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DefenseBuildingCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.StoreBuildingCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BankBuildingCoins)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CenterBuildingLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TempleBuildingLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.HealthBuildingLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AttackBuildingLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DefenseBuildingLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.BankBuildingLevel)))
	stmt.BindBinary(unsafe.Pointer(log.Row.MemberJson), int(log.Row.MemberJson_len))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.AutoAudit)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AutoAuditLevel)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Contrib)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ContribClearTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.RecruitTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WorshipTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.WorshipCnt)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.StoreSendTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.StoreSendCnt)))
	return stmt.Execute()
}

func (log *GlobalCliqueInsertLog) CommitToTLog() {
}

func (log *GlobalCliqueInsertLog) CommitToXdLog() {
}

func (log *GlobalCliqueInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalCliqueDeleteLog struct {
	db *Database
	Old *C.GlobalClique
	GoOld *GlobalClique
}

func (db *Database) newGlobalCliqueDeleteLog(old *C.GlobalClique, goOld *GlobalClique) *GlobalCliqueDeleteLog {
	return &GlobalCliqueDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *GlobalCliqueDeleteLog) Free() {
	C.Free_GlobalClique(log.Old)
}

func (log *GlobalCliqueDeleteLog) InvokeHook() {
	if g_Hooks.GlobalClique != nil {
		g_Hooks.GlobalClique.Delete(&GlobalCliqueRow{row: log.Old})
	}
}

func (log *GlobalCliqueDeleteLog) SetEventId(time.Time) {
}

func (log *GlobalCliqueDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(2)
	file.WriteInt64(log.GoOld.Id)
	file.WriteString(log.GoOld.Name)
	file.WriteString(log.GoOld.Announce)
	file.WriteInt64(log.GoOld.TotalDonateCoins)
	file.WriteInt64(log.GoOld.OwnerPid)
	file.WriteInt64(log.GoOld.OwnerLoginTime)
	file.WriteInt64(log.GoOld.MangerPid1)
	file.WriteInt64(log.GoOld.MangerPid2)
	file.WriteInt64(log.GoOld.CenterBuildingCoins)
	file.WriteInt64(log.GoOld.TempleBuildingCoins)
	file.WriteInt64(log.GoOld.HealthBuildingCoins)
	file.WriteInt64(log.GoOld.AttackBuildingCoins)
	file.WriteInt64(log.GoOld.DefenseBuildingCoins)
	file.WriteInt64(log.GoOld.StoreBuildingCoins)
	file.WriteInt64(log.GoOld.BankBuildingCoins)
	file.WriteInt16(log.GoOld.CenterBuildingLevel)
	file.WriteInt16(log.GoOld.TempleBuildingLevel)
	file.WriteInt16(log.GoOld.HealthBuildingLevel)
	file.WriteInt16(log.GoOld.AttackBuildingLevel)
	file.WriteInt16(log.GoOld.DefenseBuildingLevel)
	file.WriteInt16(log.GoOld.BankBuildingLevel)
	file.WriteBytes(log.GoOld.MemberJson)
	file.WriteInt8(log.GoOld.AutoAudit)
	file.WriteInt16(log.GoOld.AutoAuditLevel)
	file.WriteInt64(log.GoOld.Contrib)
	file.WriteInt64(log.GoOld.ContribClearTime)
	file.WriteInt64(log.GoOld.RecruitTime)
	file.WriteInt64(log.GoOld.WorshipTime)
	file.WriteInt8(log.GoOld.WorshipCnt)
	file.WriteInt64(log.GoOld.StoreSendTime)
	file.WriteInt8(log.GoOld.StoreSendCnt)
}

func (log *GlobalCliqueDeleteLog) CommitToMySql() error {
	stmt := g_SQL.GlobalClique.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *GlobalCliqueDeleteLog) CommitToTLog() {
}

func (log *GlobalCliqueDeleteLog) CommitToXdLog() {
}

func (log *GlobalCliqueDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalCliqueUpdateLog struct {
	db *Database 
	Row *C.GlobalClique
	Old *C.GlobalClique
	GoNew *GlobalClique
	GoOld *GlobalClique
}

func (db *Database) newGlobalCliqueUpdateLog(row, old *C.GlobalClique, goNew, goOld *GlobalClique) *GlobalCliqueUpdateLog {
	return &GlobalCliqueUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *GlobalCliqueUpdateLog) Free() {
	C.Free_GlobalClique(log.Old)
}

func (log *GlobalCliqueUpdateLog) InvokeHook() {
	if g_Hooks.GlobalClique != nil {
		g_Hooks.GlobalClique.Update(&GlobalCliqueRow{row: log.Row}, &GlobalCliqueRow{row: log.Old})
	}
}

func (log *GlobalCliqueUpdateLog) SetEventId(time.Time) {
}

func (log *GlobalCliqueUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(2)
	file.WriteInt64(log.GoNew.Id)
	file.WriteString(log.GoNew.Name)
	file.WriteString(log.GoNew.Announce)
	file.WriteInt64(log.GoNew.TotalDonateCoins)
	file.WriteInt64(log.GoNew.OwnerPid)
	file.WriteInt64(log.GoNew.OwnerLoginTime)
	file.WriteInt64(log.GoNew.MangerPid1)
	file.WriteInt64(log.GoNew.MangerPid2)
	file.WriteInt64(log.GoNew.CenterBuildingCoins)
	file.WriteInt64(log.GoNew.TempleBuildingCoins)
	file.WriteInt64(log.GoNew.HealthBuildingCoins)
	file.WriteInt64(log.GoNew.AttackBuildingCoins)
	file.WriteInt64(log.GoNew.DefenseBuildingCoins)
	file.WriteInt64(log.GoNew.StoreBuildingCoins)
	file.WriteInt64(log.GoNew.BankBuildingCoins)
	file.WriteInt16(log.GoNew.CenterBuildingLevel)
	file.WriteInt16(log.GoNew.TempleBuildingLevel)
	file.WriteInt16(log.GoNew.HealthBuildingLevel)
	file.WriteInt16(log.GoNew.AttackBuildingLevel)
	file.WriteInt16(log.GoNew.DefenseBuildingLevel)
	file.WriteInt16(log.GoNew.BankBuildingLevel)
	file.WriteBytes(log.GoNew.MemberJson)
	file.WriteInt8(log.GoNew.AutoAudit)
	file.WriteInt16(log.GoNew.AutoAuditLevel)
	file.WriteInt64(log.GoNew.Contrib)
	file.WriteInt64(log.GoNew.ContribClearTime)
	file.WriteInt64(log.GoNew.RecruitTime)
	file.WriteInt64(log.GoNew.WorshipTime)
	file.WriteInt8(log.GoNew.WorshipCnt)
	file.WriteInt64(log.GoNew.StoreSendTime)
	file.WriteInt8(log.GoNew.StoreSendCnt)
	file.WriteInt64(log.GoOld.Id)
	file.WriteString(log.GoOld.Name)
	file.WriteString(log.GoOld.Announce)
	file.WriteInt64(log.GoOld.TotalDonateCoins)
	file.WriteInt64(log.GoOld.OwnerPid)
	file.WriteInt64(log.GoOld.OwnerLoginTime)
	file.WriteInt64(log.GoOld.MangerPid1)
	file.WriteInt64(log.GoOld.MangerPid2)
	file.WriteInt64(log.GoOld.CenterBuildingCoins)
	file.WriteInt64(log.GoOld.TempleBuildingCoins)
	file.WriteInt64(log.GoOld.HealthBuildingCoins)
	file.WriteInt64(log.GoOld.AttackBuildingCoins)
	file.WriteInt64(log.GoOld.DefenseBuildingCoins)
	file.WriteInt64(log.GoOld.StoreBuildingCoins)
	file.WriteInt64(log.GoOld.BankBuildingCoins)
	file.WriteInt16(log.GoOld.CenterBuildingLevel)
	file.WriteInt16(log.GoOld.TempleBuildingLevel)
	file.WriteInt16(log.GoOld.HealthBuildingLevel)
	file.WriteInt16(log.GoOld.AttackBuildingLevel)
	file.WriteInt16(log.GoOld.DefenseBuildingLevel)
	file.WriteInt16(log.GoOld.BankBuildingLevel)
	file.WriteBytes(log.GoOld.MemberJson)
	file.WriteInt8(log.GoOld.AutoAudit)
	file.WriteInt16(log.GoOld.AutoAuditLevel)
	file.WriteInt64(log.GoOld.Contrib)
	file.WriteInt64(log.GoOld.ContribClearTime)
	file.WriteInt64(log.GoOld.RecruitTime)
	file.WriteInt64(log.GoOld.WorshipTime)
	file.WriteInt8(log.GoOld.WorshipCnt)
	file.WriteInt64(log.GoOld.StoreSendTime)
	file.WriteInt8(log.GoOld.StoreSendCnt)
}

func (log *GlobalCliqueUpdateLog) CommitToMySql() error {
	stmt := g_SQL.GlobalClique.Update
	stmt.BindVarchar(unsafe.Pointer(log.Row.Name), int(log.Row.Name_len))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Announce), int(log.Row.Announce_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.TotalDonateCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.OwnerPid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.OwnerLoginTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.MangerPid1)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.MangerPid2)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.CenterBuildingCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.TempleBuildingCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.HealthBuildingCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.AttackBuildingCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DefenseBuildingCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.StoreBuildingCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BankBuildingCoins)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CenterBuildingLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TempleBuildingLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.HealthBuildingLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AttackBuildingLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DefenseBuildingLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.BankBuildingLevel)))
	stmt.BindBinary(unsafe.Pointer(log.Row.MemberJson), int(log.Row.MemberJson_len))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.AutoAudit)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AutoAuditLevel)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Contrib)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ContribClearTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.RecruitTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WorshipTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.WorshipCnt)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.StoreSendTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.StoreSendCnt)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *GlobalCliqueUpdateLog) CommitToTLog() {
}

func (log *GlobalCliqueUpdateLog) CommitToXdLog() {
}

func (log *GlobalCliqueUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *GlobalCliqueInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.GlobalClique
	for crow := C.g_GlobalTables.GlobalClique; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'GlobalClique'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalClique = C.g_GlobalTables.GlobalClique.next
	}
	C.Free_GlobalClique(log.Row)
}

func (log *GlobalCliqueDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.GlobalClique
	for crow := C.g_GlobalTables.GlobalClique; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'global_clique'")
	}
	log.Old.next = C.g_GlobalTables.GlobalClique
	C.g_GlobalTables.GlobalClique = log.Old
}

func (log *GlobalCliqueUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.GlobalClique
	for crow := C.g_GlobalTables.GlobalClique; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'global_clique'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		C.g_GlobalTables.GlobalClique = log.Old
	}
	C.Free_GlobalClique(log.Row)
}

/* ========== global_clique_boat ========== */

type GlobalCliqueBoatInsertLog struct {
	db *Database 
	Row *C.GlobalCliqueBoat
	GoRow *GlobalCliqueBoat
}

func (db *Database) newGlobalCliqueBoatInsertLog(row *C.GlobalCliqueBoat, goRow *GlobalCliqueBoat) *GlobalCliqueBoatInsertLog {
	return &GlobalCliqueBoatInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *GlobalCliqueBoatInsertLog) Free() {
}

func (log *GlobalCliqueBoatInsertLog) InvokeHook() {
	if g_Hooks.GlobalCliqueBoat != nil {
		g_Hooks.GlobalCliqueBoat.Insert(&GlobalCliqueBoatRow{row: log.Row})
	}
}

func (log *GlobalCliqueBoatInsertLog) SetEventId(time.Time) {
}

func (log *GlobalCliqueBoatInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(3)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.CliqueId)
	file.WriteInt16(log.GoRow.BoatType)
	file.WriteInt64(log.GoRow.OwnerPid)
	file.WriteInt64(log.GoRow.EscortStartTimestamp)
	file.WriteInt16(log.GoRow.TotalEscortTime)
	file.WriteInt64(log.GoRow.RecoverPid)
	file.WriteInt64(log.GoRow.HijackerPid)
	file.WriteInt64(log.GoRow.HijackTimestamp)
	file.WriteInt8(log.GoRow.Status)
}

func (log *GlobalCliqueBoatInsertLog) CommitToMySql() error {
	stmt := g_SQL.GlobalCliqueBoat.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.CliqueId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.BoatType)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.OwnerPid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.EscortStartTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TotalEscortTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.RecoverPid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.HijackerPid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.HijackTimestamp)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Status)))
	return stmt.Execute()
}

func (log *GlobalCliqueBoatInsertLog) CommitToTLog() {
}

func (log *GlobalCliqueBoatInsertLog) CommitToXdLog() {
}

func (log *GlobalCliqueBoatInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalCliqueBoatDeleteLog struct {
	db *Database
	Old *C.GlobalCliqueBoat
	GoOld *GlobalCliqueBoat
}

func (db *Database) newGlobalCliqueBoatDeleteLog(old *C.GlobalCliqueBoat, goOld *GlobalCliqueBoat) *GlobalCliqueBoatDeleteLog {
	return &GlobalCliqueBoatDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *GlobalCliqueBoatDeleteLog) Free() {
	C.Free_GlobalCliqueBoat(log.Old)
}

func (log *GlobalCliqueBoatDeleteLog) InvokeHook() {
	if g_Hooks.GlobalCliqueBoat != nil {
		g_Hooks.GlobalCliqueBoat.Delete(&GlobalCliqueBoatRow{row: log.Old})
	}
}

func (log *GlobalCliqueBoatDeleteLog) SetEventId(time.Time) {
}

func (log *GlobalCliqueBoatDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(3)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.CliqueId)
	file.WriteInt16(log.GoOld.BoatType)
	file.WriteInt64(log.GoOld.OwnerPid)
	file.WriteInt64(log.GoOld.EscortStartTimestamp)
	file.WriteInt16(log.GoOld.TotalEscortTime)
	file.WriteInt64(log.GoOld.RecoverPid)
	file.WriteInt64(log.GoOld.HijackerPid)
	file.WriteInt64(log.GoOld.HijackTimestamp)
	file.WriteInt8(log.GoOld.Status)
}

func (log *GlobalCliqueBoatDeleteLog) CommitToMySql() error {
	stmt := g_SQL.GlobalCliqueBoat.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *GlobalCliqueBoatDeleteLog) CommitToTLog() {
}

func (log *GlobalCliqueBoatDeleteLog) CommitToXdLog() {
}

func (log *GlobalCliqueBoatDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalCliqueBoatUpdateLog struct {
	db *Database 
	Row *C.GlobalCliqueBoat
	Old *C.GlobalCliqueBoat
	GoNew *GlobalCliqueBoat
	GoOld *GlobalCliqueBoat
}

func (db *Database) newGlobalCliqueBoatUpdateLog(row, old *C.GlobalCliqueBoat, goNew, goOld *GlobalCliqueBoat) *GlobalCliqueBoatUpdateLog {
	return &GlobalCliqueBoatUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *GlobalCliqueBoatUpdateLog) Free() {
	C.Free_GlobalCliqueBoat(log.Old)
}

func (log *GlobalCliqueBoatUpdateLog) InvokeHook() {
	if g_Hooks.GlobalCliqueBoat != nil {
		g_Hooks.GlobalCliqueBoat.Update(&GlobalCliqueBoatRow{row: log.Row}, &GlobalCliqueBoatRow{row: log.Old})
	}
}

func (log *GlobalCliqueBoatUpdateLog) SetEventId(time.Time) {
}

func (log *GlobalCliqueBoatUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(3)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.CliqueId)
	file.WriteInt16(log.GoNew.BoatType)
	file.WriteInt64(log.GoNew.OwnerPid)
	file.WriteInt64(log.GoNew.EscortStartTimestamp)
	file.WriteInt16(log.GoNew.TotalEscortTime)
	file.WriteInt64(log.GoNew.RecoverPid)
	file.WriteInt64(log.GoNew.HijackerPid)
	file.WriteInt64(log.GoNew.HijackTimestamp)
	file.WriteInt8(log.GoNew.Status)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.CliqueId)
	file.WriteInt16(log.GoOld.BoatType)
	file.WriteInt64(log.GoOld.OwnerPid)
	file.WriteInt64(log.GoOld.EscortStartTimestamp)
	file.WriteInt16(log.GoOld.TotalEscortTime)
	file.WriteInt64(log.GoOld.RecoverPid)
	file.WriteInt64(log.GoOld.HijackerPid)
	file.WriteInt64(log.GoOld.HijackTimestamp)
	file.WriteInt8(log.GoOld.Status)
}

func (log *GlobalCliqueBoatUpdateLog) CommitToMySql() error {
	stmt := g_SQL.GlobalCliqueBoat.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.CliqueId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.BoatType)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.OwnerPid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.EscortStartTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TotalEscortTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.RecoverPid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.HijackerPid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.HijackTimestamp)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Status)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *GlobalCliqueBoatUpdateLog) CommitToTLog() {
}

func (log *GlobalCliqueBoatUpdateLog) CommitToXdLog() {
}

func (log *GlobalCliqueBoatUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *GlobalCliqueBoatInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.GlobalCliqueBoat
	for crow := C.g_GlobalTables.GlobalCliqueBoat; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'GlobalCliqueBoat'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalCliqueBoat = C.g_GlobalTables.GlobalCliqueBoat.next
	}
	C.Free_GlobalCliqueBoat(log.Row)
}

func (log *GlobalCliqueBoatDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.GlobalCliqueBoat
	for crow := C.g_GlobalTables.GlobalCliqueBoat; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'global_clique_boat'")
	}
	log.Old.next = C.g_GlobalTables.GlobalCliqueBoat
	C.g_GlobalTables.GlobalCliqueBoat = log.Old
}

func (log *GlobalCliqueBoatUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.GlobalCliqueBoat
	for crow := C.g_GlobalTables.GlobalCliqueBoat; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'global_clique_boat'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		C.g_GlobalTables.GlobalCliqueBoat = log.Old
	}
	C.Free_GlobalCliqueBoat(log.Row)
}

/* ========== global_despair_boss ========== */

type GlobalDespairBossInsertLog struct {
	db *Database 
	Row *C.GlobalDespairBoss
	GoRow *GlobalDespairBoss
}

func (db *Database) newGlobalDespairBossInsertLog(row *C.GlobalDespairBoss, goRow *GlobalDespairBoss) *GlobalDespairBossInsertLog {
	return &GlobalDespairBossInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *GlobalDespairBossInsertLog) Free() {
}

func (log *GlobalDespairBossInsertLog) InvokeHook() {
	if g_Hooks.GlobalDespairBoss != nil {
		g_Hooks.GlobalDespairBoss.Insert(&GlobalDespairBossRow{row: log.Row})
	}
}

func (log *GlobalDespairBossInsertLog) SetEventId(time.Time) {
}

func (log *GlobalDespairBossInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(4)
	file.WriteInt8(log.GoRow.CampType)
	file.WriteInt16(log.GoRow.Level)
	file.WriteInt64(log.GoRow.Timestamp)
	file.WriteInt64(log.GoRow.DeadTimestamp)
	file.WriteInt64(log.GoRow.Hp)
	file.WriteInt64(log.GoRow.MaxHp)
}

func (log *GlobalDespairBossInsertLog) CommitToMySql() error {
	stmt := g_SQL.GlobalDespairBoss.Insert
	stmt.CleanBind()
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.CampType)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DeadTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Hp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.MaxHp)))
	return stmt.Execute()
}

func (log *GlobalDespairBossInsertLog) CommitToTLog() {
}

func (log *GlobalDespairBossInsertLog) CommitToXdLog() {
}

func (log *GlobalDespairBossInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalDespairBossDeleteLog struct {
	db *Database
	Old *C.GlobalDespairBoss
	GoOld *GlobalDespairBoss
}

func (db *Database) newGlobalDespairBossDeleteLog(old *C.GlobalDespairBoss, goOld *GlobalDespairBoss) *GlobalDespairBossDeleteLog {
	return &GlobalDespairBossDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *GlobalDespairBossDeleteLog) Free() {
	C.Free_GlobalDespairBoss(log.Old)
}

func (log *GlobalDespairBossDeleteLog) InvokeHook() {
	if g_Hooks.GlobalDespairBoss != nil {
		g_Hooks.GlobalDespairBoss.Delete(&GlobalDespairBossRow{row: log.Old})
	}
}

func (log *GlobalDespairBossDeleteLog) SetEventId(time.Time) {
}

func (log *GlobalDespairBossDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(4)
	file.WriteInt8(log.GoOld.CampType)
	file.WriteInt16(log.GoOld.Level)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt64(log.GoOld.DeadTimestamp)
	file.WriteInt64(log.GoOld.Hp)
	file.WriteInt64(log.GoOld.MaxHp)
}

func (log *GlobalDespairBossDeleteLog) CommitToMySql() error {
	stmt := g_SQL.GlobalDespairBoss.Delete
	stmt.CleanBind()
	stmt.BindTinyInt(unsafe.Pointer(&(log.Old.CampType)))
	return stmt.Execute()
}

func (log *GlobalDespairBossDeleteLog) CommitToTLog() {
}

func (log *GlobalDespairBossDeleteLog) CommitToXdLog() {
}

func (log *GlobalDespairBossDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalDespairBossUpdateLog struct {
	db *Database 
	Row *C.GlobalDespairBoss
	Old *C.GlobalDespairBoss
	GoNew *GlobalDespairBoss
	GoOld *GlobalDespairBoss
}

func (db *Database) newGlobalDespairBossUpdateLog(row, old *C.GlobalDespairBoss, goNew, goOld *GlobalDespairBoss) *GlobalDespairBossUpdateLog {
	return &GlobalDespairBossUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *GlobalDespairBossUpdateLog) Free() {
	C.Free_GlobalDespairBoss(log.Old)
}

func (log *GlobalDespairBossUpdateLog) InvokeHook() {
	if g_Hooks.GlobalDespairBoss != nil {
		g_Hooks.GlobalDespairBoss.Update(&GlobalDespairBossRow{row: log.Row}, &GlobalDespairBossRow{row: log.Old})
	}
}

func (log *GlobalDespairBossUpdateLog) SetEventId(time.Time) {
}

func (log *GlobalDespairBossUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(4)
	file.WriteInt8(log.GoNew.CampType)
	file.WriteInt16(log.GoNew.Level)
	file.WriteInt64(log.GoNew.Timestamp)
	file.WriteInt64(log.GoNew.DeadTimestamp)
	file.WriteInt64(log.GoNew.Hp)
	file.WriteInt64(log.GoNew.MaxHp)
	file.WriteInt8(log.GoOld.CampType)
	file.WriteInt16(log.GoOld.Level)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt64(log.GoOld.DeadTimestamp)
	file.WriteInt64(log.GoOld.Hp)
	file.WriteInt64(log.GoOld.MaxHp)
}

func (log *GlobalDespairBossUpdateLog) CommitToMySql() error {
	stmt := g_SQL.GlobalDespairBoss.Update
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DeadTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Hp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.MaxHp)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.CampType)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *GlobalDespairBossUpdateLog) CommitToTLog() {
}

func (log *GlobalDespairBossUpdateLog) CommitToXdLog() {
}

func (log *GlobalDespairBossUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *GlobalDespairBossInsertLog) Rollback() {
	key := log.Row.CampType
	var old, prev *C.GlobalDespairBoss
	for crow := C.g_GlobalTables.GlobalDespairBoss; crow != nil; prev, crow = crow, crow.next {
		if crow.CampType == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'GlobalDespairBoss'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalDespairBoss = C.g_GlobalTables.GlobalDespairBoss.next
	}
	C.Free_GlobalDespairBoss(log.Row)
}

func (log *GlobalDespairBossDeleteLog) Rollback() {
	key := log.Old.CampType
	var old *C.GlobalDespairBoss
	for crow := C.g_GlobalTables.GlobalDespairBoss; crow != nil; crow = crow.next {
		if crow.CampType == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'global_despair_boss'")
	}
	log.Old.next = C.g_GlobalTables.GlobalDespairBoss
	C.g_GlobalTables.GlobalDespairBoss = log.Old
}

func (log *GlobalDespairBossUpdateLog) Rollback() {
	key := log.Old.CampType
	var old, prev *C.GlobalDespairBoss
	for crow := C.g_GlobalTables.GlobalDespairBoss; crow != nil; prev, crow = crow, crow.next {
		if crow.CampType == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'global_despair_boss'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		C.g_GlobalTables.GlobalDespairBoss = log.Old
	}
	C.Free_GlobalDespairBoss(log.Row)
}

/* ========== global_despair_boss_history ========== */

type GlobalDespairBossHistoryInsertLog struct {
	db *Database 
	Row *C.GlobalDespairBossHistory
	GoRow *GlobalDespairBossHistory
}

func (db *Database) newGlobalDespairBossHistoryInsertLog(row *C.GlobalDespairBossHistory, goRow *GlobalDespairBossHistory) *GlobalDespairBossHistoryInsertLog {
	return &GlobalDespairBossHistoryInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *GlobalDespairBossHistoryInsertLog) Free() {
}

func (log *GlobalDespairBossHistoryInsertLog) InvokeHook() {
	if g_Hooks.GlobalDespairBossHistory != nil {
		g_Hooks.GlobalDespairBossHistory.Insert(&GlobalDespairBossHistoryRow{row: log.Row})
	}
}

func (log *GlobalDespairBossHistoryInsertLog) SetEventId(time.Time) {
}

func (log *GlobalDespairBossHistoryInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(5)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Version)
	file.WriteInt8(log.GoRow.CampType)
	file.WriteInt16(log.GoRow.Level)
	file.WriteInt64(log.GoRow.Timestamp)
	file.WriteInt64(log.GoRow.StartTimestamp)
	file.WriteInt64(log.GoRow.DeadTimestamp)
}

func (log *GlobalDespairBossHistoryInsertLog) CommitToMySql() error {
	stmt := g_SQL.GlobalDespairBossHistory.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Version)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.CampType)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.StartTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DeadTimestamp)))
	return stmt.Execute()
}

func (log *GlobalDespairBossHistoryInsertLog) CommitToTLog() {
}

func (log *GlobalDespairBossHistoryInsertLog) CommitToXdLog() {
}

func (log *GlobalDespairBossHistoryInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalDespairBossHistoryDeleteLog struct {
	db *Database
	Old *C.GlobalDespairBossHistory
	GoOld *GlobalDespairBossHistory
}

func (db *Database) newGlobalDespairBossHistoryDeleteLog(old *C.GlobalDespairBossHistory, goOld *GlobalDespairBossHistory) *GlobalDespairBossHistoryDeleteLog {
	return &GlobalDespairBossHistoryDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *GlobalDespairBossHistoryDeleteLog) Free() {
	C.Free_GlobalDespairBossHistory(log.Old)
}

func (log *GlobalDespairBossHistoryDeleteLog) InvokeHook() {
	if g_Hooks.GlobalDespairBossHistory != nil {
		g_Hooks.GlobalDespairBossHistory.Delete(&GlobalDespairBossHistoryRow{row: log.Old})
	}
}

func (log *GlobalDespairBossHistoryDeleteLog) SetEventId(time.Time) {
}

func (log *GlobalDespairBossHistoryDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(5)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Version)
	file.WriteInt8(log.GoOld.CampType)
	file.WriteInt16(log.GoOld.Level)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt64(log.GoOld.StartTimestamp)
	file.WriteInt64(log.GoOld.DeadTimestamp)
}

func (log *GlobalDespairBossHistoryDeleteLog) CommitToMySql() error {
	stmt := g_SQL.GlobalDespairBossHistory.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *GlobalDespairBossHistoryDeleteLog) CommitToTLog() {
}

func (log *GlobalDespairBossHistoryDeleteLog) CommitToXdLog() {
}

func (log *GlobalDespairBossHistoryDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalDespairBossHistoryUpdateLog struct {
	db *Database 
	Row *C.GlobalDespairBossHistory
	Old *C.GlobalDespairBossHistory
	GoNew *GlobalDespairBossHistory
	GoOld *GlobalDespairBossHistory
}

func (db *Database) newGlobalDespairBossHistoryUpdateLog(row, old *C.GlobalDespairBossHistory, goNew, goOld *GlobalDespairBossHistory) *GlobalDespairBossHistoryUpdateLog {
	return &GlobalDespairBossHistoryUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *GlobalDespairBossHistoryUpdateLog) Free() {
	C.Free_GlobalDespairBossHistory(log.Old)
}

func (log *GlobalDespairBossHistoryUpdateLog) InvokeHook() {
	if g_Hooks.GlobalDespairBossHistory != nil {
		g_Hooks.GlobalDespairBossHistory.Update(&GlobalDespairBossHistoryRow{row: log.Row}, &GlobalDespairBossHistoryRow{row: log.Old})
	}
}

func (log *GlobalDespairBossHistoryUpdateLog) SetEventId(time.Time) {
}

func (log *GlobalDespairBossHistoryUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(5)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Version)
	file.WriteInt8(log.GoNew.CampType)
	file.WriteInt16(log.GoNew.Level)
	file.WriteInt64(log.GoNew.Timestamp)
	file.WriteInt64(log.GoNew.StartTimestamp)
	file.WriteInt64(log.GoNew.DeadTimestamp)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Version)
	file.WriteInt8(log.GoOld.CampType)
	file.WriteInt16(log.GoOld.Level)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt64(log.GoOld.StartTimestamp)
	file.WriteInt64(log.GoOld.DeadTimestamp)
}

func (log *GlobalDespairBossHistoryUpdateLog) CommitToMySql() error {
	stmt := g_SQL.GlobalDespairBossHistory.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Version)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.CampType)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.StartTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DeadTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *GlobalDespairBossHistoryUpdateLog) CommitToTLog() {
}

func (log *GlobalDespairBossHistoryUpdateLog) CommitToXdLog() {
}

func (log *GlobalDespairBossHistoryUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *GlobalDespairBossHistoryInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.GlobalDespairBossHistory
	for crow := C.g_GlobalTables.GlobalDespairBossHistory; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'GlobalDespairBossHistory'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalDespairBossHistory = C.g_GlobalTables.GlobalDespairBossHistory.next
	}
	C.Free_GlobalDespairBossHistory(log.Row)
}

func (log *GlobalDespairBossHistoryDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.GlobalDespairBossHistory
	for crow := C.g_GlobalTables.GlobalDespairBossHistory; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'global_despair_boss_history'")
	}
	log.Old.next = C.g_GlobalTables.GlobalDespairBossHistory
	C.g_GlobalTables.GlobalDespairBossHistory = log.Old
}

func (log *GlobalDespairBossHistoryUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.GlobalDespairBossHistory
	for crow := C.g_GlobalTables.GlobalDespairBossHistory; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'global_despair_boss_history'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		C.g_GlobalTables.GlobalDespairBossHistory = log.Old
	}
	C.Free_GlobalDespairBossHistory(log.Row)
}

/* ========== global_despair_land ========== */

type GlobalDespairLandInsertLog struct {
	db *Database 
	Row *C.GlobalDespairLand
	GoRow *GlobalDespairLand
}

func (db *Database) newGlobalDespairLandInsertLog(row *C.GlobalDespairLand, goRow *GlobalDespairLand) *GlobalDespairLandInsertLog {
	return &GlobalDespairLandInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *GlobalDespairLandInsertLog) Free() {
}

func (log *GlobalDespairLandInsertLog) InvokeHook() {
	if g_Hooks.GlobalDespairLand != nil {
		g_Hooks.GlobalDespairLand.Insert(&GlobalDespairLandRow{row: log.Row})
	}
}

func (log *GlobalDespairLandInsertLog) SetEventId(time.Time) {
}

func (log *GlobalDespairLandInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(6)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Version)
	file.WriteInt64(log.GoRow.Timestamp)
}

func (log *GlobalDespairLandInsertLog) CommitToMySql() error {
	stmt := g_SQL.GlobalDespairLand.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Version)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	return stmt.Execute()
}

func (log *GlobalDespairLandInsertLog) CommitToTLog() {
}

func (log *GlobalDespairLandInsertLog) CommitToXdLog() {
}

func (log *GlobalDespairLandInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalDespairLandDeleteLog struct {
	db *Database
	Old *C.GlobalDespairLand
	GoOld *GlobalDespairLand
}

func (db *Database) newGlobalDespairLandDeleteLog(old *C.GlobalDespairLand, goOld *GlobalDespairLand) *GlobalDespairLandDeleteLog {
	return &GlobalDespairLandDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *GlobalDespairLandDeleteLog) Free() {
	C.Free_GlobalDespairLand(log.Old)
}

func (log *GlobalDespairLandDeleteLog) InvokeHook() {
	if g_Hooks.GlobalDespairLand != nil {
		g_Hooks.GlobalDespairLand.Delete(&GlobalDespairLandRow{row: log.Old})
	}
}

func (log *GlobalDespairLandDeleteLog) SetEventId(time.Time) {
}

func (log *GlobalDespairLandDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(6)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Version)
	file.WriteInt64(log.GoOld.Timestamp)
}

func (log *GlobalDespairLandDeleteLog) CommitToMySql() error {
	stmt := g_SQL.GlobalDespairLand.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *GlobalDespairLandDeleteLog) CommitToTLog() {
}

func (log *GlobalDespairLandDeleteLog) CommitToXdLog() {
}

func (log *GlobalDespairLandDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalDespairLandUpdateLog struct {
	db *Database 
	Row *C.GlobalDespairLand
	Old *C.GlobalDespairLand
	GoNew *GlobalDespairLand
	GoOld *GlobalDespairLand
}

func (db *Database) newGlobalDespairLandUpdateLog(row, old *C.GlobalDespairLand, goNew, goOld *GlobalDespairLand) *GlobalDespairLandUpdateLog {
	return &GlobalDespairLandUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *GlobalDespairLandUpdateLog) Free() {
	C.Free_GlobalDespairLand(log.Old)
}

func (log *GlobalDespairLandUpdateLog) InvokeHook() {
	if g_Hooks.GlobalDespairLand != nil {
		g_Hooks.GlobalDespairLand.Update(&GlobalDespairLandRow{row: log.Row}, &GlobalDespairLandRow{row: log.Old})
	}
}

func (log *GlobalDespairLandUpdateLog) SetEventId(time.Time) {
}

func (log *GlobalDespairLandUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(6)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Version)
	file.WriteInt64(log.GoNew.Timestamp)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Version)
	file.WriteInt64(log.GoOld.Timestamp)
}

func (log *GlobalDespairLandUpdateLog) CommitToMySql() error {
	stmt := g_SQL.GlobalDespairLand.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Version)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *GlobalDespairLandUpdateLog) CommitToTLog() {
}

func (log *GlobalDespairLandUpdateLog) CommitToXdLog() {
}

func (log *GlobalDespairLandUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *GlobalDespairLandInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.GlobalDespairLand
	for crow := C.g_GlobalTables.GlobalDespairLand; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'GlobalDespairLand'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalDespairLand = C.g_GlobalTables.GlobalDespairLand.next
	}
	C.Free_GlobalDespairLand(log.Row)
}

func (log *GlobalDespairLandDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.GlobalDespairLand
	for crow := C.g_GlobalTables.GlobalDespairLand; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'global_despair_land'")
	}
	log.Old.next = C.g_GlobalTables.GlobalDespairLand
	C.g_GlobalTables.GlobalDespairLand = log.Old
}

func (log *GlobalDespairLandUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.GlobalDespairLand
	for crow := C.g_GlobalTables.GlobalDespairLand; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'global_despair_land'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		C.g_GlobalTables.GlobalDespairLand = log.Old
	}
	C.Free_GlobalDespairLand(log.Row)
}

/* ========== global_despair_land_battle_record ========== */

type GlobalDespairLandBattleRecordInsertLog struct {
	db *Database 
	Row *C.GlobalDespairLandBattleRecord
	GoRow *GlobalDespairLandBattleRecord
}

func (db *Database) newGlobalDespairLandBattleRecordInsertLog(row *C.GlobalDespairLandBattleRecord, goRow *GlobalDespairLandBattleRecord) *GlobalDespairLandBattleRecordInsertLog {
	return &GlobalDespairLandBattleRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *GlobalDespairLandBattleRecordInsertLog) Free() {
}

func (log *GlobalDespairLandBattleRecordInsertLog) InvokeHook() {
	if g_Hooks.GlobalDespairLandBattleRecord != nil {
		g_Hooks.GlobalDespairLandBattleRecord.Insert(&GlobalDespairLandBattleRecordRow{row: log.Row})
	}
}

func (log *GlobalDespairLandBattleRecordInsertLog) SetEventId(time.Time) {
}

func (log *GlobalDespairLandBattleRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(7)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.Fightnum)
	file.WriteInt64(log.GoRow.Timestamp)
	file.WriteInt16(log.GoRow.Tag)
	file.WriteInt16(log.GoRow.BattleVersion)
	file.WriteInt8(log.GoRow.CampType)
	file.WriteInt32(log.GoRow.LevelId)
	file.WriteBytes(log.GoRow.Record)
}

func (log *GlobalDespairLandBattleRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.GlobalDespairLandBattleRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Fightnum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Tag)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.BattleVersion)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.CampType)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.LevelId)))
	stmt.BindBinary(unsafe.Pointer(log.Row.Record), int(log.Row.Record_len))
	return stmt.Execute()
}

func (log *GlobalDespairLandBattleRecordInsertLog) CommitToTLog() {
}

func (log *GlobalDespairLandBattleRecordInsertLog) CommitToXdLog() {
}

func (log *GlobalDespairLandBattleRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalDespairLandBattleRecordDeleteLog struct {
	db *Database
	Old *C.GlobalDespairLandBattleRecord
	GoOld *GlobalDespairLandBattleRecord
}

func (db *Database) newGlobalDespairLandBattleRecordDeleteLog(old *C.GlobalDespairLandBattleRecord, goOld *GlobalDespairLandBattleRecord) *GlobalDespairLandBattleRecordDeleteLog {
	return &GlobalDespairLandBattleRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *GlobalDespairLandBattleRecordDeleteLog) Free() {
	C.Free_GlobalDespairLandBattleRecord(log.Old)
}

func (log *GlobalDespairLandBattleRecordDeleteLog) InvokeHook() {
	if g_Hooks.GlobalDespairLandBattleRecord != nil {
		g_Hooks.GlobalDespairLandBattleRecord.Delete(&GlobalDespairLandBattleRecordRow{row: log.Old})
	}
}

func (log *GlobalDespairLandBattleRecordDeleteLog) SetEventId(time.Time) {
}

func (log *GlobalDespairLandBattleRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(7)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Fightnum)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt16(log.GoOld.Tag)
	file.WriteInt16(log.GoOld.BattleVersion)
	file.WriteInt8(log.GoOld.CampType)
	file.WriteInt32(log.GoOld.LevelId)
	file.WriteBytes(log.GoOld.Record)
}

func (log *GlobalDespairLandBattleRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.GlobalDespairLandBattleRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *GlobalDespairLandBattleRecordDeleteLog) CommitToTLog() {
}

func (log *GlobalDespairLandBattleRecordDeleteLog) CommitToXdLog() {
}

func (log *GlobalDespairLandBattleRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalDespairLandBattleRecordUpdateLog struct {
	db *Database 
	Row *C.GlobalDespairLandBattleRecord
	Old *C.GlobalDespairLandBattleRecord
	GoNew *GlobalDespairLandBattleRecord
	GoOld *GlobalDespairLandBattleRecord
}

func (db *Database) newGlobalDespairLandBattleRecordUpdateLog(row, old *C.GlobalDespairLandBattleRecord, goNew, goOld *GlobalDespairLandBattleRecord) *GlobalDespairLandBattleRecordUpdateLog {
	return &GlobalDespairLandBattleRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *GlobalDespairLandBattleRecordUpdateLog) Free() {
	C.Free_GlobalDespairLandBattleRecord(log.Old)
}

func (log *GlobalDespairLandBattleRecordUpdateLog) InvokeHook() {
	if g_Hooks.GlobalDespairLandBattleRecord != nil {
		g_Hooks.GlobalDespairLandBattleRecord.Update(&GlobalDespairLandBattleRecordRow{row: log.Row}, &GlobalDespairLandBattleRecordRow{row: log.Old})
	}
}

func (log *GlobalDespairLandBattleRecordUpdateLog) SetEventId(time.Time) {
}

func (log *GlobalDespairLandBattleRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(7)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.Fightnum)
	file.WriteInt64(log.GoNew.Timestamp)
	file.WriteInt16(log.GoNew.Tag)
	file.WriteInt16(log.GoNew.BattleVersion)
	file.WriteInt8(log.GoNew.CampType)
	file.WriteInt32(log.GoNew.LevelId)
	file.WriteBytes(log.GoNew.Record)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Fightnum)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt16(log.GoOld.Tag)
	file.WriteInt16(log.GoOld.BattleVersion)
	file.WriteInt8(log.GoOld.CampType)
	file.WriteInt32(log.GoOld.LevelId)
	file.WriteBytes(log.GoOld.Record)
}

func (log *GlobalDespairLandBattleRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.GlobalDespairLandBattleRecord.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Fightnum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Tag)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.BattleVersion)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.CampType)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.LevelId)))
	stmt.BindBinary(unsafe.Pointer(log.Row.Record), int(log.Row.Record_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *GlobalDespairLandBattleRecordUpdateLog) CommitToTLog() {
}

func (log *GlobalDespairLandBattleRecordUpdateLog) CommitToXdLog() {
}

func (log *GlobalDespairLandBattleRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *GlobalDespairLandBattleRecordInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.GlobalDespairLandBattleRecord
	for crow := C.g_GlobalTables.GlobalDespairLandBattleRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'GlobalDespairLandBattleRecord'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalDespairLandBattleRecord = C.g_GlobalTables.GlobalDespairLandBattleRecord.next
	}
	C.Free_GlobalDespairLandBattleRecord(log.Row)
}

func (log *GlobalDespairLandBattleRecordDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.GlobalDespairLandBattleRecord
	for crow := C.g_GlobalTables.GlobalDespairLandBattleRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'global_despair_land_battle_record'")
	}
	log.Old.next = C.g_GlobalTables.GlobalDespairLandBattleRecord
	C.g_GlobalTables.GlobalDespairLandBattleRecord = log.Old
}

func (log *GlobalDespairLandBattleRecordUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.GlobalDespairLandBattleRecord
	for crow := C.g_GlobalTables.GlobalDespairLandBattleRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'global_despair_land_battle_record'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		C.g_GlobalTables.GlobalDespairLandBattleRecord = log.Old
	}
	C.Free_GlobalDespairLandBattleRecord(log.Row)
}

/* ========== global_despair_land_camp ========== */

type GlobalDespairLandCampInsertLog struct {
	db *Database 
	Row *C.GlobalDespairLandCamp
	GoRow *GlobalDespairLandCamp
}

func (db *Database) newGlobalDespairLandCampInsertLog(row *C.GlobalDespairLandCamp, goRow *GlobalDespairLandCamp) *GlobalDespairLandCampInsertLog {
	return &GlobalDespairLandCampInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *GlobalDespairLandCampInsertLog) Free() {
}

func (log *GlobalDespairLandCampInsertLog) InvokeHook() {
	if g_Hooks.GlobalDespairLandCamp != nil {
		g_Hooks.GlobalDespairLandCamp.Insert(&GlobalDespairLandCampRow{row: log.Row})
	}
}

func (log *GlobalDespairLandCampInsertLog) SetEventId(time.Time) {
}

func (log *GlobalDespairLandCampInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(8)
	file.WriteInt8(log.GoRow.CampType)
	file.WriteInt64(log.GoRow.Timestamp)
	file.WriteInt64(log.GoRow.DeadTimestamp)
	file.WriteInt16(log.GoRow.Level)
}

func (log *GlobalDespairLandCampInsertLog) CommitToMySql() error {
	stmt := g_SQL.GlobalDespairLandCamp.Insert
	stmt.CleanBind()
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.CampType)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DeadTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Level)))
	return stmt.Execute()
}

func (log *GlobalDespairLandCampInsertLog) CommitToTLog() {
}

func (log *GlobalDespairLandCampInsertLog) CommitToXdLog() {
}

func (log *GlobalDespairLandCampInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalDespairLandCampDeleteLog struct {
	db *Database
	Old *C.GlobalDespairLandCamp
	GoOld *GlobalDespairLandCamp
}

func (db *Database) newGlobalDespairLandCampDeleteLog(old *C.GlobalDespairLandCamp, goOld *GlobalDespairLandCamp) *GlobalDespairLandCampDeleteLog {
	return &GlobalDespairLandCampDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *GlobalDespairLandCampDeleteLog) Free() {
	C.Free_GlobalDespairLandCamp(log.Old)
}

func (log *GlobalDespairLandCampDeleteLog) InvokeHook() {
	if g_Hooks.GlobalDespairLandCamp != nil {
		g_Hooks.GlobalDespairLandCamp.Delete(&GlobalDespairLandCampRow{row: log.Old})
	}
}

func (log *GlobalDespairLandCampDeleteLog) SetEventId(time.Time) {
}

func (log *GlobalDespairLandCampDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(8)
	file.WriteInt8(log.GoOld.CampType)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt64(log.GoOld.DeadTimestamp)
	file.WriteInt16(log.GoOld.Level)
}

func (log *GlobalDespairLandCampDeleteLog) CommitToMySql() error {
	stmt := g_SQL.GlobalDespairLandCamp.Delete
	stmt.CleanBind()
	stmt.BindTinyInt(unsafe.Pointer(&(log.Old.CampType)))
	return stmt.Execute()
}

func (log *GlobalDespairLandCampDeleteLog) CommitToTLog() {
}

func (log *GlobalDespairLandCampDeleteLog) CommitToXdLog() {
}

func (log *GlobalDespairLandCampDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalDespairLandCampUpdateLog struct {
	db *Database 
	Row *C.GlobalDespairLandCamp
	Old *C.GlobalDespairLandCamp
	GoNew *GlobalDespairLandCamp
	GoOld *GlobalDespairLandCamp
}

func (db *Database) newGlobalDespairLandCampUpdateLog(row, old *C.GlobalDespairLandCamp, goNew, goOld *GlobalDespairLandCamp) *GlobalDespairLandCampUpdateLog {
	return &GlobalDespairLandCampUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *GlobalDespairLandCampUpdateLog) Free() {
	C.Free_GlobalDespairLandCamp(log.Old)
}

func (log *GlobalDespairLandCampUpdateLog) InvokeHook() {
	if g_Hooks.GlobalDespairLandCamp != nil {
		g_Hooks.GlobalDespairLandCamp.Update(&GlobalDespairLandCampRow{row: log.Row}, &GlobalDespairLandCampRow{row: log.Old})
	}
}

func (log *GlobalDespairLandCampUpdateLog) SetEventId(time.Time) {
}

func (log *GlobalDespairLandCampUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(8)
	file.WriteInt8(log.GoNew.CampType)
	file.WriteInt64(log.GoNew.Timestamp)
	file.WriteInt64(log.GoNew.DeadTimestamp)
	file.WriteInt16(log.GoNew.Level)
	file.WriteInt8(log.GoOld.CampType)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt64(log.GoOld.DeadTimestamp)
	file.WriteInt16(log.GoOld.Level)
}

func (log *GlobalDespairLandCampUpdateLog) CommitToMySql() error {
	stmt := g_SQL.GlobalDespairLandCamp.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DeadTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.CampType)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *GlobalDespairLandCampUpdateLog) CommitToTLog() {
}

func (log *GlobalDespairLandCampUpdateLog) CommitToXdLog() {
}

func (log *GlobalDespairLandCampUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *GlobalDespairLandCampInsertLog) Rollback() {
	key := log.Row.CampType
	var old, prev *C.GlobalDespairLandCamp
	for crow := C.g_GlobalTables.GlobalDespairLandCamp; crow != nil; prev, crow = crow, crow.next {
		if crow.CampType == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'GlobalDespairLandCamp'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalDespairLandCamp = C.g_GlobalTables.GlobalDespairLandCamp.next
	}
	C.Free_GlobalDespairLandCamp(log.Row)
}

func (log *GlobalDespairLandCampDeleteLog) Rollback() {
	key := log.Old.CampType
	var old *C.GlobalDespairLandCamp
	for crow := C.g_GlobalTables.GlobalDespairLandCamp; crow != nil; crow = crow.next {
		if crow.CampType == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'global_despair_land_camp'")
	}
	log.Old.next = C.g_GlobalTables.GlobalDespairLandCamp
	C.g_GlobalTables.GlobalDespairLandCamp = log.Old
}

func (log *GlobalDespairLandCampUpdateLog) Rollback() {
	key := log.Old.CampType
	var old, prev *C.GlobalDespairLandCamp
	for crow := C.g_GlobalTables.GlobalDespairLandCamp; crow != nil; prev, crow = crow, crow.next {
		if crow.CampType == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'global_despair_land_camp'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		C.g_GlobalTables.GlobalDespairLandCamp = log.Old
	}
	C.Free_GlobalDespairLandCamp(log.Row)
}

/* ========== global_gift_card_record ========== */

type GlobalGiftCardRecordInsertLog struct {
	db *Database 
	Row *C.GlobalGiftCardRecord
	GoRow *GlobalGiftCardRecord
}

func (db *Database) newGlobalGiftCardRecordInsertLog(row *C.GlobalGiftCardRecord, goRow *GlobalGiftCardRecord) *GlobalGiftCardRecordInsertLog {
	return &GlobalGiftCardRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *GlobalGiftCardRecordInsertLog) Free() {
}

func (log *GlobalGiftCardRecordInsertLog) InvokeHook() {
	if g_Hooks.GlobalGiftCardRecord != nil {
		g_Hooks.GlobalGiftCardRecord.Insert(&GlobalGiftCardRecordRow{row: log.Row})
	}
}

func (log *GlobalGiftCardRecordInsertLog) SetEventId(time.Time) {
}

func (log *GlobalGiftCardRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(9)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteString(log.GoRow.Code)
	file.WriteInt64(log.GoRow.Timestamp)
}

func (log *GlobalGiftCardRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.GlobalGiftCardRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Code), int(log.Row.Code_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	return stmt.Execute()
}

func (log *GlobalGiftCardRecordInsertLog) CommitToTLog() {
}

func (log *GlobalGiftCardRecordInsertLog) CommitToXdLog() {
}

func (log *GlobalGiftCardRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalGiftCardRecordDeleteLog struct {
	db *Database
	Old *C.GlobalGiftCardRecord
	GoOld *GlobalGiftCardRecord
}

func (db *Database) newGlobalGiftCardRecordDeleteLog(old *C.GlobalGiftCardRecord, goOld *GlobalGiftCardRecord) *GlobalGiftCardRecordDeleteLog {
	return &GlobalGiftCardRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *GlobalGiftCardRecordDeleteLog) Free() {
	C.Free_GlobalGiftCardRecord(log.Old)
}

func (log *GlobalGiftCardRecordDeleteLog) InvokeHook() {
	if g_Hooks.GlobalGiftCardRecord != nil {
		g_Hooks.GlobalGiftCardRecord.Delete(&GlobalGiftCardRecordRow{row: log.Old})
	}
}

func (log *GlobalGiftCardRecordDeleteLog) SetEventId(time.Time) {
}

func (log *GlobalGiftCardRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(9)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteString(log.GoOld.Code)
	file.WriteInt64(log.GoOld.Timestamp)
}

func (log *GlobalGiftCardRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.GlobalGiftCardRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *GlobalGiftCardRecordDeleteLog) CommitToTLog() {
}

func (log *GlobalGiftCardRecordDeleteLog) CommitToXdLog() {
}

func (log *GlobalGiftCardRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalGiftCardRecordUpdateLog struct {
	db *Database 
	Row *C.GlobalGiftCardRecord
	Old *C.GlobalGiftCardRecord
	GoNew *GlobalGiftCardRecord
	GoOld *GlobalGiftCardRecord
}

func (db *Database) newGlobalGiftCardRecordUpdateLog(row, old *C.GlobalGiftCardRecord, goNew, goOld *GlobalGiftCardRecord) *GlobalGiftCardRecordUpdateLog {
	return &GlobalGiftCardRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *GlobalGiftCardRecordUpdateLog) Free() {
	C.Free_GlobalGiftCardRecord(log.Old)
}

func (log *GlobalGiftCardRecordUpdateLog) InvokeHook() {
	if g_Hooks.GlobalGiftCardRecord != nil {
		g_Hooks.GlobalGiftCardRecord.Update(&GlobalGiftCardRecordRow{row: log.Row}, &GlobalGiftCardRecordRow{row: log.Old})
	}
}

func (log *GlobalGiftCardRecordUpdateLog) SetEventId(time.Time) {
}

func (log *GlobalGiftCardRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(9)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteString(log.GoNew.Code)
	file.WriteInt64(log.GoNew.Timestamp)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteString(log.GoOld.Code)
	file.WriteInt64(log.GoOld.Timestamp)
}

func (log *GlobalGiftCardRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.GlobalGiftCardRecord.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Code), int(log.Row.Code_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *GlobalGiftCardRecordUpdateLog) CommitToTLog() {
}

func (log *GlobalGiftCardRecordUpdateLog) CommitToXdLog() {
}

func (log *GlobalGiftCardRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *GlobalGiftCardRecordInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.GlobalGiftCardRecord
	for crow := C.g_GlobalTables.GlobalGiftCardRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'GlobalGiftCardRecord'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalGiftCardRecord = C.g_GlobalTables.GlobalGiftCardRecord.next
	}
	C.Free_GlobalGiftCardRecord(log.Row)
}

func (log *GlobalGiftCardRecordDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.GlobalGiftCardRecord
	for crow := C.g_GlobalTables.GlobalGiftCardRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'global_gift_card_record'")
	}
	log.Old.next = C.g_GlobalTables.GlobalGiftCardRecord
	C.g_GlobalTables.GlobalGiftCardRecord = log.Old
}

func (log *GlobalGiftCardRecordUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.GlobalGiftCardRecord
	for crow := C.g_GlobalTables.GlobalGiftCardRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'global_gift_card_record'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		C.g_GlobalTables.GlobalGiftCardRecord = log.Old
	}
	C.Free_GlobalGiftCardRecord(log.Row)
}

/* ========== global_group_buy_status ========== */

type GlobalGroupBuyStatusInsertLog struct {
	db *Database 
	Row *C.GlobalGroupBuyStatus
	GoRow *GlobalGroupBuyStatus
}

func (db *Database) newGlobalGroupBuyStatusInsertLog(row *C.GlobalGroupBuyStatus, goRow *GlobalGroupBuyStatus) *GlobalGroupBuyStatusInsertLog {
	return &GlobalGroupBuyStatusInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *GlobalGroupBuyStatusInsertLog) Free() {
}

func (log *GlobalGroupBuyStatusInsertLog) InvokeHook() {
	if g_Hooks.GlobalGroupBuyStatus != nil {
		g_Hooks.GlobalGroupBuyStatus.Insert(&GlobalGroupBuyStatusRow{row: log.Row})
	}
}

func (log *GlobalGroupBuyStatusInsertLog) SetEventId(time.Time) {
}

func (log *GlobalGroupBuyStatusInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(10)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt32(log.GoRow.Cid)
	file.WriteInt32(log.GoRow.Status)
}

func (log *GlobalGroupBuyStatusInsertLog) CommitToMySql() error {
	stmt := g_SQL.GlobalGroupBuyStatus.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Cid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Status)))
	return stmt.Execute()
}

func (log *GlobalGroupBuyStatusInsertLog) CommitToTLog() {
}

func (log *GlobalGroupBuyStatusInsertLog) CommitToXdLog() {
}

func (log *GlobalGroupBuyStatusInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalGroupBuyStatusDeleteLog struct {
	db *Database
	Old *C.GlobalGroupBuyStatus
	GoOld *GlobalGroupBuyStatus
}

func (db *Database) newGlobalGroupBuyStatusDeleteLog(old *C.GlobalGroupBuyStatus, goOld *GlobalGroupBuyStatus) *GlobalGroupBuyStatusDeleteLog {
	return &GlobalGroupBuyStatusDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *GlobalGroupBuyStatusDeleteLog) Free() {
	C.Free_GlobalGroupBuyStatus(log.Old)
}

func (log *GlobalGroupBuyStatusDeleteLog) InvokeHook() {
	if g_Hooks.GlobalGroupBuyStatus != nil {
		g_Hooks.GlobalGroupBuyStatus.Delete(&GlobalGroupBuyStatusRow{row: log.Old})
	}
}

func (log *GlobalGroupBuyStatusDeleteLog) SetEventId(time.Time) {
}

func (log *GlobalGroupBuyStatusDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(10)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt32(log.GoOld.Cid)
	file.WriteInt32(log.GoOld.Status)
}

func (log *GlobalGroupBuyStatusDeleteLog) CommitToMySql() error {
	stmt := g_SQL.GlobalGroupBuyStatus.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *GlobalGroupBuyStatusDeleteLog) CommitToTLog() {
}

func (log *GlobalGroupBuyStatusDeleteLog) CommitToXdLog() {
}

func (log *GlobalGroupBuyStatusDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalGroupBuyStatusUpdateLog struct {
	db *Database 
	Row *C.GlobalGroupBuyStatus
	Old *C.GlobalGroupBuyStatus
	GoNew *GlobalGroupBuyStatus
	GoOld *GlobalGroupBuyStatus
}

func (db *Database) newGlobalGroupBuyStatusUpdateLog(row, old *C.GlobalGroupBuyStatus, goNew, goOld *GlobalGroupBuyStatus) *GlobalGroupBuyStatusUpdateLog {
	return &GlobalGroupBuyStatusUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *GlobalGroupBuyStatusUpdateLog) Free() {
	C.Free_GlobalGroupBuyStatus(log.Old)
}

func (log *GlobalGroupBuyStatusUpdateLog) InvokeHook() {
	if g_Hooks.GlobalGroupBuyStatus != nil {
		g_Hooks.GlobalGroupBuyStatus.Update(&GlobalGroupBuyStatusRow{row: log.Row}, &GlobalGroupBuyStatusRow{row: log.Old})
	}
}

func (log *GlobalGroupBuyStatusUpdateLog) SetEventId(time.Time) {
}

func (log *GlobalGroupBuyStatusUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(10)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt32(log.GoNew.Cid)
	file.WriteInt32(log.GoNew.Status)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt32(log.GoOld.Cid)
	file.WriteInt32(log.GoOld.Status)
}

func (log *GlobalGroupBuyStatusUpdateLog) CommitToMySql() error {
	stmt := g_SQL.GlobalGroupBuyStatus.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.Cid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Status)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *GlobalGroupBuyStatusUpdateLog) CommitToTLog() {
}

func (log *GlobalGroupBuyStatusUpdateLog) CommitToXdLog() {
}

func (log *GlobalGroupBuyStatusUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *GlobalGroupBuyStatusInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.GlobalGroupBuyStatus
	for crow := C.g_GlobalTables.GlobalGroupBuyStatus; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'GlobalGroupBuyStatus'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalGroupBuyStatus = C.g_GlobalTables.GlobalGroupBuyStatus.next
	}
	C.Free_GlobalGroupBuyStatus(log.Row)
}

func (log *GlobalGroupBuyStatusDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.GlobalGroupBuyStatus
	for crow := C.g_GlobalTables.GlobalGroupBuyStatus; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'global_group_buy_status'")
	}
	log.Old.next = C.g_GlobalTables.GlobalGroupBuyStatus
	C.g_GlobalTables.GlobalGroupBuyStatus = log.Old
}

func (log *GlobalGroupBuyStatusUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.GlobalGroupBuyStatus
	for crow := C.g_GlobalTables.GlobalGroupBuyStatus; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'global_group_buy_status'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		C.g_GlobalTables.GlobalGroupBuyStatus = log.Old
	}
	C.Free_GlobalGroupBuyStatus(log.Row)
}

/* ========== global_mail ========== */

type GlobalMailInsertLog struct {
	db *Database 
	Row *C.GlobalMail
	GoRow *GlobalMail
}

func (db *Database) newGlobalMailInsertLog(row *C.GlobalMail, goRow *GlobalMail) *GlobalMailInsertLog {
	return &GlobalMailInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *GlobalMailInsertLog) Free() {
}

func (log *GlobalMailInsertLog) InvokeHook() {
	if g_Hooks.GlobalMail != nil {
		g_Hooks.GlobalMail.Insert(&GlobalMailRow{row: log.Row})
	}
}

func (log *GlobalMailInsertLog) SetEventId(time.Time) {
}

func (log *GlobalMailInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(11)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.SendTime)
	file.WriteInt64(log.GoRow.ExpireTime)
	file.WriteString(log.GoRow.Title)
	file.WriteString(log.GoRow.Content)
	file.WriteInt8(log.GoRow.Priority)
	file.WriteInt16(log.GoRow.MinLevel)
	file.WriteInt16(log.GoRow.MaxLevel)
	file.WriteInt16(log.GoRow.MinVipLevel)
	file.WriteInt16(log.GoRow.MaxVipLevel)
	file.WriteInt64(log.GoRow.MinCreateTime)
	file.WriteInt64(log.GoRow.MaxCreateTime)
}

func (log *GlobalMailInsertLog) CommitToMySql() error {
	stmt := g_SQL.GlobalMail.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SendTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ExpireTime)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Title), int(log.Row.Title_len))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Content), int(log.Row.Content_len))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Priority)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.MinLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.MaxLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.MinVipLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.MaxVipLevel)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.MinCreateTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.MaxCreateTime)))
	return stmt.Execute()
}

func (log *GlobalMailInsertLog) CommitToTLog() {
}

func (log *GlobalMailInsertLog) CommitToXdLog() {
}

func (log *GlobalMailInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalMailDeleteLog struct {
	db *Database
	Old *C.GlobalMail
	GoOld *GlobalMail
}

func (db *Database) newGlobalMailDeleteLog(old *C.GlobalMail, goOld *GlobalMail) *GlobalMailDeleteLog {
	return &GlobalMailDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *GlobalMailDeleteLog) Free() {
	C.Free_GlobalMail(log.Old)
}

func (log *GlobalMailDeleteLog) InvokeHook() {
	if g_Hooks.GlobalMail != nil {
		g_Hooks.GlobalMail.Delete(&GlobalMailRow{row: log.Old})
	}
}

func (log *GlobalMailDeleteLog) SetEventId(time.Time) {
}

func (log *GlobalMailDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(11)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.SendTime)
	file.WriteInt64(log.GoOld.ExpireTime)
	file.WriteString(log.GoOld.Title)
	file.WriteString(log.GoOld.Content)
	file.WriteInt8(log.GoOld.Priority)
	file.WriteInt16(log.GoOld.MinLevel)
	file.WriteInt16(log.GoOld.MaxLevel)
	file.WriteInt16(log.GoOld.MinVipLevel)
	file.WriteInt16(log.GoOld.MaxVipLevel)
	file.WriteInt64(log.GoOld.MinCreateTime)
	file.WriteInt64(log.GoOld.MaxCreateTime)
}

func (log *GlobalMailDeleteLog) CommitToMySql() error {
	stmt := g_SQL.GlobalMail.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *GlobalMailDeleteLog) CommitToTLog() {
}

func (log *GlobalMailDeleteLog) CommitToXdLog() {
}

func (log *GlobalMailDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalMailUpdateLog struct {
	db *Database 
	Row *C.GlobalMail
	Old *C.GlobalMail
	GoNew *GlobalMail
	GoOld *GlobalMail
}

func (db *Database) newGlobalMailUpdateLog(row, old *C.GlobalMail, goNew, goOld *GlobalMail) *GlobalMailUpdateLog {
	return &GlobalMailUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *GlobalMailUpdateLog) Free() {
	C.Free_GlobalMail(log.Old)
}

func (log *GlobalMailUpdateLog) InvokeHook() {
	if g_Hooks.GlobalMail != nil {
		g_Hooks.GlobalMail.Update(&GlobalMailRow{row: log.Row}, &GlobalMailRow{row: log.Old})
	}
}

func (log *GlobalMailUpdateLog) SetEventId(time.Time) {
}

func (log *GlobalMailUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(11)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.SendTime)
	file.WriteInt64(log.GoNew.ExpireTime)
	file.WriteString(log.GoNew.Title)
	file.WriteString(log.GoNew.Content)
	file.WriteInt8(log.GoNew.Priority)
	file.WriteInt16(log.GoNew.MinLevel)
	file.WriteInt16(log.GoNew.MaxLevel)
	file.WriteInt16(log.GoNew.MinVipLevel)
	file.WriteInt16(log.GoNew.MaxVipLevel)
	file.WriteInt64(log.GoNew.MinCreateTime)
	file.WriteInt64(log.GoNew.MaxCreateTime)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.SendTime)
	file.WriteInt64(log.GoOld.ExpireTime)
	file.WriteString(log.GoOld.Title)
	file.WriteString(log.GoOld.Content)
	file.WriteInt8(log.GoOld.Priority)
	file.WriteInt16(log.GoOld.MinLevel)
	file.WriteInt16(log.GoOld.MaxLevel)
	file.WriteInt16(log.GoOld.MinVipLevel)
	file.WriteInt16(log.GoOld.MaxVipLevel)
	file.WriteInt64(log.GoOld.MinCreateTime)
	file.WriteInt64(log.GoOld.MaxCreateTime)
}

func (log *GlobalMailUpdateLog) CommitToMySql() error {
	stmt := g_SQL.GlobalMail.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SendTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ExpireTime)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Title), int(log.Row.Title_len))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Content), int(log.Row.Content_len))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Priority)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.MinLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.MaxLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.MinVipLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.MaxVipLevel)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.MinCreateTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.MaxCreateTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *GlobalMailUpdateLog) CommitToTLog() {
}

func (log *GlobalMailUpdateLog) CommitToXdLog() {
}

func (log *GlobalMailUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *GlobalMailInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.GlobalMail
	for crow := C.g_GlobalTables.GlobalMail; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'GlobalMail'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalMail = C.g_GlobalTables.GlobalMail.next
	}
	C.Free_GlobalMail(log.Row)
}

func (log *GlobalMailDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.GlobalMail
	for crow := C.g_GlobalTables.GlobalMail; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'global_mail'")
	}
	log.Old.next = C.g_GlobalTables.GlobalMail
	C.g_GlobalTables.GlobalMail = log.Old
}

func (log *GlobalMailUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.GlobalMail
	for crow := C.g_GlobalTables.GlobalMail; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'global_mail'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		C.g_GlobalTables.GlobalMail = log.Old
	}
	C.Free_GlobalMail(log.Row)
}

/* ========== global_mail_attachments ========== */

type GlobalMailAttachmentsInsertLog struct {
	db *Database 
	Row *C.GlobalMailAttachments
	GoRow *GlobalMailAttachments
}

func (db *Database) newGlobalMailAttachmentsInsertLog(row *C.GlobalMailAttachments, goRow *GlobalMailAttachments) *GlobalMailAttachmentsInsertLog {
	return &GlobalMailAttachmentsInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *GlobalMailAttachmentsInsertLog) Free() {
}

func (log *GlobalMailAttachmentsInsertLog) InvokeHook() {
	if g_Hooks.GlobalMailAttachments != nil {
		g_Hooks.GlobalMailAttachments.Insert(&GlobalMailAttachmentsRow{row: log.Row})
	}
}

func (log *GlobalMailAttachmentsInsertLog) SetEventId(time.Time) {
}

func (log *GlobalMailAttachmentsInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(12)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.GlobalMailId)
	file.WriteInt16(log.GoRow.ItemId)
	file.WriteInt8(log.GoRow.AttachmentType)
	file.WriteInt64(log.GoRow.ItemNum)
}

func (log *GlobalMailAttachmentsInsertLog) CommitToMySql() error {
	stmt := g_SQL.GlobalMailAttachments.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.GlobalMailId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.AttachmentType)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ItemNum)))
	return stmt.Execute()
}

func (log *GlobalMailAttachmentsInsertLog) CommitToTLog() {
}

func (log *GlobalMailAttachmentsInsertLog) CommitToXdLog() {
}

func (log *GlobalMailAttachmentsInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalMailAttachmentsDeleteLog struct {
	db *Database
	Old *C.GlobalMailAttachments
	GoOld *GlobalMailAttachments
}

func (db *Database) newGlobalMailAttachmentsDeleteLog(old *C.GlobalMailAttachments, goOld *GlobalMailAttachments) *GlobalMailAttachmentsDeleteLog {
	return &GlobalMailAttachmentsDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *GlobalMailAttachmentsDeleteLog) Free() {
	C.Free_GlobalMailAttachments(log.Old)
}

func (log *GlobalMailAttachmentsDeleteLog) InvokeHook() {
	if g_Hooks.GlobalMailAttachments != nil {
		g_Hooks.GlobalMailAttachments.Delete(&GlobalMailAttachmentsRow{row: log.Old})
	}
}

func (log *GlobalMailAttachmentsDeleteLog) SetEventId(time.Time) {
}

func (log *GlobalMailAttachmentsDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(12)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.GlobalMailId)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt8(log.GoOld.AttachmentType)
	file.WriteInt64(log.GoOld.ItemNum)
}

func (log *GlobalMailAttachmentsDeleteLog) CommitToMySql() error {
	stmt := g_SQL.GlobalMailAttachments.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *GlobalMailAttachmentsDeleteLog) CommitToTLog() {
}

func (log *GlobalMailAttachmentsDeleteLog) CommitToXdLog() {
}

func (log *GlobalMailAttachmentsDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalMailAttachmentsUpdateLog struct {
	db *Database 
	Row *C.GlobalMailAttachments
	Old *C.GlobalMailAttachments
	GoNew *GlobalMailAttachments
	GoOld *GlobalMailAttachments
}

func (db *Database) newGlobalMailAttachmentsUpdateLog(row, old *C.GlobalMailAttachments, goNew, goOld *GlobalMailAttachments) *GlobalMailAttachmentsUpdateLog {
	return &GlobalMailAttachmentsUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *GlobalMailAttachmentsUpdateLog) Free() {
	C.Free_GlobalMailAttachments(log.Old)
}

func (log *GlobalMailAttachmentsUpdateLog) InvokeHook() {
	if g_Hooks.GlobalMailAttachments != nil {
		g_Hooks.GlobalMailAttachments.Update(&GlobalMailAttachmentsRow{row: log.Row}, &GlobalMailAttachmentsRow{row: log.Old})
	}
}

func (log *GlobalMailAttachmentsUpdateLog) SetEventId(time.Time) {
}

func (log *GlobalMailAttachmentsUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(12)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.GlobalMailId)
	file.WriteInt16(log.GoNew.ItemId)
	file.WriteInt8(log.GoNew.AttachmentType)
	file.WriteInt64(log.GoNew.ItemNum)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.GlobalMailId)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt8(log.GoOld.AttachmentType)
	file.WriteInt64(log.GoOld.ItemNum)
}

func (log *GlobalMailAttachmentsUpdateLog) CommitToMySql() error {
	stmt := g_SQL.GlobalMailAttachments.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.GlobalMailId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.AttachmentType)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ItemNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *GlobalMailAttachmentsUpdateLog) CommitToTLog() {
}

func (log *GlobalMailAttachmentsUpdateLog) CommitToXdLog() {
}

func (log *GlobalMailAttachmentsUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *GlobalMailAttachmentsInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.GlobalMailAttachments
	for crow := C.g_GlobalTables.GlobalMailAttachments; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'GlobalMailAttachments'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalMailAttachments = C.g_GlobalTables.GlobalMailAttachments.next
	}
	C.Free_GlobalMailAttachments(log.Row)
}

func (log *GlobalMailAttachmentsDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.GlobalMailAttachments
	for crow := C.g_GlobalTables.GlobalMailAttachments; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'global_mail_attachments'")
	}
	log.Old.next = C.g_GlobalTables.GlobalMailAttachments
	C.g_GlobalTables.GlobalMailAttachments = log.Old
}

func (log *GlobalMailAttachmentsUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.GlobalMailAttachments
	for crow := C.g_GlobalTables.GlobalMailAttachments; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'global_mail_attachments'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		C.g_GlobalTables.GlobalMailAttachments = log.Old
	}
	C.Free_GlobalMailAttachments(log.Row)
}

/* ========== global_tb_xxd_onlinecnt ========== */

type GlobalTbXxdOnlinecntInsertLog struct {
	db *Database 
	Row *C.GlobalTbXxdOnlinecnt
	GoRow *GlobalTbXxdOnlinecnt
}

func (db *Database) newGlobalTbXxdOnlinecntInsertLog(row *C.GlobalTbXxdOnlinecnt, goRow *GlobalTbXxdOnlinecnt) *GlobalTbXxdOnlinecntInsertLog {
	return &GlobalTbXxdOnlinecntInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *GlobalTbXxdOnlinecntInsertLog) Free() {
}

func (log *GlobalTbXxdOnlinecntInsertLog) InvokeHook() {
	if g_Hooks.GlobalTbXxdOnlinecnt != nil {
		g_Hooks.GlobalTbXxdOnlinecnt.Insert(&GlobalTbXxdOnlinecntRow{row: log.Row})
	}
}

func (log *GlobalTbXxdOnlinecntInsertLog) SetEventId(time.Time) {
}

func (log *GlobalTbXxdOnlinecntInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(13)
	file.WriteInt64(log.GoRow.Id)
	file.WriteString(log.GoRow.Gameappid)
	file.WriteInt64(log.GoRow.Timekey)
	file.WriteInt64(log.GoRow.Gsid)
	file.WriteInt64(log.GoRow.Onlinecntios)
	file.WriteInt64(log.GoRow.Onlinecntandroid)
	file.WriteInt64(log.GoRow.Cid)
}

func (log *GlobalTbXxdOnlinecntInsertLog) CommitToMySql() error {
	stmt := g_SQL.GlobalTbXxdOnlinecnt.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Gameappid), int(log.Row.Gameappid_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timekey)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Gsid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Onlinecntios)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Onlinecntandroid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Cid)))
	return stmt.Execute()
}

func (log *GlobalTbXxdOnlinecntInsertLog) CommitToTLog() {
}

func (log *GlobalTbXxdOnlinecntInsertLog) CommitToXdLog() {
}

func (log *GlobalTbXxdOnlinecntInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalTbXxdOnlinecntDeleteLog struct {
	db *Database
	Old *C.GlobalTbXxdOnlinecnt
	GoOld *GlobalTbXxdOnlinecnt
}

func (db *Database) newGlobalTbXxdOnlinecntDeleteLog(old *C.GlobalTbXxdOnlinecnt, goOld *GlobalTbXxdOnlinecnt) *GlobalTbXxdOnlinecntDeleteLog {
	return &GlobalTbXxdOnlinecntDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *GlobalTbXxdOnlinecntDeleteLog) Free() {
	C.Free_GlobalTbXxdOnlinecnt(log.Old)
}

func (log *GlobalTbXxdOnlinecntDeleteLog) InvokeHook() {
	if g_Hooks.GlobalTbXxdOnlinecnt != nil {
		g_Hooks.GlobalTbXxdOnlinecnt.Delete(&GlobalTbXxdOnlinecntRow{row: log.Old})
	}
}

func (log *GlobalTbXxdOnlinecntDeleteLog) SetEventId(time.Time) {
}

func (log *GlobalTbXxdOnlinecntDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(13)
	file.WriteInt64(log.GoOld.Id)
	file.WriteString(log.GoOld.Gameappid)
	file.WriteInt64(log.GoOld.Timekey)
	file.WriteInt64(log.GoOld.Gsid)
	file.WriteInt64(log.GoOld.Onlinecntios)
	file.WriteInt64(log.GoOld.Onlinecntandroid)
	file.WriteInt64(log.GoOld.Cid)
}

func (log *GlobalTbXxdOnlinecntDeleteLog) CommitToMySql() error {
	stmt := g_SQL.GlobalTbXxdOnlinecnt.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *GlobalTbXxdOnlinecntDeleteLog) CommitToTLog() {
}

func (log *GlobalTbXxdOnlinecntDeleteLog) CommitToXdLog() {
}

func (log *GlobalTbXxdOnlinecntDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type GlobalTbXxdOnlinecntUpdateLog struct {
	db *Database 
	Row *C.GlobalTbXxdOnlinecnt
	Old *C.GlobalTbXxdOnlinecnt
	GoNew *GlobalTbXxdOnlinecnt
	GoOld *GlobalTbXxdOnlinecnt
}

func (db *Database) newGlobalTbXxdOnlinecntUpdateLog(row, old *C.GlobalTbXxdOnlinecnt, goNew, goOld *GlobalTbXxdOnlinecnt) *GlobalTbXxdOnlinecntUpdateLog {
	return &GlobalTbXxdOnlinecntUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *GlobalTbXxdOnlinecntUpdateLog) Free() {
	C.Free_GlobalTbXxdOnlinecnt(log.Old)
}

func (log *GlobalTbXxdOnlinecntUpdateLog) InvokeHook() {
	if g_Hooks.GlobalTbXxdOnlinecnt != nil {
		g_Hooks.GlobalTbXxdOnlinecnt.Update(&GlobalTbXxdOnlinecntRow{row: log.Row}, &GlobalTbXxdOnlinecntRow{row: log.Old})
	}
}

func (log *GlobalTbXxdOnlinecntUpdateLog) SetEventId(time.Time) {
}

func (log *GlobalTbXxdOnlinecntUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(13)
	file.WriteInt64(log.GoNew.Id)
	file.WriteString(log.GoNew.Gameappid)
	file.WriteInt64(log.GoNew.Timekey)
	file.WriteInt64(log.GoNew.Gsid)
	file.WriteInt64(log.GoNew.Onlinecntios)
	file.WriteInt64(log.GoNew.Onlinecntandroid)
	file.WriteInt64(log.GoNew.Cid)
	file.WriteInt64(log.GoOld.Id)
	file.WriteString(log.GoOld.Gameappid)
	file.WriteInt64(log.GoOld.Timekey)
	file.WriteInt64(log.GoOld.Gsid)
	file.WriteInt64(log.GoOld.Onlinecntios)
	file.WriteInt64(log.GoOld.Onlinecntandroid)
	file.WriteInt64(log.GoOld.Cid)
}

func (log *GlobalTbXxdOnlinecntUpdateLog) CommitToMySql() error {
	stmt := g_SQL.GlobalTbXxdOnlinecnt.Update
	stmt.BindVarchar(unsafe.Pointer(log.Row.Gameappid), int(log.Row.Gameappid_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timekey)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Gsid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Onlinecntios)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Onlinecntandroid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Cid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *GlobalTbXxdOnlinecntUpdateLog) CommitToTLog() {
}

func (log *GlobalTbXxdOnlinecntUpdateLog) CommitToXdLog() {
}

func (log *GlobalTbXxdOnlinecntUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *GlobalTbXxdOnlinecntInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.GlobalTbXxdOnlinecnt
	for crow := C.g_GlobalTables.GlobalTbXxdOnlinecnt; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'GlobalTbXxdOnlinecnt'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		C.g_GlobalTables.GlobalTbXxdOnlinecnt = C.g_GlobalTables.GlobalTbXxdOnlinecnt.next
	}
	C.Free_GlobalTbXxdOnlinecnt(log.Row)
}

func (log *GlobalTbXxdOnlinecntDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.GlobalTbXxdOnlinecnt
	for crow := C.g_GlobalTables.GlobalTbXxdOnlinecnt; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'global_tb_xxd_onlinecnt'")
	}
	log.Old.next = C.g_GlobalTables.GlobalTbXxdOnlinecnt
	C.g_GlobalTables.GlobalTbXxdOnlinecnt = log.Old
}

func (log *GlobalTbXxdOnlinecntUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.GlobalTbXxdOnlinecnt
	for crow := C.g_GlobalTables.GlobalTbXxdOnlinecnt; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'global_tb_xxd_onlinecnt'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		C.g_GlobalTables.GlobalTbXxdOnlinecnt = log.Old
	}
	C.Free_GlobalTbXxdOnlinecnt(log.Row)
}

/* ========== player ========== */

type PlayerInsertLog struct {
	db *Database 
	Row *C.Player
	GoRow *Player
}

func (db *Database) newPlayerInsertLog(row *C.Player, goRow *Player) *PlayerInsertLog {
	return &PlayerInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerInsertLog) Free() {
}

func (log *PlayerInsertLog) InvokeHook() {
	if g_Hooks.Player != nil {
		g_Hooks.Player.Insert(&PlayerRow{row: log.Row})
	}
}

func (log *PlayerInsertLog) SetEventId(time.Time) {
}

func (log *PlayerInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(14)
	file.WriteInt64(log.GoRow.Id)
	file.WriteString(log.GoRow.User)
	file.WriteString(log.GoRow.Nick)
	file.WriteInt64(log.GoRow.MainRoleId)
	file.WriteInt64(log.GoRow.Cid)
}

func (log *PlayerInsertLog) CommitToMySql() error {
	stmt := g_SQL.Player.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.User), int(log.Row.User_len))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Nick), int(log.Row.Nick_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.MainRoleId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Cid)))
	return stmt.Execute()
}

func (log *PlayerInsertLog) CommitToTLog() {
}

func (log *PlayerInsertLog) CommitToXdLog() {
}

func (log *PlayerInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDeleteLog struct {
	db *Database
	Old *C.Player
	GoOld *Player
}

func (db *Database) newPlayerDeleteLog(old *C.Player, goOld *Player) *PlayerDeleteLog {
	return &PlayerDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerDeleteLog) Free() {
	C.Free_Player(log.Old)
}

func (log *PlayerDeleteLog) InvokeHook() {
	if g_Hooks.Player != nil {
		g_Hooks.Player.Delete(&PlayerRow{row: log.Old})
	}
}

func (log *PlayerDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(14)
	file.WriteInt64(log.GoOld.Id)
	file.WriteString(log.GoOld.User)
	file.WriteString(log.GoOld.Nick)
	file.WriteInt64(log.GoOld.MainRoleId)
	file.WriteInt64(log.GoOld.Cid)
}

func (log *PlayerDeleteLog) CommitToMySql() error {
	stmt := g_SQL.Player.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerDeleteLog) CommitToTLog() {
}

func (log *PlayerDeleteLog) CommitToXdLog() {
}

func (log *PlayerDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerUpdateLog struct {
	db *Database 
	Row *C.Player
	Old *C.Player
	GoNew *Player
	GoOld *Player
}

func (db *Database) newPlayerUpdateLog(row, old *C.Player, goNew, goOld *Player) *PlayerUpdateLog {
	return &PlayerUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerUpdateLog) Free() {
	C.Free_Player(log.Old)
}

func (log *PlayerUpdateLog) InvokeHook() {
	if g_Hooks.Player != nil {
		g_Hooks.Player.Update(&PlayerRow{row: log.Row}, &PlayerRow{row: log.Old})
	}
}

func (log *PlayerUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(14)
	file.WriteInt64(log.GoNew.Id)
	file.WriteString(log.GoNew.User)
	file.WriteString(log.GoNew.Nick)
	file.WriteInt64(log.GoNew.MainRoleId)
	file.WriteInt64(log.GoNew.Cid)
	file.WriteInt64(log.GoOld.Id)
	file.WriteString(log.GoOld.User)
	file.WriteString(log.GoOld.Nick)
	file.WriteInt64(log.GoOld.MainRoleId)
	file.WriteInt64(log.GoOld.Cid)
}

func (log *PlayerUpdateLog) CommitToMySql() error {
	stmt := g_SQL.Player.Update
	stmt.BindVarchar(unsafe.Pointer(log.Row.User), int(log.Row.User_len))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Nick), int(log.Row.Nick_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.MainRoleId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Cid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerUpdateLog) CommitToTLog() {
}

func (log *PlayerUpdateLog) CommitToXdLog() {
}

func (log *PlayerUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerInsertLog) Rollback() {
	if log.db.tables.Player != log.Row {
		panic("insert rollback check failed 'Player'")
	}
	log.db.tables.Player = nil
	C.Free_Player(log.Row)
}

func (log *PlayerDeleteLog) Rollback() {
	if log.db.tables.Player != nil {
		panic("delete rollback check failed 'player'")
	}
	log.db.tables.Player = log.Old
}

func (log *PlayerUpdateLog) Rollback() {
	if log.db.tables.Player != log.Row {
		panic("update rollback check failed 'player'")
	}
	log.db.tables.Player = log.Old
	C.Free_Player(log.Row)
}

/* ========== player_activity ========== */

type PlayerActivityInsertLog struct {
	db *Database 
	Row *C.PlayerActivity
	GoRow *PlayerActivity
}

func (db *Database) newPlayerActivityInsertLog(row *C.PlayerActivity, goRow *PlayerActivity) *PlayerActivityInsertLog {
	return &PlayerActivityInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerActivityInsertLog) Free() {
}

func (log *PlayerActivityInsertLog) InvokeHook() {
	if g_Hooks.PlayerActivity != nil {
		g_Hooks.PlayerActivity.Insert(&PlayerActivityRow{row: log.Row})
	}
}

func (log *PlayerActivityInsertLog) SetEventId(time.Time) {
}

func (log *PlayerActivityInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(15)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.Activity)
	file.WriteInt64(log.GoRow.LastUpdate)
}

func (log *PlayerActivityInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerActivity.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Activity)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastUpdate)))
	return stmt.Execute()
}

func (log *PlayerActivityInsertLog) CommitToTLog() {
}

func (log *PlayerActivityInsertLog) CommitToXdLog() {
}

func (log *PlayerActivityInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerActivityDeleteLog struct {
	db *Database
	Old *C.PlayerActivity
	GoOld *PlayerActivity
}

func (db *Database) newPlayerActivityDeleteLog(old *C.PlayerActivity, goOld *PlayerActivity) *PlayerActivityDeleteLog {
	return &PlayerActivityDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerActivityDeleteLog) Free() {
	C.Free_PlayerActivity(log.Old)
}

func (log *PlayerActivityDeleteLog) InvokeHook() {
	if g_Hooks.PlayerActivity != nil {
		g_Hooks.PlayerActivity.Delete(&PlayerActivityRow{row: log.Old})
	}
}

func (log *PlayerActivityDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerActivityDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(15)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Activity)
	file.WriteInt64(log.GoOld.LastUpdate)
}

func (log *PlayerActivityDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerActivity.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerActivityDeleteLog) CommitToTLog() {
}

func (log *PlayerActivityDeleteLog) CommitToXdLog() {
}

func (log *PlayerActivityDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerActivityUpdateLog struct {
	db *Database 
	Row *C.PlayerActivity
	Old *C.PlayerActivity
	GoNew *PlayerActivity
	GoOld *PlayerActivity
}

func (db *Database) newPlayerActivityUpdateLog(row, old *C.PlayerActivity, goNew, goOld *PlayerActivity) *PlayerActivityUpdateLog {
	return &PlayerActivityUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerActivityUpdateLog) Free() {
	C.Free_PlayerActivity(log.Old)
}

func (log *PlayerActivityUpdateLog) InvokeHook() {
	if g_Hooks.PlayerActivity != nil {
		g_Hooks.PlayerActivity.Update(&PlayerActivityRow{row: log.Row}, &PlayerActivityRow{row: log.Old})
	}
}

func (log *PlayerActivityUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerActivityUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(15)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.Activity)
	file.WriteInt64(log.GoNew.LastUpdate)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Activity)
	file.WriteInt64(log.GoOld.LastUpdate)
}

func (log *PlayerActivityUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerActivity.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.Activity)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastUpdate)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerActivityUpdateLog) CommitToTLog() {
}

func (log *PlayerActivityUpdateLog) CommitToXdLog() {
}

func (log *PlayerActivityUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerActivityInsertLog) Rollback() {
	if log.db.tables.PlayerActivity != log.Row {
		panic("insert rollback check failed 'PlayerActivity'")
	}
	log.db.tables.PlayerActivity = nil
	C.Free_PlayerActivity(log.Row)
}

func (log *PlayerActivityDeleteLog) Rollback() {
	if log.db.tables.PlayerActivity != nil {
		panic("delete rollback check failed 'player_activity'")
	}
	log.db.tables.PlayerActivity = log.Old
}

func (log *PlayerActivityUpdateLog) Rollback() {
	if log.db.tables.PlayerActivity != log.Row {
		panic("update rollback check failed 'player_activity'")
	}
	log.db.tables.PlayerActivity = log.Old
	C.Free_PlayerActivity(log.Row)
}

/* ========== player_addition_quest ========== */

type PlayerAdditionQuestInsertLog struct {
	db *Database 
	Row *C.PlayerAdditionQuest
	GoRow *PlayerAdditionQuest
}

func (db *Database) newPlayerAdditionQuestInsertLog(row *C.PlayerAdditionQuest, goRow *PlayerAdditionQuest) *PlayerAdditionQuestInsertLog {
	return &PlayerAdditionQuestInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerAdditionQuestInsertLog) Free() {
}

func (log *PlayerAdditionQuestInsertLog) InvokeHook() {
	if g_Hooks.PlayerAdditionQuest != nil {
		g_Hooks.PlayerAdditionQuest.Insert(&PlayerAdditionQuestRow{row: log.Row})
	}
}

func (log *PlayerAdditionQuestInsertLog) SetEventId(time.Time) {
}

func (log *PlayerAdditionQuestInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(16)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.SerialNumber)
	file.WriteInt32(log.GoRow.QuestId)
	file.WriteInt16(log.GoRow.Lock)
	file.WriteInt16(log.GoRow.Progress)
	file.WriteInt8(log.GoRow.State)
}

func (log *PlayerAdditionQuestInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerAdditionQuest.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.SerialNumber)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.QuestId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Lock)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Progress)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.State)))
	return stmt.Execute()
}

func (log *PlayerAdditionQuestInsertLog) CommitToTLog() {
}

func (log *PlayerAdditionQuestInsertLog) CommitToXdLog() {
}

func (log *PlayerAdditionQuestInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerAdditionQuestDeleteLog struct {
	db *Database
	Old *C.PlayerAdditionQuest
	GoOld *PlayerAdditionQuest
}

func (db *Database) newPlayerAdditionQuestDeleteLog(old *C.PlayerAdditionQuest, goOld *PlayerAdditionQuest) *PlayerAdditionQuestDeleteLog {
	return &PlayerAdditionQuestDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerAdditionQuestDeleteLog) Free() {
	C.Free_PlayerAdditionQuest(log.Old)
}

func (log *PlayerAdditionQuestDeleteLog) InvokeHook() {
	if g_Hooks.PlayerAdditionQuest != nil {
		g_Hooks.PlayerAdditionQuest.Delete(&PlayerAdditionQuestRow{row: log.Old})
	}
}

func (log *PlayerAdditionQuestDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerAdditionQuestDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(16)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.SerialNumber)
	file.WriteInt32(log.GoOld.QuestId)
	file.WriteInt16(log.GoOld.Lock)
	file.WriteInt16(log.GoOld.Progress)
	file.WriteInt8(log.GoOld.State)
}

func (log *PlayerAdditionQuestDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerAdditionQuest.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerAdditionQuestDeleteLog) CommitToTLog() {
}

func (log *PlayerAdditionQuestDeleteLog) CommitToXdLog() {
}

func (log *PlayerAdditionQuestDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerAdditionQuestUpdateLog struct {
	db *Database 
	Row *C.PlayerAdditionQuest
	Old *C.PlayerAdditionQuest
	GoNew *PlayerAdditionQuest
	GoOld *PlayerAdditionQuest
}

func (db *Database) newPlayerAdditionQuestUpdateLog(row, old *C.PlayerAdditionQuest, goNew, goOld *PlayerAdditionQuest) *PlayerAdditionQuestUpdateLog {
	return &PlayerAdditionQuestUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerAdditionQuestUpdateLog) Free() {
	C.Free_PlayerAdditionQuest(log.Old)
}

func (log *PlayerAdditionQuestUpdateLog) InvokeHook() {
	if g_Hooks.PlayerAdditionQuest != nil {
		g_Hooks.PlayerAdditionQuest.Update(&PlayerAdditionQuestRow{row: log.Row}, &PlayerAdditionQuestRow{row: log.Old})
	}
}

func (log *PlayerAdditionQuestUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerAdditionQuestUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(16)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.SerialNumber)
	file.WriteInt32(log.GoNew.QuestId)
	file.WriteInt16(log.GoNew.Lock)
	file.WriteInt16(log.GoNew.Progress)
	file.WriteInt8(log.GoNew.State)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.SerialNumber)
	file.WriteInt32(log.GoOld.QuestId)
	file.WriteInt16(log.GoOld.Lock)
	file.WriteInt16(log.GoOld.Progress)
	file.WriteInt8(log.GoOld.State)
}

func (log *PlayerAdditionQuestUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerAdditionQuest.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.SerialNumber)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.QuestId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Lock)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Progress)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.State)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerAdditionQuestUpdateLog) CommitToTLog() {
}

func (log *PlayerAdditionQuestUpdateLog) CommitToXdLog() {
}

func (log *PlayerAdditionQuestUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerAdditionQuestInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerAdditionQuest
	for crow := log.db.tables.PlayerAdditionQuest; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerAdditionQuest'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerAdditionQuest = log.db.tables.PlayerAdditionQuest.next
	}
	C.Free_PlayerAdditionQuest(log.Row)
}

func (log *PlayerAdditionQuestDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerAdditionQuest
	for crow := log.db.tables.PlayerAdditionQuest; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_addition_quest'")
	}
	log.Old.next = log.db.tables.PlayerAdditionQuest
	log.db.tables.PlayerAdditionQuest = log.Old
}

func (log *PlayerAdditionQuestUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerAdditionQuest
	for crow := log.db.tables.PlayerAdditionQuest; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_addition_quest'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerAdditionQuest = log.Old
	}
	C.Free_PlayerAdditionQuest(log.Row)
}

/* ========== player_any_sdk_order ========== */

type PlayerAnySdkOrderInsertLog struct {
	db *Database 
	Row *C.PlayerAnySdkOrder
	GoRow *PlayerAnySdkOrder
}

func (db *Database) newPlayerAnySdkOrderInsertLog(row *C.PlayerAnySdkOrder, goRow *PlayerAnySdkOrder) *PlayerAnySdkOrderInsertLog {
	return &PlayerAnySdkOrderInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerAnySdkOrderInsertLog) Free() {
}

func (log *PlayerAnySdkOrderInsertLog) InvokeHook() {
	if g_Hooks.PlayerAnySdkOrder != nil {
		g_Hooks.PlayerAnySdkOrder.Insert(&PlayerAnySdkOrderRow{row: log.Row})
	}
}

func (log *PlayerAnySdkOrderInsertLog) SetEventId(time.Time) {
}

func (log *PlayerAnySdkOrderInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(17)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.OrderId)
	file.WriteInt64(log.GoRow.Present)
}

func (log *PlayerAnySdkOrderInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerAnySdkOrder.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.OrderId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Present)))
	return stmt.Execute()
}

func (log *PlayerAnySdkOrderInsertLog) CommitToTLog() {
}

func (log *PlayerAnySdkOrderInsertLog) CommitToXdLog() {
}

func (log *PlayerAnySdkOrderInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerAnySdkOrderDeleteLog struct {
	db *Database
	Old *C.PlayerAnySdkOrder
	GoOld *PlayerAnySdkOrder
}

func (db *Database) newPlayerAnySdkOrderDeleteLog(old *C.PlayerAnySdkOrder, goOld *PlayerAnySdkOrder) *PlayerAnySdkOrderDeleteLog {
	return &PlayerAnySdkOrderDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerAnySdkOrderDeleteLog) Free() {
	C.Free_PlayerAnySdkOrder(log.Old)
}

func (log *PlayerAnySdkOrderDeleteLog) InvokeHook() {
	if g_Hooks.PlayerAnySdkOrder != nil {
		g_Hooks.PlayerAnySdkOrder.Delete(&PlayerAnySdkOrderRow{row: log.Old})
	}
}

func (log *PlayerAnySdkOrderDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerAnySdkOrderDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(17)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.OrderId)
	file.WriteInt64(log.GoOld.Present)
}

func (log *PlayerAnySdkOrderDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerAnySdkOrder.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerAnySdkOrderDeleteLog) CommitToTLog() {
}

func (log *PlayerAnySdkOrderDeleteLog) CommitToXdLog() {
}

func (log *PlayerAnySdkOrderDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerAnySdkOrderUpdateLog struct {
	db *Database 
	Row *C.PlayerAnySdkOrder
	Old *C.PlayerAnySdkOrder
	GoNew *PlayerAnySdkOrder
	GoOld *PlayerAnySdkOrder
}

func (db *Database) newPlayerAnySdkOrderUpdateLog(row, old *C.PlayerAnySdkOrder, goNew, goOld *PlayerAnySdkOrder) *PlayerAnySdkOrderUpdateLog {
	return &PlayerAnySdkOrderUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerAnySdkOrderUpdateLog) Free() {
	C.Free_PlayerAnySdkOrder(log.Old)
}

func (log *PlayerAnySdkOrderUpdateLog) InvokeHook() {
	if g_Hooks.PlayerAnySdkOrder != nil {
		g_Hooks.PlayerAnySdkOrder.Update(&PlayerAnySdkOrderRow{row: log.Row}, &PlayerAnySdkOrderRow{row: log.Old})
	}
}

func (log *PlayerAnySdkOrderUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerAnySdkOrderUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(17)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.OrderId)
	file.WriteInt64(log.GoNew.Present)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.OrderId)
	file.WriteInt64(log.GoOld.Present)
}

func (log *PlayerAnySdkOrderUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerAnySdkOrder.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.OrderId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Present)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerAnySdkOrderUpdateLog) CommitToTLog() {
}

func (log *PlayerAnySdkOrderUpdateLog) CommitToXdLog() {
}

func (log *PlayerAnySdkOrderUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerAnySdkOrderInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerAnySdkOrder
	for crow := log.db.tables.PlayerAnySdkOrder; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerAnySdkOrder'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerAnySdkOrder = log.db.tables.PlayerAnySdkOrder.next
	}
	C.Free_PlayerAnySdkOrder(log.Row)
}

func (log *PlayerAnySdkOrderDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerAnySdkOrder
	for crow := log.db.tables.PlayerAnySdkOrder; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_any_sdk_order'")
	}
	log.Old.next = log.db.tables.PlayerAnySdkOrder
	log.db.tables.PlayerAnySdkOrder = log.Old
}

func (log *PlayerAnySdkOrderUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerAnySdkOrder
	for crow := log.db.tables.PlayerAnySdkOrder; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_any_sdk_order'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerAnySdkOrder = log.Old
	}
	C.Free_PlayerAnySdkOrder(log.Row)
}

/* ========== player_arena ========== */

type PlayerArenaInsertLog struct {
	db *Database 
	Row *C.PlayerArena
	GoRow *PlayerArena
}

func (db *Database) newPlayerArenaInsertLog(row *C.PlayerArena, goRow *PlayerArena) *PlayerArenaInsertLog {
	return &PlayerArenaInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerArenaInsertLog) Free() {
}

func (log *PlayerArenaInsertLog) InvokeHook() {
	if g_Hooks.PlayerArena != nil {
		g_Hooks.PlayerArena.Insert(&PlayerArenaRow{row: log.Row})
	}
}

func (log *PlayerArenaInsertLog) SetEventId(time.Time) {
}

func (log *PlayerArenaInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(18)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.DailyNum)
	file.WriteInt64(log.GoRow.FailedCdTime)
	file.WriteInt64(log.GoRow.BattleTime)
	file.WriteInt16(log.GoRow.WinTimes)
	file.WriteInt32(log.GoRow.DailyAwardItem)
	file.WriteInt16(log.GoRow.NewRecordCount)
	file.WriteInt32(log.GoRow.DailyFame)
	file.WriteInt16(log.GoRow.BuyTimes)
}

func (log *PlayerArenaInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerArena.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FailedCdTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BattleTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.WinTimes)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.DailyAwardItem)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.NewRecordCount)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.DailyFame)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.BuyTimes)))
	return stmt.Execute()
}

func (log *PlayerArenaInsertLog) CommitToTLog() {
}

func (log *PlayerArenaInsertLog) CommitToXdLog() {
}

func (log *PlayerArenaInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerArenaDeleteLog struct {
	db *Database
	Old *C.PlayerArena
	GoOld *PlayerArena
}

func (db *Database) newPlayerArenaDeleteLog(old *C.PlayerArena, goOld *PlayerArena) *PlayerArenaDeleteLog {
	return &PlayerArenaDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerArenaDeleteLog) Free() {
	C.Free_PlayerArena(log.Old)
}

func (log *PlayerArenaDeleteLog) InvokeHook() {
	if g_Hooks.PlayerArena != nil {
		g_Hooks.PlayerArena.Delete(&PlayerArenaRow{row: log.Old})
	}
}

func (log *PlayerArenaDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerArenaDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(18)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.DailyNum)
	file.WriteInt64(log.GoOld.FailedCdTime)
	file.WriteInt64(log.GoOld.BattleTime)
	file.WriteInt16(log.GoOld.WinTimes)
	file.WriteInt32(log.GoOld.DailyAwardItem)
	file.WriteInt16(log.GoOld.NewRecordCount)
	file.WriteInt32(log.GoOld.DailyFame)
	file.WriteInt16(log.GoOld.BuyTimes)
}

func (log *PlayerArenaDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerArena.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerArenaDeleteLog) CommitToTLog() {
}

func (log *PlayerArenaDeleteLog) CommitToXdLog() {
}

func (log *PlayerArenaDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerArenaUpdateLog struct {
	db *Database 
	Row *C.PlayerArena
	Old *C.PlayerArena
	GoNew *PlayerArena
	GoOld *PlayerArena
}

func (db *Database) newPlayerArenaUpdateLog(row, old *C.PlayerArena, goNew, goOld *PlayerArena) *PlayerArenaUpdateLog {
	return &PlayerArenaUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerArenaUpdateLog) Free() {
	C.Free_PlayerArena(log.Old)
}

func (log *PlayerArenaUpdateLog) InvokeHook() {
	if g_Hooks.PlayerArena != nil {
		g_Hooks.PlayerArena.Update(&PlayerArenaRow{row: log.Row}, &PlayerArenaRow{row: log.Old})
	}
}

func (log *PlayerArenaUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerArenaUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(18)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.DailyNum)
	file.WriteInt64(log.GoNew.FailedCdTime)
	file.WriteInt64(log.GoNew.BattleTime)
	file.WriteInt16(log.GoNew.WinTimes)
	file.WriteInt32(log.GoNew.DailyAwardItem)
	file.WriteInt16(log.GoNew.NewRecordCount)
	file.WriteInt32(log.GoNew.DailyFame)
	file.WriteInt16(log.GoNew.BuyTimes)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.DailyNum)
	file.WriteInt64(log.GoOld.FailedCdTime)
	file.WriteInt64(log.GoOld.BattleTime)
	file.WriteInt16(log.GoOld.WinTimes)
	file.WriteInt32(log.GoOld.DailyAwardItem)
	file.WriteInt16(log.GoOld.NewRecordCount)
	file.WriteInt32(log.GoOld.DailyFame)
	file.WriteInt16(log.GoOld.BuyTimes)
}

func (log *PlayerArenaUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerArena.Update
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FailedCdTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BattleTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.WinTimes)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.DailyAwardItem)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.NewRecordCount)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.DailyFame)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.BuyTimes)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerArenaUpdateLog) CommitToTLog() {
}

func (log *PlayerArenaUpdateLog) CommitToXdLog() {
}

func (log *PlayerArenaUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerArenaInsertLog) Rollback() {
	if log.db.tables.PlayerArena != log.Row {
		panic("insert rollback check failed 'PlayerArena'")
	}
	log.db.tables.PlayerArena = nil
	C.Free_PlayerArena(log.Row)
}

func (log *PlayerArenaDeleteLog) Rollback() {
	if log.db.tables.PlayerArena != nil {
		panic("delete rollback check failed 'player_arena'")
	}
	log.db.tables.PlayerArena = log.Old
}

func (log *PlayerArenaUpdateLog) Rollback() {
	if log.db.tables.PlayerArena != log.Row {
		panic("update rollback check failed 'player_arena'")
	}
	log.db.tables.PlayerArena = log.Old
	C.Free_PlayerArena(log.Row)
}

/* ========== player_arena_rank ========== */

type PlayerArenaRankInsertLog struct {
	db *Database 
	Row *C.PlayerArenaRank
	GoRow *PlayerArenaRank
}

func (db *Database) newPlayerArenaRankInsertLog(row *C.PlayerArenaRank, goRow *PlayerArenaRank) *PlayerArenaRankInsertLog {
	return &PlayerArenaRankInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerArenaRankInsertLog) Free() {
}

func (log *PlayerArenaRankInsertLog) InvokeHook() {
	if g_Hooks.PlayerArenaRank != nil {
		g_Hooks.PlayerArenaRank.Insert(&PlayerArenaRankRow{row: log.Row})
	}
}

func (log *PlayerArenaRankInsertLog) SetEventId(time.Time) {
}

func (log *PlayerArenaRankInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(19)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.Rank)
	file.WriteInt32(log.GoRow.Rank1)
	file.WriteInt32(log.GoRow.Rank2)
	file.WriteInt32(log.GoRow.Rank3)
	file.WriteInt64(log.GoRow.Time)
}

func (log *PlayerArenaRankInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerArenaRank.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Rank)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Rank1)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Rank2)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Rank3)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Time)))
	return stmt.Execute()
}

func (log *PlayerArenaRankInsertLog) CommitToTLog() {
}

func (log *PlayerArenaRankInsertLog) CommitToXdLog() {
}

func (log *PlayerArenaRankInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerArenaRankDeleteLog struct {
	db *Database
	Old *C.PlayerArenaRank
	GoOld *PlayerArenaRank
}

func (db *Database) newPlayerArenaRankDeleteLog(old *C.PlayerArenaRank, goOld *PlayerArenaRank) *PlayerArenaRankDeleteLog {
	return &PlayerArenaRankDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerArenaRankDeleteLog) Free() {
	C.Free_PlayerArenaRank(log.Old)
}

func (log *PlayerArenaRankDeleteLog) InvokeHook() {
	if g_Hooks.PlayerArenaRank != nil {
		g_Hooks.PlayerArenaRank.Delete(&PlayerArenaRankRow{row: log.Old})
	}
}

func (log *PlayerArenaRankDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerArenaRankDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(19)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Rank)
	file.WriteInt32(log.GoOld.Rank1)
	file.WriteInt32(log.GoOld.Rank2)
	file.WriteInt32(log.GoOld.Rank3)
	file.WriteInt64(log.GoOld.Time)
}

func (log *PlayerArenaRankDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerArenaRank.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerArenaRankDeleteLog) CommitToTLog() {
}

func (log *PlayerArenaRankDeleteLog) CommitToXdLog() {
}

func (log *PlayerArenaRankDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerArenaRankUpdateLog struct {
	db *Database 
	Row *C.PlayerArenaRank
	Old *C.PlayerArenaRank
	GoNew *PlayerArenaRank
	GoOld *PlayerArenaRank
}

func (db *Database) newPlayerArenaRankUpdateLog(row, old *C.PlayerArenaRank, goNew, goOld *PlayerArenaRank) *PlayerArenaRankUpdateLog {
	return &PlayerArenaRankUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerArenaRankUpdateLog) Free() {
	C.Free_PlayerArenaRank(log.Old)
}

func (log *PlayerArenaRankUpdateLog) InvokeHook() {
	if g_Hooks.PlayerArenaRank != nil {
		g_Hooks.PlayerArenaRank.Update(&PlayerArenaRankRow{row: log.Row}, &PlayerArenaRankRow{row: log.Old})
	}
}

func (log *PlayerArenaRankUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerArenaRankUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(19)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.Rank)
	file.WriteInt32(log.GoNew.Rank1)
	file.WriteInt32(log.GoNew.Rank2)
	file.WriteInt32(log.GoNew.Rank3)
	file.WriteInt64(log.GoNew.Time)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Rank)
	file.WriteInt32(log.GoOld.Rank1)
	file.WriteInt32(log.GoOld.Rank2)
	file.WriteInt32(log.GoOld.Rank3)
	file.WriteInt64(log.GoOld.Time)
}

func (log *PlayerArenaRankUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerArenaRank.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.Rank)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Rank1)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Rank2)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Rank3)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Time)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerArenaRankUpdateLog) CommitToTLog() {
}

func (log *PlayerArenaRankUpdateLog) CommitToXdLog() {
}

func (log *PlayerArenaRankUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerArenaRankInsertLog) Rollback() {
	if log.db.tables.PlayerArenaRank != log.Row {
		panic("insert rollback check failed 'PlayerArenaRank'")
	}
	log.db.tables.PlayerArenaRank = nil
	C.Free_PlayerArenaRank(log.Row)
}

func (log *PlayerArenaRankDeleteLog) Rollback() {
	if log.db.tables.PlayerArenaRank != nil {
		panic("delete rollback check failed 'player_arena_rank'")
	}
	log.db.tables.PlayerArenaRank = log.Old
}

func (log *PlayerArenaRankUpdateLog) Rollback() {
	if log.db.tables.PlayerArenaRank != log.Row {
		panic("update rollback check failed 'player_arena_rank'")
	}
	log.db.tables.PlayerArenaRank = log.Old
	C.Free_PlayerArenaRank(log.Row)
}

/* ========== player_arena_record ========== */

type PlayerArenaRecordInsertLog struct {
	db *Database 
	Row *C.PlayerArenaRecord
	GoRow *PlayerArenaRecord
}

func (db *Database) newPlayerArenaRecordInsertLog(row *C.PlayerArenaRecord, goRow *PlayerArenaRecord) *PlayerArenaRecordInsertLog {
	return &PlayerArenaRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerArenaRecordInsertLog) Free() {
}

func (log *PlayerArenaRecordInsertLog) InvokeHook() {
	if g_Hooks.PlayerArenaRecord != nil {
		g_Hooks.PlayerArenaRecord.Insert(&PlayerArenaRecordRow{row: log.Row})
	}
}

func (log *PlayerArenaRecordInsertLog) SetEventId(time.Time) {
}

func (log *PlayerArenaRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(20)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.Mode)
	file.WriteInt32(log.GoRow.OldRank)
	file.WriteInt32(log.GoRow.NewRank)
	file.WriteInt32(log.GoRow.FightNum)
	file.WriteInt64(log.GoRow.TargetPid)
	file.WriteString(log.GoRow.TargetNick)
	file.WriteInt32(log.GoRow.TargetOldRank)
	file.WriteInt32(log.GoRow.TargetNewRank)
	file.WriteInt32(log.GoRow.TargetFightNum)
	file.WriteInt64(log.GoRow.RecordTime)
}

func (log *PlayerArenaRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerArenaRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Mode)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.OldRank)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.NewRank)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.FightNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.TargetPid)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.TargetNick), int(log.Row.TargetNick_len))
	stmt.BindInt(unsafe.Pointer(&(log.Row.TargetOldRank)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.TargetNewRank)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.TargetFightNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.RecordTime)))
	return stmt.Execute()
}

func (log *PlayerArenaRecordInsertLog) CommitToTLog() {
}

func (log *PlayerArenaRecordInsertLog) CommitToXdLog() {
}

func (log *PlayerArenaRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerArenaRecordDeleteLog struct {
	db *Database
	Old *C.PlayerArenaRecord
	GoOld *PlayerArenaRecord
}

func (db *Database) newPlayerArenaRecordDeleteLog(old *C.PlayerArenaRecord, goOld *PlayerArenaRecord) *PlayerArenaRecordDeleteLog {
	return &PlayerArenaRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerArenaRecordDeleteLog) Free() {
	C.Free_PlayerArenaRecord(log.Old)
}

func (log *PlayerArenaRecordDeleteLog) InvokeHook() {
	if g_Hooks.PlayerArenaRecord != nil {
		g_Hooks.PlayerArenaRecord.Delete(&PlayerArenaRecordRow{row: log.Old})
	}
}

func (log *PlayerArenaRecordDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerArenaRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(20)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.Mode)
	file.WriteInt32(log.GoOld.OldRank)
	file.WriteInt32(log.GoOld.NewRank)
	file.WriteInt32(log.GoOld.FightNum)
	file.WriteInt64(log.GoOld.TargetPid)
	file.WriteString(log.GoOld.TargetNick)
	file.WriteInt32(log.GoOld.TargetOldRank)
	file.WriteInt32(log.GoOld.TargetNewRank)
	file.WriteInt32(log.GoOld.TargetFightNum)
	file.WriteInt64(log.GoOld.RecordTime)
}

func (log *PlayerArenaRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerArenaRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerArenaRecordDeleteLog) CommitToTLog() {
}

func (log *PlayerArenaRecordDeleteLog) CommitToXdLog() {
}

func (log *PlayerArenaRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerArenaRecordUpdateLog struct {
	db *Database 
	Row *C.PlayerArenaRecord
	Old *C.PlayerArenaRecord
	GoNew *PlayerArenaRecord
	GoOld *PlayerArenaRecord
}

func (db *Database) newPlayerArenaRecordUpdateLog(row, old *C.PlayerArenaRecord, goNew, goOld *PlayerArenaRecord) *PlayerArenaRecordUpdateLog {
	return &PlayerArenaRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerArenaRecordUpdateLog) Free() {
	C.Free_PlayerArenaRecord(log.Old)
}

func (log *PlayerArenaRecordUpdateLog) InvokeHook() {
	if g_Hooks.PlayerArenaRecord != nil {
		g_Hooks.PlayerArenaRecord.Update(&PlayerArenaRecordRow{row: log.Row}, &PlayerArenaRecordRow{row: log.Old})
	}
}

func (log *PlayerArenaRecordUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerArenaRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(20)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.Mode)
	file.WriteInt32(log.GoNew.OldRank)
	file.WriteInt32(log.GoNew.NewRank)
	file.WriteInt32(log.GoNew.FightNum)
	file.WriteInt64(log.GoNew.TargetPid)
	file.WriteString(log.GoNew.TargetNick)
	file.WriteInt32(log.GoNew.TargetOldRank)
	file.WriteInt32(log.GoNew.TargetNewRank)
	file.WriteInt32(log.GoNew.TargetFightNum)
	file.WriteInt64(log.GoNew.RecordTime)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.Mode)
	file.WriteInt32(log.GoOld.OldRank)
	file.WriteInt32(log.GoOld.NewRank)
	file.WriteInt32(log.GoOld.FightNum)
	file.WriteInt64(log.GoOld.TargetPid)
	file.WriteString(log.GoOld.TargetNick)
	file.WriteInt32(log.GoOld.TargetOldRank)
	file.WriteInt32(log.GoOld.TargetNewRank)
	file.WriteInt32(log.GoOld.TargetFightNum)
	file.WriteInt64(log.GoOld.RecordTime)
}

func (log *PlayerArenaRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerArenaRecord.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Mode)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.OldRank)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.NewRank)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.FightNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.TargetPid)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.TargetNick), int(log.Row.TargetNick_len))
	stmt.BindInt(unsafe.Pointer(&(log.Row.TargetOldRank)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.TargetNewRank)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.TargetFightNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.RecordTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerArenaRecordUpdateLog) CommitToTLog() {
}

func (log *PlayerArenaRecordUpdateLog) CommitToXdLog() {
}

func (log *PlayerArenaRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerArenaRecordInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerArenaRecord
	for crow := log.db.tables.PlayerArenaRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerArenaRecord'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerArenaRecord = log.db.tables.PlayerArenaRecord.next
	}
	C.Free_PlayerArenaRecord(log.Row)
}

func (log *PlayerArenaRecordDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerArenaRecord
	for crow := log.db.tables.PlayerArenaRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_arena_record'")
	}
	log.Old.next = log.db.tables.PlayerArenaRecord
	log.db.tables.PlayerArenaRecord = log.Old
}

func (log *PlayerArenaRecordUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerArenaRecord
	for crow := log.db.tables.PlayerArenaRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_arena_record'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerArenaRecord = log.Old
	}
	C.Free_PlayerArenaRecord(log.Row)
}

/* ========== player_awaken_graphic ========== */

type PlayerAwakenGraphicInsertLog struct {
	db *Database 
	Row *C.PlayerAwakenGraphic
	GoRow *PlayerAwakenGraphic
}

func (db *Database) newPlayerAwakenGraphicInsertLog(row *C.PlayerAwakenGraphic, goRow *PlayerAwakenGraphic) *PlayerAwakenGraphicInsertLog {
	return &PlayerAwakenGraphicInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerAwakenGraphicInsertLog) Free() {
}

func (log *PlayerAwakenGraphicInsertLog) InvokeHook() {
	if g_Hooks.PlayerAwakenGraphic != nil {
		g_Hooks.PlayerAwakenGraphic.Insert(&PlayerAwakenGraphicRow{row: log.Row})
	}
}

func (log *PlayerAwakenGraphicInsertLog) SetEventId(time.Time) {
}

func (log *PlayerAwakenGraphicInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(21)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.RoleId)
	file.WriteInt16(log.GoRow.AttrImpl)
	file.WriteInt8(log.GoRow.Level)
}

func (log *PlayerAwakenGraphicInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerAwakenGraphic.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RoleId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AttrImpl)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Level)))
	return stmt.Execute()
}

func (log *PlayerAwakenGraphicInsertLog) CommitToTLog() {
}

func (log *PlayerAwakenGraphicInsertLog) CommitToXdLog() {
}

func (log *PlayerAwakenGraphicInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerAwakenGraphicDeleteLog struct {
	db *Database
	Old *C.PlayerAwakenGraphic
	GoOld *PlayerAwakenGraphic
}

func (db *Database) newPlayerAwakenGraphicDeleteLog(old *C.PlayerAwakenGraphic, goOld *PlayerAwakenGraphic) *PlayerAwakenGraphicDeleteLog {
	return &PlayerAwakenGraphicDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerAwakenGraphicDeleteLog) Free() {
	C.Free_PlayerAwakenGraphic(log.Old)
}

func (log *PlayerAwakenGraphicDeleteLog) InvokeHook() {
	if g_Hooks.PlayerAwakenGraphic != nil {
		g_Hooks.PlayerAwakenGraphic.Delete(&PlayerAwakenGraphicRow{row: log.Old})
	}
}

func (log *PlayerAwakenGraphicDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerAwakenGraphicDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(21)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.RoleId)
	file.WriteInt16(log.GoOld.AttrImpl)
	file.WriteInt8(log.GoOld.Level)
}

func (log *PlayerAwakenGraphicDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerAwakenGraphic.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerAwakenGraphicDeleteLog) CommitToTLog() {
}

func (log *PlayerAwakenGraphicDeleteLog) CommitToXdLog() {
}

func (log *PlayerAwakenGraphicDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerAwakenGraphicUpdateLog struct {
	db *Database 
	Row *C.PlayerAwakenGraphic
	Old *C.PlayerAwakenGraphic
	GoNew *PlayerAwakenGraphic
	GoOld *PlayerAwakenGraphic
}

func (db *Database) newPlayerAwakenGraphicUpdateLog(row, old *C.PlayerAwakenGraphic, goNew, goOld *PlayerAwakenGraphic) *PlayerAwakenGraphicUpdateLog {
	return &PlayerAwakenGraphicUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerAwakenGraphicUpdateLog) Free() {
	C.Free_PlayerAwakenGraphic(log.Old)
}

func (log *PlayerAwakenGraphicUpdateLog) InvokeHook() {
	if g_Hooks.PlayerAwakenGraphic != nil {
		g_Hooks.PlayerAwakenGraphic.Update(&PlayerAwakenGraphicRow{row: log.Row}, &PlayerAwakenGraphicRow{row: log.Old})
	}
}

func (log *PlayerAwakenGraphicUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerAwakenGraphicUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(21)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.RoleId)
	file.WriteInt16(log.GoNew.AttrImpl)
	file.WriteInt8(log.GoNew.Level)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.RoleId)
	file.WriteInt16(log.GoOld.AttrImpl)
	file.WriteInt8(log.GoOld.Level)
}

func (log *PlayerAwakenGraphicUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerAwakenGraphic.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RoleId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AttrImpl)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerAwakenGraphicUpdateLog) CommitToTLog() {
}

func (log *PlayerAwakenGraphicUpdateLog) CommitToXdLog() {
}

func (log *PlayerAwakenGraphicUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerAwakenGraphicInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerAwakenGraphic
	for crow := log.db.tables.PlayerAwakenGraphic; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerAwakenGraphic'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerAwakenGraphic = log.db.tables.PlayerAwakenGraphic.next
	}
	C.Free_PlayerAwakenGraphic(log.Row)
}

func (log *PlayerAwakenGraphicDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerAwakenGraphic
	for crow := log.db.tables.PlayerAwakenGraphic; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_awaken_graphic'")
	}
	log.Old.next = log.db.tables.PlayerAwakenGraphic
	log.db.tables.PlayerAwakenGraphic = log.Old
}

func (log *PlayerAwakenGraphicUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerAwakenGraphic
	for crow := log.db.tables.PlayerAwakenGraphic; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_awaken_graphic'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerAwakenGraphic = log.Old
	}
	C.Free_PlayerAwakenGraphic(log.Row)
}

/* ========== player_battle_pet ========== */

type PlayerBattlePetInsertLog struct {
	db *Database 
	Row *C.PlayerBattlePet
	GoRow *PlayerBattlePet
}

func (db *Database) newPlayerBattlePetInsertLog(row *C.PlayerBattlePet, goRow *PlayerBattlePet) *PlayerBattlePetInsertLog {
	return &PlayerBattlePetInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerBattlePetInsertLog) Free() {
}

func (log *PlayerBattlePetInsertLog) InvokeHook() {
	if g_Hooks.PlayerBattlePet != nil {
		g_Hooks.PlayerBattlePet.Insert(&PlayerBattlePetRow{row: log.Row})
	}
}

func (log *PlayerBattlePetInsertLog) SetEventId(time.Time) {
}

func (log *PlayerBattlePetInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(22)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.BattlePetId)
	file.WriteInt32(log.GoRow.Level)
	file.WriteInt64(log.GoRow.Exp)
	file.WriteInt32(log.GoRow.SkillLevel)
}

func (log *PlayerBattlePetInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerBattlePet.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.BattlePetId)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Exp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.SkillLevel)))
	return stmt.Execute()
}

func (log *PlayerBattlePetInsertLog) CommitToTLog() {
}

func (log *PlayerBattlePetInsertLog) CommitToXdLog() {
}

func (log *PlayerBattlePetInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerBattlePetDeleteLog struct {
	db *Database
	Old *C.PlayerBattlePet
	GoOld *PlayerBattlePet
}

func (db *Database) newPlayerBattlePetDeleteLog(old *C.PlayerBattlePet, goOld *PlayerBattlePet) *PlayerBattlePetDeleteLog {
	return &PlayerBattlePetDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerBattlePetDeleteLog) Free() {
	C.Free_PlayerBattlePet(log.Old)
}

func (log *PlayerBattlePetDeleteLog) InvokeHook() {
	if g_Hooks.PlayerBattlePet != nil {
		g_Hooks.PlayerBattlePet.Delete(&PlayerBattlePetRow{row: log.Old})
	}
}

func (log *PlayerBattlePetDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerBattlePetDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(22)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.BattlePetId)
	file.WriteInt32(log.GoOld.Level)
	file.WriteInt64(log.GoOld.Exp)
	file.WriteInt32(log.GoOld.SkillLevel)
}

func (log *PlayerBattlePetDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerBattlePet.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerBattlePetDeleteLog) CommitToTLog() {
}

func (log *PlayerBattlePetDeleteLog) CommitToXdLog() {
}

func (log *PlayerBattlePetDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerBattlePetUpdateLog struct {
	db *Database 
	Row *C.PlayerBattlePet
	Old *C.PlayerBattlePet
	GoNew *PlayerBattlePet
	GoOld *PlayerBattlePet
}

func (db *Database) newPlayerBattlePetUpdateLog(row, old *C.PlayerBattlePet, goNew, goOld *PlayerBattlePet) *PlayerBattlePetUpdateLog {
	return &PlayerBattlePetUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerBattlePetUpdateLog) Free() {
	C.Free_PlayerBattlePet(log.Old)
}

func (log *PlayerBattlePetUpdateLog) InvokeHook() {
	if g_Hooks.PlayerBattlePet != nil {
		g_Hooks.PlayerBattlePet.Update(&PlayerBattlePetRow{row: log.Row}, &PlayerBattlePetRow{row: log.Old})
	}
}

func (log *PlayerBattlePetUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerBattlePetUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(22)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.BattlePetId)
	file.WriteInt32(log.GoNew.Level)
	file.WriteInt64(log.GoNew.Exp)
	file.WriteInt32(log.GoNew.SkillLevel)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.BattlePetId)
	file.WriteInt32(log.GoOld.Level)
	file.WriteInt64(log.GoOld.Exp)
	file.WriteInt32(log.GoOld.SkillLevel)
}

func (log *PlayerBattlePetUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerBattlePet.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.BattlePetId)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Exp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.SkillLevel)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerBattlePetUpdateLog) CommitToTLog() {
}

func (log *PlayerBattlePetUpdateLog) CommitToXdLog() {
}

func (log *PlayerBattlePetUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerBattlePetInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerBattlePet
	for crow := log.db.tables.PlayerBattlePet; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerBattlePet'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerBattlePet = log.db.tables.PlayerBattlePet.next
	}
	C.Free_PlayerBattlePet(log.Row)
}

func (log *PlayerBattlePetDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerBattlePet
	for crow := log.db.tables.PlayerBattlePet; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_battle_pet'")
	}
	log.Old.next = log.db.tables.PlayerBattlePet
	log.db.tables.PlayerBattlePet = log.Old
}

func (log *PlayerBattlePetUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerBattlePet
	for crow := log.db.tables.PlayerBattlePet; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_battle_pet'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerBattlePet = log.Old
	}
	C.Free_PlayerBattlePet(log.Row)
}

/* ========== player_battle_pet_grid ========== */

type PlayerBattlePetGridInsertLog struct {
	db *Database 
	Row *C.PlayerBattlePetGrid
	GoRow *PlayerBattlePetGrid
}

func (db *Database) newPlayerBattlePetGridInsertLog(row *C.PlayerBattlePetGrid, goRow *PlayerBattlePetGrid) *PlayerBattlePetGridInsertLog {
	return &PlayerBattlePetGridInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerBattlePetGridInsertLog) Free() {
}

func (log *PlayerBattlePetGridInsertLog) InvokeHook() {
	if g_Hooks.PlayerBattlePetGrid != nil {
		g_Hooks.PlayerBattlePetGrid.Insert(&PlayerBattlePetGridRow{row: log.Row})
	}
}

func (log *PlayerBattlePetGridInsertLog) SetEventId(time.Time) {
}

func (log *PlayerBattlePetGridInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(23)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.GridId)
	file.WriteInt32(log.GoRow.BattlePetId)
}

func (log *PlayerBattlePetGridInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerBattlePetGrid.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.GridId)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.BattlePetId)))
	return stmt.Execute()
}

func (log *PlayerBattlePetGridInsertLog) CommitToTLog() {
}

func (log *PlayerBattlePetGridInsertLog) CommitToXdLog() {
}

func (log *PlayerBattlePetGridInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerBattlePetGridDeleteLog struct {
	db *Database
	Old *C.PlayerBattlePetGrid
	GoOld *PlayerBattlePetGrid
}

func (db *Database) newPlayerBattlePetGridDeleteLog(old *C.PlayerBattlePetGrid, goOld *PlayerBattlePetGrid) *PlayerBattlePetGridDeleteLog {
	return &PlayerBattlePetGridDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerBattlePetGridDeleteLog) Free() {
	C.Free_PlayerBattlePetGrid(log.Old)
}

func (log *PlayerBattlePetGridDeleteLog) InvokeHook() {
	if g_Hooks.PlayerBattlePetGrid != nil {
		g_Hooks.PlayerBattlePetGrid.Delete(&PlayerBattlePetGridRow{row: log.Old})
	}
}

func (log *PlayerBattlePetGridDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerBattlePetGridDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(23)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.GridId)
	file.WriteInt32(log.GoOld.BattlePetId)
}

func (log *PlayerBattlePetGridDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerBattlePetGrid.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerBattlePetGridDeleteLog) CommitToTLog() {
}

func (log *PlayerBattlePetGridDeleteLog) CommitToXdLog() {
}

func (log *PlayerBattlePetGridDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerBattlePetGridUpdateLog struct {
	db *Database 
	Row *C.PlayerBattlePetGrid
	Old *C.PlayerBattlePetGrid
	GoNew *PlayerBattlePetGrid
	GoOld *PlayerBattlePetGrid
}

func (db *Database) newPlayerBattlePetGridUpdateLog(row, old *C.PlayerBattlePetGrid, goNew, goOld *PlayerBattlePetGrid) *PlayerBattlePetGridUpdateLog {
	return &PlayerBattlePetGridUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerBattlePetGridUpdateLog) Free() {
	C.Free_PlayerBattlePetGrid(log.Old)
}

func (log *PlayerBattlePetGridUpdateLog) InvokeHook() {
	if g_Hooks.PlayerBattlePetGrid != nil {
		g_Hooks.PlayerBattlePetGrid.Update(&PlayerBattlePetGridRow{row: log.Row}, &PlayerBattlePetGridRow{row: log.Old})
	}
}

func (log *PlayerBattlePetGridUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerBattlePetGridUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(23)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.GridId)
	file.WriteInt32(log.GoNew.BattlePetId)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.GridId)
	file.WriteInt32(log.GoOld.BattlePetId)
}

func (log *PlayerBattlePetGridUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerBattlePetGrid.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.GridId)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.BattlePetId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerBattlePetGridUpdateLog) CommitToTLog() {
}

func (log *PlayerBattlePetGridUpdateLog) CommitToXdLog() {
}

func (log *PlayerBattlePetGridUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerBattlePetGridInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerBattlePetGrid
	for crow := log.db.tables.PlayerBattlePetGrid; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerBattlePetGrid'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerBattlePetGrid = log.db.tables.PlayerBattlePetGrid.next
	}
	C.Free_PlayerBattlePetGrid(log.Row)
}

func (log *PlayerBattlePetGridDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerBattlePetGrid
	for crow := log.db.tables.PlayerBattlePetGrid; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_battle_pet_grid'")
	}
	log.Old.next = log.db.tables.PlayerBattlePetGrid
	log.db.tables.PlayerBattlePetGrid = log.Old
}

func (log *PlayerBattlePetGridUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerBattlePetGrid
	for crow := log.db.tables.PlayerBattlePetGrid; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_battle_pet_grid'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerBattlePetGrid = log.Old
	}
	C.Free_PlayerBattlePetGrid(log.Row)
}

/* ========== player_battle_pet_state ========== */

type PlayerBattlePetStateInsertLog struct {
	db *Database 
	Row *C.PlayerBattlePetState
	GoRow *PlayerBattlePetState
}

func (db *Database) newPlayerBattlePetStateInsertLog(row *C.PlayerBattlePetState, goRow *PlayerBattlePetState) *PlayerBattlePetStateInsertLog {
	return &PlayerBattlePetStateInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerBattlePetStateInsertLog) Free() {
}

func (log *PlayerBattlePetStateInsertLog) InvokeHook() {
	if g_Hooks.PlayerBattlePetState != nil {
		g_Hooks.PlayerBattlePetState.Insert(&PlayerBattlePetStateRow{row: log.Row})
	}
}

func (log *PlayerBattlePetStateInsertLog) SetEventId(time.Time) {
}

func (log *PlayerBattlePetStateInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(24)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.UpgradeByIngotNum)
	file.WriteInt64(log.GoRow.UpgradeByIngotTime)
}

func (log *PlayerBattlePetStateInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerBattlePetState.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.UpgradeByIngotNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.UpgradeByIngotTime)))
	return stmt.Execute()
}

func (log *PlayerBattlePetStateInsertLog) CommitToTLog() {
}

func (log *PlayerBattlePetStateInsertLog) CommitToXdLog() {
}

func (log *PlayerBattlePetStateInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerBattlePetStateDeleteLog struct {
	db *Database
	Old *C.PlayerBattlePetState
	GoOld *PlayerBattlePetState
}

func (db *Database) newPlayerBattlePetStateDeleteLog(old *C.PlayerBattlePetState, goOld *PlayerBattlePetState) *PlayerBattlePetStateDeleteLog {
	return &PlayerBattlePetStateDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerBattlePetStateDeleteLog) Free() {
	C.Free_PlayerBattlePetState(log.Old)
}

func (log *PlayerBattlePetStateDeleteLog) InvokeHook() {
	if g_Hooks.PlayerBattlePetState != nil {
		g_Hooks.PlayerBattlePetState.Delete(&PlayerBattlePetStateRow{row: log.Old})
	}
}

func (log *PlayerBattlePetStateDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerBattlePetStateDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(24)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.UpgradeByIngotNum)
	file.WriteInt64(log.GoOld.UpgradeByIngotTime)
}

func (log *PlayerBattlePetStateDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerBattlePetState.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerBattlePetStateDeleteLog) CommitToTLog() {
}

func (log *PlayerBattlePetStateDeleteLog) CommitToXdLog() {
}

func (log *PlayerBattlePetStateDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerBattlePetStateUpdateLog struct {
	db *Database 
	Row *C.PlayerBattlePetState
	Old *C.PlayerBattlePetState
	GoNew *PlayerBattlePetState
	GoOld *PlayerBattlePetState
}

func (db *Database) newPlayerBattlePetStateUpdateLog(row, old *C.PlayerBattlePetState, goNew, goOld *PlayerBattlePetState) *PlayerBattlePetStateUpdateLog {
	return &PlayerBattlePetStateUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerBattlePetStateUpdateLog) Free() {
	C.Free_PlayerBattlePetState(log.Old)
}

func (log *PlayerBattlePetStateUpdateLog) InvokeHook() {
	if g_Hooks.PlayerBattlePetState != nil {
		g_Hooks.PlayerBattlePetState.Update(&PlayerBattlePetStateRow{row: log.Row}, &PlayerBattlePetStateRow{row: log.Old})
	}
}

func (log *PlayerBattlePetStateUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerBattlePetStateUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(24)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.UpgradeByIngotNum)
	file.WriteInt64(log.GoNew.UpgradeByIngotTime)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.UpgradeByIngotNum)
	file.WriteInt64(log.GoOld.UpgradeByIngotTime)
}

func (log *PlayerBattlePetStateUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerBattlePetState.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.UpgradeByIngotNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.UpgradeByIngotTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerBattlePetStateUpdateLog) CommitToTLog() {
}

func (log *PlayerBattlePetStateUpdateLog) CommitToXdLog() {
}

func (log *PlayerBattlePetStateUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerBattlePetStateInsertLog) Rollback() {
	if log.db.tables.PlayerBattlePetState != log.Row {
		panic("insert rollback check failed 'PlayerBattlePetState'")
	}
	log.db.tables.PlayerBattlePetState = nil
	C.Free_PlayerBattlePetState(log.Row)
}

func (log *PlayerBattlePetStateDeleteLog) Rollback() {
	if log.db.tables.PlayerBattlePetState != nil {
		panic("delete rollback check failed 'player_battle_pet_state'")
	}
	log.db.tables.PlayerBattlePetState = log.Old
}

func (log *PlayerBattlePetStateUpdateLog) Rollback() {
	if log.db.tables.PlayerBattlePetState != log.Row {
		panic("update rollback check failed 'player_battle_pet_state'")
	}
	log.db.tables.PlayerBattlePetState = log.Old
	C.Free_PlayerBattlePetState(log.Row)
}

/* ========== player_chest_state ========== */

type PlayerChestStateInsertLog struct {
	db *Database 
	Row *C.PlayerChestState
	GoRow *PlayerChestState
}

func (db *Database) newPlayerChestStateInsertLog(row *C.PlayerChestState, goRow *PlayerChestState) *PlayerChestStateInsertLog {
	return &PlayerChestStateInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerChestStateInsertLog) Free() {
}

func (log *PlayerChestStateInsertLog) InvokeHook() {
	if g_Hooks.PlayerChestState != nil {
		g_Hooks.PlayerChestState.Insert(&PlayerChestStateRow{row: log.Row})
	}
}

func (log *PlayerChestStateInsertLog) SetEventId(time.Time) {
}

func (log *PlayerChestStateInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(25)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.FreeCoinChestNum)
	file.WriteInt64(log.GoRow.LastFreeCoinChestAt)
	file.WriteInt32(log.GoRow.CoinChestNum)
	file.WriteInt32(log.GoRow.CoinChestTenNum)
	file.WriteInt8(log.GoRow.IsFirstCoinTen)
	file.WriteInt64(log.GoRow.CoinChestConsume)
	file.WriteInt64(log.GoRow.LastCoinChestAt)
	file.WriteInt64(log.GoRow.LastFreeIngotChestAt)
	file.WriteInt32(log.GoRow.IngotChestNum)
	file.WriteInt32(log.GoRow.IngotChestTenNum)
	file.WriteInt8(log.GoRow.IsFirstIngotTen)
	file.WriteInt64(log.GoRow.IngotChestConsume)
	file.WriteInt64(log.GoRow.LastIngotChestAt)
	file.WriteInt64(log.GoRow.LastFreePetChestAt)
}

func (log *PlayerChestStateInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerChestState.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.FreeCoinChestNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastFreeCoinChestAt)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.CoinChestNum)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.CoinChestTenNum)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.IsFirstCoinTen)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.CoinChestConsume)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastCoinChestAt)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastFreeIngotChestAt)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.IngotChestNum)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.IngotChestTenNum)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.IsFirstIngotTen)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.IngotChestConsume)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastIngotChestAt)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastFreePetChestAt)))
	return stmt.Execute()
}

func (log *PlayerChestStateInsertLog) CommitToTLog() {
}

func (log *PlayerChestStateInsertLog) CommitToXdLog() {
}

func (log *PlayerChestStateInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerChestStateDeleteLog struct {
	db *Database
	Old *C.PlayerChestState
	GoOld *PlayerChestState
}

func (db *Database) newPlayerChestStateDeleteLog(old *C.PlayerChestState, goOld *PlayerChestState) *PlayerChestStateDeleteLog {
	return &PlayerChestStateDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerChestStateDeleteLog) Free() {
	C.Free_PlayerChestState(log.Old)
}

func (log *PlayerChestStateDeleteLog) InvokeHook() {
	if g_Hooks.PlayerChestState != nil {
		g_Hooks.PlayerChestState.Delete(&PlayerChestStateRow{row: log.Old})
	}
}

func (log *PlayerChestStateDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerChestStateDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(25)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.FreeCoinChestNum)
	file.WriteInt64(log.GoOld.LastFreeCoinChestAt)
	file.WriteInt32(log.GoOld.CoinChestNum)
	file.WriteInt32(log.GoOld.CoinChestTenNum)
	file.WriteInt8(log.GoOld.IsFirstCoinTen)
	file.WriteInt64(log.GoOld.CoinChestConsume)
	file.WriteInt64(log.GoOld.LastCoinChestAt)
	file.WriteInt64(log.GoOld.LastFreeIngotChestAt)
	file.WriteInt32(log.GoOld.IngotChestNum)
	file.WriteInt32(log.GoOld.IngotChestTenNum)
	file.WriteInt8(log.GoOld.IsFirstIngotTen)
	file.WriteInt64(log.GoOld.IngotChestConsume)
	file.WriteInt64(log.GoOld.LastIngotChestAt)
	file.WriteInt64(log.GoOld.LastFreePetChestAt)
}

func (log *PlayerChestStateDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerChestState.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerChestStateDeleteLog) CommitToTLog() {
}

func (log *PlayerChestStateDeleteLog) CommitToXdLog() {
}

func (log *PlayerChestStateDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerChestStateUpdateLog struct {
	db *Database 
	Row *C.PlayerChestState
	Old *C.PlayerChestState
	GoNew *PlayerChestState
	GoOld *PlayerChestState
}

func (db *Database) newPlayerChestStateUpdateLog(row, old *C.PlayerChestState, goNew, goOld *PlayerChestState) *PlayerChestStateUpdateLog {
	return &PlayerChestStateUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerChestStateUpdateLog) Free() {
	C.Free_PlayerChestState(log.Old)
}

func (log *PlayerChestStateUpdateLog) InvokeHook() {
	if g_Hooks.PlayerChestState != nil {
		g_Hooks.PlayerChestState.Update(&PlayerChestStateRow{row: log.Row}, &PlayerChestStateRow{row: log.Old})
	}
}

func (log *PlayerChestStateUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerChestStateUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(25)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.FreeCoinChestNum)
	file.WriteInt64(log.GoNew.LastFreeCoinChestAt)
	file.WriteInt32(log.GoNew.CoinChestNum)
	file.WriteInt32(log.GoNew.CoinChestTenNum)
	file.WriteInt8(log.GoNew.IsFirstCoinTen)
	file.WriteInt64(log.GoNew.CoinChestConsume)
	file.WriteInt64(log.GoNew.LastCoinChestAt)
	file.WriteInt64(log.GoNew.LastFreeIngotChestAt)
	file.WriteInt32(log.GoNew.IngotChestNum)
	file.WriteInt32(log.GoNew.IngotChestTenNum)
	file.WriteInt8(log.GoNew.IsFirstIngotTen)
	file.WriteInt64(log.GoNew.IngotChestConsume)
	file.WriteInt64(log.GoNew.LastIngotChestAt)
	file.WriteInt64(log.GoNew.LastFreePetChestAt)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.FreeCoinChestNum)
	file.WriteInt64(log.GoOld.LastFreeCoinChestAt)
	file.WriteInt32(log.GoOld.CoinChestNum)
	file.WriteInt32(log.GoOld.CoinChestTenNum)
	file.WriteInt8(log.GoOld.IsFirstCoinTen)
	file.WriteInt64(log.GoOld.CoinChestConsume)
	file.WriteInt64(log.GoOld.LastCoinChestAt)
	file.WriteInt64(log.GoOld.LastFreeIngotChestAt)
	file.WriteInt32(log.GoOld.IngotChestNum)
	file.WriteInt32(log.GoOld.IngotChestTenNum)
	file.WriteInt8(log.GoOld.IsFirstIngotTen)
	file.WriteInt64(log.GoOld.IngotChestConsume)
	file.WriteInt64(log.GoOld.LastIngotChestAt)
	file.WriteInt64(log.GoOld.LastFreePetChestAt)
}

func (log *PlayerChestStateUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerChestState.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.FreeCoinChestNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastFreeCoinChestAt)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.CoinChestNum)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.CoinChestTenNum)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.IsFirstCoinTen)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.CoinChestConsume)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastCoinChestAt)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastFreeIngotChestAt)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.IngotChestNum)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.IngotChestTenNum)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.IsFirstIngotTen)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.IngotChestConsume)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastIngotChestAt)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastFreePetChestAt)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerChestStateUpdateLog) CommitToTLog() {
}

func (log *PlayerChestStateUpdateLog) CommitToXdLog() {
}

func (log *PlayerChestStateUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerChestStateInsertLog) Rollback() {
	if log.db.tables.PlayerChestState != log.Row {
		panic("insert rollback check failed 'PlayerChestState'")
	}
	log.db.tables.PlayerChestState = nil
	C.Free_PlayerChestState(log.Row)
}

func (log *PlayerChestStateDeleteLog) Rollback() {
	if log.db.tables.PlayerChestState != nil {
		panic("delete rollback check failed 'player_chest_state'")
	}
	log.db.tables.PlayerChestState = log.Old
}

func (log *PlayerChestStateUpdateLog) Rollback() {
	if log.db.tables.PlayerChestState != log.Row {
		panic("update rollback check failed 'player_chest_state'")
	}
	log.db.tables.PlayerChestState = log.Old
	C.Free_PlayerChestState(log.Row)
}

/* ========== player_clique_kongfu_attrib ========== */

type PlayerCliqueKongfuAttribInsertLog struct {
	db *Database 
	Row *C.PlayerCliqueKongfuAttrib
	GoRow *PlayerCliqueKongfuAttrib
}

func (db *Database) newPlayerCliqueKongfuAttribInsertLog(row *C.PlayerCliqueKongfuAttrib, goRow *PlayerCliqueKongfuAttrib) *PlayerCliqueKongfuAttribInsertLog {
	return &PlayerCliqueKongfuAttribInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerCliqueKongfuAttribInsertLog) Free() {
}

func (log *PlayerCliqueKongfuAttribInsertLog) InvokeHook() {
	if g_Hooks.PlayerCliqueKongfuAttrib != nil {
		g_Hooks.PlayerCliqueKongfuAttrib.Insert(&PlayerCliqueKongfuAttribRow{row: log.Row})
	}
}

func (log *PlayerCliqueKongfuAttribInsertLog) SetEventId(time.Time) {
}

func (log *PlayerCliqueKongfuAttribInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(26)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.Health)
	file.WriteInt32(log.GoRow.Attack)
	file.WriteInt32(log.GoRow.Defence)
}

func (log *PlayerCliqueKongfuAttribInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerCliqueKongfuAttrib.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Health)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Attack)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Defence)))
	return stmt.Execute()
}

func (log *PlayerCliqueKongfuAttribInsertLog) CommitToTLog() {
}

func (log *PlayerCliqueKongfuAttribInsertLog) CommitToXdLog() {
}

func (log *PlayerCliqueKongfuAttribInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerCliqueKongfuAttribDeleteLog struct {
	db *Database
	Old *C.PlayerCliqueKongfuAttrib
	GoOld *PlayerCliqueKongfuAttrib
}

func (db *Database) newPlayerCliqueKongfuAttribDeleteLog(old *C.PlayerCliqueKongfuAttrib, goOld *PlayerCliqueKongfuAttrib) *PlayerCliqueKongfuAttribDeleteLog {
	return &PlayerCliqueKongfuAttribDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerCliqueKongfuAttribDeleteLog) Free() {
	C.Free_PlayerCliqueKongfuAttrib(log.Old)
}

func (log *PlayerCliqueKongfuAttribDeleteLog) InvokeHook() {
	if g_Hooks.PlayerCliqueKongfuAttrib != nil {
		g_Hooks.PlayerCliqueKongfuAttrib.Delete(&PlayerCliqueKongfuAttribRow{row: log.Old})
	}
}

func (log *PlayerCliqueKongfuAttribDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerCliqueKongfuAttribDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(26)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Health)
	file.WriteInt32(log.GoOld.Attack)
	file.WriteInt32(log.GoOld.Defence)
}

func (log *PlayerCliqueKongfuAttribDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerCliqueKongfuAttrib.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerCliqueKongfuAttribDeleteLog) CommitToTLog() {
}

func (log *PlayerCliqueKongfuAttribDeleteLog) CommitToXdLog() {
}

func (log *PlayerCliqueKongfuAttribDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerCliqueKongfuAttribUpdateLog struct {
	db *Database 
	Row *C.PlayerCliqueKongfuAttrib
	Old *C.PlayerCliqueKongfuAttrib
	GoNew *PlayerCliqueKongfuAttrib
	GoOld *PlayerCliqueKongfuAttrib
}

func (db *Database) newPlayerCliqueKongfuAttribUpdateLog(row, old *C.PlayerCliqueKongfuAttrib, goNew, goOld *PlayerCliqueKongfuAttrib) *PlayerCliqueKongfuAttribUpdateLog {
	return &PlayerCliqueKongfuAttribUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerCliqueKongfuAttribUpdateLog) Free() {
	C.Free_PlayerCliqueKongfuAttrib(log.Old)
}

func (log *PlayerCliqueKongfuAttribUpdateLog) InvokeHook() {
	if g_Hooks.PlayerCliqueKongfuAttrib != nil {
		g_Hooks.PlayerCliqueKongfuAttrib.Update(&PlayerCliqueKongfuAttribRow{row: log.Row}, &PlayerCliqueKongfuAttribRow{row: log.Old})
	}
}

func (log *PlayerCliqueKongfuAttribUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerCliqueKongfuAttribUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(26)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.Health)
	file.WriteInt32(log.GoNew.Attack)
	file.WriteInt32(log.GoNew.Defence)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Health)
	file.WriteInt32(log.GoOld.Attack)
	file.WriteInt32(log.GoOld.Defence)
}

func (log *PlayerCliqueKongfuAttribUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerCliqueKongfuAttrib.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.Health)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Attack)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Defence)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerCliqueKongfuAttribUpdateLog) CommitToTLog() {
}

func (log *PlayerCliqueKongfuAttribUpdateLog) CommitToXdLog() {
}

func (log *PlayerCliqueKongfuAttribUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerCliqueKongfuAttribInsertLog) Rollback() {
	if log.db.tables.PlayerCliqueKongfuAttrib != log.Row {
		panic("insert rollback check failed 'PlayerCliqueKongfuAttrib'")
	}
	log.db.tables.PlayerCliqueKongfuAttrib = nil
	C.Free_PlayerCliqueKongfuAttrib(log.Row)
}

func (log *PlayerCliqueKongfuAttribDeleteLog) Rollback() {
	if log.db.tables.PlayerCliqueKongfuAttrib != nil {
		panic("delete rollback check failed 'player_clique_kongfu_attrib'")
	}
	log.db.tables.PlayerCliqueKongfuAttrib = log.Old
}

func (log *PlayerCliqueKongfuAttribUpdateLog) Rollback() {
	if log.db.tables.PlayerCliqueKongfuAttrib != log.Row {
		panic("update rollback check failed 'player_clique_kongfu_attrib'")
	}
	log.db.tables.PlayerCliqueKongfuAttrib = log.Old
	C.Free_PlayerCliqueKongfuAttrib(log.Row)
}

/* ========== player_coins ========== */

type PlayerCoinsInsertLog struct {
	db *Database 
	Row *C.PlayerCoins
	GoRow *PlayerCoins
}

func (db *Database) newPlayerCoinsInsertLog(row *C.PlayerCoins, goRow *PlayerCoins) *PlayerCoinsInsertLog {
	return &PlayerCoinsInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerCoinsInsertLog) Free() {
}

func (log *PlayerCoinsInsertLog) InvokeHook() {
	if g_Hooks.PlayerCoins != nil {
		g_Hooks.PlayerCoins.Insert(&PlayerCoinsRow{row: log.Row})
	}
}

func (log *PlayerCoinsInsertLog) SetEventId(time.Time) {
}

func (log *PlayerCoinsInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(27)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.BuyTime)
	file.WriteInt16(log.GoRow.DailyCount)
	file.WriteInt16(log.GoRow.BatchBought)
}

func (log *PlayerCoinsInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerCoins.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BuyTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyCount)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.BatchBought)))
	return stmt.Execute()
}

func (log *PlayerCoinsInsertLog) CommitToTLog() {
}

func (log *PlayerCoinsInsertLog) CommitToXdLog() {
}

func (log *PlayerCoinsInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerCoinsDeleteLog struct {
	db *Database
	Old *C.PlayerCoins
	GoOld *PlayerCoins
}

func (db *Database) newPlayerCoinsDeleteLog(old *C.PlayerCoins, goOld *PlayerCoins) *PlayerCoinsDeleteLog {
	return &PlayerCoinsDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerCoinsDeleteLog) Free() {
	C.Free_PlayerCoins(log.Old)
}

func (log *PlayerCoinsDeleteLog) InvokeHook() {
	if g_Hooks.PlayerCoins != nil {
		g_Hooks.PlayerCoins.Delete(&PlayerCoinsRow{row: log.Old})
	}
}

func (log *PlayerCoinsDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerCoinsDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(27)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.BuyTime)
	file.WriteInt16(log.GoOld.DailyCount)
	file.WriteInt16(log.GoOld.BatchBought)
}

func (log *PlayerCoinsDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerCoins.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerCoinsDeleteLog) CommitToTLog() {
}

func (log *PlayerCoinsDeleteLog) CommitToXdLog() {
}

func (log *PlayerCoinsDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerCoinsUpdateLog struct {
	db *Database 
	Row *C.PlayerCoins
	Old *C.PlayerCoins
	GoNew *PlayerCoins
	GoOld *PlayerCoins
}

func (db *Database) newPlayerCoinsUpdateLog(row, old *C.PlayerCoins, goNew, goOld *PlayerCoins) *PlayerCoinsUpdateLog {
	return &PlayerCoinsUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerCoinsUpdateLog) Free() {
	C.Free_PlayerCoins(log.Old)
}

func (log *PlayerCoinsUpdateLog) InvokeHook() {
	if g_Hooks.PlayerCoins != nil {
		g_Hooks.PlayerCoins.Update(&PlayerCoinsRow{row: log.Row}, &PlayerCoinsRow{row: log.Old})
	}
}

func (log *PlayerCoinsUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerCoinsUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(27)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.BuyTime)
	file.WriteInt16(log.GoNew.DailyCount)
	file.WriteInt16(log.GoNew.BatchBought)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.BuyTime)
	file.WriteInt16(log.GoOld.DailyCount)
	file.WriteInt16(log.GoOld.BatchBought)
}

func (log *PlayerCoinsUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerCoins.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BuyTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyCount)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.BatchBought)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerCoinsUpdateLog) CommitToTLog() {
}

func (log *PlayerCoinsUpdateLog) CommitToXdLog() {
}

func (log *PlayerCoinsUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerCoinsInsertLog) Rollback() {
	if log.db.tables.PlayerCoins != log.Row {
		panic("insert rollback check failed 'PlayerCoins'")
	}
	log.db.tables.PlayerCoins = nil
	C.Free_PlayerCoins(log.Row)
}

func (log *PlayerCoinsDeleteLog) Rollback() {
	if log.db.tables.PlayerCoins != nil {
		panic("delete rollback check failed 'player_coins'")
	}
	log.db.tables.PlayerCoins = log.Old
}

func (log *PlayerCoinsUpdateLog) Rollback() {
	if log.db.tables.PlayerCoins != log.Row {
		panic("update rollback check failed 'player_coins'")
	}
	log.db.tables.PlayerCoins = log.Old
	C.Free_PlayerCoins(log.Row)
}

/* ========== player_cornucopia ========== */

type PlayerCornucopiaInsertLog struct {
	db *Database 
	Row *C.PlayerCornucopia
	GoRow *PlayerCornucopia
}

func (db *Database) newPlayerCornucopiaInsertLog(row *C.PlayerCornucopia, goRow *PlayerCornucopia) *PlayerCornucopiaInsertLog {
	return &PlayerCornucopiaInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerCornucopiaInsertLog) Free() {
}

func (log *PlayerCornucopiaInsertLog) InvokeHook() {
	if g_Hooks.PlayerCornucopia != nil {
		g_Hooks.PlayerCornucopia.Insert(&PlayerCornucopiaRow{row: log.Row})
	}
}

func (log *PlayerCornucopiaInsertLog) SetEventId(time.Time) {
}

func (log *PlayerCornucopiaInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(28)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.OpenTime)
	file.WriteInt16(log.GoRow.DailyCount)
}

func (log *PlayerCornucopiaInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerCornucopia.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.OpenTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyCount)))
	return stmt.Execute()
}

func (log *PlayerCornucopiaInsertLog) CommitToTLog() {
}

func (log *PlayerCornucopiaInsertLog) CommitToXdLog() {
}

func (log *PlayerCornucopiaInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerCornucopiaDeleteLog struct {
	db *Database
	Old *C.PlayerCornucopia
	GoOld *PlayerCornucopia
}

func (db *Database) newPlayerCornucopiaDeleteLog(old *C.PlayerCornucopia, goOld *PlayerCornucopia) *PlayerCornucopiaDeleteLog {
	return &PlayerCornucopiaDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerCornucopiaDeleteLog) Free() {
	C.Free_PlayerCornucopia(log.Old)
}

func (log *PlayerCornucopiaDeleteLog) InvokeHook() {
	if g_Hooks.PlayerCornucopia != nil {
		g_Hooks.PlayerCornucopia.Delete(&PlayerCornucopiaRow{row: log.Old})
	}
}

func (log *PlayerCornucopiaDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerCornucopiaDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(28)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.OpenTime)
	file.WriteInt16(log.GoOld.DailyCount)
}

func (log *PlayerCornucopiaDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerCornucopia.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerCornucopiaDeleteLog) CommitToTLog() {
}

func (log *PlayerCornucopiaDeleteLog) CommitToXdLog() {
}

func (log *PlayerCornucopiaDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerCornucopiaUpdateLog struct {
	db *Database 
	Row *C.PlayerCornucopia
	Old *C.PlayerCornucopia
	GoNew *PlayerCornucopia
	GoOld *PlayerCornucopia
}

func (db *Database) newPlayerCornucopiaUpdateLog(row, old *C.PlayerCornucopia, goNew, goOld *PlayerCornucopia) *PlayerCornucopiaUpdateLog {
	return &PlayerCornucopiaUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerCornucopiaUpdateLog) Free() {
	C.Free_PlayerCornucopia(log.Old)
}

func (log *PlayerCornucopiaUpdateLog) InvokeHook() {
	if g_Hooks.PlayerCornucopia != nil {
		g_Hooks.PlayerCornucopia.Update(&PlayerCornucopiaRow{row: log.Row}, &PlayerCornucopiaRow{row: log.Old})
	}
}

func (log *PlayerCornucopiaUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerCornucopiaUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(28)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.OpenTime)
	file.WriteInt16(log.GoNew.DailyCount)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.OpenTime)
	file.WriteInt16(log.GoOld.DailyCount)
}

func (log *PlayerCornucopiaUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerCornucopia.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.OpenTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerCornucopiaUpdateLog) CommitToTLog() {
}

func (log *PlayerCornucopiaUpdateLog) CommitToXdLog() {
}

func (log *PlayerCornucopiaUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerCornucopiaInsertLog) Rollback() {
	if log.db.tables.PlayerCornucopia != log.Row {
		panic("insert rollback check failed 'PlayerCornucopia'")
	}
	log.db.tables.PlayerCornucopia = nil
	C.Free_PlayerCornucopia(log.Row)
}

func (log *PlayerCornucopiaDeleteLog) Rollback() {
	if log.db.tables.PlayerCornucopia != nil {
		panic("delete rollback check failed 'player_cornucopia'")
	}
	log.db.tables.PlayerCornucopia = log.Old
}

func (log *PlayerCornucopiaUpdateLog) Rollback() {
	if log.db.tables.PlayerCornucopia != log.Row {
		panic("update rollback check failed 'player_cornucopia'")
	}
	log.db.tables.PlayerCornucopia = log.Old
	C.Free_PlayerCornucopia(log.Row)
}

/* ========== player_daily_quest ========== */

type PlayerDailyQuestInsertLog struct {
	db *Database 
	Row *C.PlayerDailyQuest
	GoRow *PlayerDailyQuest
}

func (db *Database) newPlayerDailyQuestInsertLog(row *C.PlayerDailyQuest, goRow *PlayerDailyQuest) *PlayerDailyQuestInsertLog {
	return &PlayerDailyQuestInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerDailyQuestInsertLog) Free() {
}

func (log *PlayerDailyQuestInsertLog) InvokeHook() {
	if g_Hooks.PlayerDailyQuest != nil {
		g_Hooks.PlayerDailyQuest.Insert(&PlayerDailyQuestRow{row: log.Row})
	}
}

func (log *PlayerDailyQuestInsertLog) SetEventId(time.Time) {
}

func (log *PlayerDailyQuestInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(29)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.QuestId)
	file.WriteInt16(log.GoRow.FinishCount)
	file.WriteInt64(log.GoRow.LastUpdateTime)
	file.WriteInt8(log.GoRow.AwardStatus)
	file.WriteInt16(log.GoRow.Class)
}

func (log *PlayerDailyQuestInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDailyQuest.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.QuestId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.FinishCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastUpdateTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.AwardStatus)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Class)))
	return stmt.Execute()
}

func (log *PlayerDailyQuestInsertLog) CommitToTLog() {
}

func (log *PlayerDailyQuestInsertLog) CommitToXdLog() {
}

func (log *PlayerDailyQuestInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDailyQuestDeleteLog struct {
	db *Database
	Old *C.PlayerDailyQuest
	GoOld *PlayerDailyQuest
}

func (db *Database) newPlayerDailyQuestDeleteLog(old *C.PlayerDailyQuest, goOld *PlayerDailyQuest) *PlayerDailyQuestDeleteLog {
	return &PlayerDailyQuestDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerDailyQuestDeleteLog) Free() {
	C.Free_PlayerDailyQuest(log.Old)
}

func (log *PlayerDailyQuestDeleteLog) InvokeHook() {
	if g_Hooks.PlayerDailyQuest != nil {
		g_Hooks.PlayerDailyQuest.Delete(&PlayerDailyQuestRow{row: log.Old})
	}
}

func (log *PlayerDailyQuestDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerDailyQuestDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(29)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.QuestId)
	file.WriteInt16(log.GoOld.FinishCount)
	file.WriteInt64(log.GoOld.LastUpdateTime)
	file.WriteInt8(log.GoOld.AwardStatus)
	file.WriteInt16(log.GoOld.Class)
}

func (log *PlayerDailyQuestDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDailyQuest.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerDailyQuestDeleteLog) CommitToTLog() {
}

func (log *PlayerDailyQuestDeleteLog) CommitToXdLog() {
}

func (log *PlayerDailyQuestDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDailyQuestUpdateLog struct {
	db *Database 
	Row *C.PlayerDailyQuest
	Old *C.PlayerDailyQuest
	GoNew *PlayerDailyQuest
	GoOld *PlayerDailyQuest
}

func (db *Database) newPlayerDailyQuestUpdateLog(row, old *C.PlayerDailyQuest, goNew, goOld *PlayerDailyQuest) *PlayerDailyQuestUpdateLog {
	return &PlayerDailyQuestUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerDailyQuestUpdateLog) Free() {
	C.Free_PlayerDailyQuest(log.Old)
}

func (log *PlayerDailyQuestUpdateLog) InvokeHook() {
	if g_Hooks.PlayerDailyQuest != nil {
		g_Hooks.PlayerDailyQuest.Update(&PlayerDailyQuestRow{row: log.Row}, &PlayerDailyQuestRow{row: log.Old})
	}
}

func (log *PlayerDailyQuestUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerDailyQuestUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(29)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.QuestId)
	file.WriteInt16(log.GoNew.FinishCount)
	file.WriteInt64(log.GoNew.LastUpdateTime)
	file.WriteInt8(log.GoNew.AwardStatus)
	file.WriteInt16(log.GoNew.Class)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.QuestId)
	file.WriteInt16(log.GoOld.FinishCount)
	file.WriteInt64(log.GoOld.LastUpdateTime)
	file.WriteInt8(log.GoOld.AwardStatus)
	file.WriteInt16(log.GoOld.Class)
}

func (log *PlayerDailyQuestUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDailyQuest.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.QuestId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.FinishCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastUpdateTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.AwardStatus)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Class)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerDailyQuestUpdateLog) CommitToTLog() {
}

func (log *PlayerDailyQuestUpdateLog) CommitToXdLog() {
}

func (log *PlayerDailyQuestUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerDailyQuestInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerDailyQuest
	for crow := log.db.tables.PlayerDailyQuest; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerDailyQuest'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerDailyQuest = log.db.tables.PlayerDailyQuest.next
	}
	C.Free_PlayerDailyQuest(log.Row)
}

func (log *PlayerDailyQuestDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerDailyQuest
	for crow := log.db.tables.PlayerDailyQuest; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_daily_quest'")
	}
	log.Old.next = log.db.tables.PlayerDailyQuest
	log.db.tables.PlayerDailyQuest = log.Old
}

func (log *PlayerDailyQuestUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerDailyQuest
	for crow := log.db.tables.PlayerDailyQuest; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_daily_quest'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerDailyQuest = log.Old
	}
	C.Free_PlayerDailyQuest(log.Row)
}

/* ========== player_daily_quest_star_award_info ========== */

type PlayerDailyQuestStarAwardInfoInsertLog struct {
	db *Database 
	Row *C.PlayerDailyQuestStarAwardInfo
	GoRow *PlayerDailyQuestStarAwardInfo
}

func (db *Database) newPlayerDailyQuestStarAwardInfoInsertLog(row *C.PlayerDailyQuestStarAwardInfo, goRow *PlayerDailyQuestStarAwardInfo) *PlayerDailyQuestStarAwardInfoInsertLog {
	return &PlayerDailyQuestStarAwardInfoInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerDailyQuestStarAwardInfoInsertLog) Free() {
}

func (log *PlayerDailyQuestStarAwardInfoInsertLog) InvokeHook() {
	if g_Hooks.PlayerDailyQuestStarAwardInfo != nil {
		g_Hooks.PlayerDailyQuestStarAwardInfo.Insert(&PlayerDailyQuestStarAwardInfoRow{row: log.Row})
	}
}

func (log *PlayerDailyQuestStarAwardInfoInsertLog) SetEventId(time.Time) {
}

func (log *PlayerDailyQuestStarAwardInfoInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(30)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.Stars)
	file.WriteInt64(log.GoRow.Lastupdatetime)
	file.WriteString(log.GoRow.Awarded)
}

func (log *PlayerDailyQuestStarAwardInfoInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDailyQuestStarAwardInfo.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Stars)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Lastupdatetime)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Awarded), int(log.Row.Awarded_len))
	return stmt.Execute()
}

func (log *PlayerDailyQuestStarAwardInfoInsertLog) CommitToTLog() {
}

func (log *PlayerDailyQuestStarAwardInfoInsertLog) CommitToXdLog() {
}

func (log *PlayerDailyQuestStarAwardInfoInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDailyQuestStarAwardInfoDeleteLog struct {
	db *Database
	Old *C.PlayerDailyQuestStarAwardInfo
	GoOld *PlayerDailyQuestStarAwardInfo
}

func (db *Database) newPlayerDailyQuestStarAwardInfoDeleteLog(old *C.PlayerDailyQuestStarAwardInfo, goOld *PlayerDailyQuestStarAwardInfo) *PlayerDailyQuestStarAwardInfoDeleteLog {
	return &PlayerDailyQuestStarAwardInfoDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerDailyQuestStarAwardInfoDeleteLog) Free() {
	C.Free_PlayerDailyQuestStarAwardInfo(log.Old)
}

func (log *PlayerDailyQuestStarAwardInfoDeleteLog) InvokeHook() {
	if g_Hooks.PlayerDailyQuestStarAwardInfo != nil {
		g_Hooks.PlayerDailyQuestStarAwardInfo.Delete(&PlayerDailyQuestStarAwardInfoRow{row: log.Old})
	}
}

func (log *PlayerDailyQuestStarAwardInfoDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerDailyQuestStarAwardInfoDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(30)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Stars)
	file.WriteInt64(log.GoOld.Lastupdatetime)
	file.WriteString(log.GoOld.Awarded)
}

func (log *PlayerDailyQuestStarAwardInfoDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDailyQuestStarAwardInfo.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerDailyQuestStarAwardInfoDeleteLog) CommitToTLog() {
}

func (log *PlayerDailyQuestStarAwardInfoDeleteLog) CommitToXdLog() {
}

func (log *PlayerDailyQuestStarAwardInfoDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDailyQuestStarAwardInfoUpdateLog struct {
	db *Database 
	Row *C.PlayerDailyQuestStarAwardInfo
	Old *C.PlayerDailyQuestStarAwardInfo
	GoNew *PlayerDailyQuestStarAwardInfo
	GoOld *PlayerDailyQuestStarAwardInfo
}

func (db *Database) newPlayerDailyQuestStarAwardInfoUpdateLog(row, old *C.PlayerDailyQuestStarAwardInfo, goNew, goOld *PlayerDailyQuestStarAwardInfo) *PlayerDailyQuestStarAwardInfoUpdateLog {
	return &PlayerDailyQuestStarAwardInfoUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerDailyQuestStarAwardInfoUpdateLog) Free() {
	C.Free_PlayerDailyQuestStarAwardInfo(log.Old)
}

func (log *PlayerDailyQuestStarAwardInfoUpdateLog) InvokeHook() {
	if g_Hooks.PlayerDailyQuestStarAwardInfo != nil {
		g_Hooks.PlayerDailyQuestStarAwardInfo.Update(&PlayerDailyQuestStarAwardInfoRow{row: log.Row}, &PlayerDailyQuestStarAwardInfoRow{row: log.Old})
	}
}

func (log *PlayerDailyQuestStarAwardInfoUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerDailyQuestStarAwardInfoUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(30)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.Stars)
	file.WriteInt64(log.GoNew.Lastupdatetime)
	file.WriteString(log.GoNew.Awarded)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Stars)
	file.WriteInt64(log.GoOld.Lastupdatetime)
	file.WriteString(log.GoOld.Awarded)
}

func (log *PlayerDailyQuestStarAwardInfoUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDailyQuestStarAwardInfo.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.Stars)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Lastupdatetime)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Awarded), int(log.Row.Awarded_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerDailyQuestStarAwardInfoUpdateLog) CommitToTLog() {
}

func (log *PlayerDailyQuestStarAwardInfoUpdateLog) CommitToXdLog() {
}

func (log *PlayerDailyQuestStarAwardInfoUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerDailyQuestStarAwardInfoInsertLog) Rollback() {
	if log.db.tables.PlayerDailyQuestStarAwardInfo != log.Row {
		panic("insert rollback check failed 'PlayerDailyQuestStarAwardInfo'")
	}
	log.db.tables.PlayerDailyQuestStarAwardInfo = nil
	C.Free_PlayerDailyQuestStarAwardInfo(log.Row)
}

func (log *PlayerDailyQuestStarAwardInfoDeleteLog) Rollback() {
	if log.db.tables.PlayerDailyQuestStarAwardInfo != nil {
		panic("delete rollback check failed 'player_daily_quest_star_award_info'")
	}
	log.db.tables.PlayerDailyQuestStarAwardInfo = log.Old
}

func (log *PlayerDailyQuestStarAwardInfoUpdateLog) Rollback() {
	if log.db.tables.PlayerDailyQuestStarAwardInfo != log.Row {
		panic("update rollback check failed 'player_daily_quest_star_award_info'")
	}
	log.db.tables.PlayerDailyQuestStarAwardInfo = log.Old
	C.Free_PlayerDailyQuestStarAwardInfo(log.Row)
}

/* ========== player_daily_sign_in_record ========== */

type PlayerDailySignInRecordInsertLog struct {
	db *Database 
	Row *C.PlayerDailySignInRecord
	GoRow *PlayerDailySignInRecord
}

func (db *Database) newPlayerDailySignInRecordInsertLog(row *C.PlayerDailySignInRecord, goRow *PlayerDailySignInRecord) *PlayerDailySignInRecordInsertLog {
	return &PlayerDailySignInRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerDailySignInRecordInsertLog) Free() {
}

func (log *PlayerDailySignInRecordInsertLog) InvokeHook() {
	if g_Hooks.PlayerDailySignInRecord != nil {
		g_Hooks.PlayerDailySignInRecord.Insert(&PlayerDailySignInRecordRow{row: log.Row})
	}
}

func (log *PlayerDailySignInRecordInsertLog) SetEventId(time.Time) {
}

func (log *PlayerDailySignInRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(31)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.SignInTime)
}

func (log *PlayerDailySignInRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDailySignInRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SignInTime)))
	return stmt.Execute()
}

func (log *PlayerDailySignInRecordInsertLog) CommitToTLog() {
}

func (log *PlayerDailySignInRecordInsertLog) CommitToXdLog() {
}

func (log *PlayerDailySignInRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDailySignInRecordDeleteLog struct {
	db *Database
	Old *C.PlayerDailySignInRecord
	GoOld *PlayerDailySignInRecord
}

func (db *Database) newPlayerDailySignInRecordDeleteLog(old *C.PlayerDailySignInRecord, goOld *PlayerDailySignInRecord) *PlayerDailySignInRecordDeleteLog {
	return &PlayerDailySignInRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerDailySignInRecordDeleteLog) Free() {
	C.Free_PlayerDailySignInRecord(log.Old)
}

func (log *PlayerDailySignInRecordDeleteLog) InvokeHook() {
	if g_Hooks.PlayerDailySignInRecord != nil {
		g_Hooks.PlayerDailySignInRecord.Delete(&PlayerDailySignInRecordRow{row: log.Old})
	}
}

func (log *PlayerDailySignInRecordDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerDailySignInRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(31)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.SignInTime)
}

func (log *PlayerDailySignInRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDailySignInRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerDailySignInRecordDeleteLog) CommitToTLog() {
}

func (log *PlayerDailySignInRecordDeleteLog) CommitToXdLog() {
}

func (log *PlayerDailySignInRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDailySignInRecordUpdateLog struct {
	db *Database 
	Row *C.PlayerDailySignInRecord
	Old *C.PlayerDailySignInRecord
	GoNew *PlayerDailySignInRecord
	GoOld *PlayerDailySignInRecord
}

func (db *Database) newPlayerDailySignInRecordUpdateLog(row, old *C.PlayerDailySignInRecord, goNew, goOld *PlayerDailySignInRecord) *PlayerDailySignInRecordUpdateLog {
	return &PlayerDailySignInRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerDailySignInRecordUpdateLog) Free() {
	C.Free_PlayerDailySignInRecord(log.Old)
}

func (log *PlayerDailySignInRecordUpdateLog) InvokeHook() {
	if g_Hooks.PlayerDailySignInRecord != nil {
		g_Hooks.PlayerDailySignInRecord.Update(&PlayerDailySignInRecordRow{row: log.Row}, &PlayerDailySignInRecordRow{row: log.Old})
	}
}

func (log *PlayerDailySignInRecordUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerDailySignInRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(31)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.SignInTime)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.SignInTime)
}

func (log *PlayerDailySignInRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDailySignInRecord.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SignInTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerDailySignInRecordUpdateLog) CommitToTLog() {
}

func (log *PlayerDailySignInRecordUpdateLog) CommitToXdLog() {
}

func (log *PlayerDailySignInRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerDailySignInRecordInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerDailySignInRecord
	for crow := log.db.tables.PlayerDailySignInRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerDailySignInRecord'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerDailySignInRecord = log.db.tables.PlayerDailySignInRecord.next
	}
	C.Free_PlayerDailySignInRecord(log.Row)
}

func (log *PlayerDailySignInRecordDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerDailySignInRecord
	for crow := log.db.tables.PlayerDailySignInRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_daily_sign_in_record'")
	}
	log.Old.next = log.db.tables.PlayerDailySignInRecord
	log.db.tables.PlayerDailySignInRecord = log.Old
}

func (log *PlayerDailySignInRecordUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerDailySignInRecord
	for crow := log.db.tables.PlayerDailySignInRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_daily_sign_in_record'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerDailySignInRecord = log.Old
	}
	C.Free_PlayerDailySignInRecord(log.Row)
}

/* ========== player_daily_sign_in_state ========== */

type PlayerDailySignInStateInsertLog struct {
	db *Database 
	Row *C.PlayerDailySignInState
	GoRow *PlayerDailySignInState
}

func (db *Database) newPlayerDailySignInStateInsertLog(row *C.PlayerDailySignInState, goRow *PlayerDailySignInState) *PlayerDailySignInStateInsertLog {
	return &PlayerDailySignInStateInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerDailySignInStateInsertLog) Free() {
}

func (log *PlayerDailySignInStateInsertLog) InvokeHook() {
	if g_Hooks.PlayerDailySignInState != nil {
		g_Hooks.PlayerDailySignInState.Insert(&PlayerDailySignInStateRow{row: log.Row})
	}
}

func (log *PlayerDailySignInStateInsertLog) SetEventId(time.Time) {
}

func (log *PlayerDailySignInStateInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(32)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.UpdateTime)
	file.WriteInt16(log.GoRow.Record)
	file.WriteInt8(log.GoRow.SignedToday)
}

func (log *PlayerDailySignInStateInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDailySignInState.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.UpdateTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Record)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.SignedToday)))
	return stmt.Execute()
}

func (log *PlayerDailySignInStateInsertLog) CommitToTLog() {
}

func (log *PlayerDailySignInStateInsertLog) CommitToXdLog() {
}

func (log *PlayerDailySignInStateInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDailySignInStateDeleteLog struct {
	db *Database
	Old *C.PlayerDailySignInState
	GoOld *PlayerDailySignInState
}

func (db *Database) newPlayerDailySignInStateDeleteLog(old *C.PlayerDailySignInState, goOld *PlayerDailySignInState) *PlayerDailySignInStateDeleteLog {
	return &PlayerDailySignInStateDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerDailySignInStateDeleteLog) Free() {
	C.Free_PlayerDailySignInState(log.Old)
}

func (log *PlayerDailySignInStateDeleteLog) InvokeHook() {
	if g_Hooks.PlayerDailySignInState != nil {
		g_Hooks.PlayerDailySignInState.Delete(&PlayerDailySignInStateRow{row: log.Old})
	}
}

func (log *PlayerDailySignInStateDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerDailySignInStateDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(32)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.UpdateTime)
	file.WriteInt16(log.GoOld.Record)
	file.WriteInt8(log.GoOld.SignedToday)
}

func (log *PlayerDailySignInStateDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDailySignInState.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerDailySignInStateDeleteLog) CommitToTLog() {
}

func (log *PlayerDailySignInStateDeleteLog) CommitToXdLog() {
}

func (log *PlayerDailySignInStateDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDailySignInStateUpdateLog struct {
	db *Database 
	Row *C.PlayerDailySignInState
	Old *C.PlayerDailySignInState
	GoNew *PlayerDailySignInState
	GoOld *PlayerDailySignInState
}

func (db *Database) newPlayerDailySignInStateUpdateLog(row, old *C.PlayerDailySignInState, goNew, goOld *PlayerDailySignInState) *PlayerDailySignInStateUpdateLog {
	return &PlayerDailySignInStateUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerDailySignInStateUpdateLog) Free() {
	C.Free_PlayerDailySignInState(log.Old)
}

func (log *PlayerDailySignInStateUpdateLog) InvokeHook() {
	if g_Hooks.PlayerDailySignInState != nil {
		g_Hooks.PlayerDailySignInState.Update(&PlayerDailySignInStateRow{row: log.Row}, &PlayerDailySignInStateRow{row: log.Old})
	}
}

func (log *PlayerDailySignInStateUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerDailySignInStateUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(32)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.UpdateTime)
	file.WriteInt16(log.GoNew.Record)
	file.WriteInt8(log.GoNew.SignedToday)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.UpdateTime)
	file.WriteInt16(log.GoOld.Record)
	file.WriteInt8(log.GoOld.SignedToday)
}

func (log *PlayerDailySignInStateUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDailySignInState.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.UpdateTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Record)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.SignedToday)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerDailySignInStateUpdateLog) CommitToTLog() {
}

func (log *PlayerDailySignInStateUpdateLog) CommitToXdLog() {
}

func (log *PlayerDailySignInStateUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerDailySignInStateInsertLog) Rollback() {
	if log.db.tables.PlayerDailySignInState != log.Row {
		panic("insert rollback check failed 'PlayerDailySignInState'")
	}
	log.db.tables.PlayerDailySignInState = nil
	C.Free_PlayerDailySignInState(log.Row)
}

func (log *PlayerDailySignInStateDeleteLog) Rollback() {
	if log.db.tables.PlayerDailySignInState != nil {
		panic("delete rollback check failed 'player_daily_sign_in_state'")
	}
	log.db.tables.PlayerDailySignInState = log.Old
}

func (log *PlayerDailySignInStateUpdateLog) Rollback() {
	if log.db.tables.PlayerDailySignInState != log.Row {
		panic("update rollback check failed 'player_daily_sign_in_state'")
	}
	log.db.tables.PlayerDailySignInState = log.Old
	C.Free_PlayerDailySignInState(log.Row)
}

/* ========== player_despair_land_camp_state ========== */

type PlayerDespairLandCampStateInsertLog struct {
	db *Database 
	Row *C.PlayerDespairLandCampState
	GoRow *PlayerDespairLandCampState
}

func (db *Database) newPlayerDespairLandCampStateInsertLog(row *C.PlayerDespairLandCampState, goRow *PlayerDespairLandCampState) *PlayerDespairLandCampStateInsertLog {
	return &PlayerDespairLandCampStateInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerDespairLandCampStateInsertLog) Free() {
}

func (log *PlayerDespairLandCampStateInsertLog) InvokeHook() {
	if g_Hooks.PlayerDespairLandCampState != nil {
		g_Hooks.PlayerDespairLandCampState.Insert(&PlayerDespairLandCampStateRow{row: log.Row})
	}
}

func (log *PlayerDespairLandCampStateInsertLog) SetEventId(time.Time) {
}

func (log *PlayerDespairLandCampStateInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(33)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.CampType)
	file.WriteInt64(log.GoRow.Timestamp)
	file.WriteInt64(log.GoRow.Point)
	file.WriteInt64(log.GoRow.KillNum)
	file.WriteInt64(log.GoRow.DeadNum)
	file.WriteInt64(log.GoRow.DeadNumBoss)
	file.WriteInt64(log.GoRow.Hurt)
	file.WriteInt32(log.GoRow.BossBattleNum)
	file.WriteInt32(log.GoRow.DailyBossBattleNum)
	file.WriteInt64(log.GoRow.BossBattleTimestamp)
	file.WriteInt8(log.GoRow.Awarded)
}

func (log *PlayerDespairLandCampStateInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDespairLandCampState.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.CampType)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Point)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.KillNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DeadNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DeadNumBoss)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Hurt)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.BossBattleNum)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.DailyBossBattleNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BossBattleTimestamp)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Awarded)))
	return stmt.Execute()
}

func (log *PlayerDespairLandCampStateInsertLog) CommitToTLog() {
}

func (log *PlayerDespairLandCampStateInsertLog) CommitToXdLog() {
}

func (log *PlayerDespairLandCampStateInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDespairLandCampStateDeleteLog struct {
	db *Database
	Old *C.PlayerDespairLandCampState
	GoOld *PlayerDespairLandCampState
}

func (db *Database) newPlayerDespairLandCampStateDeleteLog(old *C.PlayerDespairLandCampState, goOld *PlayerDespairLandCampState) *PlayerDespairLandCampStateDeleteLog {
	return &PlayerDespairLandCampStateDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerDespairLandCampStateDeleteLog) Free() {
	C.Free_PlayerDespairLandCampState(log.Old)
}

func (log *PlayerDespairLandCampStateDeleteLog) InvokeHook() {
	if g_Hooks.PlayerDespairLandCampState != nil {
		g_Hooks.PlayerDespairLandCampState.Delete(&PlayerDespairLandCampStateRow{row: log.Old})
	}
}

func (log *PlayerDespairLandCampStateDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerDespairLandCampStateDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(33)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.CampType)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt64(log.GoOld.Point)
	file.WriteInt64(log.GoOld.KillNum)
	file.WriteInt64(log.GoOld.DeadNum)
	file.WriteInt64(log.GoOld.DeadNumBoss)
	file.WriteInt64(log.GoOld.Hurt)
	file.WriteInt32(log.GoOld.BossBattleNum)
	file.WriteInt32(log.GoOld.DailyBossBattleNum)
	file.WriteInt64(log.GoOld.BossBattleTimestamp)
	file.WriteInt8(log.GoOld.Awarded)
}

func (log *PlayerDespairLandCampStateDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDespairLandCampState.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerDespairLandCampStateDeleteLog) CommitToTLog() {
}

func (log *PlayerDespairLandCampStateDeleteLog) CommitToXdLog() {
}

func (log *PlayerDespairLandCampStateDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDespairLandCampStateUpdateLog struct {
	db *Database 
	Row *C.PlayerDespairLandCampState
	Old *C.PlayerDespairLandCampState
	GoNew *PlayerDespairLandCampState
	GoOld *PlayerDespairLandCampState
}

func (db *Database) newPlayerDespairLandCampStateUpdateLog(row, old *C.PlayerDespairLandCampState, goNew, goOld *PlayerDespairLandCampState) *PlayerDespairLandCampStateUpdateLog {
	return &PlayerDespairLandCampStateUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerDespairLandCampStateUpdateLog) Free() {
	C.Free_PlayerDespairLandCampState(log.Old)
}

func (log *PlayerDespairLandCampStateUpdateLog) InvokeHook() {
	if g_Hooks.PlayerDespairLandCampState != nil {
		g_Hooks.PlayerDespairLandCampState.Update(&PlayerDespairLandCampStateRow{row: log.Row}, &PlayerDespairLandCampStateRow{row: log.Old})
	}
}

func (log *PlayerDespairLandCampStateUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerDespairLandCampStateUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(33)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.CampType)
	file.WriteInt64(log.GoNew.Timestamp)
	file.WriteInt64(log.GoNew.Point)
	file.WriteInt64(log.GoNew.KillNum)
	file.WriteInt64(log.GoNew.DeadNum)
	file.WriteInt64(log.GoNew.DeadNumBoss)
	file.WriteInt64(log.GoNew.Hurt)
	file.WriteInt32(log.GoNew.BossBattleNum)
	file.WriteInt32(log.GoNew.DailyBossBattleNum)
	file.WriteInt64(log.GoNew.BossBattleTimestamp)
	file.WriteInt8(log.GoNew.Awarded)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.CampType)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt64(log.GoOld.Point)
	file.WriteInt64(log.GoOld.KillNum)
	file.WriteInt64(log.GoOld.DeadNum)
	file.WriteInt64(log.GoOld.DeadNumBoss)
	file.WriteInt64(log.GoOld.Hurt)
	file.WriteInt32(log.GoOld.BossBattleNum)
	file.WriteInt32(log.GoOld.DailyBossBattleNum)
	file.WriteInt64(log.GoOld.BossBattleTimestamp)
	file.WriteInt8(log.GoOld.Awarded)
}

func (log *PlayerDespairLandCampStateUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDespairLandCampState.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.CampType)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Point)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.KillNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DeadNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DeadNumBoss)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Hurt)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.BossBattleNum)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.DailyBossBattleNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BossBattleTimestamp)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Awarded)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerDespairLandCampStateUpdateLog) CommitToTLog() {
}

func (log *PlayerDespairLandCampStateUpdateLog) CommitToXdLog() {
}

func (log *PlayerDespairLandCampStateUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerDespairLandCampStateInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerDespairLandCampState
	for crow := log.db.tables.PlayerDespairLandCampState; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerDespairLandCampState'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerDespairLandCampState = log.db.tables.PlayerDespairLandCampState.next
	}
	C.Free_PlayerDespairLandCampState(log.Row)
}

func (log *PlayerDespairLandCampStateDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerDespairLandCampState
	for crow := log.db.tables.PlayerDespairLandCampState; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_despair_land_camp_state'")
	}
	log.Old.next = log.db.tables.PlayerDespairLandCampState
	log.db.tables.PlayerDespairLandCampState = log.Old
}

func (log *PlayerDespairLandCampStateUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerDespairLandCampState
	for crow := log.db.tables.PlayerDespairLandCampState; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_despair_land_camp_state'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerDespairLandCampState = log.Old
	}
	C.Free_PlayerDespairLandCampState(log.Row)
}

/* ========== player_despair_land_camp_state_history ========== */

type PlayerDespairLandCampStateHistoryInsertLog struct {
	db *Database 
	Row *C.PlayerDespairLandCampStateHistory
	GoRow *PlayerDespairLandCampStateHistory
}

func (db *Database) newPlayerDespairLandCampStateHistoryInsertLog(row *C.PlayerDespairLandCampStateHistory, goRow *PlayerDespairLandCampStateHistory) *PlayerDespairLandCampStateHistoryInsertLog {
	return &PlayerDespairLandCampStateHistoryInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerDespairLandCampStateHistoryInsertLog) Free() {
}

func (log *PlayerDespairLandCampStateHistoryInsertLog) InvokeHook() {
	if g_Hooks.PlayerDespairLandCampStateHistory != nil {
		g_Hooks.PlayerDespairLandCampStateHistory.Insert(&PlayerDespairLandCampStateHistoryRow{row: log.Row})
	}
}

func (log *PlayerDespairLandCampStateHistoryInsertLog) SetEventId(time.Time) {
}

func (log *PlayerDespairLandCampStateHistoryInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(34)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.CampType)
	file.WriteInt64(log.GoRow.Timestamp)
	file.WriteInt64(log.GoRow.Point)
	file.WriteInt64(log.GoRow.KillNum)
	file.WriteInt64(log.GoRow.DeadNum)
	file.WriteInt64(log.GoRow.DeadNumBoss)
	file.WriteInt64(log.GoRow.Hurt)
	file.WriteInt32(log.GoRow.BossBattleNum)
	file.WriteInt32(log.GoRow.DailyBossBattleNum)
	file.WriteInt64(log.GoRow.BossBattleTimestamp)
	file.WriteInt8(log.GoRow.Awarded)
}

func (log *PlayerDespairLandCampStateHistoryInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDespairLandCampStateHistory.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.CampType)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Point)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.KillNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DeadNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DeadNumBoss)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Hurt)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.BossBattleNum)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.DailyBossBattleNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BossBattleTimestamp)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Awarded)))
	return stmt.Execute()
}

func (log *PlayerDespairLandCampStateHistoryInsertLog) CommitToTLog() {
}

func (log *PlayerDespairLandCampStateHistoryInsertLog) CommitToXdLog() {
}

func (log *PlayerDespairLandCampStateHistoryInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDespairLandCampStateHistoryDeleteLog struct {
	db *Database
	Old *C.PlayerDespairLandCampStateHistory
	GoOld *PlayerDespairLandCampStateHistory
}

func (db *Database) newPlayerDespairLandCampStateHistoryDeleteLog(old *C.PlayerDespairLandCampStateHistory, goOld *PlayerDespairLandCampStateHistory) *PlayerDespairLandCampStateHistoryDeleteLog {
	return &PlayerDespairLandCampStateHistoryDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerDespairLandCampStateHistoryDeleteLog) Free() {
	C.Free_PlayerDespairLandCampStateHistory(log.Old)
}

func (log *PlayerDespairLandCampStateHistoryDeleteLog) InvokeHook() {
	if g_Hooks.PlayerDespairLandCampStateHistory != nil {
		g_Hooks.PlayerDespairLandCampStateHistory.Delete(&PlayerDespairLandCampStateHistoryRow{row: log.Old})
	}
}

func (log *PlayerDespairLandCampStateHistoryDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerDespairLandCampStateHistoryDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(34)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.CampType)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt64(log.GoOld.Point)
	file.WriteInt64(log.GoOld.KillNum)
	file.WriteInt64(log.GoOld.DeadNum)
	file.WriteInt64(log.GoOld.DeadNumBoss)
	file.WriteInt64(log.GoOld.Hurt)
	file.WriteInt32(log.GoOld.BossBattleNum)
	file.WriteInt32(log.GoOld.DailyBossBattleNum)
	file.WriteInt64(log.GoOld.BossBattleTimestamp)
	file.WriteInt8(log.GoOld.Awarded)
}

func (log *PlayerDespairLandCampStateHistoryDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDespairLandCampStateHistory.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerDespairLandCampStateHistoryDeleteLog) CommitToTLog() {
}

func (log *PlayerDespairLandCampStateHistoryDeleteLog) CommitToXdLog() {
}

func (log *PlayerDespairLandCampStateHistoryDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDespairLandCampStateHistoryUpdateLog struct {
	db *Database 
	Row *C.PlayerDespairLandCampStateHistory
	Old *C.PlayerDespairLandCampStateHistory
	GoNew *PlayerDespairLandCampStateHistory
	GoOld *PlayerDespairLandCampStateHistory
}

func (db *Database) newPlayerDespairLandCampStateHistoryUpdateLog(row, old *C.PlayerDespairLandCampStateHistory, goNew, goOld *PlayerDespairLandCampStateHistory) *PlayerDespairLandCampStateHistoryUpdateLog {
	return &PlayerDespairLandCampStateHistoryUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerDespairLandCampStateHistoryUpdateLog) Free() {
	C.Free_PlayerDespairLandCampStateHistory(log.Old)
}

func (log *PlayerDespairLandCampStateHistoryUpdateLog) InvokeHook() {
	if g_Hooks.PlayerDespairLandCampStateHistory != nil {
		g_Hooks.PlayerDespairLandCampStateHistory.Update(&PlayerDespairLandCampStateHistoryRow{row: log.Row}, &PlayerDespairLandCampStateHistoryRow{row: log.Old})
	}
}

func (log *PlayerDespairLandCampStateHistoryUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerDespairLandCampStateHistoryUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(34)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.CampType)
	file.WriteInt64(log.GoNew.Timestamp)
	file.WriteInt64(log.GoNew.Point)
	file.WriteInt64(log.GoNew.KillNum)
	file.WriteInt64(log.GoNew.DeadNum)
	file.WriteInt64(log.GoNew.DeadNumBoss)
	file.WriteInt64(log.GoNew.Hurt)
	file.WriteInt32(log.GoNew.BossBattleNum)
	file.WriteInt32(log.GoNew.DailyBossBattleNum)
	file.WriteInt64(log.GoNew.BossBattleTimestamp)
	file.WriteInt8(log.GoNew.Awarded)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.CampType)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt64(log.GoOld.Point)
	file.WriteInt64(log.GoOld.KillNum)
	file.WriteInt64(log.GoOld.DeadNum)
	file.WriteInt64(log.GoOld.DeadNumBoss)
	file.WriteInt64(log.GoOld.Hurt)
	file.WriteInt32(log.GoOld.BossBattleNum)
	file.WriteInt32(log.GoOld.DailyBossBattleNum)
	file.WriteInt64(log.GoOld.BossBattleTimestamp)
	file.WriteInt8(log.GoOld.Awarded)
}

func (log *PlayerDespairLandCampStateHistoryUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDespairLandCampStateHistory.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.CampType)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Point)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.KillNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DeadNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DeadNumBoss)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Hurt)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.BossBattleNum)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.DailyBossBattleNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BossBattleTimestamp)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Awarded)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerDespairLandCampStateHistoryUpdateLog) CommitToTLog() {
}

func (log *PlayerDespairLandCampStateHistoryUpdateLog) CommitToXdLog() {
}

func (log *PlayerDespairLandCampStateHistoryUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerDespairLandCampStateHistoryInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerDespairLandCampStateHistory
	for crow := log.db.tables.PlayerDespairLandCampStateHistory; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerDespairLandCampStateHistory'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerDespairLandCampStateHistory = log.db.tables.PlayerDespairLandCampStateHistory.next
	}
	C.Free_PlayerDespairLandCampStateHistory(log.Row)
}

func (log *PlayerDespairLandCampStateHistoryDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerDespairLandCampStateHistory
	for crow := log.db.tables.PlayerDespairLandCampStateHistory; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_despair_land_camp_state_history'")
	}
	log.Old.next = log.db.tables.PlayerDespairLandCampStateHistory
	log.db.tables.PlayerDespairLandCampStateHistory = log.Old
}

func (log *PlayerDespairLandCampStateHistoryUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerDespairLandCampStateHistory
	for crow := log.db.tables.PlayerDespairLandCampStateHistory; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_despair_land_camp_state_history'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerDespairLandCampStateHistory = log.Old
	}
	C.Free_PlayerDespairLandCampStateHistory(log.Row)
}

/* ========== player_despair_land_level_record ========== */

type PlayerDespairLandLevelRecordInsertLog struct {
	db *Database 
	Row *C.PlayerDespairLandLevelRecord
	GoRow *PlayerDespairLandLevelRecord
}

func (db *Database) newPlayerDespairLandLevelRecordInsertLog(row *C.PlayerDespairLandLevelRecord, goRow *PlayerDespairLandLevelRecord) *PlayerDespairLandLevelRecordInsertLog {
	return &PlayerDespairLandLevelRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerDespairLandLevelRecordInsertLog) Free() {
}

func (log *PlayerDespairLandLevelRecordInsertLog) InvokeHook() {
	if g_Hooks.PlayerDespairLandLevelRecord != nil {
		g_Hooks.PlayerDespairLandLevelRecord.Insert(&PlayerDespairLandLevelRecordRow{row: log.Row})
	}
}

func (log *PlayerDespairLandLevelRecordInsertLog) SetEventId(time.Time) {
}

func (log *PlayerDespairLandLevelRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(35)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.CampType)
	file.WriteInt64(log.GoRow.Timestamp)
	file.WriteInt32(log.GoRow.LevelId)
	file.WriteInt8(log.GoRow.Round)
	file.WriteInt64(log.GoRow.PassedTimestamp)
	file.WriteInt64(log.GoRow.FirstPassedTimestamp)
	file.WriteInt32(log.GoRow.FirstPassedFightnum)
	file.WriteInt8(log.GoRow.PassedAward)
	file.WriteInt8(log.GoRow.PerfectAward)
	file.WriteInt8(log.GoRow.MilestoneAward)
}

func (log *PlayerDespairLandLevelRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDespairLandLevelRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.CampType)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.LevelId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Round)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.PassedTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FirstPassedTimestamp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.FirstPassedFightnum)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.PassedAward)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.PerfectAward)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.MilestoneAward)))
	return stmt.Execute()
}

func (log *PlayerDespairLandLevelRecordInsertLog) CommitToTLog() {
}

func (log *PlayerDespairLandLevelRecordInsertLog) CommitToXdLog() {
}

func (log *PlayerDespairLandLevelRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDespairLandLevelRecordDeleteLog struct {
	db *Database
	Old *C.PlayerDespairLandLevelRecord
	GoOld *PlayerDespairLandLevelRecord
}

func (db *Database) newPlayerDespairLandLevelRecordDeleteLog(old *C.PlayerDespairLandLevelRecord, goOld *PlayerDespairLandLevelRecord) *PlayerDespairLandLevelRecordDeleteLog {
	return &PlayerDespairLandLevelRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerDespairLandLevelRecordDeleteLog) Free() {
	C.Free_PlayerDespairLandLevelRecord(log.Old)
}

func (log *PlayerDespairLandLevelRecordDeleteLog) InvokeHook() {
	if g_Hooks.PlayerDespairLandLevelRecord != nil {
		g_Hooks.PlayerDespairLandLevelRecord.Delete(&PlayerDespairLandLevelRecordRow{row: log.Old})
	}
}

func (log *PlayerDespairLandLevelRecordDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerDespairLandLevelRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(35)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.CampType)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt32(log.GoOld.LevelId)
	file.WriteInt8(log.GoOld.Round)
	file.WriteInt64(log.GoOld.PassedTimestamp)
	file.WriteInt64(log.GoOld.FirstPassedTimestamp)
	file.WriteInt32(log.GoOld.FirstPassedFightnum)
	file.WriteInt8(log.GoOld.PassedAward)
	file.WriteInt8(log.GoOld.PerfectAward)
	file.WriteInt8(log.GoOld.MilestoneAward)
}

func (log *PlayerDespairLandLevelRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDespairLandLevelRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerDespairLandLevelRecordDeleteLog) CommitToTLog() {
}

func (log *PlayerDespairLandLevelRecordDeleteLog) CommitToXdLog() {
}

func (log *PlayerDespairLandLevelRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDespairLandLevelRecordUpdateLog struct {
	db *Database 
	Row *C.PlayerDespairLandLevelRecord
	Old *C.PlayerDespairLandLevelRecord
	GoNew *PlayerDespairLandLevelRecord
	GoOld *PlayerDespairLandLevelRecord
}

func (db *Database) newPlayerDespairLandLevelRecordUpdateLog(row, old *C.PlayerDespairLandLevelRecord, goNew, goOld *PlayerDespairLandLevelRecord) *PlayerDespairLandLevelRecordUpdateLog {
	return &PlayerDespairLandLevelRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerDespairLandLevelRecordUpdateLog) Free() {
	C.Free_PlayerDespairLandLevelRecord(log.Old)
}

func (log *PlayerDespairLandLevelRecordUpdateLog) InvokeHook() {
	if g_Hooks.PlayerDespairLandLevelRecord != nil {
		g_Hooks.PlayerDespairLandLevelRecord.Update(&PlayerDespairLandLevelRecordRow{row: log.Row}, &PlayerDespairLandLevelRecordRow{row: log.Old})
	}
}

func (log *PlayerDespairLandLevelRecordUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerDespairLandLevelRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(35)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.CampType)
	file.WriteInt64(log.GoNew.Timestamp)
	file.WriteInt32(log.GoNew.LevelId)
	file.WriteInt8(log.GoNew.Round)
	file.WriteInt64(log.GoNew.PassedTimestamp)
	file.WriteInt64(log.GoNew.FirstPassedTimestamp)
	file.WriteInt32(log.GoNew.FirstPassedFightnum)
	file.WriteInt8(log.GoNew.PassedAward)
	file.WriteInt8(log.GoNew.PerfectAward)
	file.WriteInt8(log.GoNew.MilestoneAward)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.CampType)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt32(log.GoOld.LevelId)
	file.WriteInt8(log.GoOld.Round)
	file.WriteInt64(log.GoOld.PassedTimestamp)
	file.WriteInt64(log.GoOld.FirstPassedTimestamp)
	file.WriteInt32(log.GoOld.FirstPassedFightnum)
	file.WriteInt8(log.GoOld.PassedAward)
	file.WriteInt8(log.GoOld.PerfectAward)
	file.WriteInt8(log.GoOld.MilestoneAward)
}

func (log *PlayerDespairLandLevelRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDespairLandLevelRecord.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.CampType)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.LevelId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Round)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.PassedTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FirstPassedTimestamp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.FirstPassedFightnum)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.PassedAward)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.PerfectAward)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.MilestoneAward)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerDespairLandLevelRecordUpdateLog) CommitToTLog() {
}

func (log *PlayerDespairLandLevelRecordUpdateLog) CommitToXdLog() {
}

func (log *PlayerDespairLandLevelRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerDespairLandLevelRecordInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerDespairLandLevelRecord
	for crow := log.db.tables.PlayerDespairLandLevelRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerDespairLandLevelRecord'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerDespairLandLevelRecord = log.db.tables.PlayerDespairLandLevelRecord.next
	}
	C.Free_PlayerDespairLandLevelRecord(log.Row)
}

func (log *PlayerDespairLandLevelRecordDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerDespairLandLevelRecord
	for crow := log.db.tables.PlayerDespairLandLevelRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_despair_land_level_record'")
	}
	log.Old.next = log.db.tables.PlayerDespairLandLevelRecord
	log.db.tables.PlayerDespairLandLevelRecord = log.Old
}

func (log *PlayerDespairLandLevelRecordUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerDespairLandLevelRecord
	for crow := log.db.tables.PlayerDespairLandLevelRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_despair_land_level_record'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerDespairLandLevelRecord = log.Old
	}
	C.Free_PlayerDespairLandLevelRecord(log.Row)
}

/* ========== player_despair_land_state ========== */

type PlayerDespairLandStateInsertLog struct {
	db *Database 
	Row *C.PlayerDespairLandState
	GoRow *PlayerDespairLandState
}

func (db *Database) newPlayerDespairLandStateInsertLog(row *C.PlayerDespairLandState, goRow *PlayerDespairLandState) *PlayerDespairLandStateInsertLog {
	return &PlayerDespairLandStateInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerDespairLandStateInsertLog) Free() {
}

func (log *PlayerDespairLandStateInsertLog) InvokeHook() {
	if g_Hooks.PlayerDespairLandState != nil {
		g_Hooks.PlayerDespairLandState.Insert(&PlayerDespairLandStateRow{row: log.Row})
	}
}

func (log *PlayerDespairLandStateInsertLog) SetEventId(time.Time) {
}

func (log *PlayerDespairLandStateInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(36)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.Point)
	file.WriteInt64(log.GoRow.KillNum)
	file.WriteInt64(log.GoRow.DeadNum)
	file.WriteInt16(log.GoRow.DailyBattleNum)
	file.WriteInt64(log.GoRow.DailyBattleTimestamp)
	file.WriteInt16(log.GoRow.DailyBoughtBattleNum)
	file.WriteInt64(log.GoRow.DailyBoughtTimestamp)
	file.WriteInt64(log.GoRow.DailyBossBoughtTimestamp)
	file.WriteInt16(log.GoRow.DailyBossBeastBoughtBattleNum)
	file.WriteInt16(log.GoRow.DailyBossWalkingDeadBoughtBattleNum)
	file.WriteInt16(log.GoRow.DailyBossDevilBoughtBattleNum)
}

func (log *PlayerDespairLandStateInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDespairLandState.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Point)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.KillNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DeadNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyBattleNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DailyBattleTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyBoughtBattleNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DailyBoughtTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DailyBossBoughtTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyBossBeastBoughtBattleNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyBossWalkingDeadBoughtBattleNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyBossDevilBoughtBattleNum)))
	return stmt.Execute()
}

func (log *PlayerDespairLandStateInsertLog) CommitToTLog() {
}

func (log *PlayerDespairLandStateInsertLog) CommitToXdLog() {
}

func (log *PlayerDespairLandStateInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDespairLandStateDeleteLog struct {
	db *Database
	Old *C.PlayerDespairLandState
	GoOld *PlayerDespairLandState
}

func (db *Database) newPlayerDespairLandStateDeleteLog(old *C.PlayerDespairLandState, goOld *PlayerDespairLandState) *PlayerDespairLandStateDeleteLog {
	return &PlayerDespairLandStateDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerDespairLandStateDeleteLog) Free() {
	C.Free_PlayerDespairLandState(log.Old)
}

func (log *PlayerDespairLandStateDeleteLog) InvokeHook() {
	if g_Hooks.PlayerDespairLandState != nil {
		g_Hooks.PlayerDespairLandState.Delete(&PlayerDespairLandStateRow{row: log.Old})
	}
}

func (log *PlayerDespairLandStateDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerDespairLandStateDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(36)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.Point)
	file.WriteInt64(log.GoOld.KillNum)
	file.WriteInt64(log.GoOld.DeadNum)
	file.WriteInt16(log.GoOld.DailyBattleNum)
	file.WriteInt64(log.GoOld.DailyBattleTimestamp)
	file.WriteInt16(log.GoOld.DailyBoughtBattleNum)
	file.WriteInt64(log.GoOld.DailyBoughtTimestamp)
	file.WriteInt64(log.GoOld.DailyBossBoughtTimestamp)
	file.WriteInt16(log.GoOld.DailyBossBeastBoughtBattleNum)
	file.WriteInt16(log.GoOld.DailyBossWalkingDeadBoughtBattleNum)
	file.WriteInt16(log.GoOld.DailyBossDevilBoughtBattleNum)
}

func (log *PlayerDespairLandStateDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDespairLandState.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerDespairLandStateDeleteLog) CommitToTLog() {
}

func (log *PlayerDespairLandStateDeleteLog) CommitToXdLog() {
}

func (log *PlayerDespairLandStateDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDespairLandStateUpdateLog struct {
	db *Database 
	Row *C.PlayerDespairLandState
	Old *C.PlayerDespairLandState
	GoNew *PlayerDespairLandState
	GoOld *PlayerDespairLandState
}

func (db *Database) newPlayerDespairLandStateUpdateLog(row, old *C.PlayerDespairLandState, goNew, goOld *PlayerDespairLandState) *PlayerDespairLandStateUpdateLog {
	return &PlayerDespairLandStateUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerDespairLandStateUpdateLog) Free() {
	C.Free_PlayerDespairLandState(log.Old)
}

func (log *PlayerDespairLandStateUpdateLog) InvokeHook() {
	if g_Hooks.PlayerDespairLandState != nil {
		g_Hooks.PlayerDespairLandState.Update(&PlayerDespairLandStateRow{row: log.Row}, &PlayerDespairLandStateRow{row: log.Old})
	}
}

func (log *PlayerDespairLandStateUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerDespairLandStateUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(36)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.Point)
	file.WriteInt64(log.GoNew.KillNum)
	file.WriteInt64(log.GoNew.DeadNum)
	file.WriteInt16(log.GoNew.DailyBattleNum)
	file.WriteInt64(log.GoNew.DailyBattleTimestamp)
	file.WriteInt16(log.GoNew.DailyBoughtBattleNum)
	file.WriteInt64(log.GoNew.DailyBoughtTimestamp)
	file.WriteInt64(log.GoNew.DailyBossBoughtTimestamp)
	file.WriteInt16(log.GoNew.DailyBossBeastBoughtBattleNum)
	file.WriteInt16(log.GoNew.DailyBossWalkingDeadBoughtBattleNum)
	file.WriteInt16(log.GoNew.DailyBossDevilBoughtBattleNum)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.Point)
	file.WriteInt64(log.GoOld.KillNum)
	file.WriteInt64(log.GoOld.DeadNum)
	file.WriteInt16(log.GoOld.DailyBattleNum)
	file.WriteInt64(log.GoOld.DailyBattleTimestamp)
	file.WriteInt16(log.GoOld.DailyBoughtBattleNum)
	file.WriteInt64(log.GoOld.DailyBoughtTimestamp)
	file.WriteInt64(log.GoOld.DailyBossBoughtTimestamp)
	file.WriteInt16(log.GoOld.DailyBossBeastBoughtBattleNum)
	file.WriteInt16(log.GoOld.DailyBossWalkingDeadBoughtBattleNum)
	file.WriteInt16(log.GoOld.DailyBossDevilBoughtBattleNum)
}

func (log *PlayerDespairLandStateUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDespairLandState.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Point)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.KillNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DeadNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyBattleNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DailyBattleTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyBoughtBattleNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DailyBoughtTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DailyBossBoughtTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyBossBeastBoughtBattleNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyBossWalkingDeadBoughtBattleNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyBossDevilBoughtBattleNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerDespairLandStateUpdateLog) CommitToTLog() {
}

func (log *PlayerDespairLandStateUpdateLog) CommitToXdLog() {
}

func (log *PlayerDespairLandStateUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerDespairLandStateInsertLog) Rollback() {
	if log.db.tables.PlayerDespairLandState != log.Row {
		panic("insert rollback check failed 'PlayerDespairLandState'")
	}
	log.db.tables.PlayerDespairLandState = nil
	C.Free_PlayerDespairLandState(log.Row)
}

func (log *PlayerDespairLandStateDeleteLog) Rollback() {
	if log.db.tables.PlayerDespairLandState != nil {
		panic("delete rollback check failed 'player_despair_land_state'")
	}
	log.db.tables.PlayerDespairLandState = log.Old
}

func (log *PlayerDespairLandStateUpdateLog) Rollback() {
	if log.db.tables.PlayerDespairLandState != log.Row {
		panic("update rollback check failed 'player_despair_land_state'")
	}
	log.db.tables.PlayerDespairLandState = log.Old
	C.Free_PlayerDespairLandState(log.Row)
}

/* ========== player_driving_sword_event ========== */

type PlayerDrivingSwordEventInsertLog struct {
	db *Database 
	Row *C.PlayerDrivingSwordEvent
	GoRow *PlayerDrivingSwordEvent
}

func (db *Database) newPlayerDrivingSwordEventInsertLog(row *C.PlayerDrivingSwordEvent, goRow *PlayerDrivingSwordEvent) *PlayerDrivingSwordEventInsertLog {
	return &PlayerDrivingSwordEventInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerDrivingSwordEventInsertLog) Free() {
}

func (log *PlayerDrivingSwordEventInsertLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordEvent != nil {
		g_Hooks.PlayerDrivingSwordEvent.Insert(&PlayerDrivingSwordEventRow{row: log.Row})
	}
}

func (log *PlayerDrivingSwordEventInsertLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordEventInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(37)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.CloudId)
	file.WriteInt8(log.GoRow.X)
	file.WriteInt8(log.GoRow.Y)
	file.WriteInt8(log.GoRow.EventType)
	file.WriteInt8(log.GoRow.DataId)
}

func (log *PlayerDrivingSwordEventInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordEvent.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CloudId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.X)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Y)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.EventType)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DataId)))
	return stmt.Execute()
}

func (log *PlayerDrivingSwordEventInsertLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordEventInsertLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordEventInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDrivingSwordEventDeleteLog struct {
	db *Database
	Old *C.PlayerDrivingSwordEvent
	GoOld *PlayerDrivingSwordEvent
}

func (db *Database) newPlayerDrivingSwordEventDeleteLog(old *C.PlayerDrivingSwordEvent, goOld *PlayerDrivingSwordEvent) *PlayerDrivingSwordEventDeleteLog {
	return &PlayerDrivingSwordEventDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerDrivingSwordEventDeleteLog) Free() {
	C.Free_PlayerDrivingSwordEvent(log.Old)
}

func (log *PlayerDrivingSwordEventDeleteLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordEvent != nil {
		g_Hooks.PlayerDrivingSwordEvent.Delete(&PlayerDrivingSwordEventRow{row: log.Old})
	}
}

func (log *PlayerDrivingSwordEventDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordEventDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(37)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.CloudId)
	file.WriteInt8(log.GoOld.X)
	file.WriteInt8(log.GoOld.Y)
	file.WriteInt8(log.GoOld.EventType)
	file.WriteInt8(log.GoOld.DataId)
}

func (log *PlayerDrivingSwordEventDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordEvent.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerDrivingSwordEventDeleteLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordEventDeleteLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordEventDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDrivingSwordEventUpdateLog struct {
	db *Database 
	Row *C.PlayerDrivingSwordEvent
	Old *C.PlayerDrivingSwordEvent
	GoNew *PlayerDrivingSwordEvent
	GoOld *PlayerDrivingSwordEvent
}

func (db *Database) newPlayerDrivingSwordEventUpdateLog(row, old *C.PlayerDrivingSwordEvent, goNew, goOld *PlayerDrivingSwordEvent) *PlayerDrivingSwordEventUpdateLog {
	return &PlayerDrivingSwordEventUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerDrivingSwordEventUpdateLog) Free() {
	C.Free_PlayerDrivingSwordEvent(log.Old)
}

func (log *PlayerDrivingSwordEventUpdateLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordEvent != nil {
		g_Hooks.PlayerDrivingSwordEvent.Update(&PlayerDrivingSwordEventRow{row: log.Row}, &PlayerDrivingSwordEventRow{row: log.Old})
	}
}

func (log *PlayerDrivingSwordEventUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordEventUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(37)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.CloudId)
	file.WriteInt8(log.GoNew.X)
	file.WriteInt8(log.GoNew.Y)
	file.WriteInt8(log.GoNew.EventType)
	file.WriteInt8(log.GoNew.DataId)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.CloudId)
	file.WriteInt8(log.GoOld.X)
	file.WriteInt8(log.GoOld.Y)
	file.WriteInt8(log.GoOld.EventType)
	file.WriteInt8(log.GoOld.DataId)
}

func (log *PlayerDrivingSwordEventUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordEvent.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CloudId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.X)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Y)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.EventType)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DataId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerDrivingSwordEventUpdateLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordEventUpdateLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordEventUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerDrivingSwordEventInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerDrivingSwordEvent
	for crow := log.db.tables.PlayerDrivingSwordEvent; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerDrivingSwordEvent'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerDrivingSwordEvent = log.db.tables.PlayerDrivingSwordEvent.next
	}
	C.Free_PlayerDrivingSwordEvent(log.Row)
}

func (log *PlayerDrivingSwordEventDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerDrivingSwordEvent
	for crow := log.db.tables.PlayerDrivingSwordEvent; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_driving_sword_event'")
	}
	log.Old.next = log.db.tables.PlayerDrivingSwordEvent
	log.db.tables.PlayerDrivingSwordEvent = log.Old
}

func (log *PlayerDrivingSwordEventUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerDrivingSwordEvent
	for crow := log.db.tables.PlayerDrivingSwordEvent; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_driving_sword_event'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerDrivingSwordEvent = log.Old
	}
	C.Free_PlayerDrivingSwordEvent(log.Row)
}

/* ========== player_driving_sword_event_exploring ========== */

type PlayerDrivingSwordEventExploringInsertLog struct {
	db *Database 
	Row *C.PlayerDrivingSwordEventExploring
	GoRow *PlayerDrivingSwordEventExploring
}

func (db *Database) newPlayerDrivingSwordEventExploringInsertLog(row *C.PlayerDrivingSwordEventExploring, goRow *PlayerDrivingSwordEventExploring) *PlayerDrivingSwordEventExploringInsertLog {
	return &PlayerDrivingSwordEventExploringInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerDrivingSwordEventExploringInsertLog) Free() {
}

func (log *PlayerDrivingSwordEventExploringInsertLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordEventExploring != nil {
		g_Hooks.PlayerDrivingSwordEventExploring.Insert(&PlayerDrivingSwordEventExploringRow{row: log.Row})
	}
}

func (log *PlayerDrivingSwordEventExploringInsertLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordEventExploringInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(38)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.Status)
	file.WriteInt64(log.GoRow.GarrisonStart)
	file.WriteInt64(log.GoRow.GarrisonTime)
	file.WriteInt64(log.GoRow.AwardTime)
	file.WriteInt8(log.GoRow.RoleId)
	file.WriteInt16(log.GoRow.CloudId)
	file.WriteInt8(log.GoRow.X)
	file.WriteInt8(log.GoRow.Y)
	file.WriteInt8(log.GoRow.DataId)
}

func (log *PlayerDrivingSwordEventExploringInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordEventExploring.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Status)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.GarrisonStart)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.GarrisonTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.AwardTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RoleId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CloudId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.X)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Y)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DataId)))
	return stmt.Execute()
}

func (log *PlayerDrivingSwordEventExploringInsertLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordEventExploringInsertLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordEventExploringInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDrivingSwordEventExploringDeleteLog struct {
	db *Database
	Old *C.PlayerDrivingSwordEventExploring
	GoOld *PlayerDrivingSwordEventExploring
}

func (db *Database) newPlayerDrivingSwordEventExploringDeleteLog(old *C.PlayerDrivingSwordEventExploring, goOld *PlayerDrivingSwordEventExploring) *PlayerDrivingSwordEventExploringDeleteLog {
	return &PlayerDrivingSwordEventExploringDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerDrivingSwordEventExploringDeleteLog) Free() {
	C.Free_PlayerDrivingSwordEventExploring(log.Old)
}

func (log *PlayerDrivingSwordEventExploringDeleteLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordEventExploring != nil {
		g_Hooks.PlayerDrivingSwordEventExploring.Delete(&PlayerDrivingSwordEventExploringRow{row: log.Old})
	}
}

func (log *PlayerDrivingSwordEventExploringDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordEventExploringDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(38)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.Status)
	file.WriteInt64(log.GoOld.GarrisonStart)
	file.WriteInt64(log.GoOld.GarrisonTime)
	file.WriteInt64(log.GoOld.AwardTime)
	file.WriteInt8(log.GoOld.RoleId)
	file.WriteInt16(log.GoOld.CloudId)
	file.WriteInt8(log.GoOld.X)
	file.WriteInt8(log.GoOld.Y)
	file.WriteInt8(log.GoOld.DataId)
}

func (log *PlayerDrivingSwordEventExploringDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordEventExploring.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerDrivingSwordEventExploringDeleteLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordEventExploringDeleteLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordEventExploringDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDrivingSwordEventExploringUpdateLog struct {
	db *Database 
	Row *C.PlayerDrivingSwordEventExploring
	Old *C.PlayerDrivingSwordEventExploring
	GoNew *PlayerDrivingSwordEventExploring
	GoOld *PlayerDrivingSwordEventExploring
}

func (db *Database) newPlayerDrivingSwordEventExploringUpdateLog(row, old *C.PlayerDrivingSwordEventExploring, goNew, goOld *PlayerDrivingSwordEventExploring) *PlayerDrivingSwordEventExploringUpdateLog {
	return &PlayerDrivingSwordEventExploringUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerDrivingSwordEventExploringUpdateLog) Free() {
	C.Free_PlayerDrivingSwordEventExploring(log.Old)
}

func (log *PlayerDrivingSwordEventExploringUpdateLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordEventExploring != nil {
		g_Hooks.PlayerDrivingSwordEventExploring.Update(&PlayerDrivingSwordEventExploringRow{row: log.Row}, &PlayerDrivingSwordEventExploringRow{row: log.Old})
	}
}

func (log *PlayerDrivingSwordEventExploringUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordEventExploringUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(38)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.Status)
	file.WriteInt64(log.GoNew.GarrisonStart)
	file.WriteInt64(log.GoNew.GarrisonTime)
	file.WriteInt64(log.GoNew.AwardTime)
	file.WriteInt8(log.GoNew.RoleId)
	file.WriteInt16(log.GoNew.CloudId)
	file.WriteInt8(log.GoNew.X)
	file.WriteInt8(log.GoNew.Y)
	file.WriteInt8(log.GoNew.DataId)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.Status)
	file.WriteInt64(log.GoOld.GarrisonStart)
	file.WriteInt64(log.GoOld.GarrisonTime)
	file.WriteInt64(log.GoOld.AwardTime)
	file.WriteInt8(log.GoOld.RoleId)
	file.WriteInt16(log.GoOld.CloudId)
	file.WriteInt8(log.GoOld.X)
	file.WriteInt8(log.GoOld.Y)
	file.WriteInt8(log.GoOld.DataId)
}

func (log *PlayerDrivingSwordEventExploringUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordEventExploring.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Status)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.GarrisonStart)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.GarrisonTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.AwardTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RoleId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CloudId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.X)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Y)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DataId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerDrivingSwordEventExploringUpdateLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordEventExploringUpdateLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordEventExploringUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerDrivingSwordEventExploringInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerDrivingSwordEventExploring
	for crow := log.db.tables.PlayerDrivingSwordEventExploring; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerDrivingSwordEventExploring'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerDrivingSwordEventExploring = log.db.tables.PlayerDrivingSwordEventExploring.next
	}
	C.Free_PlayerDrivingSwordEventExploring(log.Row)
}

func (log *PlayerDrivingSwordEventExploringDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerDrivingSwordEventExploring
	for crow := log.db.tables.PlayerDrivingSwordEventExploring; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_driving_sword_event_exploring'")
	}
	log.Old.next = log.db.tables.PlayerDrivingSwordEventExploring
	log.db.tables.PlayerDrivingSwordEventExploring = log.Old
}

func (log *PlayerDrivingSwordEventExploringUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerDrivingSwordEventExploring
	for crow := log.db.tables.PlayerDrivingSwordEventExploring; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_driving_sword_event_exploring'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerDrivingSwordEventExploring = log.Old
	}
	C.Free_PlayerDrivingSwordEventExploring(log.Row)
}

/* ========== player_driving_sword_event_treasure ========== */

type PlayerDrivingSwordEventTreasureInsertLog struct {
	db *Database 
	Row *C.PlayerDrivingSwordEventTreasure
	GoRow *PlayerDrivingSwordEventTreasure
}

func (db *Database) newPlayerDrivingSwordEventTreasureInsertLog(row *C.PlayerDrivingSwordEventTreasure, goRow *PlayerDrivingSwordEventTreasure) *PlayerDrivingSwordEventTreasureInsertLog {
	return &PlayerDrivingSwordEventTreasureInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerDrivingSwordEventTreasureInsertLog) Free() {
}

func (log *PlayerDrivingSwordEventTreasureInsertLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordEventTreasure != nil {
		g_Hooks.PlayerDrivingSwordEventTreasure.Insert(&PlayerDrivingSwordEventTreasureRow{row: log.Row})
	}
}

func (log *PlayerDrivingSwordEventTreasureInsertLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordEventTreasureInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(39)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.Progress)
	file.WriteInt16(log.GoRow.CloudId)
	file.WriteInt8(log.GoRow.X)
	file.WriteInt8(log.GoRow.Y)
	file.WriteInt8(log.GoRow.DataId)
}

func (log *PlayerDrivingSwordEventTreasureInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordEventTreasure.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Progress)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CloudId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.X)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Y)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DataId)))
	return stmt.Execute()
}

func (log *PlayerDrivingSwordEventTreasureInsertLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordEventTreasureInsertLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordEventTreasureInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDrivingSwordEventTreasureDeleteLog struct {
	db *Database
	Old *C.PlayerDrivingSwordEventTreasure
	GoOld *PlayerDrivingSwordEventTreasure
}

func (db *Database) newPlayerDrivingSwordEventTreasureDeleteLog(old *C.PlayerDrivingSwordEventTreasure, goOld *PlayerDrivingSwordEventTreasure) *PlayerDrivingSwordEventTreasureDeleteLog {
	return &PlayerDrivingSwordEventTreasureDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerDrivingSwordEventTreasureDeleteLog) Free() {
	C.Free_PlayerDrivingSwordEventTreasure(log.Old)
}

func (log *PlayerDrivingSwordEventTreasureDeleteLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordEventTreasure != nil {
		g_Hooks.PlayerDrivingSwordEventTreasure.Delete(&PlayerDrivingSwordEventTreasureRow{row: log.Old})
	}
}

func (log *PlayerDrivingSwordEventTreasureDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordEventTreasureDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(39)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.Progress)
	file.WriteInt16(log.GoOld.CloudId)
	file.WriteInt8(log.GoOld.X)
	file.WriteInt8(log.GoOld.Y)
	file.WriteInt8(log.GoOld.DataId)
}

func (log *PlayerDrivingSwordEventTreasureDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordEventTreasure.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerDrivingSwordEventTreasureDeleteLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordEventTreasureDeleteLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordEventTreasureDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDrivingSwordEventTreasureUpdateLog struct {
	db *Database 
	Row *C.PlayerDrivingSwordEventTreasure
	Old *C.PlayerDrivingSwordEventTreasure
	GoNew *PlayerDrivingSwordEventTreasure
	GoOld *PlayerDrivingSwordEventTreasure
}

func (db *Database) newPlayerDrivingSwordEventTreasureUpdateLog(row, old *C.PlayerDrivingSwordEventTreasure, goNew, goOld *PlayerDrivingSwordEventTreasure) *PlayerDrivingSwordEventTreasureUpdateLog {
	return &PlayerDrivingSwordEventTreasureUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerDrivingSwordEventTreasureUpdateLog) Free() {
	C.Free_PlayerDrivingSwordEventTreasure(log.Old)
}

func (log *PlayerDrivingSwordEventTreasureUpdateLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordEventTreasure != nil {
		g_Hooks.PlayerDrivingSwordEventTreasure.Update(&PlayerDrivingSwordEventTreasureRow{row: log.Row}, &PlayerDrivingSwordEventTreasureRow{row: log.Old})
	}
}

func (log *PlayerDrivingSwordEventTreasureUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordEventTreasureUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(39)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.Progress)
	file.WriteInt16(log.GoNew.CloudId)
	file.WriteInt8(log.GoNew.X)
	file.WriteInt8(log.GoNew.Y)
	file.WriteInt8(log.GoNew.DataId)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.Progress)
	file.WriteInt16(log.GoOld.CloudId)
	file.WriteInt8(log.GoOld.X)
	file.WriteInt8(log.GoOld.Y)
	file.WriteInt8(log.GoOld.DataId)
}

func (log *PlayerDrivingSwordEventTreasureUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordEventTreasure.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Progress)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CloudId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.X)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Y)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DataId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerDrivingSwordEventTreasureUpdateLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordEventTreasureUpdateLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordEventTreasureUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerDrivingSwordEventTreasureInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerDrivingSwordEventTreasure
	for crow := log.db.tables.PlayerDrivingSwordEventTreasure; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerDrivingSwordEventTreasure'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerDrivingSwordEventTreasure = log.db.tables.PlayerDrivingSwordEventTreasure.next
	}
	C.Free_PlayerDrivingSwordEventTreasure(log.Row)
}

func (log *PlayerDrivingSwordEventTreasureDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerDrivingSwordEventTreasure
	for crow := log.db.tables.PlayerDrivingSwordEventTreasure; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_driving_sword_event_treasure'")
	}
	log.Old.next = log.db.tables.PlayerDrivingSwordEventTreasure
	log.db.tables.PlayerDrivingSwordEventTreasure = log.Old
}

func (log *PlayerDrivingSwordEventTreasureUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerDrivingSwordEventTreasure
	for crow := log.db.tables.PlayerDrivingSwordEventTreasure; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_driving_sword_event_treasure'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerDrivingSwordEventTreasure = log.Old
	}
	C.Free_PlayerDrivingSwordEventTreasure(log.Row)
}

/* ========== player_driving_sword_event_visiting ========== */

type PlayerDrivingSwordEventVisitingInsertLog struct {
	db *Database 
	Row *C.PlayerDrivingSwordEventVisiting
	GoRow *PlayerDrivingSwordEventVisiting
}

func (db *Database) newPlayerDrivingSwordEventVisitingInsertLog(row *C.PlayerDrivingSwordEventVisiting, goRow *PlayerDrivingSwordEventVisiting) *PlayerDrivingSwordEventVisitingInsertLog {
	return &PlayerDrivingSwordEventVisitingInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerDrivingSwordEventVisitingInsertLog) Free() {
}

func (log *PlayerDrivingSwordEventVisitingInsertLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordEventVisiting != nil {
		g_Hooks.PlayerDrivingSwordEventVisiting.Insert(&PlayerDrivingSwordEventVisitingRow{row: log.Row})
	}
}

func (log *PlayerDrivingSwordEventVisitingInsertLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordEventVisitingInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(40)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.Status)
	file.WriteInt64(log.GoRow.TargetPid)
	file.WriteBytes(log.GoRow.TargetSideState)
	file.WriteInt16(log.GoRow.CloudId)
	file.WriteInt8(log.GoRow.X)
	file.WriteInt8(log.GoRow.Y)
	file.WriteInt8(log.GoRow.DataId)
	file.WriteString(log.GoRow.TargetStatus)
}

func (log *PlayerDrivingSwordEventVisitingInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordEventVisiting.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Status)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.TargetPid)))
	stmt.BindBinary(unsafe.Pointer(log.Row.TargetSideState), int(log.Row.TargetSideState_len))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CloudId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.X)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Y)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DataId)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.TargetStatus), int(log.Row.TargetStatus_len))
	return stmt.Execute()
}

func (log *PlayerDrivingSwordEventVisitingInsertLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordEventVisitingInsertLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordEventVisitingInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDrivingSwordEventVisitingDeleteLog struct {
	db *Database
	Old *C.PlayerDrivingSwordEventVisiting
	GoOld *PlayerDrivingSwordEventVisiting
}

func (db *Database) newPlayerDrivingSwordEventVisitingDeleteLog(old *C.PlayerDrivingSwordEventVisiting, goOld *PlayerDrivingSwordEventVisiting) *PlayerDrivingSwordEventVisitingDeleteLog {
	return &PlayerDrivingSwordEventVisitingDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerDrivingSwordEventVisitingDeleteLog) Free() {
	C.Free_PlayerDrivingSwordEventVisiting(log.Old)
}

func (log *PlayerDrivingSwordEventVisitingDeleteLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordEventVisiting != nil {
		g_Hooks.PlayerDrivingSwordEventVisiting.Delete(&PlayerDrivingSwordEventVisitingRow{row: log.Old})
	}
}

func (log *PlayerDrivingSwordEventVisitingDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordEventVisitingDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(40)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.Status)
	file.WriteInt64(log.GoOld.TargetPid)
	file.WriteBytes(log.GoOld.TargetSideState)
	file.WriteInt16(log.GoOld.CloudId)
	file.WriteInt8(log.GoOld.X)
	file.WriteInt8(log.GoOld.Y)
	file.WriteInt8(log.GoOld.DataId)
	file.WriteString(log.GoOld.TargetStatus)
}

func (log *PlayerDrivingSwordEventVisitingDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordEventVisiting.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerDrivingSwordEventVisitingDeleteLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordEventVisitingDeleteLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordEventVisitingDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDrivingSwordEventVisitingUpdateLog struct {
	db *Database 
	Row *C.PlayerDrivingSwordEventVisiting
	Old *C.PlayerDrivingSwordEventVisiting
	GoNew *PlayerDrivingSwordEventVisiting
	GoOld *PlayerDrivingSwordEventVisiting
}

func (db *Database) newPlayerDrivingSwordEventVisitingUpdateLog(row, old *C.PlayerDrivingSwordEventVisiting, goNew, goOld *PlayerDrivingSwordEventVisiting) *PlayerDrivingSwordEventVisitingUpdateLog {
	return &PlayerDrivingSwordEventVisitingUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerDrivingSwordEventVisitingUpdateLog) Free() {
	C.Free_PlayerDrivingSwordEventVisiting(log.Old)
}

func (log *PlayerDrivingSwordEventVisitingUpdateLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordEventVisiting != nil {
		g_Hooks.PlayerDrivingSwordEventVisiting.Update(&PlayerDrivingSwordEventVisitingRow{row: log.Row}, &PlayerDrivingSwordEventVisitingRow{row: log.Old})
	}
}

func (log *PlayerDrivingSwordEventVisitingUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordEventVisitingUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(40)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.Status)
	file.WriteInt64(log.GoNew.TargetPid)
	file.WriteBytes(log.GoNew.TargetSideState)
	file.WriteInt16(log.GoNew.CloudId)
	file.WriteInt8(log.GoNew.X)
	file.WriteInt8(log.GoNew.Y)
	file.WriteInt8(log.GoNew.DataId)
	file.WriteString(log.GoNew.TargetStatus)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.Status)
	file.WriteInt64(log.GoOld.TargetPid)
	file.WriteBytes(log.GoOld.TargetSideState)
	file.WriteInt16(log.GoOld.CloudId)
	file.WriteInt8(log.GoOld.X)
	file.WriteInt8(log.GoOld.Y)
	file.WriteInt8(log.GoOld.DataId)
	file.WriteString(log.GoOld.TargetStatus)
}

func (log *PlayerDrivingSwordEventVisitingUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordEventVisiting.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Status)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.TargetPid)))
	stmt.BindBinary(unsafe.Pointer(log.Row.TargetSideState), int(log.Row.TargetSideState_len))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CloudId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.X)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Y)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DataId)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.TargetStatus), int(log.Row.TargetStatus_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerDrivingSwordEventVisitingUpdateLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordEventVisitingUpdateLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordEventVisitingUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerDrivingSwordEventVisitingInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerDrivingSwordEventVisiting
	for crow := log.db.tables.PlayerDrivingSwordEventVisiting; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerDrivingSwordEventVisiting'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerDrivingSwordEventVisiting = log.db.tables.PlayerDrivingSwordEventVisiting.next
	}
	C.Free_PlayerDrivingSwordEventVisiting(log.Row)
}

func (log *PlayerDrivingSwordEventVisitingDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerDrivingSwordEventVisiting
	for crow := log.db.tables.PlayerDrivingSwordEventVisiting; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_driving_sword_event_visiting'")
	}
	log.Old.next = log.db.tables.PlayerDrivingSwordEventVisiting
	log.db.tables.PlayerDrivingSwordEventVisiting = log.Old
}

func (log *PlayerDrivingSwordEventVisitingUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerDrivingSwordEventVisiting
	for crow := log.db.tables.PlayerDrivingSwordEventVisiting; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_driving_sword_event_visiting'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerDrivingSwordEventVisiting = log.Old
	}
	C.Free_PlayerDrivingSwordEventVisiting(log.Row)
}

/* ========== player_driving_sword_eventmask ========== */

type PlayerDrivingSwordEventmaskInsertLog struct {
	db *Database 
	Row *C.PlayerDrivingSwordEventmask
	GoRow *PlayerDrivingSwordEventmask
}

func (db *Database) newPlayerDrivingSwordEventmaskInsertLog(row *C.PlayerDrivingSwordEventmask, goRow *PlayerDrivingSwordEventmask) *PlayerDrivingSwordEventmaskInsertLog {
	return &PlayerDrivingSwordEventmaskInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerDrivingSwordEventmaskInsertLog) Free() {
}

func (log *PlayerDrivingSwordEventmaskInsertLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordEventmask != nil {
		g_Hooks.PlayerDrivingSwordEventmask.Insert(&PlayerDrivingSwordEventmaskRow{row: log.Row})
	}
}

func (log *PlayerDrivingSwordEventmaskInsertLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordEventmaskInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(41)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.CloudId)
	file.WriteBytes(log.GoRow.Events)
}

func (log *PlayerDrivingSwordEventmaskInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordEventmask.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CloudId)))
	stmt.BindBinary(unsafe.Pointer(log.Row.Events), int(log.Row.Events_len))
	return stmt.Execute()
}

func (log *PlayerDrivingSwordEventmaskInsertLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordEventmaskInsertLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordEventmaskInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDrivingSwordEventmaskDeleteLog struct {
	db *Database
	Old *C.PlayerDrivingSwordEventmask
	GoOld *PlayerDrivingSwordEventmask
}

func (db *Database) newPlayerDrivingSwordEventmaskDeleteLog(old *C.PlayerDrivingSwordEventmask, goOld *PlayerDrivingSwordEventmask) *PlayerDrivingSwordEventmaskDeleteLog {
	return &PlayerDrivingSwordEventmaskDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerDrivingSwordEventmaskDeleteLog) Free() {
	C.Free_PlayerDrivingSwordEventmask(log.Old)
}

func (log *PlayerDrivingSwordEventmaskDeleteLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordEventmask != nil {
		g_Hooks.PlayerDrivingSwordEventmask.Delete(&PlayerDrivingSwordEventmaskRow{row: log.Old})
	}
}

func (log *PlayerDrivingSwordEventmaskDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordEventmaskDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(41)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.CloudId)
	file.WriteBytes(log.GoOld.Events)
}

func (log *PlayerDrivingSwordEventmaskDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordEventmask.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerDrivingSwordEventmaskDeleteLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordEventmaskDeleteLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordEventmaskDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDrivingSwordEventmaskUpdateLog struct {
	db *Database 
	Row *C.PlayerDrivingSwordEventmask
	Old *C.PlayerDrivingSwordEventmask
	GoNew *PlayerDrivingSwordEventmask
	GoOld *PlayerDrivingSwordEventmask
}

func (db *Database) newPlayerDrivingSwordEventmaskUpdateLog(row, old *C.PlayerDrivingSwordEventmask, goNew, goOld *PlayerDrivingSwordEventmask) *PlayerDrivingSwordEventmaskUpdateLog {
	return &PlayerDrivingSwordEventmaskUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerDrivingSwordEventmaskUpdateLog) Free() {
	C.Free_PlayerDrivingSwordEventmask(log.Old)
}

func (log *PlayerDrivingSwordEventmaskUpdateLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordEventmask != nil {
		g_Hooks.PlayerDrivingSwordEventmask.Update(&PlayerDrivingSwordEventmaskRow{row: log.Row}, &PlayerDrivingSwordEventmaskRow{row: log.Old})
	}
}

func (log *PlayerDrivingSwordEventmaskUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordEventmaskUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(41)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.CloudId)
	file.WriteBytes(log.GoNew.Events)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.CloudId)
	file.WriteBytes(log.GoOld.Events)
}

func (log *PlayerDrivingSwordEventmaskUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordEventmask.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CloudId)))
	stmt.BindBinary(unsafe.Pointer(log.Row.Events), int(log.Row.Events_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerDrivingSwordEventmaskUpdateLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordEventmaskUpdateLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordEventmaskUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerDrivingSwordEventmaskInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerDrivingSwordEventmask
	for crow := log.db.tables.PlayerDrivingSwordEventmask; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerDrivingSwordEventmask'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerDrivingSwordEventmask = log.db.tables.PlayerDrivingSwordEventmask.next
	}
	C.Free_PlayerDrivingSwordEventmask(log.Row)
}

func (log *PlayerDrivingSwordEventmaskDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerDrivingSwordEventmask
	for crow := log.db.tables.PlayerDrivingSwordEventmask; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_driving_sword_eventmask'")
	}
	log.Old.next = log.db.tables.PlayerDrivingSwordEventmask
	log.db.tables.PlayerDrivingSwordEventmask = log.Old
}

func (log *PlayerDrivingSwordEventmaskUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerDrivingSwordEventmask
	for crow := log.db.tables.PlayerDrivingSwordEventmask; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_driving_sword_eventmask'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerDrivingSwordEventmask = log.Old
	}
	C.Free_PlayerDrivingSwordEventmask(log.Row)
}

/* ========== player_driving_sword_info ========== */

type PlayerDrivingSwordInfoInsertLog struct {
	db *Database 
	Row *C.PlayerDrivingSwordInfo
	GoRow *PlayerDrivingSwordInfo
}

func (db *Database) newPlayerDrivingSwordInfoInsertLog(row *C.PlayerDrivingSwordInfo, goRow *PlayerDrivingSwordInfo) *PlayerDrivingSwordInfoInsertLog {
	return &PlayerDrivingSwordInfoInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerDrivingSwordInfoInsertLog) Free() {
}

func (log *PlayerDrivingSwordInfoInsertLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordInfo != nil {
		g_Hooks.PlayerDrivingSwordInfo.Insert(&PlayerDrivingSwordInfoRow{row: log.Row})
	}
}

func (log *PlayerDrivingSwordInfoInsertLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordInfoInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(42)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.CurrentCloud)
	file.WriteInt16(log.GoRow.HighestCloud)
	file.WriteInt8(log.GoRow.CurrentX)
	file.WriteInt8(log.GoRow.CurrentY)
	file.WriteInt16(log.GoRow.AllowedAction)
	file.WriteInt64(log.GoRow.ActionRefreshTime)
	file.WriteInt64(log.GoRow.ActionBuyTime)
	file.WriteInt8(log.GoRow.DailyActionBought)
}

func (log *PlayerDrivingSwordInfoInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordInfo.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CurrentCloud)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.HighestCloud)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.CurrentX)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.CurrentY)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AllowedAction)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ActionRefreshTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ActionBuyTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DailyActionBought)))
	return stmt.Execute()
}

func (log *PlayerDrivingSwordInfoInsertLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordInfoInsertLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordInfoInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDrivingSwordInfoDeleteLog struct {
	db *Database
	Old *C.PlayerDrivingSwordInfo
	GoOld *PlayerDrivingSwordInfo
}

func (db *Database) newPlayerDrivingSwordInfoDeleteLog(old *C.PlayerDrivingSwordInfo, goOld *PlayerDrivingSwordInfo) *PlayerDrivingSwordInfoDeleteLog {
	return &PlayerDrivingSwordInfoDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerDrivingSwordInfoDeleteLog) Free() {
	C.Free_PlayerDrivingSwordInfo(log.Old)
}

func (log *PlayerDrivingSwordInfoDeleteLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordInfo != nil {
		g_Hooks.PlayerDrivingSwordInfo.Delete(&PlayerDrivingSwordInfoRow{row: log.Old})
	}
}

func (log *PlayerDrivingSwordInfoDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordInfoDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(42)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.CurrentCloud)
	file.WriteInt16(log.GoOld.HighestCloud)
	file.WriteInt8(log.GoOld.CurrentX)
	file.WriteInt8(log.GoOld.CurrentY)
	file.WriteInt16(log.GoOld.AllowedAction)
	file.WriteInt64(log.GoOld.ActionRefreshTime)
	file.WriteInt64(log.GoOld.ActionBuyTime)
	file.WriteInt8(log.GoOld.DailyActionBought)
}

func (log *PlayerDrivingSwordInfoDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordInfo.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerDrivingSwordInfoDeleteLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordInfoDeleteLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordInfoDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDrivingSwordInfoUpdateLog struct {
	db *Database 
	Row *C.PlayerDrivingSwordInfo
	Old *C.PlayerDrivingSwordInfo
	GoNew *PlayerDrivingSwordInfo
	GoOld *PlayerDrivingSwordInfo
}

func (db *Database) newPlayerDrivingSwordInfoUpdateLog(row, old *C.PlayerDrivingSwordInfo, goNew, goOld *PlayerDrivingSwordInfo) *PlayerDrivingSwordInfoUpdateLog {
	return &PlayerDrivingSwordInfoUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerDrivingSwordInfoUpdateLog) Free() {
	C.Free_PlayerDrivingSwordInfo(log.Old)
}

func (log *PlayerDrivingSwordInfoUpdateLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordInfo != nil {
		g_Hooks.PlayerDrivingSwordInfo.Update(&PlayerDrivingSwordInfoRow{row: log.Row}, &PlayerDrivingSwordInfoRow{row: log.Old})
	}
}

func (log *PlayerDrivingSwordInfoUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordInfoUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(42)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.CurrentCloud)
	file.WriteInt16(log.GoNew.HighestCloud)
	file.WriteInt8(log.GoNew.CurrentX)
	file.WriteInt8(log.GoNew.CurrentY)
	file.WriteInt16(log.GoNew.AllowedAction)
	file.WriteInt64(log.GoNew.ActionRefreshTime)
	file.WriteInt64(log.GoNew.ActionBuyTime)
	file.WriteInt8(log.GoNew.DailyActionBought)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.CurrentCloud)
	file.WriteInt16(log.GoOld.HighestCloud)
	file.WriteInt8(log.GoOld.CurrentX)
	file.WriteInt8(log.GoOld.CurrentY)
	file.WriteInt16(log.GoOld.AllowedAction)
	file.WriteInt64(log.GoOld.ActionRefreshTime)
	file.WriteInt64(log.GoOld.ActionBuyTime)
	file.WriteInt8(log.GoOld.DailyActionBought)
}

func (log *PlayerDrivingSwordInfoUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordInfo.Update
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CurrentCloud)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.HighestCloud)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.CurrentX)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.CurrentY)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AllowedAction)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ActionRefreshTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ActionBuyTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DailyActionBought)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerDrivingSwordInfoUpdateLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordInfoUpdateLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordInfoUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerDrivingSwordInfoInsertLog) Rollback() {
	if log.db.tables.PlayerDrivingSwordInfo != log.Row {
		panic("insert rollback check failed 'PlayerDrivingSwordInfo'")
	}
	log.db.tables.PlayerDrivingSwordInfo = nil
	C.Free_PlayerDrivingSwordInfo(log.Row)
}

func (log *PlayerDrivingSwordInfoDeleteLog) Rollback() {
	if log.db.tables.PlayerDrivingSwordInfo != nil {
		panic("delete rollback check failed 'player_driving_sword_info'")
	}
	log.db.tables.PlayerDrivingSwordInfo = log.Old
}

func (log *PlayerDrivingSwordInfoUpdateLog) Rollback() {
	if log.db.tables.PlayerDrivingSwordInfo != log.Row {
		panic("update rollback check failed 'player_driving_sword_info'")
	}
	log.db.tables.PlayerDrivingSwordInfo = log.Old
	C.Free_PlayerDrivingSwordInfo(log.Row)
}

/* ========== player_driving_sword_map ========== */

type PlayerDrivingSwordMapInsertLog struct {
	db *Database 
	Row *C.PlayerDrivingSwordMap
	GoRow *PlayerDrivingSwordMap
}

func (db *Database) newPlayerDrivingSwordMapInsertLog(row *C.PlayerDrivingSwordMap, goRow *PlayerDrivingSwordMap) *PlayerDrivingSwordMapInsertLog {
	return &PlayerDrivingSwordMapInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerDrivingSwordMapInsertLog) Free() {
}

func (log *PlayerDrivingSwordMapInsertLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordMap != nil {
		g_Hooks.PlayerDrivingSwordMap.Insert(&PlayerDrivingSwordMapRow{row: log.Row})
	}
}

func (log *PlayerDrivingSwordMapInsertLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordMapInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(43)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.CloudId)
	file.WriteBytes(log.GoRow.Shadows)
	file.WriteInt8(log.GoRow.ObstacleCount)
	file.WriteInt8(log.GoRow.ExploringCount)
	file.WriteInt8(log.GoRow.VisitingCount)
	file.WriteInt8(log.GoRow.TreasureCount)
}

func (log *PlayerDrivingSwordMapInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordMap.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CloudId)))
	stmt.BindBinary(unsafe.Pointer(log.Row.Shadows), int(log.Row.Shadows_len))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.ObstacleCount)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.ExploringCount)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.VisitingCount)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.TreasureCount)))
	return stmt.Execute()
}

func (log *PlayerDrivingSwordMapInsertLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordMapInsertLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordMapInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDrivingSwordMapDeleteLog struct {
	db *Database
	Old *C.PlayerDrivingSwordMap
	GoOld *PlayerDrivingSwordMap
}

func (db *Database) newPlayerDrivingSwordMapDeleteLog(old *C.PlayerDrivingSwordMap, goOld *PlayerDrivingSwordMap) *PlayerDrivingSwordMapDeleteLog {
	return &PlayerDrivingSwordMapDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerDrivingSwordMapDeleteLog) Free() {
	C.Free_PlayerDrivingSwordMap(log.Old)
}

func (log *PlayerDrivingSwordMapDeleteLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordMap != nil {
		g_Hooks.PlayerDrivingSwordMap.Delete(&PlayerDrivingSwordMapRow{row: log.Old})
	}
}

func (log *PlayerDrivingSwordMapDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordMapDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(43)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.CloudId)
	file.WriteBytes(log.GoOld.Shadows)
	file.WriteInt8(log.GoOld.ObstacleCount)
	file.WriteInt8(log.GoOld.ExploringCount)
	file.WriteInt8(log.GoOld.VisitingCount)
	file.WriteInt8(log.GoOld.TreasureCount)
}

func (log *PlayerDrivingSwordMapDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordMap.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerDrivingSwordMapDeleteLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordMapDeleteLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordMapDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerDrivingSwordMapUpdateLog struct {
	db *Database 
	Row *C.PlayerDrivingSwordMap
	Old *C.PlayerDrivingSwordMap
	GoNew *PlayerDrivingSwordMap
	GoOld *PlayerDrivingSwordMap
}

func (db *Database) newPlayerDrivingSwordMapUpdateLog(row, old *C.PlayerDrivingSwordMap, goNew, goOld *PlayerDrivingSwordMap) *PlayerDrivingSwordMapUpdateLog {
	return &PlayerDrivingSwordMapUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerDrivingSwordMapUpdateLog) Free() {
	C.Free_PlayerDrivingSwordMap(log.Old)
}

func (log *PlayerDrivingSwordMapUpdateLog) InvokeHook() {
	if g_Hooks.PlayerDrivingSwordMap != nil {
		g_Hooks.PlayerDrivingSwordMap.Update(&PlayerDrivingSwordMapRow{row: log.Row}, &PlayerDrivingSwordMapRow{row: log.Old})
	}
}

func (log *PlayerDrivingSwordMapUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerDrivingSwordMapUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(43)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.CloudId)
	file.WriteBytes(log.GoNew.Shadows)
	file.WriteInt8(log.GoNew.ObstacleCount)
	file.WriteInt8(log.GoNew.ExploringCount)
	file.WriteInt8(log.GoNew.VisitingCount)
	file.WriteInt8(log.GoNew.TreasureCount)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.CloudId)
	file.WriteBytes(log.GoOld.Shadows)
	file.WriteInt8(log.GoOld.ObstacleCount)
	file.WriteInt8(log.GoOld.ExploringCount)
	file.WriteInt8(log.GoOld.VisitingCount)
	file.WriteInt8(log.GoOld.TreasureCount)
}

func (log *PlayerDrivingSwordMapUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerDrivingSwordMap.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CloudId)))
	stmt.BindBinary(unsafe.Pointer(log.Row.Shadows), int(log.Row.Shadows_len))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.ObstacleCount)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.ExploringCount)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.VisitingCount)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.TreasureCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerDrivingSwordMapUpdateLog) CommitToTLog() {
}

func (log *PlayerDrivingSwordMapUpdateLog) CommitToXdLog() {
}

func (log *PlayerDrivingSwordMapUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerDrivingSwordMapInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerDrivingSwordMap
	for crow := log.db.tables.PlayerDrivingSwordMap; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerDrivingSwordMap'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerDrivingSwordMap = log.db.tables.PlayerDrivingSwordMap.next
	}
	C.Free_PlayerDrivingSwordMap(log.Row)
}

func (log *PlayerDrivingSwordMapDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerDrivingSwordMap
	for crow := log.db.tables.PlayerDrivingSwordMap; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_driving_sword_map'")
	}
	log.Old.next = log.db.tables.PlayerDrivingSwordMap
	log.db.tables.PlayerDrivingSwordMap = log.Old
}

func (log *PlayerDrivingSwordMapUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerDrivingSwordMap
	for crow := log.db.tables.PlayerDrivingSwordMap; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_driving_sword_map'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerDrivingSwordMap = log.Old
	}
	C.Free_PlayerDrivingSwordMap(log.Row)
}

/* ========== player_equipment ========== */

type PlayerEquipmentInsertLog struct {
	db *Database 
	Row *C.PlayerEquipment
	GoRow *PlayerEquipment
}

func (db *Database) newPlayerEquipmentInsertLog(row *C.PlayerEquipment, goRow *PlayerEquipment) *PlayerEquipmentInsertLog {
	return &PlayerEquipmentInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerEquipmentInsertLog) Free() {
}

func (log *PlayerEquipmentInsertLog) InvokeHook() {
	if g_Hooks.PlayerEquipment != nil {
		g_Hooks.PlayerEquipment.Insert(&PlayerEquipmentRow{row: log.Row})
	}
}

func (log *PlayerEquipmentInsertLog) SetEventId(time.Time) {
}

func (log *PlayerEquipmentInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(44)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.RoleId)
	file.WriteInt64(log.GoRow.WeaponId)
	file.WriteInt64(log.GoRow.ClothesId)
	file.WriteInt64(log.GoRow.AccessoriesId)
	file.WriteInt64(log.GoRow.ShoeId)
}

func (log *PlayerEquipmentInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEquipment.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RoleId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WeaponId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ClothesId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.AccessoriesId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ShoeId)))
	return stmt.Execute()
}

func (log *PlayerEquipmentInsertLog) CommitToTLog() {
}

func (log *PlayerEquipmentInsertLog) CommitToXdLog() {
}

func (log *PlayerEquipmentInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerEquipmentDeleteLog struct {
	db *Database
	Old *C.PlayerEquipment
	GoOld *PlayerEquipment
}

func (db *Database) newPlayerEquipmentDeleteLog(old *C.PlayerEquipment, goOld *PlayerEquipment) *PlayerEquipmentDeleteLog {
	return &PlayerEquipmentDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerEquipmentDeleteLog) Free() {
	C.Free_PlayerEquipment(log.Old)
}

func (log *PlayerEquipmentDeleteLog) InvokeHook() {
	if g_Hooks.PlayerEquipment != nil {
		g_Hooks.PlayerEquipment.Delete(&PlayerEquipmentRow{row: log.Old})
	}
}

func (log *PlayerEquipmentDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerEquipmentDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(44)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.RoleId)
	file.WriteInt64(log.GoOld.WeaponId)
	file.WriteInt64(log.GoOld.ClothesId)
	file.WriteInt64(log.GoOld.AccessoriesId)
	file.WriteInt64(log.GoOld.ShoeId)
}

func (log *PlayerEquipmentDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEquipment.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerEquipmentDeleteLog) CommitToTLog() {
}

func (log *PlayerEquipmentDeleteLog) CommitToXdLog() {
}

func (log *PlayerEquipmentDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerEquipmentUpdateLog struct {
	db *Database 
	Row *C.PlayerEquipment
	Old *C.PlayerEquipment
	GoNew *PlayerEquipment
	GoOld *PlayerEquipment
}

func (db *Database) newPlayerEquipmentUpdateLog(row, old *C.PlayerEquipment, goNew, goOld *PlayerEquipment) *PlayerEquipmentUpdateLog {
	return &PlayerEquipmentUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerEquipmentUpdateLog) Free() {
	C.Free_PlayerEquipment(log.Old)
}

func (log *PlayerEquipmentUpdateLog) InvokeHook() {
	if g_Hooks.PlayerEquipment != nil {
		g_Hooks.PlayerEquipment.Update(&PlayerEquipmentRow{row: log.Row}, &PlayerEquipmentRow{row: log.Old})
	}
}

func (log *PlayerEquipmentUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerEquipmentUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(44)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.RoleId)
	file.WriteInt64(log.GoNew.WeaponId)
	file.WriteInt64(log.GoNew.ClothesId)
	file.WriteInt64(log.GoNew.AccessoriesId)
	file.WriteInt64(log.GoNew.ShoeId)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.RoleId)
	file.WriteInt64(log.GoOld.WeaponId)
	file.WriteInt64(log.GoOld.ClothesId)
	file.WriteInt64(log.GoOld.AccessoriesId)
	file.WriteInt64(log.GoOld.ShoeId)
}

func (log *PlayerEquipmentUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEquipment.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RoleId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WeaponId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ClothesId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.AccessoriesId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ShoeId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerEquipmentUpdateLog) CommitToTLog() {
}

func (log *PlayerEquipmentUpdateLog) CommitToXdLog() {
}

func (log *PlayerEquipmentUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerEquipmentInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerEquipment
	for crow := log.db.tables.PlayerEquipment; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerEquipment'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerEquipment = log.db.tables.PlayerEquipment.next
	}
	C.Free_PlayerEquipment(log.Row)
}

func (log *PlayerEquipmentDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerEquipment
	for crow := log.db.tables.PlayerEquipment; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_equipment'")
	}
	log.Old.next = log.db.tables.PlayerEquipment
	log.db.tables.PlayerEquipment = log.Old
}

func (log *PlayerEquipmentUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerEquipment
	for crow := log.db.tables.PlayerEquipment; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_equipment'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerEquipment = log.Old
	}
	C.Free_PlayerEquipment(log.Row)
}

/* ========== player_event_award_record ========== */

type PlayerEventAwardRecordInsertLog struct {
	db *Database 
	Row *C.PlayerEventAwardRecord
	GoRow *PlayerEventAwardRecord
}

func (db *Database) newPlayerEventAwardRecordInsertLog(row *C.PlayerEventAwardRecord, goRow *PlayerEventAwardRecord) *PlayerEventAwardRecordInsertLog {
	return &PlayerEventAwardRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerEventAwardRecordInsertLog) Free() {
}

func (log *PlayerEventAwardRecordInsertLog) InvokeHook() {
	if g_Hooks.PlayerEventAwardRecord != nil {
		g_Hooks.PlayerEventAwardRecord.Insert(&PlayerEventAwardRecordRow{row: log.Row})
	}
}

func (log *PlayerEventAwardRecordInsertLog) SetEventId(time.Time) {
}

func (log *PlayerEventAwardRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(45)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteBytes(log.GoRow.RecordBytes)
	file.WriteBytes(log.GoRow.JsonEventRecord)
}

func (log *PlayerEventAwardRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEventAwardRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBinary(unsafe.Pointer(log.Row.RecordBytes), int(log.Row.RecordBytes_len))
	stmt.BindBinary(unsafe.Pointer(log.Row.JsonEventRecord), int(log.Row.JsonEventRecord_len))
	return stmt.Execute()
}

func (log *PlayerEventAwardRecordInsertLog) CommitToTLog() {
}

func (log *PlayerEventAwardRecordInsertLog) CommitToXdLog() {
}

func (log *PlayerEventAwardRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerEventAwardRecordDeleteLog struct {
	db *Database
	Old *C.PlayerEventAwardRecord
	GoOld *PlayerEventAwardRecord
}

func (db *Database) newPlayerEventAwardRecordDeleteLog(old *C.PlayerEventAwardRecord, goOld *PlayerEventAwardRecord) *PlayerEventAwardRecordDeleteLog {
	return &PlayerEventAwardRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerEventAwardRecordDeleteLog) Free() {
	C.Free_PlayerEventAwardRecord(log.Old)
}

func (log *PlayerEventAwardRecordDeleteLog) InvokeHook() {
	if g_Hooks.PlayerEventAwardRecord != nil {
		g_Hooks.PlayerEventAwardRecord.Delete(&PlayerEventAwardRecordRow{row: log.Old})
	}
}

func (log *PlayerEventAwardRecordDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerEventAwardRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(45)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteBytes(log.GoOld.RecordBytes)
	file.WriteBytes(log.GoOld.JsonEventRecord)
}

func (log *PlayerEventAwardRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEventAwardRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerEventAwardRecordDeleteLog) CommitToTLog() {
}

func (log *PlayerEventAwardRecordDeleteLog) CommitToXdLog() {
}

func (log *PlayerEventAwardRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerEventAwardRecordUpdateLog struct {
	db *Database 
	Row *C.PlayerEventAwardRecord
	Old *C.PlayerEventAwardRecord
	GoNew *PlayerEventAwardRecord
	GoOld *PlayerEventAwardRecord
}

func (db *Database) newPlayerEventAwardRecordUpdateLog(row, old *C.PlayerEventAwardRecord, goNew, goOld *PlayerEventAwardRecord) *PlayerEventAwardRecordUpdateLog {
	return &PlayerEventAwardRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerEventAwardRecordUpdateLog) Free() {
	C.Free_PlayerEventAwardRecord(log.Old)
}

func (log *PlayerEventAwardRecordUpdateLog) InvokeHook() {
	if g_Hooks.PlayerEventAwardRecord != nil {
		g_Hooks.PlayerEventAwardRecord.Update(&PlayerEventAwardRecordRow{row: log.Row}, &PlayerEventAwardRecordRow{row: log.Old})
	}
}

func (log *PlayerEventAwardRecordUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerEventAwardRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(45)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteBytes(log.GoNew.RecordBytes)
	file.WriteBytes(log.GoNew.JsonEventRecord)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteBytes(log.GoOld.RecordBytes)
	file.WriteBytes(log.GoOld.JsonEventRecord)
}

func (log *PlayerEventAwardRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEventAwardRecord.Update
	stmt.BindBinary(unsafe.Pointer(log.Row.RecordBytes), int(log.Row.RecordBytes_len))
	stmt.BindBinary(unsafe.Pointer(log.Row.JsonEventRecord), int(log.Row.JsonEventRecord_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerEventAwardRecordUpdateLog) CommitToTLog() {
}

func (log *PlayerEventAwardRecordUpdateLog) CommitToXdLog() {
}

func (log *PlayerEventAwardRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerEventAwardRecordInsertLog) Rollback() {
	if log.db.tables.PlayerEventAwardRecord != log.Row {
		panic("insert rollback check failed 'PlayerEventAwardRecord'")
	}
	log.db.tables.PlayerEventAwardRecord = nil
	C.Free_PlayerEventAwardRecord(log.Row)
}

func (log *PlayerEventAwardRecordDeleteLog) Rollback() {
	if log.db.tables.PlayerEventAwardRecord != nil {
		panic("delete rollback check failed 'player_event_award_record'")
	}
	log.db.tables.PlayerEventAwardRecord = log.Old
}

func (log *PlayerEventAwardRecordUpdateLog) Rollback() {
	if log.db.tables.PlayerEventAwardRecord != log.Row {
		panic("update rollback check failed 'player_event_award_record'")
	}
	log.db.tables.PlayerEventAwardRecord = log.Old
	C.Free_PlayerEventAwardRecord(log.Row)
}

/* ========== player_event_daily_online ========== */

type PlayerEventDailyOnlineInsertLog struct {
	db *Database 
	Row *C.PlayerEventDailyOnline
	GoRow *PlayerEventDailyOnline
}

func (db *Database) newPlayerEventDailyOnlineInsertLog(row *C.PlayerEventDailyOnline, goRow *PlayerEventDailyOnline) *PlayerEventDailyOnlineInsertLog {
	return &PlayerEventDailyOnlineInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerEventDailyOnlineInsertLog) Free() {
}

func (log *PlayerEventDailyOnlineInsertLog) InvokeHook() {
	if g_Hooks.PlayerEventDailyOnline != nil {
		g_Hooks.PlayerEventDailyOnline.Insert(&PlayerEventDailyOnlineRow{row: log.Row})
	}
}

func (log *PlayerEventDailyOnlineInsertLog) SetEventId(time.Time) {
}

func (log *PlayerEventDailyOnlineInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(46)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.Timestamp)
	file.WriteInt32(log.GoRow.TotalOnlineTime)
	file.WriteInt32(log.GoRow.AwardedOnlinetime)
}

func (log *PlayerEventDailyOnlineInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEventDailyOnline.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.TotalOnlineTime)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.AwardedOnlinetime)))
	return stmt.Execute()
}

func (log *PlayerEventDailyOnlineInsertLog) CommitToTLog() {
}

func (log *PlayerEventDailyOnlineInsertLog) CommitToXdLog() {
}

func (log *PlayerEventDailyOnlineInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerEventDailyOnlineDeleteLog struct {
	db *Database
	Old *C.PlayerEventDailyOnline
	GoOld *PlayerEventDailyOnline
}

func (db *Database) newPlayerEventDailyOnlineDeleteLog(old *C.PlayerEventDailyOnline, goOld *PlayerEventDailyOnline) *PlayerEventDailyOnlineDeleteLog {
	return &PlayerEventDailyOnlineDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerEventDailyOnlineDeleteLog) Free() {
	C.Free_PlayerEventDailyOnline(log.Old)
}

func (log *PlayerEventDailyOnlineDeleteLog) InvokeHook() {
	if g_Hooks.PlayerEventDailyOnline != nil {
		g_Hooks.PlayerEventDailyOnline.Delete(&PlayerEventDailyOnlineRow{row: log.Old})
	}
}

func (log *PlayerEventDailyOnlineDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerEventDailyOnlineDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(46)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt32(log.GoOld.TotalOnlineTime)
	file.WriteInt32(log.GoOld.AwardedOnlinetime)
}

func (log *PlayerEventDailyOnlineDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEventDailyOnline.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerEventDailyOnlineDeleteLog) CommitToTLog() {
}

func (log *PlayerEventDailyOnlineDeleteLog) CommitToXdLog() {
}

func (log *PlayerEventDailyOnlineDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerEventDailyOnlineUpdateLog struct {
	db *Database 
	Row *C.PlayerEventDailyOnline
	Old *C.PlayerEventDailyOnline
	GoNew *PlayerEventDailyOnline
	GoOld *PlayerEventDailyOnline
}

func (db *Database) newPlayerEventDailyOnlineUpdateLog(row, old *C.PlayerEventDailyOnline, goNew, goOld *PlayerEventDailyOnline) *PlayerEventDailyOnlineUpdateLog {
	return &PlayerEventDailyOnlineUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerEventDailyOnlineUpdateLog) Free() {
	C.Free_PlayerEventDailyOnline(log.Old)
}

func (log *PlayerEventDailyOnlineUpdateLog) InvokeHook() {
	if g_Hooks.PlayerEventDailyOnline != nil {
		g_Hooks.PlayerEventDailyOnline.Update(&PlayerEventDailyOnlineRow{row: log.Row}, &PlayerEventDailyOnlineRow{row: log.Old})
	}
}

func (log *PlayerEventDailyOnlineUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerEventDailyOnlineUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(46)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.Timestamp)
	file.WriteInt32(log.GoNew.TotalOnlineTime)
	file.WriteInt32(log.GoNew.AwardedOnlinetime)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt32(log.GoOld.TotalOnlineTime)
	file.WriteInt32(log.GoOld.AwardedOnlinetime)
}

func (log *PlayerEventDailyOnlineUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEventDailyOnline.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.TotalOnlineTime)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.AwardedOnlinetime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerEventDailyOnlineUpdateLog) CommitToTLog() {
}

func (log *PlayerEventDailyOnlineUpdateLog) CommitToXdLog() {
}

func (log *PlayerEventDailyOnlineUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerEventDailyOnlineInsertLog) Rollback() {
	if log.db.tables.PlayerEventDailyOnline != log.Row {
		panic("insert rollback check failed 'PlayerEventDailyOnline'")
	}
	log.db.tables.PlayerEventDailyOnline = nil
	C.Free_PlayerEventDailyOnline(log.Row)
}

func (log *PlayerEventDailyOnlineDeleteLog) Rollback() {
	if log.db.tables.PlayerEventDailyOnline != nil {
		panic("delete rollback check failed 'player_event_daily_online'")
	}
	log.db.tables.PlayerEventDailyOnline = log.Old
}

func (log *PlayerEventDailyOnlineUpdateLog) Rollback() {
	if log.db.tables.PlayerEventDailyOnline != log.Row {
		panic("update rollback check failed 'player_event_daily_online'")
	}
	log.db.tables.PlayerEventDailyOnline = log.Old
	C.Free_PlayerEventDailyOnline(log.Row)
}

/* ========== player_event_vn_daily_recharge ========== */

type PlayerEventVnDailyRechargeInsertLog struct {
	db *Database 
	Row *C.PlayerEventVnDailyRecharge
	GoRow *PlayerEventVnDailyRecharge
}

func (db *Database) newPlayerEventVnDailyRechargeInsertLog(row *C.PlayerEventVnDailyRecharge, goRow *PlayerEventVnDailyRecharge) *PlayerEventVnDailyRechargeInsertLog {
	return &PlayerEventVnDailyRechargeInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerEventVnDailyRechargeInsertLog) Free() {
}

func (log *PlayerEventVnDailyRechargeInsertLog) InvokeHook() {
	if g_Hooks.PlayerEventVnDailyRecharge != nil {
		g_Hooks.PlayerEventVnDailyRecharge.Insert(&PlayerEventVnDailyRechargeRow{row: log.Row})
	}
}

func (log *PlayerEventVnDailyRechargeInsertLog) SetEventId(time.Time) {
}

func (log *PlayerEventVnDailyRechargeInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(47)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.Page)
	file.WriteInt64(log.GoRow.Timestamp)
	file.WriteInt8(log.GoRow.Awarded)
	file.WriteInt64(log.GoRow.DailyRecharge)
	file.WriteInt64(log.GoRow.TotalRecharge)
}

func (log *PlayerEventVnDailyRechargeInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEventVnDailyRecharge.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Page)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Awarded)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DailyRecharge)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.TotalRecharge)))
	return stmt.Execute()
}

func (log *PlayerEventVnDailyRechargeInsertLog) CommitToTLog() {
}

func (log *PlayerEventVnDailyRechargeInsertLog) CommitToXdLog() {
}

func (log *PlayerEventVnDailyRechargeInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerEventVnDailyRechargeDeleteLog struct {
	db *Database
	Old *C.PlayerEventVnDailyRecharge
	GoOld *PlayerEventVnDailyRecharge
}

func (db *Database) newPlayerEventVnDailyRechargeDeleteLog(old *C.PlayerEventVnDailyRecharge, goOld *PlayerEventVnDailyRecharge) *PlayerEventVnDailyRechargeDeleteLog {
	return &PlayerEventVnDailyRechargeDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerEventVnDailyRechargeDeleteLog) Free() {
	C.Free_PlayerEventVnDailyRecharge(log.Old)
}

func (log *PlayerEventVnDailyRechargeDeleteLog) InvokeHook() {
	if g_Hooks.PlayerEventVnDailyRecharge != nil {
		g_Hooks.PlayerEventVnDailyRecharge.Delete(&PlayerEventVnDailyRechargeRow{row: log.Old})
	}
}

func (log *PlayerEventVnDailyRechargeDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerEventVnDailyRechargeDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(47)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Page)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt8(log.GoOld.Awarded)
	file.WriteInt64(log.GoOld.DailyRecharge)
	file.WriteInt64(log.GoOld.TotalRecharge)
}

func (log *PlayerEventVnDailyRechargeDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEventVnDailyRecharge.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerEventVnDailyRechargeDeleteLog) CommitToTLog() {
}

func (log *PlayerEventVnDailyRechargeDeleteLog) CommitToXdLog() {
}

func (log *PlayerEventVnDailyRechargeDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerEventVnDailyRechargeUpdateLog struct {
	db *Database 
	Row *C.PlayerEventVnDailyRecharge
	Old *C.PlayerEventVnDailyRecharge
	GoNew *PlayerEventVnDailyRecharge
	GoOld *PlayerEventVnDailyRecharge
}

func (db *Database) newPlayerEventVnDailyRechargeUpdateLog(row, old *C.PlayerEventVnDailyRecharge, goNew, goOld *PlayerEventVnDailyRecharge) *PlayerEventVnDailyRechargeUpdateLog {
	return &PlayerEventVnDailyRechargeUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerEventVnDailyRechargeUpdateLog) Free() {
	C.Free_PlayerEventVnDailyRecharge(log.Old)
}

func (log *PlayerEventVnDailyRechargeUpdateLog) InvokeHook() {
	if g_Hooks.PlayerEventVnDailyRecharge != nil {
		g_Hooks.PlayerEventVnDailyRecharge.Update(&PlayerEventVnDailyRechargeRow{row: log.Row}, &PlayerEventVnDailyRechargeRow{row: log.Old})
	}
}

func (log *PlayerEventVnDailyRechargeUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerEventVnDailyRechargeUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(47)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.Page)
	file.WriteInt64(log.GoNew.Timestamp)
	file.WriteInt8(log.GoNew.Awarded)
	file.WriteInt64(log.GoNew.DailyRecharge)
	file.WriteInt64(log.GoNew.TotalRecharge)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Page)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt8(log.GoOld.Awarded)
	file.WriteInt64(log.GoOld.DailyRecharge)
	file.WriteInt64(log.GoOld.TotalRecharge)
}

func (log *PlayerEventVnDailyRechargeUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEventVnDailyRecharge.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Page)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Awarded)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DailyRecharge)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.TotalRecharge)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerEventVnDailyRechargeUpdateLog) CommitToTLog() {
}

func (log *PlayerEventVnDailyRechargeUpdateLog) CommitToXdLog() {
}

func (log *PlayerEventVnDailyRechargeUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerEventVnDailyRechargeInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerEventVnDailyRecharge
	for crow := log.db.tables.PlayerEventVnDailyRecharge; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerEventVnDailyRecharge'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerEventVnDailyRecharge = log.db.tables.PlayerEventVnDailyRecharge.next
	}
	C.Free_PlayerEventVnDailyRecharge(log.Row)
}

func (log *PlayerEventVnDailyRechargeDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerEventVnDailyRecharge
	for crow := log.db.tables.PlayerEventVnDailyRecharge; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_event_vn_daily_recharge'")
	}
	log.Old.next = log.db.tables.PlayerEventVnDailyRecharge
	log.db.tables.PlayerEventVnDailyRecharge = log.Old
}

func (log *PlayerEventVnDailyRechargeUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerEventVnDailyRecharge
	for crow := log.db.tables.PlayerEventVnDailyRecharge; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_event_vn_daily_recharge'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerEventVnDailyRecharge = log.Old
	}
	C.Free_PlayerEventVnDailyRecharge(log.Row)
}

/* ========== player_event_vn_dragon_ball_record ========== */

type PlayerEventVnDragonBallRecordInsertLog struct {
	db *Database 
	Row *C.PlayerEventVnDragonBallRecord
	GoRow *PlayerEventVnDragonBallRecord
}

func (db *Database) newPlayerEventVnDragonBallRecordInsertLog(row *C.PlayerEventVnDragonBallRecord, goRow *PlayerEventVnDragonBallRecord) *PlayerEventVnDragonBallRecordInsertLog {
	return &PlayerEventVnDragonBallRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerEventVnDragonBallRecordInsertLog) Free() {
}

func (log *PlayerEventVnDragonBallRecordInsertLog) InvokeHook() {
	if g_Hooks.PlayerEventVnDragonBallRecord != nil {
		g_Hooks.PlayerEventVnDragonBallRecord.Insert(&PlayerEventVnDragonBallRecordRow{row: log.Row})
	}
}

func (log *PlayerEventVnDragonBallRecordInsertLog) SetEventId(time.Time) {
}

func (log *PlayerEventVnDragonBallRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(48)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.ItemId)
	file.WriteInt64(log.GoRow.Createtime)
}

func (log *PlayerEventVnDragonBallRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEventVnDragonBallRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Createtime)))
	return stmt.Execute()
}

func (log *PlayerEventVnDragonBallRecordInsertLog) CommitToTLog() {
}

func (log *PlayerEventVnDragonBallRecordInsertLog) CommitToXdLog() {
}

func (log *PlayerEventVnDragonBallRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerEventVnDragonBallRecordDeleteLog struct {
	db *Database
	Old *C.PlayerEventVnDragonBallRecord
	GoOld *PlayerEventVnDragonBallRecord
}

func (db *Database) newPlayerEventVnDragonBallRecordDeleteLog(old *C.PlayerEventVnDragonBallRecord, goOld *PlayerEventVnDragonBallRecord) *PlayerEventVnDragonBallRecordDeleteLog {
	return &PlayerEventVnDragonBallRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerEventVnDragonBallRecordDeleteLog) Free() {
	C.Free_PlayerEventVnDragonBallRecord(log.Old)
}

func (log *PlayerEventVnDragonBallRecordDeleteLog) InvokeHook() {
	if g_Hooks.PlayerEventVnDragonBallRecord != nil {
		g_Hooks.PlayerEventVnDragonBallRecord.Delete(&PlayerEventVnDragonBallRecordRow{row: log.Old})
	}
}

func (log *PlayerEventVnDragonBallRecordDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerEventVnDragonBallRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(48)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt64(log.GoOld.Createtime)
}

func (log *PlayerEventVnDragonBallRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEventVnDragonBallRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerEventVnDragonBallRecordDeleteLog) CommitToTLog() {
}

func (log *PlayerEventVnDragonBallRecordDeleteLog) CommitToXdLog() {
}

func (log *PlayerEventVnDragonBallRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerEventVnDragonBallRecordUpdateLog struct {
	db *Database 
	Row *C.PlayerEventVnDragonBallRecord
	Old *C.PlayerEventVnDragonBallRecord
	GoNew *PlayerEventVnDragonBallRecord
	GoOld *PlayerEventVnDragonBallRecord
}

func (db *Database) newPlayerEventVnDragonBallRecordUpdateLog(row, old *C.PlayerEventVnDragonBallRecord, goNew, goOld *PlayerEventVnDragonBallRecord) *PlayerEventVnDragonBallRecordUpdateLog {
	return &PlayerEventVnDragonBallRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerEventVnDragonBallRecordUpdateLog) Free() {
	C.Free_PlayerEventVnDragonBallRecord(log.Old)
}

func (log *PlayerEventVnDragonBallRecordUpdateLog) InvokeHook() {
	if g_Hooks.PlayerEventVnDragonBallRecord != nil {
		g_Hooks.PlayerEventVnDragonBallRecord.Update(&PlayerEventVnDragonBallRecordRow{row: log.Row}, &PlayerEventVnDragonBallRecordRow{row: log.Old})
	}
}

func (log *PlayerEventVnDragonBallRecordUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerEventVnDragonBallRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(48)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.ItemId)
	file.WriteInt64(log.GoNew.Createtime)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt64(log.GoOld.Createtime)
}

func (log *PlayerEventVnDragonBallRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEventVnDragonBallRecord.Update
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Createtime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerEventVnDragonBallRecordUpdateLog) CommitToTLog() {
}

func (log *PlayerEventVnDragonBallRecordUpdateLog) CommitToXdLog() {
}

func (log *PlayerEventVnDragonBallRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerEventVnDragonBallRecordInsertLog) Rollback() {
	if log.db.tables.PlayerEventVnDragonBallRecord != log.Row {
		panic("insert rollback check failed 'PlayerEventVnDragonBallRecord'")
	}
	log.db.tables.PlayerEventVnDragonBallRecord = nil
	C.Free_PlayerEventVnDragonBallRecord(log.Row)
}

func (log *PlayerEventVnDragonBallRecordDeleteLog) Rollback() {
	if log.db.tables.PlayerEventVnDragonBallRecord != nil {
		panic("delete rollback check failed 'player_event_vn_dragon_ball_record'")
	}
	log.db.tables.PlayerEventVnDragonBallRecord = log.Old
}

func (log *PlayerEventVnDragonBallRecordUpdateLog) Rollback() {
	if log.db.tables.PlayerEventVnDragonBallRecord != log.Row {
		panic("update rollback check failed 'player_event_vn_dragon_ball_record'")
	}
	log.db.tables.PlayerEventVnDragonBallRecord = log.Old
	C.Free_PlayerEventVnDragonBallRecord(log.Row)
}

/* ========== player_events_fight_power ========== */

type PlayerEventsFightPowerInsertLog struct {
	db *Database 
	Row *C.PlayerEventsFightPower
	GoRow *PlayerEventsFightPower
}

func (db *Database) newPlayerEventsFightPowerInsertLog(row *C.PlayerEventsFightPower, goRow *PlayerEventsFightPower) *PlayerEventsFightPowerInsertLog {
	return &PlayerEventsFightPowerInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerEventsFightPowerInsertLog) Free() {
}

func (log *PlayerEventsFightPowerInsertLog) InvokeHook() {
	if g_Hooks.PlayerEventsFightPower != nil {
		g_Hooks.PlayerEventsFightPower.Insert(&PlayerEventsFightPowerRow{row: log.Row})
	}
}

func (log *PlayerEventsFightPowerInsertLog) SetEventId(time.Time) {
}

func (log *PlayerEventsFightPowerInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(49)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.Lock)
}

func (log *PlayerEventsFightPowerInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEventsFightPower.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Lock)))
	return stmt.Execute()
}

func (log *PlayerEventsFightPowerInsertLog) CommitToTLog() {
}

func (log *PlayerEventsFightPowerInsertLog) CommitToXdLog() {
}

func (log *PlayerEventsFightPowerInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerEventsFightPowerDeleteLog struct {
	db *Database
	Old *C.PlayerEventsFightPower
	GoOld *PlayerEventsFightPower
}

func (db *Database) newPlayerEventsFightPowerDeleteLog(old *C.PlayerEventsFightPower, goOld *PlayerEventsFightPower) *PlayerEventsFightPowerDeleteLog {
	return &PlayerEventsFightPowerDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerEventsFightPowerDeleteLog) Free() {
	C.Free_PlayerEventsFightPower(log.Old)
}

func (log *PlayerEventsFightPowerDeleteLog) InvokeHook() {
	if g_Hooks.PlayerEventsFightPower != nil {
		g_Hooks.PlayerEventsFightPower.Delete(&PlayerEventsFightPowerRow{row: log.Old})
	}
}

func (log *PlayerEventsFightPowerDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerEventsFightPowerDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(49)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Lock)
}

func (log *PlayerEventsFightPowerDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEventsFightPower.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerEventsFightPowerDeleteLog) CommitToTLog() {
}

func (log *PlayerEventsFightPowerDeleteLog) CommitToXdLog() {
}

func (log *PlayerEventsFightPowerDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerEventsFightPowerUpdateLog struct {
	db *Database 
	Row *C.PlayerEventsFightPower
	Old *C.PlayerEventsFightPower
	GoNew *PlayerEventsFightPower
	GoOld *PlayerEventsFightPower
}

func (db *Database) newPlayerEventsFightPowerUpdateLog(row, old *C.PlayerEventsFightPower, goNew, goOld *PlayerEventsFightPower) *PlayerEventsFightPowerUpdateLog {
	return &PlayerEventsFightPowerUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerEventsFightPowerUpdateLog) Free() {
	C.Free_PlayerEventsFightPower(log.Old)
}

func (log *PlayerEventsFightPowerUpdateLog) InvokeHook() {
	if g_Hooks.PlayerEventsFightPower != nil {
		g_Hooks.PlayerEventsFightPower.Update(&PlayerEventsFightPowerRow{row: log.Row}, &PlayerEventsFightPowerRow{row: log.Old})
	}
}

func (log *PlayerEventsFightPowerUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerEventsFightPowerUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(49)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.Lock)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Lock)
}

func (log *PlayerEventsFightPowerUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEventsFightPower.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.Lock)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerEventsFightPowerUpdateLog) CommitToTLog() {
}

func (log *PlayerEventsFightPowerUpdateLog) CommitToXdLog() {
}

func (log *PlayerEventsFightPowerUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerEventsFightPowerInsertLog) Rollback() {
	if log.db.tables.PlayerEventsFightPower != log.Row {
		panic("insert rollback check failed 'PlayerEventsFightPower'")
	}
	log.db.tables.PlayerEventsFightPower = nil
	C.Free_PlayerEventsFightPower(log.Row)
}

func (log *PlayerEventsFightPowerDeleteLog) Rollback() {
	if log.db.tables.PlayerEventsFightPower != nil {
		panic("delete rollback check failed 'player_events_fight_power'")
	}
	log.db.tables.PlayerEventsFightPower = log.Old
}

func (log *PlayerEventsFightPowerUpdateLog) Rollback() {
	if log.db.tables.PlayerEventsFightPower != log.Row {
		panic("update rollback check failed 'player_events_fight_power'")
	}
	log.db.tables.PlayerEventsFightPower = log.Old
	C.Free_PlayerEventsFightPower(log.Row)
}

/* ========== player_events_ingot_record ========== */

type PlayerEventsIngotRecordInsertLog struct {
	db *Database 
	Row *C.PlayerEventsIngotRecord
	GoRow *PlayerEventsIngotRecord
}

func (db *Database) newPlayerEventsIngotRecordInsertLog(row *C.PlayerEventsIngotRecord, goRow *PlayerEventsIngotRecord) *PlayerEventsIngotRecordInsertLog {
	return &PlayerEventsIngotRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerEventsIngotRecordInsertLog) Free() {
}

func (log *PlayerEventsIngotRecordInsertLog) InvokeHook() {
	if g_Hooks.PlayerEventsIngotRecord != nil {
		g_Hooks.PlayerEventsIngotRecord.Insert(&PlayerEventsIngotRecordRow{row: log.Row})
	}
}

func (log *PlayerEventsIngotRecordInsertLog) SetEventId(time.Time) {
}

func (log *PlayerEventsIngotRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(50)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.IngotIn)
	file.WriteInt64(log.GoRow.IngotInEndTime)
	file.WriteInt64(log.GoRow.IngotOut)
	file.WriteInt64(log.GoRow.IngotOutEndTime)
}

func (log *PlayerEventsIngotRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEventsIngotRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.IngotIn)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.IngotInEndTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.IngotOut)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.IngotOutEndTime)))
	return stmt.Execute()
}

func (log *PlayerEventsIngotRecordInsertLog) CommitToTLog() {
}

func (log *PlayerEventsIngotRecordInsertLog) CommitToXdLog() {
}

func (log *PlayerEventsIngotRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerEventsIngotRecordDeleteLog struct {
	db *Database
	Old *C.PlayerEventsIngotRecord
	GoOld *PlayerEventsIngotRecord
}

func (db *Database) newPlayerEventsIngotRecordDeleteLog(old *C.PlayerEventsIngotRecord, goOld *PlayerEventsIngotRecord) *PlayerEventsIngotRecordDeleteLog {
	return &PlayerEventsIngotRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerEventsIngotRecordDeleteLog) Free() {
	C.Free_PlayerEventsIngotRecord(log.Old)
}

func (log *PlayerEventsIngotRecordDeleteLog) InvokeHook() {
	if g_Hooks.PlayerEventsIngotRecord != nil {
		g_Hooks.PlayerEventsIngotRecord.Delete(&PlayerEventsIngotRecordRow{row: log.Old})
	}
}

func (log *PlayerEventsIngotRecordDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerEventsIngotRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(50)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.IngotIn)
	file.WriteInt64(log.GoOld.IngotInEndTime)
	file.WriteInt64(log.GoOld.IngotOut)
	file.WriteInt64(log.GoOld.IngotOutEndTime)
}

func (log *PlayerEventsIngotRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEventsIngotRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerEventsIngotRecordDeleteLog) CommitToTLog() {
}

func (log *PlayerEventsIngotRecordDeleteLog) CommitToXdLog() {
}

func (log *PlayerEventsIngotRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerEventsIngotRecordUpdateLog struct {
	db *Database 
	Row *C.PlayerEventsIngotRecord
	Old *C.PlayerEventsIngotRecord
	GoNew *PlayerEventsIngotRecord
	GoOld *PlayerEventsIngotRecord
}

func (db *Database) newPlayerEventsIngotRecordUpdateLog(row, old *C.PlayerEventsIngotRecord, goNew, goOld *PlayerEventsIngotRecord) *PlayerEventsIngotRecordUpdateLog {
	return &PlayerEventsIngotRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerEventsIngotRecordUpdateLog) Free() {
	C.Free_PlayerEventsIngotRecord(log.Old)
}

func (log *PlayerEventsIngotRecordUpdateLog) InvokeHook() {
	if g_Hooks.PlayerEventsIngotRecord != nil {
		g_Hooks.PlayerEventsIngotRecord.Update(&PlayerEventsIngotRecordRow{row: log.Row}, &PlayerEventsIngotRecordRow{row: log.Old})
	}
}

func (log *PlayerEventsIngotRecordUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerEventsIngotRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(50)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.IngotIn)
	file.WriteInt64(log.GoNew.IngotInEndTime)
	file.WriteInt64(log.GoNew.IngotOut)
	file.WriteInt64(log.GoNew.IngotOutEndTime)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.IngotIn)
	file.WriteInt64(log.GoOld.IngotInEndTime)
	file.WriteInt64(log.GoOld.IngotOut)
	file.WriteInt64(log.GoOld.IngotOutEndTime)
}

func (log *PlayerEventsIngotRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerEventsIngotRecord.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.IngotIn)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.IngotInEndTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.IngotOut)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.IngotOutEndTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerEventsIngotRecordUpdateLog) CommitToTLog() {
}

func (log *PlayerEventsIngotRecordUpdateLog) CommitToXdLog() {
}

func (log *PlayerEventsIngotRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerEventsIngotRecordInsertLog) Rollback() {
	if log.db.tables.PlayerEventsIngotRecord != log.Row {
		panic("insert rollback check failed 'PlayerEventsIngotRecord'")
	}
	log.db.tables.PlayerEventsIngotRecord = nil
	C.Free_PlayerEventsIngotRecord(log.Row)
}

func (log *PlayerEventsIngotRecordDeleteLog) Rollback() {
	if log.db.tables.PlayerEventsIngotRecord != nil {
		panic("delete rollback check failed 'player_events_ingot_record'")
	}
	log.db.tables.PlayerEventsIngotRecord = log.Old
}

func (log *PlayerEventsIngotRecordUpdateLog) Rollback() {
	if log.db.tables.PlayerEventsIngotRecord != log.Row {
		panic("update rollback check failed 'player_events_ingot_record'")
	}
	log.db.tables.PlayerEventsIngotRecord = log.Old
	C.Free_PlayerEventsIngotRecord(log.Row)
}

/* ========== player_extend_level ========== */

type PlayerExtendLevelInsertLog struct {
	db *Database 
	Row *C.PlayerExtendLevel
	GoRow *PlayerExtendLevel
}

func (db *Database) newPlayerExtendLevelInsertLog(row *C.PlayerExtendLevel, goRow *PlayerExtendLevel) *PlayerExtendLevelInsertLog {
	return &PlayerExtendLevelInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerExtendLevelInsertLog) Free() {
}

func (log *PlayerExtendLevelInsertLog) InvokeHook() {
	if g_Hooks.PlayerExtendLevel != nil {
		g_Hooks.PlayerExtendLevel.Insert(&PlayerExtendLevelRow{row: log.Row})
	}
}

func (log *PlayerExtendLevelInsertLog) SetEventId(time.Time) {
}

func (log *PlayerExtendLevelInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(51)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.CoinPassTime)
	file.WriteInt64(log.GoRow.ExpPassTime)
	file.WriteInt64(log.GoRow.GhostPassTime)
	file.WriteInt64(log.GoRow.PetPassTime)
	file.WriteInt64(log.GoRow.BuddyPassTime)
	file.WriteInt8(log.GoRow.CoinDailyNum)
	file.WriteInt8(log.GoRow.ExpDailyNum)
	file.WriteInt8(log.GoRow.BuddyDailyNum)
	file.WriteInt8(log.GoRow.PetDailyNum)
	file.WriteInt8(log.GoRow.GhostDailyNum)
	file.WriteInt8(log.GoRow.RandBuddyRoleId)
	file.WriteInt8(log.GoRow.BuddyPos)
	file.WriteInt8(log.GoRow.BuddyTactical)
	file.WriteInt8(log.GoRow.RolePos)
	file.WriteInt16(log.GoRow.ExpMaxlevel)
	file.WriteInt16(log.GoRow.CoinsMaxlevel)
	file.WriteInt16(log.GoRow.CoinsBuyNum)
	file.WriteInt16(log.GoRow.ExpBuyNum)
	file.WriteInt64(log.GoRow.CoinsBuyTime)
	file.WriteInt64(log.GoRow.ExpBuyTime)
}

func (log *PlayerExtendLevelInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerExtendLevel.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.CoinPassTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ExpPassTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.GhostPassTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.PetPassTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BuddyPassTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.CoinDailyNum)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.ExpDailyNum)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.BuddyDailyNum)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.PetDailyNum)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.GhostDailyNum)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RandBuddyRoleId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.BuddyPos)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.BuddyTactical)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RolePos)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ExpMaxlevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CoinsMaxlevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CoinsBuyNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ExpBuyNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.CoinsBuyTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ExpBuyTime)))
	return stmt.Execute()
}

func (log *PlayerExtendLevelInsertLog) CommitToTLog() {
}

func (log *PlayerExtendLevelInsertLog) CommitToXdLog() {
}

func (log *PlayerExtendLevelInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerExtendLevelDeleteLog struct {
	db *Database
	Old *C.PlayerExtendLevel
	GoOld *PlayerExtendLevel
}

func (db *Database) newPlayerExtendLevelDeleteLog(old *C.PlayerExtendLevel, goOld *PlayerExtendLevel) *PlayerExtendLevelDeleteLog {
	return &PlayerExtendLevelDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerExtendLevelDeleteLog) Free() {
	C.Free_PlayerExtendLevel(log.Old)
}

func (log *PlayerExtendLevelDeleteLog) InvokeHook() {
	if g_Hooks.PlayerExtendLevel != nil {
		g_Hooks.PlayerExtendLevel.Delete(&PlayerExtendLevelRow{row: log.Old})
	}
}

func (log *PlayerExtendLevelDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerExtendLevelDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(51)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.CoinPassTime)
	file.WriteInt64(log.GoOld.ExpPassTime)
	file.WriteInt64(log.GoOld.GhostPassTime)
	file.WriteInt64(log.GoOld.PetPassTime)
	file.WriteInt64(log.GoOld.BuddyPassTime)
	file.WriteInt8(log.GoOld.CoinDailyNum)
	file.WriteInt8(log.GoOld.ExpDailyNum)
	file.WriteInt8(log.GoOld.BuddyDailyNum)
	file.WriteInt8(log.GoOld.PetDailyNum)
	file.WriteInt8(log.GoOld.GhostDailyNum)
	file.WriteInt8(log.GoOld.RandBuddyRoleId)
	file.WriteInt8(log.GoOld.BuddyPos)
	file.WriteInt8(log.GoOld.BuddyTactical)
	file.WriteInt8(log.GoOld.RolePos)
	file.WriteInt16(log.GoOld.ExpMaxlevel)
	file.WriteInt16(log.GoOld.CoinsMaxlevel)
	file.WriteInt16(log.GoOld.CoinsBuyNum)
	file.WriteInt16(log.GoOld.ExpBuyNum)
	file.WriteInt64(log.GoOld.CoinsBuyTime)
	file.WriteInt64(log.GoOld.ExpBuyTime)
}

func (log *PlayerExtendLevelDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerExtendLevel.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerExtendLevelDeleteLog) CommitToTLog() {
}

func (log *PlayerExtendLevelDeleteLog) CommitToXdLog() {
}

func (log *PlayerExtendLevelDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerExtendLevelUpdateLog struct {
	db *Database 
	Row *C.PlayerExtendLevel
	Old *C.PlayerExtendLevel
	GoNew *PlayerExtendLevel
	GoOld *PlayerExtendLevel
}

func (db *Database) newPlayerExtendLevelUpdateLog(row, old *C.PlayerExtendLevel, goNew, goOld *PlayerExtendLevel) *PlayerExtendLevelUpdateLog {
	return &PlayerExtendLevelUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerExtendLevelUpdateLog) Free() {
	C.Free_PlayerExtendLevel(log.Old)
}

func (log *PlayerExtendLevelUpdateLog) InvokeHook() {
	if g_Hooks.PlayerExtendLevel != nil {
		g_Hooks.PlayerExtendLevel.Update(&PlayerExtendLevelRow{row: log.Row}, &PlayerExtendLevelRow{row: log.Old})
	}
}

func (log *PlayerExtendLevelUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerExtendLevelUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(51)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.CoinPassTime)
	file.WriteInt64(log.GoNew.ExpPassTime)
	file.WriteInt64(log.GoNew.GhostPassTime)
	file.WriteInt64(log.GoNew.PetPassTime)
	file.WriteInt64(log.GoNew.BuddyPassTime)
	file.WriteInt8(log.GoNew.CoinDailyNum)
	file.WriteInt8(log.GoNew.ExpDailyNum)
	file.WriteInt8(log.GoNew.BuddyDailyNum)
	file.WriteInt8(log.GoNew.PetDailyNum)
	file.WriteInt8(log.GoNew.GhostDailyNum)
	file.WriteInt8(log.GoNew.RandBuddyRoleId)
	file.WriteInt8(log.GoNew.BuddyPos)
	file.WriteInt8(log.GoNew.BuddyTactical)
	file.WriteInt8(log.GoNew.RolePos)
	file.WriteInt16(log.GoNew.ExpMaxlevel)
	file.WriteInt16(log.GoNew.CoinsMaxlevel)
	file.WriteInt16(log.GoNew.CoinsBuyNum)
	file.WriteInt16(log.GoNew.ExpBuyNum)
	file.WriteInt64(log.GoNew.CoinsBuyTime)
	file.WriteInt64(log.GoNew.ExpBuyTime)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.CoinPassTime)
	file.WriteInt64(log.GoOld.ExpPassTime)
	file.WriteInt64(log.GoOld.GhostPassTime)
	file.WriteInt64(log.GoOld.PetPassTime)
	file.WriteInt64(log.GoOld.BuddyPassTime)
	file.WriteInt8(log.GoOld.CoinDailyNum)
	file.WriteInt8(log.GoOld.ExpDailyNum)
	file.WriteInt8(log.GoOld.BuddyDailyNum)
	file.WriteInt8(log.GoOld.PetDailyNum)
	file.WriteInt8(log.GoOld.GhostDailyNum)
	file.WriteInt8(log.GoOld.RandBuddyRoleId)
	file.WriteInt8(log.GoOld.BuddyPos)
	file.WriteInt8(log.GoOld.BuddyTactical)
	file.WriteInt8(log.GoOld.RolePos)
	file.WriteInt16(log.GoOld.ExpMaxlevel)
	file.WriteInt16(log.GoOld.CoinsMaxlevel)
	file.WriteInt16(log.GoOld.CoinsBuyNum)
	file.WriteInt16(log.GoOld.ExpBuyNum)
	file.WriteInt64(log.GoOld.CoinsBuyTime)
	file.WriteInt64(log.GoOld.ExpBuyTime)
}

func (log *PlayerExtendLevelUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerExtendLevel.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.CoinPassTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ExpPassTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.GhostPassTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.PetPassTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BuddyPassTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.CoinDailyNum)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.ExpDailyNum)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.BuddyDailyNum)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.PetDailyNum)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.GhostDailyNum)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RandBuddyRoleId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.BuddyPos)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.BuddyTactical)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RolePos)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ExpMaxlevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CoinsMaxlevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CoinsBuyNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ExpBuyNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.CoinsBuyTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ExpBuyTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerExtendLevelUpdateLog) CommitToTLog() {
}

func (log *PlayerExtendLevelUpdateLog) CommitToXdLog() {
}

func (log *PlayerExtendLevelUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerExtendLevelInsertLog) Rollback() {
	if log.db.tables.PlayerExtendLevel != log.Row {
		panic("insert rollback check failed 'PlayerExtendLevel'")
	}
	log.db.tables.PlayerExtendLevel = nil
	C.Free_PlayerExtendLevel(log.Row)
}

func (log *PlayerExtendLevelDeleteLog) Rollback() {
	if log.db.tables.PlayerExtendLevel != nil {
		panic("delete rollback check failed 'player_extend_level'")
	}
	log.db.tables.PlayerExtendLevel = log.Old
}

func (log *PlayerExtendLevelUpdateLog) Rollback() {
	if log.db.tables.PlayerExtendLevel != log.Row {
		panic("update rollback check failed 'player_extend_level'")
	}
	log.db.tables.PlayerExtendLevel = log.Old
	C.Free_PlayerExtendLevel(log.Row)
}

/* ========== player_extend_quest ========== */

type PlayerExtendQuestInsertLog struct {
	db *Database 
	Row *C.PlayerExtendQuest
	GoRow *PlayerExtendQuest
}

func (db *Database) newPlayerExtendQuestInsertLog(row *C.PlayerExtendQuest, goRow *PlayerExtendQuest) *PlayerExtendQuestInsertLog {
	return &PlayerExtendQuestInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerExtendQuestInsertLog) Free() {
}

func (log *PlayerExtendQuestInsertLog) InvokeHook() {
	if g_Hooks.PlayerExtendQuest != nil {
		g_Hooks.PlayerExtendQuest.Insert(&PlayerExtendQuestRow{row: log.Row})
	}
}

func (log *PlayerExtendQuestInsertLog) SetEventId(time.Time) {
}

func (log *PlayerExtendQuestInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(52)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.QuestId)
	file.WriteInt16(log.GoRow.Progress)
	file.WriteInt8(log.GoRow.State)
}

func (log *PlayerExtendQuestInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerExtendQuest.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.QuestId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Progress)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.State)))
	return stmt.Execute()
}

func (log *PlayerExtendQuestInsertLog) CommitToTLog() {
}

func (log *PlayerExtendQuestInsertLog) CommitToXdLog() {
}

func (log *PlayerExtendQuestInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerExtendQuestDeleteLog struct {
	db *Database
	Old *C.PlayerExtendQuest
	GoOld *PlayerExtendQuest
}

func (db *Database) newPlayerExtendQuestDeleteLog(old *C.PlayerExtendQuest, goOld *PlayerExtendQuest) *PlayerExtendQuestDeleteLog {
	return &PlayerExtendQuestDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerExtendQuestDeleteLog) Free() {
	C.Free_PlayerExtendQuest(log.Old)
}

func (log *PlayerExtendQuestDeleteLog) InvokeHook() {
	if g_Hooks.PlayerExtendQuest != nil {
		g_Hooks.PlayerExtendQuest.Delete(&PlayerExtendQuestRow{row: log.Old})
	}
}

func (log *PlayerExtendQuestDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerExtendQuestDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(52)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.QuestId)
	file.WriteInt16(log.GoOld.Progress)
	file.WriteInt8(log.GoOld.State)
}

func (log *PlayerExtendQuestDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerExtendQuest.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerExtendQuestDeleteLog) CommitToTLog() {
}

func (log *PlayerExtendQuestDeleteLog) CommitToXdLog() {
}

func (log *PlayerExtendQuestDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerExtendQuestUpdateLog struct {
	db *Database 
	Row *C.PlayerExtendQuest
	Old *C.PlayerExtendQuest
	GoNew *PlayerExtendQuest
	GoOld *PlayerExtendQuest
}

func (db *Database) newPlayerExtendQuestUpdateLog(row, old *C.PlayerExtendQuest, goNew, goOld *PlayerExtendQuest) *PlayerExtendQuestUpdateLog {
	return &PlayerExtendQuestUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerExtendQuestUpdateLog) Free() {
	C.Free_PlayerExtendQuest(log.Old)
}

func (log *PlayerExtendQuestUpdateLog) InvokeHook() {
	if g_Hooks.PlayerExtendQuest != nil {
		g_Hooks.PlayerExtendQuest.Update(&PlayerExtendQuestRow{row: log.Row}, &PlayerExtendQuestRow{row: log.Old})
	}
}

func (log *PlayerExtendQuestUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerExtendQuestUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(52)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.QuestId)
	file.WriteInt16(log.GoNew.Progress)
	file.WriteInt8(log.GoNew.State)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.QuestId)
	file.WriteInt16(log.GoOld.Progress)
	file.WriteInt8(log.GoOld.State)
}

func (log *PlayerExtendQuestUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerExtendQuest.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.QuestId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Progress)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.State)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerExtendQuestUpdateLog) CommitToTLog() {
}

func (log *PlayerExtendQuestUpdateLog) CommitToXdLog() {
}

func (log *PlayerExtendQuestUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerExtendQuestInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerExtendQuest
	for crow := log.db.tables.PlayerExtendQuest; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerExtendQuest'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerExtendQuest = log.db.tables.PlayerExtendQuest.next
	}
	C.Free_PlayerExtendQuest(log.Row)
}

func (log *PlayerExtendQuestDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerExtendQuest
	for crow := log.db.tables.PlayerExtendQuest; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_extend_quest'")
	}
	log.Old.next = log.db.tables.PlayerExtendQuest
	log.db.tables.PlayerExtendQuest = log.Old
}

func (log *PlayerExtendQuestUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerExtendQuest
	for crow := log.db.tables.PlayerExtendQuest; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_extend_quest'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerExtendQuest = log.Old
	}
	C.Free_PlayerExtendQuest(log.Row)
}

/* ========== player_fame ========== */

type PlayerFameInsertLog struct {
	db *Database 
	Row *C.PlayerFame
	GoRow *PlayerFame
}

func (db *Database) newPlayerFameInsertLog(row *C.PlayerFame, goRow *PlayerFame) *PlayerFameInsertLog {
	return &PlayerFameInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerFameInsertLog) Free() {
}

func (log *PlayerFameInsertLog) InvokeHook() {
	if g_Hooks.PlayerFame != nil {
		g_Hooks.PlayerFame.Insert(&PlayerFameRow{row: log.Row})
	}
}

func (log *PlayerFameInsertLog) SetEventId(time.Time) {
}

func (log *PlayerFameInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(53)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.Fame)
	file.WriteInt16(log.GoRow.Level)
	file.WriteInt32(log.GoRow.ArenaFame)
	file.WriteInt32(log.GoRow.MultLevelFame)
}

func (log *PlayerFameInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFame.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Fame)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.ArenaFame)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.MultLevelFame)))
	return stmt.Execute()
}

func (log *PlayerFameInsertLog) CommitToTLog() {
}

func (log *PlayerFameInsertLog) CommitToXdLog() {
}

func (log *PlayerFameInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerFameDeleteLog struct {
	db *Database
	Old *C.PlayerFame
	GoOld *PlayerFame
}

func (db *Database) newPlayerFameDeleteLog(old *C.PlayerFame, goOld *PlayerFame) *PlayerFameDeleteLog {
	return &PlayerFameDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerFameDeleteLog) Free() {
	C.Free_PlayerFame(log.Old)
}

func (log *PlayerFameDeleteLog) InvokeHook() {
	if g_Hooks.PlayerFame != nil {
		g_Hooks.PlayerFame.Delete(&PlayerFameRow{row: log.Old})
	}
}

func (log *PlayerFameDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerFameDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(53)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Fame)
	file.WriteInt16(log.GoOld.Level)
	file.WriteInt32(log.GoOld.ArenaFame)
	file.WriteInt32(log.GoOld.MultLevelFame)
}

func (log *PlayerFameDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFame.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerFameDeleteLog) CommitToTLog() {
}

func (log *PlayerFameDeleteLog) CommitToXdLog() {
}

func (log *PlayerFameDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerFameUpdateLog struct {
	db *Database 
	Row *C.PlayerFame
	Old *C.PlayerFame
	GoNew *PlayerFame
	GoOld *PlayerFame
}

func (db *Database) newPlayerFameUpdateLog(row, old *C.PlayerFame, goNew, goOld *PlayerFame) *PlayerFameUpdateLog {
	return &PlayerFameUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerFameUpdateLog) Free() {
	C.Free_PlayerFame(log.Old)
}

func (log *PlayerFameUpdateLog) InvokeHook() {
	if g_Hooks.PlayerFame != nil {
		g_Hooks.PlayerFame.Update(&PlayerFameRow{row: log.Row}, &PlayerFameRow{row: log.Old})
	}
}

func (log *PlayerFameUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerFameUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(53)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.Fame)
	file.WriteInt16(log.GoNew.Level)
	file.WriteInt32(log.GoNew.ArenaFame)
	file.WriteInt32(log.GoNew.MultLevelFame)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Fame)
	file.WriteInt16(log.GoOld.Level)
	file.WriteInt32(log.GoOld.ArenaFame)
	file.WriteInt32(log.GoOld.MultLevelFame)
}

func (log *PlayerFameUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFame.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.Fame)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.ArenaFame)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.MultLevelFame)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerFameUpdateLog) CommitToTLog() {
}

func (log *PlayerFameUpdateLog) CommitToXdLog() {
}

func (log *PlayerFameUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerFameInsertLog) Rollback() {
	if log.db.tables.PlayerFame != log.Row {
		panic("insert rollback check failed 'PlayerFame'")
	}
	log.db.tables.PlayerFame = nil
	C.Free_PlayerFame(log.Row)
}

func (log *PlayerFameDeleteLog) Rollback() {
	if log.db.tables.PlayerFame != nil {
		panic("delete rollback check failed 'player_fame'")
	}
	log.db.tables.PlayerFame = log.Old
}

func (log *PlayerFameUpdateLog) Rollback() {
	if log.db.tables.PlayerFame != log.Row {
		panic("update rollback check failed 'player_fame'")
	}
	log.db.tables.PlayerFame = log.Old
	C.Free_PlayerFame(log.Row)
}

/* ========== player_fashion ========== */

type PlayerFashionInsertLog struct {
	db *Database 
	Row *C.PlayerFashion
	GoRow *PlayerFashion
}

func (db *Database) newPlayerFashionInsertLog(row *C.PlayerFashion, goRow *PlayerFashion) *PlayerFashionInsertLog {
	return &PlayerFashionInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerFashionInsertLog) Free() {
}

func (log *PlayerFashionInsertLog) InvokeHook() {
	if g_Hooks.PlayerFashion != nil {
		g_Hooks.PlayerFashion.Insert(&PlayerFashionRow{row: log.Row})
	}
}

func (log *PlayerFashionInsertLog) SetEventId(time.Time) {
}

func (log *PlayerFashionInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(54)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.FashionId)
	file.WriteInt64(log.GoRow.ExpireTime)
}

func (log *PlayerFashionInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFashion.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.FashionId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ExpireTime)))
	return stmt.Execute()
}

func (log *PlayerFashionInsertLog) CommitToTLog() {
}

func (log *PlayerFashionInsertLog) CommitToXdLog() {
}

func (log *PlayerFashionInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerFashionDeleteLog struct {
	db *Database
	Old *C.PlayerFashion
	GoOld *PlayerFashion
}

func (db *Database) newPlayerFashionDeleteLog(old *C.PlayerFashion, goOld *PlayerFashion) *PlayerFashionDeleteLog {
	return &PlayerFashionDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerFashionDeleteLog) Free() {
	C.Free_PlayerFashion(log.Old)
}

func (log *PlayerFashionDeleteLog) InvokeHook() {
	if g_Hooks.PlayerFashion != nil {
		g_Hooks.PlayerFashion.Delete(&PlayerFashionRow{row: log.Old})
	}
}

func (log *PlayerFashionDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerFashionDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(54)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.FashionId)
	file.WriteInt64(log.GoOld.ExpireTime)
}

func (log *PlayerFashionDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFashion.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerFashionDeleteLog) CommitToTLog() {
}

func (log *PlayerFashionDeleteLog) CommitToXdLog() {
}

func (log *PlayerFashionDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerFashionUpdateLog struct {
	db *Database 
	Row *C.PlayerFashion
	Old *C.PlayerFashion
	GoNew *PlayerFashion
	GoOld *PlayerFashion
}

func (db *Database) newPlayerFashionUpdateLog(row, old *C.PlayerFashion, goNew, goOld *PlayerFashion) *PlayerFashionUpdateLog {
	return &PlayerFashionUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerFashionUpdateLog) Free() {
	C.Free_PlayerFashion(log.Old)
}

func (log *PlayerFashionUpdateLog) InvokeHook() {
	if g_Hooks.PlayerFashion != nil {
		g_Hooks.PlayerFashion.Update(&PlayerFashionRow{row: log.Row}, &PlayerFashionRow{row: log.Old})
	}
}

func (log *PlayerFashionUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerFashionUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(54)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.FashionId)
	file.WriteInt64(log.GoNew.ExpireTime)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.FashionId)
	file.WriteInt64(log.GoOld.ExpireTime)
}

func (log *PlayerFashionUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFashion.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.FashionId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ExpireTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerFashionUpdateLog) CommitToTLog() {
}

func (log *PlayerFashionUpdateLog) CommitToXdLog() {
}

func (log *PlayerFashionUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerFashionInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerFashion
	for crow := log.db.tables.PlayerFashion; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerFashion'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerFashion = log.db.tables.PlayerFashion.next
	}
	C.Free_PlayerFashion(log.Row)
}

func (log *PlayerFashionDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerFashion
	for crow := log.db.tables.PlayerFashion; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_fashion'")
	}
	log.Old.next = log.db.tables.PlayerFashion
	log.db.tables.PlayerFashion = log.Old
}

func (log *PlayerFashionUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerFashion
	for crow := log.db.tables.PlayerFashion; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_fashion'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerFashion = log.Old
	}
	C.Free_PlayerFashion(log.Row)
}

/* ========== player_fashion_state ========== */

type PlayerFashionStateInsertLog struct {
	db *Database 
	Row *C.PlayerFashionState
	GoRow *PlayerFashionState
}

func (db *Database) newPlayerFashionStateInsertLog(row *C.PlayerFashionState, goRow *PlayerFashionState) *PlayerFashionStateInsertLog {
	return &PlayerFashionStateInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerFashionStateInsertLog) Free() {
}

func (log *PlayerFashionStateInsertLog) InvokeHook() {
	if g_Hooks.PlayerFashionState != nil {
		g_Hooks.PlayerFashionState.Insert(&PlayerFashionStateRow{row: log.Row})
	}
}

func (log *PlayerFashionStateInsertLog) SetEventId(time.Time) {
}

func (log *PlayerFashionStateInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(55)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.UpdateTime)
	file.WriteInt16(log.GoRow.DressedFashionId)
}

func (log *PlayerFashionStateInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFashionState.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.UpdateTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DressedFashionId)))
	return stmt.Execute()
}

func (log *PlayerFashionStateInsertLog) CommitToTLog() {
}

func (log *PlayerFashionStateInsertLog) CommitToXdLog() {
}

func (log *PlayerFashionStateInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerFashionStateDeleteLog struct {
	db *Database
	Old *C.PlayerFashionState
	GoOld *PlayerFashionState
}

func (db *Database) newPlayerFashionStateDeleteLog(old *C.PlayerFashionState, goOld *PlayerFashionState) *PlayerFashionStateDeleteLog {
	return &PlayerFashionStateDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerFashionStateDeleteLog) Free() {
	C.Free_PlayerFashionState(log.Old)
}

func (log *PlayerFashionStateDeleteLog) InvokeHook() {
	if g_Hooks.PlayerFashionState != nil {
		g_Hooks.PlayerFashionState.Delete(&PlayerFashionStateRow{row: log.Old})
	}
}

func (log *PlayerFashionStateDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerFashionStateDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(55)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.UpdateTime)
	file.WriteInt16(log.GoOld.DressedFashionId)
}

func (log *PlayerFashionStateDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFashionState.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerFashionStateDeleteLog) CommitToTLog() {
}

func (log *PlayerFashionStateDeleteLog) CommitToXdLog() {
}

func (log *PlayerFashionStateDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerFashionStateUpdateLog struct {
	db *Database 
	Row *C.PlayerFashionState
	Old *C.PlayerFashionState
	GoNew *PlayerFashionState
	GoOld *PlayerFashionState
}

func (db *Database) newPlayerFashionStateUpdateLog(row, old *C.PlayerFashionState, goNew, goOld *PlayerFashionState) *PlayerFashionStateUpdateLog {
	return &PlayerFashionStateUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerFashionStateUpdateLog) Free() {
	C.Free_PlayerFashionState(log.Old)
}

func (log *PlayerFashionStateUpdateLog) InvokeHook() {
	if g_Hooks.PlayerFashionState != nil {
		g_Hooks.PlayerFashionState.Update(&PlayerFashionStateRow{row: log.Row}, &PlayerFashionStateRow{row: log.Old})
	}
}

func (log *PlayerFashionStateUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerFashionStateUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(55)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.UpdateTime)
	file.WriteInt16(log.GoNew.DressedFashionId)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.UpdateTime)
	file.WriteInt16(log.GoOld.DressedFashionId)
}

func (log *PlayerFashionStateUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFashionState.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.UpdateTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DressedFashionId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerFashionStateUpdateLog) CommitToTLog() {
}

func (log *PlayerFashionStateUpdateLog) CommitToXdLog() {
}

func (log *PlayerFashionStateUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerFashionStateInsertLog) Rollback() {
	if log.db.tables.PlayerFashionState != log.Row {
		panic("insert rollback check failed 'PlayerFashionState'")
	}
	log.db.tables.PlayerFashionState = nil
	C.Free_PlayerFashionState(log.Row)
}

func (log *PlayerFashionStateDeleteLog) Rollback() {
	if log.db.tables.PlayerFashionState != nil {
		panic("delete rollback check failed 'player_fashion_state'")
	}
	log.db.tables.PlayerFashionState = log.Old
}

func (log *PlayerFashionStateUpdateLog) Rollback() {
	if log.db.tables.PlayerFashionState != log.Row {
		panic("update rollback check failed 'player_fashion_state'")
	}
	log.db.tables.PlayerFashionState = log.Old
	C.Free_PlayerFashionState(log.Row)
}

/* ========== player_fate_box_state ========== */

type PlayerFateBoxStateInsertLog struct {
	db *Database 
	Row *C.PlayerFateBoxState
	GoRow *PlayerFateBoxState
}

func (db *Database) newPlayerFateBoxStateInsertLog(row *C.PlayerFateBoxState, goRow *PlayerFateBoxState) *PlayerFateBoxStateInsertLog {
	return &PlayerFateBoxStateInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerFateBoxStateInsertLog) Free() {
}

func (log *PlayerFateBoxStateInsertLog) InvokeHook() {
	if g_Hooks.PlayerFateBoxState != nil {
		g_Hooks.PlayerFateBoxState.Insert(&PlayerFateBoxStateRow{row: log.Row})
	}
}

func (log *PlayerFateBoxStateInsertLog) SetEventId(time.Time) {
}

func (log *PlayerFateBoxStateInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(56)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.Lock)
	file.WriteInt64(log.GoRow.StarBoxFreeDrawTimestamp)
	file.WriteInt32(log.GoRow.StarBoxDrawCount)
	file.WriteInt64(log.GoRow.MoonBoxFreeDrawTimestamp)
	file.WriteInt32(log.GoRow.MoonBoxDrawCount)
	file.WriteInt64(log.GoRow.SunBoxFreeDrawTimestamp)
	file.WriteInt32(log.GoRow.SunBoxDrawCount)
	file.WriteInt64(log.GoRow.HunyuanBoxFreeDrawTimestamp)
	file.WriteInt32(log.GoRow.HunyuanBoxDrawCount)
}

func (log *PlayerFateBoxStateInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFateBoxState.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Lock)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.StarBoxFreeDrawTimestamp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.StarBoxDrawCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.MoonBoxFreeDrawTimestamp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.MoonBoxDrawCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SunBoxFreeDrawTimestamp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.SunBoxDrawCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.HunyuanBoxFreeDrawTimestamp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.HunyuanBoxDrawCount)))
	return stmt.Execute()
}

func (log *PlayerFateBoxStateInsertLog) CommitToTLog() {
}

func (log *PlayerFateBoxStateInsertLog) CommitToXdLog() {
}

func (log *PlayerFateBoxStateInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerFateBoxStateDeleteLog struct {
	db *Database
	Old *C.PlayerFateBoxState
	GoOld *PlayerFateBoxState
}

func (db *Database) newPlayerFateBoxStateDeleteLog(old *C.PlayerFateBoxState, goOld *PlayerFateBoxState) *PlayerFateBoxStateDeleteLog {
	return &PlayerFateBoxStateDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerFateBoxStateDeleteLog) Free() {
	C.Free_PlayerFateBoxState(log.Old)
}

func (log *PlayerFateBoxStateDeleteLog) InvokeHook() {
	if g_Hooks.PlayerFateBoxState != nil {
		g_Hooks.PlayerFateBoxState.Delete(&PlayerFateBoxStateRow{row: log.Old})
	}
}

func (log *PlayerFateBoxStateDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerFateBoxStateDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(56)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Lock)
	file.WriteInt64(log.GoOld.StarBoxFreeDrawTimestamp)
	file.WriteInt32(log.GoOld.StarBoxDrawCount)
	file.WriteInt64(log.GoOld.MoonBoxFreeDrawTimestamp)
	file.WriteInt32(log.GoOld.MoonBoxDrawCount)
	file.WriteInt64(log.GoOld.SunBoxFreeDrawTimestamp)
	file.WriteInt32(log.GoOld.SunBoxDrawCount)
	file.WriteInt64(log.GoOld.HunyuanBoxFreeDrawTimestamp)
	file.WriteInt32(log.GoOld.HunyuanBoxDrawCount)
}

func (log *PlayerFateBoxStateDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFateBoxState.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerFateBoxStateDeleteLog) CommitToTLog() {
}

func (log *PlayerFateBoxStateDeleteLog) CommitToXdLog() {
}

func (log *PlayerFateBoxStateDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerFateBoxStateUpdateLog struct {
	db *Database 
	Row *C.PlayerFateBoxState
	Old *C.PlayerFateBoxState
	GoNew *PlayerFateBoxState
	GoOld *PlayerFateBoxState
}

func (db *Database) newPlayerFateBoxStateUpdateLog(row, old *C.PlayerFateBoxState, goNew, goOld *PlayerFateBoxState) *PlayerFateBoxStateUpdateLog {
	return &PlayerFateBoxStateUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerFateBoxStateUpdateLog) Free() {
	C.Free_PlayerFateBoxState(log.Old)
}

func (log *PlayerFateBoxStateUpdateLog) InvokeHook() {
	if g_Hooks.PlayerFateBoxState != nil {
		g_Hooks.PlayerFateBoxState.Update(&PlayerFateBoxStateRow{row: log.Row}, &PlayerFateBoxStateRow{row: log.Old})
	}
}

func (log *PlayerFateBoxStateUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerFateBoxStateUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(56)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.Lock)
	file.WriteInt64(log.GoNew.StarBoxFreeDrawTimestamp)
	file.WriteInt32(log.GoNew.StarBoxDrawCount)
	file.WriteInt64(log.GoNew.MoonBoxFreeDrawTimestamp)
	file.WriteInt32(log.GoNew.MoonBoxDrawCount)
	file.WriteInt64(log.GoNew.SunBoxFreeDrawTimestamp)
	file.WriteInt32(log.GoNew.SunBoxDrawCount)
	file.WriteInt64(log.GoNew.HunyuanBoxFreeDrawTimestamp)
	file.WriteInt32(log.GoNew.HunyuanBoxDrawCount)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Lock)
	file.WriteInt64(log.GoOld.StarBoxFreeDrawTimestamp)
	file.WriteInt32(log.GoOld.StarBoxDrawCount)
	file.WriteInt64(log.GoOld.MoonBoxFreeDrawTimestamp)
	file.WriteInt32(log.GoOld.MoonBoxDrawCount)
	file.WriteInt64(log.GoOld.SunBoxFreeDrawTimestamp)
	file.WriteInt32(log.GoOld.SunBoxDrawCount)
	file.WriteInt64(log.GoOld.HunyuanBoxFreeDrawTimestamp)
	file.WriteInt32(log.GoOld.HunyuanBoxDrawCount)
}

func (log *PlayerFateBoxStateUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFateBoxState.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.Lock)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.StarBoxFreeDrawTimestamp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.StarBoxDrawCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.MoonBoxFreeDrawTimestamp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.MoonBoxDrawCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SunBoxFreeDrawTimestamp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.SunBoxDrawCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.HunyuanBoxFreeDrawTimestamp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.HunyuanBoxDrawCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerFateBoxStateUpdateLog) CommitToTLog() {
}

func (log *PlayerFateBoxStateUpdateLog) CommitToXdLog() {
}

func (log *PlayerFateBoxStateUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerFateBoxStateInsertLog) Rollback() {
	if log.db.tables.PlayerFateBoxState != log.Row {
		panic("insert rollback check failed 'PlayerFateBoxState'")
	}
	log.db.tables.PlayerFateBoxState = nil
	C.Free_PlayerFateBoxState(log.Row)
}

func (log *PlayerFateBoxStateDeleteLog) Rollback() {
	if log.db.tables.PlayerFateBoxState != nil {
		panic("delete rollback check failed 'player_fate_box_state'")
	}
	log.db.tables.PlayerFateBoxState = log.Old
}

func (log *PlayerFateBoxStateUpdateLog) Rollback() {
	if log.db.tables.PlayerFateBoxState != log.Row {
		panic("update rollback check failed 'player_fate_box_state'")
	}
	log.db.tables.PlayerFateBoxState = log.Old
	C.Free_PlayerFateBoxState(log.Row)
}

/* ========== player_fight_num ========== */

type PlayerFightNumInsertLog struct {
	db *Database 
	Row *C.PlayerFightNum
	GoRow *PlayerFightNum
}

func (db *Database) newPlayerFightNumInsertLog(row *C.PlayerFightNum, goRow *PlayerFightNum) *PlayerFightNumInsertLog {
	return &PlayerFightNumInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerFightNumInsertLog) Free() {
}

func (log *PlayerFightNumInsertLog) InvokeHook() {
	if g_Hooks.PlayerFightNum != nil {
		g_Hooks.PlayerFightNum.Insert(&PlayerFightNumRow{row: log.Row})
	}
}

func (log *PlayerFightNumInsertLog) SetEventId(time.Time) {
}

func (log *PlayerFightNumInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(57)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.FightNum)
}

func (log *PlayerFightNumInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFightNum.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.FightNum)))
	return stmt.Execute()
}

func (log *PlayerFightNumInsertLog) CommitToTLog() {
}

func (log *PlayerFightNumInsertLog) CommitToXdLog() {
}

func (log *PlayerFightNumInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerFightNumDeleteLog struct {
	db *Database
	Old *C.PlayerFightNum
	GoOld *PlayerFightNum
}

func (db *Database) newPlayerFightNumDeleteLog(old *C.PlayerFightNum, goOld *PlayerFightNum) *PlayerFightNumDeleteLog {
	return &PlayerFightNumDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerFightNumDeleteLog) Free() {
	C.Free_PlayerFightNum(log.Old)
}

func (log *PlayerFightNumDeleteLog) InvokeHook() {
	if g_Hooks.PlayerFightNum != nil {
		g_Hooks.PlayerFightNum.Delete(&PlayerFightNumRow{row: log.Old})
	}
}

func (log *PlayerFightNumDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerFightNumDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(57)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.FightNum)
}

func (log *PlayerFightNumDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFightNum.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerFightNumDeleteLog) CommitToTLog() {
}

func (log *PlayerFightNumDeleteLog) CommitToXdLog() {
}

func (log *PlayerFightNumDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerFightNumUpdateLog struct {
	db *Database 
	Row *C.PlayerFightNum
	Old *C.PlayerFightNum
	GoNew *PlayerFightNum
	GoOld *PlayerFightNum
}

func (db *Database) newPlayerFightNumUpdateLog(row, old *C.PlayerFightNum, goNew, goOld *PlayerFightNum) *PlayerFightNumUpdateLog {
	return &PlayerFightNumUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerFightNumUpdateLog) Free() {
	C.Free_PlayerFightNum(log.Old)
}

func (log *PlayerFightNumUpdateLog) InvokeHook() {
	if g_Hooks.PlayerFightNum != nil {
		g_Hooks.PlayerFightNum.Update(&PlayerFightNumRow{row: log.Row}, &PlayerFightNumRow{row: log.Old})
	}
}

func (log *PlayerFightNumUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerFightNumUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(57)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.FightNum)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.FightNum)
}

func (log *PlayerFightNumUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFightNum.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.FightNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerFightNumUpdateLog) CommitToTLog() {
}

func (log *PlayerFightNumUpdateLog) CommitToXdLog() {
}

func (log *PlayerFightNumUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerFightNumInsertLog) Rollback() {
	if log.db.tables.PlayerFightNum != log.Row {
		panic("insert rollback check failed 'PlayerFightNum'")
	}
	log.db.tables.PlayerFightNum = nil
	C.Free_PlayerFightNum(log.Row)
}

func (log *PlayerFightNumDeleteLog) Rollback() {
	if log.db.tables.PlayerFightNum != nil {
		panic("delete rollback check failed 'player_fight_num'")
	}
	log.db.tables.PlayerFightNum = log.Old
}

func (log *PlayerFightNumUpdateLog) Rollback() {
	if log.db.tables.PlayerFightNum != log.Row {
		panic("update rollback check failed 'player_fight_num'")
	}
	log.db.tables.PlayerFightNum = log.Old
	C.Free_PlayerFightNum(log.Row)
}

/* ========== player_first_charge ========== */

type PlayerFirstChargeInsertLog struct {
	db *Database 
	Row *C.PlayerFirstCharge
	GoRow *PlayerFirstCharge
}

func (db *Database) newPlayerFirstChargeInsertLog(row *C.PlayerFirstCharge, goRow *PlayerFirstCharge) *PlayerFirstChargeInsertLog {
	return &PlayerFirstChargeInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerFirstChargeInsertLog) Free() {
}

func (log *PlayerFirstChargeInsertLog) InvokeHook() {
	if g_Hooks.PlayerFirstCharge != nil {
		g_Hooks.PlayerFirstCharge.Insert(&PlayerFirstChargeRow{row: log.Row})
	}
}

func (log *PlayerFirstChargeInsertLog) SetEventId(time.Time) {
}

func (log *PlayerFirstChargeInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(58)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteString(log.GoRow.ProductId)
}

func (log *PlayerFirstChargeInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFirstCharge.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.ProductId), int(log.Row.ProductId_len))
	return stmt.Execute()
}

func (log *PlayerFirstChargeInsertLog) CommitToTLog() {
}

func (log *PlayerFirstChargeInsertLog) CommitToXdLog() {
}

func (log *PlayerFirstChargeInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerFirstChargeDeleteLog struct {
	db *Database
	Old *C.PlayerFirstCharge
	GoOld *PlayerFirstCharge
}

func (db *Database) newPlayerFirstChargeDeleteLog(old *C.PlayerFirstCharge, goOld *PlayerFirstCharge) *PlayerFirstChargeDeleteLog {
	return &PlayerFirstChargeDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerFirstChargeDeleteLog) Free() {
	C.Free_PlayerFirstCharge(log.Old)
}

func (log *PlayerFirstChargeDeleteLog) InvokeHook() {
	if g_Hooks.PlayerFirstCharge != nil {
		g_Hooks.PlayerFirstCharge.Delete(&PlayerFirstChargeRow{row: log.Old})
	}
}

func (log *PlayerFirstChargeDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerFirstChargeDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(58)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteString(log.GoOld.ProductId)
}

func (log *PlayerFirstChargeDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFirstCharge.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerFirstChargeDeleteLog) CommitToTLog() {
}

func (log *PlayerFirstChargeDeleteLog) CommitToXdLog() {
}

func (log *PlayerFirstChargeDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerFirstChargeUpdateLog struct {
	db *Database 
	Row *C.PlayerFirstCharge
	Old *C.PlayerFirstCharge
	GoNew *PlayerFirstCharge
	GoOld *PlayerFirstCharge
}

func (db *Database) newPlayerFirstChargeUpdateLog(row, old *C.PlayerFirstCharge, goNew, goOld *PlayerFirstCharge) *PlayerFirstChargeUpdateLog {
	return &PlayerFirstChargeUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerFirstChargeUpdateLog) Free() {
	C.Free_PlayerFirstCharge(log.Old)
}

func (log *PlayerFirstChargeUpdateLog) InvokeHook() {
	if g_Hooks.PlayerFirstCharge != nil {
		g_Hooks.PlayerFirstCharge.Update(&PlayerFirstChargeRow{row: log.Row}, &PlayerFirstChargeRow{row: log.Old})
	}
}

func (log *PlayerFirstChargeUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerFirstChargeUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(58)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteString(log.GoNew.ProductId)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteString(log.GoOld.ProductId)
}

func (log *PlayerFirstChargeUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFirstCharge.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.ProductId), int(log.Row.ProductId_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerFirstChargeUpdateLog) CommitToTLog() {
}

func (log *PlayerFirstChargeUpdateLog) CommitToXdLog() {
}

func (log *PlayerFirstChargeUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerFirstChargeInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerFirstCharge
	for crow := log.db.tables.PlayerFirstCharge; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerFirstCharge'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerFirstCharge = log.db.tables.PlayerFirstCharge.next
	}
	C.Free_PlayerFirstCharge(log.Row)
}

func (log *PlayerFirstChargeDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerFirstCharge
	for crow := log.db.tables.PlayerFirstCharge; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_first_charge'")
	}
	log.Old.next = log.db.tables.PlayerFirstCharge
	log.db.tables.PlayerFirstCharge = log.Old
}

func (log *PlayerFirstChargeUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerFirstCharge
	for crow := log.db.tables.PlayerFirstCharge; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_first_charge'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerFirstCharge = log.Old
	}
	C.Free_PlayerFirstCharge(log.Row)
}

/* ========== player_formation ========== */

type PlayerFormationInsertLog struct {
	db *Database 
	Row *C.PlayerFormation
	GoRow *PlayerFormation
}

func (db *Database) newPlayerFormationInsertLog(row *C.PlayerFormation, goRow *PlayerFormation) *PlayerFormationInsertLog {
	return &PlayerFormationInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerFormationInsertLog) Free() {
}

func (log *PlayerFormationInsertLog) InvokeHook() {
	if g_Hooks.PlayerFormation != nil {
		g_Hooks.PlayerFormation.Insert(&PlayerFormationRow{row: log.Row})
	}
}

func (log *PlayerFormationInsertLog) SetEventId(time.Time) {
}

func (log *PlayerFormationInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(59)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.Pos0)
	file.WriteInt8(log.GoRow.Pos1)
	file.WriteInt8(log.GoRow.Pos2)
	file.WriteInt8(log.GoRow.Pos3)
	file.WriteInt8(log.GoRow.Pos4)
	file.WriteInt8(log.GoRow.Pos5)
	file.WriteInt8(log.GoRow.Pos6)
	file.WriteInt8(log.GoRow.Pos7)
	file.WriteInt8(log.GoRow.Pos8)
}

func (log *PlayerFormationInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFormation.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Pos0)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Pos1)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Pos2)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Pos3)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Pos4)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Pos5)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Pos6)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Pos7)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Pos8)))
	return stmt.Execute()
}

func (log *PlayerFormationInsertLog) CommitToTLog() {
}

func (log *PlayerFormationInsertLog) CommitToXdLog() {
}

func (log *PlayerFormationInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerFormationDeleteLog struct {
	db *Database
	Old *C.PlayerFormation
	GoOld *PlayerFormation
}

func (db *Database) newPlayerFormationDeleteLog(old *C.PlayerFormation, goOld *PlayerFormation) *PlayerFormationDeleteLog {
	return &PlayerFormationDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerFormationDeleteLog) Free() {
	C.Free_PlayerFormation(log.Old)
}

func (log *PlayerFormationDeleteLog) InvokeHook() {
	if g_Hooks.PlayerFormation != nil {
		g_Hooks.PlayerFormation.Delete(&PlayerFormationRow{row: log.Old})
	}
}

func (log *PlayerFormationDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerFormationDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(59)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.Pos0)
	file.WriteInt8(log.GoOld.Pos1)
	file.WriteInt8(log.GoOld.Pos2)
	file.WriteInt8(log.GoOld.Pos3)
	file.WriteInt8(log.GoOld.Pos4)
	file.WriteInt8(log.GoOld.Pos5)
	file.WriteInt8(log.GoOld.Pos6)
	file.WriteInt8(log.GoOld.Pos7)
	file.WriteInt8(log.GoOld.Pos8)
}

func (log *PlayerFormationDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFormation.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerFormationDeleteLog) CommitToTLog() {
}

func (log *PlayerFormationDeleteLog) CommitToXdLog() {
}

func (log *PlayerFormationDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerFormationUpdateLog struct {
	db *Database 
	Row *C.PlayerFormation
	Old *C.PlayerFormation
	GoNew *PlayerFormation
	GoOld *PlayerFormation
}

func (db *Database) newPlayerFormationUpdateLog(row, old *C.PlayerFormation, goNew, goOld *PlayerFormation) *PlayerFormationUpdateLog {
	return &PlayerFormationUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerFormationUpdateLog) Free() {
	C.Free_PlayerFormation(log.Old)
}

func (log *PlayerFormationUpdateLog) InvokeHook() {
	if g_Hooks.PlayerFormation != nil {
		g_Hooks.PlayerFormation.Update(&PlayerFormationRow{row: log.Row}, &PlayerFormationRow{row: log.Old})
	}
}

func (log *PlayerFormationUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerFormationUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(59)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.Pos0)
	file.WriteInt8(log.GoNew.Pos1)
	file.WriteInt8(log.GoNew.Pos2)
	file.WriteInt8(log.GoNew.Pos3)
	file.WriteInt8(log.GoNew.Pos4)
	file.WriteInt8(log.GoNew.Pos5)
	file.WriteInt8(log.GoNew.Pos6)
	file.WriteInt8(log.GoNew.Pos7)
	file.WriteInt8(log.GoNew.Pos8)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.Pos0)
	file.WriteInt8(log.GoOld.Pos1)
	file.WriteInt8(log.GoOld.Pos2)
	file.WriteInt8(log.GoOld.Pos3)
	file.WriteInt8(log.GoOld.Pos4)
	file.WriteInt8(log.GoOld.Pos5)
	file.WriteInt8(log.GoOld.Pos6)
	file.WriteInt8(log.GoOld.Pos7)
	file.WriteInt8(log.GoOld.Pos8)
}

func (log *PlayerFormationUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFormation.Update
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Pos0)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Pos1)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Pos2)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Pos3)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Pos4)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Pos5)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Pos6)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Pos7)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Pos8)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerFormationUpdateLog) CommitToTLog() {
}

func (log *PlayerFormationUpdateLog) CommitToXdLog() {
}

func (log *PlayerFormationUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerFormationInsertLog) Rollback() {
	if log.db.tables.PlayerFormation != log.Row {
		panic("insert rollback check failed 'PlayerFormation'")
	}
	log.db.tables.PlayerFormation = nil
	C.Free_PlayerFormation(log.Row)
}

func (log *PlayerFormationDeleteLog) Rollback() {
	if log.db.tables.PlayerFormation != nil {
		panic("delete rollback check failed 'player_formation'")
	}
	log.db.tables.PlayerFormation = log.Old
}

func (log *PlayerFormationUpdateLog) Rollback() {
	if log.db.tables.PlayerFormation != log.Row {
		panic("update rollback check failed 'player_formation'")
	}
	log.db.tables.PlayerFormation = log.Old
	C.Free_PlayerFormation(log.Row)
}

/* ========== player_func_key ========== */

type PlayerFuncKeyInsertLog struct {
	db *Database 
	Row *C.PlayerFuncKey
	GoRow *PlayerFuncKey
}

func (db *Database) newPlayerFuncKeyInsertLog(row *C.PlayerFuncKey, goRow *PlayerFuncKey) *PlayerFuncKeyInsertLog {
	return &PlayerFuncKeyInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerFuncKeyInsertLog) Free() {
}

func (log *PlayerFuncKeyInsertLog) InvokeHook() {
	if g_Hooks.PlayerFuncKey != nil {
		g_Hooks.PlayerFuncKey.Insert(&PlayerFuncKeyRow{row: log.Row})
	}
}

func (log *PlayerFuncKeyInsertLog) SetEventId(time.Time) {
}

func (log *PlayerFuncKeyInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(60)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.Key)
	file.WriteInt16(log.GoRow.PlayedKey)
	file.WriteInt64(log.GoRow.UniqueKey)
}

func (log *PlayerFuncKeyInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFuncKey.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Key)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.PlayedKey)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.UniqueKey)))
	return stmt.Execute()
}

func (log *PlayerFuncKeyInsertLog) CommitToTLog() {
}

func (log *PlayerFuncKeyInsertLog) CommitToXdLog() {
}

func (log *PlayerFuncKeyInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerFuncKeyDeleteLog struct {
	db *Database
	Old *C.PlayerFuncKey
	GoOld *PlayerFuncKey
}

func (db *Database) newPlayerFuncKeyDeleteLog(old *C.PlayerFuncKey, goOld *PlayerFuncKey) *PlayerFuncKeyDeleteLog {
	return &PlayerFuncKeyDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerFuncKeyDeleteLog) Free() {
	C.Free_PlayerFuncKey(log.Old)
}

func (log *PlayerFuncKeyDeleteLog) InvokeHook() {
	if g_Hooks.PlayerFuncKey != nil {
		g_Hooks.PlayerFuncKey.Delete(&PlayerFuncKeyRow{row: log.Old})
	}
}

func (log *PlayerFuncKeyDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerFuncKeyDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(60)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.Key)
	file.WriteInt16(log.GoOld.PlayedKey)
	file.WriteInt64(log.GoOld.UniqueKey)
}

func (log *PlayerFuncKeyDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFuncKey.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerFuncKeyDeleteLog) CommitToTLog() {
}

func (log *PlayerFuncKeyDeleteLog) CommitToXdLog() {
}

func (log *PlayerFuncKeyDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerFuncKeyUpdateLog struct {
	db *Database 
	Row *C.PlayerFuncKey
	Old *C.PlayerFuncKey
	GoNew *PlayerFuncKey
	GoOld *PlayerFuncKey
}

func (db *Database) newPlayerFuncKeyUpdateLog(row, old *C.PlayerFuncKey, goNew, goOld *PlayerFuncKey) *PlayerFuncKeyUpdateLog {
	return &PlayerFuncKeyUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerFuncKeyUpdateLog) Free() {
	C.Free_PlayerFuncKey(log.Old)
}

func (log *PlayerFuncKeyUpdateLog) InvokeHook() {
	if g_Hooks.PlayerFuncKey != nil {
		g_Hooks.PlayerFuncKey.Update(&PlayerFuncKeyRow{row: log.Row}, &PlayerFuncKeyRow{row: log.Old})
	}
}

func (log *PlayerFuncKeyUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerFuncKeyUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(60)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.Key)
	file.WriteInt16(log.GoNew.PlayedKey)
	file.WriteInt64(log.GoNew.UniqueKey)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.Key)
	file.WriteInt16(log.GoOld.PlayedKey)
	file.WriteInt64(log.GoOld.UniqueKey)
}

func (log *PlayerFuncKeyUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerFuncKey.Update
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Key)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.PlayedKey)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.UniqueKey)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerFuncKeyUpdateLog) CommitToTLog() {
}

func (log *PlayerFuncKeyUpdateLog) CommitToXdLog() {
}

func (log *PlayerFuncKeyUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerFuncKeyInsertLog) Rollback() {
	if log.db.tables.PlayerFuncKey != log.Row {
		panic("insert rollback check failed 'PlayerFuncKey'")
	}
	log.db.tables.PlayerFuncKey = nil
	C.Free_PlayerFuncKey(log.Row)
}

func (log *PlayerFuncKeyDeleteLog) Rollback() {
	if log.db.tables.PlayerFuncKey != nil {
		panic("delete rollback check failed 'player_func_key'")
	}
	log.db.tables.PlayerFuncKey = log.Old
}

func (log *PlayerFuncKeyUpdateLog) Rollback() {
	if log.db.tables.PlayerFuncKey != log.Row {
		panic("update rollback check failed 'player_func_key'")
	}
	log.db.tables.PlayerFuncKey = log.Old
	C.Free_PlayerFuncKey(log.Row)
}

/* ========== player_ghost ========== */

type PlayerGhostInsertLog struct {
	db *Database 
	Row *C.PlayerGhost
	GoRow *PlayerGhost
}

func (db *Database) newPlayerGhostInsertLog(row *C.PlayerGhost, goRow *PlayerGhost) *PlayerGhostInsertLog {
	return &PlayerGhostInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerGhostInsertLog) Free() {
}

func (log *PlayerGhostInsertLog) InvokeHook() {
	if g_Hooks.PlayerGhost != nil {
		g_Hooks.PlayerGhost.Insert(&PlayerGhostRow{row: log.Row})
	}
}

func (log *PlayerGhostInsertLog) SetEventId(time.Time) {
}

func (log *PlayerGhostInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(61)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.GhostId)
	file.WriteInt8(log.GoRow.Star)
	file.WriteInt16(log.GoRow.Level)
	file.WriteInt64(log.GoRow.Exp)
	file.WriteInt16(log.GoRow.Pos)
	file.WriteInt8(log.GoRow.IsNew)
	file.WriteInt8(log.GoRow.RoleId)
	file.WriteInt16(log.GoRow.SkillLevel)
	file.WriteInt64(log.GoRow.RelationId)
	file.WriteInt16(log.GoRow.AddGrowth)
}

func (log *PlayerGhostInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGhost.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.GhostId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Star)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Exp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Pos)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.IsNew)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RoleId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SkillLevel)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.RelationId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AddGrowth)))
	return stmt.Execute()
}

func (log *PlayerGhostInsertLog) CommitToTLog() {
}

func (log *PlayerGhostInsertLog) CommitToXdLog() {
}

func (log *PlayerGhostInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGhostDeleteLog struct {
	db *Database
	Old *C.PlayerGhost
	GoOld *PlayerGhost
}

func (db *Database) newPlayerGhostDeleteLog(old *C.PlayerGhost, goOld *PlayerGhost) *PlayerGhostDeleteLog {
	return &PlayerGhostDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerGhostDeleteLog) Free() {
	C.Free_PlayerGhost(log.Old)
}

func (log *PlayerGhostDeleteLog) InvokeHook() {
	if g_Hooks.PlayerGhost != nil {
		g_Hooks.PlayerGhost.Delete(&PlayerGhostRow{row: log.Old})
	}
}

func (log *PlayerGhostDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerGhostDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(61)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.GhostId)
	file.WriteInt8(log.GoOld.Star)
	file.WriteInt16(log.GoOld.Level)
	file.WriteInt64(log.GoOld.Exp)
	file.WriteInt16(log.GoOld.Pos)
	file.WriteInt8(log.GoOld.IsNew)
	file.WriteInt8(log.GoOld.RoleId)
	file.WriteInt16(log.GoOld.SkillLevel)
	file.WriteInt64(log.GoOld.RelationId)
	file.WriteInt16(log.GoOld.AddGrowth)
}

func (log *PlayerGhostDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGhost.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerGhostDeleteLog) CommitToTLog() {
}

func (log *PlayerGhostDeleteLog) CommitToXdLog() {
}

func (log *PlayerGhostDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGhostUpdateLog struct {
	db *Database 
	Row *C.PlayerGhost
	Old *C.PlayerGhost
	GoNew *PlayerGhost
	GoOld *PlayerGhost
}

func (db *Database) newPlayerGhostUpdateLog(row, old *C.PlayerGhost, goNew, goOld *PlayerGhost) *PlayerGhostUpdateLog {
	return &PlayerGhostUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerGhostUpdateLog) Free() {
	C.Free_PlayerGhost(log.Old)
}

func (log *PlayerGhostUpdateLog) InvokeHook() {
	if g_Hooks.PlayerGhost != nil {
		g_Hooks.PlayerGhost.Update(&PlayerGhostRow{row: log.Row}, &PlayerGhostRow{row: log.Old})
	}
}

func (log *PlayerGhostUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerGhostUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(61)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.GhostId)
	file.WriteInt8(log.GoNew.Star)
	file.WriteInt16(log.GoNew.Level)
	file.WriteInt64(log.GoNew.Exp)
	file.WriteInt16(log.GoNew.Pos)
	file.WriteInt8(log.GoNew.IsNew)
	file.WriteInt8(log.GoNew.RoleId)
	file.WriteInt16(log.GoNew.SkillLevel)
	file.WriteInt64(log.GoNew.RelationId)
	file.WriteInt16(log.GoNew.AddGrowth)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.GhostId)
	file.WriteInt8(log.GoOld.Star)
	file.WriteInt16(log.GoOld.Level)
	file.WriteInt64(log.GoOld.Exp)
	file.WriteInt16(log.GoOld.Pos)
	file.WriteInt8(log.GoOld.IsNew)
	file.WriteInt8(log.GoOld.RoleId)
	file.WriteInt16(log.GoOld.SkillLevel)
	file.WriteInt64(log.GoOld.RelationId)
	file.WriteInt16(log.GoOld.AddGrowth)
}

func (log *PlayerGhostUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGhost.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.GhostId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Star)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Exp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Pos)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.IsNew)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RoleId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SkillLevel)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.RelationId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AddGrowth)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerGhostUpdateLog) CommitToTLog() {
}

func (log *PlayerGhostUpdateLog) CommitToXdLog() {
}

func (log *PlayerGhostUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerGhostInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerGhost
	for crow := log.db.tables.PlayerGhost; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerGhost'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerGhost = log.db.tables.PlayerGhost.next
	}
	C.Free_PlayerGhost(log.Row)
}

func (log *PlayerGhostDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerGhost
	for crow := log.db.tables.PlayerGhost; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_ghost'")
	}
	log.Old.next = log.db.tables.PlayerGhost
	log.db.tables.PlayerGhost = log.Old
}

func (log *PlayerGhostUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerGhost
	for crow := log.db.tables.PlayerGhost; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_ghost'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerGhost = log.Old
	}
	C.Free_PlayerGhost(log.Row)
}

/* ========== player_ghost_equipment ========== */

type PlayerGhostEquipmentInsertLog struct {
	db *Database 
	Row *C.PlayerGhostEquipment
	GoRow *PlayerGhostEquipment
}

func (db *Database) newPlayerGhostEquipmentInsertLog(row *C.PlayerGhostEquipment, goRow *PlayerGhostEquipment) *PlayerGhostEquipmentInsertLog {
	return &PlayerGhostEquipmentInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerGhostEquipmentInsertLog) Free() {
}

func (log *PlayerGhostEquipmentInsertLog) InvokeHook() {
	if g_Hooks.PlayerGhostEquipment != nil {
		g_Hooks.PlayerGhostEquipment.Insert(&PlayerGhostEquipmentRow{row: log.Row})
	}
}

func (log *PlayerGhostEquipmentInsertLog) SetEventId(time.Time) {
}

func (log *PlayerGhostEquipmentInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(62)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.RoleId)
	file.WriteInt32(log.GoRow.GhostPower)
	file.WriteInt64(log.GoRow.Pos1)
	file.WriteInt64(log.GoRow.Pos2)
	file.WriteInt64(log.GoRow.Pos3)
	file.WriteInt64(log.GoRow.Pos4)
}

func (log *PlayerGhostEquipmentInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGhostEquipment.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RoleId)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.GhostPower)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos1)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos2)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos3)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos4)))
	return stmt.Execute()
}

func (log *PlayerGhostEquipmentInsertLog) CommitToTLog() {
}

func (log *PlayerGhostEquipmentInsertLog) CommitToXdLog() {
}

func (log *PlayerGhostEquipmentInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGhostEquipmentDeleteLog struct {
	db *Database
	Old *C.PlayerGhostEquipment
	GoOld *PlayerGhostEquipment
}

func (db *Database) newPlayerGhostEquipmentDeleteLog(old *C.PlayerGhostEquipment, goOld *PlayerGhostEquipment) *PlayerGhostEquipmentDeleteLog {
	return &PlayerGhostEquipmentDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerGhostEquipmentDeleteLog) Free() {
	C.Free_PlayerGhostEquipment(log.Old)
}

func (log *PlayerGhostEquipmentDeleteLog) InvokeHook() {
	if g_Hooks.PlayerGhostEquipment != nil {
		g_Hooks.PlayerGhostEquipment.Delete(&PlayerGhostEquipmentRow{row: log.Old})
	}
}

func (log *PlayerGhostEquipmentDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerGhostEquipmentDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(62)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.RoleId)
	file.WriteInt32(log.GoOld.GhostPower)
	file.WriteInt64(log.GoOld.Pos1)
	file.WriteInt64(log.GoOld.Pos2)
	file.WriteInt64(log.GoOld.Pos3)
	file.WriteInt64(log.GoOld.Pos4)
}

func (log *PlayerGhostEquipmentDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGhostEquipment.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerGhostEquipmentDeleteLog) CommitToTLog() {
}

func (log *PlayerGhostEquipmentDeleteLog) CommitToXdLog() {
}

func (log *PlayerGhostEquipmentDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGhostEquipmentUpdateLog struct {
	db *Database 
	Row *C.PlayerGhostEquipment
	Old *C.PlayerGhostEquipment
	GoNew *PlayerGhostEquipment
	GoOld *PlayerGhostEquipment
}

func (db *Database) newPlayerGhostEquipmentUpdateLog(row, old *C.PlayerGhostEquipment, goNew, goOld *PlayerGhostEquipment) *PlayerGhostEquipmentUpdateLog {
	return &PlayerGhostEquipmentUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerGhostEquipmentUpdateLog) Free() {
	C.Free_PlayerGhostEquipment(log.Old)
}

func (log *PlayerGhostEquipmentUpdateLog) InvokeHook() {
	if g_Hooks.PlayerGhostEquipment != nil {
		g_Hooks.PlayerGhostEquipment.Update(&PlayerGhostEquipmentRow{row: log.Row}, &PlayerGhostEquipmentRow{row: log.Old})
	}
}

func (log *PlayerGhostEquipmentUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerGhostEquipmentUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(62)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.RoleId)
	file.WriteInt32(log.GoNew.GhostPower)
	file.WriteInt64(log.GoNew.Pos1)
	file.WriteInt64(log.GoNew.Pos2)
	file.WriteInt64(log.GoNew.Pos3)
	file.WriteInt64(log.GoNew.Pos4)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.RoleId)
	file.WriteInt32(log.GoOld.GhostPower)
	file.WriteInt64(log.GoOld.Pos1)
	file.WriteInt64(log.GoOld.Pos2)
	file.WriteInt64(log.GoOld.Pos3)
	file.WriteInt64(log.GoOld.Pos4)
}

func (log *PlayerGhostEquipmentUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGhostEquipment.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RoleId)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.GhostPower)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos1)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos2)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos3)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos4)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerGhostEquipmentUpdateLog) CommitToTLog() {
}

func (log *PlayerGhostEquipmentUpdateLog) CommitToXdLog() {
}

func (log *PlayerGhostEquipmentUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerGhostEquipmentInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerGhostEquipment
	for crow := log.db.tables.PlayerGhostEquipment; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerGhostEquipment'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerGhostEquipment = log.db.tables.PlayerGhostEquipment.next
	}
	C.Free_PlayerGhostEquipment(log.Row)
}

func (log *PlayerGhostEquipmentDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerGhostEquipment
	for crow := log.db.tables.PlayerGhostEquipment; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_ghost_equipment'")
	}
	log.Old.next = log.db.tables.PlayerGhostEquipment
	log.db.tables.PlayerGhostEquipment = log.Old
}

func (log *PlayerGhostEquipmentUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerGhostEquipment
	for crow := log.db.tables.PlayerGhostEquipment; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_ghost_equipment'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerGhostEquipment = log.Old
	}
	C.Free_PlayerGhostEquipment(log.Row)
}

/* ========== player_ghost_state ========== */

type PlayerGhostStateInsertLog struct {
	db *Database 
	Row *C.PlayerGhostState
	GoRow *PlayerGhostState
}

func (db *Database) newPlayerGhostStateInsertLog(row *C.PlayerGhostState, goRow *PlayerGhostState) *PlayerGhostStateInsertLog {
	return &PlayerGhostStateInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerGhostStateInsertLog) Free() {
}

func (log *PlayerGhostStateInsertLog) InvokeHook() {
	if g_Hooks.PlayerGhostState != nil {
		g_Hooks.PlayerGhostState.Insert(&PlayerGhostStateRow{row: log.Row})
	}
}

func (log *PlayerGhostStateInsertLog) SetEventId(time.Time) {
}

func (log *PlayerGhostStateInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(63)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.TrainByIngotNum)
	file.WriteInt64(log.GoRow.TrainByIngotTime)
	file.WriteInt64(log.GoRow.LastFlushTime)
	file.WriteInt64(log.GoRow.GhostFightNum)
}

func (log *PlayerGhostStateInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGhostState.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.TrainByIngotNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.TrainByIngotTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastFlushTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.GhostFightNum)))
	return stmt.Execute()
}

func (log *PlayerGhostStateInsertLog) CommitToTLog() {
}

func (log *PlayerGhostStateInsertLog) CommitToXdLog() {
}

func (log *PlayerGhostStateInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGhostStateDeleteLog struct {
	db *Database
	Old *C.PlayerGhostState
	GoOld *PlayerGhostState
}

func (db *Database) newPlayerGhostStateDeleteLog(old *C.PlayerGhostState, goOld *PlayerGhostState) *PlayerGhostStateDeleteLog {
	return &PlayerGhostStateDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerGhostStateDeleteLog) Free() {
	C.Free_PlayerGhostState(log.Old)
}

func (log *PlayerGhostStateDeleteLog) InvokeHook() {
	if g_Hooks.PlayerGhostState != nil {
		g_Hooks.PlayerGhostState.Delete(&PlayerGhostStateRow{row: log.Old})
	}
}

func (log *PlayerGhostStateDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerGhostStateDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(63)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.TrainByIngotNum)
	file.WriteInt64(log.GoOld.TrainByIngotTime)
	file.WriteInt64(log.GoOld.LastFlushTime)
	file.WriteInt64(log.GoOld.GhostFightNum)
}

func (log *PlayerGhostStateDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGhostState.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerGhostStateDeleteLog) CommitToTLog() {
}

func (log *PlayerGhostStateDeleteLog) CommitToXdLog() {
}

func (log *PlayerGhostStateDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGhostStateUpdateLog struct {
	db *Database 
	Row *C.PlayerGhostState
	Old *C.PlayerGhostState
	GoNew *PlayerGhostState
	GoOld *PlayerGhostState
}

func (db *Database) newPlayerGhostStateUpdateLog(row, old *C.PlayerGhostState, goNew, goOld *PlayerGhostState) *PlayerGhostStateUpdateLog {
	return &PlayerGhostStateUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerGhostStateUpdateLog) Free() {
	C.Free_PlayerGhostState(log.Old)
}

func (log *PlayerGhostStateUpdateLog) InvokeHook() {
	if g_Hooks.PlayerGhostState != nil {
		g_Hooks.PlayerGhostState.Update(&PlayerGhostStateRow{row: log.Row}, &PlayerGhostStateRow{row: log.Old})
	}
}

func (log *PlayerGhostStateUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerGhostStateUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(63)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.TrainByIngotNum)
	file.WriteInt64(log.GoNew.TrainByIngotTime)
	file.WriteInt64(log.GoNew.LastFlushTime)
	file.WriteInt64(log.GoNew.GhostFightNum)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.TrainByIngotNum)
	file.WriteInt64(log.GoOld.TrainByIngotTime)
	file.WriteInt64(log.GoOld.LastFlushTime)
	file.WriteInt64(log.GoOld.GhostFightNum)
}

func (log *PlayerGhostStateUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGhostState.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.TrainByIngotNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.TrainByIngotTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastFlushTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.GhostFightNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerGhostStateUpdateLog) CommitToTLog() {
}

func (log *PlayerGhostStateUpdateLog) CommitToXdLog() {
}

func (log *PlayerGhostStateUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerGhostStateInsertLog) Rollback() {
	if log.db.tables.PlayerGhostState != log.Row {
		panic("insert rollback check failed 'PlayerGhostState'")
	}
	log.db.tables.PlayerGhostState = nil
	C.Free_PlayerGhostState(log.Row)
}

func (log *PlayerGhostStateDeleteLog) Rollback() {
	if log.db.tables.PlayerGhostState != nil {
		panic("delete rollback check failed 'player_ghost_state'")
	}
	log.db.tables.PlayerGhostState = log.Old
}

func (log *PlayerGhostStateUpdateLog) Rollback() {
	if log.db.tables.PlayerGhostState != log.Row {
		panic("update rollback check failed 'player_ghost_state'")
	}
	log.db.tables.PlayerGhostState = log.Old
	C.Free_PlayerGhostState(log.Row)
}

/* ========== player_global_chat_state ========== */

type PlayerGlobalChatStateInsertLog struct {
	db *Database 
	Row *C.PlayerGlobalChatState
	GoRow *PlayerGlobalChatState
}

func (db *Database) newPlayerGlobalChatStateInsertLog(row *C.PlayerGlobalChatState, goRow *PlayerGlobalChatState) *PlayerGlobalChatStateInsertLog {
	return &PlayerGlobalChatStateInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerGlobalChatStateInsertLog) Free() {
}

func (log *PlayerGlobalChatStateInsertLog) InvokeHook() {
	if g_Hooks.PlayerGlobalChatState != nil {
		g_Hooks.PlayerGlobalChatState.Insert(&PlayerGlobalChatStateRow{row: log.Row})
	}
}

func (log *PlayerGlobalChatStateInsertLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalChatStateInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(64)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.BanUntil)
}

func (log *PlayerGlobalChatStateInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalChatState.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BanUntil)))
	return stmt.Execute()
}

func (log *PlayerGlobalChatStateInsertLog) CommitToTLog() {
}

func (log *PlayerGlobalChatStateInsertLog) CommitToXdLog() {
}

func (log *PlayerGlobalChatStateInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalChatStateDeleteLog struct {
	db *Database
	Old *C.PlayerGlobalChatState
	GoOld *PlayerGlobalChatState
}

func (db *Database) newPlayerGlobalChatStateDeleteLog(old *C.PlayerGlobalChatState, goOld *PlayerGlobalChatState) *PlayerGlobalChatStateDeleteLog {
	return &PlayerGlobalChatStateDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerGlobalChatStateDeleteLog) Free() {
	C.Free_PlayerGlobalChatState(log.Old)
}

func (log *PlayerGlobalChatStateDeleteLog) InvokeHook() {
	if g_Hooks.PlayerGlobalChatState != nil {
		g_Hooks.PlayerGlobalChatState.Delete(&PlayerGlobalChatStateRow{row: log.Old})
	}
}

func (log *PlayerGlobalChatStateDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalChatStateDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(64)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.BanUntil)
}

func (log *PlayerGlobalChatStateDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalChatState.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerGlobalChatStateDeleteLog) CommitToTLog() {
}

func (log *PlayerGlobalChatStateDeleteLog) CommitToXdLog() {
}

func (log *PlayerGlobalChatStateDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalChatStateUpdateLog struct {
	db *Database 
	Row *C.PlayerGlobalChatState
	Old *C.PlayerGlobalChatState
	GoNew *PlayerGlobalChatState
	GoOld *PlayerGlobalChatState
}

func (db *Database) newPlayerGlobalChatStateUpdateLog(row, old *C.PlayerGlobalChatState, goNew, goOld *PlayerGlobalChatState) *PlayerGlobalChatStateUpdateLog {
	return &PlayerGlobalChatStateUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerGlobalChatStateUpdateLog) Free() {
	C.Free_PlayerGlobalChatState(log.Old)
}

func (log *PlayerGlobalChatStateUpdateLog) InvokeHook() {
	if g_Hooks.PlayerGlobalChatState != nil {
		g_Hooks.PlayerGlobalChatState.Update(&PlayerGlobalChatStateRow{row: log.Row}, &PlayerGlobalChatStateRow{row: log.Old})
	}
}

func (log *PlayerGlobalChatStateUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalChatStateUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(64)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.BanUntil)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.BanUntil)
}

func (log *PlayerGlobalChatStateUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalChatState.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BanUntil)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerGlobalChatStateUpdateLog) CommitToTLog() {
}

func (log *PlayerGlobalChatStateUpdateLog) CommitToXdLog() {
}

func (log *PlayerGlobalChatStateUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerGlobalChatStateInsertLog) Rollback() {
	if log.db.tables.PlayerGlobalChatState != log.Row {
		panic("insert rollback check failed 'PlayerGlobalChatState'")
	}
	log.db.tables.PlayerGlobalChatState = nil
	C.Free_PlayerGlobalChatState(log.Row)
}

func (log *PlayerGlobalChatStateDeleteLog) Rollback() {
	if log.db.tables.PlayerGlobalChatState != nil {
		panic("delete rollback check failed 'player_global_chat_state'")
	}
	log.db.tables.PlayerGlobalChatState = log.Old
}

func (log *PlayerGlobalChatStateUpdateLog) Rollback() {
	if log.db.tables.PlayerGlobalChatState != log.Row {
		panic("update rollback check failed 'player_global_chat_state'")
	}
	log.db.tables.PlayerGlobalChatState = log.Old
	C.Free_PlayerGlobalChatState(log.Row)
}

/* ========== player_global_clique_building ========== */

type PlayerGlobalCliqueBuildingInsertLog struct {
	db *Database 
	Row *C.PlayerGlobalCliqueBuilding
	GoRow *PlayerGlobalCliqueBuilding
}

func (db *Database) newPlayerGlobalCliqueBuildingInsertLog(row *C.PlayerGlobalCliqueBuilding, goRow *PlayerGlobalCliqueBuilding) *PlayerGlobalCliqueBuildingInsertLog {
	return &PlayerGlobalCliqueBuildingInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerGlobalCliqueBuildingInsertLog) Free() {
}

func (log *PlayerGlobalCliqueBuildingInsertLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueBuilding != nil {
		g_Hooks.PlayerGlobalCliqueBuilding.Insert(&PlayerGlobalCliqueBuildingRow{row: log.Row})
	}
}

func (log *PlayerGlobalCliqueBuildingInsertLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueBuildingInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(65)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.SilverExchangeTime)
	file.WriteInt16(log.GoRow.GoldExchangeNum)
	file.WriteInt16(log.GoRow.SilverExchangeNum)
	file.WriteInt64(log.GoRow.DonateCoinsCenterBuilding)
	file.WriteInt64(log.GoRow.DonateCoinsTempleBuilding)
	file.WriteInt64(log.GoRow.DonateCoinsBankBuilding)
	file.WriteInt64(log.GoRow.DonateCoinsHealthBuilding)
	file.WriteInt64(log.GoRow.DonateCoinsAttackBuilding)
	file.WriteInt64(log.GoRow.DonateCoinsDefenseBuilding)
	file.WriteInt64(log.GoRow.DonateCoinsStoreBuilding)
	file.WriteInt16(log.GoRow.HealthLevel)
	file.WriteInt16(log.GoRow.AttackLevel)
	file.WriteInt16(log.GoRow.DefenseLevel)
	file.WriteInt64(log.GoRow.WorshipTime)
	file.WriteInt64(log.GoRow.DonateCoinsCenterBuildingTime)
	file.WriteInt64(log.GoRow.GlodExchangeTime)
	file.WriteInt8(log.GoRow.WorshipType)
}

func (log *PlayerGlobalCliqueBuildingInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueBuilding.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SilverExchangeTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.GoldExchangeNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SilverExchangeNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DonateCoinsCenterBuilding)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DonateCoinsTempleBuilding)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DonateCoinsBankBuilding)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DonateCoinsHealthBuilding)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DonateCoinsAttackBuilding)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DonateCoinsDefenseBuilding)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DonateCoinsStoreBuilding)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.HealthLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AttackLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DefenseLevel)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WorshipTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DonateCoinsCenterBuildingTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.GlodExchangeTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.WorshipType)))
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueBuildingInsertLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueBuildingInsertLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueBuildingInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalCliqueBuildingDeleteLog struct {
	db *Database
	Old *C.PlayerGlobalCliqueBuilding
	GoOld *PlayerGlobalCliqueBuilding
}

func (db *Database) newPlayerGlobalCliqueBuildingDeleteLog(old *C.PlayerGlobalCliqueBuilding, goOld *PlayerGlobalCliqueBuilding) *PlayerGlobalCliqueBuildingDeleteLog {
	return &PlayerGlobalCliqueBuildingDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerGlobalCliqueBuildingDeleteLog) Free() {
	C.Free_PlayerGlobalCliqueBuilding(log.Old)
}

func (log *PlayerGlobalCliqueBuildingDeleteLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueBuilding != nil {
		g_Hooks.PlayerGlobalCliqueBuilding.Delete(&PlayerGlobalCliqueBuildingRow{row: log.Old})
	}
}

func (log *PlayerGlobalCliqueBuildingDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueBuildingDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(65)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.SilverExchangeTime)
	file.WriteInt16(log.GoOld.GoldExchangeNum)
	file.WriteInt16(log.GoOld.SilverExchangeNum)
	file.WriteInt64(log.GoOld.DonateCoinsCenterBuilding)
	file.WriteInt64(log.GoOld.DonateCoinsTempleBuilding)
	file.WriteInt64(log.GoOld.DonateCoinsBankBuilding)
	file.WriteInt64(log.GoOld.DonateCoinsHealthBuilding)
	file.WriteInt64(log.GoOld.DonateCoinsAttackBuilding)
	file.WriteInt64(log.GoOld.DonateCoinsDefenseBuilding)
	file.WriteInt64(log.GoOld.DonateCoinsStoreBuilding)
	file.WriteInt16(log.GoOld.HealthLevel)
	file.WriteInt16(log.GoOld.AttackLevel)
	file.WriteInt16(log.GoOld.DefenseLevel)
	file.WriteInt64(log.GoOld.WorshipTime)
	file.WriteInt64(log.GoOld.DonateCoinsCenterBuildingTime)
	file.WriteInt64(log.GoOld.GlodExchangeTime)
	file.WriteInt8(log.GoOld.WorshipType)
}

func (log *PlayerGlobalCliqueBuildingDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueBuilding.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueBuildingDeleteLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueBuildingDeleteLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueBuildingDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalCliqueBuildingUpdateLog struct {
	db *Database 
	Row *C.PlayerGlobalCliqueBuilding
	Old *C.PlayerGlobalCliqueBuilding
	GoNew *PlayerGlobalCliqueBuilding
	GoOld *PlayerGlobalCliqueBuilding
}

func (db *Database) newPlayerGlobalCliqueBuildingUpdateLog(row, old *C.PlayerGlobalCliqueBuilding, goNew, goOld *PlayerGlobalCliqueBuilding) *PlayerGlobalCliqueBuildingUpdateLog {
	return &PlayerGlobalCliqueBuildingUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerGlobalCliqueBuildingUpdateLog) Free() {
	C.Free_PlayerGlobalCliqueBuilding(log.Old)
}

func (log *PlayerGlobalCliqueBuildingUpdateLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueBuilding != nil {
		g_Hooks.PlayerGlobalCliqueBuilding.Update(&PlayerGlobalCliqueBuildingRow{row: log.Row}, &PlayerGlobalCliqueBuildingRow{row: log.Old})
	}
}

func (log *PlayerGlobalCliqueBuildingUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueBuildingUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(65)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.SilverExchangeTime)
	file.WriteInt16(log.GoNew.GoldExchangeNum)
	file.WriteInt16(log.GoNew.SilverExchangeNum)
	file.WriteInt64(log.GoNew.DonateCoinsCenterBuilding)
	file.WriteInt64(log.GoNew.DonateCoinsTempleBuilding)
	file.WriteInt64(log.GoNew.DonateCoinsBankBuilding)
	file.WriteInt64(log.GoNew.DonateCoinsHealthBuilding)
	file.WriteInt64(log.GoNew.DonateCoinsAttackBuilding)
	file.WriteInt64(log.GoNew.DonateCoinsDefenseBuilding)
	file.WriteInt64(log.GoNew.DonateCoinsStoreBuilding)
	file.WriteInt16(log.GoNew.HealthLevel)
	file.WriteInt16(log.GoNew.AttackLevel)
	file.WriteInt16(log.GoNew.DefenseLevel)
	file.WriteInt64(log.GoNew.WorshipTime)
	file.WriteInt64(log.GoNew.DonateCoinsCenterBuildingTime)
	file.WriteInt64(log.GoNew.GlodExchangeTime)
	file.WriteInt8(log.GoNew.WorshipType)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.SilverExchangeTime)
	file.WriteInt16(log.GoOld.GoldExchangeNum)
	file.WriteInt16(log.GoOld.SilverExchangeNum)
	file.WriteInt64(log.GoOld.DonateCoinsCenterBuilding)
	file.WriteInt64(log.GoOld.DonateCoinsTempleBuilding)
	file.WriteInt64(log.GoOld.DonateCoinsBankBuilding)
	file.WriteInt64(log.GoOld.DonateCoinsHealthBuilding)
	file.WriteInt64(log.GoOld.DonateCoinsAttackBuilding)
	file.WriteInt64(log.GoOld.DonateCoinsDefenseBuilding)
	file.WriteInt64(log.GoOld.DonateCoinsStoreBuilding)
	file.WriteInt16(log.GoOld.HealthLevel)
	file.WriteInt16(log.GoOld.AttackLevel)
	file.WriteInt16(log.GoOld.DefenseLevel)
	file.WriteInt64(log.GoOld.WorshipTime)
	file.WriteInt64(log.GoOld.DonateCoinsCenterBuildingTime)
	file.WriteInt64(log.GoOld.GlodExchangeTime)
	file.WriteInt8(log.GoOld.WorshipType)
}

func (log *PlayerGlobalCliqueBuildingUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueBuilding.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SilverExchangeTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.GoldExchangeNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SilverExchangeNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DonateCoinsCenterBuilding)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DonateCoinsTempleBuilding)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DonateCoinsBankBuilding)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DonateCoinsHealthBuilding)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DonateCoinsAttackBuilding)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DonateCoinsDefenseBuilding)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DonateCoinsStoreBuilding)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.HealthLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AttackLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DefenseLevel)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WorshipTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DonateCoinsCenterBuildingTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.GlodExchangeTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.WorshipType)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueBuildingUpdateLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueBuildingUpdateLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueBuildingUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerGlobalCliqueBuildingInsertLog) Rollback() {
	if log.db.tables.PlayerGlobalCliqueBuilding != log.Row {
		panic("insert rollback check failed 'PlayerGlobalCliqueBuilding'")
	}
	log.db.tables.PlayerGlobalCliqueBuilding = nil
	C.Free_PlayerGlobalCliqueBuilding(log.Row)
}

func (log *PlayerGlobalCliqueBuildingDeleteLog) Rollback() {
	if log.db.tables.PlayerGlobalCliqueBuilding != nil {
		panic("delete rollback check failed 'player_global_clique_building'")
	}
	log.db.tables.PlayerGlobalCliqueBuilding = log.Old
}

func (log *PlayerGlobalCliqueBuildingUpdateLog) Rollback() {
	if log.db.tables.PlayerGlobalCliqueBuilding != log.Row {
		panic("update rollback check failed 'player_global_clique_building'")
	}
	log.db.tables.PlayerGlobalCliqueBuilding = log.Old
	C.Free_PlayerGlobalCliqueBuilding(log.Row)
}

/* ========== player_global_clique_building_quest ========== */

type PlayerGlobalCliqueBuildingQuestInsertLog struct {
	db *Database 
	Row *C.PlayerGlobalCliqueBuildingQuest
	GoRow *PlayerGlobalCliqueBuildingQuest
}

func (db *Database) newPlayerGlobalCliqueBuildingQuestInsertLog(row *C.PlayerGlobalCliqueBuildingQuest, goRow *PlayerGlobalCliqueBuildingQuest) *PlayerGlobalCliqueBuildingQuestInsertLog {
	return &PlayerGlobalCliqueBuildingQuestInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerGlobalCliqueBuildingQuestInsertLog) Free() {
}

func (log *PlayerGlobalCliqueBuildingQuestInsertLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueBuildingQuest != nil {
		g_Hooks.PlayerGlobalCliqueBuildingQuest.Insert(&PlayerGlobalCliqueBuildingQuestRow{row: log.Row})
	}
}

func (log *PlayerGlobalCliqueBuildingQuestInsertLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueBuildingQuestInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(66)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.QuestId)
	file.WriteInt8(log.GoRow.AwardStatus)
	file.WriteInt16(log.GoRow.BuildingType)
}

func (log *PlayerGlobalCliqueBuildingQuestInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueBuildingQuest.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.QuestId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.AwardStatus)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.BuildingType)))
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueBuildingQuestInsertLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueBuildingQuestInsertLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueBuildingQuestInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalCliqueBuildingQuestDeleteLog struct {
	db *Database
	Old *C.PlayerGlobalCliqueBuildingQuest
	GoOld *PlayerGlobalCliqueBuildingQuest
}

func (db *Database) newPlayerGlobalCliqueBuildingQuestDeleteLog(old *C.PlayerGlobalCliqueBuildingQuest, goOld *PlayerGlobalCliqueBuildingQuest) *PlayerGlobalCliqueBuildingQuestDeleteLog {
	return &PlayerGlobalCliqueBuildingQuestDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerGlobalCliqueBuildingQuestDeleteLog) Free() {
	C.Free_PlayerGlobalCliqueBuildingQuest(log.Old)
}

func (log *PlayerGlobalCliqueBuildingQuestDeleteLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueBuildingQuest != nil {
		g_Hooks.PlayerGlobalCliqueBuildingQuest.Delete(&PlayerGlobalCliqueBuildingQuestRow{row: log.Old})
	}
}

func (log *PlayerGlobalCliqueBuildingQuestDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueBuildingQuestDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(66)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.QuestId)
	file.WriteInt8(log.GoOld.AwardStatus)
	file.WriteInt16(log.GoOld.BuildingType)
}

func (log *PlayerGlobalCliqueBuildingQuestDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueBuildingQuest.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueBuildingQuestDeleteLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueBuildingQuestDeleteLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueBuildingQuestDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalCliqueBuildingQuestUpdateLog struct {
	db *Database 
	Row *C.PlayerGlobalCliqueBuildingQuest
	Old *C.PlayerGlobalCliqueBuildingQuest
	GoNew *PlayerGlobalCliqueBuildingQuest
	GoOld *PlayerGlobalCliqueBuildingQuest
}

func (db *Database) newPlayerGlobalCliqueBuildingQuestUpdateLog(row, old *C.PlayerGlobalCliqueBuildingQuest, goNew, goOld *PlayerGlobalCliqueBuildingQuest) *PlayerGlobalCliqueBuildingQuestUpdateLog {
	return &PlayerGlobalCliqueBuildingQuestUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerGlobalCliqueBuildingQuestUpdateLog) Free() {
	C.Free_PlayerGlobalCliqueBuildingQuest(log.Old)
}

func (log *PlayerGlobalCliqueBuildingQuestUpdateLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueBuildingQuest != nil {
		g_Hooks.PlayerGlobalCliqueBuildingQuest.Update(&PlayerGlobalCliqueBuildingQuestRow{row: log.Row}, &PlayerGlobalCliqueBuildingQuestRow{row: log.Old})
	}
}

func (log *PlayerGlobalCliqueBuildingQuestUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueBuildingQuestUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(66)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.QuestId)
	file.WriteInt8(log.GoNew.AwardStatus)
	file.WriteInt16(log.GoNew.BuildingType)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.QuestId)
	file.WriteInt8(log.GoOld.AwardStatus)
	file.WriteInt16(log.GoOld.BuildingType)
}

func (log *PlayerGlobalCliqueBuildingQuestUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueBuildingQuest.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.QuestId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.AwardStatus)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.BuildingType)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueBuildingQuestUpdateLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueBuildingQuestUpdateLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueBuildingQuestUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerGlobalCliqueBuildingQuestInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerGlobalCliqueBuildingQuest
	for crow := log.db.tables.PlayerGlobalCliqueBuildingQuest; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerGlobalCliqueBuildingQuest'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerGlobalCliqueBuildingQuest = log.db.tables.PlayerGlobalCliqueBuildingQuest.next
	}
	C.Free_PlayerGlobalCliqueBuildingQuest(log.Row)
}

func (log *PlayerGlobalCliqueBuildingQuestDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerGlobalCliqueBuildingQuest
	for crow := log.db.tables.PlayerGlobalCliqueBuildingQuest; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_global_clique_building_quest'")
	}
	log.Old.next = log.db.tables.PlayerGlobalCliqueBuildingQuest
	log.db.tables.PlayerGlobalCliqueBuildingQuest = log.Old
}

func (log *PlayerGlobalCliqueBuildingQuestUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerGlobalCliqueBuildingQuest
	for crow := log.db.tables.PlayerGlobalCliqueBuildingQuest; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_global_clique_building_quest'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerGlobalCliqueBuildingQuest = log.Old
	}
	C.Free_PlayerGlobalCliqueBuildingQuest(log.Row)
}

/* ========== player_global_clique_daily_quest ========== */

type PlayerGlobalCliqueDailyQuestInsertLog struct {
	db *Database 
	Row *C.PlayerGlobalCliqueDailyQuest
	GoRow *PlayerGlobalCliqueDailyQuest
}

func (db *Database) newPlayerGlobalCliqueDailyQuestInsertLog(row *C.PlayerGlobalCliqueDailyQuest, goRow *PlayerGlobalCliqueDailyQuest) *PlayerGlobalCliqueDailyQuestInsertLog {
	return &PlayerGlobalCliqueDailyQuestInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerGlobalCliqueDailyQuestInsertLog) Free() {
}

func (log *PlayerGlobalCliqueDailyQuestInsertLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueDailyQuest != nil {
		g_Hooks.PlayerGlobalCliqueDailyQuest.Insert(&PlayerGlobalCliqueDailyQuestRow{row: log.Row})
	}
}

func (log *PlayerGlobalCliqueDailyQuestInsertLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueDailyQuestInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(67)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.QuestId)
	file.WriteInt16(log.GoRow.FinishCount)
	file.WriteInt64(log.GoRow.LastUpdateTime)
	file.WriteInt8(log.GoRow.AwardStatus)
	file.WriteInt16(log.GoRow.Class)
}

func (log *PlayerGlobalCliqueDailyQuestInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueDailyQuest.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.QuestId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.FinishCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastUpdateTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.AwardStatus)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Class)))
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueDailyQuestInsertLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueDailyQuestInsertLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueDailyQuestInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalCliqueDailyQuestDeleteLog struct {
	db *Database
	Old *C.PlayerGlobalCliqueDailyQuest
	GoOld *PlayerGlobalCliqueDailyQuest
}

func (db *Database) newPlayerGlobalCliqueDailyQuestDeleteLog(old *C.PlayerGlobalCliqueDailyQuest, goOld *PlayerGlobalCliqueDailyQuest) *PlayerGlobalCliqueDailyQuestDeleteLog {
	return &PlayerGlobalCliqueDailyQuestDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerGlobalCliqueDailyQuestDeleteLog) Free() {
	C.Free_PlayerGlobalCliqueDailyQuest(log.Old)
}

func (log *PlayerGlobalCliqueDailyQuestDeleteLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueDailyQuest != nil {
		g_Hooks.PlayerGlobalCliqueDailyQuest.Delete(&PlayerGlobalCliqueDailyQuestRow{row: log.Old})
	}
}

func (log *PlayerGlobalCliqueDailyQuestDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueDailyQuestDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(67)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.QuestId)
	file.WriteInt16(log.GoOld.FinishCount)
	file.WriteInt64(log.GoOld.LastUpdateTime)
	file.WriteInt8(log.GoOld.AwardStatus)
	file.WriteInt16(log.GoOld.Class)
}

func (log *PlayerGlobalCliqueDailyQuestDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueDailyQuest.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueDailyQuestDeleteLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueDailyQuestDeleteLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueDailyQuestDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalCliqueDailyQuestUpdateLog struct {
	db *Database 
	Row *C.PlayerGlobalCliqueDailyQuest
	Old *C.PlayerGlobalCliqueDailyQuest
	GoNew *PlayerGlobalCliqueDailyQuest
	GoOld *PlayerGlobalCliqueDailyQuest
}

func (db *Database) newPlayerGlobalCliqueDailyQuestUpdateLog(row, old *C.PlayerGlobalCliqueDailyQuest, goNew, goOld *PlayerGlobalCliqueDailyQuest) *PlayerGlobalCliqueDailyQuestUpdateLog {
	return &PlayerGlobalCliqueDailyQuestUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerGlobalCliqueDailyQuestUpdateLog) Free() {
	C.Free_PlayerGlobalCliqueDailyQuest(log.Old)
}

func (log *PlayerGlobalCliqueDailyQuestUpdateLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueDailyQuest != nil {
		g_Hooks.PlayerGlobalCliqueDailyQuest.Update(&PlayerGlobalCliqueDailyQuestRow{row: log.Row}, &PlayerGlobalCliqueDailyQuestRow{row: log.Old})
	}
}

func (log *PlayerGlobalCliqueDailyQuestUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueDailyQuestUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(67)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.QuestId)
	file.WriteInt16(log.GoNew.FinishCount)
	file.WriteInt64(log.GoNew.LastUpdateTime)
	file.WriteInt8(log.GoNew.AwardStatus)
	file.WriteInt16(log.GoNew.Class)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.QuestId)
	file.WriteInt16(log.GoOld.FinishCount)
	file.WriteInt64(log.GoOld.LastUpdateTime)
	file.WriteInt8(log.GoOld.AwardStatus)
	file.WriteInt16(log.GoOld.Class)
}

func (log *PlayerGlobalCliqueDailyQuestUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueDailyQuest.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.QuestId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.FinishCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastUpdateTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.AwardStatus)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Class)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueDailyQuestUpdateLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueDailyQuestUpdateLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueDailyQuestUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerGlobalCliqueDailyQuestInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerGlobalCliqueDailyQuest
	for crow := log.db.tables.PlayerGlobalCliqueDailyQuest; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerGlobalCliqueDailyQuest'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerGlobalCliqueDailyQuest = log.db.tables.PlayerGlobalCliqueDailyQuest.next
	}
	C.Free_PlayerGlobalCliqueDailyQuest(log.Row)
}

func (log *PlayerGlobalCliqueDailyQuestDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerGlobalCliqueDailyQuest
	for crow := log.db.tables.PlayerGlobalCliqueDailyQuest; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_global_clique_daily_quest'")
	}
	log.Old.next = log.db.tables.PlayerGlobalCliqueDailyQuest
	log.db.tables.PlayerGlobalCliqueDailyQuest = log.Old
}

func (log *PlayerGlobalCliqueDailyQuestUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerGlobalCliqueDailyQuest
	for crow := log.db.tables.PlayerGlobalCliqueDailyQuest; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_global_clique_daily_quest'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerGlobalCliqueDailyQuest = log.Old
	}
	C.Free_PlayerGlobalCliqueDailyQuest(log.Row)
}

/* ========== player_global_clique_escort ========== */

type PlayerGlobalCliqueEscortInsertLog struct {
	db *Database 
	Row *C.PlayerGlobalCliqueEscort
	GoRow *PlayerGlobalCliqueEscort
}

func (db *Database) newPlayerGlobalCliqueEscortInsertLog(row *C.PlayerGlobalCliqueEscort, goRow *PlayerGlobalCliqueEscort) *PlayerGlobalCliqueEscortInsertLog {
	return &PlayerGlobalCliqueEscortInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerGlobalCliqueEscortInsertLog) Free() {
}

func (log *PlayerGlobalCliqueEscortInsertLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueEscort != nil {
		g_Hooks.PlayerGlobalCliqueEscort.Insert(&PlayerGlobalCliqueEscortRow{row: log.Row})
	}
}

func (log *PlayerGlobalCliqueEscortInsertLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueEscortInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(68)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.DailyEscortTimestamp)
	file.WriteInt16(log.GoRow.DailyEscortNum)
	file.WriteInt16(log.GoRow.EscortBoatType)
	file.WriteInt8(log.GoRow.Status)
	file.WriteInt8(log.GoRow.Hijacked)
}

func (log *PlayerGlobalCliqueEscortInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueEscort.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DailyEscortTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyEscortNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.EscortBoatType)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Status)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Hijacked)))
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueEscortInsertLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueEscortInsertLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueEscortInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalCliqueEscortDeleteLog struct {
	db *Database
	Old *C.PlayerGlobalCliqueEscort
	GoOld *PlayerGlobalCliqueEscort
}

func (db *Database) newPlayerGlobalCliqueEscortDeleteLog(old *C.PlayerGlobalCliqueEscort, goOld *PlayerGlobalCliqueEscort) *PlayerGlobalCliqueEscortDeleteLog {
	return &PlayerGlobalCliqueEscortDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerGlobalCliqueEscortDeleteLog) Free() {
	C.Free_PlayerGlobalCliqueEscort(log.Old)
}

func (log *PlayerGlobalCliqueEscortDeleteLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueEscort != nil {
		g_Hooks.PlayerGlobalCliqueEscort.Delete(&PlayerGlobalCliqueEscortRow{row: log.Old})
	}
}

func (log *PlayerGlobalCliqueEscortDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueEscortDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(68)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.DailyEscortTimestamp)
	file.WriteInt16(log.GoOld.DailyEscortNum)
	file.WriteInt16(log.GoOld.EscortBoatType)
	file.WriteInt8(log.GoOld.Status)
	file.WriteInt8(log.GoOld.Hijacked)
}

func (log *PlayerGlobalCliqueEscortDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueEscort.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueEscortDeleteLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueEscortDeleteLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueEscortDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalCliqueEscortUpdateLog struct {
	db *Database 
	Row *C.PlayerGlobalCliqueEscort
	Old *C.PlayerGlobalCliqueEscort
	GoNew *PlayerGlobalCliqueEscort
	GoOld *PlayerGlobalCliqueEscort
}

func (db *Database) newPlayerGlobalCliqueEscortUpdateLog(row, old *C.PlayerGlobalCliqueEscort, goNew, goOld *PlayerGlobalCliqueEscort) *PlayerGlobalCliqueEscortUpdateLog {
	return &PlayerGlobalCliqueEscortUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerGlobalCliqueEscortUpdateLog) Free() {
	C.Free_PlayerGlobalCliqueEscort(log.Old)
}

func (log *PlayerGlobalCliqueEscortUpdateLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueEscort != nil {
		g_Hooks.PlayerGlobalCliqueEscort.Update(&PlayerGlobalCliqueEscortRow{row: log.Row}, &PlayerGlobalCliqueEscortRow{row: log.Old})
	}
}

func (log *PlayerGlobalCliqueEscortUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueEscortUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(68)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.DailyEscortTimestamp)
	file.WriteInt16(log.GoNew.DailyEscortNum)
	file.WriteInt16(log.GoNew.EscortBoatType)
	file.WriteInt8(log.GoNew.Status)
	file.WriteInt8(log.GoNew.Hijacked)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.DailyEscortTimestamp)
	file.WriteInt16(log.GoOld.DailyEscortNum)
	file.WriteInt16(log.GoOld.EscortBoatType)
	file.WriteInt8(log.GoOld.Status)
	file.WriteInt8(log.GoOld.Hijacked)
}

func (log *PlayerGlobalCliqueEscortUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueEscort.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DailyEscortTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyEscortNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.EscortBoatType)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Status)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Hijacked)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueEscortUpdateLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueEscortUpdateLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueEscortUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerGlobalCliqueEscortInsertLog) Rollback() {
	if log.db.tables.PlayerGlobalCliqueEscort != log.Row {
		panic("insert rollback check failed 'PlayerGlobalCliqueEscort'")
	}
	log.db.tables.PlayerGlobalCliqueEscort = nil
	C.Free_PlayerGlobalCliqueEscort(log.Row)
}

func (log *PlayerGlobalCliqueEscortDeleteLog) Rollback() {
	if log.db.tables.PlayerGlobalCliqueEscort != nil {
		panic("delete rollback check failed 'player_global_clique_escort'")
	}
	log.db.tables.PlayerGlobalCliqueEscort = log.Old
}

func (log *PlayerGlobalCliqueEscortUpdateLog) Rollback() {
	if log.db.tables.PlayerGlobalCliqueEscort != log.Row {
		panic("update rollback check failed 'player_global_clique_escort'")
	}
	log.db.tables.PlayerGlobalCliqueEscort = log.Old
	C.Free_PlayerGlobalCliqueEscort(log.Row)
}

/* ========== player_global_clique_escort_message ========== */

type PlayerGlobalCliqueEscortMessageInsertLog struct {
	db *Database 
	Row *C.PlayerGlobalCliqueEscortMessage
	GoRow *PlayerGlobalCliqueEscortMessage
}

func (db *Database) newPlayerGlobalCliqueEscortMessageInsertLog(row *C.PlayerGlobalCliqueEscortMessage, goRow *PlayerGlobalCliqueEscortMessage) *PlayerGlobalCliqueEscortMessageInsertLog {
	return &PlayerGlobalCliqueEscortMessageInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerGlobalCliqueEscortMessageInsertLog) Free() {
}

func (log *PlayerGlobalCliqueEscortMessageInsertLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueEscortMessage != nil {
		g_Hooks.PlayerGlobalCliqueEscortMessage.Insert(&PlayerGlobalCliqueEscortMessageRow{row: log.Row})
	}
}

func (log *PlayerGlobalCliqueEscortMessageInsertLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueEscortMessageInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(69)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.TplId)
	file.WriteString(log.GoRow.Parameters)
	file.WriteInt64(log.GoRow.Timestamp)
}

func (log *PlayerGlobalCliqueEscortMessageInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueEscortMessage.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TplId)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Parameters), int(log.Row.Parameters_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueEscortMessageInsertLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueEscortMessageInsertLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueEscortMessageInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalCliqueEscortMessageDeleteLog struct {
	db *Database
	Old *C.PlayerGlobalCliqueEscortMessage
	GoOld *PlayerGlobalCliqueEscortMessage
}

func (db *Database) newPlayerGlobalCliqueEscortMessageDeleteLog(old *C.PlayerGlobalCliqueEscortMessage, goOld *PlayerGlobalCliqueEscortMessage) *PlayerGlobalCliqueEscortMessageDeleteLog {
	return &PlayerGlobalCliqueEscortMessageDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerGlobalCliqueEscortMessageDeleteLog) Free() {
	C.Free_PlayerGlobalCliqueEscortMessage(log.Old)
}

func (log *PlayerGlobalCliqueEscortMessageDeleteLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueEscortMessage != nil {
		g_Hooks.PlayerGlobalCliqueEscortMessage.Delete(&PlayerGlobalCliqueEscortMessageRow{row: log.Old})
	}
}

func (log *PlayerGlobalCliqueEscortMessageDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueEscortMessageDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(69)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.TplId)
	file.WriteString(log.GoOld.Parameters)
	file.WriteInt64(log.GoOld.Timestamp)
}

func (log *PlayerGlobalCliqueEscortMessageDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueEscortMessage.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueEscortMessageDeleteLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueEscortMessageDeleteLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueEscortMessageDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalCliqueEscortMessageUpdateLog struct {
	db *Database 
	Row *C.PlayerGlobalCliqueEscortMessage
	Old *C.PlayerGlobalCliqueEscortMessage
	GoNew *PlayerGlobalCliqueEscortMessage
	GoOld *PlayerGlobalCliqueEscortMessage
}

func (db *Database) newPlayerGlobalCliqueEscortMessageUpdateLog(row, old *C.PlayerGlobalCliqueEscortMessage, goNew, goOld *PlayerGlobalCliqueEscortMessage) *PlayerGlobalCliqueEscortMessageUpdateLog {
	return &PlayerGlobalCliqueEscortMessageUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerGlobalCliqueEscortMessageUpdateLog) Free() {
	C.Free_PlayerGlobalCliqueEscortMessage(log.Old)
}

func (log *PlayerGlobalCliqueEscortMessageUpdateLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueEscortMessage != nil {
		g_Hooks.PlayerGlobalCliqueEscortMessage.Update(&PlayerGlobalCliqueEscortMessageRow{row: log.Row}, &PlayerGlobalCliqueEscortMessageRow{row: log.Old})
	}
}

func (log *PlayerGlobalCliqueEscortMessageUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueEscortMessageUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(69)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.TplId)
	file.WriteString(log.GoNew.Parameters)
	file.WriteInt64(log.GoNew.Timestamp)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.TplId)
	file.WriteString(log.GoOld.Parameters)
	file.WriteInt64(log.GoOld.Timestamp)
}

func (log *PlayerGlobalCliqueEscortMessageUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueEscortMessage.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TplId)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Parameters), int(log.Row.Parameters_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueEscortMessageUpdateLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueEscortMessageUpdateLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueEscortMessageUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerGlobalCliqueEscortMessageInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerGlobalCliqueEscortMessage
	for crow := log.db.tables.PlayerGlobalCliqueEscortMessage; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerGlobalCliqueEscortMessage'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerGlobalCliqueEscortMessage = log.db.tables.PlayerGlobalCliqueEscortMessage.next
	}
	C.Free_PlayerGlobalCliqueEscortMessage(log.Row)
}

func (log *PlayerGlobalCliqueEscortMessageDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerGlobalCliqueEscortMessage
	for crow := log.db.tables.PlayerGlobalCliqueEscortMessage; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_global_clique_escort_message'")
	}
	log.Old.next = log.db.tables.PlayerGlobalCliqueEscortMessage
	log.db.tables.PlayerGlobalCliqueEscortMessage = log.Old
}

func (log *PlayerGlobalCliqueEscortMessageUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerGlobalCliqueEscortMessage
	for crow := log.db.tables.PlayerGlobalCliqueEscortMessage; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_global_clique_escort_message'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerGlobalCliqueEscortMessage = log.Old
	}
	C.Free_PlayerGlobalCliqueEscortMessage(log.Row)
}

/* ========== player_global_clique_hijack ========== */

type PlayerGlobalCliqueHijackInsertLog struct {
	db *Database 
	Row *C.PlayerGlobalCliqueHijack
	GoRow *PlayerGlobalCliqueHijack
}

func (db *Database) newPlayerGlobalCliqueHijackInsertLog(row *C.PlayerGlobalCliqueHijack, goRow *PlayerGlobalCliqueHijack) *PlayerGlobalCliqueHijackInsertLog {
	return &PlayerGlobalCliqueHijackInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerGlobalCliqueHijackInsertLog) Free() {
}

func (log *PlayerGlobalCliqueHijackInsertLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueHijack != nil {
		g_Hooks.PlayerGlobalCliqueHijack.Insert(&PlayerGlobalCliqueHijackRow{row: log.Row})
	}
}

func (log *PlayerGlobalCliqueHijackInsertLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueHijackInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(70)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.DailyHijackTimestamp)
	file.WriteInt16(log.GoRow.DailyHijackNum)
	file.WriteInt16(log.GoRow.HijackedBoatType)
	file.WriteInt8(log.GoRow.Status)
	file.WriteInt64(log.GoRow.BuyTimestamp)
	file.WriteInt16(log.GoRow.BuyNum)
}

func (log *PlayerGlobalCliqueHijackInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueHijack.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DailyHijackTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyHijackNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.HijackedBoatType)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Status)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BuyTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.BuyNum)))
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueHijackInsertLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueHijackInsertLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueHijackInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalCliqueHijackDeleteLog struct {
	db *Database
	Old *C.PlayerGlobalCliqueHijack
	GoOld *PlayerGlobalCliqueHijack
}

func (db *Database) newPlayerGlobalCliqueHijackDeleteLog(old *C.PlayerGlobalCliqueHijack, goOld *PlayerGlobalCliqueHijack) *PlayerGlobalCliqueHijackDeleteLog {
	return &PlayerGlobalCliqueHijackDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerGlobalCliqueHijackDeleteLog) Free() {
	C.Free_PlayerGlobalCliqueHijack(log.Old)
}

func (log *PlayerGlobalCliqueHijackDeleteLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueHijack != nil {
		g_Hooks.PlayerGlobalCliqueHijack.Delete(&PlayerGlobalCliqueHijackRow{row: log.Old})
	}
}

func (log *PlayerGlobalCliqueHijackDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueHijackDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(70)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.DailyHijackTimestamp)
	file.WriteInt16(log.GoOld.DailyHijackNum)
	file.WriteInt16(log.GoOld.HijackedBoatType)
	file.WriteInt8(log.GoOld.Status)
	file.WriteInt64(log.GoOld.BuyTimestamp)
	file.WriteInt16(log.GoOld.BuyNum)
}

func (log *PlayerGlobalCliqueHijackDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueHijack.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueHijackDeleteLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueHijackDeleteLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueHijackDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalCliqueHijackUpdateLog struct {
	db *Database 
	Row *C.PlayerGlobalCliqueHijack
	Old *C.PlayerGlobalCliqueHijack
	GoNew *PlayerGlobalCliqueHijack
	GoOld *PlayerGlobalCliqueHijack
}

func (db *Database) newPlayerGlobalCliqueHijackUpdateLog(row, old *C.PlayerGlobalCliqueHijack, goNew, goOld *PlayerGlobalCliqueHijack) *PlayerGlobalCliqueHijackUpdateLog {
	return &PlayerGlobalCliqueHijackUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerGlobalCliqueHijackUpdateLog) Free() {
	C.Free_PlayerGlobalCliqueHijack(log.Old)
}

func (log *PlayerGlobalCliqueHijackUpdateLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueHijack != nil {
		g_Hooks.PlayerGlobalCliqueHijack.Update(&PlayerGlobalCliqueHijackRow{row: log.Row}, &PlayerGlobalCliqueHijackRow{row: log.Old})
	}
}

func (log *PlayerGlobalCliqueHijackUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueHijackUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(70)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.DailyHijackTimestamp)
	file.WriteInt16(log.GoNew.DailyHijackNum)
	file.WriteInt16(log.GoNew.HijackedBoatType)
	file.WriteInt8(log.GoNew.Status)
	file.WriteInt64(log.GoNew.BuyTimestamp)
	file.WriteInt16(log.GoNew.BuyNum)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.DailyHijackTimestamp)
	file.WriteInt16(log.GoOld.DailyHijackNum)
	file.WriteInt16(log.GoOld.HijackedBoatType)
	file.WriteInt8(log.GoOld.Status)
	file.WriteInt64(log.GoOld.BuyTimestamp)
	file.WriteInt16(log.GoOld.BuyNum)
}

func (log *PlayerGlobalCliqueHijackUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueHijack.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DailyHijackTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyHijackNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.HijackedBoatType)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Status)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BuyTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.BuyNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueHijackUpdateLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueHijackUpdateLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueHijackUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerGlobalCliqueHijackInsertLog) Rollback() {
	if log.db.tables.PlayerGlobalCliqueHijack != log.Row {
		panic("insert rollback check failed 'PlayerGlobalCliqueHijack'")
	}
	log.db.tables.PlayerGlobalCliqueHijack = nil
	C.Free_PlayerGlobalCliqueHijack(log.Row)
}

func (log *PlayerGlobalCliqueHijackDeleteLog) Rollback() {
	if log.db.tables.PlayerGlobalCliqueHijack != nil {
		panic("delete rollback check failed 'player_global_clique_hijack'")
	}
	log.db.tables.PlayerGlobalCliqueHijack = log.Old
}

func (log *PlayerGlobalCliqueHijackUpdateLog) Rollback() {
	if log.db.tables.PlayerGlobalCliqueHijack != log.Row {
		panic("update rollback check failed 'player_global_clique_hijack'")
	}
	log.db.tables.PlayerGlobalCliqueHijack = log.Old
	C.Free_PlayerGlobalCliqueHijack(log.Row)
}

/* ========== player_global_clique_info ========== */

type PlayerGlobalCliqueInfoInsertLog struct {
	db *Database 
	Row *C.PlayerGlobalCliqueInfo
	GoRow *PlayerGlobalCliqueInfo
}

func (db *Database) newPlayerGlobalCliqueInfoInsertLog(row *C.PlayerGlobalCliqueInfo, goRow *PlayerGlobalCliqueInfo) *PlayerGlobalCliqueInfoInsertLog {
	return &PlayerGlobalCliqueInfoInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerGlobalCliqueInfoInsertLog) Free() {
}

func (log *PlayerGlobalCliqueInfoInsertLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueInfo != nil {
		g_Hooks.PlayerGlobalCliqueInfo.Insert(&PlayerGlobalCliqueInfoRow{row: log.Row})
	}
}

func (log *PlayerGlobalCliqueInfoInsertLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueInfoInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(71)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.CliqueId)
	file.WriteInt64(log.GoRow.JoinTime)
	file.WriteInt64(log.GoRow.Contrib)
	file.WriteInt64(log.GoRow.ContribClearTime)
	file.WriteInt64(log.GoRow.DonateCoinsTime)
	file.WriteInt64(log.GoRow.DailyDonateCoins)
	file.WriteInt64(log.GoRow.TotalContrib)
}

func (log *PlayerGlobalCliqueInfoInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueInfo.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.CliqueId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.JoinTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Contrib)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ContribClearTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DonateCoinsTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DailyDonateCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.TotalContrib)))
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueInfoInsertLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueInfoInsertLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueInfoInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalCliqueInfoDeleteLog struct {
	db *Database
	Old *C.PlayerGlobalCliqueInfo
	GoOld *PlayerGlobalCliqueInfo
}

func (db *Database) newPlayerGlobalCliqueInfoDeleteLog(old *C.PlayerGlobalCliqueInfo, goOld *PlayerGlobalCliqueInfo) *PlayerGlobalCliqueInfoDeleteLog {
	return &PlayerGlobalCliqueInfoDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerGlobalCliqueInfoDeleteLog) Free() {
	C.Free_PlayerGlobalCliqueInfo(log.Old)
}

func (log *PlayerGlobalCliqueInfoDeleteLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueInfo != nil {
		g_Hooks.PlayerGlobalCliqueInfo.Delete(&PlayerGlobalCliqueInfoRow{row: log.Old})
	}
}

func (log *PlayerGlobalCliqueInfoDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueInfoDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(71)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.CliqueId)
	file.WriteInt64(log.GoOld.JoinTime)
	file.WriteInt64(log.GoOld.Contrib)
	file.WriteInt64(log.GoOld.ContribClearTime)
	file.WriteInt64(log.GoOld.DonateCoinsTime)
	file.WriteInt64(log.GoOld.DailyDonateCoins)
	file.WriteInt64(log.GoOld.TotalContrib)
}

func (log *PlayerGlobalCliqueInfoDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueInfo.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueInfoDeleteLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueInfoDeleteLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueInfoDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalCliqueInfoUpdateLog struct {
	db *Database 
	Row *C.PlayerGlobalCliqueInfo
	Old *C.PlayerGlobalCliqueInfo
	GoNew *PlayerGlobalCliqueInfo
	GoOld *PlayerGlobalCliqueInfo
}

func (db *Database) newPlayerGlobalCliqueInfoUpdateLog(row, old *C.PlayerGlobalCliqueInfo, goNew, goOld *PlayerGlobalCliqueInfo) *PlayerGlobalCliqueInfoUpdateLog {
	return &PlayerGlobalCliqueInfoUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerGlobalCliqueInfoUpdateLog) Free() {
	C.Free_PlayerGlobalCliqueInfo(log.Old)
}

func (log *PlayerGlobalCliqueInfoUpdateLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueInfo != nil {
		g_Hooks.PlayerGlobalCliqueInfo.Update(&PlayerGlobalCliqueInfoRow{row: log.Row}, &PlayerGlobalCliqueInfoRow{row: log.Old})
	}
}

func (log *PlayerGlobalCliqueInfoUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueInfoUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(71)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.CliqueId)
	file.WriteInt64(log.GoNew.JoinTime)
	file.WriteInt64(log.GoNew.Contrib)
	file.WriteInt64(log.GoNew.ContribClearTime)
	file.WriteInt64(log.GoNew.DonateCoinsTime)
	file.WriteInt64(log.GoNew.DailyDonateCoins)
	file.WriteInt64(log.GoNew.TotalContrib)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.CliqueId)
	file.WriteInt64(log.GoOld.JoinTime)
	file.WriteInt64(log.GoOld.Contrib)
	file.WriteInt64(log.GoOld.ContribClearTime)
	file.WriteInt64(log.GoOld.DonateCoinsTime)
	file.WriteInt64(log.GoOld.DailyDonateCoins)
	file.WriteInt64(log.GoOld.TotalContrib)
}

func (log *PlayerGlobalCliqueInfoUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueInfo.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.CliqueId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.JoinTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Contrib)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ContribClearTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DonateCoinsTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DailyDonateCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.TotalContrib)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueInfoUpdateLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueInfoUpdateLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueInfoUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerGlobalCliqueInfoInsertLog) Rollback() {
	if log.db.tables.PlayerGlobalCliqueInfo != log.Row {
		panic("insert rollback check failed 'PlayerGlobalCliqueInfo'")
	}
	log.db.tables.PlayerGlobalCliqueInfo = nil
	C.Free_PlayerGlobalCliqueInfo(log.Row)
}

func (log *PlayerGlobalCliqueInfoDeleteLog) Rollback() {
	if log.db.tables.PlayerGlobalCliqueInfo != nil {
		panic("delete rollback check failed 'player_global_clique_info'")
	}
	log.db.tables.PlayerGlobalCliqueInfo = log.Old
}

func (log *PlayerGlobalCliqueInfoUpdateLog) Rollback() {
	if log.db.tables.PlayerGlobalCliqueInfo != log.Row {
		panic("update rollback check failed 'player_global_clique_info'")
	}
	log.db.tables.PlayerGlobalCliqueInfo = log.Old
	C.Free_PlayerGlobalCliqueInfo(log.Row)
}

/* ========== player_global_clique_kongfu ========== */

type PlayerGlobalCliqueKongfuInsertLog struct {
	db *Database 
	Row *C.PlayerGlobalCliqueKongfu
	GoRow *PlayerGlobalCliqueKongfu
}

func (db *Database) newPlayerGlobalCliqueKongfuInsertLog(row *C.PlayerGlobalCliqueKongfu, goRow *PlayerGlobalCliqueKongfu) *PlayerGlobalCliqueKongfuInsertLog {
	return &PlayerGlobalCliqueKongfuInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerGlobalCliqueKongfuInsertLog) Free() {
}

func (log *PlayerGlobalCliqueKongfuInsertLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueKongfu != nil {
		g_Hooks.PlayerGlobalCliqueKongfu.Insert(&PlayerGlobalCliqueKongfuRow{row: log.Row})
	}
}

func (log *PlayerGlobalCliqueKongfuInsertLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueKongfuInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(72)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.KongfuId)
	file.WriteInt16(log.GoRow.Level)
}

func (log *PlayerGlobalCliqueKongfuInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueKongfu.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.KongfuId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Level)))
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueKongfuInsertLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueKongfuInsertLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueKongfuInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalCliqueKongfuDeleteLog struct {
	db *Database
	Old *C.PlayerGlobalCliqueKongfu
	GoOld *PlayerGlobalCliqueKongfu
}

func (db *Database) newPlayerGlobalCliqueKongfuDeleteLog(old *C.PlayerGlobalCliqueKongfu, goOld *PlayerGlobalCliqueKongfu) *PlayerGlobalCliqueKongfuDeleteLog {
	return &PlayerGlobalCliqueKongfuDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerGlobalCliqueKongfuDeleteLog) Free() {
	C.Free_PlayerGlobalCliqueKongfu(log.Old)
}

func (log *PlayerGlobalCliqueKongfuDeleteLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueKongfu != nil {
		g_Hooks.PlayerGlobalCliqueKongfu.Delete(&PlayerGlobalCliqueKongfuRow{row: log.Old})
	}
}

func (log *PlayerGlobalCliqueKongfuDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueKongfuDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(72)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.KongfuId)
	file.WriteInt16(log.GoOld.Level)
}

func (log *PlayerGlobalCliqueKongfuDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueKongfu.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueKongfuDeleteLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueKongfuDeleteLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueKongfuDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalCliqueKongfuUpdateLog struct {
	db *Database 
	Row *C.PlayerGlobalCliqueKongfu
	Old *C.PlayerGlobalCliqueKongfu
	GoNew *PlayerGlobalCliqueKongfu
	GoOld *PlayerGlobalCliqueKongfu
}

func (db *Database) newPlayerGlobalCliqueKongfuUpdateLog(row, old *C.PlayerGlobalCliqueKongfu, goNew, goOld *PlayerGlobalCliqueKongfu) *PlayerGlobalCliqueKongfuUpdateLog {
	return &PlayerGlobalCliqueKongfuUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerGlobalCliqueKongfuUpdateLog) Free() {
	C.Free_PlayerGlobalCliqueKongfu(log.Old)
}

func (log *PlayerGlobalCliqueKongfuUpdateLog) InvokeHook() {
	if g_Hooks.PlayerGlobalCliqueKongfu != nil {
		g_Hooks.PlayerGlobalCliqueKongfu.Update(&PlayerGlobalCliqueKongfuRow{row: log.Row}, &PlayerGlobalCliqueKongfuRow{row: log.Old})
	}
}

func (log *PlayerGlobalCliqueKongfuUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalCliqueKongfuUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(72)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.KongfuId)
	file.WriteInt16(log.GoNew.Level)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.KongfuId)
	file.WriteInt16(log.GoOld.Level)
}

func (log *PlayerGlobalCliqueKongfuUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalCliqueKongfu.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.KongfuId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerGlobalCliqueKongfuUpdateLog) CommitToTLog() {
}

func (log *PlayerGlobalCliqueKongfuUpdateLog) CommitToXdLog() {
}

func (log *PlayerGlobalCliqueKongfuUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerGlobalCliqueKongfuInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerGlobalCliqueKongfu
	for crow := log.db.tables.PlayerGlobalCliqueKongfu; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerGlobalCliqueKongfu'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerGlobalCliqueKongfu = log.db.tables.PlayerGlobalCliqueKongfu.next
	}
	C.Free_PlayerGlobalCliqueKongfu(log.Row)
}

func (log *PlayerGlobalCliqueKongfuDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerGlobalCliqueKongfu
	for crow := log.db.tables.PlayerGlobalCliqueKongfu; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_global_clique_kongfu'")
	}
	log.Old.next = log.db.tables.PlayerGlobalCliqueKongfu
	log.db.tables.PlayerGlobalCliqueKongfu = log.Old
}

func (log *PlayerGlobalCliqueKongfuUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerGlobalCliqueKongfu
	for crow := log.db.tables.PlayerGlobalCliqueKongfu; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_global_clique_kongfu'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerGlobalCliqueKongfu = log.Old
	}
	C.Free_PlayerGlobalCliqueKongfu(log.Row)
}

/* ========== player_global_friend ========== */

type PlayerGlobalFriendInsertLog struct {
	db *Database 
	Row *C.PlayerGlobalFriend
	GoRow *PlayerGlobalFriend
}

func (db *Database) newPlayerGlobalFriendInsertLog(row *C.PlayerGlobalFriend, goRow *PlayerGlobalFriend) *PlayerGlobalFriendInsertLog {
	return &PlayerGlobalFriendInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerGlobalFriendInsertLog) Free() {
}

func (log *PlayerGlobalFriendInsertLog) InvokeHook() {
	if g_Hooks.PlayerGlobalFriend != nil {
		g_Hooks.PlayerGlobalFriend.Insert(&PlayerGlobalFriendRow{row: log.Row})
	}
}

func (log *PlayerGlobalFriendInsertLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalFriendInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(73)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.FriendPid)
	file.WriteString(log.GoRow.FriendNick)
	file.WriteInt8(log.GoRow.FriendRoleId)
	file.WriteInt8(log.GoRow.FriendMode)
	file.WriteInt64(log.GoRow.LastChatTime)
	file.WriteInt64(log.GoRow.FriendTime)
	file.WriteInt64(log.GoRow.SendHeartTime)
	file.WriteInt8(log.GoRow.BlockMode)
}

func (log *PlayerGlobalFriendInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalFriend.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FriendPid)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.FriendNick), int(log.Row.FriendNick_len))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.FriendRoleId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.FriendMode)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastChatTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FriendTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SendHeartTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.BlockMode)))
	return stmt.Execute()
}

func (log *PlayerGlobalFriendInsertLog) CommitToTLog() {
}

func (log *PlayerGlobalFriendInsertLog) CommitToXdLog() {
}

func (log *PlayerGlobalFriendInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalFriendDeleteLog struct {
	db *Database
	Old *C.PlayerGlobalFriend
	GoOld *PlayerGlobalFriend
}

func (db *Database) newPlayerGlobalFriendDeleteLog(old *C.PlayerGlobalFriend, goOld *PlayerGlobalFriend) *PlayerGlobalFriendDeleteLog {
	return &PlayerGlobalFriendDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerGlobalFriendDeleteLog) Free() {
	C.Free_PlayerGlobalFriend(log.Old)
}

func (log *PlayerGlobalFriendDeleteLog) InvokeHook() {
	if g_Hooks.PlayerGlobalFriend != nil {
		g_Hooks.PlayerGlobalFriend.Delete(&PlayerGlobalFriendRow{row: log.Old})
	}
}

func (log *PlayerGlobalFriendDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalFriendDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(73)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.FriendPid)
	file.WriteString(log.GoOld.FriendNick)
	file.WriteInt8(log.GoOld.FriendRoleId)
	file.WriteInt8(log.GoOld.FriendMode)
	file.WriteInt64(log.GoOld.LastChatTime)
	file.WriteInt64(log.GoOld.FriendTime)
	file.WriteInt64(log.GoOld.SendHeartTime)
	file.WriteInt8(log.GoOld.BlockMode)
}

func (log *PlayerGlobalFriendDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalFriend.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerGlobalFriendDeleteLog) CommitToTLog() {
}

func (log *PlayerGlobalFriendDeleteLog) CommitToXdLog() {
}

func (log *PlayerGlobalFriendDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalFriendUpdateLog struct {
	db *Database 
	Row *C.PlayerGlobalFriend
	Old *C.PlayerGlobalFriend
	GoNew *PlayerGlobalFriend
	GoOld *PlayerGlobalFriend
}

func (db *Database) newPlayerGlobalFriendUpdateLog(row, old *C.PlayerGlobalFriend, goNew, goOld *PlayerGlobalFriend) *PlayerGlobalFriendUpdateLog {
	return &PlayerGlobalFriendUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerGlobalFriendUpdateLog) Free() {
	C.Free_PlayerGlobalFriend(log.Old)
}

func (log *PlayerGlobalFriendUpdateLog) InvokeHook() {
	if g_Hooks.PlayerGlobalFriend != nil {
		g_Hooks.PlayerGlobalFriend.Update(&PlayerGlobalFriendRow{row: log.Row}, &PlayerGlobalFriendRow{row: log.Old})
	}
}

func (log *PlayerGlobalFriendUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalFriendUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(73)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.FriendPid)
	file.WriteString(log.GoNew.FriendNick)
	file.WriteInt8(log.GoNew.FriendRoleId)
	file.WriteInt8(log.GoNew.FriendMode)
	file.WriteInt64(log.GoNew.LastChatTime)
	file.WriteInt64(log.GoNew.FriendTime)
	file.WriteInt64(log.GoNew.SendHeartTime)
	file.WriteInt8(log.GoNew.BlockMode)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.FriendPid)
	file.WriteString(log.GoOld.FriendNick)
	file.WriteInt8(log.GoOld.FriendRoleId)
	file.WriteInt8(log.GoOld.FriendMode)
	file.WriteInt64(log.GoOld.LastChatTime)
	file.WriteInt64(log.GoOld.FriendTime)
	file.WriteInt64(log.GoOld.SendHeartTime)
	file.WriteInt8(log.GoOld.BlockMode)
}

func (log *PlayerGlobalFriendUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalFriend.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FriendPid)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.FriendNick), int(log.Row.FriendNick_len))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.FriendRoleId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.FriendMode)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastChatTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FriendTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SendHeartTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.BlockMode)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerGlobalFriendUpdateLog) CommitToTLog() {
}

func (log *PlayerGlobalFriendUpdateLog) CommitToXdLog() {
}

func (log *PlayerGlobalFriendUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerGlobalFriendInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerGlobalFriend
	for crow := log.db.tables.PlayerGlobalFriend; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerGlobalFriend'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerGlobalFriend = log.db.tables.PlayerGlobalFriend.next
	}
	C.Free_PlayerGlobalFriend(log.Row)
}

func (log *PlayerGlobalFriendDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerGlobalFriend
	for crow := log.db.tables.PlayerGlobalFriend; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_global_friend'")
	}
	log.Old.next = log.db.tables.PlayerGlobalFriend
	log.db.tables.PlayerGlobalFriend = log.Old
}

func (log *PlayerGlobalFriendUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerGlobalFriend
	for crow := log.db.tables.PlayerGlobalFriend; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_global_friend'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerGlobalFriend = log.Old
	}
	C.Free_PlayerGlobalFriend(log.Row)
}

/* ========== player_global_friend_chat ========== */

type PlayerGlobalFriendChatInsertLog struct {
	db *Database 
	Row *C.PlayerGlobalFriendChat
	GoRow *PlayerGlobalFriendChat
}

func (db *Database) newPlayerGlobalFriendChatInsertLog(row *C.PlayerGlobalFriendChat, goRow *PlayerGlobalFriendChat) *PlayerGlobalFriendChatInsertLog {
	return &PlayerGlobalFriendChatInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerGlobalFriendChatInsertLog) Free() {
}

func (log *PlayerGlobalFriendChatInsertLog) InvokeHook() {
	if g_Hooks.PlayerGlobalFriendChat != nil {
		g_Hooks.PlayerGlobalFriendChat.Insert(&PlayerGlobalFriendChatRow{row: log.Row})
	}
}

func (log *PlayerGlobalFriendChatInsertLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalFriendChatInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(74)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.FriendPid)
	file.WriteInt8(log.GoRow.Mode)
	file.WriteInt64(log.GoRow.SendTime)
	file.WriteString(log.GoRow.Message)
}

func (log *PlayerGlobalFriendChatInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalFriendChat.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FriendPid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Mode)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SendTime)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Message), int(log.Row.Message_len))
	return stmt.Execute()
}

func (log *PlayerGlobalFriendChatInsertLog) CommitToTLog() {
}

func (log *PlayerGlobalFriendChatInsertLog) CommitToXdLog() {
}

func (log *PlayerGlobalFriendChatInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalFriendChatDeleteLog struct {
	db *Database
	Old *C.PlayerGlobalFriendChat
	GoOld *PlayerGlobalFriendChat
}

func (db *Database) newPlayerGlobalFriendChatDeleteLog(old *C.PlayerGlobalFriendChat, goOld *PlayerGlobalFriendChat) *PlayerGlobalFriendChatDeleteLog {
	return &PlayerGlobalFriendChatDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerGlobalFriendChatDeleteLog) Free() {
	C.Free_PlayerGlobalFriendChat(log.Old)
}

func (log *PlayerGlobalFriendChatDeleteLog) InvokeHook() {
	if g_Hooks.PlayerGlobalFriendChat != nil {
		g_Hooks.PlayerGlobalFriendChat.Delete(&PlayerGlobalFriendChatRow{row: log.Old})
	}
}

func (log *PlayerGlobalFriendChatDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalFriendChatDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(74)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.FriendPid)
	file.WriteInt8(log.GoOld.Mode)
	file.WriteInt64(log.GoOld.SendTime)
	file.WriteString(log.GoOld.Message)
}

func (log *PlayerGlobalFriendChatDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalFriendChat.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerGlobalFriendChatDeleteLog) CommitToTLog() {
}

func (log *PlayerGlobalFriendChatDeleteLog) CommitToXdLog() {
}

func (log *PlayerGlobalFriendChatDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalFriendChatUpdateLog struct {
	db *Database 
	Row *C.PlayerGlobalFriendChat
	Old *C.PlayerGlobalFriendChat
	GoNew *PlayerGlobalFriendChat
	GoOld *PlayerGlobalFriendChat
}

func (db *Database) newPlayerGlobalFriendChatUpdateLog(row, old *C.PlayerGlobalFriendChat, goNew, goOld *PlayerGlobalFriendChat) *PlayerGlobalFriendChatUpdateLog {
	return &PlayerGlobalFriendChatUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerGlobalFriendChatUpdateLog) Free() {
	C.Free_PlayerGlobalFriendChat(log.Old)
}

func (log *PlayerGlobalFriendChatUpdateLog) InvokeHook() {
	if g_Hooks.PlayerGlobalFriendChat != nil {
		g_Hooks.PlayerGlobalFriendChat.Update(&PlayerGlobalFriendChatRow{row: log.Row}, &PlayerGlobalFriendChatRow{row: log.Old})
	}
}

func (log *PlayerGlobalFriendChatUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalFriendChatUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(74)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.FriendPid)
	file.WriteInt8(log.GoNew.Mode)
	file.WriteInt64(log.GoNew.SendTime)
	file.WriteString(log.GoNew.Message)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.FriendPid)
	file.WriteInt8(log.GoOld.Mode)
	file.WriteInt64(log.GoOld.SendTime)
	file.WriteString(log.GoOld.Message)
}

func (log *PlayerGlobalFriendChatUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalFriendChat.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FriendPid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Mode)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SendTime)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Message), int(log.Row.Message_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerGlobalFriendChatUpdateLog) CommitToTLog() {
}

func (log *PlayerGlobalFriendChatUpdateLog) CommitToXdLog() {
}

func (log *PlayerGlobalFriendChatUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerGlobalFriendChatInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerGlobalFriendChat
	for crow := log.db.tables.PlayerGlobalFriendChat; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerGlobalFriendChat'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerGlobalFriendChat = log.db.tables.PlayerGlobalFriendChat.next
	}
	C.Free_PlayerGlobalFriendChat(log.Row)
}

func (log *PlayerGlobalFriendChatDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerGlobalFriendChat
	for crow := log.db.tables.PlayerGlobalFriendChat; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_global_friend_chat'")
	}
	log.Old.next = log.db.tables.PlayerGlobalFriendChat
	log.db.tables.PlayerGlobalFriendChat = log.Old
}

func (log *PlayerGlobalFriendChatUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerGlobalFriendChat
	for crow := log.db.tables.PlayerGlobalFriendChat; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_global_friend_chat'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerGlobalFriendChat = log.Old
	}
	C.Free_PlayerGlobalFriendChat(log.Row)
}

/* ========== player_global_friend_state ========== */

type PlayerGlobalFriendStateInsertLog struct {
	db *Database 
	Row *C.PlayerGlobalFriendState
	GoRow *PlayerGlobalFriendState
}

func (db *Database) newPlayerGlobalFriendStateInsertLog(row *C.PlayerGlobalFriendState, goRow *PlayerGlobalFriendState) *PlayerGlobalFriendStateInsertLog {
	return &PlayerGlobalFriendStateInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerGlobalFriendStateInsertLog) Free() {
}

func (log *PlayerGlobalFriendStateInsertLog) InvokeHook() {
	if g_Hooks.PlayerGlobalFriendState != nil {
		g_Hooks.PlayerGlobalFriendState.Insert(&PlayerGlobalFriendStateRow{row: log.Row})
	}
}

func (log *PlayerGlobalFriendStateInsertLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalFriendStateInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(75)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.DeleteDayCount)
	file.WriteInt64(log.GoRow.DeleteTime)
	file.WriteInt8(log.GoRow.ExistOfflineChat)
	file.WriteInt32(log.GoRow.PlatformFriendNum)
}

func (log *PlayerGlobalFriendStateInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalFriendState.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.DeleteDayCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DeleteTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.ExistOfflineChat)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.PlatformFriendNum)))
	return stmt.Execute()
}

func (log *PlayerGlobalFriendStateInsertLog) CommitToTLog() {
}

func (log *PlayerGlobalFriendStateInsertLog) CommitToXdLog() {
}

func (log *PlayerGlobalFriendStateInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalFriendStateDeleteLog struct {
	db *Database
	Old *C.PlayerGlobalFriendState
	GoOld *PlayerGlobalFriendState
}

func (db *Database) newPlayerGlobalFriendStateDeleteLog(old *C.PlayerGlobalFriendState, goOld *PlayerGlobalFriendState) *PlayerGlobalFriendStateDeleteLog {
	return &PlayerGlobalFriendStateDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerGlobalFriendStateDeleteLog) Free() {
	C.Free_PlayerGlobalFriendState(log.Old)
}

func (log *PlayerGlobalFriendStateDeleteLog) InvokeHook() {
	if g_Hooks.PlayerGlobalFriendState != nil {
		g_Hooks.PlayerGlobalFriendState.Delete(&PlayerGlobalFriendStateRow{row: log.Old})
	}
}

func (log *PlayerGlobalFriendStateDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalFriendStateDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(75)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.DeleteDayCount)
	file.WriteInt64(log.GoOld.DeleteTime)
	file.WriteInt8(log.GoOld.ExistOfflineChat)
	file.WriteInt32(log.GoOld.PlatformFriendNum)
}

func (log *PlayerGlobalFriendStateDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalFriendState.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerGlobalFriendStateDeleteLog) CommitToTLog() {
}

func (log *PlayerGlobalFriendStateDeleteLog) CommitToXdLog() {
}

func (log *PlayerGlobalFriendStateDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalFriendStateUpdateLog struct {
	db *Database 
	Row *C.PlayerGlobalFriendState
	Old *C.PlayerGlobalFriendState
	GoNew *PlayerGlobalFriendState
	GoOld *PlayerGlobalFriendState
}

func (db *Database) newPlayerGlobalFriendStateUpdateLog(row, old *C.PlayerGlobalFriendState, goNew, goOld *PlayerGlobalFriendState) *PlayerGlobalFriendStateUpdateLog {
	return &PlayerGlobalFriendStateUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerGlobalFriendStateUpdateLog) Free() {
	C.Free_PlayerGlobalFriendState(log.Old)
}

func (log *PlayerGlobalFriendStateUpdateLog) InvokeHook() {
	if g_Hooks.PlayerGlobalFriendState != nil {
		g_Hooks.PlayerGlobalFriendState.Update(&PlayerGlobalFriendStateRow{row: log.Row}, &PlayerGlobalFriendStateRow{row: log.Old})
	}
}

func (log *PlayerGlobalFriendStateUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalFriendStateUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(75)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.DeleteDayCount)
	file.WriteInt64(log.GoNew.DeleteTime)
	file.WriteInt8(log.GoNew.ExistOfflineChat)
	file.WriteInt32(log.GoNew.PlatformFriendNum)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.DeleteDayCount)
	file.WriteInt64(log.GoOld.DeleteTime)
	file.WriteInt8(log.GoOld.ExistOfflineChat)
	file.WriteInt32(log.GoOld.PlatformFriendNum)
}

func (log *PlayerGlobalFriendStateUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalFriendState.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.DeleteDayCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DeleteTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.ExistOfflineChat)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.PlatformFriendNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerGlobalFriendStateUpdateLog) CommitToTLog() {
}

func (log *PlayerGlobalFriendStateUpdateLog) CommitToXdLog() {
}

func (log *PlayerGlobalFriendStateUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerGlobalFriendStateInsertLog) Rollback() {
	if log.db.tables.PlayerGlobalFriendState != log.Row {
		panic("insert rollback check failed 'PlayerGlobalFriendState'")
	}
	log.db.tables.PlayerGlobalFriendState = nil
	C.Free_PlayerGlobalFriendState(log.Row)
}

func (log *PlayerGlobalFriendStateDeleteLog) Rollback() {
	if log.db.tables.PlayerGlobalFriendState != nil {
		panic("delete rollback check failed 'player_global_friend_state'")
	}
	log.db.tables.PlayerGlobalFriendState = log.Old
}

func (log *PlayerGlobalFriendStateUpdateLog) Rollback() {
	if log.db.tables.PlayerGlobalFriendState != log.Row {
		panic("update rollback check failed 'player_global_friend_state'")
	}
	log.db.tables.PlayerGlobalFriendState = log.Old
	C.Free_PlayerGlobalFriendState(log.Row)
}

/* ========== player_global_gift_card_record ========== */

type PlayerGlobalGiftCardRecordInsertLog struct {
	db *Database 
	Row *C.PlayerGlobalGiftCardRecord
	GoRow *PlayerGlobalGiftCardRecord
}

func (db *Database) newPlayerGlobalGiftCardRecordInsertLog(row *C.PlayerGlobalGiftCardRecord, goRow *PlayerGlobalGiftCardRecord) *PlayerGlobalGiftCardRecordInsertLog {
	return &PlayerGlobalGiftCardRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerGlobalGiftCardRecordInsertLog) Free() {
}

func (log *PlayerGlobalGiftCardRecordInsertLog) InvokeHook() {
	if g_Hooks.PlayerGlobalGiftCardRecord != nil {
		g_Hooks.PlayerGlobalGiftCardRecord.Insert(&PlayerGlobalGiftCardRecordRow{row: log.Row})
	}
}

func (log *PlayerGlobalGiftCardRecordInsertLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalGiftCardRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(76)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteString(log.GoRow.Code)
}

func (log *PlayerGlobalGiftCardRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalGiftCardRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Code), int(log.Row.Code_len))
	return stmt.Execute()
}

func (log *PlayerGlobalGiftCardRecordInsertLog) CommitToTLog() {
}

func (log *PlayerGlobalGiftCardRecordInsertLog) CommitToXdLog() {
}

func (log *PlayerGlobalGiftCardRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalGiftCardRecordDeleteLog struct {
	db *Database
	Old *C.PlayerGlobalGiftCardRecord
	GoOld *PlayerGlobalGiftCardRecord
}

func (db *Database) newPlayerGlobalGiftCardRecordDeleteLog(old *C.PlayerGlobalGiftCardRecord, goOld *PlayerGlobalGiftCardRecord) *PlayerGlobalGiftCardRecordDeleteLog {
	return &PlayerGlobalGiftCardRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerGlobalGiftCardRecordDeleteLog) Free() {
	C.Free_PlayerGlobalGiftCardRecord(log.Old)
}

func (log *PlayerGlobalGiftCardRecordDeleteLog) InvokeHook() {
	if g_Hooks.PlayerGlobalGiftCardRecord != nil {
		g_Hooks.PlayerGlobalGiftCardRecord.Delete(&PlayerGlobalGiftCardRecordRow{row: log.Old})
	}
}

func (log *PlayerGlobalGiftCardRecordDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalGiftCardRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(76)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteString(log.GoOld.Code)
}

func (log *PlayerGlobalGiftCardRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalGiftCardRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerGlobalGiftCardRecordDeleteLog) CommitToTLog() {
}

func (log *PlayerGlobalGiftCardRecordDeleteLog) CommitToXdLog() {
}

func (log *PlayerGlobalGiftCardRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalGiftCardRecordUpdateLog struct {
	db *Database 
	Row *C.PlayerGlobalGiftCardRecord
	Old *C.PlayerGlobalGiftCardRecord
	GoNew *PlayerGlobalGiftCardRecord
	GoOld *PlayerGlobalGiftCardRecord
}

func (db *Database) newPlayerGlobalGiftCardRecordUpdateLog(row, old *C.PlayerGlobalGiftCardRecord, goNew, goOld *PlayerGlobalGiftCardRecord) *PlayerGlobalGiftCardRecordUpdateLog {
	return &PlayerGlobalGiftCardRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerGlobalGiftCardRecordUpdateLog) Free() {
	C.Free_PlayerGlobalGiftCardRecord(log.Old)
}

func (log *PlayerGlobalGiftCardRecordUpdateLog) InvokeHook() {
	if g_Hooks.PlayerGlobalGiftCardRecord != nil {
		g_Hooks.PlayerGlobalGiftCardRecord.Update(&PlayerGlobalGiftCardRecordRow{row: log.Row}, &PlayerGlobalGiftCardRecordRow{row: log.Old})
	}
}

func (log *PlayerGlobalGiftCardRecordUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalGiftCardRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(76)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteString(log.GoNew.Code)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteString(log.GoOld.Code)
}

func (log *PlayerGlobalGiftCardRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalGiftCardRecord.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Code), int(log.Row.Code_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerGlobalGiftCardRecordUpdateLog) CommitToTLog() {
}

func (log *PlayerGlobalGiftCardRecordUpdateLog) CommitToXdLog() {
}

func (log *PlayerGlobalGiftCardRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerGlobalGiftCardRecordInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerGlobalGiftCardRecord
	for crow := log.db.tables.PlayerGlobalGiftCardRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerGlobalGiftCardRecord'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerGlobalGiftCardRecord = log.db.tables.PlayerGlobalGiftCardRecord.next
	}
	C.Free_PlayerGlobalGiftCardRecord(log.Row)
}

func (log *PlayerGlobalGiftCardRecordDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerGlobalGiftCardRecord
	for crow := log.db.tables.PlayerGlobalGiftCardRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_global_gift_card_record'")
	}
	log.Old.next = log.db.tables.PlayerGlobalGiftCardRecord
	log.db.tables.PlayerGlobalGiftCardRecord = log.Old
}

func (log *PlayerGlobalGiftCardRecordUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerGlobalGiftCardRecord
	for crow := log.db.tables.PlayerGlobalGiftCardRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_global_gift_card_record'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerGlobalGiftCardRecord = log.Old
	}
	C.Free_PlayerGlobalGiftCardRecord(log.Row)
}

/* ========== player_global_mail_state ========== */

type PlayerGlobalMailStateInsertLog struct {
	db *Database 
	Row *C.PlayerGlobalMailState
	GoRow *PlayerGlobalMailState
}

func (db *Database) newPlayerGlobalMailStateInsertLog(row *C.PlayerGlobalMailState, goRow *PlayerGlobalMailState) *PlayerGlobalMailStateInsertLog {
	return &PlayerGlobalMailStateInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerGlobalMailStateInsertLog) Free() {
}

func (log *PlayerGlobalMailStateInsertLog) InvokeHook() {
	if g_Hooks.PlayerGlobalMailState != nil {
		g_Hooks.PlayerGlobalMailState.Insert(&PlayerGlobalMailStateRow{row: log.Row})
	}
}

func (log *PlayerGlobalMailStateInsertLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalMailStateInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(77)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.MaxTimestamp)
}

func (log *PlayerGlobalMailStateInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalMailState.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.MaxTimestamp)))
	return stmt.Execute()
}

func (log *PlayerGlobalMailStateInsertLog) CommitToTLog() {
}

func (log *PlayerGlobalMailStateInsertLog) CommitToXdLog() {
}

func (log *PlayerGlobalMailStateInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalMailStateDeleteLog struct {
	db *Database
	Old *C.PlayerGlobalMailState
	GoOld *PlayerGlobalMailState
}

func (db *Database) newPlayerGlobalMailStateDeleteLog(old *C.PlayerGlobalMailState, goOld *PlayerGlobalMailState) *PlayerGlobalMailStateDeleteLog {
	return &PlayerGlobalMailStateDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerGlobalMailStateDeleteLog) Free() {
	C.Free_PlayerGlobalMailState(log.Old)
}

func (log *PlayerGlobalMailStateDeleteLog) InvokeHook() {
	if g_Hooks.PlayerGlobalMailState != nil {
		g_Hooks.PlayerGlobalMailState.Delete(&PlayerGlobalMailStateRow{row: log.Old})
	}
}

func (log *PlayerGlobalMailStateDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalMailStateDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(77)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.MaxTimestamp)
}

func (log *PlayerGlobalMailStateDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalMailState.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerGlobalMailStateDeleteLog) CommitToTLog() {
}

func (log *PlayerGlobalMailStateDeleteLog) CommitToXdLog() {
}

func (log *PlayerGlobalMailStateDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalMailStateUpdateLog struct {
	db *Database 
	Row *C.PlayerGlobalMailState
	Old *C.PlayerGlobalMailState
	GoNew *PlayerGlobalMailState
	GoOld *PlayerGlobalMailState
}

func (db *Database) newPlayerGlobalMailStateUpdateLog(row, old *C.PlayerGlobalMailState, goNew, goOld *PlayerGlobalMailState) *PlayerGlobalMailStateUpdateLog {
	return &PlayerGlobalMailStateUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerGlobalMailStateUpdateLog) Free() {
	C.Free_PlayerGlobalMailState(log.Old)
}

func (log *PlayerGlobalMailStateUpdateLog) InvokeHook() {
	if g_Hooks.PlayerGlobalMailState != nil {
		g_Hooks.PlayerGlobalMailState.Update(&PlayerGlobalMailStateRow{row: log.Row}, &PlayerGlobalMailStateRow{row: log.Old})
	}
}

func (log *PlayerGlobalMailStateUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalMailStateUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(77)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.MaxTimestamp)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.MaxTimestamp)
}

func (log *PlayerGlobalMailStateUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalMailState.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.MaxTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerGlobalMailStateUpdateLog) CommitToTLog() {
}

func (log *PlayerGlobalMailStateUpdateLog) CommitToXdLog() {
}

func (log *PlayerGlobalMailStateUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerGlobalMailStateInsertLog) Rollback() {
	if log.db.tables.PlayerGlobalMailState != log.Row {
		panic("insert rollback check failed 'PlayerGlobalMailState'")
	}
	log.db.tables.PlayerGlobalMailState = nil
	C.Free_PlayerGlobalMailState(log.Row)
}

func (log *PlayerGlobalMailStateDeleteLog) Rollback() {
	if log.db.tables.PlayerGlobalMailState != nil {
		panic("delete rollback check failed 'player_global_mail_state'")
	}
	log.db.tables.PlayerGlobalMailState = log.Old
}

func (log *PlayerGlobalMailStateUpdateLog) Rollback() {
	if log.db.tables.PlayerGlobalMailState != log.Row {
		panic("update rollback check failed 'player_global_mail_state'")
	}
	log.db.tables.PlayerGlobalMailState = log.Old
	C.Free_PlayerGlobalMailState(log.Row)
}

/* ========== player_global_world_chat_state ========== */

type PlayerGlobalWorldChatStateInsertLog struct {
	db *Database 
	Row *C.PlayerGlobalWorldChatState
	GoRow *PlayerGlobalWorldChatState
}

func (db *Database) newPlayerGlobalWorldChatStateInsertLog(row *C.PlayerGlobalWorldChatState, goRow *PlayerGlobalWorldChatState) *PlayerGlobalWorldChatStateInsertLog {
	return &PlayerGlobalWorldChatStateInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerGlobalWorldChatStateInsertLog) Free() {
}

func (log *PlayerGlobalWorldChatStateInsertLog) InvokeHook() {
	if g_Hooks.PlayerGlobalWorldChatState != nil {
		g_Hooks.PlayerGlobalWorldChatState.Insert(&PlayerGlobalWorldChatStateRow{row: log.Row})
	}
}

func (log *PlayerGlobalWorldChatStateInsertLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalWorldChatStateInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(78)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.Timestamp)
	file.WriteInt16(log.GoRow.DailyNum)
}

func (log *PlayerGlobalWorldChatStateInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalWorldChatState.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyNum)))
	return stmt.Execute()
}

func (log *PlayerGlobalWorldChatStateInsertLog) CommitToTLog() {
}

func (log *PlayerGlobalWorldChatStateInsertLog) CommitToXdLog() {
}

func (log *PlayerGlobalWorldChatStateInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalWorldChatStateDeleteLog struct {
	db *Database
	Old *C.PlayerGlobalWorldChatState
	GoOld *PlayerGlobalWorldChatState
}

func (db *Database) newPlayerGlobalWorldChatStateDeleteLog(old *C.PlayerGlobalWorldChatState, goOld *PlayerGlobalWorldChatState) *PlayerGlobalWorldChatStateDeleteLog {
	return &PlayerGlobalWorldChatStateDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerGlobalWorldChatStateDeleteLog) Free() {
	C.Free_PlayerGlobalWorldChatState(log.Old)
}

func (log *PlayerGlobalWorldChatStateDeleteLog) InvokeHook() {
	if g_Hooks.PlayerGlobalWorldChatState != nil {
		g_Hooks.PlayerGlobalWorldChatState.Delete(&PlayerGlobalWorldChatStateRow{row: log.Old})
	}
}

func (log *PlayerGlobalWorldChatStateDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalWorldChatStateDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(78)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt16(log.GoOld.DailyNum)
}

func (log *PlayerGlobalWorldChatStateDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalWorldChatState.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerGlobalWorldChatStateDeleteLog) CommitToTLog() {
}

func (log *PlayerGlobalWorldChatStateDeleteLog) CommitToXdLog() {
}

func (log *PlayerGlobalWorldChatStateDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerGlobalWorldChatStateUpdateLog struct {
	db *Database 
	Row *C.PlayerGlobalWorldChatState
	Old *C.PlayerGlobalWorldChatState
	GoNew *PlayerGlobalWorldChatState
	GoOld *PlayerGlobalWorldChatState
}

func (db *Database) newPlayerGlobalWorldChatStateUpdateLog(row, old *C.PlayerGlobalWorldChatState, goNew, goOld *PlayerGlobalWorldChatState) *PlayerGlobalWorldChatStateUpdateLog {
	return &PlayerGlobalWorldChatStateUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerGlobalWorldChatStateUpdateLog) Free() {
	C.Free_PlayerGlobalWorldChatState(log.Old)
}

func (log *PlayerGlobalWorldChatStateUpdateLog) InvokeHook() {
	if g_Hooks.PlayerGlobalWorldChatState != nil {
		g_Hooks.PlayerGlobalWorldChatState.Update(&PlayerGlobalWorldChatStateRow{row: log.Row}, &PlayerGlobalWorldChatStateRow{row: log.Old})
	}
}

func (log *PlayerGlobalWorldChatStateUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerGlobalWorldChatStateUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(78)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.Timestamp)
	file.WriteInt16(log.GoNew.DailyNum)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt16(log.GoOld.DailyNum)
}

func (log *PlayerGlobalWorldChatStateUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerGlobalWorldChatState.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerGlobalWorldChatStateUpdateLog) CommitToTLog() {
}

func (log *PlayerGlobalWorldChatStateUpdateLog) CommitToXdLog() {
}

func (log *PlayerGlobalWorldChatStateUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerGlobalWorldChatStateInsertLog) Rollback() {
	if log.db.tables.PlayerGlobalWorldChatState != log.Row {
		panic("insert rollback check failed 'PlayerGlobalWorldChatState'")
	}
	log.db.tables.PlayerGlobalWorldChatState = nil
	C.Free_PlayerGlobalWorldChatState(log.Row)
}

func (log *PlayerGlobalWorldChatStateDeleteLog) Rollback() {
	if log.db.tables.PlayerGlobalWorldChatState != nil {
		panic("delete rollback check failed 'player_global_world_chat_state'")
	}
	log.db.tables.PlayerGlobalWorldChatState = log.Old
}

func (log *PlayerGlobalWorldChatStateUpdateLog) Rollback() {
	if log.db.tables.PlayerGlobalWorldChatState != log.Row {
		panic("update rollback check failed 'player_global_world_chat_state'")
	}
	log.db.tables.PlayerGlobalWorldChatState = log.Old
	C.Free_PlayerGlobalWorldChatState(log.Row)
}

/* ========== player_hard_level ========== */

type PlayerHardLevelInsertLog struct {
	db *Database 
	Row *C.PlayerHardLevel
	GoRow *PlayerHardLevel
}

func (db *Database) newPlayerHardLevelInsertLog(row *C.PlayerHardLevel, goRow *PlayerHardLevel) *PlayerHardLevelInsertLog {
	return &PlayerHardLevelInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerHardLevelInsertLog) Free() {
}

func (log *PlayerHardLevelInsertLog) InvokeHook() {
	if g_Hooks.PlayerHardLevel != nil {
		g_Hooks.PlayerHardLevel.Insert(&PlayerHardLevelRow{row: log.Row})
	}
}

func (log *PlayerHardLevelInsertLog) SetEventId(time.Time) {
}

func (log *PlayerHardLevelInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(79)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.Lock)
	file.WriteInt32(log.GoRow.AwardLock)
}

func (log *PlayerHardLevelInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerHardLevel.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Lock)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.AwardLock)))
	return stmt.Execute()
}

func (log *PlayerHardLevelInsertLog) CommitToTLog() {
}

func (log *PlayerHardLevelInsertLog) CommitToXdLog() {
}

func (log *PlayerHardLevelInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerHardLevelDeleteLog struct {
	db *Database
	Old *C.PlayerHardLevel
	GoOld *PlayerHardLevel
}

func (db *Database) newPlayerHardLevelDeleteLog(old *C.PlayerHardLevel, goOld *PlayerHardLevel) *PlayerHardLevelDeleteLog {
	return &PlayerHardLevelDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerHardLevelDeleteLog) Free() {
	C.Free_PlayerHardLevel(log.Old)
}

func (log *PlayerHardLevelDeleteLog) InvokeHook() {
	if g_Hooks.PlayerHardLevel != nil {
		g_Hooks.PlayerHardLevel.Delete(&PlayerHardLevelRow{row: log.Old})
	}
}

func (log *PlayerHardLevelDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerHardLevelDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(79)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Lock)
	file.WriteInt32(log.GoOld.AwardLock)
}

func (log *PlayerHardLevelDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerHardLevel.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerHardLevelDeleteLog) CommitToTLog() {
}

func (log *PlayerHardLevelDeleteLog) CommitToXdLog() {
}

func (log *PlayerHardLevelDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerHardLevelUpdateLog struct {
	db *Database 
	Row *C.PlayerHardLevel
	Old *C.PlayerHardLevel
	GoNew *PlayerHardLevel
	GoOld *PlayerHardLevel
}

func (db *Database) newPlayerHardLevelUpdateLog(row, old *C.PlayerHardLevel, goNew, goOld *PlayerHardLevel) *PlayerHardLevelUpdateLog {
	return &PlayerHardLevelUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerHardLevelUpdateLog) Free() {
	C.Free_PlayerHardLevel(log.Old)
}

func (log *PlayerHardLevelUpdateLog) InvokeHook() {
	if g_Hooks.PlayerHardLevel != nil {
		g_Hooks.PlayerHardLevel.Update(&PlayerHardLevelRow{row: log.Row}, &PlayerHardLevelRow{row: log.Old})
	}
}

func (log *PlayerHardLevelUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerHardLevelUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(79)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.Lock)
	file.WriteInt32(log.GoNew.AwardLock)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Lock)
	file.WriteInt32(log.GoOld.AwardLock)
}

func (log *PlayerHardLevelUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerHardLevel.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.Lock)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.AwardLock)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerHardLevelUpdateLog) CommitToTLog() {
}

func (log *PlayerHardLevelUpdateLog) CommitToXdLog() {
}

func (log *PlayerHardLevelUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerHardLevelInsertLog) Rollback() {
	if log.db.tables.PlayerHardLevel != log.Row {
		panic("insert rollback check failed 'PlayerHardLevel'")
	}
	log.db.tables.PlayerHardLevel = nil
	C.Free_PlayerHardLevel(log.Row)
}

func (log *PlayerHardLevelDeleteLog) Rollback() {
	if log.db.tables.PlayerHardLevel != nil {
		panic("delete rollback check failed 'player_hard_level'")
	}
	log.db.tables.PlayerHardLevel = log.Old
}

func (log *PlayerHardLevelUpdateLog) Rollback() {
	if log.db.tables.PlayerHardLevel != log.Row {
		panic("update rollback check failed 'player_hard_level'")
	}
	log.db.tables.PlayerHardLevel = log.Old
	C.Free_PlayerHardLevel(log.Row)
}

/* ========== player_hard_level_record ========== */

type PlayerHardLevelRecordInsertLog struct {
	db *Database 
	Row *C.PlayerHardLevelRecord
	GoRow *PlayerHardLevelRecord
}

func (db *Database) newPlayerHardLevelRecordInsertLog(row *C.PlayerHardLevelRecord, goRow *PlayerHardLevelRecord) *PlayerHardLevelRecordInsertLog {
	return &PlayerHardLevelRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerHardLevelRecordInsertLog) Free() {
}

func (log *PlayerHardLevelRecordInsertLog) InvokeHook() {
	if g_Hooks.PlayerHardLevelRecord != nil {
		g_Hooks.PlayerHardLevelRecord.Insert(&PlayerHardLevelRecordRow{row: log.Row})
	}
}

func (log *PlayerHardLevelRecordInsertLog) SetEventId(time.Time) {
}

func (log *PlayerHardLevelRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(80)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.LevelId)
	file.WriteInt64(log.GoRow.OpenTime)
	file.WriteInt32(log.GoRow.Score)
	file.WriteInt8(log.GoRow.Round)
	file.WriteInt8(log.GoRow.DailyNum)
	file.WriteInt64(log.GoRow.LastEnterTime)
	file.WriteInt16(log.GoRow.BuyTimes)
	file.WriteInt64(log.GoRow.BuyUpdateTime)
}

func (log *PlayerHardLevelRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerHardLevelRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.LevelId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.OpenTime)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Score)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Round)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DailyNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastEnterTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.BuyTimes)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BuyUpdateTime)))
	return stmt.Execute()
}

func (log *PlayerHardLevelRecordInsertLog) CommitToTLog() {
}

func (log *PlayerHardLevelRecordInsertLog) CommitToXdLog() {
}

func (log *PlayerHardLevelRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerHardLevelRecordDeleteLog struct {
	db *Database
	Old *C.PlayerHardLevelRecord
	GoOld *PlayerHardLevelRecord
}

func (db *Database) newPlayerHardLevelRecordDeleteLog(old *C.PlayerHardLevelRecord, goOld *PlayerHardLevelRecord) *PlayerHardLevelRecordDeleteLog {
	return &PlayerHardLevelRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerHardLevelRecordDeleteLog) Free() {
	C.Free_PlayerHardLevelRecord(log.Old)
}

func (log *PlayerHardLevelRecordDeleteLog) InvokeHook() {
	if g_Hooks.PlayerHardLevelRecord != nil {
		g_Hooks.PlayerHardLevelRecord.Delete(&PlayerHardLevelRecordRow{row: log.Old})
	}
}

func (log *PlayerHardLevelRecordDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerHardLevelRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(80)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.LevelId)
	file.WriteInt64(log.GoOld.OpenTime)
	file.WriteInt32(log.GoOld.Score)
	file.WriteInt8(log.GoOld.Round)
	file.WriteInt8(log.GoOld.DailyNum)
	file.WriteInt64(log.GoOld.LastEnterTime)
	file.WriteInt16(log.GoOld.BuyTimes)
	file.WriteInt64(log.GoOld.BuyUpdateTime)
}

func (log *PlayerHardLevelRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerHardLevelRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerHardLevelRecordDeleteLog) CommitToTLog() {
}

func (log *PlayerHardLevelRecordDeleteLog) CommitToXdLog() {
}

func (log *PlayerHardLevelRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerHardLevelRecordUpdateLog struct {
	db *Database 
	Row *C.PlayerHardLevelRecord
	Old *C.PlayerHardLevelRecord
	GoNew *PlayerHardLevelRecord
	GoOld *PlayerHardLevelRecord
}

func (db *Database) newPlayerHardLevelRecordUpdateLog(row, old *C.PlayerHardLevelRecord, goNew, goOld *PlayerHardLevelRecord) *PlayerHardLevelRecordUpdateLog {
	return &PlayerHardLevelRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerHardLevelRecordUpdateLog) Free() {
	C.Free_PlayerHardLevelRecord(log.Old)
}

func (log *PlayerHardLevelRecordUpdateLog) InvokeHook() {
	if g_Hooks.PlayerHardLevelRecord != nil {
		g_Hooks.PlayerHardLevelRecord.Update(&PlayerHardLevelRecordRow{row: log.Row}, &PlayerHardLevelRecordRow{row: log.Old})
	}
}

func (log *PlayerHardLevelRecordUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerHardLevelRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(80)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.LevelId)
	file.WriteInt64(log.GoNew.OpenTime)
	file.WriteInt32(log.GoNew.Score)
	file.WriteInt8(log.GoNew.Round)
	file.WriteInt8(log.GoNew.DailyNum)
	file.WriteInt64(log.GoNew.LastEnterTime)
	file.WriteInt16(log.GoNew.BuyTimes)
	file.WriteInt64(log.GoNew.BuyUpdateTime)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.LevelId)
	file.WriteInt64(log.GoOld.OpenTime)
	file.WriteInt32(log.GoOld.Score)
	file.WriteInt8(log.GoOld.Round)
	file.WriteInt8(log.GoOld.DailyNum)
	file.WriteInt64(log.GoOld.LastEnterTime)
	file.WriteInt16(log.GoOld.BuyTimes)
	file.WriteInt64(log.GoOld.BuyUpdateTime)
}

func (log *PlayerHardLevelRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerHardLevelRecord.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.LevelId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.OpenTime)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Score)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Round)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DailyNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastEnterTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.BuyTimes)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BuyUpdateTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerHardLevelRecordUpdateLog) CommitToTLog() {
}

func (log *PlayerHardLevelRecordUpdateLog) CommitToXdLog() {
}

func (log *PlayerHardLevelRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerHardLevelRecordInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerHardLevelRecord
	for crow := log.db.tables.PlayerHardLevelRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerHardLevelRecord'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerHardLevelRecord = log.db.tables.PlayerHardLevelRecord.next
	}
	C.Free_PlayerHardLevelRecord(log.Row)
}

func (log *PlayerHardLevelRecordDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerHardLevelRecord
	for crow := log.db.tables.PlayerHardLevelRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_hard_level_record'")
	}
	log.Old.next = log.db.tables.PlayerHardLevelRecord
	log.db.tables.PlayerHardLevelRecord = log.Old
}

func (log *PlayerHardLevelRecordUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerHardLevelRecord
	for crow := log.db.tables.PlayerHardLevelRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_hard_level_record'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerHardLevelRecord = log.Old
	}
	C.Free_PlayerHardLevelRecord(log.Row)
}

/* ========== player_heart ========== */

type PlayerHeartInsertLog struct {
	db *Database 
	Row *C.PlayerHeart
	GoRow *PlayerHeart
}

func (db *Database) newPlayerHeartInsertLog(row *C.PlayerHeart, goRow *PlayerHeart) *PlayerHeartInsertLog {
	return &PlayerHeartInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerHeartInsertLog) Free() {
}

func (log *PlayerHeartInsertLog) InvokeHook() {
	if g_Hooks.PlayerHeart != nil {
		g_Hooks.PlayerHeart.Insert(&PlayerHeartRow{row: log.Row})
	}
}

func (log *PlayerHeartInsertLog) SetEventId(time.Time) {
}

func (log *PlayerHeartInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(81)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.Value)
	file.WriteInt64(log.GoRow.UpdateTime)
	file.WriteInt32(log.GoRow.AddDayCount)
	file.WriteInt64(log.GoRow.AddTime)
	file.WriteInt16(log.GoRow.RecoverDayCount)
}

func (log *PlayerHeartInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerHeart.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Value)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.UpdateTime)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.AddDayCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.AddTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.RecoverDayCount)))
	return stmt.Execute()
}

func (log *PlayerHeartInsertLog) CommitToTLog() {
}

func (log *PlayerHeartInsertLog) CommitToXdLog() {
}

func (log *PlayerHeartInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerHeartDeleteLog struct {
	db *Database
	Old *C.PlayerHeart
	GoOld *PlayerHeart
}

func (db *Database) newPlayerHeartDeleteLog(old *C.PlayerHeart, goOld *PlayerHeart) *PlayerHeartDeleteLog {
	return &PlayerHeartDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerHeartDeleteLog) Free() {
	C.Free_PlayerHeart(log.Old)
}

func (log *PlayerHeartDeleteLog) InvokeHook() {
	if g_Hooks.PlayerHeart != nil {
		g_Hooks.PlayerHeart.Delete(&PlayerHeartRow{row: log.Old})
	}
}

func (log *PlayerHeartDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerHeartDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(81)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.Value)
	file.WriteInt64(log.GoOld.UpdateTime)
	file.WriteInt32(log.GoOld.AddDayCount)
	file.WriteInt64(log.GoOld.AddTime)
	file.WriteInt16(log.GoOld.RecoverDayCount)
}

func (log *PlayerHeartDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerHeart.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerHeartDeleteLog) CommitToTLog() {
}

func (log *PlayerHeartDeleteLog) CommitToXdLog() {
}

func (log *PlayerHeartDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerHeartUpdateLog struct {
	db *Database 
	Row *C.PlayerHeart
	Old *C.PlayerHeart
	GoNew *PlayerHeart
	GoOld *PlayerHeart
}

func (db *Database) newPlayerHeartUpdateLog(row, old *C.PlayerHeart, goNew, goOld *PlayerHeart) *PlayerHeartUpdateLog {
	return &PlayerHeartUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerHeartUpdateLog) Free() {
	C.Free_PlayerHeart(log.Old)
}

func (log *PlayerHeartUpdateLog) InvokeHook() {
	if g_Hooks.PlayerHeart != nil {
		g_Hooks.PlayerHeart.Update(&PlayerHeartRow{row: log.Row}, &PlayerHeartRow{row: log.Old})
	}
}

func (log *PlayerHeartUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerHeartUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(81)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.Value)
	file.WriteInt64(log.GoNew.UpdateTime)
	file.WriteInt32(log.GoNew.AddDayCount)
	file.WriteInt64(log.GoNew.AddTime)
	file.WriteInt16(log.GoNew.RecoverDayCount)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.Value)
	file.WriteInt64(log.GoOld.UpdateTime)
	file.WriteInt32(log.GoOld.AddDayCount)
	file.WriteInt64(log.GoOld.AddTime)
	file.WriteInt16(log.GoOld.RecoverDayCount)
}

func (log *PlayerHeartUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerHeart.Update
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Value)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.UpdateTime)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.AddDayCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.AddTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.RecoverDayCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerHeartUpdateLog) CommitToTLog() {
}

func (log *PlayerHeartUpdateLog) CommitToXdLog() {
}

func (log *PlayerHeartUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerHeartInsertLog) Rollback() {
	if log.db.tables.PlayerHeart != log.Row {
		panic("insert rollback check failed 'PlayerHeart'")
	}
	log.db.tables.PlayerHeart = nil
	C.Free_PlayerHeart(log.Row)
}

func (log *PlayerHeartDeleteLog) Rollback() {
	if log.db.tables.PlayerHeart != nil {
		panic("delete rollback check failed 'player_heart'")
	}
	log.db.tables.PlayerHeart = log.Old
}

func (log *PlayerHeartUpdateLog) Rollback() {
	if log.db.tables.PlayerHeart != log.Row {
		panic("update rollback check failed 'player_heart'")
	}
	log.db.tables.PlayerHeart = log.Old
	C.Free_PlayerHeart(log.Row)
}

/* ========== player_heart_draw ========== */

type PlayerHeartDrawInsertLog struct {
	db *Database 
	Row *C.PlayerHeartDraw
	GoRow *PlayerHeartDraw
}

func (db *Database) newPlayerHeartDrawInsertLog(row *C.PlayerHeartDraw, goRow *PlayerHeartDraw) *PlayerHeartDrawInsertLog {
	return &PlayerHeartDrawInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerHeartDrawInsertLog) Free() {
}

func (log *PlayerHeartDrawInsertLog) InvokeHook() {
	if g_Hooks.PlayerHeartDraw != nil {
		g_Hooks.PlayerHeartDraw.Insert(&PlayerHeartDrawRow{row: log.Row})
	}
}

func (log *PlayerHeartDrawInsertLog) SetEventId(time.Time) {
}

func (log *PlayerHeartDrawInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(82)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.DrawType)
	file.WriteInt8(log.GoRow.DailyNum)
	file.WriteInt64(log.GoRow.DrawTime)
}

func (log *PlayerHeartDrawInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerHeartDraw.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DrawType)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DailyNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DrawTime)))
	return stmt.Execute()
}

func (log *PlayerHeartDrawInsertLog) CommitToTLog() {
}

func (log *PlayerHeartDrawInsertLog) CommitToXdLog() {
}

func (log *PlayerHeartDrawInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerHeartDrawDeleteLog struct {
	db *Database
	Old *C.PlayerHeartDraw
	GoOld *PlayerHeartDraw
}

func (db *Database) newPlayerHeartDrawDeleteLog(old *C.PlayerHeartDraw, goOld *PlayerHeartDraw) *PlayerHeartDrawDeleteLog {
	return &PlayerHeartDrawDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerHeartDrawDeleteLog) Free() {
	C.Free_PlayerHeartDraw(log.Old)
}

func (log *PlayerHeartDrawDeleteLog) InvokeHook() {
	if g_Hooks.PlayerHeartDraw != nil {
		g_Hooks.PlayerHeartDraw.Delete(&PlayerHeartDrawRow{row: log.Old})
	}
}

func (log *PlayerHeartDrawDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerHeartDrawDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(82)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.DrawType)
	file.WriteInt8(log.GoOld.DailyNum)
	file.WriteInt64(log.GoOld.DrawTime)
}

func (log *PlayerHeartDrawDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerHeartDraw.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerHeartDrawDeleteLog) CommitToTLog() {
}

func (log *PlayerHeartDrawDeleteLog) CommitToXdLog() {
}

func (log *PlayerHeartDrawDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerHeartDrawUpdateLog struct {
	db *Database 
	Row *C.PlayerHeartDraw
	Old *C.PlayerHeartDraw
	GoNew *PlayerHeartDraw
	GoOld *PlayerHeartDraw
}

func (db *Database) newPlayerHeartDrawUpdateLog(row, old *C.PlayerHeartDraw, goNew, goOld *PlayerHeartDraw) *PlayerHeartDrawUpdateLog {
	return &PlayerHeartDrawUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerHeartDrawUpdateLog) Free() {
	C.Free_PlayerHeartDraw(log.Old)
}

func (log *PlayerHeartDrawUpdateLog) InvokeHook() {
	if g_Hooks.PlayerHeartDraw != nil {
		g_Hooks.PlayerHeartDraw.Update(&PlayerHeartDrawRow{row: log.Row}, &PlayerHeartDrawRow{row: log.Old})
	}
}

func (log *PlayerHeartDrawUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerHeartDrawUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(82)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.DrawType)
	file.WriteInt8(log.GoNew.DailyNum)
	file.WriteInt64(log.GoNew.DrawTime)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.DrawType)
	file.WriteInt8(log.GoOld.DailyNum)
	file.WriteInt64(log.GoOld.DrawTime)
}

func (log *PlayerHeartDrawUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerHeartDraw.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DrawType)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DailyNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DrawTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerHeartDrawUpdateLog) CommitToTLog() {
}

func (log *PlayerHeartDrawUpdateLog) CommitToXdLog() {
}

func (log *PlayerHeartDrawUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerHeartDrawInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerHeartDraw
	for crow := log.db.tables.PlayerHeartDraw; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerHeartDraw'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerHeartDraw = log.db.tables.PlayerHeartDraw.next
	}
	C.Free_PlayerHeartDraw(log.Row)
}

func (log *PlayerHeartDrawDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerHeartDraw
	for crow := log.db.tables.PlayerHeartDraw; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_heart_draw'")
	}
	log.Old.next = log.db.tables.PlayerHeartDraw
	log.db.tables.PlayerHeartDraw = log.Old
}

func (log *PlayerHeartDrawUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerHeartDraw
	for crow := log.db.tables.PlayerHeartDraw; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_heart_draw'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerHeartDraw = log.Old
	}
	C.Free_PlayerHeartDraw(log.Row)
}

/* ========== player_heart_draw_card_record ========== */

type PlayerHeartDrawCardRecordInsertLog struct {
	db *Database 
	Row *C.PlayerHeartDrawCardRecord
	GoRow *PlayerHeartDrawCardRecord
}

func (db *Database) newPlayerHeartDrawCardRecordInsertLog(row *C.PlayerHeartDrawCardRecord, goRow *PlayerHeartDrawCardRecord) *PlayerHeartDrawCardRecordInsertLog {
	return &PlayerHeartDrawCardRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerHeartDrawCardRecordInsertLog) Free() {
}

func (log *PlayerHeartDrawCardRecordInsertLog) InvokeHook() {
	if g_Hooks.PlayerHeartDrawCardRecord != nil {
		g_Hooks.PlayerHeartDrawCardRecord.Insert(&PlayerHeartDrawCardRecordRow{row: log.Row})
	}
}

func (log *PlayerHeartDrawCardRecordInsertLog) SetEventId(time.Time) {
}

func (log *PlayerHeartDrawCardRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(83)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.AwardType)
	file.WriteInt16(log.GoRow.AwardNum)
	file.WriteInt16(log.GoRow.ItemId)
	file.WriteInt64(log.GoRow.DrawTime)
}

func (log *PlayerHeartDrawCardRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerHeartDrawCardRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.AwardType)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AwardNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DrawTime)))
	return stmt.Execute()
}

func (log *PlayerHeartDrawCardRecordInsertLog) CommitToTLog() {
}

func (log *PlayerHeartDrawCardRecordInsertLog) CommitToXdLog() {
}

func (log *PlayerHeartDrawCardRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerHeartDrawCardRecordDeleteLog struct {
	db *Database
	Old *C.PlayerHeartDrawCardRecord
	GoOld *PlayerHeartDrawCardRecord
}

func (db *Database) newPlayerHeartDrawCardRecordDeleteLog(old *C.PlayerHeartDrawCardRecord, goOld *PlayerHeartDrawCardRecord) *PlayerHeartDrawCardRecordDeleteLog {
	return &PlayerHeartDrawCardRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerHeartDrawCardRecordDeleteLog) Free() {
	C.Free_PlayerHeartDrawCardRecord(log.Old)
}

func (log *PlayerHeartDrawCardRecordDeleteLog) InvokeHook() {
	if g_Hooks.PlayerHeartDrawCardRecord != nil {
		g_Hooks.PlayerHeartDrawCardRecord.Delete(&PlayerHeartDrawCardRecordRow{row: log.Old})
	}
}

func (log *PlayerHeartDrawCardRecordDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerHeartDrawCardRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(83)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.AwardType)
	file.WriteInt16(log.GoOld.AwardNum)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt64(log.GoOld.DrawTime)
}

func (log *PlayerHeartDrawCardRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerHeartDrawCardRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerHeartDrawCardRecordDeleteLog) CommitToTLog() {
}

func (log *PlayerHeartDrawCardRecordDeleteLog) CommitToXdLog() {
}

func (log *PlayerHeartDrawCardRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerHeartDrawCardRecordUpdateLog struct {
	db *Database 
	Row *C.PlayerHeartDrawCardRecord
	Old *C.PlayerHeartDrawCardRecord
	GoNew *PlayerHeartDrawCardRecord
	GoOld *PlayerHeartDrawCardRecord
}

func (db *Database) newPlayerHeartDrawCardRecordUpdateLog(row, old *C.PlayerHeartDrawCardRecord, goNew, goOld *PlayerHeartDrawCardRecord) *PlayerHeartDrawCardRecordUpdateLog {
	return &PlayerHeartDrawCardRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerHeartDrawCardRecordUpdateLog) Free() {
	C.Free_PlayerHeartDrawCardRecord(log.Old)
}

func (log *PlayerHeartDrawCardRecordUpdateLog) InvokeHook() {
	if g_Hooks.PlayerHeartDrawCardRecord != nil {
		g_Hooks.PlayerHeartDrawCardRecord.Update(&PlayerHeartDrawCardRecordRow{row: log.Row}, &PlayerHeartDrawCardRecordRow{row: log.Old})
	}
}

func (log *PlayerHeartDrawCardRecordUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerHeartDrawCardRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(83)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.AwardType)
	file.WriteInt16(log.GoNew.AwardNum)
	file.WriteInt16(log.GoNew.ItemId)
	file.WriteInt64(log.GoNew.DrawTime)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.AwardType)
	file.WriteInt16(log.GoOld.AwardNum)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt64(log.GoOld.DrawTime)
}

func (log *PlayerHeartDrawCardRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerHeartDrawCardRecord.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.AwardType)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AwardNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DrawTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerHeartDrawCardRecordUpdateLog) CommitToTLog() {
}

func (log *PlayerHeartDrawCardRecordUpdateLog) CommitToXdLog() {
}

func (log *PlayerHeartDrawCardRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerHeartDrawCardRecordInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerHeartDrawCardRecord
	for crow := log.db.tables.PlayerHeartDrawCardRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerHeartDrawCardRecord'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerHeartDrawCardRecord = log.db.tables.PlayerHeartDrawCardRecord.next
	}
	C.Free_PlayerHeartDrawCardRecord(log.Row)
}

func (log *PlayerHeartDrawCardRecordDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerHeartDrawCardRecord
	for crow := log.db.tables.PlayerHeartDrawCardRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_heart_draw_card_record'")
	}
	log.Old.next = log.db.tables.PlayerHeartDrawCardRecord
	log.db.tables.PlayerHeartDrawCardRecord = log.Old
}

func (log *PlayerHeartDrawCardRecordUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerHeartDrawCardRecord
	for crow := log.db.tables.PlayerHeartDrawCardRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_heart_draw_card_record'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerHeartDrawCardRecord = log.Old
	}
	C.Free_PlayerHeartDrawCardRecord(log.Row)
}

/* ========== player_heart_draw_wheel_record ========== */

type PlayerHeartDrawWheelRecordInsertLog struct {
	db *Database 
	Row *C.PlayerHeartDrawWheelRecord
	GoRow *PlayerHeartDrawWheelRecord
}

func (db *Database) newPlayerHeartDrawWheelRecordInsertLog(row *C.PlayerHeartDrawWheelRecord, goRow *PlayerHeartDrawWheelRecord) *PlayerHeartDrawWheelRecordInsertLog {
	return &PlayerHeartDrawWheelRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerHeartDrawWheelRecordInsertLog) Free() {
}

func (log *PlayerHeartDrawWheelRecordInsertLog) InvokeHook() {
	if g_Hooks.PlayerHeartDrawWheelRecord != nil {
		g_Hooks.PlayerHeartDrawWheelRecord.Insert(&PlayerHeartDrawWheelRecordRow{row: log.Row})
	}
}

func (log *PlayerHeartDrawWheelRecordInsertLog) SetEventId(time.Time) {
}

func (log *PlayerHeartDrawWheelRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(84)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.AwardType)
	file.WriteInt16(log.GoRow.AwardNum)
	file.WriteInt16(log.GoRow.ItemId)
	file.WriteInt64(log.GoRow.DrawTime)
}

func (log *PlayerHeartDrawWheelRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerHeartDrawWheelRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.AwardType)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AwardNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DrawTime)))
	return stmt.Execute()
}

func (log *PlayerHeartDrawWheelRecordInsertLog) CommitToTLog() {
}

func (log *PlayerHeartDrawWheelRecordInsertLog) CommitToXdLog() {
}

func (log *PlayerHeartDrawWheelRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerHeartDrawWheelRecordDeleteLog struct {
	db *Database
	Old *C.PlayerHeartDrawWheelRecord
	GoOld *PlayerHeartDrawWheelRecord
}

func (db *Database) newPlayerHeartDrawWheelRecordDeleteLog(old *C.PlayerHeartDrawWheelRecord, goOld *PlayerHeartDrawWheelRecord) *PlayerHeartDrawWheelRecordDeleteLog {
	return &PlayerHeartDrawWheelRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerHeartDrawWheelRecordDeleteLog) Free() {
	C.Free_PlayerHeartDrawWheelRecord(log.Old)
}

func (log *PlayerHeartDrawWheelRecordDeleteLog) InvokeHook() {
	if g_Hooks.PlayerHeartDrawWheelRecord != nil {
		g_Hooks.PlayerHeartDrawWheelRecord.Delete(&PlayerHeartDrawWheelRecordRow{row: log.Old})
	}
}

func (log *PlayerHeartDrawWheelRecordDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerHeartDrawWheelRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(84)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.AwardType)
	file.WriteInt16(log.GoOld.AwardNum)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt64(log.GoOld.DrawTime)
}

func (log *PlayerHeartDrawWheelRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerHeartDrawWheelRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerHeartDrawWheelRecordDeleteLog) CommitToTLog() {
}

func (log *PlayerHeartDrawWheelRecordDeleteLog) CommitToXdLog() {
}

func (log *PlayerHeartDrawWheelRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerHeartDrawWheelRecordUpdateLog struct {
	db *Database 
	Row *C.PlayerHeartDrawWheelRecord
	Old *C.PlayerHeartDrawWheelRecord
	GoNew *PlayerHeartDrawWheelRecord
	GoOld *PlayerHeartDrawWheelRecord
}

func (db *Database) newPlayerHeartDrawWheelRecordUpdateLog(row, old *C.PlayerHeartDrawWheelRecord, goNew, goOld *PlayerHeartDrawWheelRecord) *PlayerHeartDrawWheelRecordUpdateLog {
	return &PlayerHeartDrawWheelRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerHeartDrawWheelRecordUpdateLog) Free() {
	C.Free_PlayerHeartDrawWheelRecord(log.Old)
}

func (log *PlayerHeartDrawWheelRecordUpdateLog) InvokeHook() {
	if g_Hooks.PlayerHeartDrawWheelRecord != nil {
		g_Hooks.PlayerHeartDrawWheelRecord.Update(&PlayerHeartDrawWheelRecordRow{row: log.Row}, &PlayerHeartDrawWheelRecordRow{row: log.Old})
	}
}

func (log *PlayerHeartDrawWheelRecordUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerHeartDrawWheelRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(84)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.AwardType)
	file.WriteInt16(log.GoNew.AwardNum)
	file.WriteInt16(log.GoNew.ItemId)
	file.WriteInt64(log.GoNew.DrawTime)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.AwardType)
	file.WriteInt16(log.GoOld.AwardNum)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt64(log.GoOld.DrawTime)
}

func (log *PlayerHeartDrawWheelRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerHeartDrawWheelRecord.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.AwardType)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AwardNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.DrawTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerHeartDrawWheelRecordUpdateLog) CommitToTLog() {
}

func (log *PlayerHeartDrawWheelRecordUpdateLog) CommitToXdLog() {
}

func (log *PlayerHeartDrawWheelRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerHeartDrawWheelRecordInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerHeartDrawWheelRecord
	for crow := log.db.tables.PlayerHeartDrawWheelRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerHeartDrawWheelRecord'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerHeartDrawWheelRecord = log.db.tables.PlayerHeartDrawWheelRecord.next
	}
	C.Free_PlayerHeartDrawWheelRecord(log.Row)
}

func (log *PlayerHeartDrawWheelRecordDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerHeartDrawWheelRecord
	for crow := log.db.tables.PlayerHeartDrawWheelRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_heart_draw_wheel_record'")
	}
	log.Old.next = log.db.tables.PlayerHeartDrawWheelRecord
	log.db.tables.PlayerHeartDrawWheelRecord = log.Old
}

func (log *PlayerHeartDrawWheelRecordUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerHeartDrawWheelRecord
	for crow := log.db.tables.PlayerHeartDrawWheelRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_heart_draw_wheel_record'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerHeartDrawWheelRecord = log.Old
	}
	C.Free_PlayerHeartDrawWheelRecord(log.Row)
}

/* ========== player_info ========== */

type PlayerInfoInsertLog struct {
	db *Database 
	Row *C.PlayerInfo
	GoRow *PlayerInfo
}

func (db *Database) newPlayerInfoInsertLog(row *C.PlayerInfo, goRow *PlayerInfo) *PlayerInfoInsertLog {
	return &PlayerInfoInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerInfoInsertLog) Free() {
}

func (log *PlayerInfoInsertLog) InvokeHook() {
	if g_Hooks.PlayerInfo != nil {
		g_Hooks.PlayerInfo.Insert(&PlayerInfoRow{row: log.Row})
	}
}

func (log *PlayerInfoInsertLog) SetEventId(time.Time) {
}

func (log *PlayerInfoInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(85)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.Ingot)
	file.WriteInt64(log.GoRow.Coins)
	file.WriteInt16(log.GoRow.NewMailNum)
	file.WriteInt64(log.GoRow.LastLoginTime)
	file.WriteInt64(log.GoRow.LastOfflineTime)
	file.WriteInt64(log.GoRow.TotalOnlineTime)
	file.WriteInt64(log.GoRow.FirstLoginTime)
	file.WriteInt16(log.GoRow.NewArenaReportNum)
	file.WriteInt64(log.GoRow.LastSkillFlush)
	file.WriteInt64(log.GoRow.SwordSoulFragment)
}

func (log *PlayerInfoInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerInfo.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Ingot)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Coins)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.NewMailNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastLoginTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastOfflineTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.TotalOnlineTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FirstLoginTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.NewArenaReportNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastSkillFlush)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SwordSoulFragment)))
	return stmt.Execute()
}

func (log *PlayerInfoInsertLog) CommitToTLog() {
}

func (log *PlayerInfoInsertLog) CommitToXdLog() {
}

func (log *PlayerInfoInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerInfoDeleteLog struct {
	db *Database
	Old *C.PlayerInfo
	GoOld *PlayerInfo
}

func (db *Database) newPlayerInfoDeleteLog(old *C.PlayerInfo, goOld *PlayerInfo) *PlayerInfoDeleteLog {
	return &PlayerInfoDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerInfoDeleteLog) Free() {
	C.Free_PlayerInfo(log.Old)
}

func (log *PlayerInfoDeleteLog) InvokeHook() {
	if g_Hooks.PlayerInfo != nil {
		g_Hooks.PlayerInfo.Delete(&PlayerInfoRow{row: log.Old})
	}
}

func (log *PlayerInfoDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerInfoDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(85)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.Ingot)
	file.WriteInt64(log.GoOld.Coins)
	file.WriteInt16(log.GoOld.NewMailNum)
	file.WriteInt64(log.GoOld.LastLoginTime)
	file.WriteInt64(log.GoOld.LastOfflineTime)
	file.WriteInt64(log.GoOld.TotalOnlineTime)
	file.WriteInt64(log.GoOld.FirstLoginTime)
	file.WriteInt16(log.GoOld.NewArenaReportNum)
	file.WriteInt64(log.GoOld.LastSkillFlush)
	file.WriteInt64(log.GoOld.SwordSoulFragment)
}

func (log *PlayerInfoDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerInfo.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerInfoDeleteLog) CommitToTLog() {
}

func (log *PlayerInfoDeleteLog) CommitToXdLog() {
}

func (log *PlayerInfoDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerInfoUpdateLog struct {
	db *Database 
	Row *C.PlayerInfo
	Old *C.PlayerInfo
	GoNew *PlayerInfo
	GoOld *PlayerInfo
}

func (db *Database) newPlayerInfoUpdateLog(row, old *C.PlayerInfo, goNew, goOld *PlayerInfo) *PlayerInfoUpdateLog {
	return &PlayerInfoUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerInfoUpdateLog) Free() {
	C.Free_PlayerInfo(log.Old)
}

func (log *PlayerInfoUpdateLog) InvokeHook() {
	if g_Hooks.PlayerInfo != nil {
		g_Hooks.PlayerInfo.Update(&PlayerInfoRow{row: log.Row}, &PlayerInfoRow{row: log.Old})
	}
}

func (log *PlayerInfoUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerInfoUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(85)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.Ingot)
	file.WriteInt64(log.GoNew.Coins)
	file.WriteInt16(log.GoNew.NewMailNum)
	file.WriteInt64(log.GoNew.LastLoginTime)
	file.WriteInt64(log.GoNew.LastOfflineTime)
	file.WriteInt64(log.GoNew.TotalOnlineTime)
	file.WriteInt64(log.GoNew.FirstLoginTime)
	file.WriteInt16(log.GoNew.NewArenaReportNum)
	file.WriteInt64(log.GoNew.LastSkillFlush)
	file.WriteInt64(log.GoNew.SwordSoulFragment)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.Ingot)
	file.WriteInt64(log.GoOld.Coins)
	file.WriteInt16(log.GoOld.NewMailNum)
	file.WriteInt64(log.GoOld.LastLoginTime)
	file.WriteInt64(log.GoOld.LastOfflineTime)
	file.WriteInt64(log.GoOld.TotalOnlineTime)
	file.WriteInt64(log.GoOld.FirstLoginTime)
	file.WriteInt16(log.GoOld.NewArenaReportNum)
	file.WriteInt64(log.GoOld.LastSkillFlush)
	file.WriteInt64(log.GoOld.SwordSoulFragment)
}

func (log *PlayerInfoUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerInfo.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Ingot)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Coins)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.NewMailNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastLoginTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastOfflineTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.TotalOnlineTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FirstLoginTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.NewArenaReportNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastSkillFlush)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SwordSoulFragment)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerInfoUpdateLog) CommitToTLog() {
}

func (log *PlayerInfoUpdateLog) CommitToXdLog() {
}

func (log *PlayerInfoUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerInfoInsertLog) Rollback() {
	if log.db.tables.PlayerInfo != log.Row {
		panic("insert rollback check failed 'PlayerInfo'")
	}
	log.db.tables.PlayerInfo = nil
	C.Free_PlayerInfo(log.Row)
}

func (log *PlayerInfoDeleteLog) Rollback() {
	if log.db.tables.PlayerInfo != nil {
		panic("delete rollback check failed 'player_info'")
	}
	log.db.tables.PlayerInfo = log.Old
}

func (log *PlayerInfoUpdateLog) Rollback() {
	if log.db.tables.PlayerInfo != log.Row {
		panic("update rollback check failed 'player_info'")
	}
	log.db.tables.PlayerInfo = log.Old
	C.Free_PlayerInfo(log.Row)
}

/* ========== player_is_operated ========== */

type PlayerIsOperatedInsertLog struct {
	db *Database 
	Row *C.PlayerIsOperated
	GoRow *PlayerIsOperated
}

func (db *Database) newPlayerIsOperatedInsertLog(row *C.PlayerIsOperated, goRow *PlayerIsOperated) *PlayerIsOperatedInsertLog {
	return &PlayerIsOperatedInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerIsOperatedInsertLog) Free() {
}

func (log *PlayerIsOperatedInsertLog) InvokeHook() {
	if g_Hooks.PlayerIsOperated != nil {
		g_Hooks.PlayerIsOperated.Insert(&PlayerIsOperatedRow{row: log.Row})
	}
}

func (log *PlayerIsOperatedInsertLog) SetEventId(time.Time) {
}

func (log *PlayerIsOperatedInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(86)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.OperateValue)
}

func (log *PlayerIsOperatedInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerIsOperated.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.OperateValue)))
	return stmt.Execute()
}

func (log *PlayerIsOperatedInsertLog) CommitToTLog() {
}

func (log *PlayerIsOperatedInsertLog) CommitToXdLog() {
}

func (log *PlayerIsOperatedInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerIsOperatedDeleteLog struct {
	db *Database
	Old *C.PlayerIsOperated
	GoOld *PlayerIsOperated
}

func (db *Database) newPlayerIsOperatedDeleteLog(old *C.PlayerIsOperated, goOld *PlayerIsOperated) *PlayerIsOperatedDeleteLog {
	return &PlayerIsOperatedDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerIsOperatedDeleteLog) Free() {
	C.Free_PlayerIsOperated(log.Old)
}

func (log *PlayerIsOperatedDeleteLog) InvokeHook() {
	if g_Hooks.PlayerIsOperated != nil {
		g_Hooks.PlayerIsOperated.Delete(&PlayerIsOperatedRow{row: log.Old})
	}
}

func (log *PlayerIsOperatedDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerIsOperatedDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(86)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.OperateValue)
}

func (log *PlayerIsOperatedDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerIsOperated.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerIsOperatedDeleteLog) CommitToTLog() {
}

func (log *PlayerIsOperatedDeleteLog) CommitToXdLog() {
}

func (log *PlayerIsOperatedDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerIsOperatedUpdateLog struct {
	db *Database 
	Row *C.PlayerIsOperated
	Old *C.PlayerIsOperated
	GoNew *PlayerIsOperated
	GoOld *PlayerIsOperated
}

func (db *Database) newPlayerIsOperatedUpdateLog(row, old *C.PlayerIsOperated, goNew, goOld *PlayerIsOperated) *PlayerIsOperatedUpdateLog {
	return &PlayerIsOperatedUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerIsOperatedUpdateLog) Free() {
	C.Free_PlayerIsOperated(log.Old)
}

func (log *PlayerIsOperatedUpdateLog) InvokeHook() {
	if g_Hooks.PlayerIsOperated != nil {
		g_Hooks.PlayerIsOperated.Update(&PlayerIsOperatedRow{row: log.Row}, &PlayerIsOperatedRow{row: log.Old})
	}
}

func (log *PlayerIsOperatedUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerIsOperatedUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(86)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.OperateValue)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.OperateValue)
}

func (log *PlayerIsOperatedUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerIsOperated.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.OperateValue)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerIsOperatedUpdateLog) CommitToTLog() {
}

func (log *PlayerIsOperatedUpdateLog) CommitToXdLog() {
}

func (log *PlayerIsOperatedUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerIsOperatedInsertLog) Rollback() {
	if log.db.tables.PlayerIsOperated != log.Row {
		panic("insert rollback check failed 'PlayerIsOperated'")
	}
	log.db.tables.PlayerIsOperated = nil
	C.Free_PlayerIsOperated(log.Row)
}

func (log *PlayerIsOperatedDeleteLog) Rollback() {
	if log.db.tables.PlayerIsOperated != nil {
		panic("delete rollback check failed 'player_is_operated'")
	}
	log.db.tables.PlayerIsOperated = log.Old
}

func (log *PlayerIsOperatedUpdateLog) Rollback() {
	if log.db.tables.PlayerIsOperated != log.Row {
		panic("update rollback check failed 'player_is_operated'")
	}
	log.db.tables.PlayerIsOperated = log.Old
	C.Free_PlayerIsOperated(log.Row)
}

/* ========== player_item ========== */

type PlayerItemInsertLog struct {
	db *Database 
	Row *C.PlayerItem
	GoRow *PlayerItem
}

func (db *Database) newPlayerItemInsertLog(row *C.PlayerItem, goRow *PlayerItem) *PlayerItemInsertLog {
	return &PlayerItemInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerItemInsertLog) Free() {
}

func (log *PlayerItemInsertLog) InvokeHook() {
	if g_Hooks.PlayerItem != nil {
		g_Hooks.PlayerItem.Insert(&PlayerItemRow{row: log.Row})
	}
}

func (log *PlayerItemInsertLog) SetEventId(time.Time) {
}

func (log *PlayerItemInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(87)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.ItemId)
	file.WriteInt16(log.GoRow.Num)
	file.WriteInt8(log.GoRow.IsDressed)
	file.WriteInt8(log.GoRow.BuyBackState)
	file.WriteInt64(log.GoRow.AppendixId)
	file.WriteInt16(log.GoRow.RefineLevelBak)
	file.WriteInt16(log.GoRow.RefineFailTimes)
	file.WriteInt16(log.GoRow.RefineLevel)
	file.WriteInt32(log.GoRow.Price)
}

func (log *PlayerItemInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerItem.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Num)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.IsDressed)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.BuyBackState)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.AppendixId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.RefineLevelBak)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.RefineFailTimes)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.RefineLevel)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Price)))
	return stmt.Execute()
}

func (log *PlayerItemInsertLog) CommitToTLog() {
}

func (log *PlayerItemInsertLog) CommitToXdLog() {
}

func (log *PlayerItemInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerItemDeleteLog struct {
	db *Database
	Old *C.PlayerItem
	GoOld *PlayerItem
}

func (db *Database) newPlayerItemDeleteLog(old *C.PlayerItem, goOld *PlayerItem) *PlayerItemDeleteLog {
	return &PlayerItemDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerItemDeleteLog) Free() {
	C.Free_PlayerItem(log.Old)
}

func (log *PlayerItemDeleteLog) InvokeHook() {
	if g_Hooks.PlayerItem != nil {
		g_Hooks.PlayerItem.Delete(&PlayerItemRow{row: log.Old})
	}
}

func (log *PlayerItemDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerItemDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(87)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt16(log.GoOld.Num)
	file.WriteInt8(log.GoOld.IsDressed)
	file.WriteInt8(log.GoOld.BuyBackState)
	file.WriteInt64(log.GoOld.AppendixId)
	file.WriteInt16(log.GoOld.RefineLevelBak)
	file.WriteInt16(log.GoOld.RefineFailTimes)
	file.WriteInt16(log.GoOld.RefineLevel)
	file.WriteInt32(log.GoOld.Price)
}

func (log *PlayerItemDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerItem.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerItemDeleteLog) CommitToTLog() {
}

func (log *PlayerItemDeleteLog) CommitToXdLog() {
}

func (log *PlayerItemDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerItemUpdateLog struct {
	db *Database 
	Row *C.PlayerItem
	Old *C.PlayerItem
	GoNew *PlayerItem
	GoOld *PlayerItem
}

func (db *Database) newPlayerItemUpdateLog(row, old *C.PlayerItem, goNew, goOld *PlayerItem) *PlayerItemUpdateLog {
	return &PlayerItemUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerItemUpdateLog) Free() {
	C.Free_PlayerItem(log.Old)
}

func (log *PlayerItemUpdateLog) InvokeHook() {
	if g_Hooks.PlayerItem != nil {
		g_Hooks.PlayerItem.Update(&PlayerItemRow{row: log.Row}, &PlayerItemRow{row: log.Old})
	}
}

func (log *PlayerItemUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerItemUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(87)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.ItemId)
	file.WriteInt16(log.GoNew.Num)
	file.WriteInt8(log.GoNew.IsDressed)
	file.WriteInt8(log.GoNew.BuyBackState)
	file.WriteInt64(log.GoNew.AppendixId)
	file.WriteInt16(log.GoNew.RefineLevelBak)
	file.WriteInt16(log.GoNew.RefineFailTimes)
	file.WriteInt16(log.GoNew.RefineLevel)
	file.WriteInt32(log.GoNew.Price)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt16(log.GoOld.Num)
	file.WriteInt8(log.GoOld.IsDressed)
	file.WriteInt8(log.GoOld.BuyBackState)
	file.WriteInt64(log.GoOld.AppendixId)
	file.WriteInt16(log.GoOld.RefineLevelBak)
	file.WriteInt16(log.GoOld.RefineFailTimes)
	file.WriteInt16(log.GoOld.RefineLevel)
	file.WriteInt32(log.GoOld.Price)
}

func (log *PlayerItemUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerItem.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Num)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.IsDressed)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.BuyBackState)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.AppendixId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.RefineLevelBak)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.RefineFailTimes)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.RefineLevel)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Price)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerItemUpdateLog) CommitToTLog() {
}

func (log *PlayerItemUpdateLog) CommitToXdLog() {
}

func (log *PlayerItemUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerItemInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerItem
	for crow := log.db.tables.PlayerItem; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerItem'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerItem = log.db.tables.PlayerItem.next
	}
	C.Free_PlayerItem(log.Row)
}

func (log *PlayerItemDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerItem
	for crow := log.db.tables.PlayerItem; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_item'")
	}
	log.Old.next = log.db.tables.PlayerItem
	log.db.tables.PlayerItem = log.Old
}

func (log *PlayerItemUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerItem
	for crow := log.db.tables.PlayerItem; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_item'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerItem = log.Old
	}
	C.Free_PlayerItem(log.Row)
}

/* ========== player_item_appendix ========== */

type PlayerItemAppendixInsertLog struct {
	db *Database 
	Row *C.PlayerItemAppendix
	GoRow *PlayerItemAppendix
}

func (db *Database) newPlayerItemAppendixInsertLog(row *C.PlayerItemAppendix, goRow *PlayerItemAppendix) *PlayerItemAppendixInsertLog {
	return &PlayerItemAppendixInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerItemAppendixInsertLog) Free() {
}

func (log *PlayerItemAppendixInsertLog) InvokeHook() {
	if g_Hooks.PlayerItemAppendix != nil {
		g_Hooks.PlayerItemAppendix.Insert(&PlayerItemAppendixRow{row: log.Row})
	}
}

func (log *PlayerItemAppendixInsertLog) SetEventId(time.Time) {
}

func (log *PlayerItemAppendixInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(88)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.Health)
	file.WriteInt32(log.GoRow.Cultivation)
	file.WriteInt32(log.GoRow.Speed)
	file.WriteInt32(log.GoRow.Attack)
	file.WriteInt32(log.GoRow.Defence)
	file.WriteInt32(log.GoRow.DodgeLevel)
	file.WriteInt32(log.GoRow.HitLevel)
	file.WriteInt32(log.GoRow.BlockLevel)
	file.WriteInt32(log.GoRow.CriticalLevel)
	file.WriteInt32(log.GoRow.TenacityLevel)
	file.WriteInt32(log.GoRow.DestroyLevel)
	file.WriteInt8(log.GoRow.RecastAttr)
}

func (log *PlayerItemAppendixInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerItemAppendix.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Health)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Cultivation)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Speed)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Attack)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Defence)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.DodgeLevel)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.HitLevel)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.BlockLevel)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.CriticalLevel)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.TenacityLevel)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.DestroyLevel)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RecastAttr)))
	return stmt.Execute()
}

func (log *PlayerItemAppendixInsertLog) CommitToTLog() {
}

func (log *PlayerItemAppendixInsertLog) CommitToXdLog() {
}

func (log *PlayerItemAppendixInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerItemAppendixDeleteLog struct {
	db *Database
	Old *C.PlayerItemAppendix
	GoOld *PlayerItemAppendix
}

func (db *Database) newPlayerItemAppendixDeleteLog(old *C.PlayerItemAppendix, goOld *PlayerItemAppendix) *PlayerItemAppendixDeleteLog {
	return &PlayerItemAppendixDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerItemAppendixDeleteLog) Free() {
	C.Free_PlayerItemAppendix(log.Old)
}

func (log *PlayerItemAppendixDeleteLog) InvokeHook() {
	if g_Hooks.PlayerItemAppendix != nil {
		g_Hooks.PlayerItemAppendix.Delete(&PlayerItemAppendixRow{row: log.Old})
	}
}

func (log *PlayerItemAppendixDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerItemAppendixDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(88)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Health)
	file.WriteInt32(log.GoOld.Cultivation)
	file.WriteInt32(log.GoOld.Speed)
	file.WriteInt32(log.GoOld.Attack)
	file.WriteInt32(log.GoOld.Defence)
	file.WriteInt32(log.GoOld.DodgeLevel)
	file.WriteInt32(log.GoOld.HitLevel)
	file.WriteInt32(log.GoOld.BlockLevel)
	file.WriteInt32(log.GoOld.CriticalLevel)
	file.WriteInt32(log.GoOld.TenacityLevel)
	file.WriteInt32(log.GoOld.DestroyLevel)
	file.WriteInt8(log.GoOld.RecastAttr)
}

func (log *PlayerItemAppendixDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerItemAppendix.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerItemAppendixDeleteLog) CommitToTLog() {
}

func (log *PlayerItemAppendixDeleteLog) CommitToXdLog() {
}

func (log *PlayerItemAppendixDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerItemAppendixUpdateLog struct {
	db *Database 
	Row *C.PlayerItemAppendix
	Old *C.PlayerItemAppendix
	GoNew *PlayerItemAppendix
	GoOld *PlayerItemAppendix
}

func (db *Database) newPlayerItemAppendixUpdateLog(row, old *C.PlayerItemAppendix, goNew, goOld *PlayerItemAppendix) *PlayerItemAppendixUpdateLog {
	return &PlayerItemAppendixUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerItemAppendixUpdateLog) Free() {
	C.Free_PlayerItemAppendix(log.Old)
}

func (log *PlayerItemAppendixUpdateLog) InvokeHook() {
	if g_Hooks.PlayerItemAppendix != nil {
		g_Hooks.PlayerItemAppendix.Update(&PlayerItemAppendixRow{row: log.Row}, &PlayerItemAppendixRow{row: log.Old})
	}
}

func (log *PlayerItemAppendixUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerItemAppendixUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(88)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.Health)
	file.WriteInt32(log.GoNew.Cultivation)
	file.WriteInt32(log.GoNew.Speed)
	file.WriteInt32(log.GoNew.Attack)
	file.WriteInt32(log.GoNew.Defence)
	file.WriteInt32(log.GoNew.DodgeLevel)
	file.WriteInt32(log.GoNew.HitLevel)
	file.WriteInt32(log.GoNew.BlockLevel)
	file.WriteInt32(log.GoNew.CriticalLevel)
	file.WriteInt32(log.GoNew.TenacityLevel)
	file.WriteInt32(log.GoNew.DestroyLevel)
	file.WriteInt8(log.GoNew.RecastAttr)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Health)
	file.WriteInt32(log.GoOld.Cultivation)
	file.WriteInt32(log.GoOld.Speed)
	file.WriteInt32(log.GoOld.Attack)
	file.WriteInt32(log.GoOld.Defence)
	file.WriteInt32(log.GoOld.DodgeLevel)
	file.WriteInt32(log.GoOld.HitLevel)
	file.WriteInt32(log.GoOld.BlockLevel)
	file.WriteInt32(log.GoOld.CriticalLevel)
	file.WriteInt32(log.GoOld.TenacityLevel)
	file.WriteInt32(log.GoOld.DestroyLevel)
	file.WriteInt8(log.GoOld.RecastAttr)
}

func (log *PlayerItemAppendixUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerItemAppendix.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Health)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Cultivation)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Speed)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Attack)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Defence)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.DodgeLevel)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.HitLevel)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.BlockLevel)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.CriticalLevel)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.TenacityLevel)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.DestroyLevel)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RecastAttr)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerItemAppendixUpdateLog) CommitToTLog() {
}

func (log *PlayerItemAppendixUpdateLog) CommitToXdLog() {
}

func (log *PlayerItemAppendixUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerItemAppendixInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerItemAppendix
	for crow := log.db.tables.PlayerItemAppendix; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerItemAppendix'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerItemAppendix = log.db.tables.PlayerItemAppendix.next
	}
	C.Free_PlayerItemAppendix(log.Row)
}

func (log *PlayerItemAppendixDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerItemAppendix
	for crow := log.db.tables.PlayerItemAppendix; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_item_appendix'")
	}
	log.Old.next = log.db.tables.PlayerItemAppendix
	log.db.tables.PlayerItemAppendix = log.Old
}

func (log *PlayerItemAppendixUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerItemAppendix
	for crow := log.db.tables.PlayerItemAppendix; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_item_appendix'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerItemAppendix = log.Old
	}
	C.Free_PlayerItemAppendix(log.Row)
}

/* ========== player_item_buyback ========== */

type PlayerItemBuybackInsertLog struct {
	db *Database 
	Row *C.PlayerItemBuyback
	GoRow *PlayerItemBuyback
}

func (db *Database) newPlayerItemBuybackInsertLog(row *C.PlayerItemBuyback, goRow *PlayerItemBuyback) *PlayerItemBuybackInsertLog {
	return &PlayerItemBuybackInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerItemBuybackInsertLog) Free() {
}

func (log *PlayerItemBuybackInsertLog) InvokeHook() {
	if g_Hooks.PlayerItemBuyback != nil {
		g_Hooks.PlayerItemBuyback.Insert(&PlayerItemBuybackRow{row: log.Row})
	}
}

func (log *PlayerItemBuybackInsertLog) SetEventId(time.Time) {
}

func (log *PlayerItemBuybackInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(89)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.BackId1)
	file.WriteInt64(log.GoRow.BackId2)
	file.WriteInt64(log.GoRow.BackId3)
	file.WriteInt64(log.GoRow.BackId4)
	file.WriteInt64(log.GoRow.BackId5)
	file.WriteInt64(log.GoRow.BackId6)
}

func (log *PlayerItemBuybackInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerItemBuyback.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BackId1)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BackId2)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BackId3)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BackId4)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BackId5)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BackId6)))
	return stmt.Execute()
}

func (log *PlayerItemBuybackInsertLog) CommitToTLog() {
}

func (log *PlayerItemBuybackInsertLog) CommitToXdLog() {
}

func (log *PlayerItemBuybackInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerItemBuybackDeleteLog struct {
	db *Database
	Old *C.PlayerItemBuyback
	GoOld *PlayerItemBuyback
}

func (db *Database) newPlayerItemBuybackDeleteLog(old *C.PlayerItemBuyback, goOld *PlayerItemBuyback) *PlayerItemBuybackDeleteLog {
	return &PlayerItemBuybackDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerItemBuybackDeleteLog) Free() {
	C.Free_PlayerItemBuyback(log.Old)
}

func (log *PlayerItemBuybackDeleteLog) InvokeHook() {
	if g_Hooks.PlayerItemBuyback != nil {
		g_Hooks.PlayerItemBuyback.Delete(&PlayerItemBuybackRow{row: log.Old})
	}
}

func (log *PlayerItemBuybackDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerItemBuybackDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(89)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.BackId1)
	file.WriteInt64(log.GoOld.BackId2)
	file.WriteInt64(log.GoOld.BackId3)
	file.WriteInt64(log.GoOld.BackId4)
	file.WriteInt64(log.GoOld.BackId5)
	file.WriteInt64(log.GoOld.BackId6)
}

func (log *PlayerItemBuybackDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerItemBuyback.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerItemBuybackDeleteLog) CommitToTLog() {
}

func (log *PlayerItemBuybackDeleteLog) CommitToXdLog() {
}

func (log *PlayerItemBuybackDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerItemBuybackUpdateLog struct {
	db *Database 
	Row *C.PlayerItemBuyback
	Old *C.PlayerItemBuyback
	GoNew *PlayerItemBuyback
	GoOld *PlayerItemBuyback
}

func (db *Database) newPlayerItemBuybackUpdateLog(row, old *C.PlayerItemBuyback, goNew, goOld *PlayerItemBuyback) *PlayerItemBuybackUpdateLog {
	return &PlayerItemBuybackUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerItemBuybackUpdateLog) Free() {
	C.Free_PlayerItemBuyback(log.Old)
}

func (log *PlayerItemBuybackUpdateLog) InvokeHook() {
	if g_Hooks.PlayerItemBuyback != nil {
		g_Hooks.PlayerItemBuyback.Update(&PlayerItemBuybackRow{row: log.Row}, &PlayerItemBuybackRow{row: log.Old})
	}
}

func (log *PlayerItemBuybackUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerItemBuybackUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(89)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.BackId1)
	file.WriteInt64(log.GoNew.BackId2)
	file.WriteInt64(log.GoNew.BackId3)
	file.WriteInt64(log.GoNew.BackId4)
	file.WriteInt64(log.GoNew.BackId5)
	file.WriteInt64(log.GoNew.BackId6)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.BackId1)
	file.WriteInt64(log.GoOld.BackId2)
	file.WriteInt64(log.GoOld.BackId3)
	file.WriteInt64(log.GoOld.BackId4)
	file.WriteInt64(log.GoOld.BackId5)
	file.WriteInt64(log.GoOld.BackId6)
}

func (log *PlayerItemBuybackUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerItemBuyback.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BackId1)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BackId2)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BackId3)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BackId4)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BackId5)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BackId6)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerItemBuybackUpdateLog) CommitToTLog() {
}

func (log *PlayerItemBuybackUpdateLog) CommitToXdLog() {
}

func (log *PlayerItemBuybackUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerItemBuybackInsertLog) Rollback() {
	if log.db.tables.PlayerItemBuyback != log.Row {
		panic("insert rollback check failed 'PlayerItemBuyback'")
	}
	log.db.tables.PlayerItemBuyback = nil
	C.Free_PlayerItemBuyback(log.Row)
}

func (log *PlayerItemBuybackDeleteLog) Rollback() {
	if log.db.tables.PlayerItemBuyback != nil {
		panic("delete rollback check failed 'player_item_buyback'")
	}
	log.db.tables.PlayerItemBuyback = log.Old
}

func (log *PlayerItemBuybackUpdateLog) Rollback() {
	if log.db.tables.PlayerItemBuyback != log.Row {
		panic("update rollback check failed 'player_item_buyback'")
	}
	log.db.tables.PlayerItemBuyback = log.Old
	C.Free_PlayerItemBuyback(log.Row)
}

/* ========== player_item_recast_state ========== */

type PlayerItemRecastStateInsertLog struct {
	db *Database 
	Row *C.PlayerItemRecastState
	GoRow *PlayerItemRecastState
}

func (db *Database) newPlayerItemRecastStateInsertLog(row *C.PlayerItemRecastState, goRow *PlayerItemRecastState) *PlayerItemRecastStateInsertLog {
	return &PlayerItemRecastStateInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerItemRecastStateInsertLog) Free() {
}

func (log *PlayerItemRecastStateInsertLog) InvokeHook() {
	if g_Hooks.PlayerItemRecastState != nil {
		g_Hooks.PlayerItemRecastState.Insert(&PlayerItemRecastStateRow{row: log.Row})
	}
}

func (log *PlayerItemRecastStateInsertLog) SetEventId(time.Time) {
}

func (log *PlayerItemRecastStateInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(90)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.PlayerItemId)
	file.WriteInt8(log.GoRow.SelectedAttr)
	file.WriteInt8(log.GoRow.Attr1Type)
	file.WriteInt32(log.GoRow.Attr1Value)
	file.WriteInt8(log.GoRow.Attr2Type)
	file.WriteInt32(log.GoRow.Attr2Value)
	file.WriteInt8(log.GoRow.Attr3Type)
	file.WriteInt32(log.GoRow.Attr3Value)
}

func (log *PlayerItemRecastStateInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerItemRecastState.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.PlayerItemId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.SelectedAttr)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Attr1Type)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Attr1Value)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Attr2Type)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Attr2Value)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Attr3Type)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Attr3Value)))
	return stmt.Execute()
}

func (log *PlayerItemRecastStateInsertLog) CommitToTLog() {
}

func (log *PlayerItemRecastStateInsertLog) CommitToXdLog() {
}

func (log *PlayerItemRecastStateInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerItemRecastStateDeleteLog struct {
	db *Database
	Old *C.PlayerItemRecastState
	GoOld *PlayerItemRecastState
}

func (db *Database) newPlayerItemRecastStateDeleteLog(old *C.PlayerItemRecastState, goOld *PlayerItemRecastState) *PlayerItemRecastStateDeleteLog {
	return &PlayerItemRecastStateDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerItemRecastStateDeleteLog) Free() {
	C.Free_PlayerItemRecastState(log.Old)
}

func (log *PlayerItemRecastStateDeleteLog) InvokeHook() {
	if g_Hooks.PlayerItemRecastState != nil {
		g_Hooks.PlayerItemRecastState.Delete(&PlayerItemRecastStateRow{row: log.Old})
	}
}

func (log *PlayerItemRecastStateDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerItemRecastStateDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(90)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.PlayerItemId)
	file.WriteInt8(log.GoOld.SelectedAttr)
	file.WriteInt8(log.GoOld.Attr1Type)
	file.WriteInt32(log.GoOld.Attr1Value)
	file.WriteInt8(log.GoOld.Attr2Type)
	file.WriteInt32(log.GoOld.Attr2Value)
	file.WriteInt8(log.GoOld.Attr3Type)
	file.WriteInt32(log.GoOld.Attr3Value)
}

func (log *PlayerItemRecastStateDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerItemRecastState.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerItemRecastStateDeleteLog) CommitToTLog() {
}

func (log *PlayerItemRecastStateDeleteLog) CommitToXdLog() {
}

func (log *PlayerItemRecastStateDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerItemRecastStateUpdateLog struct {
	db *Database 
	Row *C.PlayerItemRecastState
	Old *C.PlayerItemRecastState
	GoNew *PlayerItemRecastState
	GoOld *PlayerItemRecastState
}

func (db *Database) newPlayerItemRecastStateUpdateLog(row, old *C.PlayerItemRecastState, goNew, goOld *PlayerItemRecastState) *PlayerItemRecastStateUpdateLog {
	return &PlayerItemRecastStateUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerItemRecastStateUpdateLog) Free() {
	C.Free_PlayerItemRecastState(log.Old)
}

func (log *PlayerItemRecastStateUpdateLog) InvokeHook() {
	if g_Hooks.PlayerItemRecastState != nil {
		g_Hooks.PlayerItemRecastState.Update(&PlayerItemRecastStateRow{row: log.Row}, &PlayerItemRecastStateRow{row: log.Old})
	}
}

func (log *PlayerItemRecastStateUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerItemRecastStateUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(90)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.PlayerItemId)
	file.WriteInt8(log.GoNew.SelectedAttr)
	file.WriteInt8(log.GoNew.Attr1Type)
	file.WriteInt32(log.GoNew.Attr1Value)
	file.WriteInt8(log.GoNew.Attr2Type)
	file.WriteInt32(log.GoNew.Attr2Value)
	file.WriteInt8(log.GoNew.Attr3Type)
	file.WriteInt32(log.GoNew.Attr3Value)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.PlayerItemId)
	file.WriteInt8(log.GoOld.SelectedAttr)
	file.WriteInt8(log.GoOld.Attr1Type)
	file.WriteInt32(log.GoOld.Attr1Value)
	file.WriteInt8(log.GoOld.Attr2Type)
	file.WriteInt32(log.GoOld.Attr2Value)
	file.WriteInt8(log.GoOld.Attr3Type)
	file.WriteInt32(log.GoOld.Attr3Value)
}

func (log *PlayerItemRecastStateUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerItemRecastState.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.PlayerItemId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.SelectedAttr)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Attr1Type)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Attr1Value)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Attr2Type)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Attr2Value)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Attr3Type)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Attr3Value)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerItemRecastStateUpdateLog) CommitToTLog() {
}

func (log *PlayerItemRecastStateUpdateLog) CommitToXdLog() {
}

func (log *PlayerItemRecastStateUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerItemRecastStateInsertLog) Rollback() {
	if log.db.tables.PlayerItemRecastState != log.Row {
		panic("insert rollback check failed 'PlayerItemRecastState'")
	}
	log.db.tables.PlayerItemRecastState = nil
	C.Free_PlayerItemRecastState(log.Row)
}

func (log *PlayerItemRecastStateDeleteLog) Rollback() {
	if log.db.tables.PlayerItemRecastState != nil {
		panic("delete rollback check failed 'player_item_recast_state'")
	}
	log.db.tables.PlayerItemRecastState = log.Old
}

func (log *PlayerItemRecastStateUpdateLog) Rollback() {
	if log.db.tables.PlayerItemRecastState != log.Row {
		panic("update rollback check failed 'player_item_recast_state'")
	}
	log.db.tables.PlayerItemRecastState = log.Old
	C.Free_PlayerItemRecastState(log.Row)
}

/* ========== player_level_award_info ========== */

type PlayerLevelAwardInfoInsertLog struct {
	db *Database 
	Row *C.PlayerLevelAwardInfo
	GoRow *PlayerLevelAwardInfo
}

func (db *Database) newPlayerLevelAwardInfoInsertLog(row *C.PlayerLevelAwardInfo, goRow *PlayerLevelAwardInfo) *PlayerLevelAwardInfoInsertLog {
	return &PlayerLevelAwardInfoInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerLevelAwardInfoInsertLog) Free() {
}

func (log *PlayerLevelAwardInfoInsertLog) InvokeHook() {
	if g_Hooks.PlayerLevelAwardInfo != nil {
		g_Hooks.PlayerLevelAwardInfo.Insert(&PlayerLevelAwardInfoRow{row: log.Row})
	}
}

func (log *PlayerLevelAwardInfoInsertLog) SetEventId(time.Time) {
}

func (log *PlayerLevelAwardInfoInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(91)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteString(log.GoRow.Awarded)
}

func (log *PlayerLevelAwardInfoInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerLevelAwardInfo.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Awarded), int(log.Row.Awarded_len))
	return stmt.Execute()
}

func (log *PlayerLevelAwardInfoInsertLog) CommitToTLog() {
}

func (log *PlayerLevelAwardInfoInsertLog) CommitToXdLog() {
}

func (log *PlayerLevelAwardInfoInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerLevelAwardInfoDeleteLog struct {
	db *Database
	Old *C.PlayerLevelAwardInfo
	GoOld *PlayerLevelAwardInfo
}

func (db *Database) newPlayerLevelAwardInfoDeleteLog(old *C.PlayerLevelAwardInfo, goOld *PlayerLevelAwardInfo) *PlayerLevelAwardInfoDeleteLog {
	return &PlayerLevelAwardInfoDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerLevelAwardInfoDeleteLog) Free() {
	C.Free_PlayerLevelAwardInfo(log.Old)
}

func (log *PlayerLevelAwardInfoDeleteLog) InvokeHook() {
	if g_Hooks.PlayerLevelAwardInfo != nil {
		g_Hooks.PlayerLevelAwardInfo.Delete(&PlayerLevelAwardInfoRow{row: log.Old})
	}
}

func (log *PlayerLevelAwardInfoDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerLevelAwardInfoDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(91)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteString(log.GoOld.Awarded)
}

func (log *PlayerLevelAwardInfoDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerLevelAwardInfo.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerLevelAwardInfoDeleteLog) CommitToTLog() {
}

func (log *PlayerLevelAwardInfoDeleteLog) CommitToXdLog() {
}

func (log *PlayerLevelAwardInfoDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerLevelAwardInfoUpdateLog struct {
	db *Database 
	Row *C.PlayerLevelAwardInfo
	Old *C.PlayerLevelAwardInfo
	GoNew *PlayerLevelAwardInfo
	GoOld *PlayerLevelAwardInfo
}

func (db *Database) newPlayerLevelAwardInfoUpdateLog(row, old *C.PlayerLevelAwardInfo, goNew, goOld *PlayerLevelAwardInfo) *PlayerLevelAwardInfoUpdateLog {
	return &PlayerLevelAwardInfoUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerLevelAwardInfoUpdateLog) Free() {
	C.Free_PlayerLevelAwardInfo(log.Old)
}

func (log *PlayerLevelAwardInfoUpdateLog) InvokeHook() {
	if g_Hooks.PlayerLevelAwardInfo != nil {
		g_Hooks.PlayerLevelAwardInfo.Update(&PlayerLevelAwardInfoRow{row: log.Row}, &PlayerLevelAwardInfoRow{row: log.Old})
	}
}

func (log *PlayerLevelAwardInfoUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerLevelAwardInfoUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(91)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteString(log.GoNew.Awarded)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteString(log.GoOld.Awarded)
}

func (log *PlayerLevelAwardInfoUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerLevelAwardInfo.Update
	stmt.BindVarchar(unsafe.Pointer(log.Row.Awarded), int(log.Row.Awarded_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerLevelAwardInfoUpdateLog) CommitToTLog() {
}

func (log *PlayerLevelAwardInfoUpdateLog) CommitToXdLog() {
}

func (log *PlayerLevelAwardInfoUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerLevelAwardInfoInsertLog) Rollback() {
	if log.db.tables.PlayerLevelAwardInfo != log.Row {
		panic("insert rollback check failed 'PlayerLevelAwardInfo'")
	}
	log.db.tables.PlayerLevelAwardInfo = nil
	C.Free_PlayerLevelAwardInfo(log.Row)
}

func (log *PlayerLevelAwardInfoDeleteLog) Rollback() {
	if log.db.tables.PlayerLevelAwardInfo != nil {
		panic("delete rollback check failed 'player_level_award_info'")
	}
	log.db.tables.PlayerLevelAwardInfo = log.Old
}

func (log *PlayerLevelAwardInfoUpdateLog) Rollback() {
	if log.db.tables.PlayerLevelAwardInfo != log.Row {
		panic("update rollback check failed 'player_level_award_info'")
	}
	log.db.tables.PlayerLevelAwardInfo = log.Old
	C.Free_PlayerLevelAwardInfo(log.Row)
}

/* ========== player_login_award_record ========== */

type PlayerLoginAwardRecordInsertLog struct {
	db *Database 
	Row *C.PlayerLoginAwardRecord
	GoRow *PlayerLoginAwardRecord
}

func (db *Database) newPlayerLoginAwardRecordInsertLog(row *C.PlayerLoginAwardRecord, goRow *PlayerLoginAwardRecord) *PlayerLoginAwardRecordInsertLog {
	return &PlayerLoginAwardRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerLoginAwardRecordInsertLog) Free() {
}

func (log *PlayerLoginAwardRecordInsertLog) InvokeHook() {
	if g_Hooks.PlayerLoginAwardRecord != nil {
		g_Hooks.PlayerLoginAwardRecord.Insert(&PlayerLoginAwardRecordRow{row: log.Row})
	}
}

func (log *PlayerLoginAwardRecordInsertLog) SetEventId(time.Time) {
}

func (log *PlayerLoginAwardRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(92)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.ActiveDays)
	file.WriteInt32(log.GoRow.Record)
	file.WriteInt64(log.GoRow.UpdateTimestamp)
}

func (log *PlayerLoginAwardRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerLoginAwardRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.ActiveDays)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Record)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.UpdateTimestamp)))
	return stmt.Execute()
}

func (log *PlayerLoginAwardRecordInsertLog) CommitToTLog() {
}

func (log *PlayerLoginAwardRecordInsertLog) CommitToXdLog() {
}

func (log *PlayerLoginAwardRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerLoginAwardRecordDeleteLog struct {
	db *Database
	Old *C.PlayerLoginAwardRecord
	GoOld *PlayerLoginAwardRecord
}

func (db *Database) newPlayerLoginAwardRecordDeleteLog(old *C.PlayerLoginAwardRecord, goOld *PlayerLoginAwardRecord) *PlayerLoginAwardRecordDeleteLog {
	return &PlayerLoginAwardRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerLoginAwardRecordDeleteLog) Free() {
	C.Free_PlayerLoginAwardRecord(log.Old)
}

func (log *PlayerLoginAwardRecordDeleteLog) InvokeHook() {
	if g_Hooks.PlayerLoginAwardRecord != nil {
		g_Hooks.PlayerLoginAwardRecord.Delete(&PlayerLoginAwardRecordRow{row: log.Old})
	}
}

func (log *PlayerLoginAwardRecordDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerLoginAwardRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(92)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.ActiveDays)
	file.WriteInt32(log.GoOld.Record)
	file.WriteInt64(log.GoOld.UpdateTimestamp)
}

func (log *PlayerLoginAwardRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerLoginAwardRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerLoginAwardRecordDeleteLog) CommitToTLog() {
}

func (log *PlayerLoginAwardRecordDeleteLog) CommitToXdLog() {
}

func (log *PlayerLoginAwardRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerLoginAwardRecordUpdateLog struct {
	db *Database 
	Row *C.PlayerLoginAwardRecord
	Old *C.PlayerLoginAwardRecord
	GoNew *PlayerLoginAwardRecord
	GoOld *PlayerLoginAwardRecord
}

func (db *Database) newPlayerLoginAwardRecordUpdateLog(row, old *C.PlayerLoginAwardRecord, goNew, goOld *PlayerLoginAwardRecord) *PlayerLoginAwardRecordUpdateLog {
	return &PlayerLoginAwardRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerLoginAwardRecordUpdateLog) Free() {
	C.Free_PlayerLoginAwardRecord(log.Old)
}

func (log *PlayerLoginAwardRecordUpdateLog) InvokeHook() {
	if g_Hooks.PlayerLoginAwardRecord != nil {
		g_Hooks.PlayerLoginAwardRecord.Update(&PlayerLoginAwardRecordRow{row: log.Row}, &PlayerLoginAwardRecordRow{row: log.Old})
	}
}

func (log *PlayerLoginAwardRecordUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerLoginAwardRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(92)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.ActiveDays)
	file.WriteInt32(log.GoNew.Record)
	file.WriteInt64(log.GoNew.UpdateTimestamp)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.ActiveDays)
	file.WriteInt32(log.GoOld.Record)
	file.WriteInt64(log.GoOld.UpdateTimestamp)
}

func (log *PlayerLoginAwardRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerLoginAwardRecord.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.ActiveDays)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Record)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.UpdateTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerLoginAwardRecordUpdateLog) CommitToTLog() {
}

func (log *PlayerLoginAwardRecordUpdateLog) CommitToXdLog() {
}

func (log *PlayerLoginAwardRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerLoginAwardRecordInsertLog) Rollback() {
	if log.db.tables.PlayerLoginAwardRecord != log.Row {
		panic("insert rollback check failed 'PlayerLoginAwardRecord'")
	}
	log.db.tables.PlayerLoginAwardRecord = nil
	C.Free_PlayerLoginAwardRecord(log.Row)
}

func (log *PlayerLoginAwardRecordDeleteLog) Rollback() {
	if log.db.tables.PlayerLoginAwardRecord != nil {
		panic("delete rollback check failed 'player_login_award_record'")
	}
	log.db.tables.PlayerLoginAwardRecord = log.Old
}

func (log *PlayerLoginAwardRecordUpdateLog) Rollback() {
	if log.db.tables.PlayerLoginAwardRecord != log.Row {
		panic("update rollback check failed 'player_login_award_record'")
	}
	log.db.tables.PlayerLoginAwardRecord = log.Old
	C.Free_PlayerLoginAwardRecord(log.Row)
}

/* ========== player_mail ========== */

type PlayerMailInsertLog struct {
	db *Database 
	Row *C.PlayerMail
	GoRow *PlayerMail
}

func (db *Database) newPlayerMailInsertLog(row *C.PlayerMail, goRow *PlayerMail) *PlayerMailInsertLog {
	return &PlayerMailInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerMailInsertLog) Free() {
}

func (log *PlayerMailInsertLog) InvokeHook() {
	if g_Hooks.PlayerMail != nil {
		g_Hooks.PlayerMail.Insert(&PlayerMailRow{row: log.Row})
	}
}

func (log *PlayerMailInsertLog) SetEventId(time.Time) {
}

func (log *PlayerMailInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(93)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.MailId)
	file.WriteInt8(log.GoRow.State)
	file.WriteInt64(log.GoRow.SendTime)
	file.WriteString(log.GoRow.Parameters)
	file.WriteInt8(log.GoRow.HaveAttachment)
	file.WriteString(log.GoRow.Title)
	file.WriteString(log.GoRow.Content)
	file.WriteInt64(log.GoRow.ExpireTime)
	file.WriteInt8(log.GoRow.Priority)
}

func (log *PlayerMailInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMail.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.MailId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.State)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SendTime)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Parameters), int(log.Row.Parameters_len))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.HaveAttachment)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Title), int(log.Row.Title_len))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Content), int(log.Row.Content_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ExpireTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Priority)))
	return stmt.Execute()
}

func (log *PlayerMailInsertLog) CommitToTLog() {
}

func (log *PlayerMailInsertLog) CommitToXdLog() {
}

func (log *PlayerMailInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMailDeleteLog struct {
	db *Database
	Old *C.PlayerMail
	GoOld *PlayerMail
}

func (db *Database) newPlayerMailDeleteLog(old *C.PlayerMail, goOld *PlayerMail) *PlayerMailDeleteLog {
	return &PlayerMailDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerMailDeleteLog) Free() {
	C.Free_PlayerMail(log.Old)
}

func (log *PlayerMailDeleteLog) InvokeHook() {
	if g_Hooks.PlayerMail != nil {
		g_Hooks.PlayerMail.Delete(&PlayerMailRow{row: log.Old})
	}
}

func (log *PlayerMailDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerMailDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(93)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.MailId)
	file.WriteInt8(log.GoOld.State)
	file.WriteInt64(log.GoOld.SendTime)
	file.WriteString(log.GoOld.Parameters)
	file.WriteInt8(log.GoOld.HaveAttachment)
	file.WriteString(log.GoOld.Title)
	file.WriteString(log.GoOld.Content)
	file.WriteInt64(log.GoOld.ExpireTime)
	file.WriteInt8(log.GoOld.Priority)
}

func (log *PlayerMailDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMail.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerMailDeleteLog) CommitToTLog() {
}

func (log *PlayerMailDeleteLog) CommitToXdLog() {
}

func (log *PlayerMailDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMailUpdateLog struct {
	db *Database 
	Row *C.PlayerMail
	Old *C.PlayerMail
	GoNew *PlayerMail
	GoOld *PlayerMail
}

func (db *Database) newPlayerMailUpdateLog(row, old *C.PlayerMail, goNew, goOld *PlayerMail) *PlayerMailUpdateLog {
	return &PlayerMailUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerMailUpdateLog) Free() {
	C.Free_PlayerMail(log.Old)
}

func (log *PlayerMailUpdateLog) InvokeHook() {
	if g_Hooks.PlayerMail != nil {
		g_Hooks.PlayerMail.Update(&PlayerMailRow{row: log.Row}, &PlayerMailRow{row: log.Old})
	}
}

func (log *PlayerMailUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerMailUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(93)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.MailId)
	file.WriteInt8(log.GoNew.State)
	file.WriteInt64(log.GoNew.SendTime)
	file.WriteString(log.GoNew.Parameters)
	file.WriteInt8(log.GoNew.HaveAttachment)
	file.WriteString(log.GoNew.Title)
	file.WriteString(log.GoNew.Content)
	file.WriteInt64(log.GoNew.ExpireTime)
	file.WriteInt8(log.GoNew.Priority)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.MailId)
	file.WriteInt8(log.GoOld.State)
	file.WriteInt64(log.GoOld.SendTime)
	file.WriteString(log.GoOld.Parameters)
	file.WriteInt8(log.GoOld.HaveAttachment)
	file.WriteString(log.GoOld.Title)
	file.WriteString(log.GoOld.Content)
	file.WriteInt64(log.GoOld.ExpireTime)
	file.WriteInt8(log.GoOld.Priority)
}

func (log *PlayerMailUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMail.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.MailId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.State)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SendTime)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Parameters), int(log.Row.Parameters_len))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.HaveAttachment)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Title), int(log.Row.Title_len))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Content), int(log.Row.Content_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ExpireTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Priority)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerMailUpdateLog) CommitToTLog() {
}

func (log *PlayerMailUpdateLog) CommitToXdLog() {
}

func (log *PlayerMailUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerMailInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerMail
	for crow := log.db.tables.PlayerMail; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerMail'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerMail = log.db.tables.PlayerMail.next
	}
	C.Free_PlayerMail(log.Row)
}

func (log *PlayerMailDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerMail
	for crow := log.db.tables.PlayerMail; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_mail'")
	}
	log.Old.next = log.db.tables.PlayerMail
	log.db.tables.PlayerMail = log.Old
}

func (log *PlayerMailUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerMail
	for crow := log.db.tables.PlayerMail; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_mail'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerMail = log.Old
	}
	C.Free_PlayerMail(log.Row)
}

/* ========== player_mail_attachment ========== */

type PlayerMailAttachmentInsertLog struct {
	db *Database 
	Row *C.PlayerMailAttachment
	GoRow *PlayerMailAttachment
}

func (db *Database) newPlayerMailAttachmentInsertLog(row *C.PlayerMailAttachment, goRow *PlayerMailAttachment) *PlayerMailAttachmentInsertLog {
	return &PlayerMailAttachmentInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerMailAttachmentInsertLog) Free() {
}

func (log *PlayerMailAttachmentInsertLog) InvokeHook() {
	if g_Hooks.PlayerMailAttachment != nil {
		g_Hooks.PlayerMailAttachment.Insert(&PlayerMailAttachmentRow{row: log.Row})
	}
}

func (log *PlayerMailAttachmentInsertLog) SetEventId(time.Time) {
}

func (log *PlayerMailAttachmentInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(94)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.PlayerMailId)
	file.WriteInt8(log.GoRow.AttachmentType)
	file.WriteInt16(log.GoRow.ItemId)
	file.WriteInt64(log.GoRow.ItemNum)
}

func (log *PlayerMailAttachmentInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMailAttachment.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.PlayerMailId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.AttachmentType)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ItemNum)))
	return stmt.Execute()
}

func (log *PlayerMailAttachmentInsertLog) CommitToTLog() {
}

func (log *PlayerMailAttachmentInsertLog) CommitToXdLog() {
}

func (log *PlayerMailAttachmentInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMailAttachmentDeleteLog struct {
	db *Database
	Old *C.PlayerMailAttachment
	GoOld *PlayerMailAttachment
}

func (db *Database) newPlayerMailAttachmentDeleteLog(old *C.PlayerMailAttachment, goOld *PlayerMailAttachment) *PlayerMailAttachmentDeleteLog {
	return &PlayerMailAttachmentDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerMailAttachmentDeleteLog) Free() {
	C.Free_PlayerMailAttachment(log.Old)
}

func (log *PlayerMailAttachmentDeleteLog) InvokeHook() {
	if g_Hooks.PlayerMailAttachment != nil {
		g_Hooks.PlayerMailAttachment.Delete(&PlayerMailAttachmentRow{row: log.Old})
	}
}

func (log *PlayerMailAttachmentDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerMailAttachmentDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(94)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.PlayerMailId)
	file.WriteInt8(log.GoOld.AttachmentType)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt64(log.GoOld.ItemNum)
}

func (log *PlayerMailAttachmentDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMailAttachment.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerMailAttachmentDeleteLog) CommitToTLog() {
}

func (log *PlayerMailAttachmentDeleteLog) CommitToXdLog() {
}

func (log *PlayerMailAttachmentDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMailAttachmentUpdateLog struct {
	db *Database 
	Row *C.PlayerMailAttachment
	Old *C.PlayerMailAttachment
	GoNew *PlayerMailAttachment
	GoOld *PlayerMailAttachment
}

func (db *Database) newPlayerMailAttachmentUpdateLog(row, old *C.PlayerMailAttachment, goNew, goOld *PlayerMailAttachment) *PlayerMailAttachmentUpdateLog {
	return &PlayerMailAttachmentUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerMailAttachmentUpdateLog) Free() {
	C.Free_PlayerMailAttachment(log.Old)
}

func (log *PlayerMailAttachmentUpdateLog) InvokeHook() {
	if g_Hooks.PlayerMailAttachment != nil {
		g_Hooks.PlayerMailAttachment.Update(&PlayerMailAttachmentRow{row: log.Row}, &PlayerMailAttachmentRow{row: log.Old})
	}
}

func (log *PlayerMailAttachmentUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerMailAttachmentUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(94)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.PlayerMailId)
	file.WriteInt8(log.GoNew.AttachmentType)
	file.WriteInt16(log.GoNew.ItemId)
	file.WriteInt64(log.GoNew.ItemNum)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.PlayerMailId)
	file.WriteInt8(log.GoOld.AttachmentType)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt64(log.GoOld.ItemNum)
}

func (log *PlayerMailAttachmentUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMailAttachment.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.PlayerMailId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.AttachmentType)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ItemNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerMailAttachmentUpdateLog) CommitToTLog() {
}

func (log *PlayerMailAttachmentUpdateLog) CommitToXdLog() {
}

func (log *PlayerMailAttachmentUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerMailAttachmentInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerMailAttachment
	for crow := log.db.tables.PlayerMailAttachment; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerMailAttachment'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerMailAttachment = log.db.tables.PlayerMailAttachment.next
	}
	C.Free_PlayerMailAttachment(log.Row)
}

func (log *PlayerMailAttachmentDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerMailAttachment
	for crow := log.db.tables.PlayerMailAttachment; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_mail_attachment'")
	}
	log.Old.next = log.db.tables.PlayerMailAttachment
	log.db.tables.PlayerMailAttachment = log.Old
}

func (log *PlayerMailAttachmentUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerMailAttachment
	for crow := log.db.tables.PlayerMailAttachment; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_mail_attachment'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerMailAttachment = log.Old
	}
	C.Free_PlayerMailAttachment(log.Row)
}

/* ========== player_mail_attachment_lg ========== */

type PlayerMailAttachmentLgInsertLog struct {
	db *Database 
	Row *C.PlayerMailAttachmentLg
	GoRow *PlayerMailAttachmentLg
}

func (db *Database) newPlayerMailAttachmentLgInsertLog(row *C.PlayerMailAttachmentLg, goRow *PlayerMailAttachmentLg) *PlayerMailAttachmentLgInsertLog {
	return &PlayerMailAttachmentLgInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerMailAttachmentLgInsertLog) Free() {
}

func (log *PlayerMailAttachmentLgInsertLog) InvokeHook() {
	if g_Hooks.PlayerMailAttachmentLg != nil {
		g_Hooks.PlayerMailAttachmentLg.Insert(&PlayerMailAttachmentLgRow{row: log.Row})
	}
}

func (log *PlayerMailAttachmentLgInsertLog) SetEventId(time.Time) {
}

func (log *PlayerMailAttachmentLgInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(95)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.PlayerMailId)
	file.WriteInt8(log.GoRow.AttachmentType)
	file.WriteInt16(log.GoRow.ItemId)
	file.WriteInt64(log.GoRow.ItemNum)
	file.WriteInt64(log.GoRow.TakeTimestamp)
}

func (log *PlayerMailAttachmentLgInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMailAttachmentLg.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.PlayerMailId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.AttachmentType)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ItemNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.TakeTimestamp)))
	return stmt.Execute()
}

func (log *PlayerMailAttachmentLgInsertLog) CommitToTLog() {
}

func (log *PlayerMailAttachmentLgInsertLog) CommitToXdLog() {
}

func (log *PlayerMailAttachmentLgInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMailAttachmentLgDeleteLog struct {
	db *Database
	Old *C.PlayerMailAttachmentLg
	GoOld *PlayerMailAttachmentLg
}

func (db *Database) newPlayerMailAttachmentLgDeleteLog(old *C.PlayerMailAttachmentLg, goOld *PlayerMailAttachmentLg) *PlayerMailAttachmentLgDeleteLog {
	return &PlayerMailAttachmentLgDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerMailAttachmentLgDeleteLog) Free() {
	C.Free_PlayerMailAttachmentLg(log.Old)
}

func (log *PlayerMailAttachmentLgDeleteLog) InvokeHook() {
	if g_Hooks.PlayerMailAttachmentLg != nil {
		g_Hooks.PlayerMailAttachmentLg.Delete(&PlayerMailAttachmentLgRow{row: log.Old})
	}
}

func (log *PlayerMailAttachmentLgDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerMailAttachmentLgDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(95)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.PlayerMailId)
	file.WriteInt8(log.GoOld.AttachmentType)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt64(log.GoOld.ItemNum)
	file.WriteInt64(log.GoOld.TakeTimestamp)
}

func (log *PlayerMailAttachmentLgDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMailAttachmentLg.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerMailAttachmentLgDeleteLog) CommitToTLog() {
}

func (log *PlayerMailAttachmentLgDeleteLog) CommitToXdLog() {
}

func (log *PlayerMailAttachmentLgDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMailAttachmentLgUpdateLog struct {
	db *Database 
	Row *C.PlayerMailAttachmentLg
	Old *C.PlayerMailAttachmentLg
	GoNew *PlayerMailAttachmentLg
	GoOld *PlayerMailAttachmentLg
}

func (db *Database) newPlayerMailAttachmentLgUpdateLog(row, old *C.PlayerMailAttachmentLg, goNew, goOld *PlayerMailAttachmentLg) *PlayerMailAttachmentLgUpdateLog {
	return &PlayerMailAttachmentLgUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerMailAttachmentLgUpdateLog) Free() {
	C.Free_PlayerMailAttachmentLg(log.Old)
}

func (log *PlayerMailAttachmentLgUpdateLog) InvokeHook() {
	if g_Hooks.PlayerMailAttachmentLg != nil {
		g_Hooks.PlayerMailAttachmentLg.Update(&PlayerMailAttachmentLgRow{row: log.Row}, &PlayerMailAttachmentLgRow{row: log.Old})
	}
}

func (log *PlayerMailAttachmentLgUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerMailAttachmentLgUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(95)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.PlayerMailId)
	file.WriteInt8(log.GoNew.AttachmentType)
	file.WriteInt16(log.GoNew.ItemId)
	file.WriteInt64(log.GoNew.ItemNum)
	file.WriteInt64(log.GoNew.TakeTimestamp)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.PlayerMailId)
	file.WriteInt8(log.GoOld.AttachmentType)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt64(log.GoOld.ItemNum)
	file.WriteInt64(log.GoOld.TakeTimestamp)
}

func (log *PlayerMailAttachmentLgUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMailAttachmentLg.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.PlayerMailId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.AttachmentType)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ItemNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.TakeTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerMailAttachmentLgUpdateLog) CommitToTLog() {
}

func (log *PlayerMailAttachmentLgUpdateLog) CommitToXdLog() {
}

func (log *PlayerMailAttachmentLgUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerMailAttachmentLgInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerMailAttachmentLg
	for crow := log.db.tables.PlayerMailAttachmentLg; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerMailAttachmentLg'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerMailAttachmentLg = log.db.tables.PlayerMailAttachmentLg.next
	}
	C.Free_PlayerMailAttachmentLg(log.Row)
}

func (log *PlayerMailAttachmentLgDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerMailAttachmentLg
	for crow := log.db.tables.PlayerMailAttachmentLg; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_mail_attachment_lg'")
	}
	log.Old.next = log.db.tables.PlayerMailAttachmentLg
	log.db.tables.PlayerMailAttachmentLg = log.Old
}

func (log *PlayerMailAttachmentLgUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerMailAttachmentLg
	for crow := log.db.tables.PlayerMailAttachmentLg; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_mail_attachment_lg'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerMailAttachmentLg = log.Old
	}
	C.Free_PlayerMailAttachmentLg(log.Row)
}

/* ========== player_mail_lg ========== */

type PlayerMailLgInsertLog struct {
	db *Database 
	Row *C.PlayerMailLg
	GoRow *PlayerMailLg
}

func (db *Database) newPlayerMailLgInsertLog(row *C.PlayerMailLg, goRow *PlayerMailLg) *PlayerMailLgInsertLog {
	return &PlayerMailLgInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerMailLgInsertLog) Free() {
}

func (log *PlayerMailLgInsertLog) InvokeHook() {
	if g_Hooks.PlayerMailLg != nil {
		g_Hooks.PlayerMailLg.Insert(&PlayerMailLgRow{row: log.Row})
	}
}

func (log *PlayerMailLgInsertLog) SetEventId(time.Time) {
}

func (log *PlayerMailLgInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(96)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pmid)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.MailId)
	file.WriteInt8(log.GoRow.State)
	file.WriteInt64(log.GoRow.SendTime)
	file.WriteString(log.GoRow.Parameters)
	file.WriteInt8(log.GoRow.HaveAttachment)
	file.WriteString(log.GoRow.Title)
	file.WriteString(log.GoRow.Content)
}

func (log *PlayerMailLgInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMailLg.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pmid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.MailId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.State)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SendTime)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Parameters), int(log.Row.Parameters_len))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.HaveAttachment)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Title), int(log.Row.Title_len))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Content), int(log.Row.Content_len))
	return stmt.Execute()
}

func (log *PlayerMailLgInsertLog) CommitToTLog() {
}

func (log *PlayerMailLgInsertLog) CommitToXdLog() {
}

func (log *PlayerMailLgInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMailLgDeleteLog struct {
	db *Database
	Old *C.PlayerMailLg
	GoOld *PlayerMailLg
}

func (db *Database) newPlayerMailLgDeleteLog(old *C.PlayerMailLg, goOld *PlayerMailLg) *PlayerMailLgDeleteLog {
	return &PlayerMailLgDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerMailLgDeleteLog) Free() {
	C.Free_PlayerMailLg(log.Old)
}

func (log *PlayerMailLgDeleteLog) InvokeHook() {
	if g_Hooks.PlayerMailLg != nil {
		g_Hooks.PlayerMailLg.Delete(&PlayerMailLgRow{row: log.Old})
	}
}

func (log *PlayerMailLgDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerMailLgDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(96)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pmid)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.MailId)
	file.WriteInt8(log.GoOld.State)
	file.WriteInt64(log.GoOld.SendTime)
	file.WriteString(log.GoOld.Parameters)
	file.WriteInt8(log.GoOld.HaveAttachment)
	file.WriteString(log.GoOld.Title)
	file.WriteString(log.GoOld.Content)
}

func (log *PlayerMailLgDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMailLg.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerMailLgDeleteLog) CommitToTLog() {
}

func (log *PlayerMailLgDeleteLog) CommitToXdLog() {
}

func (log *PlayerMailLgDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMailLgUpdateLog struct {
	db *Database 
	Row *C.PlayerMailLg
	Old *C.PlayerMailLg
	GoNew *PlayerMailLg
	GoOld *PlayerMailLg
}

func (db *Database) newPlayerMailLgUpdateLog(row, old *C.PlayerMailLg, goNew, goOld *PlayerMailLg) *PlayerMailLgUpdateLog {
	return &PlayerMailLgUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerMailLgUpdateLog) Free() {
	C.Free_PlayerMailLg(log.Old)
}

func (log *PlayerMailLgUpdateLog) InvokeHook() {
	if g_Hooks.PlayerMailLg != nil {
		g_Hooks.PlayerMailLg.Update(&PlayerMailLgRow{row: log.Row}, &PlayerMailLgRow{row: log.Old})
	}
}

func (log *PlayerMailLgUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerMailLgUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(96)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pmid)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.MailId)
	file.WriteInt8(log.GoNew.State)
	file.WriteInt64(log.GoNew.SendTime)
	file.WriteString(log.GoNew.Parameters)
	file.WriteInt8(log.GoNew.HaveAttachment)
	file.WriteString(log.GoNew.Title)
	file.WriteString(log.GoNew.Content)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pmid)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.MailId)
	file.WriteInt8(log.GoOld.State)
	file.WriteInt64(log.GoOld.SendTime)
	file.WriteString(log.GoOld.Parameters)
	file.WriteInt8(log.GoOld.HaveAttachment)
	file.WriteString(log.GoOld.Title)
	file.WriteString(log.GoOld.Content)
}

func (log *PlayerMailLgUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMailLg.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pmid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.MailId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.State)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SendTime)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Parameters), int(log.Row.Parameters_len))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.HaveAttachment)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Title), int(log.Row.Title_len))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Content), int(log.Row.Content_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerMailLgUpdateLog) CommitToTLog() {
}

func (log *PlayerMailLgUpdateLog) CommitToXdLog() {
}

func (log *PlayerMailLgUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerMailLgInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerMailLg
	for crow := log.db.tables.PlayerMailLg; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerMailLg'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerMailLg = log.db.tables.PlayerMailLg.next
	}
	C.Free_PlayerMailLg(log.Row)
}

func (log *PlayerMailLgDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerMailLg
	for crow := log.db.tables.PlayerMailLg; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_mail_lg'")
	}
	log.Old.next = log.db.tables.PlayerMailLg
	log.db.tables.PlayerMailLg = log.Old
}

func (log *PlayerMailLgUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerMailLg
	for crow := log.db.tables.PlayerMailLg; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_mail_lg'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerMailLg = log.Old
	}
	C.Free_PlayerMailLg(log.Row)
}

/* ========== player_meditation_state ========== */

type PlayerMeditationStateInsertLog struct {
	db *Database 
	Row *C.PlayerMeditationState
	GoRow *PlayerMeditationState
}

func (db *Database) newPlayerMeditationStateInsertLog(row *C.PlayerMeditationState, goRow *PlayerMeditationState) *PlayerMeditationStateInsertLog {
	return &PlayerMeditationStateInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerMeditationStateInsertLog) Free() {
}

func (log *PlayerMeditationStateInsertLog) InvokeHook() {
	if g_Hooks.PlayerMeditationState != nil {
		g_Hooks.PlayerMeditationState.Insert(&PlayerMeditationStateRow{row: log.Row})
	}
}

func (log *PlayerMeditationStateInsertLog) SetEventId(time.Time) {
}

func (log *PlayerMeditationStateInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(97)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.AccumulateTime)
	file.WriteInt64(log.GoRow.StartTimestamp)
}

func (log *PlayerMeditationStateInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMeditationState.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.AccumulateTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.StartTimestamp)))
	return stmt.Execute()
}

func (log *PlayerMeditationStateInsertLog) CommitToTLog() {
}

func (log *PlayerMeditationStateInsertLog) CommitToXdLog() {
}

func (log *PlayerMeditationStateInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMeditationStateDeleteLog struct {
	db *Database
	Old *C.PlayerMeditationState
	GoOld *PlayerMeditationState
}

func (db *Database) newPlayerMeditationStateDeleteLog(old *C.PlayerMeditationState, goOld *PlayerMeditationState) *PlayerMeditationStateDeleteLog {
	return &PlayerMeditationStateDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerMeditationStateDeleteLog) Free() {
	C.Free_PlayerMeditationState(log.Old)
}

func (log *PlayerMeditationStateDeleteLog) InvokeHook() {
	if g_Hooks.PlayerMeditationState != nil {
		g_Hooks.PlayerMeditationState.Delete(&PlayerMeditationStateRow{row: log.Old})
	}
}

func (log *PlayerMeditationStateDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerMeditationStateDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(97)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.AccumulateTime)
	file.WriteInt64(log.GoOld.StartTimestamp)
}

func (log *PlayerMeditationStateDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMeditationState.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerMeditationStateDeleteLog) CommitToTLog() {
}

func (log *PlayerMeditationStateDeleteLog) CommitToXdLog() {
}

func (log *PlayerMeditationStateDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMeditationStateUpdateLog struct {
	db *Database 
	Row *C.PlayerMeditationState
	Old *C.PlayerMeditationState
	GoNew *PlayerMeditationState
	GoOld *PlayerMeditationState
}

func (db *Database) newPlayerMeditationStateUpdateLog(row, old *C.PlayerMeditationState, goNew, goOld *PlayerMeditationState) *PlayerMeditationStateUpdateLog {
	return &PlayerMeditationStateUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerMeditationStateUpdateLog) Free() {
	C.Free_PlayerMeditationState(log.Old)
}

func (log *PlayerMeditationStateUpdateLog) InvokeHook() {
	if g_Hooks.PlayerMeditationState != nil {
		g_Hooks.PlayerMeditationState.Update(&PlayerMeditationStateRow{row: log.Row}, &PlayerMeditationStateRow{row: log.Old})
	}
}

func (log *PlayerMeditationStateUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerMeditationStateUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(97)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.AccumulateTime)
	file.WriteInt64(log.GoNew.StartTimestamp)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.AccumulateTime)
	file.WriteInt64(log.GoOld.StartTimestamp)
}

func (log *PlayerMeditationStateUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMeditationState.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.AccumulateTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.StartTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerMeditationStateUpdateLog) CommitToTLog() {
}

func (log *PlayerMeditationStateUpdateLog) CommitToXdLog() {
}

func (log *PlayerMeditationStateUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerMeditationStateInsertLog) Rollback() {
	if log.db.tables.PlayerMeditationState != log.Row {
		panic("insert rollback check failed 'PlayerMeditationState'")
	}
	log.db.tables.PlayerMeditationState = nil
	C.Free_PlayerMeditationState(log.Row)
}

func (log *PlayerMeditationStateDeleteLog) Rollback() {
	if log.db.tables.PlayerMeditationState != nil {
		panic("delete rollback check failed 'player_meditation_state'")
	}
	log.db.tables.PlayerMeditationState = log.Old
}

func (log *PlayerMeditationStateUpdateLog) Rollback() {
	if log.db.tables.PlayerMeditationState != log.Row {
		panic("update rollback check failed 'player_meditation_state'")
	}
	log.db.tables.PlayerMeditationState = log.Old
	C.Free_PlayerMeditationState(log.Row)
}

/* ========== player_mission ========== */

type PlayerMissionInsertLog struct {
	db *Database 
	Row *C.PlayerMission
	GoRow *PlayerMission
}

func (db *Database) newPlayerMissionInsertLog(row *C.PlayerMission, goRow *PlayerMission) *PlayerMissionInsertLog {
	return &PlayerMissionInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerMissionInsertLog) Free() {
}

func (log *PlayerMissionInsertLog) InvokeHook() {
	if g_Hooks.PlayerMission != nil {
		g_Hooks.PlayerMission.Insert(&PlayerMissionRow{row: log.Row})
	}
}

func (log *PlayerMissionInsertLog) SetEventId(time.Time) {
}

func (log *PlayerMissionInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(98)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.Key)
	file.WriteInt8(log.GoRow.MaxOrder)
}

func (log *PlayerMissionInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMission.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Key)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.MaxOrder)))
	return stmt.Execute()
}

func (log *PlayerMissionInsertLog) CommitToTLog() {
}

func (log *PlayerMissionInsertLog) CommitToXdLog() {
}

func (log *PlayerMissionInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMissionDeleteLog struct {
	db *Database
	Old *C.PlayerMission
	GoOld *PlayerMission
}

func (db *Database) newPlayerMissionDeleteLog(old *C.PlayerMission, goOld *PlayerMission) *PlayerMissionDeleteLog {
	return &PlayerMissionDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerMissionDeleteLog) Free() {
	C.Free_PlayerMission(log.Old)
}

func (log *PlayerMissionDeleteLog) InvokeHook() {
	if g_Hooks.PlayerMission != nil {
		g_Hooks.PlayerMission.Delete(&PlayerMissionRow{row: log.Old})
	}
}

func (log *PlayerMissionDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerMissionDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(98)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Key)
	file.WriteInt8(log.GoOld.MaxOrder)
}

func (log *PlayerMissionDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMission.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerMissionDeleteLog) CommitToTLog() {
}

func (log *PlayerMissionDeleteLog) CommitToXdLog() {
}

func (log *PlayerMissionDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMissionUpdateLog struct {
	db *Database 
	Row *C.PlayerMission
	Old *C.PlayerMission
	GoNew *PlayerMission
	GoOld *PlayerMission
}

func (db *Database) newPlayerMissionUpdateLog(row, old *C.PlayerMission, goNew, goOld *PlayerMission) *PlayerMissionUpdateLog {
	return &PlayerMissionUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerMissionUpdateLog) Free() {
	C.Free_PlayerMission(log.Old)
}

func (log *PlayerMissionUpdateLog) InvokeHook() {
	if g_Hooks.PlayerMission != nil {
		g_Hooks.PlayerMission.Update(&PlayerMissionRow{row: log.Row}, &PlayerMissionRow{row: log.Old})
	}
}

func (log *PlayerMissionUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerMissionUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(98)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.Key)
	file.WriteInt8(log.GoNew.MaxOrder)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Key)
	file.WriteInt8(log.GoOld.MaxOrder)
}

func (log *PlayerMissionUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMission.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.Key)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.MaxOrder)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerMissionUpdateLog) CommitToTLog() {
}

func (log *PlayerMissionUpdateLog) CommitToXdLog() {
}

func (log *PlayerMissionUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerMissionInsertLog) Rollback() {
	if log.db.tables.PlayerMission != log.Row {
		panic("insert rollback check failed 'PlayerMission'")
	}
	log.db.tables.PlayerMission = nil
	C.Free_PlayerMission(log.Row)
}

func (log *PlayerMissionDeleteLog) Rollback() {
	if log.db.tables.PlayerMission != nil {
		panic("delete rollback check failed 'player_mission'")
	}
	log.db.tables.PlayerMission = log.Old
}

func (log *PlayerMissionUpdateLog) Rollback() {
	if log.db.tables.PlayerMission != log.Row {
		panic("update rollback check failed 'player_mission'")
	}
	log.db.tables.PlayerMission = log.Old
	C.Free_PlayerMission(log.Row)
}

/* ========== player_mission_level ========== */

type PlayerMissionLevelInsertLog struct {
	db *Database 
	Row *C.PlayerMissionLevel
	GoRow *PlayerMissionLevel
}

func (db *Database) newPlayerMissionLevelInsertLog(row *C.PlayerMissionLevel, goRow *PlayerMissionLevel) *PlayerMissionLevelInsertLog {
	return &PlayerMissionLevelInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerMissionLevelInsertLog) Free() {
}

func (log *PlayerMissionLevelInsertLog) InvokeHook() {
	if g_Hooks.PlayerMissionLevel != nil {
		g_Hooks.PlayerMissionLevel.Insert(&PlayerMissionLevelRow{row: log.Row})
	}
}

func (log *PlayerMissionLevelInsertLog) SetEventId(time.Time) {
}

func (log *PlayerMissionLevelInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(99)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.Lock)
	file.WriteInt32(log.GoRow.MaxLock)
	file.WriteInt32(log.GoRow.AwardLock)
}

func (log *PlayerMissionLevelInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMissionLevel.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Lock)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.MaxLock)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.AwardLock)))
	return stmt.Execute()
}

func (log *PlayerMissionLevelInsertLog) CommitToTLog() {
}

func (log *PlayerMissionLevelInsertLog) CommitToXdLog() {
}

func (log *PlayerMissionLevelInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMissionLevelDeleteLog struct {
	db *Database
	Old *C.PlayerMissionLevel
	GoOld *PlayerMissionLevel
}

func (db *Database) newPlayerMissionLevelDeleteLog(old *C.PlayerMissionLevel, goOld *PlayerMissionLevel) *PlayerMissionLevelDeleteLog {
	return &PlayerMissionLevelDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerMissionLevelDeleteLog) Free() {
	C.Free_PlayerMissionLevel(log.Old)
}

func (log *PlayerMissionLevelDeleteLog) InvokeHook() {
	if g_Hooks.PlayerMissionLevel != nil {
		g_Hooks.PlayerMissionLevel.Delete(&PlayerMissionLevelRow{row: log.Old})
	}
}

func (log *PlayerMissionLevelDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerMissionLevelDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(99)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Lock)
	file.WriteInt32(log.GoOld.MaxLock)
	file.WriteInt32(log.GoOld.AwardLock)
}

func (log *PlayerMissionLevelDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMissionLevel.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerMissionLevelDeleteLog) CommitToTLog() {
}

func (log *PlayerMissionLevelDeleteLog) CommitToXdLog() {
}

func (log *PlayerMissionLevelDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMissionLevelUpdateLog struct {
	db *Database 
	Row *C.PlayerMissionLevel
	Old *C.PlayerMissionLevel
	GoNew *PlayerMissionLevel
	GoOld *PlayerMissionLevel
}

func (db *Database) newPlayerMissionLevelUpdateLog(row, old *C.PlayerMissionLevel, goNew, goOld *PlayerMissionLevel) *PlayerMissionLevelUpdateLog {
	return &PlayerMissionLevelUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerMissionLevelUpdateLog) Free() {
	C.Free_PlayerMissionLevel(log.Old)
}

func (log *PlayerMissionLevelUpdateLog) InvokeHook() {
	if g_Hooks.PlayerMissionLevel != nil {
		g_Hooks.PlayerMissionLevel.Update(&PlayerMissionLevelRow{row: log.Row}, &PlayerMissionLevelRow{row: log.Old})
	}
}

func (log *PlayerMissionLevelUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerMissionLevelUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(99)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.Lock)
	file.WriteInt32(log.GoNew.MaxLock)
	file.WriteInt32(log.GoNew.AwardLock)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Lock)
	file.WriteInt32(log.GoOld.MaxLock)
	file.WriteInt32(log.GoOld.AwardLock)
}

func (log *PlayerMissionLevelUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMissionLevel.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.Lock)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.MaxLock)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.AwardLock)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerMissionLevelUpdateLog) CommitToTLog() {
}

func (log *PlayerMissionLevelUpdateLog) CommitToXdLog() {
}

func (log *PlayerMissionLevelUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerMissionLevelInsertLog) Rollback() {
	if log.db.tables.PlayerMissionLevel != log.Row {
		panic("insert rollback check failed 'PlayerMissionLevel'")
	}
	log.db.tables.PlayerMissionLevel = nil
	C.Free_PlayerMissionLevel(log.Row)
}

func (log *PlayerMissionLevelDeleteLog) Rollback() {
	if log.db.tables.PlayerMissionLevel != nil {
		panic("delete rollback check failed 'player_mission_level'")
	}
	log.db.tables.PlayerMissionLevel = log.Old
}

func (log *PlayerMissionLevelUpdateLog) Rollback() {
	if log.db.tables.PlayerMissionLevel != log.Row {
		panic("update rollback check failed 'player_mission_level'")
	}
	log.db.tables.PlayerMissionLevel = log.Old
	C.Free_PlayerMissionLevel(log.Row)
}

/* ========== player_mission_level_record ========== */

type PlayerMissionLevelRecordInsertLog struct {
	db *Database 
	Row *C.PlayerMissionLevelRecord
	GoRow *PlayerMissionLevelRecord
}

func (db *Database) newPlayerMissionLevelRecordInsertLog(row *C.PlayerMissionLevelRecord, goRow *PlayerMissionLevelRecord) *PlayerMissionLevelRecordInsertLog {
	return &PlayerMissionLevelRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerMissionLevelRecordInsertLog) Free() {
}

func (log *PlayerMissionLevelRecordInsertLog) InvokeHook() {
	if g_Hooks.PlayerMissionLevelRecord != nil {
		g_Hooks.PlayerMissionLevelRecord.Insert(&PlayerMissionLevelRecordRow{row: log.Row})
	}
}

func (log *PlayerMissionLevelRecordInsertLog) SetEventId(time.Time) {
}

func (log *PlayerMissionLevelRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(100)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.MissionId)
	file.WriteInt32(log.GoRow.MissionLevelId)
	file.WriteInt64(log.GoRow.OpenTime)
	file.WriteInt32(log.GoRow.Score)
	file.WriteInt8(log.GoRow.Round)
	file.WriteInt8(log.GoRow.DailyNum)
	file.WriteInt64(log.GoRow.LastEnterTime)
	file.WriteInt16(log.GoRow.EmptyShadowBits)
}

func (log *PlayerMissionLevelRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMissionLevelRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.MissionId)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.MissionLevelId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.OpenTime)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Score)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Round)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DailyNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastEnterTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.EmptyShadowBits)))
	return stmt.Execute()
}

func (log *PlayerMissionLevelRecordInsertLog) CommitToTLog() {
}

func (log *PlayerMissionLevelRecordInsertLog) CommitToXdLog() {
}

func (log *PlayerMissionLevelRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMissionLevelRecordDeleteLog struct {
	db *Database
	Old *C.PlayerMissionLevelRecord
	GoOld *PlayerMissionLevelRecord
}

func (db *Database) newPlayerMissionLevelRecordDeleteLog(old *C.PlayerMissionLevelRecord, goOld *PlayerMissionLevelRecord) *PlayerMissionLevelRecordDeleteLog {
	return &PlayerMissionLevelRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerMissionLevelRecordDeleteLog) Free() {
	C.Free_PlayerMissionLevelRecord(log.Old)
}

func (log *PlayerMissionLevelRecordDeleteLog) InvokeHook() {
	if g_Hooks.PlayerMissionLevelRecord != nil {
		g_Hooks.PlayerMissionLevelRecord.Delete(&PlayerMissionLevelRecordRow{row: log.Old})
	}
}

func (log *PlayerMissionLevelRecordDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerMissionLevelRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(100)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.MissionId)
	file.WriteInt32(log.GoOld.MissionLevelId)
	file.WriteInt64(log.GoOld.OpenTime)
	file.WriteInt32(log.GoOld.Score)
	file.WriteInt8(log.GoOld.Round)
	file.WriteInt8(log.GoOld.DailyNum)
	file.WriteInt64(log.GoOld.LastEnterTime)
	file.WriteInt16(log.GoOld.EmptyShadowBits)
}

func (log *PlayerMissionLevelRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMissionLevelRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerMissionLevelRecordDeleteLog) CommitToTLog() {
}

func (log *PlayerMissionLevelRecordDeleteLog) CommitToXdLog() {
}

func (log *PlayerMissionLevelRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMissionLevelRecordUpdateLog struct {
	db *Database 
	Row *C.PlayerMissionLevelRecord
	Old *C.PlayerMissionLevelRecord
	GoNew *PlayerMissionLevelRecord
	GoOld *PlayerMissionLevelRecord
}

func (db *Database) newPlayerMissionLevelRecordUpdateLog(row, old *C.PlayerMissionLevelRecord, goNew, goOld *PlayerMissionLevelRecord) *PlayerMissionLevelRecordUpdateLog {
	return &PlayerMissionLevelRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerMissionLevelRecordUpdateLog) Free() {
	C.Free_PlayerMissionLevelRecord(log.Old)
}

func (log *PlayerMissionLevelRecordUpdateLog) InvokeHook() {
	if g_Hooks.PlayerMissionLevelRecord != nil {
		g_Hooks.PlayerMissionLevelRecord.Update(&PlayerMissionLevelRecordRow{row: log.Row}, &PlayerMissionLevelRecordRow{row: log.Old})
	}
}

func (log *PlayerMissionLevelRecordUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerMissionLevelRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(100)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.MissionId)
	file.WriteInt32(log.GoNew.MissionLevelId)
	file.WriteInt64(log.GoNew.OpenTime)
	file.WriteInt32(log.GoNew.Score)
	file.WriteInt8(log.GoNew.Round)
	file.WriteInt8(log.GoNew.DailyNum)
	file.WriteInt64(log.GoNew.LastEnterTime)
	file.WriteInt16(log.GoNew.EmptyShadowBits)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.MissionId)
	file.WriteInt32(log.GoOld.MissionLevelId)
	file.WriteInt64(log.GoOld.OpenTime)
	file.WriteInt32(log.GoOld.Score)
	file.WriteInt8(log.GoOld.Round)
	file.WriteInt8(log.GoOld.DailyNum)
	file.WriteInt64(log.GoOld.LastEnterTime)
	file.WriteInt16(log.GoOld.EmptyShadowBits)
}

func (log *PlayerMissionLevelRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMissionLevelRecord.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.MissionId)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.MissionLevelId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.OpenTime)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Score)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Round)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DailyNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastEnterTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.EmptyShadowBits)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerMissionLevelRecordUpdateLog) CommitToTLog() {
}

func (log *PlayerMissionLevelRecordUpdateLog) CommitToXdLog() {
}

func (log *PlayerMissionLevelRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerMissionLevelRecordInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerMissionLevelRecord
	for crow := log.db.tables.PlayerMissionLevelRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerMissionLevelRecord'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerMissionLevelRecord = log.db.tables.PlayerMissionLevelRecord.next
	}
	C.Free_PlayerMissionLevelRecord(log.Row)
}

func (log *PlayerMissionLevelRecordDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerMissionLevelRecord
	for crow := log.db.tables.PlayerMissionLevelRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_mission_level_record'")
	}
	log.Old.next = log.db.tables.PlayerMissionLevelRecord
	log.db.tables.PlayerMissionLevelRecord = log.Old
}

func (log *PlayerMissionLevelRecordUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerMissionLevelRecord
	for crow := log.db.tables.PlayerMissionLevelRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_mission_level_record'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerMissionLevelRecord = log.Old
	}
	C.Free_PlayerMissionLevelRecord(log.Row)
}

/* ========== player_mission_level_state_bin ========== */

type PlayerMissionLevelStateBinInsertLog struct {
	db *Database 
	Row *C.PlayerMissionLevelStateBin
	GoRow *PlayerMissionLevelStateBin
}

func (db *Database) newPlayerMissionLevelStateBinInsertLog(row *C.PlayerMissionLevelStateBin, goRow *PlayerMissionLevelStateBin) *PlayerMissionLevelStateBinInsertLog {
	return &PlayerMissionLevelStateBinInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerMissionLevelStateBinInsertLog) Free() {
}

func (log *PlayerMissionLevelStateBinInsertLog) InvokeHook() {
	if g_Hooks.PlayerMissionLevelStateBin != nil {
		g_Hooks.PlayerMissionLevelStateBin.Insert(&PlayerMissionLevelStateBinRow{row: log.Row})
	}
}

func (log *PlayerMissionLevelStateBinInsertLog) SetEventId(time.Time) {
}

func (log *PlayerMissionLevelStateBinInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(101)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteBytes(log.GoRow.Bin)
}

func (log *PlayerMissionLevelStateBinInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMissionLevelStateBin.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBinary(unsafe.Pointer(log.Row.Bin), int(log.Row.Bin_len))
	return stmt.Execute()
}

func (log *PlayerMissionLevelStateBinInsertLog) CommitToTLog() {
}

func (log *PlayerMissionLevelStateBinInsertLog) CommitToXdLog() {
}

func (log *PlayerMissionLevelStateBinInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMissionLevelStateBinDeleteLog struct {
	db *Database
	Old *C.PlayerMissionLevelStateBin
	GoOld *PlayerMissionLevelStateBin
}

func (db *Database) newPlayerMissionLevelStateBinDeleteLog(old *C.PlayerMissionLevelStateBin, goOld *PlayerMissionLevelStateBin) *PlayerMissionLevelStateBinDeleteLog {
	return &PlayerMissionLevelStateBinDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerMissionLevelStateBinDeleteLog) Free() {
	C.Free_PlayerMissionLevelStateBin(log.Old)
}

func (log *PlayerMissionLevelStateBinDeleteLog) InvokeHook() {
	if g_Hooks.PlayerMissionLevelStateBin != nil {
		g_Hooks.PlayerMissionLevelStateBin.Delete(&PlayerMissionLevelStateBinRow{row: log.Old})
	}
}

func (log *PlayerMissionLevelStateBinDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerMissionLevelStateBinDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(101)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteBytes(log.GoOld.Bin)
}

func (log *PlayerMissionLevelStateBinDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMissionLevelStateBin.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerMissionLevelStateBinDeleteLog) CommitToTLog() {
}

func (log *PlayerMissionLevelStateBinDeleteLog) CommitToXdLog() {
}

func (log *PlayerMissionLevelStateBinDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMissionLevelStateBinUpdateLog struct {
	db *Database 
	Row *C.PlayerMissionLevelStateBin
	Old *C.PlayerMissionLevelStateBin
	GoNew *PlayerMissionLevelStateBin
	GoOld *PlayerMissionLevelStateBin
}

func (db *Database) newPlayerMissionLevelStateBinUpdateLog(row, old *C.PlayerMissionLevelStateBin, goNew, goOld *PlayerMissionLevelStateBin) *PlayerMissionLevelStateBinUpdateLog {
	return &PlayerMissionLevelStateBinUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerMissionLevelStateBinUpdateLog) Free() {
	C.Free_PlayerMissionLevelStateBin(log.Old)
}

func (log *PlayerMissionLevelStateBinUpdateLog) InvokeHook() {
	if g_Hooks.PlayerMissionLevelStateBin != nil {
		g_Hooks.PlayerMissionLevelStateBin.Update(&PlayerMissionLevelStateBinRow{row: log.Row}, &PlayerMissionLevelStateBinRow{row: log.Old})
	}
}

func (log *PlayerMissionLevelStateBinUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerMissionLevelStateBinUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(101)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteBytes(log.GoNew.Bin)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteBytes(log.GoOld.Bin)
}

func (log *PlayerMissionLevelStateBinUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMissionLevelStateBin.Update
	stmt.BindBinary(unsafe.Pointer(log.Row.Bin), int(log.Row.Bin_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerMissionLevelStateBinUpdateLog) CommitToTLog() {
}

func (log *PlayerMissionLevelStateBinUpdateLog) CommitToXdLog() {
}

func (log *PlayerMissionLevelStateBinUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerMissionLevelStateBinInsertLog) Rollback() {
	if log.db.tables.PlayerMissionLevelStateBin != log.Row {
		panic("insert rollback check failed 'PlayerMissionLevelStateBin'")
	}
	log.db.tables.PlayerMissionLevelStateBin = nil
	C.Free_PlayerMissionLevelStateBin(log.Row)
}

func (log *PlayerMissionLevelStateBinDeleteLog) Rollback() {
	if log.db.tables.PlayerMissionLevelStateBin != nil {
		panic("delete rollback check failed 'player_mission_level_state_bin'")
	}
	log.db.tables.PlayerMissionLevelStateBin = log.Old
}

func (log *PlayerMissionLevelStateBinUpdateLog) Rollback() {
	if log.db.tables.PlayerMissionLevelStateBin != log.Row {
		panic("update rollback check failed 'player_mission_level_state_bin'")
	}
	log.db.tables.PlayerMissionLevelStateBin = log.Old
	C.Free_PlayerMissionLevelStateBin(log.Row)
}

/* ========== player_mission_record ========== */

type PlayerMissionRecordInsertLog struct {
	db *Database 
	Row *C.PlayerMissionRecord
	GoRow *PlayerMissionRecord
}

func (db *Database) newPlayerMissionRecordInsertLog(row *C.PlayerMissionRecord, goRow *PlayerMissionRecord) *PlayerMissionRecordInsertLog {
	return &PlayerMissionRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerMissionRecordInsertLog) Free() {
}

func (log *PlayerMissionRecordInsertLog) InvokeHook() {
	if g_Hooks.PlayerMissionRecord != nil {
		g_Hooks.PlayerMissionRecord.Insert(&PlayerMissionRecordRow{row: log.Row})
	}
}

func (log *PlayerMissionRecordInsertLog) SetEventId(time.Time) {
}

func (log *PlayerMissionRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(102)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.TownId)
	file.WriteInt16(log.GoRow.MissionId)
	file.WriteInt64(log.GoRow.OpenTime)
}

func (log *PlayerMissionRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMissionRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TownId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.MissionId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.OpenTime)))
	return stmt.Execute()
}

func (log *PlayerMissionRecordInsertLog) CommitToTLog() {
}

func (log *PlayerMissionRecordInsertLog) CommitToXdLog() {
}

func (log *PlayerMissionRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMissionRecordDeleteLog struct {
	db *Database
	Old *C.PlayerMissionRecord
	GoOld *PlayerMissionRecord
}

func (db *Database) newPlayerMissionRecordDeleteLog(old *C.PlayerMissionRecord, goOld *PlayerMissionRecord) *PlayerMissionRecordDeleteLog {
	return &PlayerMissionRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerMissionRecordDeleteLog) Free() {
	C.Free_PlayerMissionRecord(log.Old)
}

func (log *PlayerMissionRecordDeleteLog) InvokeHook() {
	if g_Hooks.PlayerMissionRecord != nil {
		g_Hooks.PlayerMissionRecord.Delete(&PlayerMissionRecordRow{row: log.Old})
	}
}

func (log *PlayerMissionRecordDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerMissionRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(102)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.TownId)
	file.WriteInt16(log.GoOld.MissionId)
	file.WriteInt64(log.GoOld.OpenTime)
}

func (log *PlayerMissionRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMissionRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerMissionRecordDeleteLog) CommitToTLog() {
}

func (log *PlayerMissionRecordDeleteLog) CommitToXdLog() {
}

func (log *PlayerMissionRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMissionRecordUpdateLog struct {
	db *Database 
	Row *C.PlayerMissionRecord
	Old *C.PlayerMissionRecord
	GoNew *PlayerMissionRecord
	GoOld *PlayerMissionRecord
}

func (db *Database) newPlayerMissionRecordUpdateLog(row, old *C.PlayerMissionRecord, goNew, goOld *PlayerMissionRecord) *PlayerMissionRecordUpdateLog {
	return &PlayerMissionRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerMissionRecordUpdateLog) Free() {
	C.Free_PlayerMissionRecord(log.Old)
}

func (log *PlayerMissionRecordUpdateLog) InvokeHook() {
	if g_Hooks.PlayerMissionRecord != nil {
		g_Hooks.PlayerMissionRecord.Update(&PlayerMissionRecordRow{row: log.Row}, &PlayerMissionRecordRow{row: log.Old})
	}
}

func (log *PlayerMissionRecordUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerMissionRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(102)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.TownId)
	file.WriteInt16(log.GoNew.MissionId)
	file.WriteInt64(log.GoNew.OpenTime)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.TownId)
	file.WriteInt16(log.GoOld.MissionId)
	file.WriteInt64(log.GoOld.OpenTime)
}

func (log *PlayerMissionRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMissionRecord.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TownId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.MissionId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.OpenTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerMissionRecordUpdateLog) CommitToTLog() {
}

func (log *PlayerMissionRecordUpdateLog) CommitToXdLog() {
}

func (log *PlayerMissionRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerMissionRecordInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerMissionRecord
	for crow := log.db.tables.PlayerMissionRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerMissionRecord'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerMissionRecord = log.db.tables.PlayerMissionRecord.next
	}
	C.Free_PlayerMissionRecord(log.Row)
}

func (log *PlayerMissionRecordDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerMissionRecord
	for crow := log.db.tables.PlayerMissionRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_mission_record'")
	}
	log.Old.next = log.db.tables.PlayerMissionRecord
	log.db.tables.PlayerMissionRecord = log.Old
}

func (log *PlayerMissionRecordUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerMissionRecord
	for crow := log.db.tables.PlayerMissionRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_mission_record'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerMissionRecord = log.Old
	}
	C.Free_PlayerMissionRecord(log.Row)
}

/* ========== player_mission_star_award ========== */

type PlayerMissionStarAwardInsertLog struct {
	db *Database 
	Row *C.PlayerMissionStarAward
	GoRow *PlayerMissionStarAward
}

func (db *Database) newPlayerMissionStarAwardInsertLog(row *C.PlayerMissionStarAward, goRow *PlayerMissionStarAward) *PlayerMissionStarAwardInsertLog {
	return &PlayerMissionStarAwardInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerMissionStarAwardInsertLog) Free() {
}

func (log *PlayerMissionStarAwardInsertLog) InvokeHook() {
	if g_Hooks.PlayerMissionStarAward != nil {
		g_Hooks.PlayerMissionStarAward.Insert(&PlayerMissionStarAwardRow{row: log.Row})
	}
}

func (log *PlayerMissionStarAwardInsertLog) SetEventId(time.Time) {
}

func (log *PlayerMissionStarAwardInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(103)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.TownId)
	file.WriteInt8(log.GoRow.BoxType)
}

func (log *PlayerMissionStarAwardInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMissionStarAward.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TownId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.BoxType)))
	return stmt.Execute()
}

func (log *PlayerMissionStarAwardInsertLog) CommitToTLog() {
}

func (log *PlayerMissionStarAwardInsertLog) CommitToXdLog() {
}

func (log *PlayerMissionStarAwardInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMissionStarAwardDeleteLog struct {
	db *Database
	Old *C.PlayerMissionStarAward
	GoOld *PlayerMissionStarAward
}

func (db *Database) newPlayerMissionStarAwardDeleteLog(old *C.PlayerMissionStarAward, goOld *PlayerMissionStarAward) *PlayerMissionStarAwardDeleteLog {
	return &PlayerMissionStarAwardDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerMissionStarAwardDeleteLog) Free() {
	C.Free_PlayerMissionStarAward(log.Old)
}

func (log *PlayerMissionStarAwardDeleteLog) InvokeHook() {
	if g_Hooks.PlayerMissionStarAward != nil {
		g_Hooks.PlayerMissionStarAward.Delete(&PlayerMissionStarAwardRow{row: log.Old})
	}
}

func (log *PlayerMissionStarAwardDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerMissionStarAwardDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(103)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.TownId)
	file.WriteInt8(log.GoOld.BoxType)
}

func (log *PlayerMissionStarAwardDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMissionStarAward.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerMissionStarAwardDeleteLog) CommitToTLog() {
}

func (log *PlayerMissionStarAwardDeleteLog) CommitToXdLog() {
}

func (log *PlayerMissionStarAwardDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMissionStarAwardUpdateLog struct {
	db *Database 
	Row *C.PlayerMissionStarAward
	Old *C.PlayerMissionStarAward
	GoNew *PlayerMissionStarAward
	GoOld *PlayerMissionStarAward
}

func (db *Database) newPlayerMissionStarAwardUpdateLog(row, old *C.PlayerMissionStarAward, goNew, goOld *PlayerMissionStarAward) *PlayerMissionStarAwardUpdateLog {
	return &PlayerMissionStarAwardUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerMissionStarAwardUpdateLog) Free() {
	C.Free_PlayerMissionStarAward(log.Old)
}

func (log *PlayerMissionStarAwardUpdateLog) InvokeHook() {
	if g_Hooks.PlayerMissionStarAward != nil {
		g_Hooks.PlayerMissionStarAward.Update(&PlayerMissionStarAwardRow{row: log.Row}, &PlayerMissionStarAwardRow{row: log.Old})
	}
}

func (log *PlayerMissionStarAwardUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerMissionStarAwardUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(103)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.TownId)
	file.WriteInt8(log.GoNew.BoxType)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.TownId)
	file.WriteInt8(log.GoOld.BoxType)
}

func (log *PlayerMissionStarAwardUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMissionStarAward.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TownId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.BoxType)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerMissionStarAwardUpdateLog) CommitToTLog() {
}

func (log *PlayerMissionStarAwardUpdateLog) CommitToXdLog() {
}

func (log *PlayerMissionStarAwardUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerMissionStarAwardInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerMissionStarAward
	for crow := log.db.tables.PlayerMissionStarAward; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerMissionStarAward'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerMissionStarAward = log.db.tables.PlayerMissionStarAward.next
	}
	C.Free_PlayerMissionStarAward(log.Row)
}

func (log *PlayerMissionStarAwardDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerMissionStarAward
	for crow := log.db.tables.PlayerMissionStarAward; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_mission_star_award'")
	}
	log.Old.next = log.db.tables.PlayerMissionStarAward
	log.db.tables.PlayerMissionStarAward = log.Old
}

func (log *PlayerMissionStarAwardUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerMissionStarAward
	for crow := log.db.tables.PlayerMissionStarAward; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_mission_star_award'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerMissionStarAward = log.Old
	}
	C.Free_PlayerMissionStarAward(log.Row)
}

/* ========== player_money_tree ========== */

type PlayerMoneyTreeInsertLog struct {
	db *Database 
	Row *C.PlayerMoneyTree
	GoRow *PlayerMoneyTree
}

func (db *Database) newPlayerMoneyTreeInsertLog(row *C.PlayerMoneyTree, goRow *PlayerMoneyTree) *PlayerMoneyTreeInsertLog {
	return &PlayerMoneyTreeInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerMoneyTreeInsertLog) Free() {
}

func (log *PlayerMoneyTreeInsertLog) InvokeHook() {
	if g_Hooks.PlayerMoneyTree != nil {
		g_Hooks.PlayerMoneyTree.Insert(&PlayerMoneyTreeRow{row: log.Row})
	}
}

func (log *PlayerMoneyTreeInsertLog) SetEventId(time.Time) {
}

func (log *PlayerMoneyTreeInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(104)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.Total)
	file.WriteInt8(log.GoRow.WavedTimesTotal)
	file.WriteInt8(log.GoRow.WavedTimes)
	file.WriteInt64(log.GoRow.LastWavedTime)
}

func (log *PlayerMoneyTreeInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMoneyTree.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Total)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.WavedTimesTotal)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.WavedTimes)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastWavedTime)))
	return stmt.Execute()
}

func (log *PlayerMoneyTreeInsertLog) CommitToTLog() {
}

func (log *PlayerMoneyTreeInsertLog) CommitToXdLog() {
}

func (log *PlayerMoneyTreeInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMoneyTreeDeleteLog struct {
	db *Database
	Old *C.PlayerMoneyTree
	GoOld *PlayerMoneyTree
}

func (db *Database) newPlayerMoneyTreeDeleteLog(old *C.PlayerMoneyTree, goOld *PlayerMoneyTree) *PlayerMoneyTreeDeleteLog {
	return &PlayerMoneyTreeDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerMoneyTreeDeleteLog) Free() {
	C.Free_PlayerMoneyTree(log.Old)
}

func (log *PlayerMoneyTreeDeleteLog) InvokeHook() {
	if g_Hooks.PlayerMoneyTree != nil {
		g_Hooks.PlayerMoneyTree.Delete(&PlayerMoneyTreeRow{row: log.Old})
	}
}

func (log *PlayerMoneyTreeDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerMoneyTreeDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(104)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Total)
	file.WriteInt8(log.GoOld.WavedTimesTotal)
	file.WriteInt8(log.GoOld.WavedTimes)
	file.WriteInt64(log.GoOld.LastWavedTime)
}

func (log *PlayerMoneyTreeDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMoneyTree.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerMoneyTreeDeleteLog) CommitToTLog() {
}

func (log *PlayerMoneyTreeDeleteLog) CommitToXdLog() {
}

func (log *PlayerMoneyTreeDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMoneyTreeUpdateLog struct {
	db *Database 
	Row *C.PlayerMoneyTree
	Old *C.PlayerMoneyTree
	GoNew *PlayerMoneyTree
	GoOld *PlayerMoneyTree
}

func (db *Database) newPlayerMoneyTreeUpdateLog(row, old *C.PlayerMoneyTree, goNew, goOld *PlayerMoneyTree) *PlayerMoneyTreeUpdateLog {
	return &PlayerMoneyTreeUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerMoneyTreeUpdateLog) Free() {
	C.Free_PlayerMoneyTree(log.Old)
}

func (log *PlayerMoneyTreeUpdateLog) InvokeHook() {
	if g_Hooks.PlayerMoneyTree != nil {
		g_Hooks.PlayerMoneyTree.Update(&PlayerMoneyTreeRow{row: log.Row}, &PlayerMoneyTreeRow{row: log.Old})
	}
}

func (log *PlayerMoneyTreeUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerMoneyTreeUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(104)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.Total)
	file.WriteInt8(log.GoNew.WavedTimesTotal)
	file.WriteInt8(log.GoNew.WavedTimes)
	file.WriteInt64(log.GoNew.LastWavedTime)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Total)
	file.WriteInt8(log.GoOld.WavedTimesTotal)
	file.WriteInt8(log.GoOld.WavedTimes)
	file.WriteInt64(log.GoOld.LastWavedTime)
}

func (log *PlayerMoneyTreeUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMoneyTree.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.Total)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.WavedTimesTotal)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.WavedTimes)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastWavedTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerMoneyTreeUpdateLog) CommitToTLog() {
}

func (log *PlayerMoneyTreeUpdateLog) CommitToXdLog() {
}

func (log *PlayerMoneyTreeUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerMoneyTreeInsertLog) Rollback() {
	if log.db.tables.PlayerMoneyTree != log.Row {
		panic("insert rollback check failed 'PlayerMoneyTree'")
	}
	log.db.tables.PlayerMoneyTree = nil
	C.Free_PlayerMoneyTree(log.Row)
}

func (log *PlayerMoneyTreeDeleteLog) Rollback() {
	if log.db.tables.PlayerMoneyTree != nil {
		panic("delete rollback check failed 'player_money_tree'")
	}
	log.db.tables.PlayerMoneyTree = log.Old
}

func (log *PlayerMoneyTreeUpdateLog) Rollback() {
	if log.db.tables.PlayerMoneyTree != log.Row {
		panic("update rollback check failed 'player_money_tree'")
	}
	log.db.tables.PlayerMoneyTree = log.Old
	C.Free_PlayerMoneyTree(log.Row)
}

/* ========== player_month_card_award_record ========== */

type PlayerMonthCardAwardRecordInsertLog struct {
	db *Database 
	Row *C.PlayerMonthCardAwardRecord
	GoRow *PlayerMonthCardAwardRecord
}

func (db *Database) newPlayerMonthCardAwardRecordInsertLog(row *C.PlayerMonthCardAwardRecord, goRow *PlayerMonthCardAwardRecord) *PlayerMonthCardAwardRecordInsertLog {
	return &PlayerMonthCardAwardRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerMonthCardAwardRecordInsertLog) Free() {
}

func (log *PlayerMonthCardAwardRecordInsertLog) InvokeHook() {
	if g_Hooks.PlayerMonthCardAwardRecord != nil {
		g_Hooks.PlayerMonthCardAwardRecord.Insert(&PlayerMonthCardAwardRecordRow{row: log.Row})
	}
}

func (log *PlayerMonthCardAwardRecordInsertLog) SetEventId(time.Time) {
}

func (log *PlayerMonthCardAwardRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(105)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.LastUpdate)
}

func (log *PlayerMonthCardAwardRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMonthCardAwardRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastUpdate)))
	return stmt.Execute()
}

func (log *PlayerMonthCardAwardRecordInsertLog) CommitToTLog() {
}

func (log *PlayerMonthCardAwardRecordInsertLog) CommitToXdLog() {
}

func (log *PlayerMonthCardAwardRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMonthCardAwardRecordDeleteLog struct {
	db *Database
	Old *C.PlayerMonthCardAwardRecord
	GoOld *PlayerMonthCardAwardRecord
}

func (db *Database) newPlayerMonthCardAwardRecordDeleteLog(old *C.PlayerMonthCardAwardRecord, goOld *PlayerMonthCardAwardRecord) *PlayerMonthCardAwardRecordDeleteLog {
	return &PlayerMonthCardAwardRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerMonthCardAwardRecordDeleteLog) Free() {
	C.Free_PlayerMonthCardAwardRecord(log.Old)
}

func (log *PlayerMonthCardAwardRecordDeleteLog) InvokeHook() {
	if g_Hooks.PlayerMonthCardAwardRecord != nil {
		g_Hooks.PlayerMonthCardAwardRecord.Delete(&PlayerMonthCardAwardRecordRow{row: log.Old})
	}
}

func (log *PlayerMonthCardAwardRecordDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerMonthCardAwardRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(105)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.LastUpdate)
}

func (log *PlayerMonthCardAwardRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMonthCardAwardRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerMonthCardAwardRecordDeleteLog) CommitToTLog() {
}

func (log *PlayerMonthCardAwardRecordDeleteLog) CommitToXdLog() {
}

func (log *PlayerMonthCardAwardRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMonthCardAwardRecordUpdateLog struct {
	db *Database 
	Row *C.PlayerMonthCardAwardRecord
	Old *C.PlayerMonthCardAwardRecord
	GoNew *PlayerMonthCardAwardRecord
	GoOld *PlayerMonthCardAwardRecord
}

func (db *Database) newPlayerMonthCardAwardRecordUpdateLog(row, old *C.PlayerMonthCardAwardRecord, goNew, goOld *PlayerMonthCardAwardRecord) *PlayerMonthCardAwardRecordUpdateLog {
	return &PlayerMonthCardAwardRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerMonthCardAwardRecordUpdateLog) Free() {
	C.Free_PlayerMonthCardAwardRecord(log.Old)
}

func (log *PlayerMonthCardAwardRecordUpdateLog) InvokeHook() {
	if g_Hooks.PlayerMonthCardAwardRecord != nil {
		g_Hooks.PlayerMonthCardAwardRecord.Update(&PlayerMonthCardAwardRecordRow{row: log.Row}, &PlayerMonthCardAwardRecordRow{row: log.Old})
	}
}

func (log *PlayerMonthCardAwardRecordUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerMonthCardAwardRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(105)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.LastUpdate)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.LastUpdate)
}

func (log *PlayerMonthCardAwardRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMonthCardAwardRecord.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastUpdate)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerMonthCardAwardRecordUpdateLog) CommitToTLog() {
}

func (log *PlayerMonthCardAwardRecordUpdateLog) CommitToXdLog() {
}

func (log *PlayerMonthCardAwardRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerMonthCardAwardRecordInsertLog) Rollback() {
	if log.db.tables.PlayerMonthCardAwardRecord != log.Row {
		panic("insert rollback check failed 'PlayerMonthCardAwardRecord'")
	}
	log.db.tables.PlayerMonthCardAwardRecord = nil
	C.Free_PlayerMonthCardAwardRecord(log.Row)
}

func (log *PlayerMonthCardAwardRecordDeleteLog) Rollback() {
	if log.db.tables.PlayerMonthCardAwardRecord != nil {
		panic("delete rollback check failed 'player_month_card_award_record'")
	}
	log.db.tables.PlayerMonthCardAwardRecord = log.Old
}

func (log *PlayerMonthCardAwardRecordUpdateLog) Rollback() {
	if log.db.tables.PlayerMonthCardAwardRecord != log.Row {
		panic("update rollback check failed 'player_month_card_award_record'")
	}
	log.db.tables.PlayerMonthCardAwardRecord = log.Old
	C.Free_PlayerMonthCardAwardRecord(log.Row)
}

/* ========== player_month_card_info ========== */

type PlayerMonthCardInfoInsertLog struct {
	db *Database 
	Row *C.PlayerMonthCardInfo
	GoRow *PlayerMonthCardInfo
}

func (db *Database) newPlayerMonthCardInfoInsertLog(row *C.PlayerMonthCardInfo, goRow *PlayerMonthCardInfo) *PlayerMonthCardInfoInsertLog {
	return &PlayerMonthCardInfoInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerMonthCardInfoInsertLog) Free() {
}

func (log *PlayerMonthCardInfoInsertLog) InvokeHook() {
	if g_Hooks.PlayerMonthCardInfo != nil {
		g_Hooks.PlayerMonthCardInfo.Insert(&PlayerMonthCardInfoRow{row: log.Row})
	}
}

func (log *PlayerMonthCardInfoInsertLog) SetEventId(time.Time) {
}

func (log *PlayerMonthCardInfoInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(106)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.Starttime)
	file.WriteInt64(log.GoRow.Endtime)
	file.WriteInt64(log.GoRow.Buytimes)
	file.WriteInt64(log.GoRow.Presenttotal)
}

func (log *PlayerMonthCardInfoInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMonthCardInfo.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Starttime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Endtime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Buytimes)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Presenttotal)))
	return stmt.Execute()
}

func (log *PlayerMonthCardInfoInsertLog) CommitToTLog() {
}

func (log *PlayerMonthCardInfoInsertLog) CommitToXdLog() {
}

func (log *PlayerMonthCardInfoInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMonthCardInfoDeleteLog struct {
	db *Database
	Old *C.PlayerMonthCardInfo
	GoOld *PlayerMonthCardInfo
}

func (db *Database) newPlayerMonthCardInfoDeleteLog(old *C.PlayerMonthCardInfo, goOld *PlayerMonthCardInfo) *PlayerMonthCardInfoDeleteLog {
	return &PlayerMonthCardInfoDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerMonthCardInfoDeleteLog) Free() {
	C.Free_PlayerMonthCardInfo(log.Old)
}

func (log *PlayerMonthCardInfoDeleteLog) InvokeHook() {
	if g_Hooks.PlayerMonthCardInfo != nil {
		g_Hooks.PlayerMonthCardInfo.Delete(&PlayerMonthCardInfoRow{row: log.Old})
	}
}

func (log *PlayerMonthCardInfoDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerMonthCardInfoDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(106)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.Starttime)
	file.WriteInt64(log.GoOld.Endtime)
	file.WriteInt64(log.GoOld.Buytimes)
	file.WriteInt64(log.GoOld.Presenttotal)
}

func (log *PlayerMonthCardInfoDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMonthCardInfo.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerMonthCardInfoDeleteLog) CommitToTLog() {
}

func (log *PlayerMonthCardInfoDeleteLog) CommitToXdLog() {
}

func (log *PlayerMonthCardInfoDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMonthCardInfoUpdateLog struct {
	db *Database 
	Row *C.PlayerMonthCardInfo
	Old *C.PlayerMonthCardInfo
	GoNew *PlayerMonthCardInfo
	GoOld *PlayerMonthCardInfo
}

func (db *Database) newPlayerMonthCardInfoUpdateLog(row, old *C.PlayerMonthCardInfo, goNew, goOld *PlayerMonthCardInfo) *PlayerMonthCardInfoUpdateLog {
	return &PlayerMonthCardInfoUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerMonthCardInfoUpdateLog) Free() {
	C.Free_PlayerMonthCardInfo(log.Old)
}

func (log *PlayerMonthCardInfoUpdateLog) InvokeHook() {
	if g_Hooks.PlayerMonthCardInfo != nil {
		g_Hooks.PlayerMonthCardInfo.Update(&PlayerMonthCardInfoRow{row: log.Row}, &PlayerMonthCardInfoRow{row: log.Old})
	}
}

func (log *PlayerMonthCardInfoUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerMonthCardInfoUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(106)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.Starttime)
	file.WriteInt64(log.GoNew.Endtime)
	file.WriteInt64(log.GoNew.Buytimes)
	file.WriteInt64(log.GoNew.Presenttotal)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.Starttime)
	file.WriteInt64(log.GoOld.Endtime)
	file.WriteInt64(log.GoOld.Buytimes)
	file.WriteInt64(log.GoOld.Presenttotal)
}

func (log *PlayerMonthCardInfoUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMonthCardInfo.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Starttime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Endtime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Buytimes)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Presenttotal)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerMonthCardInfoUpdateLog) CommitToTLog() {
}

func (log *PlayerMonthCardInfoUpdateLog) CommitToXdLog() {
}

func (log *PlayerMonthCardInfoUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerMonthCardInfoInsertLog) Rollback() {
	if log.db.tables.PlayerMonthCardInfo != log.Row {
		panic("insert rollback check failed 'PlayerMonthCardInfo'")
	}
	log.db.tables.PlayerMonthCardInfo = nil
	C.Free_PlayerMonthCardInfo(log.Row)
}

func (log *PlayerMonthCardInfoDeleteLog) Rollback() {
	if log.db.tables.PlayerMonthCardInfo != nil {
		panic("delete rollback check failed 'player_month_card_info'")
	}
	log.db.tables.PlayerMonthCardInfo = log.Old
}

func (log *PlayerMonthCardInfoUpdateLog) Rollback() {
	if log.db.tables.PlayerMonthCardInfo != log.Row {
		panic("update rollback check failed 'player_month_card_info'")
	}
	log.db.tables.PlayerMonthCardInfo = log.Old
	C.Free_PlayerMonthCardInfo(log.Row)
}

/* ========== player_multi_level_info ========== */

type PlayerMultiLevelInfoInsertLog struct {
	db *Database 
	Row *C.PlayerMultiLevelInfo
	GoRow *PlayerMultiLevelInfo
}

func (db *Database) newPlayerMultiLevelInfoInsertLog(row *C.PlayerMultiLevelInfo, goRow *PlayerMultiLevelInfo) *PlayerMultiLevelInfoInsertLog {
	return &PlayerMultiLevelInfoInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerMultiLevelInfoInsertLog) Free() {
}

func (log *PlayerMultiLevelInfoInsertLog) InvokeHook() {
	if g_Hooks.PlayerMultiLevelInfo != nil {
		g_Hooks.PlayerMultiLevelInfo.Insert(&PlayerMultiLevelInfoRow{row: log.Row})
	}
}

func (log *PlayerMultiLevelInfoInsertLog) SetEventId(time.Time) {
}

func (log *PlayerMultiLevelInfoInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(107)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.BuddyRoleId)
	file.WriteInt8(log.GoRow.BuddyRow)
	file.WriteInt8(log.GoRow.TacticalGrid)
	file.WriteInt8(log.GoRow.DailyNum)
	file.WriteInt64(log.GoRow.BattleTime)
	file.WriteInt32(log.GoRow.Lock)
}

func (log *PlayerMultiLevelInfoInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMultiLevelInfo.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.BuddyRoleId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.BuddyRow)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.TacticalGrid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DailyNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BattleTime)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Lock)))
	return stmt.Execute()
}

func (log *PlayerMultiLevelInfoInsertLog) CommitToTLog() {
}

func (log *PlayerMultiLevelInfoInsertLog) CommitToXdLog() {
}

func (log *PlayerMultiLevelInfoInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMultiLevelInfoDeleteLog struct {
	db *Database
	Old *C.PlayerMultiLevelInfo
	GoOld *PlayerMultiLevelInfo
}

func (db *Database) newPlayerMultiLevelInfoDeleteLog(old *C.PlayerMultiLevelInfo, goOld *PlayerMultiLevelInfo) *PlayerMultiLevelInfoDeleteLog {
	return &PlayerMultiLevelInfoDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerMultiLevelInfoDeleteLog) Free() {
	C.Free_PlayerMultiLevelInfo(log.Old)
}

func (log *PlayerMultiLevelInfoDeleteLog) InvokeHook() {
	if g_Hooks.PlayerMultiLevelInfo != nil {
		g_Hooks.PlayerMultiLevelInfo.Delete(&PlayerMultiLevelInfoRow{row: log.Old})
	}
}

func (log *PlayerMultiLevelInfoDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerMultiLevelInfoDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(107)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.BuddyRoleId)
	file.WriteInt8(log.GoOld.BuddyRow)
	file.WriteInt8(log.GoOld.TacticalGrid)
	file.WriteInt8(log.GoOld.DailyNum)
	file.WriteInt64(log.GoOld.BattleTime)
	file.WriteInt32(log.GoOld.Lock)
}

func (log *PlayerMultiLevelInfoDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMultiLevelInfo.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerMultiLevelInfoDeleteLog) CommitToTLog() {
}

func (log *PlayerMultiLevelInfoDeleteLog) CommitToXdLog() {
}

func (log *PlayerMultiLevelInfoDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerMultiLevelInfoUpdateLog struct {
	db *Database 
	Row *C.PlayerMultiLevelInfo
	Old *C.PlayerMultiLevelInfo
	GoNew *PlayerMultiLevelInfo
	GoOld *PlayerMultiLevelInfo
}

func (db *Database) newPlayerMultiLevelInfoUpdateLog(row, old *C.PlayerMultiLevelInfo, goNew, goOld *PlayerMultiLevelInfo) *PlayerMultiLevelInfoUpdateLog {
	return &PlayerMultiLevelInfoUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerMultiLevelInfoUpdateLog) Free() {
	C.Free_PlayerMultiLevelInfo(log.Old)
}

func (log *PlayerMultiLevelInfoUpdateLog) InvokeHook() {
	if g_Hooks.PlayerMultiLevelInfo != nil {
		g_Hooks.PlayerMultiLevelInfo.Update(&PlayerMultiLevelInfoRow{row: log.Row}, &PlayerMultiLevelInfoRow{row: log.Old})
	}
}

func (log *PlayerMultiLevelInfoUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerMultiLevelInfoUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(107)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.BuddyRoleId)
	file.WriteInt8(log.GoNew.BuddyRow)
	file.WriteInt8(log.GoNew.TacticalGrid)
	file.WriteInt8(log.GoNew.DailyNum)
	file.WriteInt64(log.GoNew.BattleTime)
	file.WriteInt32(log.GoNew.Lock)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.BuddyRoleId)
	file.WriteInt8(log.GoOld.BuddyRow)
	file.WriteInt8(log.GoOld.TacticalGrid)
	file.WriteInt8(log.GoOld.DailyNum)
	file.WriteInt64(log.GoOld.BattleTime)
	file.WriteInt32(log.GoOld.Lock)
}

func (log *PlayerMultiLevelInfoUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerMultiLevelInfo.Update
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.BuddyRoleId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.BuddyRow)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.TacticalGrid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DailyNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BattleTime)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Lock)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerMultiLevelInfoUpdateLog) CommitToTLog() {
}

func (log *PlayerMultiLevelInfoUpdateLog) CommitToXdLog() {
}

func (log *PlayerMultiLevelInfoUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerMultiLevelInfoInsertLog) Rollback() {
	if log.db.tables.PlayerMultiLevelInfo != log.Row {
		panic("insert rollback check failed 'PlayerMultiLevelInfo'")
	}
	log.db.tables.PlayerMultiLevelInfo = nil
	C.Free_PlayerMultiLevelInfo(log.Row)
}

func (log *PlayerMultiLevelInfoDeleteLog) Rollback() {
	if log.db.tables.PlayerMultiLevelInfo != nil {
		panic("delete rollback check failed 'player_multi_level_info'")
	}
	log.db.tables.PlayerMultiLevelInfo = log.Old
}

func (log *PlayerMultiLevelInfoUpdateLog) Rollback() {
	if log.db.tables.PlayerMultiLevelInfo != log.Row {
		panic("update rollback check failed 'player_multi_level_info'")
	}
	log.db.tables.PlayerMultiLevelInfo = log.Old
	C.Free_PlayerMultiLevelInfo(log.Row)
}

/* ========== player_new_year_consume_record ========== */

type PlayerNewYearConsumeRecordInsertLog struct {
	db *Database 
	Row *C.PlayerNewYearConsumeRecord
	GoRow *PlayerNewYearConsumeRecord
}

func (db *Database) newPlayerNewYearConsumeRecordInsertLog(row *C.PlayerNewYearConsumeRecord, goRow *PlayerNewYearConsumeRecord) *PlayerNewYearConsumeRecordInsertLog {
	return &PlayerNewYearConsumeRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerNewYearConsumeRecordInsertLog) Free() {
}

func (log *PlayerNewYearConsumeRecordInsertLog) InvokeHook() {
	if g_Hooks.PlayerNewYearConsumeRecord != nil {
		g_Hooks.PlayerNewYearConsumeRecord.Insert(&PlayerNewYearConsumeRecordRow{row: log.Row})
	}
}

func (log *PlayerNewYearConsumeRecordInsertLog) SetEventId(time.Time) {
}

func (log *PlayerNewYearConsumeRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(108)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteString(log.GoRow.ConsumeRecord)
}

func (log *PlayerNewYearConsumeRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerNewYearConsumeRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.ConsumeRecord), int(log.Row.ConsumeRecord_len))
	return stmt.Execute()
}

func (log *PlayerNewYearConsumeRecordInsertLog) CommitToTLog() {
}

func (log *PlayerNewYearConsumeRecordInsertLog) CommitToXdLog() {
}

func (log *PlayerNewYearConsumeRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerNewYearConsumeRecordDeleteLog struct {
	db *Database
	Old *C.PlayerNewYearConsumeRecord
	GoOld *PlayerNewYearConsumeRecord
}

func (db *Database) newPlayerNewYearConsumeRecordDeleteLog(old *C.PlayerNewYearConsumeRecord, goOld *PlayerNewYearConsumeRecord) *PlayerNewYearConsumeRecordDeleteLog {
	return &PlayerNewYearConsumeRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerNewYearConsumeRecordDeleteLog) Free() {
	C.Free_PlayerNewYearConsumeRecord(log.Old)
}

func (log *PlayerNewYearConsumeRecordDeleteLog) InvokeHook() {
	if g_Hooks.PlayerNewYearConsumeRecord != nil {
		g_Hooks.PlayerNewYearConsumeRecord.Delete(&PlayerNewYearConsumeRecordRow{row: log.Old})
	}
}

func (log *PlayerNewYearConsumeRecordDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerNewYearConsumeRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(108)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteString(log.GoOld.ConsumeRecord)
}

func (log *PlayerNewYearConsumeRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerNewYearConsumeRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerNewYearConsumeRecordDeleteLog) CommitToTLog() {
}

func (log *PlayerNewYearConsumeRecordDeleteLog) CommitToXdLog() {
}

func (log *PlayerNewYearConsumeRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerNewYearConsumeRecordUpdateLog struct {
	db *Database 
	Row *C.PlayerNewYearConsumeRecord
	Old *C.PlayerNewYearConsumeRecord
	GoNew *PlayerNewYearConsumeRecord
	GoOld *PlayerNewYearConsumeRecord
}

func (db *Database) newPlayerNewYearConsumeRecordUpdateLog(row, old *C.PlayerNewYearConsumeRecord, goNew, goOld *PlayerNewYearConsumeRecord) *PlayerNewYearConsumeRecordUpdateLog {
	return &PlayerNewYearConsumeRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerNewYearConsumeRecordUpdateLog) Free() {
	C.Free_PlayerNewYearConsumeRecord(log.Old)
}

func (log *PlayerNewYearConsumeRecordUpdateLog) InvokeHook() {
	if g_Hooks.PlayerNewYearConsumeRecord != nil {
		g_Hooks.PlayerNewYearConsumeRecord.Update(&PlayerNewYearConsumeRecordRow{row: log.Row}, &PlayerNewYearConsumeRecordRow{row: log.Old})
	}
}

func (log *PlayerNewYearConsumeRecordUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerNewYearConsumeRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(108)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteString(log.GoNew.ConsumeRecord)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteString(log.GoOld.ConsumeRecord)
}

func (log *PlayerNewYearConsumeRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerNewYearConsumeRecord.Update
	stmt.BindVarchar(unsafe.Pointer(log.Row.ConsumeRecord), int(log.Row.ConsumeRecord_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerNewYearConsumeRecordUpdateLog) CommitToTLog() {
}

func (log *PlayerNewYearConsumeRecordUpdateLog) CommitToXdLog() {
}

func (log *PlayerNewYearConsumeRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerNewYearConsumeRecordInsertLog) Rollback() {
	if log.db.tables.PlayerNewYearConsumeRecord != log.Row {
		panic("insert rollback check failed 'PlayerNewYearConsumeRecord'")
	}
	log.db.tables.PlayerNewYearConsumeRecord = nil
	C.Free_PlayerNewYearConsumeRecord(log.Row)
}

func (log *PlayerNewYearConsumeRecordDeleteLog) Rollback() {
	if log.db.tables.PlayerNewYearConsumeRecord != nil {
		panic("delete rollback check failed 'player_new_year_consume_record'")
	}
	log.db.tables.PlayerNewYearConsumeRecord = log.Old
}

func (log *PlayerNewYearConsumeRecordUpdateLog) Rollback() {
	if log.db.tables.PlayerNewYearConsumeRecord != log.Row {
		panic("update rollback check failed 'player_new_year_consume_record'")
	}
	log.db.tables.PlayerNewYearConsumeRecord = log.Old
	C.Free_PlayerNewYearConsumeRecord(log.Row)
}

/* ========== player_npc_talk_record ========== */

type PlayerNpcTalkRecordInsertLog struct {
	db *Database 
	Row *C.PlayerNpcTalkRecord
	GoRow *PlayerNpcTalkRecord
}

func (db *Database) newPlayerNpcTalkRecordInsertLog(row *C.PlayerNpcTalkRecord, goRow *PlayerNpcTalkRecord) *PlayerNpcTalkRecordInsertLog {
	return &PlayerNpcTalkRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerNpcTalkRecordInsertLog) Free() {
}

func (log *PlayerNpcTalkRecordInsertLog) InvokeHook() {
	if g_Hooks.PlayerNpcTalkRecord != nil {
		g_Hooks.PlayerNpcTalkRecord.Insert(&PlayerNpcTalkRecordRow{row: log.Row})
	}
}

func (log *PlayerNpcTalkRecordInsertLog) SetEventId(time.Time) {
}

func (log *PlayerNpcTalkRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(109)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.NpcId)
	file.WriteInt16(log.GoRow.TownId)
	file.WriteInt16(log.GoRow.QuestId)
}

func (log *PlayerNpcTalkRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerNpcTalkRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.NpcId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TownId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.QuestId)))
	return stmt.Execute()
}

func (log *PlayerNpcTalkRecordInsertLog) CommitToTLog() {
}

func (log *PlayerNpcTalkRecordInsertLog) CommitToXdLog() {
}

func (log *PlayerNpcTalkRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerNpcTalkRecordDeleteLog struct {
	db *Database
	Old *C.PlayerNpcTalkRecord
	GoOld *PlayerNpcTalkRecord
}

func (db *Database) newPlayerNpcTalkRecordDeleteLog(old *C.PlayerNpcTalkRecord, goOld *PlayerNpcTalkRecord) *PlayerNpcTalkRecordDeleteLog {
	return &PlayerNpcTalkRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerNpcTalkRecordDeleteLog) Free() {
	C.Free_PlayerNpcTalkRecord(log.Old)
}

func (log *PlayerNpcTalkRecordDeleteLog) InvokeHook() {
	if g_Hooks.PlayerNpcTalkRecord != nil {
		g_Hooks.PlayerNpcTalkRecord.Delete(&PlayerNpcTalkRecordRow{row: log.Old})
	}
}

func (log *PlayerNpcTalkRecordDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerNpcTalkRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(109)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.NpcId)
	file.WriteInt16(log.GoOld.TownId)
	file.WriteInt16(log.GoOld.QuestId)
}

func (log *PlayerNpcTalkRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerNpcTalkRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerNpcTalkRecordDeleteLog) CommitToTLog() {
}

func (log *PlayerNpcTalkRecordDeleteLog) CommitToXdLog() {
}

func (log *PlayerNpcTalkRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerNpcTalkRecordUpdateLog struct {
	db *Database 
	Row *C.PlayerNpcTalkRecord
	Old *C.PlayerNpcTalkRecord
	GoNew *PlayerNpcTalkRecord
	GoOld *PlayerNpcTalkRecord
}

func (db *Database) newPlayerNpcTalkRecordUpdateLog(row, old *C.PlayerNpcTalkRecord, goNew, goOld *PlayerNpcTalkRecord) *PlayerNpcTalkRecordUpdateLog {
	return &PlayerNpcTalkRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerNpcTalkRecordUpdateLog) Free() {
	C.Free_PlayerNpcTalkRecord(log.Old)
}

func (log *PlayerNpcTalkRecordUpdateLog) InvokeHook() {
	if g_Hooks.PlayerNpcTalkRecord != nil {
		g_Hooks.PlayerNpcTalkRecord.Update(&PlayerNpcTalkRecordRow{row: log.Row}, &PlayerNpcTalkRecordRow{row: log.Old})
	}
}

func (log *PlayerNpcTalkRecordUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerNpcTalkRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(109)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.NpcId)
	file.WriteInt16(log.GoNew.TownId)
	file.WriteInt16(log.GoNew.QuestId)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.NpcId)
	file.WriteInt16(log.GoOld.TownId)
	file.WriteInt16(log.GoOld.QuestId)
}

func (log *PlayerNpcTalkRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerNpcTalkRecord.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.NpcId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TownId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.QuestId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerNpcTalkRecordUpdateLog) CommitToTLog() {
}

func (log *PlayerNpcTalkRecordUpdateLog) CommitToXdLog() {
}

func (log *PlayerNpcTalkRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerNpcTalkRecordInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerNpcTalkRecord
	for crow := log.db.tables.PlayerNpcTalkRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerNpcTalkRecord'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerNpcTalkRecord = log.db.tables.PlayerNpcTalkRecord.next
	}
	C.Free_PlayerNpcTalkRecord(log.Row)
}

func (log *PlayerNpcTalkRecordDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerNpcTalkRecord
	for crow := log.db.tables.PlayerNpcTalkRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_npc_talk_record'")
	}
	log.Old.next = log.db.tables.PlayerNpcTalkRecord
	log.db.tables.PlayerNpcTalkRecord = log.Old
}

func (log *PlayerNpcTalkRecordUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerNpcTalkRecord
	for crow := log.db.tables.PlayerNpcTalkRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_npc_talk_record'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerNpcTalkRecord = log.Old
	}
	C.Free_PlayerNpcTalkRecord(log.Row)
}

/* ========== player_opened_town_treasure ========== */

type PlayerOpenedTownTreasureInsertLog struct {
	db *Database 
	Row *C.PlayerOpenedTownTreasure
	GoRow *PlayerOpenedTownTreasure
}

func (db *Database) newPlayerOpenedTownTreasureInsertLog(row *C.PlayerOpenedTownTreasure, goRow *PlayerOpenedTownTreasure) *PlayerOpenedTownTreasureInsertLog {
	return &PlayerOpenedTownTreasureInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerOpenedTownTreasureInsertLog) Free() {
}

func (log *PlayerOpenedTownTreasureInsertLog) InvokeHook() {
	if g_Hooks.PlayerOpenedTownTreasure != nil {
		g_Hooks.PlayerOpenedTownTreasure.Insert(&PlayerOpenedTownTreasureRow{row: log.Row})
	}
}

func (log *PlayerOpenedTownTreasureInsertLog) SetEventId(time.Time) {
}

func (log *PlayerOpenedTownTreasureInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(110)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.TownId)
}

func (log *PlayerOpenedTownTreasureInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerOpenedTownTreasure.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TownId)))
	return stmt.Execute()
}

func (log *PlayerOpenedTownTreasureInsertLog) CommitToTLog() {
}

func (log *PlayerOpenedTownTreasureInsertLog) CommitToXdLog() {
}

func (log *PlayerOpenedTownTreasureInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerOpenedTownTreasureDeleteLog struct {
	db *Database
	Old *C.PlayerOpenedTownTreasure
	GoOld *PlayerOpenedTownTreasure
}

func (db *Database) newPlayerOpenedTownTreasureDeleteLog(old *C.PlayerOpenedTownTreasure, goOld *PlayerOpenedTownTreasure) *PlayerOpenedTownTreasureDeleteLog {
	return &PlayerOpenedTownTreasureDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerOpenedTownTreasureDeleteLog) Free() {
	C.Free_PlayerOpenedTownTreasure(log.Old)
}

func (log *PlayerOpenedTownTreasureDeleteLog) InvokeHook() {
	if g_Hooks.PlayerOpenedTownTreasure != nil {
		g_Hooks.PlayerOpenedTownTreasure.Delete(&PlayerOpenedTownTreasureRow{row: log.Old})
	}
}

func (log *PlayerOpenedTownTreasureDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerOpenedTownTreasureDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(110)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.TownId)
}

func (log *PlayerOpenedTownTreasureDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerOpenedTownTreasure.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerOpenedTownTreasureDeleteLog) CommitToTLog() {
}

func (log *PlayerOpenedTownTreasureDeleteLog) CommitToXdLog() {
}

func (log *PlayerOpenedTownTreasureDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerOpenedTownTreasureUpdateLog struct {
	db *Database 
	Row *C.PlayerOpenedTownTreasure
	Old *C.PlayerOpenedTownTreasure
	GoNew *PlayerOpenedTownTreasure
	GoOld *PlayerOpenedTownTreasure
}

func (db *Database) newPlayerOpenedTownTreasureUpdateLog(row, old *C.PlayerOpenedTownTreasure, goNew, goOld *PlayerOpenedTownTreasure) *PlayerOpenedTownTreasureUpdateLog {
	return &PlayerOpenedTownTreasureUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerOpenedTownTreasureUpdateLog) Free() {
	C.Free_PlayerOpenedTownTreasure(log.Old)
}

func (log *PlayerOpenedTownTreasureUpdateLog) InvokeHook() {
	if g_Hooks.PlayerOpenedTownTreasure != nil {
		g_Hooks.PlayerOpenedTownTreasure.Update(&PlayerOpenedTownTreasureRow{row: log.Row}, &PlayerOpenedTownTreasureRow{row: log.Old})
	}
}

func (log *PlayerOpenedTownTreasureUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerOpenedTownTreasureUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(110)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.TownId)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.TownId)
}

func (log *PlayerOpenedTownTreasureUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerOpenedTownTreasure.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TownId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerOpenedTownTreasureUpdateLog) CommitToTLog() {
}

func (log *PlayerOpenedTownTreasureUpdateLog) CommitToXdLog() {
}

func (log *PlayerOpenedTownTreasureUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerOpenedTownTreasureInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerOpenedTownTreasure
	for crow := log.db.tables.PlayerOpenedTownTreasure; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerOpenedTownTreasure'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerOpenedTownTreasure = log.db.tables.PlayerOpenedTownTreasure.next
	}
	C.Free_PlayerOpenedTownTreasure(log.Row)
}

func (log *PlayerOpenedTownTreasureDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerOpenedTownTreasure
	for crow := log.db.tables.PlayerOpenedTownTreasure; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_opened_town_treasure'")
	}
	log.Old.next = log.db.tables.PlayerOpenedTownTreasure
	log.db.tables.PlayerOpenedTownTreasure = log.Old
}

func (log *PlayerOpenedTownTreasureUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerOpenedTownTreasure
	for crow := log.db.tables.PlayerOpenedTownTreasure; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_opened_town_treasure'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerOpenedTownTreasure = log.Old
	}
	C.Free_PlayerOpenedTownTreasure(log.Row)
}

/* ========== player_physical ========== */

type PlayerPhysicalInsertLog struct {
	db *Database 
	Row *C.PlayerPhysical
	GoRow *PlayerPhysical
}

func (db *Database) newPlayerPhysicalInsertLog(row *C.PlayerPhysical, goRow *PlayerPhysical) *PlayerPhysicalInsertLog {
	return &PlayerPhysicalInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerPhysicalInsertLog) Free() {
}

func (log *PlayerPhysicalInsertLog) InvokeHook() {
	if g_Hooks.PlayerPhysical != nil {
		g_Hooks.PlayerPhysical.Insert(&PlayerPhysicalRow{row: log.Row})
	}
}

func (log *PlayerPhysicalInsertLog) SetEventId(time.Time) {
}

func (log *PlayerPhysicalInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(111)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.Value)
	file.WriteInt64(log.GoRow.UpdateTime)
	file.WriteInt64(log.GoRow.BuyCount)
	file.WriteInt64(log.GoRow.BuyUpdateTime)
	file.WriteInt16(log.GoRow.DailyCount)
}

func (log *PlayerPhysicalInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerPhysical.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Value)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.UpdateTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BuyCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BuyUpdateTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyCount)))
	return stmt.Execute()
}

func (log *PlayerPhysicalInsertLog) CommitToTLog() {
}

func (log *PlayerPhysicalInsertLog) CommitToXdLog() {
}

func (log *PlayerPhysicalInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerPhysicalDeleteLog struct {
	db *Database
	Old *C.PlayerPhysical
	GoOld *PlayerPhysical
}

func (db *Database) newPlayerPhysicalDeleteLog(old *C.PlayerPhysical, goOld *PlayerPhysical) *PlayerPhysicalDeleteLog {
	return &PlayerPhysicalDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerPhysicalDeleteLog) Free() {
	C.Free_PlayerPhysical(log.Old)
}

func (log *PlayerPhysicalDeleteLog) InvokeHook() {
	if g_Hooks.PlayerPhysical != nil {
		g_Hooks.PlayerPhysical.Delete(&PlayerPhysicalRow{row: log.Old})
	}
}

func (log *PlayerPhysicalDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerPhysicalDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(111)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.Value)
	file.WriteInt64(log.GoOld.UpdateTime)
	file.WriteInt64(log.GoOld.BuyCount)
	file.WriteInt64(log.GoOld.BuyUpdateTime)
	file.WriteInt16(log.GoOld.DailyCount)
}

func (log *PlayerPhysicalDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerPhysical.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerPhysicalDeleteLog) CommitToTLog() {
}

func (log *PlayerPhysicalDeleteLog) CommitToXdLog() {
}

func (log *PlayerPhysicalDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerPhysicalUpdateLog struct {
	db *Database 
	Row *C.PlayerPhysical
	Old *C.PlayerPhysical
	GoNew *PlayerPhysical
	GoOld *PlayerPhysical
}

func (db *Database) newPlayerPhysicalUpdateLog(row, old *C.PlayerPhysical, goNew, goOld *PlayerPhysical) *PlayerPhysicalUpdateLog {
	return &PlayerPhysicalUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerPhysicalUpdateLog) Free() {
	C.Free_PlayerPhysical(log.Old)
}

func (log *PlayerPhysicalUpdateLog) InvokeHook() {
	if g_Hooks.PlayerPhysical != nil {
		g_Hooks.PlayerPhysical.Update(&PlayerPhysicalRow{row: log.Row}, &PlayerPhysicalRow{row: log.Old})
	}
}

func (log *PlayerPhysicalUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerPhysicalUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(111)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.Value)
	file.WriteInt64(log.GoNew.UpdateTime)
	file.WriteInt64(log.GoNew.BuyCount)
	file.WriteInt64(log.GoNew.BuyUpdateTime)
	file.WriteInt16(log.GoNew.DailyCount)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.Value)
	file.WriteInt64(log.GoOld.UpdateTime)
	file.WriteInt64(log.GoOld.BuyCount)
	file.WriteInt64(log.GoOld.BuyUpdateTime)
	file.WriteInt16(log.GoOld.DailyCount)
}

func (log *PlayerPhysicalUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerPhysical.Update
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Value)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.UpdateTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BuyCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BuyUpdateTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DailyCount)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerPhysicalUpdateLog) CommitToTLog() {
}

func (log *PlayerPhysicalUpdateLog) CommitToXdLog() {
}

func (log *PlayerPhysicalUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerPhysicalInsertLog) Rollback() {
	if log.db.tables.PlayerPhysical != log.Row {
		panic("insert rollback check failed 'PlayerPhysical'")
	}
	log.db.tables.PlayerPhysical = nil
	C.Free_PlayerPhysical(log.Row)
}

func (log *PlayerPhysicalDeleteLog) Rollback() {
	if log.db.tables.PlayerPhysical != nil {
		panic("delete rollback check failed 'player_physical'")
	}
	log.db.tables.PlayerPhysical = log.Old
}

func (log *PlayerPhysicalUpdateLog) Rollback() {
	if log.db.tables.PlayerPhysical != log.Row {
		panic("update rollback check failed 'player_physical'")
	}
	log.db.tables.PlayerPhysical = log.Old
	C.Free_PlayerPhysical(log.Row)
}

/* ========== player_purchase_record ========== */

type PlayerPurchaseRecordInsertLog struct {
	db *Database 
	Row *C.PlayerPurchaseRecord
	GoRow *PlayerPurchaseRecord
}

func (db *Database) newPlayerPurchaseRecordInsertLog(row *C.PlayerPurchaseRecord, goRow *PlayerPurchaseRecord) *PlayerPurchaseRecordInsertLog {
	return &PlayerPurchaseRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerPurchaseRecordInsertLog) Free() {
}

func (log *PlayerPurchaseRecordInsertLog) InvokeHook() {
	if g_Hooks.PlayerPurchaseRecord != nil {
		g_Hooks.PlayerPurchaseRecord.Insert(&PlayerPurchaseRecordRow{row: log.Row})
	}
}

func (log *PlayerPurchaseRecordInsertLog) SetEventId(time.Time) {
}

func (log *PlayerPurchaseRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(112)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.ItemId)
	file.WriteInt16(log.GoRow.Num)
	file.WriteInt64(log.GoRow.Timestamp)
}

func (log *PlayerPurchaseRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerPurchaseRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Num)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	return stmt.Execute()
}

func (log *PlayerPurchaseRecordInsertLog) CommitToTLog() {
}

func (log *PlayerPurchaseRecordInsertLog) CommitToXdLog() {
}

func (log *PlayerPurchaseRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerPurchaseRecordDeleteLog struct {
	db *Database
	Old *C.PlayerPurchaseRecord
	GoOld *PlayerPurchaseRecord
}

func (db *Database) newPlayerPurchaseRecordDeleteLog(old *C.PlayerPurchaseRecord, goOld *PlayerPurchaseRecord) *PlayerPurchaseRecordDeleteLog {
	return &PlayerPurchaseRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerPurchaseRecordDeleteLog) Free() {
	C.Free_PlayerPurchaseRecord(log.Old)
}

func (log *PlayerPurchaseRecordDeleteLog) InvokeHook() {
	if g_Hooks.PlayerPurchaseRecord != nil {
		g_Hooks.PlayerPurchaseRecord.Delete(&PlayerPurchaseRecordRow{row: log.Old})
	}
}

func (log *PlayerPurchaseRecordDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerPurchaseRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(112)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt16(log.GoOld.Num)
	file.WriteInt64(log.GoOld.Timestamp)
}

func (log *PlayerPurchaseRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerPurchaseRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerPurchaseRecordDeleteLog) CommitToTLog() {
}

func (log *PlayerPurchaseRecordDeleteLog) CommitToXdLog() {
}

func (log *PlayerPurchaseRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerPurchaseRecordUpdateLog struct {
	db *Database 
	Row *C.PlayerPurchaseRecord
	Old *C.PlayerPurchaseRecord
	GoNew *PlayerPurchaseRecord
	GoOld *PlayerPurchaseRecord
}

func (db *Database) newPlayerPurchaseRecordUpdateLog(row, old *C.PlayerPurchaseRecord, goNew, goOld *PlayerPurchaseRecord) *PlayerPurchaseRecordUpdateLog {
	return &PlayerPurchaseRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerPurchaseRecordUpdateLog) Free() {
	C.Free_PlayerPurchaseRecord(log.Old)
}

func (log *PlayerPurchaseRecordUpdateLog) InvokeHook() {
	if g_Hooks.PlayerPurchaseRecord != nil {
		g_Hooks.PlayerPurchaseRecord.Update(&PlayerPurchaseRecordRow{row: log.Row}, &PlayerPurchaseRecordRow{row: log.Old})
	}
}

func (log *PlayerPurchaseRecordUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerPurchaseRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(112)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.ItemId)
	file.WriteInt16(log.GoNew.Num)
	file.WriteInt64(log.GoNew.Timestamp)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt16(log.GoOld.Num)
	file.WriteInt64(log.GoOld.Timestamp)
}

func (log *PlayerPurchaseRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerPurchaseRecord.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Num)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerPurchaseRecordUpdateLog) CommitToTLog() {
}

func (log *PlayerPurchaseRecordUpdateLog) CommitToXdLog() {
}

func (log *PlayerPurchaseRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerPurchaseRecordInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerPurchaseRecord
	for crow := log.db.tables.PlayerPurchaseRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerPurchaseRecord'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerPurchaseRecord = log.db.tables.PlayerPurchaseRecord.next
	}
	C.Free_PlayerPurchaseRecord(log.Row)
}

func (log *PlayerPurchaseRecordDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerPurchaseRecord
	for crow := log.db.tables.PlayerPurchaseRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_purchase_record'")
	}
	log.Old.next = log.db.tables.PlayerPurchaseRecord
	log.db.tables.PlayerPurchaseRecord = log.Old
}

func (log *PlayerPurchaseRecordUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerPurchaseRecord
	for crow := log.db.tables.PlayerPurchaseRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_purchase_record'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerPurchaseRecord = log.Old
	}
	C.Free_PlayerPurchaseRecord(log.Row)
}

/* ========== player_push_notify_switch ========== */

type PlayerPushNotifySwitchInsertLog struct {
	db *Database 
	Row *C.PlayerPushNotifySwitch
	GoRow *PlayerPushNotifySwitch
}

func (db *Database) newPlayerPushNotifySwitchInsertLog(row *C.PlayerPushNotifySwitch, goRow *PlayerPushNotifySwitch) *PlayerPushNotifySwitchInsertLog {
	return &PlayerPushNotifySwitchInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerPushNotifySwitchInsertLog) Free() {
}

func (log *PlayerPushNotifySwitchInsertLog) InvokeHook() {
	if g_Hooks.PlayerPushNotifySwitch != nil {
		g_Hooks.PlayerPushNotifySwitch.Insert(&PlayerPushNotifySwitchRow{row: log.Row})
	}
}

func (log *PlayerPushNotifySwitchInsertLog) SetEventId(time.Time) {
}

func (log *PlayerPushNotifySwitchInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(113)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.Options)
}

func (log *PlayerPushNotifySwitchInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerPushNotifySwitch.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Options)))
	return stmt.Execute()
}

func (log *PlayerPushNotifySwitchInsertLog) CommitToTLog() {
}

func (log *PlayerPushNotifySwitchInsertLog) CommitToXdLog() {
}

func (log *PlayerPushNotifySwitchInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerPushNotifySwitchDeleteLog struct {
	db *Database
	Old *C.PlayerPushNotifySwitch
	GoOld *PlayerPushNotifySwitch
}

func (db *Database) newPlayerPushNotifySwitchDeleteLog(old *C.PlayerPushNotifySwitch, goOld *PlayerPushNotifySwitch) *PlayerPushNotifySwitchDeleteLog {
	return &PlayerPushNotifySwitchDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerPushNotifySwitchDeleteLog) Free() {
	C.Free_PlayerPushNotifySwitch(log.Old)
}

func (log *PlayerPushNotifySwitchDeleteLog) InvokeHook() {
	if g_Hooks.PlayerPushNotifySwitch != nil {
		g_Hooks.PlayerPushNotifySwitch.Delete(&PlayerPushNotifySwitchRow{row: log.Old})
	}
}

func (log *PlayerPushNotifySwitchDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerPushNotifySwitchDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(113)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.Options)
}

func (log *PlayerPushNotifySwitchDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerPushNotifySwitch.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerPushNotifySwitchDeleteLog) CommitToTLog() {
}

func (log *PlayerPushNotifySwitchDeleteLog) CommitToXdLog() {
}

func (log *PlayerPushNotifySwitchDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerPushNotifySwitchUpdateLog struct {
	db *Database 
	Row *C.PlayerPushNotifySwitch
	Old *C.PlayerPushNotifySwitch
	GoNew *PlayerPushNotifySwitch
	GoOld *PlayerPushNotifySwitch
}

func (db *Database) newPlayerPushNotifySwitchUpdateLog(row, old *C.PlayerPushNotifySwitch, goNew, goOld *PlayerPushNotifySwitch) *PlayerPushNotifySwitchUpdateLog {
	return &PlayerPushNotifySwitchUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerPushNotifySwitchUpdateLog) Free() {
	C.Free_PlayerPushNotifySwitch(log.Old)
}

func (log *PlayerPushNotifySwitchUpdateLog) InvokeHook() {
	if g_Hooks.PlayerPushNotifySwitch != nil {
		g_Hooks.PlayerPushNotifySwitch.Update(&PlayerPushNotifySwitchRow{row: log.Row}, &PlayerPushNotifySwitchRow{row: log.Old})
	}
}

func (log *PlayerPushNotifySwitchUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerPushNotifySwitchUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(113)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.Options)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.Options)
}

func (log *PlayerPushNotifySwitchUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerPushNotifySwitch.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Options)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerPushNotifySwitchUpdateLog) CommitToTLog() {
}

func (log *PlayerPushNotifySwitchUpdateLog) CommitToXdLog() {
}

func (log *PlayerPushNotifySwitchUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerPushNotifySwitchInsertLog) Rollback() {
	if log.db.tables.PlayerPushNotifySwitch != log.Row {
		panic("insert rollback check failed 'PlayerPushNotifySwitch'")
	}
	log.db.tables.PlayerPushNotifySwitch = nil
	C.Free_PlayerPushNotifySwitch(log.Row)
}

func (log *PlayerPushNotifySwitchDeleteLog) Rollback() {
	if log.db.tables.PlayerPushNotifySwitch != nil {
		panic("delete rollback check failed 'player_push_notify_switch'")
	}
	log.db.tables.PlayerPushNotifySwitch = log.Old
}

func (log *PlayerPushNotifySwitchUpdateLog) Rollback() {
	if log.db.tables.PlayerPushNotifySwitch != log.Row {
		panic("update rollback check failed 'player_push_notify_switch'")
	}
	log.db.tables.PlayerPushNotifySwitch = log.Old
	C.Free_PlayerPushNotifySwitch(log.Row)
}

/* ========== player_pve_state ========== */

type PlayerPveStateInsertLog struct {
	db *Database 
	Row *C.PlayerPveState
	GoRow *PlayerPveState
}

func (db *Database) newPlayerPveStateInsertLog(row *C.PlayerPveState, goRow *PlayerPveState) *PlayerPveStateInsertLog {
	return &PlayerPveStateInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerPveStateInsertLog) Free() {
}

func (log *PlayerPveStateInsertLog) InvokeHook() {
	if g_Hooks.PlayerPveState != nil {
		g_Hooks.PlayerPveState.Insert(&PlayerPveStateRow{row: log.Row})
	}
}

func (log *PlayerPveStateInsertLog) SetEventId(time.Time) {
}

func (log *PlayerPveStateInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(114)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.MaxPassedFloor)
	file.WriteInt16(log.GoRow.MaxAwardedFloor)
	file.WriteInt16(log.GoRow.UnpassedFloorEnemyNum)
	file.WriteInt64(log.GoRow.EnterTime)
	file.WriteInt8(log.GoRow.DailyNum)
}

func (log *PlayerPveStateInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerPveState.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.MaxPassedFloor)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.MaxAwardedFloor)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.UnpassedFloorEnemyNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.EnterTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DailyNum)))
	return stmt.Execute()
}

func (log *PlayerPveStateInsertLog) CommitToTLog() {
}

func (log *PlayerPveStateInsertLog) CommitToXdLog() {
}

func (log *PlayerPveStateInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerPveStateDeleteLog struct {
	db *Database
	Old *C.PlayerPveState
	GoOld *PlayerPveState
}

func (db *Database) newPlayerPveStateDeleteLog(old *C.PlayerPveState, goOld *PlayerPveState) *PlayerPveStateDeleteLog {
	return &PlayerPveStateDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerPveStateDeleteLog) Free() {
	C.Free_PlayerPveState(log.Old)
}

func (log *PlayerPveStateDeleteLog) InvokeHook() {
	if g_Hooks.PlayerPveState != nil {
		g_Hooks.PlayerPveState.Delete(&PlayerPveStateRow{row: log.Old})
	}
}

func (log *PlayerPveStateDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerPveStateDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(114)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.MaxPassedFloor)
	file.WriteInt16(log.GoOld.MaxAwardedFloor)
	file.WriteInt16(log.GoOld.UnpassedFloorEnemyNum)
	file.WriteInt64(log.GoOld.EnterTime)
	file.WriteInt8(log.GoOld.DailyNum)
}

func (log *PlayerPveStateDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerPveState.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerPveStateDeleteLog) CommitToTLog() {
}

func (log *PlayerPveStateDeleteLog) CommitToXdLog() {
}

func (log *PlayerPveStateDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerPveStateUpdateLog struct {
	db *Database 
	Row *C.PlayerPveState
	Old *C.PlayerPveState
	GoNew *PlayerPveState
	GoOld *PlayerPveState
}

func (db *Database) newPlayerPveStateUpdateLog(row, old *C.PlayerPveState, goNew, goOld *PlayerPveState) *PlayerPveStateUpdateLog {
	return &PlayerPveStateUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerPveStateUpdateLog) Free() {
	C.Free_PlayerPveState(log.Old)
}

func (log *PlayerPveStateUpdateLog) InvokeHook() {
	if g_Hooks.PlayerPveState != nil {
		g_Hooks.PlayerPveState.Update(&PlayerPveStateRow{row: log.Row}, &PlayerPveStateRow{row: log.Old})
	}
}

func (log *PlayerPveStateUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerPveStateUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(114)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.MaxPassedFloor)
	file.WriteInt16(log.GoNew.MaxAwardedFloor)
	file.WriteInt16(log.GoNew.UnpassedFloorEnemyNum)
	file.WriteInt64(log.GoNew.EnterTime)
	file.WriteInt8(log.GoNew.DailyNum)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.MaxPassedFloor)
	file.WriteInt16(log.GoOld.MaxAwardedFloor)
	file.WriteInt16(log.GoOld.UnpassedFloorEnemyNum)
	file.WriteInt64(log.GoOld.EnterTime)
	file.WriteInt8(log.GoOld.DailyNum)
}

func (log *PlayerPveStateUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerPveState.Update
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.MaxPassedFloor)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.MaxAwardedFloor)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.UnpassedFloorEnemyNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.EnterTime)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.DailyNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerPveStateUpdateLog) CommitToTLog() {
}

func (log *PlayerPveStateUpdateLog) CommitToXdLog() {
}

func (log *PlayerPveStateUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerPveStateInsertLog) Rollback() {
	if log.db.tables.PlayerPveState != log.Row {
		panic("insert rollback check failed 'PlayerPveState'")
	}
	log.db.tables.PlayerPveState = nil
	C.Free_PlayerPveState(log.Row)
}

func (log *PlayerPveStateDeleteLog) Rollback() {
	if log.db.tables.PlayerPveState != nil {
		panic("delete rollback check failed 'player_pve_state'")
	}
	log.db.tables.PlayerPveState = log.Old
}

func (log *PlayerPveStateUpdateLog) Rollback() {
	if log.db.tables.PlayerPveState != log.Row {
		panic("update rollback check failed 'player_pve_state'")
	}
	log.db.tables.PlayerPveState = log.Old
	C.Free_PlayerPveState(log.Row)
}

/* ========== player_qq_vip_gift ========== */

type PlayerQqVipGiftInsertLog struct {
	db *Database 
	Row *C.PlayerQqVipGift
	GoRow *PlayerQqVipGift
}

func (db *Database) newPlayerQqVipGiftInsertLog(row *C.PlayerQqVipGift, goRow *PlayerQqVipGift) *PlayerQqVipGiftInsertLog {
	return &PlayerQqVipGiftInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerQqVipGiftInsertLog) Free() {
}

func (log *PlayerQqVipGiftInsertLog) InvokeHook() {
	if g_Hooks.PlayerQqVipGift != nil {
		g_Hooks.PlayerQqVipGift.Insert(&PlayerQqVipGiftRow{row: log.Row})
	}
}

func (log *PlayerQqVipGiftInsertLog) SetEventId(time.Time) {
}

func (log *PlayerQqVipGiftInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(115)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.Qqvip)
	file.WriteInt16(log.GoRow.Surper)
}

func (log *PlayerQqVipGiftInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerQqVipGift.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Qqvip)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Surper)))
	return stmt.Execute()
}

func (log *PlayerQqVipGiftInsertLog) CommitToTLog() {
}

func (log *PlayerQqVipGiftInsertLog) CommitToXdLog() {
}

func (log *PlayerQqVipGiftInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerQqVipGiftDeleteLog struct {
	db *Database
	Old *C.PlayerQqVipGift
	GoOld *PlayerQqVipGift
}

func (db *Database) newPlayerQqVipGiftDeleteLog(old *C.PlayerQqVipGift, goOld *PlayerQqVipGift) *PlayerQqVipGiftDeleteLog {
	return &PlayerQqVipGiftDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerQqVipGiftDeleteLog) Free() {
	C.Free_PlayerQqVipGift(log.Old)
}

func (log *PlayerQqVipGiftDeleteLog) InvokeHook() {
	if g_Hooks.PlayerQqVipGift != nil {
		g_Hooks.PlayerQqVipGift.Delete(&PlayerQqVipGiftRow{row: log.Old})
	}
}

func (log *PlayerQqVipGiftDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerQqVipGiftDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(115)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.Qqvip)
	file.WriteInt16(log.GoOld.Surper)
}

func (log *PlayerQqVipGiftDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerQqVipGift.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerQqVipGiftDeleteLog) CommitToTLog() {
}

func (log *PlayerQqVipGiftDeleteLog) CommitToXdLog() {
}

func (log *PlayerQqVipGiftDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerQqVipGiftUpdateLog struct {
	db *Database 
	Row *C.PlayerQqVipGift
	Old *C.PlayerQqVipGift
	GoNew *PlayerQqVipGift
	GoOld *PlayerQqVipGift
}

func (db *Database) newPlayerQqVipGiftUpdateLog(row, old *C.PlayerQqVipGift, goNew, goOld *PlayerQqVipGift) *PlayerQqVipGiftUpdateLog {
	return &PlayerQqVipGiftUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerQqVipGiftUpdateLog) Free() {
	C.Free_PlayerQqVipGift(log.Old)
}

func (log *PlayerQqVipGiftUpdateLog) InvokeHook() {
	if g_Hooks.PlayerQqVipGift != nil {
		g_Hooks.PlayerQqVipGift.Update(&PlayerQqVipGiftRow{row: log.Row}, &PlayerQqVipGiftRow{row: log.Old})
	}
}

func (log *PlayerQqVipGiftUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerQqVipGiftUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(115)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.Qqvip)
	file.WriteInt16(log.GoNew.Surper)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.Qqvip)
	file.WriteInt16(log.GoOld.Surper)
}

func (log *PlayerQqVipGiftUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerQqVipGift.Update
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Qqvip)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Surper)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerQqVipGiftUpdateLog) CommitToTLog() {
}

func (log *PlayerQqVipGiftUpdateLog) CommitToXdLog() {
}

func (log *PlayerQqVipGiftUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerQqVipGiftInsertLog) Rollback() {
	if log.db.tables.PlayerQqVipGift != log.Row {
		panic("insert rollback check failed 'PlayerQqVipGift'")
	}
	log.db.tables.PlayerQqVipGift = nil
	C.Free_PlayerQqVipGift(log.Row)
}

func (log *PlayerQqVipGiftDeleteLog) Rollback() {
	if log.db.tables.PlayerQqVipGift != nil {
		panic("delete rollback check failed 'player_qq_vip_gift'")
	}
	log.db.tables.PlayerQqVipGift = log.Old
}

func (log *PlayerQqVipGiftUpdateLog) Rollback() {
	if log.db.tables.PlayerQqVipGift != log.Row {
		panic("update rollback check failed 'player_qq_vip_gift'")
	}
	log.db.tables.PlayerQqVipGift = log.Old
	C.Free_PlayerQqVipGift(log.Row)
}

/* ========== player_quest ========== */

type PlayerQuestInsertLog struct {
	db *Database 
	Row *C.PlayerQuest
	GoRow *PlayerQuest
}

func (db *Database) newPlayerQuestInsertLog(row *C.PlayerQuest, goRow *PlayerQuest) *PlayerQuestInsertLog {
	return &PlayerQuestInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerQuestInsertLog) Free() {
}

func (log *PlayerQuestInsertLog) InvokeHook() {
	if g_Hooks.PlayerQuest != nil {
		g_Hooks.PlayerQuest.Insert(&PlayerQuestRow{row: log.Row})
	}
}

func (log *PlayerQuestInsertLog) SetEventId(time.Time) {
}

func (log *PlayerQuestInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(116)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.QuestId)
	file.WriteInt8(log.GoRow.State)
}

func (log *PlayerQuestInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerQuest.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.QuestId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.State)))
	return stmt.Execute()
}

func (log *PlayerQuestInsertLog) CommitToTLog() {
}

func (log *PlayerQuestInsertLog) CommitToXdLog() {
}

func (log *PlayerQuestInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerQuestDeleteLog struct {
	db *Database
	Old *C.PlayerQuest
	GoOld *PlayerQuest
}

func (db *Database) newPlayerQuestDeleteLog(old *C.PlayerQuest, goOld *PlayerQuest) *PlayerQuestDeleteLog {
	return &PlayerQuestDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerQuestDeleteLog) Free() {
	C.Free_PlayerQuest(log.Old)
}

func (log *PlayerQuestDeleteLog) InvokeHook() {
	if g_Hooks.PlayerQuest != nil {
		g_Hooks.PlayerQuest.Delete(&PlayerQuestRow{row: log.Old})
	}
}

func (log *PlayerQuestDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerQuestDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(116)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.QuestId)
	file.WriteInt8(log.GoOld.State)
}

func (log *PlayerQuestDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerQuest.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerQuestDeleteLog) CommitToTLog() {
}

func (log *PlayerQuestDeleteLog) CommitToXdLog() {
}

func (log *PlayerQuestDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerQuestUpdateLog struct {
	db *Database 
	Row *C.PlayerQuest
	Old *C.PlayerQuest
	GoNew *PlayerQuest
	GoOld *PlayerQuest
}

func (db *Database) newPlayerQuestUpdateLog(row, old *C.PlayerQuest, goNew, goOld *PlayerQuest) *PlayerQuestUpdateLog {
	return &PlayerQuestUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerQuestUpdateLog) Free() {
	C.Free_PlayerQuest(log.Old)
}

func (log *PlayerQuestUpdateLog) InvokeHook() {
	if g_Hooks.PlayerQuest != nil {
		g_Hooks.PlayerQuest.Update(&PlayerQuestRow{row: log.Row}, &PlayerQuestRow{row: log.Old})
	}
}

func (log *PlayerQuestUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerQuestUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(116)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.QuestId)
	file.WriteInt8(log.GoNew.State)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.QuestId)
	file.WriteInt8(log.GoOld.State)
}

func (log *PlayerQuestUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerQuest.Update
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.QuestId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.State)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerQuestUpdateLog) CommitToTLog() {
}

func (log *PlayerQuestUpdateLog) CommitToXdLog() {
}

func (log *PlayerQuestUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerQuestInsertLog) Rollback() {
	if log.db.tables.PlayerQuest != log.Row {
		panic("insert rollback check failed 'PlayerQuest'")
	}
	log.db.tables.PlayerQuest = nil
	C.Free_PlayerQuest(log.Row)
}

func (log *PlayerQuestDeleteLog) Rollback() {
	if log.db.tables.PlayerQuest != nil {
		panic("delete rollback check failed 'player_quest'")
	}
	log.db.tables.PlayerQuest = log.Old
}

func (log *PlayerQuestUpdateLog) Rollback() {
	if log.db.tables.PlayerQuest != log.Row {
		panic("update rollback check failed 'player_quest'")
	}
	log.db.tables.PlayerQuest = log.Old
	C.Free_PlayerQuest(log.Row)
}

/* ========== player_rainbow_level ========== */

type PlayerRainbowLevelInsertLog struct {
	db *Database 
	Row *C.PlayerRainbowLevel
	GoRow *PlayerRainbowLevel
}

func (db *Database) newPlayerRainbowLevelInsertLog(row *C.PlayerRainbowLevel, goRow *PlayerRainbowLevel) *PlayerRainbowLevelInsertLog {
	return &PlayerRainbowLevelInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerRainbowLevelInsertLog) Free() {
}

func (log *PlayerRainbowLevelInsertLog) InvokeHook() {
	if g_Hooks.PlayerRainbowLevel != nil {
		g_Hooks.PlayerRainbowLevel.Insert(&PlayerRainbowLevelRow{row: log.Row})
	}
}

func (log *PlayerRainbowLevelInsertLog) SetEventId(time.Time) {
}

func (log *PlayerRainbowLevelInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(117)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.ResetTimestamp)
	file.WriteInt32(log.GoRow.ResetNum)
	file.WriteInt16(log.GoRow.Segment)
	file.WriteInt8(log.GoRow.Order)
	file.WriteInt8(log.GoRow.Status)
	file.WriteInt16(log.GoRow.BestSegment)
	file.WriteInt8(log.GoRow.BestOrder)
	file.WriteInt64(log.GoRow.BestRecordTimestamp)
	file.WriteInt16(log.GoRow.MaxOpenSegment)
	file.WriteInt16(log.GoRow.MaxPassSegment)
	file.WriteInt8(log.GoRow.AutoFightNum)
	file.WriteInt16(log.GoRow.BuyTimes)
	file.WriteInt64(log.GoRow.AutoFightTime)
	file.WriteInt64(log.GoRow.BuyTimestamp)
}

func (log *PlayerRainbowLevelInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerRainbowLevel.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ResetTimestamp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.ResetNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Segment)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Order)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Status)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.BestSegment)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.BestOrder)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BestRecordTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.MaxOpenSegment)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.MaxPassSegment)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.AutoFightNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.BuyTimes)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.AutoFightTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BuyTimestamp)))
	return stmt.Execute()
}

func (log *PlayerRainbowLevelInsertLog) CommitToTLog() {
}

func (log *PlayerRainbowLevelInsertLog) CommitToXdLog() {
}

func (log *PlayerRainbowLevelInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerRainbowLevelDeleteLog struct {
	db *Database
	Old *C.PlayerRainbowLevel
	GoOld *PlayerRainbowLevel
}

func (db *Database) newPlayerRainbowLevelDeleteLog(old *C.PlayerRainbowLevel, goOld *PlayerRainbowLevel) *PlayerRainbowLevelDeleteLog {
	return &PlayerRainbowLevelDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerRainbowLevelDeleteLog) Free() {
	C.Free_PlayerRainbowLevel(log.Old)
}

func (log *PlayerRainbowLevelDeleteLog) InvokeHook() {
	if g_Hooks.PlayerRainbowLevel != nil {
		g_Hooks.PlayerRainbowLevel.Delete(&PlayerRainbowLevelRow{row: log.Old})
	}
}

func (log *PlayerRainbowLevelDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerRainbowLevelDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(117)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.ResetTimestamp)
	file.WriteInt32(log.GoOld.ResetNum)
	file.WriteInt16(log.GoOld.Segment)
	file.WriteInt8(log.GoOld.Order)
	file.WriteInt8(log.GoOld.Status)
	file.WriteInt16(log.GoOld.BestSegment)
	file.WriteInt8(log.GoOld.BestOrder)
	file.WriteInt64(log.GoOld.BestRecordTimestamp)
	file.WriteInt16(log.GoOld.MaxOpenSegment)
	file.WriteInt16(log.GoOld.MaxPassSegment)
	file.WriteInt8(log.GoOld.AutoFightNum)
	file.WriteInt16(log.GoOld.BuyTimes)
	file.WriteInt64(log.GoOld.AutoFightTime)
	file.WriteInt64(log.GoOld.BuyTimestamp)
}

func (log *PlayerRainbowLevelDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerRainbowLevel.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerRainbowLevelDeleteLog) CommitToTLog() {
}

func (log *PlayerRainbowLevelDeleteLog) CommitToXdLog() {
}

func (log *PlayerRainbowLevelDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerRainbowLevelUpdateLog struct {
	db *Database 
	Row *C.PlayerRainbowLevel
	Old *C.PlayerRainbowLevel
	GoNew *PlayerRainbowLevel
	GoOld *PlayerRainbowLevel
}

func (db *Database) newPlayerRainbowLevelUpdateLog(row, old *C.PlayerRainbowLevel, goNew, goOld *PlayerRainbowLevel) *PlayerRainbowLevelUpdateLog {
	return &PlayerRainbowLevelUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerRainbowLevelUpdateLog) Free() {
	C.Free_PlayerRainbowLevel(log.Old)
}

func (log *PlayerRainbowLevelUpdateLog) InvokeHook() {
	if g_Hooks.PlayerRainbowLevel != nil {
		g_Hooks.PlayerRainbowLevel.Update(&PlayerRainbowLevelRow{row: log.Row}, &PlayerRainbowLevelRow{row: log.Old})
	}
}

func (log *PlayerRainbowLevelUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerRainbowLevelUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(117)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.ResetTimestamp)
	file.WriteInt32(log.GoNew.ResetNum)
	file.WriteInt16(log.GoNew.Segment)
	file.WriteInt8(log.GoNew.Order)
	file.WriteInt8(log.GoNew.Status)
	file.WriteInt16(log.GoNew.BestSegment)
	file.WriteInt8(log.GoNew.BestOrder)
	file.WriteInt64(log.GoNew.BestRecordTimestamp)
	file.WriteInt16(log.GoNew.MaxOpenSegment)
	file.WriteInt16(log.GoNew.MaxPassSegment)
	file.WriteInt8(log.GoNew.AutoFightNum)
	file.WriteInt16(log.GoNew.BuyTimes)
	file.WriteInt64(log.GoNew.AutoFightTime)
	file.WriteInt64(log.GoNew.BuyTimestamp)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.ResetTimestamp)
	file.WriteInt32(log.GoOld.ResetNum)
	file.WriteInt16(log.GoOld.Segment)
	file.WriteInt8(log.GoOld.Order)
	file.WriteInt8(log.GoOld.Status)
	file.WriteInt16(log.GoOld.BestSegment)
	file.WriteInt8(log.GoOld.BestOrder)
	file.WriteInt64(log.GoOld.BestRecordTimestamp)
	file.WriteInt16(log.GoOld.MaxOpenSegment)
	file.WriteInt16(log.GoOld.MaxPassSegment)
	file.WriteInt8(log.GoOld.AutoFightNum)
	file.WriteInt16(log.GoOld.BuyTimes)
	file.WriteInt64(log.GoOld.AutoFightTime)
	file.WriteInt64(log.GoOld.BuyTimestamp)
}

func (log *PlayerRainbowLevelUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerRainbowLevel.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.ResetTimestamp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.ResetNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Segment)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Order)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Status)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.BestSegment)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.BestOrder)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BestRecordTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.MaxOpenSegment)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.MaxPassSegment)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.AutoFightNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.BuyTimes)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.AutoFightTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BuyTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerRainbowLevelUpdateLog) CommitToTLog() {
}

func (log *PlayerRainbowLevelUpdateLog) CommitToXdLog() {
}

func (log *PlayerRainbowLevelUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerRainbowLevelInsertLog) Rollback() {
	if log.db.tables.PlayerRainbowLevel != log.Row {
		panic("insert rollback check failed 'PlayerRainbowLevel'")
	}
	log.db.tables.PlayerRainbowLevel = nil
	C.Free_PlayerRainbowLevel(log.Row)
}

func (log *PlayerRainbowLevelDeleteLog) Rollback() {
	if log.db.tables.PlayerRainbowLevel != nil {
		panic("delete rollback check failed 'player_rainbow_level'")
	}
	log.db.tables.PlayerRainbowLevel = log.Old
}

func (log *PlayerRainbowLevelUpdateLog) Rollback() {
	if log.db.tables.PlayerRainbowLevel != log.Row {
		panic("update rollback check failed 'player_rainbow_level'")
	}
	log.db.tables.PlayerRainbowLevel = log.Old
	C.Free_PlayerRainbowLevel(log.Row)
}

/* ========== player_rainbow_level_state_bin ========== */

type PlayerRainbowLevelStateBinInsertLog struct {
	db *Database 
	Row *C.PlayerRainbowLevelStateBin
	GoRow *PlayerRainbowLevelStateBin
}

func (db *Database) newPlayerRainbowLevelStateBinInsertLog(row *C.PlayerRainbowLevelStateBin, goRow *PlayerRainbowLevelStateBin) *PlayerRainbowLevelStateBinInsertLog {
	return &PlayerRainbowLevelStateBinInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerRainbowLevelStateBinInsertLog) Free() {
}

func (log *PlayerRainbowLevelStateBinInsertLog) InvokeHook() {
	if g_Hooks.PlayerRainbowLevelStateBin != nil {
		g_Hooks.PlayerRainbowLevelStateBin.Insert(&PlayerRainbowLevelStateBinRow{row: log.Row})
	}
}

func (log *PlayerRainbowLevelStateBinInsertLog) SetEventId(time.Time) {
}

func (log *PlayerRainbowLevelStateBinInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(118)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteBytes(log.GoRow.Bin)
}

func (log *PlayerRainbowLevelStateBinInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerRainbowLevelStateBin.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBinary(unsafe.Pointer(log.Row.Bin), int(log.Row.Bin_len))
	return stmt.Execute()
}

func (log *PlayerRainbowLevelStateBinInsertLog) CommitToTLog() {
}

func (log *PlayerRainbowLevelStateBinInsertLog) CommitToXdLog() {
}

func (log *PlayerRainbowLevelStateBinInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerRainbowLevelStateBinDeleteLog struct {
	db *Database
	Old *C.PlayerRainbowLevelStateBin
	GoOld *PlayerRainbowLevelStateBin
}

func (db *Database) newPlayerRainbowLevelStateBinDeleteLog(old *C.PlayerRainbowLevelStateBin, goOld *PlayerRainbowLevelStateBin) *PlayerRainbowLevelStateBinDeleteLog {
	return &PlayerRainbowLevelStateBinDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerRainbowLevelStateBinDeleteLog) Free() {
	C.Free_PlayerRainbowLevelStateBin(log.Old)
}

func (log *PlayerRainbowLevelStateBinDeleteLog) InvokeHook() {
	if g_Hooks.PlayerRainbowLevelStateBin != nil {
		g_Hooks.PlayerRainbowLevelStateBin.Delete(&PlayerRainbowLevelStateBinRow{row: log.Old})
	}
}

func (log *PlayerRainbowLevelStateBinDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerRainbowLevelStateBinDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(118)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteBytes(log.GoOld.Bin)
}

func (log *PlayerRainbowLevelStateBinDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerRainbowLevelStateBin.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerRainbowLevelStateBinDeleteLog) CommitToTLog() {
}

func (log *PlayerRainbowLevelStateBinDeleteLog) CommitToXdLog() {
}

func (log *PlayerRainbowLevelStateBinDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerRainbowLevelStateBinUpdateLog struct {
	db *Database 
	Row *C.PlayerRainbowLevelStateBin
	Old *C.PlayerRainbowLevelStateBin
	GoNew *PlayerRainbowLevelStateBin
	GoOld *PlayerRainbowLevelStateBin
}

func (db *Database) newPlayerRainbowLevelStateBinUpdateLog(row, old *C.PlayerRainbowLevelStateBin, goNew, goOld *PlayerRainbowLevelStateBin) *PlayerRainbowLevelStateBinUpdateLog {
	return &PlayerRainbowLevelStateBinUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerRainbowLevelStateBinUpdateLog) Free() {
	C.Free_PlayerRainbowLevelStateBin(log.Old)
}

func (log *PlayerRainbowLevelStateBinUpdateLog) InvokeHook() {
	if g_Hooks.PlayerRainbowLevelStateBin != nil {
		g_Hooks.PlayerRainbowLevelStateBin.Update(&PlayerRainbowLevelStateBinRow{row: log.Row}, &PlayerRainbowLevelStateBinRow{row: log.Old})
	}
}

func (log *PlayerRainbowLevelStateBinUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerRainbowLevelStateBinUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(118)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteBytes(log.GoNew.Bin)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteBytes(log.GoOld.Bin)
}

func (log *PlayerRainbowLevelStateBinUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerRainbowLevelStateBin.Update
	stmt.BindBinary(unsafe.Pointer(log.Row.Bin), int(log.Row.Bin_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerRainbowLevelStateBinUpdateLog) CommitToTLog() {
}

func (log *PlayerRainbowLevelStateBinUpdateLog) CommitToXdLog() {
}

func (log *PlayerRainbowLevelStateBinUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerRainbowLevelStateBinInsertLog) Rollback() {
	if log.db.tables.PlayerRainbowLevelStateBin != log.Row {
		panic("insert rollback check failed 'PlayerRainbowLevelStateBin'")
	}
	log.db.tables.PlayerRainbowLevelStateBin = nil
	C.Free_PlayerRainbowLevelStateBin(log.Row)
}

func (log *PlayerRainbowLevelStateBinDeleteLog) Rollback() {
	if log.db.tables.PlayerRainbowLevelStateBin != nil {
		panic("delete rollback check failed 'player_rainbow_level_state_bin'")
	}
	log.db.tables.PlayerRainbowLevelStateBin = log.Old
}

func (log *PlayerRainbowLevelStateBinUpdateLog) Rollback() {
	if log.db.tables.PlayerRainbowLevelStateBin != log.Row {
		panic("update rollback check failed 'player_rainbow_level_state_bin'")
	}
	log.db.tables.PlayerRainbowLevelStateBin = log.Old
	C.Free_PlayerRainbowLevelStateBin(log.Row)
}

/* ========== player_role ========== */

type PlayerRoleInsertLog struct {
	db *Database 
	Row *C.PlayerRole
	GoRow *PlayerRole
}

func (db *Database) newPlayerRoleInsertLog(row *C.PlayerRole, goRow *PlayerRole) *PlayerRoleInsertLog {
	return &PlayerRoleInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerRoleInsertLog) Free() {
}

func (log *PlayerRoleInsertLog) InvokeHook() {
	if g_Hooks.PlayerRole != nil {
		g_Hooks.PlayerRole.Insert(&PlayerRoleRow{row: log.Row})
	}
}

func (log *PlayerRoleInsertLog) SetEventId(time.Time) {
}

func (log *PlayerRoleInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(119)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.RoleId)
	file.WriteInt16(log.GoRow.Level)
	file.WriteInt64(log.GoRow.Exp)
	file.WriteInt32(log.GoRow.FriendshipLevel)
	file.WriteInt16(log.GoRow.Status)
	file.WriteInt16(log.GoRow.CoopId)
	file.WriteInt64(log.GoRow.Timestamp)
	file.WriteInt32(log.GoRow.FightNum)
}

func (log *PlayerRoleInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerRole.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RoleId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Exp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.FriendshipLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Status)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CoopId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.FightNum)))
	return stmt.Execute()
}

func (log *PlayerRoleInsertLog) CommitToTLog() {
}

func (log *PlayerRoleInsertLog) CommitToXdLog() {
}

func (log *PlayerRoleInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerRoleDeleteLog struct {
	db *Database
	Old *C.PlayerRole
	GoOld *PlayerRole
}

func (db *Database) newPlayerRoleDeleteLog(old *C.PlayerRole, goOld *PlayerRole) *PlayerRoleDeleteLog {
	return &PlayerRoleDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerRoleDeleteLog) Free() {
	C.Free_PlayerRole(log.Old)
}

func (log *PlayerRoleDeleteLog) InvokeHook() {
	if g_Hooks.PlayerRole != nil {
		g_Hooks.PlayerRole.Delete(&PlayerRoleRow{row: log.Old})
	}
}

func (log *PlayerRoleDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerRoleDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(119)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.RoleId)
	file.WriteInt16(log.GoOld.Level)
	file.WriteInt64(log.GoOld.Exp)
	file.WriteInt32(log.GoOld.FriendshipLevel)
	file.WriteInt16(log.GoOld.Status)
	file.WriteInt16(log.GoOld.CoopId)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt32(log.GoOld.FightNum)
}

func (log *PlayerRoleDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerRole.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerRoleDeleteLog) CommitToTLog() {
}

func (log *PlayerRoleDeleteLog) CommitToXdLog() {
}

func (log *PlayerRoleDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerRoleUpdateLog struct {
	db *Database 
	Row *C.PlayerRole
	Old *C.PlayerRole
	GoNew *PlayerRole
	GoOld *PlayerRole
}

func (db *Database) newPlayerRoleUpdateLog(row, old *C.PlayerRole, goNew, goOld *PlayerRole) *PlayerRoleUpdateLog {
	return &PlayerRoleUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerRoleUpdateLog) Free() {
	C.Free_PlayerRole(log.Old)
}

func (log *PlayerRoleUpdateLog) InvokeHook() {
	if g_Hooks.PlayerRole != nil {
		g_Hooks.PlayerRole.Update(&PlayerRoleRow{row: log.Row}, &PlayerRoleRow{row: log.Old})
	}
}

func (log *PlayerRoleUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerRoleUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(119)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.RoleId)
	file.WriteInt16(log.GoNew.Level)
	file.WriteInt64(log.GoNew.Exp)
	file.WriteInt32(log.GoNew.FriendshipLevel)
	file.WriteInt16(log.GoNew.Status)
	file.WriteInt16(log.GoNew.CoopId)
	file.WriteInt64(log.GoNew.Timestamp)
	file.WriteInt32(log.GoNew.FightNum)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.RoleId)
	file.WriteInt16(log.GoOld.Level)
	file.WriteInt64(log.GoOld.Exp)
	file.WriteInt32(log.GoOld.FriendshipLevel)
	file.WriteInt16(log.GoOld.Status)
	file.WriteInt16(log.GoOld.CoopId)
	file.WriteInt64(log.GoOld.Timestamp)
	file.WriteInt32(log.GoOld.FightNum)
}

func (log *PlayerRoleUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerRole.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RoleId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Exp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.FriendshipLevel)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Status)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.CoopId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.FightNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerRoleUpdateLog) CommitToTLog() {
}

func (log *PlayerRoleUpdateLog) CommitToXdLog() {
}

func (log *PlayerRoleUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerRoleInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerRole
	for crow := log.db.tables.PlayerRole; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerRole'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerRole = log.db.tables.PlayerRole.next
	}
	C.Free_PlayerRole(log.Row)
}

func (log *PlayerRoleDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerRole
	for crow := log.db.tables.PlayerRole; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_role'")
	}
	log.Old.next = log.db.tables.PlayerRole
	log.db.tables.PlayerRole = log.Old
}

func (log *PlayerRoleUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerRole
	for crow := log.db.tables.PlayerRole; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_role'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerRole = log.Old
	}
	C.Free_PlayerRole(log.Row)
}

/* ========== player_sealedbook ========== */

type PlayerSealedbookInsertLog struct {
	db *Database 
	Row *C.PlayerSealedbook
	GoRow *PlayerSealedbook
}

func (db *Database) newPlayerSealedbookInsertLog(row *C.PlayerSealedbook, goRow *PlayerSealedbook) *PlayerSealedbookInsertLog {
	return &PlayerSealedbookInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerSealedbookInsertLog) Free() {
}

func (log *PlayerSealedbookInsertLog) InvokeHook() {
	if g_Hooks.PlayerSealedbook != nil {
		g_Hooks.PlayerSealedbook.Insert(&PlayerSealedbookRow{row: log.Row})
	}
}

func (log *PlayerSealedbookInsertLog) SetEventId(time.Time) {
}

func (log *PlayerSealedbookInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(120)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteBytes(log.GoRow.SealedRecord)
}

func (log *PlayerSealedbookInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerSealedbook.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBinary(unsafe.Pointer(log.Row.SealedRecord), int(log.Row.SealedRecord_len))
	return stmt.Execute()
}

func (log *PlayerSealedbookInsertLog) CommitToTLog() {
}

func (log *PlayerSealedbookInsertLog) CommitToXdLog() {
}

func (log *PlayerSealedbookInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerSealedbookDeleteLog struct {
	db *Database
	Old *C.PlayerSealedbook
	GoOld *PlayerSealedbook
}

func (db *Database) newPlayerSealedbookDeleteLog(old *C.PlayerSealedbook, goOld *PlayerSealedbook) *PlayerSealedbookDeleteLog {
	return &PlayerSealedbookDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerSealedbookDeleteLog) Free() {
	C.Free_PlayerSealedbook(log.Old)
}

func (log *PlayerSealedbookDeleteLog) InvokeHook() {
	if g_Hooks.PlayerSealedbook != nil {
		g_Hooks.PlayerSealedbook.Delete(&PlayerSealedbookRow{row: log.Old})
	}
}

func (log *PlayerSealedbookDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerSealedbookDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(120)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteBytes(log.GoOld.SealedRecord)
}

func (log *PlayerSealedbookDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerSealedbook.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerSealedbookDeleteLog) CommitToTLog() {
}

func (log *PlayerSealedbookDeleteLog) CommitToXdLog() {
}

func (log *PlayerSealedbookDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerSealedbookUpdateLog struct {
	db *Database 
	Row *C.PlayerSealedbook
	Old *C.PlayerSealedbook
	GoNew *PlayerSealedbook
	GoOld *PlayerSealedbook
}

func (db *Database) newPlayerSealedbookUpdateLog(row, old *C.PlayerSealedbook, goNew, goOld *PlayerSealedbook) *PlayerSealedbookUpdateLog {
	return &PlayerSealedbookUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerSealedbookUpdateLog) Free() {
	C.Free_PlayerSealedbook(log.Old)
}

func (log *PlayerSealedbookUpdateLog) InvokeHook() {
	if g_Hooks.PlayerSealedbook != nil {
		g_Hooks.PlayerSealedbook.Update(&PlayerSealedbookRow{row: log.Row}, &PlayerSealedbookRow{row: log.Old})
	}
}

func (log *PlayerSealedbookUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerSealedbookUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(120)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteBytes(log.GoNew.SealedRecord)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteBytes(log.GoOld.SealedRecord)
}

func (log *PlayerSealedbookUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerSealedbook.Update
	stmt.BindBinary(unsafe.Pointer(log.Row.SealedRecord), int(log.Row.SealedRecord_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerSealedbookUpdateLog) CommitToTLog() {
}

func (log *PlayerSealedbookUpdateLog) CommitToXdLog() {
}

func (log *PlayerSealedbookUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerSealedbookInsertLog) Rollback() {
	if log.db.tables.PlayerSealedbook != log.Row {
		panic("insert rollback check failed 'PlayerSealedbook'")
	}
	log.db.tables.PlayerSealedbook = nil
	C.Free_PlayerSealedbook(log.Row)
}

func (log *PlayerSealedbookDeleteLog) Rollback() {
	if log.db.tables.PlayerSealedbook != nil {
		panic("delete rollback check failed 'player_sealedbook'")
	}
	log.db.tables.PlayerSealedbook = log.Old
}

func (log *PlayerSealedbookUpdateLog) Rollback() {
	if log.db.tables.PlayerSealedbook != log.Row {
		panic("update rollback check failed 'player_sealedbook'")
	}
	log.db.tables.PlayerSealedbook = log.Old
	C.Free_PlayerSealedbook(log.Row)
}

/* ========== player_send_heart_record ========== */

type PlayerSendHeartRecordInsertLog struct {
	db *Database 
	Row *C.PlayerSendHeartRecord
	GoRow *PlayerSendHeartRecord
}

func (db *Database) newPlayerSendHeartRecordInsertLog(row *C.PlayerSendHeartRecord, goRow *PlayerSendHeartRecord) *PlayerSendHeartRecordInsertLog {
	return &PlayerSendHeartRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerSendHeartRecordInsertLog) Free() {
}

func (log *PlayerSendHeartRecordInsertLog) InvokeHook() {
	if g_Hooks.PlayerSendHeartRecord != nil {
		g_Hooks.PlayerSendHeartRecord.Insert(&PlayerSendHeartRecordRow{row: log.Row})
	}
}

func (log *PlayerSendHeartRecordInsertLog) SetEventId(time.Time) {
}

func (log *PlayerSendHeartRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(121)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.FriendPid)
	file.WriteInt64(log.GoRow.SendHeartTime)
}

func (log *PlayerSendHeartRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerSendHeartRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FriendPid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SendHeartTime)))
	return stmt.Execute()
}

func (log *PlayerSendHeartRecordInsertLog) CommitToTLog() {
}

func (log *PlayerSendHeartRecordInsertLog) CommitToXdLog() {
}

func (log *PlayerSendHeartRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerSendHeartRecordDeleteLog struct {
	db *Database
	Old *C.PlayerSendHeartRecord
	GoOld *PlayerSendHeartRecord
}

func (db *Database) newPlayerSendHeartRecordDeleteLog(old *C.PlayerSendHeartRecord, goOld *PlayerSendHeartRecord) *PlayerSendHeartRecordDeleteLog {
	return &PlayerSendHeartRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerSendHeartRecordDeleteLog) Free() {
	C.Free_PlayerSendHeartRecord(log.Old)
}

func (log *PlayerSendHeartRecordDeleteLog) InvokeHook() {
	if g_Hooks.PlayerSendHeartRecord != nil {
		g_Hooks.PlayerSendHeartRecord.Delete(&PlayerSendHeartRecordRow{row: log.Old})
	}
}

func (log *PlayerSendHeartRecordDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerSendHeartRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(121)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.FriendPid)
	file.WriteInt64(log.GoOld.SendHeartTime)
}

func (log *PlayerSendHeartRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerSendHeartRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerSendHeartRecordDeleteLog) CommitToTLog() {
}

func (log *PlayerSendHeartRecordDeleteLog) CommitToXdLog() {
}

func (log *PlayerSendHeartRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerSendHeartRecordUpdateLog struct {
	db *Database 
	Row *C.PlayerSendHeartRecord
	Old *C.PlayerSendHeartRecord
	GoNew *PlayerSendHeartRecord
	GoOld *PlayerSendHeartRecord
}

func (db *Database) newPlayerSendHeartRecordUpdateLog(row, old *C.PlayerSendHeartRecord, goNew, goOld *PlayerSendHeartRecord) *PlayerSendHeartRecordUpdateLog {
	return &PlayerSendHeartRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerSendHeartRecordUpdateLog) Free() {
	C.Free_PlayerSendHeartRecord(log.Old)
}

func (log *PlayerSendHeartRecordUpdateLog) InvokeHook() {
	if g_Hooks.PlayerSendHeartRecord != nil {
		g_Hooks.PlayerSendHeartRecord.Update(&PlayerSendHeartRecordRow{row: log.Row}, &PlayerSendHeartRecordRow{row: log.Old})
	}
}

func (log *PlayerSendHeartRecordUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerSendHeartRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(121)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.FriendPid)
	file.WriteInt64(log.GoNew.SendHeartTime)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.FriendPid)
	file.WriteInt64(log.GoOld.SendHeartTime)
}

func (log *PlayerSendHeartRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerSendHeartRecord.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FriendPid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.SendHeartTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerSendHeartRecordUpdateLog) CommitToTLog() {
}

func (log *PlayerSendHeartRecordUpdateLog) CommitToXdLog() {
}

func (log *PlayerSendHeartRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerSendHeartRecordInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerSendHeartRecord
	for crow := log.db.tables.PlayerSendHeartRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerSendHeartRecord'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerSendHeartRecord = log.db.tables.PlayerSendHeartRecord.next
	}
	C.Free_PlayerSendHeartRecord(log.Row)
}

func (log *PlayerSendHeartRecordDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerSendHeartRecord
	for crow := log.db.tables.PlayerSendHeartRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_send_heart_record'")
	}
	log.Old.next = log.db.tables.PlayerSendHeartRecord
	log.db.tables.PlayerSendHeartRecord = log.Old
}

func (log *PlayerSendHeartRecordUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerSendHeartRecord
	for crow := log.db.tables.PlayerSendHeartRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_send_heart_record'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerSendHeartRecord = log.Old
	}
	C.Free_PlayerSendHeartRecord(log.Row)
}

/* ========== player_skill ========== */

type PlayerSkillInsertLog struct {
	db *Database 
	Row *C.PlayerSkill
	GoRow *PlayerSkill
}

func (db *Database) newPlayerSkillInsertLog(row *C.PlayerSkill, goRow *PlayerSkill) *PlayerSkillInsertLog {
	return &PlayerSkillInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerSkillInsertLog) Free() {
}

func (log *PlayerSkillInsertLog) InvokeHook() {
	if g_Hooks.PlayerSkill != nil {
		g_Hooks.PlayerSkill.Insert(&PlayerSkillRow{row: log.Row})
	}
}

func (log *PlayerSkillInsertLog) SetEventId(time.Time) {
}

func (log *PlayerSkillInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(122)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.RoleId)
	file.WriteInt16(log.GoRow.SkillId)
	file.WriteInt32(log.GoRow.SkillTrnlv)
}

func (log *PlayerSkillInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerSkill.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RoleId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SkillId)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.SkillTrnlv)))
	return stmt.Execute()
}

func (log *PlayerSkillInsertLog) CommitToTLog() {
}

func (log *PlayerSkillInsertLog) CommitToXdLog() {
}

func (log *PlayerSkillInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerSkillDeleteLog struct {
	db *Database
	Old *C.PlayerSkill
	GoOld *PlayerSkill
}

func (db *Database) newPlayerSkillDeleteLog(old *C.PlayerSkill, goOld *PlayerSkill) *PlayerSkillDeleteLog {
	return &PlayerSkillDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerSkillDeleteLog) Free() {
	C.Free_PlayerSkill(log.Old)
}

func (log *PlayerSkillDeleteLog) InvokeHook() {
	if g_Hooks.PlayerSkill != nil {
		g_Hooks.PlayerSkill.Delete(&PlayerSkillRow{row: log.Old})
	}
}

func (log *PlayerSkillDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerSkillDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(122)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.RoleId)
	file.WriteInt16(log.GoOld.SkillId)
	file.WriteInt32(log.GoOld.SkillTrnlv)
}

func (log *PlayerSkillDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerSkill.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerSkillDeleteLog) CommitToTLog() {
}

func (log *PlayerSkillDeleteLog) CommitToXdLog() {
}

func (log *PlayerSkillDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerSkillUpdateLog struct {
	db *Database 
	Row *C.PlayerSkill
	Old *C.PlayerSkill
	GoNew *PlayerSkill
	GoOld *PlayerSkill
}

func (db *Database) newPlayerSkillUpdateLog(row, old *C.PlayerSkill, goNew, goOld *PlayerSkill) *PlayerSkillUpdateLog {
	return &PlayerSkillUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerSkillUpdateLog) Free() {
	C.Free_PlayerSkill(log.Old)
}

func (log *PlayerSkillUpdateLog) InvokeHook() {
	if g_Hooks.PlayerSkill != nil {
		g_Hooks.PlayerSkill.Update(&PlayerSkillRow{row: log.Row}, &PlayerSkillRow{row: log.Old})
	}
}

func (log *PlayerSkillUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerSkillUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(122)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.RoleId)
	file.WriteInt16(log.GoNew.SkillId)
	file.WriteInt32(log.GoNew.SkillTrnlv)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.RoleId)
	file.WriteInt16(log.GoOld.SkillId)
	file.WriteInt32(log.GoOld.SkillTrnlv)
}

func (log *PlayerSkillUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerSkill.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RoleId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SkillId)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.SkillTrnlv)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerSkillUpdateLog) CommitToTLog() {
}

func (log *PlayerSkillUpdateLog) CommitToXdLog() {
}

func (log *PlayerSkillUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerSkillInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerSkill
	for crow := log.db.tables.PlayerSkill; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerSkill'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerSkill = log.db.tables.PlayerSkill.next
	}
	C.Free_PlayerSkill(log.Row)
}

func (log *PlayerSkillDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerSkill
	for crow := log.db.tables.PlayerSkill; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_skill'")
	}
	log.Old.next = log.db.tables.PlayerSkill
	log.db.tables.PlayerSkill = log.Old
}

func (log *PlayerSkillUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerSkill
	for crow := log.db.tables.PlayerSkill; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_skill'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerSkill = log.Old
	}
	C.Free_PlayerSkill(log.Row)
}

/* ========== player_state ========== */

type PlayerStateInsertLog struct {
	db *Database 
	Row *C.PlayerState
	GoRow *PlayerState
}

func (db *Database) newPlayerStateInsertLog(row *C.PlayerState, goRow *PlayerState) *PlayerStateInsertLog {
	return &PlayerStateInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerStateInsertLog) Free() {
}

func (log *PlayerStateInsertLog) InvokeHook() {
	if g_Hooks.PlayerState != nil {
		g_Hooks.PlayerState.Insert(&PlayerStateRow{row: log.Row})
	}
}

func (log *PlayerStateInsertLog) SetEventId(time.Time) {
}

func (log *PlayerStateInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(123)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.BanStartTime)
	file.WriteInt64(log.GoRow.BanEndTime)
}

func (log *PlayerStateInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerState.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BanStartTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BanEndTime)))
	return stmt.Execute()
}

func (log *PlayerStateInsertLog) CommitToTLog() {
}

func (log *PlayerStateInsertLog) CommitToXdLog() {
}

func (log *PlayerStateInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerStateDeleteLog struct {
	db *Database
	Old *C.PlayerState
	GoOld *PlayerState
}

func (db *Database) newPlayerStateDeleteLog(old *C.PlayerState, goOld *PlayerState) *PlayerStateDeleteLog {
	return &PlayerStateDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerStateDeleteLog) Free() {
	C.Free_PlayerState(log.Old)
}

func (log *PlayerStateDeleteLog) InvokeHook() {
	if g_Hooks.PlayerState != nil {
		g_Hooks.PlayerState.Delete(&PlayerStateRow{row: log.Old})
	}
}

func (log *PlayerStateDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerStateDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(123)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.BanStartTime)
	file.WriteInt64(log.GoOld.BanEndTime)
}

func (log *PlayerStateDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerState.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerStateDeleteLog) CommitToTLog() {
}

func (log *PlayerStateDeleteLog) CommitToXdLog() {
}

func (log *PlayerStateDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerStateUpdateLog struct {
	db *Database 
	Row *C.PlayerState
	Old *C.PlayerState
	GoNew *PlayerState
	GoOld *PlayerState
}

func (db *Database) newPlayerStateUpdateLog(row, old *C.PlayerState, goNew, goOld *PlayerState) *PlayerStateUpdateLog {
	return &PlayerStateUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerStateUpdateLog) Free() {
	C.Free_PlayerState(log.Old)
}

func (log *PlayerStateUpdateLog) InvokeHook() {
	if g_Hooks.PlayerState != nil {
		g_Hooks.PlayerState.Update(&PlayerStateRow{row: log.Row}, &PlayerStateRow{row: log.Old})
	}
}

func (log *PlayerStateUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerStateUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(123)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.BanStartTime)
	file.WriteInt64(log.GoNew.BanEndTime)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.BanStartTime)
	file.WriteInt64(log.GoOld.BanEndTime)
}

func (log *PlayerStateUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerState.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BanStartTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BanEndTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerStateUpdateLog) CommitToTLog() {
}

func (log *PlayerStateUpdateLog) CommitToXdLog() {
}

func (log *PlayerStateUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerStateInsertLog) Rollback() {
	if log.db.tables.PlayerState != log.Row {
		panic("insert rollback check failed 'PlayerState'")
	}
	log.db.tables.PlayerState = nil
	C.Free_PlayerState(log.Row)
}

func (log *PlayerStateDeleteLog) Rollback() {
	if log.db.tables.PlayerState != nil {
		panic("delete rollback check failed 'player_state'")
	}
	log.db.tables.PlayerState = log.Old
}

func (log *PlayerStateUpdateLog) Rollback() {
	if log.db.tables.PlayerState != log.Row {
		panic("update rollback check failed 'player_state'")
	}
	log.db.tables.PlayerState = log.Old
	C.Free_PlayerState(log.Row)
}

/* ========== player_sword_soul ========== */

type PlayerSwordSoulInsertLog struct {
	db *Database 
	Row *C.PlayerSwordSoul
	GoRow *PlayerSwordSoul
}

func (db *Database) newPlayerSwordSoulInsertLog(row *C.PlayerSwordSoul, goRow *PlayerSwordSoul) *PlayerSwordSoulInsertLog {
	return &PlayerSwordSoulInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerSwordSoulInsertLog) Free() {
}

func (log *PlayerSwordSoulInsertLog) InvokeHook() {
	if g_Hooks.PlayerSwordSoul != nil {
		g_Hooks.PlayerSwordSoul.Insert(&PlayerSwordSoulRow{row: log.Row})
	}
}

func (log *PlayerSwordSoulInsertLog) SetEventId(time.Time) {
}

func (log *PlayerSwordSoulInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(124)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.Pos)
	file.WriteInt16(log.GoRow.SwordSoulId)
	file.WriteInt32(log.GoRow.Exp)
	file.WriteInt8(log.GoRow.Level)
}

func (log *PlayerSwordSoulInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerSwordSoul.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Pos)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SwordSoulId)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Exp)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Level)))
	return stmt.Execute()
}

func (log *PlayerSwordSoulInsertLog) CommitToTLog() {
}

func (log *PlayerSwordSoulInsertLog) CommitToXdLog() {
}

func (log *PlayerSwordSoulInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerSwordSoulDeleteLog struct {
	db *Database
	Old *C.PlayerSwordSoul
	GoOld *PlayerSwordSoul
}

func (db *Database) newPlayerSwordSoulDeleteLog(old *C.PlayerSwordSoul, goOld *PlayerSwordSoul) *PlayerSwordSoulDeleteLog {
	return &PlayerSwordSoulDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerSwordSoulDeleteLog) Free() {
	C.Free_PlayerSwordSoul(log.Old)
}

func (log *PlayerSwordSoulDeleteLog) InvokeHook() {
	if g_Hooks.PlayerSwordSoul != nil {
		g_Hooks.PlayerSwordSoul.Delete(&PlayerSwordSoulRow{row: log.Old})
	}
}

func (log *PlayerSwordSoulDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerSwordSoulDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(124)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.Pos)
	file.WriteInt16(log.GoOld.SwordSoulId)
	file.WriteInt32(log.GoOld.Exp)
	file.WriteInt8(log.GoOld.Level)
}

func (log *PlayerSwordSoulDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerSwordSoul.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerSwordSoulDeleteLog) CommitToTLog() {
}

func (log *PlayerSwordSoulDeleteLog) CommitToXdLog() {
}

func (log *PlayerSwordSoulDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerSwordSoulUpdateLog struct {
	db *Database 
	Row *C.PlayerSwordSoul
	Old *C.PlayerSwordSoul
	GoNew *PlayerSwordSoul
	GoOld *PlayerSwordSoul
}

func (db *Database) newPlayerSwordSoulUpdateLog(row, old *C.PlayerSwordSoul, goNew, goOld *PlayerSwordSoul) *PlayerSwordSoulUpdateLog {
	return &PlayerSwordSoulUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerSwordSoulUpdateLog) Free() {
	C.Free_PlayerSwordSoul(log.Old)
}

func (log *PlayerSwordSoulUpdateLog) InvokeHook() {
	if g_Hooks.PlayerSwordSoul != nil {
		g_Hooks.PlayerSwordSoul.Update(&PlayerSwordSoulRow{row: log.Row}, &PlayerSwordSoulRow{row: log.Old})
	}
}

func (log *PlayerSwordSoulUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerSwordSoulUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(124)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.Pos)
	file.WriteInt16(log.GoNew.SwordSoulId)
	file.WriteInt32(log.GoNew.Exp)
	file.WriteInt8(log.GoNew.Level)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.Pos)
	file.WriteInt16(log.GoOld.SwordSoulId)
	file.WriteInt32(log.GoOld.Exp)
	file.WriteInt8(log.GoOld.Level)
}

func (log *PlayerSwordSoulUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerSwordSoul.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Pos)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SwordSoulId)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Exp)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerSwordSoulUpdateLog) CommitToTLog() {
}

func (log *PlayerSwordSoulUpdateLog) CommitToXdLog() {
}

func (log *PlayerSwordSoulUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerSwordSoulInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerSwordSoul
	for crow := log.db.tables.PlayerSwordSoul; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerSwordSoul'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerSwordSoul = log.db.tables.PlayerSwordSoul.next
	}
	C.Free_PlayerSwordSoul(log.Row)
}

func (log *PlayerSwordSoulDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerSwordSoul
	for crow := log.db.tables.PlayerSwordSoul; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_sword_soul'")
	}
	log.Old.next = log.db.tables.PlayerSwordSoul
	log.db.tables.PlayerSwordSoul = log.Old
}

func (log *PlayerSwordSoulUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerSwordSoul
	for crow := log.db.tables.PlayerSwordSoul; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_sword_soul'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerSwordSoul = log.Old
	}
	C.Free_PlayerSwordSoul(log.Row)
}

/* ========== player_sword_soul_equipment ========== */

type PlayerSwordSoulEquipmentInsertLog struct {
	db *Database 
	Row *C.PlayerSwordSoulEquipment
	GoRow *PlayerSwordSoulEquipment
}

func (db *Database) newPlayerSwordSoulEquipmentInsertLog(row *C.PlayerSwordSoulEquipment, goRow *PlayerSwordSoulEquipment) *PlayerSwordSoulEquipmentInsertLog {
	return &PlayerSwordSoulEquipmentInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerSwordSoulEquipmentInsertLog) Free() {
}

func (log *PlayerSwordSoulEquipmentInsertLog) InvokeHook() {
	if g_Hooks.PlayerSwordSoulEquipment != nil {
		g_Hooks.PlayerSwordSoulEquipment.Insert(&PlayerSwordSoulEquipmentRow{row: log.Row})
	}
}

func (log *PlayerSwordSoulEquipmentInsertLog) SetEventId(time.Time) {
}

func (log *PlayerSwordSoulEquipmentInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(125)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.RoleId)
	file.WriteInt8(log.GoRow.IsEquippedXuanyuan)
	file.WriteInt64(log.GoRow.TypeBits)
	file.WriteInt64(log.GoRow.Pos1)
	file.WriteInt64(log.GoRow.Pos2)
	file.WriteInt64(log.GoRow.Pos3)
	file.WriteInt64(log.GoRow.Pos4)
	file.WriteInt64(log.GoRow.Pos5)
	file.WriteInt64(log.GoRow.Pos6)
	file.WriteInt64(log.GoRow.Pos7)
	file.WriteInt64(log.GoRow.Pos8)
	file.WriteInt64(log.GoRow.Pos9)
}

func (log *PlayerSwordSoulEquipmentInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerSwordSoulEquipment.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RoleId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.IsEquippedXuanyuan)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.TypeBits)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos1)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos2)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos3)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos4)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos5)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos6)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos7)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos8)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos9)))
	return stmt.Execute()
}

func (log *PlayerSwordSoulEquipmentInsertLog) CommitToTLog() {
}

func (log *PlayerSwordSoulEquipmentInsertLog) CommitToXdLog() {
}

func (log *PlayerSwordSoulEquipmentInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerSwordSoulEquipmentDeleteLog struct {
	db *Database
	Old *C.PlayerSwordSoulEquipment
	GoOld *PlayerSwordSoulEquipment
}

func (db *Database) newPlayerSwordSoulEquipmentDeleteLog(old *C.PlayerSwordSoulEquipment, goOld *PlayerSwordSoulEquipment) *PlayerSwordSoulEquipmentDeleteLog {
	return &PlayerSwordSoulEquipmentDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerSwordSoulEquipmentDeleteLog) Free() {
	C.Free_PlayerSwordSoulEquipment(log.Old)
}

func (log *PlayerSwordSoulEquipmentDeleteLog) InvokeHook() {
	if g_Hooks.PlayerSwordSoulEquipment != nil {
		g_Hooks.PlayerSwordSoulEquipment.Delete(&PlayerSwordSoulEquipmentRow{row: log.Old})
	}
}

func (log *PlayerSwordSoulEquipmentDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerSwordSoulEquipmentDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(125)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.RoleId)
	file.WriteInt8(log.GoOld.IsEquippedXuanyuan)
	file.WriteInt64(log.GoOld.TypeBits)
	file.WriteInt64(log.GoOld.Pos1)
	file.WriteInt64(log.GoOld.Pos2)
	file.WriteInt64(log.GoOld.Pos3)
	file.WriteInt64(log.GoOld.Pos4)
	file.WriteInt64(log.GoOld.Pos5)
	file.WriteInt64(log.GoOld.Pos6)
	file.WriteInt64(log.GoOld.Pos7)
	file.WriteInt64(log.GoOld.Pos8)
	file.WriteInt64(log.GoOld.Pos9)
}

func (log *PlayerSwordSoulEquipmentDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerSwordSoulEquipment.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerSwordSoulEquipmentDeleteLog) CommitToTLog() {
}

func (log *PlayerSwordSoulEquipmentDeleteLog) CommitToXdLog() {
}

func (log *PlayerSwordSoulEquipmentDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerSwordSoulEquipmentUpdateLog struct {
	db *Database 
	Row *C.PlayerSwordSoulEquipment
	Old *C.PlayerSwordSoulEquipment
	GoNew *PlayerSwordSoulEquipment
	GoOld *PlayerSwordSoulEquipment
}

func (db *Database) newPlayerSwordSoulEquipmentUpdateLog(row, old *C.PlayerSwordSoulEquipment, goNew, goOld *PlayerSwordSoulEquipment) *PlayerSwordSoulEquipmentUpdateLog {
	return &PlayerSwordSoulEquipmentUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerSwordSoulEquipmentUpdateLog) Free() {
	C.Free_PlayerSwordSoulEquipment(log.Old)
}

func (log *PlayerSwordSoulEquipmentUpdateLog) InvokeHook() {
	if g_Hooks.PlayerSwordSoulEquipment != nil {
		g_Hooks.PlayerSwordSoulEquipment.Update(&PlayerSwordSoulEquipmentRow{row: log.Row}, &PlayerSwordSoulEquipmentRow{row: log.Old})
	}
}

func (log *PlayerSwordSoulEquipmentUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerSwordSoulEquipmentUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(125)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.RoleId)
	file.WriteInt8(log.GoNew.IsEquippedXuanyuan)
	file.WriteInt64(log.GoNew.TypeBits)
	file.WriteInt64(log.GoNew.Pos1)
	file.WriteInt64(log.GoNew.Pos2)
	file.WriteInt64(log.GoNew.Pos3)
	file.WriteInt64(log.GoNew.Pos4)
	file.WriteInt64(log.GoNew.Pos5)
	file.WriteInt64(log.GoNew.Pos6)
	file.WriteInt64(log.GoNew.Pos7)
	file.WriteInt64(log.GoNew.Pos8)
	file.WriteInt64(log.GoNew.Pos9)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.RoleId)
	file.WriteInt8(log.GoOld.IsEquippedXuanyuan)
	file.WriteInt64(log.GoOld.TypeBits)
	file.WriteInt64(log.GoOld.Pos1)
	file.WriteInt64(log.GoOld.Pos2)
	file.WriteInt64(log.GoOld.Pos3)
	file.WriteInt64(log.GoOld.Pos4)
	file.WriteInt64(log.GoOld.Pos5)
	file.WriteInt64(log.GoOld.Pos6)
	file.WriteInt64(log.GoOld.Pos7)
	file.WriteInt64(log.GoOld.Pos8)
	file.WriteInt64(log.GoOld.Pos9)
}

func (log *PlayerSwordSoulEquipmentUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerSwordSoulEquipment.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RoleId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.IsEquippedXuanyuan)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.TypeBits)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos1)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos2)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos3)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos4)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos5)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos6)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos7)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos8)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos9)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerSwordSoulEquipmentUpdateLog) CommitToTLog() {
}

func (log *PlayerSwordSoulEquipmentUpdateLog) CommitToXdLog() {
}

func (log *PlayerSwordSoulEquipmentUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerSwordSoulEquipmentInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerSwordSoulEquipment
	for crow := log.db.tables.PlayerSwordSoulEquipment; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerSwordSoulEquipment'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerSwordSoulEquipment = log.db.tables.PlayerSwordSoulEquipment.next
	}
	C.Free_PlayerSwordSoulEquipment(log.Row)
}

func (log *PlayerSwordSoulEquipmentDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerSwordSoulEquipment
	for crow := log.db.tables.PlayerSwordSoulEquipment; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_sword_soul_equipment'")
	}
	log.Old.next = log.db.tables.PlayerSwordSoulEquipment
	log.db.tables.PlayerSwordSoulEquipment = log.Old
}

func (log *PlayerSwordSoulEquipmentUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerSwordSoulEquipment
	for crow := log.db.tables.PlayerSwordSoulEquipment; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_sword_soul_equipment'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerSwordSoulEquipment = log.Old
	}
	C.Free_PlayerSwordSoulEquipment(log.Row)
}

/* ========== player_sword_soul_state ========== */

type PlayerSwordSoulStateInsertLog struct {
	db *Database 
	Row *C.PlayerSwordSoulState
	GoRow *PlayerSwordSoulState
}

func (db *Database) newPlayerSwordSoulStateInsertLog(row *C.PlayerSwordSoulState, goRow *PlayerSwordSoulState) *PlayerSwordSoulStateInsertLog {
	return &PlayerSwordSoulStateInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerSwordSoulStateInsertLog) Free() {
}

func (log *PlayerSwordSoulStateInsertLog) InvokeHook() {
	if g_Hooks.PlayerSwordSoulState != nil {
		g_Hooks.PlayerSwordSoulState.Insert(&PlayerSwordSoulStateRow{row: log.Row})
	}
}

func (log *PlayerSwordSoulStateInsertLog) SetEventId(time.Time) {
}

func (log *PlayerSwordSoulStateInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(126)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.BoxState)
	file.WriteInt16(log.GoRow.Num)
	file.WriteInt64(log.GoRow.UpdateTime)
	file.WriteInt64(log.GoRow.IngotNum)
	file.WriteInt8(log.GoRow.SupersoulAdditionalPossible)
	file.WriteInt16(log.GoRow.IngotYelloOne)
	file.WriteInt64(log.GoRow.LastIngotDrawTime)
	file.WriteInt32(log.GoRow.SwordSoulNum)
}

func (log *PlayerSwordSoulStateInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerSwordSoulState.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.BoxState)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Num)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.UpdateTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.IngotNum)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.SupersoulAdditionalPossible)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.IngotYelloOne)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastIngotDrawTime)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.SwordSoulNum)))
	return stmt.Execute()
}

func (log *PlayerSwordSoulStateInsertLog) CommitToTLog() {
}

func (log *PlayerSwordSoulStateInsertLog) CommitToXdLog() {
}

func (log *PlayerSwordSoulStateInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerSwordSoulStateDeleteLog struct {
	db *Database
	Old *C.PlayerSwordSoulState
	GoOld *PlayerSwordSoulState
}

func (db *Database) newPlayerSwordSoulStateDeleteLog(old *C.PlayerSwordSoulState, goOld *PlayerSwordSoulState) *PlayerSwordSoulStateDeleteLog {
	return &PlayerSwordSoulStateDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerSwordSoulStateDeleteLog) Free() {
	C.Free_PlayerSwordSoulState(log.Old)
}

func (log *PlayerSwordSoulStateDeleteLog) InvokeHook() {
	if g_Hooks.PlayerSwordSoulState != nil {
		g_Hooks.PlayerSwordSoulState.Delete(&PlayerSwordSoulStateRow{row: log.Old})
	}
}

func (log *PlayerSwordSoulStateDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerSwordSoulStateDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(126)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.BoxState)
	file.WriteInt16(log.GoOld.Num)
	file.WriteInt64(log.GoOld.UpdateTime)
	file.WriteInt64(log.GoOld.IngotNum)
	file.WriteInt8(log.GoOld.SupersoulAdditionalPossible)
	file.WriteInt16(log.GoOld.IngotYelloOne)
	file.WriteInt64(log.GoOld.LastIngotDrawTime)
	file.WriteInt32(log.GoOld.SwordSoulNum)
}

func (log *PlayerSwordSoulStateDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerSwordSoulState.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerSwordSoulStateDeleteLog) CommitToTLog() {
}

func (log *PlayerSwordSoulStateDeleteLog) CommitToXdLog() {
}

func (log *PlayerSwordSoulStateDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerSwordSoulStateUpdateLog struct {
	db *Database 
	Row *C.PlayerSwordSoulState
	Old *C.PlayerSwordSoulState
	GoNew *PlayerSwordSoulState
	GoOld *PlayerSwordSoulState
}

func (db *Database) newPlayerSwordSoulStateUpdateLog(row, old *C.PlayerSwordSoulState, goNew, goOld *PlayerSwordSoulState) *PlayerSwordSoulStateUpdateLog {
	return &PlayerSwordSoulStateUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerSwordSoulStateUpdateLog) Free() {
	C.Free_PlayerSwordSoulState(log.Old)
}

func (log *PlayerSwordSoulStateUpdateLog) InvokeHook() {
	if g_Hooks.PlayerSwordSoulState != nil {
		g_Hooks.PlayerSwordSoulState.Update(&PlayerSwordSoulStateRow{row: log.Row}, &PlayerSwordSoulStateRow{row: log.Old})
	}
}

func (log *PlayerSwordSoulStateUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerSwordSoulStateUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(126)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.BoxState)
	file.WriteInt16(log.GoNew.Num)
	file.WriteInt64(log.GoNew.UpdateTime)
	file.WriteInt64(log.GoNew.IngotNum)
	file.WriteInt8(log.GoNew.SupersoulAdditionalPossible)
	file.WriteInt16(log.GoNew.IngotYelloOne)
	file.WriteInt64(log.GoNew.LastIngotDrawTime)
	file.WriteInt32(log.GoNew.SwordSoulNum)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.BoxState)
	file.WriteInt16(log.GoOld.Num)
	file.WriteInt64(log.GoOld.UpdateTime)
	file.WriteInt64(log.GoOld.IngotNum)
	file.WriteInt8(log.GoOld.SupersoulAdditionalPossible)
	file.WriteInt16(log.GoOld.IngotYelloOne)
	file.WriteInt64(log.GoOld.LastIngotDrawTime)
	file.WriteInt32(log.GoOld.SwordSoulNum)
}

func (log *PlayerSwordSoulStateUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerSwordSoulState.Update
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.BoxState)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Num)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.UpdateTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.IngotNum)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.SupersoulAdditionalPossible)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.IngotYelloOne)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastIngotDrawTime)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.SwordSoulNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerSwordSoulStateUpdateLog) CommitToTLog() {
}

func (log *PlayerSwordSoulStateUpdateLog) CommitToXdLog() {
}

func (log *PlayerSwordSoulStateUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerSwordSoulStateInsertLog) Rollback() {
	if log.db.tables.PlayerSwordSoulState != log.Row {
		panic("insert rollback check failed 'PlayerSwordSoulState'")
	}
	log.db.tables.PlayerSwordSoulState = nil
	C.Free_PlayerSwordSoulState(log.Row)
}

func (log *PlayerSwordSoulStateDeleteLog) Rollback() {
	if log.db.tables.PlayerSwordSoulState != nil {
		panic("delete rollback check failed 'player_sword_soul_state'")
	}
	log.db.tables.PlayerSwordSoulState = log.Old
}

func (log *PlayerSwordSoulStateUpdateLog) Rollback() {
	if log.db.tables.PlayerSwordSoulState != log.Row {
		panic("update rollback check failed 'player_sword_soul_state'")
	}
	log.db.tables.PlayerSwordSoulState = log.Old
	C.Free_PlayerSwordSoulState(log.Row)
}

/* ========== player_taoyuan_bless ========== */

type PlayerTaoyuanBlessInsertLog struct {
	db *Database 
	Row *C.PlayerTaoyuanBless
	GoRow *PlayerTaoyuanBless
}

func (db *Database) newPlayerTaoyuanBlessInsertLog(row *C.PlayerTaoyuanBless, goRow *PlayerTaoyuanBless) *PlayerTaoyuanBlessInsertLog {
	return &PlayerTaoyuanBlessInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerTaoyuanBlessInsertLog) Free() {
}

func (log *PlayerTaoyuanBlessInsertLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanBless != nil {
		g_Hooks.PlayerTaoyuanBless.Insert(&PlayerTaoyuanBlessRow{row: log.Row})
	}
}

func (log *PlayerTaoyuanBlessInsertLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanBlessInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(127)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.DailyBlessTimes)
	file.WriteInt64(log.GoRow.LastBlessTime)
	file.WriteInt32(log.GoRow.DailyBeBlessTimes)
	file.WriteInt64(log.GoRow.LastBeBlessTime)
	file.WriteInt64(log.GoRow.BlessPid1)
	file.WriteInt64(log.GoRow.BlessPid2)
	file.WriteInt64(log.GoRow.BlessPid3)
	file.WriteInt64(log.GoRow.BlessPid4)
	file.WriteInt64(log.GoRow.BlessPid5)
}

func (log *PlayerTaoyuanBlessInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanBless.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.DailyBlessTimes)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastBlessTime)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.DailyBeBlessTimes)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastBeBlessTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BlessPid1)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BlessPid2)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BlessPid3)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BlessPid4)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BlessPid5)))
	return stmt.Execute()
}

func (log *PlayerTaoyuanBlessInsertLog) CommitToTLog() {
}

func (log *PlayerTaoyuanBlessInsertLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanBlessInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTaoyuanBlessDeleteLog struct {
	db *Database
	Old *C.PlayerTaoyuanBless
	GoOld *PlayerTaoyuanBless
}

func (db *Database) newPlayerTaoyuanBlessDeleteLog(old *C.PlayerTaoyuanBless, goOld *PlayerTaoyuanBless) *PlayerTaoyuanBlessDeleteLog {
	return &PlayerTaoyuanBlessDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerTaoyuanBlessDeleteLog) Free() {
	C.Free_PlayerTaoyuanBless(log.Old)
}

func (log *PlayerTaoyuanBlessDeleteLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanBless != nil {
		g_Hooks.PlayerTaoyuanBless.Delete(&PlayerTaoyuanBlessRow{row: log.Old})
	}
}

func (log *PlayerTaoyuanBlessDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanBlessDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(127)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.DailyBlessTimes)
	file.WriteInt64(log.GoOld.LastBlessTime)
	file.WriteInt32(log.GoOld.DailyBeBlessTimes)
	file.WriteInt64(log.GoOld.LastBeBlessTime)
	file.WriteInt64(log.GoOld.BlessPid1)
	file.WriteInt64(log.GoOld.BlessPid2)
	file.WriteInt64(log.GoOld.BlessPid3)
	file.WriteInt64(log.GoOld.BlessPid4)
	file.WriteInt64(log.GoOld.BlessPid5)
}

func (log *PlayerTaoyuanBlessDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanBless.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerTaoyuanBlessDeleteLog) CommitToTLog() {
}

func (log *PlayerTaoyuanBlessDeleteLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanBlessDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTaoyuanBlessUpdateLog struct {
	db *Database 
	Row *C.PlayerTaoyuanBless
	Old *C.PlayerTaoyuanBless
	GoNew *PlayerTaoyuanBless
	GoOld *PlayerTaoyuanBless
}

func (db *Database) newPlayerTaoyuanBlessUpdateLog(row, old *C.PlayerTaoyuanBless, goNew, goOld *PlayerTaoyuanBless) *PlayerTaoyuanBlessUpdateLog {
	return &PlayerTaoyuanBlessUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerTaoyuanBlessUpdateLog) Free() {
	C.Free_PlayerTaoyuanBless(log.Old)
}

func (log *PlayerTaoyuanBlessUpdateLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanBless != nil {
		g_Hooks.PlayerTaoyuanBless.Update(&PlayerTaoyuanBlessRow{row: log.Row}, &PlayerTaoyuanBlessRow{row: log.Old})
	}
}

func (log *PlayerTaoyuanBlessUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanBlessUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(127)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.DailyBlessTimes)
	file.WriteInt64(log.GoNew.LastBlessTime)
	file.WriteInt32(log.GoNew.DailyBeBlessTimes)
	file.WriteInt64(log.GoNew.LastBeBlessTime)
	file.WriteInt64(log.GoNew.BlessPid1)
	file.WriteInt64(log.GoNew.BlessPid2)
	file.WriteInt64(log.GoNew.BlessPid3)
	file.WriteInt64(log.GoNew.BlessPid4)
	file.WriteInt64(log.GoNew.BlessPid5)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.DailyBlessTimes)
	file.WriteInt64(log.GoOld.LastBlessTime)
	file.WriteInt32(log.GoOld.DailyBeBlessTimes)
	file.WriteInt64(log.GoOld.LastBeBlessTime)
	file.WriteInt64(log.GoOld.BlessPid1)
	file.WriteInt64(log.GoOld.BlessPid2)
	file.WriteInt64(log.GoOld.BlessPid3)
	file.WriteInt64(log.GoOld.BlessPid4)
	file.WriteInt64(log.GoOld.BlessPid5)
}

func (log *PlayerTaoyuanBlessUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanBless.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.DailyBlessTimes)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastBlessTime)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.DailyBeBlessTimes)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastBeBlessTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BlessPid1)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BlessPid2)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BlessPid3)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BlessPid4)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.BlessPid5)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerTaoyuanBlessUpdateLog) CommitToTLog() {
}

func (log *PlayerTaoyuanBlessUpdateLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanBlessUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerTaoyuanBlessInsertLog) Rollback() {
	if log.db.tables.PlayerTaoyuanBless != log.Row {
		panic("insert rollback check failed 'PlayerTaoyuanBless'")
	}
	log.db.tables.PlayerTaoyuanBless = nil
	C.Free_PlayerTaoyuanBless(log.Row)
}

func (log *PlayerTaoyuanBlessDeleteLog) Rollback() {
	if log.db.tables.PlayerTaoyuanBless != nil {
		panic("delete rollback check failed 'player_taoyuan_bless'")
	}
	log.db.tables.PlayerTaoyuanBless = log.Old
}

func (log *PlayerTaoyuanBlessUpdateLog) Rollback() {
	if log.db.tables.PlayerTaoyuanBless != log.Row {
		panic("update rollback check failed 'player_taoyuan_bless'")
	}
	log.db.tables.PlayerTaoyuanBless = log.Old
	C.Free_PlayerTaoyuanBless(log.Row)
}

/* ========== player_taoyuan_fileds ========== */

type PlayerTaoyuanFiledsInsertLog struct {
	db *Database 
	Row *C.PlayerTaoyuanFileds
	GoRow *PlayerTaoyuanFileds
}

func (db *Database) newPlayerTaoyuanFiledsInsertLog(row *C.PlayerTaoyuanFileds, goRow *PlayerTaoyuanFileds) *PlayerTaoyuanFiledsInsertLog {
	return &PlayerTaoyuanFiledsInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerTaoyuanFiledsInsertLog) Free() {
}

func (log *PlayerTaoyuanFiledsInsertLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanFileds != nil {
		g_Hooks.PlayerTaoyuanFileds.Insert(&PlayerTaoyuanFiledsRow{row: log.Row})
	}
}

func (log *PlayerTaoyuanFiledsInsertLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanFiledsInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(128)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.FiledId)
	file.WriteInt16(log.GoRow.FiledStatus)
	file.WriteInt16(log.GoRow.PlantId)
	file.WriteInt64(log.GoRow.GrowTime)
}

func (log *PlayerTaoyuanFiledsInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanFileds.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.FiledId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.FiledStatus)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.PlantId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.GrowTime)))
	return stmt.Execute()
}

func (log *PlayerTaoyuanFiledsInsertLog) CommitToTLog() {
}

func (log *PlayerTaoyuanFiledsInsertLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanFiledsInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTaoyuanFiledsDeleteLog struct {
	db *Database
	Old *C.PlayerTaoyuanFileds
	GoOld *PlayerTaoyuanFileds
}

func (db *Database) newPlayerTaoyuanFiledsDeleteLog(old *C.PlayerTaoyuanFileds, goOld *PlayerTaoyuanFileds) *PlayerTaoyuanFiledsDeleteLog {
	return &PlayerTaoyuanFiledsDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerTaoyuanFiledsDeleteLog) Free() {
	C.Free_PlayerTaoyuanFileds(log.Old)
}

func (log *PlayerTaoyuanFiledsDeleteLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanFileds != nil {
		g_Hooks.PlayerTaoyuanFileds.Delete(&PlayerTaoyuanFiledsRow{row: log.Old})
	}
}

func (log *PlayerTaoyuanFiledsDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanFiledsDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(128)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.FiledId)
	file.WriteInt16(log.GoOld.FiledStatus)
	file.WriteInt16(log.GoOld.PlantId)
	file.WriteInt64(log.GoOld.GrowTime)
}

func (log *PlayerTaoyuanFiledsDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanFileds.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerTaoyuanFiledsDeleteLog) CommitToTLog() {
}

func (log *PlayerTaoyuanFiledsDeleteLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanFiledsDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTaoyuanFiledsUpdateLog struct {
	db *Database 
	Row *C.PlayerTaoyuanFileds
	Old *C.PlayerTaoyuanFileds
	GoNew *PlayerTaoyuanFileds
	GoOld *PlayerTaoyuanFileds
}

func (db *Database) newPlayerTaoyuanFiledsUpdateLog(row, old *C.PlayerTaoyuanFileds, goNew, goOld *PlayerTaoyuanFileds) *PlayerTaoyuanFiledsUpdateLog {
	return &PlayerTaoyuanFiledsUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerTaoyuanFiledsUpdateLog) Free() {
	C.Free_PlayerTaoyuanFileds(log.Old)
}

func (log *PlayerTaoyuanFiledsUpdateLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanFileds != nil {
		g_Hooks.PlayerTaoyuanFileds.Update(&PlayerTaoyuanFiledsRow{row: log.Row}, &PlayerTaoyuanFiledsRow{row: log.Old})
	}
}

func (log *PlayerTaoyuanFiledsUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanFiledsUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(128)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.FiledId)
	file.WriteInt16(log.GoNew.FiledStatus)
	file.WriteInt16(log.GoNew.PlantId)
	file.WriteInt64(log.GoNew.GrowTime)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.FiledId)
	file.WriteInt16(log.GoOld.FiledStatus)
	file.WriteInt16(log.GoOld.PlantId)
	file.WriteInt64(log.GoOld.GrowTime)
}

func (log *PlayerTaoyuanFiledsUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanFileds.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.FiledId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.FiledStatus)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.PlantId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.GrowTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerTaoyuanFiledsUpdateLog) CommitToTLog() {
}

func (log *PlayerTaoyuanFiledsUpdateLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanFiledsUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerTaoyuanFiledsInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerTaoyuanFileds
	for crow := log.db.tables.PlayerTaoyuanFileds; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerTaoyuanFileds'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerTaoyuanFileds = log.db.tables.PlayerTaoyuanFileds.next
	}
	C.Free_PlayerTaoyuanFileds(log.Row)
}

func (log *PlayerTaoyuanFiledsDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerTaoyuanFileds
	for crow := log.db.tables.PlayerTaoyuanFileds; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_taoyuan_fileds'")
	}
	log.Old.next = log.db.tables.PlayerTaoyuanFileds
	log.db.tables.PlayerTaoyuanFileds = log.Old
}

func (log *PlayerTaoyuanFiledsUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerTaoyuanFileds
	for crow := log.db.tables.PlayerTaoyuanFileds; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_taoyuan_fileds'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerTaoyuanFileds = log.Old
	}
	C.Free_PlayerTaoyuanFileds(log.Row)
}

/* ========== player_taoyuan_heart ========== */

type PlayerTaoyuanHeartInsertLog struct {
	db *Database 
	Row *C.PlayerTaoyuanHeart
	GoRow *PlayerTaoyuanHeart
}

func (db *Database) newPlayerTaoyuanHeartInsertLog(row *C.PlayerTaoyuanHeart, goRow *PlayerTaoyuanHeart) *PlayerTaoyuanHeartInsertLog {
	return &PlayerTaoyuanHeartInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerTaoyuanHeartInsertLog) Free() {
}

func (log *PlayerTaoyuanHeartInsertLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanHeart != nil {
		g_Hooks.PlayerTaoyuanHeart.Insert(&PlayerTaoyuanHeartRow{row: log.Row})
	}
}

func (log *PlayerTaoyuanHeartInsertLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanHeartInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(129)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.Level)
	file.WriteInt64(log.GoRow.Exp)
}

func (log *PlayerTaoyuanHeartInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanHeart.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Exp)))
	return stmt.Execute()
}

func (log *PlayerTaoyuanHeartInsertLog) CommitToTLog() {
}

func (log *PlayerTaoyuanHeartInsertLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanHeartInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTaoyuanHeartDeleteLog struct {
	db *Database
	Old *C.PlayerTaoyuanHeart
	GoOld *PlayerTaoyuanHeart
}

func (db *Database) newPlayerTaoyuanHeartDeleteLog(old *C.PlayerTaoyuanHeart, goOld *PlayerTaoyuanHeart) *PlayerTaoyuanHeartDeleteLog {
	return &PlayerTaoyuanHeartDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerTaoyuanHeartDeleteLog) Free() {
	C.Free_PlayerTaoyuanHeart(log.Old)
}

func (log *PlayerTaoyuanHeartDeleteLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanHeart != nil {
		g_Hooks.PlayerTaoyuanHeart.Delete(&PlayerTaoyuanHeartRow{row: log.Old})
	}
}

func (log *PlayerTaoyuanHeartDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanHeartDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(129)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.Level)
	file.WriteInt64(log.GoOld.Exp)
}

func (log *PlayerTaoyuanHeartDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanHeart.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerTaoyuanHeartDeleteLog) CommitToTLog() {
}

func (log *PlayerTaoyuanHeartDeleteLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanHeartDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTaoyuanHeartUpdateLog struct {
	db *Database 
	Row *C.PlayerTaoyuanHeart
	Old *C.PlayerTaoyuanHeart
	GoNew *PlayerTaoyuanHeart
	GoOld *PlayerTaoyuanHeart
}

func (db *Database) newPlayerTaoyuanHeartUpdateLog(row, old *C.PlayerTaoyuanHeart, goNew, goOld *PlayerTaoyuanHeart) *PlayerTaoyuanHeartUpdateLog {
	return &PlayerTaoyuanHeartUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerTaoyuanHeartUpdateLog) Free() {
	C.Free_PlayerTaoyuanHeart(log.Old)
}

func (log *PlayerTaoyuanHeartUpdateLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanHeart != nil {
		g_Hooks.PlayerTaoyuanHeart.Update(&PlayerTaoyuanHeartRow{row: log.Row}, &PlayerTaoyuanHeartRow{row: log.Old})
	}
}

func (log *PlayerTaoyuanHeartUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanHeartUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(129)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.Level)
	file.WriteInt64(log.GoNew.Exp)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.Level)
	file.WriteInt64(log.GoOld.Exp)
}

func (log *PlayerTaoyuanHeartUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanHeart.Update
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Exp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerTaoyuanHeartUpdateLog) CommitToTLog() {
}

func (log *PlayerTaoyuanHeartUpdateLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanHeartUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerTaoyuanHeartInsertLog) Rollback() {
	if log.db.tables.PlayerTaoyuanHeart != log.Row {
		panic("insert rollback check failed 'PlayerTaoyuanHeart'")
	}
	log.db.tables.PlayerTaoyuanHeart = nil
	C.Free_PlayerTaoyuanHeart(log.Row)
}

func (log *PlayerTaoyuanHeartDeleteLog) Rollback() {
	if log.db.tables.PlayerTaoyuanHeart != nil {
		panic("delete rollback check failed 'player_taoyuan_heart'")
	}
	log.db.tables.PlayerTaoyuanHeart = log.Old
}

func (log *PlayerTaoyuanHeartUpdateLog) Rollback() {
	if log.db.tables.PlayerTaoyuanHeart != log.Row {
		panic("update rollback check failed 'player_taoyuan_heart'")
	}
	log.db.tables.PlayerTaoyuanHeart = log.Old
	C.Free_PlayerTaoyuanHeart(log.Row)
}

/* ========== player_taoyuan_item ========== */

type PlayerTaoyuanItemInsertLog struct {
	db *Database 
	Row *C.PlayerTaoyuanItem
	GoRow *PlayerTaoyuanItem
}

func (db *Database) newPlayerTaoyuanItemInsertLog(row *C.PlayerTaoyuanItem, goRow *PlayerTaoyuanItem) *PlayerTaoyuanItemInsertLog {
	return &PlayerTaoyuanItemInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerTaoyuanItemInsertLog) Free() {
}

func (log *PlayerTaoyuanItemInsertLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanItem != nil {
		g_Hooks.PlayerTaoyuanItem.Insert(&PlayerTaoyuanItemRow{row: log.Row})
	}
}

func (log *PlayerTaoyuanItemInsertLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanItemInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(130)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.ItemId)
	file.WriteInt16(log.GoRow.Num)
}

func (log *PlayerTaoyuanItemInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanItem.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Num)))
	return stmt.Execute()
}

func (log *PlayerTaoyuanItemInsertLog) CommitToTLog() {
}

func (log *PlayerTaoyuanItemInsertLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanItemInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTaoyuanItemDeleteLog struct {
	db *Database
	Old *C.PlayerTaoyuanItem
	GoOld *PlayerTaoyuanItem
}

func (db *Database) newPlayerTaoyuanItemDeleteLog(old *C.PlayerTaoyuanItem, goOld *PlayerTaoyuanItem) *PlayerTaoyuanItemDeleteLog {
	return &PlayerTaoyuanItemDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerTaoyuanItemDeleteLog) Free() {
	C.Free_PlayerTaoyuanItem(log.Old)
}

func (log *PlayerTaoyuanItemDeleteLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanItem != nil {
		g_Hooks.PlayerTaoyuanItem.Delete(&PlayerTaoyuanItemRow{row: log.Old})
	}
}

func (log *PlayerTaoyuanItemDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanItemDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(130)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt16(log.GoOld.Num)
}

func (log *PlayerTaoyuanItemDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanItem.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerTaoyuanItemDeleteLog) CommitToTLog() {
}

func (log *PlayerTaoyuanItemDeleteLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanItemDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTaoyuanItemUpdateLog struct {
	db *Database 
	Row *C.PlayerTaoyuanItem
	Old *C.PlayerTaoyuanItem
	GoNew *PlayerTaoyuanItem
	GoOld *PlayerTaoyuanItem
}

func (db *Database) newPlayerTaoyuanItemUpdateLog(row, old *C.PlayerTaoyuanItem, goNew, goOld *PlayerTaoyuanItem) *PlayerTaoyuanItemUpdateLog {
	return &PlayerTaoyuanItemUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerTaoyuanItemUpdateLog) Free() {
	C.Free_PlayerTaoyuanItem(log.Old)
}

func (log *PlayerTaoyuanItemUpdateLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanItem != nil {
		g_Hooks.PlayerTaoyuanItem.Update(&PlayerTaoyuanItemRow{row: log.Row}, &PlayerTaoyuanItemRow{row: log.Old})
	}
}

func (log *PlayerTaoyuanItemUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanItemUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(130)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.ItemId)
	file.WriteInt16(log.GoNew.Num)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt16(log.GoOld.Num)
}

func (log *PlayerTaoyuanItemUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanItem.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Num)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerTaoyuanItemUpdateLog) CommitToTLog() {
}

func (log *PlayerTaoyuanItemUpdateLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanItemUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerTaoyuanItemInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerTaoyuanItem
	for crow := log.db.tables.PlayerTaoyuanItem; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerTaoyuanItem'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerTaoyuanItem = log.db.tables.PlayerTaoyuanItem.next
	}
	C.Free_PlayerTaoyuanItem(log.Row)
}

func (log *PlayerTaoyuanItemDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerTaoyuanItem
	for crow := log.db.tables.PlayerTaoyuanItem; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_taoyuan_item'")
	}
	log.Old.next = log.db.tables.PlayerTaoyuanItem
	log.db.tables.PlayerTaoyuanItem = log.Old
}

func (log *PlayerTaoyuanItemUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerTaoyuanItem
	for crow := log.db.tables.PlayerTaoyuanItem; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_taoyuan_item'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerTaoyuanItem = log.Old
	}
	C.Free_PlayerTaoyuanItem(log.Row)
}

/* ========== player_taoyuan_message ========== */

type PlayerTaoyuanMessageInsertLog struct {
	db *Database 
	Row *C.PlayerTaoyuanMessage
	GoRow *PlayerTaoyuanMessage
}

func (db *Database) newPlayerTaoyuanMessageInsertLog(row *C.PlayerTaoyuanMessage, goRow *PlayerTaoyuanMessage) *PlayerTaoyuanMessageInsertLog {
	return &PlayerTaoyuanMessageInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerTaoyuanMessageInsertLog) Free() {
}

func (log *PlayerTaoyuanMessageInsertLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanMessage != nil {
		g_Hooks.PlayerTaoyuanMessage.Insert(&PlayerTaoyuanMessageRow{row: log.Row})
	}
}

func (log *PlayerTaoyuanMessageInsertLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanMessageInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(131)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteString(log.GoRow.Nick)
	file.WriteInt32(log.GoRow.Exp)
}

func (log *PlayerTaoyuanMessageInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanMessage.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Nick), int(log.Row.Nick_len))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Exp)))
	return stmt.Execute()
}

func (log *PlayerTaoyuanMessageInsertLog) CommitToTLog() {
}

func (log *PlayerTaoyuanMessageInsertLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanMessageInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTaoyuanMessageDeleteLog struct {
	db *Database
	Old *C.PlayerTaoyuanMessage
	GoOld *PlayerTaoyuanMessage
}

func (db *Database) newPlayerTaoyuanMessageDeleteLog(old *C.PlayerTaoyuanMessage, goOld *PlayerTaoyuanMessage) *PlayerTaoyuanMessageDeleteLog {
	return &PlayerTaoyuanMessageDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerTaoyuanMessageDeleteLog) Free() {
	C.Free_PlayerTaoyuanMessage(log.Old)
}

func (log *PlayerTaoyuanMessageDeleteLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanMessage != nil {
		g_Hooks.PlayerTaoyuanMessage.Delete(&PlayerTaoyuanMessageRow{row: log.Old})
	}
}

func (log *PlayerTaoyuanMessageDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanMessageDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(131)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteString(log.GoOld.Nick)
	file.WriteInt32(log.GoOld.Exp)
}

func (log *PlayerTaoyuanMessageDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanMessage.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerTaoyuanMessageDeleteLog) CommitToTLog() {
}

func (log *PlayerTaoyuanMessageDeleteLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanMessageDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTaoyuanMessageUpdateLog struct {
	db *Database 
	Row *C.PlayerTaoyuanMessage
	Old *C.PlayerTaoyuanMessage
	GoNew *PlayerTaoyuanMessage
	GoOld *PlayerTaoyuanMessage
}

func (db *Database) newPlayerTaoyuanMessageUpdateLog(row, old *C.PlayerTaoyuanMessage, goNew, goOld *PlayerTaoyuanMessage) *PlayerTaoyuanMessageUpdateLog {
	return &PlayerTaoyuanMessageUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerTaoyuanMessageUpdateLog) Free() {
	C.Free_PlayerTaoyuanMessage(log.Old)
}

func (log *PlayerTaoyuanMessageUpdateLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanMessage != nil {
		g_Hooks.PlayerTaoyuanMessage.Update(&PlayerTaoyuanMessageRow{row: log.Row}, &PlayerTaoyuanMessageRow{row: log.Old})
	}
}

func (log *PlayerTaoyuanMessageUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanMessageUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(131)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteString(log.GoNew.Nick)
	file.WriteInt32(log.GoNew.Exp)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteString(log.GoOld.Nick)
	file.WriteInt32(log.GoOld.Exp)
}

func (log *PlayerTaoyuanMessageUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanMessage.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Nick), int(log.Row.Nick_len))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Exp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerTaoyuanMessageUpdateLog) CommitToTLog() {
}

func (log *PlayerTaoyuanMessageUpdateLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanMessageUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerTaoyuanMessageInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerTaoyuanMessage
	for crow := log.db.tables.PlayerTaoyuanMessage; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerTaoyuanMessage'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerTaoyuanMessage = log.db.tables.PlayerTaoyuanMessage.next
	}
	C.Free_PlayerTaoyuanMessage(log.Row)
}

func (log *PlayerTaoyuanMessageDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerTaoyuanMessage
	for crow := log.db.tables.PlayerTaoyuanMessage; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_taoyuan_message'")
	}
	log.Old.next = log.db.tables.PlayerTaoyuanMessage
	log.db.tables.PlayerTaoyuanMessage = log.Old
}

func (log *PlayerTaoyuanMessageUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerTaoyuanMessage
	for crow := log.db.tables.PlayerTaoyuanMessage; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_taoyuan_message'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerTaoyuanMessage = log.Old
	}
	C.Free_PlayerTaoyuanMessage(log.Row)
}

/* ========== player_taoyuan_product ========== */

type PlayerTaoyuanProductInsertLog struct {
	db *Database 
	Row *C.PlayerTaoyuanProduct
	GoRow *PlayerTaoyuanProduct
}

func (db *Database) newPlayerTaoyuanProductInsertLog(row *C.PlayerTaoyuanProduct, goRow *PlayerTaoyuanProduct) *PlayerTaoyuanProductInsertLog {
	return &PlayerTaoyuanProductInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerTaoyuanProductInsertLog) Free() {
}

func (log *PlayerTaoyuanProductInsertLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanProduct != nil {
		g_Hooks.PlayerTaoyuanProduct.Insert(&PlayerTaoyuanProductRow{row: log.Row})
	}
}

func (log *PlayerTaoyuanProductInsertLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanProductInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(132)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.SkillId1)
	file.WriteInt16(log.GoRow.SkillId2)
	file.WriteInt16(log.GoRow.SameTimeWineQueue)
	file.WriteInt64(log.GoRow.MakeWineTimes)
	file.WriteInt16(log.GoRow.SameTimeFoodQueue)
	file.WriteInt64(log.GoRow.MakeFoodTimes)
	file.WriteInt16(log.GoRow.FoodQueue1)
	file.WriteInt64(log.GoRow.FoodQueue1StartTimestamp)
	file.WriteInt64(log.GoRow.FoodQueue1EndTimestamp)
	file.WriteInt16(log.GoRow.FoodQueue2)
	file.WriteInt64(log.GoRow.FoodQueue2StartTimestamp)
	file.WriteInt64(log.GoRow.FoodQueue2EndTimestamp)
	file.WriteInt16(log.GoRow.FoodQueue3)
	file.WriteInt64(log.GoRow.FoodQueue3StartTimestamp)
	file.WriteInt64(log.GoRow.FoodQueue3EndTimestamp)
	file.WriteInt16(log.GoRow.FoodQueue4)
	file.WriteInt64(log.GoRow.FoodQueue4StartTimestamp)
	file.WriteInt64(log.GoRow.FoodQueue4EndTimestamp)
	file.WriteInt16(log.GoRow.WineQueue1)
	file.WriteInt64(log.GoRow.WineQueue1StartTimestamp)
	file.WriteInt64(log.GoRow.WineQueue1EndTimestamp)
	file.WriteInt16(log.GoRow.WineQueue2)
	file.WriteInt64(log.GoRow.WineQueue2StartTimestamp)
	file.WriteInt64(log.GoRow.WineQueue2EndTimestamp)
	file.WriteInt16(log.GoRow.WineQueue3)
	file.WriteInt64(log.GoRow.WineQueue3StartTimestamp)
	file.WriteInt64(log.GoRow.WineQueue3EndTimestamp)
	file.WriteInt16(log.GoRow.WineQueue4)
	file.WriteInt64(log.GoRow.WineQueue4StartTimestamp)
	file.WriteInt64(log.GoRow.WineQueue4EndTimestamp)
	file.WriteInt8(log.GoRow.IsOpenWine)
	file.WriteInt8(log.GoRow.IsOpenFood)
}

func (log *PlayerTaoyuanProductInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanProduct.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SkillId1)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SkillId2)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SameTimeWineQueue)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.MakeWineTimes)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SameTimeFoodQueue)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.MakeFoodTimes)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.FoodQueue1)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FoodQueue1StartTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FoodQueue1EndTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.FoodQueue2)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FoodQueue2StartTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FoodQueue2EndTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.FoodQueue3)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FoodQueue3StartTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FoodQueue3EndTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.FoodQueue4)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FoodQueue4StartTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FoodQueue4EndTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.WineQueue1)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WineQueue1StartTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WineQueue1EndTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.WineQueue2)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WineQueue2StartTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WineQueue2EndTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.WineQueue3)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WineQueue3StartTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WineQueue3EndTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.WineQueue4)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WineQueue4StartTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WineQueue4EndTimestamp)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.IsOpenWine)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.IsOpenFood)))
	return stmt.Execute()
}

func (log *PlayerTaoyuanProductInsertLog) CommitToTLog() {
}

func (log *PlayerTaoyuanProductInsertLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanProductInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTaoyuanProductDeleteLog struct {
	db *Database
	Old *C.PlayerTaoyuanProduct
	GoOld *PlayerTaoyuanProduct
}

func (db *Database) newPlayerTaoyuanProductDeleteLog(old *C.PlayerTaoyuanProduct, goOld *PlayerTaoyuanProduct) *PlayerTaoyuanProductDeleteLog {
	return &PlayerTaoyuanProductDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerTaoyuanProductDeleteLog) Free() {
	C.Free_PlayerTaoyuanProduct(log.Old)
}

func (log *PlayerTaoyuanProductDeleteLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanProduct != nil {
		g_Hooks.PlayerTaoyuanProduct.Delete(&PlayerTaoyuanProductRow{row: log.Old})
	}
}

func (log *PlayerTaoyuanProductDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanProductDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(132)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.SkillId1)
	file.WriteInt16(log.GoOld.SkillId2)
	file.WriteInt16(log.GoOld.SameTimeWineQueue)
	file.WriteInt64(log.GoOld.MakeWineTimes)
	file.WriteInt16(log.GoOld.SameTimeFoodQueue)
	file.WriteInt64(log.GoOld.MakeFoodTimes)
	file.WriteInt16(log.GoOld.FoodQueue1)
	file.WriteInt64(log.GoOld.FoodQueue1StartTimestamp)
	file.WriteInt64(log.GoOld.FoodQueue1EndTimestamp)
	file.WriteInt16(log.GoOld.FoodQueue2)
	file.WriteInt64(log.GoOld.FoodQueue2StartTimestamp)
	file.WriteInt64(log.GoOld.FoodQueue2EndTimestamp)
	file.WriteInt16(log.GoOld.FoodQueue3)
	file.WriteInt64(log.GoOld.FoodQueue3StartTimestamp)
	file.WriteInt64(log.GoOld.FoodQueue3EndTimestamp)
	file.WriteInt16(log.GoOld.FoodQueue4)
	file.WriteInt64(log.GoOld.FoodQueue4StartTimestamp)
	file.WriteInt64(log.GoOld.FoodQueue4EndTimestamp)
	file.WriteInt16(log.GoOld.WineQueue1)
	file.WriteInt64(log.GoOld.WineQueue1StartTimestamp)
	file.WriteInt64(log.GoOld.WineQueue1EndTimestamp)
	file.WriteInt16(log.GoOld.WineQueue2)
	file.WriteInt64(log.GoOld.WineQueue2StartTimestamp)
	file.WriteInt64(log.GoOld.WineQueue2EndTimestamp)
	file.WriteInt16(log.GoOld.WineQueue3)
	file.WriteInt64(log.GoOld.WineQueue3StartTimestamp)
	file.WriteInt64(log.GoOld.WineQueue3EndTimestamp)
	file.WriteInt16(log.GoOld.WineQueue4)
	file.WriteInt64(log.GoOld.WineQueue4StartTimestamp)
	file.WriteInt64(log.GoOld.WineQueue4EndTimestamp)
	file.WriteInt8(log.GoOld.IsOpenWine)
	file.WriteInt8(log.GoOld.IsOpenFood)
}

func (log *PlayerTaoyuanProductDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanProduct.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerTaoyuanProductDeleteLog) CommitToTLog() {
}

func (log *PlayerTaoyuanProductDeleteLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanProductDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTaoyuanProductUpdateLog struct {
	db *Database 
	Row *C.PlayerTaoyuanProduct
	Old *C.PlayerTaoyuanProduct
	GoNew *PlayerTaoyuanProduct
	GoOld *PlayerTaoyuanProduct
}

func (db *Database) newPlayerTaoyuanProductUpdateLog(row, old *C.PlayerTaoyuanProduct, goNew, goOld *PlayerTaoyuanProduct) *PlayerTaoyuanProductUpdateLog {
	return &PlayerTaoyuanProductUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerTaoyuanProductUpdateLog) Free() {
	C.Free_PlayerTaoyuanProduct(log.Old)
}

func (log *PlayerTaoyuanProductUpdateLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanProduct != nil {
		g_Hooks.PlayerTaoyuanProduct.Update(&PlayerTaoyuanProductRow{row: log.Row}, &PlayerTaoyuanProductRow{row: log.Old})
	}
}

func (log *PlayerTaoyuanProductUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanProductUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(132)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.SkillId1)
	file.WriteInt16(log.GoNew.SkillId2)
	file.WriteInt16(log.GoNew.SameTimeWineQueue)
	file.WriteInt64(log.GoNew.MakeWineTimes)
	file.WriteInt16(log.GoNew.SameTimeFoodQueue)
	file.WriteInt64(log.GoNew.MakeFoodTimes)
	file.WriteInt16(log.GoNew.FoodQueue1)
	file.WriteInt64(log.GoNew.FoodQueue1StartTimestamp)
	file.WriteInt64(log.GoNew.FoodQueue1EndTimestamp)
	file.WriteInt16(log.GoNew.FoodQueue2)
	file.WriteInt64(log.GoNew.FoodQueue2StartTimestamp)
	file.WriteInt64(log.GoNew.FoodQueue2EndTimestamp)
	file.WriteInt16(log.GoNew.FoodQueue3)
	file.WriteInt64(log.GoNew.FoodQueue3StartTimestamp)
	file.WriteInt64(log.GoNew.FoodQueue3EndTimestamp)
	file.WriteInt16(log.GoNew.FoodQueue4)
	file.WriteInt64(log.GoNew.FoodQueue4StartTimestamp)
	file.WriteInt64(log.GoNew.FoodQueue4EndTimestamp)
	file.WriteInt16(log.GoNew.WineQueue1)
	file.WriteInt64(log.GoNew.WineQueue1StartTimestamp)
	file.WriteInt64(log.GoNew.WineQueue1EndTimestamp)
	file.WriteInt16(log.GoNew.WineQueue2)
	file.WriteInt64(log.GoNew.WineQueue2StartTimestamp)
	file.WriteInt64(log.GoNew.WineQueue2EndTimestamp)
	file.WriteInt16(log.GoNew.WineQueue3)
	file.WriteInt64(log.GoNew.WineQueue3StartTimestamp)
	file.WriteInt64(log.GoNew.WineQueue3EndTimestamp)
	file.WriteInt16(log.GoNew.WineQueue4)
	file.WriteInt64(log.GoNew.WineQueue4StartTimestamp)
	file.WriteInt64(log.GoNew.WineQueue4EndTimestamp)
	file.WriteInt8(log.GoNew.IsOpenWine)
	file.WriteInt8(log.GoNew.IsOpenFood)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.SkillId1)
	file.WriteInt16(log.GoOld.SkillId2)
	file.WriteInt16(log.GoOld.SameTimeWineQueue)
	file.WriteInt64(log.GoOld.MakeWineTimes)
	file.WriteInt16(log.GoOld.SameTimeFoodQueue)
	file.WriteInt64(log.GoOld.MakeFoodTimes)
	file.WriteInt16(log.GoOld.FoodQueue1)
	file.WriteInt64(log.GoOld.FoodQueue1StartTimestamp)
	file.WriteInt64(log.GoOld.FoodQueue1EndTimestamp)
	file.WriteInt16(log.GoOld.FoodQueue2)
	file.WriteInt64(log.GoOld.FoodQueue2StartTimestamp)
	file.WriteInt64(log.GoOld.FoodQueue2EndTimestamp)
	file.WriteInt16(log.GoOld.FoodQueue3)
	file.WriteInt64(log.GoOld.FoodQueue3StartTimestamp)
	file.WriteInt64(log.GoOld.FoodQueue3EndTimestamp)
	file.WriteInt16(log.GoOld.FoodQueue4)
	file.WriteInt64(log.GoOld.FoodQueue4StartTimestamp)
	file.WriteInt64(log.GoOld.FoodQueue4EndTimestamp)
	file.WriteInt16(log.GoOld.WineQueue1)
	file.WriteInt64(log.GoOld.WineQueue1StartTimestamp)
	file.WriteInt64(log.GoOld.WineQueue1EndTimestamp)
	file.WriteInt16(log.GoOld.WineQueue2)
	file.WriteInt64(log.GoOld.WineQueue2StartTimestamp)
	file.WriteInt64(log.GoOld.WineQueue2EndTimestamp)
	file.WriteInt16(log.GoOld.WineQueue3)
	file.WriteInt64(log.GoOld.WineQueue3StartTimestamp)
	file.WriteInt64(log.GoOld.WineQueue3EndTimestamp)
	file.WriteInt16(log.GoOld.WineQueue4)
	file.WriteInt64(log.GoOld.WineQueue4StartTimestamp)
	file.WriteInt64(log.GoOld.WineQueue4EndTimestamp)
	file.WriteInt8(log.GoOld.IsOpenWine)
	file.WriteInt8(log.GoOld.IsOpenFood)
}

func (log *PlayerTaoyuanProductUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanProduct.Update
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SkillId1)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SkillId2)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SameTimeWineQueue)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.MakeWineTimes)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SameTimeFoodQueue)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.MakeFoodTimes)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.FoodQueue1)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FoodQueue1StartTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FoodQueue1EndTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.FoodQueue2)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FoodQueue2StartTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FoodQueue2EndTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.FoodQueue3)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FoodQueue3StartTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FoodQueue3EndTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.FoodQueue4)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FoodQueue4StartTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.FoodQueue4EndTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.WineQueue1)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WineQueue1StartTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WineQueue1EndTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.WineQueue2)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WineQueue2StartTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WineQueue2EndTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.WineQueue3)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WineQueue3StartTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WineQueue3EndTimestamp)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.WineQueue4)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WineQueue4StartTimestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.WineQueue4EndTimestamp)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.IsOpenWine)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.IsOpenFood)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerTaoyuanProductUpdateLog) CommitToTLog() {
}

func (log *PlayerTaoyuanProductUpdateLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanProductUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerTaoyuanProductInsertLog) Rollback() {
	if log.db.tables.PlayerTaoyuanProduct != log.Row {
		panic("insert rollback check failed 'PlayerTaoyuanProduct'")
	}
	log.db.tables.PlayerTaoyuanProduct = nil
	C.Free_PlayerTaoyuanProduct(log.Row)
}

func (log *PlayerTaoyuanProductDeleteLog) Rollback() {
	if log.db.tables.PlayerTaoyuanProduct != nil {
		panic("delete rollback check failed 'player_taoyuan_product'")
	}
	log.db.tables.PlayerTaoyuanProduct = log.Old
}

func (log *PlayerTaoyuanProductUpdateLog) Rollback() {
	if log.db.tables.PlayerTaoyuanProduct != log.Row {
		panic("update rollback check failed 'player_taoyuan_product'")
	}
	log.db.tables.PlayerTaoyuanProduct = log.Old
	C.Free_PlayerTaoyuanProduct(log.Row)
}

/* ========== player_taoyuan_purchase_record ========== */

type PlayerTaoyuanPurchaseRecordInsertLog struct {
	db *Database 
	Row *C.PlayerTaoyuanPurchaseRecord
	GoRow *PlayerTaoyuanPurchaseRecord
}

func (db *Database) newPlayerTaoyuanPurchaseRecordInsertLog(row *C.PlayerTaoyuanPurchaseRecord, goRow *PlayerTaoyuanPurchaseRecord) *PlayerTaoyuanPurchaseRecordInsertLog {
	return &PlayerTaoyuanPurchaseRecordInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerTaoyuanPurchaseRecordInsertLog) Free() {
}

func (log *PlayerTaoyuanPurchaseRecordInsertLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanPurchaseRecord != nil {
		g_Hooks.PlayerTaoyuanPurchaseRecord.Insert(&PlayerTaoyuanPurchaseRecordRow{row: log.Row})
	}
}

func (log *PlayerTaoyuanPurchaseRecordInsertLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanPurchaseRecordInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(133)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.ItemId)
	file.WriteInt16(log.GoRow.Num)
	file.WriteInt64(log.GoRow.Timestamp)
}

func (log *PlayerTaoyuanPurchaseRecordInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanPurchaseRecord.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Num)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	return stmt.Execute()
}

func (log *PlayerTaoyuanPurchaseRecordInsertLog) CommitToTLog() {
}

func (log *PlayerTaoyuanPurchaseRecordInsertLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanPurchaseRecordInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTaoyuanPurchaseRecordDeleteLog struct {
	db *Database
	Old *C.PlayerTaoyuanPurchaseRecord
	GoOld *PlayerTaoyuanPurchaseRecord
}

func (db *Database) newPlayerTaoyuanPurchaseRecordDeleteLog(old *C.PlayerTaoyuanPurchaseRecord, goOld *PlayerTaoyuanPurchaseRecord) *PlayerTaoyuanPurchaseRecordDeleteLog {
	return &PlayerTaoyuanPurchaseRecordDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerTaoyuanPurchaseRecordDeleteLog) Free() {
	C.Free_PlayerTaoyuanPurchaseRecord(log.Old)
}

func (log *PlayerTaoyuanPurchaseRecordDeleteLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanPurchaseRecord != nil {
		g_Hooks.PlayerTaoyuanPurchaseRecord.Delete(&PlayerTaoyuanPurchaseRecordRow{row: log.Old})
	}
}

func (log *PlayerTaoyuanPurchaseRecordDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanPurchaseRecordDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(133)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt16(log.GoOld.Num)
	file.WriteInt64(log.GoOld.Timestamp)
}

func (log *PlayerTaoyuanPurchaseRecordDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanPurchaseRecord.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerTaoyuanPurchaseRecordDeleteLog) CommitToTLog() {
}

func (log *PlayerTaoyuanPurchaseRecordDeleteLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanPurchaseRecordDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTaoyuanPurchaseRecordUpdateLog struct {
	db *Database 
	Row *C.PlayerTaoyuanPurchaseRecord
	Old *C.PlayerTaoyuanPurchaseRecord
	GoNew *PlayerTaoyuanPurchaseRecord
	GoOld *PlayerTaoyuanPurchaseRecord
}

func (db *Database) newPlayerTaoyuanPurchaseRecordUpdateLog(row, old *C.PlayerTaoyuanPurchaseRecord, goNew, goOld *PlayerTaoyuanPurchaseRecord) *PlayerTaoyuanPurchaseRecordUpdateLog {
	return &PlayerTaoyuanPurchaseRecordUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerTaoyuanPurchaseRecordUpdateLog) Free() {
	C.Free_PlayerTaoyuanPurchaseRecord(log.Old)
}

func (log *PlayerTaoyuanPurchaseRecordUpdateLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanPurchaseRecord != nil {
		g_Hooks.PlayerTaoyuanPurchaseRecord.Update(&PlayerTaoyuanPurchaseRecordRow{row: log.Row}, &PlayerTaoyuanPurchaseRecordRow{row: log.Old})
	}
}

func (log *PlayerTaoyuanPurchaseRecordUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanPurchaseRecordUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(133)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.ItemId)
	file.WriteInt16(log.GoNew.Num)
	file.WriteInt64(log.GoNew.Timestamp)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt16(log.GoOld.Num)
	file.WriteInt64(log.GoOld.Timestamp)
}

func (log *PlayerTaoyuanPurchaseRecordUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanPurchaseRecord.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Num)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Timestamp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerTaoyuanPurchaseRecordUpdateLog) CommitToTLog() {
}

func (log *PlayerTaoyuanPurchaseRecordUpdateLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanPurchaseRecordUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerTaoyuanPurchaseRecordInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerTaoyuanPurchaseRecord
	for crow := log.db.tables.PlayerTaoyuanPurchaseRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerTaoyuanPurchaseRecord'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerTaoyuanPurchaseRecord = log.db.tables.PlayerTaoyuanPurchaseRecord.next
	}
	C.Free_PlayerTaoyuanPurchaseRecord(log.Row)
}

func (log *PlayerTaoyuanPurchaseRecordDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerTaoyuanPurchaseRecord
	for crow := log.db.tables.PlayerTaoyuanPurchaseRecord; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_taoyuan_purchase_record'")
	}
	log.Old.next = log.db.tables.PlayerTaoyuanPurchaseRecord
	log.db.tables.PlayerTaoyuanPurchaseRecord = log.Old
}

func (log *PlayerTaoyuanPurchaseRecordUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerTaoyuanPurchaseRecord
	for crow := log.db.tables.PlayerTaoyuanPurchaseRecord; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_taoyuan_purchase_record'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerTaoyuanPurchaseRecord = log.Old
	}
	C.Free_PlayerTaoyuanPurchaseRecord(log.Row)
}

/* ========== player_taoyuan_quest ========== */

type PlayerTaoyuanQuestInsertLog struct {
	db *Database 
	Row *C.PlayerTaoyuanQuest
	GoRow *PlayerTaoyuanQuest
}

func (db *Database) newPlayerTaoyuanQuestInsertLog(row *C.PlayerTaoyuanQuest, goRow *PlayerTaoyuanQuest) *PlayerTaoyuanQuestInsertLog {
	return &PlayerTaoyuanQuestInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerTaoyuanQuestInsertLog) Free() {
}

func (log *PlayerTaoyuanQuestInsertLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanQuest != nil {
		g_Hooks.PlayerTaoyuanQuest.Insert(&PlayerTaoyuanQuestRow{row: log.Row})
	}
}

func (log *PlayerTaoyuanQuestInsertLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanQuestInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(134)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.LastQuestFlashTime)
	file.WriteInt16(log.GoRow.Quest1ItemId)
	file.WriteInt16(log.GoRow.Quest1ItemNum)
	file.WriteInt16(log.GoRow.Quest1TotalExp)
	file.WriteInt64(log.GoRow.Quest1TotalCoins)
	file.WriteInt64(log.GoRow.Quest1FinishTime)
	file.WriteInt16(log.GoRow.Quest2ItemId)
	file.WriteInt16(log.GoRow.Quest2ItemNum)
	file.WriteInt16(log.GoRow.Quest2TotalExp)
	file.WriteInt64(log.GoRow.Quest2TotalCoins)
	file.WriteInt64(log.GoRow.Quest2FinishTime)
	file.WriteInt16(log.GoRow.Quest3ItemId)
	file.WriteInt16(log.GoRow.Quest3ItemNum)
	file.WriteInt16(log.GoRow.Quest3TotalExp)
	file.WriteInt64(log.GoRow.Quest3TotalCoins)
	file.WriteInt64(log.GoRow.Quest3FinishTime)
	file.WriteInt16(log.GoRow.Quest4ItemId)
	file.WriteInt16(log.GoRow.Quest4ItemNum)
	file.WriteInt16(log.GoRow.Quest4TotalExp)
	file.WriteInt64(log.GoRow.Quest4TotalCoins)
	file.WriteInt64(log.GoRow.Quest4FinishTime)
	file.WriteInt16(log.GoRow.Quest5ItemId)
	file.WriteInt16(log.GoRow.Quest5ItemNum)
	file.WriteInt16(log.GoRow.Quest5TotalExp)
	file.WriteInt64(log.GoRow.Quest5TotalCoins)
	file.WriteInt64(log.GoRow.Quest5FinishTime)
}

func (log *PlayerTaoyuanQuestInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanQuest.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastQuestFlashTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest1ItemId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest1ItemNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest1TotalExp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Quest1TotalCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Quest1FinishTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest2ItemId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest2ItemNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest2TotalExp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Quest2TotalCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Quest2FinishTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest3ItemId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest3ItemNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest3TotalExp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Quest3TotalCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Quest3FinishTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest4ItemId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest4ItemNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest4TotalExp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Quest4TotalCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Quest4FinishTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest5ItemId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest5ItemNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest5TotalExp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Quest5TotalCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Quest5FinishTime)))
	return stmt.Execute()
}

func (log *PlayerTaoyuanQuestInsertLog) CommitToTLog() {
}

func (log *PlayerTaoyuanQuestInsertLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanQuestInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTaoyuanQuestDeleteLog struct {
	db *Database
	Old *C.PlayerTaoyuanQuest
	GoOld *PlayerTaoyuanQuest
}

func (db *Database) newPlayerTaoyuanQuestDeleteLog(old *C.PlayerTaoyuanQuest, goOld *PlayerTaoyuanQuest) *PlayerTaoyuanQuestDeleteLog {
	return &PlayerTaoyuanQuestDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerTaoyuanQuestDeleteLog) Free() {
	C.Free_PlayerTaoyuanQuest(log.Old)
}

func (log *PlayerTaoyuanQuestDeleteLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanQuest != nil {
		g_Hooks.PlayerTaoyuanQuest.Delete(&PlayerTaoyuanQuestRow{row: log.Old})
	}
}

func (log *PlayerTaoyuanQuestDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanQuestDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(134)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.LastQuestFlashTime)
	file.WriteInt16(log.GoOld.Quest1ItemId)
	file.WriteInt16(log.GoOld.Quest1ItemNum)
	file.WriteInt16(log.GoOld.Quest1TotalExp)
	file.WriteInt64(log.GoOld.Quest1TotalCoins)
	file.WriteInt64(log.GoOld.Quest1FinishTime)
	file.WriteInt16(log.GoOld.Quest2ItemId)
	file.WriteInt16(log.GoOld.Quest2ItemNum)
	file.WriteInt16(log.GoOld.Quest2TotalExp)
	file.WriteInt64(log.GoOld.Quest2TotalCoins)
	file.WriteInt64(log.GoOld.Quest2FinishTime)
	file.WriteInt16(log.GoOld.Quest3ItemId)
	file.WriteInt16(log.GoOld.Quest3ItemNum)
	file.WriteInt16(log.GoOld.Quest3TotalExp)
	file.WriteInt64(log.GoOld.Quest3TotalCoins)
	file.WriteInt64(log.GoOld.Quest3FinishTime)
	file.WriteInt16(log.GoOld.Quest4ItemId)
	file.WriteInt16(log.GoOld.Quest4ItemNum)
	file.WriteInt16(log.GoOld.Quest4TotalExp)
	file.WriteInt64(log.GoOld.Quest4TotalCoins)
	file.WriteInt64(log.GoOld.Quest4FinishTime)
	file.WriteInt16(log.GoOld.Quest5ItemId)
	file.WriteInt16(log.GoOld.Quest5ItemNum)
	file.WriteInt16(log.GoOld.Quest5TotalExp)
	file.WriteInt64(log.GoOld.Quest5TotalCoins)
	file.WriteInt64(log.GoOld.Quest5FinishTime)
}

func (log *PlayerTaoyuanQuestDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanQuest.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerTaoyuanQuestDeleteLog) CommitToTLog() {
}

func (log *PlayerTaoyuanQuestDeleteLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanQuestDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTaoyuanQuestUpdateLog struct {
	db *Database 
	Row *C.PlayerTaoyuanQuest
	Old *C.PlayerTaoyuanQuest
	GoNew *PlayerTaoyuanQuest
	GoOld *PlayerTaoyuanQuest
}

func (db *Database) newPlayerTaoyuanQuestUpdateLog(row, old *C.PlayerTaoyuanQuest, goNew, goOld *PlayerTaoyuanQuest) *PlayerTaoyuanQuestUpdateLog {
	return &PlayerTaoyuanQuestUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerTaoyuanQuestUpdateLog) Free() {
	C.Free_PlayerTaoyuanQuest(log.Old)
}

func (log *PlayerTaoyuanQuestUpdateLog) InvokeHook() {
	if g_Hooks.PlayerTaoyuanQuest != nil {
		g_Hooks.PlayerTaoyuanQuest.Update(&PlayerTaoyuanQuestRow{row: log.Row}, &PlayerTaoyuanQuestRow{row: log.Old})
	}
}

func (log *PlayerTaoyuanQuestUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerTaoyuanQuestUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(134)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.LastQuestFlashTime)
	file.WriteInt16(log.GoNew.Quest1ItemId)
	file.WriteInt16(log.GoNew.Quest1ItemNum)
	file.WriteInt16(log.GoNew.Quest1TotalExp)
	file.WriteInt64(log.GoNew.Quest1TotalCoins)
	file.WriteInt64(log.GoNew.Quest1FinishTime)
	file.WriteInt16(log.GoNew.Quest2ItemId)
	file.WriteInt16(log.GoNew.Quest2ItemNum)
	file.WriteInt16(log.GoNew.Quest2TotalExp)
	file.WriteInt64(log.GoNew.Quest2TotalCoins)
	file.WriteInt64(log.GoNew.Quest2FinishTime)
	file.WriteInt16(log.GoNew.Quest3ItemId)
	file.WriteInt16(log.GoNew.Quest3ItemNum)
	file.WriteInt16(log.GoNew.Quest3TotalExp)
	file.WriteInt64(log.GoNew.Quest3TotalCoins)
	file.WriteInt64(log.GoNew.Quest3FinishTime)
	file.WriteInt16(log.GoNew.Quest4ItemId)
	file.WriteInt16(log.GoNew.Quest4ItemNum)
	file.WriteInt16(log.GoNew.Quest4TotalExp)
	file.WriteInt64(log.GoNew.Quest4TotalCoins)
	file.WriteInt64(log.GoNew.Quest4FinishTime)
	file.WriteInt16(log.GoNew.Quest5ItemId)
	file.WriteInt16(log.GoNew.Quest5ItemNum)
	file.WriteInt16(log.GoNew.Quest5TotalExp)
	file.WriteInt64(log.GoNew.Quest5TotalCoins)
	file.WriteInt64(log.GoNew.Quest5FinishTime)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.LastQuestFlashTime)
	file.WriteInt16(log.GoOld.Quest1ItemId)
	file.WriteInt16(log.GoOld.Quest1ItemNum)
	file.WriteInt16(log.GoOld.Quest1TotalExp)
	file.WriteInt64(log.GoOld.Quest1TotalCoins)
	file.WriteInt64(log.GoOld.Quest1FinishTime)
	file.WriteInt16(log.GoOld.Quest2ItemId)
	file.WriteInt16(log.GoOld.Quest2ItemNum)
	file.WriteInt16(log.GoOld.Quest2TotalExp)
	file.WriteInt64(log.GoOld.Quest2TotalCoins)
	file.WriteInt64(log.GoOld.Quest2FinishTime)
	file.WriteInt16(log.GoOld.Quest3ItemId)
	file.WriteInt16(log.GoOld.Quest3ItemNum)
	file.WriteInt16(log.GoOld.Quest3TotalExp)
	file.WriteInt64(log.GoOld.Quest3TotalCoins)
	file.WriteInt64(log.GoOld.Quest3FinishTime)
	file.WriteInt16(log.GoOld.Quest4ItemId)
	file.WriteInt16(log.GoOld.Quest4ItemNum)
	file.WriteInt16(log.GoOld.Quest4TotalExp)
	file.WriteInt64(log.GoOld.Quest4TotalCoins)
	file.WriteInt64(log.GoOld.Quest4FinishTime)
	file.WriteInt16(log.GoOld.Quest5ItemId)
	file.WriteInt16(log.GoOld.Quest5ItemNum)
	file.WriteInt16(log.GoOld.Quest5TotalExp)
	file.WriteInt64(log.GoOld.Quest5TotalCoins)
	file.WriteInt64(log.GoOld.Quest5FinishTime)
}

func (log *PlayerTaoyuanQuestUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTaoyuanQuest.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastQuestFlashTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest1ItemId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest1ItemNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest1TotalExp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Quest1TotalCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Quest1FinishTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest2ItemId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest2ItemNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest2TotalExp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Quest2TotalCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Quest2FinishTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest3ItemId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest3ItemNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest3TotalExp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Quest3TotalCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Quest3FinishTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest4ItemId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest4ItemNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest4TotalExp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Quest4TotalCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Quest4FinishTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest5ItemId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest5ItemNum)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Quest5TotalExp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Quest5TotalCoins)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Quest5FinishTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerTaoyuanQuestUpdateLog) CommitToTLog() {
}

func (log *PlayerTaoyuanQuestUpdateLog) CommitToXdLog() {
}

func (log *PlayerTaoyuanQuestUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerTaoyuanQuestInsertLog) Rollback() {
	if log.db.tables.PlayerTaoyuanQuest != log.Row {
		panic("insert rollback check failed 'PlayerTaoyuanQuest'")
	}
	log.db.tables.PlayerTaoyuanQuest = nil
	C.Free_PlayerTaoyuanQuest(log.Row)
}

func (log *PlayerTaoyuanQuestDeleteLog) Rollback() {
	if log.db.tables.PlayerTaoyuanQuest != nil {
		panic("delete rollback check failed 'player_taoyuan_quest'")
	}
	log.db.tables.PlayerTaoyuanQuest = log.Old
}

func (log *PlayerTaoyuanQuestUpdateLog) Rollback() {
	if log.db.tables.PlayerTaoyuanQuest != log.Row {
		panic("update rollback check failed 'player_taoyuan_quest'")
	}
	log.db.tables.PlayerTaoyuanQuest = log.Old
	C.Free_PlayerTaoyuanQuest(log.Row)
}

/* ========== player_tb_xxd_roleinfo ========== */

type PlayerTbXxdRoleinfoInsertLog struct {
	db *Database 
	Row *C.PlayerTbXxdRoleinfo
	GoRow *PlayerTbXxdRoleinfo
}

func (db *Database) newPlayerTbXxdRoleinfoInsertLog(row *C.PlayerTbXxdRoleinfo, goRow *PlayerTbXxdRoleinfo) *PlayerTbXxdRoleinfoInsertLog {
	return &PlayerTbXxdRoleinfoInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerTbXxdRoleinfoInsertLog) Free() {
}

func (log *PlayerTbXxdRoleinfoInsertLog) InvokeHook() {
	if g_Hooks.PlayerTbXxdRoleinfo != nil {
		g_Hooks.PlayerTbXxdRoleinfo.Insert(&PlayerTbXxdRoleinfoRow{row: log.Row})
	}
}

func (log *PlayerTbXxdRoleinfoInsertLog) SetEventId(time.Time) {
}

func (log *PlayerTbXxdRoleinfoInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(135)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteString(log.GoRow.Gameappid)
	file.WriteString(log.GoRow.Openid)
	file.WriteInt64(log.GoRow.Regtime)
	file.WriteInt16(log.GoRow.Level)
	file.WriteInt16(log.GoRow.IFriends)
	file.WriteInt64(log.GoRow.Moneyios)
	file.WriteInt64(log.GoRow.Moneyandroid)
	file.WriteInt64(log.GoRow.Diamondios)
	file.WriteInt64(log.GoRow.Diamondandroid)
}

func (log *PlayerTbXxdRoleinfoInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTbXxdRoleinfo.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Gameappid), int(log.Row.Gameappid_len))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Openid), int(log.Row.Openid_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Regtime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.IFriends)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Moneyios)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Moneyandroid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Diamondios)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Diamondandroid)))
	return stmt.Execute()
}

func (log *PlayerTbXxdRoleinfoInsertLog) CommitToTLog() {
}

func (log *PlayerTbXxdRoleinfoInsertLog) CommitToXdLog() {
}

func (log *PlayerTbXxdRoleinfoInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTbXxdRoleinfoDeleteLog struct {
	db *Database
	Old *C.PlayerTbXxdRoleinfo
	GoOld *PlayerTbXxdRoleinfo
}

func (db *Database) newPlayerTbXxdRoleinfoDeleteLog(old *C.PlayerTbXxdRoleinfo, goOld *PlayerTbXxdRoleinfo) *PlayerTbXxdRoleinfoDeleteLog {
	return &PlayerTbXxdRoleinfoDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerTbXxdRoleinfoDeleteLog) Free() {
	C.Free_PlayerTbXxdRoleinfo(log.Old)
}

func (log *PlayerTbXxdRoleinfoDeleteLog) InvokeHook() {
	if g_Hooks.PlayerTbXxdRoleinfo != nil {
		g_Hooks.PlayerTbXxdRoleinfo.Delete(&PlayerTbXxdRoleinfoRow{row: log.Old})
	}
}

func (log *PlayerTbXxdRoleinfoDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerTbXxdRoleinfoDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(135)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteString(log.GoOld.Gameappid)
	file.WriteString(log.GoOld.Openid)
	file.WriteInt64(log.GoOld.Regtime)
	file.WriteInt16(log.GoOld.Level)
	file.WriteInt16(log.GoOld.IFriends)
	file.WriteInt64(log.GoOld.Moneyios)
	file.WriteInt64(log.GoOld.Moneyandroid)
	file.WriteInt64(log.GoOld.Diamondios)
	file.WriteInt64(log.GoOld.Diamondandroid)
}

func (log *PlayerTbXxdRoleinfoDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTbXxdRoleinfo.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerTbXxdRoleinfoDeleteLog) CommitToTLog() {
}

func (log *PlayerTbXxdRoleinfoDeleteLog) CommitToXdLog() {
}

func (log *PlayerTbXxdRoleinfoDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTbXxdRoleinfoUpdateLog struct {
	db *Database 
	Row *C.PlayerTbXxdRoleinfo
	Old *C.PlayerTbXxdRoleinfo
	GoNew *PlayerTbXxdRoleinfo
	GoOld *PlayerTbXxdRoleinfo
}

func (db *Database) newPlayerTbXxdRoleinfoUpdateLog(row, old *C.PlayerTbXxdRoleinfo, goNew, goOld *PlayerTbXxdRoleinfo) *PlayerTbXxdRoleinfoUpdateLog {
	return &PlayerTbXxdRoleinfoUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerTbXxdRoleinfoUpdateLog) Free() {
	C.Free_PlayerTbXxdRoleinfo(log.Old)
}

func (log *PlayerTbXxdRoleinfoUpdateLog) InvokeHook() {
	if g_Hooks.PlayerTbXxdRoleinfo != nil {
		g_Hooks.PlayerTbXxdRoleinfo.Update(&PlayerTbXxdRoleinfoRow{row: log.Row}, &PlayerTbXxdRoleinfoRow{row: log.Old})
	}
}

func (log *PlayerTbXxdRoleinfoUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerTbXxdRoleinfoUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(135)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteString(log.GoNew.Gameappid)
	file.WriteString(log.GoNew.Openid)
	file.WriteInt64(log.GoNew.Regtime)
	file.WriteInt16(log.GoNew.Level)
	file.WriteInt16(log.GoNew.IFriends)
	file.WriteInt64(log.GoNew.Moneyios)
	file.WriteInt64(log.GoNew.Moneyandroid)
	file.WriteInt64(log.GoNew.Diamondios)
	file.WriteInt64(log.GoNew.Diamondandroid)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteString(log.GoOld.Gameappid)
	file.WriteString(log.GoOld.Openid)
	file.WriteInt64(log.GoOld.Regtime)
	file.WriteInt16(log.GoOld.Level)
	file.WriteInt16(log.GoOld.IFriends)
	file.WriteInt64(log.GoOld.Moneyios)
	file.WriteInt64(log.GoOld.Moneyandroid)
	file.WriteInt64(log.GoOld.Diamondios)
	file.WriteInt64(log.GoOld.Diamondandroid)
}

func (log *PlayerTbXxdRoleinfoUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTbXxdRoleinfo.Update
	stmt.BindVarchar(unsafe.Pointer(log.Row.Gameappid), int(log.Row.Gameappid_len))
	stmt.BindVarchar(unsafe.Pointer(log.Row.Openid), int(log.Row.Openid_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Regtime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.IFriends)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Moneyios)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Moneyandroid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Diamondios)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Diamondandroid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerTbXxdRoleinfoUpdateLog) CommitToTLog() {
}

func (log *PlayerTbXxdRoleinfoUpdateLog) CommitToXdLog() {
}

func (log *PlayerTbXxdRoleinfoUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerTbXxdRoleinfoInsertLog) Rollback() {
	if log.db.tables.PlayerTbXxdRoleinfo != log.Row {
		panic("insert rollback check failed 'PlayerTbXxdRoleinfo'")
	}
	log.db.tables.PlayerTbXxdRoleinfo = nil
	C.Free_PlayerTbXxdRoleinfo(log.Row)
}

func (log *PlayerTbXxdRoleinfoDeleteLog) Rollback() {
	if log.db.tables.PlayerTbXxdRoleinfo != nil {
		panic("delete rollback check failed 'player_tb_xxd_roleinfo'")
	}
	log.db.tables.PlayerTbXxdRoleinfo = log.Old
}

func (log *PlayerTbXxdRoleinfoUpdateLog) Rollback() {
	if log.db.tables.PlayerTbXxdRoleinfo != log.Row {
		panic("update rollback check failed 'player_tb_xxd_roleinfo'")
	}
	log.db.tables.PlayerTbXxdRoleinfo = log.Old
	C.Free_PlayerTbXxdRoleinfo(log.Row)
}

/* ========== player_team_info ========== */

type PlayerTeamInfoInsertLog struct {
	db *Database 
	Row *C.PlayerTeamInfo
	GoRow *PlayerTeamInfo
}

func (db *Database) newPlayerTeamInfoInsertLog(row *C.PlayerTeamInfo, goRow *PlayerTeamInfo) *PlayerTeamInfoInsertLog {
	return &PlayerTeamInfoInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerTeamInfoInsertLog) Free() {
}

func (log *PlayerTeamInfoInsertLog) InvokeHook() {
	if g_Hooks.PlayerTeamInfo != nil {
		g_Hooks.PlayerTeamInfo.Insert(&PlayerTeamInfoRow{row: log.Row})
	}
}

func (log *PlayerTeamInfoInsertLog) SetEventId(time.Time) {
}

func (log *PlayerTeamInfoInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(136)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt32(log.GoRow.Relationship)
	file.WriteInt16(log.GoRow.HealthLv)
	file.WriteInt16(log.GoRow.AttackLv)
	file.WriteInt16(log.GoRow.DefenceLv)
}

func (log *PlayerTeamInfoInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTeamInfo.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Relationship)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.HealthLv)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AttackLv)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DefenceLv)))
	return stmt.Execute()
}

func (log *PlayerTeamInfoInsertLog) CommitToTLog() {
}

func (log *PlayerTeamInfoInsertLog) CommitToXdLog() {
}

func (log *PlayerTeamInfoInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTeamInfoDeleteLog struct {
	db *Database
	Old *C.PlayerTeamInfo
	GoOld *PlayerTeamInfo
}

func (db *Database) newPlayerTeamInfoDeleteLog(old *C.PlayerTeamInfo, goOld *PlayerTeamInfo) *PlayerTeamInfoDeleteLog {
	return &PlayerTeamInfoDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerTeamInfoDeleteLog) Free() {
	C.Free_PlayerTeamInfo(log.Old)
}

func (log *PlayerTeamInfoDeleteLog) InvokeHook() {
	if g_Hooks.PlayerTeamInfo != nil {
		g_Hooks.PlayerTeamInfo.Delete(&PlayerTeamInfoRow{row: log.Old})
	}
}

func (log *PlayerTeamInfoDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerTeamInfoDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(136)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Relationship)
	file.WriteInt16(log.GoOld.HealthLv)
	file.WriteInt16(log.GoOld.AttackLv)
	file.WriteInt16(log.GoOld.DefenceLv)
}

func (log *PlayerTeamInfoDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTeamInfo.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerTeamInfoDeleteLog) CommitToTLog() {
}

func (log *PlayerTeamInfoDeleteLog) CommitToXdLog() {
}

func (log *PlayerTeamInfoDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTeamInfoUpdateLog struct {
	db *Database 
	Row *C.PlayerTeamInfo
	Old *C.PlayerTeamInfo
	GoNew *PlayerTeamInfo
	GoOld *PlayerTeamInfo
}

func (db *Database) newPlayerTeamInfoUpdateLog(row, old *C.PlayerTeamInfo, goNew, goOld *PlayerTeamInfo) *PlayerTeamInfoUpdateLog {
	return &PlayerTeamInfoUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerTeamInfoUpdateLog) Free() {
	C.Free_PlayerTeamInfo(log.Old)
}

func (log *PlayerTeamInfoUpdateLog) InvokeHook() {
	if g_Hooks.PlayerTeamInfo != nil {
		g_Hooks.PlayerTeamInfo.Update(&PlayerTeamInfoRow{row: log.Row}, &PlayerTeamInfoRow{row: log.Old})
	}
}

func (log *PlayerTeamInfoUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerTeamInfoUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(136)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt32(log.GoNew.Relationship)
	file.WriteInt16(log.GoNew.HealthLv)
	file.WriteInt16(log.GoNew.AttackLv)
	file.WriteInt16(log.GoNew.DefenceLv)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt32(log.GoOld.Relationship)
	file.WriteInt16(log.GoOld.HealthLv)
	file.WriteInt16(log.GoOld.AttackLv)
	file.WriteInt16(log.GoOld.DefenceLv)
}

func (log *PlayerTeamInfoUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTeamInfo.Update
	stmt.BindInt(unsafe.Pointer(&(log.Row.Relationship)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.HealthLv)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AttackLv)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.DefenceLv)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerTeamInfoUpdateLog) CommitToTLog() {
}

func (log *PlayerTeamInfoUpdateLog) CommitToXdLog() {
}

func (log *PlayerTeamInfoUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerTeamInfoInsertLog) Rollback() {
	if log.db.tables.PlayerTeamInfo != log.Row {
		panic("insert rollback check failed 'PlayerTeamInfo'")
	}
	log.db.tables.PlayerTeamInfo = nil
	C.Free_PlayerTeamInfo(log.Row)
}

func (log *PlayerTeamInfoDeleteLog) Rollback() {
	if log.db.tables.PlayerTeamInfo != nil {
		panic("delete rollback check failed 'player_team_info'")
	}
	log.db.tables.PlayerTeamInfo = log.Old
}

func (log *PlayerTeamInfoUpdateLog) Rollback() {
	if log.db.tables.PlayerTeamInfo != log.Row {
		panic("update rollback check failed 'player_team_info'")
	}
	log.db.tables.PlayerTeamInfo = log.Old
	C.Free_PlayerTeamInfo(log.Row)
}

/* ========== player_totem ========== */

type PlayerTotemInsertLog struct {
	db *Database 
	Row *C.PlayerTotem
	GoRow *PlayerTotem
}

func (db *Database) newPlayerTotemInsertLog(row *C.PlayerTotem, goRow *PlayerTotem) *PlayerTotemInsertLog {
	return &PlayerTotemInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerTotemInsertLog) Free() {
}

func (log *PlayerTotemInsertLog) InvokeHook() {
	if g_Hooks.PlayerTotem != nil {
		g_Hooks.PlayerTotem.Insert(&PlayerTotemRow{row: log.Row})
	}
}

func (log *PlayerTotemInsertLog) SetEventId(time.Time) {
}

func (log *PlayerTotemInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(137)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.TotemId)
	file.WriteInt8(log.GoRow.Level)
	file.WriteInt16(log.GoRow.SkillId)
}

func (log *PlayerTotemInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTotem.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TotemId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SkillId)))
	return stmt.Execute()
}

func (log *PlayerTotemInsertLog) CommitToTLog() {
}

func (log *PlayerTotemInsertLog) CommitToXdLog() {
}

func (log *PlayerTotemInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTotemDeleteLog struct {
	db *Database
	Old *C.PlayerTotem
	GoOld *PlayerTotem
}

func (db *Database) newPlayerTotemDeleteLog(old *C.PlayerTotem, goOld *PlayerTotem) *PlayerTotemDeleteLog {
	return &PlayerTotemDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerTotemDeleteLog) Free() {
	C.Free_PlayerTotem(log.Old)
}

func (log *PlayerTotemDeleteLog) InvokeHook() {
	if g_Hooks.PlayerTotem != nil {
		g_Hooks.PlayerTotem.Delete(&PlayerTotemRow{row: log.Old})
	}
}

func (log *PlayerTotemDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerTotemDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(137)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.TotemId)
	file.WriteInt8(log.GoOld.Level)
	file.WriteInt16(log.GoOld.SkillId)
}

func (log *PlayerTotemDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTotem.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerTotemDeleteLog) CommitToTLog() {
}

func (log *PlayerTotemDeleteLog) CommitToXdLog() {
}

func (log *PlayerTotemDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTotemUpdateLog struct {
	db *Database 
	Row *C.PlayerTotem
	Old *C.PlayerTotem
	GoNew *PlayerTotem
	GoOld *PlayerTotem
}

func (db *Database) newPlayerTotemUpdateLog(row, old *C.PlayerTotem, goNew, goOld *PlayerTotem) *PlayerTotemUpdateLog {
	return &PlayerTotemUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerTotemUpdateLog) Free() {
	C.Free_PlayerTotem(log.Old)
}

func (log *PlayerTotemUpdateLog) InvokeHook() {
	if g_Hooks.PlayerTotem != nil {
		g_Hooks.PlayerTotem.Update(&PlayerTotemRow{row: log.Row}, &PlayerTotemRow{row: log.Old})
	}
}

func (log *PlayerTotemUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerTotemUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(137)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.TotemId)
	file.WriteInt8(log.GoNew.Level)
	file.WriteInt16(log.GoNew.SkillId)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.TotemId)
	file.WriteInt8(log.GoOld.Level)
	file.WriteInt16(log.GoOld.SkillId)
}

func (log *PlayerTotemUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTotem.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TotemId)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SkillId)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerTotemUpdateLog) CommitToTLog() {
}

func (log *PlayerTotemUpdateLog) CommitToXdLog() {
}

func (log *PlayerTotemUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerTotemInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerTotem
	for crow := log.db.tables.PlayerTotem; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerTotem'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerTotem = log.db.tables.PlayerTotem.next
	}
	C.Free_PlayerTotem(log.Row)
}

func (log *PlayerTotemDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerTotem
	for crow := log.db.tables.PlayerTotem; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_totem'")
	}
	log.Old.next = log.db.tables.PlayerTotem
	log.db.tables.PlayerTotem = log.Old
}

func (log *PlayerTotemUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerTotem
	for crow := log.db.tables.PlayerTotem; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_totem'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerTotem = log.Old
	}
	C.Free_PlayerTotem(log.Row)
}

/* ========== player_totem_info ========== */

type PlayerTotemInfoInsertLog struct {
	db *Database 
	Row *C.PlayerTotemInfo
	GoRow *PlayerTotemInfo
}

func (db *Database) newPlayerTotemInfoInsertLog(row *C.PlayerTotemInfo, goRow *PlayerTotemInfo) *PlayerTotemInfoInsertLog {
	return &PlayerTotemInfoInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerTotemInfoInsertLog) Free() {
}

func (log *PlayerTotemInfoInsertLog) InvokeHook() {
	if g_Hooks.PlayerTotemInfo != nil {
		g_Hooks.PlayerTotemInfo.Insert(&PlayerTotemInfoRow{row: log.Row})
	}
}

func (log *PlayerTotemInfoInsertLog) SetEventId(time.Time) {
}

func (log *PlayerTotemInfoInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(138)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.IngotCallDailyNum)
	file.WriteInt64(log.GoRow.IngotCallTimestamp)
	file.WriteInt32(log.GoRow.RockRuneNum)
	file.WriteInt32(log.GoRow.JadeRuneNum)
	file.WriteInt64(log.GoRow.Pos1)
	file.WriteInt64(log.GoRow.Pos2)
	file.WriteInt64(log.GoRow.Pos3)
	file.WriteInt64(log.GoRow.Pos4)
	file.WriteInt64(log.GoRow.Pos5)
	file.WriteInt16(log.GoRow.IngotDrawTimes)
}

func (log *PlayerTotemInfoInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTotemInfo.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.IngotCallDailyNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.IngotCallTimestamp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.RockRuneNum)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.JadeRuneNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos1)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos2)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos3)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos4)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos5)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.IngotDrawTimes)))
	return stmt.Execute()
}

func (log *PlayerTotemInfoInsertLog) CommitToTLog() {
}

func (log *PlayerTotemInfoInsertLog) CommitToXdLog() {
}

func (log *PlayerTotemInfoInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTotemInfoDeleteLog struct {
	db *Database
	Old *C.PlayerTotemInfo
	GoOld *PlayerTotemInfo
}

func (db *Database) newPlayerTotemInfoDeleteLog(old *C.PlayerTotemInfo, goOld *PlayerTotemInfo) *PlayerTotemInfoDeleteLog {
	return &PlayerTotemInfoDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerTotemInfoDeleteLog) Free() {
	C.Free_PlayerTotemInfo(log.Old)
}

func (log *PlayerTotemInfoDeleteLog) InvokeHook() {
	if g_Hooks.PlayerTotemInfo != nil {
		g_Hooks.PlayerTotemInfo.Delete(&PlayerTotemInfoRow{row: log.Old})
	}
}

func (log *PlayerTotemInfoDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerTotemInfoDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(138)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.IngotCallDailyNum)
	file.WriteInt64(log.GoOld.IngotCallTimestamp)
	file.WriteInt32(log.GoOld.RockRuneNum)
	file.WriteInt32(log.GoOld.JadeRuneNum)
	file.WriteInt64(log.GoOld.Pos1)
	file.WriteInt64(log.GoOld.Pos2)
	file.WriteInt64(log.GoOld.Pos3)
	file.WriteInt64(log.GoOld.Pos4)
	file.WriteInt64(log.GoOld.Pos5)
	file.WriteInt16(log.GoOld.IngotDrawTimes)
}

func (log *PlayerTotemInfoDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTotemInfo.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerTotemInfoDeleteLog) CommitToTLog() {
}

func (log *PlayerTotemInfoDeleteLog) CommitToXdLog() {
}

func (log *PlayerTotemInfoDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTotemInfoUpdateLog struct {
	db *Database 
	Row *C.PlayerTotemInfo
	Old *C.PlayerTotemInfo
	GoNew *PlayerTotemInfo
	GoOld *PlayerTotemInfo
}

func (db *Database) newPlayerTotemInfoUpdateLog(row, old *C.PlayerTotemInfo, goNew, goOld *PlayerTotemInfo) *PlayerTotemInfoUpdateLog {
	return &PlayerTotemInfoUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerTotemInfoUpdateLog) Free() {
	C.Free_PlayerTotemInfo(log.Old)
}

func (log *PlayerTotemInfoUpdateLog) InvokeHook() {
	if g_Hooks.PlayerTotemInfo != nil {
		g_Hooks.PlayerTotemInfo.Update(&PlayerTotemInfoRow{row: log.Row}, &PlayerTotemInfoRow{row: log.Old})
	}
}

func (log *PlayerTotemInfoUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerTotemInfoUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(138)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.IngotCallDailyNum)
	file.WriteInt64(log.GoNew.IngotCallTimestamp)
	file.WriteInt32(log.GoNew.RockRuneNum)
	file.WriteInt32(log.GoNew.JadeRuneNum)
	file.WriteInt64(log.GoNew.Pos1)
	file.WriteInt64(log.GoNew.Pos2)
	file.WriteInt64(log.GoNew.Pos3)
	file.WriteInt64(log.GoNew.Pos4)
	file.WriteInt64(log.GoNew.Pos5)
	file.WriteInt16(log.GoNew.IngotDrawTimes)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.IngotCallDailyNum)
	file.WriteInt64(log.GoOld.IngotCallTimestamp)
	file.WriteInt32(log.GoOld.RockRuneNum)
	file.WriteInt32(log.GoOld.JadeRuneNum)
	file.WriteInt64(log.GoOld.Pos1)
	file.WriteInt64(log.GoOld.Pos2)
	file.WriteInt64(log.GoOld.Pos3)
	file.WriteInt64(log.GoOld.Pos4)
	file.WriteInt64(log.GoOld.Pos5)
	file.WriteInt16(log.GoOld.IngotDrawTimes)
}

func (log *PlayerTotemInfoUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTotemInfo.Update
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.IngotCallDailyNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.IngotCallTimestamp)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.RockRuneNum)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.JadeRuneNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos1)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos2)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos3)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos4)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pos5)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.IngotDrawTimes)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerTotemInfoUpdateLog) CommitToTLog() {
}

func (log *PlayerTotemInfoUpdateLog) CommitToXdLog() {
}

func (log *PlayerTotemInfoUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerTotemInfoInsertLog) Rollback() {
	if log.db.tables.PlayerTotemInfo != log.Row {
		panic("insert rollback check failed 'PlayerTotemInfo'")
	}
	log.db.tables.PlayerTotemInfo = nil
	C.Free_PlayerTotemInfo(log.Row)
}

func (log *PlayerTotemInfoDeleteLog) Rollback() {
	if log.db.tables.PlayerTotemInfo != nil {
		panic("delete rollback check failed 'player_totem_info'")
	}
	log.db.tables.PlayerTotemInfo = log.Old
}

func (log *PlayerTotemInfoUpdateLog) Rollback() {
	if log.db.tables.PlayerTotemInfo != log.Row {
		panic("update rollback check failed 'player_totem_info'")
	}
	log.db.tables.PlayerTotemInfo = log.Old
	C.Free_PlayerTotemInfo(log.Row)
}

/* ========== player_town ========== */

type PlayerTownInsertLog struct {
	db *Database 
	Row *C.PlayerTown
	GoRow *PlayerTown
}

func (db *Database) newPlayerTownInsertLog(row *C.PlayerTown, goRow *PlayerTown) *PlayerTownInsertLog {
	return &PlayerTownInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerTownInsertLog) Free() {
}

func (log *PlayerTownInsertLog) InvokeHook() {
	if g_Hooks.PlayerTown != nil {
		g_Hooks.PlayerTown.Insert(&PlayerTownRow{row: log.Row})
	}
}

func (log *PlayerTownInsertLog) SetEventId(time.Time) {
}

func (log *PlayerTownInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(139)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.TownId)
	file.WriteInt32(log.GoRow.Lock)
	file.WriteInt16(log.GoRow.AtPosX)
	file.WriteInt16(log.GoRow.AtPosY)
}

func (log *PlayerTownInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTown.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TownId)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Lock)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AtPosX)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AtPosY)))
	return stmt.Execute()
}

func (log *PlayerTownInsertLog) CommitToTLog() {
}

func (log *PlayerTownInsertLog) CommitToXdLog() {
}

func (log *PlayerTownInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTownDeleteLog struct {
	db *Database
	Old *C.PlayerTown
	GoOld *PlayerTown
}

func (db *Database) newPlayerTownDeleteLog(old *C.PlayerTown, goOld *PlayerTown) *PlayerTownDeleteLog {
	return &PlayerTownDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerTownDeleteLog) Free() {
	C.Free_PlayerTown(log.Old)
}

func (log *PlayerTownDeleteLog) InvokeHook() {
	if g_Hooks.PlayerTown != nil {
		g_Hooks.PlayerTown.Delete(&PlayerTownRow{row: log.Old})
	}
}

func (log *PlayerTownDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerTownDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(139)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.TownId)
	file.WriteInt32(log.GoOld.Lock)
	file.WriteInt16(log.GoOld.AtPosX)
	file.WriteInt16(log.GoOld.AtPosY)
}

func (log *PlayerTownDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTown.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerTownDeleteLog) CommitToTLog() {
}

func (log *PlayerTownDeleteLog) CommitToXdLog() {
}

func (log *PlayerTownDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTownUpdateLog struct {
	db *Database 
	Row *C.PlayerTown
	Old *C.PlayerTown
	GoNew *PlayerTown
	GoOld *PlayerTown
}

func (db *Database) newPlayerTownUpdateLog(row, old *C.PlayerTown, goNew, goOld *PlayerTown) *PlayerTownUpdateLog {
	return &PlayerTownUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerTownUpdateLog) Free() {
	C.Free_PlayerTown(log.Old)
}

func (log *PlayerTownUpdateLog) InvokeHook() {
	if g_Hooks.PlayerTown != nil {
		g_Hooks.PlayerTown.Update(&PlayerTownRow{row: log.Row}, &PlayerTownRow{row: log.Old})
	}
}

func (log *PlayerTownUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerTownUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(139)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.TownId)
	file.WriteInt32(log.GoNew.Lock)
	file.WriteInt16(log.GoNew.AtPosX)
	file.WriteInt16(log.GoNew.AtPosY)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.TownId)
	file.WriteInt32(log.GoOld.Lock)
	file.WriteInt16(log.GoOld.AtPosX)
	file.WriteInt16(log.GoOld.AtPosY)
}

func (log *PlayerTownUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTown.Update
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TownId)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.Lock)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AtPosX)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.AtPosY)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerTownUpdateLog) CommitToTLog() {
}

func (log *PlayerTownUpdateLog) CommitToXdLog() {
}

func (log *PlayerTownUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerTownInsertLog) Rollback() {
	if log.db.tables.PlayerTown != log.Row {
		panic("insert rollback check failed 'PlayerTown'")
	}
	log.db.tables.PlayerTown = nil
	C.Free_PlayerTown(log.Row)
}

func (log *PlayerTownDeleteLog) Rollback() {
	if log.db.tables.PlayerTown != nil {
		panic("delete rollback check failed 'player_town'")
	}
	log.db.tables.PlayerTown = log.Old
}

func (log *PlayerTownUpdateLog) Rollback() {
	if log.db.tables.PlayerTown != log.Row {
		panic("update rollback check failed 'player_town'")
	}
	log.db.tables.PlayerTown = log.Old
	C.Free_PlayerTown(log.Row)
}

/* ========== player_trader_refresh_state ========== */

type PlayerTraderRefreshStateInsertLog struct {
	db *Database 
	Row *C.PlayerTraderRefreshState
	GoRow *PlayerTraderRefreshState
}

func (db *Database) newPlayerTraderRefreshStateInsertLog(row *C.PlayerTraderRefreshState, goRow *PlayerTraderRefreshState) *PlayerTraderRefreshStateInsertLog {
	return &PlayerTraderRefreshStateInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerTraderRefreshStateInsertLog) Free() {
}

func (log *PlayerTraderRefreshStateInsertLog) InvokeHook() {
	if g_Hooks.PlayerTraderRefreshState != nil {
		g_Hooks.PlayerTraderRefreshState.Insert(&PlayerTraderRefreshStateRow{row: log.Row})
	}
}

func (log *PlayerTraderRefreshStateInsertLog) SetEventId(time.Time) {
}

func (log *PlayerTraderRefreshStateInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(140)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.LastUpdateTime)
	file.WriteInt64(log.GoRow.AutoUpdateTime)
	file.WriteInt16(log.GoRow.TraderId)
	file.WriteInt16(log.GoRow.RefreshNum)
}

func (log *PlayerTraderRefreshStateInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTraderRefreshState.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastUpdateTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.AutoUpdateTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TraderId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.RefreshNum)))
	return stmt.Execute()
}

func (log *PlayerTraderRefreshStateInsertLog) CommitToTLog() {
}

func (log *PlayerTraderRefreshStateInsertLog) CommitToXdLog() {
}

func (log *PlayerTraderRefreshStateInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTraderRefreshStateDeleteLog struct {
	db *Database
	Old *C.PlayerTraderRefreshState
	GoOld *PlayerTraderRefreshState
}

func (db *Database) newPlayerTraderRefreshStateDeleteLog(old *C.PlayerTraderRefreshState, goOld *PlayerTraderRefreshState) *PlayerTraderRefreshStateDeleteLog {
	return &PlayerTraderRefreshStateDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerTraderRefreshStateDeleteLog) Free() {
	C.Free_PlayerTraderRefreshState(log.Old)
}

func (log *PlayerTraderRefreshStateDeleteLog) InvokeHook() {
	if g_Hooks.PlayerTraderRefreshState != nil {
		g_Hooks.PlayerTraderRefreshState.Delete(&PlayerTraderRefreshStateRow{row: log.Old})
	}
}

func (log *PlayerTraderRefreshStateDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerTraderRefreshStateDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(140)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.LastUpdateTime)
	file.WriteInt64(log.GoOld.AutoUpdateTime)
	file.WriteInt16(log.GoOld.TraderId)
	file.WriteInt16(log.GoOld.RefreshNum)
}

func (log *PlayerTraderRefreshStateDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTraderRefreshState.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerTraderRefreshStateDeleteLog) CommitToTLog() {
}

func (log *PlayerTraderRefreshStateDeleteLog) CommitToXdLog() {
}

func (log *PlayerTraderRefreshStateDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTraderRefreshStateUpdateLog struct {
	db *Database 
	Row *C.PlayerTraderRefreshState
	Old *C.PlayerTraderRefreshState
	GoNew *PlayerTraderRefreshState
	GoOld *PlayerTraderRefreshState
}

func (db *Database) newPlayerTraderRefreshStateUpdateLog(row, old *C.PlayerTraderRefreshState, goNew, goOld *PlayerTraderRefreshState) *PlayerTraderRefreshStateUpdateLog {
	return &PlayerTraderRefreshStateUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerTraderRefreshStateUpdateLog) Free() {
	C.Free_PlayerTraderRefreshState(log.Old)
}

func (log *PlayerTraderRefreshStateUpdateLog) InvokeHook() {
	if g_Hooks.PlayerTraderRefreshState != nil {
		g_Hooks.PlayerTraderRefreshState.Update(&PlayerTraderRefreshStateRow{row: log.Row}, &PlayerTraderRefreshStateRow{row: log.Old})
	}
}

func (log *PlayerTraderRefreshStateUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerTraderRefreshStateUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(140)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.LastUpdateTime)
	file.WriteInt64(log.GoNew.AutoUpdateTime)
	file.WriteInt16(log.GoNew.TraderId)
	file.WriteInt16(log.GoNew.RefreshNum)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.LastUpdateTime)
	file.WriteInt64(log.GoOld.AutoUpdateTime)
	file.WriteInt16(log.GoOld.TraderId)
	file.WriteInt16(log.GoOld.RefreshNum)
}

func (log *PlayerTraderRefreshStateUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTraderRefreshState.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.LastUpdateTime)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.AutoUpdateTime)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TraderId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.RefreshNum)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerTraderRefreshStateUpdateLog) CommitToTLog() {
}

func (log *PlayerTraderRefreshStateUpdateLog) CommitToXdLog() {
}

func (log *PlayerTraderRefreshStateUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerTraderRefreshStateInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerTraderRefreshState
	for crow := log.db.tables.PlayerTraderRefreshState; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerTraderRefreshState'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerTraderRefreshState = log.db.tables.PlayerTraderRefreshState.next
	}
	C.Free_PlayerTraderRefreshState(log.Row)
}

func (log *PlayerTraderRefreshStateDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerTraderRefreshState
	for crow := log.db.tables.PlayerTraderRefreshState; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_trader_refresh_state'")
	}
	log.Old.next = log.db.tables.PlayerTraderRefreshState
	log.db.tables.PlayerTraderRefreshState = log.Old
}

func (log *PlayerTraderRefreshStateUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerTraderRefreshState
	for crow := log.db.tables.PlayerTraderRefreshState; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_trader_refresh_state'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerTraderRefreshState = log.Old
	}
	C.Free_PlayerTraderRefreshState(log.Row)
}

/* ========== player_trader_store_state ========== */

type PlayerTraderStoreStateInsertLog struct {
	db *Database 
	Row *C.PlayerTraderStoreState
	GoRow *PlayerTraderStoreState
}

func (db *Database) newPlayerTraderStoreStateInsertLog(row *C.PlayerTraderStoreState, goRow *PlayerTraderStoreState) *PlayerTraderStoreStateInsertLog {
	return &PlayerTraderStoreStateInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerTraderStoreStateInsertLog) Free() {
}

func (log *PlayerTraderStoreStateInsertLog) InvokeHook() {
	if g_Hooks.PlayerTraderStoreState != nil {
		g_Hooks.PlayerTraderStoreState.Insert(&PlayerTraderStoreStateRow{row: log.Row})
	}
}

func (log *PlayerTraderStoreStateInsertLog) SetEventId(time.Time) {
}

func (log *PlayerTraderStoreStateInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(141)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt16(log.GoRow.TraderId)
	file.WriteInt32(log.GoRow.GridId)
	file.WriteInt16(log.GoRow.ItemId)
	file.WriteInt16(log.GoRow.Num)
	file.WriteInt64(log.GoRow.Cost)
	file.WriteInt8(log.GoRow.Stock)
	file.WriteInt8(log.GoRow.GoodsType)
}

func (log *PlayerTraderStoreStateInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTraderStoreState.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TraderId)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.GridId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Num)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Cost)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Stock)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.GoodsType)))
	return stmt.Execute()
}

func (log *PlayerTraderStoreStateInsertLog) CommitToTLog() {
}

func (log *PlayerTraderStoreStateInsertLog) CommitToXdLog() {
}

func (log *PlayerTraderStoreStateInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTraderStoreStateDeleteLog struct {
	db *Database
	Old *C.PlayerTraderStoreState
	GoOld *PlayerTraderStoreState
}

func (db *Database) newPlayerTraderStoreStateDeleteLog(old *C.PlayerTraderStoreState, goOld *PlayerTraderStoreState) *PlayerTraderStoreStateDeleteLog {
	return &PlayerTraderStoreStateDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerTraderStoreStateDeleteLog) Free() {
	C.Free_PlayerTraderStoreState(log.Old)
}

func (log *PlayerTraderStoreStateDeleteLog) InvokeHook() {
	if g_Hooks.PlayerTraderStoreState != nil {
		g_Hooks.PlayerTraderStoreState.Delete(&PlayerTraderStoreStateRow{row: log.Old})
	}
}

func (log *PlayerTraderStoreStateDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerTraderStoreStateDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(141)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.TraderId)
	file.WriteInt32(log.GoOld.GridId)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt16(log.GoOld.Num)
	file.WriteInt64(log.GoOld.Cost)
	file.WriteInt8(log.GoOld.Stock)
	file.WriteInt8(log.GoOld.GoodsType)
}

func (log *PlayerTraderStoreStateDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTraderStoreState.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerTraderStoreStateDeleteLog) CommitToTLog() {
}

func (log *PlayerTraderStoreStateDeleteLog) CommitToXdLog() {
}

func (log *PlayerTraderStoreStateDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerTraderStoreStateUpdateLog struct {
	db *Database 
	Row *C.PlayerTraderStoreState
	Old *C.PlayerTraderStoreState
	GoNew *PlayerTraderStoreState
	GoOld *PlayerTraderStoreState
}

func (db *Database) newPlayerTraderStoreStateUpdateLog(row, old *C.PlayerTraderStoreState, goNew, goOld *PlayerTraderStoreState) *PlayerTraderStoreStateUpdateLog {
	return &PlayerTraderStoreStateUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerTraderStoreStateUpdateLog) Free() {
	C.Free_PlayerTraderStoreState(log.Old)
}

func (log *PlayerTraderStoreStateUpdateLog) InvokeHook() {
	if g_Hooks.PlayerTraderStoreState != nil {
		g_Hooks.PlayerTraderStoreState.Update(&PlayerTraderStoreStateRow{row: log.Row}, &PlayerTraderStoreStateRow{row: log.Old})
	}
}

func (log *PlayerTraderStoreStateUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerTraderStoreStateUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(141)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt16(log.GoNew.TraderId)
	file.WriteInt32(log.GoNew.GridId)
	file.WriteInt16(log.GoNew.ItemId)
	file.WriteInt16(log.GoNew.Num)
	file.WriteInt64(log.GoNew.Cost)
	file.WriteInt8(log.GoNew.Stock)
	file.WriteInt8(log.GoNew.GoodsType)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt16(log.GoOld.TraderId)
	file.WriteInt32(log.GoOld.GridId)
	file.WriteInt16(log.GoOld.ItemId)
	file.WriteInt16(log.GoOld.Num)
	file.WriteInt64(log.GoOld.Cost)
	file.WriteInt8(log.GoOld.Stock)
	file.WriteInt8(log.GoOld.GoodsType)
}

func (log *PlayerTraderStoreStateUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerTraderStoreState.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.TraderId)))
	stmt.BindInt(unsafe.Pointer(&(log.Row.GridId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.ItemId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Num)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Cost)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.Stock)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.GoodsType)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerTraderStoreStateUpdateLog) CommitToTLog() {
}

func (log *PlayerTraderStoreStateUpdateLog) CommitToXdLog() {
}

func (log *PlayerTraderStoreStateUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerTraderStoreStateInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerTraderStoreState
	for crow := log.db.tables.PlayerTraderStoreState; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerTraderStoreState'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerTraderStoreState = log.db.tables.PlayerTraderStoreState.next
	}
	C.Free_PlayerTraderStoreState(log.Row)
}

func (log *PlayerTraderStoreStateDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerTraderStoreState
	for crow := log.db.tables.PlayerTraderStoreState; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_trader_store_state'")
	}
	log.Old.next = log.db.tables.PlayerTraderStoreState
	log.db.tables.PlayerTraderStoreState = log.Old
}

func (log *PlayerTraderStoreStateUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerTraderStoreState
	for crow := log.db.tables.PlayerTraderStoreState; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_trader_store_state'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerTraderStoreState = log.Old
	}
	C.Free_PlayerTraderStoreState(log.Row)
}

/* ========== player_use_skill ========== */

type PlayerUseSkillInsertLog struct {
	db *Database 
	Row *C.PlayerUseSkill
	GoRow *PlayerUseSkill
}

func (db *Database) newPlayerUseSkillInsertLog(row *C.PlayerUseSkill, goRow *PlayerUseSkill) *PlayerUseSkillInsertLog {
	return &PlayerUseSkillInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerUseSkillInsertLog) Free() {
}

func (log *PlayerUseSkillInsertLog) InvokeHook() {
	if g_Hooks.PlayerUseSkill != nil {
		g_Hooks.PlayerUseSkill.Insert(&PlayerUseSkillRow{row: log.Row})
	}
}

func (log *PlayerUseSkillInsertLog) SetEventId(time.Time) {
}

func (log *PlayerUseSkillInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(142)
	file.WriteInt64(log.GoRow.Id)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt8(log.GoRow.RoleId)
	file.WriteInt16(log.GoRow.SkillId1)
	file.WriteInt16(log.GoRow.SkillId4)
	file.WriteInt16(log.GoRow.SkillId2)
	file.WriteInt16(log.GoRow.SkillId3)
}

func (log *PlayerUseSkillInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerUseSkill.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RoleId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SkillId1)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SkillId4)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SkillId2)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SkillId3)))
	return stmt.Execute()
}

func (log *PlayerUseSkillInsertLog) CommitToTLog() {
}

func (log *PlayerUseSkillInsertLog) CommitToXdLog() {
}

func (log *PlayerUseSkillInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerUseSkillDeleteLog struct {
	db *Database
	Old *C.PlayerUseSkill
	GoOld *PlayerUseSkill
}

func (db *Database) newPlayerUseSkillDeleteLog(old *C.PlayerUseSkill, goOld *PlayerUseSkill) *PlayerUseSkillDeleteLog {
	return &PlayerUseSkillDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerUseSkillDeleteLog) Free() {
	C.Free_PlayerUseSkill(log.Old)
}

func (log *PlayerUseSkillDeleteLog) InvokeHook() {
	if g_Hooks.PlayerUseSkill != nil {
		g_Hooks.PlayerUseSkill.Delete(&PlayerUseSkillRow{row: log.Old})
	}
}

func (log *PlayerUseSkillDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerUseSkillDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(142)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.RoleId)
	file.WriteInt16(log.GoOld.SkillId1)
	file.WriteInt16(log.GoOld.SkillId4)
	file.WriteInt16(log.GoOld.SkillId2)
	file.WriteInt16(log.GoOld.SkillId3)
}

func (log *PlayerUseSkillDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerUseSkill.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Id)))
	return stmt.Execute()
}

func (log *PlayerUseSkillDeleteLog) CommitToTLog() {
}

func (log *PlayerUseSkillDeleteLog) CommitToXdLog() {
}

func (log *PlayerUseSkillDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerUseSkillUpdateLog struct {
	db *Database 
	Row *C.PlayerUseSkill
	Old *C.PlayerUseSkill
	GoNew *PlayerUseSkill
	GoOld *PlayerUseSkill
}

func (db *Database) newPlayerUseSkillUpdateLog(row, old *C.PlayerUseSkill, goNew, goOld *PlayerUseSkill) *PlayerUseSkillUpdateLog {
	return &PlayerUseSkillUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerUseSkillUpdateLog) Free() {
	C.Free_PlayerUseSkill(log.Old)
}

func (log *PlayerUseSkillUpdateLog) InvokeHook() {
	if g_Hooks.PlayerUseSkill != nil {
		g_Hooks.PlayerUseSkill.Update(&PlayerUseSkillRow{row: log.Row}, &PlayerUseSkillRow{row: log.Old})
	}
}

func (log *PlayerUseSkillUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerUseSkillUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(142)
	file.WriteInt64(log.GoNew.Id)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt8(log.GoNew.RoleId)
	file.WriteInt16(log.GoNew.SkillId1)
	file.WriteInt16(log.GoNew.SkillId4)
	file.WriteInt16(log.GoNew.SkillId2)
	file.WriteInt16(log.GoNew.SkillId3)
	file.WriteInt64(log.GoOld.Id)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt8(log.GoOld.RoleId)
	file.WriteInt16(log.GoOld.SkillId1)
	file.WriteInt16(log.GoOld.SkillId4)
	file.WriteInt16(log.GoOld.SkillId2)
	file.WriteInt16(log.GoOld.SkillId3)
}

func (log *PlayerUseSkillUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerUseSkill.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindTinyInt(unsafe.Pointer(&(log.Row.RoleId)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SkillId1)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SkillId4)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SkillId2)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.SkillId3)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Id)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerUseSkillUpdateLog) CommitToTLog() {
}

func (log *PlayerUseSkillUpdateLog) CommitToXdLog() {
}

func (log *PlayerUseSkillUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerUseSkillInsertLog) Rollback() {
	key := log.Row.Id
	var old, prev *C.PlayerUseSkill
	for crow := log.db.tables.PlayerUseSkill; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("insert rollback check failed 'PlayerUseSkill'")
	}
	if prev != nil {
		prev.next = old.next
	} else {
		log.db.tables.PlayerUseSkill = log.db.tables.PlayerUseSkill.next
	}
	C.Free_PlayerUseSkill(log.Row)
}

func (log *PlayerUseSkillDeleteLog) Rollback() {
	key := log.Old.Id
	var old *C.PlayerUseSkill
	for crow := log.db.tables.PlayerUseSkill; crow != nil; crow = crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != nil {
		panic("delete rollback check failed 'player_use_skill'")
	}
	log.Old.next = log.db.tables.PlayerUseSkill
	log.db.tables.PlayerUseSkill = log.Old
}

func (log *PlayerUseSkillUpdateLog) Rollback() {
	key := log.Old.Id
	var old, prev *C.PlayerUseSkill
	for crow := log.db.tables.PlayerUseSkill; crow != nil; prev, crow = crow, crow.next {
		if crow.Id == key {
			old = crow
			break
		}
	}
	if old != log.Row {
		panic("update rollback check failed 'player_use_skill'")
	}
	log.Old.next = old.next
	if prev != nil {
		prev.next = log.Old
	} else {
		log.db.tables.PlayerUseSkill = log.Old
	}
	C.Free_PlayerUseSkill(log.Row)
}

/* ========== player_vip ========== */

type PlayerVipInsertLog struct {
	db *Database 
	Row *C.PlayerVip
	GoRow *PlayerVip
}

func (db *Database) newPlayerVipInsertLog(row *C.PlayerVip, goRow *PlayerVip) *PlayerVipInsertLog {
	return &PlayerVipInsertLog{db:db, Row:row, GoRow:goRow}
}

func (log *PlayerVipInsertLog) Free() {
}

func (log *PlayerVipInsertLog) InvokeHook() {
	if g_Hooks.PlayerVip != nil {
		g_Hooks.PlayerVip.Insert(&PlayerVipRow{row: log.Row})
	}
}

func (log *PlayerVipInsertLog) SetEventId(time.Time) {
}

func (log *PlayerVipInsertLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(1)
	file.WriteInt32(143)
	file.WriteInt64(log.GoRow.Pid)
	file.WriteInt64(log.GoRow.Ingot)
	file.WriteInt16(log.GoRow.Level)
	file.WriteString(log.GoRow.CardId)
	file.WriteInt64(log.GoRow.PresentExp)
}

func (log *PlayerVipInsertLog) CommitToMySql() error {
	stmt := g_SQL.PlayerVip.Insert
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Ingot)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.CardId), int(log.Row.CardId_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.PresentExp)))
	return stmt.Execute()
}

func (log *PlayerVipInsertLog) CommitToTLog() {
}

func (log *PlayerVipInsertLog) CommitToXdLog() {
}

func (log *PlayerVipInsertLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerVipDeleteLog struct {
	db *Database
	Old *C.PlayerVip
	GoOld *PlayerVip
}

func (db *Database) newPlayerVipDeleteLog(old *C.PlayerVip, goOld *PlayerVip) *PlayerVipDeleteLog {
	return &PlayerVipDeleteLog{db:db, Old:old, GoOld:goOld}
}

func (log *PlayerVipDeleteLog) Free() {
	C.Free_PlayerVip(log.Old)
}

func (log *PlayerVipDeleteLog) InvokeHook() {
	if g_Hooks.PlayerVip != nil {
		g_Hooks.PlayerVip.Delete(&PlayerVipRow{row: log.Old})
	}
}

func (log *PlayerVipDeleteLog) SetEventId(time.Time) {
}

func (log *PlayerVipDeleteLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(2)
	file.WriteInt32(143)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.Ingot)
	file.WriteInt16(log.GoOld.Level)
	file.WriteString(log.GoOld.CardId)
	file.WriteInt64(log.GoOld.PresentExp)
}

func (log *PlayerVipDeleteLog) CommitToMySql() error {
	stmt := g_SQL.PlayerVip.Delete
	stmt.CleanBind()
	stmt.BindBigInt(unsafe.Pointer(&(log.Old.Pid)))
	return stmt.Execute()
}

func (log *PlayerVipDeleteLog) CommitToTLog() {
}

func (log *PlayerVipDeleteLog) CommitToXdLog() {
}

func (log *PlayerVipDeleteLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

type PlayerVipUpdateLog struct {
	db *Database 
	Row *C.PlayerVip
	Old *C.PlayerVip
	GoNew *PlayerVip
	GoOld *PlayerVip
}

func (db *Database) newPlayerVipUpdateLog(row, old *C.PlayerVip, goNew, goOld *PlayerVip) *PlayerVipUpdateLog {
	return &PlayerVipUpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}
}

func (log *PlayerVipUpdateLog) Free() {
	C.Free_PlayerVip(log.Old)
}

func (log *PlayerVipUpdateLog) InvokeHook() {
	if g_Hooks.PlayerVip != nil {
		g_Hooks.PlayerVip.Update(&PlayerVipRow{row: log.Row}, &PlayerVipRow{row: log.Old})
	}
}

func (log *PlayerVipUpdateLog) SetEventId(time.Time) {
}

func (log *PlayerVipUpdateLog) CommitToFile(file *SyncFile) {
	file.WriteUint8(3)
	file.WriteInt32(143)
	file.WriteInt64(log.GoNew.Pid)
	file.WriteInt64(log.GoNew.Ingot)
	file.WriteInt16(log.GoNew.Level)
	file.WriteString(log.GoNew.CardId)
	file.WriteInt64(log.GoNew.PresentExp)
	file.WriteInt64(log.GoOld.Pid)
	file.WriteInt64(log.GoOld.Ingot)
	file.WriteInt16(log.GoOld.Level)
	file.WriteString(log.GoOld.CardId)
	file.WriteInt64(log.GoOld.PresentExp)
}

func (log *PlayerVipUpdateLog) CommitToMySql() error {
	stmt := g_SQL.PlayerVip.Update
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Ingot)))
	stmt.BindSmallInt(unsafe.Pointer(&(log.Row.Level)))
	stmt.BindVarchar(unsafe.Pointer(log.Row.CardId), int(log.Row.CardId_len))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.PresentExp)))
	stmt.BindBigInt(unsafe.Pointer(&(log.Row.Pid)))
	stmt.CleanBind()
	return stmt.Execute()
}

func (log *PlayerVipUpdateLog) CommitToTLog() {
}

func (log *PlayerVipUpdateLog) CommitToXdLog() {
}

func (log *PlayerVipUpdateLog) GetTransType() int8 {
	return TRANS_TYPE_MYSQL
}

func (log *PlayerVipInsertLog) Rollback() {
	if log.db.tables.PlayerVip != log.Row {
		panic("insert rollback check failed 'PlayerVip'")
	}
	log.db.tables.PlayerVip = nil
	C.Free_PlayerVip(log.Row)
}

func (log *PlayerVipDeleteLog) Rollback() {
	if log.db.tables.PlayerVip != nil {
		panic("delete rollback check failed 'player_vip'")
	}
	log.db.tables.PlayerVip = log.Old
}

func (log *PlayerVipUpdateLog) Rollback() {
	if log.db.tables.PlayerVip != log.Row {
		panic("update rollback check failed 'player_vip'")
	}
	log.db.tables.PlayerVip = log.Old
	C.Free_PlayerVip(log.Row)
}

