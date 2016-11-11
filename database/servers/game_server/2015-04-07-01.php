<?php 


$this->AddSQL("

CREATE TABLE IF NOT EXISTS `sealed_book` (
    `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '物品ID',
    `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型(1- - 伙伴；1 - 魂侍；5 - 怪物；7 - 魂侍技1； 8 - 魂侍技2)',
    `item_id` smallint(4) NOT NULL DEFAULT '0' COMMENT '物品id',
    `health` int(11) NOT NULL DEFAULT '0' COMMENT '生命',
    `attack` int(11) NOT NULL DEFAULT '0' COMMENT '攻击',
    `defense` int(11) NOT NULL DEFAULT '0' COMMENT '防御',
    `cultivation` int(11) NOT NULL DEFAULT '0' COMMENT '内力',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='天书配置';
");

?>




