<?php

$this->AddSQL("

CREATE TABLE `fate_box` (
	`id` INT (11) NOT NULL AUTO_INCREMENT COMMENT 'id'
	,`name` VARCHAR(20) NOT NULL comment '命锁名称'
	,`sign` VARCHAR(20) NOT NULL comment '标识符'
	,`level` SMALLINT (6) NOT NULL comment '要求等级'
	,`require_lock` INT (11) NOT NULL comment '命锁宝箱权值'
	,`award_lock` INT (11) NOT NULL comment '奖励命锁宝箱权值'
  	,UNIQUE KEY `sign` (`sign`)
	,PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '命锁宝箱';

CREATE TABLE `player_fate_box_state` (
	`pid` BIGINT (20) NOT NULL COMMENT 'pid'
	,`lock` INT (11) NOT NULL DEFAULT '0' COMMENT '玩家命锁权值'
	,`star_box_free_draw_timestamp` BIGINT (20) NOT NULL DEFAULT '0' COMMENT '星辉命匣免费抽取时间戳'
	,`star_box_draw_count` INT (11) NOT NULL DEFAULT '0' COMMENT '星辉宝箱抽奖次数'
	,`moon_box_free_draw_timestamp` BIGINT (20) NOT NULL DEFAULT '0' COMMENT '月影命匣免费抽取时间戳'
	,`moon_box_draw_count` INT (11) NOT NULL DEFAULT '0' COMMENT '月影宝箱抽奖次数'
	,`sun_box_free_draw_timestamp` BIGINT (20) NOT NULL DEFAULT '0' COMMENT '日耀命匣免费抽取时间戳'
	,`sun_box_draw_count` INT (11) NOT NULL DEFAULT '0' COMMENT '日耀宝箱抽奖次数'
	,KEY `idx_pid`(`pid`)
	,PRIMARY KEY (`pid`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '玩家命锁宝箱状态';

DELETE FROM `chest`;
DELETE FROM `chest_item`;

INSERT INTO `func` (`name`, `sign`, `lock`, `level`, `unique_key`, `need_play`, `type`) 
	values ('命锁宝箱', 'FUNC_FATE_BOX', '1400', '0', '1048576', '1', '1');

DELETE FROM  `func` where `name`='神龙宝箱';

INSERT INTO `fate_box` (`id`, `name`, `level`, `require_lock`, `award_lock`, `sign`)
	VALUES
	(1,'星辉宝箱',20,0,1000,'StarBox'),
	(2,'月影宝箱',50,1000,2000,'MoonBox'),
	(3,'日耀宝箱',999,999999,0,'SunBox');


");

