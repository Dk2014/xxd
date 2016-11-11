<?php

$this->AddSQL("alter table `arena_award_box` add column frame int(11) not null default 0 comment '奖励声望'");


$this->AddSQL("create table `frame_system`(
	`id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '功能ID',
	`name` varchar(20) NOT NULL COMMENT '系统名称',
	`sign` varchar(20) NOT NULL COMMENT '唯一标示',
	`max_frame` int(11) NOT NULL DEFAULT '0' COMMENT '最大产出声望',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
");

$this->AddSQL("create table `frame_level`(
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`level` smallint(6) NOT NULL COMMENT '等级',
	`required_frame` int(11) NOT NULL COMMENT '要求声望',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
");


$this->AddSQL("create table `player_frame`(
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`frame` int(11) NOT NULL DEFAULT '0' COMMENT '总声望',
	`level` smallint(6) NOT NULL DEFAULT '1' COMMENT '声望等级',
	`arena_frame` int(11) NOT NULL DEFAULT '0' COMMENT '比武场声望',
	`mult_level_frame` int(11) NOT NULL DEFAULT '0' COMMENT '多人关卡声望',
	PRIMARY KEY (`id`),
KEY `idx_pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
");
?>
