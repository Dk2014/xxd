<?php
db_execute($db, "

alter table `skill` change column `type` `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型(0 - 系统；1 - 主角；5 - 怪物；7 - 魂侍技1； 8 - 魂侍技2)';

");

?>
