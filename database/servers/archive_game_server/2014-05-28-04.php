<?php
db_execute($db, "

DROP TABLE IF EXISTS `player_friend`;
CREATE TABLE `player_friend` (
  `id` bigint(20) NOT NULL COMMENT '好友关系ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `friend_pid` bigint(20) NOT NULL COMMENT '好友ID',
  `friend_nick` varchar(50) NOT NULL DEFAULT '' COMMENT '玩家昵称',
  `friend_role_id` tinyint(4) NOT NULL COMMENT '好友角色ID',
  `friend_mode` tinyint(4) NOT NULL COMMENT '好友关系:0陌生人,1仅关注,2仅被关注,3互相关注(好友)',
  `last_chat_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后聊天时间',
  `friend_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '成为好友时间',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家好友列表';

DROP TABLE IF EXISTS `player_friend_chat`;
CREATE TABLE `player_friend_chat` (
  `id` bigint(20) NOT NULL COMMENT '消息ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `friend_pid` bigint(20) NOT NULL COMMENT '对方玩家ID',
  `mode` tinyint(4) NOT NULL COMMENT '1发送，2接收',
  `send_time` bigint(20) NOT NULL COMMENT '发送时间戳',
  `message` varchar(140) NOT NULL COMMENT '消息内容',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家聊天记录';
");
?>