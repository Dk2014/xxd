<?php

include("extend_level_type.php");

$pconfig = array(
	'title'   => $extend_level_type[$_GET['m']],
	'table'   => 'extend_level',
	'links'   => array(	
		array('text' => '资源关卡', 'id' => 'extend_level&m=1'),
		array('text' => '伙伴关卡', 'id' => 'extend_level&m=9'),
		array('text' => '灵宠关卡', 'id' => 'extend_level&m=10'),
		array('text' => '魂侍关卡', 'id' => 'extend_level&m=11'),
),
	
	'columns' => array(
		array('name' => 'max_level', 'text' => '开放等级', 'width' => '180px'),
	),
	
	'opera' => array(
		array('type' => 'link', 'text' => '关卡', 'render' => 'leveltable'),
	),

	'not_delete' => true,
	'where' => 'sql_where',
	'insert' => 'sql_insert',
);

?>