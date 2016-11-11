<?php
include "arena_links.php";

$pconfig = array(
	'title'   => '',
	'table'   => 'arena_rank_gap',
	'links'   =>$arena_links,
	
	'columns' => array(
		array('name' => 'rank', 'text' => '排名', 'width' => '180px'),
		array('name' => 'gap', 'text' => '间隔', 'width' => '180px'),
	),
	

);

?>
