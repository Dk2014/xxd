<?php

db_execute($db, "
ALTER TABLE `quest_activity_center` ADD COLUMN `mail_title` varchar(60) COMMENT '补发奖励邮件标题';
ALTER TABLE `quest_activity_center` ADD COLUMN `mail_content` text COMMENT '补发奖励邮件内容,{val}对应权值';
");
?>
