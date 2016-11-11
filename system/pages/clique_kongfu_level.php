<?php

include "clique_buildings_link.php";

$pconfig = array(
	'title'   => '帮派武功等级',
	'table'   => 'clique_kongfu_level',
	//'links'   => $clique_buildings_link,
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'consume_contrib', 'text' => '消耗贡献', 'width' => '200px'),
		array('name' => 'level', 'text' => '等级', 'width' => '200px'),

	),
	
	'opera' => array(
	),
	

	'where' 		=> 'sql_where',
	'insert' 		=> 'sql_insert',
);

function sql_where($params){
	if(isset($params['m'])) {
		return "where  kongfu_id={$params['m']} ";
	}
	die("need kongfu_id");
}

function sql_insert($params){
	if(isset($params['m'])) {
		return " kongfu_id={$params['m']} ";
	}
	die("need kongfu_id");
}


?>
