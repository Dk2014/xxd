<?php

function GetImport_arena_award_box() {
	return array();
}

function GetCode_arena_award_box($db) {
	$configs = db_query($db, "select * from arena_rank_gap order by rank asc");
	$code = "//自动生成请勿修改\n";
	$code .= "func GetArenaGapByRank(rank int32) (gap int32) {\n";
	foreach($configs as $config) {
		$code .= "	if rank <= {$config['rank']} {\n";
		$code .= "		return {$config['gap']}\n";
		$code .= "	}\n";
	}
	$code .= "	return 1\n";
	$code .= "}\n";

	return $code;
}

?>
