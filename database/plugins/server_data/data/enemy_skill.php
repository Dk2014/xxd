<?php

$generate = function($db) {
	$battleAutoFile = dirname(__FILE__)."/../../../../server/src/game_server/battle/battle_auto.go";

	unlink($battleAutoFile);

	$code  = "package battle\n\n";
	$code .= "import \"math/rand\"\n\n";

	$role_skills = db_query($db, "select id, role_id, type, child_type, name, config from skill where type not in (3,4,6)");

	$code .= "func createRoleSkill(id int) *Skill {\n";
	$code .= "	switch id {\n";
	foreach ($role_skills as $skill) {
		$code .= "	case ".$skill['id'].": return role_skill_".$skill['id']." // ".$skill['name']."\n";
	}
	$code .= "	}\n";
	$code .= "	return nil\n";
	$code .= "}\n\n";

	$totem_skills = db_query($db, "select * from  skill where type=9 order by role_id");
	$code .= "func createTotemSkill(id int16, level int16) (*Skill, *SkillInfo) {\n";
	$code .= "	switch id {\n";
	foreach ($totem_skills as $skill) {
		$code .= "		case ".$skill['id'].": return role_skill_".$skill['id'];
		$code .= ", &SkillInfo{SkillTrnLv: level-1} ";
		$code .= " // ".$skill['name']."\n";
	}
	$code .= "	}\n";
	$code .= "	return nil, nil\n";
	$code .= "}\n\n";

	$ghost_skills = db_query($db, "select id, role_id,`type`, child_type, name, config from skill where type in(7,8) order by role_id, `required_level`,type");
	if (count($ghost_skills) > 0) {
		$ghostSkillTable = array();
		foreach ($ghost_skills as $skill) {
			$ghostSkillTable[$skill['role_id']][$skill['type']] = $skill;
		}

		$code .= "func createGhostSkill(fighter *Fighter) (skill *Skill, skillInfo *SkillInfo) {\n";
		$code .= "	id := int(fighter.useGhostSkillId)\n";
		$code .= "	switch id {\n";
		foreach ($ghost_skills as $skill) {
			if ($skill['type'] != 7) {
				continue;
			}

			$s = $ghostSkillTable[$skill['role_id']];
			$skill2 = null;
			if (isset($s[8])) {
				$skill2 = $s[8];
			}

			$code .= "	case ".$skill['id'].":	// {$skill['name']}\n";
			$code .= "		skill = role_skill_{$skill['id']}\n";
			if(isset($skill2)) {
				$code .= "		skillInfo = &SkillInfo{SkillId:{$skill['id']}, SkillId2:{$skill2['id']}}\n";
			} else {
				$code .= "		skillInfo = &SkillInfo{SkillId:{$skill['id']},}\n";
			}
		}
		$code .= "	}\n";
		$code .= "	ghost := fighter.GetMainGhost()\n";
		$code .= "	if ghost != nil {\n";
		$code .= "		skillInfo.SkillTrnLv = ghost.GhostSkillLv\n";
		$code .= "		skillInfo.SkillTrnLv2 = ghost.GhostSkillLv\n";
		$code .= "	}\n";

		$code .= "	return \n";
		$code .= "}\n\n";
	}

	foreach ($role_skills as $skill) {
		$code .= "// ".$skill['name']."\n";
		$code .= "var role_skill_{$skill['id']} = &Skill{\n";
		$code .= gen_skill_code("role", $skill, true);
		$code .= "}\n\n";
	}


	save_code($battleAutoFile, $code);
};

$buff_id = 0;

// 去除undefined属性
function fix_skill_config_undefined(&$config) {
	if (!isset($config))
		return;

	if (!isset($config->TrnlvRate))
		$config->TrnlvRate = 0;
}

// 去除undefined属性
function fix_skill_buff_config_undefined(&$buff) {
	if (!isset($buff))
		return;

	if (!isset($buff->TrnlvRate))
		$buff->TrnlvRate = 0;

	if (!isset($buff->RequireSkillLevel))
		$buff->RequireSkillLevel = 0;

	if (!isset($buff->BuffUncleanable))
		$buff->BuffUncleanable= false;
}

