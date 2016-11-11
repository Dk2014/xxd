<?php
db_execute($db,"
alter table player_hard_level_state add column round int(11) not null default 0 comment '累积回合数';
"
);
?>

