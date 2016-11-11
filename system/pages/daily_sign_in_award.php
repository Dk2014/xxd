<?php

$links = array(
	array('text' => '新玩家签到', 'id' => 'daily_sign_in_award&m=0'),
	array('text' => '常规签到', 'id' => 'daily_sign_in_award&m=1'),
);

$pconfig = array(
	'title'   => '每日签到奖励',
	'table'   => 'daily_sign_in_award',
	'links'   => $links,

	'columns' => array(
		array('name' => 'award_type', 'text' => '奖励类型', 'width' => '80px'),
		array('name' => 'award_id', 'text' => '奖励物品ID', 'width' => '80px'),
		array('name' => 'num', 'text' => '奖励数量', 'width' => '80px'),
		array('name' => 'vip_double', 'text' => 'VIP双倍奖励', 'width' => '80px'),

	),

	'where' 		=> 'sql_where',
	'insert'  		=> 'sql_insert',
	'js'			=> 'js_function',
	'location' => 'location',
);

?>
