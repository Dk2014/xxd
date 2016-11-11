<?php

$generate = function($db) {
	$battleAutoFile = dirname(__FILE__)."/../../../../server/src/game_server/battle/battle_enemy_passive_skill_auto.go";

	unlink($battleAutoFile);

	$code  = "package battle\n\n";
	//$code .= "import \"math/rand\"\n\n";

	$enemy_passive_skill_config = db_query($db, "select `id`, `config` from `enemy_passive_skill`");

	$code .= "func getEnemyPassiveSkill(enemy *Fighter) ([]*Skill) {\n";
	$code .= "	switch enemy.RoleId {\n";
	foreach ($enemy_passive_skill_config as $config) {
		$code .= "	case ".$config['id'].":\n";

		$skill_config = json_decode($config['config'], true);
		if ($skill_config && is_array($skill_config) && count($skill_config) > 0) {
			$code_skills = join(",", array_map(function($skill_id){ return "role_skill_{$skill_id}";}, $skill_config));
			$code .= "	return []*Skill{{$code_skills}}\n";

		} else {
			$code .="	return nil\n";	
		}
	$code .= "	}\n";
	$code .= "	return nil\n";
	$code .= "}\n\n";
	}
}

?>