function gen_skill_code($skill_type, $skill, $for_role) {
	global $buff_id;

	$code = "";

	$config = json_decode($skill['config']);

	fix_skill_config_undefined($config);

	$code .= "	SkillId: {$skill['id']},\n";
	$code .= "	ChildType: {$skill['child_type']},\n";
	if(is_ghost_skill($skill['type'])) {
		$code .= "	IsGhostSkill: true,\n";
	}

	// 角色绝招专有
	if ($for_role) {
		if (isset($config->DecPower) && $config->DecPower != 0)
			$code .= "	DecPower: {$config->DecPower},\n";

		if (isset($config->IncPower) && $config->IncPower != 0)
			$code .= "	IncPower: {$config->IncPower},\n";

		if (isset($config->AppendSpecialType) && $config->AppendSpecialType > 0) {
			$typeVal = get_AppendSpecialType_const($config->AppendSpecialType);
			$code .= "	AppendSpecialType: {$typeVal},\n";
			$code .= "	AppendSpecialValue: {$config->AppendSpecialValue},\n";
		}
		//$code .= "	IsRoleSkill: true,\n";
	}

	// 魂侍绝招专有
	if (!$for_role) {
		// 伤害增加值
		if (isset($config->HurtAdd) && $config->HurtAdd > 0)
			$code .= "	HurtAdd: {$config->HurtAdd},\n";
		// 伤害增加百分比
		if (isset($config->HurtAddRate) && $config->HurtAddRate > 0)
			$code .= "	HurtAddRate: ".($config->HurtAddRate * 0.01).",\n";
		// 治疗量加值
		if (isset($config->CureAdd) && $config->CureAdd)
			$code .= "	CureAdd: {$config->CureAdd},\n";
		// 治疗增加百分比          
		if (isset($config->CureAddRate) && $config->CureAddRate)
			$code .= "	CureAddRate: ".($config->CureAddRate * 0.01).",\n";
		// 覆盖buff
		if (isset($config->GhostOverrideTargetBuff) && $config->GhostOverrideTargetBuff)
			$code .= " OverrideTargetBuff: true,\n";
		// 覆盖buff
		if (isset($config->GhostOverrideSelfBuff) && $config->GhostOverrideSelfBuff)
			$code .= " OverrideSelfBuff: true,\n";
		// 覆盖buff
		if (isset($config->GhostOverrideBuddyBuff) && $config->GhostOverrideBuddyBuff)
			$code .= " OverrideBuddyBuff: true,\n";
		$code .= "	IsGhostSkill: true,\n";
	}

	// 固定值
	if ($for_role) {
		$fixedVal = (isset($config->DefaultAttack)) ? $config->DefaultAttack : 0;
		$code .= "	FixedValue: {$fixedVal},\n";
	}

	// 破甲
	if (isset($config->SunderAttack) && $config->SunderAttack != 0)
		$code .= "	SunderAttack: {$config->SunderAttack},\n";
	// 暴击
	if (isset($config->Critial) && $config->Critial > 0)
		$code .= "	Critial: {$config->Critial},\n";
	// 无视防御
	if (isset($config->ReduceDefend) && $config->ReduceDefend > 0)
		$code .= "	ReduceDefend: ".($config->ReduceDefend / 100.0).",\n";
	// 百分百命中
	if (isset($config->MustHit) && $config->MustHit)
		$code .= "	NotMiss: true,\n";

	// 攻击次数
	if (isset($config->AttackNum) && $config->AttackNum != 0 && $config->AttackNum != 1)
		$code .= "	AttackNum: {$config->AttackNum},\n";

	if (isset($config->TargetMode)) {
		$findTarget = "";
		$findTargetMode = "";
		$ghostAddRate = "";
		$attackRangeRatio = "";

		switch ($config->TargetMode) {
			// 单体
		case 0:
			$findTarget = "findOneTarget";
			$findTargetMode = "SKILL_ATTACK_MODE_SINGLE";
			$ghostAddRate = "5";
			$attackRangeRatio = "1";
			break;
			// 全体
		case 1:
			$findTarget = "findAllTargets";
			$findTargetMode = "SKILL_ATTACK_MODE_AOE";
			$ghostAddRate = "1";
			$attackRangeRatio = "0.6";
			break;
			// 横向
		case 2:
			$findTarget = "findRowTarget";
			$findTargetMode = "SKILL_ATTACK_MODE_AOE";
			$ghostAddRate = "2";
			$attackRangeRatio = "0.75";
			break;
			// 纵向
		case 3:
			$findTarget = "findColTarget";
			$findTargetMode = "SKILL_ATTACK_MODE_AOE";
			$ghostAddRate = "2";
			$attackRangeRatio = "0.75";
			break;
			// 后排横向
		case 5:
			$findTarget = "findLastRowTargets";
			$findTargetMode = "SKILL_ATTACK_MODE_AOE";
			$ghostAddRate = "2";
			$attackRangeRatio = "0.75";
			break;
			// 血量最少
		case 6:
			$findTarget = "findLeastHealthTarget";
			$findTargetMode = "SKILL_ATTACK_MODE_SINGLE";
			$ghostAddRate = "5";
			$attackRangeRatio = "1";
			break;
			// 单体重复2次
		case 7:
			$findTarget = "findOneTarget2";
			$findTargetMode = "SKILL_ATTACK_MODE_SINGLE_TWICE";
			$ghostAddRate = "2";
			$attackRangeRatio = "1";
			break;
			// 纵向攻击，并额外随机攻击另一列
		case 9:
			$findTarget = "findColTargetWithOneColRandom";
			$findTargetMode = "SKILL_ATTACK_MODE_AOE";
			$ghostAddRate = "2";
			$attackRangeRatio = "0.75";
			break;
			// 随机攻击两个目标
		case 10:
			$findTarget = "findTowRandomTargets";
			$findTargetMode = "SKILL_ATTACK_MODE_AOE";
			$ghostAddRate = "2";
			$attackRangeRatio = "0.75";
			break;
			// 单体攻击，并额外随机攻击另一人
		case 11:
			$findTarget = "findOneTargetWithOneRandom";
			$findTargetMode = "SKILL_ATTACK_MODE_AOE";
			$ghostAddRate = "2";
			$attackRangeRatio = "0.75";
			break;
			// 血量最多
		case 12:
			$findTarget = "findMostHealthTarget";
			$findTargetMode = "SKILL_ATTACK_MODE_SINGLE";
			$ghostAddRate = "5";
			$attackRangeRatio = "1";
			break;
			// V字攻击
		case 13:
			$findTarget = "findVAttackTarget";
			$findTargetMode = "SKILL_ATTACK_MODE_AOE";
			$ghostAddRate = "2";
			$attackRangeRatio = "0.75";
			break;
			// 后排单体攻击
		case 14:
			$findTarget = "findOneTargetFromBack";
			$findTargetMode = "SKILL_ATTACK_MODE_SINGLE";
			$ghostAddRate = "5";
			$attackRangeRatio = "1";
			break;
			// 十字攻击
		case 15:
			$findTarget = "findCrossTarget";
			$findTargetMode = "SKILL_ATTACK_MODE_AOE";
			$ghostAddRate = "2";
			$attackRangeRatio = "0.75";
			break;
			// 横排穿透
		case 16:
			$findTarget = "findRowPenetrateTarget";
			$findTargetMode = "SKILL_ATTACK_MODE_AOE";
			$ghostAddRate = "2";
			$attackRangeRatio = "0.75";
			break;
			// 十字攻击(固定)
		case 17:
			$findTarget = "findFixCrossTarget";
			$findTargetMode = "SKILL_ATTACK_MODE_AOE";
			$ghostAddRate = "2";
			$attackRangeRatio = "0.75";
			break;

		//前排横向(固定)
		case 18:
			$findTarget = "findFixFrontendRow";
			$findTargetMode = "SKILL_ATTACK_MODE_AOE";
			$ghostAddRate = "2";
			$attackRangeRatio = "0.75";
			break;

		//中排横向(固定)
		case 19:
			$findTarget = "findFixMiddleRow";
			$findTargetMode = "SKILL_ATTACK_MODE_AOE";
			$ghostAddRate = "2";
			$attackRangeRatio = "0.75";
			break;

		//后排横向(固定)
		case 20:
			$findTarget = "findFixBackendRow";
			$findTargetMode = "SKILL_ATTACK_MODE_AOE";
			$ghostAddRate = "2";
			$attackRangeRatio = "0.75";
			break;
		}


		if ($findTarget != "") {
			$code .= "	GhostAddRate: $ghostAddRate,\n";
			$code .= "	AttackMode: $findTargetMode,\n";
			$code .= "	AttackRangeRatio: $attackRangeRatio,\n";

			//if ($skill['role_id'] < 0 && $for_role) {
			$code .= "	_FindTargets: $findTarget,\n";
			//} else {
			//	$code .= "	_FindTargets: buddyFindTarget($findTarget),\n";
			//}
		}
	}

	if (isset($config->TargetMode) && isset($config->AttackMode) && $config->TargetMode != 8 && $config->TargetMode != 4) {
		switch ($config->AttackMode) {
		case 0:
			//普通攻击
			$code .= "	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },\n";
			break;
		case 1:
			//自爆
			$code .= "	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { health := f.Health; f.Health = 0; return float64(health) * ".($config->KillSelfRate / 100.0)." + force },\n";
			$code .= "	IsTotemSkill: false, \n";
			break;
		case 2:
			//玩家攻击
			$code .= "	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + {$config->DefaultAttack} + f.Cultivation * ".($config->Cul2AtkRate / 100.0)." + float64(skillTrnlv) * {$config->TrnlvRate} },\n";
			$code .= "	IsTotemSkill: false, \n";
			$code .= "	IsRoleSkill: true, \n";
			break;
		case 3:
			//魂侍攻击
			$code .= "	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + {$config->DefaultAttack} + f.Cultivation * ".($config->Cul2AtkRate / 100.0)." + float64(skillTrnlv) * {$config->TrnlvRate} },\n";
			$code .= "	IsTotemSkill: false, \n";
			break;
		case 4:
			//灵宠发飙
			$code .= "	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + {$config->DefaultAttack} + float64(f.NaturalSkillLv) * {$config->TrnlvRate} },\n";
			$code .= "	IsTotemSkill: false, \n";
			break;
		case 5:
			//阵印（无伤害）
			$code .= "	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return  0   },\n";
			$code .= "	IsTotemSkill: true, \n";
		}
	}

	// 触发事件
	if (isset($config->EventTriggerType) && $config->EventTriggerType > 0 ) {
		$fightEventConst = get_fight_event_const($config->EventTriggerType);
		$buffConst = get_buff_const($config->EventTriggerBuff);

		$code .= "	_EventTrigger: func(f, t *Fighter) {\n";
		$targetFighter = 'f'; //默认检查攻击者的Buff
		if ($config->EventTriggerTarget != 0) { //检查防守者的Buff
			$targetFighter = 't'; 
		}
		$code .= "		if {$targetFighter}.HasBuff({$buffConst}) {\n";
		$code .= "			t.triggerEvent = {$fightEventConst}\n";
		$code .= "		}\n";
		$code .= "	},\n";

		if (isset($config->TriggerTargetBuff) && $config->TriggerTargetBuff) {
			$code .= "	TriggerTargetBuff: true ,\n";
		}
	}

	// 目标BUFF
	if (isset($config->TargetBuffs) && count($config->TargetBuffs) > 0) {
		$code .= "	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {\n";
		$code .= "		buffs := make([]*Buff, 0, ".count($config->TargetBuffs).")\n";

		$code .= "		buffCount := 0.0\n";
		$code .= "		_ = buffCount\n"; 

		$rate_buff = array();
		foreach ($config->TargetBuffs as $buff) {

			fix_skill_buff_config_undefined($buff);

			if($buff->RequireSkillLevel > 0) {
				$code .= "		if skillTrnlv >= {$buff->RequireSkillLevel} {\n";
			}

			if ((isset($buff->CountRate) && $buff->CountRate > 0) || (isset($buff->ValueCountRate) && $buff->ValueCountRate > 0)) {
				$buff_id += 1;
				$curr_buff_id = $buff_id;
				$code .= "		buffCount = t.GetBuffCount($curr_buff_id)\n";
			}

			if ($buff->Rate < 100 || (isset($buff->CountRate) && $buff->CountRate > 0)) {
				// 概率发生的buff
				$count_rate_code = '';
				if (isset($buff->CountRate) && $buff->CountRate > 0) {
					$count_rate_code = "- (float64(buffCount) * ".$buff->CountRate." / 100.0)";
				}

				$levelVar = get_buff_levelvar($buff->Type);

				$code .= "		if rand.Float64() < ".($buff->Rate / 100.0).(is_bad_buff($buff) ? " - (t.{$levelVar[0]} + (t.{$levelVar[1]} * {$levelVar[2]} / f.probLevel)) * 0.01" : "")." $count_rate_code  {\n";
				$code .= "			buffs = append_buff(buffs, ".get_buff_code($skill['id'], "t", $buff, false, "").")\n";

				// 计数器只在buff生效时更新
				if ((isset($buff->CountRate) && $buff->CountRate > 0) || (isset($buff->ValueCountRate) && $buff->ValueCountRate > 0)) {
					$code .= "			t.AddBuffCount($curr_buff_id)\n";
				}

				$code .= "		}\n";
			} else {
				$code .= "		buffs = append_buff(buffs, ".get_buff_code($skill['id'], "t", $buff, false, "").")\n";

				// 计数器只在buff生效时更新
				if ((isset($buff->CountRate) && $buff->CountRate > 0) || (isset($buff->ValueCountRate) && $buff->ValueCountRate > 0)) {
					$code .= "t.AddBuffCount($curr_buff_id)\n";
				}
			}

			if($buff->RequireSkillLevel > 0) {
				$code .= "		}\n";
			}

		}

		$code .= "		return buffs\n";
		$code .= "	},\n";
	}

	// 自身BUFF
	if (isset($config->SelfBuffs) && count($config->SelfBuffs) > 0) {
		if ($config->TargetMode != 8) {
			$code .= "	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16".($for_role ? ", ghost *GhostSkill" : ""). ") []*Buff {\n";
			$code .= "		buffs := make([]*Buff, 0, ".count($config->SelfBuffs).")\n";
		} else {
			$code .= "	_ApplyEffect: func() {\n";
		}

		$code .= "		buffCount := 0.0\n";
		$code .= "		_ = buffCount\n"; 

		foreach ($config->SelfBuffs as $buff) {

			fix_skill_buff_config_undefined($buff);

			if($buff->RequireSkillLevel > 0) {
				$code .= "		if skillTrnlv >= {$buff->RequireSkillLevel} {\n";
			}

			if ((isset($buff->CountRate) && $buff->CountRate > 0) || (isset($buff->ValueCountRate) && $buff->ValueCountRate > 0)) {
				$buff_id += 1;
				$curr_buff_id = $buff_id;
				$code .= "		buffCount = f.GetBuffCount($curr_buff_id)\n";
			}

			if ($buff->Rate < 100 || (isset($buff->CountRate) && $buff->CountRate > 0)) {
				// 概率发生的buff
				$count_rate_code = '';
				if (isset($buff->CountRate) && $buff->CountRate > 0) {
					$count_rate_code = "- (float64(buffCount) * ".$buff->CountRate." / 100.0)";
				}

				$code .= "		if rand.Float64() < ".($buff->Rate / 100.0)." $count_rate_code {\n";

				if ($config->TargetMode != 8)
					$code .= "			buffs = append_buff(buffs, ";

				$code .= get_buff_code($skill['id'], "f", $buff, $for_role, "5");

				if ($config->TargetMode != 8)
					$code .= ")";

				$code .= "\n";

				// 计数器只在buff生效时更新
				if ((isset($buff->CountRate) && $buff->CountRate > 0) || (isset($buff->ValueCountRate) && $buff->ValueCountRate > 0)) {
					$code .= "			f.AddBuffCount($curr_buff_id)\n";
				}

				$code .= "		}\n";
			} else {
				if ($config->TargetMode == 8) {
					$code .= "			".get_buff_code($skill['id'], "f", $buff, $for_role, "5");
				} else {
					$code .= "		buffs = append_buff(buffs, ".get_buff_code($skill['id'], "f", $buff, $for_role, "5").")\n";
				}

				// 计数器只在buff生效时更新
				if ((isset($buff->CountRate) && $buff->CountRate > 0) || (isset($buff->ValueCountRate) && $buff->ValueCountRate > 0)) {
					$code .= "f.AddBuffCount($curr_buff_id)\n";
				}
			}
			if($buff->RequireSkillLevel > 0) {
				$code .= "		}\n";
			}

		}

		if ($config->TargetMode != 8)
			$code .= "		return buffs\n";

		$code .= "	},\n";
	}

	// 伙伴BUFF
	if (isset($config->BuddyBuffs) && count($config->BuddyBuffs) > 0) {
		$code .= "	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16".($for_role ? ", ghost *GhostSkill" : ""). ") []*Buff {\n";
		$code .= "		buffs := make([]*Buff, 0, ".count($config->BuddyBuffs).")\n";

		$code .= "		buffCount := 0.0\n";
		$code .= "		_ = buffCount\n"; 

		$i = 0;
		foreach ($config->BuddyBuffs as $buff) {

			fix_skill_buff_config_undefined($buff);

			if($buff->RequireSkillLevel > 0) {
				$code .= "		if skillTrnlv >= {$buff->RequireSkillLevel} {\n";
			}

			$i ++;
			$ghostAddRate = "";
			//buff选择目标方式
			switch ($buff->TargetMode) {
			case 0:
				$code .= "		buddies{$i} := f.getBuddies()\n";
				$ghostAddRate = "1";
				break;
			case 1:
				$code .= "		buddies{$i} := findBuddyWithoutSelf(f)\n";
				$ghostAddRate = "5";
				break;
			case 2:
				$code .= "		buddies{$i} := []*Fighter{findLeastHealthBuddy(f)}\n";
				$ghostAddRate = "5";
				break;
			case 3:
				$code .= "		buddies{$i} := findBattlePetBuddy(f)\n";
				$ghostAddRate = "5";
				break;
			case 4:
				$code .= "		buddies{$i} := findOurBuddies(f)\n";
				$ghostAddRate = "5";
				break;
			case 5:
				$code .= "		buddies{$i} := findOurRoleByFixRow(f, fixFrontendRowIndex)\n";
				$ghostAddRate = "5";
				break;
			case 6:
				$code .= "		buddies{$i} := findOurRoleByFixRow(f, fixMiddleRowIndex)\n";
				$ghostAddRate = "5";
				break;
			case 7:
				$code .= "		buddies{$i} := findOurRoleByFixRow(f, fixBackendRowIndex)\n";
				$ghostAddRate = "5";
				break;
			case 8:
				$code .= "		buddies{$i} := []*Fighter{findDeadFighter(f.side.Fighters)}\n";
				$ghostAddRate = "5";
				break;

			}

			$sunder_buff_code = '';
			if ($buff->Type == 22) {
				$sunder_buff_code = "&& buddy{$i}.sunderValue > 0";
			}

			$buff_to_dead = 'false /*dead target*/';
			if ($buff->TargetMode == 8) {
				$buff_to_dead = 'true /*buff to death*/';
			}

			$code .= "		for _, buddy{$i} := range buddies{$i} {\n";

			$code .= "			if buddy{$i} != nil && ( $buff_to_dead || (buddy{$i}.Health > MIN_HEALTH $sunder_buff_code)) {\n";

			if ($buff->Rate < 100 || (isset($buff->CountRate) && $buff->CountRate > 0)) {
				// 概率发生的buff
				$count_rate_code = '';
				if (isset($buff->CountRate) && $buff->CountRate > 0) {
					$count_rate_code = "- (float64(buffCount) * ".$buff->CountRate." / 100.0)";
				}

				$levelVar = get_buff_levelvar($buff->Type);

				$code .= "		if rand.Float64() < ".($buff->Rate / 100.0).(is_bad_buff($buff) ? " - (buddy{$i}.{$levelVar[0]} + (buddy{$i}.{$levelVar[1]} * {$levelVar[2]} / f.probLevel)) * 0.01" : "")." $count_rate_code  {\n";
			}

			if ((isset($buff->CountRate) && $buff->CountRate > 0) || (isset($buff->ValueCountRate) && $buff->ValueCountRate > 0)) {
				$buff_id += 1;
				$curr_buff_id = $buff_id;
				$code .= "				buffCount = buddy{$i}.GetBuffCount($curr_buff_id)\n";
			}

			$code .= "				buffs = append_buff(buffs, ".get_buff_code($skill['id'], "buddy{$i}", $buff, $for_role, $ghostAddRate).")\n";

			// 计数器只在buff生效时更新
			if ((isset($buff->CountRate) && $buff->CountRate > 0) || (isset($buff->ValueCountRate) && $buff->ValueCountRate > 0)) {
				$code .= "					buddy{$i}.AddBuffCount($curr_buff_id)\n";
			}

			if ($buff->Rate < 100 || (isset($buff->CountRate) && $buff->CountRate > 0)) {
				$code .= "				}\n";
			}
			$code .= "			}\n";
			//if ($buff->TargetMode != 2) {
				$code .= "		}\n";
			//}

			if($buff->RequireSkillLevel > 0) {
				$code .= "		}\n";
			}
		}
		$code .= "		return buffs\n";
		$code .= "	},\n";
	}
	$i = 0;
	if(isset($config->EnemyCalls) && count($config->EnemyCalls)>0){
		$code .= "	CallEnemys:[]CallInfo{";
		foreach ($config->EnemyCalls as $call) {
			$code .= 'CallInfo{ '.$call->Enemys.','.$call->Position.'}';
			if($i < count($config->EnemyCalls)-1){
				$code .=',';
			}
			$i++;
		}
		$code .= "},\n";
	}

	return $code;
}

