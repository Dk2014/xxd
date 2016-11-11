<?php
db_execute($db, "

DROP TABLE IF EXISTS `ghost_exchange`;


CREATE TABLE `ghost_exchange` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `target_ghost` smallint(6) NOT NULL DEFAULT '0' COMMENT '兑换的目标魂侍',
  `ghost_id1` smallint(6) NOT NULL DEFAULT '0' COMMENT '魂侍id1',
  `ghost_level1` smallint(6) NOT NULL DEFAULT '0' COMMENT '魂侍等级1',
  `ghost_num1` smallint(6) NOT NULL DEFAULT '0' COMMENT '魂侍数量1',
  `ghost_id2` smallint(6) NOT NULL DEFAULT '0' COMMENT '魂侍id2',
  `ghost_level2` smallint(6) NOT NULL DEFAULT '0' COMMENT '魂侍等级2',
  `ghost_num2` smallint(6) NOT NULL DEFAULT '0' COMMENT '魂侍数量2',
  `ghost_id3` smallint(6) NOT NULL DEFAULT '0' COMMENT '魂侍id3',
  `ghost_level3` smallint(6) NOT NULL DEFAULT '0' COMMENT '魂侍等级3',
  `ghost_num3` smallint(6) NOT NULL DEFAULT '0' COMMENT '魂侍数量3',
  `ghost_id4` smallint(6) NOT NULL DEFAULT '0' COMMENT '魂侍id4',
  `ghost_level4` smallint(6) NOT NULL DEFAULT '0' COMMENT '魂侍等级4',
  `ghost_num4` smallint(6) NOT NULL DEFAULT '0' COMMENT '魂侍数量4',
  `ghost_id5` smallint(6) NOT NULL DEFAULT '0' COMMENT '魂侍id5',
  `ghost_level5` smallint(6) NOT NULL DEFAULT '0' COMMENT '魂侍等级5',
  `ghost_num5` smallint(6) NOT NULL DEFAULT '0' COMMENT '魂侍数量5',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='魂侍兑换表';


");
?>