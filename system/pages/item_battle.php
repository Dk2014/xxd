<?php

include "item_links.php";

$pconfig = array(
	'title'   => '战斗道具',
	'table'   => 'item',
	'links'   => $links,
	
	'columns' => array(
		array('name' => 'name', 'text' => '道具名称', 'width' => '100px'),
		array('name' => 'quality', 'text' => '品质', 'width' => '60px'),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '150px'),
		array('name' => 'music_sign', 'text' => '音乐资源', 'width' => '150px'),
		array('name' => 'price', 'text' => '价格', 'width' => '60px'),
		array('name' => 'desc', 'text' => '道具描述', 'width' => '500px'),
		array('name' => 'show_mode', 'text' => '展现方式','width' => '100px'),
	),
	
	'foot'          => 'foot',
	'where' 		=> 'sql_where',
	'insert'		=> 'sql_insert',
	'not_delete' => true,
	'opera' => array(
		array('type' => 'link', 'text' => '道具配置', 'render' => 'op_render'),
	),
);
?>
