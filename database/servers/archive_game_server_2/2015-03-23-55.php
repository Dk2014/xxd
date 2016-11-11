<?php 
$this->AddSQL("
REPLACE INTO `skill` (`id`, `name`, `type`, `child_type`, `child_kind`, `sign`, `music_sign`, `role_id`, `required_level`, `info`, `jump_attack`, `display_param`, `config`, `quality`, `can_add_level`, `parent_skill_id`, `order`, `target`, `warcry`, `cheat_id`, `auto_learn_level`, `required_fame_level`, `required_friendship_level`, `info_vars`)
VALUES
	(1164, '魂力果实', 1, 5, 2, 'HunLiGuoShi', 'HunLiGuoShi', 5, 45, '增加魂力100，并增加我方全体500+\$a点命中1回合', 0, 0, '{\"TargetMode\":4,\"AttackMode\":2,\"KillSelfRate\":1,\"DefaultAttack\":2000,\"Cul2AtkRate\":60,\"TrnlvRate\":0,\"DecPower\":0,\"IncPower\":0,\"HurtAdd\":0,\"HurtAddRate\":0,\"CureAdd\":0,\"CureAddRate\":0,\"Critial\":0,\"ReduceDefend\":0,\"SunderAttack\":5,\"MustHit\":false,\"GhostOverrideBuddyBuff\":false,\"GhostOverrideSelfBuff\":false,\"GhostOverrideTargetBuff\":false,\"AttackNum\":1,\"SelfBuffs\":[],\"TargetBuffs\":[],\"BuddyBuffs\":[{\"Type\":27,\"Keep\":0,\"Override\":1,\"Rate\":100,\"CountRate\":0,\"BuffSign\":0,\"BaseValue\":100,\"RawValueRate\":0,\"AttackRate\":0,\"SkillForceRate\":0,\"HurtRate\":0,\"Cul2ValueRate\":0,\"TrnlvRate\":0,\"ValueCountRate\":0,\"TargetMode\":0},{\"Type\":15,\"Keep\":1,\"Override\":1,\"Rate\":100,\"CountRate\":0,\"BuffSign\":0,\"BaseValue\":500,\"RawValueRate\":0,\"AttackRate\":0,\"SkillForceRate\":0,\"HurtRate\":0,\"Cul2ValueRate\":0,\"TrnlvRate\":10,\"ValueCountRate\":0,\"TargetMode\":0}],\"EnemyCalls\":[],\"AppendSpecialType\":0,\"AppendSpecialValue\":0,\"EventTriggerType\":0,\"EventTriggerTarget\":0,\"EventTriggerBuff\":6,\"TriggerTargetBuff\":false}', 0, 1, 1164, 0, 2, '魂力在沸腾', 0, 1, 9, 0, '{\"a\":{\"Cul2AtkRate\":\"0\",\"TrnLvRate\":\"10\"}}');

REPLACE INTO `skill` (`id`, `name`, `type`, `child_type`, `child_kind`, `sign`, `music_sign`, `role_id`, `required_level`, `info`, `jump_attack`, `display_param`, `config`, `quality`, `can_add_level`, `parent_skill_id`, `order`, `target`, `warcry`, `cheat_id`, `auto_learn_level`, `required_fame_level`, `required_friendship_level`, `info_vars`)
	VALUES
		(1101, '致胜千里附属', 8, 1, 1, 'SanQianLuoShuiJianM', 'SanQianLuoShuiJian', 8, 1, '', 0, 0, '{\"TargetMode\":1,\"AttackMode\":3,\"KillSelfRate\":1,\"DefaultAttack\":2000,\"Cul2AtkRate\":150,\"TrnlvRate\":50,\"DecPower\":0,\"IncPower\":0,\"HurtAdd\":0,\"HurtAddRate\":0,\"CureAdd\":0,\"CureAddRate\":0,\"Critial\":0,\"ReduceDefend\":0,\"SunderAttack\":5,\"MustHit\":false,\"GhostOverrideBuddyBuff\":false,\"GhostOverrideSelfBuff\":false,\"GhostOverrideTargetBuff\":false,\"AttackNum\":1,\"SelfBuffs\":[],\"TargetBuffs\":[],\"BuddyBuffs\":[{\"Type\":2,\"Keep\":2,\"Override\":1,\"Rate\":100,\"CountRate\":0,\"BuffSign\":0,\"BaseValue\":2000,\"RawValueRate\":0,\"AttackRate\":0,\"SkillForceRate\":0,\"HurtRate\":0,\"Cul2ValueRate\":0,\"TrnlvRate\":40,\"ValueCountRate\":0,\"TargetMode\":0}],\"EnemyCalls\":[],\"AppendSpecialType\":0,\"AppendSpecialValue\":0,\"EventTriggerType\":0,\"EventTriggerTarget\":0,\"EventTriggerBuff\":6,\"TriggerTargetBuff\":false}', 0, 0, 0, 0, 2, '', 0, 1, 0, 0, '{}');

REPLACE INTO `skill` (`id`, `name`, `type`, `child_type`, `child_kind`, `sign`, `music_sign`, `role_id`, `required_level`, `info`, `jump_attack`, `display_param`, `config`, `quality`, `can_add_level`, `parent_skill_id`, `order`, `target`, `warcry`, `cheat_id`, `auto_learn_level`, `required_fame_level`, `required_friendship_level`, `info_vars`)
	VALUES
		(1523, '铁壁', 1, 5, 2, 'TieBi', 'TieBuShan', 9, 60, '提升自己1000+\$a点格挡，并吸引火力，持续2回合', 0, 0, '{\"TargetMode\":4,\"AttackMode\":0,\"KillSelfRate\":1,\"DefaultAttack\":0,\"Cul2AtkRate\":0,\"TrnlvRate\":0,\"DecPower\":0,\"IncPower\":0,\"HurtAdd\":0,\"HurtAddRate\":0,\"CureAdd\":0,\"CureAddRate\":0,\"Critial\":0,\"ReduceDefend\":0,\"SunderAttack\":0,\"MustHit\":false,\"GhostOverrideBuddyBuff\":false,\"GhostOverrideSelfBuff\":false,\"GhostOverrideTargetBuff\":false,\"AttackNum\":1,\"SelfBuffs\":[{\"Type\":11,\"Keep\":2,\"Override\":1,\"Rate\":100,\"CountRate\":0,\"BuffSign\":0,\"BaseValue\":1000,\"RawValueRate\":0,\"AttackRate\":0,\"SkillForceRate\":0,\"HurtRate\":0,\"Cul2ValueRate\":0,\"TrnlvRate\":6,\"ValueCountRate\":0},{\"Type\":19,\"Keep\":2,\"Override\":1,\"Rate\":100,\"CountRate\":0,\"BuffSign\":0,\"BaseValue\":1,\"RawValueRate\":0,\"AttackRate\":0,\"SkillForceRate\":0,\"HurtRate\":0,\"Cul2ValueRate\":0,\"TrnlvRate\":0,\"ValueCountRate\":0}],\"TargetBuffs\":[],\"BuddyBuffs\":[],\"EnemyCalls\":[],\"AppendSpecialType\":0,\"AppendSpecialValue\":0,\"EventTriggerType\":0,\"EventTriggerTarget\":0,\"EventTriggerBuff\":6,\"TriggerTargetBuff\":false}', 0, 0, 1523, 0, 1, '铜墙铁壁', 0, 1, 10, 0, '{\n    \"a\": {\n        \"Cul2AtkRate\": \"\",\n        \"TrnLvRate\": \"6\"\n    }\n}');


REPLACE INTO `skill` (`id`, `name`, `type`, `child_type`, `child_kind`, `sign`, `music_sign`, `role_id`, `required_level`, `info`, `jump_attack`, `display_param`, `config`, `quality`, `can_add_level`, `parent_skill_id`, `order`, `target`, `warcry`, `cheat_id`, `auto_learn_level`, `required_fame_level`, `required_friendship_level`, `info_vars`)
	VALUES
		(1528, '神兵卷', 1, 1, 2, 'ShenBingJuan', 'LingZhuZhou', 12, 110, '为我方全体增加1000+\$a点攻击，500+\$b点爆击，持续2回合，可叠加2次', 0, 0, '{\"TargetMode\":4,\"AttackMode\":0,\"KillSelfRate\":1,\"DefaultAttack\":0,\"Cul2AtkRate\":0,\"TrnlvRate\":0,\"DecPower\":0,\"IncPower\":0,\"HurtAdd\":0,\"HurtAddRate\":0,\"CureAdd\":0,\"CureAddRate\":0,\"Critial\":0,\"ReduceDefend\":0,\"SunderAttack\":0,\"MustHit\":false,\"GhostOverrideBuddyBuff\":false,\"GhostOverrideSelfBuff\":false,\"GhostOverrideTargetBuff\":false,\"AttackNum\":1,\"SelfBuffs\":[],\"TargetBuffs\":[],\"BuddyBuffs\":[{\"Type\":2,\"Keep\":2,\"Override\":2,\"Rate\":100,\"CountRate\":0,\"BuffSign\":0,\"BaseValue\":1000,\"RawValueRate\":0,\"AttackRate\":0,\"SkillForceRate\":0,\"HurtRate\":0,\"Cul2ValueRate\":0,\"TrnlvRate\":20,\"ValueCountRate\":0,\"TargetMode\":0},{\"Type\":14,\"Keep\":2,\"Override\":2,\"Rate\":100,\"CountRate\":0,\"BuffSign\":0,\"BaseValue\":500,\"RawValueRate\":0,\"AttackRate\":0,\"SkillForceRate\":0,\"HurtRate\":0,\"Cul2ValueRate\":0,\"TrnlvRate\":10,\"ValueCountRate\":0,\"TargetMode\":0}],\"EnemyCalls\":[],\"AppendSpecialType\":0,\"AppendSpecialValue\":0,\"EventTriggerType\":0,\"EventTriggerTarget\":0,\"EventTriggerBuff\":6,\"TriggerTargetBuff\":false}', 0, 0, 1528, 0, 1, '神兵符卷', 0, 1, 12, 0, '{\n    \"a\": {\n        \"Cul2AtkRate\": \"0\",\n        \"TrnLvRate\": \"20\"\n    },\n    \"b\": {\n        \"Cul2AtkRate\": \"0\",\n        \"TrnLvRate\": \"10\"\n    }\n}');


");

?>