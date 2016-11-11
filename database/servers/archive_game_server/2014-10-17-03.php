<?php

db_execute($db, "
ALTER TABLE `quest_activity_center` ADD COLUMN `sign` varchar(40) COMMENT '活动标识' ;
");
?>