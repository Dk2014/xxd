<?php

function GetImport_push_notify() {
	return array("core/i18l");
}

function GetCode_push_notify($db) {
	$notifications = db_query($db, "select * from push_notify where `trigger_time`='-1'");
	$code = "";

	$code .= "//自动生产请勿修改\n";
	$code .= "func Load(){\n";
	$code .= "	mapPushNotify = make(map[int32]*PushNotify)\n";
	foreach($notifications as $n) {
		$code .= "	mapPushNotify[{$n['id']}] = &PushNotify {\n";
		$code .= "		Content: i18l.T.Tran(\"{$n['content']}\"),\n";
		$code .= "		Name: i18l.T.Tran(\"{$n['name']}\"),\n";
		$code .= "	}\n";
		$code .= "\n";
	}
	$code .= "}\n";




	return $code;
}

?>
