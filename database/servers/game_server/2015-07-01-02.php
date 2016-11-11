<?php
$this->AddSQL("
    ALTER TABLE `fate_box` ADD `fixed_prop` smallint(6)  NOT NULL DEFAULT '0' COMMENT '固定道具';
");
?>
