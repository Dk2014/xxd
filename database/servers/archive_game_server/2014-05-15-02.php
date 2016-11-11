<?php
db_execute($db, "

DROP TABLE IF EXISTS `ghost_level`;


CREATE TABLE `ghost_level` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `level` smallint(6) NOT NULL DEFAULT '0' COMMENT '魂侍等级',
  `exp` bigint(20) NOT NULL DEFAULT '0' COMMENT '升级所需经验',
  `need_crystal_num` bigint(20) NOT NULL DEFAULT '0' COMMENT '所需水晶数量',
  `min_add_exp` bigint(20) NOT NULL DEFAULT '0' COMMENT '最小经验加值',
  `max_add_exp` bigint(20) NOT NULL DEFAULT '0' COMMENT '最大经验加值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='魂侍等级表';
	

");
?>