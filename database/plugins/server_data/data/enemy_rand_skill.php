<?php

$generate = function($db) {
	$configlist = db_query($db, "select * from enemy_boss_script");


	$code = "package battle\n\n";
	$code.= generateCode_CreateBossTrigger($configlist);
	$code .= generateCode_GetBattlePetCatchRate($db);
	save_code(dirname(__FILE__)."/../../../../server/src/game_server/battle/battle_boss_skill_auto.go", $code);


	$playerBattleSkill = db_query($db, "select `pet_id`, `skill` from battle_pet");
	$configlist = db_query($db, "select * from enemy_role");
	$code = "package battle\n\n";
	$code .= generateCode_CreateEnemySkill($configlist, $playerBattleSkill);
	save_code(dirname(__FILE__)."/../../../../server/src/game_server/battle/battle_enemy_skill_auto.go", $code);
};

function generateCode_GetBattlePetCatchRate($db) {
	$all_rows = db_query($db, "select * from battle_pet_catch_rate order by enemy_id, health_rate asc");
	$all_config = array();
	foreach($all_rows as $row) {
		$all_config[$row['enemy_id']][] = array('health_rate'=> $row['health_rate'], 'rate'=>$row['rate']);
	}


	$code = "//获取灵宠概率\n";
	$code .= "func getBattlePetCatchRate(pet *Fighter) int {\n";
	$code .= "	health_rate := float64(pet.Health) / float64(pet.MaxHealth)\n";
	$code .= "	_ = health_rate \n";
	$code .= "	switch pet.RoleId {\n";
	foreach($all_config as $enemy_id => $config) {
		$code .= generateCode_BattleCatchRateCond($enemy_id, $config);
	}
	$code .="	default: return 0 \n";
	$code .="	}\n";
	$code .= "	return 0\n";
	$code .= "}\n";


	return $code;
}

function generateCode_BattleCatchRateCond($pet_id, &$configs) {
	$code = "	case {$pet_id}:\n";
	foreach($configs as $config ) {
		$code .=<<<T
		if health_rate <= {$config['health_rate']} {
			return {$config['rate']}
		}

T;
	}
	return $code;
}

function generateCode_CreateBossTrigger(&$configlist) {
	$code = "";
	$itemFuncList = "";
	foreach ($configlist as $key => $item) {
		$code .= "	case {$item['boss_id']}:\n";
		$code .= "		return skillTriggerCondition_{$item['boss_id']}\n";
		$itemFuncList .= generateCode_skillTriggerCondition($item);
	}

	return <<<C

func createTriggerSkill(roleId int) funcBossTriggerSkill {
	switch roleId {
{$code}		
	}
	return nil
}

{$itemFuncList}

C;
}

function generateCode_skillTriggerCondition($item) {
	$idx = 0;

	$config = json_decode($item['config']);

	$round_map = array();
	$list_round = array();
	$health_map = array();
	$list_health = array();

	for ($i=0; $i < count($config); $i++) { 
		$rs = $config[$i];
		switch ($rs->condition ) {
		case 1:
			array_push($list_health, $rs->val);
			$health_map[$rs->val] = $rs;
			break;
		case 2:
			array_push($list_round, $rs->val);
			$round_map[$rs->val] = $rs;
			break;
		default:
			continue;
		}
	}

	$code = "";

	sort($list_health, SORT_NUMERIC);
	for ($i=0; $i < count($list_health); $i++) { 
		$val = $list_health[$i];
		$rs = $health_map[$val];
		$idx += 1;

		$code .=<<<T

	if f.Health <= {$val} {
		skill = append(skill, triggerSkill{{$idx}, {$rs->skill_id}, {$rs->power}})
	}

T;
	}

	sort($list_round, SORT_NUMERIC);
	$list_round = array_reverse($list_round);
	for ($i=0; $i < count($list_round); $i++) {
		$val = $list_round[$i];
		$rs = $round_map[$val];
		$idx += 1;
		$code .=<<<T

	if f.battle.GetRounds() >= {$val} - 1 {
		skill = append(skill, triggerSkill{{$idx}, {$rs->skill_id}, {$rs->power}})
	}

T;
	}


	return <<<C

func skillTriggerCondition_{$item['boss_id']}(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}
{$code}	
	return skill
}

C;
}


function generateCode_CreateEnemySkill(&$configlist, &$playerBattlePetSkill) {
	$pet_id = -1;

	$playerPetSkillCode = "";
	foreach($playerBattlePetSkill as $petSkillInfo) {
		$playerPetSkillCode .=  "	case {$petSkillInfo['pet_id']}:\n";
		$playerPetSkillCode .= "		return {$petSkillInfo['skill']}, int(f.NaturalSkillLv)\n";
	}

	$code = "";
	$itemFuncList = "";
	foreach ($configlist as $key => $item) {
		if (empty($item['skill_config'])) {
			continue;	
		} 
		$code .= "	case {$item['id']}:\n";
		$code .= "		return randSkillWithEnemy_{$item['id']}(f)\n";
		$itemFuncList .= generateCode_enemy_skill($item);
	}

	return <<<C

import (
	"math/rand"
)

func createEnemySkill(f *Fighter) (int, /*force for enemy and training level for battle pet*/ int) {
	//玩家的灵宠
	if f.IsBattlePet && f.PlayerId > 0 {
		switch f.RoleId {
{$playerPetSkillCode}
			default:
				return 0,0
		}
	}

	switch f.RoleId {
{$code}		
	}
	return 0, 0
}

{$itemFuncList}

C;
}

function generateCode_enemy_skill($item) {

	$config = json_decode($item['skill_config']);

	$code = "";
	$curr_rand = 0;
	for ($i=0; $i < count($config); $i++) { 
		$rs = $config[$i];
		$rand = $curr_rand+$rs->rand;
		$if_str = ($curr_rand == 0) ? '	if' : ' else if';
		$code .=<<<T
{$if_str} randNum > {$curr_rand} && randNum <= {$rand} {
		skillId = {$rs->skill_id}
		skillForce = {$rs->power}
	}
T;
		$curr_rand = $rand;
	}

	return <<<C

func randSkillWithEnemy_{$item['id']}(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

{$code}	


	return
}

C;
}

?>
