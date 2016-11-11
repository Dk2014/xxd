<?php
db_execute($db, "
create table `func_prediction` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`order` smallint(6) NOT NULL COMMENT '顺序',
	`type` tinyint(4) NOT NULL COMMENT '类型 0--等级触发 1--新功能权值触发',
	`condition_value` smallint(6) NOT NULL COMMENT '触发条件 等级 或 功能权值',
	`sign` varchar(30) NOT NULL COMMENT '资源标识',
	`summary` varchar(30) NOT NULL COMMENT '下一功能描述',
	`tips` varchar(1024) NOT NULL COMMENT 'tips',
PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='功能预告';
");
