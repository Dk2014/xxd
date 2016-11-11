<?php

db_execute($db, "

ALTER TABLE `player_mail_attachment`
MODIFY COLUMN `item_num`  bigint(20) NOT NULL DEFAULT 0 COMMENT '数量' AFTER `item_id`;

");
?>
