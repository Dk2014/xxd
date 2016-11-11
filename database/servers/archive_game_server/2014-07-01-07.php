<?php

db_execute($db, "

ALTER TABLE `item` CHANGE COLUMN `appendix_level` `appendix_level` int(11) NOT NULL DEFAULT '0' COMMENT '追加属性等级';


");
?>
