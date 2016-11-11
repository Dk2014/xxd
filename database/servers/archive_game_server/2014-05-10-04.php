<?php

db_execute($db, "

  ALTER TABLE `player_formation` ADD pos3  tinyint(4) NOT NULL DEFAULT '-1' COMMENT '站位3';
  ALTER TABLE `player_formation` ADD pos4  tinyint(4) NOT NULL DEFAULT '-1' COMMENT '站位4';
  ALTER TABLE `player_formation` ADD pos5  tinyint(4) NOT NULL DEFAULT '-1' COMMENT '站位5';
");
?>
