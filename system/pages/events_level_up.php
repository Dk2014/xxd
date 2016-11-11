<?php

$pconfig = array(
	'title'   => '运营等级活动',
	'table'   => 'events_level_up',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'require_level', 'text' => '需要等级', 'width' => '80px'),
		array('name' => 'ingot', 'text' => '奖励元宝', 'width' => '80px'),
		array('name' => 'coins', 'text' => '奖励铜钱', 'width' => '80px'),

		array('name' => 'item1_id', 'text' => '奖励物品1', 'width' => '80px'),
		array('name' => 'item1_num', 'text' => '物品1数量', 'width' => '80px'),

		array('name' => 'item2_id', 'text' => '奖励物品2', 'width' => '80px'),
		array('name' => 'item2_num', 'text' => '物品2数量', 'width' => '80px'),

		array('name' => 'item3_id', 'text' => '奖励物品3', 'width' => '80px'),
		array('name' => 'item3_num', 'text' => '物品3数量', 'width' => '80px'),
	),	

	'where' 		=> 'sql_where',
	'js' 			=> 'js_function',
);


function sql_where($params) {
	return " order by require_level asc";
}

?>