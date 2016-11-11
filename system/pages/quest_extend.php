<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'type' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data'	 	=> array(
								// 任务类型
								/* 0 */ '进入城镇',
								/* 1 */ 'NPC对话',
								/* 2 */ '进入关卡',
								/* 3 */ '通关关卡',
								/* 4 */ '消灭敌人'
							),
	),
	'desc' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
	),
	'mission_drama_talk' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
	),
	'drama_mode' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array('无', '完成后播放'),
	),
	'related_town' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'	=> array('params'=>array()),
	),
	'town_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range' =>array('params' => array()),
	),
	'town_npc_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range' =>array('params' => array()),
	),
	'mission_level_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range' =>array('params' => array()),
	),
	'show_black_curtain' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array('无','有'),
	),
	'enemy_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		//'tip_render'=> array('params' => array()),
		'data' 		=> array(),
		'range' => array('params' => array()),
	),
	'award_item1_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'		=> array('params' => array()),
	),
	'award_item2_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range' =>array('params' => array()),
	),
	'award_item3_id' => array(
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

	'award_town_key' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range' =>array('params' => array()),
	),
	'award_role_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range' =>array('params' => array()),
	),
	'award_func_key' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range' =>array('params' => array()),
	),
	'award_mission_level_lock' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range' =>array('params' => array()),
	),
	'enemy_num' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range' =>array('params' => array()),
	),
	'auto_fight' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array('否','是'),
	),
	'pass_mission_share' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array('否','是'),
	),
);

$allNpc = array('0'=>'无');
$allMissionLevel = array();
$allEnemy = get_all_enemy();
$allItem = get_all_item();
$allMissionLevelLock = array();
$allEnemyIds = array();
$allRole = array();
$allFuncKey = array();

//城镇
$sql = "select `id`,`town_id`,`name` from `town_npc`";
$rows = db_query($db, $sql);
foreach ($rows as $key => $value) {
	$allNpc[$value['id']] = $allTown[$value['town_id']].' - '.$value['name'];
}

// 区域关卡
$sql = "select `id`,`mission_id`,`name`, `lock` from `mission_level` where `mission_id`>0 && `parent_type` = 0";
$rows = db_query($db, $sql);
$allMissionLevel[0] = '无';
$allMissionLevelLock[0] = '无';
foreach ($rows as $key => $value) {
	$allMissionLevel[$value['id']] = $value['name'];
	$allMissionLevelLock[$value['lock']] = $value['name'];
}

// 怪物组
$sql = "
SELECT `half`.`shaded_mission_id` != 0 `is_shaded`
	,`half`.`name` `level_name`
	,`shaded`.`name` `shaded_name`
	,`half`.`order` `order`
	,`half`.`id` `id`
FROM (
	SELECT `level`.`name`
		,`enemy`.`id`
		,`enemy`.`order`
		,`enemy`.`shaded_mission_id`
	FROM `mission_level` `level`
	LEFT JOIN `mission_enemy` `enemy` ON `level`.`mission_id` > 0
		AND `level`.`parent_type` = 0
		AND `enemy`.`mission_level_id` = `level`.`id`
	ORDER BY `level`.`lock`
		,`enemy`.`order`
	) `half`
LEFT JOIN `shaded_mission` `shaded` ON `half`.`shaded_mission_id` = `shaded`.`id`;
";
$rows = db_query($db, $sql);
$allEnemyIds[0] = '无';
foreach ($rows as $_ => $value) {
	if ($value['id'] > 0 && $value['order'] > 0) {
		$allEnemyIds[$value['id']] =
			$value['level_name']
			.($value['is_shaded']>0?'('.$value['shaded_name'].')':'')
			.'第'.$value['order'].'组';
	}
}

// 角色
$sql = "select `id`,`name` from `role`";
$rows = db_query($db, $sql);
$allRole[0] = '无';
foreach ($rows as $key => $value) {
	$allRole[$value['id']] = $value['name'];
}

