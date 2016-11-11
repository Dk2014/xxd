<?php

$this->AddSQL("

-- 删除高阶灵宠后台数值
delete from `battle_pet` where `pet_id`!=`parent_pet_id`;

-- 删除废弃表格及废弃字段，通过编译错误，清理golang相关代码

alter table `battle_pet`
	drop `parent_pet_id`,
	drop `star`,
	drop `force`,
	drop `health`,
	drop `attack`,
	drop `defence`,
	drop `speed`;

alter table `player_battle_pet_grid`
	drop `level`,
	drop `exp`;

alter table `player_battle_pet`
	drop `star`,
	drop `parent_pet_id`;

drop table `battle_pet_grid_level`;

drop table `battle_pet_grid_attribute`;

drop table `battle_pet_soul_exchange`;

drop table `battle_pet_grid_upgrade_price`;

");

