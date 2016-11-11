<?php 
$this->AddSQL("
ALTER TABLE `player_event_award_record` ADD COLUMN `json_event_record` mediumblob COMMENT 'json模板配置的玩家活动状态';
");
?>