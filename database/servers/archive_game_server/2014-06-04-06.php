<?php
db_execute($db, 
"

CREATE TABLE `mission_talk` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '对话ID',
  `enemy_id` int(11) NOT NULL COMMENT '副本敌人ID',
  `content` varchar(1024) DEFAULT '' COMMENT '对话内容',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='副本战场对话';
"
);

?>
