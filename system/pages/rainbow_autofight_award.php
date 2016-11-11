<?php



$pconfig = array(
	'title'   => '彩虹关扫荡宝箱',
	'table'   => 'rainbow_level_award',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'order', 'text' => '品质顺序', 'width' => '80px'),
		array('name' => 'award_type', 'text' => '奖励类型', 'width' => '80px'),
		array('name' => 'award_chance', 'text' => '奖励概率%', 'width' => '80px'),
		array('name' => 'award_num', 'text' => '奖励数量', 'width' => '80px'),
		array('name' => 'item_id', 'text' => '物品ID(非物品填 无 )', 'width' => '150px'),
	),
	
	'opera' => array(
	),
	
	'js'=>'js_function',
	'insert' => 'sql_insert',
	'where' => 'sql_where',
);

?>
