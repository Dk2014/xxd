<?php

$this->AddSQL("
ALTER TABLE `player_role` ADD COLUMN `timestamp` bigint(20) NOT NULL DEFAULT '0' COMMENT '等级变化时间';
");

?>
