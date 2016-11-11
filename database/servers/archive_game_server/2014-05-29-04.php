<?php
db_execute($db, 
"
DROP TABLE IF EXISTS `player_friend_state`;
CREATE TABLE `player_friend_state` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `delete_day_count` int(11) NOT NULL DEFAULT '0' COMMENT '每日删除好友数量',
  `delete_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '删除好友时间',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家好友功能状态数据';

"
);

?>