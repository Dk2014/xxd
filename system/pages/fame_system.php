<?php
include "fame_links.php";

$pconfig = array(
	'title'   => '声望系统',
	'table'   => 'fame_system',
	'links'  => $fame_links, 

	'columns' => array(
		array('name' => 'name', 'text' => '系统名', 'width' => '150px'),
		array('name' => 'sign', 'text' => '唯一标识', 'width' => '150px'),
		array('name' => 'max_fame', 'text' => '最大产出声望', 'width' => '150px'),
	),

	'opera' => array(),

);
?>


