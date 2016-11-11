<?php
$this->AddSQL("
    drop table if exists `player_add_ingot_record`;
    CREATE TABLE `player_add_ingot_record` (
      `pid` bigint(20) AUTO_INCREMENT NOT NULL COMMENT '玩家PID',
      `ingot` bigint(20) NOT NULL COMMENT '元宝数',
      `timestamp` bigint(20) NOT NULL COMMENT '时间戳',
    PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='充值排行榜期内玩家充值记录';
");
?>
