<?php
db_execute($db, "
ALTER TABLE `mission_level_box` ADD COLUMN `must_in_first` tinyint(4) default 0 COMMENT '是否第一次通关必然获得，0-否 1-是';
ALTER TABLE `player_hard_level_record` ADD COLUMN `pass_of_times` int(12) default 0 COMMENT '通关次数';
");