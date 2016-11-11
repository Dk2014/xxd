<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'description' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
	),
	'conversion_reciving_quest' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
	),
	'conversion_recived_quest' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
	),

	'conversion_finish_quest' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
	),
	
	'showup_main_quest' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'     => array('params' => array()),
	),
	'disappear_main_quest' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'     => array('params' => array()),
	),

	'publish_npc' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'     => array('params' => array()),
	),
	
	'type' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array( 
			'1' => 'NPC对话', 
			'2' => '消灭敌人',
			'3' => '通关关卡',
			'4' => '收集物品',
			'5' => '展示物品',
			'6' => '区域评星',
			'7' => '招募伙伴',
		),
	),


	'npc_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'     => array('params' => array()),
	),
	'mission_level_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'     => array('params' => array()),
	),
	'enemy_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'     => array('params' => array()),
	),

	'mission_enemy_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'     => array('params' => array()),
	),
	'quest_item_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'     => array('params' => array()),
	),

	'require_item_type' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(
			'1' => '装备',
			'2' => '普通物品',
		),
	),
	'require_item_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'     => array('params' => array()),
	),


	'award_item_1' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'     => array('params' => array()),
	),
	'award_item_2' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'     => array('params' => array()),
	),
	'role_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'     => array('params' => array()),
	),
	'npc_role' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'     => array('params' => array()),
	),
	'mission_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'     => array('params' => array()),
	),
	'require_serial_number' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'     => array('params' => array()),
	),
	'serial_number' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'     => array('params' => array()),
	),
);

$allMission = get_all_mission();
$allItem = get_item_exclude_equip();
$allItemWithEquip = get_all_item_by_types(array(11,10), true); //获取所有普通物品
$allNPCRole = get_all_npc_role();
$allRole = get_all_role();
$allNPCs = get_all_npc();
$allNPCs[0] = '无';
$allQuests = get_all_quest();
$allQuests[0] = '一直有效';
$allMissionLevel = get_all_level_by_type(0);
$allEnemyGroup = get_all_enemy_group();
$allEnemy = get_all_enemy_name();
$allSerial = get_all_addition_quest_serial();

