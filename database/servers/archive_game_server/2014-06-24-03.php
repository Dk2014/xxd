<?php

db_execute($db, "

CREATE TABLE `player_chest_state` (
`pid` bigint(20) NOT NULL COMMENT '玩家id',
`coin_chest_num` int(11) NOT NULL COMMENT '开青铜宝箱次数',
`coin_chest_consume` bigint(20) NOT NULL COMMENT '开青铜宝箱花费铜钱数',
`last_coin_chest_at` bigint(20) NOT NULL COMMENT '上次开青铜宝箱时间',
`ingot_chest_num` int(11) NOT NULL COMMENT '开神龙宝箱次数',
`ingot_chest_consume` bigint(20) NOT NULL COMMENT '开神龙宝箱花费元宝数',
`last_ingot_chest_at` bigint(20) NOT NULL COMMENT '上次开神龙宝箱时间',
PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家宝箱状态';
");

?>

