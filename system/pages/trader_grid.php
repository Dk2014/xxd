<?php

$pconfig = array(
	'title'   => '货物配置',
	'table'   => 'trader_grid',
	'links'   => array(),

	'columns' => array(
		array('name' => 'money_type', 'text' => '货币类型', 'width' => '80px'),
	),

	'opera' => array(
		array('type' => 'link', 'text' => '', 'render' => 'trader_grid_config'),
		array('type' => 'text', 'text' => '', 'render' => 'trader_grid_config_list'),
	),


	'where' 		=> 'sql_where',
	'insert'  		=> 'sql_insert',
	'location'		=> 'location',
);
?>
