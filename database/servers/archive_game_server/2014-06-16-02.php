
<?php
db_execute($db, 
"

ALTER TABLE `equipment_refine`
MODIFY COLUMN `level1_consume_coins`  bigint(20) NOT NULL COMMENT '精练到1级消耗的铜钱' AFTER `orange_crystal_num`,
MODIFY COLUMN `level2_consume_coins`  bigint(20) NOT NULL COMMENT '精练到2级消耗的铜钱' AFTER `level1_consume_coins`,
MODIFY COLUMN `level3_consume_coins`  bigint(20) NOT NULL COMMENT '精练到3级消耗的铜钱' AFTER `level2_consume_coins`,
MODIFY COLUMN `level4_consume_coins`  bigint(20) NOT NULL COMMENT '精练到4级消耗的铜钱' AFTER `level3_consume_coins`,
MODIFY COLUMN `level5_consume_coins`  bigint(20) NOT NULL COMMENT '精练到5级消耗的铜钱' AFTER `level4_consume_coins`;

"
);

?>

