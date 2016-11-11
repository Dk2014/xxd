
<?php
db_execute($db, 
"

ALTER TABLE `multi_level` ADD COLUMN `award_lock` int(11) NOT NULL DEFAULT 0 COMMENT '奖励权值';

"
);

?>

