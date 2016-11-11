<?php 
$this->AddSQL("
alter table quest add column `pass_mission_share` tinyint(4) not null default '0' comment '通关区域分享';
");
?>
