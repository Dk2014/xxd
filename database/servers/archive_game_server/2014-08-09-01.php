<?php
db_execute($db, "
alter table `player` change column `user` `user` varchar(150) NOT NULL COMMENT '平台传递过来的用户标识';

");

?>
