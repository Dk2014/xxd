<?php
db_execute($db,"

drop table if exists `trader_refresh_price`;
create table `trader_refresh_price` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`trader_id` smallint(6) NOT NULL  COMMENT '随机商人ID',
	`time` smallint(6) NOT NULL COMMENT '点击次数',
	`price` bigint(20) NOT NULL COMMENT '价格',
UNIQUE KEY `time` (`trader_id`, `time`),
PRIMARY KEY (`id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='随机商店商人额外对话';
"
);
?>
