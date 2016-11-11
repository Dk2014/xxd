package mail_dat

import (
	"strconv"
	"bytes"
	
)

const( 
	HAS_BEEN_READ = 1
	UN_READ = 0
	HAVE_ATTACHMENT = 1
	NO_ATTACHMENT = 0
	NO_ATTACHMENT_SAVE_DAYS = 3
	HAVE_ATTACHMENT_SAVE_DAYS = 7
	CLIQUE_HAVE_ATTACHMENT_SAVE_DAYS = 14
	HEART_MAIL_SAVE_DAYS = 3
	HEART_MAIL_ID = 2
	CLIQUE_SALARY_OWNER_MAIL_ID = 37
	CLIQUE_SALARY_MANAGER_MAIL_ID = 38
	CLIQUE_SALARY_MAIL_OWNER_COIN_BASE = 500
	CLIQUE_SALARY_MAIL_MANAGER_COIN_BASE = 100
)
const( 
	ATTACHMENT_ITEM = 0
	ATTACHMENT_COINS = 1
	ATTACHMENT_INGOT = 2
	ATTACHMENT_HEART = 3
	ATTACHMENT_FORMATION_EXP = 4
	ATTACHMENT_SINGLE_ROLE_EXP = 5
	ATTACHMENT_GHOST = 6
	ATTACHMENT_SWORD_SOUL = 7
	ATTACHMENT_HEART_FROM_FRIEND = 8
	ATTACHMENT_RELATIONSHIP = 9
	ATTACHMENT_FAME = 10
	ATTACHMENT_TOTEM = 11
)
const( 
	AUTO_DELETE_MAGIC_NUM = 10
	AUTO_DELETE_AFTER_EXPIRED = 0
	AUTO_DELETE_AFTER_READ_WITHOUT_ATTACHMENT = 1
)






var g_attachments =  map[int16][]*Attachment {
	2: []*Attachment{&Attachment{8,0,1},},
	3: []*Attachment{&Attachment{0,31,1},&Attachment{0,32,1},},
	12: []*Attachment{&Attachment{3,0,5},},
	13: []*Attachment{&Attachment{0,418,1},&Attachment{0,423,10},},
	28: []*Attachment{&Attachment{2,0,18},&Attachment{3,0,2},&Attachment{0,306,1},},
	29: []*Attachment{&Attachment{0,432,10},&Attachment{0,510,1},&Attachment{3,0,10},},
	30: []*Attachment{&Attachment{0,418,1},&Attachment{0,423,10},},
	32: []*Attachment{&Attachment{2,0,100},},
	33: []*Attachment{&Attachment{2,0,150},},
	34: []*Attachment{&Attachment{1,0,50000},},
	35: []*Attachment{&Attachment{1,0,80000},},
	48: []*Attachment{&Attachment{0,61,1},&Attachment{0,303,5},&Attachment{0,259,10},&Attachment{0,306,5},&Attachment{2,0,200},},
}
// 背包已满提示
type MailBagFull struct {
	Func string // 功能
	Attachments []*Attachment
}

func (this MailBagFull) GetMailId() int32 {
	return int32(1)
}


func (this MailBagFull) GetSendTime() int64 {
	return 0
}

func (this MailBagFull) GetExpireTime() int64 {
	return 1
}
func (this MailBagFull) GetTitle() string {
	return ""
}
func (this MailBagFull) GetPriority() int8 {
	return 0
}
func (this MailBagFull) GetContent() string {
	return ""
}
func (this MailBagFull) GetMinLevel() int16 {
	return 0
}
func (this MailBagFull) GetMaxLevel() int16 {
	return 0
}
func (this MailBagFull) GetMinVIPLevel() int16 {
	return 0
}
func (this MailBagFull) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailBagFull) GetMinCreateTime() int64 {
	return 0
}
func (this MailBagFull) GetMaxCreateTime() int64 {
	return 0
}
func (this MailBagFull) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Func))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailBagFull) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[1]; ok {
		return attachments
	}
	return this.Attachments
}

// 爱心赠送邮件
type MailHeart struct {
	Who string // 发送者
	Attachments []*Attachment
}

func (this MailHeart) GetMailId() int32 {
	return int32(2)
}


func (this MailHeart) GetSendTime() int64 {
	return 0
}

func (this MailHeart) GetExpireTime() int64 {
	return 1
}
func (this MailHeart) GetTitle() string {
	return ""
}
func (this MailHeart) GetPriority() int8 {
	return 0
}
func (this MailHeart) GetContent() string {
	return ""
}
func (this MailHeart) GetMinLevel() int16 {
	return 0
}
func (this MailHeart) GetMaxLevel() int16 {
	return 0
}
func (this MailHeart) GetMinVIPLevel() int16 {
	return 0
}
func (this MailHeart) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailHeart) GetMinCreateTime() int64 {
	return 0
}
func (this MailHeart) GetMaxCreateTime() int64 {
	return 0
}
func (this MailHeart) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Who))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailHeart) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[2]; ok {
		return attachments
	}
	return this.Attachments
}

// 测试邮件
type MailTestMail struct {
	P1 string //  参数1
	P2 string //  参数2
	Attachments []*Attachment
}

func (this MailTestMail) GetMailId() int32 {
	return int32(3)
}


func (this MailTestMail) GetSendTime() int64 {
	return 0
}

func (this MailTestMail) GetExpireTime() int64 {
	return 1
}
func (this MailTestMail) GetTitle() string {
	return ""
}
func (this MailTestMail) GetPriority() int8 {
	return 0
}
func (this MailTestMail) GetContent() string {
	return ""
}
func (this MailTestMail) GetMinLevel() int16 {
	return 0
}
func (this MailTestMail) GetMaxLevel() int16 {
	return 0
}
func (this MailTestMail) GetMinVIPLevel() int16 {
	return 0
}
func (this MailTestMail) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailTestMail) GetMinCreateTime() int64 {
	return 0
}
func (this MailTestMail) GetMaxCreateTime() int64 {
	return 0
}
func (this MailTestMail) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.P1))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.P2))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailTestMail) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[3]; ok {
		return attachments
	}
	return this.Attachments
}

// 多人关卡战斗奖励
type MailMultiLevel struct {
	Name string // 关卡名称
	Attachments []*Attachment
}

func (this MailMultiLevel) GetMailId() int32 {
	return int32(4)
}


func (this MailMultiLevel) GetSendTime() int64 {
	return 0
}

func (this MailMultiLevel) GetExpireTime() int64 {
	return 1
}
func (this MailMultiLevel) GetTitle() string {
	return ""
}
func (this MailMultiLevel) GetPriority() int8 {
	return 0
}
func (this MailMultiLevel) GetContent() string {
	return ""
}
func (this MailMultiLevel) GetMinLevel() int16 {
	return 0
}
func (this MailMultiLevel) GetMaxLevel() int16 {
	return 0
}
func (this MailMultiLevel) GetMinVIPLevel() int16 {
	return 0
}
func (this MailMultiLevel) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailMultiLevel) GetMinCreateTime() int64 {
	return 0
}
func (this MailMultiLevel) GetMaxCreateTime() int64 {
	return 0
}
func (this MailMultiLevel) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Name))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailMultiLevel) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[4]; ok {
		return attachments
	}
	return this.Attachments
}

// 魂侍背包已满提示
type MailGhostBagFull struct {
	Func string // 功能
	Attachments []*Attachment
}

func (this MailGhostBagFull) GetMailId() int32 {
	return int32(6)
}


func (this MailGhostBagFull) GetSendTime() int64 {
	return 0
}

func (this MailGhostBagFull) GetExpireTime() int64 {
	return 0
}
func (this MailGhostBagFull) GetTitle() string {
	return ""
}
func (this MailGhostBagFull) GetPriority() int8 {
	return 0
}
func (this MailGhostBagFull) GetContent() string {
	return ""
}
func (this MailGhostBagFull) GetMinLevel() int16 {
	return 0
}
func (this MailGhostBagFull) GetMaxLevel() int16 {
	return 0
}
func (this MailGhostBagFull) GetMinVIPLevel() int16 {
	return 0
}
func (this MailGhostBagFull) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailGhostBagFull) GetMinCreateTime() int64 {
	return 0
}
func (this MailGhostBagFull) GetMaxCreateTime() int64 {
	return 0
}
func (this MailGhostBagFull) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Func))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailGhostBagFull) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[6]; ok {
		return attachments
	}
	return this.Attachments
}

// 剑心背包已满提示
type MailSwordSoulBagFull struct {
	Func string // 功能
	Attachments []*Attachment
}

func (this MailSwordSoulBagFull) GetMailId() int32 {
	return int32(7)
}


func (this MailSwordSoulBagFull) GetSendTime() int64 {
	return 0
}

