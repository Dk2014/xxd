<?php
db_execute($db, "
RENAME TABLE `mission_level_meng_yao` TO `mission_level_recovery_meng_yao`;
ALTER TABLE `mission_level_recovery_meng_yao`  MODIFY  my_dir varchar(20) NOT NULL COMMENT '关卡梦妖朝向 r=>右, rb=>右下方, b=>下, lb=>左下方, l=>左, lt=>左上方, t=>上, rt=>右上方';
");
?>