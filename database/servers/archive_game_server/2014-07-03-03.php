
<?php
db_execute($db, "

ALTER TABLE `arena_award_box` ADD COLUMN `item2_id` smallint(6) NOT NULL COMMENT '物品2';
ALTER TABLE `arena_award_box` ADD COLUMN `item2_num` smallint(6) NOT NULL COMMENT '物品2数量';

");
?>
