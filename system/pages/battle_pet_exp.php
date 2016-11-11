<?php
include "battle_pet_links.php";

function sql_where($params) {
	return " order by `level`";
}

$pconfig = array(
	'title'       => '灵宠等级',
	'table'       => 'battle_pet_exp',
	'links'       => $battle_pet_links,
	'columns'     => array(
		array('name' => 'level', 'text' => '等级',   'width' => '50px'),
		array('name' => 'exp', 'text' => '升级所需经验','width' => '100px'),
		array('name' => 'need_soul_num', 'text' => '每次所需灵魄数量','width' => '120px'),
		array('name' => 'min_add_exp', 'text' => '最小经验加值','width' => '100px'),
		array('name' => 'max_add_exp', 'text' => '最大经验加值','width' => '100px'),
	),
	'where' => 'sql_where',
);
