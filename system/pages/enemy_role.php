<?php

function enemy_role_skill_id_render($row, $key) {
	global $allEnemySkill;
	$html = '';
	if (isset($allEnemySkill[$row[$key]])){
		$html.= '<a href="?p=skill_enemy">'.$allEnemySkill[$row[$key]].'</a>';
	}
	return $html;
}

function enemy_role_skill_id_editor($row, $key){
	global $allEnemySkills;
	$code  = '<select name="'.$key.'[]" >';
	$code .= '<option value="-1" selected="selected">无</option>';
	foreach ($allEnemySkills as $value){
		if ($value['id'] == $row[$key]) {
			$code .= '<option value="'.$value['id'].'"  selected="selected">'.$value['name'].'</option>';
		} else {
			$code .= '<option value="'.$value['id'].'" >'.$value['name'].'</option>';
		}
	}
	$code .= '</select>';
	return $code;
}

function enemy_role_skill_id_render1($row) {
	return enemy_role_skill_id_render($row, 'skill_id');
}

function enemy_role_skill_id_editor1($row) {
	return enemy_role_skill_id_editor($row, 'skill_id');
}

function enemy_role_skill_id_render2($row) {
	return enemy_role_skill_id_render($row, 'skill2_id');
}

function enemy_role_skill_id_editor2($row) {
	return enemy_role_skill_id_editor($row, 'skill2_id');
}

function addr_extra($params) {
	$url = "";
	if (isset($params['m'])) 
		$url .= "&m={$params['m']}";
	if (isset($params['town_id']))
		$url .= "&town_id={$params['town_id']}";
	if (isset($params['type'])) 
		$url .= "&type={$params['type']}";
	if (isset($params['challenge_mission_set'])) 
		$url .= "&challenge_mission_set={$params['challenge_mission_set']}";
	
	return $url;
}

function get_call_enemys(){
	global $db;
	$sql = "select `config` from `skill` where `config` like '%\"EnemyCalls\": [{%' or `config` like '%\"EnemyCalls\":[{%';";
	$res = db_query($db, $sql);
	$enemyIds = array();
	foreach ($res as $value) {
		$str = $value['config'];
		$skill_config = json_decode($str);
		foreach ($skill_config->EnemyCalls as $enemy) {
			$enemyIds[$enemy->Enemys] = $enemy->Enemys;
		}
	}
	return array_keys($enemyIds);
}

