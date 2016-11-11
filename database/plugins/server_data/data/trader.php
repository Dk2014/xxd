<?php

function GetImport_trader() {
	return array(
		"core/mysql",
		"core/i18l",
	);
}



function GetCode_trader($db) {
	$traders = db_query($db, "SELECT * FROM trader");
	$code = '';

	$code .= "//出现时间变量\n";
	$code .="var (\n";
		$code .= "	MapTraderSchedule map[int16][]*TraderSchedule //出现时间配置\n";
	$code .=")\n\n";

	$code .= "//初始化加随机商人出现时间和刷新时间\n";
	$code .= "//每个商人默认会有一个24小时出现永不过期的配置\n";
	$code .= "func loadTraderSchedule(db *mysql.Connection) {\n";
	$code .= "	MapTraderSchedule = make(map[int16][]*TraderSchedule)\n";
	foreach ($traders as $trader) {
		$code .= "	//{$trader['name']}配置\n";
		if($trader['sign'] == 'XunYouShangRen') {
			$code .= "	//{$trader['name']}出现时间需要特别处理\n";

			//TODO 恢复正常出现时间
			$code .= "	MapTraderSchedule[{$trader['id']}] = make([]*TraderSchedule, 0, 2)\n";
			$code .= "	MapTraderSchedule[{$trader['id']}] = append(MapTraderSchedule[{$trader['id']}], &TraderSchedule{Expire: 0, Showup: 36000, Disappear: 79200})\n";
			
			//FIXME 测试用， 巡游商人没60分钟出现一次，持续59分钟
			/*
			$duration = 59*60;
			$code .= "	MapTraderSchedule[{$trader['id']}] = make([]*TraderSchedule, 0, 24)\n";
			for($i=3600; $i < 86400; $i+=3600) {
				$disappear = $i + $duration;
				$code .= "	MapTraderSchedule[{$trader['id']}] = append(MapTraderSchedule[{$trader['id']}], &TraderSchedule{Expire: 0, Showup: {$i}, Disappear: {$disappear}})\n";
			}
			*/
		} else {
			$code .= "	MapTraderSchedule[{$trader['id']}] = make([]*TraderSchedule, 0, 1)\n";
			$code .= "	MapTraderSchedule[{$trader['id']}] = append(MapTraderSchedule[{$trader['id']}], &TraderSchedule{Expire: 0, Showup: 0, Disappear: 86400})\n";
		}
	}
	$code .= "}\n";

	//查询刷新时间计划API
	$code .= "//根据商人ID查询刷新时间计划\n";
	$code .= "func TraderRefreshSchedule(traderId int16) []int64{\n";
	$code .= "	switch traderId {\n";
	foreach ($traders as $trader) {
		$code .= "	case {$trader['id']} :\n";
		$code .= "		return {$trader['sign']}Refresh\n";
	}
	$code .="	default:\n";
	$code .="		panic(\"undefine trader\")\n";
	$code .="	}\n";
	$code .= "}\n";

	//查询商人名称API
	$code .= "//根据商人ID查询商人名称\n";
	$code .= "func GetTraderNameById(traderId int16) string{\n";
	$code .= "	switch traderId {\n";
	foreach ($traders as $trader) {
		$code .= "	case {$trader['id']} :\n";
		$code .= "		return i18l.T.Tran(\"{$trader['name']}\")\n";
	}
	$code .="	default:\n";
	$code .="		panic(\"undefine trader\")\n";
	$code .="	}\n";
	$code .= "}\n";




	return $code;
}
