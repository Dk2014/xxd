<?php
db_execute($db, "

ALTER TABLE `player_ghost` DROP column `pos`;
ALTER TABLE `player_ghost` ADD COLUMN `is_equip` tinyint(1) DEFAULT '0' COMMENT '是否装备';


CREATE TABLE `player_ghost_equipment` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `main_ghost_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '主魂侍id',
  `assist_ghost_id1` bigint(20) NOT NULL DEFAULT '0' COMMENT '辅魂侍id1',
  `assist_ghost_id2` bigint(20) NOT NULL DEFAULT '0' COMMENT '辅魂侍id2',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家魂侍装备表';


ALTER TABLE `player_ghost` CHANGE dge_level dodge_level int(11) NOT NULL DEFAULT '0' COMMENT '闪避等级';
ALTER TABLE `player_ghost` CHANGE cri_level crit_level  int(11) NOT NULL DEFAULT '0' COMMENT '暴击等级';
ALTER TABLE `player_ghost` CHANGE blk_level block_level int(11) NOT NULL DEFAULT '0' COMMENT '格挡等级';

");
?>