<?php

db_execute($db, "
UPDATE `func` SET `lock` = 1450 WHERE `sign` = 'FUNC_BATTLE_PET';
");

?>
