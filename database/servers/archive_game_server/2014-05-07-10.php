<?php

db_execute($db, "
	ALTER TABLE player_roles RENAME TO player_role;
");
?>
