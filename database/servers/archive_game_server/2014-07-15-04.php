<?php
db_execute($db,"

alter table `player_extend_level` add column  `buddy_pos` tinyint(4) NOT NULL COMMENT '随机的伙伴角色位置';
alter table `player_extend_level` add column  `buddy_tactical` tinyint(4) NOT NULL COMMENT '伙伴关卡队伍战术';

"
);
?>
