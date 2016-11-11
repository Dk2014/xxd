<?php

$equal_sql="
update `skill` set `child_kind`=(`child_kind`-2+1)%2+2 where `type`=1 and `role_id`=-2 and (`child_kind`=2 or `child_kind`=3);
"; // this sql's behaviour is equally to the below, leave it here for debug and 提升B格


$this->AddSQL("
-- swap two skill types, 体术(进阶) and 奥义
update `skill` set `child_kind`=5 where `type`=1 and `role_id`=-2 and `child_kind`=3;
update `skill` set `child_kind`=3 where `type`=1 and `role_id`=-2 and `child_kind`=2;
update `skill` set `child_kind`=2 where `type`=1 and `role_id`=-2 and `child_kind`=5;
");

$this->AddSQL("
-- change a few of skills's type
-- these two sql is from document our designer given to us. the `id` is auto incrementation column, so will be different if we recreate the data of this table, careful.
update `skill` set `child_kind`=3 where `id` in (1153, 1157, 1155);
update `skill` set `child_kind`=4 where `id` in (1227, 1226, 1139, 89, 1156, 1225);
");

/*$this->AddSQL*/("
-- now we changed lots of skill's type, we must unarm player's using skill
update `player_use_skill` set `skill_id1`=0, `skill_id2`=0, `skill_id3`=0
	where `role_id`=1 or `role_id`=2;
-- later we will choose the best we think for them
");

$this->AddSQL("
-- add 咒术 grid
alter table `player_use_skill` change column `skill_id0` `skill_id4` smallint(6) NOT NULL DEFAULT '0' COMMENT '招式4';
");

