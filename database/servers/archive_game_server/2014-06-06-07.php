
<?php
db_execute($db, 
"
  update `mission_level` set `award_box` = 1 where `parent_type` = 0;
"

);
?>