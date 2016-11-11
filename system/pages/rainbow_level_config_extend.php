<?php

$LEVEL_TYPE = 12;

$extend_columns = array(
/*   '字段' => 配置 */
	'type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(/* 0 */'普通', /* 1 */'精英', /* 2 */'Boss'),
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

function editor_box_dir($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_box_dir($column_name, $column_val, $row, $data){
	return $data[$column_val];
}


function map_enemy_render($row) {
	global $LEVEL_TYPE;
	$html = html_get_link("怪物组", "?p=mission_enemy&m=".$row['id']."&bt={$LEVEL_TYPE}", false);
	$html .= " | ";
	$html .= html_get_link("宝箱配置", "?p=rainbow_level_award&m=".$row['id'], false);
	return $html;
}

function sql_insert($params) {
	global $LEVEL_TYPE;
	return " `parent_id` = {$params['m']}, `parent_type`= {$LEVEL_TYPE}, `award_lock`=0, `mission_id`='0', `lock`='0', `daily_num`='-1', `physical`='0', `award_key`='0'";
}

/*
function location($params){
	return "";
	global $db;
	$html = '';

	$sql = "select `mission_id` from `hard_level` where id = {$params['m']}";
	$hard_level = db_query($db, $sql);
	if (count($hard_level) == 0) {
		return $html;
	}
	
	
	$sql = "select `name` from mission where id = {$hard_level[0]['mission_id']}";
	$mission_level = db_query($db, $sql);
	if (count($mission_level) == 0) {
		return $html;
	}
	
	$html .= html_get_link("难度关卡", "?p=hard_level", false).' -> ';
	$html .= $mission_level[0]['name'] . "的阴影";
	
	return $html;
}
 */

function sql_where($params) {
	global $LEVEL_TYPE;
	if (!isset($params['m'])){
		return '';
	}

	return " where `parent_id`={$params['m']} and `parent_type`='{$LEVEL_TYPE}' order by `order`";
}

?>
