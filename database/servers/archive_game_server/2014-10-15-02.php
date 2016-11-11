<?php

db_execute($db, "
DROP TABLE IF EXISTS `quest_activity_center`;
CREATE TABLE `quest_activity_center`(
    `id` smallint AUTO_INCREMENT ,
    `relative` smallint NOT NULL COMMENT '关联的活动',
    `weight` int(8) NOT NULL DEFAULT 0 COMMENT '活动权值',
    `name` varchar(60) NOT NULL COMMENT '活动名称(列表左侧)',
    `title` varchar(100) NOT NULL COMMENT '活动标题(列表右侧)',
    `content` text NOT NULL COMMENT '活动描述',
    `start` tinyint(4) DEFAULT 0 COMMENT '活动开始天数点',
    `end` tinyint(4) DEFAULT 0 COMMENT '活动结束天数点',
    `dispose_days` smallint DEFAULT 0 COMMENT '活动销毁天数点',
    `is_go` tinyint(4) COMMENT '是否前往',
    `tag` tinyint(4) COMMENT '活动标签(1:最新,2:限时,3:推荐)',
    `is_mail` tinyint(4) DEFAULT 0 COMMENT '活动结束是否补发奖励',
    PRIMARY KEY(`id`)
    )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='任务活动中心';
");
?>
