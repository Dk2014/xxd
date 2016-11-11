<?php
define("MAX_VIP_LEVEL", 15); //VIP最大等级

$vip_arr = array(
	0	=>	0);
$i = 1;
for(; $i <= MAX_VIP_LEVEL; $i++){
	$vip_arr[$i] = $i;
}

$extend_columns = array(
	'require_vip_level' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
	),
	'award_vip_level1' => 'require_vip_level',
	'award_vip_level2' => 'require_vip_level',
	'award_vip_level3' => 'require_vip_level',
	'award_vip_level4' => 'require_vip_level',
	'award_vip_level5' => 'require_vip_level',

	'award_vip_item1_id' => array(
			'editor' => array('params' => array()),
			'render' => array('params' => array()),
			'data' => array(),
			'range' => array('params' => array()),
		),
	'award_vip_item2_id' => array(
			'editor' => array('params' => array()),
			'render' => array('params' => array()),
			'data' => array(),
			'range' => array('params' => array()),
		),
	'award_vip_item3_id' => array(
			'editor' => array('params' => array()),
			'render' => array('params' => array()),
			'data' => array(),
			'range' => array('params' => array()),
		),
	'award_vip_item4_id' => array(
			'editor' => array('params' => array()),
			'render' => array('params' => array()),
			'data' => array(),
			'range' => array('params' => array()),
		),
	'award_vip_item5_id' => array(
			'editor' => array('params' => array()),
			'render' => array('params' => array()),
			'data' => array(),
			'range' => array('params' => array()),
		),
);

$extend_columns['require_vip_level']['data'] = $vip_arr;

function editor_require_vip_level($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_require_vip_level($column_name, $column_val, $row, $data){
	global $vip_arr;
	return $vip_arr[$column_val];
}

$all_items = get_all_item();

function editor_award_vip_item1_id($column_name, $column_val, $row, $data){
	global $all_items;

	$item_name = '';
	if (isset($all_items[$column_val])) {
		$item_name = $all_items[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_vip_item1_id($column_name, $column_val, $row, $data){
	global $all_items;

	return (isset($all_items[$column_val]) ? $all_items[$column_val] : "无");
}

function range_award_vip_item1_id(){
	global $all_items;

	$all_items_temp = $all_items;
	$all_items_temp[0] = '无';
	return $all_items_temp;
}

function editor_award_vip_item2_id($column_name, $column_val, $row, $data){
	global $all_items;

	$item_name = '';
	if (isset($all_items[$column_val])) {
		$item_name = $all_items[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_vip_item2_id($column_name, $column_val, $row, $data){
	global $all_items;

	return (isset($all_items[$column_val]) ? $all_items[$column_val] : "无");
}

function range_award_vip_item2_id(){
	global $all_items;

	$all_items_temp = $all_items;
	$all_items_temp[0] = '无';
	return $all_items_temp;
}

function editor_award_vip_item3_id($column_name, $column_val, $row, $data){
	global $all_items;

	$item_name = '';
	if (isset($all_items[$column_val])) {
		$item_name = $all_items[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_vip_item3_id($column_name, $column_val, $row, $data){
	global $all_items;

	return (isset($all_items[$column_val]) ? $all_items[$column_val] : "无");
}

function range_award_vip_item3_id(){
	global $all_items;

	$all_items_temp = $all_items;
	$all_items_temp[0] = '无';
	return $all_items_temp;
}

function editor_award_vip_item4_id($column_name, $column_val, $row, $data){
	global $all_items;

	$item_name = '';
	if (isset($all_items[$column_val])) {
		$item_name = $all_items[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_vip_item4_id($column_name, $column_val, $row, $data){
	global $all_items;

	return (isset($all_items[$column_val]) ? $all_items[$column_val] : "无");
}

function range_award_vip_item4_id(){
	global $all_items;

	$all_items_temp = $all_items;
	$all_items_temp[0] = '无';
	return $all_items_temp;
}

function editor_award_vip_item5_id($column_name, $column_val, $row, $data){
	global $all_items;

	$item_name = '';
	if (isset($all_items[$column_val])) {
		$item_name = $all_items[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_vip_item5_id($column_name, $column_val, $row, $data){
	global $all_items;

	return (isset($all_items[$column_val]) ? $all_items[$column_val] : "无");
}

function range_award_vip_item5_id(){
	global $all_items;

	$all_items_temp = $all_items;
	$all_items_temp[0] = '无';
	return $all_items_temp;
}

function sql_where(){
	return "order by `require_vip_level` ASC, `require_vip_count` ASC";
}

function js_function($params) {
	global $all_items;

	$html  = choose_single_item($all_items, 'award_vip_item1_id');
	$html .= choose_single_item($all_items, 'award_vip_item2_id');
	$html .= choose_single_item($all_items, 'award_vip_item3_id');
	$html .= choose_single_item($all_items, 'award_vip_item4_id');
	$html .= choose_single_item($all_items, 'award_vip_item5_id');

	return $html;
}


?>