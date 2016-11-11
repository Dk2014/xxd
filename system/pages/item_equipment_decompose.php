<?php

include "item_equipment_links.php";

$pconfig = array(
	'title'   => '装备分解',
	'table'   => 'equipment_decompose',
	'links'   => $links,
	
	'columns' => array(
		array('name' => 'quality', 'text' => '品质', 'width' => '80px'),
		array('name' => 'level', 'text' => '等级下限', 'width' => '80px'),
		array('name' => 'fragment_num', 'text' => '获得部位碎片数量', 'width' => '80px'),
		array('name' => 'crystal_id', 'text' => '获得结晶', 'width' => '80px'),
		array('name' => 'crystal_num', 'text' => '获得结晶数量', 'width' => '80px'),
		array('name' => 'dragon_ball', 'text' => '获得龙珠', 'width' => '80px'),
		array('name' => 'dragon_ball_num', 'text' => '获得龙珠数量', 'width' => '80px'),
	),
	'js'			=> 'jsFunction',
);
