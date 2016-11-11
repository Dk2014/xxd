<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'talk' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	'sold_out_talk' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	'deal_talk' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);

function editor_talk($column_name, $column_val, $row, $data) {
	return html_get_textarea($column_name, $column_val, true, 7, 70);
}

function render_talk($column_name, $column_val, $row, $data) {
	return html_get_textarea($column_name, $column_val, false, 7, 70);
}
function editor_sold_out_talk($column_name, $column_val, $row, $data) {
	return html_get_textarea($column_name, $column_val, true, 7, 70);
}

function render_sold_out_talk($column_name, $column_val, $row, $data) {
	return html_get_textarea($column_name, $column_val, false, 7, 70);
}

function editor_deal_talk($column_name, $column_val, $row, $data) {
	return html_get_textarea($column_name, $column_val, true, 7, 70);
}

function render_deal_talk($column_name, $column_val, $row, $data) {
	return html_get_textarea($column_name, $column_val, false, 7, 70);
}


function trader_town_position($row) {
	return html_get_link(" 城镇坐标 |", "?p=trader_position&m={$row['id']}", false);
}

function trader_config($row) {
	return html_get_link(" 货物配置 |", "?p=trader_grid&m={$row['id']}", false);
}

function trader_extra_talk($row) {
	return html_get_link(" 额外对话 |", "?p=trader_extra_talk&m={$row['id']}", true);
}

function trader_refresh_price($row) {
	return html_get_link(" 刷新价格 ", "?p=trader_refresh_price&m={$row['id']}", true);
}





?>
