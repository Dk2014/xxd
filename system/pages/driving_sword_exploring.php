<?php

$pconfig = array(
	'title'   => '云海御剑探险类',
	'table'   => 'driving_sword_exploring',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'event_id', 'text' => '探险山id', 'width' => '90px'),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '90px'),
		array('name' => 'award_item1', 'text' => '奖励品1', 'width' => '90px'),
		array('name' => 'award_num1', 'text' => '奖励品1数量', 'width' => '90px'),
		array('name' => 'award_item2', 'text' => '奖励品2', 'width' => '90px'),
		array('name' => 'award_num2', 'text' => '奖励品2数量', 'width' => '90px'),
		array('name' => 'award_item3', 'text' => '奖励品3', 'width' => '90px'),
		array('name' => 'award_num3', 'text' => '奖励品3数量', 'width' => '90px'),
		array('name' => 'award_coin_num', 'text' => '奖励铜钱数量', 'width' => '90px'),
		array('name' => 'garrison_item', 'text' => '每轮驻守奖励品', 'width' => '90px'),
		array('name' => 'garrison_num', 'text' => '每轮驻守奖励品数量', 'width' => '90px'),
		array('name' => 'garrison_coin_num', 'text' => '每轮驻守奖励铜钱数量', 'width' => '90px'),
	),

	'opera' => array(
		array('type' => 'link', 'text' => '关联关卡', 'render' => 'sweet_monsters'),
	),
	
	'where' => 'sql_where',
	'insert' => 'sql_insert',
	'js' => 'js_function',
);

function sweet_monsters($row) {
	return html_get_link("关联关卡", "?p=driving_sword_exploring_config&m={$row['id']}", false);
;
}

