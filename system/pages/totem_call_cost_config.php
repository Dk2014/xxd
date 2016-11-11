<?php

include "common_links.php";


$pconfig = array(
	'title' => '阵印元宝召唤价格',
	'table' => 'totem_call_cost_config',
	'links' => $totem_links,

	'columns' => array(
		array('name' => 'times', 'text' => '次数', 'width' => '80px'),
		array('name' => 'cost', 'text' => '价格', 'width' => '80px'),
	),
);
