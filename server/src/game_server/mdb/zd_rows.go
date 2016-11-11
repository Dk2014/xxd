package mdb

/*
#include "zd_cgo.h"
*/
import "C"
import "unsafe"

/* ========== global_announcement ========== */

// 公告列表
type GlobalAnnouncementRow struct {
	row *C.GlobalAnnouncement
	isBreak bool
	cached_Parameters bool
	cacheof_Parameters string
	cached_Content bool
	cacheof_Content string
}

func (this *GlobalAnnouncementRow) reset(row *C.GlobalAnnouncement) {	this.row = row
	this.isBreak = false
	this.cached_Parameters = false
	this.cached_Content = false
}

func (this *GlobalAnnouncementRow) Break() {
	this.isBreak = true
}

// 公告ID
func (this *GlobalAnnouncementRow) Id() int64 {
	return int64(this.row.Id)
}

// 创建时间戳
func (this *GlobalAnnouncementRow) ExpireTime() int64 {
	return int64(this.row.ExpireTime)
}

// 公告模版ID
func (this *GlobalAnnouncementRow) TplId() int32 {
	return int32(this.row.TplId)
}

// 模版参数
func (this *GlobalAnnouncementRow) Parameters() string {
	if !this.cached_Parameters {
		this.cached_Parameters = true
		this.cacheof_Parameters = C.GoStringN(this.row.Parameters, this.row.Parameters_len)
	}
	return this.cacheof_Parameters
}

// 公共内容，有则忽略模版
func (this *GlobalAnnouncementRow) Content() string {
	if !this.cached_Content {
		this.cached_Content = true
		this.cacheof_Content = C.GoStringN(this.row.Content, this.row.Content_len)
	}
	return this.cacheof_Content
}

// 发送时间
func (this *GlobalAnnouncementRow) SendTime() int64 {
	return int64(this.row.SendTime)
}

// 间隔时间
func (this *GlobalAnnouncementRow) SpacingTime() int64 {
	return int64(this.row.SpacingTime)
}

func (this *GlobalAnnouncementRow) GoObject() *GlobalAnnouncement {
	return &GlobalAnnouncement{
		Id: int64(this.row.Id),
		ExpireTime: int64(this.row.ExpireTime),
		TplId: int32(this.row.TplId),
		Parameters: this.Parameters(),
		Content: this.Content(),
		SendTime: int64(this.row.SendTime),
		SpacingTime: int64(this.row.SpacingTime),
	}
}

/* ========== global_arena_rank ========== */

// 全局比武场数据
type GlobalArenaRankRow struct {
	row *C.GlobalArenaRank
	isBreak bool
}

func (this *GlobalArenaRankRow) reset(row *C.GlobalArenaRank) {	this.row = row
	this.isBreak = false
}

func (this *GlobalArenaRankRow) Break() {
	this.isBreak = true
}

// 排名
func (this *GlobalArenaRankRow) Rank() int32 {
	return int32(this.row.Rank)
}

// 玩家ID
func (this *GlobalArenaRankRow) Pid() int64 {
	return int64(this.row.Pid)
}

func (this *GlobalArenaRankRow) GoObject() *GlobalArenaRank {
	return &GlobalArenaRank{
		Rank: int32(this.row.Rank),
		Pid: int64(this.row.Pid),
	}
}

/* ========== global_clique ========== */

// 帮派信息
type GlobalCliqueRow struct {
	row *C.GlobalClique
	isBreak bool
	cached_Name bool
	cacheof_Name string
	cached_Announce bool
	cacheof_Announce string
	cached_MemberJson bool
	cacheof_MemberJson []byte
}

func (this *GlobalCliqueRow) reset(row *C.GlobalClique) {	this.row = row
	this.isBreak = false
	this.cached_Name = false
	this.cached_Announce = false
	this.cached_MemberJson = false
}

func (this *GlobalCliqueRow) Break() {
	this.isBreak = true
}

// 主键
func (this *GlobalCliqueRow) Id() int64 {
	return int64(this.row.Id)
}

// 名字
func (this *GlobalCliqueRow) Name() string {
	if !this.cached_Name {
		this.cached_Name = true
		this.cacheof_Name = C.GoStringN(this.row.Name, this.row.Name_len)
	}
	return this.cacheof_Name
}

// 公告
func (this *GlobalCliqueRow) Announce() string {
	if !this.cached_Announce {
		this.cached_Announce = true
		this.cacheof_Announce = C.GoStringN(this.row.Announce, this.row.Announce_len)
	}
	return this.cacheof_Announce
}

// 累积贡献铜币
func (this *GlobalCliqueRow) TotalDonateCoins() int64 {
	return int64(this.row.TotalDonateCoins)
}

// 帮主
func (this *GlobalCliqueRow) OwnerPid() int64 {
	return int64(this.row.OwnerPid)
}

// 帮主最近登录时间
func (this *GlobalCliqueRow) OwnerLoginTime() int64 {
	return int64(this.row.OwnerLoginTime)
}

// 副帮主1
func (this *GlobalCliqueRow) MangerPid1() int64 {
	return int64(this.row.MangerPid1)
}

// 副帮主2
func (this *GlobalCliqueRow) MangerPid2() int64 {
	return int64(this.row.MangerPid2)
}

// 总部当前贡献铜币
func (this *GlobalCliqueRow) CenterBuildingCoins() int64 {
	return int64(this.row.CenterBuildingCoins)
}

// 宗祠贡献铜币
func (this *GlobalCliqueRow) TempleBuildingCoins() int64 {
	return int64(this.row.TempleBuildingCoins)
}

// 回春堂当前贡献铜币
func (this *GlobalCliqueRow) HealthBuildingCoins() int64 {
	return int64(this.row.HealthBuildingCoins)
}

// 神兵堂当前贡献铜币
func (this *GlobalCliqueRow) AttackBuildingCoins() int64 {
	return int64(this.row.AttackBuildingCoins)
}

// 金刚堂当前贡献铜币
func (this *GlobalCliqueRow) DefenseBuildingCoins() int64 {
	return int64(this.row.DefenseBuildingCoins)
}

// 战备仓库当前贡献铜币
func (this *GlobalCliqueRow) StoreBuildingCoins() int64 {
	return int64(this.row.StoreBuildingCoins)
}

// 钱庄当前贡献铜币
func (this *GlobalCliqueRow) BankBuildingCoins() int64 {
	return int64(this.row.BankBuildingCoins)
}

// 总部当前等级
func (this *GlobalCliqueRow) CenterBuildingLevel() int16 {
	return int16(this.row.CenterBuildingLevel)
}

// 宗祠等级
func (this *GlobalCliqueRow) TempleBuildingLevel() int16 {
	return int16(this.row.TempleBuildingLevel)
}

// 回春堂当前等级
func (this *GlobalCliqueRow) HealthBuildingLevel() int16 {
	return int16(this.row.HealthBuildingLevel)
}

// 神兵堂当前等级
func (this *GlobalCliqueRow) AttackBuildingLevel() int16 {
	return int16(this.row.AttackBuildingLevel)
}

// 金刚堂当前等级
func (this *GlobalCliqueRow) DefenseBuildingLevel() int16 {
	return int16(this.row.DefenseBuildingLevel)
}

// 钱庄当前等级
func (this *GlobalCliqueRow) BankBuildingLevel() int16 {
	return int16(this.row.BankBuildingLevel)
}

// 成员列表存储 pid 的json 数组
func (this *GlobalCliqueRow) MemberJson() []byte {
	if !this.cached_MemberJson {
		this.cached_MemberJson = true
		this.cacheof_MemberJson = C.GoBytes(unsafe.Pointer(this.row.MemberJson), this.row.MemberJson_len)
	}
	return this.cacheof_MemberJson
}

// 自动审核
func (this *GlobalCliqueRow) AutoAudit() int8 {
	return int8(this.row.AutoAudit)
}

// 自动审核等级下限
func (this *GlobalCliqueRow) AutoAuditLevel() int16 {
	return int16(this.row.AutoAuditLevel)
}

// 帮派贡献
func (this *GlobalCliqueRow) Contrib() int64 {
	return int64(this.row.Contrib)
}

// 帮派贡献清除时间
func (this *GlobalCliqueRow) ContribClearTime() int64 {
	return int64(this.row.ContribClearTime)
}

// 招募公告时间
func (this *GlobalCliqueRow) RecruitTime() int64 {
	return int64(this.row.RecruitTime)
}

// 上香时间
func (this *GlobalCliqueRow) WorshipTime() int64 {
	return int64(this.row.WorshipTime)
}

// 当日上香次数
func (this *GlobalCliqueRow) WorshipCnt() int8 {
	return int8(this.row.WorshipCnt)
}

// 战备物资发送时间
func (this *GlobalCliqueRow) StoreSendTime() int64 {
	return int64(this.row.StoreSendTime)
}

// 战备物资发送次数
func (this *GlobalCliqueRow) StoreSendCnt() int8 {
	return int8(this.row.StoreSendCnt)
}

func (this *GlobalCliqueRow) GoObject() *GlobalClique {
	return &GlobalClique{
		Id: int64(this.row.Id),
		Name: this.Name(),
		Announce: this.Announce(),
		TotalDonateCoins: int64(this.row.TotalDonateCoins),
		OwnerPid: int64(this.row.OwnerPid),
		OwnerLoginTime: int64(this.row.OwnerLoginTime),
		MangerPid1: int64(this.row.MangerPid1),
		MangerPid2: int64(this.row.MangerPid2),
		CenterBuildingCoins: int64(this.row.CenterBuildingCoins),
		TempleBuildingCoins: int64(this.row.TempleBuildingCoins),
		HealthBuildingCoins: int64(this.row.HealthBuildingCoins),
		AttackBuildingCoins: int64(this.row.AttackBuildingCoins),
		DefenseBuildingCoins: int64(this.row.DefenseBuildingCoins),
		StoreBuildingCoins: int64(this.row.StoreBuildingCoins),
		BankBuildingCoins: int64(this.row.BankBuildingCoins),
		CenterBuildingLevel: int16(this.row.CenterBuildingLevel),
		TempleBuildingLevel: int16(this.row.TempleBuildingLevel),
		HealthBuildingLevel: int16(this.row.HealthBuildingLevel),
		AttackBuildingLevel: int16(this.row.AttackBuildingLevel),
		DefenseBuildingLevel: int16(this.row.DefenseBuildingLevel),
		BankBuildingLevel: int16(this.row.BankBuildingLevel),
		MemberJson: this.MemberJson(),
		AutoAudit: int8(this.row.AutoAudit),
		AutoAuditLevel: int16(this.row.AutoAuditLevel),
		Contrib: int64(this.row.Contrib),
		ContribClearTime: int64(this.row.ContribClearTime),
		RecruitTime: int64(this.row.RecruitTime),
		WorshipTime: int64(this.row.WorshipTime),
		WorshipCnt: int8(this.row.WorshipCnt),
		StoreSendTime: int64(this.row.StoreSendTime),
		StoreSendCnt: int8(this.row.StoreSendCnt),
	}
}

/* ========== global_clique_boat ========== */

// 镖船信息
type GlobalCliqueBoatRow struct {
	row *C.GlobalCliqueBoat
	isBreak bool
}

func (this *GlobalCliqueBoatRow) reset(row *C.GlobalCliqueBoat) {	this.row = row
	this.isBreak = false
}

func (this *GlobalCliqueBoatRow) Break() {
	this.isBreak = true
}

// ID
func (this *GlobalCliqueBoatRow) Id() int64 {
	return int64(this.row.Id)
}

// 所属帮派ID
func (this *GlobalCliqueBoatRow) CliqueId() int64 {
	return int64(this.row.CliqueId)
}

// 镖船类型
func (this *GlobalCliqueBoatRow) BoatType() int16 {
	return int16(this.row.BoatType)
}

// 护送人PID
func (this *GlobalCliqueBoatRow) OwnerPid() int64 {
	return int64(this.row.OwnerPid)
}

// 本次送镖时间戳
func (this *GlobalCliqueBoatRow) EscortStartTimestamp() int64 {
	return int64(this.row.EscortStartTimestamp)
}

// 累计送镖时间
func (this *GlobalCliqueBoatRow) TotalEscortTime() int16 {
	return int16(this.row.TotalEscortTime)
}

// 夺回玩家PID
func (this *GlobalCliqueBoatRow) RecoverPid() int64 {
	return int64(this.row.RecoverPid)
}

// 劫持人PID
func (this *GlobalCliqueBoatRow) HijackerPid() int64 {
	return int64(this.row.HijackerPid)
}

// 本次送镖时间戳
func (this *GlobalCliqueBoatRow) HijackTimestamp() int64 {
	return int64(this.row.HijackTimestamp)
}

// 状态 0--运送中 1--劫持中 2--夺回中 3--劫持完成 
func (this *GlobalCliqueBoatRow) Status() int8 {
	return int8(this.row.Status)
}

func (this *GlobalCliqueBoatRow) GoObject() *GlobalCliqueBoat {
	return &GlobalCliqueBoat{
		Id: int64(this.row.Id),
		CliqueId: int64(this.row.CliqueId),
		BoatType: int16(this.row.BoatType),
		OwnerPid: int64(this.row.OwnerPid),
		EscortStartTimestamp: int64(this.row.EscortStartTimestamp),
		TotalEscortTime: int16(this.row.TotalEscortTime),
		RecoverPid: int64(this.row.RecoverPid),
		HijackerPid: int64(this.row.HijackerPid),
		HijackTimestamp: int64(this.row.HijackTimestamp),
		Status: int8(this.row.Status),
	}
}

/* ========== global_despair_boss ========== */

// 绝望关卡boss
type GlobalDespairBossRow struct {
	row *C.GlobalDespairBoss
	isBreak bool
}

func (this *GlobalDespairBossRow) reset(row *C.GlobalDespairBoss) {	this.row = row
	this.isBreak = false
}

func (this *GlobalDespairBossRow) Break() {
	this.isBreak = true
}

// 阵营类型
func (this *GlobalDespairBossRow) CampType() int8 {
	return int8(this.row.CampType)
}

// boss等级
func (this *GlobalDespairBossRow) Level() int16 {
	return int16(this.row.Level)
}

// boss出现时间
func (this *GlobalDespairBossRow) Timestamp() int64 {
	return int64(this.row.Timestamp)
}

// boss死亡时间
func (this *GlobalDespairBossRow) DeadTimestamp() int64 {
	return int64(this.row.DeadTimestamp)
}

// boss当前血量
func (this *GlobalDespairBossRow) Hp() int64 {
	return int64(this.row.Hp)
}

// boss最大血量
func (this *GlobalDespairBossRow) MaxHp() int64 {
	return int64(this.row.MaxHp)
}

func (this *GlobalDespairBossRow) GoObject() *GlobalDespairBoss {
	return &GlobalDespairBoss{
		CampType: int8(this.row.CampType),
		Level: int16(this.row.Level),
		Timestamp: int64(this.row.Timestamp),
		DeadTimestamp: int64(this.row.DeadTimestamp),
		Hp: int64(this.row.Hp),
		MaxHp: int64(this.row.MaxHp),
	}
}

/* ========== global_despair_boss_history ========== */

// 绝望关卡boss历史记录
type GlobalDespairBossHistoryRow struct {
	row *C.GlobalDespairBossHistory
	isBreak bool
}

func (this *GlobalDespairBossHistoryRow) reset(row *C.GlobalDespairBossHistory) {	this.row = row
	this.isBreak = false
}

func (this *GlobalDespairBossHistoryRow) Break() {
	this.isBreak = true
}

// 
func (this *GlobalDespairBossHistoryRow) Id() int64 {
	return int64(this.row.Id)
}

// 绝望之地version
func (this *GlobalDespairBossHistoryRow) Version() int64 {
	return int64(this.row.Version)
}

// 阵营类型
func (this *GlobalDespairBossHistoryRow) CampType() int8 {
	return int8(this.row.CampType)
}

// boss等级
func (this *GlobalDespairBossHistoryRow) Level() int16 {
	return int16(this.row.Level)
}

// 记录产生时间
func (this *GlobalDespairBossHistoryRow) Timestamp() int64 {
	return int64(this.row.Timestamp)
}

// boss出现时间
func (this *GlobalDespairBossHistoryRow) StartTimestamp() int64 {
	return int64(this.row.StartTimestamp)
}

// boss死亡时间
func (this *GlobalDespairBossHistoryRow) DeadTimestamp() int64 {
	return int64(this.row.DeadTimestamp)
}

func (this *GlobalDespairBossHistoryRow) GoObject() *GlobalDespairBossHistory {
	return &GlobalDespairBossHistory{
		Id: int64(this.row.Id),
		Version: int64(this.row.Version),
		CampType: int8(this.row.CampType),
		Level: int16(this.row.Level),
		Timestamp: int64(this.row.Timestamp),
		StartTimestamp: int64(this.row.StartTimestamp),
		DeadTimestamp: int64(this.row.DeadTimestamp),
	}
}

/* ========== global_despair_land ========== */

// 绝望之地状态
type GlobalDespairLandRow struct {
	row *C.GlobalDespairLand
	isBreak bool
}

func (this *GlobalDespairLandRow) reset(row *C.GlobalDespairLand) {	this.row = row
	this.isBreak = false
}

func (this *GlobalDespairLandRow) Break() {
	this.isBreak = true
}

// 主键ID
func (this *GlobalDespairLandRow) Id() int64 {
	return int64(this.row.Id)
}

// 绝望之地version
func (this *GlobalDespairLandRow) Version() int64 {
	return int64(this.row.Version)
}

// 绝望之地开始时间戳
func (this *GlobalDespairLandRow) Timestamp() int64 {
	return int64(this.row.Timestamp)
}

func (this *GlobalDespairLandRow) GoObject() *GlobalDespairLand {
	return &GlobalDespairLand{
		Id: int64(this.row.Id),
		Version: int64(this.row.Version),
		Timestamp: int64(this.row.Timestamp),
	}
}

/* ========== global_despair_land_battle_record ========== */

// 玩家战报记录
type GlobalDespairLandBattleRecordRow struct {
	row *C.GlobalDespairLandBattleRecord
	isBreak bool
	cached_Record bool
	cacheof_Record []byte
}

func (this *GlobalDespairLandBattleRecordRow) reset(row *C.GlobalDespairLandBattleRecord) {	this.row = row
	this.isBreak = false
	this.cached_Record = false
}

func (this *GlobalDespairLandBattleRecordRow) Break() {
	this.isBreak = true
}

