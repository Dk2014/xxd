<?php
db_execute($db, 
	"
	ALTER TABLE `ghost_level` CHANGE need_crystal_num need_fruit_num int(11) NOT NULL DEFAULT '0' COMMENT '所需影界果实数量';
	"
);
?>