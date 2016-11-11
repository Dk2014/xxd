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

	'show_origin' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => $origin_source_type,
	),
);

$allType = array();
$ids = '(';
$got = 0;

$sql = "select `id`, `name` from `item_type` where `name` in ('资源');";

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

function range_quality(){
	global $all_quality;
	return $all_quality;
}

function render_show_origin($column_name, $column_val, $row, $data) {
	return $data[$column_val];
}

function editor_show_origin($column_name, $column_val, $row, $data) {
	return html_get_select($column_name,$data,$column_val);
}
 
function range_type_id(){
	global $allType;
	return $allType;
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
