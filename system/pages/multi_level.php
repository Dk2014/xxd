<?php

$pconfig = array(
	'title'   => '多人关卡',
	'table'   => 'multi_level',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'name', 'text' => '关卡名称', 'width' => '150px'),
		array('name' => 'require_level', 'text' => '主角等级要求', 'width' => '80px'),
		array('name' => 'lock', 'text' => '关卡权值', 'width' => '80px'),
		array('name' => 'award_lock', 'text' => '奖励权值', 'width' => '80px'),
		array('name' => 'sign_war', 'text' => '战斗资源标识', 'width' => '150px'),
		array('name' => 'music', 'text' => '音乐资源标识', 'width' => '150px'),
		array('name' => 'award_exp', 'text' => '奖励经验', 'width' => '80px'),
		array('name' => 'award_fame', 'text' => '奖励声望', 'width' => '80px'),
		array('name' => 'award_coin', 'text' => '奖励铜钱', 'width' => '80px'),
		array('name' => 'award_relationship', 'text' => '奖励友情', 'width' => '80px'),
		array('name' => 'award_item1_id', 'text' => '奖励物品1 id', 'width' => '80px'),
		array('name' => 'award_item1_num', 'text' => '物品1数量', 'width' => '80px'),
		array('name' => 'award_item2_id', 'text' => '奖励物品2 id', 'width' => '80px'),
		array('name' => 'award_item2_num', 'text' => '物品2数量', 'width' => '80px'),
		array('name' => 'award_item3_id', 'text' => '奖励物品3 id', 'width' => '80px'),
		array('name' => 'award_item3_num', 'text' => '物品3数量', 'width' => '80px'),
	),
	
	'opera' => array(
		array('type' => 'link', 'text' => '', 'render' => 'opera_render'),
	),
	
	'js'			=>'js_function',
	'where' 		=> 'sql_where',

	'not_delete' 		=> true,
);

?>