function sql_where($params) {
	global $db;

	$where = array();

	$have_params = false;

	if (isset($params['town_id'])) {
		$enemyIds = array();

		$sql = "select * from (select ml.name,ml.id as mlid from mission m left join mission_level ml on ml.mission_id=m.id and m.town_id={$params['town_id']}) m1 left join
		 (select m.mission_level_id,m.monster1_id,m.monster2_id,m.monster3_id,m.monster4_id,m.monster5_id,
		 	e.* from mission_enemy m left join enemy_deploy_form e on m.id=e.parent_id and e.battle_type!=3) m2 on m1.mlid = m2.mission_level_id;";
	
		$res = db_query($db, $sql);
		foreach ($res as $v) {
			for ($i=1; $i < 6; $i++) {
				$field = 'monster'.$i.'_id';
				if ($v[$field] > 0 ) {
					array_push($enemyIds, (int)$v[$field]);
				}
			}

			for ($i=1; $i < 16; $i++) {
				$field = 'pos'.$i;
				if ($v[$field] > 0 ) {
					array_push($enemyIds, (int)$v[$field]);
				}
			}			
		}

		// array_unique($enemyIds);

		if (count($enemyIds) >0){
			$where[] = "`id` in (".implode(',', $enemyIds).")";
		}

		$have_params = true;
	}
	else if (isset($params['m'])) {
		$have_params = true;
		return " WHERE id = ".$params['m'];
	}

	if (isset($params['type'])) {
		switch ($params['type']) {
			case 'notuse':

				$enemyIds = array();

				$sql = "SELECT * FROM enemy_deploy_form;";
				$res = db_query($db, $sql);
				foreach ($res as $v) {
					for ($i=1; $i < 16; $i++) { 
						$field = 'pos'.$i;
						if ( $v[$field] > 0 ) {
							array_push($enemyIds, (int)$v[$field]);
						}
					}
				}

				$sql = "SELECT * FROM mission_enemy;";
				$res = db_query($db, $sql);
				foreach ($res as $v) {
					for ($i=1; $i < 6; $i++) { 
						$field = "monster{$i}_id";
						if ($v[$field] > 0 ) {
							array_push($enemyIds, (int)$v[$field]);
						}
					}
				}


				$sql = "select pet_id from battle_pet;";
				$rows = db_query($db, $sql);
				foreach($rows as $row) {
					if($row['pet_id']) 
						array_push($enemyIds, $row['pet_id']);
				}

				array_unique($enemyIds);
				$enemyIds = array_merge($enemyIds, get_call_enemys());

				if (count($enemyIds) >0){
					$where[] = " `id` not in (".implode(',', $enemyIds).") ";
				}

				break;
			case 'multi_level':
				$col_pattern = '/pos[1-9][0-9]*/';
				$enemyIds = array();
				//对人战场和怪物的关联在 战斗场地部署这里。 battle_type=3 则是多人关卡
				$sql = "SELECT * FROM enemy_deploy_form where battle_type=3;";
				$deploys = db_query($db, $sql);
				foreach($deploys as $deploy) {
					foreach($deploy as $k => $v) {
						//选出 pos1 pos2 ... pos15 这些列，v=0 是没有设置怪物
						if(!preg_match($col_pattern, $k) || $v == "0") {
							continue;
						}
						if(!in_array($v, $enemyIds)) {
							array_push($enemyIds, $v);
						}
					}
				}
				if (count($enemyIds) >0){
					$where[] = " `id` in (".implode(',', $enemyIds).") ";
				}
				break;
			case 'rainbow_level':

				$have_params = true;
				$col_pattern = '/monster[1-9][0-9]*_id/';
				$enemyIds = array();
				$sql = "select m.id as id from mission_level as m join hard_level as r on m.parent_id = r.id where m.parent_type = {$params['m2']};";
				$rows = db_query($db, $sql);
				$resource_level_ids = array();
				foreach($rows as $row) {
					$resource_level_ids[] = $row['id'];
				}
				if (count($resource_level_ids) < 1) {
					break;
				}
				$sql = "select * from mission_enemy where `mission_level_id` in (".implode(',', $resource_level_ids).") ";
				$rows = db_query($db, $sql);
				foreach($rows as $enemy) {
					foreach($enemy as $k => $v) {
						if(!preg_match($col_pattern, $k) || $v == "0") {
							continue;
						}
						if(!in_array($v, $enemyIds)) {
							array_push($enemyIds, $v);
						}
					}
				}

				$sql = "select * from enemy_deploy_form ef where ef.battle_type = {$params['m2']}";
				$rows = db_query($db, $sql);
				foreach($rows as $enemy) {
					foreach($enemy as $k => $v) {
						if(!preg_match('/pos[1-9]/', $k) || $v == "0") {
							continue;
						}
						if(!in_array($v, $enemyIds)) {
							array_push($enemyIds, $v);
						}
					}
				}
				if (count($enemyIds) >0){
					$where[] = " `id` in (".implode(',', $enemyIds).") ";
				}


				break;
			case 'hard_level':
				$have_params = true;
				$col_pattern = '/monster[1-9][0-9]*_id/';
				$enemyIds = array();
				$sql = "select m.id as id from mission_level as m join hard_level as r on m.parent_id = r.id where m.parent_type = {$params['m2']};";
				$rows = db_query($db, $sql);
				$resource_level_ids = array();
				foreach($rows as $row) {
					$resource_level_ids[] = $row['id'];
				}
				if (count($resource_level_ids) < 1) {
					break;
				}
				$sql = "select * from mission_enemy where `mission_level_id` in (".implode(',', $resource_level_ids).") ";
				$rows = db_query($db, $sql);
				foreach($rows as $enemy) {
					foreach($enemy as $k => $v) {
						if(!preg_match($col_pattern, $k) || $v == "0") {
							continue;
						}
						if(!in_array($v, $enemyIds)) {
							array_push($enemyIds, $v);
						}
					}
				}

				$sql = "select * from enemy_deploy_form ef where ef.battle_type = {$params['m2']}";
				$rows = db_query($db, $sql);
				foreach($rows as $enemy) {
					foreach($enemy as $k => $v) {
						if(!preg_match('/pos[1-9]/', $k) || $v == "0") {
							continue;
						}
						if(!in_array($v, $enemyIds)) {
							array_push($enemyIds, $v);
						}
					}
				}
				if (count($enemyIds) >0){
					$where[] = " `id` in (".implode(',', $enemyIds).") ";
				}
				break;
			case 'extend_level':
				$have_params = true;
				$col_pattern = '/monster[1-9][0-9]*_id/';
				$enemyIds = array();
				//查找处所有资源关卡
				$sql = "select m.id as id from mission_level as m join extend_level as r on m.parent_id = r.id where m.parent_type = {$params['m2']};";
				$rows = db_query($db, $sql);
				$resource_level_ids = array();
				foreach($rows as $row) {
					$resource_level_ids[] = $row['id'];
				}
				if (count($resource_level_ids) < 1) {
					break;
				}
				$sql = "select * from mission_enemy where `mission_level_id` in (".implode(',', $resource_level_ids).") ";
				$rows = db_query($db, $sql);
				foreach($rows as $enemy) {
					foreach($enemy as $k => $v) {
						if(!preg_match($col_pattern, $k) || $v == "0") {
							continue;
						}
						if(!in_array($v, $enemyIds)) {
							array_push($enemyIds, $v);
						}
					}
				}

				$sql = "select * from enemy_deploy_form ef where ef.battle_type = {$params['m2']}";
				$rows = db_query($db, $sql);
				foreach($rows as $enemy) {
					foreach($enemy as $k => $v) {
						if(!preg_match('/pos[1-9]/', $k) || $v == "0") {
							continue;
						}
						if(!in_array($v, $enemyIds)) {
							array_push($enemyIds, $v);
						}
					}
				}
				if (count($enemyIds) >0){
					$where[] = " `id` in (".implode(',', $enemyIds).") ";
				}
				break;
			case 'battle_pet':
				$sql = "select pet_id from battle_pet;";
				$enemyIds = array();
				$rows = db_query($db, $sql);
				foreach($rows as $row) {
					if($row['pet_id']) 
						array_push($enemyIds, $row['pet_id']);
				}
				if (count($enemyIds) >0){
					$where[] = " `id` in (".implode(',', $enemyIds).") ";
				}
				array_unique($enemyIds);
				break;
			case 'pet_virtual_env':
				$have_params = true;
				$col_pattern = '/monster[1-9][0-9]*_id/';
				$enemyIds = array();
				//查找处所有灵宠幻境关卡
				$sql = "select m.id as id  from mission_level as m join pve_level as r on m.parent_id = r.id where m.parent_type=13;";
				$rows = db_query($db, $sql);
				$pve_level_ids = array();
				foreach($rows as $row) {
					$pve_level_ids[] = $row['id'];
				}
				if (count($pve_level_ids) < 1) {
					break;
				}
				$sql = "select * from mission_enemy where `mission_level_id` in (".implode(',', $pve_level_ids).") ";
				$rows = db_query($db, $sql);
				foreach($rows as $enemy) {
					foreach($enemy as $k => $v) {
						if(!preg_match($col_pattern, $k) || $v == "0") {
							continue;
						}
						if(!in_array($v, $enemyIds)) {
							array_push($enemyIds, $v);
						}
					}
				}

				$sql = "select * from enemy_deploy_form ef where ef.battle_type = 13";
				$rows = db_query($db, $sql);
				foreach($rows as $enemy) {
					foreach($enemy as $k => $v) {
						if(!preg_match('/pos[1-9]/', $k) || $v == "0") {
							continue;
						}
						if(!in_array($v, $enemyIds)) {
							array_push($enemyIds, $v);
						}
					}
				}
				if (count($enemyIds) >0){
					$where[] = " `id` in (".implode(',', $enemyIds).") ";
				}
				break;
			case 'fate_box':
				$have_params = true;
				$col_pattern = '/monster[1-9][0-9]*_id/';
				$enemyIds = array();
				$sql = "select m.id as id from mission_level as m where m.parent_type = {$params['m2']};";
				$rows = db_query($db, $sql);
				$resource_level_ids = array();
				foreach($rows as $row) {
					$resource_level_ids[] = $row['id'];
				}
				if (count($resource_level_ids) < 1) {
					break;
				}
				$sql = "select * from mission_enemy where `mission_level_id` in (".implode(',', $resource_level_ids).") ";
				$rows = db_query($db, $sql);
				foreach($rows as $enemy) {
					foreach($enemy as $k => $v) {
						if(!preg_match($col_pattern, $k) || $v == "0") {
							continue;
						}
						if(!in_array($v, $enemyIds)) {
							array_push($enemyIds, $v);
						}
					}
				}

				$sql = "select * from enemy_deploy_form ef where ef.battle_type = {$params['m2']}";
				$rows = db_query($db, $sql);
				foreach($rows as $enemy) {
					foreach($enemy as $k => $v) {
						if(!preg_match('/pos[1-9]/', $k) || $v == "0") {
							continue;
						}
						if(!in_array($v, $enemyIds)) {
							array_push($enemyIds, $v);
						}
					}
				}
				if (count($enemyIds) >0){
					$where[] = " `id` in (".implode(',', $enemyIds).") ";
				}

				break;
			case 'call_enemys':
			 $enemyIds = get_call_enemys();
			 if (count($enemyIds) >0){
					$where[] = " `id` in (".implode(',', $enemyIds).") ";
				}

				break;
			default:
				break;
		}
	}

	$where_sql = '';

	if (count($where)) 
		$where_sql = "WHERE ". implode(" AND ", $where);
	else if((!$have_params && isset($params['order']) && $params['order']!='level' && !isset($params['type'])))
		return $where_sql = " ORDER BY `id` DESC";
	else if(!isset($params['order']) && !isset($params['type']))
		return $where_sql = " ORDER BY `id` DESC";
	else
		$where_sql = " WHERE 1=1";

		$where_sql .= " ORDER BY `level`";
	return $where_sql;
}


