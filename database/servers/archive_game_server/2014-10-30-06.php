<?php

db_execute($db, "
alter table player_battle_pet modify column parent_pet_id int(11) NOT NULL COMMENT '父灵宠ID(怪物ID)';
alter table player_battle_pet modify column battle_pet_id int(11) NOT NULL COMMENT '灵宠ID(怪物ID）';

alter table player_battle_pet_grid drop column pet_id;
alter table player_battle_pet_grid add column battle_pet_id int(11) NOT NULL DEFAULT '0' COMMENT '灵宠ID';

");
?>
