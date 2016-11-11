<?php

$extend_columns = array(
/*   '字段' => 配置 */

	'award_type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('1'=>'铜钱','2'=>'元宝','3'=>'道具'),
	),
	'item_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);


function render_award_type($column_name, $column_val, $row, $data) {
	return $data[$column_val];
}

function editor_award_type($column_name, $column_val, $row, $data) {
	return html_get_select($column_name,$data,$column_val);
}


$all_items = get_all_item();

function editor_item_id($column_name, $column_val, $row, $data){
	global $all_items;

	$item_name = '';
	if (isset($all_items[$column_val])) {
		$item_name = $all_items[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_item_id($column_name, $column_val, $row, $data){
	global $all_items;

	return (isset($all_items[$column_val]) ? $all_items[$column_val] : "无");
}


function sql_insert($params) {
	return "`heart_draw_id` = {$params['m']}";
}


function sql_where($params) {
	if (!isset($params['m'])){
		return '';
	}

	return "where `heart_draw_id`={$params['m']} ";
}

function js_function($params) {
	global $all_items;

	return choose_single_item($all_items, 'item_id');
}

?>