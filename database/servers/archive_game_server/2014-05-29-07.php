<?php
db_execute($db, 
"
    DELETE FROM `sword_soul_level` WHERE `level` = 0;
"
);

?>