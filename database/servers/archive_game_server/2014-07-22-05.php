<?php
db_execute($db,"

alter table `daily_quest` add column `level_type` tinyint(4) NOT NULL DEFAULT '-1' comment '关卡类型; -1 无; 0-区域关卡;1-资源关卡;2-通天塔;8-难度关卡;9-伙伴关卡;10-灵宠关卡;11-魂侍关卡';
alter table `daily_quest` add column `level_sub_type` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '关卡子类型(-1--无;1--铜钱关卡;2--经验关卡)';

"
);
?>
