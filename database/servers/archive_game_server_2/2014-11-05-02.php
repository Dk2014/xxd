<?php

$this->AddSQL("

ALTER TABLE `quest_activity_center` add  COLUMN `is_relative` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否为相对时间';

");

?>