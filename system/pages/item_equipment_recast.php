<?php

include "item_equipment_links.php";

$pconfig = array(
	'title'   => '装备重铸',
	'table'   => 'equipment_recast',
	'links'   => $links,
	
	'columns' => array(
		array('name' => 'level', 'text' => '等级下限', 'width' => '80px'),
		array('name' => 'quality', 'text' => '品质', 'width' => '80px'),
		array('name' => 'fragment_num', 'text' => '需要部位碎片数量', 'width' => '80px'),
		array('name' => 'blue_crystal_num', 'text' => '需要蓝色结晶数量', 'width' => '80px'),
		array('name' => 'purple_crystal_num', 'text' => '需要紫色结晶数量', 'width' => '80px'),
		array('name' => 'golden_crystal_num', 'text' => '需要金色结晶数量', 'width' => '80px'),
		array('name' => 'orange_crystal_num', 'text' => '需要橙色结晶数量', 'width' => '80px'),
		array('name' => 'consume_coin', 'text' => '消耗的铜钱', 'width' => '80px'),
	),
);