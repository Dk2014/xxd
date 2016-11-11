<?php


$extend_columns = array(
/*   '字段' => 配置 */
	'monster1_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'tip_render'=> array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
		'monster2_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'tip_render'=> array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'monster3_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'tip_render'=> array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'monster4_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'tip_render'=> array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'monster5_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'tip_render'=> array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),

	'is_boss' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array("否","是"),
	),

	'boss_dir' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array('朝左', '朝右'),
	),
);

// 怪物上阵配置一定要指定bt参数(battleType)
if (!isset($_GET['bt'])) {
	exit("怪物上阵配置没有指定战斗类型参数");
}

$g_battleType = $_GET['bt'];


function render_is_boss($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_is_boss($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

$sql = "select *, (select name from skill where id = skill_id) skill_name from `enemy_role` order by `level` asc;";
$allEnemyRole_tmp = db_query($db, $sql);

foreach ($allEnemyRole_tmp as $value)
{
	$allEnemyRole[$value['id']] = $value;
}

function render_monster1_id($column_name, $column_val, $row, $data){
	global $allEnemyRole;

	if ($column_val == 0) {
		return "无";
	} else if(isset($allEnemyRole[$column_val])) {
		return html_get_link($allEnemyRole[$column_val]['name'], '?p=enemy_role&m='.$column_val, true);
	}
}

function editor_monster1_id($column_name, $column_val, $row, $data){
	global $allEnemyRole;

	return enemy_role_editor($allEnemyRole, $row, $column_name);
}

function tip_render_monster1_id($column_name, $column_val, $row, $data){
	global $allEnemyRole;

	if (isset($allEnemyRole[$column_val])) {
		$role = $allEnemyRole[$column_val];

		return mission_enemy_tip_render($role);
	}
}

function range_monster1_id(){
	global $allEnemyRole;
	$tempEnemys = array(0 => '无');
	foreach ($allEnemyRole as $item) {
		$tempEnemys[$item['id']] = $item['name'];
	}
	return $tempEnemys;
}

function render_monster2_id($column_name, $column_val, $row, $data){
	global $allEnemyRole;

	if ($column_val == 0) {
		return "无";
	} else if(isset($allEnemyRole[$column_val])) {
		return html_get_link($allEnemyRole[$column_val]['name'], '?p=enemy_role&m='.$column_val, true);
	}
}

function editor_monster2_id($column_name, $column_val, $row, $data){
	global $allEnemyRole;

	return enemy_role_editor($allEnemyRole, $row, $column_name);
}

function tip_render_monster2_id($column_name, $column_val, $row, $data){
	global $allEnemyRole;

	if (isset($allEnemyRole[$column_val])) {
		$role = $allEnemyRole[$column_val];

		return mission_enemy_tip_render($role);
	}
}

function range_monster2_id(){
	global $allEnemyRole;
	$tempEnemys = array(0 => '无');
	foreach ($allEnemyRole as $item) {
		$tempEnemys[$item['id']] = $item['name'];
	}
	return $tempEnemys;
}

function render_monster3_id($column_name, $column_val, $row, $data){
	global $allEnemyRole;

	if ($column_val == 0) {
		return "无";
	} else if(isset($allEnemyRole[$column_val])) {
		return html_get_link($allEnemyRole[$column_val]['name'], '?p=enemy_role&m='.$column_val, true);
	}
}

function editor_monster3_id($column_name, $column_val, $row, $data){
	global $allEnemyRole;

	return enemy_role_editor($allEnemyRole, $row, $column_name);
}

function tip_render_monster3_id($column_name, $column_val, $row, $data){
	global $allEnemyRole;

	if (isset($allEnemyRole[$column_val])) {
		$role = $allEnemyRole[$column_val];

		return mission_enemy_tip_render($role);
	}
}

function range_monster3_id(){
	global $allEnemyRole;
	$tempEnemys = array(0 => '无');
	foreach ($allEnemyRole as $item) {
		$tempEnemys[$item['id']] = $item['name'];
	}
	return $tempEnemys;
}

function render_monster4_id($column_name, $column_val, $row, $data){
	global $allEnemyRole;

	if ($column_val == 0) {
		return "无";
	} else if(isset($allEnemyRole[$column_val])) {
		return html_get_link($allEnemyRole[$column_val]['name'], '?p=enemy_role&m='.$column_val, true);
	}
}

function editor_monster4_id($column_name, $column_val, $row, $data){
	global $allEnemyRole;

	return enemy_role_editor($allEnemyRole, $row, $column_name);
}

function tip_render_monster4_id($column_name, $column_val, $row, $data){
	global $allEnemyRole;

	if (isset($allEnemyRole[$column_val])) {
		$role = $allEnemyRole[$column_val];

		return mission_enemy_tip_render($role);
	}
}

function range_monster4_id(){
	global $allEnemyRole;
	$tempEnemys = array(0 => '无');
	foreach ($allEnemyRole as $item) {
		$tempEnemys[$item['id']] = $item['name'];
	}
	return $tempEnemys;
}

function render_monster5_id($column_name, $column_val, $row, $data){
	global $allEnemyRole;

	if ($column_val == 0) {
		return "无";
	} else if(isset($allEnemyRole[$column_val])) {
		return html_get_link($allEnemyRole[$column_val]['name'], '?p=enemy_role&m='.$column_val, true);
	}
}

function editor_monster5_id($column_name, $column_val, $row, $data){
	global $allEnemyRole;

	return enemy_role_editor($allEnemyRole, $row, $column_name);
}

function tip_render_monster5_id($column_name, $column_val, $row, $data){
	global $allEnemyRole;

	if (isset($allEnemyRole[$column_val])) {
		$role = $allEnemyRole[$column_val];

		return mission_enemy_tip_render($role);
	}
}

function range_monster5_id(){
	global $allEnemyRole;
	$tempEnemys = array(0 => '无');
	foreach ($allEnemyRole as $item) {
		$tempEnemys[$item['id']] = $item['name'];
	}
	return $tempEnemys;
}

function editor_boss_dir($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_boss_dir($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function location($params){
	//var_dump($params);
	//die();
	global $db;
	$html = '';
	$battle_type = $params['bt'];
	switch ($battle_type) {
	case 8: //难度关卡
		$sql = "select * from `mission_level` where id = ".$_GET['m'].";";
		$levels = db_query($db, $sql);
		$html .= html_get_link("难度关卡", "?p=hard_level", false) . ' -> ';
		$html .= html_get_link($levels[0]['name'], "?p=hard_level_config&m=".$levels[0]['parent_id'], false);
		return $html;

	default:

		$sql = "select * from `mission_level` where id = ".$_GET['m'].";";
		$levels = db_query($db, $sql);

		$sql = "select town_id, name from `mission` where id = ".$levels[0]["mission_id"].";";

		$tid = db_query($db, $sql);
		if (count($tid)>0 && $tid[0]['town_id']!=0) {
			$html .= html_get_link("副本", '?p=mission&m='.$tid[0]['town_id'], false).' -> ';
			$html .= html_get_link($tid[0]['name'], '?p=mission_level&m='.$levels[0]['mission_id'], false).' -> '.$levels[0]['name'];	
		}

		return $html;
	}
}

function sql_where($params) {
	$shaded_mission_id = (isset($params['m2']) ? $params['m2'] : 0);
	return "where `mission_level_id`={$params['m']} AND `shaded_mission_id`={$shaded_mission_id} ORDER BY `order`";
}

function sql_insert($params) {
	$shaded_mission_id = (isset($params['m2']) ? $params['m2'] : 0);
	return "`mission_level_id` = {$params['m']}, `shaded_mission_id`={$shaded_mission_id}";
}

function jsFunction($params) {
	global $allEnemyRole;

	$html = '
		$("select").change( function(){
			$(this).css("color","#0033FF");
});';

$html .= enemy_auto_complete($allEnemyRole);

return $html;
}

function del_check($rowid) {
	global $db;
	$sql = "select * from enemy_deploy_form where battle_type = 0 and parent_id = {$rowid}";
	$datas = db_query($db, $sql);
	if (count($datas) > 0) {
		return false;
	}
	return true;
}

function mission_enemy_opera($row) {
	global $allEnemyRole, $db, $g_battleType;

	$html = '';
	$html .= '<table border="0" width="200"><tr><td width="180">';
	$html .= '<a href="?p=enemy_deploy_form&m='.$row['id'].'&m2='.$g_battleType.'&bt='.$_GET['bt'].'" target="_blank">怪物阵型</a>';
	$html .= ' | <a href="?p=mission_talk&m='.$row['id'].'" target="_blank">对话</a>';
	$html .= ' | <a href="?p=level_battle_pet&m='.$row['id'].'&bt='.$_GET['bt'].'" target="_blank">灵宠</a>';
	$html .= enemy_group_table($allEnemyRole, $row, "select * from enemy_deploy_form where parent_id = ".$row['id']." and `battle_type`= ". $g_battleType);
	$html .= '</td>';
	$html .= '</tr></table>';

	return $html;
}
?>
