<?php
db_execute($db, "
alter table  `player_rainbow_level` add column `best_segment` smallint(6) NOT NULL DEFAULT '0' COMMENT '最好记录段数';
alter table  `player_rainbow_level` add column `best_order` tinyint(4) NOT NULL DEFAULT '0' COMMENT '最好记录关卡顺序';
alter table  `player_rainbow_level` add column `best_record_timestamp` bigint(20) NOT NULL DEFAULT '0' COMMENT '最好记录时间戳';
")
?>
