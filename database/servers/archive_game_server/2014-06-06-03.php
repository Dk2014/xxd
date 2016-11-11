
<?php
db_execute($db, 
"

ALTER TABLE `mission_level` CHANGE `parent_type` `parent_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '关联关卡类型(0--无;1--资源关卡;2--通天塔)';

"

);
?>