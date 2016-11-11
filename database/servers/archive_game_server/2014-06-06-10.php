
<?php
db_execute($db, 
"
  ALTER TABLE `mission_level_small_box` ADD COLUMN   `box_dir` tinyint(4) NOT NULL COMMENT '宝箱朝向(0--左;1--右)';
"

);
?>