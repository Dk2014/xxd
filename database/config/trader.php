<?php
$config = array(
	//请勿改变顺序
	array('const'), //商人ID常量
	array('data', 'RefreshableTrader' => '[]int16'),//可刷新的商人列表
	//请勿改变顺序

	//消费类型
	array( 'const',
		'COINS_TYPE' => 0, //铜币
		'INGOT_TYPE' => 1, //元宝
		'HEART_TYPE' => 2, //爱心
		'DRAGON_COIN_TYPE' => 3, //龙币
	),

	array('data', 'YingHaiJiShiRefresh' => '[]int64', 12*60*60, 18*60*60, 21*3600), // 12,18,21点刷新
	array('data', 'XunYouShangRenRefresh' => '[]int64', 10*60*60, 22*60*60), //10, 22点刷新
	array('data', 'HeiShiLaoDaRefresh' => '[]int64'), //废弃
	//测试用
	//array('data', 'YingHaiJiShiRefresh' => '[]int64'), 
	//array('data', 'XunYouShangRenRefresh' => '[]int64'),


	array('const',
		'ITEM' => 0, //物品
		'SWORD_SOUL' => 1, //剑心
		'GHOST' => 2, //魂侍
		'PET' => 3, //灵宠
		'HEART' => 4, //爱心
		'COINS' => 5, //铜币
		'INGOT' => 6, //元宝
		'PHYSICAL' => 7, //体力
		'EQUIPMENT' => 8, //装备
	),

	array('const', 
		'NOTIFY_IN_ADVANCE' => 300,	//提前 5分钟通知
		'YINGHAISHANGREN_REQUIRE_LEVEL' => 14 //瀛海商人要求14级开放
	), 

);

$db = db_connect();
$traders = db_query($db, 'select * from trader');
foreach($traders as $trader) {
	$config[0][$trader['sign']] = intval($trader['id']);
	if( $trader['sign'] != 'XunYouShangRen') {
		$config[1][] = intval($trader['id']);
	}
}

////瀛海集市和黑市老大每个10分钟刷新一次
//for($i=600; $i<86400; $i=$i+600) {
//	$config[2][] = $i;
//	$config[4][] = $i;
//}
////巡游商人每隔20分钟刷新一次
//for($i=3600; $i<86400; $i=$i+3600) {
//	$config[3][] = $i;
//}
?>
