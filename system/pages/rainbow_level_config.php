<?php

$pconfig = array(
	'title'   => '战场配置',
	'table'   => 'mission_level',
	'links'   => array(),
	'columns' => array(
		array('name' => 'order', 'text' => '顺序', 'width' => '50'),
		array('name' => 'name', 'text' => '战场名称', 'width' => '150'),
		array('name' => 'type', 'text' => '战场类型', 'width' => '90px',),
		array('name' => 'sign', 'text' => '场景资源标识', 'width' => '120'),
		array('name' => 'sign_war', 'text' => '战斗资源标识', 'width' => '150'),
		array('name' => 'flip_horizontal', 'text' => '水平翻转', 'width' => '50'),
		array('name' => 'music', 'text' => '音乐资源标识', 'width' => '150'),
		array('name' => 'enter_x', 'text' => '出生点x', 'width' => '50'),
		array('name' => 'enter_y', 'text' => '出生点y', 'width' => '50'),
		array('name' => 'box_x', 'text' => '宝箱x', 'width' => '50'),
		array('name' => 'box_y', 'text' => '宝箱y', 'width' => '50'),
		array('name' => 'box_dir', 'text' => '宝箱朝向', 'width' => '50'),
		array('name' => 'award_exp', 'text' => '奖励经验', 'width' => '80'),
		array('name' => 'award_coin', 'text' => '奖励铜钱', 'width' => '80'),
		array('name' => 'award_relationship', 'text' => '奖励友情', 'width' => '80'),
		array('name' => 'award_box', 'text' => '奖励宝箱', 'width' => '80'),
	),
	'opera' => array(
		array('type' => 'link', 'text' => '地图怪物', 'render' => 'map_enemy_render'),
	),
	'insert' => 'sql_insert',
	'where' => 'sql_where',
	//'location' => 'location',
);
?>
