<?php
db_execute($db, "
alter table `enemy_role` add column `show_shader` tinyint(4) DEFAULT '0' COMMENT '是否显示阴影';
");

?>
