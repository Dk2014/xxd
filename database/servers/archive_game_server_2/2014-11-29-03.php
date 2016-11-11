<?php

//朱媛媛 治愈之风2级 表现改变
$this->AddSQL("
update `skill` set `target`='1' where `id`='1142';
");

//青丘山开启钥匙改为4把
$this->AddSQL("
update `mission` set `keys`=4 where `id`='8';
");

//多人关卡30级开放
$this->AddSQL("
update `level_func` set `level`='30', `need_play`=1 where `id`='5';
");

//深渊关卡产出调整
$this->AddSQL("
update `mission_level_box` set `item_id`='418' where `id`='1883';
update `mission_level_box` set `item_id`='419' where `id`='1963';
update `mission_level_box` set `item_id`='420' where `id`='2013';
");

//胧月活动更新
$this->AddSQL("
	update `quest_activity_center` set `content`='初始等级[level]级，东瀛女忍者，性情阴晴不定，花费[ingot]即可邀请胧月加入队伍(赠送专属紫色武器及魂侍将通过邮件发放）', `is_relative`='1', `start`='0', `end`='604800', `dispose`='604800', `weight`='0' where `id`='24';
");


?>
