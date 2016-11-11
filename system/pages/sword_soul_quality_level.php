<?php
include_once "common_links.php";

function location($params){
	global $db;
	$html = '';

	$sql = "select `id`, `name` from sword_soul_quality where id = {$params['m']}";
	$sword_soul_quality = db_query($db, $sql);
	if (count($sword_soul_quality) == 0) {
		return $html;
	}
	
	$html = '<a href="?p=sword_soul_quality&m='.$sword_soul_quality[0]['id'].'">'.$sword_soul_quality[0]['name'].'</a> -> ' . $html;

	return $html;
}

function sql_where($params) {
	if (!isset($params['m'])){
		return '';
	}

	return "where `quality_id`={$params['m']}";
}

function sql_insert($params) {
	return "`quality_id` = {$params['m']}";
}

$pconfig = array(
	'title'   => '剑心品质等级',
	'table'   => 'sword_soul_quality_level',
	'links'   => $sword_soul_links,
	'columns' => array(
		array('name' => 'level', 'text' => '等级'),
		array('name' => 'exp', 'text' => '升级到下一级所需经验'),
	),
	'where'	  => 'sql_where',
	'insert'  => 'sql_insert',
	'location' => 'location',
);
?>