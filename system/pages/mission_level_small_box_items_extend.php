<?php


$extend_columns = array(
/*   '字段' => 配置 */
	'award_type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(/* 0 */'铜钱', /* 1 */'道具'),
	),
	'probability' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	'item_number' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	//'item_number' => 'probability',
	'item_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range'=>array('params'=>array()),
	),
);

function range_item_id(){
	global $all_items;
	$all_items_temp = $all_items;
	$all_items_temp[0] = '无';
	return $all_items_temp;
}

function get_items() {
	global $db;
	$items = array();

	$tmp = db_query($db, "select `id`, `name` from `item` where type_id != 11");
	foreach ($tmp as $value)
	{
		$items[$value['id']] = $value['name'];
	}

	return $items;
}


$all_items = get_items();

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

	return (isset($all_items[$column_val]) ? $all_items[$column_val] : "0");
}

function render_award_type($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_award_type($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_item_number($column_name, $column_val, $row, $data){
	return $column_val;
}

function editor_item_number($column_name, $column_val, $row, $data){
	if (!isset($column_val))
		$column_val = 0;
	return html_get_input($column_name, $column_val);
}

function render_probability($column_name, $column_val, $row, $data){
	return $column_val;
}

function editor_probability($column_name, $column_val, $row, $data){
	if (!isset($column_val))
		$column_val = 0;
	return html_get_input($column_name, $column_val);
}

function sql_insert($params) {
	return "`small_box_id` = {$params['m']}";
}


function location($params){
	//return "<span>location in mission_level_small_box_items_extend.php</span>";
	global $db;
	$html = '';


	$sql = "select `id`, `mission_level_id` from mission_level_small_box where id = {$params['m']}";
	$mission_level_small_box = db_query($db, $sql);
	if (count($mission_level_small_box) == 0) {
		return $html;
	}

	$sql = "select `id`, `mission_id`, `name` from mission_level where id = {$mission_level_small_box[0]['mission_level_id']}";
	$mission_level = db_query($db, $sql);
	if (count($mission_level) == 0) {
		return $html;
	}

	$sql = "select `id`, `name`, `town_id` from mission where id = {$mission_level[0]['mission_id']}";
	$mission = db_query($db, $sql);
	if (count($mission) == 0) {
		return $html;
	}
	
	
	$sql = "select `id`, `name` from town where id = {$mission[0]['town_id']}";
	$town = db_query($db, $sql);
	if (count($town) == 0) {
		return $html;
	}
	
	
	$html .= html_get_link("所有城镇", "?p=town", false).' -> ';
	$html .= html_get_link($town[0]['name'], "?p=mission&m=".$town[0]['id'], false).' -> ';
	$html .= html_get_link($mission[0]['name'], "?p=mission_level&m=".$mission[0]['id'], false) . '-> ';
	$html .= html_get_link("{$mission_level[0]['name']}", "?p=mission_level_small_box&m=".$mission_level_small_box[0]['mission_level_id'], false);

	return $html;
}


function sql_where($params) {
	if (!isset($params['m'])){
		return '';
	}

	return "where `small_box_id`={$params['m']} ";
}

function js_function($params) {
	global $all_items;

	return choose_single_item($all_items, 'item_id');
}

?>