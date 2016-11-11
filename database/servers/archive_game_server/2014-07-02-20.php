<?php
db_execute($db, "

ALTER TABLE `ghost`
CHANGE COLUMN `defense` `defence`  int(11) NOT NULL DEFAULT 0 COMMENT '防御' AFTER `attack`;

");
?>
