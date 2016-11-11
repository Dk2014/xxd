<?php
db_execute($db, 
"
ALTER TABLE `player_friend_state` ADD COLUMN `exist_offline_chat` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0没有离线消息，1有离线消息';

"

);
?>