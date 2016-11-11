<?php

$this->AddSQL("

CREATE TABLE IF NOT EXISTS `game_scene` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '剧情ID',
  `order` int(11) NOT NULL COMMENT '剧情排序',
  `name` varchar(32) NOT NULL COMMENT '剧情标题',
  `quest_order` int(11) NOT NULL COMMENT '关联任务',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='主线任务剧情';

");

?>
