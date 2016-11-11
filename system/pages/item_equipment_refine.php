<?php

include "item_equipment_links.php";

$pconfig = array(
	'title'   => '装备精练',
	'table'   => 'equipment_refine',
	'links'   => $links,
	
	'columns' => array(
		array('name' => 'level', 'text' => '等级下限', 'width' => '80px'),
		array('name' => 'quality', 'text' => '品质', 'width' => '80px'),
		array('name' => 'fragment_num', 'text' => '需要部位碎片数量', 'width' => '80px'),
		array('name' => 'blue_crystal_num', 'text' => '需要蓝色结晶数量', 'width' => '80px'),
		array('name' => 'purple_crystal_num', 'text' => '需要紫色结晶数量', 'width' => '80px'),
		array('name' => 'golden_crystal_num', 'text' => '需要金色结晶数量', 'width' => '80px'),
		array('name' => 'orange_crystal_num', 'text' => '需要橙色结晶数量', 'width' => '80px'),
		array('name' => 'level1_consume_coin', 'text' => '精练到1级消耗的铜钱', 'width' => '80px'),
		array('name' => 'level2_consume_coin', 'text' => '精练到2级消耗的铜钱', 'width' => '80px'),
		array('name' => 'level3_consume_coin', 'text' => '精练到3级消耗的铜钱', 'width' => '80px'),
		array('name' => 'level4_consume_coin', 'text' => '精练到4级消耗的铜钱', 'width' => '80px'),
		array('name' => 'level5_consume_coin', 'text' => '精练到5级消耗的铜钱', 'width' => '80px'),
		array('name' => 'level6_consume_coin', 'text' => '精练到6级消耗的铜钱', 'width' => '80px'),
		array('name' => 'level7_consume_coin', 'text' => '精练到7级消耗的铜钱', 'width' => '80px'),
		array('name' => 'level8_consume_coin', 'text' => '精练到8级消耗的铜钱', 'width' => '80px'),
		array('name' => 'level9_consume_coin', 'text' => '精练到9级消耗的铜钱', 'width' => '80px'),
		array('name' => 'level10_consume_coin', 'text' => '精练到10级消耗的铜钱', 'width' => '80px'),
	),
);
