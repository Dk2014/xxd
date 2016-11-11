
<?php
db_execute($db, 
"
ALTER TABLE `mission_level` ADD COLUMN `parent_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '关联关卡类型(0--无;1--资源关卡;)' after `mission_id`;
"
);
?>