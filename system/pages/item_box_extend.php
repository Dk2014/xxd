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

	'func_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	'can_batch' => array(
			'editor' => array('params' => array()),
			'render' => array('params' => array()),
			'data' => array(0 => '否',1 => '是'),
	),
);

$all_funcs = array();
$tmp = db_query($db, "select `id`, `name` from `func` where `lock`<30000");
foreach ($tmp as $value)
{
	$all_funcs[$value['id']] = $value['name'];
}

$allType = array();
$ids = '(';
$got = 0;

$sql = "select `id`, `name` from `item_type` where `name` in ('礼包','资源道具');";

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

function render_can_batch($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_can_batch($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function range_quality(){
	global $all_quality;
	return $all_quality;
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

function render_func_id($column_name, $column_val, $row, $data) {
	global $all_funcs;
	return normal_render($all_funcs, $row['func_id']);
}

function editor_func_id($column_name, $column_val, $row, $data) {
	global $all_funcs;

	$item_name = '';
	$item_id = $row['func_id'];

	if (isset($all_funcs[$item_id])) {
		$item_name = $all_funcs[$item_id];
	}

	return editor_single_item($item_name, $item_id, "func_id");
}

function js_function($params) {
	global $all_funcs;
	$html = choose_single_item($all_funcs, 'func_id');
	return $html;
}

function sql_where($params) {
	global $ids;
	global $got;

	if ($got == 1) {
		$ids = substr($ids,1,strlen($ids)-2);
		return " where `type_id` = {$ids} order by `type_id`,`quality`,`level`";
	} else if ($got > 1) {
		return " where `type_id` in {$ids} order by `type_id`,`quality`,`level`";
	}
	return "";
}

function item_box_content_render($row) {
	return "<a href='?p=item_box_content&m={$row['id']}'>宝箱内容</a>";
}
?>