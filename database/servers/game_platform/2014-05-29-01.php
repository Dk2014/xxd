<?php
db_execute($db, "

		CREATE TABLE `server` (
		  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '游戏服ID',
		  `type` tinyint(4) NOT NULL COMMENT '平台类型[1-手Q,2-微信]',
		  `name` varchar(50) NOT NULL COMMENT '名称',
		  `status` tinyint(4) NOT NULL COMMENT '状态[0-维护,1-通畅,2-繁忙,3-拥挤]',
		  `is_new` tinyint(4) NOT NULL DEFAULT '0' COMMENT '新服标记[0-否,1-是]',
		  `is_hot` tinyint(4) NOT NULL DEFAULT '0' COMMENT '推荐服[0-否,1-是]',
		  `open_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '开服时间',
		  PRIMARY KEY (`id`),
		  KEY `type` (`type`)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='游戏服配置表';

		CREATE TABLE `account` (
		  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '账号ID',
		  
		  `type` tinyint(4) NOT NULL COMMENT '平台类型[1-手Q,2-微信]',
		  `openid` varchar(50) NOT NULL COMMENT '玩家唯一标识',

		  `sid` int(11) NOT NULL COMMENT '游戏服ID',
		  `nick` varchar(50) NOT NULL COMMENT '玩家昵称',
		  `role_id` tinyint(4) DEFAULT '0' COMMENT '角色模板ID',
		  `login_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '登陆时间',

		  PRIMARY KEY (`id`),
		  KEY `openid` (`openid`),
		  UNIQUE KEY `sid_openid_nick` (`sid`, `openid`, `nick`)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='平台账号';
");
?>