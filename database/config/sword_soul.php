<?php
$config = array(
	array('const',
			
		'MAX_DRAW_NUM' => 100,  	// 最大拔剑次数
		'MAX_INGOT_DRAW_NUM' => 20,  	// 最大拔剑次数
		'RECOVERY_VALUE' => 1, 		//定时回复次数
		'RECOVERY_TIME' => 60*5 , 	//定时回复时间间隔 五分钟
		
		'BAG_ON_BODY_NUM'   => 9,  // 用于穿戴的背包数量
		'BAG_START'         => 0,  // 剑心背包的起始序号
		'BAG_END'           => 35, // 剑心背包的结束序号
		'BAG_NUM'           => 36, // 剑心背包数量
		
		'BOX_A'   => 0,  // 箱子A序号
		'BOX_B'   => 1,  // 箱子B序号
		'BOX_C'   => 2,  // 箱子C序号
		'BOX_D'   => 3,  // 箱子D序号
		'BOX_E'   => 4,  // 箱子E序号
		'BOX_NUM' => 5,  // 箱子数目
		
		'QIAN_LONG_ID' => 37,  // 潜龙剑心

		'RUBBISH_AWARD_COIN' => 500, // 拔剑没有获得东西，奖励的铜钱
	),

	// 剑心箱配置
	array( 'type', 'box_config',
		'key'  => 'int8',    // 箱子key
		'coin' => 'int64',   // 消耗铜钱
		'name' => 'client',  // 剑名
	),

	array( 'data', 'box_configs' => '[]box_config',
		array(1, 1000, "败亡铁剑"),
		array(2, 2000, "败亡阔剑"),
		array(4, 3000, "败亡魔剑"),
		array(8, 4000, "败亡宝剑"),
		array(16, 8000, "败亡神剑"),
	),

	// 等级对应剑心装备数
	array('between', 'level' => 'int16', 'GetOpenedPosNum' => 'int', 'assert'=>true, 'default' => 1,
		array('(,9]', 1),
		array('[10,19]', 2),
		array('[20,29]', 3),
		array('[30,39]', 4),
		array('[40,49]', 5),
		array('[50,59]', 6),
		array('[60,69]', 7),
		array('[70,79]', 8),
		array('[80,)', 9),
	),

	//角色等级对应剑心最高等级限制
	array('between', 'level' => 'int16', 'SwordSoulMaxLevel' => 'int8', 'assert'=> true, 'default'=> 1,
		array('(,20]', 4),
		array('[21,40]', 8),
		array('[41,60]', 12),
		array('[61,80]', 16),
		array('[81,)', 20),
	),

	//元宝拔剑元宝价格
	array('between', 'times' => 'int64', 'GetSwordDrawPriceInIngot' => 'int64', 'assert' => true, 'default' => 0,
		array('[1,2]', 50),
		array('[3,5]', 75),
		array('[6,10]', 100),
		array('[11,20]', 200),
	),

	//元宝拔剑铜钱价格
	array('between', 'times' => 'int64', 'GetSwordDrawPriceInIngotCoin' => 'int64', 'assert' => true, 'default' => 0,
		array('[1,2]', 4000),
		array('[3,5]', 4000),
		array('[6,10]', 4000),
		array('[11,20]', 4000),
	),
);
?>
