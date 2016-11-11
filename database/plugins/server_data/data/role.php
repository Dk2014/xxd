<?php

function GetImport_role() {
  return array();
}

function GetCode_role($db) {
  return "";
}

/* old example */

// function GetImport_demo() {
// 	return array("fmt");
// }
//
// function GetCode_demo($db) {
// 	return get_const($db);
// }
//
//
// function get_const($db)
// {
// 	$code = <<<CONST
// 		// 定义常量
// const(
// 		TEST_VAL1 = 1
// 		TEST_VAL2 = 2
// 	)
// CONST;
//
// 	return $code;
// }
?>