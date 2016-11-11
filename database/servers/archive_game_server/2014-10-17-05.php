<?php

db_execute($db, "
alter table `ghost_star` add column costs bigint(20) NOT NULL DEFAULT '0' COMMENT '费用';
");
?>
