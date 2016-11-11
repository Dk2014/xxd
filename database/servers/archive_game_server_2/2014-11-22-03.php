<?php
$this->AddSQL("
alter table `ghost` add column  `production_info` varchar(50) DEFAULT NULL COMMENT '产出描述';
");
?>