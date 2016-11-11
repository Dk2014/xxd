<?php

$this->AddSQL("
create table mission_level_drama (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
  `name` varchar(30) NOT NULL COMMENT '剧情名称',
  `mission_level_id` int(11) NOT NULL DEFAULT '0' COMMENT '任务区域关卡ID',
  `quest_id` smallint(6) NOT NULL COMMENT '关联任务ID',
  `quest_state` tinyint(4) DEFAULT NULL COMMENT '任务状态',
  `area_x` smallint(6) NOT NULL COMMENT '区域X坐标',
  `area_y` smallint(6) NOT NULL COMMENT '区域Y坐标',
  `area_width` smallint(6) NOT NULL COMMENT '区域',
  `area_height` smallint(6) NOT NULL COMMENT '区域',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='关卡剧情';
");

