<?php
include "common_links.php";

$pconfig = array(
	'title'   => '元宝培养魂侍价格配置',
	'table'   => 'ghost_train_price',
	'links'   => $ghost_links,
	
	'columns' => array(
		array('name' => 'times', 'text' => '今日次数', 'width' => '100px'),
		array('name' => 'cost', 'text' => '魂力果实单价', 'width' => '100px'),

	),
	'where' 		=> 'sql_where',
);

function sql_where($params) {
	return ' order by `times` asc';
}

?>
