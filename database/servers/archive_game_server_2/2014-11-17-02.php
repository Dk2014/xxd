<?php
$this->AddSQL("
ALTER TABLE `events_group_buy` MODIFY COLUMN `base_price` smallint(6) NOT NULL COMMENT '团购物品低价';
");

?>
