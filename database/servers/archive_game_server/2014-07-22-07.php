<?php
db_execute($db,"
alter table player_info modify column `first_login_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '玩家注册时间';
"
);
?>
