<?php
db_execute($db, 
	"
	CREATE TABLE `func` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(10) NOT NULL COMMENT '功能名称',
  `sign` varchar(30) NOT NULL COMMENT '功能标识',
  `lock` smallint(6) NOT NULL COMMENT '功能权值',
  `level` smallint(6) NOT NULL DEFAULT '0' COMMENT '开启等级',
  `unique_key` bigint(20) NOT NULL DEFAULT '0' COMMENT '唯一权值',
  `need_play` tinyint(4) DEFAULT '1' COMMENT '是否需要播放 0不需要 1需要',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='功能权值配置';

	"
);
?>