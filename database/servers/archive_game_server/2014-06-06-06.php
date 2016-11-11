
<?php
db_execute($db, 
"
  ALTER TABLE `mission_level` ADD COLUMN `award_box` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否奖励宝箱' after `award_coin`;
"

);
?>