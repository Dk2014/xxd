<?php

$pconfig = array(
	'title'   => '七天登录奖励',
	'table'   => 'login_award',
	'links'   => array(),

	'columns' => array(
		array('name' => 'requrie_active_days', 'text' => '累计活跃天数', 'width' => '80px'),
		array('name' => 'award_type', 'text' => '奖励类型', 'width' => '80px'),
		array('name' => 'award_id', 'text' => '奖励物品ID', 'width' => '80px'),
		array('name' => 'award_num', 'text' => '奖励数量', 'width' => '80px'),

	),

	'where' 		=> 'sql_where',
	'js'			=> 'js_function',
);

?>
