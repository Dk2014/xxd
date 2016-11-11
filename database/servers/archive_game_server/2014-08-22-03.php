<?php
db_execute($db, "

alter table `role_level` add column `sunder_end_defend_rate` int(11) NOT NULL DEFAULT '0' COMMENT '破甲后减防（百分比）' after `sunder_end_hurt_rate`;
alter table `enemy_role` add column `sunder_end_defend_rate` int(11) NOT NULL DEFAULT '0' COMMENT '破甲后减防（百分比）' after `sunder_end_hurt_rate`;

");

?>
