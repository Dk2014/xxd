<?php
db_execute($db, "


create table player_purchase_record (
	`id` bigint(20) NOT NULL COMMENT '主键ID',
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`item_id` smallint(6) NOT NULL COMMENT '物品ID',
	`num` smallint(6) NOT NULL COMMENT '已购买数量',
	`timestamp` bigint(20) NOT NULL COMMENT '上次购买时间',
	PRIMARY KEY (`id`),
KEY `idx_pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家物品购买记录 仅记录限购商品' ;

create table purchase_limit (
	`id` int(11) NOT NULL COMMENT '主键ID',
	`item_id` smallint(6) NOT NULL COMMENT '物品ID',
	`num` smallint(6) NOT NULL COMMENT '购买次数限制',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='物品购买次数限制' ;
");

?>
