<?php

$this->AddSQL("

ALTER TABLE `skill` add  COLUMN `auto_learn_level` tinyint(4) NOT NULL DEFAULT 1 COMMENT '是否达到等级习得';

");

?>