
<?php
db_execute($db, 
"
    ALTER TABLE `sword_soul_level`
    MODIFY COLUMN `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键';
"
);

?>

