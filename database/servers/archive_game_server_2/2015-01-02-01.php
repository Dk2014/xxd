<?php
$this->AddSQL("
alter table `func` add column `type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '开放类型 1--权值 2--等级';
");
?>
