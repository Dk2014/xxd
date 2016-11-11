<?php

$pconfig = array(
	'title'   => '功能预告',
	'table'   => 'func_prediction',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'order', 'text' => '顺序', 'width' => '80px'),
		array('name' => 'type', 'text' => '类型', 'width' => '80px'),
		array('name' => 'condition_value', 'text' => '触发条件', 'width' => '80px'),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '80px'),
		array('name' => 'summary', 'text' => '开启条件描述', 'width' => '80px'),
		array('name' => 'tips', 'text' => 'tips', 'width' => '80px'),

	),
	
	'opera' => array(
	),
	'where' 		=> 'sql_where',
	
);

?>
