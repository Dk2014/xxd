<?php
db_execute($db, "

alter table `global_announcement` add column `content` varchar(1024) NOT NULL COMMENT '公共内容，有则忽略模版';


");

?>
