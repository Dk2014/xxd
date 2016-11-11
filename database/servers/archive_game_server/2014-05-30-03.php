
<?php
db_execute($db, 
"
    ALTER TABLE `sword_soul_quality_level`
    MODIFY COLUMN `exp`  int(11) NOT NULL COMMENT '升到下一级所需的经验' AFTER `level`;
"
);

?>

