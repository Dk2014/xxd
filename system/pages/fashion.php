<?php


$pconfig = array(
	'title'   => '时装',
	'table'   => 'fashion',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'name', 'text' => '时装名称', 'width' => '100px'),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '100px'),
		array('name' => 'item_sign', 'text' => '关联物品资源标识', 'width' => '100px'),
		array('name' => 'level', 'text' => '需求等级', 'width' => '60px'),
		array('name' => 'health', 'text' => '生命', 'width' => '50px'),
		array('name' => 'speed', 'text' => '速度', 'width' => '50px'),
		array('name' => 'attack', 'text' => '攻击', 'width' => '50px'),
		array('name' => 'defence', 'text' => '防御', 'width' => '50px'),
		array('name' => 'cultivation', 'text' => '内力', 'width' => '50px'),
		array('name' => 'source', 'text' => '获得来源', 'width' => '150px'),
		array('name' => 'desc', 'text' => '时装描述', 'width' => '80px'),
	),
	'opera' => array(
		array('type' => 'link', 'render' => 'fashion_exchange_opera'),
	),
	'not_delete' => true,
);
?>
