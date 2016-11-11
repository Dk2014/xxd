<?php

include "item_links.php";

$pconfig = array(
	'title'   => '时装兑换',
	'table'   => 'fashion_exchange',
	'links' => array(),
	
	'columns' => array(
		array('name' => 'item_id', 'text' => '物品', 'width' => '70px'),
		//array('name' => 'valid_hours', 'text' => '有效时间（小时，0永久有效）', 'width' => '50px'),
	),
	'where' 		=> 'sql_where',
	'insert' 		=> 'sql_insert',
	'js'			=> 'js_function',
	'not_delete' => true,
);
?>
