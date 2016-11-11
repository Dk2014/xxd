<?php
db_execute($db, "
DROP TABLE IF EXISTS `mission_level_meng_yao`;
CREATE TABLE `mission_level_meng_yao`(
`id` int(10) NOT NULL AUTO_INCREMENT COMMENT '关卡梦妖标示',
`mission_level_id` int(11) NOT NULL COMMENT '关卡id',
`my_x` int(11) NOT NULL COMMENT '关卡梦妖x坐标',
`my_y` int(11) NOT NULL COMMENT '关卡梦妖y坐标',
`probability` tinyint(4) NOT NULL COMMENT '梦妖出现概率',
`my_effect` tinyint(4) NOT NULL COMMENT '关卡梦妖效果，1-恢复灵宠使用次数 2-恢复全体生命 3-恢复伙伴绝招次数 ',
`my_dir` tinyint(4) NOT NULL COMMENT '关卡梦妖朝向 1=>右, 2=>右下方, 3=>下, 4=>左下方, 5=>左, 6=>左上方, 7=>上, 8=>右上方',
`talk` text NOT NULL COMMENT '梦妖对话内容',
 PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='关卡内的梦妖配置';
");
?>