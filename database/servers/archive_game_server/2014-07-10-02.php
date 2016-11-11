<?php
db_execute($db,"

ALTER TABLE `sword_soul`
ADD COLUMN `fragment_id`  smallint(6) NULL COMMENT '兑换需要的碎片物品id' AFTER `fragment_num`;

"
);
?>

