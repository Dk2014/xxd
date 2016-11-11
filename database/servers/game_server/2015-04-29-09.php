<?php

$this->AddSQL("
	ALTER TABLE `global_clique` ADD COLUMN `contri` bigint(20) NOT NULL DEFAULT 0 COMMENT '帮派贡献';
");

?>