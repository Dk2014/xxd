<?php

$this->AddSQL("

alter table `clique_center_building_level_info` add column `cur_level_desc` text COMMENT '当前总舵描述';
alter table `clique_center_building_level_info` add column `next_level_desc` text COMMENT '下一等级总舵描述';

alter table `clique_building_bank` add column `cur_level_desc` text COMMENT '当前钱庄描述';
alter table `clique_building_bank` add column `next_level_desc` text COMMENT '下一等级钱庄描述';


");


?>
