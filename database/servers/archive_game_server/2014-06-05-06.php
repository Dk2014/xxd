
<?php
db_execute($db, 
"
ALTER TABLE `enemy_deploy_form` MODIFY COLUMN `battle_type` tinyint(4) NOT NULL COMMENT '战场类型(0--关卡;1--资源关卡;2--极限关卡)';

"

);
?>