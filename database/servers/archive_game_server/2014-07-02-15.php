<?php
db_execute($db, "

INSERT INTO `skill` (`name`, `type`, `child_type`, `sign`, `music_sign`, `role_id`, `required_level`, `info`, `jump_attack`, `display_param`, `config`, `quality`, `can_add_level`, `parent_skill_id`, `skill_level`, `order`) VALUES
('t', 7, 1, NULL, NULL, 1, 1, 'desc', 0, 0, '{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}', 0, 0, 0, 0, 0);

");
?>

