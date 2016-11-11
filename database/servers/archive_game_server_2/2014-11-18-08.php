<?php
$this->AddSQL("

ALTER TABLE `global_tb_xxd_onlinecnt` CHANGE `id` `id` BIGINT(20)  NOT NULL  AUTO_INCREMENT  COMMENT '主键ID';

");

?>
