<?php

$this->AddSQL("
update player_item set refine_level=10 where refine_level>10;
");

?>
