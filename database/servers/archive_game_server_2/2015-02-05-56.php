<?php

$this->ADDSQL(
"DELETE FROM `player_daily_quest` WHERE quest_id NOT IN (SELECT id FROM daily_quest);"
);

?>