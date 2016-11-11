<?php

include "item_links.php";

$pconfig = array(
	'title'   => '灵宠契约球',
	'table'   => 'item',
	'links'   => $links,
	
	'columns' => array(
		array('name' => 'name', 'text' => '契约球名称', 'width' => '150px'),
		array('name' => 'type_id', 'text' => '类型ID', 'width' => '70px'),
		array('name' => 'quality', 'text' => '品质', 'width' => '50px'),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '100px'),
		array('name' => 'price', 'text' => '价格', 'width' => '50px'),
		array('name' => 'desc', 'text' => '材料描述', 'width' => '80px'),
		array('name' => 'panel', 'text' => '指向面板', 'width' => '80px'),
	),
	
	'where' 		=> 'sql_where',
	'not_delete' => true,
);
?>
