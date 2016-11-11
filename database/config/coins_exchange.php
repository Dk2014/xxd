<?php
$config = array(
	array( 'const',
		'INGOT_TO_CONIS_CRIT_RATE' => 30, //购买铜币暴击概率%
		'DAILY_COINS_BUY_COUNT' => 1, //每日购买铜币次数限制
	),

	array('const',
		'FIRST_TIME_BATCH_BUY_COINS_NUM' => 3, 	//每日首次批量购买铜币次数
		'BATCH_BUY_COINS_NUM' => 5,		//默认批量购买铜币次数

		'COINS_BATCH_BUY_REQUIRED_LEVEL' => 0	//批量购买铜币要求等级
	),
);
?>
