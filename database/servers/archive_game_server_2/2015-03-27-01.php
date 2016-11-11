<?php 
$this->AddSQL("
alter table `item` add column can_sell tinyint(4) not null default 0 comment '是否可以出售';
");
?>
