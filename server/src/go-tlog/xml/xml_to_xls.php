<?php
$db = mysql_connect ( '42.120.22.64', "rlms", "xxdrlms" );
$ok = mysql_query('use xxd_dev20140919;' , $db);

$string = file_get_contents("log_template.xml");
$xml = simplexml_load_string($string);
$str = "";
$tx="imodeid\tvmodename\timodetype\timodetypname\t\n金钱流向ID\t金钱流向名称\t经前流向分类(0-产出;1-消耗)\t金钱流向分类(产出或消耗)\t\n"; 
$str .= $tx;
foreach ($xml->macrosgroup as $key => $value) {
	if($value->attributes()["name"] == "MoneyFlowReason")
	foreach ($value as $k => $v) {
		$str .= $v["value"]."\t".$v["desc"]."\t\t\t\n";
	}
}

header("Content-type:application/vnd.ms-excel");
header("Content-Disposition:attachment;filename=金钱产出消耗配置表xxd_rpt_tbmoneymodeconf.xls");
file_put_contents("金钱产出消耗配置表xxd_rpt_tbmoneymodeconf.xls", iconv('utf-8', 'gb2312', $str));

$string = file_get_contents("log_template.xml");
$xml = simplexml_load_string($string);
$str = "";
$tx="igoodstype\tigoodsid\tvgoodsname\tvgoodstypename\tigoodssubtype\tvgoodssubtypename\tiprice\tioriginalgoodstype\titerm\tioriginalterm\t\n"; 
$str .= $tx;
$tx="物品类型ID\t物品ID\t物品名称\t物品类型名称\t物品子类型ID(无此项留空)\t物品子类型名称(无此项留空)\t价格(无此项留空)\t物品原始类型(无此项留空)\t物品期限(无此项填0)\t物品原始期限(无此项填0)\t\n";
$str .= $tx;

foreach ($xml->macrosgroup as $key => $value) {
	if($value->attributes()["name"] == "Item")
	foreach ($value as $k => $v) {
		$str .= "12\t".$v["value"]."\t".$v["desc"]."\t物品\t\t\t\t\t\t\t\n";
	}
	if($value->attributes()["name"] == "Ghost")
	foreach ($value as $k => $v) {
		$str .= "-1\t".$v["value"]."\t".$v["desc"]."\t魂侍(完整)\t\t\t\t\t\t\t\n";
	}
	if($value->attributes()["name"] == "Sword")
	foreach ($value as $k => $v) {
		$str .= "-2\t".$v["value"]."\t".$v["desc"]."\t剑心(完整)\t\t\t\t\t\t\t\n";
	}

}

header("Content-type:application/vnd.ms-excel");
header("Content-Disposition:attachment;filename=物品名称配置表xxd_rpt_tbgoodsconf.xls");
file_put_contents("物品名称配置表xxd_rpt_tbgoodsconf.xls", iconv('utf-8', 'gb2312', $str));

$string = file_get_contents("log_template.xml");
$xml = simplexml_load_string($string);
$str = "";
$tx="iflowtype\tvtypename\tisubtype\tvsubtypename\t\n"; 
$str .= $tx;
$tx="物品流向类型ID\t物品流向类型名称\t物品流向类型子ID\t物品流向类型子名称\t\n";
$str .= $tx;

foreach ($xml->macrosgroup as $key => $value) {
	if($value->attributes()["name"] == "ItemFlowReason")
	foreach ($value as $k => $v) {
		$str .= $v["value"]."\t".$v["desc"]."\t\t\t\n";
	}
}

header("Content-type:application/vnd.ms-excel");
header("Content-Disposition:attachment;filename=道具流向配置表xxd_rpt_tbconfitemflowconf.xls");
file_put_contents("道具流向配置表xxd_rpt_tbconfitemflowconf.xls", iconv('utf-8', 'gb2312', $str));


$string = file_get_contents("log_template.xml");
$xml = simplexml_load_string($string);

$type_arr = array();
$item = array();
$sql = 'select * from item_type';
$result = mysql_query($sql, $db);
do {
	$row = mysql_fetch_object($result);
	$type_arr[$row->id] = $row->name;
} while($row);
$sql = 'select * from item';
$result = mysql_query($sql, $db);
do {
	$row = mysql_fetch_object($result);
	if (strlen($row->id)) {
		$item[$row->id] = array('name'=>$row->name,'type_id'=>$row->type_id);
	}
} while($row);

$str = "";
$tx="igoodstype\tigoodsid\tvgoodsname\tvgoodstypename\tigoodssubtype\tvgoodssubtypename\tiprice\tioriginalgoodstype\titerm\tioriginalterm\t\n"; 
$str .= $tx;
$tx="物品类型ID\t物品ID\t物品名称\t物品类型名称\t物品子类型ID(无此项留空)\t物品子类型名称(无此项留空)\t价格(无此项留空)\t物品原始类型(无此项留空)\t物品期限(无此项填0)\t物品原始期限(无此项填0)\t\n";
$str .= $tx;

foreach ($xml->macrosgroup as $key => $value) {
	if($value->attributes()["name"] == "Item")
	foreach ($item as $k => $v) {
		$str .= "{$v['type_id']}\t".$k."\t".$v["name"]."\t{$type_arr[$v['type_id']]}\t\t\t\t\t\t\t\n";
	}
	if($value->attributes()["name"] == "Ghost")
	foreach ($value as $k => $v) {
		$str .= "-1\t".$v["value"]."\t".$v["desc"]."\t魂侍(完整)\t\t\t\t\t\t\t\n";
	}
	if($value->attributes()["name"] == "Sword")
	foreach ($value as $k => $v) {
		$str .= "-2\t".$v["value"]."\t".$v["desc"]."\t剑心(完整)\t\t\t\t\t\t\t\n";
	}

}

header("Content-type:application/vnd.ms-excel");
header("Content-Disposition:attachment;filename=itemmoneyflowid.xls");
file_put_contents("itemmoneyflowid.xls", iconv('utf-8', 'gb2312', $str));
?>