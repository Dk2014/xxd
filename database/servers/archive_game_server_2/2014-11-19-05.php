<?php
$this->AddSQL("
create table `ghost_train_price` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',	
  `times` int(11) NOT NULL COMMENT '购买次数',	
  `cost` int(11) NOT NULL COMMENT '一个果实单价',
PRIMARY KEY(`id`),
UNIQUE KEY `times` (`times`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='魂侍元宝培养价格梯度';
");

$this-> AddSQL("
	alter table `player_ghost_state` add column `train_by_ingot_num`  int(11) NOT NULL DEFAULT '0' COMMENT '今日使用元宝培养次数';
");

$this->AddSQL("
	 alter table `player_ghost_state` add column `train_by_ingot_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最近一次使用元宝培养时间';
");

?>
