package xdgm_server

import (
	"game_server/api/protocol/role_api"
)

// 应答
type XDGM_RSP struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type XDGM_MAIL struct {
	MailTitle   string     `json:"mailtitle"`   // 邮件标题
	MailContent string     `json:"mailcontent"` // 邮件内容
	SendTime    int64      `json:"sendtime"`    // 邮件发送时间
	ItemDetail  []MailItem `json:"itemdetail"`
}

type MailItem struct {
	ItemId  int16 `json:"itemid"`  // 邮件赠送道具1ID
	ItemNum int64 `json:"itemnum"` // 邮件赠送道具1数量
}

// 玩家信息查询请求
type XDGM_GET_PLAYER_INFO_REQ struct {
	Players string `json:"players"`
}

type XDGM_GET_PLAYER_INFO_DATA struct {
	Players map[string]XDGM_PLAYER_INFO `json:"players"`
}

type XDGM_PLAYER_INFO struct {
	RoleName      string                      `json:"rolename"`      // 角色名称
	Pid           int64                       `json:"pid"`           // 玩家ID
	Level         int16                       `json:"level"`         // 当前等级
	Vip           int16                       `json:"vip"`           // 当前VIP等级
	Exp           int64                       `json:"exp"`           // 当前经验
	Coin          int64                       `json:"coin"`          // 当前铜钱
	Ingot         int64                       `json:"ingot"`         // 当前元宝数量
	Physical      int16                       `json:"physical"`      // 当前体力值
	RegisterTime  int64                       `json:"registertime"`  // 注册时间
	LastLoginTime int64                       `json:"lastlogintime"` // 玩家最后登录时间
	Mails         []XDGM_MAIL                 `json:"mails"`         //邮件信息
	Rank          int32                       `json:"rank"`          //比武场排名
	PlayerInfo    []role_api.PlayerInfo_Roles `json:"playerinfo"`    //角色信息
}

// 帮派信息查询请求
type XDGM_CLIQUE_INFO_REQ struct {
	ServerId int   `json:"server_id"` //互动服ID
	Limit    int16 `json:"limit"`     //结果集合size
	Offset   int16 `json:"offset"`    //跳过前X个
}

// 角色封号操作
type XDGM_LOCK_PLAYER_REQ struct {
	Manager   string `json:"manager"`    // 操作者
	Players   string `json:"players"`    // 角色ID
	BlockTime int64  `json:"block_time"` // 封禁时间-1:永久;0:解封;>0:封禁时长
}

// 角色禁言操作
type XDGM_GAG_PLAYER_REQ struct {
	Manager   string `json:"manager"`    // 操作者
	Players   string `json:"players"`    // 角色ID
	BlockTime int64  `json:"block_time"` // 封禁时间-1:永久;0:解封;>0:封禁时长
}

type XDGM_GET_NORMAL_EVENT_INFO_REQ struct {
	ServerId int   `json:"server_id"` //游戏服ID
	Limit    int16 `json:"limit"`
	Offset   int16 `json:"offset"`
}

type XDGM_GET_JSON_EVENT_INFO_REQ struct {
	ServerId int   `json:"server_id"` //游戏服ID
	Limit    int16 `json:"limit"`
	Offset   int16 `json:"offset"`
}

type XDGM_EVENT_AWARD_INFO_REQ struct {
	ServerId  int   `json:"server_id"`  //游戏服ID
	EventSign int16 `json:"event_sign"` //正常活动传event_id,json配置的传type
	Page      int32 `json:"page"`       // 活动期数，正常活动则传0
}

type XDGM_UPDATE_EVENTS_INFO_REQ struct {
	ServerId   int    `json:"server_id"`
	EventsInfo string `json:"events_info"`
}

// 单人邮件发送
type XDGM_SEND_MAIL_REQ struct {
	Manager string `json:"manager"`  // 操作者
	Players string `json:"players"`  // 角色ID
	Title   string `json:"title"`    //邮件标题
	Content string `json:"content"`  //邮件标题
	Attach  string `json:"attach"`   //物品列表，json格式
	Endtime int64  `json:"end_time"` //过期时间
}

