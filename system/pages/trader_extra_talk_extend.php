<?php
$extend_columns = array(
/*   '字段' => 配置 */
	'talk' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);

function render_talk($column_name, $column_val, $row, $data) {
	return html_get_textarea($column_name, $column_val, false, 7, 70);
}
function editor_talk($column_name, $column_val, $row, $data) {
	return html_get_textarea($column_name, $column_val, true, 7, 70);
}

function sql_where($params) {
	return "where trader_id={$params['m']}";
}

function sql_insert($params) {
	return "trader_id={$params['m']}";
}

?>
