<?php
$this->AddSQL("
 ALTER TABLE `skill` ADD COLUMN `required_frame_level` int(11) NOT NULL DEFAULT '0' COMMENT '需要声望等级';
");

?>
