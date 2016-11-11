package battle

// 战场类型
const (
	BT_MISSION_LEVEL          = 0  // 区域关卡
	BT_RESOURCE_LEVEL         = 1  // 资源关卡
	BT_TOWER_LEVEL            = 2  // 通天塔
	BT_MULTI_LEVEL            = 3  // 多人关卡
	BT_ARENA                  = 4  // 比武场
	BT_HARD_LEVEL             = 8  // 难度关卡
	BT_BUDDY_LEVEL            = 9  // 伙伴关卡
	BT_PET_LEVEL              = 10 // 灵宠关卡
	BT_GHOST_LEVEL            = 11 // 魂侍关卡
	BT_RAINBOW_LEVEL          = 12 // 彩虹关卡
	BT_PVE_LEVEL              = 13 // 灵宠幻境关卡
	BT_FATE_BOX_LEVEL         = 14 // 命锁宝箱关卡
	BT_DRIVING_LEVEL          = 15 // 仙山探险
	BT_DRIVING_SWORD_BF_LEVEL = 16 // 深海御剑拜访关卡
	BT_HIJACK_BOAT            = 17 // 劫持镖船
	BT_RECOVER_BOAT           = 18 // 夺回镖船
)

// 战斗对象的类型
const (
	FK_PLAYER = 0 // 玩家
	FK_BUDDY  = 1 // 伙伴
	FK_ENEMY  = 2 // 敌人
)

// 战斗角色种族
const (
	FP_SPIRIT = 0 // 仙灵
	FP_HUMAN  = 1 // 人畜
	FP_DEVIL  = 2 // 妖魔
)

// 出手者或被攻击目标的类型
const (
	FT_ATK = 0 // 攻方
	FT_DEF = 1 // 守方
)

// 战斗时发生的事件类型
const (
	FE_NONE    = 0 // 无
	FE_DODGE   = 1 // 闪避
	FE_CRIT    = 2 // 暴击
	FE_BLOCK   = 3 // 格挡
	FE_SQUELCH = 4 // 反击
)

// buff类型（跟客户端同步）
const (
	BUFF_POWER              = 0  // 精气
	BUFF_SPEED              = 1  // 速度
	BUFF_ATTACK             = 2  // 攻击
	BUFF_DEFEND             = 3  // 防御
	BUFF_HEALTH             = 4  // 生命
	BUFF_DIZZINESS          = 5  // 眩晕
	BUFF_POISONING          = 6  // 中毒
	BUFF_CLEAN_BAD          = 7  // 清除负面buff
	BUFF_CLEAN_GOOD         = 8  // 清除增益buff
	BUFF_REDUCE_HURT        = 9  // 伤害减免
	BUFF_RANDOM             = 10 // 混乱
	BUFF_BLOCK              = 11 // 格挡概率
	BUFF_BLOCK_LEVEL        = 12 // 格挡概率等级
	BUFF_DODGE_LEVEL        = 13 // 闪避概率等级
	BUFF_CRITIAL_LEVEL      = 14 // 暴击等级
	BUFF_HIT_LEVEL          = 15 // 命中等级
	BUFF_HURT_ADD           = 16 // 伤害加值（百分数）
	BUFF_MAX_HEALTH         = 17 // 增加最大生命
	BUFF_KEEPER_REDUCE_HURT = 18 // 守卫者免伤
	BUFF_ATTRACT_FIRE       = 19 // 吸引火力
	BUFF_DESTROY_LEVEL      = 20 // 破击
	BUFF_TENACITY_LEVEL     = 21 // 韧性
	BUFF_SUNDER             = 22 // 护甲
	BUFF_SLEEP              = 23 // 睡眠
	BUFF_DISABLE_SKILL      = 24 // 禁用绝招
	BUFF_DIRECT_REDUCE_HURT = 25 // 直接免伤
	BUFF_ABSORB_HURT        = 26 // 吸收伤害
	BUFF_GHOST_POWER        = 27 // 魂力
	BUFF_PET_LIVE_ROUND     = 28 //灵宠存活回合
	BUFF_BUDDY_SKILL        = 29 //伙伴进阶技能使用次数
	BUFF_CLEAR_ABSORB_HURT  = 30 //清除伤害吸收护盾

	BUFF_SLEEP_LEVEL         = 31 //睡眠抗性等级
	BUFF_DIZZINESS_LEVEL     = 32 //眩晕抗性等级
	BUFF_RANDOM_LEVEL        = 33 //混乱抗性等级
	BUFF_DISABLE_SKILL_LEVEL = 34 //封魔抗性等级
	BUFF_POISONING_LEVEL     = 35 //中毒抗性等级
	BUFF_RECOVER_BUDDY_SKILL = 36 //恢复伙伴进阶技能使用次数
	BUFF_MAKE_POWER_FULL     = 37 //精力恢复至满
	BUFF_DOGE                = 38 //闪避
	BUFF_HIT                 = 39 //命中
	BUFF_CRITIAL             = 40 //暴击
	BUFF_TENACITY            = 41 //韧性
	BUFF_TAKE_SUNDER         = 42 //破甲

	BUFF_DEFEND_PERSENT = 43 // 防御系数 (原防御 + (原防御 * buff_value))

	BUFF_SUNDER_STATE = 44 // 破甲状态

	BUFF_HEALTH_PERCENT = 45 // 生命百分比

	BUFF_ALL_RESIST = 46 // 所有抗性等级(睡眠，眩晕，混乱，封魔，中毒等)

	BUFF_REBOTH_HEALTH = 47 // 复活并指定血量

	BUFF_REBOTH_HEALTH_PERCENT = 48 // 复活并指定血量百分比

	BUFF_GHOST_SHIELD = 127 // 魂侍护盾
	// 添加类型后需要做的：
	// 0. buff的apply和revert增加类型判断
	// 1. addBuff增加对应的判断逻辑
	// 2. database的代码生成添加类型
	// 	database/plugins/server_data/data/enemy_skill.php
	// 	system/pages/skill_editor.php
	// 	没有 value 概念的需要特殊处理 get_buff_value

	// 3. 通讯协议增加类型
	// 4. 客户端处理

)

