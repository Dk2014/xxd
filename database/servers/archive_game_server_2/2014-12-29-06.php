<?php

$this->AddSQL("
ALTER TABLE `resource_origin` ADD COLUMN `description` varchar(1024) NOT NULL DEFAULT '' COMMENT '描述';
");

?>
