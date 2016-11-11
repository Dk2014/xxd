<?php

$this->AddSQL("
	ALTER TABLE `event_vn_dragon_ball_exchange` ADD COLUMN `sign` varchar(30) COMMENT '资源标识';
", false, array('soha'));
?>