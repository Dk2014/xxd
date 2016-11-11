<?php
db_execute($db,"
alter table `player_daily_sign_in_state` add column `signed_today` tinyint(4) comment '今天是否已签到';
"
);
?>
