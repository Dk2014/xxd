<?php
db_execute($db, "
ALTER TABLE `daily_quest` drop column `type`;
ALTER TABLE `daily_quest` ADD COLUMN `order` int(11) default 0 COMMENT '显示优先级';
");
?>
