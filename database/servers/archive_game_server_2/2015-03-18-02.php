<?php
$this->AddSQL("
CREATE TABLE `player_totem` (
	`id` bigint(20) NOT NULL COMMENT '主键',
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`totem_id` smallint(6) NOT NULL COMMENT '阵印ID',
	`level` tinyint(4) NOT NULL DEFAULT '1' COMMENT '等级',
	`skill_id` smallint(6) NOT NULL COMMENT '技能',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家阵印表';

CREATE TABLE `player_totem_info` (
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`ingot_call_daily_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '每日元宝召唤次数',
	`ingot_call_timestamp` bigint(20) NOT NULL DEFAULT '0' COMMENT '元宝召唤时间戳',
	`rock_rune_num` int(11) NOT NULL DEFAULT '0' COMMENT '石附文数量',
	`jade_rune_num` int(11) NOT NULL DEFAULT '0' COMMENT '玉附文数量',
	`pos1` bigint(20) NOT NULL DEFAULT '0' COMMENT '装备位置1的阵印id',
	`pos2` bigint(20) NOT NULL DEFAULT '0' COMMENT '装备位置2的阵印id',
	`pos3` bigint(20) NOT NULL DEFAULT '0' COMMENT '装备位置3的阵印id',
	`pos4` bigint(20) NOT NULL DEFAULT '0' COMMENT '装备位置4的阵印id',
	`pos5` bigint(20) NOT NULL DEFAULT '0' COMMENT '装备位置4的阵印id',
	PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家阵印装备表';

CREATE TABLE `totem` (
	`id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键',
	`name` varchar(30) NOT NULL DEFAULT '' COMMENT '阵印名称',
	`sign` varchar(30) NOT NULL DEFAULT '' COMMENT '资源标识',
	`quality` tinyint(4) NOT NULL DEFAULT '0' COMMENT '阵印品质',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='阵印';

CREATE TABLE `totem_level_info` (
	`id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键',
	`quality` tinyint(4) NOT NULL DEFAULT '0' COMMENT '阵印品质',
	`level` tinyint(4) NOT NULL DEFAULT '1' COMMENT '等级',
	`health` int(11) NOT NULL COMMENT '生命 - health',
	`attack` int(11) NOT NULL COMMENT '普攻 - attack',
	`defence` int(11) NOT NULL COMMENT '普防 - defence',
	`cultivation` int(11) NOT NULL COMMENT '内力 - cultivation',
	`rock_rune_rate` tinyint(4) NOT NULL COMMENT '分解的石符文概率',
	`rock_rune_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '分解的石符文数量',
	`jade_rune_rate` tinyint(4) NOT NULL COMMENT '分解的玉符文概率',
	`jade_rune_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '分解的玉符文数量',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='阵印等级信息';

CREATE TABLE `totem_call_cost_config` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
	`times` smallint(6) NOT NULL COMMENT '召唤次数',
	`cost` int(11) NOT NULL COMMENT '单价',
	PRIMARY KEY (`id`),
UNIQUE KEY `times` (`times`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='阵印元宝召唤价格';

");
?>

