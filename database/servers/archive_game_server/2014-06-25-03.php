<?php

db_execute($db, "

ALTER TABLE `player_chest_state`
ADD COLUMN `coin_chest_ten_num`  int(11) NOT NULL COMMENT '今日青铜宝箱十连抽次数' AFTER `coin_chest_num`,
ADD COLUMN `ingot_chest_ten_num`  int(11) NOT NULL COMMENT '今日神龙宝箱十连抽次数' AFTER `ingot_chest_num`;

");
?>
