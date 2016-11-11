<?php
db_execute($db, "

ALTER TABLE `ghost_umbra`
ADD COLUMN `fragment_min_num`  smallint(6) NOT NULL COMMENT '碎片随机最小值' AFTER `fragment_probability`,
ADD COLUMN `fragment_max_num`  smallint(6) NOT NULL COMMENT '碎片随机最大值' AFTER `fragment_min_num`;

");