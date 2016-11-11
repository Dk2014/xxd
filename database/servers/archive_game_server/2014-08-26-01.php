<?php
db_execute($db, "
delete from item where type_id=9 and id in (300, 301);
");

?>
