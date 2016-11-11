<?php
$this->AddSQL("
  alter table `totem` drop column `can_rand` ;
  alter table `totem` drop column `rand_need_times`;
  alter table `player_totem_info` drop column `ingot_draw_times`;
  alter table `totem_level_info` add column `upgrade_need_rock` smallint(6) not null default 0 comment '升级所需石符文';
	alter table `totem_level_info` add column `upgrade_need_jade` smallint(6) not null default 0 comment '升级所需玉符文';
  ",false,array("soha"));
?>