<?php

$pconfig = array(
	'title'   => '下载奖励',
	'table'   => 'town_treasure',
	'links'   => array(),
	
	'columns' => array(
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
	),
	
	'opera' => array(
	),
	
	'where'	  => 'sql_where',
	'insert'  => 'sql_insert',
	'location' => 'location',
	'js' =>	'jsFunction',
);
