<?php

$extend_columns = array(
	/* '字段' => 配置 */
	'type' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array('等级', '功能权值'),
	),
	'tips' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(),
	),	
);

function render_type($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_type($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_tips($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function editor_tips($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function sql_where($params) {
	return " order by `order` desc";
}
