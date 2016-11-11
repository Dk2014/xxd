<?php

$extend_columns = array(
	/* '字段' => 配置 */
	'question' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(),
	),
	'answer' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(),
	),	
);

function sql_where($params) {
	return " ORDER BY `order` ASC";
}

function editor_question($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 2, 30);
}

function render_question($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 2, 30);
}


function editor_answer($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 5, 40);
}

function render_answer($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 5, 40);
}

?>