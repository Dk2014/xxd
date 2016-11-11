<?php

include "item_equipment_links.php";

$pconfig = array(
	'title'   => '特殊装备追加属性',
	'table'   => 'equipment_appendix',
	'links'   => $links,
	
	'columns' => array(
		array('name' => 'level', 'text' => '物品', 'width' => '150px'),
		array('name' => 'health', 'text' => '生命','width' => '50px',),
		array('name' => 'attack', 'text' => '攻击', 'width' => '50px'),
		array('name' => 'defence', 'text' => '防御', 'width' => '50px'),
		array('name' => 'speed', 'text' => '速度', 'width' => '50px'),
		array('name' => 'cultivation', 'text' => '内力', 'width' => '50px'),
		array('name' => 'hit_level', 'text' => '命中', 'width' => '50px'),
		array('name' => 'dodge_level', 'text' => '闪避', 'width' => '50px'),
		array('name' => 'critical_level', 'text' => '暴击', 'width' => '50px'),
		array('name' => 'tenacity_level', 'text' => '韧性', 'width' => '50px'),
		array('name' => 'block_level', 'text' => '格挡', 'width' => '50px'),
		array('name' => 'destroy_level', 'text' => '破击', 'width' => '50px'),
	),

	'where' 		=> 'sql_where',
);
?>