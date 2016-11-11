<?php
db_execute($db, "

ALTER TABLE `enemy_role` CHANGE COLUMN `skill_id` `skill_id` smallint(6) DEFAULT '0' COMMENT '绝招ID';
ALTER TABLE `enemy_role` CHANGE COLUMN `skill_force` `skill_force` int(11) DEFAULT '0' COMMENT '绝招威力';

ALTER TABLE `enemy_role` CHANGE COLUMN `skill2_id` `skill2_id` smallint(6) DEFAULT '0' COMMENT '绝招2 ID';
ALTER TABLE `enemy_role` CHANGE COLUMN `skill2_force`  `skill2_force` int(11) DEFAULT '0' COMMENT '绝招2 威力';

ALTER TABLE `enemy_role` CHANGE COLUMN `release_num`  `release_num`  tinyint(4) DEFAULT '0' COMMENT '释放次数';
ALTER TABLE `enemy_role` CHANGE COLUMN `recover_round_num`  `recover_round_num` tinyint(4) DEFAULT '0' COMMENT '恢复回合数';
ALTER TABLE `enemy_role` CHANGE COLUMN `common_attack_num`  `common_attack_num` tinyint(4) DEFAULT '0'  COMMENT '入场普通攻击次数';
ALTER TABLE `enemy_role` CHANGE COLUMN `skill_wait`  `skill_wait` tinyint(4) DEFAULT '0'  COMMENT '绝招蓄力回合';

");
?>