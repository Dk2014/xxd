<?php
$pconfig = array(
	'title'   => 'NPC对话',
	'table'   => 'npc_talk',
	'links'   => array(),
	'columns' => array(
		array('name' => 'npc_id', 'text' => 'NPC ID', 'width' => '80px'),
		array('name' => 'type', 'text' => '类型', 'width' => '80px'),
		array('name' => 'quest_id', 'text' => '关联任务', 'width' => '80px'),
		array('name' => 'conversion', 'text' => '对话内容', 'width' => '80px'),
		array('name' => 'award_item_id', 'text' => '奖励物品ID', 'width' => '80px'),
		array('name' => 'award_item_num', 'text' => '奖励物品数量', 'width' => '80px'),
	),
	'insert' => 'sql_insert',
	'where' => 'sql_where',
	'location' => 'location',
	'js'	        => 'js_function',
);
?>