// 机率等级调整系数
const (
	CRITIAL_LEVEL_ARG       = 0.025 // 暴击等级调整系数
	DODGE_LEVEL_ARG         = 0.025 // 闪避等级调整系数
	HIT_LEVEL_ARG           = 0.025 // 命中等级调整系数
	BLOCK_LEVEL_ARG         = 0.025 // 格挡等级调整系数
	CRITIAL_HURT_LEVEL_ARG  = 0.025 // 暴击伤害调整系数
	TENACITY_LEVEL_ARG      = 0.025 // 韧性等级调整系数
	DESTROY_LEVEL_ARG       = 0.025 // 破击等级调整系数
	SLEEP_LEVEL_ARG         = 0.025 // 睡眠等级调整系数
	DIZZINESS_LEVEL_ARG     = 0.025 // 眩晕等级调整系数
	RANDOM_LEVEL_ARG        = 0.025 // 混乱等级调整系数
	DISABLE_SKILL_LEVEL_ARG = 0.025 // 封魔等级调整系数
	POISONING_LEVEL_ARG     = 0.025 // 中毒等级调整系数
	SPIRIT_HURT_ARG         = 0.025 // 仙灵伤害调整系数
	HUMAN_HURT_ARG          = 0.025 // 人畜伤害调整系数
	DEVIL_HURT_ARG          = 0.025 // 妖魔伤害调整系数
)

// 战斗状态结果返回
const (
	NOT_END             = 0 // 还没结束
	ATK_WIN             = 1 // 攻击方胜利
	DEF_WIN             = 2 // 守方胜利
	ATK_NEXT            = 3 // 攻击方切换下一组
	DEF_NEXT            = 4 // 防守方切换下一组
	TRIGGER_CALL_ENEMYS = 5 //触发怪物召唤技能
	WAITING             = 6 //等待其他玩家操作
)

const (
	MAX_FORM_ROLES = 15 // 最大布阵人数
	MAX_COLS       = 5
	MAX_ROUND      = 50  // 最大回合数
	MIN_HEALTH     = 0   // 最小生命值
	MAX_SKILLS     = 255 // 最大绝招ID
	ROLE_INIT_HIT  = 100 // 初始命中100%
)

const (
	JOB_ATTACKE = 1 // 进攻者
	JOB_DAMAGE  = 2 // 破坏者
	JOB_KEEP    = 3 // 守卫者
	JOB_CURE    = 4 // 治疗者
	JOB_SUPPORT = 5 // 支援者
	JOB_BLOCK   = 6 // 阻碍者
)

//技能类型
const (
	SKILL_KIND_ATTACK        = 1 //进攻
	SKILL_KIND_DEFEND        = 3 //防御
	SKILL_KIND_CURE          = 4 //治疗
	SKILL_KIND_SUPPORT       = 5 //辅助
	SKILL_KIND_REDUCE_SUNDER = 6 //破甲
)

const (
	DEFAULT_SKILL_ID          = -3  // 默认绝招ID
	YE_KAI_ATTACK_UP_SKILL_ID = 999 // 叶开在其他伙伴阵亡时，攻击力增加50%，持续5回合用的特殊绝招ID
	SUNDER_DIZZINESS_SKILL_ID = 998 // 破甲眩晕用的特殊绝招ID
	ROUND_BUFF_SKILL_ID       = 997 // 回合狂暴逻辑用到的特殊绝招ID
	GHOST_SHIELD_SKILL_ID     = 996 // 魂侍护盾特殊绝招ID

	PET_LEVEL_MAIN_ROLE_SKILL_ID = 1041 // 灵宠关卡主角特殊技（聚气）
)

const (
	SKILL_TYPE_ROLE  = 1 // 绝招大类 角色招式
	SKILL_TYPE_ENEMY = 5 // 绝招大类 怪绝招
	SKILL_EMPTY      = 0 // 无绝招
	MAX_FORM_NUM     = 3 // 最大上阵人数
)

// for go build
const (
	SWORD_SOUL_SKILL_ID = 0
)

const (
	FULL_GHOST_POWER         = 100
	ATTACK_ADD_GHOST_POWER   = 5  //主动攻击增加魂力
	ATTACKED_ADD_GHOST_POWER = 5  //被攻击增加魂力
	KILL_ADD_GHOST_POWER     = 10 //击杀获对手增加魂力
)

const (
	APPEND_SPECIAL_TYPE_NONE               = 0
	APPEND_SPECIAL_TYPE_ATTACKED_INC_POWER = 1
)

const (
	RECOVER_POWER = 3 //主角每次自动增加3点精气
)