func (this MailSwordSoulBagFull) GetExpireTime() int64 {
	return 1
}
func (this MailSwordSoulBagFull) GetTitle() string {
	return ""
}
func (this MailSwordSoulBagFull) GetPriority() int8 {
	return 0
}
func (this MailSwordSoulBagFull) GetContent() string {
	return ""
}
func (this MailSwordSoulBagFull) GetMinLevel() int16 {
	return 0
}
func (this MailSwordSoulBagFull) GetMaxLevel() int16 {
	return 0
}
func (this MailSwordSoulBagFull) GetMinVIPLevel() int16 {
	return 0
}
func (this MailSwordSoulBagFull) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailSwordSoulBagFull) GetMinCreateTime() int64 {
	return 0
}
func (this MailSwordSoulBagFull) GetMaxCreateTime() int64 {
	return 0
}
func (this MailSwordSoulBagFull) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Func))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailSwordSoulBagFull) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[7]; ok {
		return attachments
	}
	return this.Attachments
}

// 充值成功
type MailRecharge struct {
	Time1 string // 时间
	Num string // 充值元宝数量
	Attachments []*Attachment
}

func (this MailRecharge) GetMailId() int32 {
	return int32(8)
}


func (this MailRecharge) GetSendTime() int64 {
	return 0
}

func (this MailRecharge) GetExpireTime() int64 {
	return 0
}
func (this MailRecharge) GetTitle() string {
	return ""
}
func (this MailRecharge) GetPriority() int8 {
	return 0
}
func (this MailRecharge) GetContent() string {
	return ""
}
func (this MailRecharge) GetMinLevel() int16 {
	return 0
}
func (this MailRecharge) GetMaxLevel() int16 {
	return 0
}
func (this MailRecharge) GetMinVIPLevel() int16 {
	return 0
}
func (this MailRecharge) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailRecharge) GetMinCreateTime() int64 {
	return 0
}
func (this MailRecharge) GetMaxCreateTime() int64 {
	return 0
}
func (this MailRecharge) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Time1))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Num))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailRecharge) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[8]; ok {
		return attachments
	}
	return this.Attachments
}

// 购买道具成功提示
type MailPurchaseTips struct {
	Source string // 来源
	ItemName string // 道具名称
	Func string // 功能
	Attachments []*Attachment
}

func (this MailPurchaseTips) GetMailId() int32 {
	return int32(9)
}


func (this MailPurchaseTips) GetSendTime() int64 {
	return 0
}

func (this MailPurchaseTips) GetExpireTime() int64 {
	return 0
}
func (this MailPurchaseTips) GetTitle() string {
	return ""
}
func (this MailPurchaseTips) GetPriority() int8 {
	return 0
}
func (this MailPurchaseTips) GetContent() string {
	return ""
}
func (this MailPurchaseTips) GetMinLevel() int16 {
	return 0
}
func (this MailPurchaseTips) GetMaxLevel() int16 {
	return 0
}
func (this MailPurchaseTips) GetMinVIPLevel() int16 {
	return 0
}
func (this MailPurchaseTips) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailPurchaseTips) GetMinCreateTime() int64 {
	return 0
}
func (this MailPurchaseTips) GetMaxCreateTime() int64 {
	return 0
}
func (this MailPurchaseTips) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Source))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.ItemName))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Func))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailPurchaseTips) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[9]; ok {
		return attachments
	}
	return this.Attachments
}

// 神龙宝箱获得提示
type MailDrawTips struct {
	ItemName string // 道具名称
	Func string // 功能
	Attachments []*Attachment
}

func (this MailDrawTips) GetMailId() int32 {
	return int32(10)
}


func (this MailDrawTips) GetSendTime() int64 {
	return 0
}

func (this MailDrawTips) GetExpireTime() int64 {
	return 0
}
func (this MailDrawTips) GetTitle() string {
	return ""
}
func (this MailDrawTips) GetPriority() int8 {
	return 0
}
func (this MailDrawTips) GetContent() string {
	return ""
}
func (this MailDrawTips) GetMinLevel() int16 {
	return 0
}
func (this MailDrawTips) GetMaxLevel() int16 {
	return 0
}
func (this MailDrawTips) GetMinVIPLevel() int16 {
	return 0
}
func (this MailDrawTips) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailDrawTips) GetMinCreateTime() int64 {
	return 0
}
func (this MailDrawTips) GetMaxCreateTime() int64 {
	return 0
}
func (this MailDrawTips) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.ItemName))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Func))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailDrawTips) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[10]; ok {
		return attachments
	}
	return this.Attachments
}

// 邀请好友奖励
type MailPlatformFriendAward struct {
	FriendNum string // 人数
	ItemNum string // 数字
	ItemName string // 道具
	Attachments []*Attachment
}

func (this MailPlatformFriendAward) GetMailId() int32 {
	return int32(11)
}


func (this MailPlatformFriendAward) GetSendTime() int64 {
	return 0
}

func (this MailPlatformFriendAward) GetExpireTime() int64 {
	return 1
}
func (this MailPlatformFriendAward) GetTitle() string {
	return ""
}
func (this MailPlatformFriendAward) GetPriority() int8 {
	return 0
}
func (this MailPlatformFriendAward) GetContent() string {
	return ""
}
func (this MailPlatformFriendAward) GetMinLevel() int16 {
	return 0
}
func (this MailPlatformFriendAward) GetMaxLevel() int16 {
	return 0
}
func (this MailPlatformFriendAward) GetMinVIPLevel() int16 {
	return 0
}
func (this MailPlatformFriendAward) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailPlatformFriendAward) GetMinCreateTime() int64 {
	return 0
}
func (this MailPlatformFriendAward) GetMaxCreateTime() int64 {
	return 0
}
func (this MailPlatformFriendAward) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.FriendNum))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.ItemNum))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.ItemName))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailPlatformFriendAward) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[11]; ok {
		return attachments
	}
	return this.Attachments
}

// 感谢参与封测
type MailWelcomeBeta struct {
	Attachments []*Attachment
}

func (this MailWelcomeBeta) GetMailId() int32 {
	return int32(12)
}


func (this MailWelcomeBeta) GetSendTime() int64 {
	return 0
}

func (this MailWelcomeBeta) GetExpireTime() int64 {
	return 0
}
func (this MailWelcomeBeta) GetTitle() string {
	return ""
}
func (this MailWelcomeBeta) GetPriority() int8 {
	return 1
}
func (this MailWelcomeBeta) GetContent() string {
	return ""
}
func (this MailWelcomeBeta) GetMinLevel() int16 {
	return 0
}
func (this MailWelcomeBeta) GetMaxLevel() int16 {
	return 0
}
func (this MailWelcomeBeta) GetMinVIPLevel() int16 {
	return 0
}
func (this MailWelcomeBeta) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailWelcomeBeta) GetMinCreateTime() int64 {
	return 0
}
func (this MailWelcomeBeta) GetMaxCreateTime() int64 {
	return 0
}
func (this MailWelcomeBeta) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailWelcomeBeta) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[12]; ok {
		return attachments
	}
	return this.Attachments
}

// 比武场宝箱领取提醒
type MailWarTips struct {
	Nun string // 排名
	Attachments []*Attachment
}

func (this MailWarTips) GetMailId() int32 {
	return int32(13)
}


func (this MailWarTips) GetSendTime() int64 {
	return 0
}

func (this MailWarTips) GetExpireTime() int64 {
	return 1
}
func (this MailWarTips) GetTitle() string {
	return ""
}
func (this MailWarTips) GetPriority() int8 {
	return 0
}
func (this MailWarTips) GetContent() string {
	return ""
}
func (this MailWarTips) GetMinLevel() int16 {
	return 0
}
func (this MailWarTips) GetMaxLevel() int16 {
	return 0
}
func (this MailWarTips) GetMinVIPLevel() int16 {
	return 0
}
func (this MailWarTips) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailWarTips) GetMinCreateTime() int64 {
	return 0
}
func (this MailWarTips) GetMaxCreateTime() int64 {
	return 0
}
func (this MailWarTips) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Nun))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailWarTips) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[13]; ok {
		return attachments
	}
	return this.Attachments
}

// 仙尊每日爱心
type MailVIPHeart struct {
	Attachments []*Attachment
}

func (this MailVIPHeart) GetMailId() int32 {
	return int32(14)
}


func (this MailVIPHeart) GetSendTime() int64 {
	return 0
}

func (this MailVIPHeart) GetExpireTime() int64 {
	return 1
}
func (this MailVIPHeart) GetTitle() string {
	return ""
}
func (this MailVIPHeart) GetPriority() int8 {
	return 0
}
func (this MailVIPHeart) GetContent() string {
	return ""
}
func (this MailVIPHeart) GetMinLevel() int16 {
	return 0
}
func (this MailVIPHeart) GetMaxLevel() int16 {
	return 0
}
func (this MailVIPHeart) GetMinVIPLevel() int16 {
	return 0
}
func (this MailVIPHeart) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailVIPHeart) GetMinCreateTime() int64 {
	return 0
}
func (this MailVIPHeart) GetMaxCreateTime() int64 {
	return 0
}
func (this MailVIPHeart) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailVIPHeart) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[14]; ok {
		return attachments
	}
	return this.Attachments
}

