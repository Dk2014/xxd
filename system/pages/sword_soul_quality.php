<?php
include_once "common_links.php";

function color_render($row){
	return '<span style="color:#'.preg_replace('/0x/i', "", $row['color']).'">'.$row['color'].'</a>';
}

function quality_name_render($row) {
	return '<a href="?p=sword_soul_quality_level&m='.$row['id'].'">'.$row['name'].'</a>';
}


$pconfig = array(
	'title'   => '剑心品质',
	'table'   => 'sword_soul_quality',
	'links'   => $sword_soul_links,
	'columns' => array(
		array('name' => 'name', 'text' => '名称', 
			'render' => 'quality_name_render', 'width' => '120px',
		),
		array('name' => 'sign', 'text' => '程序标示'),
		array('name' => 'color', 'text' => '颜色',
			'render' => 'color_render',
		),
	),
	'not_delete' => true,
	'not_new' => true,
);
?>