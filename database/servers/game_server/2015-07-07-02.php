<?php
$this->AddSQL("
  alter table `totem` add column `can_rand` tinyint(4) not null default 0 comment '是否可抽取';
  alter table `totem` add column `rand_need_times` smallint(6) not null default 0 comment '正常抽取需求元宝抽取次数';
  alter table `player_totem_info` add column `ingot_draw_times` smallint(6) not null default 0 comment '玩家元宝抽取次数';
  ",false,array("soha"));
?>