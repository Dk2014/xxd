<?php
$this->AddSQL("
  CREATE TABLE IF NOT EXISTS `monster_property_addition`(
  `id` int(12) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `monster_id` int(12)  NOT NULL DEFAULT '0' COMMENT '怪物ID',
  `level` int(12) NOT NULL DEFAULT '0' COMMENT '等级',
  `shengming` int(12) NOT NULL DEFAULT '0' COMMENT  '生命',
  `gongji` int(12) NOT NULL DEFAULT '0' COMMENT '攻击',
  `fangyu` int(12) NOT NULL DEFAULT '0' COMMENT '防御',
  `neili` int(12) NOT NULL DEFAULT '0' COMMENT '内力',
  `sudu`  int(12) NOT NULL DEFAULT '0' COMMENT '速度',
  `hujia`  int(12) NOT NULL DEFAULT '0' COMMENT '护甲上限',
  `pojiaqian` int(12) NOT NULL DEFAULT '0'COMMENT '破甲前百分比',
  `pojiahou` int(12) NOT NULL DEFAULT '0' COMMENT '破甲后百分比',
  `pojiahoujianfang` int(12) NOT NULL DEFAULT '0' COMMENT '破甲后减防百分比',
  `gongjipojia` int(12) NOT NULL DEFAULT '0' COMMENT '攻击破甲值',
  `baoji` int(12) NOT NULL DEFAULT '0' COMMENT '暴击%',
  `gedang`  int(12) NOT NULL DEFAULT '0' COMMENT '格挡%',
  `mingzhong` int(12)  NOT NULL DEFAULT '0' COMMENT '命中%',
  `shanbi` int(12) NOT NULL DEFAULT '0' COMMENT '闪避%',
  `bisha` int(12) NOT NULL DEFAULT '0' COMMENT '必杀%',
  `renxing` int(12) NOT NULL DEFAULT '0' COMMENT '韧性%',
  `poji` int(12) NOT NULL DEFAULT '0' COMMENT '破击%',
  `shuimian` int(12) NOT NULL DEFAULT '0' COMMENT '睡眠抗性%',
  `xuanyun` int(12) NOT NULL DEFAULT '0' COMMENT '眩晕抗性%',
  `hunluan` int(12) NOT NULL DEFAULT '0' COMMENT '混乱抗性%',
  `fengmo` int(12) NOT NULL DEFAULT '0' COMMENT '封魔抗性%',
  `zhongdu` int(12) NOT NULL DEFAULT '0' COMMENT '中毒抗性%',
  PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='怪物属性加成';
");
?>
