<?php

db_execute($db, "
ALTER TABLE `dbupgrade_version` MODIFY COLUMN `version` int(11) DEFAULT 0;
");

?>
