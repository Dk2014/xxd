package battle

// 绝招信息
type SkillInfo struct {
	SkillId       int
	SkillJob      int
	DecPower      int
	SkillId2      int     // 第二绝招
	SkillForce    float64 // 绝招威力，怪用到
	SkillForce2   float64 // 第二绝招威力，怪用到
	SkillTrnLv    int16   // 绝招训练等级
	SkillTrnLv2   int16   // 第二绝招训练等级
	Rhythm        int     // 释放节奏（状态值），例如：停停放放 = 1、2、3、4，1、2的时候停，3、4的时候放，入场不放绝招这个值等于0，入场可放绝招这个值等于RecoverRhythm
	UseRhythm     int     // 释放节奏（恢复后可以放几次）
	RecoverRhythm int     // 恢复节奏（放完后需要几回合恢复）
	MaxReleaseNum int     //最大技能释放次数（0则无限制）
	ReleaseNum    int     //剩余技能释放次数
}

type FindTargetFunc func(f *Fighter) []*Fighter
type GetAttackFunc func(f *Fighter, force float64, trnlv int16) float64
type BuffToBuddyFunc func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff
type BuffToSelfFunc func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff
type GhostBuffToBuddyFunc func(f *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff
type GhostBuffToSelfFunc func(f *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff
type BuffToTargetFunc func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff
type SetEventTrigger func(f, t *Fighter)

//召唤怪物的信息
type CallInfo struct {
	Enemy    int  //怪物id
	Position int8 //怪物站位
}

// 攻击模式
const (
	SKILL_ATTACK_MODE_AOE          = 1 // 范围
	SKILL_ATTACK_MODE_SINGLE       = 2 // 单体
	SKILL_ATTACK_MODE_SINGLE_TWICE = 3 // 单体两次
)

// 绝招配置
type Skill struct {
	SkillId            int              // 模板数据技能ID
	ChildType          int              // 绝招类型
	DecPower           int              // 减少精气量
	IncPower           int              // 增加精气量
	FixedValue         int              // 固定值
	NotMiss            bool             // 是否不会miss
	Critial            float64          // 爆击
	ReduceDefend       float64          // 无视防御
	SunderAttack       int              // 破甲值
	LevelSunderAttack  int              // 绝招等级破甲参数
	AttackMode         int              // 攻击模式（范围、单体）
	RecordCount        bool             // 是否记录攻击次数
	GhostAddRate       int              // 魂侍攻击、破甲、治疗的叠加转换率
	AttackRangeRatio   float64          // 攻击范围系数(根据攻击范围决定，比如单体、全体等)
	_FindTargets       FindTargetFunc   // 获取攻击目标的方式
	_GetAttack         GetAttackFunc    // 计算攻击值
	_EventTrigger      SetEventTrigger  // 计算必定发生的战斗事件
	_BuffToBuddy       BuffToBuddyFunc  // 对伙伴施加buff
	_BuffToSelf        BuffToSelfFunc   // 攻击后对自己释放buff
	_BuffToTarget      BuffToTargetFunc // 对被攻击者施加buff
	AttackNum          int              // 攻击次数
	AppendSpecialType  int              // 附加特殊属性
	AppendSpecialValue float64          // 附加特殊属性值
	TriggerTargetBuff  bool             // 目标Buff仅在触发事件生效时起作用
	CallEnemys         []CallInfo       //技能召唤出来的怪物组

	//TODO 技能类型统一使用一个字段
	IsTotemSkill bool // 图腾技能
	IsRoleSkill  bool // 伙伴技能技能
	IsGhostSkill bool // 魂侍技能技能
}

type GhostSkill struct {
	SkillId            int                  // 模板数据技能ID
	ChildType          int                  // 绝招类型
	NotMiss            bool                 // 是否不会miss
	Critial            float64              // 爆击
	ReduceDefend       float64              // 无视防御
	SunderAttack       int                  // 破甲值加值
	HurtAdd            float64              // 伤害增加值
	HurtAddRate        float64              // 伤害增加百分比
	CureAdd            float64              // 治疗量加值
	CureAddRate        float64              // 治疗增加百分比
	BuffKeepAdd        int                  // buff持续回合增加
	AttackMode         int                  // 攻击模式（范围、单体）
	RecordCount        bool                 // 是否记录攻击次数
	GhostAddRate       int                  // 魂侍攻击、破甲、治疗的叠加转换率
	AttackRangeRatio   float64              // 攻击范围系数(根据攻击范围决定，比如单体、全体等)
	_FindTargets       FindTargetFunc       // 获取攻击目标的方式
	_GetAttack         GetAttackFunc        // 计算攻击值
	_BuffToBuddy       GhostBuffToBuddyFunc // 对伙伴施加buff
	_BuffToSelf        GhostBuffToSelfFunc  // 攻击后对自己释放buff
	_BuffToTarget      BuffToTargetFunc     // 对被攻击者施加buff
	OverrideTargetBuff bool                 // 是否覆盖buff
	OverrideSelfBuff   bool                 // 是否覆盖buff
	OverrideBuddyBuff  bool                 // 是否覆盖buff
	AttackNum          int                  // 攻击次数
}

//type GhostSkill struct {
//	SkillId           int                  // 模板数据技能ID
//	ChildType         int                  // 绝招类型
//	NotMiss           bool                 // 是否不会miss
//	Critial           float64              // 爆击
//	ReduceDefend      float64              // 无视防御
//	SunderAttack      int                  // 破甲值
//	LevelSunderAttack int                  // 绝招等级破甲参数
//	AttackMode        int                  // 攻击模式（范围、单体）
//	RecordCount       bool                 // 是否记录攻击次数
//	GhostAddRate      int                  // 魂侍攻击、破甲、治疗的叠加转换率
//	AttackRangeRatio  float64              // 攻击范围系数(根据攻击范围决定，比如单体、全体等)
//	_FindTargets      FindTargetFunc       // 获取攻击目标的方式
//	_GetAttack        GetAttackFunc        // 计算攻击值
//	_BuffToBuddy      GhostBuffToBuddyFunc // 对伙伴施加buff
//	_BuffToSelf       GhostBuffToSelfFunc  // 攻击后对自己释放buff
//	_BuffToTarget     BuffToTargetFunc     // 对被攻击者施加buff
//}

// 普通攻击的SKillInfo，用来获得一个0值的SkillForce
var normalSkillInfo = &SkillInfo{}

var dummyGhostSkill = &GhostSkill{}

var swordSoulSkillInfo = &SkillInfo{}

// 主角和怪的普通攻击
var normalAttack = &Skill{
	SkillId:          DEFAULT_SKILL_ID,
	IncPower:         0,
	SunderAttack:     5,
	AttackRangeRatio: 1,
	AttackMode:       SKILL_ATTACK_MODE_SINGLE,
	_FindTargets:     findOneTarget,
	_GetAttack: func(f *Fighter, _ float64, _ int16) float64 {
		return f.Attack
	},
}

// 伙伴普通攻击
var buddyNormalAttack = &Skill{
	SkillId:          DEFAULT_SKILL_ID,
	IncPower:         2,
	SunderAttack:     5,
	AttackRangeRatio: 1,
	AttackMode:       SKILL_ATTACK_MODE_SINGLE,
	_FindTargets:     findOneTarget,
	_GetAttack: func(f *Fighter, _ float64, _ int16) float64 {
		return f.Attack
	},
}

var swordSoulAttack = &Skill{
	SkillId:          SWORD_SOUL_SKILL_ID,
	IncPower:         0,
	SunderAttack:     0,
	ReduceDefend:     1,
	AttackRangeRatio: 1,
	AttackMode:       SKILL_ATTACK_MODE_AOE,
	_FindTargets:     findAllTargets,
	_GetAttack: func(f *Fighter, _ float64, _ int16) float64 {
		return float64(f.SwordSoulValue)
	},
}
