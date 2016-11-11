<?php
$this->ADDSQL("
update player_item set refine_level=10 where refine_level=5;
update player_item set refine_level=8 where refine_level=4;
update player_item set refine_level=6 where refine_level=3;
update player_item set refine_level=4 where refine_level=2;
update player_item set refine_level=2 where refine_level=1;
"
);
?>
