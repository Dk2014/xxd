<?php
include "arena_links.php";

$pconfig = array(
	'title'   => '比武场奖励',
	'table'   => 'arena_award_box',
	'links'   =>$arena_links,
	
	'columns' => array(
		array('name' => 'max_rank', 'text' => '最大排名', 'width' => '100px'),
		array('name' => 'ingot', 'text' => '元宝', 'width' => '180px'),
		array('name' => 'fame', 'text' => '声望', 'width' => '180px'),
		array('name' => 'coins', 'text' => '奖励铜钱', 'width' => '180px'),
		array('name' => 'item_id', 'text' => '奖励物品', 'width' => '180px'),
		array('name' => 'item_num', 'text' => '奖励物品数量', 'width' => '180px'),
		array('name' => 'item2_id', 'text' => '奖励物品2', 'width' => '180px'),
		array('name' => 'item2_num', 'text' => '奖励物品2数量', 'width' => '180px'),

	),
	

	'js' => 'js_function',
);

?>
