<?php
db_execute($db, "

alter table `skill` change column `child_type` `child_kind` tinyint(4) DEFAULT '1' NOT NULL COMMENT '绝招子类型 1－基本 2-进阶 3-奥义' after `effect`;

alter table `skill` change column `effect` `child_type` smallint(6) DEFAULT '1' NOT NULL COMMENT '绝招作用 1-攻击 3-防御 4-治疗 5-辅助 6-破甲';
");

?>
