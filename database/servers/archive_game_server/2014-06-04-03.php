
<?php
db_execute($db, 
"
ALTER TABLE `enemy_role` ADD COLUMN `skill_config` text COMMENT  '技能配置';
"
);
?>