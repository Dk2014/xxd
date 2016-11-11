<?php

db_execute($db, "

ALTER TABLE `sword_soul`
ADD COLUMN `sign`  varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '资源标识' AFTER `name`;

");
?>
