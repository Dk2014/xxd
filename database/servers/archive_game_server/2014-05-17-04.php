<?php
db_execute($db, "

ALTER TABLE `ghost` CHANGE heath health int(11) NOT NULL DEFAULT '0' COMMENT '生命';


");
?>
