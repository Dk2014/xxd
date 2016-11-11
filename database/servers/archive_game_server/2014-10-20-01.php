<?php

db_execute($db, "
DROP TABLE IF EXISTS `player_event_award_record`;
CREATE TABLE `player_event_award_record`(
    `pid` bigint(20) NOT NULL COMMENT '用户ID',
    `record_bytes` mediumblob  COMMENT '奖励领取状态',
    PRIMARY KEY(`pid`)
  )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家活动奖励领取记录';
");
?>