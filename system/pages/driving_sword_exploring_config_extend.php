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

function sql_insert($params) {
	return "`parent_id` = {$params['m']}, `parent_type`= 15, `award_lock`=0, `mission_id`='0', `lock`='0', `order`='0', `daily_num`='-1'";
}

function sql_where($params) {
	return " where `parent_id`={$params['m']} and `parent_type`='15'";
}
function map_enemy_render($row) {
	$html = html_get_link("怪物组", "?p=mission_enemy&m=".$row['id']."&bt=15", false);
	return $html;
}
?>
