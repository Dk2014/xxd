<?php

db_execute($db, "

ALTER TABLE `mission_enemy` ADD COLUMN `best_round` tinyint(4) NOT NULL COMMENT '最好的通关回合数';
  
");
?>