// 升级活动（废弃）
type MailShenJiHuoDong struct {
	Attachments []*Attachment
}

func (this MailShenJiHuoDong) GetMailId() int32 {
	return int32(15)
}


func (this MailShenJiHuoDong) GetSendTime() int64 {
	return 0
}

func (this MailShenJiHuoDong) GetExpireTime() int64 {
	return 1
}
func (this MailShenJiHuoDong) GetTitle() string {
	return ""
}
func (this MailShenJiHuoDong) GetPriority() int8 {
	return 1
}
func (this MailShenJiHuoDong) GetContent() string {
	return ""
}
func (this MailShenJiHuoDong) GetMinLevel() int16 {
	return 0
}
func (this MailShenJiHuoDong) GetMaxLevel() int16 {
	return 0
}
func (this MailShenJiHuoDong) GetMinVIPLevel() int16 {
	return 0
}
func (this MailShenJiHuoDong) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailShenJiHuoDong) GetMinCreateTime() int64 {
	return 0
}
func (this MailShenJiHuoDong) GetMaxCreateTime() int64 {
	return 0
}
func (this MailShenJiHuoDong) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailShenJiHuoDong) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[15]; ok {
		return attachments
	}
	return this.Attachments
}

// 升级活动奖励（废弃）
type MailHuoDongJiangLi struct {
	Attachments []*Attachment
}

func (this MailHuoDongJiangLi) GetMailId() int32 {
	return int32(17)
}


func (this MailHuoDongJiangLi) GetSendTime() int64 {
	return 0
}

func (this MailHuoDongJiangLi) GetExpireTime() int64 {
	return 1
}
func (this MailHuoDongJiangLi) GetTitle() string {
	return ""
}
func (this MailHuoDongJiangLi) GetPriority() int8 {
	return 0
}
func (this MailHuoDongJiangLi) GetContent() string {
	return ""
}
func (this MailHuoDongJiangLi) GetMinLevel() int16 {
	return 0
}
func (this MailHuoDongJiangLi) GetMaxLevel() int16 {
	return 0
}
func (this MailHuoDongJiangLi) GetMinVIPLevel() int16 {
	return 0
}
func (this MailHuoDongJiangLi) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailHuoDongJiangLi) GetMinCreateTime() int64 {
	return 0
}
func (this MailHuoDongJiangLi) GetMaxCreateTime() int64 {
	return 0
}
func (this MailHuoDongJiangLi) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailHuoDongJiangLi) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[17]; ok {
		return attachments
	}
	return this.Attachments
}

// 仙尊签到奖励
type MailVIPSignAward struct {
	Attachments []*Attachment
}

func (this MailVIPSignAward) GetMailId() int32 {
	return int32(18)
}


func (this MailVIPSignAward) GetSendTime() int64 {
	return 0
}

func (this MailVIPSignAward) GetExpireTime() int64 {
	return 1
}
func (this MailVIPSignAward) GetTitle() string {
	return ""
}
func (this MailVIPSignAward) GetPriority() int8 {
	return 0
}
func (this MailVIPSignAward) GetContent() string {
	return ""
}
func (this MailVIPSignAward) GetMinLevel() int16 {
	return 0
}
func (this MailVIPSignAward) GetMaxLevel() int16 {
	return 0
}
func (this MailVIPSignAward) GetMinVIPLevel() int16 {
	return 0
}
func (this MailVIPSignAward) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailVIPSignAward) GetMinCreateTime() int64 {
	return 0
}
func (this MailVIPSignAward) GetMaxCreateTime() int64 {
	return 0
}
func (this MailVIPSignAward) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailVIPSignAward) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[18]; ok {
		return attachments
	}
	return this.Attachments
}

// 战力活动
type MailZhanLiHuoDong struct {
	Attachments []*Attachment
}

func (this MailZhanLiHuoDong) GetMailId() int32 {
	return int32(19)
}


func (this MailZhanLiHuoDong) GetSendTime() int64 {
	return 0
}

func (this MailZhanLiHuoDong) GetExpireTime() int64 {
	return 0
}
func (this MailZhanLiHuoDong) GetTitle() string {
	return ""
}
func (this MailZhanLiHuoDong) GetPriority() int8 {
	return 1
}
func (this MailZhanLiHuoDong) GetContent() string {
	return ""
}
func (this MailZhanLiHuoDong) GetMinLevel() int16 {
	return 0
}
func (this MailZhanLiHuoDong) GetMaxLevel() int16 {
	return 0
}
func (this MailZhanLiHuoDong) GetMinVIPLevel() int16 {
	return 0
}
func (this MailZhanLiHuoDong) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailZhanLiHuoDong) GetMinCreateTime() int64 {
	return 0
}
func (this MailZhanLiHuoDong) GetMaxCreateTime() int64 {
	return 0
}
func (this MailZhanLiHuoDong) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailZhanLiHuoDong) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[19]; ok {
		return attachments
	}
	return this.Attachments
}

// 战力活动奖励
type MailZhanLiHuoDongJiangLi struct {
	Attachments []*Attachment
}

func (this MailZhanLiHuoDongJiangLi) GetMailId() int32 {
	return int32(20)
}


func (this MailZhanLiHuoDongJiangLi) GetSendTime() int64 {
	return 0
}

func (this MailZhanLiHuoDongJiangLi) GetExpireTime() int64 {
	return 1
}
func (this MailZhanLiHuoDongJiangLi) GetTitle() string {
	return ""
}
func (this MailZhanLiHuoDongJiangLi) GetPriority() int8 {
	return 0
}
func (this MailZhanLiHuoDongJiangLi) GetContent() string {
	return ""
}
func (this MailZhanLiHuoDongJiangLi) GetMinLevel() int16 {
	return 0
}
func (this MailZhanLiHuoDongJiangLi) GetMaxLevel() int16 {
	return 0
}
func (this MailZhanLiHuoDongJiangLi) GetMinVIPLevel() int16 {
	return 0
}
func (this MailZhanLiHuoDongJiangLi) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailZhanLiHuoDongJiangLi) GetMinCreateTime() int64 {
	return 0
}
func (this MailZhanLiHuoDongJiangLi) GetMaxCreateTime() int64 {
	return 0
}
func (this MailZhanLiHuoDongJiangLi) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailZhanLiHuoDongJiangLi) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[20]; ok {
		return attachments
	}
	return this.Attachments
}

// 仙尊特权奖励
type MailXianZun struct {
	VipNum string // vip等级
	Attachments []*Attachment
}

func (this MailXianZun) GetMailId() int32 {
	return int32(22)
}


func (this MailXianZun) GetSendTime() int64 {
	return 0
}

func (this MailXianZun) GetExpireTime() int64 {
	return 1
}
func (this MailXianZun) GetTitle() string {
	return ""
}
func (this MailXianZun) GetPriority() int8 {
	return 1
}
func (this MailXianZun) GetContent() string {
	return ""
}
func (this MailXianZun) GetMinLevel() int16 {
	return 0
}
func (this MailXianZun) GetMaxLevel() int16 {
	return 0
}
func (this MailXianZun) GetMinVIPLevel() int16 {
	return 0
}
func (this MailXianZun) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailXianZun) GetMinCreateTime() int64 {
	return 0
}
func (this MailXianZun) GetMaxCreateTime() int64 {
	return 0
}
func (this MailXianZun) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.VipNum))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailXianZun) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[22]; ok {
		return attachments
	}
	return this.Attachments
}

// 道具获得提示
type MailDaoJuHuoDe struct {
	ItemName string // 道具名称
	Func string // 功能
	Attachments []*Attachment
}

func (this MailDaoJuHuoDe) GetMailId() int32 {
	return int32(23)
}


func (this MailDaoJuHuoDe) GetSendTime() int64 {
	return 0
}

func (this MailDaoJuHuoDe) GetExpireTime() int64 {
	return 0
}
func (this MailDaoJuHuoDe) GetTitle() string {
	return ""
}
func (this MailDaoJuHuoDe) GetPriority() int8 {
	return 0
}
func (this MailDaoJuHuoDe) GetContent() string {
	return ""
}
func (this MailDaoJuHuoDe) GetMinLevel() int16 {
	return 0
}
func (this MailDaoJuHuoDe) GetMaxLevel() int16 {
	return 0
}
func (this MailDaoJuHuoDe) GetMinVIPLevel() int16 {
	return 0
}
func (this MailDaoJuHuoDe) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailDaoJuHuoDe) GetMinCreateTime() int64 {
	return 0
}
func (this MailDaoJuHuoDe) GetMaxCreateTime() int64 {
	return 0
}
func (this MailDaoJuHuoDe) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.ItemName))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Func))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailDaoJuHuoDe) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[23]; ok {
		return attachments
	}
	return this.Attachments
}

