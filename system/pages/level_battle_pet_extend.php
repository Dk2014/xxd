<?php
$extend_columns = array(
/*   '字段' => 配置 */
	'battle_pet_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(),
		'tip_render'=> array('params' => array()),
	),
);

$allEnemy = get_all_enemy();

$battlePet = array();
$battlePetName = array();
$sql = "select p.id,p.pet_id,e.name from `battle_pet` p left join `enemy_role` e on e.id = p.pet_id where e.id > 0";
$rows = db_query($db, $sql);
foreach ($rows as $key => $value) {
	$battlePet[$value['id']] = $value['pet_id'];
	$battlePetName[$value['id']]=$value['name'];
}

function render_battle_pet_id($column_name, $column_val, $row, $data){
	global $allEnemy,$battlePet;

	$petId = ($column_val>0)?$battlePet[$column_val]:0;

	if ($petId == 0) {
		return "无";
	} else if(isset($allEnemy[$petId])) {
		return html_get_link($allEnemy[$petId]['name'], '?p=enemy_role&m='.$petId, true);
	}
}

function editor_battle_pet_id($column_name, $column_val, $row, $data){
	global $allEnemy, $battlePetName;

	return html_get_select($column_name, $battlePetName, $column_val);
}

function tip_render_battle_pet_id($column_name, $column_val, $row, $data){
	global $allEnemy,$battlePet;

	$petId = ($column_val>0)?$battlePet[$column_val]:0;

	if (isset($allEnemy[$petId])) {
		$role = $allEnemy[$petId];

		return mission_enemy_tip_render($role);
	}
}

function location($params){
	global $db;
	$html = '';
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
	$html .= '第'.$missionEnemy[0]['order'].'波怪物组';
	
	return $html;
}

function sql_where($params) {
	return "where `mission_enemy_id`={$params['m']}";
}

function sql_insert($params) {
	return "`mission_enemy_id` = {$params['m']}";
}

?>