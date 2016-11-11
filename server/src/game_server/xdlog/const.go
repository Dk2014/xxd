package xdlog

//ConsumeKind消耗类型
const (
	RECENT  = 1 //实时
	TIMER   = 1 //时长
	FOREVER = 1 //永久
)

//EventType事件类型
const (
	ET_DEBUG                                  = 0   //DEBUG获得
	ET_BUY_ITEM                               = 1   //购买游戏内物品
	ET_SELL_ITEM                              = 2   //游戏内出售物品
	ET_BUY_BACK_ITEM                          = 3   //游戏内物品回购
	ET_ITEM_DECOMPOSE                         = 4   //分解装备
	ET_OPEN_ITEM_CHEST                        = 5   //打开物品宝箱
	ET_GHOST_COMPOSE                          = 6   //魂侍合成
	ET_MAIL_TAKE_ATTACHMENT                   = 7   //获取邮件附件
	ET_MISSION_STAR_AWARD                     = 8   //副本星级奖励
	ET_BATTLE_PET_UPGRADE                     = 9   //灵宠升级
	ET_OPEN_FUNC                              = 10  //功能开放
	ET_MISSION_LEVEL_AWARD                    = 11  //通关奖励
	ET_MISSION_LEVEL_SMALL_BOX                = 12  //打开副本内小宝箱
	ET_AUTO_FIGHT                             = 13  //关卡扫荡
	ET_MISSION_LEVEL_CATCH_PET                = 14  //关卡灵宠抓取
	ET_MISSION_LEVEL_USE_ITEM                 = 15  //副本中使用物品
	ET_SHADED_AWARD                           = 16  //影之间隙奖励
	ET_MISSION_LEVEL_RANDOM_AWARD             = 17  //通关奖励随机宝箱
	ET_MISSION_OPEN                           = 18  //开启区域
	ET_DISPOSE_EVENT                          = 19  //掉线自动奖励
	ET_ITEM_ROLE_USE                          = 20  //角色使用物品
	ET_ARENA_AWARD_BOX                        = 21  //比武场奖励
	ET_PET_VIRTUAL                            = 22  //灵宠幻境
	ET_ADDITION_QUEST                         = 23  //支线任务
	ET_DAILY_QUEST                            = 24  //每日任务
	ET_EXTEND_QUEST                           = 25  //扩展任务
	ET_QUEST                                  = 26  //主线任务
	ET_GHOST_SKILL                            = 27  //魂侍技能
	ET_PET_VIRTUAL_AUTO_FIGHT                 = 28  //灵宠幻境扫荡
	ET_PET_VIRTUAL_WIN                        = 29  //灵宠幻境胜利
	ET_PET_VIRTUAL_LOSE                       = 30  //灵宠幻境失败
	ET_MISSION_LEVEL_WIN                      = 31  //关卡胜利
	ET_DAILY_SIGN                             = 32  //每日签到
	ET_STAR_FATE_BOX_FREE                     = 33  //免费星辉命锁宝箱
	ET_STAR_FATE_BOX_ONE                      = 34  //命锁星辉宝箱一次
	ET_STAR_FATE_BOX_TEN                      = 35  //命锁星辉宝箱十次
	ET_MOON_FATE_BOX_FREE                     = 36  //免费月影命锁宝箱
	ET_MOON_FATE_BOX_ONE                      = 37  //命锁月影宝箱一次
	ET_MOON_FATE_BOX_TEN                      = 38  //命锁月影宝箱十次
	ET_SUN_FATE_BOX_FREE                      = 39  //免费日耀命锁宝箱
	ET_SUN_FATE_BOX_ONE                       = 40  //命锁日耀宝箱一次
	ET_SUN_FATE_BOX_TEN                       = 41  //命锁日耀宝箱十次
	ET_HUNYUAN_FATE_BOX_FREE                  = 42  //免费混元命锁宝箱
	ET_HUNYUAN_FATE_BOX_ONE                   = 43  //命锁混元宝箱一次
	ET_HUNYUAN_FATE_BOX_TEN                   = 44  //命锁混元宝箱十次
	ET_HEART_DRAW                             = 45  //爱心抽奖
	ET_CLOUD_CLIMB                            = 46  //开启云海御剑云层
	ET_EVENT_CENTER                           = 47  //活动中心
	ET_MOUNTAIN_TREASURE_OPEN                 = 48  //开启仙山宝箱
	ET_AWARD_GARRISON                         = 49  //驻守奖励
	ET_TOWN_TREASURES                         = 50  //城镇奖励
	ET_NPC_TALK                               = 51  //npc对话奖励
	ET_DRIVING_VISIT_AWARD                    = 52  //仙山拜访
	ET_DRIVING_SWORD_WIN                      = 53  //云海御剑胜利
	ET_GHOST_TRAIN                            = 54  //魂侍培养
	ET_GHOST_UPGRADE                          = 55  //魂侍升星
	ET_FLUSH_GHOST_ATTR                       = 56  //魂侍洗点
	ET_EQUIPMENT_RECAST                       = 57  //装备重铸
	ET_USE_ITEM                               = 58  //使用物品
	ET_MULTI_LEVEL                            = 59  //多人关卡
	ET_RAINBOW                                = 60  //彩虹关卡
	ET_FRIEND_SHIP                            = 61  //羁绊
	ET_TOTEM                                  = 62  //阵印
	ET_TRADER                                 = 63  //神秘商人
	ET_PET_SKILL                              = 64  //灵宠技能
	ET_BUY_HARD_LEVEL_TIMES                   = 65  //深渊关卡次数购买
	ET_MONEY_TREE                             = 66  //摇钱树
	ET_BUY_PHYSICAL                           = 67  //购买体力
	ET_ROLE_SKILL                             = 68  //角色技能
	ET_ARENA_CD                               = 69  //比武场CD
	ET_WORLD_CHAT                             = 70  //世界聊天
	ET_MISSION_LEVEL_RELIVE                   = 71  //战斗复活
	ET_CLIQUE_DONATE                          = 72  //帮派捐赠
	ET_CLIQUE_DONATE_BANK_TRADE               = 73  //帮派钱庄交易操作
	ET_CLIQUE_TEMPLE                          = 74  //帮派上香
	ET_CLIQUE_BOAT                            = 75  //帮派运镖
	ET_CLIQUE                                 = 76  //帮派相关
	ET_DRIVING_ACTION_TIMES                   = 77  //云海购买行动点
	ET_SEALEDBOOK                             = 78  //天书
	ET_EQUIPMENT_REFINE                       = 79  //装备强化
	ET_CORNUCOPIA                             = 80  //聚宝盆
	ET_BUY_CONINS                             = 81  //开元通宝
	ET_MONTH_CARD                             = 82  //月卡
	ET_DRAW_SWORD_SOUL                        = 83  //剑山拔剑
	ET_EXCHANGE_SWORD_SOUL                    = 84  //碎片兑换剑心
	ET_CHARGE                                 = 85  //充值
	ET_GHOST_BAPTIZE                          = 86  //魂侍洗炼
	ET_EXCHANGE_GHOST_CRYSTAL                 = 87  //兑换魂侍水晶
	ET_EVENT_CENTER_LEVEL_AWARD               = 88  //活动中心_等级活动
	ET_EVENT_CENTER_STRONG_AWARD              = 89  //活动中心_战力活动
	ET_EVENT_CENTER_ARENA_RANK_AWARDS         = 90  //活动中心_比武场活动
	ET_EVENT_CENTER_PHYSICAL_AWARDS           = 91  //活动中心_活跃度活动
	ET_EVENT_CENTER_DINNER_AWARDS             = 92  //活动中心_午餐活动
	ET_EVENT_CENTER_SUPPER_AWARDS             = 93  //活动中心_晚餐活动
	ET_EVENT_CENTER_QQVIP_GIFT_AWARDS         = 94  //活动中心_QQ会员特权活动
	ET_EVENT_CENTER_TOTAL_LOGIN               = 95  //活动中心_连续登陆活动
	ET_EVENT_CENTER_VIP_CLUB                  = 96  //活动中心_仙尊俱乐部活动
	ET_EVENT_CENTER_SHARE_AWARDS              = 97  //活动中心_分享活动
	ET_EVENT_CENTER_TEN_DRAW                  = 98  //活动中心_十连抽活动
	ET_EVENT_CENTER_TOTAL_CONSUME             = 99  //活动中心_累计消费活动
	ET_EVENT_CENTER_FIRST_RECHARGE_DAILY      = 100 //活动中心_每日首冲活动
	ET_EVENT_CENTER_SEVEN_DAY_AWARDS          = 101 //活动中心_七日登陆活动
	ET_EVENT_CENTER_TOTAL_RECHARGE            = 102 //活动中心_累计充值活动
	ET_EVENT_CENTER_JSON_ARENA_RANK           = 103 //活动中心_比武场排名高级活动
	ET_EVENT_CENTER_JSON_FIRST_RECHARGE_DAILY = 104 //活动中心_每日首冲高级活动
	ET_EVENT_CENTER_JSON_NEW_YEAR             = 105 //活动中心_春节高级活动
	ET_EVENT_CENTER_JSON_SINGLE_CONSUME       = 106 //活动中心_单笔消费高级活动
	ET_EVENT_CENTER_JSON_TEN_DRAW             = 107 //活动中心_十连抽高级活动
	ET_EVENT_CENTER_JSON_TOTAL_CONSUME        = 108 //活动中心_累计消费高级活动
	ET_EVENT_CENTER_JSON_TOTAL_RECHARGE       = 109 //活动中心_累计充值高级活动
	ET_CHARGE_PRESENT                         = 110 //充值返利
	ET_BUY_RESOURCE_MISSION_LEVEL_TIMES       = 111 //购买资源关卡次数
	ET_EVENT_CENTER_JSON_SINGLE_RECHARGE      = 112 //活动中心_单笔充值高级活动
	ET_CHARGE_FRIST                           = 113 //首冲奖励
	ET_EVENT_CENTER_JSON_TOTAL_CONSUME_COIN   = 114 //活动中心_累计消费铜钱高级活动
	ET_BUY_BOSS_LEVEL_TIMES                   = 115 //购买BOSS关卡次数
	ET_ROLE_AWAKEN                            = 116 //觉醒属性
)

//LoginType登录、登出类型
const (
	LOGIN  = 0 //登录
	LOGOUT = 1 //登出
)

//MissionAction副本动作类型
const (
	MA_ENTER  = 0 //进入
	MA_FINISH = 1 //完成
	MA_AUTO   = 2 //扫荡
)

//MissionType
const (
	MT_NORMAL     = 0  //普通副本
	MT_RESOURCE   = 1  //资源关卡
	MT_HARD_LEVEL = 8  //深渊关卡
	MT_BUDDY      = 9  //伙伴关卡
	MT_PET        = 10 //灵宠关卡
	MT_GHOST      = 11 //魂侍关卡
)

//MoneyType货币类型
const (
	MT_INGOT               = 0 //元宝
	MT_COIN                = 1 //铜币
	MT_HEART               = 2 //爱心
	MT_SWORD_SOUL_FRAGMENT = 3 //剑心碎片
)

//PayType付费类型
const (
	ISPAY  = 0 //付费玩家
	NOTPAY = 1 //非付费玩家
)
