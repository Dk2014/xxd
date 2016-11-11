<?php
include "item_links.php";

$extend_columns = array(
/*   '字段' => 配置 */
	'target_item_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),

	'item_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
);

$allItem = array();
$res = db_query($db, "select `id`, `name` from `item` where `type_id` != 10;");
foreach ($res as $value)
{
	$allItem[$value['id']] = $value['name'];
}

$useItem = array();
$res = db_query($db, "select `id`, `name` from `item` where `type_id` = 12;");
foreach ($res as $value)
{
	$useItem[$value['id']] = $value['name'];
}

function range_target_item_id(){
	global $allItem;
	return $allItem;
}

function range_item_id(){
	global $useItem;
	return $useItem;
}

function render_target_item_id($column_name, $column_val, $row, $data) {
	global $allItem;
	return $allItem[$column_val];
}

function editor_target_item_id($column_name, $column_val, $row, $data) {
	global $allItem;
	return html_get_select($column_name,$allItem,$column_val);
}

function render_item_id($column_name, $column_val, $row, $data) {
	global $useItem;
	return $useItem[$column_val];
}

function editor_item_id($column_name, $column_val, $row, $data) {
	global $useItem;
	return html_get_select($column_name,$useItem,$column_val);
}

?>