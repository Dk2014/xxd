<?php

$extend_columns = array(
/*   '字段' => 配置 */
'building' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
'attrib_type' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
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

//注意和常量同步 database/config/clique.php
$all_attrib_types = array(
	1 => '生命',
	2 => '攻击',
	3 => '防御',
);

function range_attrib_type() {
	global $all_attrib_types;
	return $all_attrib_types;
}

function render_attrib_type($column_name, $column_val, $row, $data) {
	global $all_attrib_types;
	return $all_attrib_types[$column_val];
}

function editor_attrib_type($column_name, $column_val, $row, $data) {
	global $all_attrib_types;
	return html_get_select($column_name,$all_attrib_types,$column_val);
}

?>
