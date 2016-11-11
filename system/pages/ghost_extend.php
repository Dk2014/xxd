<?php


$extend_columns = array(
	/*'字段' => 配置 */

	'town_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
				'range' => array('params' => array()),
	),
	'role_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
),
	'fragment_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'init_star' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'quality' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'production_info' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
);

function render_production_info($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function editor_production_info($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}


function range_init_star(){
	global $all_stars;
	return $all_stars;
}

function range_fragment_id(){
	global $all_ghost_fragments;
	return $all_ghost_fragments;
}

function range_role_id(){
	global $all_roles;
	return $all_roles;
}

function range_quality() {
	global $qualityTypes;
	return $qualityTypes;
}

function range_town_id(){
	global $all_towns;
	$temp_towns = $all_towns;
	$temp_towns[0] = '无';
	return $temp_towns;
}

function _ghost_get_all_town() {
	global $db;
	$all_towns = array();

	$tmp = db_query($db, "select `id`, `name` from `town` order by `id`");
	foreach ($tmp as $value)
	{
		$all_towns[$value['id']] = $value['name'];
	}

	return $all_towns;
}

$all_towns = _ghost_get_all_town();
$all_towns[0] = '无';

function render_town_id($column_name, $column_val, $row, $data) {
	global $all_towns;
	return $all_towns[$column_val];
}

function editor_town_id($column_name, $column_val, $row, $data) {
	global $all_towns;
	return html_get_select($column_name,$all_towns,$column_val);
}

?>
