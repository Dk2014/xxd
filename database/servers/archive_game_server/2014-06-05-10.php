
<?php
db_execute($db, 
"

ALTER TABLE `mission_level_small_box` change column `award_chance` `probability` tinyint(4) NOT NULL COMMENT '出现概率';
ALTER TABLE `mission_level_small_box_items` change column `award_chance` `probability` tinyint(4) NOT NULL COMMENT '出现概率';
ALTER TABLE `mission_level_small_box_items` ADD COLUMN `item_number` int(11) NOT NULL DEFAULT '0' COMMENT '奖励数量';
ALTER TABLE `mission_level_small_box_items` ADD COLUMN `award_type` tinyint(4) NOT NULL COMMENT '奖励类型(0--铜钱;1--道具;2--装备)';

"
);

?>