// 七日新手礼活动（废弃）
type MailQiriXinShouLi struct {
	Attachments []*Attachment
}

func (this MailQiriXinShouLi) GetMailId() int32 {
	return int32(24)
}


func (this MailQiriXinShouLi) GetSendTime() int64 {
	return 0
}

func (this MailQiriXinShouLi) GetExpireTime() int64 {
	return 0
}
func (this MailQiriXinShouLi) GetTitle() string {
	return ""
}
func (this MailQiriXinShouLi) GetPriority() int8 {
	return 0
}
func (this MailQiriXinShouLi) GetContent() string {
	return ""
}
func (this MailQiriXinShouLi) GetMinLevel() int16 {
	return 0
}
func (this MailQiriXinShouLi) GetMaxLevel() int16 {
	return 0
}
func (this MailQiriXinShouLi) GetMinVIPLevel() int16 {
	return 0
}
func (this MailQiriXinShouLi) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailQiriXinShouLi) GetMinCreateTime() int64 {
	return 0
}
func (this MailQiriXinShouLi) GetMaxCreateTime() int64 {
	return 0
}
func (this MailQiriXinShouLi) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailQiriXinShouLi) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[24]; ok {
		return attachments
	}
	return this.Attachments
}

// 打坐小提醒
type MailDaZuoTiXing struct {
	Attachments []*Attachment
}

func (this MailDaZuoTiXing) GetMailId() int32 {
	return int32(25)
}


func (this MailDaZuoTiXing) GetSendTime() int64 {
	return 0
}

func (this MailDaZuoTiXing) GetExpireTime() int64 {
	return 0
}
func (this MailDaZuoTiXing) GetTitle() string {
	return ""
}
func (this MailDaZuoTiXing) GetPriority() int8 {
	return 0
}
func (this MailDaZuoTiXing) GetContent() string {
	return ""
}
func (this MailDaZuoTiXing) GetMinLevel() int16 {
	return 0
}
func (this MailDaZuoTiXing) GetMaxLevel() int16 {
	return 0
}
func (this MailDaZuoTiXing) GetMinVIPLevel() int16 {
	return 0
}
func (this MailDaZuoTiXing) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailDaZuoTiXing) GetMinCreateTime() int64 {
	return 0
}
func (this MailDaZuoTiXing) GetMaxCreateTime() int64 {
	return 0
}
func (this MailDaZuoTiXing) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailDaZuoTiXing) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[25]; ok {
		return attachments
	}
	return this.Attachments
}

// 注册好礼
type MailZhuCe struct {
	Attachments []*Attachment
}

func (this MailZhuCe) GetMailId() int32 {
	return int32(26)
}


func (this MailZhuCe) GetSendTime() int64 {
	return 0
}

func (this MailZhuCe) GetExpireTime() int64 {
	return 1
}
func (this MailZhuCe) GetTitle() string {
	return ""
}
func (this MailZhuCe) GetPriority() int8 {
	return 0
}
func (this MailZhuCe) GetContent() string {
	return ""
}
func (this MailZhuCe) GetMinLevel() int16 {
	return 0
}
func (this MailZhuCe) GetMaxLevel() int16 {
	return 0
}
func (this MailZhuCe) GetMinVIPLevel() int16 {
	return 0
}
func (this MailZhuCe) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailZhuCe) GetMinCreateTime() int64 {
	return 0
}
func (this MailZhuCe) GetMaxCreateTime() int64 {
	return 0
}
func (this MailZhuCe) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailZhuCe) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[26]; ok {
		return attachments
	}
	return this.Attachments
}

// 侠客见面礼
type MailNewbieGift struct {
	Attachments []*Attachment
}

func (this MailNewbieGift) GetMailId() int32 {
	return int32(28)
}


func (this MailNewbieGift) GetSendTime() int64 {
	return 0
}

func (this MailNewbieGift) GetExpireTime() int64 {
	return 1
}
func (this MailNewbieGift) GetTitle() string {
	return ""
}
func (this MailNewbieGift) GetPriority() int8 {
	return 0
}
func (this MailNewbieGift) GetContent() string {
	return ""
}
func (this MailNewbieGift) GetMinLevel() int16 {
	return 0
}
func (this MailNewbieGift) GetMaxLevel() int16 {
	return 0
}
func (this MailNewbieGift) GetMinVIPLevel() int16 {
	return 0
}
func (this MailNewbieGift) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailNewbieGift) GetMinCreateTime() int64 {
	return 0
}
func (this MailNewbieGift) GetMaxCreateTime() int64 {
	return 0
}
func (this MailNewbieGift) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailNewbieGift) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[28]; ok {
		return attachments
	}
	return this.Attachments
}

// 首充豪华礼
type MailShouChongHaoHuaLi struct {
	Attachments []*Attachment
}

func (this MailShouChongHaoHuaLi) GetMailId() int32 {
	return int32(29)
}


func (this MailShouChongHaoHuaLi) GetSendTime() int64 {
	return 0
}

func (this MailShouChongHaoHuaLi) GetExpireTime() int64 {
	return 1
}
func (this MailShouChongHaoHuaLi) GetTitle() string {
	return ""
}
func (this MailShouChongHaoHuaLi) GetPriority() int8 {
	return 0
}
func (this MailShouChongHaoHuaLi) GetContent() string {
	return ""
}
func (this MailShouChongHaoHuaLi) GetMinLevel() int16 {
	return 0
}
func (this MailShouChongHaoHuaLi) GetMaxLevel() int16 {
	return 0
}
func (this MailShouChongHaoHuaLi) GetMinVIPLevel() int16 {
	return 0
}
func (this MailShouChongHaoHuaLi) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailShouChongHaoHuaLi) GetMinCreateTime() int64 {
	return 0
}
func (this MailShouChongHaoHuaLi) GetMaxCreateTime() int64 {
	return 0
}
func (this MailShouChongHaoHuaLi) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailShouChongHaoHuaLi) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[29]; ok {
		return attachments
	}
	return this.Attachments
}

// 胧月参上
type MailYaoQingLongYue struct {
	Attachments []*Attachment
}

func (this MailYaoQingLongYue) GetMailId() int32 {
	return int32(30)
}


func (this MailYaoQingLongYue) GetSendTime() int64 {
	return 0
}

func (this MailYaoQingLongYue) GetExpireTime() int64 {
	return 1
}
func (this MailYaoQingLongYue) GetTitle() string {
	return ""
}
func (this MailYaoQingLongYue) GetPriority() int8 {
	return 0
}
func (this MailYaoQingLongYue) GetContent() string {
	return ""
}
func (this MailYaoQingLongYue) GetMinLevel() int16 {
	return 0
}
func (this MailYaoQingLongYue) GetMaxLevel() int16 {
	return 0
}
func (this MailYaoQingLongYue) GetMinVIPLevel() int16 {
	return 0
}
func (this MailYaoQingLongYue) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailYaoQingLongYue) GetMinCreateTime() int64 {
	return 0
}
func (this MailYaoQingLongYue) GetMaxCreateTime() int64 {
	return 0
}
func (this MailYaoQingLongYue) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailYaoQingLongYue) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[30]; ok {
		return attachments
	}
	return this.Attachments
}

// 新春红包活动奖励
type MailRedPaper struct {
	Attachments []*Attachment
}

func (this MailRedPaper) GetMailId() int32 {
	return int32(31)
}


func (this MailRedPaper) GetSendTime() int64 {
	return 0
}

func (this MailRedPaper) GetExpireTime() int64 {
	return 0
}
func (this MailRedPaper) GetTitle() string {
	return ""
}
func (this MailRedPaper) GetPriority() int8 {
	return 0
}
func (this MailRedPaper) GetContent() string {
	return ""
}
func (this MailRedPaper) GetMinLevel() int16 {
	return 0
}
func (this MailRedPaper) GetMaxLevel() int16 {
	return 0
}
func (this MailRedPaper) GetMinVIPLevel() int16 {
	return 0
}
func (this MailRedPaper) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailRedPaper) GetMinCreateTime() int64 {
	return 0
}
func (this MailRedPaper) GetMaxCreateTime() int64 {
	return 0
}
func (this MailRedPaper) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailRedPaper) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[31]; ok {
		return attachments
	}
	return this.Attachments
}

// 会员独享新手礼包
type MailQQVip struct {
	Attachments []*Attachment
}

func (this MailQQVip) GetMailId() int32 {
	return int32(32)
}


