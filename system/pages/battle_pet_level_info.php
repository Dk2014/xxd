<?php
include "battle_pet_links.php";

$pconfig = array(
	'title'   => '灵宠等级数据',
	'table'   => 'battle_pet_level_info',
	'links'   => $battle_pet_links,
        	
	'columns' => array(
		array('name' => 'level', 'text' => '等级', 'width' => '150px'),

		array('name' => 'health', 'text' => '生命', 'width' => '40px'),
		array('name' => 'attack', 'text' => '攻击', 'width' => '40px'),
		array('name' => 'defence', 'text' => '防御', 'width' => '40px'),
		array('name' => 'speed', 'text' => '速度', 'width' => '40px'),

		array('name' => 'sunder_max_value', 'text' => '护甲上限','width' => '60px', 'default' => '100'),
	    array('name' => 'sunder_min_hurt_rate', 'text' => '破甲前百分比','width' => '60px', 'default' => '100'),
        array('name' => 'sunder_end_hurt_rate', 'text' => '破甲后百分比','width' => '60px', 'default' => '150'),
        array('name' => 'sunder_end_defend_rate', 'text' => '破甲后减防百分比','width' => '60px', 'default' => '20'),
	),
	
	'where' => 'sql_where',
    'insert' => 'sql_insert',
);
