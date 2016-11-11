<?php
include "item_links.php";

$extend_columns = array(
/*   '字段' => 配置 */

	'quality' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'act_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
);


function range_quality(){
	global $all_quality;
	return $all_quality;
}

function render_quality($column_name, $column_val, $row, $data) {
	global $all_quality;
	return $all_quality[$column_val];
}

function editor_quality($column_name, $column_val, $row, $data) {
	global $all_quality;
	return html_get_select($column_name,$all_quality,$column_val);
}

function range_act_id(){
	global $all_act;
	return $all_act;
}

function render_act_id($column_name, $column_val, $row, $data){
	global $all_act;
	return $all_act[$column_val];
}

function editor_act_id($column_name, $column_val, $row, $data){
	global $all_act;
	return html_get_select($column_name, $all_act, $column_val);
}

function sql_where($params) {
	return " where `type_id` = 21 order by `level`,`quality`";
}


function sql_insert($params) {
	return " `can_use` = 1, `type_id` = 21 ";
}

function op_render($row){
	return '<a href="javascript:;" onclick="open_editor('.$row['id'].')">配置</a>';
}

function foot() {
	require dirname(__FILE__).'/battle_item_editor.php';
}

function item_reflect_config_render($row) {
	return "<a href='?p=item_reflect_config&m={$row['id']}'>功能参数</a>";
}

?>
