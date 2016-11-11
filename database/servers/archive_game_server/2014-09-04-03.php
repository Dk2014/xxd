<?php
db_execute($db, "
alter table  `rainbow_level_award` modify column `award_type`  tinyint(4) NOT NULL COMMENT '奖励类型(0--铜钱;1--道具;2--装备 3--经验 4--经验倍数 5--铜钱倍数 6--恢复伙伴技能 7--恢复魂侍技能 8--恢复灵宠状态 9--主角精气 10--百分比生命 11-增加魂力)';
")
?>
