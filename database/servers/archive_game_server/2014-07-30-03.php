<?php
db_execute($db, "
	
	CREATE TABLE `server_info` (
  `version` varchar(10) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

");
?>
