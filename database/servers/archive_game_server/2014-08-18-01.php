<?php
db_execute($db, "

alter table `mail` change column `auto_delete` `expire_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '邮件删除时机 0-默认过期删除 1-无附件已阅读自动删除 >1 指定时间删除';

alter table `player_mail` add column `expire_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '邮件删除时机';

");

?>
