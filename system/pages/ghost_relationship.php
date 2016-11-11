<?php
include "common_links.php";
$pconfig = array(
	'title'   => '魂侍连锁',
	'table'   => 'ghost_relationship',
	'links'       => $ghost_links,
	'columns' => array(
		array('name' => 'ghost_a', 'text' => '魂侍A', 'width' => '80px'),
		array('name' => 'ghost_a_star', 'text' => '魂侍A所需星级', 'width' => '80px'),
		array('name' => 'ghost_b', 'text' => '魂侍B', 'width' => '80px'),
		array('name' => 'ghost_b_star', 'text' => '魂侍B所需星级', 'width' => '80px'),

	),
);

?>