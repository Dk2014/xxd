<?php
include "item_equipment_links.php";

$extend_columns = array(
    /*'字段' => 配置 */
	'dragon_ball' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range' =>array('params' => array()),
	),
	'crystal_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range' =>array('params' => array()),
	),
	'quality' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
);

function range_quality(){
	global $all_quality;
	return $all_quality;
}

function render_quality($column_name, $column_val, $row, $data) {
	global $all_quality;
	return $all_quality[$column_val];
}

function editor_quality($column_name, $column_val, $row, $data) {
	global $all_quality;
	return html_get_select($column_name,$all_quality,$column_val);
}

function editor_fragment_id($column_name, $column_val, $row, $data){
	global $allItem;

	$item_name = '';
	if (isset($allItem[$column_val])) {
		$item_name = $allItem[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_fragment_id($column_name, $column_val, $row, $data){
	global $allItem;

	return (isset($allItem[$column_val]) ? $allItem[$column_val] : "0");
}

function editor_dragon_ball($column_name, $column_val, $row, $data){
	global $allItem;

	$item_name = '';
	if (isset($allItem[$column_val])) {
		$item_name = $allItem[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_dragon_ball($column_name, $column_val, $row, $data){
	global $allItem;

	return (isset($allItem[$column_val]) ? $allItem[$column_val] : "无");
}

function editor_crystal_id($column_name, $column_val, $row, $data){
	global $allItem;

	$item_name = '';
	if (isset($allItem[$column_val])) {
		$item_name = $allItem[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_crystal_id($column_name, $column_val, $row, $data){
	global $allItem;

	return (isset($allItem[$column_val]) ? $allItem[$column_val] : "无");
}

$allItem = get_item_exclude_equip();

function jsFunction($params) {
	global $allItem;

	$html = '';

	$html .= choose_single_item($allItem, 'crystal_id');
	$html .= choose_single_item($allItem, 'dragon_ball');
	return $html;
}
?>
