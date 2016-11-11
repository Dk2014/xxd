<?php
$this->AddSQL("
    ALTER TABLE `player_extend_level` ADD COLUMN `coins_buy_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '元宝购买铜钱关卡次数';
    ALTER TABLE `player_extend_level` ADD COLUMN `exp_buy_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '元宝购买经验关卡次数';
");
?>
