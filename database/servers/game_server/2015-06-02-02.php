
<?php

$this->AddSQL("
        alter table `player_ghost` add column `add_growth` smallint(6) not null default '0' comment '魂侍洗练添加的成长值';

        create table `ghost_baptize` (
            `id` tinyint(4) not null auto_increment comment '魂侍洗炼ID',
            `star` tinyint(4) not null comment '魂侍星级',
            `max_add_growth` tinyint(4) not null default '0' comment '最大增加成长值',
            `probablity1` tinyint(4) not null default '0' comment '概率1',
            `min_add_growth1` tinyint(4) not null comment '最小随机添加成长值1',
            `max_add_growth1` tinyint(4) not null comment '最大随机添加成长值1',
            `probablity2` tinyint(4) not null default '0' comment '概率2',
            `min_add_growth2` tinyint(4) not null comment '最小随机添加成长值2',
            `max_add_growth2` tinyint(4) not null comment '最大随机添加成长值2',
            primary key (`id`),
            unique key `star` (`star`)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='魂侍洗炼表';
");

?>
