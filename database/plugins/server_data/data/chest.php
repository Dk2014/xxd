<?php

function GetImport_chest() {
	return array();
}

function GetCode_chest($db) {
	$code = '';
	$code .= getFateBoxConst($db);

	$code .= "const (\n";
	$code .= "	SPECIAL_DRAW_COUNT = 20 // 前20次开宝箱特殊处理\n";
	$code .= ")\n";


	return $code;
}

function getFateBoxConst($db) {
	$fateBoxQuery =  db_query($db, "select * from fate_box");
	$code = '';
	$code = "/*宝箱常量*/\n";
	$code .= "const (\n";
	foreach($fateBoxQuery as  $fateBox) {
		$code .= "	{$fateBox['sign']} = {$fateBox['id']} //{$fateBox['name']}\n";
		$code .= "	{$fateBox['sign']}Lock = {$fateBox['award_lock']} //{$fateBox['name']}权值\n";
	}
	$code .= ")\n";
	return $code;
}


?>
