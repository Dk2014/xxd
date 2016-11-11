<?php
$config = array(
	array('const',
		'TARGET_TYPE_SELF'            =>1, // 道具使用者
		'TARGET_TYPE_US_ALL'          =>2, // 我方全体
		'TARGET_TYPE_US_LEAST_HTH'    =>3, // 我方血量最少
		'TARGET_TYPE_ENEMY_ALL'       =>4, // 敌方全体
		'TARGET_TYPE_ENEMY_LEAST_HTH' =>5, // 敌方血量最少
		'TARGET_TYPE_US_DEAD'         =>6, // 我方阵亡角色
		'TARGET_TYPE_BATTLE_PET'      =>7, // 灵宠
	),

	array('const',
		'EFFECT_MODE_POWER' => 0, //精气
		'EFFECT_MODE_HEALTH' => 1, //生命
		'EFFECT_MODE_RELIVE' => 3, //复活
		'EFFECT_MODE_CATCH_PET' => 4, //灵宠捕捉
	),

	array('const',
		'BATTLE_RECOVER_POWER' => 3, //战场主角自动恢复精气，修改时同时修改 battle/battle_const.go 里面的 RECOVER_POWER
	)
);
?>
