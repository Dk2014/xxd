<?php

$this->AddSQL("
	ALTER TABLE `clique_building_bank` ADD COLUMN `desc` text COMMENT '对应等级描述';
	ALTER TABLE `clique_center_building_level_info` ADD COLUMN `desc` text COMMENT '对应等级描述';
	");

?>