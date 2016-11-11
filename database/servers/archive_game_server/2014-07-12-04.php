<?php
db_execute($db,"
alter table `player_town` add column  `at_pos_x` smallint(6) NOT NULL COMMENT '当前城镇的X轴位置';
alter table `player_town` add column  `at_pos_y` smallint(6) NOT NULL COMMENT '当前城镇的y轴位置';
"
);
?>
