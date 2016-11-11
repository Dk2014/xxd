// this is auto-genrated file,
// Don't modify this file manually

package idip_server

type IDIP_COMMON_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
}



// IDIP消息头
type IDIP_REQ_HEAD struct {
	PacketLen uint32 // 包长
	Cmdid uint32 // 命令ID
	Seqid uint32 // 流水号
	ServiceName string // 服务名
	SendTime uint32 // 发送时间YYYYMMDD对应的整数
	Version uint32 // 版本号
	Authenticate string // 加密串
	Result int32 // 错误码,返回码类型：0：处理成功，需要解开包体获得详细信息,1：处理成功，但包体返回为空，不需要处理包体（eg：查询用户角色，用户角色不存在等），-1: 网络通信异常,-2：超时,-3：数据库操作异常,-4：API返回异常,-5：服务器忙,-6：其他错误,小于-100 ：用户自定义错误，需要填写szRetErrMsg
	RetErrMsg string // 错误信息

}



// 角色装备信息对象
type SEquipInfo struct {
	EquipId uint64 // 装备ID
	EquipName string // 装备名称
	Level uint32 // 装备精炼等级
	IsBattle uint8 // 上阵1/未上阵0

}



// 角色灵宠信息对象
type SRolePetInfo struct {
	PetId uint64 // 灵宠ID
	PetName string // 灵宠名称
	Level uint32 // 灵宠等级
	IsBattle uint8 // 上阵1/未上阵0

}



// 剑心信息对象
type SSwordInfo struct {
	Slot uint32 // 开启槽数
	SwordName string // 剑心名
	Level uint32 // 剑心等级
	IsBattle uint8 // 上阵1/未上阵0

}



// 魂侍信息对象
type SSoulInfo struct {
	Slot uint32 // 开启槽数
	RoleName string // 角色名称
	MajorSoul uint32 // 主魂侍
	MajorSoulLevel uint32 // 魂侍等级
	Minor1Soul uint32 // 辅魂侍1
	MinorSoul1Level uint32 // 辅魂侍1等级
	Minor2Soul uint32 // 辅魂侍2
	MinorSoul2Level uint32 // 辅魂侍2等级
	Minor3Soul uint32 // 辅魂侍3
	MinorSoul3Level uint32 // 辅魂侍3等级
	IsBattle uint8 // 上阵1/未上阵0

}



// 任务进度信息对象
type STaskProgressInfo struct {
	TaskName string // 任务名称
	Type uint8 // 任务类型（1：主线任务 2：每日任务）
	Status uint8 // 任务状态（1:正在进行中,2:完成未领取奖励,3:完成已领取奖励）

}



// 背包存量信息对象
type SBagInfo struct {
	ItemName string // 道具名称
	ItemId uint64 // 道具ID（包括灵宠、魂侍、剑心、消耗品道具等）
	ItemNum uint32 // 道具存量

}



// 走马灯公告信息对象
type SGameLampNoticeInfo struct {
	BeginTime string // 公告生效时间
	EndTime string // 公告结束时间
	Freq int32 // 滚动频率
	NoticeId string // 公告ID
	NoticeContent string // 公告内容

}



// 封号请求
type IDIP_DO_BAN_USR_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	OpenId string // openid
	BanTime uint32 // 封号时长0 永久，**秒）
	Source uint32 // 渠道号，由前端生成，不需要填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 封号应答
type IDIP_DO_BAN_USR_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息

	} `json:"body"`

}



// 解封请求
type IDIP_DO_UNBAN_USR_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	OpenId string // openid
	Source uint32 // 渠道号，由前端生成，不需要填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 解封应答
type IDIP_DO_UNBAN_USR_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息

	} `json:"body"`

}



// 删除道具请求
type IDIP_DO_DEL_ITEM_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	OpenId string // openid
	ItemId uint64 // 道具ID
	ItemNum uint32 // 删除数量
	Source uint32 // 渠道号，由前端生成，不需要填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 删除道具应答
