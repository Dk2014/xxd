<?php


$this->AddSQL("
alter table `player_vip` add column  `present_exp` bigint(20) DEFAULT 0 NOT NULL COMMENT '赠送vip经验';
");

?>
