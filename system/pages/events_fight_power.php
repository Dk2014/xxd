<?php


$pconfig = array(
	'title'   => '战力活动',
	'table'   => 'events_fight_power',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'lock', 'text' => '档位', 'width' => '80px'),
		array('name' => 'fight', 'text' => '战力', 'width' => '80px'),
		array('name' => 'ingot', 'text' => '奖励元宝', 'width' => '80px'),
		array('name' => 'item1_id', 'text' => '物品1', 'width' => '80px'),
		array('name' => 'item1_num', 'text' => '物品1数量', 'width' => '80px'),
		array('name' => 'item2_id', 'text' => '物品2', 'width' => '80px'),
		array('name' => 'item2_num', 'text' => '物品2数量', 'width' => '80px'),
		array('name' => 'item3_id', 'text' => '物品3', 'width' => '80px'),
		array('name' => 'item3_num', 'text' => '物品3数量', 'width' => '80px'),

	),
	

	'where' 		=> 'sql_where',
	'js' 			=> 'js_function',
);


function sql_where($params) {
	return " order by `lock` asc";
}
?>