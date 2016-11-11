<?php
db_execute($db,"
alter table `player_trader_store_state` add column `grid_id` int(11) NOT NULL  COMMENT '格子ID';
"
);
?>
