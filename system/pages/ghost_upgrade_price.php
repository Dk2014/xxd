<?php
include "common_links.php";

$pconfig = array(
	'title'   => '元宝升级魂侍价格配置',
	'table'   => 'ghost_upgrade_price',
	'links'   => $ghost_links,
	
	'columns' => array(
		array('name' => 'quality', 'text' => '品质', 'width' => '100px'),
		array('name' => 'cost', 'text' => '碎片单价', 'width' => '100px'),

	),
	'where' 		=> 'sql_where',
);

?>
