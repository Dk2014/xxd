<?php

$pconfig = array(
	'title'   => '问答',
	'table'   => 'faq',
	'links'   => array(),
	
	'columns' => array(
		// array('name' => 'id', 'text' => '主键', 'width' => '80px'),
		array('name' => 'order', 'text' => '顺序', 'width' => '80px'),
		array('name' => 'question', 'text' => '问题', 'width' => '80px'),
		array('name' => 'answer', 'text' => '回答', 'width' => '80px'),

	),
	
	'where' 		=> 'sql_where',

);


?>