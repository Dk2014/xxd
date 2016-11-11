<?php
die("This page is DEPRECATED!");

include "battle_pet_links.php";

function sql_where($params) {
	return " order by `grid_id`, `level`";
}

$pconfig = array(
	'title'       => '灵宠等级属性',
	'table'       => 'battle_pet_grid_attribute',
	'links'       => $battle_pet_links,
	'columns'     => array(
		array('name' => 'grid_id', 'text' => '格子',   'width' => '50px'),
		array('name' => 'level', 'text' => '等级',   'width' => '50px'),
		array('name' => 'health', 'text' => '生命', 'width' => '40px'),
		array('name' => 'attack', 'text' => '攻击', 'width' => '40px'),
		array('name' => 'defence', 'text' => '防御', 'width' => '40px'),
		array('name' => 'speed', 'text' => '速度', 'width' => '40px'),
	),
	'where' => 'sql_where',
);
?>
