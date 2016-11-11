<?php
die("This page is DEPRECATED!");

include "battle_pet_links.php";

function sql_where($params) {
	return " order by `level`";
}

$pconfig = array(
	'title'       => '灵宠格等级数值',
	'table'       => 'battle_pet_grid_level',
	'links'       => $battle_pet_links,
	'columns'     => array(
		array('name' => 'level', 'text' => '格子等级',   'width' => '50px'),
		array('name' => 'require_level', 'text' => '主角等级',   'width' => '50px'),
		array('name' => 'cost_soul_num', 'text' => '消耗灵魂数量',   'width' => '50px'),
		array('name' => 'exp', 'text' => '升到下一级所需经验','width' => '300px'),
		array('name' => 'min_add_exp', 'text' => '最小经验','width' => '300px'),
		array('name' => 'max_add_exp', 'text' => '最大经验','width' => '300px'),
	),
	'where' => 'sql_where',
);
?>
