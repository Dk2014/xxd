
<?php
db_execute($db, "
alter table `hard_level` drop column `mission_id`;
alter table `hard_level` add column `town_id` smallint(6) NOT NULL COMMENT '城镇ID';
alter table `hard_level` add column `hard_level_lock`  int(11) NOT NULL DEFAULT '0' COMMENT '难度关卡权值';
alter table `hard_level` add column `award_hard_level_lock` int(11) not null COMMENT '难度关卡奖励权值'; 
alter table `hard_level` change column `lock` `mission_level_lock` int(11) NOT NULL COMMENT '区域关卡功能权值';


drop table if exists `player_hard_level`;
CREATE TABLE `player_hard_level` (
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`lock` int(11) DEFAULT '0' COMMENT '难度关卡功能权值',
	PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='难度关卡功能权值';


");
?>
