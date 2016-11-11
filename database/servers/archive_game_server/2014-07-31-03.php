<?php
db_execute($db, "
alter table  `player_global_friend_state` add column `platform_friend_num` int(11) not null default '0' comment '平台好友历史最大值';

alter table `platform_friend_award` modify column `num` int(11) NOT NULL COMMENT '物品数量';
");

?>
