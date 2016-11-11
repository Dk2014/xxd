create table `gift_code` (
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`version` bigint(20) NOT NULL DEFAULT '0' COMMENT '兑换码版本号',
	`code` char(10) NOT NULL  COMMENT '兑换码',
	`type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型 0--可被多个玩家分别使用多次 1--用完后其他玩家不能使用',
	`server_id` int(11) NOT NULL COMMENT '服务器ID 0--所有服务器可以用 其他--指定服务器可用',
	`effect_timestamp` bigint(20) NOT NULL DEFAULT '0' COMMENT '生效时间',
	`expire_timestamp` bigint(20) NOT NULL DEFAULT '0' COMMENT '过期时间',
	`item_id`  smallint(6) NOT NULL COMMENT '兑换物品',
	`disable`  tinyint(4) NOT NULL DEFAULT '0' COMMENT '取消兑换码',
	`content`	varchar(1024) comment '兑换码使用后邮件标题',
	PRIMARY KEY (`id`),
	UNIQUE KEY `code_idx` (`server_id`, `code`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='兑换码存储表';

-- [{"item_type":1, "item_id":2, "num":10}]

alter table gift_code add column `config` text not null comment "兑换码配置";