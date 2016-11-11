<?php

db_execute($db, "
rename table battle_pet_grid_exp to battle_pet_grid_level;

");
?>
