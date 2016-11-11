<?php

db_execute($db, "

ALTER TABLE `skill` ADD COLUMN `music_sign` varchar(30) DEFAULT NULL COMMENT '音乐资源标识' AFTER `sign`;

");
?>