function get_buff_levelvar($type) {
	switch ($type) {
	case 23:
		return array("Sleep","SleepLevel", "SLEEP_LEVEL_ARG");
	case 5:
		return array("Dizziness","DizzinessLevel", "DIZZINESS_LEVEL_ARG");
	case 10:
		return array("Random","RandomLevel", "RANDOM_LEVEL_ARG");
	case 24:
		return array("DisableSkill","DisableSkillLevel", "DISABLE_SKILL_LEVEL_ARG");
	case 6:
		return array("Poisoning","PoisoningLevel", "POISONING_LEVEL_ARG");
	}

	return array();
}

// Deprecated!!!
function gen_buff_variable($buffs) {
	$code = '';

	// 检查是否需要生成变量
	$hasWillLevel = false;
	foreach ($buffs as $buff) {
		if ((isset($buff->CountRate) && $buff->CountRate > 0) || (isset($buff->ValueCountRate) && $buff->ValueCountRate > 0)) {
			$code .= "		buffCount := 0.0\n";
		}

		$levelVar = get_buff_levelvar($buff->Type);
		if ($buff->Rate < 100 || $buff->CountRate > 0) {
			if ($hasWillLevel == false) {
				$code .= "		buffLevel := t.{$levelVar[1]}\n";
				//$code .= "		if t.FameLevel > 0 && f.FameLevel > 0 && t.FameLevel > f.FameLevel {\n";
				//$code .= "			buffLevel += float64(t.FameLevel - f.FameLevel) * 100\n";
				//$code .= "		}\n";
				$hasWillLevel = true;
			}
		}
	}

	return $code;
}

