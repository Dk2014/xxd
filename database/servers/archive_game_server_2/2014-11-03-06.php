<?php


$this->AddSQL("
TRUNCATE TABLE `event_multiply_config`;
ALTER TABLE `event_multiply_config` MODIFY COLUMN `condition` int(11) COMMENT '加成的事件id'; 
ALTER TABLE `event_multiply_config` MODIFY COLUMN `times` decimal(4,2) COMMENT '加成的倍数';
");


?>