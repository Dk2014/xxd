<?php

include "chest_links.php";

$pconfig = array(
	'title'   => '宝箱物品',
	'table'   => 'chest_item',
	'links'   => array(),
	
	'columns' => array(
		
		array('name' => 'type', 'text' => '类型', 'width' => '80px'),
		array('name' => 'item_id', 'text' => '物品', 'width' => '80px'),
		array('name' => 'item_num', 'text' => '数量', 'width' => '80px'),

	),
	
	'where' 		=> 'sql_where',
	'insert'  		=> 'sql_insert',
	'js'	        => 'js_function',
);


?>