func (this MailQQVip) GetSendTime() int64 {
	return 0
}

func (this MailQQVip) GetExpireTime() int64 {
	return 0
}
func (this MailQQVip) GetTitle() string {
	return ""
}
func (this MailQQVip) GetPriority() int8 {
	return 0
}
func (this MailQQVip) GetContent() string {
	return ""
}
func (this MailQQVip) GetMinLevel() int16 {
	return 0
}
func (this MailQQVip) GetMaxLevel() int16 {
	return 0
}
func (this MailQQVip) GetMinVIPLevel() int16 {
	return 0
}
func (this MailQQVip) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailQQVip) GetMinCreateTime() int64 {
	return 0
}
func (this MailQQVip) GetMaxCreateTime() int64 {
	return 0
}
func (this MailQQVip) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailQQVip) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[32]; ok {
		return attachments
	}
	return this.Attachments
}

// 超级会员尊贵新手礼包
type MailQQSvip struct {
	Attachments []*Attachment
}

func (this MailQQSvip) GetMailId() int32 {
	return int32(33)
}


func (this MailQQSvip) GetSendTime() int64 {
	return 0
}

func (this MailQQSvip) GetExpireTime() int64 {
	return 0
}
func (this MailQQSvip) GetTitle() string {
	return ""
}
func (this MailQQSvip) GetPriority() int8 {
	return 0
}
func (this MailQQSvip) GetContent() string {
	return ""
}
func (this MailQQSvip) GetMinLevel() int16 {
	return 0
}
func (this MailQQSvip) GetMaxLevel() int16 {
	return 0
}
func (this MailQQSvip) GetMinVIPLevel() int16 {
	return 0
}
func (this MailQQSvip) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailQQSvip) GetMinCreateTime() int64 {
	return 0
}
func (this MailQQSvip) GetMaxCreateTime() int64 {
	return 0
}
func (this MailQQSvip) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailQQSvip) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[33]; ok {
		return attachments
	}
	return this.Attachments
}

// 会员开通/续费礼包
type MailQQvipXuFei struct {
	Attachments []*Attachment
}

func (this MailQQvipXuFei) GetMailId() int32 {
	return int32(34)
}


func (this MailQQvipXuFei) GetSendTime() int64 {
	return 0
}

func (this MailQQvipXuFei) GetExpireTime() int64 {
	return 0
}
func (this MailQQvipXuFei) GetTitle() string {
	return ""
}
func (this MailQQvipXuFei) GetPriority() int8 {
	return 0
}
func (this MailQQvipXuFei) GetContent() string {
	return ""
}
func (this MailQQvipXuFei) GetMinLevel() int16 {
	return 0
}
func (this MailQQvipXuFei) GetMaxLevel() int16 {
	return 0
}
func (this MailQQvipXuFei) GetMinVIPLevel() int16 {
	return 0
}
func (this MailQQvipXuFei) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailQQvipXuFei) GetMinCreateTime() int64 {
	return 0
}
func (this MailQQvipXuFei) GetMaxCreateTime() int64 {
	return 0
}
func (this MailQQvipXuFei) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailQQvipXuFei) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[34]; ok {
		return attachments
	}
	return this.Attachments
}

// 超级会员开通/续费礼包
type MailQQSvipXuFei struct {
	Attachments []*Attachment
}

func (this MailQQSvipXuFei) GetMailId() int32 {
	return int32(35)
}


func (this MailQQSvipXuFei) GetSendTime() int64 {
	return 0
}

func (this MailQQSvipXuFei) GetExpireTime() int64 {
	return 0
}
func (this MailQQSvipXuFei) GetTitle() string {
	return ""
}
func (this MailQQSvipXuFei) GetPriority() int8 {
	return 0
}
func (this MailQQSvipXuFei) GetContent() string {
	return ""
}
func (this MailQQSvipXuFei) GetMinLevel() int16 {
	return 0
}
func (this MailQQSvipXuFei) GetMaxLevel() int16 {
	return 0
}
func (this MailQQSvipXuFei) GetMinVIPLevel() int16 {
	return 0
}
func (this MailQQSvipXuFei) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailQQSvipXuFei) GetMinCreateTime() int64 {
	return 0
}
func (this MailQQSvipXuFei) GetMaxCreateTime() int64 {
	return 0
}
func (this MailQQSvipXuFei) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailQQSvipXuFei) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[35]; ok {
		return attachments
	}
	return this.Attachments
}

// 阵印背包已满提示
type MailTotemBagFull struct {
	Num string // 数量
	Attachments []*Attachment
}

func (this MailTotemBagFull) GetMailId() int32 {
	return int32(36)
}


func (this MailTotemBagFull) GetSendTime() int64 {
	return 0
}

func (this MailTotemBagFull) GetExpireTime() int64 {
	return 0
}
func (this MailTotemBagFull) GetTitle() string {
	return ""
}
func (this MailTotemBagFull) GetPriority() int8 {
	return 0
}
func (this MailTotemBagFull) GetContent() string {
	return ""
}
func (this MailTotemBagFull) GetMinLevel() int16 {
	return 0
}
func (this MailTotemBagFull) GetMaxLevel() int16 {
	return 0
}
func (this MailTotemBagFull) GetMinVIPLevel() int16 {
	return 0
}
func (this MailTotemBagFull) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailTotemBagFull) GetMinCreateTime() int64 {
	return 0
}
func (this MailTotemBagFull) GetMaxCreateTime() int64 {
	return 0
}
func (this MailTotemBagFull) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Num))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailTotemBagFull) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[36]; ok {
		return attachments
	}
	return this.Attachments
}

// 帮派邮件
type Mailcliquesalaryowner struct {
	Name string // 帮派名
	Num string // 帮派人数
	Coins string // 奖励铜钱
	Attachments []*Attachment
}

func (this Mailcliquesalaryowner) GetMailId() int32 {
	return int32(37)
}


func (this Mailcliquesalaryowner) GetSendTime() int64 {
	return 0
}

func (this Mailcliquesalaryowner) GetExpireTime() int64 {
	return 14
}
func (this Mailcliquesalaryowner) GetTitle() string {
	return ""
}
func (this Mailcliquesalaryowner) GetPriority() int8 {
	return 0
}
func (this Mailcliquesalaryowner) GetContent() string {
	return ""
}
func (this Mailcliquesalaryowner) GetMinLevel() int16 {
	return 0
}
func (this Mailcliquesalaryowner) GetMaxLevel() int16 {
	return 0
}
func (this Mailcliquesalaryowner) GetMinVIPLevel() int16 {
	return 0
}
func (this Mailcliquesalaryowner) GetMaxVIPLevel() int16 {
	return 0
}
func (this Mailcliquesalaryowner) GetMinCreateTime() int64 {
	return 0
}
func (this Mailcliquesalaryowner) GetMaxCreateTime() int64 {
	return 0
}
func (this Mailcliquesalaryowner) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Name))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Num))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Coins))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this Mailcliquesalaryowner) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[37]; ok {
		return attachments
	}
	return this.Attachments
}

// 帮派邮件
type Mailcliquesalarymanager struct {
	Name string // 帮派名
	Num string // 帮派人数
	Coins string // 奖励铜钱
	Attachments []*Attachment
}

func (this Mailcliquesalarymanager) GetMailId() int32 {
	return int32(38)
}


func (this Mailcliquesalarymanager) GetSendTime() int64 {
	return 0
}

func (this Mailcliquesalarymanager) GetExpireTime() int64 {
	return 14
}
func (this Mailcliquesalarymanager) GetTitle() string {
	return ""
}
func (this Mailcliquesalarymanager) GetPriority() int8 {
	return 0
}
func (this Mailcliquesalarymanager) GetContent() string {
	return ""
}
func (this Mailcliquesalarymanager) GetMinLevel() int16 {
	return 0
}
func (this Mailcliquesalarymanager) GetMaxLevel() int16 {
	return 0
}
func (this Mailcliquesalarymanager) GetMinVIPLevel() int16 {
	return 0
}
func (this Mailcliquesalarymanager) GetMaxVIPLevel() int16 {
	return 0
}
func (this Mailcliquesalarymanager) GetMinCreateTime() int64 {
	return 0
}
func (this Mailcliquesalarymanager) GetMaxCreateTime() int64 {
	return 0
}
func (this Mailcliquesalarymanager) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Name))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Num))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Coins))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this Mailcliquesalarymanager) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[38]; ok {
		return attachments
	}
	return this.Attachments
}

// 帮派邮件
type Mailcliqueleave struct {
	Name string // 帮派名
	Attachments []*Attachment
}

func (this Mailcliqueleave) GetMailId() int32 {
	return int32(39)
}


func (this Mailcliqueleave) GetSendTime() int64 {
	return 0
}

