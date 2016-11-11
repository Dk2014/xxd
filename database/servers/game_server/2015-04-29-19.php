<?php

$this->AddSQL("
	ALTER TABLE `clique_building_kongfu` ADD COLUMN `desc` text COMMENT '对应等级描述';
	");


?>
