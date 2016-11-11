
<?php
db_execute($db, 
"
ALTER TABLE `quest` CHANGE COLUMN `enemy_num` `enemy_num` int(10) unsigned NOT NULL COMMENT '敌人组数';

"

);
?>