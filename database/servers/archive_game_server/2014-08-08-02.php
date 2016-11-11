<?php
db_execute($db, "
alter table `player_npc_talk_record` add column quest_id smallint(6) DEFAULT '-1' NOT NULL COMMENT '任务ID  首次对话作为特殊任务ID -1';

alter table `player_npc_talk_record` comment='玩家与NPC对话奖励记录';
");

?>
