<?php
db_execute($db, "
drop table if exists `hard_level`;
create table `hard_level`(
	`id` smallint(6) NOT NULL AUTO_INCREMENT,
	`mission_id` smallint(6)  NOT NULL COMMENT '关联的区域ID',
	`lock` int(11) NOT NULL COMMENT '进入的权值',
	`desc` varchar(100) NOT NULL COMMENT '关卡描述',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='难度关卡';

ALTER TABLE `mission_level` CHANGE COLUMN  `parent_type` 
	`parent_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '关联关卡类型(0--无;1--资源关卡;2--通天塔;8--难度关卡)';

ALTER TABLE `enemy_deploy_form` CHANGE COLUMN `battle_type` 
	`battle_type` tinyint(4) NOT NULL COMMENT '战场类型(0--关卡;1--资源关卡;2--极限关卡;3--多人关卡;8--难度关卡)';

DROP TABLE IF EXISTS `player_hard_level_record`;
CREATE TABLE `player_hard_level_record` (
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`level_id` int(11) NOT NULL COMMENT '开启的关卡ID',
	`open_time` bigint(20) NOT NULL COMMENT '关卡开启时间',
	`score` int(11) NOT NULL DEFAULT '0' COMMENT '得分',
	`round` tinyint(4) NOT NULL DEFAULT '0' COMMENT '通关回合数',
	`daily_num` tinyint(4) NOT NULL COMMENT '当日已进入关卡的次数',
	`last_enter_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后一次进入时间',
	PRIMARY KEY (`id`),
KEY `idx_pid` (`pid`)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='难度关卡记录';

");
?>
