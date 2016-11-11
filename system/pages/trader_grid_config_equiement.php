<?php


$pconfig = array(
	'title'   => '',
	'table'   => 'trader_grid_config_equiement',
	'links'   => array(),

	'columns' => array(
		array('name' => 'min_level', 'text' => '等级下限', 'width' => '80px'),
		array('name' => 'max_level', 'text' => '等级上限', 'width' => '80px'),
		array('name' => 'cost', 'text' => '价格', 'width' => '80px'),
		array('name' => 'item_id', 'text' => '物品ID', 'width' => '80px'),

	),

	'where' 		=> 'sql_where',
	'insert'  		=> 'sql_insert',
	'js' => 'jsFunction',

);
?>
