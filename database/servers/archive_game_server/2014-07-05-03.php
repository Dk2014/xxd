
<?php
db_execute($db, "
update `hard_level` set `town_id`=1 where `id` between 1 and 5;
");
?>
