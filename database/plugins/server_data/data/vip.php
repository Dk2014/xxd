<?php

function GetImport_vip() {
	return array();
}

function getPrivilege($db) {
	$code = "\n /* 特权常量说明 \n";
	$code .= "const(\n";
	$privileges =  db_query($db, "select * from vip_privilege");
	foreach($privileges as  $privilege) {
		$code .= "	{$privilege['sign']} = {$privilege['id']} // {$privilege['name']}\n";
	}
	$code .= ")\n";
	$code .= "*/\n";
	return $code;
}

function getSpecialPrivilege() {
	$code = "// 非所有版本都存在的特权 手动配置\n";
	$code .= "const(\n"; 
	$code .= "	FUDITEQUAN_SPECIAL = 45\n";
	$code .= ")\n";
	return $code;
}

function GetCode_vip($db) {
	$code = getSpecialPrivilege();
	$code .= "\n\n";
	$code .= getPrivilege($db);
	return $code;
}
?>
