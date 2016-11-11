<?php
db_execute($db,"
update `func` set `lock`='1400' where `id`='4';
update `func` set `lock`='3000' where `id`='5';
"
);
?>