function get_buff_code($id, $target, $buff, $for_self_or_buddy, $ghostAddRate) {
	$buff_const = get_buff_const($buff->Type);
	$buff_value = get_buff_value($target, $buff, $for_self_or_buddy, $ghostAddRate);
	$buff_uncleanable = $buff->BuffUncleanable ? "true" : "false";
	return <<<BUFF
{$target}.addBuff(f, {$buff_const}, {$buff_value}, {$buff->Keep}, {$id}, {$buff->Override}, {$buff_uncleanable})
BUFF;
	//return $target.".addBuff(f, ".get_buff_const($buff->Type).", ".get_buff_value($target, $buff, $for_self_or_buddy, $ghostAddRate).", ".$buff->Keep.", ".$id.", ".$buff->Override.")";
}

function get_fight_event_const($type) {
	switch($type) {
	case 0:
		return 'FE_NONE';
	case 1:
		return 'FE_DODGE';
	case 2:
		return 'FE_CRIT';
	case 3:
		return 'FE_BLOCK';
	case 4:
		return 'FE_SQUELCH';
	}
	die('undefine fight event num');
}

function get_buff_const($type) {
	switch ($type) {	
	case 0 : // 精气
		return "BUFF_POWER";
	case 1 : // 速度
		return "BUFF_SPEED";
	case 2 : // 攻击
		return "BUFF_ATTACK";
	case 3 : // 防御
		return "BUFF_DEFEND";
	case 4 : // 生命
		return "BUFF_HEALTH";
	case 5 : // 眩晕
		return "BUFF_DIZZINESS";
	case 6 : // 中毒
		return "BUFF_POISONING";
	case 7 : // 清除负面buff
		return "BUFF_CLEAN_BAD";
	case 8 : // 清除增益buff
		return "BUFF_CLEAN_GOOD";
	case 9 : // 伤害减免
		return "BUFF_REDUCE_HURT";
	case 10: // 混乱
		return "BUFF_RANDOM";
	case 11: // 格挡概率
		return "BUFF_BLOCK";
	case 12: // 格挡概率等级
		return "BUFF_BLOCK_LEVEL";
	case 13: // 闪避概率等级
		return "BUFF_DODGE_LEVEL";
	case 14: // 暴击等级
		return "BUFF_CRITIAL_LEVEL";
	case 15: // 命中等级
		return "BUFF_HIT_LEVEL";
	case 16: // 伤害百分比
		return "BUFF_HURT_ADD";
	case 17: // 最大生命值
		return "BUFF_MAX_HEALTH";
	case 18: // 守卫者免伤
		return "BUFF_KEEPER_REDUCE_HURT";
	case 19: // 吸引火力
		return "BUFF_ATTRACT_FIRE";
	case 20: // 破击
		return "BUFF_DESTROY_LEVEL";
	case 21: // 韧性
		return "BUFF_TENACITY_LEVEL";
	case 22: // 护甲
		return "BUFF_SUNDER";
	case 23: // 意志
		return "BUFF_SLEEP";
	case 24: // 禁用绝招
		return "BUFF_DISABLE_SKILL";
	case 25: // 直接免伤
		return "BUFF_DIRECT_REDUCE_HURT";
	case 26: // 伤害吸收
		return "BUFF_ABSORB_HURT";		
	case 27: // 魂力
		return "BUFF_GHOST_POWER";		
	case 28: // 灵宠存活回合
		return "BUFF_PET_LIVE_ROUND";		
	case 29: // 伙伴技能使用次数
		return "BUFF_BUDDY_SKILL";		
	case 30: // 清除伤害吸收
		return "BUFF_CLEAR_ABSORB_HURT";		
	case 31:
		return "BUFF_SLEEP_LEVEL";		
	case 32: 
		return "BUFF_DIZZINESS_LEVEL";		
	case 33: 
		return "BUFF_RANDOM_LEVEL";		
	case 34: 
		return "BUFF_DISABLE_SKILL_LEVEL";		
	case 35: 
		return "BUFF_POISONING_LEVEL";		
	case 36: 
		return "BUFF_RECOVER_BUDDY_SKILL";		
	case 37:
		return "BUFF_MAKE_POWER_FULL";
	case 38:
		return "BUFF_DOGE";
	case 39:
		return "BUFF_HIT";
	case 40:
		return "BUFF_CRITIAL";
	case 41:
		return "BUFF_TENACITY";
	case 42:
		return "BUFF_TAKE_SUNDER";
	case 43:
		return "BUFF_DEFEND_PERSENT";
	case 44:
		return "BUFF_SUNDER_STATE";
	case 45:
		return "BUFF_HEALTH_PERCENT";
	case 46:
		return "BUFF_ALL_RESIST";
	case 47:
		return "BUFF_REBOTH_HEALTH";
	case 48:
		return "BUFF_REBOTH_HEALTH_PERCENT";

	default:
		die("unkonw buff {$type}");
	}

	return "";
}

