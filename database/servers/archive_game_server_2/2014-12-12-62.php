<?php 
$this->AddSQL(
"update player_town set `lock`=100150 where pid in (select pid from player_quest where `quest_id`=1016);"
);
?>

