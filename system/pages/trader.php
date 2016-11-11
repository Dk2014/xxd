<?php

$pconfig = array(
	'title'   => '随机商店商人',
	'table'   => 'trader',
	'links'   => array(),

	'columns' => array(
		array('name' => 'name', 'text' => 'NPC名称', 'width' => '80px'),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '80px'),
		array('name' => 'talk', 'text' => '日常对话', 'width' => '80px'),
		array('name' => 'sold_out_talk', 'text' => '售罄对话', 'width' => '80px'),
		array('name' => 'deal_talk', 'text' => '成交对话', 'width' => '80px'),

	),

	'opera' => array(
		array('type' => 'link', 'text' => '城镇坐标', 'render' => 'trader_town_position'),
		array('type' => 'link', 'text' => '货物配置', 'render' => 'trader_config'),
		array('type' => 'link', 'text' => '额外对话', 'render' => 'trader_extra_talk'),
		array('type' => 'link', 'text' => '刷新价格', 'render' => 'trader_refresh_price'),
	),
	'not_delete' 	=> true,
);

?>
