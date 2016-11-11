<?php

db_execute($db, "

ALTER TABLE `heart_draw_award` ADD COLUMN `chance` tinyint(4) NOT NULL COMMENT '抽奖概率%';

");
?>
