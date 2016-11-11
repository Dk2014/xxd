<?php
db_execute($db,"
alter table `player_trader_store_state` add column `goods_type` tinyint(4) not null default '0' comment '货物类型0-物品 1-爱心 2-剑心 3-魂侍';
alter table `trader_grid_config` add column `goods_type` tinyint(4) not null default '0' comment '货物类型0-物品 1-爱心 2-剑心 3-魂侍';
"
);
?>
