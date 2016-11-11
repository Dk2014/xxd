<?php
$this->AddSQL("
alter table player_rainbow_level modify column  `reset_num`  int(11) NOT NULL DEFAULT '0' COMMENT '已重置次数';
");

?>
