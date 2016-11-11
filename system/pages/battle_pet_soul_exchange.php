<?php
die("This page is DEPRECATED!");

include "battle_pet_links.php";

function sql_where($params) {
	return " order by `quality`, `star`";
}

$pconfig = array(
	'title'       => '灵宠灵魄兑换',
	'table'       => 'battle_pet_soul_exchange',
	'links'       => $battle_pet_links,
	'columns'     => array(
		array('name' => 'star', 'text' => '星', 'width' => '80px'),
		array('name' => 'quality', 'text' => '品质', 'width' => '80px'),
		array('name' => 'soul_num', 'text' => '灵魄数量', 'width' => '80px'),
	),
	'where' => 'sql_where',
);
?>
