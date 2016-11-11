<?php
db_execute($db,"


update extend_level set level_type = 1;

INSERT INTO `func` (`name`, `sign`, `lock`, `level`, `unique_key`, `need_play`) VALUES
 ('伙伴关卡', 'FUNC_BUDDY_LEVEL', 6000, 0, 256, 1),
 ('灵宠关卡', 'FUNC_PET_LEVEL', 7000, 0, 512, 1),
 ('魂侍关卡', 'FUNC_GHOST_LEVEL', 8000, 0, 1024, 1);

"
);
?>
