<?php
db_execute($db,"


update `func` set `lock` = 6000 where `id` in(11,12,13);

"
);
?>
