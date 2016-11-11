<?php
db_execute($db, 
"

alter table `mission_level` add column `parent_id` smallint(6) DEFAULT '0' COMMENT '关联关卡的外键' after `mission_id`;		

CREATE TABLE `resource_level` (
	`id` smallint(6) NOT NULL AUTO_INCREMENT,
	`max_level` smallint(6) NOT NULL COMMENT '允许开放的等级上限',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资源关卡';


CREATE TABLE `player_resource_level_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `max_level` smallint(6) NOT NULL COMMENT '允许开放的等级上限',
  `level_id` int(11) NOT NULL COMMENT '关卡ID',
  `daily_num` tinyint(4) NOT NULL COMMENT '当日已进入关卡的次数',
  `last_enter_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后一次进入时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家资源关卡记录';

"
);

?>