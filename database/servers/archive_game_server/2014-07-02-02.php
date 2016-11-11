<?php
db_execute($db, "

DROP TABLE `ghost`;
CREATE TABLE `ghost` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '魂侍名称',
  `sign` varchar(30) NOT NULL DEFAULT '' COMMENT '资源标识',
  `town_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '城镇id（影界id）',
  `role_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '可装备角色id',
  `unique_key` smallint(6) NOT NULL DEFAULT '0' COMMENT '每个影界中魂侍的唯一标记',
  `init_star` tinyint(4) NOT NULL DEFAULT '0' COMMENT '初始星级',
  `health` int(11) NOT NULL DEFAULT '0' COMMENT '生命',
  `attack` int(11) NOT NULL DEFAULT '0' COMMENT '攻击',
  `defense` int(11) NOT NULL DEFAULT '0' COMMENT '防御',
  `speed` int(11) NOT NULL DEFAULT '0' COMMENT '速度',
  `desc` varchar(300) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='魂侍';


DROP TABLE `ghost_exchange`;
DROP TABLE `ghost_item_exchange`;

CREATE TABLE `ghost_star` (
  `id` bigint(20) NOT NULL COMMENT '主键',
  `star` tinyint(4) NOT NULL DEFAULT '0' COMMENT '星级',
  `need_fragment_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '需要碎片数量',
  `growth` smallint(6) NOT NULL DEFAULT '0' COMMENT '成长值',
  `color` tinyint(4)  NOT NULL DEFAULT '0' COMMENT '颜色',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='魂侍进阶表';

DROP TABLE `player_ghost`;
CREATE TABLE `player_ghost` (
  `id` bigint(20) NOT NULL COMMENT '主键',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `ghost_id` smallint(6) NOT NULL COMMENT '魂侍ID',
  `star` tinyint(4) NOT NULL DEFAULT '0' COMMENT '星级',
  `level` smallint(6) NOT NULL DEFAULT '1' COMMENT '魂侍等级',
  `exp` bigint(20) NOT NULL DEFAULT '0' COMMENT '魂侍经验',
  `pos` smallint(6) NOT NULL COMMENT '位置',
  `is_new` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否新魂侍',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家魂侍表';

DROP TABLE `player_ghost_equipment`;
CREATE TABLE `player_ghost_equipment` (
  `id` bigint(20) NOT NULL COMMENT '主键',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `ghost_power` tinyint(4) NOT NULL COMMENT '魂力',
  `pos1` bigint(20) NOT NULL COMMENT '装备位置1的魂侍id',
  `pos2` bigint(20) NOT NULL COMMENT '装备位置2的魂侍id',
  `pos3` bigint(20) NOT NULL COMMENT '装备位置3的魂侍id',
  `pos4` bigint(20) NOT NULL COMMENT '装备位置4的魂侍id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家魂侍装备表';

DROP TABLE `player_ghost_state`;
CREATE TABLE `player_ghost_state` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `purify_day_count` bigint(20) DEFAULT '0' COMMENT '每日净化次数',
  `ghost_unique_key` smallint(6) NOT NULL DEFAULT '0' COMMENT '获得金魂的信息',
  `purify_update_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '净化时间',
  `ghost_mission_key` int(11) NOT NULL DEFAULT '0' COMMENT '开启影界最大权值',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家魂侍状态表';

");
?>
