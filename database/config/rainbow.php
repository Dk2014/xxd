<?php
$config = array(
	array( 'const',
		
		'MAX_BATTLE_ROUND' => 15, // 允许最大战斗回合数

		'INIT_LEVEL_ORDER' => 1, // 初始段关卡
		'INIT_SEGMENT_NUM' => 1, // 初始段数字

		'LEVEL_NUM_PER_SEGMENT' => 7, // 每段多少个关卡
		'MAX_RAINBOW_SEGMENT_NUM' => 63, // 彩虹关卡最多 63 段
		'MIN_SEGMENT_NUM_CAN_JUMP' => 2, // 最小可跳转关卡

		'RAINBOW_LEVEL_STATUS_NEVER_PASS' => 0, //未打败Boss
		'RAINBOW_LEVEL_STATUS_NEVER_AWARD' => 1, //未领取奖励
		'RAINBOW_LEVEL_STATUS_NO_MORE' => 2, //没有后续关卡

		'DAILY_RESET_NUM' => 1, //每日可重置次数

		'AWARD_CHOICE_NUM' => 5, //奖励选择数
		'AWARD_NUM' => 2, //奖励数
		'DAILY_AUTO_FIGHT_NUM' => 3, //每日扫荡次数3次 如果以后扫荡作为vip特权则将起已到VIP特权次数接口中
	),

	array( 'const',
		'AWARD_TYPE_COIN' 		=> 0,  // 铜钱
		'AWARD_TYPE_ITEM' 		=> 1,  // 道具
		'AWARD_TYPE_EQUIMENT' 	=> 2,  // 装备
		'AWARD_TYPE_EXP' 		=> 3,  // 经验	
		'AWARD_TYPE_MULTIPLE_EXP' 	=> 4,  // 经验倍数
		'AWARD_TYPE_MULTIPLE_COIN' => 5,  // 铜钱倍数
		'AWARD_TYPE_RECOVER_BUDDY_SKILL' => 6,  // 恢复伙伴技能
		'AWARD_TYPE_RECOVER_GHOST' => 7,  // 恢复魂侍技能
		'AWARD_TYPE_RECOVER_PET' => 8,  // 恢复灵宠状态
		'AWARD_TYPE_ADD_POWER' => 9,  // 主角增加精气（
		'AWARD_TYPE_ADD_HEALTH_BY_PERSENT' => 10,  //  百分比恢复生命（上阵）
		'AWARD_TYPE_AAD_GHOST_POWER' => 11,  // 增加魂力（上阵）
	)
);
?>
