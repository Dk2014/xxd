<?php
$config = array(

	array( 'const',
		'MAX_DAILY_NUM'      => 5, // 每日挑战次数

		'AWARD_ANYINGGUOSHI_WIN_HIGH_RANK' => 5, // 战胜高于自己排名的玩家奖励暗影果实数
		'AWARD_ANYINGGUOSHI_WIN_LOW_RANK'  => 3, // 战胜低于自己排名的玩家奖励暗影果实数
		'AWARD_ANYINGGUOSHI_LOSE_LOW_RANK'  => 1, // 挑战失败奖励龙币数

		'LOSE_CD_TIME_SECONDS' => 10*60, // 战败后冷却时间间隔


		'ARENA_ATTACK_SUCC' 	=> 1,	// 挑战成功
		'ARENA_ATTACK_FAILED' 	=> 2,	// 挑战失败
		'ARENA_ATTACKED_SUCC' 	=> 3,	// 被挑战,挑战方成功
		'ARENA_ATTACKED_FAILED' => 4,	// 被挑战,挑战方失败

		'CD_TIME_ONE_MIN_COST_INGOT' => 1, // 1元宝消除1min

		'MAX_RESERVED_RECORD' => 20, //最多返回x条纪录给客户端

		'GETBACK_AWARD_BOX_DAY1_COST_INGOT' => _S(10, 10, 20, 10), // 找回前一天比武场的宝箱要花费10元宝
		'GETBACK_AWARD_BOX_DAY2_COST_INGOT' => _S(20, 20, 40, 20), // 找回前两天比武场的宝箱要花费20元宝
		'GETBACK_AWARD_BOX_DAY3_COST_INGOT' => _S(40, 40, 80, 40), // 找回前三天比武场的宝箱要花费40元宝
	),
);
?>
