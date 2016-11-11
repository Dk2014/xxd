
<?php
db_execute($db, 
"

ALTER TABLE `equipment_recast`
CHANGE COLUMN `consume_coins` `consume_coin`  bigint(20) NOT NULL COMMENT '消耗的铜钱' AFTER `orange_crystal_num`;

ALTER TABLE `equipment_refine`
CHANGE COLUMN `level1_consume_coins` `level1_consume_coin`  bigint(20) NOT NULL COMMENT '精练到1级消耗的铜钱' AFTER `orange_crystal_num`,
CHANGE COLUMN `level2_consume_coins` `level2_consume_coin`  bigint(20) NOT NULL COMMENT '精练到2级消耗的铜钱' AFTER `level1_consume_coin`,
CHANGE COLUMN `level3_consume_coins` `level3_consume_coin`  bigint(20) NOT NULL COMMENT '精练到3级消耗的铜钱' AFTER `level2_consume_coin`,
CHANGE COLUMN `level4_consume_coins` `level4_consume_coin`  bigint(20) NOT NULL COMMENT '精练到4级消耗的铜钱' AFTER `level3_consume_coin`,
CHANGE COLUMN `level5_consume_coins` `level5_consume_coin`  bigint(20) NOT NULL COMMENT '精练到5级消耗的铜钱' AFTER `level4_consume_coin`;

"
);

?>

