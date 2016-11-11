<?php

$this->AddSQL("
 alter table `player_global_clique_info`  add  column   `worship_type` tinyint(4)  NOT NULL DEFAULT '0' COMMENT '上香类型';
");
?>
