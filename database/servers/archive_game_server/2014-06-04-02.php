
<?php
db_execute($db, 
"
    TRUNCATE TABLE `sword_soul_quality`;
    INSERT INTO `sword_soul_quality` (`id`, `name`, `sign`, `init_exp`, `price`, `color`) VALUES
    (0, '杂物', 'SUNDRY', 0, 400, '0xc5c3b7'),
    (1, '特殊', 'SPECIAL', 1000, 0, '0xdb2e00'),
    (2, '优良', 'FINE', 50, 0, '0x22ac38'),
    (3, '精良', 'EXCELLENT', 200, 0, '0x00a0e9'),
    (4, '传奇', 'LEGEND', 500, 0, '0xc301c3'),
    (5, '神器', 'ARTIFACT', 1500, 0, '0xfff100'),
    (6, '唯一', 'ONLY', 1080, 0, '0xf39700');
"
);
?>