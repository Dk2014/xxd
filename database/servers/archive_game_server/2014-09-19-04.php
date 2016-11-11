<?php
db_execute($db,"
UPDATE `func` SET `need_play` = 1 WHERE `sign` = 'FUNC_ARENA';
");
?>