<?php
db_execute($db,"

DROP TABLE IF EXISTS `player_hard_level_state`;
CREATE TABLE `player_hard_level_state` (
	`id` bigint(20) NOT NULL COMMENT 'ID',
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`mission_level_id` int(11) NOT NULL COMMENT '难度关卡ID',
	`battle_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后一次战败时间',
	`pos1` int(11) NOT NULL DEFAULT '0' COMMENT '位置1上的敌人生命',
	`pos2` int(11) NOT NULL DEFAULT '0' COMMENT '位置2上的敌人生命',
	`pos3` int(11) NOT NULL DEFAULT '0' COMMENT '位置3上的敌人生命',
	`pos4` int(11) NOT NULL DEFAULT '0' COMMENT '位置4上的敌人生命',
	`pos5` int(11) NOT NULL DEFAULT '0' COMMENT '位置5上的敌人生命',
	`pos6` int(11) NOT NULL DEFAULT '0' COMMENT '位置6上的敌人生命',
	`pos7` int(11) NOT NULL DEFAULT '0' COMMENT '位置7上的敌人生命',
	`pos8` int(11) NOT NULL DEFAULT '0' COMMENT '位置8上的敌人生命',
	`pos9` int(11) NOT NULL DEFAULT '0' COMMENT '位置9上的敌人生命',
	`pos10` int(11) NOT NULL DEFAULT '0' COMMENT '位置10上的敌人生命',
	`pos11` int(11) NOT NULL DEFAULT '0' COMMENT '位置11上的敌人生命',
	`pos12` int(11) NOT NULL DEFAULT '0' COMMENT '位置12上的敌人生命',
	`pos13` int(11) NOT NULL DEFAULT '0' COMMENT '位置13上的敌人生命',
	`pos14` int(11) NOT NULL DEFAULT '0' COMMENT '位置14上的敌人生命',
	`pos15` int(11) NOT NULL DEFAULT '0' COMMENT '位置15上的敌人生命',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='难度关卡失败状态记录';
");
?>
