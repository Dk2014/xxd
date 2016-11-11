<?php

db_execute($db, "

CREATE TABLE `dbupgrade_version` (
  `version` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

");

?>
