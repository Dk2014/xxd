<?php

$pconfig = array(
	'title'   => '龙珠兑换配置',
	'table'   => 'item_dragon_ball_config',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'item_id', 'text' => '物品', 'width' => '120px'),
		array('name' => 'item_num', 'text' => '物品数量', 'width' => '120px'),
		array('name' => 'rate', 'text' => '概率', 'width' => '120px'),

	),
	
	'where' 		=> 'sql_where',
	'js'=>'js_function',
	'insert' => 'sql_insert',
);

?>
