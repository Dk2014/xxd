<?php
db_execute($db,"
update `func` set `lock`='1400' where `id`='5';
update `func` set `lock`='1800' where `id`='11';
"
);
?>
