<?php
db_execute($db, "

create table  `level_func`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`name` varchar(10) NOT NULL COMMENT '功能名称',
	`sign` varchar(30) NOT NULL COMMENT '功能标识',
	`level` smallint(6) NOT NULL DEFAULT '0' COMMENT '开启等级',
	`need_play` tinyint(4) DEFAULT '0' COMMENT '是否需要播放 0不需要 1需要',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='等级功能配置';
");

?>