// 主键ID
func (this *GlobalDespairLandBattleRecordRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家id
func (this *GlobalDespairLandBattleRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 玩家战力（当时）
func (this *GlobalDespairLandBattleRecordRow) Fightnum() int32 {
	return int32(this.row.Fightnum)
}

// 战斗时间
func (this *GlobalDespairLandBattleRecordRow) Timestamp() int64 {
	return int64(this.row.Timestamp)
}

// 1-绝望之地最早通关 2-绝望之地最近通关 4-绝望之地最低战力通关
func (this *GlobalDespairLandBattleRecordRow) Tag() int16 {
	return int16(this.row.Tag)
}

// 战场服务端版本
func (this *GlobalDespairLandBattleRecordRow) BattleVersion() int16 {
	return int16(this.row.BattleVersion)
}

// 战场类型
func (this *GlobalDespairLandBattleRecordRow) CampType() int8 {
	return int8(this.row.CampType)
}

// 关卡id
func (this *GlobalDespairLandBattleRecordRow) LevelId() int32 {
	return int32(this.row.LevelId)
}

// 战报
func (this *GlobalDespairLandBattleRecordRow) Record() []byte {
	if !this.cached_Record {
		this.cached_Record = true
		this.cacheof_Record = C.GoBytes(unsafe.Pointer(this.row.Record), this.row.Record_len)
	}
	return this.cacheof_Record
}

func (this *GlobalDespairLandBattleRecordRow) GoObject() *GlobalDespairLandBattleRecord {
	return &GlobalDespairLandBattleRecord{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		Fightnum: int32(this.row.Fightnum),
		Timestamp: int64(this.row.Timestamp),
		Tag: int16(this.row.Tag),
		BattleVersion: int16(this.row.BattleVersion),
		CampType: int8(this.row.CampType),
		LevelId: int32(this.row.LevelId),
		Record: this.Record(),
	}
}

/* ========== global_despair_land_camp ========== */

// 绝望之地阵营状态
type GlobalDespairLandCampRow struct {
	row *C.GlobalDespairLandCamp
	isBreak bool
}

func (this *GlobalDespairLandCampRow) reset(row *C.GlobalDespairLandCamp) {	this.row = row
	this.isBreak = false
}

func (this *GlobalDespairLandCampRow) Break() {
	this.isBreak = true
}

// 阵营
func (this *GlobalDespairLandCampRow) CampType() int8 {
	return int8(this.row.CampType)
}

// boss开始时间戳
func (this *GlobalDespairLandCampRow) Timestamp() int64 {
	return int64(this.row.Timestamp)
}

// boss死亡时间戳
func (this *GlobalDespairLandCampRow) DeadTimestamp() int64 {
	return int64(this.row.DeadTimestamp)
}

// boss等级也即进度等级
func (this *GlobalDespairLandCampRow) Level() int16 {
	return int16(this.row.Level)
}

func (this *GlobalDespairLandCampRow) GoObject() *GlobalDespairLandCamp {
	return &GlobalDespairLandCamp{
		CampType: int8(this.row.CampType),
		Timestamp: int64(this.row.Timestamp),
		DeadTimestamp: int64(this.row.DeadTimestamp),
		Level: int16(this.row.Level),
	}
}

/* ========== global_gift_card_record ========== */

// 礼品卡兑换记录
type GlobalGiftCardRecordRow struct {
	row *C.GlobalGiftCardRecord
	isBreak bool
	cached_Code bool
	cacheof_Code string
}

func (this *GlobalGiftCardRecordRow) reset(row *C.GlobalGiftCardRecord) {	this.row = row
	this.isBreak = false
	this.cached_Code = false
}

func (this *GlobalGiftCardRecordRow) Break() {
	this.isBreak = true
}

// 主键
func (this *GlobalGiftCardRecordRow) Id() int64 {
	return int64(this.row.Id)
}

// 兑换玩家PID
func (this *GlobalGiftCardRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 兑换码
func (this *GlobalGiftCardRecordRow) Code() string {
	if !this.cached_Code {
		this.cached_Code = true
		this.cacheof_Code = C.GoStringN(this.row.Code, this.row.Code_len)
	}
	return this.cacheof_Code
}

// 兑换时间
func (this *GlobalGiftCardRecordRow) Timestamp() int64 {
	return int64(this.row.Timestamp)
}

func (this *GlobalGiftCardRecordRow) GoObject() *GlobalGiftCardRecord {
	return &GlobalGiftCardRecord{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		Code: this.Code(),
		Timestamp: int64(this.row.Timestamp),
	}
}

/* ========== global_group_buy_status ========== */

// hd服务器记录团购状态信息
type GlobalGroupBuyStatusRow struct {
	row *C.GlobalGroupBuyStatus
	isBreak bool
}

func (this *GlobalGroupBuyStatusRow) reset(row *C.GlobalGroupBuyStatus) {	this.row = row
	this.isBreak = false
}

func (this *GlobalGroupBuyStatusRow) Break() {
	this.isBreak = true
}

// ID标识
func (this *GlobalGroupBuyStatusRow) Id() int64 {
	return int64(this.row.Id)
}

// 外键，指定对应得团购物品记录id
func (this *GlobalGroupBuyStatusRow) Cid() int32 {
	return int32(this.row.Cid)
}

// 当前团购状态，即购买总数
func (this *GlobalGroupBuyStatusRow) Status() int32 {
	return int32(this.row.Status)
}

func (this *GlobalGroupBuyStatusRow) GoObject() *GlobalGroupBuyStatus {
	return &GlobalGroupBuyStatus{
		Id: int64(this.row.Id),
		Cid: int32(this.row.Cid),
		Status: int32(this.row.Status),
	}
}

/* ========== global_mail ========== */

// 全局邮件
type GlobalMailRow struct {
	row *C.GlobalMail
	isBreak bool
	cached_Title bool
	cacheof_Title string
	cached_Content bool
	cacheof_Content string
}

func (this *GlobalMailRow) reset(row *C.GlobalMail) {	this.row = row
	this.isBreak = false
	this.cached_Title = false
	this.cached_Content = false
}

func (this *GlobalMailRow) Break() {
	this.isBreak = true
}

// 公告ID
func (this *GlobalMailRow) Id() int64 {
	return int64(this.row.Id)
}

// 发送时间戳
func (this *GlobalMailRow) SendTime() int64 {
	return int64(this.row.SendTime)
}

// 创建时间戳
func (this *GlobalMailRow) ExpireTime() int64 {
	return int64(this.row.ExpireTime)
}

// 标题
func (this *GlobalMailRow) Title() string {
	if !this.cached_Title {
		this.cached_Title = true
		this.cacheof_Title = C.GoStringN(this.row.Title, this.row.Title_len)
	}
	return this.cacheof_Title
}

// 内容
func (this *GlobalMailRow) Content() string {
	if !this.cached_Content {
		this.cached_Content = true
		this.cacheof_Content = C.GoStringN(this.row.Content, this.row.Content_len)
	}
	return this.cacheof_Content
}

// 优先级
func (this *GlobalMailRow) Priority() int8 {
	return int8(this.row.Priority)
}

// 要求最低等级
func (this *GlobalMailRow) MinLevel() int16 {
	return int16(this.row.MinLevel)
}

// 要求最高等级
func (this *GlobalMailRow) MaxLevel() int16 {
	return int16(this.row.MaxLevel)
}

// 要求VIP最低等级
func (this *GlobalMailRow) MinVipLevel() int16 {
	return int16(this.row.MinVipLevel)
}

// 要求VIP最等级
func (this *GlobalMailRow) MaxVipLevel() int16 {
	return int16(this.row.MaxVipLevel)
}

// 要求最早注册时间
func (this *GlobalMailRow) MinCreateTime() int64 {
	return int64(this.row.MinCreateTime)
}

// 要求最晚注册时间
func (this *GlobalMailRow) MaxCreateTime() int64 {
	return int64(this.row.MaxCreateTime)
}

func (this *GlobalMailRow) GoObject() *GlobalMail {
	return &GlobalMail{
		Id: int64(this.row.Id),
		SendTime: int64(this.row.SendTime),
		ExpireTime: int64(this.row.ExpireTime),
		Title: this.Title(),
		Content: this.Content(),
		Priority: int8(this.row.Priority),
		MinLevel: int16(this.row.MinLevel),
		MaxLevel: int16(this.row.MaxLevel),
		MinVipLevel: int16(this.row.MinVipLevel),
		MaxVipLevel: int16(this.row.MaxVipLevel),
		MinCreateTime: int64(this.row.MinCreateTime),
		MaxCreateTime: int64(this.row.MaxCreateTime),
	}
}

/* ========== global_mail_attachments ========== */

// 全局邮件附件
type GlobalMailAttachmentsRow struct {
	row *C.GlobalMailAttachments
	isBreak bool
}

func (this *GlobalMailAttachmentsRow) reset(row *C.GlobalMailAttachments) {	this.row = row
	this.isBreak = false
}

func (this *GlobalMailAttachmentsRow) Break() {
	this.isBreak = true
}

// 附件邮件ID
func (this *GlobalMailAttachmentsRow) Id() int64 {
	return int64(this.row.Id)
}

// 全局邮件表主键
func (this *GlobalMailAttachmentsRow) GlobalMailId() int64 {
	return int64(this.row.GlobalMailId)
}

// 附件ID
func (this *GlobalMailAttachmentsRow) ItemId() int16 {
	return int16(this.row.ItemId)
}

// 附件类型
func (this *GlobalMailAttachmentsRow) AttachmentType() int8 {
	return int8(this.row.AttachmentType)
}

// 数量
func (this *GlobalMailAttachmentsRow) ItemNum() int64 {
	return int64(this.row.ItemNum)
}

func (this *GlobalMailAttachmentsRow) GoObject() *GlobalMailAttachments {
	return &GlobalMailAttachments{
		Id: int64(this.row.Id),
		GlobalMailId: int64(this.row.GlobalMailId),
		ItemId: int16(this.row.ItemId),
		AttachmentType: int8(this.row.AttachmentType),
		ItemNum: int64(this.row.ItemNum),
	}
}

/* ========== global_tb_xxd_onlinecnt ========== */

// 腾讯经分在线玩家统计日志
type GlobalTbXxdOnlinecntRow struct {
	row *C.GlobalTbXxdOnlinecnt
	isBreak bool
	cached_Gameappid bool
	cacheof_Gameappid string
}

func (this *GlobalTbXxdOnlinecntRow) reset(row *C.GlobalTbXxdOnlinecnt) {	this.row = row
	this.isBreak = false
	this.cached_Gameappid = false
}

func (this *GlobalTbXxdOnlinecntRow) Break() {
	this.isBreak = true
}

// 主键ID
func (this *GlobalTbXxdOnlinecntRow) Id() int64 {
	return int64(this.row.Id)
}

// 平台分配的AppID
func (this *GlobalTbXxdOnlinecntRow) Gameappid() string {
	if !this.cached_Gameappid {
		this.cached_Gameappid = true
		this.cacheof_Gameappid = C.GoStringN(this.row.Gameappid, this.row.Gameappid_len)
	}
	return this.cacheof_Gameappid
}

// 当前时间除以60s，下取整
func (this *GlobalTbXxdOnlinecntRow) Timekey() int64 {
	return int64(this.row.Timekey)
}

// 游戏服务器编号
func (this *GlobalTbXxdOnlinecntRow) Gsid() int64 {
	return int64(this.row.Gsid)
}

// ios在线人数
func (this *GlobalTbXxdOnlinecntRow) Onlinecntios() int64 {
	return int64(this.row.Onlinecntios)
}

// android在线人数
func (this *GlobalTbXxdOnlinecntRow) Onlinecntandroid() int64 {
	return int64(this.row.Onlinecntandroid)
}

// 渠道ID
func (this *GlobalTbXxdOnlinecntRow) Cid() int64 {
	return int64(this.row.Cid)
}

func (this *GlobalTbXxdOnlinecntRow) GoObject() *GlobalTbXxdOnlinecnt {
	return &GlobalTbXxdOnlinecnt{
		Id: int64(this.row.Id),
		Gameappid: this.Gameappid(),
		Timekey: int64(this.row.Timekey),
		Gsid: int64(this.row.Gsid),
		Onlinecntios: int64(this.row.Onlinecntios),
		Onlinecntandroid: int64(this.row.Onlinecntandroid),
		Cid: int64(this.row.Cid),
	}
}

/* ========== player ========== */

// 玩家基础信息
type PlayerRow struct {
	row *C.Player
	isBreak bool
	cached_User bool
	cacheof_User string
	cached_Nick bool
	cacheof_Nick string
}

func (this *PlayerRow) reset(row *C.Player) {	this.row = row
	this.isBreak = false
	this.cached_User = false
	this.cached_Nick = false
}

func (this *PlayerRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerRow) Id() int64 {
	return int64(this.row.Id)
}

// 平台传递过来的用户标识
func (this *PlayerRow) User() string {
	if !this.cached_User {
		this.cached_User = true
		this.cacheof_User = C.GoStringN(this.row.User, this.row.User_len)
	}
	return this.cacheof_User
}

// 玩家昵称
func (this *PlayerRow) Nick() string {
	if !this.cached_Nick {
		this.cached_Nick = true
		this.cacheof_Nick = C.GoStringN(this.row.Nick, this.row.Nick_len)
	}
	return this.cacheof_Nick
}

// 主角ID
func (this *PlayerRow) MainRoleId() int64 {
	return int64(this.row.MainRoleId)
}

// 渠道ID
func (this *PlayerRow) Cid() int64 {
	return int64(this.row.Cid)
}

func (this *PlayerRow) GoObject() *Player {
	return &Player{
		Id: int64(this.row.Id),
		User: this.User(),
		Nick: this.Nick(),
		MainRoleId: int64(this.row.MainRoleId),
		Cid: int64(this.row.Cid),
	}
}

/* ========== player_activity ========== */

// 玩家活跃度
type PlayerActivityRow struct {
	row *C.PlayerActivity
	isBreak bool
}

func (this *PlayerActivityRow) reset(row *C.PlayerActivity) {	this.row = row
	this.isBreak = false
}

func (this *PlayerActivityRow) Break() {
	this.isBreak = true
}

// 用户ID
func (this *PlayerActivityRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 用户活跃度
func (this *PlayerActivityRow) Activity() int32 {
	return int32(this.row.Activity)
}

// 最后一次更新时间戳
func (this *PlayerActivityRow) LastUpdate() int64 {
	return int64(this.row.LastUpdate)
}

func (this *PlayerActivityRow) GoObject() *PlayerActivity {
	return &PlayerActivity{
		Pid: int64(this.row.Pid),
		Activity: int32(this.row.Activity),
		LastUpdate: int64(this.row.LastUpdate),
	}
}

/* ========== player_addition_quest ========== */

// 玩家支线任务
type PlayerAdditionQuestRow struct {
	row *C.PlayerAdditionQuest
	isBreak bool
}

func (this *PlayerAdditionQuestRow) reset(row *C.PlayerAdditionQuest) {	this.row = row
	this.isBreak = false
}

func (this *PlayerAdditionQuestRow) Break() {
	this.isBreak = true
}

// ID
func (this *PlayerAdditionQuestRow) Id() int64 {
	return int64(this.row.Id)
}

// 用户ID
func (this *PlayerAdditionQuestRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 任务系列ID
func (this *PlayerAdditionQuestRow) SerialNumber() int32 {
	return int32(this.row.SerialNumber)
}

// 当前任务ID
func (this *PlayerAdditionQuestRow) QuestId() int32 {
	return int32(this.row.QuestId)
}

// 任务链权值
func (this *PlayerAdditionQuestRow) Lock() int16 {
	return int16(this.row.Lock)
}

// 任务进度
func (this *PlayerAdditionQuestRow) Progress() int16 {
	return int16(this.row.Progress)
}

// 0--未完成 1--已完成 2--已奖励 3--已放弃
func (this *PlayerAdditionQuestRow) State() int8 {
	return int8(this.row.State)
}

func (this *PlayerAdditionQuestRow) GoObject() *PlayerAdditionQuest {
	return &PlayerAdditionQuest{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		SerialNumber: int32(this.row.SerialNumber),
		QuestId: int32(this.row.QuestId),
		Lock: int16(this.row.Lock),
		Progress: int16(this.row.Progress),
		State: int8(this.row.State),
	}
}

/* ========== player_any_sdk_order ========== */

// anysdk 订单处理纪录
type PlayerAnySdkOrderRow struct {
	row *C.PlayerAnySdkOrder
	isBreak bool
}

func (this *PlayerAnySdkOrderRow) reset(row *C.PlayerAnySdkOrder) {	this.row = row
	this.isBreak = false
}

func (this *PlayerAnySdkOrderRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerAnySdkOrderRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerAnySdkOrderRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 内部订单ID
func (this *PlayerAnySdkOrderRow) OrderId() int64 {
	return int64(this.row.OrderId)
}

// 赠送元宝元宝
func (this *PlayerAnySdkOrderRow) Present() int64 {
	return int64(this.row.Present)
}

func (this *PlayerAnySdkOrderRow) GoObject() *PlayerAnySdkOrder {
	return &PlayerAnySdkOrder{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		OrderId: int64(this.row.OrderId),
		Present: int64(this.row.Present),
	}
}

/* ========== player_arena ========== */

// 玩家比武场数据
type PlayerArenaRow struct {
	row *C.PlayerArena
	isBreak bool
}

func (this *PlayerArenaRow) reset(row *C.PlayerArena) {	this.row = row
	this.isBreak = false
}

func (this *PlayerArenaRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerArenaRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 今日已挑战次数
func (this *PlayerArenaRow) DailyNum() int16 {
	return int16(this.row.DailyNum)
}

// 战败CD结束时间
func (this *PlayerArenaRow) FailedCdTime() int64 {
	return int64(this.row.FailedCdTime)
}

// 最近一次挑战时间
func (this *PlayerArenaRow) BattleTime() int64 {
	return int64(this.row.BattleTime)
}

// >0 连胜场次; 0 保持不变; -1 下降趋势
func (this *PlayerArenaRow) WinTimes() int16 {
	return int16(this.row.WinTimes)
}

// 今日获得暗影果实累计
func (this *PlayerArenaRow) DailyAwardItem() int32 {
	return int32(this.row.DailyAwardItem)
}

// 新战报计数
func (this *PlayerArenaRow) NewRecordCount() int16 {
	return int16(this.row.NewRecordCount)
}

// 每日奖励声望
func (this *PlayerArenaRow) DailyFame() int32 {
	return int32(this.row.DailyFame)
}

// 今日购买次数
func (this *PlayerArenaRow) BuyTimes() int16 {
	return int16(this.row.BuyTimes)
}

func (this *PlayerArenaRow) GoObject() *PlayerArena {
	return &PlayerArena{
		Pid: int64(this.row.Pid),
		DailyNum: int16(this.row.DailyNum),
		FailedCdTime: int64(this.row.FailedCdTime),
		BattleTime: int64(this.row.BattleTime),
		WinTimes: int16(this.row.WinTimes),
		DailyAwardItem: int32(this.row.DailyAwardItem),
		NewRecordCount: int16(this.row.NewRecordCount),
		DailyFame: int32(this.row.DailyFame),
		BuyTimes: int16(this.row.BuyTimes),
	}
}

/* ========== player_arena_rank ========== */

// 玩家比武场最近排名记录
type PlayerArenaRankRow struct {
	row *C.PlayerArenaRank
	isBreak bool
}

func (this *PlayerArenaRankRow) reset(row *C.PlayerArenaRank) {	this.row = row
	this.isBreak = false
}

func (this *PlayerArenaRankRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerArenaRankRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 昨天排名
func (this *PlayerArenaRankRow) Rank() int32 {
	return int32(this.row.Rank)
}

// 1天前排名
func (this *PlayerArenaRankRow) Rank1() int32 {
	return int32(this.row.Rank1)
}

// 2天前排名
func (this *PlayerArenaRankRow) Rank2() int32 {
	return int32(this.row.Rank2)
}

// 3天前排名
func (this *PlayerArenaRankRow) Rank3() int32 {
	return int32(this.row.Rank3)
}

// 宝箱刷新时间
func (this *PlayerArenaRankRow) Time() int64 {
	return int64(this.row.Time)
}

func (this *PlayerArenaRankRow) GoObject() *PlayerArenaRank {
	return &PlayerArenaRank{
		Pid: int64(this.row.Pid),
		Rank: int32(this.row.Rank),
		Rank1: int32(this.row.Rank1),
		Rank2: int32(this.row.Rank2),
		Rank3: int32(this.row.Rank3),
		Time: int64(this.row.Time),
	}
}

/* ========== player_arena_record ========== */

// 玩家比武场记录
type PlayerArenaRecordRow struct {
	row *C.PlayerArenaRecord
	isBreak bool
	cached_TargetNick bool
	cacheof_TargetNick string
}

func (this *PlayerArenaRecordRow) reset(row *C.PlayerArenaRecord) {	this.row = row
	this.isBreak = false
	this.cached_TargetNick = false
}

func (this *PlayerArenaRecordRow) Break() {
	this.isBreak = true
}

// 记录ID
func (this *PlayerArenaRecordRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerArenaRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 记录类型，0无数据，1挑战成功，2挑战失败，3被挑战且成功，4被挑战且失败
func (this *PlayerArenaRecordRow) Mode() int8 {
	return int8(this.row.Mode)
}

// 原排位
func (this *PlayerArenaRecordRow) OldRank() int32 {
	return int32(this.row.OldRank)
}

// 新排位
func (this *PlayerArenaRecordRow) NewRank() int32 {
	return int32(this.row.NewRank)
}

// 战力
func (this *PlayerArenaRecordRow) FightNum() int32 {
	return int32(this.row.FightNum)
}

// 对手玩家ID
func (this *PlayerArenaRecordRow) TargetPid() int64 {
	return int64(this.row.TargetPid)
}

// 对手昵称
func (this *PlayerArenaRecordRow) TargetNick() string {
	if !this.cached_TargetNick {
		this.cached_TargetNick = true
		this.cacheof_TargetNick = C.GoStringN(this.row.TargetNick, this.row.TargetNick_len)
	}
	return this.cacheof_TargetNick
}

// 对手原排位
func (this *PlayerArenaRecordRow) TargetOldRank() int32 {
	return int32(this.row.TargetOldRank)
}

// 对手新排位
func (this *PlayerArenaRecordRow) TargetNewRank() int32 {
	return int32(this.row.TargetNewRank)
}

// 对手战力
func (this *PlayerArenaRecordRow) TargetFightNum() int32 {
	return int32(this.row.TargetFightNum)
}

// 记录时间
func (this *PlayerArenaRecordRow) RecordTime() int64 {
	return int64(this.row.RecordTime)
}

func (this *PlayerArenaRecordRow) GoObject() *PlayerArenaRecord {
	return &PlayerArenaRecord{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		Mode: int8(this.row.Mode),
		OldRank: int32(this.row.OldRank),
		NewRank: int32(this.row.NewRank),
		FightNum: int32(this.row.FightNum),
		TargetPid: int64(this.row.TargetPid),
		TargetNick: this.TargetNick(),
		TargetOldRank: int32(this.row.TargetOldRank),
		TargetNewRank: int32(this.row.TargetNewRank),
		TargetFightNum: int32(this.row.TargetFightNum),
		RecordTime: int64(this.row.RecordTime),
	}
}

/* ========== player_awaken_graphic ========== */

// 玩家觉醒图谱
type PlayerAwakenGraphicRow struct {
	row *C.PlayerAwakenGraphic
	isBreak bool
}

func (this *PlayerAwakenGraphicRow) reset(row *C.PlayerAwakenGraphic) {	this.row = row
	this.isBreak = false
}

func (this *PlayerAwakenGraphicRow) Break() {
	this.isBreak = true
}

// 主键ID
func (this *PlayerAwakenGraphicRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerAwakenGraphicRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 角色ID
func (this *PlayerAwakenGraphicRow) RoleId() int8 {
	return int8(this.row.RoleId)
}

// 觉醒属性实例ID
func (this *PlayerAwakenGraphicRow) AttrImpl() int16 {
	return int16(this.row.AttrImpl)
}

// 觉醒等级
func (this *PlayerAwakenGraphicRow) Level() int8 {
	return int8(this.row.Level)
}

func (this *PlayerAwakenGraphicRow) GoObject() *PlayerAwakenGraphic {
	return &PlayerAwakenGraphic{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		RoleId: int8(this.row.RoleId),
		AttrImpl: int16(this.row.AttrImpl),
		Level: int8(this.row.Level),
	}
}

/* ========== player_battle_pet ========== */

// 玩家灵宠数据
type PlayerBattlePetRow struct {
	row *C.PlayerBattlePet
	isBreak bool
}

func (this *PlayerBattlePetRow) reset(row *C.PlayerBattlePet) {	this.row = row
	this.isBreak = false
}

func (this *PlayerBattlePetRow) Break() {
	this.isBreak = true
}

// ID
func (this *PlayerBattlePetRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerBattlePetRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 灵宠ID(怪物ID）
func (this *PlayerBattlePetRow) BattlePetId() int32 {
	return int32(this.row.BattlePetId)
}

// 灵宠等级
func (this *PlayerBattlePetRow) Level() int32 {
	return int32(this.row.Level)
}

// 当前经验
func (this *PlayerBattlePetRow) Exp() int64 {
	return int64(this.row.Exp)
}

// 技能等级
func (this *PlayerBattlePetRow) SkillLevel() int32 {
	return int32(this.row.SkillLevel)
}

func (this *PlayerBattlePetRow) GoObject() *PlayerBattlePet {
	return &PlayerBattlePet{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		BattlePetId: int32(this.row.BattlePetId),
		Level: int32(this.row.Level),
		Exp: int64(this.row.Exp),
		SkillLevel: int32(this.row.SkillLevel),
	}
}

/* ========== player_battle_pet_grid ========== */

// 
type PlayerBattlePetGridRow struct {
	row *C.PlayerBattlePetGrid
	isBreak bool
}

func (this *PlayerBattlePetGridRow) reset(row *C.PlayerBattlePetGrid) {	this.row = row
	this.isBreak = false
}

func (this *PlayerBattlePetGridRow) Break() {
	this.isBreak = true
}

// id
func (this *PlayerBattlePetGridRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerBattlePetGridRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 灵宠格子ID
func (this *PlayerBattlePetGridRow) GridId() int8 {
	return int8(this.row.GridId)
}

// 灵宠ID
func (this *PlayerBattlePetGridRow) BattlePetId() int32 {
	return int32(this.row.BattlePetId)
}

func (this *PlayerBattlePetGridRow) GoObject() *PlayerBattlePetGrid {
	return &PlayerBattlePetGrid{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		GridId: int8(this.row.GridId),
		BattlePetId: int32(this.row.BattlePetId),
	}
}

/* ========== player_battle_pet_state ========== */

// 玩家灵宠状态
type PlayerBattlePetStateRow struct {
	row *C.PlayerBattlePetState
	isBreak bool
}

func (this *PlayerBattlePetStateRow) reset(row *C.PlayerBattlePetState) {	this.row = row
	this.isBreak = false
}

func (this *PlayerBattlePetStateRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerBattlePetStateRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 今日使用元宝升级次数
func (this *PlayerBattlePetStateRow) UpgradeByIngotNum() int32 {
	return int32(this.row.UpgradeByIngotNum)
}

// 最近一次使用元宝升级时间
func (this *PlayerBattlePetStateRow) UpgradeByIngotTime() int64 {
	return int64(this.row.UpgradeByIngotTime)
}

func (this *PlayerBattlePetStateRow) GoObject() *PlayerBattlePetState {
	return &PlayerBattlePetState{
		Pid: int64(this.row.Pid),
		UpgradeByIngotNum: int32(this.row.UpgradeByIngotNum),
		UpgradeByIngotTime: int64(this.row.UpgradeByIngotTime),
	}
}

/* ========== player_chest_state ========== */

// 玩家宝箱状态
type PlayerChestStateRow struct {
	row *C.PlayerChestState
	isBreak bool
}

func (this *PlayerChestStateRow) reset(row *C.PlayerChestState) {	this.row = row
	this.isBreak = false
}

func (this *PlayerChestStateRow) Break() {
	this.isBreak = true
}

// 玩家id
func (this *PlayerChestStateRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 每日免费青铜宝箱数
func (this *PlayerChestStateRow) FreeCoinChestNum() int32 {
	return int32(this.row.FreeCoinChestNum)
}

// 上次开免费青铜宝箱时间
func (this *PlayerChestStateRow) LastFreeCoinChestAt() int64 {
	return int64(this.row.LastFreeCoinChestAt)
}

// 今天开青铜宝箱次数
func (this *PlayerChestStateRow) CoinChestNum() int32 {
	return int32(this.row.CoinChestNum)
}

// 今日青铜宝箱十连抽次数
func (this *PlayerChestStateRow) CoinChestTenNum() int32 {
	return int32(this.row.CoinChestTenNum)
}

// 是否第一次青龙宝箱十连抽
func (this *PlayerChestStateRow) IsFirstCoinTen() int8 {
	return int8(this.row.IsFirstCoinTen)
}

// 今天开青铜宝箱花费铜钱数
func (this *PlayerChestStateRow) CoinChestConsume() int64 {
	return int64(this.row.CoinChestConsume)
}

// 上次开消费青铜宝箱时间
func (this *PlayerChestStateRow) LastCoinChestAt() int64 {
	return int64(this.row.LastCoinChestAt)
}

// 上次开免费神龙宝箱时间
func (this *PlayerChestStateRow) LastFreeIngotChestAt() int64 {
	return int64(this.row.LastFreeIngotChestAt)
}

// 今天开神龙宝箱次数
func (this *PlayerChestStateRow) IngotChestNum() int32 {
	return int32(this.row.IngotChestNum)
}

// 今日神龙宝箱十连抽次数
func (this *PlayerChestStateRow) IngotChestTenNum() int32 {
	return int32(this.row.IngotChestTenNum)
}

// 是否第一次神龙宝箱十连抽
func (this *PlayerChestStateRow) IsFirstIngotTen() int8 {
	return int8(this.row.IsFirstIngotTen)
}

// 今天开神龙宝箱花费元宝数
func (this *PlayerChestStateRow) IngotChestConsume() int64 {
	return int64(this.row.IngotChestConsume)
}

// 上次开消费神龙宝箱时间
func (this *PlayerChestStateRow) LastIngotChestAt() int64 {
	return int64(this.row.LastIngotChestAt)
}

// 上次开免费灵兽宝箱时间
func (this *PlayerChestStateRow) LastFreePetChestAt() int64 {
	return int64(this.row.LastFreePetChestAt)
}

func (this *PlayerChestStateRow) GoObject() *PlayerChestState {
	return &PlayerChestState{
		Pid: int64(this.row.Pid),
		FreeCoinChestNum: int32(this.row.FreeCoinChestNum),
		LastFreeCoinChestAt: int64(this.row.LastFreeCoinChestAt),
		CoinChestNum: int32(this.row.CoinChestNum),
		CoinChestTenNum: int32(this.row.CoinChestTenNum),
		IsFirstCoinTen: int8(this.row.IsFirstCoinTen),
		CoinChestConsume: int64(this.row.CoinChestConsume),
		LastCoinChestAt: int64(this.row.LastCoinChestAt),
		LastFreeIngotChestAt: int64(this.row.LastFreeIngotChestAt),
		IngotChestNum: int32(this.row.IngotChestNum),
		IngotChestTenNum: int32(this.row.IngotChestTenNum),
		IsFirstIngotTen: int8(this.row.IsFirstIngotTen),
		IngotChestConsume: int64(this.row.IngotChestConsume),
		LastIngotChestAt: int64(this.row.LastIngotChestAt),
		LastFreePetChestAt: int64(this.row.LastFreePetChestAt),
	}
}

/* ========== player_clique_kongfu_attrib ========== */

// 玩家帮派武功属性加成
type PlayerCliqueKongfuAttribRow struct {
	row *C.PlayerCliqueKongfuAttrib
	isBreak bool
}

func (this *PlayerCliqueKongfuAttribRow) reset(row *C.PlayerCliqueKongfuAttrib) {	this.row = row
	this.isBreak = false
}

func (this *PlayerCliqueKongfuAttribRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerCliqueKongfuAttribRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 生命
func (this *PlayerCliqueKongfuAttribRow) Health() int32 {
	return int32(this.row.Health)
}

// 攻击
func (this *PlayerCliqueKongfuAttribRow) Attack() int32 {
	return int32(this.row.Attack)
}

// 防御
func (this *PlayerCliqueKongfuAttribRow) Defence() int32 {
	return int32(this.row.Defence)
}

func (this *PlayerCliqueKongfuAttribRow) GoObject() *PlayerCliqueKongfuAttrib {
	return &PlayerCliqueKongfuAttrib{
		Pid: int64(this.row.Pid),
		Health: int32(this.row.Health),
		Attack: int32(this.row.Attack),
		Defence: int32(this.row.Defence),
	}
}

/* ========== player_coins ========== */

// 玩家铜币兑换表
type PlayerCoinsRow struct {
	row *C.PlayerCoins
	isBreak bool
}

func (this *PlayerCoinsRow) reset(row *C.PlayerCoins) {	this.row = row
	this.isBreak = false
}

func (this *PlayerCoinsRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerCoinsRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 购买更新时间
func (this *PlayerCoinsRow) BuyTime() int64 {
	return int64(this.row.BuyTime)
}

// 当天购买次数
func (this *PlayerCoinsRow) DailyCount() int16 {
	return int16(this.row.DailyCount)
}

// 玩家批量购买铜币次数
func (this *PlayerCoinsRow) BatchBought() int16 {
	return int16(this.row.BatchBought)
}

func (this *PlayerCoinsRow) GoObject() *PlayerCoins {
	return &PlayerCoins{
		Pid: int64(this.row.Pid),
		BuyTime: int64(this.row.BuyTime),
		DailyCount: int16(this.row.DailyCount),
		BatchBought: int16(this.row.BatchBought),
	}
}

/* ========== player_cornucopia ========== */

// 玩家铜币兑换表
type PlayerCornucopiaRow struct {
	row *C.PlayerCornucopia
	isBreak bool
}

func (this *PlayerCornucopiaRow) reset(row *C.PlayerCornucopia) {	this.row = row
	this.isBreak = false
}

func (this *PlayerCornucopiaRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerCornucopiaRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 开启更新时间
func (this *PlayerCornucopiaRow) OpenTime() int64 {
	return int64(this.row.OpenTime)
}

// 当天开启次数
func (this *PlayerCornucopiaRow) DailyCount() int16 {
	return int16(this.row.DailyCount)
}

func (this *PlayerCornucopiaRow) GoObject() *PlayerCornucopia {
	return &PlayerCornucopia{
		Pid: int64(this.row.Pid),
		OpenTime: int64(this.row.OpenTime),
		DailyCount: int16(this.row.DailyCount),
	}
}

/* ========== player_daily_quest ========== */

// 玩家每日任务
type PlayerDailyQuestRow struct {
	row *C.PlayerDailyQuest
	isBreak bool
}

func (this *PlayerDailyQuestRow) reset(row *C.PlayerDailyQuest) {	this.row = row
	this.isBreak = false
}

func (this *PlayerDailyQuestRow) Break() {
	this.isBreak = true
}

// 记录ID
func (this *PlayerDailyQuestRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerDailyQuestRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 任务ID
func (this *PlayerDailyQuestRow) QuestId() int16 {
	return int16(this.row.QuestId)
}

// 完成数量
func (this *PlayerDailyQuestRow) FinishCount() int16 {
	return int16(this.row.FinishCount)
}

// 最近一次更新时间
func (this *PlayerDailyQuestRow) LastUpdateTime() int64 {
	return int64(this.row.LastUpdateTime)
}

// 奖励状态; 0 未奖励； 1已奖励
func (this *PlayerDailyQuestRow) AwardStatus() int8 {
	return int8(this.row.AwardStatus)
}

// 每日任务类别
func (this *PlayerDailyQuestRow) Class() int16 {
	return int16(this.row.Class)
}

func (this *PlayerDailyQuestRow) GoObject() *PlayerDailyQuest {
	return &PlayerDailyQuest{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		QuestId: int16(this.row.QuestId),
		FinishCount: int16(this.row.FinishCount),
		LastUpdateTime: int64(this.row.LastUpdateTime),
		AwardStatus: int8(this.row.AwardStatus),
		Class: int16(this.row.Class),
	}
}

/* ========== player_daily_quest_star_award_info ========== */

// 玩家每日任务星数奖励状态
type PlayerDailyQuestStarAwardInfoRow struct {
	row *C.PlayerDailyQuestStarAwardInfo
	isBreak bool
	cached_Awarded bool
	cacheof_Awarded string
}

func (this *PlayerDailyQuestStarAwardInfoRow) reset(row *C.PlayerDailyQuestStarAwardInfo) {	this.row = row
	this.isBreak = false
	this.cached_Awarded = false
}

func (this *PlayerDailyQuestStarAwardInfoRow) Break() {
	this.isBreak = true
}

// 玩家id
func (this *PlayerDailyQuestStarAwardInfoRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 当日星数
func (this *PlayerDailyQuestStarAwardInfoRow) Stars() int32 {
	return int32(this.row.Stars)
}

// 最后修改时间
func (this *PlayerDailyQuestStarAwardInfoRow) Lastupdatetime() int64 {
	return int64(this.row.Lastupdatetime)
}

// 已领阶段
func (this *PlayerDailyQuestStarAwardInfoRow) Awarded() string {
	if !this.cached_Awarded {
		this.cached_Awarded = true
		this.cacheof_Awarded = C.GoStringN(this.row.Awarded, this.row.Awarded_len)
	}
	return this.cacheof_Awarded
}

func (this *PlayerDailyQuestStarAwardInfoRow) GoObject() *PlayerDailyQuestStarAwardInfo {
	return &PlayerDailyQuestStarAwardInfo{
		Pid: int64(this.row.Pid),
		Stars: int32(this.row.Stars),
		Lastupdatetime: int64(this.row.Lastupdatetime),
		Awarded: this.Awarded(),
	}
}

/* ========== player_daily_sign_in_record ========== */

// 玩家每日签到记录
type PlayerDailySignInRecordRow struct {
	row *C.PlayerDailySignInRecord
	isBreak bool
}

func (this *PlayerDailySignInRecordRow) reset(row *C.PlayerDailySignInRecord) {	this.row = row
	this.isBreak = false
}

func (this *PlayerDailySignInRecordRow) Break() {
	this.isBreak = true
}

// ID
func (this *PlayerDailySignInRecordRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerDailySignInRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 签到时间
func (this *PlayerDailySignInRecordRow) SignInTime() int64 {
	return int64(this.row.SignInTime)
}

func (this *PlayerDailySignInRecordRow) GoObject() *PlayerDailySignInRecord {
	return &PlayerDailySignInRecord{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		SignInTime: int64(this.row.SignInTime),
	}
}

/* ========== player_daily_sign_in_state ========== */

// 玩家最近七日每日签到状态
type PlayerDailySignInStateRow struct {
	row *C.PlayerDailySignInState
	isBreak bool
}

func (this *PlayerDailySignInStateRow) reset(row *C.PlayerDailySignInState) {	this.row = row
	this.isBreak = false
}

func (this *PlayerDailySignInStateRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerDailySignInStateRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 最新签到时间
func (this *PlayerDailySignInStateRow) UpdateTime() int64 {
	return int64(this.row.UpdateTime)
}

// 签到记录
func (this *PlayerDailySignInStateRow) Record() int16 {
	return int16(this.row.Record)
}

// 今天是否已签到
func (this *PlayerDailySignInStateRow) SignedToday() int8 {
	return int8(this.row.SignedToday)
}

func (this *PlayerDailySignInStateRow) GoObject() *PlayerDailySignInState {
	return &PlayerDailySignInState{
		Pid: int64(this.row.Pid),
		UpdateTime: int64(this.row.UpdateTime),
		Record: int16(this.row.Record),
		SignedToday: int8(this.row.SignedToday),
	}
}

/* ========== player_despair_land_camp_state ========== */

// 玩家绝望之地每周进度
type PlayerDespairLandCampStateRow struct {
	row *C.PlayerDespairLandCampState
	isBreak bool
}

func (this *PlayerDespairLandCampStateRow) reset(row *C.PlayerDespairLandCampState) {	this.row = row
	this.isBreak = false
}

func (this *PlayerDespairLandCampStateRow) Break() {
	this.isBreak = true
}

// 主键标识
func (this *PlayerDespairLandCampStateRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家id
func (this *PlayerDespairLandCampStateRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 阵营
func (this *PlayerDespairLandCampStateRow) CampType() int8 {
	return int8(this.row.CampType)
}

// 记录产生时间戳
func (this *PlayerDespairLandCampStateRow) Timestamp() int64 {
	return int64(this.row.Timestamp)
}

// 周讨伐点
func (this *PlayerDespairLandCampStateRow) Point() int64 {
	return int64(this.row.Point)
}

// 周击杀怪物数
func (this *PlayerDespairLandCampStateRow) KillNum() int64 {
	return int64(this.row.KillNum)
}

// 周角色死亡次数
func (this *PlayerDespairLandCampStateRow) DeadNum() int64 {
	return int64(this.row.DeadNum)
}

// 周boss战总角色死亡次数
func (this *PlayerDespairLandCampStateRow) DeadNumBoss() int64 {
	return int64(this.row.DeadNumBoss)
}

// 对boss伤害
func (this *PlayerDespairLandCampStateRow) Hurt() int64 {
	return int64(this.row.Hurt)
}

// 参与boss战斗次数
func (this *PlayerDespairLandCampStateRow) BossBattleNum() int32 {
	return int32(this.row.BossBattleNum)
}

// 今日参与boss战斗次数
func (this *PlayerDespairLandCampStateRow) DailyBossBattleNum() int32 {
	return int32(this.row.DailyBossBattleNum)
}

// 最近一次boss战斗结束时间戳
func (this *PlayerDespairLandCampStateRow) BossBattleTimestamp() int64 {
	return int64(this.row.BossBattleTimestamp)
}

// 是否领取奖励
func (this *PlayerDespairLandCampStateRow) Awarded() int8 {
	return int8(this.row.Awarded)
}

func (this *PlayerDespairLandCampStateRow) GoObject() *PlayerDespairLandCampState {
	return &PlayerDespairLandCampState{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		CampType: int8(this.row.CampType),
		Timestamp: int64(this.row.Timestamp),
		Point: int64(this.row.Point),
		KillNum: int64(this.row.KillNum),
		DeadNum: int64(this.row.DeadNum),
		DeadNumBoss: int64(this.row.DeadNumBoss),
		Hurt: int64(this.row.Hurt),
		BossBattleNum: int32(this.row.BossBattleNum),
		DailyBossBattleNum: int32(this.row.DailyBossBattleNum),
		BossBattleTimestamp: int64(this.row.BossBattleTimestamp),
		Awarded: int8(this.row.Awarded),
	}
}

/* ========== player_despair_land_camp_state_history ========== */

// 玩家绝望之地每周进度
type PlayerDespairLandCampStateHistoryRow struct {
	row *C.PlayerDespairLandCampStateHistory
	isBreak bool
}

func (this *PlayerDespairLandCampStateHistoryRow) reset(row *C.PlayerDespairLandCampStateHistory) {	this.row = row
	this.isBreak = false
}

func (this *PlayerDespairLandCampStateHistoryRow) Break() {
	this.isBreak = true
}

// 主键标识
func (this *PlayerDespairLandCampStateHistoryRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家id
func (this *PlayerDespairLandCampStateHistoryRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 阵营
func (this *PlayerDespairLandCampStateHistoryRow) CampType() int8 {
	return int8(this.row.CampType)
}

// 记录产生时间戳
func (this *PlayerDespairLandCampStateHistoryRow) Timestamp() int64 {
	return int64(this.row.Timestamp)
}

// 周讨伐点
func (this *PlayerDespairLandCampStateHistoryRow) Point() int64 {
	return int64(this.row.Point)
}

// 周击杀怪物数
func (this *PlayerDespairLandCampStateHistoryRow) KillNum() int64 {
	return int64(this.row.KillNum)
}

// 周角色死亡次数
func (this *PlayerDespairLandCampStateHistoryRow) DeadNum() int64 {
	return int64(this.row.DeadNum)
}

// 周boss战总角色死亡次数
func (this *PlayerDespairLandCampStateHistoryRow) DeadNumBoss() int64 {
	return int64(this.row.DeadNumBoss)
}

// 对boss伤害
func (this *PlayerDespairLandCampStateHistoryRow) Hurt() int64 {
	return int64(this.row.Hurt)
}

// 参与boss战斗次数
func (this *PlayerDespairLandCampStateHistoryRow) BossBattleNum() int32 {
	return int32(this.row.BossBattleNum)
}

// 今日参与boss战斗次数
func (this *PlayerDespairLandCampStateHistoryRow) DailyBossBattleNum() int32 {
	return int32(this.row.DailyBossBattleNum)
}

// 最近一次boss战斗结束时间戳
func (this *PlayerDespairLandCampStateHistoryRow) BossBattleTimestamp() int64 {
	return int64(this.row.BossBattleTimestamp)
}

// 是否领取奖励
func (this *PlayerDespairLandCampStateHistoryRow) Awarded() int8 {
	return int8(this.row.Awarded)
}

func (this *PlayerDespairLandCampStateHistoryRow) GoObject() *PlayerDespairLandCampStateHistory {
	return &PlayerDespairLandCampStateHistory{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		CampType: int8(this.row.CampType),
		Timestamp: int64(this.row.Timestamp),
		Point: int64(this.row.Point),
		KillNum: int64(this.row.KillNum),
		DeadNum: int64(this.row.DeadNum),
		DeadNumBoss: int64(this.row.DeadNumBoss),
		Hurt: int64(this.row.Hurt),
		BossBattleNum: int32(this.row.BossBattleNum),
		DailyBossBattleNum: int32(this.row.DailyBossBattleNum),
		BossBattleTimestamp: int64(this.row.BossBattleTimestamp),
		Awarded: int8(this.row.Awarded),
	}
}

/* ========== player_despair_land_level_record ========== */

// 玩家绝望之地状态
type PlayerDespairLandLevelRecordRow struct {
	row *C.PlayerDespairLandLevelRecord
	isBreak bool
}

func (this *PlayerDespairLandLevelRecordRow) reset(row *C.PlayerDespairLandLevelRecord) {	this.row = row
	this.isBreak = false
}

func (this *PlayerDespairLandLevelRecordRow) Break() {
	this.isBreak = true
}

// 主键ID
func (this *PlayerDespairLandLevelRecordRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家id
func (this *PlayerDespairLandLevelRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 阵营
func (this *PlayerDespairLandLevelRecordRow) CampType() int8 {
	return int8(this.row.CampType)
}

// 记录产生时间戳
func (this *PlayerDespairLandLevelRecordRow) Timestamp() int64 {
	return int64(this.row.Timestamp)
}

// 关卡ID
func (this *PlayerDespairLandLevelRecordRow) LevelId() int32 {
	return int32(this.row.LevelId)
}

// 最小通关回合数
func (this *PlayerDespairLandLevelRecordRow) Round() int8 {
	return int8(this.row.Round)
}

// 最近通关时间
func (this *PlayerDespairLandLevelRecordRow) PassedTimestamp() int64 {
	return int64(this.row.PassedTimestamp)
}

// 首次通关时间
func (this *PlayerDespairLandLevelRecordRow) FirstPassedTimestamp() int64 {
	return int64(this.row.FirstPassedTimestamp)
}

// 首次通关战力
func (this *PlayerDespairLandLevelRecordRow) FirstPassedFightnum() int32 {
	return int32(this.row.FirstPassedFightnum)
}

// 首次通奖励
func (this *PlayerDespairLandLevelRecordRow) PassedAward() int8 {
	return int8(this.row.PassedAward)
}

// 三星通奖励
func (this *PlayerDespairLandLevelRecordRow) PerfectAward() int8 {
	return int8(this.row.PerfectAward)
}

// 阶段通奖励
func (this *PlayerDespairLandLevelRecordRow) MilestoneAward() int8 {
	return int8(this.row.MilestoneAward)
}

func (this *PlayerDespairLandLevelRecordRow) GoObject() *PlayerDespairLandLevelRecord {
	return &PlayerDespairLandLevelRecord{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		CampType: int8(this.row.CampType),
		Timestamp: int64(this.row.Timestamp),
		LevelId: int32(this.row.LevelId),
		Round: int8(this.row.Round),
		PassedTimestamp: int64(this.row.PassedTimestamp),
		FirstPassedTimestamp: int64(this.row.FirstPassedTimestamp),
		FirstPassedFightnum: int32(this.row.FirstPassedFightnum),
		PassedAward: int8(this.row.PassedAward),
		PerfectAward: int8(this.row.PerfectAward),
		MilestoneAward: int8(this.row.MilestoneAward),
	}
}

/* ========== player_despair_land_state ========== */

// 玩家绝望之地状态
type PlayerDespairLandStateRow struct {
	row *C.PlayerDespairLandState
	isBreak bool
}

func (this *PlayerDespairLandStateRow) reset(row *C.PlayerDespairLandState) {	this.row = row
	this.isBreak = false
}

func (this *PlayerDespairLandStateRow) Break() {
	this.isBreak = true
}

// 玩家id
func (this *PlayerDespairLandStateRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 总讨伐点
func (this *PlayerDespairLandStateRow) Point() int64 {
	return int64(this.row.Point)
}

// 总击杀怪物数
func (this *PlayerDespairLandStateRow) KillNum() int64 {
	return int64(this.row.KillNum)
}

// 总阵亡次数
func (this *PlayerDespairLandStateRow) DeadNum() int64 {
	return int64(this.row.DeadNum)
}

// 今日讨伐次数
func (this *PlayerDespairLandStateRow) DailyBattleNum() int16 {
	return int16(this.row.DailyBattleNum)
}

// 今日讨伐时间
func (this *PlayerDespairLandStateRow) DailyBattleTimestamp() int64 {
	return int64(this.row.DailyBattleTimestamp)
}

// 今日购买讨伐次数
func (this *PlayerDespairLandStateRow) DailyBoughtBattleNum() int16 {
	return int16(this.row.DailyBoughtBattleNum)
}

// 今日购买讨伐次数时间
func (this *PlayerDespairLandStateRow) DailyBoughtTimestamp() int64 {
	return int64(this.row.DailyBoughtTimestamp)
}

// 今日购买boss讨伐次数时间
func (this *PlayerDespairLandStateRow) DailyBossBoughtTimestamp() int64 {
	return int64(this.row.DailyBossBoughtTimestamp)
}

// 今日购买血兽boss讨伐次数
func (this *PlayerDespairLandStateRow) DailyBossBeastBoughtBattleNum() int16 {
	return int16(this.row.DailyBossBeastBoughtBattleNum)
}

// 今日购买尸鬼boss讨伐次数
func (this *PlayerDespairLandStateRow) DailyBossWalkingDeadBoughtBattleNum() int16 {
	return int16(this.row.DailyBossWalkingDeadBoughtBattleNum)
}

// 今日购买影魔boss讨伐次数
func (this *PlayerDespairLandStateRow) DailyBossDevilBoughtBattleNum() int16 {
	return int16(this.row.DailyBossDevilBoughtBattleNum)
}

func (this *PlayerDespairLandStateRow) GoObject() *PlayerDespairLandState {
	return &PlayerDespairLandState{
		Pid: int64(this.row.Pid),
		Point: int64(this.row.Point),
		KillNum: int64(this.row.KillNum),
		DeadNum: int64(this.row.DeadNum),
		DailyBattleNum: int16(this.row.DailyBattleNum),
		DailyBattleTimestamp: int64(this.row.DailyBattleTimestamp),
		DailyBoughtBattleNum: int16(this.row.DailyBoughtBattleNum),
		DailyBoughtTimestamp: int64(this.row.DailyBoughtTimestamp),
		DailyBossBoughtTimestamp: int64(this.row.DailyBossBoughtTimestamp),
		DailyBossBeastBoughtBattleNum: int16(this.row.DailyBossBeastBoughtBattleNum),
		DailyBossWalkingDeadBoughtBattleNum: int16(this.row.DailyBossWalkingDeadBoughtBattleNum),
		DailyBossDevilBoughtBattleNum: int16(this.row.DailyBossDevilBoughtBattleNum),
	}
}

/* ========== player_driving_sword_event ========== */

// 玩家云海事件列表
type PlayerDrivingSwordEventRow struct {
	row *C.PlayerDrivingSwordEvent
	isBreak bool
}

func (this *PlayerDrivingSwordEventRow) reset(row *C.PlayerDrivingSwordEvent) {	this.row = row
	this.isBreak = false
}

func (this *PlayerDrivingSwordEventRow) Break() {
	this.isBreak = true
}

// 主键ID
func (this *PlayerDrivingSwordEventRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerDrivingSwordEventRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 云层id
func (this *PlayerDrivingSwordEventRow) CloudId() int16 {
	return int16(this.row.CloudId)
}

// 事件坐标x
func (this *PlayerDrivingSwordEventRow) X() int8 {
	return int8(this.row.X)
}

// 事件坐标y
func (this *PlayerDrivingSwordEventRow) Y() int8 {
	return int8(this.row.Y)
}

// 事件类型
func (this *PlayerDrivingSwordEventRow) EventType() int8 {
	return int8(this.row.EventType)
}

// 事件模版数据id
func (this *PlayerDrivingSwordEventRow) DataId() int8 {
	return int8(this.row.DataId)
}

func (this *PlayerDrivingSwordEventRow) GoObject() *PlayerDrivingSwordEvent {
	return &PlayerDrivingSwordEvent{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		CloudId: int16(this.row.CloudId),
		X: int8(this.row.X),
		Y: int8(this.row.Y),
		EventType: int8(this.row.EventType),
		DataId: int8(this.row.DataId),
	}
}

/* ========== player_driving_sword_event_exploring ========== */

// 玩家云海探险事件信息
type PlayerDrivingSwordEventExploringRow struct {
	row *C.PlayerDrivingSwordEventExploring
	isBreak bool
}

func (this *PlayerDrivingSwordEventExploringRow) reset(row *C.PlayerDrivingSwordEventExploring) {	this.row = row
	this.isBreak = false
}

func (this *PlayerDrivingSwordEventExploringRow) Break() {
	this.isBreak = true
}

// 主键ID
func (this *PlayerDrivingSwordEventExploringRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerDrivingSwordEventExploringRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 事件状态
func (this *PlayerDrivingSwordEventExploringRow) Status() int8 {
	return int8(this.row.Status)
}

// 驻守开始时间
func (this *PlayerDrivingSwordEventExploringRow) GarrisonStart() int64 {
	return int64(this.row.GarrisonStart)
}

// 已驻守时间
func (this *PlayerDrivingSwordEventExploringRow) GarrisonTime() int64 {
	return int64(this.row.GarrisonTime)
}

// 已领奖时间
func (this *PlayerDrivingSwordEventExploringRow) AwardTime() int64 {
	return int64(this.row.AwardTime)
}

// 派驻角色id
func (this *PlayerDrivingSwordEventExploringRow) RoleId() int8 {
	return int8(this.row.RoleId)
}

// 云层id
func (this *PlayerDrivingSwordEventExploringRow) CloudId() int16 {
	return int16(this.row.CloudId)
}

// 事件坐标x
func (this *PlayerDrivingSwordEventExploringRow) X() int8 {
	return int8(this.row.X)
}

// 事件坐标y
func (this *PlayerDrivingSwordEventExploringRow) Y() int8 {
	return int8(this.row.Y)
}

// 事件模版数据id
func (this *PlayerDrivingSwordEventExploringRow) DataId() int8 {
	return int8(this.row.DataId)
}

func (this *PlayerDrivingSwordEventExploringRow) GoObject() *PlayerDrivingSwordEventExploring {
	return &PlayerDrivingSwordEventExploring{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		Status: int8(this.row.Status),
		GarrisonStart: int64(this.row.GarrisonStart),
		GarrisonTime: int64(this.row.GarrisonTime),
		AwardTime: int64(this.row.AwardTime),
		RoleId: int8(this.row.RoleId),
		CloudId: int16(this.row.CloudId),
		X: int8(this.row.X),
		Y: int8(this.row.Y),
		DataId: int8(this.row.DataId),
	}
}

/* ========== player_driving_sword_event_treasure ========== */

// 玩家云海宝藏事件信息
type PlayerDrivingSwordEventTreasureRow struct {
	row *C.PlayerDrivingSwordEventTreasure
	isBreak bool
}

func (this *PlayerDrivingSwordEventTreasureRow) reset(row *C.PlayerDrivingSwordEventTreasure) {	this.row = row
	this.isBreak = false
}

func (this *PlayerDrivingSwordEventTreasureRow) Break() {
	this.isBreak = true
}

// 主键ID
func (this *PlayerDrivingSwordEventTreasureRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerDrivingSwordEventTreasureRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 开箱进度
func (this *PlayerDrivingSwordEventTreasureRow) Progress() int8 {
	return int8(this.row.Progress)
}

// 云层id
func (this *PlayerDrivingSwordEventTreasureRow) CloudId() int16 {
	return int16(this.row.CloudId)
}

// 事件坐标x
func (this *PlayerDrivingSwordEventTreasureRow) X() int8 {
	return int8(this.row.X)
}

// 事件坐标y
func (this *PlayerDrivingSwordEventTreasureRow) Y() int8 {
	return int8(this.row.Y)
}

// 事件模版数据id
func (this *PlayerDrivingSwordEventTreasureRow) DataId() int8 {
	return int8(this.row.DataId)
}

func (this *PlayerDrivingSwordEventTreasureRow) GoObject() *PlayerDrivingSwordEventTreasure {
	return &PlayerDrivingSwordEventTreasure{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		Progress: int8(this.row.Progress),
		CloudId: int16(this.row.CloudId),
		X: int8(this.row.X),
		Y: int8(this.row.Y),
		DataId: int8(this.row.DataId),
	}
}

/* ========== player_driving_sword_event_visiting ========== */

// 玩家云海拜访事件信息
type PlayerDrivingSwordEventVisitingRow struct {
	row *C.PlayerDrivingSwordEventVisiting
	isBreak bool
	cached_TargetSideState bool
	cacheof_TargetSideState []byte
	cached_TargetStatus bool
	cacheof_TargetStatus string
}

func (this *PlayerDrivingSwordEventVisitingRow) reset(row *C.PlayerDrivingSwordEventVisiting) {	this.row = row
	this.isBreak = false
	this.cached_TargetSideState = false
	this.cached_TargetStatus = false
}

func (this *PlayerDrivingSwordEventVisitingRow) Break() {
	this.isBreak = true
}

// 主键ID
func (this *PlayerDrivingSwordEventVisitingRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerDrivingSwordEventVisitingRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 事件状态,1代表已经通过还没领奖，2代表已领奖
func (this *PlayerDrivingSwordEventVisitingRow) Status() int8 {
	return int8(this.row.Status)
}

// 拜访的玩家
func (this *PlayerDrivingSwordEventVisitingRow) TargetPid() int64 {
	return int64(this.row.TargetPid)
}

// 拜访玩家战斗状态记录
func (this *PlayerDrivingSwordEventVisitingRow) TargetSideState() []byte {
	if !this.cached_TargetSideState {
		this.cached_TargetSideState = true
		this.cacheof_TargetSideState = C.GoBytes(unsafe.Pointer(this.row.TargetSideState), this.row.TargetSideState_len)
	}
	return this.cacheof_TargetSideState
}

// 云层id
func (this *PlayerDrivingSwordEventVisitingRow) CloudId() int16 {
	return int16(this.row.CloudId)
}

// 事件坐标x
func (this *PlayerDrivingSwordEventVisitingRow) X() int8 {
	return int8(this.row.X)
}

// 事件坐标y
func (this *PlayerDrivingSwordEventVisitingRow) Y() int8 {
	return int8(this.row.Y)
}

// 事件模版数据id
func (this *PlayerDrivingSwordEventVisitingRow) DataId() int8 {
	return int8(this.row.DataId)
}

// 拜访玩家相关信息记录
func (this *PlayerDrivingSwordEventVisitingRow) TargetStatus() string {
	if !this.cached_TargetStatus {
		this.cached_TargetStatus = true
		this.cacheof_TargetStatus = C.GoStringN(this.row.TargetStatus, this.row.TargetStatus_len)
	}
	return this.cacheof_TargetStatus
}

func (this *PlayerDrivingSwordEventVisitingRow) GoObject() *PlayerDrivingSwordEventVisiting {
	return &PlayerDrivingSwordEventVisiting{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		Status: int8(this.row.Status),
		TargetPid: int64(this.row.TargetPid),
		TargetSideState: this.TargetSideState(),
		CloudId: int16(this.row.CloudId),
		X: int8(this.row.X),
		Y: int8(this.row.Y),
		DataId: int8(this.row.DataId),
		TargetStatus: this.TargetStatus(),
	}
}

/* ========== player_driving_sword_eventmask ========== */

// 玩家云海事件地图
type PlayerDrivingSwordEventmaskRow struct {
	row *C.PlayerDrivingSwordEventmask
	isBreak bool
	cached_Events bool
	cacheof_Events []byte
}

func (this *PlayerDrivingSwordEventmaskRow) reset(row *C.PlayerDrivingSwordEventmask) {	this.row = row
	this.isBreak = false
	this.cached_Events = false
}

func (this *PlayerDrivingSwordEventmaskRow) Break() {
	this.isBreak = true
}

// 主键ID
func (this *PlayerDrivingSwordEventmaskRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerDrivingSwordEventmaskRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 云层id
func (this *PlayerDrivingSwordEventmaskRow) CloudId() int16 {
	return int16(this.row.CloudId)
}

// 事件比特阵列
func (this *PlayerDrivingSwordEventmaskRow) Events() []byte {
	if !this.cached_Events {
		this.cached_Events = true
		this.cacheof_Events = C.GoBytes(unsafe.Pointer(this.row.Events), this.row.Events_len)
	}
	return this.cacheof_Events
}

func (this *PlayerDrivingSwordEventmaskRow) GoObject() *PlayerDrivingSwordEventmask {
	return &PlayerDrivingSwordEventmask{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		CloudId: int16(this.row.CloudId),
		Events: this.Events(),
	}
}

/* ========== player_driving_sword_info ========== */

// 玩家御剑基本数据
type PlayerDrivingSwordInfoRow struct {
	row *C.PlayerDrivingSwordInfo
	isBreak bool
}

func (this *PlayerDrivingSwordInfoRow) reset(row *C.PlayerDrivingSwordInfo) {	this.row = row
	this.isBreak = false
}

func (this *PlayerDrivingSwordInfoRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerDrivingSwordInfoRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 当前所在云层
func (this *PlayerDrivingSwordInfoRow) CurrentCloud() int16 {
	return int16(this.row.CurrentCloud)
}

// 最高开启云层
func (this *PlayerDrivingSwordInfoRow) HighestCloud() int16 {
	return int16(this.row.HighestCloud)
}

// 当前坐标x
func (this *PlayerDrivingSwordInfoRow) CurrentX() int8 {
	return int8(this.row.CurrentX)
}

// 当前坐标y
func (this *PlayerDrivingSwordInfoRow) CurrentY() int8 {
	return int8(this.row.CurrentY)
}

// 行动点
func (this *PlayerDrivingSwordInfoRow) AllowedAction() int16 {
	return int16(this.row.AllowedAction)
}

// 行动点刷新时间
func (this *PlayerDrivingSwordInfoRow) ActionRefreshTime() int64 {
	return int64(this.row.ActionRefreshTime)
}

// 行动点购买更新时间
func (this *PlayerDrivingSwordInfoRow) ActionBuyTime() int64 {
	return int64(this.row.ActionBuyTime)
}

// 行动点当天购买次数
func (this *PlayerDrivingSwordInfoRow) DailyActionBought() int8 {
	return int8(this.row.DailyActionBought)
}

func (this *PlayerDrivingSwordInfoRow) GoObject() *PlayerDrivingSwordInfo {
	return &PlayerDrivingSwordInfo{
		Pid: int64(this.row.Pid),
		CurrentCloud: int16(this.row.CurrentCloud),
		HighestCloud: int16(this.row.HighestCloud),
		CurrentX: int8(this.row.CurrentX),
		CurrentY: int8(this.row.CurrentY),
		AllowedAction: int16(this.row.AllowedAction),
		ActionRefreshTime: int64(this.row.ActionRefreshTime),
		ActionBuyTime: int64(this.row.ActionBuyTime),
		DailyActionBought: int8(this.row.DailyActionBought),
	}
}

/* ========== player_driving_sword_map ========== */

// 玩家云海地图
type PlayerDrivingSwordMapRow struct {
	row *C.PlayerDrivingSwordMap
	isBreak bool
	cached_Shadows bool
	cacheof_Shadows []byte
}

func (this *PlayerDrivingSwordMapRow) reset(row *C.PlayerDrivingSwordMap) {	this.row = row
	this.isBreak = false
	this.cached_Shadows = false
}

func (this *PlayerDrivingSwordMapRow) Break() {
	this.isBreak = true
}

// 主键ID
func (this *PlayerDrivingSwordMapRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerDrivingSwordMapRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 云层id
func (this *PlayerDrivingSwordMapRow) CloudId() int16 {
	return int16(this.row.CloudId)
}

// 阴影比特阵列
func (this *PlayerDrivingSwordMapRow) Shadows() []byte {
	if !this.cached_Shadows {
		this.cached_Shadows = true
		this.cacheof_Shadows = C.GoBytes(unsafe.Pointer(this.row.Shadows), this.row.Shadows_len)
	}
	return this.cacheof_Shadows
}

// 障碍总数
func (this *PlayerDrivingSwordMapRow) ObstacleCount() int8 {
	return int8(this.row.ObstacleCount)
}

// 探险总数
func (this *PlayerDrivingSwordMapRow) ExploringCount() int8 {
	return int8(this.row.ExploringCount)
}

// 拜访总数
func (this *PlayerDrivingSwordMapRow) VisitingCount() int8 {
	return int8(this.row.VisitingCount)
}

// 宝藏总数
func (this *PlayerDrivingSwordMapRow) TreasureCount() int8 {
	return int8(this.row.TreasureCount)
}

func (this *PlayerDrivingSwordMapRow) GoObject() *PlayerDrivingSwordMap {
	return &PlayerDrivingSwordMap{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		CloudId: int16(this.row.CloudId),
		Shadows: this.Shadows(),
		ObstacleCount: int8(this.row.ObstacleCount),
		ExploringCount: int8(this.row.ExploringCount),
		VisitingCount: int8(this.row.VisitingCount),
		TreasureCount: int8(this.row.TreasureCount),
	}
}

/* ========== player_equipment ========== */

// 玩家装备表
type PlayerEquipmentRow struct {
	row *C.PlayerEquipment
	isBreak bool
}

func (this *PlayerEquipmentRow) reset(row *C.PlayerEquipment) {	this.row = row
	this.isBreak = false
}

func (this *PlayerEquipmentRow) Break() {
	this.isBreak = true
}

// 主键ID
func (this *PlayerEquipmentRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerEquipmentRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 角色ID
func (this *PlayerEquipmentRow) RoleId() int8 {
	return int8(this.row.RoleId)
}

// 武器的player_item表主键ID
func (this *PlayerEquipmentRow) WeaponId() int64 {
	return int64(this.row.WeaponId)
}

// 战袍的player_item表主键ID
func (this *PlayerEquipmentRow) ClothesId() int64 {
	return int64(this.row.ClothesId)
}

// 饰品的player_item表主键ID
func (this *PlayerEquipmentRow) AccessoriesId() int64 {
	return int64(this.row.AccessoriesId)
}

// 靴子的player_item表主键ID
func (this *PlayerEquipmentRow) ShoeId() int64 {
	return int64(this.row.ShoeId)
}

func (this *PlayerEquipmentRow) GoObject() *PlayerEquipment {
	return &PlayerEquipment{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		RoleId: int8(this.row.RoleId),
		WeaponId: int64(this.row.WeaponId),
		ClothesId: int64(this.row.ClothesId),
		AccessoriesId: int64(this.row.AccessoriesId),
		ShoeId: int64(this.row.ShoeId),
	}
}

/* ========== player_event_award_record ========== */

// 玩家活动奖励领取记录
type PlayerEventAwardRecordRow struct {
	row *C.PlayerEventAwardRecord
	isBreak bool
	cached_RecordBytes bool
	cacheof_RecordBytes []byte
	cached_JsonEventRecord bool
	cacheof_JsonEventRecord []byte
}

func (this *PlayerEventAwardRecordRow) reset(row *C.PlayerEventAwardRecord) {	this.row = row
	this.isBreak = false
	this.cached_RecordBytes = false
	this.cached_JsonEventRecord = false
}

func (this *PlayerEventAwardRecordRow) Break() {
	this.isBreak = true
}

// 用户ID
func (this *PlayerEventAwardRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 奖励领取状态
func (this *PlayerEventAwardRecordRow) RecordBytes() []byte {
	if !this.cached_RecordBytes {
		this.cached_RecordBytes = true
		this.cacheof_RecordBytes = C.GoBytes(unsafe.Pointer(this.row.RecordBytes), this.row.RecordBytes_len)
	}
	return this.cacheof_RecordBytes
}

// json模板配置的玩家活动状态
func (this *PlayerEventAwardRecordRow) JsonEventRecord() []byte {
	if !this.cached_JsonEventRecord {
		this.cached_JsonEventRecord = true
		this.cacheof_JsonEventRecord = C.GoBytes(unsafe.Pointer(this.row.JsonEventRecord), this.row.JsonEventRecord_len)
	}
	return this.cacheof_JsonEventRecord
}

func (this *PlayerEventAwardRecordRow) GoObject() *PlayerEventAwardRecord {
	return &PlayerEventAwardRecord{
		Pid: int64(this.row.Pid),
		RecordBytes: this.RecordBytes(),
		JsonEventRecord: this.JsonEventRecord(),
	}
}

/* ========== player_event_daily_online ========== */

// 
type PlayerEventDailyOnlineRow struct {
	row *C.PlayerEventDailyOnline
	isBreak bool
}

func (this *PlayerEventDailyOnlineRow) reset(row *C.PlayerEventDailyOnline) {	this.row = row
	this.isBreak = false
}

func (this *PlayerEventDailyOnlineRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerEventDailyOnlineRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 记录创建时间戳
func (this *PlayerEventDailyOnlineRow) Timestamp() int64 {
	return int64(this.row.Timestamp)
}

// 今日累计在线时长
func (this *PlayerEventDailyOnlineRow) TotalOnlineTime() int32 {
	return int32(this.row.TotalOnlineTime)
}

// 今日已奖励在线时长
func (this *PlayerEventDailyOnlineRow) AwardedOnlinetime() int32 {
	return int32(this.row.AwardedOnlinetime)
}

func (this *PlayerEventDailyOnlineRow) GoObject() *PlayerEventDailyOnline {
	return &PlayerEventDailyOnline{
		Pid: int64(this.row.Pid),
		Timestamp: int64(this.row.Timestamp),
		TotalOnlineTime: int32(this.row.TotalOnlineTime),
		AwardedOnlinetime: int32(this.row.AwardedOnlinetime),
	}
}

/* ========== player_event_vn_daily_recharge ========== */

// 
type PlayerEventVnDailyRechargeRow struct {
	row *C.PlayerEventVnDailyRecharge
	isBreak bool
}

func (this *PlayerEventVnDailyRechargeRow) reset(row *C.PlayerEventVnDailyRecharge) {	this.row = row
	this.isBreak = false
}

func (this *PlayerEventVnDailyRechargeRow) Break() {
	this.isBreak = true
}

// 
func (this *PlayerEventVnDailyRechargeRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerEventVnDailyRechargeRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 玩家ID
func (this *PlayerEventVnDailyRechargeRow) Page() int32 {
	return int32(this.row.Page)
}

// 记录创建时间戳
func (this *PlayerEventVnDailyRechargeRow) Timestamp() int64 {
	return int64(this.row.Timestamp)
}

// 今日是否已领取奖励
func (this *PlayerEventVnDailyRechargeRow) Awarded() int8 {
	return int8(this.row.Awarded)
}

// 今日累计充值
func (this *PlayerEventVnDailyRechargeRow) DailyRecharge() int64 {
	return int64(this.row.DailyRecharge)
}

// 活动累计充值
func (this *PlayerEventVnDailyRechargeRow) TotalRecharge() int64 {
	return int64(this.row.TotalRecharge)
}

func (this *PlayerEventVnDailyRechargeRow) GoObject() *PlayerEventVnDailyRecharge {
	return &PlayerEventVnDailyRecharge{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		Page: int32(this.row.Page),
		Timestamp: int64(this.row.Timestamp),
		Awarded: int8(this.row.Awarded),
		DailyRecharge: int64(this.row.DailyRecharge),
		TotalRecharge: int64(this.row.TotalRecharge),
	}
}

/* ========== player_event_vn_dragon_ball_record ========== */

// 越南龙珠活动玩家获取记录
type PlayerEventVnDragonBallRecordRow struct {
	row *C.PlayerEventVnDragonBallRecord
	isBreak bool
}

func (this *PlayerEventVnDragonBallRecordRow) reset(row *C.PlayerEventVnDragonBallRecord) {	this.row = row
	this.isBreak = false
}

func (this *PlayerEventVnDragonBallRecordRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerEventVnDragonBallRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 获取的物品ID
func (this *PlayerEventVnDragonBallRecordRow) ItemId() int16 {
	return int16(this.row.ItemId)
}

// 获取时间戳
func (this *PlayerEventVnDragonBallRecordRow) Createtime() int64 {
	return int64(this.row.Createtime)
}

func (this *PlayerEventVnDragonBallRecordRow) GoObject() *PlayerEventVnDragonBallRecord {
	return &PlayerEventVnDragonBallRecord{
		Pid: int64(this.row.Pid),
		ItemId: int16(this.row.ItemId),
		Createtime: int64(this.row.Createtime),
	}
}

/* ========== player_events_fight_power ========== */

// 玩家战力运营活动记录
type PlayerEventsFightPowerRow struct {
	row *C.PlayerEventsFightPower
	isBreak bool
}

func (this *PlayerEventsFightPowerRow) reset(row *C.PlayerEventsFightPower) {	this.row = row
	this.isBreak = false
}

func (this *PlayerEventsFightPowerRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerEventsFightPowerRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 当前已奖励的战力档位
func (this *PlayerEventsFightPowerRow) Lock() int32 {
	return int32(this.row.Lock)
}

func (this *PlayerEventsFightPowerRow) GoObject() *PlayerEventsFightPower {
	return &PlayerEventsFightPower{
		Pid: int64(this.row.Pid),
		Lock: int32(this.row.Lock),
	}
}

/* ========== player_events_ingot_record ========== */

// 玩家元宝充值和消耗活动记录
type PlayerEventsIngotRecordRow struct {
	row *C.PlayerEventsIngotRecord
	isBreak bool
}

func (this *PlayerEventsIngotRecordRow) reset(row *C.PlayerEventsIngotRecord) {	this.row = row
	this.isBreak = false
}

func (this *PlayerEventsIngotRecordRow) Break() {
	this.isBreak = true
}

// 
func (this *PlayerEventsIngotRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 充值元宝总数
func (this *PlayerEventsIngotRecordRow) IngotIn() int64 {
	return int64(this.row.IngotIn)
}

// 累计充值活动结束时间戳
func (this *PlayerEventsIngotRecordRow) IngotInEndTime() int64 {
	return int64(this.row.IngotInEndTime)
}

// 消耗元宝总数
func (this *PlayerEventsIngotRecordRow) IngotOut() int64 {
	return int64(this.row.IngotOut)
}

// 消耗元宝活动结束时间戳
func (this *PlayerEventsIngotRecordRow) IngotOutEndTime() int64 {
	return int64(this.row.IngotOutEndTime)
}

func (this *PlayerEventsIngotRecordRow) GoObject() *PlayerEventsIngotRecord {
	return &PlayerEventsIngotRecord{
		Pid: int64(this.row.Pid),
		IngotIn: int64(this.row.IngotIn),
		IngotInEndTime: int64(this.row.IngotInEndTime),
		IngotOut: int64(this.row.IngotOut),
		IngotOutEndTime: int64(this.row.IngotOutEndTime),
	}
}

/* ========== player_extend_level ========== */

// 玩家活动关卡状态
type PlayerExtendLevelRow struct {
	row *C.PlayerExtendLevel
	isBreak bool
}

func (this *PlayerExtendLevelRow) reset(row *C.PlayerExtendLevel) {	this.row = row
	this.isBreak = false
}

func (this *PlayerExtendLevelRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerExtendLevelRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 铜钱关卡通关时间
func (this *PlayerExtendLevelRow) CoinPassTime() int64 {
	return int64(this.row.CoinPassTime)
}

// 经验关卡通关时间
func (this *PlayerExtendLevelRow) ExpPassTime() int64 {
	return int64(this.row.ExpPassTime)
}

// 魂侍关卡通关时间
func (this *PlayerExtendLevelRow) GhostPassTime() int64 {
	return int64(this.row.GhostPassTime)
}

// 灵宠关卡通关时间
func (this *PlayerExtendLevelRow) PetPassTime() int64 {
	return int64(this.row.PetPassTime)
}

// 伙伴关卡通关时间
func (this *PlayerExtendLevelRow) BuddyPassTime() int64 {
	return int64(this.row.BuddyPassTime)
}

// 经验关卡今日进入次数
func (this *PlayerExtendLevelRow) CoinDailyNum() int8 {
	return int8(this.row.CoinDailyNum)
}

// 铜钱关卡今日进入次数
func (this *PlayerExtendLevelRow) ExpDailyNum() int8 {
	return int8(this.row.ExpDailyNum)
}

// 伙伴关卡今日进入次数
func (this *PlayerExtendLevelRow) BuddyDailyNum() int8 {
	return int8(this.row.BuddyDailyNum)
}

// 灵宠关卡今日进入次数
func (this *PlayerExtendLevelRow) PetDailyNum() int8 {
	return int8(this.row.PetDailyNum)
}

// 魂侍关卡今日进入次数
func (this *PlayerExtendLevelRow) GhostDailyNum() int8 {
	return int8(this.row.GhostDailyNum)
}

// 随机的伙伴角色ID
func (this *PlayerExtendLevelRow) RandBuddyRoleId() int8 {
	return int8(this.row.RandBuddyRoleId)
}

// 随机的伙伴角色位置
func (this *PlayerExtendLevelRow) BuddyPos() int8 {
	return int8(this.row.BuddyPos)
}

// 伙伴关卡队伍战术
func (this *PlayerExtendLevelRow) BuddyTactical() int8 {
	return int8(this.row.BuddyTactical)
}

// 主角站位
func (this *PlayerExtendLevelRow) RolePos() int8 {
	return int8(this.row.RolePos)
}

// 经验关卡通关了的最大等级
func (this *PlayerExtendLevelRow) ExpMaxlevel() int16 {
	return int16(this.row.ExpMaxlevel)
}

// 经验关卡通关了的最大等级
func (this *PlayerExtendLevelRow) CoinsMaxlevel() int16 {
	return int16(this.row.CoinsMaxlevel)
}

// 元宝购买铜钱关卡次数
func (this *PlayerExtendLevelRow) CoinsBuyNum() int16 {
	return int16(this.row.CoinsBuyNum)
}

// 元宝购买经验关卡次数
func (this *PlayerExtendLevelRow) ExpBuyNum() int16 {
	return int16(this.row.ExpBuyNum)
}

// 元宝购买铜钱关卡时间
func (this *PlayerExtendLevelRow) CoinsBuyTime() int64 {
	return int64(this.row.CoinsBuyTime)
}

// 元宝购买经验关卡时间
func (this *PlayerExtendLevelRow) ExpBuyTime() int64 {
	return int64(this.row.ExpBuyTime)
}

func (this *PlayerExtendLevelRow) GoObject() *PlayerExtendLevel {
	return &PlayerExtendLevel{
		Pid: int64(this.row.Pid),
		CoinPassTime: int64(this.row.CoinPassTime),
		ExpPassTime: int64(this.row.ExpPassTime),
		GhostPassTime: int64(this.row.GhostPassTime),
		PetPassTime: int64(this.row.PetPassTime),
		BuddyPassTime: int64(this.row.BuddyPassTime),
		CoinDailyNum: int8(this.row.CoinDailyNum),
		ExpDailyNum: int8(this.row.ExpDailyNum),
		BuddyDailyNum: int8(this.row.BuddyDailyNum),
		PetDailyNum: int8(this.row.PetDailyNum),
		GhostDailyNum: int8(this.row.GhostDailyNum),
		RandBuddyRoleId: int8(this.row.RandBuddyRoleId),
		BuddyPos: int8(this.row.BuddyPos),
		BuddyTactical: int8(this.row.BuddyTactical),
		RolePos: int8(this.row.RolePos),
		ExpMaxlevel: int16(this.row.ExpMaxlevel),
		CoinsMaxlevel: int16(this.row.CoinsMaxlevel),
		CoinsBuyNum: int16(this.row.CoinsBuyNum),
		ExpBuyNum: int16(this.row.ExpBuyNum),
		CoinsBuyTime: int64(this.row.CoinsBuyTime),
		ExpBuyTime: int64(this.row.ExpBuyTime),
	}
}

/* ========== player_extend_quest ========== */

// 伙伴任务
type PlayerExtendQuestRow struct {
	row *C.PlayerExtendQuest
	isBreak bool
}

func (this *PlayerExtendQuestRow) reset(row *C.PlayerExtendQuest) {	this.row = row
	this.isBreak = false
}

func (this *PlayerExtendQuestRow) Break() {
	this.isBreak = true
}

// ID
func (this *PlayerExtendQuestRow) Id() int64 {
	return int64(this.row.Id)
}

// 用户ID
func (this *PlayerExtendQuestRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 任务ID
func (this *PlayerExtendQuestRow) QuestId() int32 {
	return int32(this.row.QuestId)
}

// 任务进度
func (this *PlayerExtendQuestRow) Progress() int16 {
	return int16(this.row.Progress)
}

// 0--未完成 1--已完成 2--已奖励
func (this *PlayerExtendQuestRow) State() int8 {
	return int8(this.row.State)
}

func (this *PlayerExtendQuestRow) GoObject() *PlayerExtendQuest {
	return &PlayerExtendQuest{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		QuestId: int32(this.row.QuestId),
		Progress: int16(this.row.Progress),
		State: int8(this.row.State),
	}
}

/* ========== player_fame ========== */

// 
type PlayerFameRow struct {
	row *C.PlayerFame
	isBreak bool
}

func (this *PlayerFameRow) reset(row *C.PlayerFame) {	this.row = row
	this.isBreak = false
}

func (this *PlayerFameRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerFameRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 总声望
func (this *PlayerFameRow) Fame() int32 {
	return int32(this.row.Fame)
}

// 声望等级
func (this *PlayerFameRow) Level() int16 {
	return int16(this.row.Level)
}

// 比武场关卡声望
func (this *PlayerFameRow) ArenaFame() int32 {
	return int32(this.row.ArenaFame)
}

// 多人关卡声望
func (this *PlayerFameRow) MultLevelFame() int32 {
	return int32(this.row.MultLevelFame)
}

func (this *PlayerFameRow) GoObject() *PlayerFame {
	return &PlayerFame{
		Pid: int64(this.row.Pid),
		Fame: int32(this.row.Fame),
		Level: int16(this.row.Level),
		ArenaFame: int32(this.row.ArenaFame),
		MultLevelFame: int32(this.row.MultLevelFame),
	}
}

/* ========== player_fashion ========== */

// 玩家时装
type PlayerFashionRow struct {
	row *C.PlayerFashion
	isBreak bool
}

func (this *PlayerFashionRow) reset(row *C.PlayerFashion) {	this.row = row
	this.isBreak = false
}

func (this *PlayerFashionRow) Break() {
	this.isBreak = true
}

// ID
func (this *PlayerFashionRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerFashionRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 时装模版ID
func (this *PlayerFashionRow) FashionId() int16 {
	return int16(this.row.FashionId)
}

// 过期时间 0--永远有效
func (this *PlayerFashionRow) ExpireTime() int64 {
	return int64(this.row.ExpireTime)
}

func (this *PlayerFashionRow) GoObject() *PlayerFashion {
	return &PlayerFashion{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		FashionId: int16(this.row.FashionId),
		ExpireTime: int64(this.row.ExpireTime),
	}
}

/* ========== player_fashion_state ========== */

// 玩家时装状态
type PlayerFashionStateRow struct {
	row *C.PlayerFashionState
	isBreak bool
}

func (this *PlayerFashionStateRow) reset(row *C.PlayerFashionState) {	this.row = row
	this.isBreak = false
}

func (this *PlayerFashionStateRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerFashionStateRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 状态更新时间
func (this *PlayerFashionStateRow) UpdateTime() int64 {
	return int64(this.row.UpdateTime)
}

// 当前装备的时装
func (this *PlayerFashionStateRow) DressedFashionId() int16 {
	return int16(this.row.DressedFashionId)
}

func (this *PlayerFashionStateRow) GoObject() *PlayerFashionState {
	return &PlayerFashionState{
		Pid: int64(this.row.Pid),
		UpdateTime: int64(this.row.UpdateTime),
		DressedFashionId: int16(this.row.DressedFashionId),
	}
}

/* ========== player_fate_box_state ========== */

// 玩家命锁宝箱状态
type PlayerFateBoxStateRow struct {
	row *C.PlayerFateBoxState
	isBreak bool
}

func (this *PlayerFateBoxStateRow) reset(row *C.PlayerFateBoxState) {	this.row = row
	this.isBreak = false
}

func (this *PlayerFateBoxStateRow) Break() {
	this.isBreak = true
}

// pid
func (this *PlayerFateBoxStateRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 玩家命锁权值
func (this *PlayerFateBoxStateRow) Lock() int32 {
	return int32(this.row.Lock)
}

// 星辉命匣免费抽取时间戳
func (this *PlayerFateBoxStateRow) StarBoxFreeDrawTimestamp() int64 {
	return int64(this.row.StarBoxFreeDrawTimestamp)
}

// 星辉宝箱抽奖次数
func (this *PlayerFateBoxStateRow) StarBoxDrawCount() int32 {
	return int32(this.row.StarBoxDrawCount)
}

// 月影命匣免费抽取时间戳
func (this *PlayerFateBoxStateRow) MoonBoxFreeDrawTimestamp() int64 {
	return int64(this.row.MoonBoxFreeDrawTimestamp)
}

// 月影宝箱抽奖次数
func (this *PlayerFateBoxStateRow) MoonBoxDrawCount() int32 {
	return int32(this.row.MoonBoxDrawCount)
}

// 日耀命匣免费抽取时间戳
func (this *PlayerFateBoxStateRow) SunBoxFreeDrawTimestamp() int64 {
	return int64(this.row.SunBoxFreeDrawTimestamp)
}

// 日耀宝箱抽奖次数
func (this *PlayerFateBoxStateRow) SunBoxDrawCount() int32 {
	return int32(this.row.SunBoxDrawCount)
}

// 混元命匣免费抽取时间戳
func (this *PlayerFateBoxStateRow) HunyuanBoxFreeDrawTimestamp() int64 {
	return int64(this.row.HunyuanBoxFreeDrawTimestamp)
}

// 混元宝箱抽奖次数
func (this *PlayerFateBoxStateRow) HunyuanBoxDrawCount() int32 {
	return int32(this.row.HunyuanBoxDrawCount)
}

func (this *PlayerFateBoxStateRow) GoObject() *PlayerFateBoxState {
	return &PlayerFateBoxState{
		Pid: int64(this.row.Pid),
		Lock: int32(this.row.Lock),
		StarBoxFreeDrawTimestamp: int64(this.row.StarBoxFreeDrawTimestamp),
		StarBoxDrawCount: int32(this.row.StarBoxDrawCount),
		MoonBoxFreeDrawTimestamp: int64(this.row.MoonBoxFreeDrawTimestamp),
		MoonBoxDrawCount: int32(this.row.MoonBoxDrawCount),
		SunBoxFreeDrawTimestamp: int64(this.row.SunBoxFreeDrawTimestamp),
		SunBoxDrawCount: int32(this.row.SunBoxDrawCount),
		HunyuanBoxFreeDrawTimestamp: int64(this.row.HunyuanBoxFreeDrawTimestamp),
		HunyuanBoxDrawCount: int32(this.row.HunyuanBoxDrawCount),
	}
}

/* ========== player_fight_num ========== */

// 玩家战斗力
type PlayerFightNumRow struct {
	row *C.PlayerFightNum
	isBreak bool
}

func (this *PlayerFightNumRow) reset(row *C.PlayerFightNum) {	this.row = row
	this.isBreak = false
}

func (this *PlayerFightNumRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerFightNumRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 战力力
func (this *PlayerFightNumRow) FightNum() int32 {
	return int32(this.row.FightNum)
}

func (this *PlayerFightNumRow) GoObject() *PlayerFightNum {
	return &PlayerFightNum{
		Pid: int64(this.row.Pid),
		FightNum: int32(this.row.FightNum),
	}
}

/* ========== player_first_charge ========== */

// 玩家首冲表
type PlayerFirstChargeRow struct {
	row *C.PlayerFirstCharge
	isBreak bool
	cached_ProductId bool
	cacheof_ProductId string
}

func (this *PlayerFirstChargeRow) reset(row *C.PlayerFirstCharge) {	this.row = row
	this.isBreak = false
	this.cached_ProductId = false
}

func (this *PlayerFirstChargeRow) Break() {
	this.isBreak = true
}

// 
func (this *PlayerFirstChargeRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerFirstChargeRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 
func (this *PlayerFirstChargeRow) ProductId() string {
	if !this.cached_ProductId {
		this.cached_ProductId = true
		this.cacheof_ProductId = C.GoStringN(this.row.ProductId, this.row.ProductId_len)
	}
	return this.cacheof_ProductId
}

func (this *PlayerFirstChargeRow) GoObject() *PlayerFirstCharge {
	return &PlayerFirstCharge{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		ProductId: this.ProductId(),
	}
}

/* ========== player_formation ========== */

// 玩家阵型站位
type PlayerFormationRow struct {
	row *C.PlayerFormation
	isBreak bool
}

func (this *PlayerFormationRow) reset(row *C.PlayerFormation) {	this.row = row
	this.isBreak = false
}

func (this *PlayerFormationRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerFormationRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 站位0
func (this *PlayerFormationRow) Pos0() int8 {
	return int8(this.row.Pos0)
}

// 站位1
func (this *PlayerFormationRow) Pos1() int8 {
	return int8(this.row.Pos1)
}

// 站位2
func (this *PlayerFormationRow) Pos2() int8 {
	return int8(this.row.Pos2)
}

// 站位3
func (this *PlayerFormationRow) Pos3() int8 {
	return int8(this.row.Pos3)
}

// 站位4
func (this *PlayerFormationRow) Pos4() int8 {
	return int8(this.row.Pos4)
}

// 站位5
func (this *PlayerFormationRow) Pos5() int8 {
	return int8(this.row.Pos5)
}

// 站位6
func (this *PlayerFormationRow) Pos6() int8 {
	return int8(this.row.Pos6)
}

// 站位7
func (this *PlayerFormationRow) Pos7() int8 {
	return int8(this.row.Pos7)
}

// 站位8
func (this *PlayerFormationRow) Pos8() int8 {
	return int8(this.row.Pos8)
}

func (this *PlayerFormationRow) GoObject() *PlayerFormation {
	return &PlayerFormation{
		Pid: int64(this.row.Pid),
		Pos0: int8(this.row.Pos0),
		Pos1: int8(this.row.Pos1),
		Pos2: int8(this.row.Pos2),
		Pos3: int8(this.row.Pos3),
		Pos4: int8(this.row.Pos4),
		Pos5: int8(this.row.Pos5),
		Pos6: int8(this.row.Pos6),
		Pos7: int8(this.row.Pos7),
		Pos8: int8(this.row.Pos8),
	}
}

/* ========== player_func_key ========== */

// 玩家功能开放表
type PlayerFuncKeyRow struct {
	row *C.PlayerFuncKey
	isBreak bool
}

func (this *PlayerFuncKeyRow) reset(row *C.PlayerFuncKey) {	this.row = row
	this.isBreak = false
}

func (this *PlayerFuncKeyRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerFuncKeyRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 功能权值
func (this *PlayerFuncKeyRow) Key() int16 {
	return int16(this.row.Key)
}

// 已播放提示的功能
func (this *PlayerFuncKeyRow) PlayedKey() int16 {
	return int16(this.row.PlayedKey)
}

// 已开启功能的唯一权值
func (this *PlayerFuncKeyRow) UniqueKey() int64 {
	return int64(this.row.UniqueKey)
}

func (this *PlayerFuncKeyRow) GoObject() *PlayerFuncKey {
	return &PlayerFuncKey{
		Pid: int64(this.row.Pid),
		Key: int16(this.row.Key),
		PlayedKey: int16(this.row.PlayedKey),
		UniqueKey: int64(this.row.UniqueKey),
	}
}

/* ========== player_ghost ========== */

// 玩家魂侍表
type PlayerGhostRow struct {
	row *C.PlayerGhost
	isBreak bool
}

func (this *PlayerGhostRow) reset(row *C.PlayerGhost) {	this.row = row
	this.isBreak = false
}

func (this *PlayerGhostRow) Break() {
	this.isBreak = true
}

// 主键
func (this *PlayerGhostRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerGhostRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 魂侍ID
func (this *PlayerGhostRow) GhostId() int16 {
	return int16(this.row.GhostId)
}

// 星级
func (this *PlayerGhostRow) Star() int8 {
	return int8(this.row.Star)
}

// 魂侍等级
func (this *PlayerGhostRow) Level() int16 {
	return int16(this.row.Level)
}

// 魂侍经验
func (this *PlayerGhostRow) Exp() int64 {
	return int64(this.row.Exp)
}

// 位置
func (this *PlayerGhostRow) Pos() int16 {
	return int16(this.row.Pos)
}

// 是否新魂侍
func (this *PlayerGhostRow) IsNew() int8 {
	return int8(this.row.IsNew)
}

// 魂侍装备在那个角色身上 0 未装备
func (this *PlayerGhostRow) RoleId() int8 {
	return int8(this.row.RoleId)
}

// 技能等级
func (this *PlayerGhostRow) SkillLevel() int16 {
	return int16(this.row.SkillLevel)
}

// 连锁玩家魂侍ID
func (this *PlayerGhostRow) RelationId() int64 {
	return int64(this.row.RelationId)
}

// 魂侍洗练添加的成长值
func (this *PlayerGhostRow) AddGrowth() int16 {
	return int16(this.row.AddGrowth)
}

func (this *PlayerGhostRow) GoObject() *PlayerGhost {
	return &PlayerGhost{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		GhostId: int16(this.row.GhostId),
		Star: int8(this.row.Star),
		Level: int16(this.row.Level),
		Exp: int64(this.row.Exp),
		Pos: int16(this.row.Pos),
		IsNew: int8(this.row.IsNew),
		RoleId: int8(this.row.RoleId),
		SkillLevel: int16(this.row.SkillLevel),
		RelationId: int64(this.row.RelationId),
		AddGrowth: int16(this.row.AddGrowth),
	}
}

/* ========== player_ghost_equipment ========== */

// 玩家魂侍装备表
type PlayerGhostEquipmentRow struct {
	row *C.PlayerGhostEquipment
	isBreak bool
}

func (this *PlayerGhostEquipmentRow) reset(row *C.PlayerGhostEquipment) {	this.row = row
	this.isBreak = false
}

func (this *PlayerGhostEquipmentRow) Break() {
	this.isBreak = true
}

// 主键
func (this *PlayerGhostEquipmentRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerGhostEquipmentRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 角色ID
func (this *PlayerGhostEquipmentRow) RoleId() int8 {
	return int8(this.row.RoleId)
}

// 魂力
func (this *PlayerGhostEquipmentRow) GhostPower() int32 {
	return int32(this.row.GhostPower)
}

// 装备位置1的魂侍id
func (this *PlayerGhostEquipmentRow) Pos1() int64 {
	return int64(this.row.Pos1)
}

// 装备位置2的魂侍id
func (this *PlayerGhostEquipmentRow) Pos2() int64 {
	return int64(this.row.Pos2)
}

// 装备位置3的魂侍id
func (this *PlayerGhostEquipmentRow) Pos3() int64 {
	return int64(this.row.Pos3)
}

// 装备位置4的魂侍id
func (this *PlayerGhostEquipmentRow) Pos4() int64 {
	return int64(this.row.Pos4)
}

func (this *PlayerGhostEquipmentRow) GoObject() *PlayerGhostEquipment {
	return &PlayerGhostEquipment{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		RoleId: int8(this.row.RoleId),
		GhostPower: int32(this.row.GhostPower),
		Pos1: int64(this.row.Pos1),
		Pos2: int64(this.row.Pos2),
		Pos3: int64(this.row.Pos3),
		Pos4: int64(this.row.Pos4),
	}
}

/* ========== player_ghost_state ========== */

// 玩家魂侍相关状态
type PlayerGhostStateRow struct {
	row *C.PlayerGhostState
	isBreak bool
}

func (this *PlayerGhostStateRow) reset(row *C.PlayerGhostState) {	this.row = row
	this.isBreak = false
}

func (this *PlayerGhostStateRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerGhostStateRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 今日使用元宝培养次数
func (this *PlayerGhostStateRow) TrainByIngotNum() int32 {
	return int32(this.row.TrainByIngotNum)
}

// 最近一次使用元宝培养时间
func (this *PlayerGhostStateRow) TrainByIngotTime() int64 {
	return int64(this.row.TrainByIngotTime)
}

// 最近一次洗点时间
func (this *PlayerGhostStateRow) LastFlushTime() int64 {
	return int64(this.row.LastFlushTime)
}

// 玩家魂侍总战力
func (this *PlayerGhostStateRow) GhostFightNum() int64 {
	return int64(this.row.GhostFightNum)
}

func (this *PlayerGhostStateRow) GoObject() *PlayerGhostState {
	return &PlayerGhostState{
		Pid: int64(this.row.Pid),
		TrainByIngotNum: int32(this.row.TrainByIngotNum),
		TrainByIngotTime: int64(this.row.TrainByIngotTime),
		LastFlushTime: int64(this.row.LastFlushTime),
		GhostFightNum: int64(this.row.GhostFightNum),
	}
}

/* ========== player_global_chat_state ========== */

// 禁言状态
type PlayerGlobalChatStateRow struct {
	row *C.PlayerGlobalChatState
	isBreak bool
}

func (this *PlayerGlobalChatStateRow) reset(row *C.PlayerGlobalChatState) {	this.row = row
	this.isBreak = false
}

func (this *PlayerGlobalChatStateRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerGlobalChatStateRow) Pid() int64 {
	return int64(this.row.Pid)
}

// <=0 没有禁言 >0禁言
func (this *PlayerGlobalChatStateRow) BanUntil() int64 {
	return int64(this.row.BanUntil)
}

func (this *PlayerGlobalChatStateRow) GoObject() *PlayerGlobalChatState {
	return &PlayerGlobalChatState{
		Pid: int64(this.row.Pid),
		BanUntil: int64(this.row.BanUntil),
	}
}

/* ========== player_global_clique_building ========== */

// 玩家帮派建筑信息
type PlayerGlobalCliqueBuildingRow struct {
	row *C.PlayerGlobalCliqueBuilding
	isBreak bool
}

func (this *PlayerGlobalCliqueBuildingRow) reset(row *C.PlayerGlobalCliqueBuilding) {	this.row = row
	this.isBreak = false
}

func (this *PlayerGlobalCliqueBuildingRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerGlobalCliqueBuildingRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 
func (this *PlayerGlobalCliqueBuildingRow) SilverExchangeTime() int64 {
	return int64(this.row.SilverExchangeTime)
}

// 每日金劵兑换数
func (this *PlayerGlobalCliqueBuildingRow) GoldExchangeNum() int16 {
	return int16(this.row.GoldExchangeNum)
}

// 每日银劵兑换数
func (this *PlayerGlobalCliqueBuildingRow) SilverExchangeNum() int16 {
	return int16(this.row.SilverExchangeNum)
}

// 总部累积贡献铜币
func (this *PlayerGlobalCliqueBuildingRow) DonateCoinsCenterBuilding() int64 {
	return int64(this.row.DonateCoinsCenterBuilding)
}

// 宗祠累积贡献铜币
func (this *PlayerGlobalCliqueBuildingRow) DonateCoinsTempleBuilding() int64 {
	return int64(this.row.DonateCoinsTempleBuilding)
}

// 钱庄累积贡献铜币
func (this *PlayerGlobalCliqueBuildingRow) DonateCoinsBankBuilding() int64 {
	return int64(this.row.DonateCoinsBankBuilding)
}

// 回春堂累积贡献铜币
func (this *PlayerGlobalCliqueBuildingRow) DonateCoinsHealthBuilding() int64 {
	return int64(this.row.DonateCoinsHealthBuilding)
}

// 神兵堂累积贡献铜币
func (this *PlayerGlobalCliqueBuildingRow) DonateCoinsAttackBuilding() int64 {
	return int64(this.row.DonateCoinsAttackBuilding)
}

// 金刚堂累积贡献铜币
func (this *PlayerGlobalCliqueBuildingRow) DonateCoinsDefenseBuilding() int64 {
	return int64(this.row.DonateCoinsDefenseBuilding)
}

// 战备仓库累积贡献铜币
func (this *PlayerGlobalCliqueBuildingRow) DonateCoinsStoreBuilding() int64 {
	return int64(this.row.DonateCoinsStoreBuilding)
}

// 回春等级
func (this *PlayerGlobalCliqueBuildingRow) HealthLevel() int16 {
	return int16(this.row.HealthLevel)
}

// 神兵等级
func (this *PlayerGlobalCliqueBuildingRow) AttackLevel() int16 {
	return int16(this.row.AttackLevel)
}

// 金刚等级
func (this *PlayerGlobalCliqueBuildingRow) DefenseLevel() int16 {
	return int16(this.row.DefenseLevel)
}

// 上香时间
func (this *PlayerGlobalCliqueBuildingRow) WorshipTime() int64 {
	return int64(this.row.WorshipTime)
}

// 上次总舵捐钱时间戳
func (this *PlayerGlobalCliqueBuildingRow) DonateCoinsCenterBuildingTime() int64 {
	return int64(this.row.DonateCoinsCenterBuildingTime)
}

// 金卷购买时间戳
func (this *PlayerGlobalCliqueBuildingRow) GlodExchangeTime() int64 {
	return int64(this.row.GlodExchangeTime)
}

// 上香类型
func (this *PlayerGlobalCliqueBuildingRow) WorshipType() int8 {
	return int8(this.row.WorshipType)
}

func (this *PlayerGlobalCliqueBuildingRow) GoObject() *PlayerGlobalCliqueBuilding {
	return &PlayerGlobalCliqueBuilding{
		Pid: int64(this.row.Pid),
		SilverExchangeTime: int64(this.row.SilverExchangeTime),
		GoldExchangeNum: int16(this.row.GoldExchangeNum),
		SilverExchangeNum: int16(this.row.SilverExchangeNum),
		DonateCoinsCenterBuilding: int64(this.row.DonateCoinsCenterBuilding),
		DonateCoinsTempleBuilding: int64(this.row.DonateCoinsTempleBuilding),
		DonateCoinsBankBuilding: int64(this.row.DonateCoinsBankBuilding),
		DonateCoinsHealthBuilding: int64(this.row.DonateCoinsHealthBuilding),
		DonateCoinsAttackBuilding: int64(this.row.DonateCoinsAttackBuilding),
		DonateCoinsDefenseBuilding: int64(this.row.DonateCoinsDefenseBuilding),
		DonateCoinsStoreBuilding: int64(this.row.DonateCoinsStoreBuilding),
		HealthLevel: int16(this.row.HealthLevel),
		AttackLevel: int16(this.row.AttackLevel),
		DefenseLevel: int16(this.row.DefenseLevel),
		WorshipTime: int64(this.row.WorshipTime),
		DonateCoinsCenterBuildingTime: int64(this.row.DonateCoinsCenterBuildingTime),
		GlodExchangeTime: int64(this.row.GlodExchangeTime),
		WorshipType: int8(this.row.WorshipType),
	}
}

/* ========== player_global_clique_building_quest ========== */

// 玩家帮派建筑任务
type PlayerGlobalCliqueBuildingQuestRow struct {
	row *C.PlayerGlobalCliqueBuildingQuest
	isBreak bool
}

func (this *PlayerGlobalCliqueBuildingQuestRow) reset(row *C.PlayerGlobalCliqueBuildingQuest) {	this.row = row
	this.isBreak = false
}

func (this *PlayerGlobalCliqueBuildingQuestRow) Break() {
	this.isBreak = true
}

// 记录ID
func (this *PlayerGlobalCliqueBuildingQuestRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerGlobalCliqueBuildingQuestRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 任务ID
func (this *PlayerGlobalCliqueBuildingQuestRow) QuestId() int16 {
	return int16(this.row.QuestId)
}

// 奖励状态; 0 未奖励； 1已奖励
func (this *PlayerGlobalCliqueBuildingQuestRow) AwardStatus() int8 {
	return int8(this.row.AwardStatus)
}

// 建筑类别
func (this *PlayerGlobalCliqueBuildingQuestRow) BuildingType() int16 {
	return int16(this.row.BuildingType)
}

func (this *PlayerGlobalCliqueBuildingQuestRow) GoObject() *PlayerGlobalCliqueBuildingQuest {
	return &PlayerGlobalCliqueBuildingQuest{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		QuestId: int16(this.row.QuestId),
		AwardStatus: int8(this.row.AwardStatus),
		BuildingType: int16(this.row.BuildingType),
	}
}

/* ========== player_global_clique_daily_quest ========== */

// 玩家每日帮派任务
type PlayerGlobalCliqueDailyQuestRow struct {
	row *C.PlayerGlobalCliqueDailyQuest
	isBreak bool
}

func (this *PlayerGlobalCliqueDailyQuestRow) reset(row *C.PlayerGlobalCliqueDailyQuest) {	this.row = row
	this.isBreak = false
}

func (this *PlayerGlobalCliqueDailyQuestRow) Break() {
	this.isBreak = true
}

// 记录ID
func (this *PlayerGlobalCliqueDailyQuestRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerGlobalCliqueDailyQuestRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 任务ID
func (this *PlayerGlobalCliqueDailyQuestRow) QuestId() int16 {
	return int16(this.row.QuestId)
}

// 完成数量
func (this *PlayerGlobalCliqueDailyQuestRow) FinishCount() int16 {
	return int16(this.row.FinishCount)
}

// 最近一次更新时间
func (this *PlayerGlobalCliqueDailyQuestRow) LastUpdateTime() int64 {
	return int64(this.row.LastUpdateTime)
}

// 奖励状态; 0 未奖励； 1已奖励
func (this *PlayerGlobalCliqueDailyQuestRow) AwardStatus() int8 {
	return int8(this.row.AwardStatus)
}

// 每日任务类别
func (this *PlayerGlobalCliqueDailyQuestRow) Class() int16 {
	return int16(this.row.Class)
}

func (this *PlayerGlobalCliqueDailyQuestRow) GoObject() *PlayerGlobalCliqueDailyQuest {
	return &PlayerGlobalCliqueDailyQuest{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		QuestId: int16(this.row.QuestId),
		FinishCount: int16(this.row.FinishCount),
		LastUpdateTime: int64(this.row.LastUpdateTime),
		AwardStatus: int8(this.row.AwardStatus),
		Class: int16(this.row.Class),
	}
}

/* ========== player_global_clique_escort ========== */

// 玩家帮派运镖信息
type PlayerGlobalCliqueEscortRow struct {
	row *C.PlayerGlobalCliqueEscort
	isBreak bool
}

func (this *PlayerGlobalCliqueEscortRow) reset(row *C.PlayerGlobalCliqueEscort) {	this.row = row
	this.isBreak = false
}

func (this *PlayerGlobalCliqueEscortRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerGlobalCliqueEscortRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 今日首次送镖时间戳
func (this *PlayerGlobalCliqueEscortRow) DailyEscortTimestamp() int64 {
	return int64(this.row.DailyEscortTimestamp)
}

// 送镖次数
func (this *PlayerGlobalCliqueEscortRow) DailyEscortNum() int16 {
	return int16(this.row.DailyEscortNum)
}

// 当前标船类型
func (this *PlayerGlobalCliqueEscortRow) EscortBoatType() int16 {
	return int16(this.row.EscortBoatType)
}

// 0--无运镖 1--运镖中 2--等待领取奖励
func (this *PlayerGlobalCliqueEscortRow) Status() int8 {
	return int8(this.row.Status)
}

// 0--没有被完成过劫持 1--被完成过劫持 
func (this *PlayerGlobalCliqueEscortRow) Hijacked() int8 {
	return int8(this.row.Hijacked)
}

func (this *PlayerGlobalCliqueEscortRow) GoObject() *PlayerGlobalCliqueEscort {
	return &PlayerGlobalCliqueEscort{
		Pid: int64(this.row.Pid),
		DailyEscortTimestamp: int64(this.row.DailyEscortTimestamp),
		DailyEscortNum: int16(this.row.DailyEscortNum),
		EscortBoatType: int16(this.row.EscortBoatType),
		Status: int8(this.row.Status),
		Hijacked: int8(this.row.Hijacked),
	}
}

/* ========== player_global_clique_escort_message ========== */

// 玩家帮派运镖消息
type PlayerGlobalCliqueEscortMessageRow struct {
	row *C.PlayerGlobalCliqueEscortMessage
	isBreak bool
	cached_Parameters bool
	cacheof_Parameters string
}

func (this *PlayerGlobalCliqueEscortMessageRow) reset(row *C.PlayerGlobalCliqueEscortMessage) {	this.row = row
	this.isBreak = false
	this.cached_Parameters = false
}

func (this *PlayerGlobalCliqueEscortMessageRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerGlobalCliqueEscortMessageRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerGlobalCliqueEscortMessageRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 模版ID
func (this *PlayerGlobalCliqueEscortMessageRow) TplId() int16 {
	return int16(this.row.TplId)
}

// 模版参数
func (this *PlayerGlobalCliqueEscortMessageRow) Parameters() string {
	if !this.cached_Parameters {
		this.cached_Parameters = true
		this.cacheof_Parameters = C.GoStringN(this.row.Parameters, this.row.Parameters_len)
	}
	return this.cacheof_Parameters
}

// 发送时间戳
func (this *PlayerGlobalCliqueEscortMessageRow) Timestamp() int64 {
	return int64(this.row.Timestamp)
}

func (this *PlayerGlobalCliqueEscortMessageRow) GoObject() *PlayerGlobalCliqueEscortMessage {
	return &PlayerGlobalCliqueEscortMessage{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		TplId: int16(this.row.TplId),
		Parameters: this.Parameters(),
		Timestamp: int64(this.row.Timestamp),
	}
}

/* ========== player_global_clique_hijack ========== */

// 玩家劫镖信息
type PlayerGlobalCliqueHijackRow struct {
	row *C.PlayerGlobalCliqueHijack
	isBreak bool
}

func (this *PlayerGlobalCliqueHijackRow) reset(row *C.PlayerGlobalCliqueHijack) {	this.row = row
	this.isBreak = false
}

func (this *PlayerGlobalCliqueHijackRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerGlobalCliqueHijackRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 劫持时间戳
func (this *PlayerGlobalCliqueHijackRow) DailyHijackTimestamp() int64 {
	return int64(this.row.DailyHijackTimestamp)
}

// 劫持次数
func (this *PlayerGlobalCliqueHijackRow) DailyHijackNum() int16 {
	return int16(this.row.DailyHijackNum)
}

// 待奖励的标船类型
func (this *PlayerGlobalCliqueHijackRow) HijackedBoatType() int16 {
	return int16(this.row.HijackedBoatType)
}

// 0--无劫镖 1--劫镖中 2--等待领取奖励
func (this *PlayerGlobalCliqueHijackRow) Status() int8 {
	return int8(this.row.Status)
}

// 购买时间戳
func (this *PlayerGlobalCliqueHijackRow) BuyTimestamp() int64 {
	return int64(this.row.BuyTimestamp)
}

// 今日购买次数
func (this *PlayerGlobalCliqueHijackRow) BuyNum() int16 {
	return int16(this.row.BuyNum)
}

func (this *PlayerGlobalCliqueHijackRow) GoObject() *PlayerGlobalCliqueHijack {
	return &PlayerGlobalCliqueHijack{
		Pid: int64(this.row.Pid),
		DailyHijackTimestamp: int64(this.row.DailyHijackTimestamp),
		DailyHijackNum: int16(this.row.DailyHijackNum),
		HijackedBoatType: int16(this.row.HijackedBoatType),
		Status: int8(this.row.Status),
		BuyTimestamp: int64(this.row.BuyTimestamp),
		BuyNum: int16(this.row.BuyNum),
	}
}

/* ========== player_global_clique_info ========== */

// 玩家帮派信息
type PlayerGlobalCliqueInfoRow struct {
	row *C.PlayerGlobalCliqueInfo
	isBreak bool
}

func (this *PlayerGlobalCliqueInfoRow) reset(row *C.PlayerGlobalCliqueInfo) {	this.row = row
	this.isBreak = false
}

func (this *PlayerGlobalCliqueInfoRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerGlobalCliqueInfoRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 帮派名称
func (this *PlayerGlobalCliqueInfoRow) CliqueId() int64 {
	return int64(this.row.CliqueId)
}

// 加入帮派时间
func (this *PlayerGlobalCliqueInfoRow) JoinTime() int64 {
	return int64(this.row.JoinTime)
}

// 帮派贡献
func (this *PlayerGlobalCliqueInfoRow) Contrib() int64 {
	return int64(this.row.Contrib)
}

// 帮派贡献结算时间(每周一清理贡献)
func (this *PlayerGlobalCliqueInfoRow) ContribClearTime() int64 {
	return int64(this.row.ContribClearTime)
}

// 贡献铜币时间戳(限制捐献大小)
func (this *PlayerGlobalCliqueInfoRow) DonateCoinsTime() int64 {
	return int64(this.row.DonateCoinsTime)
}

// 当日贡献铜币
func (this *PlayerGlobalCliqueInfoRow) DailyDonateCoins() int64 {
	return int64(this.row.DailyDonateCoins)
}

// 玩家帮派总贡献
func (this *PlayerGlobalCliqueInfoRow) TotalContrib() int64 {
	return int64(this.row.TotalContrib)
}

func (this *PlayerGlobalCliqueInfoRow) GoObject() *PlayerGlobalCliqueInfo {
	return &PlayerGlobalCliqueInfo{
		Pid: int64(this.row.Pid),
		CliqueId: int64(this.row.CliqueId),
		JoinTime: int64(this.row.JoinTime),
		Contrib: int64(this.row.Contrib),
		ContribClearTime: int64(this.row.ContribClearTime),
		DonateCoinsTime: int64(this.row.DonateCoinsTime),
		DailyDonateCoins: int64(this.row.DailyDonateCoins),
		TotalContrib: int64(this.row.TotalContrib),
	}
}

/* ========== player_global_clique_kongfu ========== */

// 玩家帮派武学
type PlayerGlobalCliqueKongfuRow struct {
	row *C.PlayerGlobalCliqueKongfu
	isBreak bool
}

func (this *PlayerGlobalCliqueKongfuRow) reset(row *C.PlayerGlobalCliqueKongfu) {	this.row = row
	this.isBreak = false
}

func (this *PlayerGlobalCliqueKongfuRow) Break() {
	this.isBreak = true
}

// ID
func (this *PlayerGlobalCliqueKongfuRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerGlobalCliqueKongfuRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 武功ID
func (this *PlayerGlobalCliqueKongfuRow) KongfuId() int32 {
	return int32(this.row.KongfuId)
}

// 武功等级
func (this *PlayerGlobalCliqueKongfuRow) Level() int16 {
	return int16(this.row.Level)
}

func (this *PlayerGlobalCliqueKongfuRow) GoObject() *PlayerGlobalCliqueKongfu {
	return &PlayerGlobalCliqueKongfu{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		KongfuId: int32(this.row.KongfuId),
		Level: int16(this.row.Level),
	}
}

/* ========== player_global_friend ========== */

// 玩家好友列表
type PlayerGlobalFriendRow struct {
	row *C.PlayerGlobalFriend
	isBreak bool
	cached_FriendNick bool
	cacheof_FriendNick string
}

func (this *PlayerGlobalFriendRow) reset(row *C.PlayerGlobalFriend) {	this.row = row
	this.isBreak = false
	this.cached_FriendNick = false
}

func (this *PlayerGlobalFriendRow) Break() {
	this.isBreak = true
}

// 好友关系ID
func (this *PlayerGlobalFriendRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerGlobalFriendRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 好友ID
func (this *PlayerGlobalFriendRow) FriendPid() int64 {
	return int64(this.row.FriendPid)
}

// 玩家昵称
func (this *PlayerGlobalFriendRow) FriendNick() string {
	if !this.cached_FriendNick {
		this.cached_FriendNick = true
		this.cacheof_FriendNick = C.GoStringN(this.row.FriendNick, this.row.FriendNick_len)
	}
	return this.cacheof_FriendNick
}

// 好友角色ID
func (this *PlayerGlobalFriendRow) FriendRoleId() int8 {
	return int8(this.row.FriendRoleId)
}

// 好友关系:0陌生人,1仅关注,2仅被关注,3互相关注(好友)
func (this *PlayerGlobalFriendRow) FriendMode() int8 {
	return int8(this.row.FriendMode)
}

// 最后聊天时间
func (this *PlayerGlobalFriendRow) LastChatTime() int64 {
	return int64(this.row.LastChatTime)
}

// 成为好友时间
func (this *PlayerGlobalFriendRow) FriendTime() int64 {
	return int64(this.row.FriendTime)
}

// 上次送爱心时间
func (this *PlayerGlobalFriendRow) SendHeartTime() int64 {
	return int64(this.row.SendHeartTime)
}

// 黑名单状态:0-否,1-是
func (this *PlayerGlobalFriendRow) BlockMode() int8 {
	return int8(this.row.BlockMode)
}

func (this *PlayerGlobalFriendRow) GoObject() *PlayerGlobalFriend {
	return &PlayerGlobalFriend{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		FriendPid: int64(this.row.FriendPid),
		FriendNick: this.FriendNick(),
		FriendRoleId: int8(this.row.FriendRoleId),
		FriendMode: int8(this.row.FriendMode),
		LastChatTime: int64(this.row.LastChatTime),
		FriendTime: int64(this.row.FriendTime),
		SendHeartTime: int64(this.row.SendHeartTime),
		BlockMode: int8(this.row.BlockMode),
	}
}

/* ========== player_global_friend_chat ========== */

// 玩家聊天记录
type PlayerGlobalFriendChatRow struct {
	row *C.PlayerGlobalFriendChat
	isBreak bool
	cached_Message bool
	cacheof_Message string
}

func (this *PlayerGlobalFriendChatRow) reset(row *C.PlayerGlobalFriendChat) {	this.row = row
	this.isBreak = false
	this.cached_Message = false
}

func (this *PlayerGlobalFriendChatRow) Break() {
	this.isBreak = true
}

// 消息ID
func (this *PlayerGlobalFriendChatRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerGlobalFriendChatRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 对方玩家ID
func (this *PlayerGlobalFriendChatRow) FriendPid() int64 {
	return int64(this.row.FriendPid)
}

// 1发送，2接收
func (this *PlayerGlobalFriendChatRow) Mode() int8 {
	return int8(this.row.Mode)
}

// 发送时间戳
func (this *PlayerGlobalFriendChatRow) SendTime() int64 {
	return int64(this.row.SendTime)
}

// 消息内容
func (this *PlayerGlobalFriendChatRow) Message() string {
	if !this.cached_Message {
		this.cached_Message = true
		this.cacheof_Message = C.GoStringN(this.row.Message, this.row.Message_len)
	}
	return this.cacheof_Message
}

func (this *PlayerGlobalFriendChatRow) GoObject() *PlayerGlobalFriendChat {
	return &PlayerGlobalFriendChat{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		FriendPid: int64(this.row.FriendPid),
		Mode: int8(this.row.Mode),
		SendTime: int64(this.row.SendTime),
		Message: this.Message(),
	}
}

/* ========== player_global_friend_state ========== */

// 玩家好友功能状态数据
type PlayerGlobalFriendStateRow struct {
	row *C.PlayerGlobalFriendState
	isBreak bool
}

func (this *PlayerGlobalFriendStateRow) reset(row *C.PlayerGlobalFriendState) {	this.row = row
	this.isBreak = false
}

func (this *PlayerGlobalFriendStateRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerGlobalFriendStateRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 每日删除好友数量
func (this *PlayerGlobalFriendStateRow) DeleteDayCount() int32 {
	return int32(this.row.DeleteDayCount)
}

// 删除好友时间
func (this *PlayerGlobalFriendStateRow) DeleteTime() int64 {
	return int64(this.row.DeleteTime)
}

// 0没有离线消息，1有离线消息
func (this *PlayerGlobalFriendStateRow) ExistOfflineChat() int8 {
	return int8(this.row.ExistOfflineChat)
}

// 平台好友历史最大值
func (this *PlayerGlobalFriendStateRow) PlatformFriendNum() int32 {
	return int32(this.row.PlatformFriendNum)
}

func (this *PlayerGlobalFriendStateRow) GoObject() *PlayerGlobalFriendState {
	return &PlayerGlobalFriendState{
		Pid: int64(this.row.Pid),
		DeleteDayCount: int32(this.row.DeleteDayCount),
		DeleteTime: int64(this.row.DeleteTime),
		ExistOfflineChat: int8(this.row.ExistOfflineChat),
		PlatformFriendNum: int32(this.row.PlatformFriendNum),
	}
}

/* ========== player_global_gift_card_record ========== */

// 玩家兑换记录
type PlayerGlobalGiftCardRecordRow struct {
	row *C.PlayerGlobalGiftCardRecord
	isBreak bool
	cached_Code bool
	cacheof_Code string
}

func (this *PlayerGlobalGiftCardRecordRow) reset(row *C.PlayerGlobalGiftCardRecord) {	this.row = row
	this.isBreak = false
	this.cached_Code = false
}

func (this *PlayerGlobalGiftCardRecordRow) Break() {
	this.isBreak = true
}

// 主键
func (this *PlayerGlobalGiftCardRecordRow) Id() int64 {
	return int64(this.row.Id)
}

// 兑换玩家PID
func (this *PlayerGlobalGiftCardRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 兑换码
func (this *PlayerGlobalGiftCardRecordRow) Code() string {
	if !this.cached_Code {
		this.cached_Code = true
		this.cacheof_Code = C.GoStringN(this.row.Code, this.row.Code_len)
	}
	return this.cacheof_Code
}

func (this *PlayerGlobalGiftCardRecordRow) GoObject() *PlayerGlobalGiftCardRecord {
	return &PlayerGlobalGiftCardRecord{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		Code: this.Code(),
	}
}

/* ========== player_global_mail_state ========== */

// 玩家全局邮件记录
type PlayerGlobalMailStateRow struct {
	row *C.PlayerGlobalMailState
	isBreak bool
}

func (this *PlayerGlobalMailStateRow) reset(row *C.PlayerGlobalMailState) {	this.row = row
	this.isBreak = false
}

func (this *PlayerGlobalMailStateRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerGlobalMailStateRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 发送时间戳
func (this *PlayerGlobalMailStateRow) MaxTimestamp() int64 {
	return int64(this.row.MaxTimestamp)
}

func (this *PlayerGlobalMailStateRow) GoObject() *PlayerGlobalMailState {
	return &PlayerGlobalMailState{
		Pid: int64(this.row.Pid),
		MaxTimestamp: int64(this.row.MaxTimestamp),
	}
}

/* ========== player_global_world_chat_state ========== */

// 玩家世界聊天状态
type PlayerGlobalWorldChatStateRow struct {
	row *C.PlayerGlobalWorldChatState
	isBreak bool
}

func (this *PlayerGlobalWorldChatStateRow) reset(row *C.PlayerGlobalWorldChatState) {	this.row = row
	this.isBreak = false
}

func (this *PlayerGlobalWorldChatStateRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerGlobalWorldChatStateRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 最近一次发送聊天时间戳
func (this *PlayerGlobalWorldChatStateRow) Timestamp() int64 {
	return int64(this.row.Timestamp)
}

// 今日次数
func (this *PlayerGlobalWorldChatStateRow) DailyNum() int16 {
	return int16(this.row.DailyNum)
}

func (this *PlayerGlobalWorldChatStateRow) GoObject() *PlayerGlobalWorldChatState {
	return &PlayerGlobalWorldChatState{
		Pid: int64(this.row.Pid),
		Timestamp: int64(this.row.Timestamp),
		DailyNum: int16(this.row.DailyNum),
	}
}

/* ========== player_hard_level ========== */

// 难度关卡功能权值
type PlayerHardLevelRow struct {
	row *C.PlayerHardLevel
	isBreak bool
}

func (this *PlayerHardLevelRow) reset(row *C.PlayerHardLevel) {	this.row = row
	this.isBreak = false
}

func (this *PlayerHardLevelRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerHardLevelRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 难度关卡功能权值
func (this *PlayerHardLevelRow) Lock() int32 {
	return int32(this.row.Lock)
}

// 已获得过奖励关卡的最大lock
func (this *PlayerHardLevelRow) AwardLock() int32 {
	return int32(this.row.AwardLock)
}

func (this *PlayerHardLevelRow) GoObject() *PlayerHardLevel {
	return &PlayerHardLevel{
		Pid: int64(this.row.Pid),
		Lock: int32(this.row.Lock),
		AwardLock: int32(this.row.AwardLock),
	}
}

/* ========== player_hard_level_record ========== */

// 难度关卡记录
type PlayerHardLevelRecordRow struct {
	row *C.PlayerHardLevelRecord
	isBreak bool
}

func (this *PlayerHardLevelRecordRow) reset(row *C.PlayerHardLevelRecord) {	this.row = row
	this.isBreak = false
}

func (this *PlayerHardLevelRecordRow) Break() {
	this.isBreak = true
}

// 记录ID
func (this *PlayerHardLevelRecordRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerHardLevelRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 开启的关卡ID
func (this *PlayerHardLevelRecordRow) LevelId() int32 {
	return int32(this.row.LevelId)
}

// 关卡开启时间
func (this *PlayerHardLevelRecordRow) OpenTime() int64 {
	return int64(this.row.OpenTime)
}

// 得分
func (this *PlayerHardLevelRecordRow) Score() int32 {
	return int32(this.row.Score)
}

// 通关回合数
func (this *PlayerHardLevelRecordRow) Round() int8 {
	return int8(this.row.Round)
}

// 当日已进入关卡的次数
func (this *PlayerHardLevelRecordRow) DailyNum() int8 {
	return int8(this.row.DailyNum)
}

// 最后一次进入时间
func (this *PlayerHardLevelRecordRow) LastEnterTime() int64 {
	return int64(this.row.LastEnterTime)
}

// 深渊关卡今日购买次数
func (this *PlayerHardLevelRecordRow) BuyTimes() int16 {
	return int16(this.row.BuyTimes)
}

// 深渊关卡上次购买时间戳
func (this *PlayerHardLevelRecordRow) BuyUpdateTime() int64 {
	return int64(this.row.BuyUpdateTime)
}

func (this *PlayerHardLevelRecordRow) GoObject() *PlayerHardLevelRecord {
	return &PlayerHardLevelRecord{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		LevelId: int32(this.row.LevelId),
		OpenTime: int64(this.row.OpenTime),
		Score: int32(this.row.Score),
		Round: int8(this.row.Round),
		DailyNum: int8(this.row.DailyNum),
		LastEnterTime: int64(this.row.LastEnterTime),
		BuyTimes: int16(this.row.BuyTimes),
		BuyUpdateTime: int64(this.row.BuyUpdateTime),
	}
}

/* ========== player_heart ========== */

// 玩家爱心表
type PlayerHeartRow struct {
	row *C.PlayerHeart
	isBreak bool
}

func (this *PlayerHeartRow) reset(row *C.PlayerHeart) {	this.row = row
	this.isBreak = false
}

func (this *PlayerHeartRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerHeartRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 爱心值
func (this *PlayerHeartRow) Value() int16 {
	return int16(this.row.Value)
}

// 最后更新时间
func (this *PlayerHeartRow) UpdateTime() int64 {
	return int64(this.row.UpdateTime)
}

// 今日好友赠送数量
func (this *PlayerHeartRow) AddDayCount() int32 {
	return int32(this.row.AddDayCount)
}

// 接受好友赠送爱心时间
func (this *PlayerHeartRow) AddTime() int64 {
	return int64(this.row.AddTime)
}

// 今日恢复数量
func (this *PlayerHeartRow) RecoverDayCount() int16 {
	return int16(this.row.RecoverDayCount)
}

func (this *PlayerHeartRow) GoObject() *PlayerHeart {
	return &PlayerHeart{
		Pid: int64(this.row.Pid),
		Value: int16(this.row.Value),
		UpdateTime: int64(this.row.UpdateTime),
		AddDayCount: int32(this.row.AddDayCount),
		AddTime: int64(this.row.AddTime),
		RecoverDayCount: int16(this.row.RecoverDayCount),
	}
}

/* ========== player_heart_draw ========== */

// 玩家爱心抽奖
type PlayerHeartDrawRow struct {
	row *C.PlayerHeartDraw
	isBreak bool
}

func (this *PlayerHeartDrawRow) reset(row *C.PlayerHeartDraw) {	this.row = row
	this.isBreak = false
}

func (this *PlayerHeartDrawRow) Break() {
	this.isBreak = true
}

// 
func (this *PlayerHeartDrawRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerHeartDrawRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 抽奖类型（1-大转盘；2-刮刮卡）
func (this *PlayerHeartDrawRow) DrawType() int8 {
	return int8(this.row.DrawType)
}

// 当日已抽次数
func (this *PlayerHeartDrawRow) DailyNum() int8 {
	return int8(this.row.DailyNum)
}

// 最近一次抽奖时间
func (this *PlayerHeartDrawRow) DrawTime() int64 {
	return int64(this.row.DrawTime)
}

func (this *PlayerHeartDrawRow) GoObject() *PlayerHeartDraw {
	return &PlayerHeartDraw{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		DrawType: int8(this.row.DrawType),
		DailyNum: int8(this.row.DailyNum),
		DrawTime: int64(this.row.DrawTime),
	}
}

/* ========== player_heart_draw_card_record ========== */

// 玩家爱心刮刮卡抽奖记录
type PlayerHeartDrawCardRecordRow struct {
	row *C.PlayerHeartDrawCardRecord
	isBreak bool
}

func (this *PlayerHeartDrawCardRecordRow) reset(row *C.PlayerHeartDrawCardRecord) {	this.row = row
	this.isBreak = false
}

func (this *PlayerHeartDrawCardRecordRow) Break() {
	this.isBreak = true
}

// 
func (this *PlayerHeartDrawCardRecordRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerHeartDrawCardRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 奖品类型（1-铜钱；2-元宝；3-道具）
func (this *PlayerHeartDrawCardRecordRow) AwardType() int8 {
	return int8(this.row.AwardType)
}

// 奖品数量
func (this *PlayerHeartDrawCardRecordRow) AwardNum() int16 {
	return int16(this.row.AwardNum)
}

// 道具奖品ID
func (this *PlayerHeartDrawCardRecordRow) ItemId() int16 {
	return int16(this.row.ItemId)
}

// 抽奖时间
func (this *PlayerHeartDrawCardRecordRow) DrawTime() int64 {
	return int64(this.row.DrawTime)
}

func (this *PlayerHeartDrawCardRecordRow) GoObject() *PlayerHeartDrawCardRecord {
	return &PlayerHeartDrawCardRecord{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		AwardType: int8(this.row.AwardType),
		AwardNum: int16(this.row.AwardNum),
		ItemId: int16(this.row.ItemId),
		DrawTime: int64(this.row.DrawTime),
	}
}

/* ========== player_heart_draw_wheel_record ========== */

// 玩家爱心大转盘抽奖记录
type PlayerHeartDrawWheelRecordRow struct {
	row *C.PlayerHeartDrawWheelRecord
	isBreak bool
}

func (this *PlayerHeartDrawWheelRecordRow) reset(row *C.PlayerHeartDrawWheelRecord) {	this.row = row
	this.isBreak = false
}

func (this *PlayerHeartDrawWheelRecordRow) Break() {
	this.isBreak = true
}

// 
func (this *PlayerHeartDrawWheelRecordRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerHeartDrawWheelRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 奖品类型（1-铜钱；2-元宝；3-道具）
func (this *PlayerHeartDrawWheelRecordRow) AwardType() int8 {
	return int8(this.row.AwardType)
}

// 奖品数量
func (this *PlayerHeartDrawWheelRecordRow) AwardNum() int16 {
	return int16(this.row.AwardNum)
}

// 道具奖品ID
func (this *PlayerHeartDrawWheelRecordRow) ItemId() int16 {
	return int16(this.row.ItemId)
}

// 抽奖时间
func (this *PlayerHeartDrawWheelRecordRow) DrawTime() int64 {
	return int64(this.row.DrawTime)
}

func (this *PlayerHeartDrawWheelRecordRow) GoObject() *PlayerHeartDrawWheelRecord {
	return &PlayerHeartDrawWheelRecord{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		AwardType: int8(this.row.AwardType),
		AwardNum: int16(this.row.AwardNum),
		ItemId: int16(this.row.ItemId),
		DrawTime: int64(this.row.DrawTime),
	}
}

/* ========== player_info ========== */

// 玩家信息表
type PlayerInfoRow struct {
	row *C.PlayerInfo
	isBreak bool
}

func (this *PlayerInfoRow) reset(row *C.PlayerInfo) {	this.row = row
	this.isBreak = false
}

func (this *PlayerInfoRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerInfoRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 元宝
func (this *PlayerInfoRow) Ingot() int64 {
	return int64(this.row.Ingot)
}

// 铜钱
func (this *PlayerInfoRow) Coins() int64 {
	return int64(this.row.Coins)
}

// 新邮件数
func (this *PlayerInfoRow) NewMailNum() int16 {
	return int16(this.row.NewMailNum)
}

// 上次登录时间
func (this *PlayerInfoRow) LastLoginTime() int64 {
	return int64(this.row.LastLoginTime)
}

// 上次下线时间
func (this *PlayerInfoRow) LastOfflineTime() int64 {
	return int64(this.row.LastOfflineTime)
}

// 总在线时间
func (this *PlayerInfoRow) TotalOnlineTime() int64 {
	return int64(this.row.TotalOnlineTime)
}

// 玩家注册时间
func (this *PlayerInfoRow) FirstLoginTime() int64 {
	return int64(this.row.FirstLoginTime)
}

// 玩家离线比武场战报数
func (this *PlayerInfoRow) NewArenaReportNum() int16 {
	return int16(this.row.NewArenaReportNum)
}

// 最近一次洗点时间
func (this *PlayerInfoRow) LastSkillFlush() int64 {
	return int64(this.row.LastSkillFlush)
}

// 剑心碎片数量
func (this *PlayerInfoRow) SwordSoulFragment() int64 {
	return int64(this.row.SwordSoulFragment)
}

func (this *PlayerInfoRow) GoObject() *PlayerInfo {
	return &PlayerInfo{
		Pid: int64(this.row.Pid),
		Ingot: int64(this.row.Ingot),
		Coins: int64(this.row.Coins),
		NewMailNum: int16(this.row.NewMailNum),
		LastLoginTime: int64(this.row.LastLoginTime),
		LastOfflineTime: int64(this.row.LastOfflineTime),
		TotalOnlineTime: int64(this.row.TotalOnlineTime),
		FirstLoginTime: int64(this.row.FirstLoginTime),
		NewArenaReportNum: int16(this.row.NewArenaReportNum),
		LastSkillFlush: int64(this.row.LastSkillFlush),
		SwordSoulFragment: int64(this.row.SwordSoulFragment),
	}
}

/* ========== player_is_operated ========== */

// 记录玩家是否第一次操作
type PlayerIsOperatedRow struct {
	row *C.PlayerIsOperated
	isBreak bool
}

func (this *PlayerIsOperatedRow) reset(row *C.PlayerIsOperated) {	this.row = row
	this.isBreak = false
}

func (this *PlayerIsOperatedRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerIsOperatedRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 操作值
func (this *PlayerIsOperatedRow) OperateValue() int64 {
	return int64(this.row.OperateValue)
}

func (this *PlayerIsOperatedRow) GoObject() *PlayerIsOperated {
	return &PlayerIsOperated{
		Pid: int64(this.row.Pid),
		OperateValue: int64(this.row.OperateValue),
	}
}

/* ========== player_item ========== */

// 玩家物品
type PlayerItemRow struct {
	row *C.PlayerItem
	isBreak bool
}

func (this *PlayerItemRow) reset(row *C.PlayerItem) {	this.row = row
	this.isBreak = false
}

func (this *PlayerItemRow) Break() {
	this.isBreak = true
}

// 主键ID
func (this *PlayerItemRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerItemRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 物品ID
func (this *PlayerItemRow) ItemId() int16 {
	return int16(this.row.ItemId)
}

// 数量
func (this *PlayerItemRow) Num() int16 {
	return int16(this.row.Num)
}

// 是否被装备
func (this *PlayerItemRow) IsDressed() int8 {
	return int8(this.row.IsDressed)
}

// 记录物品是否在回购栏
func (this *PlayerItemRow) BuyBackState() int8 {
	return int8(this.row.BuyBackState)
}

// 附加属性ID
func (this *PlayerItemRow) AppendixId() int64 {
	return int64(this.row.AppendixId)
}

// 精炼等级备份
func (this *PlayerItemRow) RefineLevelBak() int16 {
	return int16(this.row.RefineLevelBak)
}

// 精炼失败次数
func (this *PlayerItemRow) RefineFailTimes() int16 {
	return int16(this.row.RefineFailTimes)
}

// 精炼等级
func (this *PlayerItemRow) RefineLevel() int16 {
	return int16(this.row.RefineLevel)
}

// 装备精炼价格
func (this *PlayerItemRow) Price() int32 {
	return int32(this.row.Price)
}

func (this *PlayerItemRow) GoObject() *PlayerItem {
	return &PlayerItem{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		ItemId: int16(this.row.ItemId),
		Num: int16(this.row.Num),
		IsDressed: int8(this.row.IsDressed),
		BuyBackState: int8(this.row.BuyBackState),
		AppendixId: int64(this.row.AppendixId),
		RefineLevelBak: int16(this.row.RefineLevelBak),
		RefineFailTimes: int16(this.row.RefineFailTimes),
		RefineLevel: int16(this.row.RefineLevel),
		Price: int32(this.row.Price),
	}
}

/* ========== player_item_appendix ========== */

// 玩家装备追加属性表
type PlayerItemAppendixRow struct {
	row *C.PlayerItemAppendix
	isBreak bool
}

func (this *PlayerItemAppendixRow) reset(row *C.PlayerItemAppendix) {	this.row = row
	this.isBreak = false
}

func (this *PlayerItemAppendixRow) Break() {
	this.isBreak = true
}

// ID
func (this *PlayerItemAppendixRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerItemAppendixRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 生命
func (this *PlayerItemAppendixRow) Health() int32 {
	return int32(this.row.Health)
}

// 内力
func (this *PlayerItemAppendixRow) Cultivation() int32 {
	return int32(this.row.Cultivation)
}

// 速度
func (this *PlayerItemAppendixRow) Speed() int32 {
	return int32(this.row.Speed)
}

// 攻击
func (this *PlayerItemAppendixRow) Attack() int32 {
	return int32(this.row.Attack)
}

// 防御
func (this *PlayerItemAppendixRow) Defence() int32 {
	return int32(this.row.Defence)
}

// 闪避
func (this *PlayerItemAppendixRow) DodgeLevel() int32 {
	return int32(this.row.DodgeLevel)
}

// 命中
func (this *PlayerItemAppendixRow) HitLevel() int32 {
	return int32(this.row.HitLevel)
}

// 格挡
func (this *PlayerItemAppendixRow) BlockLevel() int32 {
	return int32(this.row.BlockLevel)
}

// 暴击
func (this *PlayerItemAppendixRow) CriticalLevel() int32 {
	return int32(this.row.CriticalLevel)
}

// 韧性
func (this *PlayerItemAppendixRow) TenacityLevel() int32 {
	return int32(this.row.TenacityLevel)
}

// 破击
func (this *PlayerItemAppendixRow) DestroyLevel() int32 {
	return int32(this.row.DestroyLevel)
}

// 重铸属性
func (this *PlayerItemAppendixRow) RecastAttr() int8 {
	return int8(this.row.RecastAttr)
}

func (this *PlayerItemAppendixRow) GoObject() *PlayerItemAppendix {
	return &PlayerItemAppendix{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		Health: int32(this.row.Health),
		Cultivation: int32(this.row.Cultivation),
		Speed: int32(this.row.Speed),
		Attack: int32(this.row.Attack),
		Defence: int32(this.row.Defence),
		DodgeLevel: int32(this.row.DodgeLevel),
		HitLevel: int32(this.row.HitLevel),
		BlockLevel: int32(this.row.BlockLevel),
		CriticalLevel: int32(this.row.CriticalLevel),
		TenacityLevel: int32(this.row.TenacityLevel),
		DestroyLevel: int32(this.row.DestroyLevel),
		RecastAttr: int8(this.row.RecastAttr),
	}
}

/* ========== player_item_buyback ========== */

// 玩家物品回购表
type PlayerItemBuybackRow struct {
	row *C.PlayerItemBuyback
	isBreak bool
}

func (this *PlayerItemBuybackRow) reset(row *C.PlayerItemBuyback) {	this.row = row
	this.isBreak = false
}

func (this *PlayerItemBuybackRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerItemBuybackRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 回购格子1,player_item表主键ID
func (this *PlayerItemBuybackRow) BackId1() int64 {
	return int64(this.row.BackId1)
}

// 回购格子2,player_item表主键ID
func (this *PlayerItemBuybackRow) BackId2() int64 {
	return int64(this.row.BackId2)
}

// 回购格子3,player_item表主键ID
func (this *PlayerItemBuybackRow) BackId3() int64 {
	return int64(this.row.BackId3)
}

// 回购格子4,player_item表主键ID
func (this *PlayerItemBuybackRow) BackId4() int64 {
	return int64(this.row.BackId4)
}

// 回购格子5,player_item表主键ID
func (this *PlayerItemBuybackRow) BackId5() int64 {
	return int64(this.row.BackId5)
}

// 回购格子6,player_item表主键ID
func (this *PlayerItemBuybackRow) BackId6() int64 {
	return int64(this.row.BackId6)
}

func (this *PlayerItemBuybackRow) GoObject() *PlayerItemBuyback {
	return &PlayerItemBuyback{
		Pid: int64(this.row.Pid),
		BackId1: int64(this.row.BackId1),
		BackId2: int64(this.row.BackId2),
		BackId3: int64(this.row.BackId3),
		BackId4: int64(this.row.BackId4),
		BackId5: int64(this.row.BackId5),
		BackId6: int64(this.row.BackId6),
	}
}

/* ========== player_item_recast_state ========== */

// 玩家装备重铸状态
type PlayerItemRecastStateRow struct {
	row *C.PlayerItemRecastState
	isBreak bool
}

func (this *PlayerItemRecastStateRow) reset(row *C.PlayerItemRecastState) {	this.row = row
	this.isBreak = false
}

func (this *PlayerItemRecastStateRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerItemRecastStateRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 玩家装备ID
func (this *PlayerItemRecastStateRow) PlayerItemId() int64 {
	return int64(this.row.PlayerItemId)
}

// 选中的属性
func (this *PlayerItemRecastStateRow) SelectedAttr() int8 {
	return int8(this.row.SelectedAttr)
}

// 重铸属性1类型
func (this *PlayerItemRecastStateRow) Attr1Type() int8 {
	return int8(this.row.Attr1Type)
}

// 重铸属性1数值
func (this *PlayerItemRecastStateRow) Attr1Value() int32 {
	return int32(this.row.Attr1Value)
}

// 重铸属性2类型
func (this *PlayerItemRecastStateRow) Attr2Type() int8 {
	return int8(this.row.Attr2Type)
}

// 重铸属性2数值
func (this *PlayerItemRecastStateRow) Attr2Value() int32 {
	return int32(this.row.Attr2Value)
}

// 重铸属性3类型
func (this *PlayerItemRecastStateRow) Attr3Type() int8 {
	return int8(this.row.Attr3Type)
}

// 重铸属性3数值
func (this *PlayerItemRecastStateRow) Attr3Value() int32 {
	return int32(this.row.Attr3Value)
}

func (this *PlayerItemRecastStateRow) GoObject() *PlayerItemRecastState {
	return &PlayerItemRecastState{
		Pid: int64(this.row.Pid),
		PlayerItemId: int64(this.row.PlayerItemId),
		SelectedAttr: int8(this.row.SelectedAttr),
		Attr1Type: int8(this.row.Attr1Type),
		Attr1Value: int32(this.row.Attr1Value),
		Attr2Type: int8(this.row.Attr2Type),
		Attr2Value: int32(this.row.Attr2Value),
		Attr3Type: int8(this.row.Attr3Type),
		Attr3Value: int32(this.row.Attr3Value),
	}
}

/* ========== player_level_award_info ========== */

// 玩家等级奖励领取状态表
type PlayerLevelAwardInfoRow struct {
	row *C.PlayerLevelAwardInfo
	isBreak bool
	cached_Awarded bool
	cacheof_Awarded string
}

func (this *PlayerLevelAwardInfoRow) reset(row *C.PlayerLevelAwardInfo) {	this.row = row
	this.isBreak = false
	this.cached_Awarded = false
}

func (this *PlayerLevelAwardInfoRow) Break() {
	this.isBreak = true
}

// 玩家id
func (this *PlayerLevelAwardInfoRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 玩家已领取状态
func (this *PlayerLevelAwardInfoRow) Awarded() string {
	if !this.cached_Awarded {
		this.cached_Awarded = true
		this.cacheof_Awarded = C.GoStringN(this.row.Awarded, this.row.Awarded_len)
	}
	return this.cacheof_Awarded
}

func (this *PlayerLevelAwardInfoRow) GoObject() *PlayerLevelAwardInfo {
	return &PlayerLevelAwardInfo{
		Pid: int64(this.row.Pid),
		Awarded: this.Awarded(),
	}
}

/* ========== player_login_award_record ========== */

// 玩家奖励领取情况
type PlayerLoginAwardRecordRow struct {
	row *C.PlayerLoginAwardRecord
	isBreak bool
}

func (this *PlayerLoginAwardRecordRow) reset(row *C.PlayerLoginAwardRecord) {	this.row = row
	this.isBreak = false
}

func (this *PlayerLoginAwardRecordRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerLoginAwardRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 累计活跃天数
func (this *PlayerLoginAwardRecordRow) ActiveDays() int32 {
	return int32(this.row.ActiveDays)
}

// 七天奖励领取记录
func (this *PlayerLoginAwardRecordRow) Record() int32 {
	return int32(this.row.Record)
}

// 更新时间戳
func (this *PlayerLoginAwardRecordRow) UpdateTimestamp() int64 {
	return int64(this.row.UpdateTimestamp)
}

func (this *PlayerLoginAwardRecordRow) GoObject() *PlayerLoginAwardRecord {
	return &PlayerLoginAwardRecord{
		Pid: int64(this.row.Pid),
		ActiveDays: int32(this.row.ActiveDays),
		Record: int32(this.row.Record),
		UpdateTimestamp: int64(this.row.UpdateTimestamp),
	}
}

/* ========== player_mail ========== */

// 玩家邮件表
type PlayerMailRow struct {
	row *C.PlayerMail
	isBreak bool
	cached_Parameters bool
	cacheof_Parameters string
	cached_Title bool
	cacheof_Title string
	cached_Content bool
	cacheof_Content string
}

func (this *PlayerMailRow) reset(row *C.PlayerMail) {	this.row = row
	this.isBreak = false
	this.cached_Parameters = false
	this.cached_Title = false
	this.cached_Content = false
}

func (this *PlayerMailRow) Break() {
	this.isBreak = true
}

// 玩家邮件ID
func (this *PlayerMailRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerMailRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 邮件模版ID
func (this *PlayerMailRow) MailId() int32 {
	return int32(this.row.MailId)
}

// 0未读，1已读
func (this *PlayerMailRow) State() int8 {
	return int8(this.row.State)
}

// 发送时间戳
func (this *PlayerMailRow) SendTime() int64 {
	return int64(this.row.SendTime)
}

// 模版参数
func (this *PlayerMailRow) Parameters() string {
	if !this.cached_Parameters {
		this.cached_Parameters = true
		this.cacheof_Parameters = C.GoStringN(this.row.Parameters, this.row.Parameters_len)
	}
	return this.cacheof_Parameters
}

// 是否有附件
func (this *PlayerMailRow) HaveAttachment() int8 {
	return int8(this.row.HaveAttachment)
}

// 标题
func (this *PlayerMailRow) Title() string {
	if !this.cached_Title {
		this.cached_Title = true
		this.cacheof_Title = C.GoStringN(this.row.Title, this.row.Title_len)
	}
	return this.cacheof_Title
}

// 内容
func (this *PlayerMailRow) Content() string {
	if !this.cached_Content {
		this.cached_Content = true
		this.cacheof_Content = C.GoStringN(this.row.Content, this.row.Content_len)
	}
	return this.cacheof_Content
}

// 邮件删除时机
func (this *PlayerMailRow) ExpireTime() int64 {
	return int64(this.row.ExpireTime)
}

// 优先级
func (this *PlayerMailRow) Priority() int8 {
	return int8(this.row.Priority)
}

func (this *PlayerMailRow) GoObject() *PlayerMail {
	return &PlayerMail{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		MailId: int32(this.row.MailId),
		State: int8(this.row.State),
		SendTime: int64(this.row.SendTime),
		Parameters: this.Parameters(),
		HaveAttachment: int8(this.row.HaveAttachment),
		Title: this.Title(),
		Content: this.Content(),
		ExpireTime: int64(this.row.ExpireTime),
		Priority: int8(this.row.Priority),
	}
}

/* ========== player_mail_attachment ========== */

// 玩家邮件附件表
type PlayerMailAttachmentRow struct {
	row *C.PlayerMailAttachment
	isBreak bool
}

func (this *PlayerMailAttachmentRow) reset(row *C.PlayerMailAttachment) {	this.row = row
	this.isBreak = false
}

func (this *PlayerMailAttachmentRow) Break() {
	this.isBreak = true
}

// 玩家邮件附件ID
func (this *PlayerMailAttachmentRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerMailAttachmentRow) Pid() int64 {
	return int64(this.row.Pid)
}

// player_mail 主键ID
func (this *PlayerMailAttachmentRow) PlayerMailId() int64 {
	return int64(this.row.PlayerMailId)
}

// 附件类型
func (this *PlayerMailAttachmentRow) AttachmentType() int8 {
	return int8(this.row.AttachmentType)
}

// 物品
func (this *PlayerMailAttachmentRow) ItemId() int16 {
	return int16(this.row.ItemId)
}

// 数量
func (this *PlayerMailAttachmentRow) ItemNum() int64 {
	return int64(this.row.ItemNum)
}

func (this *PlayerMailAttachmentRow) GoObject() *PlayerMailAttachment {
	return &PlayerMailAttachment{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		PlayerMailId: int64(this.row.PlayerMailId),
		AttachmentType: int8(this.row.AttachmentType),
		ItemId: int16(this.row.ItemId),
		ItemNum: int64(this.row.ItemNum),
	}
}

/* ========== player_mail_attachment_lg ========== */

// 玩家领取邮件附件表
type PlayerMailAttachmentLgRow struct {
	row *C.PlayerMailAttachmentLg
	isBreak bool
}

func (this *PlayerMailAttachmentLgRow) reset(row *C.PlayerMailAttachmentLg) {	this.row = row
	this.isBreak = false
}

func (this *PlayerMailAttachmentLgRow) Break() {
	this.isBreak = true
}

// 玩家邮件附件ID
func (this *PlayerMailAttachmentLgRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerMailAttachmentLgRow) Pid() int64 {
	return int64(this.row.Pid)
}

// player_mail 主键ID
func (this *PlayerMailAttachmentLgRow) PlayerMailId() int64 {
	return int64(this.row.PlayerMailId)
}

// 附件类型
func (this *PlayerMailAttachmentLgRow) AttachmentType() int8 {
	return int8(this.row.AttachmentType)
}

// 物品
func (this *PlayerMailAttachmentLgRow) ItemId() int16 {
	return int16(this.row.ItemId)
}

// 数量
func (this *PlayerMailAttachmentLgRow) ItemNum() int64 {
	return int64(this.row.ItemNum)
}

// 附件领取时间
func (this *PlayerMailAttachmentLgRow) TakeTimestamp() int64 {
	return int64(this.row.TakeTimestamp)
}

func (this *PlayerMailAttachmentLgRow) GoObject() *PlayerMailAttachmentLg {
	return &PlayerMailAttachmentLg{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		PlayerMailId: int64(this.row.PlayerMailId),
		AttachmentType: int8(this.row.AttachmentType),
		ItemId: int16(this.row.ItemId),
		ItemNum: int64(this.row.ItemNum),
		TakeTimestamp: int64(this.row.TakeTimestamp),
	}
}

/* ========== player_mail_lg ========== */

// 玩家已删除邮件表
type PlayerMailLgRow struct {
	row *C.PlayerMailLg
	isBreak bool
	cached_Parameters bool
	cacheof_Parameters string
	cached_Title bool
	cacheof_Title string
	cached_Content bool
	cacheof_Content string
}

func (this *PlayerMailLgRow) reset(row *C.PlayerMailLg) {	this.row = row
	this.isBreak = false
	this.cached_Parameters = false
	this.cached_Title = false
	this.cached_Content = false
}

func (this *PlayerMailLgRow) Break() {
	this.isBreak = true
}

// ID
func (this *PlayerMailLgRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家邮件ID
func (this *PlayerMailLgRow) Pmid() int64 {
	return int64(this.row.Pmid)
}

// 玩家ID
func (this *PlayerMailLgRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 邮件模版ID
func (this *PlayerMailLgRow) MailId() int32 {
	return int32(this.row.MailId)
}

// 0未读，1已读
func (this *PlayerMailLgRow) State() int8 {
	return int8(this.row.State)
}

// 发送时间戳
func (this *PlayerMailLgRow) SendTime() int64 {
	return int64(this.row.SendTime)
}

// 模版参数
func (this *PlayerMailLgRow) Parameters() string {
	if !this.cached_Parameters {
		this.cached_Parameters = true
		this.cacheof_Parameters = C.GoStringN(this.row.Parameters, this.row.Parameters_len)
	}
	return this.cacheof_Parameters
}

// 是否有附件
func (this *PlayerMailLgRow) HaveAttachment() int8 {
	return int8(this.row.HaveAttachment)
}

// 标题
func (this *PlayerMailLgRow) Title() string {
	if !this.cached_Title {
		this.cached_Title = true
		this.cacheof_Title = C.GoStringN(this.row.Title, this.row.Title_len)
	}
	return this.cacheof_Title
}

// 内容
func (this *PlayerMailLgRow) Content() string {
	if !this.cached_Content {
		this.cached_Content = true
		this.cacheof_Content = C.GoStringN(this.row.Content, this.row.Content_len)
	}
	return this.cacheof_Content
}

func (this *PlayerMailLgRow) GoObject() *PlayerMailLg {
	return &PlayerMailLg{
		Id: int64(this.row.Id),
		Pmid: int64(this.row.Pmid),
		Pid: int64(this.row.Pid),
		MailId: int32(this.row.MailId),
		State: int8(this.row.State),
		SendTime: int64(this.row.SendTime),
		Parameters: this.Parameters(),
		HaveAttachment: int8(this.row.HaveAttachment),
		Title: this.Title(),
		Content: this.Content(),
	}
}

/* ========== player_meditation_state ========== */

// 玩家打坐状态
type PlayerMeditationStateRow struct {
	row *C.PlayerMeditationState
	isBreak bool
}

func (this *PlayerMeditationStateRow) reset(row *C.PlayerMeditationState) {	this.row = row
	this.isBreak = false
}

func (this *PlayerMeditationStateRow) Break() {
	this.isBreak = true
}

// player id
func (this *PlayerMeditationStateRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 光明钥匙奖励累积时间
func (this *PlayerMeditationStateRow) AccumulateTime() int32 {
	return int32(this.row.AccumulateTime)
}

// 打坐开始时间 0-未未打坐状态
func (this *PlayerMeditationStateRow) StartTimestamp() int64 {
	return int64(this.row.StartTimestamp)
}

func (this *PlayerMeditationStateRow) GoObject() *PlayerMeditationState {
	return &PlayerMeditationState{
		Pid: int64(this.row.Pid),
		AccumulateTime: int32(this.row.AccumulateTime),
		StartTimestamp: int64(this.row.StartTimestamp),
	}
}

/* ========== player_mission ========== */

// 玩家区域数据
type PlayerMissionRow struct {
	row *C.PlayerMission
	isBreak bool
}

func (this *PlayerMissionRow) reset(row *C.PlayerMission) {	this.row = row
	this.isBreak = false
}

func (this *PlayerMissionRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerMissionRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 拥有的区域钥匙数
func (this *PlayerMissionRow) Key() int32 {
	return int32(this.row.Key)
}

// 已开启区域的最大序号
func (this *PlayerMissionRow) MaxOrder() int8 {
	return int8(this.row.MaxOrder)
}

func (this *PlayerMissionRow) GoObject() *PlayerMission {
	return &PlayerMission{
		Pid: int64(this.row.Pid),
		Key: int32(this.row.Key),
		MaxOrder: int8(this.row.MaxOrder),
	}
}

/* ========== player_mission_level ========== */

// 玩家区域关卡数据
type PlayerMissionLevelRow struct {
	row *C.PlayerMissionLevel
	isBreak bool
}

func (this *PlayerMissionLevelRow) reset(row *C.PlayerMissionLevel) {	this.row = row
	this.isBreak = false
}

func (this *PlayerMissionLevelRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerMissionLevelRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 当前的关卡权值
func (this *PlayerMissionLevelRow) Lock() int32 {
	return int32(this.row.Lock)
}

// 已开启的关卡最大权值
func (this *PlayerMissionLevelRow) MaxLock() int32 {
	return int32(this.row.MaxLock)
}

// 已获得过奖励关卡的最大lock
func (this *PlayerMissionLevelRow) AwardLock() int32 {
	return int32(this.row.AwardLock)
}

func (this *PlayerMissionLevelRow) GoObject() *PlayerMissionLevel {
	return &PlayerMissionLevel{
		Pid: int64(this.row.Pid),
		Lock: int32(this.row.Lock),
		MaxLock: int32(this.row.MaxLock),
		AwardLock: int32(this.row.AwardLock),
	}
}

/* ========== player_mission_level_record ========== */

// 关卡记录
type PlayerMissionLevelRecordRow struct {
	row *C.PlayerMissionLevelRecord
	isBreak bool
}

func (this *PlayerMissionLevelRecordRow) reset(row *C.PlayerMissionLevelRecord) {	this.row = row
	this.isBreak = false
}

func (this *PlayerMissionLevelRecordRow) Break() {
	this.isBreak = true
}

// 记录ID
func (this *PlayerMissionLevelRecordRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerMissionLevelRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 区域ID
func (this *PlayerMissionLevelRecordRow) MissionId() int16 {
	return int16(this.row.MissionId)
}

// 开启的关卡ID
func (this *PlayerMissionLevelRecordRow) MissionLevelId() int32 {
	return int32(this.row.MissionLevelId)
}

// 关卡开启时间
func (this *PlayerMissionLevelRecordRow) OpenTime() int64 {
	return int64(this.row.OpenTime)
}

// boss战得分
func (this *PlayerMissionLevelRecordRow) Score() int32 {
	return int32(this.row.Score)
}

// 通关回合数
func (this *PlayerMissionLevelRecordRow) Round() int8 {
	return int8(this.row.Round)
}

// 当日已进入关卡的次数
func (this *PlayerMissionLevelRecordRow) DailyNum() int8 {
	return int8(this.row.DailyNum)
}

// 最后一次进入时间
func (this *PlayerMissionLevelRecordRow) LastEnterTime() int64 {
	return int64(this.row.LastEnterTime)
}

// 清剿过的影之间隙
func (this *PlayerMissionLevelRecordRow) EmptyShadowBits() int16 {
	return int16(this.row.EmptyShadowBits)
}

func (this *PlayerMissionLevelRecordRow) GoObject() *PlayerMissionLevelRecord {
	return &PlayerMissionLevelRecord{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		MissionId: int16(this.row.MissionId),
		MissionLevelId: int32(this.row.MissionLevelId),
		OpenTime: int64(this.row.OpenTime),
		Score: int32(this.row.Score),
		Round: int8(this.row.Round),
		DailyNum: int8(this.row.DailyNum),
		LastEnterTime: int64(this.row.LastEnterTime),
		EmptyShadowBits: int16(this.row.EmptyShadowBits),
	}
}

/* ========== player_mission_level_state_bin ========== */

// 玩家区域关卡状态保存
type PlayerMissionLevelStateBinRow struct {
	row *C.PlayerMissionLevelStateBin
	isBreak bool
	cached_Bin bool
	cacheof_Bin []byte
}

func (this *PlayerMissionLevelStateBinRow) reset(row *C.PlayerMissionLevelStateBin) {	this.row = row
	this.isBreak = false
	this.cached_Bin = false
}

func (this *PlayerMissionLevelStateBinRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerMissionLevelStateBinRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 状态MissionLevelState
func (this *PlayerMissionLevelStateBinRow) Bin() []byte {
	if !this.cached_Bin {
		this.cached_Bin = true
		this.cacheof_Bin = C.GoBytes(unsafe.Pointer(this.row.Bin), this.row.Bin_len)
	}
	return this.cacheof_Bin
}

func (this *PlayerMissionLevelStateBinRow) GoObject() *PlayerMissionLevelStateBin {
	return &PlayerMissionLevelStateBin{
		Pid: int64(this.row.Pid),
		Bin: this.Bin(),
	}
}

/* ========== player_mission_record ========== */

// 区域记录
type PlayerMissionRecordRow struct {
	row *C.PlayerMissionRecord
	isBreak bool
}

func (this *PlayerMissionRecordRow) reset(row *C.PlayerMissionRecord) {	this.row = row
	this.isBreak = false
}

func (this *PlayerMissionRecordRow) Break() {
	this.isBreak = true
}

// 记录ID
func (this *PlayerMissionRecordRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerMissionRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 城镇ID
func (this *PlayerMissionRecordRow) TownId() int16 {
	return int16(this.row.TownId)
}

// 开启的区域ID
func (this *PlayerMissionRecordRow) MissionId() int16 {
	return int16(this.row.MissionId)
}

// 开启的区域时间
func (this *PlayerMissionRecordRow) OpenTime() int64 {
	return int64(this.row.OpenTime)
}

func (this *PlayerMissionRecordRow) GoObject() *PlayerMissionRecord {
	return &PlayerMissionRecord{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		TownId: int16(this.row.TownId),
		MissionId: int16(this.row.MissionId),
		OpenTime: int64(this.row.OpenTime),
	}
}

/* ========== player_mission_star_award ========== */

// 玩家区域评星领奖记录
type PlayerMissionStarAwardRow struct {
	row *C.PlayerMissionStarAward
	isBreak bool
}

func (this *PlayerMissionStarAwardRow) reset(row *C.PlayerMissionStarAward) {	this.row = row
	this.isBreak = false
}

func (this *PlayerMissionStarAwardRow) Break() {
	this.isBreak = true
}

// 主键
func (this *PlayerMissionStarAwardRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerMissionStarAwardRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 阵印ID
func (this *PlayerMissionStarAwardRow) TownId() int16 {
	return int16(this.row.TownId)
}

// 宝箱类型 1:铜 2:银 3:金
func (this *PlayerMissionStarAwardRow) BoxType() int8 {
	return int8(this.row.BoxType)
}

func (this *PlayerMissionStarAwardRow) GoObject() *PlayerMissionStarAward {
	return &PlayerMissionStarAward{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		TownId: int16(this.row.TownId),
		BoxType: int8(this.row.BoxType),
	}
}

/* ========== player_money_tree ========== */

// 玩家摇钱树记录
type PlayerMoneyTreeRow struct {
	row *C.PlayerMoneyTree
	isBreak bool
}

func (this *PlayerMoneyTreeRow) reset(row *C.PlayerMoneyTree) {	this.row = row
	this.isBreak = false
}

func (this *PlayerMoneyTreeRow) Break() {
	this.isBreak = true
}

// 玩家id
func (this *PlayerMoneyTreeRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 摇钱树铜钱总额
func (this *PlayerMoneyTreeRow) Total() int32 {
	return int32(this.row.Total)
}

// 总的摇树次数
func (this *PlayerMoneyTreeRow) WavedTimesTotal() int8 {
	return int8(this.row.WavedTimesTotal)
}

// 当天已经摇过的次数
func (this *PlayerMoneyTreeRow) WavedTimes() int8 {
	return int8(this.row.WavedTimes)
}

// 上次摇树的时间戳
func (this *PlayerMoneyTreeRow) LastWavedTime() int64 {
	return int64(this.row.LastWavedTime)
}

func (this *PlayerMoneyTreeRow) GoObject() *PlayerMoneyTree {
	return &PlayerMoneyTree{
		Pid: int64(this.row.Pid),
		Total: int32(this.row.Total),
		WavedTimesTotal: int8(this.row.WavedTimesTotal),
		WavedTimes: int8(this.row.WavedTimes),
		LastWavedTime: int64(this.row.LastWavedTime),
	}
}

/* ========== player_month_card_award_record ========== */

// 玩家最后领取月卡时间
type PlayerMonthCardAwardRecordRow struct {
	row *C.PlayerMonthCardAwardRecord
	isBreak bool
}

func (this *PlayerMonthCardAwardRecordRow) reset(row *C.PlayerMonthCardAwardRecord) {	this.row = row
	this.isBreak = false
}

func (this *PlayerMonthCardAwardRecordRow) Break() {
	this.isBreak = true
}

// 用户ID
func (this *PlayerMonthCardAwardRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 最后一次更新时间戳
func (this *PlayerMonthCardAwardRecordRow) LastUpdate() int64 {
	return int64(this.row.LastUpdate)
}

func (this *PlayerMonthCardAwardRecordRow) GoObject() *PlayerMonthCardAwardRecord {
	return &PlayerMonthCardAwardRecord{
		Pid: int64(this.row.Pid),
		LastUpdate: int64(this.row.LastUpdate),
	}
}

/* ========== player_month_card_info ========== */

// 玩家月卡信息
type PlayerMonthCardInfoRow struct {
	row *C.PlayerMonthCardInfo
	isBreak bool
}

func (this *PlayerMonthCardInfoRow) reset(row *C.PlayerMonthCardInfo) {	this.row = row
	this.isBreak = false
}

func (this *PlayerMonthCardInfoRow) Break() {
	this.isBreak = true
}

// 用户ID
func (this *PlayerMonthCardInfoRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 月卡开始时间
func (this *PlayerMonthCardInfoRow) Starttime() int64 {
	return int64(this.row.Starttime)
}

// 月卡结束时间
func (this *PlayerMonthCardInfoRow) Endtime() int64 {
	return int64(this.row.Endtime)
}

// 购买次数
func (this *PlayerMonthCardInfoRow) Buytimes() int64 {
	return int64(this.row.Buytimes)
}

// 赠送总金额
func (this *PlayerMonthCardInfoRow) Presenttotal() int64 {
	return int64(this.row.Presenttotal)
}

func (this *PlayerMonthCardInfoRow) GoObject() *PlayerMonthCardInfo {
	return &PlayerMonthCardInfo{
		Pid: int64(this.row.Pid),
		Starttime: int64(this.row.Starttime),
		Endtime: int64(this.row.Endtime),
		Buytimes: int64(this.row.Buytimes),
		Presenttotal: int64(this.row.Presenttotal),
	}
}

/* ========== player_multi_level_info ========== */

// 玩家多人关卡信息
type PlayerMultiLevelInfoRow struct {
	row *C.PlayerMultiLevelInfo
	isBreak bool
}

func (this *PlayerMultiLevelInfoRow) reset(row *C.PlayerMultiLevelInfo) {	this.row = row
	this.isBreak = false
}

func (this *PlayerMultiLevelInfoRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerMultiLevelInfoRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 上阵伙伴角色模板ID
func (this *PlayerMultiLevelInfoRow) BuddyRoleId() int8 {
	return int8(this.row.BuddyRoleId)
}

// 上阵伙伴所在行（1或2)
func (this *PlayerMultiLevelInfoRow) BuddyRow() int8 {
	return int8(this.row.BuddyRow)
}

// 战术
func (this *PlayerMultiLevelInfoRow) TacticalGrid() int8 {
	return int8(this.row.TacticalGrid)
}

// 今日已战斗次数
func (this *PlayerMultiLevelInfoRow) DailyNum() int8 {
	return int8(this.row.DailyNum)
}

// 最近一次战斗时间
func (this *PlayerMultiLevelInfoRow) BattleTime() int64 {
	return int64(this.row.BattleTime)
}

// 关卡开启权值
func (this *PlayerMultiLevelInfoRow) Lock() int32 {
	return int32(this.row.Lock)
}

func (this *PlayerMultiLevelInfoRow) GoObject() *PlayerMultiLevelInfo {
	return &PlayerMultiLevelInfo{
		Pid: int64(this.row.Pid),
		BuddyRoleId: int8(this.row.BuddyRoleId),
		BuddyRow: int8(this.row.BuddyRow),
		TacticalGrid: int8(this.row.TacticalGrid),
		DailyNum: int8(this.row.DailyNum),
		BattleTime: int64(this.row.BattleTime),
		Lock: int32(this.row.Lock),
	}
}

/* ========== player_new_year_consume_record ========== */

// 春节玩家消费记录
type PlayerNewYearConsumeRecordRow struct {
	row *C.PlayerNewYearConsumeRecord
	isBreak bool
	cached_ConsumeRecord bool
	cacheof_ConsumeRecord string
}

func (this *PlayerNewYearConsumeRecordRow) reset(row *C.PlayerNewYearConsumeRecord) {	this.row = row
	this.isBreak = false
	this.cached_ConsumeRecord = false
}

func (this *PlayerNewYearConsumeRecordRow) Break() {
	this.isBreak = true
}

// 玩家id
func (this *PlayerNewYearConsumeRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 玩家消费情况
func (this *PlayerNewYearConsumeRecordRow) ConsumeRecord() string {
	if !this.cached_ConsumeRecord {
		this.cached_ConsumeRecord = true
		this.cacheof_ConsumeRecord = C.GoStringN(this.row.ConsumeRecord, this.row.ConsumeRecord_len)
	}
	return this.cacheof_ConsumeRecord
}

func (this *PlayerNewYearConsumeRecordRow) GoObject() *PlayerNewYearConsumeRecord {
	return &PlayerNewYearConsumeRecord{
		Pid: int64(this.row.Pid),
		ConsumeRecord: this.ConsumeRecord(),
	}
}

/* ========== player_npc_talk_record ========== */

// 玩家与NPC对话奖励记录
type PlayerNpcTalkRecordRow struct {
	row *C.PlayerNpcTalkRecord
	isBreak bool
}

func (this *PlayerNpcTalkRecordRow) reset(row *C.PlayerNpcTalkRecord) {	this.row = row
	this.isBreak = false
}

func (this *PlayerNpcTalkRecordRow) Break() {
	this.isBreak = true
}

// 记录ID
func (this *PlayerNpcTalkRecordRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerNpcTalkRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// NPC ID
func (this *PlayerNpcTalkRecordRow) NpcId() int32 {
	return int32(this.row.NpcId)
}

// 相关城镇
func (this *PlayerNpcTalkRecordRow) TownId() int16 {
	return int16(this.row.TownId)
}

// 任务ID  首次对话作为特殊任务ID -1
func (this *PlayerNpcTalkRecordRow) QuestId() int16 {
	return int16(this.row.QuestId)
}

func (this *PlayerNpcTalkRecordRow) GoObject() *PlayerNpcTalkRecord {
	return &PlayerNpcTalkRecord{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		NpcId: int32(this.row.NpcId),
		TownId: int16(this.row.TownId),
		QuestId: int16(this.row.QuestId),
	}
}

/* ========== player_opened_town_treasure ========== */

// 玩家已开启的城镇宝箱
type PlayerOpenedTownTreasureRow struct {
	row *C.PlayerOpenedTownTreasure
	isBreak bool
}

func (this *PlayerOpenedTownTreasureRow) reset(row *C.PlayerOpenedTownTreasure) {	this.row = row
	this.isBreak = false
}

func (this *PlayerOpenedTownTreasureRow) Break() {
	this.isBreak = true
}

// 主键
func (this *PlayerOpenedTownTreasureRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerOpenedTownTreasureRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 城镇ID
func (this *PlayerOpenedTownTreasureRow) TownId() int16 {
	return int16(this.row.TownId)
}

func (this *PlayerOpenedTownTreasureRow) GoObject() *PlayerOpenedTownTreasure {
	return &PlayerOpenedTownTreasure{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		TownId: int16(this.row.TownId),
	}
}

/* ========== player_physical ========== */

// 玩家体力表
type PlayerPhysicalRow struct {
	row *C.PlayerPhysical
	isBreak bool
}

func (this *PlayerPhysicalRow) reset(row *C.PlayerPhysical) {	this.row = row
	this.isBreak = false
}

func (this *PlayerPhysicalRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerPhysicalRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 体力值
func (this *PlayerPhysicalRow) Value() int16 {
	return int16(this.row.Value)
}

// 最后更新时间
func (this *PlayerPhysicalRow) UpdateTime() int64 {
	return int64(this.row.UpdateTime)
}

// 购买次数
func (this *PlayerPhysicalRow) BuyCount() int64 {
	return int64(this.row.BuyCount)
}

// 购买次数更新时间
func (this *PlayerPhysicalRow) BuyUpdateTime() int64 {
	return int64(this.row.BuyUpdateTime)
}

// 当天购买次数
func (this *PlayerPhysicalRow) DailyCount() int16 {
	return int16(this.row.DailyCount)
}

func (this *PlayerPhysicalRow) GoObject() *PlayerPhysical {
	return &PlayerPhysical{
		Pid: int64(this.row.Pid),
		Value: int16(this.row.Value),
		UpdateTime: int64(this.row.UpdateTime),
		BuyCount: int64(this.row.BuyCount),
		BuyUpdateTime: int64(this.row.BuyUpdateTime),
		DailyCount: int16(this.row.DailyCount),
	}
}

/* ========== player_purchase_record ========== */

// 玩家物品购买记录 仅记录限购商品
type PlayerPurchaseRecordRow struct {
	row *C.PlayerPurchaseRecord
	isBreak bool
}

func (this *PlayerPurchaseRecordRow) reset(row *C.PlayerPurchaseRecord) {	this.row = row
	this.isBreak = false
}

func (this *PlayerPurchaseRecordRow) Break() {
	this.isBreak = true
}

// 主键ID
func (this *PlayerPurchaseRecordRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerPurchaseRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 物品ID
func (this *PlayerPurchaseRecordRow) ItemId() int16 {
	return int16(this.row.ItemId)
}

// 已购买数量
func (this *PlayerPurchaseRecordRow) Num() int16 {
	return int16(this.row.Num)
}

// 上次购买时间
func (this *PlayerPurchaseRecordRow) Timestamp() int64 {
	return int64(this.row.Timestamp)
}

func (this *PlayerPurchaseRecordRow) GoObject() *PlayerPurchaseRecord {
	return &PlayerPurchaseRecord{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		ItemId: int16(this.row.ItemId),
		Num: int16(this.row.Num),
		Timestamp: int64(this.row.Timestamp),
	}
}

/* ========== player_push_notify_switch ========== */

// 玩家推送通知开关列表
type PlayerPushNotifySwitchRow struct {
	row *C.PlayerPushNotifySwitch
	isBreak bool
}

func (this *PlayerPushNotifySwitchRow) reset(row *C.PlayerPushNotifySwitch) {	this.row = row
	this.isBreak = false
}

func (this *PlayerPushNotifySwitchRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerPushNotifySwitchRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 推送通知开关
func (this *PlayerPushNotifySwitchRow) Options() int64 {
	return int64(this.row.Options)
}

func (this *PlayerPushNotifySwitchRow) GoObject() *PlayerPushNotifySwitch {
	return &PlayerPushNotifySwitch{
		Pid: int64(this.row.Pid),
		Options: int64(this.row.Options),
	}
}

/* ========== player_pve_state ========== */

// 玩家灵宠状态
type PlayerPveStateRow struct {
	row *C.PlayerPveState
	isBreak bool
}

func (this *PlayerPveStateRow) reset(row *C.PlayerPveState) {	this.row = row
	this.isBreak = false
}

func (this *PlayerPveStateRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerPveStateRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 已通关最大层数
func (this *PlayerPveStateRow) MaxPassedFloor() int16 {
	return int16(this.row.MaxPassedFloor)
}

// 已奖励最大层数
func (this *PlayerPveStateRow) MaxAwardedFloor() int16 {
	return int16(this.row.MaxAwardedFloor)
}

// 未通关关卡杀敌数
func (this *PlayerPveStateRow) UnpassedFloorEnemyNum() int16 {
	return int16(this.row.UnpassedFloorEnemyNum)
}

// 最后一次进入关卡次数
func (this *PlayerPveStateRow) EnterTime() int64 {
	return int64(this.row.EnterTime)
}

// 今日进入次数
func (this *PlayerPveStateRow) DailyNum() int8 {
	return int8(this.row.DailyNum)
}

func (this *PlayerPveStateRow) GoObject() *PlayerPveState {
	return &PlayerPveState{
		Pid: int64(this.row.Pid),
		MaxPassedFloor: int16(this.row.MaxPassedFloor),
		MaxAwardedFloor: int16(this.row.MaxAwardedFloor),
		UnpassedFloorEnemyNum: int16(this.row.UnpassedFloorEnemyNum),
		EnterTime: int64(this.row.EnterTime),
		DailyNum: int8(this.row.DailyNum),
	}
}

/* ========== player_qq_vip_gift ========== */

// 玩家qq服务礼包领取记录
type PlayerQqVipGiftRow struct {
	row *C.PlayerQqVipGift
	isBreak bool
}

func (this *PlayerQqVipGiftRow) reset(row *C.PlayerQqVipGift) {	this.row = row
	this.isBreak = false
}

func (this *PlayerQqVipGiftRow) Break() {
	this.isBreak = true
}

// 
func (this *PlayerQqVipGiftRow) Pid() int64 {
	return int64(this.row.Pid)
}

// qq会员的礼包领取记录，1代表开通礼包已领取，2代表续费礼包已领取
func (this *PlayerQqVipGiftRow) Qqvip() int16 {
	return int16(this.row.Qqvip)
}

// qq超级会员礼包领取记录，值同上
func (this *PlayerQqVipGiftRow) Surper() int16 {
	return int16(this.row.Surper)
}

func (this *PlayerQqVipGiftRow) GoObject() *PlayerQqVipGift {
	return &PlayerQqVipGift{
		Pid: int64(this.row.Pid),
		Qqvip: int16(this.row.Qqvip),
		Surper: int16(this.row.Surper),
	}
}

/* ========== player_quest ========== */

// 玩家任务
type PlayerQuestRow struct {
	row *C.PlayerQuest
	isBreak bool
}

func (this *PlayerQuestRow) reset(row *C.PlayerQuest) {	this.row = row
	this.isBreak = false
}

func (this *PlayerQuestRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerQuestRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 当前任务ID
func (this *PlayerQuestRow) QuestId() int16 {
	return int16(this.row.QuestId)
}

// 
func (this *PlayerQuestRow) State() int8 {
	return int8(this.row.State)
}

func (this *PlayerQuestRow) GoObject() *PlayerQuest {
	return &PlayerQuest{
		Pid: int64(this.row.Pid),
		QuestId: int16(this.row.QuestId),
		State: int8(this.row.State),
	}
}

/* ========== player_rainbow_level ========== */

// 玩家彩虹关卡状态
type PlayerRainbowLevelRow struct {
	row *C.PlayerRainbowLevel
	isBreak bool
}

func (this *PlayerRainbowLevelRow) reset(row *C.PlayerRainbowLevel) {	this.row = row
	this.isBreak = false
}

func (this *PlayerRainbowLevelRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerRainbowLevelRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 重置彩虹关卡时间戳
func (this *PlayerRainbowLevelRow) ResetTimestamp() int64 {
	return int64(this.row.ResetTimestamp)
}

// 已重置次数
func (this *PlayerRainbowLevelRow) ResetNum() int32 {
	return int32(this.row.ResetNum)
}

// 段数
func (this *PlayerRainbowLevelRow) Segment() int16 {
	return int16(this.row.Segment)
}

// 彩虹关段内第X关顺序[1,7]
func (this *PlayerRainbowLevelRow) Order() int8 {
	return int8(this.row.Order)
}

// 状态 0--打败Boss  1--未领取奖励 2--已奖励未进入下一关卡
func (this *PlayerRainbowLevelRow) Status() int8 {
	return int8(this.row.Status)
}

// 最好记录段数
func (this *PlayerRainbowLevelRow) BestSegment() int16 {
	return int16(this.row.BestSegment)
}

// 最好记录关卡顺序
func (this *PlayerRainbowLevelRow) BestOrder() int8 {
	return int8(this.row.BestOrder)
}

// 最好记录时间戳
func (this *PlayerRainbowLevelRow) BestRecordTimestamp() int64 {
	return int64(this.row.BestRecordTimestamp)
}

// 可跳转的最大段数
func (this *PlayerRainbowLevelRow) MaxOpenSegment() int16 {
	return int16(this.row.MaxOpenSegment)
}

// 打通的最高段数
func (this *PlayerRainbowLevelRow) MaxPassSegment() int16 {
	return int16(this.row.MaxPassSegment)
}

// 今日扫荡次数
func (this *PlayerRainbowLevelRow) AutoFightNum() int8 {
	return int8(this.row.AutoFightNum)
}

// 今日购买次数
func (this *PlayerRainbowLevelRow) BuyTimes() int16 {
	return int16(this.row.BuyTimes)
}

// 扫荡时间
func (this *PlayerRainbowLevelRow) AutoFightTime() int64 {
	return int64(this.row.AutoFightTime)
}

// 购买彩虹关卡次数时间戳
func (this *PlayerRainbowLevelRow) BuyTimestamp() int64 {
	return int64(this.row.BuyTimestamp)
}

func (this *PlayerRainbowLevelRow) GoObject() *PlayerRainbowLevel {
	return &PlayerRainbowLevel{
		Pid: int64(this.row.Pid),
		ResetTimestamp: int64(this.row.ResetTimestamp),
		ResetNum: int32(this.row.ResetNum),
		Segment: int16(this.row.Segment),
		Order: int8(this.row.Order),
		Status: int8(this.row.Status),
		BestSegment: int16(this.row.BestSegment),
		BestOrder: int8(this.row.BestOrder),
		BestRecordTimestamp: int64(this.row.BestRecordTimestamp),
		MaxOpenSegment: int16(this.row.MaxOpenSegment),
		MaxPassSegment: int16(this.row.MaxPassSegment),
		AutoFightNum: int8(this.row.AutoFightNum),
		BuyTimes: int16(this.row.BuyTimes),
		AutoFightTime: int64(this.row.AutoFightTime),
		BuyTimestamp: int64(this.row.BuyTimestamp),
	}
}

/* ========== player_rainbow_level_state_bin ========== */

// 玩家彩虹关卡状态
type PlayerRainbowLevelStateBinRow struct {
	row *C.PlayerRainbowLevelStateBin
	isBreak bool
	cached_Bin bool
	cacheof_Bin []byte
}

func (this *PlayerRainbowLevelStateBinRow) reset(row *C.PlayerRainbowLevelStateBin) {	this.row = row
	this.isBreak = false
	this.cached_Bin = false
}

func (this *PlayerRainbowLevelStateBinRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerRainbowLevelStateBinRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 彩虹状态
func (this *PlayerRainbowLevelStateBinRow) Bin() []byte {
	if !this.cached_Bin {
		this.cached_Bin = true
		this.cacheof_Bin = C.GoBytes(unsafe.Pointer(this.row.Bin), this.row.Bin_len)
	}
	return this.cacheof_Bin
}

func (this *PlayerRainbowLevelStateBinRow) GoObject() *PlayerRainbowLevelStateBin {
	return &PlayerRainbowLevelStateBin{
		Pid: int64(this.row.Pid),
		Bin: this.Bin(),
	}
}

/* ========== player_role ========== */

// 玩家角色数据
type PlayerRoleRow struct {
	row *C.PlayerRole
	isBreak bool
}

func (this *PlayerRoleRow) reset(row *C.PlayerRole) {	this.row = row
	this.isBreak = false
}

func (this *PlayerRoleRow) Break() {
	this.isBreak = true
}

// 玩家角色ID
func (this *PlayerRoleRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerRoleRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 角色模板ID
func (this *PlayerRoleRow) RoleId() int8 {
	return int8(this.row.RoleId)
}

// 等级
func (this *PlayerRoleRow) Level() int16 {
	return int16(this.row.Level)
}

// 经验
func (this *PlayerRoleRow) Exp() int64 {
	return int64(this.row.Exp)
}

// 角色的羁绊等级
func (this *PlayerRoleRow) FriendshipLevel() int32 {
	return int32(this.row.FriendshipLevel)
}

// 伙伴状态，0表示正常，1表示在客栈
func (this *PlayerRoleRow) Status() int16 {
	return int16(this.row.Status)
}

// 协力组合
func (this *PlayerRoleRow) CoopId() int16 {
	return int16(this.row.CoopId)
}

// 等级变化时间
func (this *PlayerRoleRow) Timestamp() int64 {
	return int64(this.row.Timestamp)
}

// 角色战力
func (this *PlayerRoleRow) FightNum() int32 {
	return int32(this.row.FightNum)
}

func (this *PlayerRoleRow) GoObject() *PlayerRole {
	return &PlayerRole{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		RoleId: int8(this.row.RoleId),
		Level: int16(this.row.Level),
		Exp: int64(this.row.Exp),
		FriendshipLevel: int32(this.row.FriendshipLevel),
		Status: int16(this.row.Status),
		CoopId: int16(this.row.CoopId),
		Timestamp: int64(this.row.Timestamp),
		FightNum: int32(this.row.FightNum),
	}
}

/* ========== player_sealedbook ========== */

// 玩家天书记录
type PlayerSealedbookRow struct {
	row *C.PlayerSealedbook
	isBreak bool
	cached_SealedRecord bool
	cacheof_SealedRecord []byte
}

func (this *PlayerSealedbookRow) reset(row *C.PlayerSealedbook) {	this.row = row
	this.isBreak = false
	this.cached_SealedRecord = false
}

func (this *PlayerSealedbookRow) Break() {
	this.isBreak = true
}

// 用户ID
func (this *PlayerSealedbookRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 玩家天书记录
func (this *PlayerSealedbookRow) SealedRecord() []byte {
	if !this.cached_SealedRecord {
		this.cached_SealedRecord = true
		this.cacheof_SealedRecord = C.GoBytes(unsafe.Pointer(this.row.SealedRecord), this.row.SealedRecord_len)
	}
	return this.cacheof_SealedRecord
}

func (this *PlayerSealedbookRow) GoObject() *PlayerSealedbook {
	return &PlayerSealedbook{
		Pid: int64(this.row.Pid),
		SealedRecord: this.SealedRecord(),
	}
}

/* ========== player_send_heart_record ========== */

// 玩家好友列表
type PlayerSendHeartRecordRow struct {
	row *C.PlayerSendHeartRecord
	isBreak bool
}

func (this *PlayerSendHeartRecordRow) reset(row *C.PlayerSendHeartRecord) {	this.row = row
	this.isBreak = false
}

func (this *PlayerSendHeartRecordRow) Break() {
	this.isBreak = true
}

// ID
func (this *PlayerSendHeartRecordRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerSendHeartRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 好友ID
func (this *PlayerSendHeartRecordRow) FriendPid() int64 {
	return int64(this.row.FriendPid)
}

// 上次送爱心时间
func (this *PlayerSendHeartRecordRow) SendHeartTime() int64 {
	return int64(this.row.SendHeartTime)
}

func (this *PlayerSendHeartRecordRow) GoObject() *PlayerSendHeartRecord {
	return &PlayerSendHeartRecord{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		FriendPid: int64(this.row.FriendPid),
		SendHeartTime: int64(this.row.SendHeartTime),
	}
}

/* ========== player_skill ========== */

// 玩家角色绝招表
type PlayerSkillRow struct {
	row *C.PlayerSkill
	isBreak bool
}

func (this *PlayerSkillRow) reset(row *C.PlayerSkill) {	this.row = row
	this.isBreak = false
}

func (this *PlayerSkillRow) Break() {
	this.isBreak = true
}

// 主键ID
func (this *PlayerSkillRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerSkillRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 角色ID
func (this *PlayerSkillRow) RoleId() int8 {
	return int8(this.row.RoleId)
}

// 绝招ID
func (this *PlayerSkillRow) SkillId() int16 {
	return int16(this.row.SkillId)
}

// 技能等级
func (this *PlayerSkillRow) SkillTrnlv() int32 {
	return int32(this.row.SkillTrnlv)
}

func (this *PlayerSkillRow) GoObject() *PlayerSkill {
	return &PlayerSkill{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		RoleId: int8(this.row.RoleId),
		SkillId: int16(this.row.SkillId),
		SkillTrnlv: int32(this.row.SkillTrnlv),
	}
}

/* ========== player_state ========== */

// 玩家状态
type PlayerStateRow struct {
	row *C.PlayerState
	isBreak bool
}

func (this *PlayerStateRow) reset(row *C.PlayerState) {	this.row = row
	this.isBreak = false
}

func (this *PlayerStateRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerStateRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 被冻结的时间
func (this *PlayerStateRow) BanStartTime() int64 {
	return int64(this.row.BanStartTime)
}

// 被冻结时长, -1 无冻结；0 永久
func (this *PlayerStateRow) BanEndTime() int64 {
	return int64(this.row.BanEndTime)
}

func (this *PlayerStateRow) GoObject() *PlayerState {
	return &PlayerState{
		Pid: int64(this.row.Pid),
		BanStartTime: int64(this.row.BanStartTime),
		BanEndTime: int64(this.row.BanEndTime),
	}
}

/* ========== player_sword_soul ========== */

// 玩家剑心数据
type PlayerSwordSoulRow struct {
	row *C.PlayerSwordSoul
	isBreak bool
}

func (this *PlayerSwordSoulRow) reset(row *C.PlayerSwordSoul) {	this.row = row
	this.isBreak = false
}

func (this *PlayerSwordSoulRow) Break() {
	this.isBreak = true
}

// 玩家物品ID
func (this *PlayerSwordSoulRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerSwordSoulRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 是否已装备  1-已装备 0-未装备
func (this *PlayerSwordSoulRow) Pos() int8 {
	return int8(this.row.Pos)
}

// 剑心ID
func (this *PlayerSwordSoulRow) SwordSoulId() int16 {
	return int16(this.row.SwordSoulId)
}

// 当前经验
func (this *PlayerSwordSoulRow) Exp() int32 {
	return int32(this.row.Exp)
}

// 等级
func (this *PlayerSwordSoulRow) Level() int8 {
	return int8(this.row.Level)
}

func (this *PlayerSwordSoulRow) GoObject() *PlayerSwordSoul {
	return &PlayerSwordSoul{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		Pos: int8(this.row.Pos),
		SwordSoulId: int16(this.row.SwordSoulId),
		Exp: int32(this.row.Exp),
		Level: int8(this.row.Level),
	}
}

/* ========== player_sword_soul_equipment ========== */

// 玩家剑心装备表
type PlayerSwordSoulEquipmentRow struct {
	row *C.PlayerSwordSoulEquipment
	isBreak bool
}

func (this *PlayerSwordSoulEquipmentRow) reset(row *C.PlayerSwordSoulEquipment) {	this.row = row
	this.isBreak = false
}

func (this *PlayerSwordSoulEquipmentRow) Break() {
	this.isBreak = true
}

// 主键ID
func (this *PlayerSwordSoulEquipmentRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerSwordSoulEquipmentRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 角色ID
func (this *PlayerSwordSoulEquipmentRow) RoleId() int8 {
	return int8(this.row.RoleId)
}

// 
func (this *PlayerSwordSoulEquipmentRow) IsEquippedXuanyuan() int8 {
	return int8(this.row.IsEquippedXuanyuan)
}

// 所有装备剑心类型的位标记
func (this *PlayerSwordSoulEquipmentRow) TypeBits() int64 {
	return int64(this.row.TypeBits)
}

// 装备位置1的剑心
func (this *PlayerSwordSoulEquipmentRow) Pos1() int64 {
	return int64(this.row.Pos1)
}

// 装备位置2的剑心
func (this *PlayerSwordSoulEquipmentRow) Pos2() int64 {
	return int64(this.row.Pos2)
}

// 装备位置3的剑心
func (this *PlayerSwordSoulEquipmentRow) Pos3() int64 {
	return int64(this.row.Pos3)
}

// 装备位置4的剑心
func (this *PlayerSwordSoulEquipmentRow) Pos4() int64 {
	return int64(this.row.Pos4)
}

// 装备位置5的剑心
func (this *PlayerSwordSoulEquipmentRow) Pos5() int64 {
	return int64(this.row.Pos5)
}

// 装备位置6的剑心
func (this *PlayerSwordSoulEquipmentRow) Pos6() int64 {
	return int64(this.row.Pos6)
}

// 装备位置7的剑心
func (this *PlayerSwordSoulEquipmentRow) Pos7() int64 {
	return int64(this.row.Pos7)
}

// 装备位置8的剑心
func (this *PlayerSwordSoulEquipmentRow) Pos8() int64 {
	return int64(this.row.Pos8)
}

// 装备位置9的剑心
func (this *PlayerSwordSoulEquipmentRow) Pos9() int64 {
	return int64(this.row.Pos9)
}

func (this *PlayerSwordSoulEquipmentRow) GoObject() *PlayerSwordSoulEquipment {
	return &PlayerSwordSoulEquipment{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		RoleId: int8(this.row.RoleId),
		IsEquippedXuanyuan: int8(this.row.IsEquippedXuanyuan),
		TypeBits: int64(this.row.TypeBits),
		Pos1: int64(this.row.Pos1),
		Pos2: int64(this.row.Pos2),
		Pos3: int64(this.row.Pos3),
		Pos4: int64(this.row.Pos4),
		Pos5: int64(this.row.Pos5),
		Pos6: int64(this.row.Pos6),
		Pos7: int64(this.row.Pos7),
		Pos8: int64(this.row.Pos8),
		Pos9: int64(this.row.Pos9),
	}
}

/* ========== player_sword_soul_state ========== */

// 玩家拔剑状态
type PlayerSwordSoulStateRow struct {
	row *C.PlayerSwordSoulState
	isBreak bool
}

func (this *PlayerSwordSoulStateRow) reset(row *C.PlayerSwordSoulState) {	this.row = row
	this.isBreak = false
}

func (this *PlayerSwordSoulStateRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerSwordSoulStateRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 开箱子的状态(位操作)
func (this *PlayerSwordSoulStateRow) BoxState() int8 {
	return int8(this.row.BoxState)
}

// 当前可拔剑次数
func (this *PlayerSwordSoulStateRow) Num() int16 {
	return int16(this.row.Num)
}

// 更新时间
func (this *PlayerSwordSoulStateRow) UpdateTime() int64 {
	return int64(this.row.UpdateTime)
}

// 
func (this *PlayerSwordSoulStateRow) IngotNum() int64 {
	return int64(this.row.IngotNum)
}

// 
func (this *PlayerSwordSoulStateRow) SupersoulAdditionalPossible() int8 {
	return int8(this.row.SupersoulAdditionalPossible)
}

// 
func (this *PlayerSwordSoulStateRow) IngotYelloOne() int16 {
	return int16(this.row.IngotYelloOne)
}

// 
func (this *PlayerSwordSoulStateRow) LastIngotDrawTime() int64 {
	return int64(this.row.LastIngotDrawTime)
}

// 玩家剑心值
func (this *PlayerSwordSoulStateRow) SwordSoulNum() int32 {
	return int32(this.row.SwordSoulNum)
}

func (this *PlayerSwordSoulStateRow) GoObject() *PlayerSwordSoulState {
	return &PlayerSwordSoulState{
		Pid: int64(this.row.Pid),
		BoxState: int8(this.row.BoxState),
		Num: int16(this.row.Num),
		UpdateTime: int64(this.row.UpdateTime),
		IngotNum: int64(this.row.IngotNum),
		SupersoulAdditionalPossible: int8(this.row.SupersoulAdditionalPossible),
		IngotYelloOne: int16(this.row.IngotYelloOne),
		LastIngotDrawTime: int64(this.row.LastIngotDrawTime),
		SwordSoulNum: int32(this.row.SwordSoulNum),
	}
}

/* ========== player_taoyuan_bless ========== */

// 玩家桃源祈福
type PlayerTaoyuanBlessRow struct {
	row *C.PlayerTaoyuanBless
	isBreak bool
}

func (this *PlayerTaoyuanBlessRow) reset(row *C.PlayerTaoyuanBless) {	this.row = row
	this.isBreak = false
}

func (this *PlayerTaoyuanBlessRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerTaoyuanBlessRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 每日祈福次数
func (this *PlayerTaoyuanBlessRow) DailyBlessTimes() int32 {
	return int32(this.row.DailyBlessTimes)
}

// 最后一次祈福时间
func (this *PlayerTaoyuanBlessRow) LastBlessTime() int64 {
	return int64(this.row.LastBlessTime)
}

// 每日被祈福次数
func (this *PlayerTaoyuanBlessRow) DailyBeBlessTimes() int32 {
	return int32(this.row.DailyBeBlessTimes)
}

// 最后一次被祈福时间
func (this *PlayerTaoyuanBlessRow) LastBeBlessTime() int64 {
	return int64(this.row.LastBeBlessTime)
}

// 祈福玩家1
func (this *PlayerTaoyuanBlessRow) BlessPid1() int64 {
	return int64(this.row.BlessPid1)
}

// 祈福玩家2
func (this *PlayerTaoyuanBlessRow) BlessPid2() int64 {
	return int64(this.row.BlessPid2)
}

// 祈福玩家3
func (this *PlayerTaoyuanBlessRow) BlessPid3() int64 {
	return int64(this.row.BlessPid3)
}

// 祈福玩家4
func (this *PlayerTaoyuanBlessRow) BlessPid4() int64 {
	return int64(this.row.BlessPid4)
}

// 祈福玩家5
func (this *PlayerTaoyuanBlessRow) BlessPid5() int64 {
	return int64(this.row.BlessPid5)
}

func (this *PlayerTaoyuanBlessRow) GoObject() *PlayerTaoyuanBless {
	return &PlayerTaoyuanBless{
		Pid: int64(this.row.Pid),
		DailyBlessTimes: int32(this.row.DailyBlessTimes),
		LastBlessTime: int64(this.row.LastBlessTime),
		DailyBeBlessTimes: int32(this.row.DailyBeBlessTimes),
		LastBeBlessTime: int64(this.row.LastBeBlessTime),
		BlessPid1: int64(this.row.BlessPid1),
		BlessPid2: int64(this.row.BlessPid2),
		BlessPid3: int64(this.row.BlessPid3),
		BlessPid4: int64(this.row.BlessPid4),
		BlessPid5: int64(this.row.BlessPid5),
	}
}

/* ========== player_taoyuan_fileds ========== */

// 玩家桃源田地作物
type PlayerTaoyuanFiledsRow struct {
	row *C.PlayerTaoyuanFileds
	isBreak bool
}

func (this *PlayerTaoyuanFiledsRow) reset(row *C.PlayerTaoyuanFileds) {	this.row = row
	this.isBreak = false
}

func (this *PlayerTaoyuanFiledsRow) Break() {
	this.isBreak = true
}

// 
func (this *PlayerTaoyuanFiledsRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerTaoyuanFiledsRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 对应田地编号
func (this *PlayerTaoyuanFiledsRow) FiledId() int16 {
	return int16(this.row.FiledId)
}

// 对应田地状态0-普通,1-黑土地
func (this *PlayerTaoyuanFiledsRow) FiledStatus() int16 {
	return int16(this.row.FiledStatus)
}

// 种植植物ID
func (this *PlayerTaoyuanFiledsRow) PlantId() int16 {
	return int16(this.row.PlantId)
}

// 种植时间
func (this *PlayerTaoyuanFiledsRow) GrowTime() int64 {
	return int64(this.row.GrowTime)
}

func (this *PlayerTaoyuanFiledsRow) GoObject() *PlayerTaoyuanFileds {
	return &PlayerTaoyuanFileds{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		FiledId: int16(this.row.FiledId),
		FiledStatus: int16(this.row.FiledStatus),
		PlantId: int16(this.row.PlantId),
		GrowTime: int64(this.row.GrowTime),
	}
}

/* ========== player_taoyuan_heart ========== */

// 玩家桃源之心
type PlayerTaoyuanHeartRow struct {
	row *C.PlayerTaoyuanHeart
	isBreak bool
}

func (this *PlayerTaoyuanHeartRow) reset(row *C.PlayerTaoyuanHeart) {	this.row = row
	this.isBreak = false
}

func (this *PlayerTaoyuanHeartRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerTaoyuanHeartRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 桃源之心等级
func (this *PlayerTaoyuanHeartRow) Level() int16 {
	return int16(this.row.Level)
}

// 桃源之心经验
func (this *PlayerTaoyuanHeartRow) Exp() int64 {
	return int64(this.row.Exp)
}

func (this *PlayerTaoyuanHeartRow) GoObject() *PlayerTaoyuanHeart {
	return &PlayerTaoyuanHeart{
		Pid: int64(this.row.Pid),
		Level: int16(this.row.Level),
		Exp: int64(this.row.Exp),
	}
}

/* ========== player_taoyuan_item ========== */

// 玩家桃源物品
type PlayerTaoyuanItemRow struct {
	row *C.PlayerTaoyuanItem
	isBreak bool
}

func (this *PlayerTaoyuanItemRow) reset(row *C.PlayerTaoyuanItem) {	this.row = row
	this.isBreak = false
}

func (this *PlayerTaoyuanItemRow) Break() {
	this.isBreak = true
}

// 主键ID
func (this *PlayerTaoyuanItemRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerTaoyuanItemRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 物品ID
func (this *PlayerTaoyuanItemRow) ItemId() int16 {
	return int16(this.row.ItemId)
}

// 数量
func (this *PlayerTaoyuanItemRow) Num() int16 {
	return int16(this.row.Num)
}

func (this *PlayerTaoyuanItemRow) GoObject() *PlayerTaoyuanItem {
	return &PlayerTaoyuanItem{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		ItemId: int16(this.row.ItemId),
		Num: int16(this.row.Num),
	}
}

/* ========== player_taoyuan_message ========== */

// 祈福消息列表
type PlayerTaoyuanMessageRow struct {
	row *C.PlayerTaoyuanMessage
	isBreak bool
	cached_Nick bool
	cacheof_Nick string
}

func (this *PlayerTaoyuanMessageRow) reset(row *C.PlayerTaoyuanMessage) {	this.row = row
	this.isBreak = false
	this.cached_Nick = false
}

func (this *PlayerTaoyuanMessageRow) Break() {
	this.isBreak = true
}

// 主键ID
func (this *PlayerTaoyuanMessageRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerTaoyuanMessageRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 祈福玩家名
func (this *PlayerTaoyuanMessageRow) Nick() string {
	if !this.cached_Nick {
		this.cached_Nick = true
		this.cacheof_Nick = C.GoStringN(this.row.Nick, this.row.Nick_len)
	}
	return this.cacheof_Nick
}

// 纳福获得桃源之心经验
func (this *PlayerTaoyuanMessageRow) Exp() int32 {
	return int32(this.row.Exp)
}

func (this *PlayerTaoyuanMessageRow) GoObject() *PlayerTaoyuanMessage {
	return &PlayerTaoyuanMessage{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		Nick: this.Nick(),
		Exp: int32(this.row.Exp),
	}
}

/* ========== player_taoyuan_product ========== */

// 玩家桃源食坊、酒坊
type PlayerTaoyuanProductRow struct {
	row *C.PlayerTaoyuanProduct
	isBreak bool
}

func (this *PlayerTaoyuanProductRow) reset(row *C.PlayerTaoyuanProduct) {	this.row = row
	this.isBreak = false
}

func (this *PlayerTaoyuanProductRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerTaoyuanProductRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 桃源酒坊技能
func (this *PlayerTaoyuanProductRow) SkillId1() int16 {
	return int16(this.row.SkillId1)
}

// 桃源食坊技能
func (this *PlayerTaoyuanProductRow) SkillId2() int16 {
	return int16(this.row.SkillId2)
}

// 酿酒队列
func (this *PlayerTaoyuanProductRow) SameTimeWineQueue() int16 {
	return int16(this.row.SameTimeWineQueue)
}

// 酿酒次数
func (this *PlayerTaoyuanProductRow) MakeWineTimes() int64 {
	return int64(this.row.MakeWineTimes)
}

// 烹饪队列
func (this *PlayerTaoyuanProductRow) SameTimeFoodQueue() int16 {
	return int16(this.row.SameTimeFoodQueue)
}

// 烹饪次数
func (this *PlayerTaoyuanProductRow) MakeFoodTimes() int64 {
	return int64(this.row.MakeFoodTimes)
}

// 队列1美食id
func (this *PlayerTaoyuanProductRow) FoodQueue1() int16 {
	return int16(this.row.FoodQueue1)
}

// 队列1美食开始时间
func (this *PlayerTaoyuanProductRow) FoodQueue1StartTimestamp() int64 {
	return int64(this.row.FoodQueue1StartTimestamp)
}

// 队列1美食结束时间
func (this *PlayerTaoyuanProductRow) FoodQueue1EndTimestamp() int64 {
	return int64(this.row.FoodQueue1EndTimestamp)
}

// 队列2美食id
func (this *PlayerTaoyuanProductRow) FoodQueue2() int16 {
	return int16(this.row.FoodQueue2)
}

// 队列2美食开始时间
func (this *PlayerTaoyuanProductRow) FoodQueue2StartTimestamp() int64 {
	return int64(this.row.FoodQueue2StartTimestamp)
}

// 队列2美食结束时间
func (this *PlayerTaoyuanProductRow) FoodQueue2EndTimestamp() int64 {
	return int64(this.row.FoodQueue2EndTimestamp)
}

// 队列3美食id
func (this *PlayerTaoyuanProductRow) FoodQueue3() int16 {
	return int16(this.row.FoodQueue3)
}

// 队列3美食开始时间
func (this *PlayerTaoyuanProductRow) FoodQueue3StartTimestamp() int64 {
	return int64(this.row.FoodQueue3StartTimestamp)
}

// 队列3美食结束时间
func (this *PlayerTaoyuanProductRow) FoodQueue3EndTimestamp() int64 {
	return int64(this.row.FoodQueue3EndTimestamp)
}

// 队列4美食id
func (this *PlayerTaoyuanProductRow) FoodQueue4() int16 {
	return int16(this.row.FoodQueue4)
}

// 队列4美食开始时间
func (this *PlayerTaoyuanProductRow) FoodQueue4StartTimestamp() int64 {
	return int64(this.row.FoodQueue4StartTimestamp)
}

// 队列4美食结束时间
func (this *PlayerTaoyuanProductRow) FoodQueue4EndTimestamp() int64 {
	return int64(this.row.FoodQueue4EndTimestamp)
}

// 队列1美酒id
func (this *PlayerTaoyuanProductRow) WineQueue1() int16 {
	return int16(this.row.WineQueue1)
}

// 队列1美酒开始时间
func (this *PlayerTaoyuanProductRow) WineQueue1StartTimestamp() int64 {
	return int64(this.row.WineQueue1StartTimestamp)
}

// 队列1美酒结束时间
func (this *PlayerTaoyuanProductRow) WineQueue1EndTimestamp() int64 {
	return int64(this.row.WineQueue1EndTimestamp)
}

// 队列2美酒id
func (this *PlayerTaoyuanProductRow) WineQueue2() int16 {
	return int16(this.row.WineQueue2)
}

// 队列2美酒开始时间
func (this *PlayerTaoyuanProductRow) WineQueue2StartTimestamp() int64 {
	return int64(this.row.WineQueue2StartTimestamp)
}

// 队列2美酒结束时间
func (this *PlayerTaoyuanProductRow) WineQueue2EndTimestamp() int64 {
	return int64(this.row.WineQueue2EndTimestamp)
}

// 队列3美酒id
func (this *PlayerTaoyuanProductRow) WineQueue3() int16 {
	return int16(this.row.WineQueue3)
}

// 队列3美酒开始时间
func (this *PlayerTaoyuanProductRow) WineQueue3StartTimestamp() int64 {
	return int64(this.row.WineQueue3StartTimestamp)
}

// 队列3美酒结束时间
func (this *PlayerTaoyuanProductRow) WineQueue3EndTimestamp() int64 {
	return int64(this.row.WineQueue3EndTimestamp)
}

// 队列4美酒id
func (this *PlayerTaoyuanProductRow) WineQueue4() int16 {
	return int16(this.row.WineQueue4)
}

// 队列4美酒开始时间
func (this *PlayerTaoyuanProductRow) WineQueue4StartTimestamp() int64 {
	return int64(this.row.WineQueue4StartTimestamp)
}

// 队列4美酒结束时间
func (this *PlayerTaoyuanProductRow) WineQueue4EndTimestamp() int64 {
	return int64(this.row.WineQueue4EndTimestamp)
}

// 酒坊是否开启
func (this *PlayerTaoyuanProductRow) IsOpenWine() int8 {
	return int8(this.row.IsOpenWine)
}

// 食坊是否开启
func (this *PlayerTaoyuanProductRow) IsOpenFood() int8 {
	return int8(this.row.IsOpenFood)
}

func (this *PlayerTaoyuanProductRow) GoObject() *PlayerTaoyuanProduct {
	return &PlayerTaoyuanProduct{
		Pid: int64(this.row.Pid),
		SkillId1: int16(this.row.SkillId1),
		SkillId2: int16(this.row.SkillId2),
		SameTimeWineQueue: int16(this.row.SameTimeWineQueue),
		MakeWineTimes: int64(this.row.MakeWineTimes),
		SameTimeFoodQueue: int16(this.row.SameTimeFoodQueue),
		MakeFoodTimes: int64(this.row.MakeFoodTimes),
		FoodQueue1: int16(this.row.FoodQueue1),
		FoodQueue1StartTimestamp: int64(this.row.FoodQueue1StartTimestamp),
		FoodQueue1EndTimestamp: int64(this.row.FoodQueue1EndTimestamp),
		FoodQueue2: int16(this.row.FoodQueue2),
		FoodQueue2StartTimestamp: int64(this.row.FoodQueue2StartTimestamp),
		FoodQueue2EndTimestamp: int64(this.row.FoodQueue2EndTimestamp),
		FoodQueue3: int16(this.row.FoodQueue3),
		FoodQueue3StartTimestamp: int64(this.row.FoodQueue3StartTimestamp),
		FoodQueue3EndTimestamp: int64(this.row.FoodQueue3EndTimestamp),
		FoodQueue4: int16(this.row.FoodQueue4),
		FoodQueue4StartTimestamp: int64(this.row.FoodQueue4StartTimestamp),
		FoodQueue4EndTimestamp: int64(this.row.FoodQueue4EndTimestamp),
		WineQueue1: int16(this.row.WineQueue1),
		WineQueue1StartTimestamp: int64(this.row.WineQueue1StartTimestamp),
		WineQueue1EndTimestamp: int64(this.row.WineQueue1EndTimestamp),
		WineQueue2: int16(this.row.WineQueue2),
		WineQueue2StartTimestamp: int64(this.row.WineQueue2StartTimestamp),
		WineQueue2EndTimestamp: int64(this.row.WineQueue2EndTimestamp),
		WineQueue3: int16(this.row.WineQueue3),
		WineQueue3StartTimestamp: int64(this.row.WineQueue3StartTimestamp),
		WineQueue3EndTimestamp: int64(this.row.WineQueue3EndTimestamp),
		WineQueue4: int16(this.row.WineQueue4),
		WineQueue4StartTimestamp: int64(this.row.WineQueue4StartTimestamp),
		WineQueue4EndTimestamp: int64(this.row.WineQueue4EndTimestamp),
		IsOpenWine: int8(this.row.IsOpenWine),
		IsOpenFood: int8(this.row.IsOpenFood),
	}
}

/* ========== player_taoyuan_purchase_record ========== */

// 玩家桃源物品购买记录 仅记录限购商品
type PlayerTaoyuanPurchaseRecordRow struct {
	row *C.PlayerTaoyuanPurchaseRecord
	isBreak bool
}

func (this *PlayerTaoyuanPurchaseRecordRow) reset(row *C.PlayerTaoyuanPurchaseRecord) {	this.row = row
	this.isBreak = false
}

func (this *PlayerTaoyuanPurchaseRecordRow) Break() {
	this.isBreak = true
}

// 主键ID
func (this *PlayerTaoyuanPurchaseRecordRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerTaoyuanPurchaseRecordRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 物品ID
func (this *PlayerTaoyuanPurchaseRecordRow) ItemId() int16 {
	return int16(this.row.ItemId)
}

// 已购买数量
func (this *PlayerTaoyuanPurchaseRecordRow) Num() int16 {
	return int16(this.row.Num)
}

// 上次购买时间
func (this *PlayerTaoyuanPurchaseRecordRow) Timestamp() int64 {
	return int64(this.row.Timestamp)
}

func (this *PlayerTaoyuanPurchaseRecordRow) GoObject() *PlayerTaoyuanPurchaseRecord {
	return &PlayerTaoyuanPurchaseRecord{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		ItemId: int16(this.row.ItemId),
		Num: int16(this.row.Num),
		Timestamp: int64(this.row.Timestamp),
	}
}

/* ========== player_taoyuan_quest ========== */

// 玩家桃源愿望
type PlayerTaoyuanQuestRow struct {
	row *C.PlayerTaoyuanQuest
	isBreak bool
}

func (this *PlayerTaoyuanQuestRow) reset(row *C.PlayerTaoyuanQuest) {	this.row = row
	this.isBreak = false
}

func (this *PlayerTaoyuanQuestRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerTaoyuanQuestRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 最后一次桃源愿望刷新时间
func (this *PlayerTaoyuanQuestRow) LastQuestFlashTime() int64 {
	return int64(this.row.LastQuestFlashTime)
}

// 桃源愿望1物品ID
func (this *PlayerTaoyuanQuestRow) Quest1ItemId() int16 {
	return int16(this.row.Quest1ItemId)
}

// 桃源愿望1物品需求数量
func (this *PlayerTaoyuanQuestRow) Quest1ItemNum() int16 {
	return int16(this.row.Quest1ItemNum)
}

// 桃源愿望1获得总经验
func (this *PlayerTaoyuanQuestRow) Quest1TotalExp() int16 {
	return int16(this.row.Quest1TotalExp)
}

// 桃源愿望1获得总铜钱
func (this *PlayerTaoyuanQuestRow) Quest1TotalCoins() int64 {
	return int64(this.row.Quest1TotalCoins)
}

// 桃源愿望1完成时间
func (this *PlayerTaoyuanQuestRow) Quest1FinishTime() int64 {
	return int64(this.row.Quest1FinishTime)
}

// 桃源愿望2物品ID
func (this *PlayerTaoyuanQuestRow) Quest2ItemId() int16 {
	return int16(this.row.Quest2ItemId)
}

// 桃源愿望2物品需求数量
func (this *PlayerTaoyuanQuestRow) Quest2ItemNum() int16 {
	return int16(this.row.Quest2ItemNum)
}

// 桃源愿望2获得总经验
func (this *PlayerTaoyuanQuestRow) Quest2TotalExp() int16 {
	return int16(this.row.Quest2TotalExp)
}

// 桃源愿望2获得总铜钱
func (this *PlayerTaoyuanQuestRow) Quest2TotalCoins() int64 {
	return int64(this.row.Quest2TotalCoins)
}

// 桃源愿望2完成时间
func (this *PlayerTaoyuanQuestRow) Quest2FinishTime() int64 {
	return int64(this.row.Quest2FinishTime)
}

// 桃源愿望3物品ID
func (this *PlayerTaoyuanQuestRow) Quest3ItemId() int16 {
	return int16(this.row.Quest3ItemId)
}

// 桃源愿望3物品需求数量
func (this *PlayerTaoyuanQuestRow) Quest3ItemNum() int16 {
	return int16(this.row.Quest3ItemNum)
}

// 桃源愿望3获得总经验
func (this *PlayerTaoyuanQuestRow) Quest3TotalExp() int16 {
	return int16(this.row.Quest3TotalExp)
}

// 桃源愿望3获得总铜钱
func (this *PlayerTaoyuanQuestRow) Quest3TotalCoins() int64 {
	return int64(this.row.Quest3TotalCoins)
}

// 桃源愿望3完成时间
func (this *PlayerTaoyuanQuestRow) Quest3FinishTime() int64 {
	return int64(this.row.Quest3FinishTime)
}

// 桃源愿望4物品ID
func (this *PlayerTaoyuanQuestRow) Quest4ItemId() int16 {
	return int16(this.row.Quest4ItemId)
}

// 桃源愿望4物品需求数量
func (this *PlayerTaoyuanQuestRow) Quest4ItemNum() int16 {
	return int16(this.row.Quest4ItemNum)
}

// 桃源愿望4获得总经验
func (this *PlayerTaoyuanQuestRow) Quest4TotalExp() int16 {
	return int16(this.row.Quest4TotalExp)
}

// 桃源愿望4获得总铜钱
func (this *PlayerTaoyuanQuestRow) Quest4TotalCoins() int64 {
	return int64(this.row.Quest4TotalCoins)
}

// 桃源愿望4完成时间
func (this *PlayerTaoyuanQuestRow) Quest4FinishTime() int64 {
	return int64(this.row.Quest4FinishTime)
}

// 桃源愿望5物品ID
func (this *PlayerTaoyuanQuestRow) Quest5ItemId() int16 {
	return int16(this.row.Quest5ItemId)
}

// 桃源愿望5物品需求数量
func (this *PlayerTaoyuanQuestRow) Quest5ItemNum() int16 {
	return int16(this.row.Quest5ItemNum)
}

// 桃源愿望5获得总经验
func (this *PlayerTaoyuanQuestRow) Quest5TotalExp() int16 {
	return int16(this.row.Quest5TotalExp)
}

// 桃源愿望5获得总铜钱
func (this *PlayerTaoyuanQuestRow) Quest5TotalCoins() int64 {
	return int64(this.row.Quest5TotalCoins)
}

// 桃源愿望5完成时间
func (this *PlayerTaoyuanQuestRow) Quest5FinishTime() int64 {
	return int64(this.row.Quest5FinishTime)
}

func (this *PlayerTaoyuanQuestRow) GoObject() *PlayerTaoyuanQuest {
	return &PlayerTaoyuanQuest{
		Pid: int64(this.row.Pid),
		LastQuestFlashTime: int64(this.row.LastQuestFlashTime),
		Quest1ItemId: int16(this.row.Quest1ItemId),
		Quest1ItemNum: int16(this.row.Quest1ItemNum),
		Quest1TotalExp: int16(this.row.Quest1TotalExp),
		Quest1TotalCoins: int64(this.row.Quest1TotalCoins),
		Quest1FinishTime: int64(this.row.Quest1FinishTime),
		Quest2ItemId: int16(this.row.Quest2ItemId),
		Quest2ItemNum: int16(this.row.Quest2ItemNum),
		Quest2TotalExp: int16(this.row.Quest2TotalExp),
		Quest2TotalCoins: int64(this.row.Quest2TotalCoins),
		Quest2FinishTime: int64(this.row.Quest2FinishTime),
		Quest3ItemId: int16(this.row.Quest3ItemId),
		Quest3ItemNum: int16(this.row.Quest3ItemNum),
		Quest3TotalExp: int16(this.row.Quest3TotalExp),
		Quest3TotalCoins: int64(this.row.Quest3TotalCoins),
		Quest3FinishTime: int64(this.row.Quest3FinishTime),
		Quest4ItemId: int16(this.row.Quest4ItemId),
		Quest4ItemNum: int16(this.row.Quest4ItemNum),
		Quest4TotalExp: int16(this.row.Quest4TotalExp),
		Quest4TotalCoins: int64(this.row.Quest4TotalCoins),
		Quest4FinishTime: int64(this.row.Quest4FinishTime),
		Quest5ItemId: int16(this.row.Quest5ItemId),
		Quest5ItemNum: int16(this.row.Quest5ItemNum),
		Quest5TotalExp: int16(this.row.Quest5TotalExp),
		Quest5TotalCoins: int64(this.row.Quest5TotalCoins),
		Quest5FinishTime: int64(this.row.Quest5FinishTime),
	}
}

/* ========== player_tb_xxd_roleinfo ========== */

// 腾讯经分用户信息表
type PlayerTbXxdRoleinfoRow struct {
	row *C.PlayerTbXxdRoleinfo
	isBreak bool
	cached_Gameappid bool
	cacheof_Gameappid string
	cached_Openid bool
	cacheof_Openid string
}

func (this *PlayerTbXxdRoleinfoRow) reset(row *C.PlayerTbXxdRoleinfo) {	this.row = row
	this.isBreak = false
	this.cached_Gameappid = false
	this.cached_Openid = false
}

func (this *PlayerTbXxdRoleinfoRow) Break() {
	this.isBreak = true
}

// 玩家id
func (this *PlayerTbXxdRoleinfoRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 平台分配的AppID
func (this *PlayerTbXxdRoleinfoRow) Gameappid() string {
	if !this.cached_Gameappid {
		this.cached_Gameappid = true
		this.cacheof_Gameappid = C.GoStringN(this.row.Gameappid, this.row.Gameappid_len)
	}
	return this.cacheof_Gameappid
}

// 玩家平台唯一标识
func (this *PlayerTbXxdRoleinfoRow) Openid() string {
	if !this.cached_Openid {
		this.cached_Openid = true
		this.cacheof_Openid = C.GoStringN(this.row.Openid, this.row.Openid_len)
	}
	return this.cacheof_Openid
}

// 注册时间
func (this *PlayerTbXxdRoleinfoRow) Regtime() int64 {
	return int64(this.row.Regtime)
}

// 玩家等级
func (this *PlayerTbXxdRoleinfoRow) Level() int16 {
	return int16(this.row.Level)
}

// 玩家好友数
func (this *PlayerTbXxdRoleinfoRow) IFriends() int16 {
	return int16(this.row.IFriends)
}

// ios金钱存量
func (this *PlayerTbXxdRoleinfoRow) Moneyios() int64 {
	return int64(this.row.Moneyios)
}

// android金钱存量
func (this *PlayerTbXxdRoleinfoRow) Moneyandroid() int64 {
	return int64(this.row.Moneyandroid)
}

// ios钻石存量
func (this *PlayerTbXxdRoleinfoRow) Diamondios() int64 {
	return int64(this.row.Diamondios)
}

// android钻石存量
func (this *PlayerTbXxdRoleinfoRow) Diamondandroid() int64 {
	return int64(this.row.Diamondandroid)
}

func (this *PlayerTbXxdRoleinfoRow) GoObject() *PlayerTbXxdRoleinfo {
	return &PlayerTbXxdRoleinfo{
		Pid: int64(this.row.Pid),
		Gameappid: this.Gameappid(),
		Openid: this.Openid(),
		Regtime: int64(this.row.Regtime),
		Level: int16(this.row.Level),
		IFriends: int16(this.row.IFriends),
		Moneyios: int64(this.row.Moneyios),
		Moneyandroid: int64(this.row.Moneyandroid),
		Diamondios: int64(this.row.Diamondios),
		Diamondandroid: int64(this.row.Diamondandroid),
	}
}

/* ========== player_team_info ========== */

// 玩家队伍相关信息
type PlayerTeamInfoRow struct {
	row *C.PlayerTeamInfo
	isBreak bool
}

func (this *PlayerTeamInfoRow) reset(row *C.PlayerTeamInfo) {	this.row = row
	this.isBreak = false
}

func (this *PlayerTeamInfoRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerTeamInfoRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 友情点数
func (this *PlayerTeamInfoRow) Relationship() int32 {
	return int32(this.row.Relationship)
}

// 生命项等级
func (this *PlayerTeamInfoRow) HealthLv() int16 {
	return int16(this.row.HealthLv)
}

// 攻击项等级
func (this *PlayerTeamInfoRow) AttackLv() int16 {
	return int16(this.row.AttackLv)
}

// 防御项等级
func (this *PlayerTeamInfoRow) DefenceLv() int16 {
	return int16(this.row.DefenceLv)
}

func (this *PlayerTeamInfoRow) GoObject() *PlayerTeamInfo {
	return &PlayerTeamInfo{
		Pid: int64(this.row.Pid),
		Relationship: int32(this.row.Relationship),
		HealthLv: int16(this.row.HealthLv),
		AttackLv: int16(this.row.AttackLv),
		DefenceLv: int16(this.row.DefenceLv),
	}
}

/* ========== player_totem ========== */

// 玩家阵印表
type PlayerTotemRow struct {
	row *C.PlayerTotem
	isBreak bool
}

func (this *PlayerTotemRow) reset(row *C.PlayerTotem) {	this.row = row
	this.isBreak = false
}

func (this *PlayerTotemRow) Break() {
	this.isBreak = true
}

// 主键
func (this *PlayerTotemRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerTotemRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 阵印ID
func (this *PlayerTotemRow) TotemId() int16 {
	return int16(this.row.TotemId)
}

// 等级
func (this *PlayerTotemRow) Level() int8 {
	return int8(this.row.Level)
}

// 技能
func (this *PlayerTotemRow) SkillId() int16 {
	return int16(this.row.SkillId)
}

func (this *PlayerTotemRow) GoObject() *PlayerTotem {
	return &PlayerTotem{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		TotemId: int16(this.row.TotemId),
		Level: int8(this.row.Level),
		SkillId: int16(this.row.SkillId),
	}
}

/* ========== player_totem_info ========== */

// 玩家阵印装备表
type PlayerTotemInfoRow struct {
	row *C.PlayerTotemInfo
	isBreak bool
}

func (this *PlayerTotemInfoRow) reset(row *C.PlayerTotemInfo) {	this.row = row
	this.isBreak = false
}

func (this *PlayerTotemInfoRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerTotemInfoRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 每日元宝召唤次数
func (this *PlayerTotemInfoRow) IngotCallDailyNum() int16 {
	return int16(this.row.IngotCallDailyNum)
}

// 元宝召唤时间戳
func (this *PlayerTotemInfoRow) IngotCallTimestamp() int64 {
	return int64(this.row.IngotCallTimestamp)
}

// 石附文数量
func (this *PlayerTotemInfoRow) RockRuneNum() int32 {
	return int32(this.row.RockRuneNum)
}

// 玉附文数量
func (this *PlayerTotemInfoRow) JadeRuneNum() int32 {
	return int32(this.row.JadeRuneNum)
}

// 装备位置1的阵印id
func (this *PlayerTotemInfoRow) Pos1() int64 {
	return int64(this.row.Pos1)
}

// 装备位置2的阵印id
func (this *PlayerTotemInfoRow) Pos2() int64 {
	return int64(this.row.Pos2)
}

// 装备位置3的阵印id
func (this *PlayerTotemInfoRow) Pos3() int64 {
	return int64(this.row.Pos3)
}

// 装备位置4的阵印id
func (this *PlayerTotemInfoRow) Pos4() int64 {
	return int64(this.row.Pos4)
}

// 装备位置4的阵印id
func (this *PlayerTotemInfoRow) Pos5() int64 {
	return int64(this.row.Pos5)
}

// 玩家元宝抽取次数
func (this *PlayerTotemInfoRow) IngotDrawTimes() int16 {
	return int16(this.row.IngotDrawTimes)
}

func (this *PlayerTotemInfoRow) GoObject() *PlayerTotemInfo {
	return &PlayerTotemInfo{
		Pid: int64(this.row.Pid),
		IngotCallDailyNum: int16(this.row.IngotCallDailyNum),
		IngotCallTimestamp: int64(this.row.IngotCallTimestamp),
		RockRuneNum: int32(this.row.RockRuneNum),
		JadeRuneNum: int32(this.row.JadeRuneNum),
		Pos1: int64(this.row.Pos1),
		Pos2: int64(this.row.Pos2),
		Pos3: int64(this.row.Pos3),
		Pos4: int64(this.row.Pos4),
		Pos5: int64(this.row.Pos5),
		IngotDrawTimes: int16(this.row.IngotDrawTimes),
	}
}

/* ========== player_town ========== */

// 玩家城镇数据
type PlayerTownRow struct {
	row *C.PlayerTown
	isBreak bool
}

func (this *PlayerTownRow) reset(row *C.PlayerTown) {	this.row = row
	this.isBreak = false
}

func (this *PlayerTownRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerTownRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 当前玩家所处的城镇ID
func (this *PlayerTownRow) TownId() int16 {
	return int16(this.row.TownId)
}

// 当前拥有的城镇权值
func (this *PlayerTownRow) Lock() int32 {
	return int32(this.row.Lock)
}

// 当前城镇的X轴位置
func (this *PlayerTownRow) AtPosX() int16 {
	return int16(this.row.AtPosX)
}

// 当前城镇的y轴位置
func (this *PlayerTownRow) AtPosY() int16 {
	return int16(this.row.AtPosY)
}

func (this *PlayerTownRow) GoObject() *PlayerTown {
	return &PlayerTown{
		Pid: int64(this.row.Pid),
		TownId: int16(this.row.TownId),
		Lock: int32(this.row.Lock),
		AtPosX: int16(this.row.AtPosX),
		AtPosY: int16(this.row.AtPosY),
	}
}

/* ========== player_trader_refresh_state ========== */

// 玩家随机商店手动刷新次数状态
type PlayerTraderRefreshStateRow struct {
	row *C.PlayerTraderRefreshState
	isBreak bool
}

func (this *PlayerTraderRefreshStateRow) reset(row *C.PlayerTraderRefreshState) {	this.row = row
	this.isBreak = false
}

func (this *PlayerTraderRefreshStateRow) Break() {
	this.isBreak = true
}

// ID
func (this *PlayerTraderRefreshStateRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerTraderRefreshStateRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 最近一次*手动*刷新时间
func (this *PlayerTraderRefreshStateRow) LastUpdateTime() int64 {
	return int64(this.row.LastUpdateTime)
}

// 最近一次*自动*刷新时间
func (this *PlayerTraderRefreshStateRow) AutoUpdateTime() int64 {
	return int64(this.row.AutoUpdateTime)
}

// 随机商人ID
func (this *PlayerTraderRefreshStateRow) TraderId() int16 {
	return int16(this.row.TraderId)
}

// 已刷新次数
func (this *PlayerTraderRefreshStateRow) RefreshNum() int16 {
	return int16(this.row.RefreshNum)
}

func (this *PlayerTraderRefreshStateRow) GoObject() *PlayerTraderRefreshState {
	return &PlayerTraderRefreshState{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		LastUpdateTime: int64(this.row.LastUpdateTime),
		AutoUpdateTime: int64(this.row.AutoUpdateTime),
		TraderId: int16(this.row.TraderId),
		RefreshNum: int16(this.row.RefreshNum),
	}
}

/* ========== player_trader_store_state ========== */

// 玩家随机商店状态
type PlayerTraderStoreStateRow struct {
	row *C.PlayerTraderStoreState
	isBreak bool
}

func (this *PlayerTraderStoreStateRow) reset(row *C.PlayerTraderStoreState) {	this.row = row
	this.isBreak = false
}

func (this *PlayerTraderStoreStateRow) Break() {
	this.isBreak = true
}

// ID
func (this *PlayerTraderStoreStateRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerTraderStoreStateRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 随机商人ID
func (this *PlayerTraderStoreStateRow) TraderId() int16 {
	return int16(this.row.TraderId)
}

// 格子ID
func (this *PlayerTraderStoreStateRow) GridId() int32 {
	return int32(this.row.GridId)
}

// 物品ID
func (this *PlayerTraderStoreStateRow) ItemId() int16 {
	return int16(this.row.ItemId)
}

// 物品数量
func (this *PlayerTraderStoreStateRow) Num() int16 {
	return int16(this.row.Num)
}

// 价格
func (this *PlayerTraderStoreStateRow) Cost() int64 {
	return int64(this.row.Cost)
}

// 剩余可购买次数
func (this *PlayerTraderStoreStateRow) Stock() int8 {
	return int8(this.row.Stock)
}

// 货物类型0-物品 1-爱心 2-剑心 3-魂侍
func (this *PlayerTraderStoreStateRow) GoodsType() int8 {
	return int8(this.row.GoodsType)
}

func (this *PlayerTraderStoreStateRow) GoObject() *PlayerTraderStoreState {
	return &PlayerTraderStoreState{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		TraderId: int16(this.row.TraderId),
		GridId: int32(this.row.GridId),
		ItemId: int16(this.row.ItemId),
		Num: int16(this.row.Num),
		Cost: int64(this.row.Cost),
		Stock: int8(this.row.Stock),
		GoodsType: int8(this.row.GoodsType),
	}
}

/* ========== player_use_skill ========== */

// 玩家角色当前使用的绝招表
type PlayerUseSkillRow struct {
	row *C.PlayerUseSkill
	isBreak bool
}

func (this *PlayerUseSkillRow) reset(row *C.PlayerUseSkill) {	this.row = row
	this.isBreak = false
}

func (this *PlayerUseSkillRow) Break() {
	this.isBreak = true
}

// 主键ID
func (this *PlayerUseSkillRow) Id() int64 {
	return int64(this.row.Id)
}

// 玩家ID
func (this *PlayerUseSkillRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 角色ID
func (this *PlayerUseSkillRow) RoleId() int8 {
	return int8(this.row.RoleId)
}

// 招式1
func (this *PlayerUseSkillRow) SkillId1() int16 {
	return int16(this.row.SkillId1)
}

// 招式4
func (this *PlayerUseSkillRow) SkillId4() int16 {
	return int16(this.row.SkillId4)
}

// 招式2
func (this *PlayerUseSkillRow) SkillId2() int16 {
	return int16(this.row.SkillId2)
}

// 招式3
func (this *PlayerUseSkillRow) SkillId3() int16 {
	return int16(this.row.SkillId3)
}

func (this *PlayerUseSkillRow) GoObject() *PlayerUseSkill {
	return &PlayerUseSkill{
		Id: int64(this.row.Id),
		Pid: int64(this.row.Pid),
		RoleId: int8(this.row.RoleId),
		SkillId1: int16(this.row.SkillId1),
		SkillId4: int16(this.row.SkillId4),
		SkillId2: int16(this.row.SkillId2),
		SkillId3: int16(this.row.SkillId3),
	}
}

/* ========== player_vip ========== */

// 玩家VIP卡信息
type PlayerVipRow struct {
	row *C.PlayerVip
	isBreak bool
	cached_CardId bool
	cacheof_CardId string
}

func (this *PlayerVipRow) reset(row *C.PlayerVip) {	this.row = row
	this.isBreak = false
	this.cached_CardId = false
}

func (this *PlayerVipRow) Break() {
	this.isBreak = true
}

// 玩家ID
func (this *PlayerVipRow) Pid() int64 {
	return int64(this.row.Pid)
}

// 累计充值元宝数
func (this *PlayerVipRow) Ingot() int64 {
	return int64(this.row.Ingot)
}

// VIP等级
func (this *PlayerVipRow) Level() int16 {
	return int16(this.row.Level)
}

// VIP卡编号
func (this *PlayerVipRow) CardId() string {
	if !this.cached_CardId {
		this.cached_CardId = true
		this.cacheof_CardId = C.GoStringN(this.row.CardId, this.row.CardId_len)
	}
	return this.cacheof_CardId
}

// 赠送vip经验
func (this *PlayerVipRow) PresentExp() int64 {
	return int64(this.row.PresentExp)
}

func (this *PlayerVipRow) GoObject() *PlayerVip {
	return &PlayerVip{
		Pid: int64(this.row.Pid),
		Ingot: int64(this.row.Ingot),
		Level: int16(this.row.Level),
		CardId: this.CardId(),
		PresentExp: int64(this.row.PresentExp),
	}
}