//发送游戏走马灯
type XDGM_SEND_GAME_LAMP_REQ struct {
	Manager   string `json:"manager"`    // 操作者
	Content   string `json:"content"`    //跑马灯内容
	BeginTime int64  `json:"begin_time"` //开始时间
	EndTime   int64  `json:"end_time"`   //结束时间
	Interval  int32  `json:"interval"`   //跑马灯间隔
}

//查询游戏走马灯
type XDGM_SEARCH_GAME_LAMP_REQ struct {
}

//删除游戏走马灯
type XDGM_DEL_GAME_LAMP_REQ struct {
	Manager  string `json:"manager"` // 操作者
	NoticeId int64  `json:"notice_id"`
}

// 群发邮件
type XDGM_SEND_MAIL_ALL_REQ struct {
	Manager     string `json:"manager"`       // 操作者
	Title       string `json:"title"`         //邮件标题
	Content     string `json:"content"`       //邮件标题
	Attach      string `json:"attach"`        //物品列表，json格式
	BeginTime   int64  `json:"begin_time"`    //开始时间
	Endtime     int64  `json:"end_time"`      //过期时间
	MinLevel    int16  `json:"min_level"`     //最小等级
	MaxLevel    int16  `json:"max_level"`     //最大等级
	MinVipLevel int16  `json:"min_vip_level"` //最小vip等级
	MaxVipLevel int16  `json:"max_vip_level"` //最大vip等级
}

// 更新活动奖励
type XDGM_UPDATE_EVENT_AWARDS_REQ struct {
	ServerId    int    `json:"server_id"`
	EventAwards string `json:"event_awards"`
}

// 拉取纯文案活动
type XDGM_GET_TEXT_EVENTS_REQ struct {
	ServerId int `json:"server_id"`
}

//增加玩家VIP经验
type XDGM_ADD_VIP_EXP_REQ struct {
	Manager string `json:"manager"` // 操作者
	Players string `json:"players"` // 角色ID
	Value   int32  `json:"value"`   // vip经验
}

//修改充值返利规则
type XDGM_SET_PAYMENTS_PRESENT_REQ struct {
	Manager   string `json:"manager"`    // 操作者
	Servers   string `json:"servers"`    // 角色ID
	Rule      string `json:"rule"`       // vip经验
	BeginTime int64  `json:"begin_time"` //返利开始时间
	EndTime   int64  `json:"end_time"`   //返利结束时间
}

//生成兑换码请求
type XDGM_GEN_GIFT_CODE_REQ struct {
	Num             int16              `json:"num"`              //数量
	Servers         []int              `json:"servers"`          //目标服务器 servers 为空则对所有服务器有效
	Type            int8               `json:"type"`             //0--独占型 1--共享型号
	EffectTimestamp int64              `json:"effect_timestamp"` //生效时间戳
	ExpireTimestamp int64              `json:"expire_timestamp"` //过期时间戳
	Content         string             `json:"content"`          //使用兑换码后邮件内容
	Config          []GIFT_CODE_CONFIG `json:"config"`           //使用兑换码后邮件内容
}

type GIFT_CODE_CONFIG struct {
	ItemId  int16 `json:"item_id"`  // 赠送道具ID
	ItemNum int64 `json:"item_num"` // 赠送道具数量
}

//撤销兑换码
type XDGM_CANCEL_GIFT_CODE_REQ struct {
	ServerId int   `json:"server_id"` //服务器ID < 0 只证据 version 取消
	Version  int64 `json:"version"`   //兑换码批次
}

//查询兑换码
type XDGM_QUERY_GIFT_CODE_REQ struct {
	ServerId int   `json:"server_id"` //服务器ID < 0 只证据 version 取消
	Version  int64 `json:"version"`   //兑换码批次
	Limit    int   `json:"limit"`
	Offset   int   `json:"offset"`
}

//发布激活码变化
type XDGM_RELEASE_GIFT_CODE_REQ struct {
	ServerId int `json:"server_id"` //服务器ID < 0 只证据 version 取消
}

//patch公告
type XDGM_PATCH_NOTICE_REQ struct {
	Manager string `json:"manager"`     //操作者
	Content string `json:"content"`     //公告内容
	Time    int64  `json:"notice_time"` //公告右下角显示时间
}

//ip封解查操作
type XDGM_BLACK_IP_REQ struct {
	Manager string `json:"manager"` //操作者
	Mode    string `json:"mode"`    //模式，0-封，1-解，2-查
	Ip      string `json:"ip"`      //ip
}
