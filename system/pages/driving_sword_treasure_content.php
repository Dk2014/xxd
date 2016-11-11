<?php

$pconfig = array(
	'title'   => '云海御剑宝藏山宝物列表',
	'table'   => 'driving_sword_treasure_content',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'order', 'text' => '奖励次序', 'width' => '90px'),
		array('name' => 'award_item', 'text' => '奖励物品', 'width' => '90px'),
		array('name' => 'award_num', 'text' => '奖励数量', 'width' => '90px'),
		array('name' => 'award_coins', 'text' => '奖励铜币数量', 'width' => '90px'),
	),
	
	'where' => 'sql_where',
	'insert' => 'sql_insert',
	'js' => 'js_function',
);

