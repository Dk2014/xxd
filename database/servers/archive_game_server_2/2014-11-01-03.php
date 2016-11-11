<?php


//玩家打坐状态表去掉自增
$this->AddSQL("alter table `player_chest_state` add column  `last_free_pet_chest_at` bigint(20) NOT NULL DEFAULT 0 COMMENT '上次开免费灵兽宝箱时间';");

?>
