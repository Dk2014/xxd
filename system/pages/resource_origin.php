<?php
include "item_links.php";
$pconfig = array(
	'title'   => '物品产出',
	'table'   => 'resource_origin',
	'links'   => $links,

	'columns' => array(
		array('name' => 'item_id', 'text' => '物品', 'width' => '80px'),
		array('name' => 'origin_type', 'text' => '产出关卡类型', 'width' => '80px'),
		array('name' => 'origin_key', 'text' => '产出关卡ID', 'width' => '80px'),
		array('name' => 'description', 'text' => '描述', 'width' => '200px'),
	),

	'opera' => array(),

	'js'	        => 'js_function',

);


?>
