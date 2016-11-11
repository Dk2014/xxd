<?php

db_execute($db, "

CREATE TABLE `faq` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `order` int(11) NOT NULL COMMENT '顺序',
  `question` varchar(512) NOT NULL COMMENT '问题',
  `answer` varchar(1024) NOT NULL COMMENT '回答',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='问答';

");
?>

