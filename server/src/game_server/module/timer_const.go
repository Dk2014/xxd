package module

/*
  所有定时器相关的定义都在这里进行
*/

const (
	TIMER_NONE = iota + 1

	// 在下面添加定时器标识

	TIMER_PHYSICAL            // 体力恢复定时器
	TIMER_HEART               // 爱心恢复定时器
	TIMER_MULTI_LEVEL_AUTO    // 多人关卡随机
	TIMER_SWORD_SOUL          // 剑山拔剑次数定时恢复
	TIMER_DAILY_QUEST         // 每日任务零点刷新
	TIMER_DRAW_CHEST_BY_CONIS // 免费青龙宝箱
	TIMER_DRAW_CHEST_BY_INGOT // 免费神龙宝箱
	TIMER_ESCORT_START        // 开始运镖
	TIMER_HIJACK_START        // 开始劫镖
)
