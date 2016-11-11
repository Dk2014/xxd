<?php
include "battle_pet_links.php";


$pconfig = array(
	'title'   => '灵宠绝招训练',
	'table'   => 'battle_pet_skill_training',
	'links'   => $battle_pet_links,
	
	'columns' => array(
		array('name' => 'level', 'text' => '绝招等级', 'width' => '150px'),
		array('name' => 'cost_coins', 'text' => '下一级所需金币数', 'width' => '150px'),
	),

	'opera' => array(),
	
	'where' => 'sql_where',
);