func (this Mailcliqueleave) GetExpireTime() int64 {
	return 0
}
func (this Mailcliqueleave) GetTitle() string {
	return ""
}
func (this Mailcliqueleave) GetPriority() int8 {
	return 0
}
func (this Mailcliqueleave) GetContent() string {
	return ""
}
func (this Mailcliqueleave) GetMinLevel() int16 {
	return 0
}
func (this Mailcliqueleave) GetMaxLevel() int16 {
	return 0
}
func (this Mailcliqueleave) GetMinVIPLevel() int16 {
	return 0
}
func (this Mailcliqueleave) GetMaxVIPLevel() int16 {
	return 0
}
func (this Mailcliqueleave) GetMinCreateTime() int64 {
	return 0
}
func (this Mailcliqueleave) GetMaxCreateTime() int64 {
	return 0
}
func (this Mailcliqueleave) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Name))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this Mailcliqueleave) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[39]; ok {
		return attachments
	}
	return this.Attachments
}

// 帮派邮件
type Mailcliquebemanger struct {
	Attachments []*Attachment
}

func (this Mailcliquebemanger) GetMailId() int32 {
	return int32(40)
}


func (this Mailcliquebemanger) GetSendTime() int64 {
	return 0
}

func (this Mailcliquebemanger) GetExpireTime() int64 {
	return 0
}
func (this Mailcliquebemanger) GetTitle() string {
	return ""
}
func (this Mailcliquebemanger) GetPriority() int8 {
	return 0
}
func (this Mailcliquebemanger) GetContent() string {
	return ""
}
func (this Mailcliquebemanger) GetMinLevel() int16 {
	return 0
}
func (this Mailcliquebemanger) GetMaxLevel() int16 {
	return 0
}
func (this Mailcliquebemanger) GetMinVIPLevel() int16 {
	return 0
}
func (this Mailcliquebemanger) GetMaxVIPLevel() int16 {
	return 0
}
func (this Mailcliquebemanger) GetMinCreateTime() int64 {
	return 0
}
func (this Mailcliquebemanger) GetMaxCreateTime() int64 {
	return 0
}
func (this Mailcliquebemanger) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this Mailcliquebemanger) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[40]; ok {
		return attachments
	}
	return this.Attachments
}

// 帮派邮件
type Mailcliquecancelmanager struct {
	Attachments []*Attachment
}

func (this Mailcliquecancelmanager) GetMailId() int32 {
	return int32(41)
}


func (this Mailcliquecancelmanager) GetSendTime() int64 {
	return 0
}

func (this Mailcliquecancelmanager) GetExpireTime() int64 {
	return 0
}
func (this Mailcliquecancelmanager) GetTitle() string {
	return ""
}
func (this Mailcliquecancelmanager) GetPriority() int8 {
	return 0
}
func (this Mailcliquecancelmanager) GetContent() string {
	return ""
}
func (this Mailcliquecancelmanager) GetMinLevel() int16 {
	return 0
}
func (this Mailcliquecancelmanager) GetMaxLevel() int16 {
	return 0
}
func (this Mailcliquecancelmanager) GetMinVIPLevel() int16 {
	return 0
}
func (this Mailcliquecancelmanager) GetMaxVIPLevel() int16 {
	return 0
}
func (this Mailcliquecancelmanager) GetMinCreateTime() int64 {
	return 0
}
func (this Mailcliquecancelmanager) GetMaxCreateTime() int64 {
	return 0
}
func (this Mailcliquecancelmanager) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this Mailcliquecancelmanager) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[41]; ok {
		return attachments
	}
	return this.Attachments
}

// 帮派邮件
type Mailcliquecancelowner struct {
	Name string // 新任帮主
	Attachments []*Attachment
}

func (this Mailcliquecancelowner) GetMailId() int32 {
	return int32(42)
}


func (this Mailcliquecancelowner) GetSendTime() int64 {
	return 0
}

func (this Mailcliquecancelowner) GetExpireTime() int64 {
	return 0
}
func (this Mailcliquecancelowner) GetTitle() string {
	return ""
}
func (this Mailcliquecancelowner) GetPriority() int8 {
	return 0
}
func (this Mailcliquecancelowner) GetContent() string {
	return ""
}
func (this Mailcliquecancelowner) GetMinLevel() int16 {
	return 0
}
func (this Mailcliquecancelowner) GetMaxLevel() int16 {
	return 0
}
func (this Mailcliquecancelowner) GetMinVIPLevel() int16 {
	return 0
}
func (this Mailcliquecancelowner) GetMaxVIPLevel() int16 {
	return 0
}
func (this Mailcliquecancelowner) GetMinCreateTime() int64 {
	return 0
}
func (this Mailcliquecancelowner) GetMaxCreateTime() int64 {
	return 0
}
func (this Mailcliquecancelowner) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Name))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this Mailcliquecancelowner) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[42]; ok {
		return attachments
	}
	return this.Attachments
}

// 帮派邮件
type Mailcliquechangeowner struct {
	Name string // 老帮主
	Attachments []*Attachment
}

func (this Mailcliquechangeowner) GetMailId() int32 {
	return int32(43)
}


func (this Mailcliquechangeowner) GetSendTime() int64 {
	return 0
}

func (this Mailcliquechangeowner) GetExpireTime() int64 {
	return 0
}
func (this Mailcliquechangeowner) GetTitle() string {
	return ""
}
func (this Mailcliquechangeowner) GetPriority() int8 {
	return 0
}
func (this Mailcliquechangeowner) GetContent() string {
	return ""
}
func (this Mailcliquechangeowner) GetMinLevel() int16 {
	return 0
}
func (this Mailcliquechangeowner) GetMaxLevel() int16 {
	return 0
}
func (this Mailcliquechangeowner) GetMinVIPLevel() int16 {
	return 0
}
func (this Mailcliquechangeowner) GetMaxVIPLevel() int16 {
	return 0
}
func (this Mailcliquechangeowner) GetMinCreateTime() int64 {
	return 0
}
func (this Mailcliquechangeowner) GetMaxCreateTime() int64 {
	return 0
}
func (this Mailcliquechangeowner) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Name))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this Mailcliquechangeowner) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[43]; ok {
		return attachments
	}
	return this.Attachments
}

// 帮派邮件
type Mailcliquedestory struct {
	Name string // 帮派名
	Attachments []*Attachment
}

func (this Mailcliquedestory) GetMailId() int32 {
	return int32(44)
}


func (this Mailcliquedestory) GetSendTime() int64 {
	return 0
}

func (this Mailcliquedestory) GetExpireTime() int64 {
	return 0
}
func (this Mailcliquedestory) GetTitle() string {
	return ""
}
func (this Mailcliquedestory) GetPriority() int8 {
	return 0
}
func (this Mailcliquedestory) GetContent() string {
	return ""
}
func (this Mailcliquedestory) GetMinLevel() int16 {
	return 0
}
func (this Mailcliquedestory) GetMaxLevel() int16 {
	return 0
}
func (this Mailcliquedestory) GetMinVIPLevel() int16 {
	return 0
}
func (this Mailcliquedestory) GetMaxVIPLevel() int16 {
	return 0
}
func (this Mailcliquedestory) GetMinCreateTime() int64 {
	return 0
}
func (this Mailcliquedestory) GetMaxCreateTime() int64 {
	return 0
}
func (this Mailcliquedestory) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Name))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this Mailcliquedestory) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[44]; ok {
		return attachments
	}
	return this.Attachments
}

// 上香祈福
type Mailzongcishangxiangqifujiang struct {
	Num string // 铜钱数量
	Attachments []*Attachment
}

func (this Mailzongcishangxiangqifujiang) GetMailId() int32 {
	return int32(45)
}


func (this Mailzongcishangxiangqifujiang) GetSendTime() int64 {
	return 0
}

func (this Mailzongcishangxiangqifujiang) GetExpireTime() int64 {
	return 0
}
func (this Mailzongcishangxiangqifujiang) GetTitle() string {
	return ""
}
func (this Mailzongcishangxiangqifujiang) GetPriority() int8 {
	return 0
}
func (this Mailzongcishangxiangqifujiang) GetContent() string {
	return ""
}
func (this Mailzongcishangxiangqifujiang) GetMinLevel() int16 {
	return 0
}
func (this Mailzongcishangxiangqifujiang) GetMaxLevel() int16 {
	return 0
}
func (this Mailzongcishangxiangqifujiang) GetMinVIPLevel() int16 {
	return 0
}
func (this Mailzongcishangxiangqifujiang) GetMaxVIPLevel() int16 {
	return 0
}
func (this Mailzongcishangxiangqifujiang) GetMinCreateTime() int64 {
	return 0
}
func (this Mailzongcishangxiangqifujiang) GetMaxCreateTime() int64 {
	return 0
}
func (this Mailzongcishangxiangqifujiang) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Num))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this Mailzongcishangxiangqifujiang) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[45]; ok {
		return attachments
	}
	return this.Attachments
}

