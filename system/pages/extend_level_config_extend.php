<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(/* 0 */'普通', /* 1 */'精英', /* 2 */'Boss'),
	),

	'sub_type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(/* 0 */'无', /* 1 */'铜钱关卡', /* 2 */'经验关卡'),
	),	

	'daily_num' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),

	'box_dir' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('朝左', '朝右'),
	),

	'flip_horizontal' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('否', '是'),
	),

	'award_box' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('否', '是'),
	),
);


function render_sub_type($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_sub_type($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_award_box($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_award_box($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_type($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_type($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_flip_horizontal($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_flip_horizontal($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_daily_num($column_name, $column_val, $row, $data){
	if ($column_val < 0) {
		return "不限";
	}

	return $column_val;
}

function editor_daily_num($column_name, $column_val, $row, $data){
	if ($column_val == 0)
		$column_val = -1;
	return html_get_input($column_name, $column_val);
}

function editor_box_dir($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_box_dir($column_name, $column_val, $row, $data){
	return $data[$column_val];
}


function map_enemy_render($row) {
	$html = html_get_link("怪物组", "?p=mission_enemy&m=".$row['id']."&bt={$_GET['m2']}", false);
	$html .= " | ";
	$html .= html_get_link("宝箱配置", "?p=mission_level_box&m=".$row['id'], false);

	return $html;
}

function sql_insert($params) {
	return "`mission_id` = 0, `type` = 0, `lock` = 0, `award_lock`=0, `parent_type`={$params['m2']}, `parent_id` = {$params['m']}, `order`='0'";
}

function location($params){
	return '';
}

function sql_where($params) {
	return " where `parent_type`={$params['m2']} and `parent_id`={$params['m']}";
}

?>
