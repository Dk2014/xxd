<?php

$this->AddSQL("
ALTER TABLE `clique_building_bank` CHANGE COLUMN `sliver_coupon_coins`  `silver_coupon_coins` int(11) NOT NULL DEFAULT '0' COMMENT '银劵底价';
ALTER TABLE `clique_building_bank` CHANGE COLUMN  `sliver_coupon_num` `silver_coupon_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '银劵限购';
ALTER TABLE `clique_building_bank` CHANGE COLUMN  `sliver_coupon_sold` `silver_coupon_sold` int(11) NOT NULL DEFAULT '0' COMMENT '银劵售价';



");


?>
