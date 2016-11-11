<?php
db_execute($db, "

alter table `role_level` change column `will` `sleep` int(11) NOT NULL DEFAULT '0' COMMENT '睡眠抗性';
alter table `role_level` add column `dizziness` int(11) NOT NULL COMMENT '眩晕抗性' after `sleep`;
alter table `role_level` add column `random` int(11) NOT NULL COMMENT '混乱抗性' after `dizziness`;
alter table `role_level` add column `disable_skill` int(11) NOT NULL COMMENT '封魔抗性' after `random`;
alter table `role_level` add column `poisoning` int(11) NOT NULL COMMENT '中毒抗性' after `disable_skill`;

alter table `enemy_role` change column `will` `sleep` int(11) NOT NULL DEFAULT '0' COMMENT '睡眠抗性' after `sunder_end_hurt_rate`;
alter table `enemy_role` add column `dizziness` int(11) NOT NULL COMMENT '眩晕抗性' after `sleep`;
alter table `enemy_role` add column `random` int(11) NOT NULL COMMENT '混乱抗性' after `dizziness`;
alter table `enemy_role` add column `disable_skill` int(11) NOT NULL COMMENT '封魔抗性' after `random`;
alter table `enemy_role` add column `poisoning` int(11) NOT NULL COMMENT '中毒抗性' after `disable_skill`;

");

?>
