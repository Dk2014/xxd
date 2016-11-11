<?php
$this->AddSQL("
ALTER TABLE `player_item` MODIFY `refine_level` SMALLINT(6) DEFAULT 0 COMMENT '精炼等级';
");
?>