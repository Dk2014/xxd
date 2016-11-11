<?php

$this->AddSQL("

alter table `pve_level` modify column `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '要求等级';

");

?>
