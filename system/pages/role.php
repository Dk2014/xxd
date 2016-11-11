<?php
include "role_common.php";

$pconfig = array(
	'title'   => '玩家角色数据',
	'table'   => 'role',
	'links'   => $role_links,
	
	'columns' => array(
		array('name' => 'name', 'text' => '角色名称', 'width' => '80px'),
		array('name' => 'type', 'text' => '类型', 'width' => '80px'),
		array('name' => 'scale', 'text' => '缩放比', 'width' => '80px'),
		array('name' => 'is_special', 'text' => '是否特殊伙伴', 'width' => '90px'),
		array('name' => 'buddy_level', 'text' => '伙伴等级', 'width' => '80px'),
		array('name' => 'mission_lock', 'text' => '解锁副本权值', 'width' => '80px'),
		array('name' => 'skill_id1', 'text' => '初始招式1', 'width' => '80px'),
		array('name' => 'skill_id2', 'text' => '初始招式2', 'width' => '80px'),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '80px'),

	),
	
	'opera' => array(
		array('type' => 'link', 'text' => '', 'render' => 'buddy_data'),
	),
	
	// 'js' 			=> '',
	'where' 		=> 'sql_where',
	// 'insert'  		=> 'sql_insert',
	// 'location' 		=> '',
	// 'not_delete' 	=> true,
	// 'not_new' 		=> true,
	// 'not_edit' 		=> true,
);
?>
