package channel_dat

import (
	"strconv"
	"bytes"
	
)

const( 
	MESSAGE_TPL_DRAWSWORDSOUL = 1
	MESSAGE_TPL_FATEBOXEQUIPMENT = 2
	MESSAGE_TPL_FATEBOXGHOSTFRAME = 3
	MESSAGE_TPL_RAINBOWLEVELGHOST = 4
	MESSAGE_TPL_COMPOSEGHOST = 5
	MESSAGE_TPL_CALLTOTEM = 6
	MESSAGE_TPL_FOUNDCLIQUE = 8
	MESSAGE_TPL_CLIQUERECRUITMENT = 9
	MESSAGE_TPL_CLIQUEASSIGNMANGER = 10
	MESSAGE_TPL_CLIQUEASSIGNOWNER = 11
	MESSAGE_TPL_CLIQUEELECTOWNER = 12
	MESSAGE_TPL_CLIQUEMEMBERLEAVE = 13
	MESSAGE_TPL_CLIQUEKICKMEMBER = 14
	MESSAGE_TPL_CLIQUENEWANNC = 15
	MESSAGE_TPL_COMMONCHAT = 16
	MESSAGE_TPL_CLIQUEFIREMANGER = 17
	MESSAGE_TPL_CLIQUENEWMEMBER = 18
	MESSAGE_TPL_CLIQUEBUILDINGLEVELUP = 19
	MESSAGE_TPL_CLIQUEBOATHIJACKED = 20
	MESSAGE_TPL_CLIQUEBOATRECOVER = 21
	MESSAGE_TPL_BOATRECOVERED = 22
	MESSAGE_TPL_BOATRECOVEREDBYHERO = 23
	MESSAGE_TPL_BOATHIJACKFINISHED = 24
	MESSAGE_TPL_BOATESCORTFINISHED = 25
	MESSAGE_TPL_BOATHIJACKINGFINISHED = 26
	MESSAGE_TPL_BOATHIJACKING = 27
	MESSAGE_TPL_DESPAIRLANDCLEAR = 28
	MESSAGE_TPL_DESPAIRLANDPERFECTCLEAR = 29
	MESSAGE_TPL_CLIQUESTORESEND = 30
	MESSAGE_TPL_DESPAIRLANDBOSSKILL = 31
)
const( 
	WORLD_CHAT_CD_TIME = 5
	WORLD_CHAT_COST = 5
	WORLD_CHAT_MAX_CONTENT_LEN = 60
	WORLD_CHAT_DAILY_FREE_NUM = 5
	WORLD_CHAT_SERVER_OPEN_LEVEL = 14
	WORLD_CHAT_CLIENT_BUFF_SIZE = 50
	WORLD_CHAT_CLIENT_PAGE_SIZE = 20
	WORLD_CHAT_CLIENT_BUFF_EXPIRE_TIME = 600
)
const( 
	PARAM_TYPE_STRING = 1
	PARAM_TYPE_ITEM = 2
	PARAM_TYPE_PLAYER = 3
	PARAM_TYPE_CLIQUE = 4
	PARAM_TYPE_CLIQUE_BOAT = 5
)






type MessageDrawSwordSoul struct {
	Player ParamPlayer // 玩家
	Item ParamItem // 道具
}

func (this MessageDrawSwordSoul) GetTplId() int16 {
	return 1
}

func (this MessageDrawSwordSoul) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Player.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Item.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageFateBoxEquipment struct {
	Player ParamPlayer // 玩家
	Item ParamItem // 道具
}

func (this MessageFateBoxEquipment) GetTplId() int16 {
	return 2
}

func (this MessageFateBoxEquipment) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Player.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Item.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageFateBoxGhostFrame struct {
	Player ParamPlayer // 玩家
	Item ParamItem // 道具
}

func (this MessageFateBoxGhostFrame) GetTplId() int16 {
	return 3
}

func (this MessageFateBoxGhostFrame) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Player.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Item.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageRainbowLevelGhost struct {
	Player ParamPlayer // 玩家
	Level ParamString // 关卡
	Item ParamItem // 道具
}

