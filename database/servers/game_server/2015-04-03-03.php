<?php

$this->AddSQL("ALTER TABLE `player_ghost` ADD COLUMN `relation_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '连锁玩家魂侍ID';");

?>