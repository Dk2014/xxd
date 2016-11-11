<?php
db_execute($db,"
alter table `player_info` add column `init_global_srv` tinyint(4) not null default '0' comment '是否在互动服已初始化. 0 - 没有';
"
);
?>