function render_require_item_type($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_require_item_type($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_type($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_type($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function editor_require_serial_number($column_name, $column_val, $row, $data){
	global $allSerial;
	$name = '';
	if (isset($allSerial[$column_val])) {
		$name = $allSerial[$column_val];
	}

	return editor_single_item($name, $column_val, $column_name);
}
function range_serial_number(){
	global $allSerial;
	$temp_serials = $allSerial;
	$temp_serials[0] = '无';
	return $temp_serials;
}

function range_require_serial_number(){
	global $allSerial;
	$temp_serials = $allSerial;
	$temp_serials[0] = '无';
	return $temp_serials;
}

function range_publish_npc(){
	global $allNPCs;
	return $allNPCs;
}

function range_npc_id(){
	global $allNPCs;
	return $allNPCs;
}

function render_require_serial_number($column_name, $column_val, $row, $data){
	global $allSerial;

	return (isset($allSerial[$column_val]) ? $allSerial[$column_val] : "0");
}

function editor_serial_number($column_name, $column_val, $row, $data){
	global $allSerial;
	$name = '';
	if (isset($allSerial[$column_val])) {
		$name = $allSerial[$column_val];
	}

	return editor_single_item($name, $column_val, $column_name);
}

function render_serial_number($column_name, $column_val, $row, $data){
	global $allSerial;

	return (isset($allSerial[$column_val]) ? $allSerial[$column_val] : "0");
}

function editor_mission_id($column_name, $column_val, $row, $data){
	global $allMission;
	$enemyName = '';
	if (isset($allMission[$column_val])) {
		$enemyName = $allMission[$column_val];
	}

	return editor_single_item($enemyName, $column_val, $column_name);
}

function render_mission_id($column_name, $column_val, $row, $data){
	global $allMission;

	return (isset($allMission[$column_val]) ? $allMission[$column_val] : "0");
}

function range_mission_id(){
	global $allMission;
	$temp_mission = $allMission;
	$temp_mission[0] = '无';
	return $temp_mission;
}

function editor_role_id($column_name, $column_val, $row, $data){
	global $allRole;
	$enemyName = '';
	if (isset($allRole[$column_val])) {
		$enemyName = $allRole[$column_val];
	}

	return editor_single_item($enemyName, $column_val, $column_name);
}

function render_role_id($column_name, $column_val, $row, $data){
	global $allRole;

	return (isset($allRole[$column_val]) ? $allRole[$column_val] : "0");
}

function range_role_id(){
	global $allRole;
	$temp_role = $allRole;
	$temp_role[0] = '无';
	return $temp_role;
}

function editor_npc_role($column_name, $column_val, $row, $data){
	global $allNPCRole;
	$enemyName = '';
	if (isset($allNPCRole[$column_val])) {
		$enemyName = $allNPCRole[$column_val];
	}

	return editor_single_item($enemyName, $column_val, $column_name);
}

function render_npc_role($column_name, $column_val, $row, $data){
	global $allNPCRole;

	return (isset($allNPCRole[$column_val]) ? $allNPCRole[$column_val] : "0");
}

function range_npc_role(){
	global $allNPCRole;
	$temp_npc_role = $allNPCRole;
	$temp_npc_role[0] = '无';
	return $temp_npc_role;
}

function editor_enemy_id($column_name, $column_val, $row, $data){
	global $allEnemy;
	$enemyName = '';
	if (isset($allEnemy[$column_val])) {
		$enemyName = $allEnemy[$column_val];
	}

	return editor_single_item($enemyName, $column_val, $column_name);
}

function render_enemy_id($column_name, $column_val, $row, $data){
	global $allEnemy;

	return (isset($allEnemy[$column_val]) ? $allEnemy[$column_val] : "0");
}

function range_enemy_id(){
	global $allEnemy;
	$temp_enemy = $allEnemy;
	$temp_enemy[0] = '无';
	return $temp_enemy;
}

function editor_mission_enemy_id($column_name, $column_val, $row, $data){
	global $allEnemyGroup;
	$enemyGroupName = '';
	if (isset($allEnemyGroup[$column_val])) {
		$enemyGroupName = $allEnemyGroup[$column_val];
	}

	return editor_single_item($enemyGroupName, $column_val, $column_name);
}

function render_mission_enemy_id($column_name, $column_val, $row, $data){
	global $allEnemyGroup;

	return (isset($allEnemyGroup[$column_val]) ? $allEnemyGroup[$column_val] : "0");
}

function range_mission_enemy_id(){
	global $allEnemyGroup;
	return $allEnemyGroup;
}

function editor_mission_level_id($column_name, $column_val, $row, $data){
	global $allMissionLevel;
	$mission_level_name = '';
	if (isset($allMissionLevel[$column_val])) {
		$mission_level_name = $allMissionLevel[$column_val];
	}

	return editor_single_item($mission_level_name, $column_val, $column_name);
}

function render_mission_level_id($column_name, $column_val, $row, $data){
	global $allMissionLevel;

	return (isset($allMissionLevel[$column_val]) ? $allMissionLevel[$column_val] : "0");
}

function range_mission_level_id(){
	global $allMissionLevel;
	$temp_mission_level = $allMissionLevel;
	$temp_mission_level[0] = '无';
	return $temp_mission_level;
}

function editor_publish_npc($column_name, $column_val, $row, $data){
	global $allNPCs;
	$role_name = '';
	if (isset($allNPCs[$column_val])) {
		$role_name = $allNPCs[$column_val];
	}

	return editor_single_item($role_name, $column_val, $column_name);
}

function render_publish_npc($column_name, $column_val, $row, $data){
	global $allNPCs;

	return (isset($allNPCs[$column_val]) ? $allNPCs[$column_val] : "0");
}

function editor_npc_id($column_name, $column_val, $row, $data){
	global $allNPCs;
	$role_name = '';
	if (isset($allNPCs[$column_val])) {
		$role_name = $allNPCs[$column_val];
	}

	return editor_single_item($role_name, $column_val, $column_name);
}

function render_npc_id($column_name, $column_val, $row, $data){
	global $allNPCs;

	return (isset($allNPCs[$column_val]) ? $allNPCs[$column_val] : "0");
}

function editor_require_item_id($column_name, $column_val, $row, $data){
	global $allItemWithEquip;

	$item_name = '';
	if (isset($allItemWithEquip[$column_val])) {
		$item_name = $allItemWithEquip[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_require_item_id($column_name, $column_val, $row, $data){
	global $allItemWithEquip;

	return (isset($allItemWithEquip[$column_val]) ? $allItemWithEquip[$column_val] : "0");
}

function range_require_item_id(){
	global $allItemWithEquip;
	$temp_item_equip = $allItemWithEquip;
	$temp_item_equip[0] = '无';
	return $temp_item_equip;
}

function editor_quest_item_id($column_name, $column_val, $row, $data){
	global $allItem;

	$item_name = '';
	if (isset($allItem[$column_val])) {
		$item_name = $allItem[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_quest_item_id($column_name, $column_val, $row, $data){
	global $allItem;

	return (isset($allItem[$column_val]) ? $allItem[$column_val] : "0");
}

function range_quest_item_id(){
	global $allItem;
	$temp_item = $allItem;
	$temp_item[0] = '无';
	return $temp_item;
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

function editor_description($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_description($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}
function editor_conversion_reciving_quest($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_conversion_reciving_quest($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function editor_conversion_recived_quest($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_conversion_recived_quest($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function editor_conversion_finish_quest($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_conversion_finish_quest($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function editor_showup_main_quest($column_name, $column_val, $row, $data){
	global $allQuests;
	$quest_name = '';
	if (isset($allQuests[$column_val])) {
		$quest_name = $allQuests[$column_val];
	}

	return editor_single_item($quest_name, $column_val, $column_name);
}

function render_showup_main_quest($column_name, $column_val, $row, $data){
	global $allQuests;

	return (isset($allQuests[$column_val]) ? $allQuests[$column_val] : "0");
}

function range_showup_main_quest(){
	global $allQuests;
	return $allQuests;
}

function range_disappear_main_quest(){
	global $allQuests;
	return $allQuests;
}

function editor_disappear_main_quest($column_name, $column_val, $row, $data){
	global $allQuests;
	$quest_name = '';
	if (isset($allQuests[$column_val])) {
		$quest_name = $allQuests[$column_val];
	}

	return editor_single_item($quest_name, $column_val, $column_name);
}

function render_disappear_main_quest($column_name, $column_val, $row, $data){
	global $allQuests;

	return (isset($allQuests[$column_val]) ? $allQuests[$column_val] : "0");
}

function sql_where($params) {
	if(isset($params['m'])) {
		if (intval($params['m']) == 0 )  {
			return "where award_lock=0 order by `require_level`, `require_lock`";
		}
		return " where serial_number={$params['m']} order by `require_level`, `require_lock`";
	}
	return " order by `serial_number`,`require_level`, `require_lock`";
}


function jsFunction($params) {
	global $allItem;
	global $allItemWithEquip;
	global $allNPCs;
	global $allEnemyGroup;
	global $allMissionLevel;
	global $allQuests;
	global $allEnemy;
	global $allRole;
	global $allNPCRole;
	global $allMission;
	global $allSerial;


	$html = '
$("select").change( function(){
		$(this).css("color","#0033FF");
});';

	$html .= choose_single_item($allNPCs, 'publish_npc');
	$html .= choose_single_item($allRole, 'role_id');
	$html .= choose_single_item($allMission, 'mission_id');
	$html .= choose_single_item($allSerial, 'serial_number');
	$html .= choose_single_item($allSerial, 'require_serial_number');
	$html .= choose_single_item($allNPCRole, 'npc_role');
	$html .= choose_single_item($allNPCs, 'npc_id');
	$html .= choose_single_item($allQuests, 'showup_main_quest');
	$html .= choose_single_item($allQuests, 'disappear_main_quest');
	$html .= choose_single_item($allEnemyGroup, 'mission_enemy_id');
	$html .= choose_single_item($allEnemy, 'enemy_id');
	$html .= choose_single_item($allMissionLevel, 'mission_level_id');
	$html .= choose_single_item($allItemWithEquip, 'require_item_id');
	$html .= choose_single_item($allItem, 'quest_item_id');
	$html .= choose_single_item($allItem, 'award_item_1');
	$html .= choose_single_item($allItem, 'award_item_2');
	return $html;
}
?>
