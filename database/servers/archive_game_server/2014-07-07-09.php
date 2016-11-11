<?php
db_execute($db, "

ALTER TABLE `player_chest_state`
ADD COLUMN `is_first_coin_ten`  tinyint(4) NOT NULL COMMENT '是否第一次青龙宝箱十连抽' AFTER `coin_chest_ten_num`,
ADD COLUMN `is_first_ingot_ten`  tinyint(4) NOT NULL COMMENT '是否第一次神龙宝箱十连抽' AFTER `ingot_chest_ten_num`;

");
?>
