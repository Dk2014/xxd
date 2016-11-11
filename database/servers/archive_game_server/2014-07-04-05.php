<?php
db_execute($db, "


drop table if exists `vip_privilege`;
CREATE TABLE `vip_privilege` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`name` varchar(20) NOT NULL COMMENT '特权名称',
	`sign` varchar(20) NOT NULL COMMENT '唯一标识',
	`tip` varchar(200) NOT NULL COMMENT '特权描述',
	PRIMARY KEY (`id`),
UNIQUE KEY `sign` (`sign`)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家VIP特权表';

drop table if exists `vip_privilege_config`;
create table `vip_privilege_config` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`privilege_id` int(11) not null comment '特权ID',
	`level` smallint(6) not null comment 'VIP等级',
	`times` smallint(6) NOT NULL DEFAULT '0' COMMENT '特权次数',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家VIP特权表配置表';

");
?>
