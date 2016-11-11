<?php
db_execute($db,"
alter table `mission_talk` add column `town_id` smallint(6) NOT NULL COMMENT '城镇ID';
alter table `mission_talk` add column `quest_id` smallint(6) NOT NULL COMMENT '任务';
"
);
?>
