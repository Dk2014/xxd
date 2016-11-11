<?php

include "item_links.php";

$pconfig = array(
	'title'   => '秘宝',
	'table'   => 'item',
	'links'   => $links,
	
	'columns' => array(
		array('name' => 'name', 'text' => '道具名称', 'width' => '100px'),
		array('name' => 'quality', 'text' => '品质', 'width' => '60px'),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '150px'),
		array('name' => 'act_id', 'text' => '触发功能', 'width' => '150px'),
		array('name' => 'price', 'text' => '价格', 'width' => '60px'),
		array('name' => 'desc', 'text' => '描述', 'width' => '500px'),
	),
	'opera' => array(
		array('type' => 'link', 'text' => '功能参数', 'render' => 'item_reflect_config_render'),
	),
	
	'foot'          => 'foot',
	'where' 		=> 'sql_where',
	'insert'		=> 'sql_insert',
	'not_delete' => true,
);
?>
