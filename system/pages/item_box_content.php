<?php
$pconfig = array(
	'title'   => '宝箱数据',
	'table'   => 'item_box_content',
	'links'   => $links,
	'columns' => array(		
		array('name' => 'mode', 'text' => '随机方式','width' => '70px'),
		array('name' => 'type', 'text' => '奖励种类','width' => '70px'),
		array('name' => 'min_num', 'text' => '最少数量', 'width' => '70px'),
		array('name' => 'max_num', 'text' => '最多数量', 'width' => '70px'),
		array('name' => 'probability', 'text' => '概率(%)', 'width' => '70px'),
		array('name' => 'get_item_id', 'text' => '获得的物品名称','width' => '150px'),
		array('name' => 'item_desc', 'text' => '随机物品集的描述', 'width' => '150px'),
		array('name' => 'item_id_set', 'text' => '随机的物品集'),		
	),
	'location' => 'location',
	'where' => 'sql_where',
	'insert' => 'sql_insert',
	'js' => 'js_function',
);
?>