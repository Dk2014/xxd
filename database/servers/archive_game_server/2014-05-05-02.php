<?php

db_execute($db, "
DROP TABLE IF EXISTS `player`;

CREATE TABLE `player` (
  `id` bigint(20) NOT NULL COMMENT '玩家ID',
  `user` varchar(150) NOT NULL COMMENT '平台传递过来的用户标识',
  `nick` varchar(50) NOT NULL COMMENT '玩家昵称',
  PRIMARY KEY (`id`),
  KEY `ix_player_sign` (`user`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家基础信息';

");
?>