<?php

db_execute($db, "

ALTER TABLE `battle_pet` ADD COLUMN `item_battle_pet_id` int(11) NOT NULL COMMENT '灵宠契约球';


");
?>
