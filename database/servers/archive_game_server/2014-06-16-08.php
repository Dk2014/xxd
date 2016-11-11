
<?php
db_execute($db, 
"

ALTER TABLE `player_multi_level_info` ADD COLUMN `lock` int(11) NOT NULL DEFAULT 0 COMMENT '关卡开启权值';

"
);

?>

