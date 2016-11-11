<?php
$config = array(
	array( 'const',
		/*货币类型*/
		'BUY_PHYSICAL_PRICE' => 50, // 购买体力价格
		'BUY_PHYSICAL_VALUE' => 200, // 购买体力后获得的体力点数
		'PHYSICAL_RECOVER_SECOND' => 5*60, // 体力恢复时间间隔
		'PER_RECOVER_PHYSICAL_VALUE' => 1, // 每次恢复体力
		'DAILY_PHYSICAL_BUY_COUNT' => 1,  // 每日允许购买次数
		'MAX_PHYSICAL_VALUE_BY_TIME' => 200, // 玩家自动回复体力上限为200
		'MAX_PHYSICAL_VALUE' => 2000, // 玩家体力上限
	),
);
?>
