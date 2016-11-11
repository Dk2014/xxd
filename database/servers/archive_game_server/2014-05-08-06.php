<?php

db_execute($db, "

alter table `mission_level` modify `type` tinyint(4) NOT NULL COMMENT '关卡类型(0--普通;1--精英;2--Boss)';
alter table `enemy_deploy_form` modify `battle_type` tinyint(4) NOT NULL COMMENT '战场类型(0--关卡;)';

");
?>
