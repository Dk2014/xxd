<?php
include "skill_links.php";

$pconfig = array(
	'title'   => '敌人招式',
	'table'   => 'skill',
	'links'   => $links,
	'pagesize' => 100000,

	'columns' => array(
		array('name' => 'order', 'text' => '排序字段','width' => '80px'),
		array('name' => 'child_type', 'text' => '绝招作用', 'width' => '60px'),
		array('name' => 'name', 'text' => '绝招名称','width' => '200px'),
		array('name' => 'sign', 'text' => '资源标识','width' => '200px'),
		array('name' => 'music_sign', 'text' => '音乐资源标识', 'width' => '100px'),
		array('name' => 'jump_attack', 'text' => '是否跳跃攻击','width' => '80px'),
		array('name' => 'target', 'text' => '特效展现方式', 'width' => '60px'),
	),

	'opera' => array(
		array('type' => 'link', 'text' => '绝招配置', 'render' => 'skill_opera'),
	    array('type' => 'link', 'text' => '绝招描述', 'render' => 'skill_info'),
	),

	'where'   => 'sql_where',
	'insert'  => 'sql_insert',
	'foot' => 'foot',
);
?>
