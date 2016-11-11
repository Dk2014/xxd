<?php

$pconfig = array(
	'title'   => '每日任务',
	'table'   => 'daily_quest',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'order', 'text' => '显示顺序', 'width' => '80px'),
		array('name' => 'name', 'text' => '任务标题', 'width' => '80px'),
		array('name' => 'require_open_day', 'text' => '开放日', 'width' => '180px'),
		array('name' => 'require_min_level', 'text' => '要求最低等级', 'width' => '50px'),
		array('name' => 'require_max_level', 'text' => '要求最高等级', 'width' => '50px'),
		

		array('name' => 'class', 'text' => '类别', 'width' => '50px'),
		array('name' => 'level_type', 'text' => '关卡', 'width' => '50px'),
		array('name' => 'level_sub_type', 'text' => '关卡子类', 'width' => '50px'),

		array('name' => 'require_count', 'text' => '要求完成数量', 'width' => '80px'),
		array('name' => 'award_exp', 'text' => '奖励经验', 'width' => '80px'),
		array('name' => 'award_ingot', 'text' => '奖励元宝', 'width' => '80px'),
		array('name' => 'award_coins', 'text' => '奖励铜钱', 'width' => '80px'),
		array('name' => 'award_physical', 'text' => '奖励体力', 'width' => '80px'),
		array('name' => 'award_item1_id', 'text' => '奖励物品1', 'width' => '80px'),
		array('name' => 'award_item1_num', 'text' => '奖励物品1数量', 'width' => '80px'),
		array('name' => 'award_item2_id', 'text' => '奖励物品2', 'width' => '80px'),
		array('name' => 'award_item2_num', 'text' => '奖励物品2数量', 'width' => '80px'),
		array('name' => 'award_item3_id', 'text' => '奖励物品3', 'width' => '80px'),
		array('name' => 'award_item3_num', 'text' => '奖励物品3数量', 'width' => '80px'),
		array('name' => 'award_item4_id', 'text' => '奖励物品4', 'width' => '80px'),
		array('name' => 'award_item4_num', 'text' => '奖励物品4数量', 'width' => '80px'),
		array('name' => 'award_stars', 'text' => '奖励星数', 'width' => '80px'),
		array('name' => 'desc', 'text' => '简介', 'width' => '180px'),
	),


	'where' 		=> 'sql_where',
	'js'			=> 'jsFunction',

);
?>