type IDIP_DO_DEL_ITEM_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息

	} `json:"body"`

}



// 修改角色经验请求
type IDIP_DO_UPDATE_ROLE_EXP_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	OpenId string // openid
	RoleId uint64 // 角色ID
	Value int32 // 修改值（正加负减）
	Source uint32 // 渠道号，由前端生成，不需要填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 修改角色经验应答
type IDIP_DO_UPDATE_ROLE_EXP_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息

	} `json:"body"`

}



// 修改用户元宝请求
type IDIP_DO_UPDATE_ROLE_GOLD_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	OpenId string // openid
	RoleId uint64 // 角色ID
	Value int32 // 数量（正加负减）
	Source uint32 // 渠道号，由前端生成，不需要填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 修改用户元宝应答
type IDIP_DO_UPDATE_ROLE_GOLD_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息

	} `json:"body"`

}



// 修改用户体力请求
type IDIP_DO_UPDATE_ROLE_PHYSICAL_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	OpenId string // openid
	RoleId uint64 // 角色ID
	Value int32 // 数量（正加负减）
	Source uint32 // 渠道号，由前端生成，不需要填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 修改用户体力应答
type IDIP_DO_UPDATE_ROLE_PHYSICAL_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息

	} `json:"body"`

}



// 修改角色铜钱请求
type IDIP_DO_UPDATE_ROLE_COIN_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	OpenId string // openid
	RoleId uint64 // 角色ID
	Value int32 // 数量（正加负减）
	Source uint32 // 渠道号，由前端生成，不需要填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 修改角色铜钱应答
type IDIP_DO_UPDATE_ROLE_COIN_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息

	} `json:"body"`

}



// 修改角色爱心请求
type IDIP_DO_UPDATE_ROLE_HEART_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	OpenId string // openid
	RoleId uint64 // 角色ID
	Value int32 // 数量（正加负减）
	Source uint32 // 渠道号，由前端生成，不需要填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 修改角色爱心应答
type IDIP_DO_UPDATE_ROLE_HEART_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息

	} `json:"body"`

}



// 发放道具请求
type IDIP_DO_SEND_ITEM_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	OpenId string // openid
	RoleId uint64 // 角色ID
	ItemId uint64 // 道具ID
	ItemNum uint32 // 数量
	Source uint32 // 渠道号，由前端生成，不需要填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 发放道具应答
type IDIP_DO_SEND_ITEM_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息

	} `json:"body"`

}



// 查询当前个人信息请求
type IDIP_QUERY_USR_INFO_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	OpenId string // openid

	} `json:"body"`

}



// 查询当前个人信息应答
type IDIP_QUERY_USR_INFO_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Level uint32 // 当前等级
	Vip uint32 // 当前VIP等级
	Exp uint32 // 当前经验
	Coin uint32 // 当前铜钱
	Gold uint32 // 当前元宝数量
	Physical uint32 // 当前体力值
	MaxPhysical uint32 // 体力值上限
	MaxBag uint32 // 背包上限值
	RegisterTime uint64 // 注册时间
	IsOnline uint8 // 是否在线（0在线，1离线）
	AccStatus uint8 // 帐号状态（0 正常，1封号）
	BanEndTime uint64 // 封号截至时间
	ArmyId uint64 // 所在公会
	RankInArmy uint32 // 在公会中的排名
	ArmyRank uint32 // 公会排名
	PassProgress uint32 // 当前关卡进度
	PvpRank uint32 // 个人PVP排名
	PvpScore uint32 // PVP积分数量
	LastLoginTime string // 玩家最后登录时间
	RoleName string // 角色名称

	} `json:"body"`

}



// 查询角色装备信息请求
type IDIP_QUERY_ROLE_EQUIP_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	OpenId string // openid
	RoleId uint64 // 角色ID

	} `json:"body"`

}



