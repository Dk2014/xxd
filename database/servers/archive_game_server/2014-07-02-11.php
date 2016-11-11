<?php
db_execute($db, "

INSERT INTO `ghost_tip` ( `tip`) VALUES
('战斗中，角色每次行动与受伤都会获得魂力\r\n当魂力满100时，可以召唤魂侍'),
('每只魂侍在同个关卡内只能被触发一次');

");
?>

