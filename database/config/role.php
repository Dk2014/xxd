<?php
$config = array(
	array( 'const',
    "MAIN_ROLE" => 1, // 主角
    "MAX_ROLE_NUM" => 5, // 最大角色数
    "MAX_ROLE_NUM_USING" => 9, // 最大在列伙伴数
    "PLAYER_BOY_ROLE_ID" 	=> 1, // 玩家男主角
    "PLAYER_GIRL_ROLE_ID"	=> 2, // 玩家女主角
    "MAX_ROLE_LEVEL"	=> 140, // 最大等级数
	"MAX_FRISHP_LEVEL"  => 10, // 最大羁绊等级数

	"ROLE_STATUS_NOMAL" => 0, // 正常伙伴状态
	"ROLE_STATUS_ININN" => 1, // 处于客栈的状态

    "FIGHT_FOR_ALL"        => 0, // 战斗属性加成标识
	"FIGHT_FOR_ROLE_LEVEL" => 1, // 角色等级属性标识
	"FIGHT_FOR_SKILL"      => 2, // 绝招加成标识
	"FIGHT_FOR_EQUIP"      => 3, // 装备加成标识
	"FIGHT_FOR_GHOST"      => 4, // 魂侍加成标识
	"FIGHT_FOR_REALM"      => 5, // 境界加成标识
	"FIGHT_FOR_FRIENDSHIP" => 6, // 羁绊加成标识
	"FIGHT_FOR_TEAMSHIP"   => 7, // 伙伴配合加成标识
	),

	//array('switch', 
	//	'num' => 'int', 
	//	'getMainRoleCoopAttributeDemo' => array('health'=>'int','attack' =>'int', 'defence'=>'int', 'speed'=>'int', 'cultivation'=>'int'), 
	//	false,
	//	'1' => array(1,1,1,1,1,1),

	//),
); 
?>
