<?php

db_execute($db, "
DROP TABLE IF EXISTS `skill_content`;

CREATE TABLE `skill_content` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `skill_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '绝招ID',
  `release_num` int(11) NOT NULL DEFAULT '0' COMMENT '释放次数',
  `recover_round_num` int(11) NOT NULL DEFAULT '0' COMMENT '恢复回合数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='绝招数据表';
");
?>