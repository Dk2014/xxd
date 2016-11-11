<?php
db_execute($db, "

INSERT INTO `sword_soul_type` (`id`, `name`, `sign`) VALUES
(1, '生命', 'MAX_HEALTH'),
(2, '攻击', 'ATTACK'),
(3, '防御', 'DEFEND'),
(4, '内力', 'CULTIVATION'),
(5, '速度', 'SPEED'),
(6, '命中等级', 'HIT_LEVEL'),
(7, '闪避等级', 'DODGE_LEVEL'),
(8, '暴击等级', 'CRITIAL_LEVEL'),
(9, '挡格等级', 'BLOCK_LEVEL'),
(10, '破击等级', 'DESTROY_LEVEL'),
(11, '韧性等级', 'TENACITY_LEVEL'),
(12, '意志等级', 'WILL_LEVEL'),
(13, '化魂', 'GHOST_SKILL_RATE'),
(14, '范围伤害免伤', 'AOE_REDUCE'),
(15, '单体伤害免伤', 'SINGLE_REDUCE'),
(16, '必杀等级', 'CRITIAL_HURT_LEVEL'),
(17, '护甲上限', 'SUNDER_MAX_VALUE'),
(18, '抵抗混乱等级', 'ANTI_RANDOM_LEVEL'),
(99, '剑心经验', 'EXP');

");
?>
