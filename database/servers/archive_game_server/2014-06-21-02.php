<?php

db_execute($db, "

update `item` set `equip_role_id` = -1 where `equip_role_id` = -2;

");
?>
