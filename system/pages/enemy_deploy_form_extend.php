<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'0' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(),
	),
);



function enemy_group_form_opera($row, $is_edit = false) {
	global $allEnemyRole;

	$role_exp = 0;
	$coins = 0;
	
	if (!$is_edit) {
		return enemy_group_table_one($allEnemyRole, $row, $row, $row['id'], $role_exp, $coins);
	} else {
		return enemy_group_table_editor($allEnemyRole, $row);
	}
}


function location($params){
	global $db;
	$html = '';

	if (isset($params['m']) && isset($params['m2'])) {
		$sql = "select `id` from `enemy_deploy_form` where parent_id = {$params['m']};";
		$rs = db_query($db, $sql);
		if (count($rs) == 0) {
			// $sql = "INSERT INTO `enemy_deploy_form` set `parent_id`={$params['m']},`battle_type`={$params['m2']};";
			// db_execute($db, $sql);
		}
	}

	switch($params['m2']){
		case 0:
			$sql = "select `mission_level_id`, `order` from `mission_enemy` where id = {$params['m']};";
			$missionEnemy = db_query($db, $sql);
			if (count($missionEnemy) == 0) {
				return $html;
			}

			$sql = "select `mission_id`,`name` from `mission_level` where id = {$missionEnemy[0]['mission_level_id']}";
			$mission_level = db_query($db, $sql);
			if (count($mission_level) == 0) {
				return $html;
			}

			$sql = "select `id`, `name`, `town_id` from mission where id = {$mission_level[0]['mission_id']}";
			$mission = db_query($db, $sql);
			if (count($mission) == 0) {
				return $html;
			}
		
			$sql = "select `id`, `name` from town where id = {$mission[0]['town_id']}";
			$town = db_query($db, $sql);
			if (count($town) == 0) {
				return $html;
			}
			$html .= html_get_link("所有城镇", "?p=town", false)." -> ";
			$html .= html_get_link($town[0]['name'], '?p=mission&m='.$town[0]['id'], false)." -> ";
			$html .= html_get_link($mission[0]['name'], '?p=mission_level&m='.$mission[0]['id'], false)." -> ";
			$html .= html_get_link($mission_level[0]['name'], '?p=mission_enemy&bt='.$params['bt'].'&m='.$missionEnemy[0]['mission_level_id'], false)." -> ";
			$html .= '第'.$missionEnemy[0]['order'].'波怪物阵型';
		break;
	}
	return $html;
}

function sql_where($params) {
	return "where `parent_id`={$params['m']} and `battle_type`={$params['m2']}";
}

function sql_insert($params) {
	return "`parent_id` = {$params['m']},`battle_type`={$params['m2']}";
}

function jsFunction($params) {
	global $allEnemyRole;
	
	$html = '$("select").change( function(){
		$(this).css("color","#0033FF");
	});';

	$html .= enemy_auto_complete($allEnemyRole);

	return $html;
}

$sql = "select *, (select name from skill where id = skill_id) skill_name from `enemy_role` order by `level` asc;";
$allEnemyRole_tmp = db_query($db, $sql);

foreach ($allEnemyRole_tmp as $value)
{
	$allEnemyRole[$value['id']] = $value;
}


?>