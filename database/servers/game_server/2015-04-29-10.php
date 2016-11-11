<?php

$this->AddSQL("
	ALTER TABLE `global_clique` ADD COLUMN `contrib_clear_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '帮派贡献清除时间';
");

?>