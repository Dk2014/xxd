<?php

$extend_columns = array(
/*   '字段' => 配置 */

	'item_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	'item2_id' => 'item_id',
);


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

function js_function($params) {
	global $all_items;

	$html = choose_single_item($all_items, 'item_id');
	$html.=choose_single_item($all_items, 'item2_id');
	return $html;
}

?>