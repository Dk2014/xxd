<?php

db_execute($db, "

ALTER TABLE `mission_level` modify `award_lock` int(11) NOT NULL COMMENT '通关奖励权值';
  
");
?>