//绝招
$sql = "select `id`,`name` from skill where `type` = 5;";
$allEnemySkills = db_query($db, $sql);
foreach ($allEnemySkills as $value){
	$allEnemySkill[$value['id']] = $value['name'];
}



$i = 0;
$linksTown[$i]['text'] = '最新录入';
$linksTown[$i]['id'] = "enemy_role";



$linksTown[++$i]['text'] = "未使用";
$linksTown[$i]['id'] = "enemy_role&type=notuse";

$linksTown[++$i]['text'] = "难度关卡";
$linksTown[$i]['id'] = "enemy_role&type=hard_level&m2=8";

//多人关卡
$linksTown[++$i]['text'] = "多人关卡";
$linksTown[$i]['id'] = "enemy_role&type=multi_level";

//极限关卡
$linksTown[++$i]['text'] = "彩虹关卡";
$linksTown[$i]['id'] = "enemy_role&type=rainbow_level&m2=12";

// 命锁关卡
$linksTown[++$i]['text'] = "命锁关卡";
$linksTown[$i]['id'] = "enemy_role&type=fate_box&m2=14";

//资源关卡
$linksTown[++$i]['text'] = "资源关卡";
$linksTown[$i]['id'] = "enemy_role&type=extend_level&m2=1";

$linksTown[++$i]['text'] = "伙伴关卡";
$linksTown[$i]['id'] = "enemy_role&type=extend_level&m2=9";


