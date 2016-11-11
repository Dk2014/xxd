<?php

db_execute($db, "
CREATE TABLE IF NOT EXISTS `player_activity`(
    `pid` bigint(20) COMMENT '用户ID',
    `activity` int(11) DEFAULT 0 COMMENT '用户活跃度',
    `last_update` bigint(20) DEFAULT 0 COMMENT '最后一次更新时间戳',
    PRIMARY KEY(`pid`)
    )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家活跃度';
");

?>
