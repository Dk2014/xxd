<?php
db_execute($db,"
alter table player_sword_soul_state modify column num smallint(6) not null comment '当前可拔剑次数';
"
);
?>
