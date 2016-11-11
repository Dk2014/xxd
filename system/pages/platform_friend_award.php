<?php
$pconfig = array(
	'title'   => '平台好友奖励',
	'table'   => 'platform_friend_award',
	'links'   => array(),

		
	'columns' => array(
		array('name' => 'name', 'text' => '物品展示名称', 'width' => '80px'),
		array('name' => 'require_friend_num', 'text' => '平台好友数', 'width' => '80px'),
		array('name' => 'award_type', 'text' => '奖励类型', 'width' => '80px'),
		array('name' => 'award_id', 'text' => '物品ID', 'width' => '80px'),
		array('name' => 'num', 'text' => '物品数量', 'width' => '80px'),
	),

	'js'	        => 'js_function',

);

?>
