<?php

db_execute($db, "
alter table ghost add column   `quality` tinyint(4) NOT NULL DEFAULT '0' COMMENT '技能品质';
");

?>
