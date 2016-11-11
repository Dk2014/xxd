<?php

$this->AddSQL("
alter table player_global_clique_info drop column `silver_exchange_time` ;
alter table player_global_clique_info drop column `gold_exchange_num`;
alter table player_global_clique_info drop column `silver_exchange_num` ;
alter table player_global_clique_info drop column `donate_coins_center_building`;
alter table player_global_clique_info drop column `donate_coins_temple_building`;
alter table player_global_clique_info drop column `donate_coins_bank_building`; 
alter table player_global_clique_info drop column `donate_coins_health_building`;
alter table player_global_clique_info drop column `donate_coins_attack_building`;
alter table player_global_clique_info drop column `donate_coins_defense_building`;
alter table player_global_clique_info drop column `health_level`;
alter table player_global_clique_info drop column `attack_level`;
alter table player_global_clique_info drop column `defense_level`;
alter table player_global_clique_info drop column `worship_time`;
alter table player_global_clique_info drop column `donate_coins_center_building_time`; 
alter table player_global_clique_info drop column `glod_exchange_time`; 
alter table player_global_clique_info drop column `worship_type`; 

");

?>
