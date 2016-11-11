<?php
$this->AddSQL("
ALTER TABLE `rainbow_level_award` ADD COLUMN `autofight_box` tinyint(4) DEFAULT 0 COMMENT '是否是扫荡宝箱';
");

?>