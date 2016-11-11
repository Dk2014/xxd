<?php
db_execute($db,"
alter table `player_coins` modify column `batch_bought` smallint(6) not null default '0' comment '玩家批量购买铜币次数';
"
);
?>

