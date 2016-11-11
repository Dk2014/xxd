<?php
$this->AddSQL("
 ALTER TABLE `events_share_awards` ADD COLUMN `heart` smallint(6) DEFAULT 0 COMMENT '爱心';
");

?>