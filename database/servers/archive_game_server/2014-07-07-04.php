<?php
db_execute($db, "

ALTER TABLE player_arena_rank RENAME TO player_global_arena_rank;

");
?>
