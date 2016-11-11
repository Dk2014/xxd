<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(/* 0 */'普通', /* 1 */'精英', /* 2 */'Boss'),
	),


	'flip_horizontal' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('否', '是'),
	),

);


function render_flip_horizontal($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_flip_horizontal($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_type($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_type($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function map_enemy_render($row) {
	$html = html_get_link("怪物组", "?p=mission_enemy&m=".$row['id']."&bt=14", false);
	$html .= " | ";
	$html .= html_get_link("宝箱配置", "?p=mission_level_box&m=".$row['id'], false);
	$html .= " | ";
	$html .= html_get_link("小宝箱", "?p=mission_level_small_box&m=".$row['id'], false);
	$html .= " | ";
	$html .= html_get_link("关卡评星", "?p=level_star&m=" . $row['id'], false);

	return $html;
}

function sql_insert($params) {
	return "`parent_id` = {$params['m']}, `parent_type`= 14, `award_lock`=0, `mission_id`='0', `lock`='0', `order`='0'";
}

function sql_where($params) {
	if (!isset($params['m'])){
		return '';
	}

	return " where `parent_id`={$params['m']} and `parent_type`='14'";
}

?>
