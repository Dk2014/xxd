<?php
db_execute($db, "
	ALTER TABLE `town_npc` ADD COLUMN `showup_quest` smallint(6) NOT NULL DEFAULT '0' COMMENT '出现任务ID -1无效; 0一直有效' ;
	ALTER TABLE `town_npc` ADD COLUMN `disappear_quest` smallint(6) NOT NULL DEFAULT '0' COMMENT '消失任务ID -1无效; 0一直有效 ' ;
");
?>