$linksTown[++$i]['text'] = "灵宠关卡";
$linksTown[$i]['id'] = "enemy_role&type=extend_level&m2=10";


$linksTown[++$i]['text'] = "魂侍关卡";
$linksTown[$i]['id'] = "enemy_role&type=extend_level&m2=11";


//灵宠
$linksTown[++$i]['text'] = "灵宠";
$linksTown[$i]['id'] = "enemy_role&type=battle_pet";

//灵宠幻境
$linksTown[++$i]['text'] = "灵宠幻境";
$linksTown[$i]['id'] = "enemy_role&type=pet_virtual_env";

// 怪物召唤
$linksTown[++$i]['text'] = "怪物召唤";
$linksTown[$i]['id'] = "enemy_role&type=call_enemys";

//城镇
$sql = "select `id`,`name` from `town`";
$allTowns = db_query($db, $sql);


$i++;
foreach ($allTowns as $value){
	$linksTown[$i]['text'] = $value['name'];
	$linksTown[$i]['id'] = "enemy_role&town_id=".$value['id'];
	$i++;
}

$linksTown[++$i]['text'] = '按等级分怪物';
$linksTown[$i]['id'] = "enemy_role&order=level";

$pconfig = array(
	'title'   => '敌人角色数据',
	'table'   => 'enemy_role',
	'links'   => $linksTown,
	'columns' => array(
		array('name' => 'name', 'text' => '角色名称', 'width' => '90px'),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '120px'),
		array('name' => 'prop', 'text' => '种族', 'width' => '120px'),
		array('name' => 'is_boss', 'text' => 'Boss', 'width' => '40px'),
		array('name' => 'show_shader', 'text' => '阴影附身', 'width' => '40px'),
		array('name' => 'jump_attack', 'text' => '跳跃攻击', 'width' => '60px'),
		array('name' => 'body_size', 'text' => '体型', 'width' => '40px'),
		array('name' => 'scale_size', 'text' => '缩放比%', 'width' => '40px'),
		array('name' => 'level', 'text' => '等级', 'width' => '30px'),
		array('name' => 'health', 'text' => '生命', 'width' => '60px'),
		array('name' => 'attack', 'text' => '攻击', 'width' => '60px'),
		array('name' => 'defence', 'text' => '防御', 'width' => '60px'),
		array('name' => 'cultivation', 'text' => '内力', 'width' => '60px'),
		array('name' => 'speed', 'text' => '速度', 'width' => '60px'),
		array('name' => 'sunder_max_value', 'text' => '护甲上限','width' => '60px', 'default' => '100'),
		array('name' => 'sunder_min_hurt_rate', 'text' => '破甲前百分比','width' => '60px', 'default' => '100'),
		array('name' => 'sunder_end_hurt_rate', 'text' => '破甲后百分比','width' => '60px', 'default' => '150'),
		array('name' => 'sunder_end_defend_rate', 'text' => '破甲后减防百分比','width' => '60px', 'default' => '20'),
		array('name' => 'sunder_attack', 'text' => '攻击破甲值','width' => '60px', 'default' => '5'),

		// array('name' => 'skill_id', 'text' => '绝招1种类','width' => '90px',
		// 	'render' => "enemy_role_skill_id_render1",
		// 	'editor' => "enemy_role_skill_id_editor1",
		// 	),
		// array('name' => 'skill_force', 'text' => '绝招1威力','width' => '60px'),

		// array('name' => 'skill2_id', 'text' => '绝招2种类','width' => '90px',
		// 	'render' => "enemy_role_skill_id_render2",
		// 	'editor' => "enemy_role_skill_id_editor2",
		// 	),
		// array('name' => 'skill2_force', 'text' => '绝招2威力','width' => '60px'),

		// array('name' => 'release_num', 'text' => '释放次数', 'width' => '60px', 'default' => '1'),
		// array('name' => 'recover_round_num', 'text' => '恢复回合数', 'width' => '60px', 'default' => '2'),
		// array('name' => 'common_attack_num', 'text' => '入场普通攻击次数', 'width' => '60px', 'default' => '0'),
		// array('name' => 'skill_wait', 'text' => '绝招蓄力回合','width' => '60px', 'default' => '0'),
		
		array('name' => 'critial', 'text' => '暴击%','width' => '60px'),
		array('name' => 'block', 'text' => '格挡%','width' => '60px'),
		array('name' => 'hit', 'text' => '命中%','width' => '60px'),
		array('name' => 'dodge', 'text' => '闪避%','width' => '60px'),
		array('name' => 'critial_hurt', 'text' => '必杀%','width' => '60px'),
		array('name' => 'toughness', 'text' => '韧性%','width' => '60px'),
		array('name' => 'destroy', 'text' => '破击%','width' => '60px'),
		array('name' => 'sleep',  'text' => '睡眠抗性%', 'width' => '60px'),
	    array('name' => 'dizziness',  'text' => '眩晕抗性%', 'width' => '60px'),
	    array('name' => 'random',  'text' => '混乱抗性%', 'width' => '60px'),
	    array('name' => 'disable_skill',  'text' => '封魔抗性%', 'width' => '60px'),
	    array('name' => 'poisoning',  'text' => '中毒抗性%', 'width' => '60px'),
	),
	'where'   => 'sql_where',
	'addr'     => 'addr_extra',
	'not_delete' => true,
	'foot' => 'foot',
	'opera' => array(
    array('type' => 'link', 'text' => 'boss脚本配置', 'render' => 'op_render'),
    array('type' => 'link', 'text' => '', 'render' => 'monster_property_add'),
	),
	'batch_update_key' => 'id',
	'batch_update_password' => '789',
);
if(count($_GET)>1 && !isset($_GET['page']) && !isset($_GET['order'])){
	$pconfig['pagesize'] = 100000;
}
?>
