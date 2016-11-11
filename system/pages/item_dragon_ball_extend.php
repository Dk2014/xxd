<?php
include "item_links.php";

$extend_columns = array(
/*   '字段' => 配置 */

	'quality' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'can_use' => array(
			'editor' => array('params' => array()),
			'render' => array('params' => array()),
			'data' => array(0 => '否',1 => '是'),
	),
);

function render_can_use($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_can_use($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function range_quality(){
	global $all_quality;
	return $all_quality;
}

function render_quality($column_name, $column_val, $row, $data) {
	global $all_quality;
	return $all_quality[$column_val];
}

function editor_quality($column_name, $column_val, $row, $data) {
	global $all_quality;
	return html_get_select($column_name,$all_quality,$column_val);
}

function sql_where($params) {
	return " where `type_id` = 20 ";
}


function sql_insert($params) {
	return " `type_id` = 20 ";
}

function op_render($row){
	return html_get_link("配置", "?p=item_dragon_ball_config&m=".$row['id'], true);
}

?>
