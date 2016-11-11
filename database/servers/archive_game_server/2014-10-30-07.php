<?php

db_execute($db, "

alter table battle_pet_grid_exp add column cost_soul_num smallint(6) NOT NULL DEFAULT '1' COMMENT '升级消耗灵魂';	
alter table battle_pet_grid_exp add column   `min_add_exp` bigint(20) NOT NULL DEFAULT '0' COMMENT '最小经验加值';
alter table battle_pet_grid_exp add column   `max_add_exp` bigint(20) NOT NULL DEFAULT '0' COMMENT '最大经验加值';
alter table battle_pet_grid_exp add column `require_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '要求等级';

");
?>