// 查询角色装备信息应答
type IDIP_QUERY_ROLE_EQUIP_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	EquipList_count uint32 // 角色装备信息列表的最大数量
	EquipList []SEquipInfo // 角色装备信息列表

	} `json:"body"`

}



// 查询角色灵宠信息请求
type IDIP_QUERY_ROLE_PET_INFO_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	OpenId string // openid
	RoleId uint64 // 角色ID

	} `json:"body"`

}



// 查询角色灵宠信息应答
type IDIP_QUERY_ROLE_PET_INFO_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	RolePetList_count uint32 // 角色灵宠信息列表的最大数量
	RolePetList []SRolePetInfo // 角色灵宠信息列表

	} `json:"body"`

}



// 查询剑心信息请求
type IDIP_QUERY_SWORD_INFO_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	OpenId string // openid
	RoleId uint64 // 角色ID

	} `json:"body"`

}



// 查询剑心信息应答
type IDIP_QUERY_SWORD_INFO_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	SwordList_count uint32 // 剑心信息列表的最大数量
	SwordList []SSwordInfo // 剑心信息列表

	} `json:"body"`

}



// 查询魂侍信息请求
type IDIP_QUERY_SOUL_INFO_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	OpenId string // openid
	RoleId uint64 // 角色ID

	} `json:"body"`

}



// 查询魂侍信息应答
type IDIP_QUERY_SOUL_INFO_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	SoulList_count uint32 // 魂侍信息列表的最大数量
	SoulList []SSoulInfo // 魂侍信息列表

	} `json:"body"`

}



// 查询任务进度请求
type IDIP_QUERY_TASK_INFO_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	OpenId string // openid
	BeginTime uint64 // 开始时间
	EndTime uint64 // 结束时间

	} `json:"body"`

}



// 查询任务进度应答
type IDIP_QUERY_TASK_INFO_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	TaskProgressList_count uint32 // 任务进度信息列表的最大数量
	TaskProgressList []STaskProgressInfo // 任务进度信息列表

	} `json:"body"`

}



// 查询背包存量信息请求
type IDIP_QUERY_BAG_INFO_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	OpenId string // openid
	BeginTime uint64 // 开始时间
	EndTime uint64 // 结束时间

	} `json:"body"`

}



// 查询背包存量信息应答
type IDIP_QUERY_BAG_INFO_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	BagList_count uint32 // 背包存量信息列表的最大数量
	BagList []SBagInfo // 背包存量信息列表

	} `json:"body"`

}



// 发放邮件请求
type IDIP_DO_SEND_MAIL_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	OpenId string // openid
	MailTitle string // 邮件标题
	MailContent string // 邮件内容
	SendTime uint64 // 邮件发送时间
	EndTime uint64 // 邮件自动删除时间
	ItemId1 uint64 // 邮件赠送道具1ID
	ItemNum1 uint32 // 邮件赠送道具1数量
	ItemId2 uint64 // 邮件赠送道具2ID
	ItemNum2 uint32 // 邮件赠送道具2数量
	ItemId3 uint64 // 邮件赠送道具3ID
	ItemNum3 uint32 // 邮件赠送道具3数量
	ItemId4 uint64 // 邮件赠送道具4ID
	ItemNum4 uint32 // 邮件赠送道具4数量
	Source uint32 // 渠道号，由前端生成，不需要填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 发放邮件应答
type IDIP_DO_SEND_MAIL_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息

	} `json:"body"`

}



// 更新新闻公告请求
type IDIP_DO_UPDATE_NOTICE_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	Title string // 新闻标题
	Content string // 新闻内容
	Time uint64 // 更新时间
	Source uint32 // 渠道号，由前端生成，不需要填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 更新新闻公告应答
type IDIP_DO_UPDATE_NOTICE_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息

	} `json:"body"`

}



