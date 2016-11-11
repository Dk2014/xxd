<?php

$extend_columns = array(
/*   '字段' => 配置 */

	'award_item1_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'award_item2_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'award_item3_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
);

$all_items = get_all_item();

function editor_award_item1_id($column_name, $column_val, $row, $data){
	global $all_items;

	$item_name = '';
	if (isset($all_items[$column_val])) {
		$item_name = $all_items[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item1_id($column_name, $column_val, $row, $data){
	global $all_items;

	return (isset($all_items[$column_val]) ? $all_items[$column_val] : "无");
}

function range_award_item1_id(){
	global $all_items;
	$temp_all_items = $all_items;
	$temp_all_items[0] = '无';
	return $temp_all_items;
}
function editor_award_item2_id($column_name, $column_val, $row, $data){
	global $all_items;

	$item_name = '';
	if (isset($all_items[$column_val])) {
		$item_name = $all_items[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item2_id($column_name, $column_val, $row, $data){
	global $all_items;

	return (isset($all_items[$column_val]) ? $all_items[$column_val] : "无");
}

function range_award_item2_id(){
	global $all_items;
	$temp_all_items = $all_items;
	$temp_all_items[0] = '无';
	return $temp_all_items;
}
function editor_award_item3_id($column_name, $column_val, $row, $data){
	global $all_items;

	$item_name = '';
	if (isset($all_items[$column_val])) {
		$item_name = $all_items[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item3_id($column_name, $column_val, $row, $data){
	global $all_items;

	return (isset($all_items[$column_val]) ? $all_items[$column_val] : "无");
}

function range_award_item3_id(){
	global $all_items;
	$temp_all_items = $all_items;
	$temp_all_items[0] = '无';
	return $temp_all_items;
}

function opera_render($row) {
	$html = '<a href="?p=enemy_deploy_form&m='.$row['id'].'&m2=3" target="_blank">怪物布阵</a>';
	return $html;
}

function js_function($params) {
	global $all_items;

	$html = choose_single_item($all_items, 'award_item1_id');
	$html .= choose_single_item($all_items, 'award_item2_id');
	$html .= choose_single_item($all_items, 'award_item3_id');
	return $html;
}

function sql_where($row) {
	return " order by `require_level`";
}

?>
