<?php
db_execute($db, "

update `ghost_train` set exp = exp/100, min_add_exp = min_add_exp/100, max_add_exp = max_add_exp/100;

");
?>
