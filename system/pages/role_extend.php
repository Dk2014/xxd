<?php

$extend_columns = array(
	'is_special' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'skill_id1' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'skill_id2' =>  array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
);


// 角色类型
$roleTypes[1] = '主角';
$roleTypes[2] = '伙伴';

// 是否特殊伙伴
$isSpecial[0] = '不是特殊伙伴';
$isSpecial[1] = '是特殊伙伴';


// 绝招 
$sql = "select * from `skill` where `type` = 1";
$skills = db_query($db,$sql);
foreach ($skills as $value) {
	$skill[$value['id']] = $value['name'];
}

function range_type(){
	global $roleTypes;
	return $roleTypes;
}
 
function range_is_special(){
	global $isSpecial;
	return $isSpecial;
}
function range_skill_id1(){
	global $skill;
	$skill_temp = $skill;
	$skill_temp[0] = '无';
	return $skill_temp;
}
function range_skill_id2(){
	global $skill;
	$skill_temp = $skill;
	$skill_temp[0] = '无';
	return $skill_temp;
}

function sql_where($params) {
	return " order by `type` desc,`buddy_level`";
}

function render_type($column_name, $column_val, $row, $data){
	switch($row['type']){
		case "1":
			return '主角';
		case "2":
			return '伙伴';
	}
}

function editor_type($column_name, $column_val, $row, $data){

	global $roleTypes;

	$code = '';
	if ($row != null) {
		$code .= '<select name="type[]" >';
		$code .= '<option value="0" selected="selected">无</option>';
	foreach ($roleTypes as $key => $value){

		if ($key == $row['type']) {
			$code .= '<option value="'.$key.'"  selected="selected">'.$value.'</option>';
		} else {
			$code .= '<option value="'.$key.'" >'.$value.'</option>';
		}
	}
		$code .= '</select>';
	} else {
		$code .= '<select name="type[]" >';
		$code .= '<option value="1" >主角</option>';
		$code .= '<option value="2" selected="selected">伙伴</option>';
		$code .= '</select>';
	}
	
	return $code;
}

function render_is_special($column_name, $column_val, $row, $data){
	switch($column_val){
		case "0":
			return '不是特殊伙伴';
		case "1":
			return '是特殊伙伴';
	}
}

function editor_is_special($column_name, $column_val, $row, $data){
	global $isSpecial;

	$code = '';
	if ($row != null) {
		$code .= '<select name="is_special[]" >';
		$code .= '<option value="0" selected="selected">无</option>';
	foreach ($isSpecial as $key => $value){

		if ($key == $row['is_special']) {
			$code .= '<option value="'.$key.'"  selected="selected">'.$value.'</option>';
		} else {
			$code .= '<option value="'.$key.'" >'.$value.'</option>';
		}
	}
		$code .= '</select>';
	} else {
		$code .= '<select name="is_special[]" >';
		$code .= '<option value="0" >不是特殊伙伴</option>';
		$code .= '<option value="1" selected="selected">是特殊伙伴</option>';
		$code .= '</select>';
	}
	
	return $code;
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

function rendere_mission_lock($column_name, $column_val, $row, $data) {
	global $mission;
	$html = '';
	if (isset($mission[$row['mission_lock']])){
		$html .= '<a>'.$mission[$row['mission_lock']].'</a>';
	}
	return $html;
}

function editor_mission_lock($column_name, $column_val, $row, $data){
	global $missions;
	$code  = '<select name="mission_lock[]" >';
	$code .= '<option value="0" selected="selected">无</option>';
	foreach ($missions as $value){
		if ($value['lock'] == $row['mission_lock']) {
			$code .= '<option value="'.$value['lock'].'"  selected="selected">'.$value['name'].'</option>';
		} else {
			$code .= '<option value="'.$value['lock'].'" >'.$value['name'].'</option>';
		}
	}
	$code .= '</select>';
	return $code;
}


function buddy_data($row) {
    $rtn='';
	switch($row['type']){
		case "1":
			$rtn.='<a href="?p=main_role_level&m='.$row['id'].'">主角等级数据</a>';
			break;
		case "2": 
			$rtn.='<a href="?p=main_role_level&m='.$row['id'].'">伙伴等级数据</a>';
			break;
	}
	$rtn.=' | <a href="?p=role_friendship&m='.$row['id'].'">角色羁绊数据</a>';
	$rtn.=' | <a href="?p=awaken_graphic&m='.$row['id'].'">觉醒图谱</a>';
	return $rtn;
}

?>
