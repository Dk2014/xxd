<?php
$pconfig = array(
	'title'   => '功能参数',
	'table'   => 'item_reflect_config',
	'links'   => $links,
	'columns' => array(		
		array('name' => 'award_coin_min', 'text' => '最少奖励铜钱', 'width' => '70px'),
		array('name' => 'award_coin_max', 'text' => '最多奖励铜钱', 'width' => '70px'),
	),
	'location' => 'location',
	'where' => 'sql_where',
	'insert' => 'sql_insert',
);
?>
