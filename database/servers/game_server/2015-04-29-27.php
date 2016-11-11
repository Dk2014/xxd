<?php

$this->AddSQL("
 alter table `ancestral_hall`  add  column   `desc` text COMMENT '对应等级描述';
");
?>
