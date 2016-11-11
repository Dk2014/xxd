<?php
include "item_equipment_links.php";

$pconfig = array(
	'title'   => '装备共鸣',
	'table'   => 'equipment_resonance',
	'links'   =>$links,
        	
	'columns' => array(
		array('name' => 'require_level', 'text' => '精炼等级', 'width' => '150px'),
		array('name' => 'health', 'text' => '生命', 'width' => '40px'),
		array('name' => 'attack', 'text' => '攻击', 'width' => '40px'),
		array('name' => 'defence', 'text' => '防御', 'width' => '40px'),

	),
	
	'where' => 'sql_where',
);
