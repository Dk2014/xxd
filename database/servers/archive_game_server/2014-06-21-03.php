<?php

db_execute($db, "

ALTER TABLE `item`
MODIFY COLUMN `appendix_level`  int(11) NOT NULL COMMENT '追加属性等级' AFTER `appendix_num`;

");
?>

