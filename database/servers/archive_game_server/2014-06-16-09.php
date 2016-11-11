
<?php
db_execute($db, 
"

ALTER TABLE `multi_level` ADD COLUMN `lock` int(11) NOT NULL DEFAULT 0 COMMENT '关卡开启权值';

"
);

?>

