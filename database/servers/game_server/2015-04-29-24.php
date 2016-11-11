<?php

$this->AddSQL("

alter table `player_global_clique_info` change column  `donate_coins_attck_building` `donate_coins_attack_building` bigint(20) NOT NULL DEFAULT '0' COMMENT '神兵堂累积贡献铜币';

alter table `global_clique` change column `attck_building_coins` `attack_building_coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '神兵堂当前贡献铜币';

alter table `global_clique` change column `attck_building_level` `attack_building_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '神兵堂当前等级';





");


?>
