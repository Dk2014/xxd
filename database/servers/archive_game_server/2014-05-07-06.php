<?php

db_execute($db, "
DROP TABLE IF EXISTS `role`;

CREATE TABLE `role` (
  `id` tinyint(4) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(10) NOT NULL COMMENT '角色名称',
  `sign` varchar(20) DEFAULT NULL COMMENT '资源标识',
  `type` tinyint(4) NOT NULL COMMENT '类型：1.主角，2.伙伴',
  `is_special` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否特殊伙伴 0不是 1是',
  `skill_id1` smallint(6) NOT NULL DEFAULT '0' COMMENT '默认绝招1',
  `skill_id2` smallint(6) NOT NULL DEFAULT '0' COMMENT '默认绝招2',
  `buddy_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '伙伴等级',
  `mission_lock` int(11) DEFAULT '0' COMMENT '解锁副本权值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

");
?>
