<?php

$extend_columns = array(
	'desc' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);
//查找出所有城镇区域
$sql = "select `id`, `name` from `mission`;";
$missionRows = db_query($db, $sql);
$allMissions = array();
foreach($missionRows as $row) {
	$allMissions[$row['id']] = $row['name'];
}

function render_mission_id($column_name, $column_val, $row, $data) {
	global $allMissions;
	$html = '';
	if (isset($allMissions[$row['mission_id']])){
		$html .= $allMissions[$row['mission_id']];
	}
	return $html;
}

function editor_mission_id($column_name, $column_val, $row, $data){
	global $missionRows;
	$code = '';
	$code .= '<select name="mission_id[]" >';
	$code .= '<option value="0" selected="selected">无</option>';
	if ($row != null) {
		foreach ($missionRows as $value){
			if ($value['id'] == $row['mission_id']) {
				$code .= '<option value="'.$value['id'].'"  selected="selected">'.$value['name'].'</option>';
			} else {
				$code .= '<option value="'.$value['id'].'" >'.$value['name'].'</option>';
			}
		}
	} else {
		foreach ($missionRows as $value){
			$code .= '<option value="'.$value['id'].'" >'.$value['name'].'</option>';
		}
	}
	$code .= '</select>';
	return $code;
}


function editor_desc($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_desc($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}


function leveltable($row) {
	//var_dump($row);
	//die();
	return html_get_link("关卡", "?p=hard_level_config&m=".$row['id'], true);
}

function sql_insert($params) {
	return " `town_id`={$params['m']} ";
}

function sql_where($params) {
	return "where `town_id`={$params['m']} ";
}

function location($params) {
	global $db;
	$html = '';
	$sql = "select `id`, `name` from town where id = {$params['m']}";
	$town = db_query($db, $sql);
	if (count($town) == 0) {
		return $html;
	}
	return html_get_link("所有城镇", "?p=town", false).' -> '.$town[0]['name'];
}




?>
