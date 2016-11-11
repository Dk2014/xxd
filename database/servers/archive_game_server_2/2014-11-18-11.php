<?php
$this->AddSQL("
alter table `player_rainbow_level` add column  `buy_timestamp` bigint(20) NOT NULL DEFAULT '0' COMMENT '购买彩虹关卡次数时间戳';
");
?>