<?php
db_execute($db, "
alter table `item` modify column `show_mode` tinyint(4) NOT NULL DEFAULT '0' COMMENT '物品使用后表现形式 0-无 1-单体 2-全体 3-纵向 4-横向';
")
?>
