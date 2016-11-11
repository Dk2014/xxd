<?php
db_execute($db, "
alter table `platform_friend_award` add column `name` varchar(30) NOT NULL DEFAULT '' COMMENT '奖励名称' after `id`;
");
?>


