<?php
include "role_common.php";

$pconfig = array(
	'title'   => '招募伙伴',
	'table'   => 'recruit_buddy',
	'links'   => $role_links,
	
	'columns' => array(
		array('name' => 'role_id', 'text' => '关联角色', 'width' => '80px'),
		array('name' => 'init_level', 'text' => '初始等级', 'width' => '80px'),
		array('name' => 'related_npc', 'text' => 'NPC角色', 'width' => '80px'),
		array('name' => 'quest_id', 'text' => '开启任务', 'width' => '80px'),
		array('name' => 'favourite_item', 'text' => '喜好品', 'width' => '80px'),
		array('name' => 'favourite_count', 'text' => '喜好品数量', 'width' => '80px'),
		array('name' => 'description', 'text' => '简介', 'width' => '180px'),
	),


	'where' 		=> 'sql_where',
	'js'			=> 'jsFunction',

);
?>
