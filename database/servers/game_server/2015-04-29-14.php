<?php
$this->AddSQL("
ALTER TABLE `global_clique` CHANGE COLUMN `contri`   `contrib` bigint(20) NOT NULL DEFAULT '0' COMMENT '帮派贡献';
");
?>
