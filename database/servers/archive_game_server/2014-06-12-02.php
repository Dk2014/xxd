<?php

db_execute($db, "

ALTER TABLE `player_sword_soul_equipment`
CHANGE COLUMN `pos0` `pos1`  bigint(20) NOT NULL COMMENT '装备位置1的剑心' AFTER `role_id`,
CHANGE COLUMN `pos1` `pos2`  bigint(20) NOT NULL COMMENT '装备位置2的剑心' AFTER `pos1`,
CHANGE COLUMN `pos2` `pos3`  bigint(20) NOT NULL COMMENT '装备位置3的剑心' AFTER `pos2`,
CHANGE COLUMN `pos3` `pos4`  bigint(20) NOT NULL COMMENT '装备位置4的剑心' AFTER `pos3`,
CHANGE COLUMN `pos4` `pos5`  bigint(20) NOT NULL COMMENT '装备位置5的剑心' AFTER `pos4`,
CHANGE COLUMN `pos5` `pos6`  bigint(20) NOT NULL COMMENT '装备位置6的剑心' AFTER `pos5`,
CHANGE COLUMN `pos6` `pos7`  bigint(20) NOT NULL COMMENT '装备位置7的剑心' AFTER `pos6`,
CHANGE COLUMN `pos7` `pos8`  bigint(20) NOT NULL COMMENT '装备位置8的剑心' AFTER `pos7`,
CHANGE COLUMN `pos8` `pos9`  bigint(20) NOT NULL COMMENT '装备位置9的剑心' AFTER `pos8`;

ALTER TABLE `player_sword_soul_equipment`
ADD COLUMN `is_equipped_xuanyuan`  tinyint(4) NOT NULL DEFAULT 0 AFTER `role_id`;

");
?>