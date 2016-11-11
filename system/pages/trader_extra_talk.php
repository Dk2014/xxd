<?php

$pconfig = array(
	'title'   => '随机商人附加对话',
	'table'   => 'trader_extra_talk',
	'links'   => array(),

	'columns' => array(
		array('name' => 'time', 'text' => '点击次数', 'width' => '80px'),
		array('name' => 'talk', 'text' => '对话', 'width' => '80px'),

	),

	'where' 		=> 'sql_where',
	'insert'  		=> 'sql_insert',

);
?>
