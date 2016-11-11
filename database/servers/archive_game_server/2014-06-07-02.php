
<?php
db_execute($db, 
"
ALTER TABLE `player_info` DROP `heart_num`;

ALTER TABLE `player_info` ADD COLUMN `new_mail_num`  smallint(6) NOT NULL DEFAULT '0'  COMMENT '新邮件数';
"

);
?>