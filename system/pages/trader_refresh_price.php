<?php

$pconfig = array(
	'title'   => '随机商人刷新价格',
	'table'   => 'trader_refresh_price',
	'links'   => array(),

	'columns' => array(
		array('name' => 'time', 'text' => '次数', 'width' => '80px'),
		array('name' => 'price', 'text' => '价格', 'width' => '80px'),
	),


	'where' 		=> 'sql_where',
	'insert'  		=> 'sql_insert',
);
?>
