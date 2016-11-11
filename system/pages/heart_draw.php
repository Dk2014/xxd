<?php

$pconfig = array(
	'title'   => '爱心抽奖',
	'table'   => 'heart_draw',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'draw_type', 'text' => '抽奖类型', 'width' => '100px'),
		array('name' => 'daily_num', 'text' => '每日可抽奖次数', 'width' => '150px'),
		array('name' => 'cost_heart', 'text' => '每次抽奖消耗爱心数', 'width' => '150px'),
	),
	
	'opera' => array(
		array('type' => 'link', 'text' => '', 'render' => 'op_render'),
	),
	
	'not_delete' 	=> true,
);

?>