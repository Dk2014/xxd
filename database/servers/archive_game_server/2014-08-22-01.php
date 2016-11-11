<?php
db_execute($db, "
alter table `quest` add column `mission_drama_talk` varchar(1024) DEFAULT '' COMMENT '区域关卡剧情对话';

");

?>
