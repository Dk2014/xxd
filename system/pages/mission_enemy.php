<?php

$pconfig = array(
	'title'   => '关卡怪物',
	'table'   => 'mission_enemy',
	'links'   => array(),
	'location'=> 'location',
	'js' => 'jsFunction',
	'columns' => array(
		array('name' => 'order', 'text' => '顺序', 'width' => '50px'),
		array('name' => 'best_round', 'text' => '最好的通关回合数', 'width' => '50px'),
		array('name' => 'monster_num', 'text' => '上阵怪物数量', 'width' => '50px'),
		array('name' => 'enter_x', 'text' => '出生点x坐标', 'width' => '50px'),
		array('name' => 'enter_y', 'text' => '出生点y坐标', 'width' => '50px'),
		array('name' => 'monster1_id', 'text' => '怪物1 ID', 'width' => '80px'),
		array('name' => 'monster1_chance', 'text' => '1出现概率%', 'width' => '50px'),
		array('name' => 'monster2_id', 'text' => '怪物2 ID', 'width' => '80px'),
		array('name' => 'monster2_chance', 'text' => '2出现概率%', 'width' => '50px'),
		array('name' => 'monster3_id', 'text' => '怪物3 ID', 'width' => '80px'),
		array('name' => 'monster3_chance', 'text' => '3出现概率%', 'width' => '50px'),
		array('name' => 'monster4_id', 'text' => '怪物4 ID', 'width' => '80px'),
		array('name' => 'monster4_chance', 'text' => '4出现概率%', 'width' => '50px'),
		array('name' => 'monster5_id', 'text' => '怪物5 ID', 'width' => '80px'),
		array('name' => 'monster5_chance', 'text' => '5出现概率%', 'width' => '50px'),
		array('name' => 'is_boss', 'text' => 'Boss', 'width' => '80px'),
		array('name' => 'boss_dir', 'text' => 'boss朝向', 'width' => '80px'),
	),
	'where'	  => 'sql_where',
	'insert'  => 'sql_insert',
	'opera' => array(
		array('type' => 'link', 'text' => '副本敌人', 'render' => 'mission_enemy_opera'),
	),
	'del_check' => 'del_check',
);
?>
