<?php
$this->AddSQL("

drop tables if exists `buy_hard_level_times_config`;
   create table `buy_hard_level_times_config` (
       `id` bigint(20) not null auto_increment comment '主键',
       `times` bigint(20) not null default '0' comment '购买次数',
       `cost` bigint(20) not null default '0' comment '购买价格',
       primary key(`id`)
)engine=innodb default charset=utf8mb4 comment='购买深渊关卡次数设置';

drop tables if exists `buy_boss_level_times_config`;
  create table `buy_boss_level_times_config` (
      `id` bigint(20) not null auto_increment comment '主键',
      `times` bigint(20) not null default '0' comment '购买次数',
      `cost` bigint(20) not null default '0' comment '购买价格',
      primary key(`id`)
)engine=innodb default charset=utf8mb4 comment='购买boss关卡次数设置';

    insert into `buy_boss_level_times_config` (`id`, `times`, `cost`) values(1, 1, 10);
    insert into `buy_boss_level_times_config` (`id`, `times`, `cost`) values(2, 2, 20);
    insert into `buy_boss_level_times_config` (`id`, `times`, `cost`) values(3, 3, 30);
    insert into `buy_boss_level_times_config` (`id`, `times`, `cost`) values(4, 4, 40);
    insert into `buy_boss_level_times_config` (`id`, `times`, `cost`) values(5, 5, 50);
    insert into `buy_boss_level_times_config` (`id`, `times`, `cost`) values(6, 6, 60);
    insert into `buy_boss_level_times_config` (`id`, `times`, `cost`) values(7, 7, 70);
    insert into `buy_boss_level_times_config` (`id`, `times`, `cost`) values(8, 8, 80);
    insert into `buy_boss_level_times_config` (`id`, `times`, `cost`) values(9, 9, 90);
    insert into `buy_boss_level_times_config` (`id`, `times`, `cost`) values(10, 10, 100);
    insert into `buy_boss_level_times_config` (`id`, `times`, `cost`) values(11, 11, 110);
    insert into `buy_boss_level_times_config` (`id`, `times`, `cost`) values(12, 12, 120);
    insert into `buy_boss_level_times_config` (`id`, `times`, `cost`) values(13, 13, 130);
    insert into `buy_boss_level_times_config` (`id`, `times`, `cost`) values(14, 14, 140);
    insert into `buy_boss_level_times_config` (`id`, `times`, `cost`) values(15, 15, 150);

    insert into `buy_hard_level_times_config` (`id`, `times`, `cost`) values(1, 1, 10);
    insert into `buy_hard_level_times_config` (`id`, `times`, `cost`) values(2, 2, 20);
    insert into `buy_hard_level_times_config` (`id`, `times`, `cost`) values(3, 3, 30);
    insert into `buy_hard_level_times_config` (`id`, `times`, `cost`) values(4, 4, 40);
    insert into `buy_hard_level_times_config` (`id`, `times`, `cost`) values(5, 5, 50);
    insert into `buy_hard_level_times_config` (`id`, `times`, `cost`) values(6, 6, 60);
    insert into `buy_hard_level_times_config` (`id`, `times`, `cost`) values(7, 7, 70);
    insert into `buy_hard_level_times_config` (`id`, `times`, `cost`) values(8, 8, 80);
    insert into `buy_hard_level_times_config` (`id`, `times`, `cost`) values(9, 9, 90);
    insert into `buy_hard_level_times_config` (`id`, `times`, `cost`) values(10, 10, 100);
    insert into `buy_hard_level_times_config` (`id`, `times`, `cost`) values(11, 11, 110);
    insert into `buy_hard_level_times_config` (`id`, `times`, `cost`) values(12, 12, 120);
    insert into `buy_hard_level_times_config` (`id`, `times`, `cost`) values(13, 13, 130);
    insert into `buy_hard_level_times_config` (`id`, `times`, `cost`) values(14, 14, 140);
    insert into `buy_hard_level_times_config` (`id`, `times`, `cost`) values(15, 15, 150);
");
?>
