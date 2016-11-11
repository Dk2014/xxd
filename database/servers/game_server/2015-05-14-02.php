<?php

$this->AddSQL("
	ALTER TABLE `player_global_clique_info` ADD COLUMN `total_contrib` bigint(20) NOT NULL DEFAULT 0 COMMENT '玩家帮派总贡献';
	UPDATE `player_global_clique_info` SET `total_contrib` = `contrib`;
");
?>