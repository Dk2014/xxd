<?php

// pr2版本玩家数据升级

// 全局通告
$this->AddSQL("
alter table player_mail add column priority tinyint(4) NOT NULL DEFAULT '0' COMMENT '优先级';
alter table global_mail add column priority tinyint(4) NOT NULL DEFAULT '0' COMMENT '优先级';  
alter table `global_announcement` add column `spacing_time` bigint(20) DEFAULT 0 NOT NULL COMMENT '间隔时间';
");

// 增加活动中心表
$this->AddSQL("CREATE TABLE `player_activity` (
  `pid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `activity` int(11) DEFAULT '0' COMMENT '用户活跃度',
  `last_update` bigint(20) DEFAULT '0' COMMENT '最后一次更新时间戳',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家活跃度';");

// 增加灵宠表
$this->AddSQL("
alter table `player_battle_pet` drop column `ball_num`;
alter table `player_battle_pet` drop column `activated`;
alter table `player_battle_pet` add column `star` tinyint(4) NOT NULL DEFAULT '1' COMMENT '星级[1,5]';
alter table `player_battle_pet` add column `parent_pet_id` int(11) NOT NULL DEFAULT '0' COMMENT '父灵宠ID';

create table `player_battle_pet_grid` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `pet_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '灵宠ID',
  `grid_id` tinyint(4) NOT NULL COMMENT '灵宠格子ID',
  `level` smallint(6) NOT NULL DEFAULT '1' COMMENT '格子等级',
  `exp` int(11) NOT NULL DEFAULT '0' COMMENT '格子经验',
  PRIMARY KEY (`id`),
KEY `idx_pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家灵宠';


drop table `player_battle_pet_config`;
alter table player_battle_pet modify column parent_pet_id int(11) NOT NULL COMMENT '父灵宠ID(怪物ID)';
alter table player_battle_pet modify column battle_pet_id int(11) NOT NULL COMMENT '灵宠ID(怪物ID）';
alter table player_battle_pet_grid drop column pet_id;
alter table player_battle_pet_grid add column battle_pet_id int(11) NOT NULL DEFAULT '0' COMMENT '灵宠ID';
alter table player_battle_pet_grid modify column   `exp` bigint(20) NOT NULL DEFAULT '0' COMMENT '格子经验';
  ");

// 增加玩家活动信息表
$this->AddSQL("CREATE TABLE `player_event_award_record` (
  `pid` bigint(20) NOT NULL COMMENT '用户ID',
  `record_bytes` mediumblob COMMENT '奖励领取状态',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家活动奖励领取记录';");


$this->AddSQL("CREATE TABLE `player_meditation_state` (
  `pid` bigint(20) NOT NULL,
  `accumulate_time` int(11) NOT NULL COMMENT '光明钥匙奖励累积时间',
  `start_timestamp` bigint(20) NOT NULL DEFAULT '0' COMMENT '打坐开始时间 0-未未打坐状态',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家打坐状态';");


$this->AddSQL("CREATE TABLE `player_month_card_award_record` (
  `pid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `last_update` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后一次更新时间戳',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家最后领取月卡时间';");


$this->AddSQL("CREATE TABLE `player_month_card_info` (
  `pid` bigint(20) NOT NULL COMMENT '用户ID',
  `starttime` bigint(20) NOT NULL DEFAULT '0' COMMENT '月卡开始时间',
  `endtime` bigint(20) NOT NULL DEFAULT '0' COMMENT '月卡结束时间',
  `buytimes` bigint(10) NOT NULL DEFAULT '0' COMMENT '购买次数',
  `presenttotal` bigint(10) NOT NULL DEFAULT '0' COMMENT '赠送总金额',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家月卡信息';");

$this->AddSQL("CREATE TABLE `player_push_notify_switch` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `options` bigint(20) NOT NULL DEFAULT '0' COMMENT '推送通知开关',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家推送通知开关列表';");

?>