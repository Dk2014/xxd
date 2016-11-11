package player_api

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
	Info(*net.Session, *Info_In)
	Relogin(*net.Session, *Relogin_In)
	BuyPhysical(*net.Session, *BuyPhysical_In)
	SetPlayKey(*net.Session, *SetPlayKey_In)
	GetTime(*net.Session, *GetTime_In)
	FromPlatformLogin(*net.Session, *FromPlatformLogin_In)
	BuyCoins(*net.Session, *BuyCoins_In)
	GetLoginInfo(*net.Session, *GetLoginInfo_In)
	CrossLoginGameServer(*net.Session, *CrossLoginGameServer_In)
	GlobalLogin(*net.Session, *GlobalLogin_In)
	GetIngot(*net.Session, *GetIngot_In)
	SystemFame(*net.Session, *SystemFame_In)
	GetRanks(*net.Session, *GetRanks_In)
}

type OutHandler interface {
	Info(*net.Session, *Info_Out)
	Relogin(*net.Session, *Relogin_Out)
	BuyPhysical(*net.Session, *BuyPhysical_Out)
	SetPlayKey(*net.Session, *SetPlayKey_Out)
	GetTime(*net.Session, *GetTime_Out)
	FromPlatformLogin(*net.Session, *FromPlatformLogin_Out)
	BuyCoins(*net.Session, *BuyCoins_Out)
	GetLoginInfo(*net.Session, *GetLoginInfo_Out)
	CrossLoginGameServer(*net.Session, *CrossLoginGameServer_Out)
	NotifyGlobalServerInfo(*net.Session, *NotifyGlobalServerInfo_Out)
	GlobalLogin(*net.Session, *GlobalLogin_Out)
	GetIngot(*net.Session, *GetIngot_Out)
	SystemFame(*net.Session, *SystemFame_Out)
	GetRanks(*net.Session, *GetRanks_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 3:
		request := new(Info_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(Relogin_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(BuyPhysical_In)
		request.Decode(buffer)
		return request
	case 6:
		request := new(SetPlayKey_In)
		request.Decode(buffer)
		return request
	case 7:
		request := new(GetTime_In)
		request.Decode(buffer)
		return request
	case 8:
		request := new(FromPlatformLogin_In)
		request.Decode(buffer)
		return request
	case 9:
		request := new(BuyCoins_In)
		request.Decode(buffer)
		return request
	case 10:
		request := new(GetLoginInfo_In)
		request.Decode(buffer)
		return request
	case 20:
		request := new(CrossLoginGameServer_In)
		request.Decode(buffer)
		return request
	case 26:
		request := new(GlobalLogin_In)
		request.Decode(buffer)
		return request
	case 27:
		request := new(GetIngot_In)
		request.Decode(buffer)
		return request
	case 28:
		request := new(SystemFame_In)
		request.Decode(buffer)
		return request
	case 29:
		request := new(GetRanks_In)
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
	case 3:
		request := new(Info_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(Relogin_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(BuyPhysical_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(SetPlayKey_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(GetTime_Out)
		request.Decode(buffer)
		return request
	case 8:
		request := new(FromPlatformLogin_Out)
		request.Decode(buffer)
		return request
	case 9:
		request := new(BuyCoins_Out)
		request.Decode(buffer)
		return request
	case 10:
		request := new(GetLoginInfo_Out)
		request.Decode(buffer)
		return request
	case 20:
		request := new(CrossLoginGameServer_Out)
		request.Decode(buffer)
		return request
	case 25:
		request := new(NotifyGlobalServerInfo_Out)
		request.Decode(buffer)
		return request
	case 26:
		request := new(GlobalLogin_Out)
		request.Decode(buffer)
		return request
	case 27:
		request := new(GetIngot_Out)
		request.Decode(buffer)
		return request
	case 28:
		request := new(SystemFame_Out)
		request.Decode(buffer)
		return request
	case 29:
		request := new(GetRanks_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type LoginStatus int8

const (
	LOGIN_STATUS_FAILED           LoginStatus = 0
	LOGIN_STATUS_SUCCEED          LoginStatus = 1
	LOGIN_STATUS_FIRST_TIME       LoginStatus = 2
	LOGIN_STATUS_WAIT_CLOSE       LoginStatus = 3
	LOGIN_STATUS_RESTORE_SUCCEED  LoginStatus = 4
	LOGIN_STATUS_BAN              LoginStatus = 5
	LOGIN_STATUS_RELOGIN          LoginStatus = 6
	LOGIN_STATUS_INVALID_PAYTOKEN LoginStatus = 7
)

type PlatformType int8

const (
	PLATFORM_TYPE_MOBILE_IOS_WEIXIN     PlatformType = 1
	PLATFORM_TYPE_MOBILE_IOS_QQ         PlatformType = 2
	PLATFORM_TYPE_MOBILE_IOS_GUEST      PlatformType = 5
	PLATFORM_TYPE_MOBILE_ANDROID_WEIXIN PlatformType = 17
	PLATFORM_TYPE_MOBILE_ANDROID_QQ     PlatformType = 18
	PLATFORM_TYPE_MOBILE_ANDROID_GUEST  PlatformType = 21
)

type ChannelId int8

const (
	CHANNEL_ID_MOBILE            ChannelId = 0
	CHANNEL_ID_MOBILE_IOS_VN     ChannelId = 1
	CHANNEL_ID_MOBILE_ANDROID_VN ChannelId = 2
	CHANNEL_ID_MOBILE_IOS_TW     ChannelId = 3
	CHANNEL_ID_MOBILE_ANDROID_TW ChannelId = 4
)

type RankType int8

const (
	RANK_TYPE_FIGHTNUM          RankType = 0
	RANK_TYPE_LEVEL             RankType = 1
	RANK_TYPE_MAINROLE_FIGHTNUM RankType = 2
	RANK_TYPE_FAME              RankType = 4
	RANK_TYPE_BUDDY_FIGHTNUM    RankType = 5
	RANK_TYPE_GHOST_FIGHTNUM    RankType = 6
	RANK_TYPE_SWORD_SOUL_NUM    RankType = 7
	RANK_TYPE_NULL              RankType = 8
)

type Info_In struct {
}

func (this *Info_In) Process(session *net.Session) {
	g_InHandler.Info(session, this)
}

func (this *Info_In) TypeName() string {
	return "player.info.in"
}

func (this *Info_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 3
}

type Info_Out struct {
	Nickname                    []byte `json:"nickname"`
	TownId                      int16  `json:"town_id"`
	RoleId                      int8   `json:"role_id"`
	RoleLv                      int16  `json:"role_lv"`
	RoleExp                     int64  `json:"role_exp"`
	VipLevel                    int16  `json:"vip_level"`
	Fame                        int32  `json:"fame"`
	Ingot                       int64  `json:"ingot"`
	Coins                       int64  `json:"coins"`
	Fragments                   int64  `json:"fragments"`
	HeartNum                    int16  `json:"heart_num"`
	Physical                    int16  `json:"physical"`
	PhysicalBuyCount            int16  `json:"physical_buy_count"`
	CoinsBuyCount               int16  `json:"coins_buy_count"`
	BatchBoughtCoins            int16  `json:"batch_bought_coins"`
	CornucopiaCount             int16  `json:"cornucopia_count"`
	FuncKey                     int16  `json:"func_key"`
	PlayedKey                   int16  `json:"played_key"`
	TownLock                    int32  `json:"town_lock"`
	MissionKey                  int32  `json:"mission_key"`
	MissionMaxOrder             int8   `json:"mission_max_order"`
	MissionLevelMaxLock         int32  `json:"mission_level_max_lock"`
	MissionLevelAwardLock       int32  `json:"mission_level_award_lock"`
	HardLevelLock               int32  `json:"hard_level_lock"`
	QuestId                     int16  `json:"quest_id"`
	QuestState                  int8   `json:"quest_state"`
	ArenaReportNum              int16  `json:"arena_report_num"`
	NextFreeStarBoxTimestamp    int64  `json:"next_free_star_box_timestamp"`
	NextFreeMoonBoxTimestamp    int64  `json:"next_free_moon_box_timestamp"`
	NextFreeSunBoxTimestamp     int64  `json:"next_free_sun_box_timestamp"`
	NextFreeHunyuanBoxTimestamp int64  `json:"next_free_hunyuan_box_timestamp"`
	QqVipStatus                 int16  `json:"qq_vip_status"`
	ServerOpenTime              int64  `json:"server_open_time"`
}

func (this *Info_Out) Process(session *net.Session) {
	g_OutHandler.Info(session, this)
}

func (this *Info_Out) TypeName() string {
	return "player.info.out"
}

func (this *Info_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 3
}

func (this *Info_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Relogin_In struct {
}

func (this *Relogin_In) Process(session *net.Session) {
	g_InHandler.Relogin(session, this)
}

func (this *Relogin_In) TypeName() string {
	return "player.relogin.in"
}

func (this *Relogin_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 4
}

type Relogin_Out struct {
}

func (this *Relogin_Out) Process(session *net.Session) {
	g_OutHandler.Relogin(session, this)
}

func (this *Relogin_Out) TypeName() string {
	return "player.relogin.out"
}

func (this *Relogin_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 4
}

func (this *Relogin_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type BuyPhysical_In struct {
}

func (this *BuyPhysical_In) Process(session *net.Session) {
	g_InHandler.BuyPhysical(session, this)
}

func (this *BuyPhysical_In) TypeName() string {
	return "player.buy_physical.in"
}

func (this *BuyPhysical_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 5
}

type BuyPhysical_Out struct {
	PhysicalBuyCount int16 `json:"physical_buy_count"`
}

func (this *BuyPhysical_Out) Process(session *net.Session) {
	g_OutHandler.BuyPhysical(session, this)
}

func (this *BuyPhysical_Out) TypeName() string {
	return "player.buy_physical.out"
}

func (this *BuyPhysical_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 5
}

func (this *BuyPhysical_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SetPlayKey_In struct {
	Lock int16 `json:"lock"`
}

func (this *SetPlayKey_In) Process(session *net.Session) {
	g_InHandler.SetPlayKey(session, this)
}

func (this *SetPlayKey_In) TypeName() string {
	return "player.set_play_key.in"
}

func (this *SetPlayKey_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 6
}

type SetPlayKey_Out struct {
}

func (this *SetPlayKey_Out) Process(session *net.Session) {
	g_OutHandler.SetPlayKey(session, this)
}

func (this *SetPlayKey_Out) TypeName() string {
	return "player.set_play_key.out"
}

func (this *SetPlayKey_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 6
}

func (this *SetPlayKey_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetTime_In struct {
}

func (this *GetTime_In) Process(session *net.Session) {
	g_InHandler.GetTime(session, this)
}

func (this *GetTime_In) TypeName() string {
	return "player.get_time.in"
}

func (this *GetTime_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 7
}

type GetTime_Out struct {
	UnixTime int64 `json:"unix_time"`
}

func (this *GetTime_Out) Process(session *net.Session) {
	g_OutHandler.GetTime(session, this)
}

func (this *GetTime_Out) TypeName() string {
	return "player.get_time.out"
}

func (this *GetTime_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 7
}

func (this *GetTime_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type FromPlatformLogin_In struct {
	PlatformType   PlatformType `json:"platform_type"`
	ChannelId      ChannelId    `json:"channel_id"`
	Username       []byte       `json:"username"`
	Nickname       []byte       `json:"nickname"`
	RoleId         int8         `json:"role_id"`
	Hashcode       []byte       `json:"hashcode"`
	Unixtime       int64        `json:"unixtime"`
	Restore        bool         `json:"restore"`
	RecvCount      int64        `json:"recvCount"`
	Gsid           int32        `json:"gsid"`
	Openkey        []byte       `json:"openkey"`
	PayToken       []byte       `json:"pay_token"`
	Pfkey          []byte       `json:"pfkey"`
	Zoneid         int32        `json:"zoneid"`
	Pf             []byte       `json:"pf"`
	Channel        int32        `json:"channel"`
	TelecomOper    []byte       `json:"telecom_oper"`
	ClientVersion  []byte       `json:"client_version"`
	SystemHardware []byte       `json:"system_hardware"`
	Network        []byte       `json:"network"`
}

func (this *FromPlatformLogin_In) Process(session *net.Session) {
	g_InHandler.FromPlatformLogin(session, this)
}

func (this *FromPlatformLogin_In) TypeName() string {
	return "player.from_platform_login.in"
}

func (this *FromPlatformLogin_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 8
}

type FromPlatformLogin_Out struct {
	Status   LoginStatus `json:"status"`
	PlayerId int64       `json:"player_id"`
	BanTime  int64       `json:"ban_time"`
}

func (this *FromPlatformLogin_Out) Process(session *net.Session) {
	g_OutHandler.FromPlatformLogin(session, this)
}

func (this *FromPlatformLogin_Out) TypeName() string {
	return "player.from_platform_login.out"
}

func (this *FromPlatformLogin_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 8
}

func (this *FromPlatformLogin_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type BuyCoins_In struct {
	Number int16 `json:"number"`
}

func (this *BuyCoins_In) Process(session *net.Session) {
	g_InHandler.BuyCoins(session, this)
}

func (this *BuyCoins_In) TypeName() string {
	return "player.buy_coins.in"
}

func (this *BuyCoins_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 9
}

type BuyCoins_Out struct {
	Result []BuyCoins_Out_Result `json:"result"`
}

type BuyCoins_Out_Result struct {
	Ingot int64 `json:"ingot"`
	Coins int64 `json:"coins"`
	Crit  bool  `json:"crit"`
}

func (this *BuyCoins_Out) Process(session *net.Session) {
	g_OutHandler.BuyCoins(session, this)
}

func (this *BuyCoins_Out) TypeName() string {
	return "player.buy_coins.out"
}

func (this *BuyCoins_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 9
}

func (this *BuyCoins_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetLoginInfo_In struct {
}

func (this *GetLoginInfo_In) Process(session *net.Session) {
	g_InHandler.GetLoginInfo(session, this)
}

func (this *GetLoginInfo_In) TypeName() string {
	return "player.get_login_info.in"
}

func (this *GetLoginInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 10
}

type GetLoginInfo_Out struct {
	Hash []byte `json:"hash"`
	Time int64  `json:"time"`
}

func (this *GetLoginInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetLoginInfo(session, this)
}

func (this *GetLoginInfo_Out) TypeName() string {
	return "player.get_login_info.out"
}

func (this *GetLoginInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 10
}

func (this *GetLoginInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CrossLoginGameServer_In struct {
	Pid       int64  `json:"pid"`
	Openid    []byte `json:"openid"`
	Nick      []byte `json:"nick"`
	RoleId    int8   `json:"role_id"`
	RoleLevel int16  `json:"role_level"`
	Time      int64  `json:"time"`
	Hash      []byte `json:"hash"`
}

func (this *CrossLoginGameServer_In) Process(session *net.Session) {
	g_InHandler.CrossLoginGameServer(session, this)
}

func (this *CrossLoginGameServer_In) TypeName() string {
	return "player.cross_login_game_server.in"
}

func (this *CrossLoginGameServer_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 20
}

type CrossLoginGameServer_Out struct {
	Result bool `json:"result"`
}

func (this *CrossLoginGameServer_Out) Process(session *net.Session) {
	g_OutHandler.CrossLoginGameServer(session, this)
}

func (this *CrossLoginGameServer_Out) TypeName() string {
	return "player.cross_login_game_server.out"
}

func (this *CrossLoginGameServer_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 20
}

func (this *CrossLoginGameServer_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyGlobalServerInfo_Out struct {
	Gsid int32  `json:"gsid"`
	Addr []byte `json:"addr"`
}

func (this *NotifyGlobalServerInfo_Out) Process(session *net.Session) {
	g_OutHandler.NotifyGlobalServerInfo(session, this)
}

func (this *NotifyGlobalServerInfo_Out) TypeName() string {
	return "player.notify_global_server_info.out"
}

func (this *NotifyGlobalServerInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 25
}

func (this *NotifyGlobalServerInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GlobalLogin_In struct {
	Pid          int64        `json:"pid"`
	Openid       []byte       `json:"openid"`
	Nick         []byte       `json:"nick"`
	RoleId       int8         `json:"role_id"`
	RoleLevel    int16        `json:"role_level"`
	FightNum     int32        `json:"fight_num"`
	Time         int64        `json:"time"`
	Hash         []byte       `json:"hash"`
	Gsid         int32        `json:"gsid"`
	PlatformType PlatformType `json:"platform_type"`
	ChannelId    ChannelId    `json:"channel_id"`
	Username     []byte       `json:"username"`
	Openkey      []byte       `json:"openkey"`
	PayToken     []byte       `json:"pay_token"`
	Pfkey        []byte       `json:"pfkey"`
	Zoneid       int32        `json:"zoneid"`
	Pf           []byte       `json:"pf"`
}

func (this *GlobalLogin_In) Process(session *net.Session) {
	g_InHandler.GlobalLogin(session, this)
}

func (this *GlobalLogin_In) TypeName() string {
	return "player.global_login.in"
}

func (this *GlobalLogin_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 26
}

type GlobalLogin_Out struct {
	Paytoken bool `json:"paytoken"`
	Result   bool `json:"result"`
}

func (this *GlobalLogin_Out) Process(session *net.Session) {
	g_OutHandler.GlobalLogin(session, this)
}

func (this *GlobalLogin_Out) TypeName() string {
	return "player.global_login.out"
}

func (this *GlobalLogin_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 26
}

func (this *GlobalLogin_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetIngot_In struct {
}

func (this *GetIngot_In) Process(session *net.Session) {
	g_InHandler.GetIngot(session, this)
}

func (this *GetIngot_In) TypeName() string {
	return "player.get_ingot.in"
}

func (this *GetIngot_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 27
}

type GetIngot_Out struct {
	Ingot int64 `json:"ingot"`
}

func (this *GetIngot_Out) Process(session *net.Session) {
	g_OutHandler.GetIngot(session, this)
}

func (this *GetIngot_Out) TypeName() string {
	return "player.get_ingot.out"
}

func (this *GetIngot_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 27
}

func (this *GetIngot_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SystemFame_In struct {
	System int16 `json:"system"`
}

func (this *SystemFame_In) Process(session *net.Session) {
	g_InHandler.SystemFame(session, this)
}

func (this *SystemFame_In) TypeName() string {
	return "player.system_fame.in"
}

func (this *SystemFame_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 28
}

type SystemFame_Out struct {
	Fame   int32 `json:"fame"`
	System int16 `json:"system"`
}

func (this *SystemFame_Out) Process(session *net.Session) {
	g_OutHandler.SystemFame(session, this)
}

func (this *SystemFame_Out) TypeName() string {
	return "player.system_fame.out"
}

func (this *SystemFame_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 28
}

func (this *SystemFame_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetRanks_In struct {
	Flag      RankType `json:"flag"`
	PageIndex int32    `json:"page_index"`
}

func (this *GetRanks_In) Process(session *net.Session) {
	g_InHandler.GetRanks(session, this)
}

func (this *GetRanks_In) TypeName() string {
	return "player.get_ranks.in"
}

func (this *GetRanks_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 29
}

type GetRanks_Out struct {
	Ranks   []GetRanks_Out_Ranks `json:"ranks"`
	HasNext bool                 `json:"has_next"`
}

type GetRanks_Out_Ranks struct {
	Pid      int64                       `json:"pid"`
	Nickname []byte                      `json:"nickname"`
	Rank     int32                       `json:"rank"`
	Values   []GetRanks_Out_Ranks_Values `json:"values"`
}

type GetRanks_Out_Ranks_Values struct {
	Flag  RankType `json:"flag"`
	Value int64    `json:"value"`
}

func (this *GetRanks_Out) Process(session *net.Session) {
	g_OutHandler.GetRanks(session, this)
}

func (this *GetRanks_Out) TypeName() string {
	return "player.get_ranks.out"
}

func (this *GetRanks_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 0, 29
}

func (this *GetRanks_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *Info_In) Decode(buffer *net.Buffer) {
}

func (this *Info_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(3)
}

func (this *Info_In) ByteSize() int {
	size := 2
	return size
}

func (this *Info_Out) Decode(buffer *net.Buffer) {
	this.Nickname = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.TownId = int16(buffer.ReadUint16LE())
	this.RoleId = int8(buffer.ReadUint8())
	this.RoleLv = int16(buffer.ReadUint16LE())
	this.RoleExp = int64(buffer.ReadUint64LE())
	this.VipLevel = int16(buffer.ReadUint16LE())
	this.Fame = int32(buffer.ReadUint32LE())
	this.Ingot = int64(buffer.ReadUint64LE())
	this.Coins = int64(buffer.ReadUint64LE())
	this.Fragments = int64(buffer.ReadUint64LE())
	this.HeartNum = int16(buffer.ReadUint16LE())
	this.Physical = int16(buffer.ReadUint16LE())
	this.PhysicalBuyCount = int16(buffer.ReadUint16LE())
	this.CoinsBuyCount = int16(buffer.ReadUint16LE())
	this.BatchBoughtCoins = int16(buffer.ReadUint16LE())
	this.CornucopiaCount = int16(buffer.ReadUint16LE())
	this.FuncKey = int16(buffer.ReadUint16LE())
	this.PlayedKey = int16(buffer.ReadUint16LE())
	this.TownLock = int32(buffer.ReadUint32LE())
	this.MissionKey = int32(buffer.ReadUint32LE())
	this.MissionMaxOrder = int8(buffer.ReadUint8())
	this.MissionLevelMaxLock = int32(buffer.ReadUint32LE())
	this.MissionLevelAwardLock = int32(buffer.ReadUint32LE())
	this.HardLevelLock = int32(buffer.ReadUint32LE())
	this.QuestId = int16(buffer.ReadUint16LE())
	this.QuestState = int8(buffer.ReadUint8())
	this.ArenaReportNum = int16(buffer.ReadUint16LE())
	this.NextFreeStarBoxTimestamp = int64(buffer.ReadUint64LE())
	this.NextFreeMoonBoxTimestamp = int64(buffer.ReadUint64LE())
	this.NextFreeSunBoxTimestamp = int64(buffer.ReadUint64LE())
	this.NextFreeHunyuanBoxTimestamp = int64(buffer.ReadUint64LE())
	this.QqVipStatus = int16(buffer.ReadUint16LE())
	this.ServerOpenTime = int64(buffer.ReadUint64LE())
}

func (this *Info_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(3)
	buffer.WriteUint16LE(uint16(len(this.Nickname)))
	buffer.WriteBytes(this.Nickname)
	buffer.WriteUint16LE(uint16(this.TownId))
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.RoleLv))
	buffer.WriteUint64LE(uint64(this.RoleExp))
	buffer.WriteUint16LE(uint16(this.VipLevel))
	buffer.WriteUint32LE(uint32(this.Fame))
	buffer.WriteUint64LE(uint64(this.Ingot))
	buffer.WriteUint64LE(uint64(this.Coins))
	buffer.WriteUint64LE(uint64(this.Fragments))
	buffer.WriteUint16LE(uint16(this.HeartNum))
	buffer.WriteUint16LE(uint16(this.Physical))
	buffer.WriteUint16LE(uint16(this.PhysicalBuyCount))
	buffer.WriteUint16LE(uint16(this.CoinsBuyCount))
	buffer.WriteUint16LE(uint16(this.BatchBoughtCoins))
	buffer.WriteUint16LE(uint16(this.CornucopiaCount))
	buffer.WriteUint16LE(uint16(this.FuncKey))
	buffer.WriteUint16LE(uint16(this.PlayedKey))
	buffer.WriteUint32LE(uint32(this.TownLock))
	buffer.WriteUint32LE(uint32(this.MissionKey))
	buffer.WriteUint8(uint8(this.MissionMaxOrder))
	buffer.WriteUint32LE(uint32(this.MissionLevelMaxLock))
	buffer.WriteUint32LE(uint32(this.MissionLevelAwardLock))
	buffer.WriteUint32LE(uint32(this.HardLevelLock))
	buffer.WriteUint16LE(uint16(this.QuestId))
	buffer.WriteUint8(uint8(this.QuestState))
	buffer.WriteUint16LE(uint16(this.ArenaReportNum))
	buffer.WriteUint64LE(uint64(this.NextFreeStarBoxTimestamp))
	buffer.WriteUint64LE(uint64(this.NextFreeMoonBoxTimestamp))
	buffer.WriteUint64LE(uint64(this.NextFreeSunBoxTimestamp))
	buffer.WriteUint64LE(uint64(this.NextFreeHunyuanBoxTimestamp))
	buffer.WriteUint16LE(uint16(this.QqVipStatus))
	buffer.WriteUint64LE(uint64(this.ServerOpenTime))
}

func (this *Info_Out) ByteSize() int {
	size := 131
	size += len(this.Nickname)
	return size
}

func (this *Relogin_In) Decode(buffer *net.Buffer) {
}

func (this *Relogin_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(4)
}

func (this *Relogin_In) ByteSize() int {
	size := 2
	return size
}

func (this *Relogin_Out) Decode(buffer *net.Buffer) {
}

func (this *Relogin_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(4)
}

func (this *Relogin_Out) ByteSize() int {
	size := 2
	return size
}

func (this *BuyPhysical_In) Decode(buffer *net.Buffer) {
}

func (this *BuyPhysical_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(5)
}

func (this *BuyPhysical_In) ByteSize() int {
	size := 2
	return size
}

func (this *BuyPhysical_Out) Decode(buffer *net.Buffer) {
	this.PhysicalBuyCount = int16(buffer.ReadUint16LE())
}

func (this *BuyPhysical_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(5)
	buffer.WriteUint16LE(uint16(this.PhysicalBuyCount))
}

func (this *BuyPhysical_Out) ByteSize() int {
	size := 4
	return size
}

func (this *SetPlayKey_In) Decode(buffer *net.Buffer) {
	this.Lock = int16(buffer.ReadUint16LE())
}

func (this *SetPlayKey_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(6)
	buffer.WriteUint16LE(uint16(this.Lock))
}

func (this *SetPlayKey_In) ByteSize() int {
	size := 4
	return size
}

func (this *SetPlayKey_Out) Decode(buffer *net.Buffer) {
}

func (this *SetPlayKey_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(6)
}

func (this *SetPlayKey_Out) ByteSize() int {
	size := 2
	return size
}

func (this *GetTime_In) Decode(buffer *net.Buffer) {
}

func (this *GetTime_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(7)
}

func (this *GetTime_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetTime_Out) Decode(buffer *net.Buffer) {
	this.UnixTime = int64(buffer.ReadUint64LE())
}

func (this *GetTime_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(7)
	buffer.WriteUint64LE(uint64(this.UnixTime))
}

func (this *GetTime_Out) ByteSize() int {
	size := 10
	return size
}

func (this *FromPlatformLogin_In) Decode(buffer *net.Buffer) {
	this.PlatformType = PlatformType(buffer.ReadUint8())
	this.ChannelId = ChannelId(buffer.ReadUint8())
	this.Username = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Nickname = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.RoleId = int8(buffer.ReadUint8())
	this.Hashcode = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Unixtime = int64(buffer.ReadUint64LE())
	this.Restore = buffer.ReadUint8() == 1
	this.RecvCount = int64(buffer.ReadUint64LE())
	this.Gsid = int32(buffer.ReadUint32LE())
	this.Openkey = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.PayToken = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Pfkey = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Zoneid = int32(buffer.ReadUint32LE())
	this.Pf = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Channel = int32(buffer.ReadUint32LE())
	this.TelecomOper = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.ClientVersion = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.SystemHardware = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Network = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *FromPlatformLogin_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(8)
	buffer.WriteUint8(uint8(this.PlatformType))
	buffer.WriteUint8(uint8(this.ChannelId))
	buffer.WriteUint16LE(uint16(len(this.Username)))
	buffer.WriteBytes(this.Username)
	buffer.WriteUint16LE(uint16(len(this.Nickname)))
	buffer.WriteBytes(this.Nickname)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(len(this.Hashcode)))
	buffer.WriteBytes(this.Hashcode)
	buffer.WriteUint64LE(uint64(this.Unixtime))
	if this.Restore {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint64LE(uint64(this.RecvCount))
	buffer.WriteUint32LE(uint32(this.Gsid))
	buffer.WriteUint16LE(uint16(len(this.Openkey)))
	buffer.WriteBytes(this.Openkey)
	buffer.WriteUint16LE(uint16(len(this.PayToken)))
	buffer.WriteBytes(this.PayToken)
	buffer.WriteUint16LE(uint16(len(this.Pfkey)))
	buffer.WriteBytes(this.Pfkey)
	buffer.WriteUint32LE(uint32(this.Zoneid))
	buffer.WriteUint16LE(uint16(len(this.Pf)))
	buffer.WriteBytes(this.Pf)
	buffer.WriteUint32LE(uint32(this.Channel))
	buffer.WriteUint16LE(uint16(len(this.TelecomOper)))
	buffer.WriteBytes(this.TelecomOper)
	buffer.WriteUint16LE(uint16(len(this.ClientVersion)))
	buffer.WriteBytes(this.ClientVersion)
	buffer.WriteUint16LE(uint16(len(this.SystemHardware)))
	buffer.WriteBytes(this.SystemHardware)
	buffer.WriteUint16LE(uint16(len(this.Network)))
	buffer.WriteBytes(this.Network)
}

func (this *FromPlatformLogin_In) ByteSize() int {
	size := 56
	size += len(this.Username)
	size += len(this.Nickname)
	size += len(this.Hashcode)
	size += len(this.Openkey)
	size += len(this.PayToken)
	size += len(this.Pfkey)
	size += len(this.Pf)
	size += len(this.TelecomOper)
	size += len(this.ClientVersion)
	size += len(this.SystemHardware)
	size += len(this.Network)
	return size
}

func (this *FromPlatformLogin_Out) Decode(buffer *net.Buffer) {
	this.Status = LoginStatus(buffer.ReadUint8())
	this.PlayerId = int64(buffer.ReadUint64LE())
	this.BanTime = int64(buffer.ReadUint64LE())
}

func (this *FromPlatformLogin_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(8)
	buffer.WriteUint8(uint8(this.Status))
	buffer.WriteUint64LE(uint64(this.PlayerId))
	buffer.WriteUint64LE(uint64(this.BanTime))
}

func (this *FromPlatformLogin_Out) ByteSize() int {
	size := 19
	return size
}

func (this *BuyCoins_In) Decode(buffer *net.Buffer) {
	this.Number = int16(buffer.ReadUint16LE())
}

func (this *BuyCoins_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(9)
	buffer.WriteUint16LE(uint16(this.Number))
}

func (this *BuyCoins_In) ByteSize() int {
	size := 4
	return size
}

func (this *BuyCoins_Out) Decode(buffer *net.Buffer) {
	this.Result = make([]BuyCoins_Out_Result, buffer.ReadUint8())
	for i := 0; i < len(this.Result); i++ {
		this.Result[i].Decode(buffer)
	}
}

func (this *BuyCoins_Out_Result) Decode(buffer *net.Buffer) {
	this.Ingot = int64(buffer.ReadUint64LE())
	this.Coins = int64(buffer.ReadUint64LE())
	this.Crit = buffer.ReadUint8() == 1
}

func (this *BuyCoins_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(9)
	buffer.WriteUint8(uint8(len(this.Result)))
	for i := 0; i < len(this.Result); i++ {
		this.Result[i].Encode(buffer)
	}
}

func (this *BuyCoins_Out_Result) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Ingot))
	buffer.WriteUint64LE(uint64(this.Coins))
	if this.Crit {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *BuyCoins_Out) ByteSize() int {
	size := 3
	size += len(this.Result) * 17
	return size
}

func (this *GetLoginInfo_In) Decode(buffer *net.Buffer) {
}

func (this *GetLoginInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(10)
}

func (this *GetLoginInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetLoginInfo_Out) Decode(buffer *net.Buffer) {
	this.Hash = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Time = int64(buffer.ReadUint64LE())
}

func (this *GetLoginInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(10)
	buffer.WriteUint16LE(uint16(len(this.Hash)))
	buffer.WriteBytes(this.Hash)
	buffer.WriteUint64LE(uint64(this.Time))
}

func (this *GetLoginInfo_Out) ByteSize() int {
	size := 12
	size += len(this.Hash)
	return size
}

func (this *CrossLoginGameServer_In) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.Openid = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Nick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.RoleId = int8(buffer.ReadUint8())
	this.RoleLevel = int16(buffer.ReadUint16LE())
	this.Time = int64(buffer.ReadUint64LE())
	this.Hash = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *CrossLoginGameServer_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(20)
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Openid)))
	buffer.WriteBytes(this.Openid)
	buffer.WriteUint16LE(uint16(len(this.Nick)))
	buffer.WriteBytes(this.Nick)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.RoleLevel))
	buffer.WriteUint64LE(uint64(this.Time))
	buffer.WriteUint16LE(uint16(len(this.Hash)))
	buffer.WriteBytes(this.Hash)
}

func (this *CrossLoginGameServer_In) ByteSize() int {
	size := 27
	size += len(this.Openid)
	size += len(this.Nick)
	size += len(this.Hash)
	return size
}

func (this *CrossLoginGameServer_Out) Decode(buffer *net.Buffer) {
	this.Result = buffer.ReadUint8() == 1
}

func (this *CrossLoginGameServer_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(20)
	if this.Result {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *CrossLoginGameServer_Out) ByteSize() int {
	size := 3
	return size
}

func (this *NotifyGlobalServerInfo_Out) Decode(buffer *net.Buffer) {
	this.Gsid = int32(buffer.ReadUint32LE())
	this.Addr = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *NotifyGlobalServerInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(25)
	buffer.WriteUint32LE(uint32(this.Gsid))
	buffer.WriteUint16LE(uint16(len(this.Addr)))
	buffer.WriteBytes(this.Addr)
}

func (this *NotifyGlobalServerInfo_Out) ByteSize() int {
	size := 8
	size += len(this.Addr)
	return size
}

func (this *GlobalLogin_In) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.Openid = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Nick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.RoleId = int8(buffer.ReadUint8())
	this.RoleLevel = int16(buffer.ReadUint16LE())
	this.FightNum = int32(buffer.ReadUint32LE())
	this.Time = int64(buffer.ReadUint64LE())
	this.Hash = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Gsid = int32(buffer.ReadUint32LE())
	this.PlatformType = PlatformType(buffer.ReadUint8())
	this.ChannelId = ChannelId(buffer.ReadUint8())
	this.Username = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Openkey = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.PayToken = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Pfkey = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Zoneid = int32(buffer.ReadUint32LE())
	this.Pf = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *GlobalLogin_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(26)
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Openid)))
	buffer.WriteBytes(this.Openid)
	buffer.WriteUint16LE(uint16(len(this.Nick)))
	buffer.WriteBytes(this.Nick)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.RoleLevel))
	buffer.WriteUint32LE(uint32(this.FightNum))
	buffer.WriteUint64LE(uint64(this.Time))
	buffer.WriteUint16LE(uint16(len(this.Hash)))
	buffer.WriteBytes(this.Hash)
	buffer.WriteUint32LE(uint32(this.Gsid))
	buffer.WriteUint8(uint8(this.PlatformType))
	buffer.WriteUint8(uint8(this.ChannelId))
	buffer.WriteUint16LE(uint16(len(this.Username)))
	buffer.WriteBytes(this.Username)
	buffer.WriteUint16LE(uint16(len(this.Openkey)))
	buffer.WriteBytes(this.Openkey)
	buffer.WriteUint16LE(uint16(len(this.PayToken)))
	buffer.WriteBytes(this.PayToken)
	buffer.WriteUint16LE(uint16(len(this.Pfkey)))
	buffer.WriteBytes(this.Pfkey)
	buffer.WriteUint32LE(uint32(this.Zoneid))
	buffer.WriteUint16LE(uint16(len(this.Pf)))
	buffer.WriteBytes(this.Pf)
}

func (this *GlobalLogin_In) ByteSize() int {
	size := 51
	size += len(this.Openid)
	size += len(this.Nick)
	size += len(this.Hash)
	size += len(this.Username)
	size += len(this.Openkey)
	size += len(this.PayToken)
	size += len(this.Pfkey)
	size += len(this.Pf)
	return size
}

func (this *GlobalLogin_Out) Decode(buffer *net.Buffer) {
	this.Paytoken = buffer.ReadUint8() == 1
	this.Result = buffer.ReadUint8() == 1
}

func (this *GlobalLogin_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(26)
	if this.Paytoken {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	if this.Result {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *GlobalLogin_Out) ByteSize() int {
	size := 4
	return size
}

func (this *GetIngot_In) Decode(buffer *net.Buffer) {
}

func (this *GetIngot_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(27)
}

func (this *GetIngot_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetIngot_Out) Decode(buffer *net.Buffer) {
	this.Ingot = int64(buffer.ReadUint64LE())
}

func (this *GetIngot_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(27)
	buffer.WriteUint64LE(uint64(this.Ingot))
}

func (this *GetIngot_Out) ByteSize() int {
	size := 10
	return size
}

func (this *SystemFame_In) Decode(buffer *net.Buffer) {
	this.System = int16(buffer.ReadUint16LE())
}

func (this *SystemFame_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(28)
	buffer.WriteUint16LE(uint16(this.System))
}

func (this *SystemFame_In) ByteSize() int {
	size := 4
	return size
}

func (this *SystemFame_Out) Decode(buffer *net.Buffer) {
	this.Fame = int32(buffer.ReadUint32LE())
	this.System = int16(buffer.ReadUint16LE())
}

func (this *SystemFame_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(28)
	buffer.WriteUint32LE(uint32(this.Fame))
	buffer.WriteUint16LE(uint16(this.System))
}

func (this *SystemFame_Out) ByteSize() int {
	size := 8
	return size
}

func (this *GetRanks_In) Decode(buffer *net.Buffer) {
	this.Flag = RankType(buffer.ReadUint8())
	this.PageIndex = int32(buffer.ReadUint32LE())
}

func (this *GetRanks_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(29)
	buffer.WriteUint8(uint8(this.Flag))
	buffer.WriteUint32LE(uint32(this.PageIndex))
}

func (this *GetRanks_In) ByteSize() int {
	size := 7
	return size
}

func (this *GetRanks_Out) Decode(buffer *net.Buffer) {
	this.Ranks = make([]GetRanks_Out_Ranks, buffer.ReadUint8())
	for i := 0; i < len(this.Ranks); i++ {
		this.Ranks[i].Decode(buffer)
	}
	this.HasNext = buffer.ReadUint8() == 1
}

func (this *GetRanks_Out_Ranks) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.Nickname = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Rank = int32(buffer.ReadUint32LE())
	this.Values = make([]GetRanks_Out_Ranks_Values, buffer.ReadUint8())
	for i := 0; i < len(this.Values); i++ {
		this.Values[i].Decode(buffer)
	}
}

func (this *GetRanks_Out_Ranks_Values) Decode(buffer *net.Buffer) {
	this.Flag = RankType(buffer.ReadUint8())
	this.Value = int64(buffer.ReadUint64LE())
}

func (this *GetRanks_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(0)
	buffer.WriteUint8(29)
	buffer.WriteUint8(uint8(len(this.Ranks)))
	for i := 0; i < len(this.Ranks); i++ {
		this.Ranks[i].Encode(buffer)
	}
	if this.HasNext {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *GetRanks_Out_Ranks) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Nickname)))
	buffer.WriteBytes(this.Nickname)
	buffer.WriteUint32LE(uint32(this.Rank))
	buffer.WriteUint8(uint8(len(this.Values)))
	for i := 0; i < len(this.Values); i++ {
		this.Values[i].Encode(buffer)
	}
}

func (this *GetRanks_Out_Ranks_Values) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Flag))
	buffer.WriteUint64LE(uint64(this.Value))
}

func (this *GetRanks_Out) ByteSize() int {
	size := 4
	for i := 0; i < len(this.Ranks); i++ {
		size += this.Ranks[i].ByteSize()
	}
	return size
}

func (this *GetRanks_Out_Ranks) ByteSize() int {
	size := 15
	size += len(this.Nickname)
	size += len(this.Values) * 9
	return size
}
