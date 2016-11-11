<?php

include "item_links.php";

$pconfig = array(
	'title'   => '资源数据',
	'table'   => 'item',
	'links'   => $links,
	
	'columns' => array(
		array('name' => 'name', 'text' => '材料名称', 'width' => '150px'),
		array('name' => 'type_id', 'text' => '类型ID', 'width' => '70px'),
		array('name' => 'quality', 'text' => '品质', 'width' => '50px'),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '100px'),
		array('name' => 'level', 'text' => '需求等级', 'width' => '60px'),
		array('name' => 'desc', 'text' => '材料描述', 'width' => '80px'),
		array('name' => 'show_origin', 'text' => '显示产出', 'width' => '80px'),
	),
	
	'where' 		=> 'sql_where',
	'not_delete' => true,
);
?>
