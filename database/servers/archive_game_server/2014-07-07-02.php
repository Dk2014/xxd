<?php
db_execute($db, "
create unique index unique_privilege_each_level on vip_privilege_config (privilege_id, level);
");
?>
