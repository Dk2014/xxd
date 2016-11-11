<?php
db_execute($db, 
	"
		ALTER TABLE `item` ADD COLUMN `equip_role_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '可装备角色ID';

	"
);

?>