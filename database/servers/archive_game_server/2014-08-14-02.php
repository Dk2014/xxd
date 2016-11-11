<?php
db_execute($db, "
alter table `player_mail` add column `title` varchar(30) NOT NULL DEFAULT '' COMMENT '标题';

alter table `player_mail` add column `content` varchar(1024) NOT NULL DEFAULT '' COMMENT '内容';
");

?>
