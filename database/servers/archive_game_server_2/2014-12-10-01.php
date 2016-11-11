<?php

$this->AddSQL("
CREATE TABLE `sword_soul_probability` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '箱子ID',
  `box_name` varchar(16) NOT NULL COMMENT '箱子名称',
  `order` tinyint(4) NOT NULL COMMENT '箱子顺序',
  `level_up` tinyint(4) NOT NULL COMMENT '升箱概率',
  `get_exp` tinyint(4) NOT NULL COMMENT '获得潜龙概率',
  `get_rubbish` tinyint(4) NOT NULL COMMENT '获得杂物概率',
  `green` tinyint(4) NOT NULL COMMENT '优秀剑心概率',
  `blue` tinyint(4) NOT NULL COMMENT '精良剑心概率',
  `purple` tinyint(4) NOT NULL COMMENT '传奇剑心概率',
  `yello` tinyint(4) NOT NULL COMMENT '神器剑心概率',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COMMENT='剑山拔剑概率';

INSERT INTO `sword_soul_probability` (`box_name`, `order`, `level_up`, `get_exp`, `get_rubbish`, `green`, `blue`, `purple`, `yello`) values
('BOX_A', 1, 40, 0, 75, 25, 0, 0, 0),
('BOX_B', 2, 40, 0, 60, 35, 5, 0, 0),
('BOX_C', 3, 40, 0, 39, 45, 15, 1, 0),
('BOX_D', 4, 40, 20, 0, 0, 60, 20, 0),
('BOX_E', 5, 0, 30, 0, 0, 30, 30, 10),
('VIP_BOX_D', 7, 40, 70, 0, 0, 0, 30, 0),
('VIP_BOX_E', 8, 0, 50, 0, 0, 0, 40, 10);

ALTER TABLE `player_sword_soul_state`
	ADD `ingot_num` bigint(20),
	ADD `supersoul_additional_possible` tinyint(4),
	ADD `ingot_yello_one` smallint(6),
	ADD `last_ingot_draw_time` bigint(20);
UPDATE `player_sword_soul_state` SET `ingot_num`=0, `supersoul_additional_possible`=0, `ingot_yello_one`=0, `last_ingot_draw_time`=0;
");

