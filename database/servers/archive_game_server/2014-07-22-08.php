<?php
db_execute($db,"
alter table `daily_sign_in_award` drop column `daily_sign_in_id`;
"
);
?>
