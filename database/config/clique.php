<?php
$config = array(
	//!!!建筑特权常量 请勿移动!!!
	array( 'const', ),
	//!!!建筑特权常量 请勿移动!!!
	
	//!!!镖船常量 请勿移动!!!
	array('const',),
	//!!!镖船常量 请勿移动!!!
	
	array( 'const',
		//创建帮派消耗元宝
		"CREATE_CLIQUE_COST" => 100, 

		//聊天信息保留上限
		"MAX_RESORE_MESSAGE_NUM" => 30, 

		//帮主和副帮主可以消耗元宝（20元宝）在世界频道发送招募公告，玩家点击公告可以开启帮派申请界面
		"CLIQUE_RECRUITMENT_COST" => 20, 

		// 帮派建筑物最大等级
		"CLIQUE_BUILDING_MAX_LEVEL" => 12, 

		//帮主可弹劾离线时间
		"CLIQUE_OWNER_MACX_OFFLINE_TIME" => 13 * 86400, 

		//钱庄劵种类
		"CLIQUE_BUILDING_BANK_SILVER" => 0, // 银劵
		"CLIQUE_BUILDING_BANK_GOLD" => 1, // 金卷

		"CLIQUE_BUILDING_BANK_SOLD_TIMESPAN" => 24 * 3600, // 出售间隔时间

		//超过24小时的申请入帮申请删除
		"CLIQUE_TIME_OUT_APPLY" => 86400,

		//招募CD时间
		"CLIQUE_RECRUITMENT_CD" => 3 * 60, 

		//白檀香 500铜钱
		"WHITESANDALWOOD" => 500,
		//苏合香  30元宝
		"STORAX" => 30 ,
		//天木香 90 元宝
		"DAYS" => 90,
		
		//上香到8次可以领取祈福奖励
		"WORSHCOUNT" => 8,
		//祈福奖励基数
		"PRAY_BASE" => 2000,

		// 帮派贡献汇率 money / CLIQUE_CONTRIB_RATE == contrib
		"CLIQUE_CONTRIB_RATE" => 10,

		// 帮派贡献汇率 money / CLIQUE_FAME_RATE == fame
		"CLIQUE_FAME_RATE" => 100,

		// 帮派武功等级限制：不能超过对应建筑等级*10
		"CLIQUE_KONGFU_LEVEL_LIIMT_FACTOR" => 10,

		// 帮派每日上线限制
		"CLIQUE_WORSHIP_DAILY_NUM" => 8,
		// 客户端每页展现条数
		"CLIQUE_CLIENT_NUM" => 8,

		//5个小时秒数
		"CLIQUE_DAILY_QUEST_REFRESHTIME" => 18000,
		//'0'=>'上香', '1'=>'护送运镖', '2'=>'铜钱捐献'
		"CLIQUE_DAYLY_QUEST_WORSHIP" => 0,
		"CLIQUE_DAYLY_QUEST_SEVERAL" => 1,
		"CLIQUE_DAYLY_QUEST_DONATE" => 2,
		// 0 无可领；1可领取; 2已奖励)
		"CLIQUE_QUEST_STATUS_NO_AWARD" =>0,
		"CLIQUE_QUEST_STATUS_CAN_AWARD" =>1,
		"CLIQUE_QUEST_STATUS_RECIVED_AWARD" =>2,
		//领取状态  0 成功， 1,失败 2,可以继续领取
		"CLIQUE_AWARD_STATUS_SUCCESS" =>0,
		"CLIQUE_AWARD_STATUS_FAILED" =>1,
		"CLIQUE_AWARD_STATUS_CONTINUE" =>2,
	),
	array( 'const',
		'KONGFU_ATTRIB_HEALTH' => 1,
		'KONGFU_ATTRIB_ATTACK' => 2,
		'KONGFU_ATTRIB_DEFANCE' => 3,
	),
	array( 'const',
		'CLIQUE_LIST_PAGE_SIZE' => 8,
		),
	array( 'const',
		'HIJACK_TIME' => 10 * 60 ,  //劫持需要10分钟
		'ESCORT_TIME' => 20 * 60 ,  //护送需要20分钟
		//'HIJACK_TIME' => 5 * 60 ,  //劫持需要5分钟
		//'ESCORT_TIME' => 10 * 60 ,  //护送需要10分钟

		'ESCORT_DAILY_NUM' => 3 ,  //每天可以运镖3次
		'HIJACK_DAILY_NUM' => 2 ,  //每天可以劫镖2次
		'ESCORT_START_HOUR' => 9 * 3600 ,  //运镖开始时间
		'ESCORT_END_HOUR' => 23.5 * 3600  ,  //运镖结束时间
		'HIJACKED' => 1 ,  //player_global_escort.hijacked 
		'INGOT_BOAT_ID' => 5 ,  //宝船ID
		'INGOT_BOAT_REQUIRE_VIP_LEVEL' => 1 ,  //宝船要求VIP等级
	),
);

$db = db_connect();
$privileges = db_query($db, 'select * from clique_building ');
foreach($privileges as $p) {
	 $config[0][ 'CLIQUE_BUILDING_' . $p['biaoshi']] = intval($p['id']);
}
$boats = db_query($db, 'select * from clique_boat ');
foreach($boats as $boat) {
	$config[1]['Boat_' . $boat['sign']] = intval($boat['id']) ;
}

?>
