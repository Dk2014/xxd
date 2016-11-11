<?php
include "item_links.php";

$extend_columns = array(
/*   '字段' => 配置 */
	'type_id' => array(
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
	'can_use' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('不能','能'),
	),

	'panel' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
);

$allType = array();
$ids = '(';
$got = 0;

$sql = "select `id`, `name` from `item_type` where `name` in ('材料','特殊材料','功能道具');";

$allTypeTmp = db_query($db, $sql);

foreach ($allTypeTmp as $value)
{
	$allType[$value['id']] = $value['name'];
	$ids =$ids.strval($value['id']).',';
	$got += 1;
}

if ($got > 0) {
	$ids = substr($ids, 0, strlen($ids)-1);
	$ids = $ids.')';
}

function range_type_id(){
	global $allType;
	return $allType;
}
function range_quality(){
	global $all_quality;
	return $all_quality;
}

function range_panel(){
	global $panels;
	return $panels;
}

function render_type_id($column_name, $column_val, $row, $data) {
	global $allType;
	return $allType[$column_val];
}

function editor_type_id($column_name, $column_val, $row, $data) {
	global $allType;
	return html_get_select($column_name,$allType,$column_val);
}

function render_quality($column_name, $column_val, $row, $data) {
	global $all_quality;
	return $all_quality[$column_val];
}

function editor_quality($column_name, $column_val, $row, $data) {
	global $all_quality;
	return html_get_select($column_name,$all_quality,$column_val);
}

function render_can_use($column_name, $column_val, $row, $data) {
	return $data[$column_val];
}

function editor_can_use($column_name, $column_val, $row, $data) {
	return html_get_select($column_name,$data,$column_val);
}

function render_panel($column_name, $column_val, $row, $data) {
	global $panels;
	return $panels[$column_val];
}

function editor_panel($column_name, $column_val, $row, $data) {
	global $panels;
	return html_get_select($column_name,$panels,$column_val);
}

function sql_where($params) {
	global $ids;
	global $got;

	if ($got == 1) {
		$ids = substr($ids,1,strlen($ids)-2);
		return " where `type_id` = {$ids} order by `type_id`,`level`,`quality`";
	} else if ($got > 1) {
		return " where `type_id` in {$ids} order by `type_id`,`level`,`quality`";
	}
	return "";
}

?>