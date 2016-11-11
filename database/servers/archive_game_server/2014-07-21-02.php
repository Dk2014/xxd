<?php
db_execute($db,"
alter table `trader_grid` drop column `stock`;
"
);
?>
