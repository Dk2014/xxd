<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'description' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
	),

	'related_npc' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'		=> array('params' => array()),
	),
	
	'role_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'		=> array('params' => array()),
	),

	'quest_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'		=> array('params' => array()),
	),
	
	'description' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
	),
	

	'favourite_item' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'		=> array('params' => array()),
	),

);


function editor_quest_id($column_name, $column_val, $row, $data){
	global $allQuests;
	$quest_name = '';
	if (isset($allQuests[$column_val])) {
		$quest_name = $allQuests[$column_val];
	}

	return editor_single_item($quest_name, $column_val, $column_name);
}

function render_quest_id($column_name, $column_val, $row, $data){
	global $allQuests;

	return (isset($allQuests[$column_val]) ? $allQuests[$column_val] : "0");
}


function editor_related_npc($column_name, $column_val, $row, $data){
	global $allNPCs;
	$role_name = '';
	if (isset($allNPCs[$column_val])) {
		$role_name = $allNPCs[$column_val];
	}

	return editor_single_item($role_name, $column_val, $column_name);
}

function render_related_npc($column_name, $column_val, $row, $data){
	global $allNPCs;

	return (isset($allNPCs[$column_val]) ? $allNPCs[$column_val] : "0");
}


function editor_role_id($column_name, $column_val, $row, $data){
	global $allRoles;
	$role_name = '';
	if (isset($allRoles[$column_val])) {
		$role_name = $allRoles[$column_val];
	}

	return editor_single_item($role_name, $column_val, $column_name);
}

function render_role_id($column_name, $column_val, $row, $data){
	global $allRoles;

	return (isset($allRoles[$column_val]) ? $allRoles[$column_val] : "0");
}

function editor_favourite_item($column_name, $column_val, $row, $data){
	global $allItem;

	$item_name = '';
	if (isset($allItem[$column_val])) {
		$item_name = $allItem[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_favourite_item($column_name, $column_val, $row, $data){
	global $allItem;

	return (isset($allItem[$column_val]) ? $allItem[$column_val] : "0");
}


function editor_description($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_description($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function sql_where($params) {
	return "";
}

$allRoles = get_all_role();
$allItem = get_item_exclude_equip();
$allQuests = get_all_quest();
$allNPCs = get_all_npc_role();

function jsFunction($params) {
	global $allItem;
	global $allQuests;
	global $allRoles;
	global $allNPCs;

	$html = '
$("select").change( function(){
		$(this).css("color","#0033FF");
});';

	$html .= choose_single_item($allQuests, 'quest_id');
	$html .= choose_single_item($allRoles, 'role_id');
	$html .= choose_single_item($allNPCs, 'related_npc');
	$html .= choose_single_item($allItem, 'favourite_item');
	return $html;
}
?>
