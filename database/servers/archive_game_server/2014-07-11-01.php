<?php
db_execute($db,"
alter table `role_realm_level` modify column  `need_realm_class` smallint(6) NOT NULL DEFAULT '0' COMMENT '升级所需阶级';
"
);
?>
