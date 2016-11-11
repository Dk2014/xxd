<?php
db_execute($db, "
	
CREATE TABLE `tower_level` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `floor` smallint(6) NOT NULL COMMENT '楼层',
  PRIMARY KEY (`id`),
  UNIQUE KEY (`floor`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='极限关卡通天塔';

");
?>