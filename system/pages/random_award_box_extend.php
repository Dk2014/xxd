<?php


$extend_columns = array(
/*   '字段' => 配置 */
	'award_type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(/* 0 */'铜钱', /* 1 */'道具', /* 2 */'装备', /* 3 */'契约球'),
	),
	'award_chance' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	'award_num' => 'award_chance',
	'item_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range'=>array('params' => array()),
	),
	'must_in_first' => array(
			'editor' => array('params' => array()),
			'render' => array('params' => array()),
			'data' => array(0=>'否', 1=>'是'),
	), 
);

function editor_must_in_first($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_must_in_first($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

$all_items = get_all_item();

function range_item_id(){
	global $all_items;
	$all_items_temp = $all_items;
	$all_items_temp[0] = '无';
	return $all_items_temp;
}

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

function render_award_type($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_award_type($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_award_chance($column_name, $column_val, $row, $data){
	return $column_val;
}

function editor_award_chance($column_name, $column_val, $row, $data){
	if (!isset($column_val))
		$column_val = 0;
	return html_get_input($column_name, $column_val);
}

function sql_insert($params) {
	return "`mission_level_id` = {$params['m']}";
}

function location($params){
	global $db;
	$html = '';

	$sql = "select `mission_id`, `name` from mission_level where id = {$params['m']}";
	$mission_level = db_query($db, $sql);
	if (count($mission_level) == 0) {
		return $html;
	}
	
	
	$sql = "select `id`, `name` from mission where id = {$mission_level[0]['mission_id']}";
	$mission = db_query($db, $sql);
	if (count($mission) == 0) {
		return $html;
	}
	
	$html .= html_get_link("所有城镇", "?p=town", false).' -> ';
	$html .= html_get_link($mission[0]['name'], "?p=mission&m=".$mission[0]['id'], false).' -> ';
	$html .= html_get_link($mission_level[0]['name'], "?p=mission_level&m=".$mission[0]['id'], false).' -> ';
	$html .= $mission_level[0]['name']."关卡宝箱";
	
	return $html;
}

function sql_where($params) {
	if (!isset($params['m'])){
		return '';
	}

	return "where `mission_level_id`={$params['m']} ";
}

function js_function($params) {
	global $all_items;

	return choose_single_item($all_items, 'item_id');
}

?>