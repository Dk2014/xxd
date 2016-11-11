<?php
db_execute($db, "
alter table `vip_privilege_config` add column `unit` varchar(4) not null default '' comment '特权次数单位';
");
?>
