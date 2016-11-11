<?php
db_execute($db,"
update func set need_play=0 where id in (6, 10, 11, 12, 13);
"
);
?>
