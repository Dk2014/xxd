<?php
db_execute($db, "

alter table `mail` add column `auto_delete` tinyint(4) default 0 not null comment '自动删除世纪 0-过期删除 1-无附件已阅读删除';

");

?>
