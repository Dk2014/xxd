<?php

$pconfig = array(
	'title'   => 'VIP'.$_GET['m'].'级升级奖励',
	'table'   => 'vip_levelup_gift',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'item_id', 'text' => '物品ID', 'width' => '80px'),
		array('name' => 'item_num', 'text' => '物品数量', 'width' => '80px'),

	),
	
	'opera' => array(
	),
	

	'where' 		=> 'sql_where',
	'insert'  		=> 'sql_insert',
	'js'  			=> 'js_function',
);

?>




