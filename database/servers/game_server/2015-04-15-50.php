<?php
$this->AddSQL("create index item_id_refine on player_item (item_id,refine_level_bak)");
?>