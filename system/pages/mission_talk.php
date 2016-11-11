<?php

$pconfig = array(
	'title'   => '战场对话',
	'table'   => 'mission_talk',
	'links'   => array(),
	'columns' => array(
		array('name' => 'town_id', 'text' => '城镇', 'width' => '80px'),
		array('name' => 'quest_id', 'text' => '任务', 'width' => '80px'),
		array('name' => 'content',  'text' => '对话内容', 'width' => '60px'),

	),
	'where'	  => 'sql_where',
	'insert'  => 'sql_insert',
	'js'	        => 'js_function',
);
?>
