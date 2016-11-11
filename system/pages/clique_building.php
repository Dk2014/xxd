<?php

include "clique_buildings_link.php";

$pconfig = array(
	'title'   => '帮派建筑物',
	'table'   => 'clique_building',
	'links'   => $clique_buildings_link,
	
	'columns' => array(
		array('name' => 'name', 'text' => '建筑名称', 'width' => '400px'),
		array('name' => 'biaoshi', 'text' => '建筑标识', 'width' => '400px'),
		array('name' => 'order', 'text' => '优先权排序', 'width' => '100px'),

	),
	
	'opera' => array(
		//array('type' => 'link', 'text' => '描述设置', 'render' => 'descr'),
	),
	
	'foot' => 'foot',
	'where' 		=> 'sql_where',
);

function sql_where(){
	return ' ORDER BY `order` ASC';
}

?>