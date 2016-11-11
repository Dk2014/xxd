<?php
die("This page is DEPRECATED!");

include "battle_pet_links.php";

$pconfig = array(
	'title'   => '元宝升级灵宠格子价格配置',
	'table'   => 'battle_pet_grid_upgrade_price',
	'links'   => $battle_pet_links,
	
	'columns' => array(
		array('name' => 'times', 'text' => '今日次数', 'width' => '100px'),
		array('name' => 'cost', 'text' => '魂魄单价', 'width' => '100px'),

	),
	'where' 		=> 'sql_where',
);

function sql_where($params) {
	return ' order by `times` asc';
}

?>
