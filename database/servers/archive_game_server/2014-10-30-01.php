<?php

db_execute($db, "
  alter table ghost modify column `quality` tinyint(4) NOT NULL DEFAULT '0' COMMENT '魂侍品质';
");

?>
