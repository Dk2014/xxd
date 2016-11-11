<?php
db_execute($db,"
alter table skill add column target tinyint(4) not null default '1' comment '攻击目标（客户端展示用）';
"
);
?>
