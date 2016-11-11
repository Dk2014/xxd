<?php

$this->AddSQL("
ALTER TABLE `item` MODIFY `desc` varchar(200)  COMMENT '物品描述';
");

?>