// 更新游戏内走马灯请求
type IDIP_DO_UPDATE_GAME_LAMP_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	Freq uint32 // 滚动频率：**秒
	LampContent string // 走马灯内容
	BeginTime uint64 // 开始时间
	EndTime uint64 // 结束时间:不配置，则表示不限制结束时间
	Source uint32 // 渠道号，由前端生成，不需要填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 更新游戏内走马灯应答
type IDIP_DO_UPDATE_GAME_LAMP_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息

	} `json:"body"`

}



// 发送邮件（AQ）请求
type IDIP_AQ_DO_SEND_MAIL_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	OpenId string // openid
	MailContent string // 邮件内容
	Reason string // 操作原因
	Source uint32 // 渠道号，由前端生成，不需填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 发送邮件（AQ）应答
type IDIP_AQ_DO_SEND_MAIL_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 结果：0 成功，其它失败
	RetMsg string // 返回消息

	} `json:"body"`

}



// 封号（AQ）请求
type IDIP_AQ_DO_BAN_USR_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	OpenId string // openid
	Time uint32 // 封停时长（秒）
	Reason string // 操作原因
	Source uint32 // 渠道号，由前端生成，不需要填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 封号（AQ）应答
type IDIP_AQ_DO_BAN_USR_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 操作结果
	RetMsg string // 返回消息

	} `json:"body"`

}



// 解除处罚（AQ）请求
type IDIP_AQ_DO_RELIEVE_PUNISH_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 大区（1微信，2手Q）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台（0ios，1安卓）
	OpenId string // openid
	RelieveBan uint8 // 解除封号（0 否，1 是）
	Reason string // 操作原因
	Source uint32 // 渠道号，由前端生成，不需要填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 解除处罚（AQ）应答
type IDIP_AQ_DO_RELIEVE_PUNISH_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 操作结果
	RetMsg string // 返回消息

	} `json:"body"`

}



// 查询比武场请求
type IDIP_QUERY_COMPETE_INFO_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 服务器：微信（1），手Q（2）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台：IOS（0），安卓（1）
	OpenId string // openid

	} `json:"body"`

}



// 查询比武场应答
type IDIP_QUERY_COMPETE_INFO_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	RoleId string // 用户ID
	Rank int32 // 当前排名

	} `json:"body"`

}



// 设置魂侍等级请求
type IDIP_DO_SET_SOUL_LEVEL_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 服务器：微信（1），手Q（2）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台：IOS（0），安卓（1）
	OpenId string // openid
	SoulId uint64 // 魂侍ID
	Value int32 // 等级设置：填1则表示1级，2则表示2级
	Source uint32 // 渠道号，由前端生成，不需要填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 设置魂侍等级应答
type IDIP_DO_SET_SOUL_LEVEL_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息

	} `json:"body"`

}



// 设置剑心等级请求
type IDIP_DO_SET_SWORD_LEVEL_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 服务器：微信（1），手Q（2）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台：IOS（0），安卓（1）
	OpenId string // openid
	RoleId uint64 // 人物ID
	Pos uint32 // 剑心位置ID
	Value int32 // 等级设置：填1则表示1级，2则表示2级
	Source uint32 // 渠道号，由前端生成，不需要填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 设置剑心等级应答
type IDIP_DO_SET_SWORD_LEVEL_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息

	} `json:"body"`

}



// 设置宠物激活请求
type IDIP_DO_SET_PET_ACTIVE_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 服务器：微信（1），手Q（2）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台：IOS（0），安卓（1）
	OpenId string // openid
	PetId uint64 // 宠物ID
	Source uint32 // 渠道号，由前端生成，不需要填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 设置宠物激活应答
type IDIP_DO_SET_PET_ACTIVE_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息

	} `json:"body"`

}



// 清零数据请求
type IDIP_DO_CLEAR_DATA_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 服务器：微信（1），手Q（2）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台：IOS（0），安卓（1）
	OpenId string // openid
	ClearCoin uint8 // 铜钱清零：否（0），是（1）
	ClearPhysical uint8 // 体力清零：否（0），是（1）
	ClearHeart uint8 // 爱心清零：否（0），是（1） 
	Type uint32 // 需求类型
	Reason string // 操作原因
	Source uint32 // 渠道号，由前端生成，不需填写
	Serial string // 流水号，由前端生成，不需填写

	} `json:"body"`

}



