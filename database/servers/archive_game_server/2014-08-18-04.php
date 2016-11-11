<?php
db_execute($db, "
alter table `skill` add column `warcry` varchar(1024) NOT NULL DEFAULT '' COMMENT '战吼';

");

?>
