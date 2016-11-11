<?php
db_execute($db, "
alter table `enemy_role` add column `prop` tinyint(4) DEFAULT '0' COMMENT '属性' after `sign`;
");

?>
