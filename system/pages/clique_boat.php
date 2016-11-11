<?php

$pconfig = array(
	'title'   => '帮派镖船',
	'table'   => 'clique_boat',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'name', 'text' => '名称', 'width' => '150px'),
		array('name' => 'color', 'text' => '颜色', 'width' => '150px'),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '150px'),
		array('name' => 'rate', 'text' => '概率', 'width' => '80px'),
		array('name' => 'escort_time', 'text' => '送镖时间（分钟）', 'width' => '80px'),
		array('name' => 'select_cost_ingot', 'text' => '元宝直选价格(0则不可元宝选择)', 'width' => '80px'),

		array('name' => 'award_fame', 'text' => '奖励声望', 'width' => '80px'),
		array('name' => 'award_coins', 'text' => '奖励铜钱', 'width' => '80px'),
		array('name' => 'award_clique_contrib', 'text' => '奖励帮派贡献', 'width' => '80px'),

		array('name' => 'hijack_lose_fame', 'text' => '劫持损失声望', 'width' => '80px'),
		array('name' => 'hijack_lose_coins', 'text' => '劫持损失铜钱', 'width' => '80px'),
		array('name' => 'hijack_lose_clique_contrib', 'text' => '劫持损失帮派贡献', 'width' => '80px'),


	),
	

);

?>
