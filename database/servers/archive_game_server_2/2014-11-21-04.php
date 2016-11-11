<?php
$this->AddSQL("
alter table `daily_quest` add column  `award_ingot` int(11) NOT NULL DEFAULT '0' COMMENT '奖励元宝';
");
?>