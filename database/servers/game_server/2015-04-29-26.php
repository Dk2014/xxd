<?php

$this->AddSQL("

alter table `clique_building_kongfu` add column `cur_kongfu_desc` text COMMENT '当前武功描述';
alter table `clique_building_kongfu` add column `next_kongfu_desc` text COMMENT '下一等级武功描述';
alter table `clique_center_building_level_info` drop column `next_level_desc`;


");


?>
