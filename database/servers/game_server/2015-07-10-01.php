<?php

$this->AddSQL("
CREATE TABLE `event_vn_dragon_ball_exchange`(
`id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '兑换物品标识',
`quality` tinyint(4) NOT NULL DEFAULT 0 COMMENT '品质',
`name` varchar(100) NOT NULL COMMENT '名称',
`desc` varchar(200) NOT NULL COMMENT '描述',
PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='越南龙珠活动兑换列表';
", false, array('soha'));
?>
