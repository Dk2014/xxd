<?php

include "item_links.php";

$pconfig = array(
	'title'   => '越南龙珠活动',
	'table'   => 'event_vn_dragon_ball_exchange',
	'links'   => $links,
	
	'columns' => array(
		array('name' => 'name', 'text' => '道具名称', 'width' => '100px'),
		array('name' => 'quality', 'text' => '品质', 'width' => '60px'),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '150px'),
		array('name' => 'desc', 'text' => '道具描述', 'width' => '500px'),
	),
	
	'not_delete' => true,
);
?>
