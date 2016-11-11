<?php


$pconfig = array(
	'title'   => '关卡小宝箱物品',
	'table'   => 'mission_level_small_box_items',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'probability', 'text' => '出现概率%', 'width' => '80px'),
		array('name' => 'award_type', 'text' => '奖励类型', 'width' => '80px'),
		array('name' => 'item_number', 'text' => '奖励数量', 'width' => '80px'),
		array('name' => 'item_id', 'text' => '物品ID(仅物品奖励填写)', 'width' => '150px'),
	),
	
	'opera' => array(
	),
	
	'js'=>'js_function',
	'insert' => 'sql_insert',
	'where' => 'sql_where',
	'location' => 'location',
);


?>