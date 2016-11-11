<?php

$this->AddSQL("
insert into player_global_clique_building (select `pid`, `silver_exchange_time`, `gold_exchange_num`, `silver_exchange_num`, `donate_coins_center_building`, `donate_coins_temple_building`, `donate_coins_bank_building`, `donate_coins_health_building`, `donate_coins_attack_building`, `donate_coins_defense_building`, `health_level`, `attack_level`, `defense_level`, `worship_time`, `donate_coins_center_building_time`, `glod_exchange_time`, `worship_type` from player_global_clique_info);

");

?>
