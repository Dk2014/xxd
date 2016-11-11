<?php
//比武场 改为18级开放
$this->AddSQL("update `func` set `level`=18, `lock`=0, `type`=2 where `id`=10 and name='比武场';");

//资源关卡 改为 25级开放
$this->AddSQL("update `func` set `level`=25, `lock`=0, `type`=2 where `id`=16 and name='资源关卡';");

//活动关卡修改 sign
$this->AddSQL("update `func` set `sign`='FUNC_ACTIVE_LEVLE' where `sign`='LEVEL_FUNC_HUO_DONG_GUAN_KA';");

//瀛海集市
$this->AddSQL("update `func` set `sign`='FUNC_TRADER' where `sign`='LEVEL_FUNC_YIN_HAI_JI_SHI';");

//扫荡
$this->AddSQL("update `func` set `sign`='FUNC_DIRECTLY_AWARD' where `sign`='LEVEL_FUNC_SAO_DANG';");

//打坐
$this->AddSQL("update `func` set `sign`='FUNC_MEDITATION' where `sign`='LEVEL_FUNC_DA_ZUO';");

//灵宠幻境
$this->AddSQL("update `func` set `sign`='FUNC_PET_VIRTUAL_ENV' where `sign`='LEVEL_FUNC_PET_VIRTUAL_ENV';");


//世界聊天
$this->AddSQL("update `func` set `sign`='FUNC_WORLD_CHANNEL' where `sign`='LEVEL_FUNC_SHI_JIE_LIAO_TIAN';");

//多人关卡
$this->AddSQL("update `func` set `sign`='FUNC_MULTI_LEVEL' where `sign`='LEVEL_FUNC_DUO_REN_GUAN_KA';");

?>
