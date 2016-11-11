<?php

db_execute($db, "
ALTER TABLE `events_qqvip_gift_awards` ADD COLUMN `title` varchar(30) COMMENT '奖励名称';
");

?>