<?php

$this->AddSQL("ALTER TABLE `player_item` CHANGE `refine_level` `refine_level_bak` smallint(6) COMMENT '精炼等级备份';");
$this->AddSQL("ALTER TABLE `player_item` ADD `refine_level` smallint(6) DEFAULT 0 COMMENT '精炼等级';");
$this->AddSQL("create index item_id_idx on player_item (item_id,refine_level_bak);");


?>
