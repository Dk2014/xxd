<?php
db_execute($db, "

INSERT INTO `ghost_passive_skill` ( `level`, `name`, `sign`, `desc`) VALUES
( 20, '魂侍护盾', 'HunShiShouHu', '当生命少于30%时，自动施放吸收伤害的魂侍护盾，持续3回合。魂侍护盾吸收的伤害等于魂侍生命值的50%，与其他魂侍合计。魂侍护盾每次战斗只能触发1次。'),
( 30, '初始魂力', 'HunShiFaDongLv', '初始魂力5，与其他魂侍合计'),
( 40, '魂侍技能2级', 'HunShiJiNengErJi', '魂侍技能升级为2级'),
( 50, '魂侍技能2级', 'HunShiShouHuErJi', '当生命少于30%时，自动施放吸收伤害的魂侍护盾，持续3回合。魂侍护盾吸收的伤害等于魂侍生命值的100%，与其他魂侍合计。魂侍护盾每次战斗只能触发1次。'),
( 60, '初始魂力2级', 'HunShiFaDongLvErJi', '初始魂力10，与其他魂侍合计'),
( 70, '魂侍技能3级', 'HunShiJiNengSanJi', '魂侍技能升级为3级');

");
?>



