<?php 
$this->AddSQL("
alter table `driving_sword_treasure_content` add column `award_coins` int(11) not null default '0' comment '奖励铜币';
");
?>
