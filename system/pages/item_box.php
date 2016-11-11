<?php
$pconfig = array(
	'title'   => '宝箱数据',
	'table'   => 'item',
	'links'   => $links,
	'columns' => array(
		array('name' => 'name', 'text' => '宝箱名称', 'width' => '150px'),
		array('name' => 'type_id', 'text' => '类型ID','width' => '70px',),
		array('name' => 'quality', 'text' => '品质','width' => '50px',),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '150px'),
		array('name' => 'price', 'text' => '价格', 'width' => '70px'),
		array('name' => 'level', 'text' => '需求等级', 'width' => '70px'),
		array('name' => 'func_id', 'text' => '使用的功能限制','width' => '100px'),		
		array('name' => 'use_ingot', 'text' => '开启的元宝价格', 'width' => '70px'),
		array('name' => 'valid_hours', 'text' => '有效小时数', 'width' => '70px'),
		array('name' => 'renew_ingot', 'text' => '续费的元宝价格', 'width' => '70px'),
		array('name' => 'desc', 'text' => '材料描述', 'width' => '200px'),
		array('name' => 'can_batch', 'text' => '可否批量使用', 'width' => '100px'),
	),
	'opera' => array(
		array('type' => 'link', 'text' => '宝箱内容', 'render' => 'item_box_content_render'),
	),
	'where' => 'sql_where',
	'not_delete' => true,
	'js' => 'js_function',
);
?>