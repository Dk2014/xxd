<?php
$config = array(
	array( 'const',

		'LEVEL_BOX_AWARD_COUNT' => 2, // 默认奖励开宝箱次数
		'MAX_LEVEL_BOX_AWARD_COUNT' => 3, // 关卡奖励开宝箱次数
		'RANDDOM_BOX_AWARD_COUNT' => 1, //开随机奖励宝箱次数
		'LEVEL_PHYSICAL_MIN' => 3, // 区域关卡最少扣除的体力
		'LEVEL_HARD_PHYSICAL_MIN' => 1, //深渊关卡失败扣除的体力

		'LEVEL_BOX_AWARD_NIL' 		=> -1, // 非法奖励类型
		'LEVEL_BOX_AWARD_COIN' 		=> 0,  // 铜钱
		'LEVEL_BOX_AWARD_ITEM' 		=> 1,  // 道具
		'LEVEL_BOX_AWARD_EQUIMENT' 	=> 2,  // 装备
		'LEVEL_BOX_AWARD_EXP' 		=> 3,  // 经验	
		'LEVEL_BOX_AWARD_MULTIPLE_EXP' 	=> 4,  // 经验倍数
		'LEVEL_BOX_AWARD_MULTIPLE_COIN' => 5,  // 铜钱倍数
		'LEVEL_BOX_AWARD_PET_BALL' => 6,  // 契约球
	),

	array( 'const',
		'EXTEND_LEVEL_TYPE_RESOURCE' => 1,
		'EXTEND_LEVEL_TYPE_BUDDY' => 9,
		'EXTEND_LEVEL_TYPE_PET' => 10,
		'EXTEND_LEVEL_TYPE_GHOST' => 11,
		),


	array('data', 'ExtendLevelCoinOpenDay' => '[]int', 1, 2,3,4, 5, 6,0), // 星期1，3，5，星期天
	array('data', 'ExtendLevelExpOpenDay' => '[]int', 1,2,3, 4,5, 6, 0), // 星期2，4，6，星期天
	array('data', 'ExtendLevelBuddyOpenDay' => '[]int', 3, 6, 0), 
	array('data', 'ExtendLevelPetOpenDay' => '[]int', 2, 5, 0),
	array('data', 'ExtendLevelGhostOpenDay' => '[]int', 1, 4, 0),

	array( 'const',
		'RESOURCE_COIN_LEVEL' => 1,
		'RESOURCE_EXP_LEVEL' => 2,

		'EXTeND_LEVEL_PET_MAX_ROUND_NUM' => 30, // 灵宠关卡允许的最大回合数
		
		'EXTEND_LEVEL_PASS_CD_TIME' => 10*60, // 10min

		'EXTEND_LEVEL_BUDDY_MAX_DAILY_NUM' => 5, // 伙伴关卡每日挑战上限
		'EXTEND_LEVEL_RESOURCE_MAX_DAILY_NUM' => 2, // 资源关卡每日挑战上限
		'EXTEND_LEVEL_PET_MAX_DAILY_NUM' => 5, // 灵宠关卡每日挑战上限
		'EXTEND_LEVEL_GHOST_MAX_DAILY_NUM' => 5, // 魂侍关卡每日挑战上限
	),

	array( 'const',
		'ONE_STAR' 	=> 1,
		'TWO_STAR' 	=> 2, 
		'Three_STAR' 	=> 3,
	),

	array( 'const',
		'LEVEL_TYPE_MISSION' 	=> 	0, //区域关卡
		'LEVEL_TYPE_HARD'	=> 	8, //难度关卡
	),

	array( 'const',
		'NORMAL' => 0, //普通
		'ELITE' =>  1, //精英
		'BOSS' =>   2, //BOSS
		'ITEM_PET_BALL' => 250, //契约球
		'ITEM_PET_BALL_RATE' => 50 //契约球捕捉成功率
	),

	array( "const", 
		"FREE_RELIVE_LEVEL" => 20, // 免费复活的等级
		"FIRST_RELIVE_INGOT" => 10, // 第一次复活需要的元宝数

	),
	array("const",
	'RECOVER_BATTLE_PET'     => 1, //恢复灵宠
	'RECOVER_MEMBERS_HELATH' => 2, //恢复队友生命
	'RECOVER_SKILL_RELEASE'  => 3, //恢复伙伴技能
	),

	array( "const", 
		"PVE_LEVEL_DAILY_NUM" => 1, // 每日次数
		"PVE_LEVEL_INIT_ENEMY_GROUP_NUM" => 2, // 灵兽关卡初始加载怪物组数量
		"PVE_LEVEL_MAX_ROUND" => 15, // 灵兽关卡最大回合数
	),
	array( "const",
		"HARD_LEVEL_BUY_INGOT" => 50,
		"HARD_LEVEL_MAX_BUY_TIME" => 1,
	),
	array( "const",
		"CHEST_TYPE_COPPER" => 1,
		"CHEST_TYPE_SILVER" => 2,
		"CHEST_TYPE_GOLD" =>3,
	)

);
?>
