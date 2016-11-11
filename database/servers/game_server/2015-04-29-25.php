<?php

$this->AddSQL("

alter table `clique_center_building_level_info` add column `next_level_desc` text COMMENT '下一等级描述';
alter table `clique_kongfu` change column   `building` `building` int(11) NOT NULL COMMENT '所属建筑';


");


?>
