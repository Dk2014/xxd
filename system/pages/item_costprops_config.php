<?php

$pconfig = array(
	'title'   => '消耗道具配置',
	'table'   => 'item_costprops',
	'links'   => $links,
	
	'columns' => array(
		array('name' => 'type', 'text' => '类型', 'width' => '80px'),
		array('name' => 'value', 'text' => '值', 'width' => '120px'),

	),
	
	'where' 		=> 'sql_where',
	'insert' => 'sql_insert',
);

?>