<?php

include "item_links.php";

$pconfig = array(
	'title'   => '时装',
	'table'   => 'item',
	'links'   => $links,
	
	'columns' => array(
		array('name' => 'name', 'text' => '名称', 'width' => '150px'),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '100px'),
		array('name' => 'quality', 'text' => '品质','width' => '50px',),
		array('name' => 'valid_hours', 'text' => '有效时间（小时，0永久有效）', 'width' => '50px'),
		array('name' => 'price', 'text' => '价格', 'width' => '50px'),
		array('name' => 'level', 'text' => '需求等级', 'width' => '60px'),
		array('name' => 'desc', 'text' => '描述', 'width' => '80px'),
	),

	'where' 		=> 'sql_where',
	'insert'		=> 'sql_insert',
	'not_delete' => true,
);
?>
