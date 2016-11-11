<?php
db_execute($db, "

CREATE TABLE `global_tb_xxd_onlinecnt` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `gameappid` varchar(100) NOT NULL COMMENT '平台分配的AppID',
  `timekey` bigint(20) NOT NULL COMMENT '当前时间除以60s，下取整',
  `gsid` bigint(20) NOT NULL COMMENT '游戏服务器编号',
  `onlinecntios` bigint(20) NOT NULL COMMENT 'ios在线人数',
  `onlinecntandroid` bigint(20) NOT NULL COMMENT 'android在线人数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='腾讯经分在线玩家统计日志';


CREATE TABLE `player_tb_xxd_roleinfo` (
  `pid` bigint(20) NOT NULL COMMENT '玩家id',
  `gameappid` varchar(100) NOT NULL COMMENT '平台分配的AppID',
  `openid` varchar(50) NOT NULL COMMENT '玩家平台唯一标识',
  `regtime` bigint(20) NOT NULL COMMENT '注册时间',
  `level` smallint(6) NOT NULL COMMENT '玩家等级',
  `iFriends` smallint(6) NOT NULL COMMENT '玩家好友数',
  `moneyios` bigint(20) NOT NULL COMMENT 'ios金钱存量',
  `moneyandroid` bigint(20) NOT NULL COMMENT 'android金钱存量',
  `diamondios` bigint(20) NOT NULL COMMENT 'ios钻石存量',
  `diamondandroid` bigint(20) NOT NULL COMMENT 'android钻石存量',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='腾讯经分用户信息表';


");

?>
