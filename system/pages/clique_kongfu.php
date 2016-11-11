<?php

include "clique_buildings_link.php";

$pconfig = array(
	'title'   => '帮派武功',
	'table'   => 'clique_kongfu',
	'links'   => $clique_buildings_link,
	
	'columns' => array(
		array('name' => 'name', 'text' => '武功名称', 'width' => '200px'),
		//array('name' => 'sign', 'text' => '资源表识', 'width' => '200px'),
		array('name' => 'building', 'text' => '所属建筑', 'width' => '80px'),
		array('name' => 'require_building_level', 'text' => '要求建筑等级等级', 'width' => '80px'),
		array('name' => 'attrib_type', 'text' => '属性类型', 'width' => '80px'),
		array('name' => 'attrib_value', 'text' => '加成值', 'width' => '80px'),
	),
	
	'opera' => array(
		array('type' => 'link', 'text' => 'xx', 'render' => 'kongfu_opera'),
	),
	

	'where' 		=> 'sql_where',
);

function sql_where(){
	return ' ORDER BY `building`, `require_building_level` ASC';
}

function kongfu_opera($row) {
	return html_get_link("心法等级", "?p=clique_kongfu_level&m=".$row['id'], false);
}

?>
