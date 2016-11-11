<?php

db_execute($db, "
CREATE TABLE IF NOT EXISTS `player_month_card_award_record`(
    `pid` bigint(20) COMMENT '用户ID',
    `last_update` bigint(20) DEFAULT 0 COMMENT '最后一次更新时间戳' NOT NULL,
    PRIMARY KEY(`pid`)
    )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家最后领取月卡时间';
");

?>