// 帮派邮件
type Mailcliqueleavebank struct {
	Attachments []*Attachment
}

func (this Mailcliqueleavebank) GetMailId() int32 {
	return int32(46)
}


func (this Mailcliqueleavebank) GetSendTime() int64 {
	return 0
}

func (this Mailcliqueleavebank) GetExpireTime() int64 {
	return 0
}
func (this Mailcliqueleavebank) GetTitle() string {
	return ""
}
func (this Mailcliqueleavebank) GetPriority() int8 {
	return 0
}
func (this Mailcliqueleavebank) GetContent() string {
	return ""
}
func (this Mailcliqueleavebank) GetMinLevel() int16 {
	return 0
}
func (this Mailcliqueleavebank) GetMaxLevel() int16 {
	return 0
}
func (this Mailcliqueleavebank) GetMinVIPLevel() int16 {
	return 0
}
func (this Mailcliqueleavebank) GetMaxVIPLevel() int16 {
	return 0
}
func (this Mailcliqueleavebank) GetMinCreateTime() int64 {
	return 0
}
func (this Mailcliqueleavebank) GetMaxCreateTime() int64 {
	return 0
}
func (this Mailcliqueleavebank) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this Mailcliqueleavebank) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[46]; ok {
		return attachments
	}
	return this.Attachments
}

// 帮派邮件
type Mailcliquejoin struct {
	Name string // 帮派名
	Attachments []*Attachment
}

func (this Mailcliquejoin) GetMailId() int32 {
	return int32(47)
}


func (this Mailcliquejoin) GetSendTime() int64 {
	return 0
}

func (this Mailcliquejoin) GetExpireTime() int64 {
	return 0
}
func (this Mailcliquejoin) GetTitle() string {
	return ""
}
func (this Mailcliquejoin) GetPriority() int8 {
	return 0
}
func (this Mailcliquejoin) GetContent() string {
	return ""
}
func (this Mailcliquejoin) GetMinLevel() int16 {
	return 0
}
func (this Mailcliquejoin) GetMaxLevel() int16 {
	return 0
}
func (this Mailcliquejoin) GetMinVIPLevel() int16 {
	return 0
}
func (this Mailcliquejoin) GetMaxVIPLevel() int16 {
	return 0
}
func (this Mailcliquejoin) GetMinCreateTime() int64 {
	return 0
}
func (this Mailcliquejoin) GetMaxCreateTime() int64 {
	return 0
}
func (this Mailcliquejoin) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Name))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this Mailcliquejoin) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[47]; ok {
		return attachments
	}
	return this.Attachments
}

// 豪华等级礼包
type MailDengJiLiBao struct {
	Attachments []*Attachment
}

func (this MailDengJiLiBao) GetMailId() int32 {
	return int32(48)
}


func (this MailDengJiLiBao) GetSendTime() int64 {
	return 0
}

func (this MailDengJiLiBao) GetExpireTime() int64 {
	return 0
}
func (this MailDengJiLiBao) GetTitle() string {
	return ""
}
func (this MailDengJiLiBao) GetPriority() int8 {
	return 0
}
func (this MailDengJiLiBao) GetContent() string {
	return ""
}
func (this MailDengJiLiBao) GetMinLevel() int16 {
	return 0
}
func (this MailDengJiLiBao) GetMaxLevel() int16 {
	return 0
}
func (this MailDengJiLiBao) GetMinVIPLevel() int16 {
	return 0
}
func (this MailDengJiLiBao) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailDengJiLiBao) GetMinCreateTime() int64 {
	return 0
}
func (this MailDengJiLiBao) GetMaxCreateTime() int64 {
	return 0
}
func (this MailDengJiLiBao) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailDengJiLiBao) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[48]; ok {
		return attachments
	}
	return this.Attachments
}

// Boss讨伐奖励
type MailBossTaoFa struct {
	BossName string // Boss名称
	Hurt string // 伤害
	Exp string // 经验数量
	Coins string // 铜钱数量
	Fame string // 声望数量
	Attachments []*Attachment
}

func (this MailBossTaoFa) GetMailId() int32 {
	return int32(49)
}


func (this MailBossTaoFa) GetSendTime() int64 {
	return 0
}

func (this MailBossTaoFa) GetExpireTime() int64 {
	return 0
}
func (this MailBossTaoFa) GetTitle() string {
	return ""
}
func (this MailBossTaoFa) GetPriority() int8 {
	return 0
}
func (this MailBossTaoFa) GetContent() string {
	return ""
}
func (this MailBossTaoFa) GetMinLevel() int16 {
	return 0
}
func (this MailBossTaoFa) GetMaxLevel() int16 {
	return 0
}
func (this MailBossTaoFa) GetMinVIPLevel() int16 {
	return 0
}
func (this MailBossTaoFa) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailBossTaoFa) GetMinCreateTime() int64 {
	return 0
}
func (this MailBossTaoFa) GetMaxCreateTime() int64 {
	return 0
}
func (this MailBossTaoFa) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.BossName))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Hurt))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Exp))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Coins))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Fame))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailBossTaoFa) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[49]; ok {
		return attachments
	}
	return this.Attachments
}

// Boss讨伐精英奖励
type MailBossTaoFaJingYing struct {
	BossName string // Boss名称
	Rank string // 伤害排名
	Hurt string // 伤害
	Exp string // 经验数量
	Coins string // 铜钱数量
	Fame string // 声望数量
	Attachments []*Attachment
}

func (this MailBossTaoFaJingYing) GetMailId() int32 {
	return int32(50)
}


func (this MailBossTaoFaJingYing) GetSendTime() int64 {
	return 0
}

func (this MailBossTaoFaJingYing) GetExpireTime() int64 {
	return 0
}
func (this MailBossTaoFaJingYing) GetTitle() string {
	return ""
}
func (this MailBossTaoFaJingYing) GetPriority() int8 {
	return 0
}
func (this MailBossTaoFaJingYing) GetContent() string {
	return ""
}
func (this MailBossTaoFaJingYing) GetMinLevel() int16 {
	return 0
}
func (this MailBossTaoFaJingYing) GetMaxLevel() int16 {
	return 0
}
func (this MailBossTaoFaJingYing) GetMinVIPLevel() int16 {
	return 0
}
func (this MailBossTaoFaJingYing) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailBossTaoFaJingYing) GetMinCreateTime() int64 {
	return 0
}
func (this MailBossTaoFaJingYing) GetMaxCreateTime() int64 {
	return 0
}
func (this MailBossTaoFaJingYing) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.BossName))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Rank))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Hurt))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Exp))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Coins))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Fame))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailBossTaoFaJingYing) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[50]; ok {
		return attachments
	}
	return this.Attachments
}

// 讨伐胜利
type MailTaoFaShengLi struct {
	CampName string // 敌人军团名称
	Point string // 讨伐点数量
	Hope string // 希望之光数量
	Attachments []*Attachment
}

func (this MailTaoFaShengLi) GetMailId() int32 {
	return int32(51)
}


func (this MailTaoFaShengLi) GetSendTime() int64 {
	return 0
}

func (this MailTaoFaShengLi) GetExpireTime() int64 {
	return 0
}
func (this MailTaoFaShengLi) GetTitle() string {
	return ""
}
func (this MailTaoFaShengLi) GetPriority() int8 {
	return 0
}
func (this MailTaoFaShengLi) GetContent() string {
	return ""
}
func (this MailTaoFaShengLi) GetMinLevel() int16 {
	return 0
}
func (this MailTaoFaShengLi) GetMaxLevel() int16 {
	return 0
}
func (this MailTaoFaShengLi) GetMinVIPLevel() int16 {
	return 0
}
func (this MailTaoFaShengLi) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailTaoFaShengLi) GetMinCreateTime() int64 {
	return 0
}
func (this MailTaoFaShengLi) GetMaxCreateTime() int64 {
	return 0
}
func (this MailTaoFaShengLi) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.CampName))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Point))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Hope))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailTaoFaShengLi) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[51]; ok {
		return attachments
	}
	return this.Attachments
}

// 帮派宝箱福利
type MailCliqueBangPaiBaoXiang struct {
	Attachments []*Attachment
}

func (this MailCliqueBangPaiBaoXiang) GetMailId() int32 {
	return int32(52)
}


func (this MailCliqueBangPaiBaoXiang) GetSendTime() int64 {
	return 0
}

