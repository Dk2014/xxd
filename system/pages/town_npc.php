<?php

$pconfig = array(
	'title'   => '城镇NPC',
	'table'   => 'town_npc',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'npc_role', 'text' => '角色', 'width' => '80px'),
		array('name' => 'name', 'text' => 'NPC名称', 'width' => '80px'),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '80px'),
		array('name' => 'x', 'text' => 'x轴坐标', 'width' => '80px'),
		array('name' => 'y', 'text' => 'y轴坐标', 'width' => '80px'),
		array('name' => 'direction', 'text' => '朝向', 'width' => '80px'),
		array('name' => 'showup_quest', 'text' => '出现任务', 'width' => '200px'),
		array('name' => 'disappear_quest', 'text' => '消失任务', 'width' => '200px'),
		array('name' => 'talk', 'text' => '对话', 'width' => '200px'),

	),
	
	'opera' => array(
	),
	
	'where'	  => 'sql_where',
	'insert'  => 'sql_insert',
	'location' => 'location',
	'js' =>	'jsFunction',
);

?>
