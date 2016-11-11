<?php


// 此文件修复线上8服的开服时间配置错误的问题。
// 目前线上已修复，注释修复代码的原因是为了线上所有服能进行平滑数据升级，所以占用了此版本号
// $this->AddSQL("
// delete  from `player_daily_sign_in_record`;
// update `player_daily_sign_in_state` set `record`=0, `update_time`=0, `signed_today`=0;
// ");

?>
