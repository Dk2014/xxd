<?php
$config = array(
	array('const',
		/*货币类型*/
		'COINS' => 0, //金币
		'INGOT' => 1, //元宝
		'SWORDSOULFRAGMENT'=>2,//剑心碎片

		// 功能开放是否需要广播
		'NO_NEED_PLAY' => 0,
		'NEED_PLAY'    => 1,
	),
	array('const',
		/*直接赠送接口*/
		'TPRESENT_DEFAULT_DISCOUNTID' => 'UM140805211423554', //默认活动ID
		'TPRESENT_DEFAULT_GIFTID'     => '2964073651PID201408052114235636', //默认礼包ID
	),

	array('const',
		'RESET_ARENA_TIMES_IN_HOUR'         => 5, // 比武场次数重置
		'RESET_ARENA_AWARD_BOX_IN_HOUR'     => 21, // 每天比武场宝箱结算时间为21点
		'RESET_GHOST_LEVEL_TIMES_IN_HOUR'   => 5, // 魂侍试炼次数重置
		'RESET_PET_LEVEL_TIMES_IN_HOUR'     => 5,
		'RESET_BUDDY_LEVEL_TIMES_IN_HOUR'   => 5,
		'RESET_EXP_LEVEL_TIMES_IN_HOUR'     => 5,
		'RESET_COIN_LEVEL_TIMES_IN_HOUR'    => 5,
		'RESET_RAINBOW_LEVEL_TIMES_IN_HOUR' => 5,
		'RESET_HARD_LEVEL_TIMES_IN_HOUR'    => 5,
		'RESET_MAIN_LEVEL_TIMES_IN_HOUR'    => 5,
		'RESET_MULTI_LEVEL_TIMES_IN_HOUR'   => 5,
		'RESET_DAILY_QUEST_IN_HOUR'         => 5,
		'RESET_COIN_BOX_TIMES_IN_HOUR'      => 5,
		'RESET_INGOT_BOX_TIMES_IN_HOUR'     => 5,
		'RESET_HEART_DRAW_TIMES_IN_HOUR'    => 5,
		'RESET_BUY_PHYSICAL_TIMES_IN_HOUR'  => 5,
		'RESET_PLAYER_ACTIVITY_IN_HOUR'     => 5, //玩家活跃度重置时间
		'RESET_PVE_LEVEL_IN_HOUR'     => 5, //灵宠环境重置时间
		'RESET_GHOST_TRAIN_INFO_IN_HOUR'     => 5, //重置魂侍元宝培养次数
		'RESET_BATTLE_PET_UPGRADE_INFO_IN_HOUR'     => 5, //重置灵宠元宝培升阶数
		'RESET_BUY_HARD_LEVEL_TIMES_IN_HOUR' => 5, //重置玩家购买深渊关卡次数的时间点
		'RESET_BUY_TOTEM_INGOT_CALL_TIMES_IN_HOUR' => 5, //重置玩家元宝购买阵印次数时间点
		'RESET_DRIVING_SWORD_ACTION_TIMES_IN_HOUR' => 5, //重置玩家云海御剑行动点时间
		'RESET_BUY_RESOURCE_LEVEL_TIMES_IN_HOUR' => 5, //重置玩家购买资源关卡次数的时间点
		'RESET_BUY_BOSS_LEVEL_TIMES_IN_HOUR' => 5, //重置玩家购买BOSS关卡次数的时间点
	),

	array('const',
		'MAX_PLAYER_ACTIVITY' => 1500, //玩家活跃度上限
	),

	array('const',
		'ARENA_LOSE_AWARD_FAME' => 5, //比武场战败获得5点声望
		'ARENA_WIN_AWARD_FAME' => 10, //比武场胜利获得10点声望

	),

	array('const',
		'MONTH_CARD_PER_TIMES' => 30, //每次月卡信息对应的天数
	),

	array('const',
		'FUNC_OPEN_TYPE_BY_LOCK' => 1, //功能开放方式：权值
		'FUNC_OPEN_TYPE_BY_LEVEL' => 2, //功能开放方式：等级
	),

	array('const',
		'MONTH_CARD_DURATION' => 86400*30, //月卡有效时间 30 自然日
		'MONTH_CARD_RENEW_TIME_BEFORE_EXPIRE' => 86400*3, //月卡过期前X秒可续
	),

	array('const',
		'RANK_EVERY_PAGE_NUM' => 7,    //排行榜每页7条记录
		'RANK_OPEN_MIN_LEVEL' => 15,   //排行榜开启等级
        'RANK_INGOT_DESPLAY_NUM'=>100, //充值排行榜总共显示前100名玩家
	),
);
?>
