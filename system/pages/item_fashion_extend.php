<?php
include "item_equipment_links.php";

$extend_columns = array(
/*   '字段' => 配置 */
	'desc' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	'quality' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
);

function editor_desc($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_desc($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
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
	return " where `type_id`='16'";
}

function sql_insert($params) {
	return " `type_id`='16'";
}


?>
