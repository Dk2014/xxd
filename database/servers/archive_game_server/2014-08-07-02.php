<?php
db_execute($db, "
alter table `npc_talk` modify  `type` tinyint(4) NOT NULL COMMENT '对话类型 0--首次对话； 1--任务对话';
");

?>
