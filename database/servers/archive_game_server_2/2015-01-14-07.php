<?php
$this->AddSQL("ALTER TABLE `random_award_box` ADD COLUMN `must_in_first` tinyint(4) DEFAULT '0' COMMENT '是否第一次通关必然获得，0-否 1-是';");

?>