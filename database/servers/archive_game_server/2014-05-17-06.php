<?php
db_execute($db, "

ALTER TABLE `mission_enemy` change `boss_id` `is_boss` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否是boss 0否,1是';

");
?>
