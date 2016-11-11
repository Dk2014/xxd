<?php

db_execute($db, "

ALTER TABLE `item` ADD COLUMN `refine_base` int(11) DEFAULT '0' COMMENT '精炼基础值';

");

?>
