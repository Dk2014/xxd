<?php
include "item_equipment_links.php";

$extend_columns = array(
/*   '字段' => 配置 */
	'type_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),

	'equip_type_id' => array(
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

	'equip_role_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	
	'appendix_level' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
);

$allType = array();
$ids = '(';
$got = 0;

$sql = "select `id`, `name` from `item_type` where `name` in ('武器','战袍','靴子','饰品');";

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

$roles = db_query($db,"select * from `role` where type = 2");
foreach ($roles as $value) {
	$role[$value['id']] = $value['name'];
}
$role[0] = "无";
$role[-1] = "主角";

$appendixLevels = db_query($db,"select * from `equipment_appendix` order by `level` asc");
$appendixLevel[0] = "0";
foreach ($appendixLevels as $value) {
	$appendixLevel[$value['level']] = $value['level'];
}

function range_type_id(){
	global $allType;
	return $allType;
}

function range_quality(){
	global $all_quality;
	return $all_quality;
}

function range_equip_role_id(){
	global $role;
	return $role;
}

function range_equip_type_id(){
	global $equipmentLevelType;
	return $equipmentLevelType;
}

function range_appendix_level(){
	global $appendixLevel;
	return $appendixLevel;
}

function render_type_id($column_name, $column_val, $row, $data) {
	global $allType;
	return $allType[$column_val];
}

function editor_type_id($column_name, $column_val, $row, $data) {
	global $allType;
	return html_get_select($column_name,$allType,$column_val);
}

function render_equip_type_id($column_name, $column_val, $row, $data) {
	global $equipmentLevelType;
	return $equipmentLevelType[$column_val];
}

function editor_equip_type_id($column_name, $column_val, $row, $data) {
	global $equipmentLevelType;
	return html_get_select($column_name,$equipmentLevelType,$column_val);
}

function render_quality($column_name, $column_val, $row, $data) {
	global $all_quality;
	return $all_quality[$column_val];
}

function editor_quality($column_name, $column_val, $row, $data) {
	global $all_quality;
	return html_get_select($column_name,$all_quality,$column_val);
}

function render_equip_role_id($column_name, $column_val, $row, $data) {
	global $role;
	return $role[$column_val];
}

function editor_equip_role_id($column_name, $column_val, $row, $data) {
	global $role;
	return html_get_select($column_name,$role,$column_val);
}


function render_appendix_level($column_name, $column_val, $row, $data) {
	global $appendixLevel;
	return $appendixLevel[$column_val];
}

function editor_appendix_level($column_name, $column_val, $row, $data) {
	global $appendixLevel;
	return html_get_select($column_name,$appendixLevel,$column_val);
}


function sql_where($params) {
	global $ids;
	global $got;

	if ($got == 1) {
		$ids = substr($ids,1,strlen($ids)-2);
		return " where `type_id` = {$ids} order by `equip_type_id`,`quality`,`type_id`,`id` ASC";
	} else if ($got > 1) {
		return " where `type_id` in {$ids} order by `equip_type_id`,`quality`,`type_id`,`id` ASC";
	}
	return "";
}

?>