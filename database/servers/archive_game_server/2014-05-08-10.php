<?php

db_execute($db, "
DROP TABLE IF EXISTS `player_mission`;
CREATE TABLE `player_mission` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `key` int(11) NOT NULL COMMENT '拥有的区域钥匙数',
  `last_key` int(11) NOT NULL COMMENT '最近一次开启区域的钥匙数',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家区域数据';

DROP TABLE IF EXISTS `player_mission_record`;
CREATE TABLE `player_mission_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `mission_id` int(11) NOT NULL COMMENT '开启的区域ID',
  `open_time` int(11) NOT NULL COMMENT '区域开启时间',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='区域记录';

DROP TABLE IF EXISTS `player_mission_level`;
CREATE TABLE `player_mission_level` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `lock` int(11) NOT NULL COMMENT '当前的关卡权值',
  `last_lock` int(11) NOT NULL COMMENT '最近一次开启的关卡权值',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家区域关卡数据';

DROP TABLE IF EXISTS `player_mission_level_record`;
CREATE TABLE `player_mission_level_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `mission_level_id` int(11) NOT NULL COMMENT '开启的关卡ID',
  `open_time` int(11) NOT NULL COMMENT '关卡开启时间',
  `star` tinyint(4) NOT NULL DEFAULT '0' COMMENT '通关星数',
  `round` tinyint(4) NOT NULL DEFAULT '0' COMMENT '通关回合数',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='关卡记录';


alter table `mission_level` modify `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '区域关卡ID';
alter table `mission_level` modify `mission_id` smallint(6) NOT NULL COMMENT '区域ID';

");
?>