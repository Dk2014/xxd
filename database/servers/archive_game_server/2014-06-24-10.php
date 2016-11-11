<?php

db_execute($db, "

INSERT INTO `sword_soul_type` (`id`, `name`, `sign`) VALUES
(1, '攻击', 'ATTACK '),
(2, '防御', 'DEFENCE '),
(3, '生命', 'HEALTH '),
(4, '速度', 'SPEED '),
(5, '内力', 'CULTIVATION '),
(6, '命中', 'HIT_LEVEL '),
(7, '暴击', 'CRITICAL_LEVEL'),
(8, '格挡', 'BLOCK_LEVEL '),
(9, '破击', 'DESTROY_LEVEL '),
(10, '韧性', 'TENACITY_LEVEL'),
(11, '闪避', 'DODGE_LEVEL'),
(12, '护甲', 'SUNDER_MAX_VALUE'),
(13, '剑心经验', 'EXP');

");
?>
