
<?php
db_execute($db, 
	"
		ALTER TABLE `player_item` ADD COLUMN `buy_back_state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '记录物品是否在回购栏';

	"
);

?>