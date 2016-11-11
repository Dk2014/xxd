<?php
db_execute($db, "

ALTER TABLE `battle_item_config` ADD COLUMN `can_use_level` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否可以在关卡中使用';

");
?>