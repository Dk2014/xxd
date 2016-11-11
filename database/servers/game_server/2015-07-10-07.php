<?php

$this->AddSQL("
DROP TABLE IF EXISTS `monster_property_addition`;

CREATE TABLE IF NOT EXISTS `monster_property_addition`(
    `id` int(12) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `monster_id` int(12)  NOT NULL DEFAULT '0' COMMENT '怪物ID',
    `level` int(12) NOT NULL DEFAULT '0' COMMENT '等级',
    `health` int(12) NOT NULL DEFAULT '0' COMMENT  '生命',
    `attack` int(12) NOT NULL DEFAULT '0' COMMENT '攻击',
    `defence` int(12) NOT NULL DEFAULT '0' COMMENT '防御',
    `cultivation` int(12) NOT NULL DEFAULT '0' COMMENT '内力',
    `speed`  int(12) NOT NULL DEFAULT '0' COMMENT '速度',
    `sunder_max_value`  int(12) NOT NULL DEFAULT '0' COMMENT '护甲上限',
    `sunder_min_hurt_rate` int(12) NOT NULL DEFAULT '0'COMMENT '破甲前百分比',
    `sunder_end_hurt_rate` int(12) NOT NULL DEFAULT '0' COMMENT '破甲后百分比',
    `sunder_end_defend_rate` int(12) NOT NULL DEFAULT '0' COMMENT '破甲后减防百分比',
    `sunder_attack` int(12) NOT NULL DEFAULT '0' COMMENT '攻击破甲值',
    `critial` int(12) NOT NULL DEFAULT '0' COMMENT '暴击%',
    `block`  int(12) NOT NULL DEFAULT '0' COMMENT '格挡%',
    `hit` int(12)  NOT NULL DEFAULT '0' COMMENT '命中%',
    `dodge` int(12) NOT NULL DEFAULT '0' COMMENT '闪避%',
    `critial_hurt` int(12) NOT NULL DEFAULT '0' COMMENT '必杀%',
    `toughness` int(12) NOT NULL DEFAULT '0' COMMENT '韧性%',
    `destroy` int(12) NOT NULL DEFAULT '0' COMMENT '破击%',
    `sleep` int(12) NOT NULL DEFAULT '0' COMMENT '睡眠抗性%',
    `dizziness` int(12) NOT NULL DEFAULT '0' COMMENT '眩晕抗性%',
    `random` int(12) NOT NULL DEFAULT '0' COMMENT '混乱抗性%',
    `disable_skill` int(12) NOT NULL DEFAULT '0' COMMENT '封魔抗性%',
    `poisoning` int(12) NOT NULL DEFAULT '0' COMMENT '中毒抗性%',
    PRIMARY KEY (`id`)

)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='怪物属性加成';
");
?>
