<?php

db_execute($db, "

ALTER TABLE `mission_level` ADD COLUMN `sub_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '关卡子类型(0--无;1--铜钱关卡;2--经验关卡)';

");
?>