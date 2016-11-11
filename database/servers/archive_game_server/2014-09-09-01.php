<?php
db_execute($db, "
delete from rainbow_level where `segment` in (4,5,6,7);
")
?>
