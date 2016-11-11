<?php

$this->AddSQL("

alter table `pve_level` add column `level` smallint(6) NOT NULL DEFAULT '0' COMMENT '要求等级';

");

?>
