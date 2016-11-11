<?php
include "skill_links.php";

$extend_columns = array(
/*   '字段' => 配置 */
	//'child_type' => array(
	//	'editor' => array('params' => array()),
	//	'render' => array('params' => array()),
	//	'data' => array(),
	//),

	'child_type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	
	'jump_attack' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('否','是'),
	),
	'target' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(1=>'单体', 2=>'全体', 3=>'横向', 4=>'纵向'),
	),
);

// 绝招子类型 
$skillTypes[0] = '无';
$skillTypes[1] = '进攻';
$skillTypes[3] = '防御';
$skillTypes[4] = '治疗';
$skillTypes[5] = '辅助';
$skillTypes[6] = '破甲';

function range_child_type(){
	global $skillTypes;
	return $skillTypes;
}

function render_child_type($column_name, $column_val, $row, $data){
	global $skillTypes;
	return $skillTypes[$row['child_type']];
}

function editor_child_type($column_name, $column_val, $row, $data){
	global $skillTypes;
	return html_get_select($column_name,$skillTypes,$column_val);
}

function render_target($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_target($column_name, $column_val, $row, $data){
	return html_get_select($column_name,$data,$column_val);
}

function render_jump_attack($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_jump_attack($column_name, $column_val, $row, $data){
	return html_get_select($column_name,$data,$column_val);
}

function foot() {
	require dirname(__FILE__).'/skill_editor.php';
	require dirname(__FILE__).'/skill_info_editor.php';
}

function skill_opera($row) {
	return '<a href="javascript:;" onclick="open_editor(\''.$row['name'].'\', \''.$row['info'].'\', \''.$row['id'].'\', false)">配置绝招 </a>';
}

function skill_info($row) {
    return '<a href="javascript:;" onclick="edit_skill_info(\''.$row['id'].'\')">绝招描述</a>';
}

function sql_where($params) {
	return " where `type` = 5 ORDER BY `order`";
}

function sql_insert($params) {
	return "`type` = 5";
}

?>
