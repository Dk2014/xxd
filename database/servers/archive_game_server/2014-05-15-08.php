<?php
db_execute($db, "
	ALTER TABLE `item` DROP COLUMN `name_prefix`;
");
?>