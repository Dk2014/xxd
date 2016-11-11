<?php
db_execute($db, "
alter table rainbow_level_award add column `order` tinyint(4) NOT NULL COMMENT '品质顺序';
");
?>
