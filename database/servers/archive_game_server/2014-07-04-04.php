<?php
db_execute($db, "
alter table `player_coins` add column
batch_bought tinyint(1) not null default '0'
comment '是否进行过批量购买';
");
?>
