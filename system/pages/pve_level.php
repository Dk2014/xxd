<?php

$pconfig = array(
	'title'   => '灵宠幻境',
	'table'   => 'pve_level',
	'links'   => array(),

	'columns' => array(
		array('name' => 'floor', 'text' => '层数', 'width' => '80px'),
		array('name' => 'level', 'text' => '要求等级', 'width' => '80px'),
		array('name' => 'moster_num', 'text' => '怪物数量', 'width' => '80px'),
		array('name' => 'award_item', 'text' => '首通奖励物品', 'width' => '80px'),
		array('name' => 'award_num', 'text' => '首通奖励数量', 'width' => '80px'),
		array('name' => 'basic_award_num', 'text' => '基础奖励数量', 'width' => '80px'),
		array('name' => 'award_factor', 'text' => '奖励系数', 'width' => '80px'),
	),


	'js'			=> 'js_function',
	'opera' => array(
		array('type' => 'link', 'text' => '怪物配置', 'render' => 'leveltable'),
	),

);
?>
