<?php

$pconfig = array(
	'title'   => '扩展任务',
	'table'   => 'extend_quest',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'type', 'text' => '任务类型', 'width' => '80px'),
		array('name' => 'required_quest', 'text' => '前置主线任务', 'width' => '80px'),
		array('name' => 'related_npc', 'text' => '关联NPC', 'width' => '80px'),
		array('name' => 'related_mission', 'text' => '关联区域', 'width' => '80px'),
		array('name' => 'required_progress', 'text' => '要求进度', 'width' => '80px'),
		array('name' => 'award_exp', 'text' => '奖励经验', 'width' => '80px'),
		array('name' => 'award_coins', 'text' => '奖励铜钱', 'width' => '80px'),
		array('name' => 'award_item_1', 'text' => '奖励道具1', 'width' => '80px'),
		array('name' => 'award_num_1', 'text' => '奖励道具1数量', 'width' => '80px'),
		array('name' => 'award_item_2', 'text' => '奖励道具2', 'width' => '80px'),
		array('name' => 'award_num_2', 'text' => '奖励道具2数量', 'width' => '80px'),
		array('name' => 'award_item_3', 'text' => '奖励道具3', 'width' => '80px'),
		array('name' => 'award_num_3', 'text' => '奖励道具3数量', 'width' => '80px'),
		array('name' => 'description', 'text' => '简介', 'width' => '180px'),
	),


	'where' 		=> 'sql_where',
	'js'			=> 'jsFunction',

);
?>
