<?php

$this->AddSQL("

ALTER TABLE `event_multiply_config` MODIFY COLUMN `times` float(4,2) COMMENT '加成的倍数';

");

?>