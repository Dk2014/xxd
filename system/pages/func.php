<?php
include "func_links.php";

$pconfig = array(
	'title'   => '功能开放',
	'table'   => 'func',
	'links'   => $func_links,
	
	'columns' => array(
		array('name' => 'name', 'text' => '功能名称', 'width' => '80px', 'readonly' => true),
		array('name' => 'sign', 'text' => '功能标识', 'width' => '80px', 'readonly' => true),
		array('name' => 'type', 'text' => '开放类型', 'width' => '80px'),
		array('name' => 'lock', 'text' => '功能权值', 'width' => '80px'),
		array('name' => 'level', 'text' => '开启等级', 'width' => '80px'),
		array('name' => 'need_play', 'text' => '是否需要播放', 'width' => '80px'),
		array('name' => 'unique_key', 'text' => '唯一值', 'width' => '80px', 'readonly' => true),

	),
	
	'opera' => array(
	),
	
	'not_delete' => true,
	'not_new' => true,
	//'not_edit' => true,
	'where' 		=> 'sql_where',
);

?>
