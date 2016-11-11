<?php 
$this->AddSQL("
alter table `item` modify column `can_sell` tinyint(4) not null default '1' comment '是否可以出售';
");
?>