func (this MessageRainbowLevelGhost) GetTplId() int16 {
	return 4
}

func (this MessageRainbowLevelGhost) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Player.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Level.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Item.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageComposeGhost struct {
	Player ParamPlayer // 玩家
	Item ParamItem // 道具
}

func (this MessageComposeGhost) GetTplId() int16 {
	return 5
}

func (this MessageComposeGhost) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Player.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Item.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageCallTotem struct {
	Player ParamPlayer // 玩家
	Item ParamItem // 道具
}

func (this MessageCallTotem) GetTplId() int16 {
	return 6
}

func (this MessageCallTotem) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Player.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Item.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageFoundClique struct {
	Player ParamPlayer // 玩家
	Clique ParamClique // 帮派
	DummyLink ParamClique // 点击申请
}

func (this MessageFoundClique) GetTplId() int16 {
	return 8
}

func (this MessageFoundClique) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Player.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Clique.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.DummyLink.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageCliqueRecruitment struct {
	Clique ParamClique // 帮派
	DummyLink ParamClique // 点击申请
}

func (this MessageCliqueRecruitment) GetTplId() int16 {
	return 9
}

func (this MessageCliqueRecruitment) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Clique.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.DummyLink.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageCliqueAssignManger struct {
	Player ParamPlayer // 玩家
}

func (this MessageCliqueAssignManger) GetTplId() int16 {
	return 10
}

func (this MessageCliqueAssignManger) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Player.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageCliqueAssignOwner struct {
	OldOwner ParamPlayer // 原帮主
	NewOwner ParamPlayer // 新帮主
}

func (this MessageCliqueAssignOwner) GetTplId() int16 {
	return 11
}

func (this MessageCliqueAssignOwner) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.OldOwner.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.NewOwner.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageCliqueElectOwner struct {
	NewOwner ParamPlayer // 新班主
	Player2 ParamPlayer // 原帮主
}

func (this MessageCliqueElectOwner) GetTplId() int16 {
	return 12
}

func (this MessageCliqueElectOwner) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.NewOwner.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Player2.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageCliqueMemberLeave struct {
	Player ParamPlayer // 玩家
}

func (this MessageCliqueMemberLeave) GetTplId() int16 {
	return 13
}

func (this MessageCliqueMemberLeave) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Player.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageCliqueKickMember struct {
	Manger ParamPlayer // 管理员
	Member ParamPlayer // 成员
}

func (this MessageCliqueKickMember) GetTplId() int16 {
	return 14
}

func (this MessageCliqueKickMember) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Manger.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Member.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageCliqueNewAnnc struct {
}

func (this MessageCliqueNewAnnc) GetTplId() int16 {
	return 15
}

func (this MessageCliqueNewAnnc) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString("]")
	return b.Bytes()
}

type MessageCommonChat struct {
	Content ParamString // 内容
}

func (this MessageCommonChat) GetTplId() int16 {
	return 16
}

func (this MessageCommonChat) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Content.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageCliqueFireManger struct {
	Manger ParamPlayer // 管理员
}

func (this MessageCliqueFireManger) GetTplId() int16 {
	return 17
}

func (this MessageCliqueFireManger) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Manger.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageCliqueNewMember struct {
	NewMember ParamPlayer // 新成员
}

func (this MessageCliqueNewMember) GetTplId() int16 {
	return 18
}

func (this MessageCliqueNewMember) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.NewMember.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageCliqueBuildingLevelUP struct {
	BuildType ParamString // 帮派建筑名称
	CliqueLevel ParamString // 帮派等级
}

func (this MessageCliqueBuildingLevelUP) GetTplId() int16 {
	return 19
}

func (this MessageCliqueBuildingLevelUP) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.BuildType.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.CliqueLevel.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageCliqueBoatHijacked struct {
	Clique ParamClique // 敌对帮派
	Hijacker ParamPlayer // 劫持者
	BoatOwner ParamPlayer // 我帮玩家
	DummyLink ParamCliqueBoat // 镖船
}

