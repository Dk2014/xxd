<?php
$this->AddSQL("
ALTER TABLE `player_fate_box_state` ADD hunyuan_box_free_draw_timestamp bigint(20) NOT NULL DEFAULT 0 COMMENT '混元命匣免费抽取时间戳';
ALTER TABLE `player_fate_box_state` ADD hunyuan_box_draw_count int(11) NOT NULL DEFAULT 0 COMMENT '混元宝箱抽奖次数';
");
?>