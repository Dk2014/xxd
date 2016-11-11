<?php

include "chest_links.php";

$pconfig = array(
	'title'   => '神龙宝箱',
	'table'   => 'chest',
	'links'   => $links,
	
	'columns' => array(
		array('name' => 'quality', 'text' => '宝箱品质', 'width' => '80px'),
		array('name' => 'probability', 'text' => '概率（%）', 'width' => '80px'),
	),
	
	'opera' => array(
		array('type' => 'link', 'text' => '设置物品', 'render' => 'itemtable'),
	),

	'where' 		=> 'sql_where',
	'insert'  		=> 'sql_insert',
);

function sql_where($params) {
	return " where `type` = 2";
}

function sql_insert($params) {
	return "`type` = 2";
}

function itemtable($row) {
	return html_get_link("设置物品", "?p=chest_item&m=".$row['id'], false);
}

?>