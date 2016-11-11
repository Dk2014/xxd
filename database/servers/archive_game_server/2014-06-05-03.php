
<?php
db_execute($db, 
"
ALTER TABLE `item` ADD COLUMN `appendix_num` tinyint(4) NOT NULL DEFAULT '0' COMMENT '追加属性数';

DROP TABLE IF EXISTS `equipment_appendix`;
CREATE TABLE `equipment_appendix` (
   `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT 'ID',
   `level` int(11) NOT NULL COMMENT '等级',
   `health` int(11) DEFAULT '0' COMMENT '生命',
  `cultivation` int(11) DEFAULT '0' COMMENT '内力',
  `speed` int(11) DEFAULT '0' COMMENT '速度',
  `attack` int(11) DEFAULT '0' COMMENT '攻击',
  `defence` int(11) DEFAULT '0' COMMENT '防御',
  `dodge_level` int(11) DEFAULT '0' COMMENT '闪避',
  `hit_level` int(11) DEFAULT '0' COMMENT '命中',
  `block_level` int(11) DEFAULT '0' COMMENT '格挡',
  `critical_level` int(11) DEFAULT '0' COMMENT '暴击',
  `tenacity_level` int(11) DEFAULT '0' COMMENT '韧性',
  `destroy_level` int(11) DEFAULT '0' COMMENT '破击',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='装备追加属性表';

"

);
?>