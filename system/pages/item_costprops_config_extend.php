<?php

include "item_links.php";

$extend_columns = array(
/*   '字段' => 配置 */
	'type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('0' => '经验', '1' => '体力', '2' => '友情'),
	),
);


function render_type($column_name, $column_val, $row, $data) {
	return $data[$column_val];
}

function editor_type($column_name, $column_val, $row, $data) {
	return html_get_select($column_name,$data,$column_val);
}

function sql_where($params) {
	return " where `item_id` = {$params['m']} ";
}

function sql_insert($params) {
	return "`item_id` = {$params['m']}";
}

function item_box_content_render($row) {
	return "<a href='?p=item_box_content&m={$row['id']}'>配置</a>"; }
?>
