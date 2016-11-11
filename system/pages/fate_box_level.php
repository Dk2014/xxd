<?php

$pconfig = array(
	'title'   => '命锁关卡',
	'table'   => 'mission_level',
	'links'   => array(),
	'columns' => array(
		array('name' => 'name', 'text' => '关卡名称', 'width' => '150'),
		array('name' => 'type', 'text' => '关卡类型', 'width' => '90px',),
		array('name' => 'sign', 'text' => '场景资源标识', 'width' => '120'),
		array('name' => 'sign_war', 'text' => '战斗资源标识', 'width' => '120'),
		array('name' => 'flip_horizontal', 'text' => '水平翻转', 'width' => '50'),
		array('name' => 'music', 'text' => '音乐资源标识', 'width' => '150'),
	),
	'opera' => array(
		array('type' => 'link', 'text' => '地图怪物', 'render' => 'map_enemy_render'),
	),
	'insert' => 'sql_insert',
	'where' => 'sql_where',
	'not_delete' => true,
);
?>
