<?php
$extend_columns = array(
/*   '字段' => 配置 */
	'content' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	'expire_time' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	'effect_time' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);

function editor_content($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_content($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function editor_effect_time($column_name, $column_val, $row, $data){
	return html_get_input($column_name, $column_val, 120);
}

function render_effect_time($column_name, $column_val, $row, $data){
	return  date('Y-m-d H:m:s', $column_val);
}

function editor_expire_time($column_name, $column_val, $row, $data){
	return html_get_input($column_name, $column_val, 120);
}

function render_expire_time($column_name, $column_val, $row, $data){
	if ($column_val == '0') {
		return '永远有效';
	}
	return  date('Y-m-d H:m:s', $column_val);
}

?>
