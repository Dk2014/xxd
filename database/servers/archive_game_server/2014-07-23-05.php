<?php
db_execute($db,"

CREATE TABLE `item_costprops` (
  `item_id` smallint(6) NOT NULL COMMENT '道具ID',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '消耗类型； 0 - 经验； 1 - 体力',
  `value` int(11) NOT NULL DEFAULT '0' COMMENT '值',
  PRIMARY KEY (`item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='消耗道具';

"
);
?>
