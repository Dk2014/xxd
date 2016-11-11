<?php
//难度关卡
$this->AddSQL("update `func` set `sign`='FUNC_HARD_LEVEL', `level`=0, `lock`=1050, `type`=1 where `id`=14 and name='难度关卡';");

//资源关卡
$this->AddSQL("update `func` set `sign`='FUNC_RESOURCE_LEVEL', `level`=0, `lock`=1100, `type`=1 where `id`=16 and name='资源关卡'");

//比武场
$this->AddSQL("update `func` set `lock`=1300 where `id`=10 and name='比武场'");

//神龙宝箱
$this->AddSQL("update `func` set `lock`=1400 where `id`=13 and name='神龙宝箱';");

//极限关卡
$this->AddSQL("update `func` set `level`=20 where `id`=19 and name='极限关卡';");


?>
