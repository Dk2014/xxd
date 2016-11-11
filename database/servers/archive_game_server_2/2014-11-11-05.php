<?php

$this->AddSQL("
	drop table player_frame;
");

$this->AddSQL("create table `player_frame`(
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`frame` int(11) NOT NULL DEFAULT '0' COMMENT '总声望',
	`level` smallint(6) NOT NULL DEFAULT '1' COMMENT '声望等级',
	`arena_frame` int(11) NOT NULL DEFAULT '0' COMMENT '比武场声望',
	`mult_level_frame` int(11) NOT NULL DEFAULT '0' COMMENT '多人关卡声望',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
");

//给老号初始化数据
$this->AddSQL("insert into `player_frame` (`pid`) select `id` from `player`;");


?>
