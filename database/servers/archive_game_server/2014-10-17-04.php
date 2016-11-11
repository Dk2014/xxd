<?php

db_execute($db, "
alter table mail add column priority tinyint(4) NOT NULL DEFAULT '0' COMMENT '优先级';

alter table player_mail add column priority tinyint(4) NOT NULL DEFAULT '0' COMMENT '优先级';

alter table global_mail add column priority tinyint(4) NOT NULL DEFAULT '0' COMMENT '优先级';
");
?>
