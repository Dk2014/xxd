
<?php
db_execute($db, "

DROP TABLE IF EXISTS `player_ghost_equipment`;

ALTER TABLE `player_ghost` ADD COLUMN `pos` smallint(6) NOT NULL COMMENT '位置';

ALTER TABLE `player_ghost` DROP COLUMN `is_equip`;

");
?>