// 功能权值
$sql = "select `lock`,`name` from `func` where `type`=1";
$rows = db_query($db, $sql);
$allFuncKey[0] = '无';
foreach ($rows as $key => $value) {
	$allFuncKey[$value['lock']] = $value['name'];
}

function range_award_role_id(){
	global $allRole;
	return $allRole;
}

function range_award_mission_level_lock(){
	global $allMissionLevelLock;
	return $allMissionLevelLock;
}

function range_award_town_key(){
	global $allTownLock;
	return $allTownLock;
}

function range_award_func_key(){
	global $allFuncKey;
	return $allFuncKey;
}

function range_enemy_num(){
	global $allEnemyIds;
	return $allEnemyIds;
}

function range_mission_level_id(){
	global $allMissionLevel;
	return $allMissionLevel;
}
function range_enemy_id(){
	global $allEnemy;
	$all_enemy_temp = array();
	$all_enemy_temp[0] = '无';
	foreach ($allEnemy as $key => $value) {
		$all_enemy_temp[$key] = $value['name'];
	}
	return $all_enemy_temp;
}

function range_town_npc_id(){
	global $allNpc;
	return $allNpc;
}

function range_town_id(){
	global $allTown;
	return $allTown;
}

function editor_show_black_curtain($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_show_black_curtain($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_pass_mission_share($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_pass_mission_share($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_auto_fight($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_auto_fight($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_enemy_num($column_name, $column_val, $row, $data){
	global $allEnemyIds;
	$id = $row['enemy_num'];
	$name = '无';
	if (isset($allEnemyIds[$id])) {
		$name = $allEnemyIds[$id];
	}
	return editor_single_item($name, $id, "enemy_num");
}

function render_enemy_num($column_name, $column_val, $row, $data){
	global $allEnemyIds;
	return $allEnemyIds[$column_val];
}

function editor_award_mission_level_lock($column_name, $column_val, $row, $data){
	global $allMissionLevelLock;
	$id = $row['award_mission_level_lock'];
	$name = '无';
	if (isset($allMissionLevelLock[$id])) {
		$name = $allMissionLevelLock[$id];
	}
	return editor_single_item($name, $id, "award_mission_level_lock");
}

function render_award_mission_level_lock($column_name, $column_val, $row, $data){
	global $allMissionLevelLock;
	return $allMissionLevelLock[$column_val];
}

function editor_award_func_key($column_name, $column_val, $row, $data){
	global $allFuncKey;
	return html_get_select($column_name, $allFuncKey, $column_val);
}

function render_award_func_key($column_name, $column_val, $row, $data){
	global $allFuncKey;
	return $allFuncKey[$column_val];
}

function editor_award_role_id($column_name, $column_val, $row, $data){
	global $allRole;
	return html_get_select($column_name, $allRole, $column_val);
}

function render_award_role_id($column_name, $column_val, $row, $data){
	global $allRole;
	return $allRole[$column_val];
}
function range_award_item1_id(){
	global $allItem;
	$all_item_temp = $allItem;
	$all_item_temp[0] = '无';
	return $all_item_temp;
}
function range_award_item2_id(){
	global $allItem;
	$all_item_temp = $allItem;
	$all_item_temp[0] = '无';
	return $all_item_temp;
}
function range_award_item3_id(){
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

function editor_award_item1_id($column_name, $column_val, $row, $data){
	global $allItem;

	$item_name = '';
	if (isset($allItem[$column_val])) {
		$item_name = $allItem[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item1_id($column_name, $column_val, $row, $data){
	global $allItem;

	return (isset($allItem[$column_val]) ? $allItem[$column_val] : "无");
}
function range_related_town(){
	global $allTown;
	return $allTown;
}
function editor_award_item2_id($column_name, $column_val, $row, $data){
	global $allItem;

	$item_name = '';
	if (isset($allItem[$column_val])) {
		$item_name = $allItem[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item2_id($column_name, $column_val, $row, $data){
	global $allItem;

	return (isset($allItem[$column_val]) ? $allItem[$column_val] : "无");
}
function editor_award_item3_id($column_name, $column_val, $row, $data){
	global $allItem;

	$item_name = '';
	if (isset($allItem[$column_val])) {
		$item_name = $allItem[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item3_id($column_name, $column_val, $row, $data){
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

function render_enemy_id($column_name, $column_val, $row, $data){
	global $allEnemy;

	if ($column_val == 0) {
		return "无";
	} else if(isset($allEnemy[$column_val])) {
		return $allEnemy[$column_val]['name'];
		//return html_get_link($allEnemy[$column_val]['name'], '?p=enemy_role&m='.$column_val, true);
	}
}

function editor_enemy_id($column_name, $column_val, $row, $data){
	global $allEnemy;
	return enemy_role_editor($allEnemy, $row, $column_name);
}

//function tip_render_enemy_id($column_name, $column_val, $row, $data){
//	global $allEnemy;
//
//	if (isset($allEnemy[$column_val])) {
//		$role = $allEnemy[$column_val];
//
//		return mission_enemy_tip_render($role);
//	}
//}

function render_award_town_key($column_name, $column_val, $row, $data){
	global $allTownLock;
	return $allTownLock[$column_val];
}

function editor_award_town_key($column_name, $column_val, $row, $data){
	global $allTownLock;
	return html_get_select($column_name, $allTownLock, $column_val);
}

function render_type($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_type($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_related_town($column_name, $column_val, $row, $data){
	global $allTown;
	return $allTown[$column_val];
}

function editor_related_town($column_name, $column_val, $row, $data){
	global $allTown;
	return html_get_select($column_name, $allTown, $column_val);
}


function render_town_id($column_name, $column_val, $row, $data){
	global $allTown;
	return $allTown[$column_val];
}

function editor_town_id($column_name, $column_val, $row, $data){
	global $allTown;
	return html_get_select($column_name, $allTown, $column_val);
}

function render_town_npc_id($column_name, $column_val, $row, $data){
	global $allNpc;
	return $allNpc[$column_val];
}

function editor_town_npc_id($column_name, $column_val, $row, $data){
	global $allNpc;
	return html_get_select($column_name, $allNpc, $column_val);
}

function render_mission_level_id($column_name, $column_val, $row, $data){
	global $allMissionLevel;
	return $allMissionLevel[$column_val];
	//return html_get_link($allMissionLevel[$column_val], '?p=mission_enemy&m='.$column_val, true);
}

function editor_mission_level_id($column_name, $column_val, $row, $data){
	global $allMissionLevel;
	$missionLevelId = $row['mission_level_id'];
	$missionLevelName = '';
	if ( isset($allMissionLevel[$missionLevelId])) {
		$missionLevelName = $allMissionLevel[$missionLevelId];
	}
	return editor_single_item($missionLevelName, $missionLevelId, "mission_level_id");
	//return html_get_select($column_name, $allMissionLevel, $column_val);
}

function render_drama_mode($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_drama_mode($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function editor_desc($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_desc($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function editor_mission_drama_talk($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_mission_drama_talk($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function sql_where($params) {
	$where = "";
	if (isset($params['m'])) {
		$town_id = $params['m'];
		$where = " where related_town={$town_id}";
	}
	return " {$where} ORDER BY `order`";
}

function jsFunction($params) {
	global $allEnemy, $allItem, $allMissionLevel, $allEnemyIds, $allMissionLevelLock;

	$html = '
		$("select").change( function(){
			$(this).css("color","#0033FF");
});';

$html .= enemy_auto_complete($allEnemy);
$html .= choose_something($allEnemyIds, 'enemy_num');
$html .= choose_something($allMissionLevelLock, 'award_mission_level_lock');
$html .= choose_something($allMissionLevel, 'mission_level_id');
$html .= choose_single_item($allItem, 'award_item1_id');
$html .= choose_single_item($allItem, 'award_item2_id');
$html .= choose_single_item($allItem, 'award_item3_id');
$html .= choose_single_item($allItem, 'award_item4_id');
return $html;
}
?>
