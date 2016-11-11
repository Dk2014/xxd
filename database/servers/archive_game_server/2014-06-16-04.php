
<?php
db_execute($db, 
"

ALTER TABLE `equipment_refine_level`
MODIFY COLUMN `level`  tinyint(4) NOT NULL COMMENT '精练级别' AFTER `id`,
MODIFY COLUMN `quality`  tinyint(4) NOT NULL COMMENT '品质' AFTER `level`,
MODIFY COLUMN `probability`  tinyint(4) NOT NULL COMMENT '精练成功概率' AFTER `quality`;

"
);

?>

