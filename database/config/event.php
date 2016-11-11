<?php
$config = array(
	array('const',
	'MAX_LOGIN_AWARD_DAY' => 7, //等级奖励最多累计天数
	),

	array('const',
	'NOT_SEND_MAIL' => 0,
	'SEND_MAIL' =>1,
	),	
	
	array('const',
	'LOGIN_AWARD_TYPE_ITEM' => 1,
	'LOGIN_AWARD_TYPE_HEART' => 2,
	'LOGIN_AWARD_TYPE_COINS' => 3,
	'LOGIN_AWARD_TYPE_INGOT' => 4,
	'LOGIN_AWARD_TYPE_SWORD_SOUL' => 5,
	'LOGIN_AWARD_TYPE_GHOST' =>  6,
	'LOGIN_AWARD_TYPE_PET' => 7,
	),
	
	//活动列表
	array('const',),
	
	//活动奖励方式 0:活动中心领取 1:跳转 2:纯文字
	array('const',
	'NOT_GO'=>0,
	'YES_GO'=>1,
	'NEVER_GO'=>2,
	),
	
	array('const',
	'NOT_DISPOSE'=>0,
	'NOT_END' => 1,
	),
	
	//午餐活动的时间点
	array('const',
	'DINNER_DISPLAY_IN_HOUR' => 8,
	'DINNER_START_IN_HOUR' => 12,
	'DINNER_END_IN_HOUR' => 14,
	
	'SUPPER_DISPLAY_IN_HOUR' => 14,
	'SUPPER_START_IN_HOUR' => 18,
	'SUPPER_END_IN_HOUR' => 20,
	),

	//月卡id
	array('const',
		'MONTH_CARD_PRODUCTID'=> '1100',
	),
	
	//加成配置
	array('const',
		'MISSION_COINS' => 1,
		'MISSION_EXPS' => 2,
		'QQVIP_COINS' => 3,
		'QQVIP_EXPS' => 4,
		'SUPERQQ_COINS' => 5,
		'SUPERQQ_EXPS' => 6,
	),
	
	
	array('const',
		'SRD_LOCK' => 120130, //水溶洞一的权值 通关后开启七天登录活动 
		'MAX_LOGIN_AWARDS_DAY' => 5, //累计登录活动 5天之后改邮件发送
	),

	array('const',
		'MAX_SHARE_TIMES' => 30, //最大的有效分享奖励次数
		),
	array('const',
		'INVALIDINDEX' => -1, //每日首冲奖励无匹配奖励
		'CONTINUE' => 0,    //每日活动,本次领取成功，下一天活动继续
		),
	array('const',
		'JSON_EVENT_TEN_DRAW' =>257, // json配置十连抽活动
		'JSON_EVENT_TOTAL_RECHARGE' => 258, // json配置累计充值活动
		'JSON_EVENT_TOTAL_CONSUME' => 259, // json配置累计消费活动
		'JSON_EVENT_FIRST_RECHAGE_DAILY' => 260, // json配置每日小额充值
		'JSON_EVENT_SINGLE_CONSUME' => 261,		// json配置单条重复消耗
		'JSON_EVENT_ARENA_RANK' => 262,			//json配置比武场排名
		'JSON_EVENT_NEW_YEAR' => 263, //json配置新年红包活动
		),
	array('const',
		'EVENT_NEW_YEAR_RATE' => 0.1, //春节红包返回10%
		'EVENT_NEW_YEAR_EMAIL_LIMIT_TIME' => 10, //10个月
		'EVENT_NEW_YEAR_INVAILD_PAGE' => 1, //春节红包有效的期数 下年需要进行修改
		),
	array('const',
		'QQ_VIP' => 0, // QQ会员
		'SUPER_QQ_VIP' => 1, // 超级QQ会员
		)
);
$db = db_connect();
	$events = db_query($db, 'select id,sign from quest_activity_center order by `weight` ASC');
	foreach($events as $event){
		$config[3][$event['sign']] = intval($event['id']);
	}
?>
