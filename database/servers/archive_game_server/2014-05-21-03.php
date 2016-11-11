<?php
db_execute($db, "
	ALTER TABLE `role_level` DROP column `exp`;
");
?>