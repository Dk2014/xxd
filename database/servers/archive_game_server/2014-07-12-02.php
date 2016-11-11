<?php
db_execute($db,"
delete from sword_soul_level where sword_soul_id in (30, 33, 36);
delete from sword_soul where id in (30, 33, 36);
"
);
?>
