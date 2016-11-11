<?php
db_execute($db, "
alter table `quest` add column `related_town` smallint(6) NOT NULL DEFAULT '1' COMMENT '任务所属城镇ID，展示用';

");

?>
