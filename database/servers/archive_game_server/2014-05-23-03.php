<?php
db_execute($db, 
	"
		ALTER TABLE `item` CHANGE `sign` `sign` varchar(30) DEFAULT NULL COMMENT '资源标识' AFTER `price`;

	"
);

?>