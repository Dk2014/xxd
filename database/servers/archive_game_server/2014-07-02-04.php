<?php
db_execute($db, "

DROP TABLE IF EXISTS `ghost_star`;
CREATE TABLE `ghost_star` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `star` tinyint(4) NOT NULL DEFAULT '0' COMMENT '星级',
  `need_fragment_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '需要碎片数量',
  `growth` smallint(6) NOT NULL DEFAULT '0' COMMENT '成长值',
  `color` tinyint(4)  NOT NULL DEFAULT '0' COMMENT '颜色',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='魂侍进阶表';

");
?>
