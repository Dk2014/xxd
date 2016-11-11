<?php
db_execute($db, "

ALTER TABLE `ghost`
ADD COLUMN `fragment_id`  smallint(6) NOT NULL COMMENT '对应碎片物品id' AFTER `role_id`;

");
?>