<?php

$pconfig = array(
	'title'   => '灵宠关卡配置',
	'table'   => 'level_battle_pet',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'battle_pet_id', 'text' => '出现灵宠', 'width' => '80px'),

		array('name' => 'rate', 'text' =>'出现概率%', 'width' => '80px'),
		array('name' => 'live_round', 'text' => '出现后存活回合数', 'width' => '80px')
	),
	
	'opera' => array(),
	

	'where' 		=> 'sql_where',
	'insert'  		=> 'sql_insert',
	'location' => 'location',
);

?>