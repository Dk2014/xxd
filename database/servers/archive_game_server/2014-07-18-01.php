<?php
db_execute($db,"

alter table `player_extend_level` add column  `role_pos` tinyint(4) NOT NULL COMMENT '主角站位';

"
);
?>
