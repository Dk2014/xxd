<?php

$pconfig = array(
	'title'   => '影之间隙',
	'table'   => 'shaded_mission',
	'links'   => array(),
	'columns' => array(
		array('name' => 'name', 'text' => '名称', 'width' => '150'),
		array('name' => 'order', 'text' => '位序', 'width' => '50'),
		
		array('name' => 'sign', 'text' => '场景资源标识', 'width' => '120'),
		array('name' => 'sign_war', 'text' => '战斗资源标识', 'width' => '120'),
		array('name' => 'flip_horizontal', 'text' => '水平翻转', 'width' => '50'),
		array('name' => 'music', 'text' => '音乐资源标识', 'width' => '150'),

		array('name' => 'enter_x', 'text' => '出生点x', 'width' => '50'),
		array('name' => 'enter_y', 'text' => '出生点y', 'width' => '50'),
		array('name' => 'exit_x', 'text' => '退出传送x', 'width' => '50'),
		array('name' => 'exit_y', 'text' => '退出传送y', 'width' => '50'),
		array('name' => 'mission_link_x', 'text' => '区域关卡入口x', 'width' => '50'),
		array('name' => 'mission_link_y', 'text' => '区域关卡入口y', 'width' => '50'),
		array('name' => 'mission_back_x', 'text' => '区域关卡返回x', 'width' => '50'),
		array('name' => 'mission_back_y', 'text' => '区域关卡返回y', 'width' => '50'),

		array('name' => 'box_x', 'text' => '宝箱x', 'width' => '50'),
		array('name' => 'box_y', 'text' => '宝箱y', 'width' => '50'),
		array('name' => 'box_dir', 'text' => '宝箱朝向', 'width' => '50'),

		array('name' => 'award_exp', 'text' => '奖励经验', 'width' => '50'),
		array('name' => 'award_coin', 'text' => '奖励铜钱', 'width' => '50'),
		array('name' => 'award_item1', 'text' => '奖励物品1', 'width' => '50'),
		array('name' => 'award_item1_num', 'text' => '奖励物品1数量', 'width' => '50'),
		array('name' => 'award_item2', 'text' => '奖励物品2', 'width' => '50'),
		array('name' => 'award_item2_num', 'text' => '奖励物品2数量', 'width' => '50'),
		array('name' => 'award_item3', 'text' => '奖励物品3', 'width' => '50'),
		array('name' => 'award_item3_num', 'text' => '奖励物品3数量', 'width' => '50'),
	),
	'opera' => array(
		array('type' => 'link', 'text' => '地图怪物', 'render' => 'map_enemy_render'),
	),
	'insert' => 'sql_insert',
	'where' => 'sql_where',
	'js'=>'js_function',
);