function get_buff_raw($type) {
	switch ($type) {
	case 1 : // 速度
		return "Speed";
	case 2 : // 攻击
		return "Attack";
	case 3 : // 防御
		return "Defend";
	case 11: // 格挡概率
		return "Block";
	case 12: // 格挡概率等级
		return "BlockLevel";
	case 13: // 闪避概率等级
		return "DodgeLevel";
	case 14: // 暴击等级
		return "CritialLevel";
	case 15: // 命中等级
		return "HitLevel";
	}

	return "";
}

function get_AppendSpecialType_const($type) {
	switch ($type) {	
	case 0 : // 无
		return "APPEND_SPECIAL_TYPE_NONE";
	case 1 : // 每命中敌人增加精气
		return "APPEND_SPECIAL_TYPE_ATTACKED_INC_POWER";
	}

	return "";
}

function get_buff_value($target, $buff, $for_self_or_buddy, $ghostAddRate) {
	if ($buff->Type == 5 ||  //晕眩
		$buff->Type == 7 || //清除buff
		$buff->Type == 8 || //清除debuff
		$buff->Type == 10 ||  //混乱
		$buff->Type == 30 ||  //清除伤害吸收
		$buff->Type == 24 ||  //禁魔
		$buff->Type == 44 ||  //破甲状态
		/*没有Value的请在上面加*/
		false) 
		return '1';

	$code = '';

	if ($buff->BuffSign != 0)
		$code .= "-";

	$code .= "(buff_value(";

	$has_value = false;

	if ($buff->BaseValue != 0 
		|| $buff->Type == 43) { //百分比降低防御特殊处理，允许 BaseValue==0 考虑允许所有技能的BaseValue==0
		$code .= $buff->BaseValue;
		$has_value = true;
	}

	if ($buff->RawValueRate != 0) {
		$raw = get_buff_raw($buff->Type);

		if ($raw != '') {
			if ($has_value) $code .= " + ";

			$code .= $target.".raw.".$raw." * ".($buff->RawValueRate / 100.0);

			$has_value = true;
		}
	}

	if ($buff->AttackRate != 0) {
		if ($has_value) $code .= " + ";

		$code .= "f.Attack";

		if ($buff->AttackRate != 100)
			$code .= " * ".($buff->AttackRate / 100.0);

		$has_value = true;
	}

	if ($buff->SkillForceRate != 0) {
		if ($has_value) $code .= " + ";

		$code .= "force";

		if ($buff->SkillForceRate != 100)
			$code .= " * ".($buff->SkillForceRate / 100.0);

		$has_value = true;
	}

	if ($buff->TrnlvRate != 0) {
		if ($has_value) $code .= " + ";

		$code .= "float64(skillTrnlv)";

		if ($buff->TrnlvRate != 1)
			$code .= " * ".($buff->TrnlvRate);

		$has_value = true;
	}

	if ($buff->HurtRate != 0) {
		if ($has_value) $code .= " + ";

		$code .= "float64(hurt)";

		if ($buff->HurtRate != 100)
			$code .= " * ".($buff->HurtRate / 100.0);

		$has_value = true;
	}

	if (isset($buff->Cul2ValueRate) && $buff->Cul2ValueRate != 0) {
		if ($has_value) $code .= " + ";

		$code .= "f.Cultivation";

		if ($buff->Cul2ValueRate != 100)
			$code .= " * ".($buff->Cul2ValueRate / 100.0);

		$has_value = true;
	}

	if (isset($buff->ValueCountRate) && $buff->ValueCountRate != 0) {
		if ($has_value) $code .= " - ";

		$code .= "buffCount * ".$buff->ValueCountRate;

		$has_value = true;
	}

	if ($buff->Type == 4 && $for_self_or_buddy) {
		if ($has_value) $code .= " + ";

		$code .= "ghost.CureAdd * $ghostAddRate";
	}

	$code .= "))";

	if ($buff->Type == 4 && $for_self_or_buddy) {
		$code .= " * (1 + ghost.CureAddRate)";
	}

	return $code;
}

function is_bad_buff($buff) {
	$defence_attribute = get_buff_levelvar($buff->Type);
	return !empty($defence_attribute); // Why we should defend a good buff?!?!?!
}

function is_ghost_skill($skill_type) {
	if ($skill_type == 8 || $skill_type==7) {
		return true;
	}
	return false;
}
?>
