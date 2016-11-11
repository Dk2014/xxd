<?php
include "item_links.php";

function location($params){
	global $db;

	$box_name = db_query($db, "select `name` from `item` where `id` = {$params['m']}");

	$html = '<a href="?p=item_box">'.$box_name[0]['name'].'</a> -> ';

	return $html;
}

function sql_where($params) {
	if (isset($params['m']))
		$where = " where `item_id`={$params['m']}";

	return $where;
}

function sql_insert($params) {
	return "`item_id` = {$params['m']}";
}

