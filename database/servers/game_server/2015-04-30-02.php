<?php
$this->AddSQL("
ALTER TABLE `player_global_clique_info` change `exchange_time` `silver_exchange_time` bigint(20) NOT NULL DEFAULT 0;
");
?>