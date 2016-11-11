<?php
include "common_links.php";

$pconfig = array(
	'title'   => '技能训练价格',
	'table'   => 'ghost_skill_train_price',
	'links'   => $ghost_links,
	
	'columns' => array(
		array('name' => 'level', 'text' => '技能等级', 'width' => '100px'),
		array('name' => 'cost', 'text' => '训练价格', 'width' => '100px'),

	),
	'where' 		=> 'sql_where',
);

function sql_where($params) {
	return ' order by `level` ';
}

?>
