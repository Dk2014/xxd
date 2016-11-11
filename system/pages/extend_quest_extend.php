<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'description' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
	),
	
	'required_quest' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'     => array('params' => array()),
	),
	
	'type' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		//任务类型 1--通关区域评星 2--连续登录 3--元宝购买
		'data' 		=> array( 
			'1' => '通关区域评星', 
			'2' => '连续登陆',
			'3' => '元宝购买',
		),
	),
	'related_mission' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'		=> array('params' => array()),
	),
	'related_npc' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'		=> array('params' => array()),
	),

	'award_item_1' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'		=> array('params' => array()),
	),
	'award_item_2' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range' =>array('params' => array()),
	),
	'award_item_3' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range' =>array('params' => array()),
	),
	'award_item4_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range' =>array('params' => array()),
	),

);

$allItem = get_item_exclude_equip();


function render_type($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_type($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function range_related_mission(){
	global $allMission;
	$tempMissions = $allMission;
	$tempMissions[0] = '无';
	return $tempMissions;
}

function editor_related_mission($column_name, $column_val, $row, $data){
	global $allMission;
	$mission_name = '';
	if (isset($allMission[$column_val])) {
		$mission_name = $allMission[$column_val];
	}

	return editor_single_item($mission_name, $column_val, $column_name);
}

function render_related_mission($column_name, $column_val, $row, $data){
	global $allMission;
	return (isset($allMission[$column_val]) ? $allMission[$column_val] : "0");
}

function range_related_npc(){
	global $allNPCs;
	$tempNPCs = $allNPCs;
	$tempNPCs[0] = '无';
	return $tempNPCs;
}

function editor_related_npc($column_name, $column_val, $row, $data){
	global $allNPCs;
	$role_name = '';
	if (isset($allNPCs[$column_val])) {
		$role_name = $allNPCs[$column_val];
	}

	return editor_single_item($role_name, $column_val, $column_name);
}

function render_related_npc($column_name, $column_val, $row, $data){
	global $allNPCs;

	return (isset($allNPCs[$column_val]) ? $allNPCs[$column_val] : "0");
}

function editor_award_item_1($column_name, $column_val, $row, $data){
	global $allItem;

	$item_name = '';
	if (isset($allItem[$column_val])) {
		$item_name = $allItem[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item_1($column_name, $column_val, $row, $data){
	global $allItem;

	return (isset($allItem[$column_val]) ? $allItem[$column_val] : "0");
}

function editor_award_item_2($column_name, $column_val, $row, $data){
	global $allItem;

	$item_name = '';
	if (isset($allItem[$column_val])) {
		$item_name = $allItem[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item_2($column_name, $column_val, $row, $data){
	global $allItem;

	return (isset($allItem[$column_val]) ? $allItem[$column_val] : "无");
}

function editor_award_item_3($column_name, $column_val, $row, $data){
	global $allItem;

	$item_name = '';
	if (isset($allItem[$column_val])) {
		$item_name = $allItem[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item_3($column_name, $column_val, $row, $data){
	global $allItem;

	return (isset($allItem[$column_val]) ? $allItem[$column_val] : "无");
}
function editor_award_item4_id($column_name, $column_val, $row, $data){
	global $allItem;

	$item_name = '';
	if (isset($allItem[$column_val])) {
		$item_name = $allItem[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item4_id($column_name, $column_val, $row, $data){
	global $allItem;

	return (isset($allItem[$column_val]) ? $allItem[$column_val] : "无");
}
function range_award_item_1(){
	global $allItem;
	$all_item_temp = $allItem;
	$all_item_temp[0] = '无';
	return $all_item_temp;
}
function range_award_item_2(){
	global $allItem;
	$all_item_temp = $allItem;
	$all_item_temp[0] = '无';
	return $all_item_temp;
}
function range_award_item_3(){
	global $allItem;
	$all_item_temp = $allItem;
	$all_item_temp[0] = '无';
	return $all_item_temp;
}
function range_award_item4_id(){
	global $allItem;
	$all_item_temp = $allItem;
	$all_item_temp[0] = '无';
	return $all_item_temp;
}

function editor_description($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_description($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function range_required_quest(){
	global $allQuests;
	$tempQuests = $allQuests;
	$tempQuests[0] = '无';
	return $tempQuests;
}

function editor_required_quest($column_name, $column_val, $row, $data){
	global $allQuests;
	$quest_name = '';
	if (isset($allQuests[$column_val])) {
		$quest_name = $allQuests[$column_val];
	}

	return editor_single_item($quest_name, $column_val, $column_name);
}

function render_required_quest($column_name, $column_val, $row, $data){
	global $allQuests;

	return (isset($allQuests[$column_val]) ? $allQuests[$column_val] : "0");
}

function sql_where($params) {
	return "";
}

$allNPCs = get_all_npc_role();
$allMission = get_all_mission();
$allQuests = get_all_quest();

function jsFunction($params) {
	global $allItem;
	global $allNPCs;
	global $allQuests;

	global $allMission;

	$html = '
$("select").change( function(){
		$(this).css("color","#0033FF");
});';

	$html .= choose_single_item($allNPCs, 'related_npc');
	$html .= choose_single_item($allQuests, 'required_quest');
	$html .= choose_single_item($allMission, 'related_mission');
	$html .= choose_single_item($allNPCs, 'related_npc');
	$html .= choose_single_item($allItem, 'award_item_1');
	$html .= choose_single_item($allItem, 'award_item_2');
	$html .= choose_single_item($allItem, 'award_item_3');
	return $html;
}
?>
