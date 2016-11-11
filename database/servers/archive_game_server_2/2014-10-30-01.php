<?php

//修复旧灵宠数据
$this->AddSQL("update `player_battle_pet` set `parent_pet_id`=`battle_pet_id` where `parent_pet_id`='0'");

?>
