<?php
$config = array(
	array( 'const',
    'BAG_GHOST_NUM'       => 36, // 魂侍背包所放置的魂侍数量
    'BAG_GHOST_POS_START' => 0,  // 魂侍背包位置初始值
    'BAG_GHOST_POS_END'   => 35, // 魂侍背包位置结束值
   
	'COLOR_BLUE'   => 0, // 品质 - 蓝色
	'COLOR_PURPLE' => 1, // 品质 - 紫色
	'COLOR_GOLD'   => 2, // 品质 - 金色
	'COLOR_ORANGE' => 3, // 品质 - 橙色 

	'EQUIP_POS1' => 0,   // '装备位置0'
	'EQUIP_POS2' => 1,   // '装备位置1'
	'EQUIP_POS3' => 2,   // '装备位置2'
	'EQUIP_POS4' => 3,   // '装备位置3'

	'ATTRIBUTE_HEALTH' => 1,
	'ATTRIBUTE_ATTACK' => 2,
	'ATTRIBUTE_DEFEND' => 3,
	'ATTRIBUTE_SPEED'  => 4,
	
	'EQUIPPED' => -1,  // 已装备
	'UNEQUIPPED' => 0,  // 未装备

	'NO_ROLE' => 0, //没有角色使用改魂侍
	
	//!!! DESERT 
	'FOR_MAIN_ROLE' => -1, //主角使用

	'GHOST_DECOMPOSE_NUM' => 10, // 魂侍兑换碎片数目

	'MAX_GHOST_LEVEL' => 100, //魂侍最高等级

	'SHADOW_NUT_PRICE' => 10, //影届果实价格
	'GHOST_FRAGMENT_PRICE' => 35, //魂侍碎片价格

	//魂侍元宝培养碎片价格
	'COLOR_BLUE_FRAME_PRICE' => 40, 
	'COLOR_PURPLE_FRAME_PRICE' => 50, 
	'COLOR_GOLD_FRAME_PRICE' => 60, 
	'COLOR_ORANGE_FRAME_PRICE' => 9999,  //暂时没有橙色魂侍

	'FLUSH_ATTR_CD' => 30*60, //魂侍洗点冷却时间
	'FlUSH_ATTR_BACK_PERCENT' => _S(1.0, 1.0, 0.8, 0.8), //魂侍洗点返还百分比

	'RELATION_GHOST_ADD_PER' => 0.1, // 连锁魂侍加成

        //魂侍洗炼const
        'BAPTIZE_FRAME_NUM' => 10, // 魂侍洗炼消耗碎片数量
        'BAPTIZE_PLAYER_MIN_LEVEL' => 80, // 魂侍洗炼玩家最低等级
	),
); 
?>
