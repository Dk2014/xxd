<?php

include "item_links.php";

$pconfig = array(
	'title'   => '物品兑换',
	'table'   => 'item_exchange',
	'links'   => $links,
	
	'columns' => array(
		array('name' => 'target_item_id', 'text' => '目标物品', 'width' => '130px'),
		array('name' => 'target_item_num', 'text' => '目标物品数量', 'width' => '70px'),
		array('name' => 'item_id', 'text' => '物品', 'width' => '70px'),
		array('name' => 'item_num', 'text' => '物品数量', 'width' => '50px'),
	),

);
?>