func (this MailCliqueBangPaiBaoXiang) GetExpireTime() int64 {
	return 0
}
func (this MailCliqueBangPaiBaoXiang) GetTitle() string {
	return ""
}
func (this MailCliqueBangPaiBaoXiang) GetPriority() int8 {
	return 0
}
func (this MailCliqueBangPaiBaoXiang) GetContent() string {
	return ""
}
func (this MailCliqueBangPaiBaoXiang) GetMinLevel() int16 {
	return 0
}
func (this MailCliqueBangPaiBaoXiang) GetMaxLevel() int16 {
	return 0
}
func (this MailCliqueBangPaiBaoXiang) GetMinVIPLevel() int16 {
	return 0
}
func (this MailCliqueBangPaiBaoXiang) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailCliqueBangPaiBaoXiang) GetMinCreateTime() int64 {
	return 0
}
func (this MailCliqueBangPaiBaoXiang) GetMaxCreateTime() int64 {
	return 0
}
func (this MailCliqueBangPaiBaoXiang) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailCliqueBangPaiBaoXiang) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[52]; ok {
		return attachments
	}
	return this.Attachments
}

// 帮派军粮福利
type MailCliqueBangPaiJunLiang struct {
	Attachments []*Attachment
}

func (this MailCliqueBangPaiJunLiang) GetMailId() int32 {
	return int32(53)
}


func (this MailCliqueBangPaiJunLiang) GetSendTime() int64 {
	return 0
}

func (this MailCliqueBangPaiJunLiang) GetExpireTime() int64 {
	return 0
}
func (this MailCliqueBangPaiJunLiang) GetTitle() string {
	return ""
}
func (this MailCliqueBangPaiJunLiang) GetPriority() int8 {
	return 0
}
func (this MailCliqueBangPaiJunLiang) GetContent() string {
	return ""
}
func (this MailCliqueBangPaiJunLiang) GetMinLevel() int16 {
	return 0
}
func (this MailCliqueBangPaiJunLiang) GetMaxLevel() int16 {
	return 0
}
func (this MailCliqueBangPaiJunLiang) GetMinVIPLevel() int16 {
	return 0
}
func (this MailCliqueBangPaiJunLiang) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailCliqueBangPaiJunLiang) GetMinCreateTime() int64 {
	return 0
}
func (this MailCliqueBangPaiJunLiang) GetMaxCreateTime() int64 {
	return 0
}
func (this MailCliqueBangPaiJunLiang) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailCliqueBangPaiJunLiang) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[53]; ok {
		return attachments
	}
	return this.Attachments
}

// 您在比武场被挑战
type MailPassiveArenaWin struct {
	Name string // 挑战者
	Attachments []*Attachment
}

func (this MailPassiveArenaWin) GetMailId() int32 {
	return int32(54)
}


func (this MailPassiveArenaWin) GetSendTime() int64 {
	return 0
}

func (this MailPassiveArenaWin) GetExpireTime() int64 {
	return 0
}
func (this MailPassiveArenaWin) GetTitle() string {
	return ""
}
func (this MailPassiveArenaWin) GetPriority() int8 {
	return 0
}
func (this MailPassiveArenaWin) GetContent() string {
	return ""
}
func (this MailPassiveArenaWin) GetMinLevel() int16 {
	return 0
}
func (this MailPassiveArenaWin) GetMaxLevel() int16 {
	return 0
}
func (this MailPassiveArenaWin) GetMinVIPLevel() int16 {
	return 0
}
func (this MailPassiveArenaWin) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailPassiveArenaWin) GetMinCreateTime() int64 {
	return 0
}
func (this MailPassiveArenaWin) GetMaxCreateTime() int64 {
	return 0
}
func (this MailPassiveArenaWin) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Name))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailPassiveArenaWin) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[54]; ok {
		return attachments
	}
	return this.Attachments
}

// 您在比武场被挑战
type MailPassiveArenaLose struct {
	Name string // 挑战者
	Rank string // 排名
	Attachments []*Attachment
}

func (this MailPassiveArenaLose) GetMailId() int32 {
	return int32(55)
}


func (this MailPassiveArenaLose) GetSendTime() int64 {
	return 0
}

func (this MailPassiveArenaLose) GetExpireTime() int64 {
	return 0
}
func (this MailPassiveArenaLose) GetTitle() string {
	return ""
}
func (this MailPassiveArenaLose) GetPriority() int8 {
	return 0
}
func (this MailPassiveArenaLose) GetContent() string {
	return ""
}
func (this MailPassiveArenaLose) GetMinLevel() int16 {
	return 0
}
func (this MailPassiveArenaLose) GetMaxLevel() int16 {
	return 0
}
func (this MailPassiveArenaLose) GetMinVIPLevel() int16 {
	return 0
}
func (this MailPassiveArenaLose) GetMaxVIPLevel() int16 {
	return 0
}
func (this MailPassiveArenaLose) GetMinCreateTime() int64 {
	return 0
}
func (this MailPassiveArenaLose) GetMaxCreateTime() int64 {
	return 0
}
func (this MailPassiveArenaLose) GetParameters() string {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Name))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Rank))
	b.WriteString("]")
	return string(b.Bytes())
}

func (this MailPassiveArenaLose) GetAttachments() []*Attachment {
	if attachments, ok := g_attachments[55]; ok {
		return attachments
	}
	return this.Attachments
}

// 只用于调试
func NewMailTemplete(mailId int32) Mailer {
	switch mailId {
	case 1:
		return &MailBagFull{Func: "abc", }
	case 2:
		return &MailHeart{Who: "abc", }
	case 3:
		return &MailTestMail{P1: "abc", P2: "abc", }
	case 4:
		return &MailMultiLevel{Name: "abc", }
	case 6:
		return &MailGhostBagFull{Func: "abc", }
	case 7:
		return &MailSwordSoulBagFull{Func: "abc", }
	case 8:
		return &MailRecharge{Time1: "abc", Num: "abc", }
	case 9:
		return &MailPurchaseTips{Source: "abc", ItemName: "abc", Func: "abc", }
	case 10:
		return &MailDrawTips{ItemName: "abc", Func: "abc", }
	case 11:
		return &MailPlatformFriendAward{FriendNum: "abc", ItemNum: "abc", ItemName: "abc", }
	case 12:
		return &MailWelcomeBeta{}
	case 13:
		return &MailWarTips{Nun: "abc", }
	case 14:
		return &MailVIPHeart{}
	case 15:
		return &MailShenJiHuoDong{}
	case 17:
		return &MailHuoDongJiangLi{}
	case 18:
		return &MailVIPSignAward{}
	case 19:
		return &MailZhanLiHuoDong{}
	case 20:
		return &MailZhanLiHuoDongJiangLi{}
	case 22:
		return &MailXianZun{VipNum: "abc", }
	case 23:
		return &MailDaoJuHuoDe{ItemName: "abc", Func: "abc", }
	case 24:
		return &MailQiriXinShouLi{}
	case 25:
		return &MailDaZuoTiXing{}
	case 26:
		return &MailZhuCe{}
	case 28:
		return &MailNewbieGift{}
	case 29:
		return &MailShouChongHaoHuaLi{}
	case 30:
		return &MailYaoQingLongYue{}
	case 31:
		return &MailRedPaper{}
	case 32:
		return &MailQQVip{}
	case 33:
		return &MailQQSvip{}
	case 34:
		return &MailQQvipXuFei{}
	case 35:
		return &MailQQSvipXuFei{}
	case 36:
		return &MailTotemBagFull{Num: "abc", }
	case 37:
		return &Mailcliquesalaryowner{Name: "abc", Num: "abc", Coins: "abc", }
	case 38:
		return &Mailcliquesalarymanager{Name: "abc", Num: "abc", Coins: "abc", }
	case 39:
		return &Mailcliqueleave{Name: "abc", }
	case 40:
		return &Mailcliquebemanger{}
	case 41:
		return &Mailcliquecancelmanager{}
	case 42:
		return &Mailcliquecancelowner{Name: "abc", }
	case 43:
		return &Mailcliquechangeowner{Name: "abc", }
	case 44:
		return &Mailcliquedestory{Name: "abc", }
	case 45:
		return &Mailzongcishangxiangqifujiang{Num: "abc", }
	case 46:
		return &Mailcliqueleavebank{}
	case 47:
		return &Mailcliquejoin{Name: "abc", }
	case 48:
		return &MailDengJiLiBao{}
	case 49:
		return &MailBossTaoFa{BossName: "abc", Hurt: "abc", Exp: "abc", Coins: "abc", Fame: "abc", }
	case 50:
		return &MailBossTaoFaJingYing{BossName: "abc", Rank: "abc", Hurt: "abc", Exp: "abc", Coins: "abc", Fame: "abc", }
	case 51:
		return &MailTaoFaShengLi{CampName: "abc", Point: "abc", Hope: "abc", }
	case 52:
		return &MailCliqueBangPaiBaoXiang{}
	case 53:
		return &MailCliqueBangPaiJunLiang{}
	case 54:
		return &MailPassiveArenaWin{Name: "abc", }
	case 55:
		return &MailPassiveArenaLose{Name: "abc", Rank: "abc", }
	}
	return nil
}


