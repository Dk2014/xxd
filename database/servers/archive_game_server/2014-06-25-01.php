<?php

db_execute($db, "

ALTER TABLE `player_chest_state`
CHANGE COLUMN `last_coin_chest_at` `last_free_coin_chest_at`  bigint(20) NOT NULL COMMENT '上次开免费青铜宝箱时间' ,
MODIFY COLUMN `coin_chest_num`  int(11) NOT NULL COMMENT '今天开青铜宝箱次数' AFTER `last_free_coin_chest_at`,
MODIFY COLUMN `coin_chest_consume`  bigint(20) NOT NULL COMMENT '今天开青铜宝箱花费铜钱数' AFTER `coin_chest_num`,
CHANGE COLUMN `last_ingot_chest_at` `last_free_ingot_chest_at`  bigint(20) NOT NULL COMMENT '上次开免费神龙宝箱时间' AFTER `coin_chest_consume`,
MODIFY COLUMN `ingot_chest_num`  int(11) NOT NULL COMMENT '今天开神龙宝箱次数' AFTER `last_free_ingot_chest_at`,
MODIFY COLUMN `ingot_chest_consume`  bigint(20) NOT NULL COMMENT '今天开神龙宝箱花费元宝数' AFTER `ingot_chest_num`,
ADD COLUMN `free_coin_chest_num`  int(11) NOT NULL COMMENT '每日免费青铜宝箱数' AFTER `pid`;

");
?>
