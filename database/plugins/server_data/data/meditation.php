<?php

function GetImport_meditation() {
	return array();
}

function GetCode_meditation($db) {
	$code = "\n//自动生成请勿修改\n//打坐经验\n";
	$code .= "func GetMetationExpByLevel(level int16) int32 {\n";
	$rows = db_query($db, "select * from `meditation_exp` order by `require_level` desc");
	foreach($rows as $row) {
		$code .= "	if level >= {$row["require_level"]} {\n";
		$code .= "		return {$row["exp"]}\n";
		$code .= "	}\n";
	}
	$code .= "	return 0\n";
	$code .= "}\n";

	return $code;
}

?>
