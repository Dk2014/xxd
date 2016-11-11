<?php

$this->AddSQL("ALTER TABLE `server_info` ADD COLUMN `event_json_info` text COMMENT '活动配置信息';");

?>