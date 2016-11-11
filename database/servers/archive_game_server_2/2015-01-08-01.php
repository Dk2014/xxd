<?php

$this->AddSQL("
ALTER TABLE `chest` ADD COLUMN `fix_award_count` SMALLINT(6) NOT NULL DEFAULT '0' COMMENT '固奖励';
");

