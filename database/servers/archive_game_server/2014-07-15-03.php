<?php
db_execute($db,"

alter table `player_extend_level` add column  `rand_buddy_role_id` tinyint(4) NOT NULL COMMENT '随机的伙伴角色ID';

"
);
?>
