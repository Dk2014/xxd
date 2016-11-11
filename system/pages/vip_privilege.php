<?php
include "vip_common.php";

function sql_where($params) {
	return "order by `order`";
}


$pconfig = array(
	'title'   => 'VIP特权',
	'table'   => 'vip_privilege',
	'links'   => $links,

	'columns' => array(
		array('name' => 'order', 'text' => '顺序', 'width' => '180px'),
		array('name' => 'name', 'text' => '特权名称', 'width' => '180px'),
		array('name' => 'sign', 'text' => '唯一标识', 'width' => '180px'),
		array('name' => 'tip', 'text' => '特权描述', 'width' => '80px'),
	),

	'where' => 'sql_where',

);
