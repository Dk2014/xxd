<?php

$this->AddSQL("ALTER TABLE `player_extend_level` drop column  `max_level` ;");
$this->AddSQL("alter table `player_extend_level` add column `exp_maxlevel` smallint(6) NOT NULL DEFAULT '0' COMMENT '经验关卡通关了的最大等级';");
$this->AddSQL("alter table `player_extend_level` add column `coins_maxlevel` smallint(6) NOT NULL DEFAULT '0' COMMENT '经验关卡通关了的最大等级';");
$this->AddSQL("drop index item_id_idx on player_item;");

?>

