<?php
include_once "common_links.php";

function location($params){
	global $db;
	$html = '';

	$sql = "select `id`, `name` from sword_soul where id = {$params['m']}";
	$sword_soul = db_query($db, $sql);
	if (count($sword_soul) == 0) {
		return $html;
	}
	
	$html = '<a href="?p=sword_soul&m='.$sword_soul[0]['id'].'">'.$sword_soul[0]['name'].'</a> -> ' . $html;

	return $html;
}

function sql_where($params) {
	if (!isset($params['m'])){
		return '';
	}

	return "where `sword_soul_id`={$params['m']}";
}

function sql_insert($params) {
	return "`sword_soul_id` = {$params['m']}";
}

$pconfig = array(
	'title'   => '剑心等级',
	'table'   => 'sword_soul_level',
	'links'   => $sword_soul_links,
	'columns' => array(
		array('name' => 'level', 'text' => '等级'),
		array('name' => 'value', 'text' => '属性加值'),
	),
	'where'	  => 'sql_where',
	'insert'  => 'sql_insert',
	'location' => 'location',
);
?>