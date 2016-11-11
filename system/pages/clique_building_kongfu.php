<?php

include "clique_buildings_link.php";

$pconfig = array(
	'title'   => '帮派武功建筑',
	'table'   => 'clique_building_kongfu',
	'links'   => $clique_buildings_link,
	
	'columns' => array(
		array('name' => 'building', 'text' => '建筑', 'width' => '80px'),
		array('name' => 'level', 'text' => '等级', 'width' => '80px'),
		array('name' => 'upgrade_coins', 'text' => '升级铜钱', 'width' => '80px'),
		array('name' => 'desc', 'text' => '描述', 'width' => '200px'),
		array('name' => 'cur_kongfu_desc', 'text' => '当前武功描述', 'width' => '200px'),
		array('name' => 'next_kongfu_desc', 'text' => '下一等级武功描述', 'width' => '200px'),

	),
	
	'opera' => array(
	),
	

	'where' 		=> 'sql_where',
);

function sql_where($params){
	if(isset($params['m'])) {
		return "where building={$params['m']} order by level ASC";
	}
	return ' ORDER BY `building`, `level` ASC';
}


?>
