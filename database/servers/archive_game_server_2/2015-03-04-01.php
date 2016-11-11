<?php
$this->AddSQL(
"ALTER TABLE `player_role` ADD COLUMN `status` smallint(6) DEFAULT 0 COMMENT '伙伴状态，0表示正常，1表示在客栈';"
);
?>