<?php

db_execute($db, "
alter table `player_sword_soul` modify column `pos` tinyint(4) NOT NULL COMMENT '是否已装备  1-已装备 0-未装备';

update `player_sword_soul` set `pos`='0' where `pos`!='-1';

update `player_sword_soul` set `pos`='1' where `pos`='-1';

");
?>
