<?php

function GetImport_sword_soul() {
	return array();
}

function GetCode_sword_soul($db) {
	$code = "";
	$code .= getSwordSoulType($db);
	$code .= getSwordSoulQuality($db);
	$code .= genSwordSoulAllProbability($db);
	return $code;
}

function getSwordSoulType($db) {
	$sword_soul_type = db_query($db, 'SELECT id,sign FROM sword_soul_type');

	$code = "const (\n";
	foreach ($sword_soul_type as $key => $value) {
		$code .= "\tTYPE_".$value['sign']." =  {$value['id']}\n";
	}
	$code .= "\tTYPE_NUM = ".count($sword_soul_type)."\n";
	$code .= ")\n";

	return $code;
}

function getSwordSoulQuality($db) {
	$sword_soul_q = db_query($db, 'SELECT id,sign FROM sword_soul_quality');

	$code = "const (\n";
	foreach ($sword_soul_q as $key => $value) {
		$code .= "\tQUALITY_".$value['sign']." =  {$value['id']}\n";
	}
	$code .= "\tQUALITY_NUM = ".count($sword_soul_q)."\n";
	$code .= ")\n";

	return $code;
}

function genSwordSoulAllProbability($db) {
	$db_table = db_query($db, 'select * from `sword_soul_probability` where `id` in (select max(`id`) from `sword_soul_probability` group by `order`) order by `order`');
	$code = "";
	$code .= genNextBoxProbability($db_table);
	$code .= genSwordSoulProbability($db_table);
	return $code;
}

function genNextBoxProbability($db_table) {
	$code = 'var NextBoxProbability = []float64{';
	for ($i = 0; $i < 5; $i++) {
		$code .= 'float64('.$db_table[$i]['level_up'].')/100, ';
	}
	$code .= "}\n";
	$code .= 'var NextBoxProbability_VIP = []float64{';
	for ($i = 5; $i <7; $i++) {
		$code .='float64('.$db_table[$i]['level_up'].')/100, ';
	}
	$code .= "}\n";

	return $code;
}

function genSwordSoulProbability($db_table) {
	$code = "// Order: rubbish, exp, green, blue, purple, yello\n";
	$code .= "var SwordSoulProbability = [][]float64{\n";
	for ($i = 0; $i <5; $i++) {
		$total = $db_table[$i]['get_rubbish'] + $db_table[$i]['get_exp'] + $db_table[$i]['green'] + $db_table[$i]['blue'] + $db_table[$i]['purple'] + $db_table[$i]['yello'];

		$code .= "\t".'{';
		$code .= 'float64('.$db_table[$i]['get_rubbish'].')/'.$total.', ';
		$code .= 'float64('.$db_table[$i]['get_exp'].')/'.$total.', ';
		$code .= 'float64('.$db_table[$i]['green'].')/'.$total.', ';
		$code .= 'float64('.$db_table[$i]['blue'].')/'.$total.', ';
		$code .= 'float64('.$db_table[$i]['purple'].')/'.$total.', ';
		$code .= 'float64('.$db_table[$i]['yello'].')/'.$total.', ';
		$code .= "},\n";
	}
	$code .= "}\n";

	$code .= "var SwordSoulProbability_VIP = [][]float64{\n";
	for ($i = 5; $i <7; $i++) {
		$total = $db_table[$i]['get_rubbish'] + $db_table[$i]['get_exp'] + $db_table[$i]['green'] + $db_table[$i]['blue'] + $db_table[$i]['purple'] + $db_table[$i]['yello'];

		$code .= "\t".'{';
		$code .= 'float64('.$db_table[$i]['get_rubbish'].')/'.$total.', ';
		$code .= 'float64('.$db_table[$i]['get_exp'].')/'.$total.', ';
		$code .= 'float64('.$db_table[$i]['green'].')/'.$total.', ';
		$code .= 'float64('.$db_table[$i]['blue'].')/'.$total.', ';
		$code .= 'float64('.$db_table[$i]['purple'].')/'.$total.', ';
		$code .= 'float64('.$db_table[$i]['yello'].')/'.$total.', ';
		$code .= "},\n";
	}
	$code .= "}\n";

	return $code;
}

?>
