<?php
include "fame_links.php";


$pconfig = array(
	'title'   => '声望等级',
	'table'   => 'fame_level',
	'links'  => $fame_links, 

	'columns' => array(
		array('name' => 'level', 'text' => '声望等级', 'width' => '150px'),
		array('name' => 'required_fame', 'text' => '累计声望要求', 'width' => '150px'),
	),

	'opera' => array(),

);
?>


