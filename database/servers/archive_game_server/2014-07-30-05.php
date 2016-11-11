<?php
db_execute($db, "
	
alter table `role_realm_class` add column `add_cultivation` int(11) NOT NULL DEFAULT '0' COMMENT '增加内力';

");
?>
