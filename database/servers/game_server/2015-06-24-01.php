<?php
$this->AddSQL("
DROP TABLE IF EXISTS `buy_resource_times_config`;
CREATE TABLE `buy_resource_times_config`(
	`id` int(12) NOT NULL AUTO_INCREMENT,
	`times` int(12) NOT NULL COMMENT '购买次数',
	`cost` int(12) NOT NULL COMMENT '购买所需元宝',
	PRIMARY KEY (`id`),
	UNIQUE `times` (`times`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='购买体力消耗元宝配置';
");
?>
