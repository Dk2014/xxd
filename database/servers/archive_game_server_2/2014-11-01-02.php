<?php


//玩家打坐状态表去掉自增
$this->AddSQL("alter table `player_meditation_state` modify column `pid` bigint(20) NOT NULL COMMENT 'player id'");

?>
