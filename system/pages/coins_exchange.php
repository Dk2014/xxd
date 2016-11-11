<?php

function sql_where($params) {
	return "order by `unique_key` asc";
}

$pconfig = array(
	'title'   => '铜币购买',
	'table'   => 'coins_exchange',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'unique_key', 'text' => '购买次数', 'width' => '150'),
		array('name' => 'ingot', 'text' => '消耗元宝', 'width' => '80'),
		array('name' => 'coins', 'text' => '获得铜币', 'width' => '80'),
	),
);
?>
