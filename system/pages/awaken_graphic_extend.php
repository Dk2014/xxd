<?php

$extend_columns = array(
	'attr_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'direct' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
);

$awaken_ls = array();
$awaken_ls_t = db_query($db, "select `id`, `name` from `awaken_attr`;");
foreach ($awaken_ls_t as $row) {
	$awaken_ls[$row['id']] = $row['name'];
}

function range_attr_id() {
	global $awaken_ls;
	return $awaken_ls;
}

function editor_attr_id($col_nam, $col_val, $row, $data) {
	global $awaken_ls;
	
	$code = '<select name="attr_id[]">';
	foreach ($awaken_ls as $k => $v) {
		if ($k == $row['attr_id']) {
			$code .= "<option value=\"{$k}\" selected>{$v}</option>";
		} else {
			$code .= "<option value=\"{$k}\">{$v}</option>";
		}
	}
	$code .= '</select>';
	return $code;
}

function render_attr_id($col_nam, $col_val, $row, $data) {
	global $awaken_ls;
	return $awaken_ls[$row['attr_id']];
}

$directions = array(
	0 => '左',
	1 => '上',
	2 => '右',
	3 => '下',
);

function range_direct() {
	global $directions;
	return $directions;
}

function editor_direct($col_nam, $col_val, $row, $data) {
	global $directions;
	
	$code = '<select name="direct[]">';
	foreach ($directions as $k => $v) {
		if ($k == $row['direct']) {
			$code .= "<option value=\"{$k}\" selected>{$v}</option>";
		} else {
			$code .= "<option value=\"{$k}\">{$v}</option>";
		}
	}
	$code .= '</select>';
	return $code;
}

function render_direct($col_nam, $col_val, $row, $data) {
	global $directions;
	return $directions[$row['direct']];
}
