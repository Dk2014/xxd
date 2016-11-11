<?php
include "role_common.php";

$pconfig = array(
	'title'   => 'npc角色',
	'table'   => 'npc_role',
	'links'   => $role_links,
	
	'columns' => array(
		array('name' => 'name', 'text' => '角色名称', 'width' => '80px'),
	),
	
	// 'js' 			=> '',
	//'where' 		=> 'sql_where',
	// 'insert'  		=> 'sql_insert',
	// 'location' 		=> '',
	// 'not_delete' 	=> true,
	// 'not_new' 		=> true,
	// 'not_edit' 		=> true,
);
?>
