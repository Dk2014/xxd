CREATE TABLE `payment_log` (
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`order_id` varchar(50) NOT NULL COMMENT 'AnySDK订单号',
	`product_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '购买数量',
	`amount` varchar(128) NOT NULL DEFAULT '' COMMENT '支付金额',
	`pay_status` char(10) NOT NULL  COMMENT '支付状态',
	`pay_time` char(40) NOT NULL  COMMENT '支付时间YYYY-mm-dd HH:ii:ss',
	`user_id` varchar(128) NOT NULL  COMMENT '渠道用户ID',
	`order_type` char(10) NOT NULL  COMMENT '支付方式',
	`game_user_id` bigint(20) NOT NULL COMMENT 'XXD游戏内部PID',
	`server_id` bigint(20) NOT NULL COMMENT 'XXD游戏内部游戏进程ID',
	`product_name` char(128) NOT NULL  COMMENT '渠道产品名称',
	`product_id` varchar(128) NOT NULL  COMMENT '渠道产品ID',
	`private_data` varchar(128) NOT NULL  DEFAULT '' COMMENT '私有数据',
	`channel_number` varchar(128) NOT NULL  COMMENT '渠道编号',
	`sign` char(128) NOT NULL  COMMENT '加密签名',
	`source` text COMMENT '渠道请求anysdk参数',
	`enhanced_sign` char(128) NOT NULL  COMMENT '增强加密签名',
	PRIMARY KEY (`id`),
	UNIQUE KEY `order_id` (`order_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='anysdk支付通知日志';

CREATE TABLE `pending_queue` (
	`id` bigint(20) NOT NULL COMMENT 'ID',
	`try_timestamp` bigint(20) NOT NULL COMMENT '最近尝试时间戳',
	`try` tinyint(4) NOT NULL DEFAULT '0' COMMENT '尝试次数',
	`product_id` varchar(128) NOT NULL  COMMENT '渠道产品ID',
	`game_user_id` bigint(20) NOT NULL COMMENT 'XXD游戏内部PID',
	`amount` varchar(128) NOT NULL COMMENT '游戏充值金额',
	`ip` varchar(64) NOT NULL  DEFAULT '' COMMENT 'IP',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='anysdk支付待处理队列';

CREATE TABLE `failed_record` (
	`id` bigint(20) NOT NULL COMMENT 'ID',
	`fail_time` tinyint(4) NOT NULL DEFAULT '0' COMMENT '尝试次数',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='anysdk支付待失败纪录队列';

alter table pending_queue add column `public` tinyint(4) NOT NULL DEFAULT 0 COMMENT '1-正式订单 2-测试环境订单';

create table `app_store_payment_log` (
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`receipt_hash` char(32) NOT NULL COMMENT 'ios 收据 md5 hash',
	`receipt` text COMMENT 'ios 收据',
	`status` tinyint(8) NOT NULL  DEFAULT '0' COMMENT '0--等待验证 1--验证失败 2--验证成功',
	`ip` varchar(64) NOT NULL  DEFAULT '' COMMENT 'IP',
	`game_user_id` bigint(20) NOT NULL COMMENT 'XXD游戏内部PID',
	`nickname` varchar(32) NOT NULL DEFAULT "" COMMENT '玩家昵称',
	`openid` varchar(32) NOT NULL DEFAULT "" COMMENT '玩家平台openid',

	`try_timestamp` bigint(20) NOT NULL DEFAULT '0' COMMENT '最近尝试时间戳',
	`try` tinyint(4) NOT NULL DEFAULT '0' COMMENT '尝试次数',
	`transaction_id` varchar(32) NOT NULL DEFAULT '0' COMMENT '事务ID',
	`product_id` varchar(128) NOT NULL  DEFAULT '' COMMENT '渠道产品ID',
	`quantity` int(11) NOT NULL  DEFAULT '1' COMMENT '购买数量',
	`purchase_date_ms` bigint(20) NOT NULL  DEFAULT 0 COMMENT '购买时间戳(ms)',
	`verify_result` text COMMENT 'ios 验证返回',
PRIMARY KEY (`id`),
UNIQUE KEY `receipt_hash` (`receipt_hash`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='app store payment log';

create table `app_store_pending_queue` (
	`id` bigint(20) NOT NULL COMMENT 'ID',
	`try_timestamp` bigint(20) NOT NULL COMMENT '最近尝试时间戳',
	`try` tinyint(4) NOT NULL DEFAULT '0' COMMENT '尝试次数',
	`product_id` varchar(128) NOT NULL  COMMENT '渠道产品ID',
	`receipt_hash` char(32) NOT NULL COMMENT 'ios 收据 md5 hash',
	`transaction_id` varchar(32) NOT NULL DEFAULT '0' COMMENT '事务ID',
	`purchase_date_ms` bigint(20) NOT NULL  DEFAULT 0 COMMENT '购买时间戳(ms)',
	`game_user_id` bigint(20) NOT NULL COMMENT 'XXD游戏内部PID',
	`nickname` varchar(32) NOT NULL DEFAULT "" COMMENT '玩家昵称',
	`openid` varchar(32) NOT NULL DEFAULT "" COMMENT '玩家平台openid',
	`quantity` int(11) NOT NULL  DEFAULT '1' COMMENT '购买数量',
	`ip` varchar(64) NOT NULL  DEFAULT '' COMMENT 'IP',
	`is_delivered` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否已处理',
PRIMARY KEY (`id`),
UNIQUE KEY `receipt_hash` (`receipt_hash`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='app store payment deliver queue';


-- TODO create index on status
create table `google_play_payment_log` (
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`token_hash` char(32) NOT NULL COMMENT 'google play token md5 hash',
	`token` text COMMENT 'google 支付token',
	`status` tinyint(8) NOT NULL  DEFAULT '0' COMMENT '0--等待验证 1--验证失败 2--验证成功',
	`ip` varchar(64) NOT NULL  DEFAULT '' COMMENT 'IP',
	`game_user_id` bigint(20) NOT NULL COMMENT 'XXD游戏内部PID',
	`nickname` varchar(32) NOT NULL DEFAULT "" COMMENT '玩家昵称',
	`openid` varchar(32) NOT NULL DEFAULT "" COMMENT '玩家平台openid',
	`transaction_id` varchar(32) NOT NULL DEFAULT '0' COMMENT '事务ID',
	`product_id` varchar(128) NOT NULL  DEFAULT '' COMMENT '渠道产品ID',
	`quantity` int(11) NOT NULL  DEFAULT '1' COMMENT '购买数量',
	`purchase_date_ms` bigint(20) NOT NULL  DEFAULT 0 COMMENT '购买时间戳(ms)',
	`origin_req` text  COMMENT '客户端原始请求',
PRIMARY KEY (`id`),
UNIQUE KEY `token_hash` (`token_hash`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='google play payment log';

create table `google_play_pending_queue` (
	`id` bigint(20) NOT NULL COMMENT 'ID',
	`try_timestamp` bigint(20) NOT NULL COMMENT '最近尝试时间戳',
	`try` tinyint(4) NOT NULL DEFAULT '0' COMMENT '尝试次数',
	`product_id` varchar(128) NOT NULL  COMMENT '渠道产品ID',
	`token_hash` char(32) NOT NULL COMMENT 'google play token md5 hash',
	`transaction_id` varchar(32) NOT NULL DEFAULT '0' COMMENT '订单ID',
	`purchase_date_ms` bigint(20) NOT NULL  DEFAULT 0 COMMENT '购买时间戳(ms)',
	`game_user_id` bigint(20) NOT NULL COMMENT 'XXD游戏内部PID',
	`nickname` varchar(32) NOT NULL DEFAULT "" COMMENT '玩家昵称',
	`openid` varchar(32) NOT NULL DEFAULT "" COMMENT '玩家平台openid',
	`quantity` int(11) NOT NULL  DEFAULT '1' COMMENT '购买数量',
	`ip` varchar(64) NOT NULL  DEFAULT '' COMMENT 'IP',
	`is_delivered` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否已被正确处理',
PRIMARY KEY (`id`),
UNIQUE KEY `token_hash` (`token_hash`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='google play payment deliver queue';

create table `wegames_payment_log` (
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`order_id` varchar(128) NOT NULL COMMENT '外部订单ID',
	`platform_uid` varchar(128) NOT NULL COMMENT '玩家openid',
	`pay_amount` varchar(12) NOT NULL COMMENT '支付金额(当地支付货币金额)',
	`twd_money` varchar(12) NOT NULL COMMENT '支付金额(等价台币支付金额，以此为准)',
	`game_code` varchar(32) NOT NULL DEFAULT '' COMMENT '游戏代码',
	`server_code` int(11) NOT NULL COMMENT '游戏区ID',
	`game_money` bigint(20) NOT NULL COMMENT '购买获得游戏币',
	`present_game_money` bigint(20) NOT NULL COMMENT '赠送获得游戏币',
	`other_item` varchar(32) NOT NULL DEFAULT '' COMMENT '项目用途', 
	`virtual_items` text COMMENT '游戏内道具 json 格式', 
	`time` varchar(32) NOT NULL DEFAULT '' COMMENT '购买时间戳', 
	`sign` char(32) NOT NULL COMMENT '加密签名', 
UNIQUE KEY `order_id` (`order_id`),
PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='wegames payment log';

create table `wegames_pending_queue` (
	`id` bigint(20) NOT NULL COMMENT 'ID',
	`try_timestamp` bigint(20) NOT NULL COMMENT '最近尝试时间戳',
	`try` tinyint(4) NOT NULL DEFAULT '0' COMMENT '尝试次数',
	`order_id` varchar(128) NOT NULL COMMENT '外部订单ID',
	`platform_uid` varchar(128) NOT NULL COMMENT '玩家openid',
	`server_code` int(11) NOT NULL COMMENT '游戏区ID',
	`twd_money` varchar(12) NOT NULL COMMENT '支付金额(台币)',
	`game_money` bigint(20) NOT NULL COMMENT '购买获得游戏币',
	`present_game_money` bigint(20) NOT NULL COMMENT '赠送获得游戏币',
	`other_item` varchar(32) NOT NULL DEFAULT '' COMMENT '项目用途', 
	`virtual_items` text COMMENT '游戏内道具 json 格式', 
	`time` varchar(32) NOT NULL DEFAULT '' COMMENT '购买时间戳', 
PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='wegames payment deliver queue';
