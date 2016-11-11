<?php
db_execute($db,"
alter table `mission_enemy` drop column `talk`;
"
);
?>