func (this MessageCliqueBoatHijacked) GetTplId() int16 {
	return 20
}

func (this MessageCliqueBoatHijacked) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Clique.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Hijacker.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.BoatOwner.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.DummyLink.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageCliqueBoatRecover struct {
	Fighter ParamPlayer // 夺回者
	BoatOwner ParamPlayer // 镖船主
}

func (this MessageCliqueBoatRecover) GetTplId() int16 {
	return 21
}

func (this MessageCliqueBoatRecover) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Fighter.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.BoatOwner.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageBoatRecovered struct {
	Fighter ParamPlayer // 夺回者
}

func (this MessageBoatRecovered) GetTplId() int16 {
	return 22
}

func (this MessageBoatRecovered) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Fighter.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageBoatRecoveredByHero struct {
	Fighter ParamPlayer // 夺回者
}

func (this MessageBoatRecoveredByHero) GetTplId() int16 {
	return 23
}

func (this MessageBoatRecoveredByHero) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Fighter.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageBoatHijackFinished struct {
	Coins ParamString // 铜钱
	Fame ParamString // 声望
	Contrib ParamString // 贡献
}

func (this MessageBoatHijackFinished) GetTplId() int16 {
	return 24
}

func (this MessageBoatHijackFinished) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Coins.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Fame.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Contrib.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageBoatEscortFinished struct {
	Coins ParamString // 铜钱
	Fame ParamString // 声望
	Contrib ParamString // 贡献
}

func (this MessageBoatEscortFinished) GetTplId() int16 {
	return 25
}

func (this MessageBoatEscortFinished) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Coins.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Fame.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Contrib.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageBoatHijackingFinished struct {
	Hijacker ParamPlayer // 劫持者
	Coins ParamString // 铜钱
	Fame ParamString // 声望
	Contrib ParamString // 贡献
}

func (this MessageBoatHijackingFinished) GetTplId() int16 {
	return 26
}

func (this MessageBoatHijackingFinished) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Hijacker.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Coins.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Fame.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Contrib.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageBoatHijacking struct {
	HijackerClique ParamClique // 劫持者帮派
	Hijacker ParamPlayer // 劫持者
}

func (this MessageBoatHijacking) GetTplId() int16 {
	return 27
}

func (this MessageBoatHijacking) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.HijackerClique.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Hijacker.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageDespairLandClear struct {
	Name ParamPlayer // 玩家名称
	MissionName ParamString // 区域名称
	LevelName ParamString // 关卡名称
}

func (this MessageDespairLandClear) GetTplId() int16 {
	return 28
}

func (this MessageDespairLandClear) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Name.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.MissionName.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.LevelName.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageDespairLandPerfectClear struct {
	Name ParamPlayer // 玩家名称
	MissionName ParamString // 区域名称
	LevelName ParamString // 关卡名称
}

func (this MessageDespairLandPerfectClear) GetTplId() int16 {
	return 29
}

func (this MessageDespairLandPerfectClear) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Name.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.MissionName.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.LevelName.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageCliqueStoreSend struct {
	Job ParamString // 帮派职位
	Name ParamPlayer // 发放者
	ItemName ParamString // 战备类型
}

func (this MessageCliqueStoreSend) GetTplId() int16 {
	return 30
}

func (this MessageCliqueStoreSend) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Job.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Name.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.ItemName.ToString()))
	b.WriteString("]")
	return b.Bytes()
}

type MessageDespairLandBossKill struct {
	Camp ParamString // 阵营
	Boss ParamString // 首领
}

func (this MessageDespairLandBossKill) GetTplId() int16 {
	return 31
}

func (this MessageDespairLandBossKill) GetParameters() []byte {
	b := new(bytes.Buffer)
	b.WriteString("[")
	b.WriteString(strconv.Quote(this.Camp.ToString()))
	b.WriteString(",")
	b.WriteString(strconv.Quote(this.Boss.ToString()))
	b.WriteString("]")
	return b.Bytes()
}


