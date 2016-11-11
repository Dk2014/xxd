<?php

include "item_links.php";

$pconfig = array(
	'title'   => '材料数据',
	'table'   => 'item',
	'links'   => $links,
	
	'columns' => array(
		array('name' => 'name', 'text' => '材料名称', 'width' => '150px'),
		array('name' => 'type_id', 'text' => '类型ID', 'width' => '70px'),
		array('name' => 'quality', 'text' => '品质', 'width' => '50px'),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '100px'),
		array('name' => 'price', 'text' => '价格', 'width' => '50px'),
		array('name' => 'level', 'text' => '需求等级', 'width' => '60px'),
		array('name' => 'desc', 'text' => '材料描述', 'width' => '80px'),
		array('name' => 'can_use', 'text' => '能否使用', 'width' => '80px'),
		array('name' => 'panel', 'text' => '指向面板', 'width' => '80px'),
	),
	
	'where' 		=> 'sql_where',
	'not_delete' => true,
);
?>
