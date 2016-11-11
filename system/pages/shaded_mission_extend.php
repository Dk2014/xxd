<?php

function map_enemy_render($row) {
	$html = html_get_link("悄悄等着吓你一跳的小怪兽们", "?p=mission_enemy&m={$row['mission_level_id']}&m2={$row['id']}&bt=0", false);
	return $html;
}

function sql_insert($params) {
	return "`mission_level_id` = {$params['m']}";
}

function sql_where($params) {
	return " where `mission_level_id` = {$params['m']} order by `order` asc";
}

$extend_columns = array(

	'flip_horizontal' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('否', '是'),
	),

	'box_dir' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('朝左', '朝右'),
	),

	'award_item1' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	
	'award_item2' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	
	'award_item3' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	
);

function render_flip_horizontal($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_flip_horizontal($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function editor_box_dir($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_box_dir($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

$all_items = get_all_item();

function range_award_item1(){
	global $all_items;
	$all_items_temp = $all_items;
	$all_items_temp[0] = '无';
	return $all_items_temp;
}

function editor_award_item1($column_name, $column_val, $row, $data){
	global $all_items;

	$item_name = '';
	if (isset($all_items[$column_val])) {
		$item_name = $all_items[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item1($column_name, $column_val, $row, $data){
	global $all_items;

	return (isset($all_items[$column_val]) ? $all_items[$column_val] : "无");
}

function range_award_item2(){
	global $all_items;
	$all_items_temp = $all_items;
	$all_items_temp[0] = '无';
	return $all_items_temp;
}

function editor_award_item2($column_name, $column_val, $row, $data){
	global $all_items;

	$item_name = '';
	if (isset($all_items[$column_val])) {
		$item_name = $all_items[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item2($column_name, $column_val, $row, $data){
	global $all_items;

	return (isset($all_items[$column_val]) ? $all_items[$column_val] : "无");
}

function range_award_item3(){
	global $all_items;
	$all_items_temp = $all_items;
	$all_items_temp[0] = '无';
	return $all_items_temp;
}

function editor_award_item3($column_name, $column_val, $row, $data){
	global $all_items;

	$item_name = '';
	if (isset($all_items[$column_val])) {
		$item_name = $all_items[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item3($column_name, $column_val, $row, $data){
	global $all_items;

	return (isset($all_items[$column_val]) ? $all_items[$column_val] : "无");
}

function js_function($params) {
	global $all_items;

	return choose_single_item($all_items, 'award_item1').choose_single_item($all_items, 'award_item2').choose_single_item($all_items, 'award_item3');
}

