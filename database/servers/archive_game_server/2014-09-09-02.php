<?php
db_execute($db, "
alter table  `mission_level_box` change column `award_type` `award_type` tinyint(4) NOT NULL COMMENT '奖励类型(0--铜钱;1--道具;2--装备; 3--经验;4--经验倍数; 5--铜钱倍数,6--契约球)';
")
?>
