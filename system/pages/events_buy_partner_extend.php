<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'patner_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range' =>array('params' => array()),
	),
	'skill_id1' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range' =>array('params' => array()),
	),

	'skill_id2' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range' =>array('params' => array()),
	),

);

// 人物
$allRole = array();
$sql = "select `id` ,`name` from `role`";
$rows = db_query($db, $sql);
$allRole[0] = '无';
foreach ($rows as $key => $value) {
	$allRole[$value['id']] = $value['name'];
}

// 绝招 
$sql = "select * from `skill` where `type` = 1";
$skills = db_query($db,$sql);
foreach ($skills as $value) {
	$skill[$value['id']] = $value['name'];
}

function range_patner_id(){
	global $allRole;
	return $allRole;
}

function editor_patner_id($column_name, $column_val, $row, $data){
	global $allRole;
	return html_get_select($column_name, $allRole, $column_val);
}

function render_patner_id($column_name, $column_val, $row, $data){
	global $allRole;
	return $allRole[$column_val];
}

function render_skill_id1($column_name, $column_val, $row, $data) {
	global $skill;
	$html = '';
	if (isset($skill[$row['skill_id1']])){
		$html .= '<a>'.$skill[$row['skill_id1']].'</a>';
	}
	return $html;
}

function editor_skill_id1($column_name, $column_val, $row, $data){
	global $skills;
	$code  = '<select name="skill_id1[]" >';
	$code .= '<option value="0" selected="selected">无</option>';
	foreach ($skills as $value){
		if ($value['id'] == $row['skill_id1']) {
			$code .= '<option value="'.$value['id'].'"  selected="selected">'.$value['name'].'</option>';
		} else {
			$code .= '<option value="'.$value['id'].'" >'.$value['name'].'</option>';
		}
	}
	$code .= '</select>';
	return $code;
}

function render_skill_id2($column_name, $column_val, $row, $data) {
	global $skill;
	$html = '';
	if (isset($skill[$row['skill_id2']])){
		$html .= '<a>'.$skill[$row['skill_id2']].'</a>';
	}
	return $html;
}

function editor_skill_id2($column_name, $column_val, $row, $data){
	global $skills;
	$code  = '<select name="skill_id2[]" >';
	$code .= '<option value="0" selected="selected">无</option>';
	foreach ($skills as $value){
		if ($value['id'] == $row['skill_id2']) {
			$code .= '<option value="'.$value['id'].'"  selected="selected">'.$value['name'].'</option>';
		} else {
			$code .= '<option value="'.$value['id'].'" >'.$value['name'].'</option>';
		}
	}
	$code .= '</select>';
	return $code;
}

?>