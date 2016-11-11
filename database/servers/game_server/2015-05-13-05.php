<?php
//帮派送镖个人信息
$this->AddSQL("
alter table `clique_boat` add column `color` varchar(10) default '' comment '颜色';
");
