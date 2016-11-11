<?php
db_execute($db,"
alter table `player_daily_sign_in_state` drop column `index`;
"
);
?>
