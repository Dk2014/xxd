CREATE TABLE `global_tb_xxd_onlinecnt` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `gameappid` varchar(100) NOT NULL COMMENT '平台分配的AppID',
  `timekey` bigint(20) NOT NULL COMMENT '当前时间除以60s，下取整',
  `gsid` bigint(20) NOT NULL COMMENT '游戏服务器编号',
  `onlinecntios` bigint(20) NOT NULL COMMENT 'ios在线人数',
  `onlinecntandroid` bigint(20) NOT NULL COMMENT 'android在线人数',
  `cid`  bigint(20) NOT NULL DEFAULT 0 COMMENT '渠道ID',
  PRIMARY KEY (`id`),
  KEY `timekey` (`timekey`)
) ENGINE=InnoDB AUTO_INCREMENT=293532 DEFAULT CHARSET=utf8mb4 COMMENT='在线玩家统计日志'   