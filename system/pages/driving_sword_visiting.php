<?php

$pconfig = array(
	'title'   => '云海御剑拜访类',
	'table'   => 'driving_sword_visiting',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'event_id', 'text' => '拜访山id', 'width' => '90px'),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '90px'),
		array('name' => 'award_item1', 'text' => '奖励品1', 'width' => '90px'),
		array('name' => 'award_num1', 'text' => '奖励品1数量', 'width' => '90px'),
		array('name' => 'award_item2', 'text' => '奖励品2', 'width' => '90px'),
		array('name' => 'award_num2', 'text' => '奖励品2数量', 'width' => '90px'),
		array('name' => 'award_item3', 'text' => '奖励品3', 'width' => '90px'),
		array('name' => 'award_num3', 'text' => '奖励品3数量', 'width' => '90px'),
		array('name' => 'award_coin_num', 'text' => '奖励铜钱数量', 'width' => '90px'),
	),
	
	'where' => 'sql_where',
	'insert' => 'sql_insert',
	'js' => 'js_function',
);