// 清零数据应答
type IDIP_DO_CLEAR_DATA_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 结果：0 成功，其它失败
	RetMsg string // 返回消息

	} `json:"body"`

}



// 发放全区全服邮件请求
type IDIP_DO_SEND_MAIL_ALL_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 服务器：微信（1），手Q（2）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台：IOS（0），安卓（1）
	MailTitle string // 邮件标题
	MailContent string // 邮件内容
	SendTime uint64 // 邮件发送时间
	EndTime uint64 // 邮件自动删除时间
	ItemId1 uint64 //  邮件赠送道具1ID
	ItemNum1 uint32 // 邮件赠送道具1数量
	ItemId2 uint64 //  邮件赠送道具2ID
	ItemNum2 uint32 // 邮件赠送道具2数量
	ItemId3 uint64 //  邮件赠送道具3ID
	ItemNum3 uint32 // 邮件赠送道具3数量
	ItemId4 uint64 //  邮件赠送道具4ID
	ItemNum4 uint32 // 邮件赠送道具4数量
	Source uint32 // 渠道号，由前端生成，不需要填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 发放全区全服邮件应答
type IDIP_DO_SEND_MAIL_ALL_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 结果（0）成功
	RetMsg string // 返回消息

	} `json:"body"`

}



// 查询走马灯公告请求
type IDIP_QUERY_GAME_LAMP_NOTICE_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 服务器：微信（1），手Q（2）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台：IOS（0），安卓（1）

	} `json:"body"`

}



// 查询走马灯公告应答
type IDIP_QUERY_GAME_LAMP_NOTICE_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	GameLampNoticeList_count uint32 // 走马灯公告信息列表的最大数量
	GameLampNoticeList []SGameLampNoticeInfo // 走马灯公告信息列表

	} `json:"body"`

}



// 删除公告请求
type IDIP_DO_DEL_NOTICE_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 服务器：微信（1），手Q（2）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台：IOS（0），安卓（1）
	NoticeId uint64 // 公告ID
	Source uint32 // 渠道号，由前端生成，不需要填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 删除公告应答
type IDIP_DO_DEL_NOTICE_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 结果
	RetMsg string // 返回消息

	} `json:"body"`

}



// 修改vip等级请求
type IDIP_DO_UPDATE_VIP_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 服务器：微信（1），手Q（2），游客（5）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台：IOS（0），安卓（1）
	OpenId string // openid
	Value uint32 // 修改数量
	Source uint32 // 渠道号，由前端生成，不需要填写
	Serial string // 流水号，由前端生成，不需要填写

	} `json:"body"`

}



// 修改vip等级应答
type IDIP_DO_UPDATE_VIP_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Result int32 // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息

	} `json:"body"`

}



// 查询深渊挂关卡进度请求
type IDIP_QUERY_ABYSSPASS_STATE_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 服务器：微信（1），手Q（2），游客（5）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台：IOS（0），安卓（1）
	OpenId string // openid

	} `json:"body"`

}



// 查询深渊挂关卡进度应答
type IDIP_QUERY_ABYSSPASS_STATE_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Status int32 // 深渊关卡进度

	} `json:"body"`

}



// 查询战力请求
type IDIP_QUERY_FIGHT_REQ struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	AreaId uint32 // 服务器：微信（1），手Q（2），游客（5）
	Partition uint32 // 小区ID
	PlatId uint8 // 平台：IOS（0），安卓（1）
	OpenId string // openid

	} `json:"body"`

}



// 查询战力应答
type IDIP_QUERY_FIGHT_RSP struct {
	Head IDIP_REQ_HEAD `json:"head"`
	Body struct {
	Fight int32 // 玩家战力

	} `json:"body"`

}

