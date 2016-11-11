<?php
include_once "common_links.php";

$pconfig = array(
	'title'   => '剑心类型',
	'table'   => 'sword_soul_type',
	'links'   => $sword_soul_links,
	'columns' => array(
		array('name' => 'name', 'text' => '名称'),
		array('name' => 'sign', 'text' => '程序标示'),
	),
	'not_delete' => true,
	'not_new' => true,
);
?>