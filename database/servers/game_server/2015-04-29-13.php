<?php
$this->AddSQL("
	ALTER TABLE `global_clique` ADD COLUMN `recruit_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '招募公告时间';
");
?>
