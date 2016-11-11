<?php

function GetImport_player() {
	return array();
}


/* 声望系统 */
function getfameSystem($db){
	$code = "\n // 声望系统常量 \n";
	$code .= "const(\n";
	$systems = db_query($db,'SELECT `id`, `name`, `sign` FROM `fame_system`;');
	foreach ($systems as $system) {
		$code .= "	{$system['sign']}_SYSTEM = ".$system['id']." // {$system['name']}\n";
	}
	$code .= ")\n";
	return $code;
}

/* 功能权值常量 */
function getFuncLock($db){
	$code = "\n // 功能权值常量 \n";
	$code .= "const(\n";
	$funcs = db_query($db,'SELECT `name`, `sign`,`lock` FROM `func` WHERE `type`=1 ORDER BY `lock` ASC;');
	foreach ($funcs as $func) {
		$code .= "	{$func['sign']} = ".$func['lock']." // {$func['name']}\n";
	}
	$code .= ")\n";
	return $code;
}

/*按等级开放的功能*/
function getLevelFunc($db) {
	$code = "\n // 等级功能常量 \n";
	$code .= "const(\n";
	$funcs = db_query($db,'SELECT `id`, `name`, `sign`,`level` FROM `func` where `type`=2 ORDER BY `level`;');
	foreach ($funcs as $func) {
		$code .= "	{$func['sign']} = ".$func['id']." // {$func['name']} 要求等级 {$func['level']}\n";
	}
	$code .= ")\n";
	return $code;
}

/*功能 sign*/
function getFuncSign($db) {
	$code = "\n // 功能常量 \n";
	$code .= "const(\n";
	$funcs = db_query($db,'SELECT `id`, `name`, `type`, `sign`, `lock`, `level` FROM `func` ORDER BY `lock`, `level`;');
	foreach ($funcs as $func) {
		$type = "要求";
		switch ($func['type']) {
		case 1:
			$type .= "权值 {$func['lock']}";
			break;
		case 2:
			$type .= "等级 {$func['level']}";
			break;
		default:
			die('位置的开放类型');
		}
		$code .= "	{$func['sign']} = ".$func['id']." // {$func['name']} {$type}\n";
	}
	$code .= ")\n";
	return $code;
}

/*开元通宝 */
function getIngotToCoinsSign() {
        $code = "\n //开元通宝 \n";
	$code .= "var(\n";
	$code .= "	INGOT_TO_COINS_DOUBLE_NUMS = [6]int{10, 15, 20, 30, 40, 50}\n";
	$code .= ")\n";
	return $code;
}


function GetCode_player($db) {
	$html = '';
	$html .= getFuncSign($db);
	$html .= "\n\n";
	$html .= getfameSystem($db);
	$html .= "\n\n";
	$html .= getIngotToCoinsSign();
	return $html;
}
?>
