<?php

db_execute($db, "

ALTER TABLE `player_chest_state`
ADD COLUMN `last_coin_chest_at`  bigint(20) NOT NULL COMMENT '上次开消费青铜宝箱时间' AFTER `coin_chest_consume`,
ADD COLUMN `last_ingot_chest_at`  bigint(20) NOT NULL COMMENT '上次开消费神龙宝箱时间' AFTER `ingot_chest_consume`;

");
?>
