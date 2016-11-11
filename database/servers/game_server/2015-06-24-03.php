<?php
$this->AddSQL("
    ALTER TABLE `player_extend_level` ADD COLUMN `coins_buy_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '元宝购买铜钱关卡时间';
    ALTER TABLE `player_extend_level` ADD COLUMN `exp_buy_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '元宝购买经验关卡时间';
");
?>
