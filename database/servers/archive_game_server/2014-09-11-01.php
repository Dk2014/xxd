<?php
db_execute($db, "
alter table  `mission_level` add column `award_pet`  tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否出现灵宠 0--否 1--是';
")
?>
