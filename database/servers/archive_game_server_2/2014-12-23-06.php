<?php

$this->AddSQL("

CREATE TABLE `role_friendship` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `friendship_level` tinyint(4) NOT NULL COMMENT '羁绊等级',
  `required_role_level` int(11) NOT NULL COMMENT '所需角色最小等级',
  `favourite_item` smallint(6) DEFAULT NULL COMMENT '喜好品ID',
  `favourite_count` int(11) DEFAULT NULL COMMENT '喜好品需求量',
  `level_color` varchar(20) DEFAULT NULL COMMENT '名称颜色',
  `display_graph` varchar(20) DEFAULT NULL COMMENT '资源标识',
  `relationship_name` varchar(20) DEFAULT NULL COMMENT '羁绊名称',
  `health` int(11) NOT NULL COMMENT '生命 - health',
  `attack` int(11) NOT NULL COMMENT '普攻 - attack',
  `defend` int(11) NOT NULL COMMENT '普防 - defence',
  `cultivation` int(11) NOT NULL COMMENT '内力 - cultivation',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色羁绊数据';

ALTER TABLE `skill` ADD `required_friendship_level` int(11) NOT NULL DEFAULT '0' COMMENT '需要羁绊等级';

ALTER TABLE `player_role` ADD `friendship_level` int(11) NOT NULL DEFAULT '1' COMMENT '角色的羁绊等级';

INSERT INTO `level_func` (`name`, `sign`, `level`) VALUES ('伙伴羁绊', 'LEVEL_FUNC_HUO_BAN_JI_BAN', '28');

");

