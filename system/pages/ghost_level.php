<?php
include "common_links.php";

function sql_where($params) {
	return " order by `level`";
}

$pconfig = array(
	'title'       => '魂侍等级',
	'table'       => 'ghost_level',
	'links'       => $ghost_links,
	'columns'     => array(
		array('name' => 'level', 'text' => '等级',   'width' => '50px'),
		array('name' => 'exp', 'text' => '升级所需经验','width' => '100px'),
		array('name' => 'need_fruit_num', 'text' => '每次所需影界果实数量','width' => '120px'),
		array('name' => 'min_add_exp', 'text' => '最小经验加值','width' => '100px'),
		array('name' => 'max_add_exp', 'text' => '最大经验加值'),
	),
	'where' => 'sql_where',
);
?>