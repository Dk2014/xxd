<?php
include "common_links.php";

function desc_editor($row){
	$code  = '<textarea name="desc[]" rows="3" cols="20" >';
	$code .= $row['desc'];
	$code .= '</textarea>';
	return $code;
}

function desc_render($row){
	$code  = '<textarea disabled="disabled" name="desc[]" rows="3" cols="20">';
	$code .= $row['desc'];
	$code .= '</textarea>';
	return $code;
}

$pconfig = array(
	'title'   => '魂侍被动技能',
	'table'   => 'ghost_passive_skill',
	'links'   => $ghost_links,
	'columns' => array(
		array('name' => 'star', 'text' => '星级','width' => '100px'),
		array('name' => 'name', 'text' => '被动技能名称','width' => '100px'),
		array('name' => 'sign', 'text' => '图标标识','width' => '100px'),

		array('name' => 'desc', 'text' => '描述',
			'render' => 'desc_render',
			'editor' => 'desc_editor',
		),
	),
);
?>
