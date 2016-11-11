<?php

include "item_links.php";

$pconfig = array(
	'title'   => '物品限购',
	'table'   => 'purchase_limit',
	'links'   => $links,
	
	'columns' => array(
		array('name' => 'item_id', 'text' => '物品名称', 'width' => '150px'),
		array('name' => 'num', 'text' => '限购数量', 'width' => '80px'),
	),
	'js'	        => 'js_function',

);
?>
