<?php
$this->AddSQL("
alter table player_rainbow_level add column  `auto_fight_num`  tinyint(4) NOT NULL DEFAULT '0' COMMENT '今日扫荡次数';
");

?>
