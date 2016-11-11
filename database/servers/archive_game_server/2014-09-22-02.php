<?php
db_execute($db, "
alter table `fashion` add column `item_sign` varchar(30) DEFAULT NULL COMMENT '对应物品资源标识';
");
?>
