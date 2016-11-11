<?php 

$this->AddSQL("
         alter table sealed_book add column `coins` int(11) NOT NULL DEFAULT '0' COMMENT '消耗铜钱';
        ");
?>
