<?php
include "role_common.php";


$pconfig = array(
	'title'   => '玩家角色等级经验',
	'table'   => 'role_level_exp',
	'links'   => $role_links,
	
	'columns' => array(
		array('name' => 'level', 'text' => '等级', 'width' => '150'),
		array('name' => 'exp', 'text' => '经验', 'width' => '80'),
	),
	'pagesize' => 99999,
);
?>
