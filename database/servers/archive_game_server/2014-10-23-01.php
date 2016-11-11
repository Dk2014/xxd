<?php

db_execute($db, "
CREATE TABLE `event_multiply_config`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `condition` varchar(30) NOT NULL COMMENT '条件',
    `times` tinyint(4) NOT NULL COMMENT '倍数',
    PRIMARY KEY(`id`),
    UNIQUE KEY(`condition`)
    )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='翻倍活动配置';
");

?>
