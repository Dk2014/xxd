<?php
db_execute($db, "

DROP TABLE IF EXISTS `ghost_train_price`;

CREATE TABLE `ghost_train_price` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `num` smallint(6) NOT NULL DEFAULT '0' COMMENT '次数',
  `price_ingot` bigint(20) NOT NULL DEFAULT '0' COMMENT '元宝培养价格',
  `price_coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '铜钱培养价格',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='魂侍培养价格表';

");
?>