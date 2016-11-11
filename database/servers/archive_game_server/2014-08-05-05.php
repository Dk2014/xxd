<?php
db_execute($db, "
alter table ghost add column `music_sign` varchar(30) DEFAULT NULL COMMENT '音乐资源标识';
");

?>
