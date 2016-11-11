<?php

$this->AddSQL("
alter table `pve_level` add column `basic_award_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '基础奖励';

alter table `pve_level` add column `award_factor` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励系数';

");

?>
