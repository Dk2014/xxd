<?php

function sql_where($params) {	
	return " where `treasure_id`={$params['m']} order by `order` asc ";
}

function sql_insert($params) {
	return " `treasure_id`={$params['m']}";
}

$extend_columns = array(
	'award_item' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range'=>array('params' => array()),
	),
);

$all_items = get_all_item();

function editor_award_item($column_name, $column_val, $row, $data){
	global $all_items;

	$item_name = '';
	if (isset($all_items[$column_val])) {
		$item_name = $all_items[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function range_award_item(){
	global $all_items;
	$all_items_temp = $all_items;
	$all_items_temp[0] = '无';
	return $all_items_temp;
}

function render_award_item($column_name, $column_val, $row, $data){
	global $all_items;

	return (isset($all_items[$column_val]) ? $all_items[$column_val] : "无");
}

function js_function($params) {
	global $all_items;

	return choose_single_item($all_items, 'award_item');
}

