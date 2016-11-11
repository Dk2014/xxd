<?php

$extend_columns = array(
/*   '字段' => 配置 */
'building' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
'desc' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(),
	),
'cur_kongfu_desc' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(),
	),
'next_kongfu_desc' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(),
	),
);

$all_buildings = array();
$res = db_query($db, "select `id`, `name` from `clique_building`;");
foreach($res as $row) {
	$all_buildings[$row['id']] = $row['name'];
}

function range_building() {
	global $all_buildings;
	return $all_buildings;
}

function render_building($column_name, $column_val, $row, $data) {
	global $all_buildings;
	return $all_buildings[$column_val];
}

function editor_building($column_name, $column_val, $row, $data) {
	global $all_buildings;
	return html_get_select($column_name,$all_buildings,$column_val);
}

function editor_desc($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_desc($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function editor_cur_kongfu_desc($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_cur_kongfu_desc($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}


function editor_next_kongfu_desc($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_next_kongfu_desc($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}

?>
