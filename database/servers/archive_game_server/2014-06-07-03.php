<?php
db_execute($db, 
"
ALTER TABLE `player_info` ADD COLUMN `last_login_time` bigint(20) NOT NULL COMMENT '上次登录时间';
ALTER TABLE `player_info` ADD COLUMN `last_offline_time` bigint(20) NOT NULL COMMENT '上次下线时间';
ALTER TABLE `player_info` ADD COLUMN `total_online_time` bigint(20) NOT NULL COMMENT '总在线时间';
"

);
?>