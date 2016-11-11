<?php
$this->AddSQL("
CREATE TABLE IF NOT EXISTS `equipment_refine_new`(
`id` int(12) NOT NULL COMMENT '主键',
`equ_type` tinyint(4) NOT NULL COMMENT '装备类型(武器,防具等)',
`equ_kind` tinyint(4) NOT NULL COMMENT '装备种类(鹰扬,玄武等)',
`base_val` int(11) NOT NULL DEFAULT 0 COMMENT '装备基础强度',
`base_price` int(11) NOT NULL DEFAULT 0 COMMENT '装备基础价格',
`incre_val` int(11) NOT NULL DEFAULT 0 COMMENT '强化单位提升属性',
`incre_price` int(11) NOT NULL DEFAULT 0 COMMENT '强化单位价格',
PRIMARY KEY(`id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='新的按公式计算的装备强化表';
");
?>