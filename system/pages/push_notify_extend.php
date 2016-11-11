<?php
$extend_columns = array(
/*   '字段' => 配置 */
	'content' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	'trigger_time' => array(
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


function editor_trigger_time($column_name, $column_val, $row, $data){
	return html_get_input($column_name, $column_val, 120);
}

function render_trigger_time($column_name, $column_val, $row, $data){
	if ($column_val == '-1') {
		return '不定时推送';
	} else if ($column_val == '-2') {
		return '永不推送';
	}
	return gmdate("H:i:s", $column_val);
}

?>
