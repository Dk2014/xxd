<?php
db_execute($db, "
alter table  `server_info` add column `event_version` int(11) DEFAULT '0' COMMENT '运营活动数据版本号';
")
?>
