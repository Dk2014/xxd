<?php
db_execute($db, "
alter table `player_login_award_record` modify column  `record` int(11) NOT NULL DEFAULT '0' COMMENT '七天奖励领取记录';
")
?>
