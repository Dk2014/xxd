<?php
$this->AddSQL("
create table `player_ghost_state` (
	  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
	  `upgrade_by_ingot_num` int(11) NOT NULL DEFAULT 0 COMMENT '今日使用元宝升升星次数',
	  `upgrade_by_ingot_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '最近一次使用元宝升星时间',
PRIMARY KEY(`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家魂侍相关状态';
");

?>
