<?php

$this->AddSQL("

ALTER TABLE `skill` add  COLUMN `cheat_id` smallint(6) DEFAULT 0 COMMENT '学习所需秘籍id';

");

?>
