<?php
include "common_links.php";

function tip_editor($row){
	$code  = '<textarea name="tip[]" rows="7" cols="40" >';
	$code .= $row['tip'];
	$code .= '</textarea>';
	return $code;
}

function tip_render($row){
	$code  = '<textarea disabled="disabled" name="tip[]" rows="7" cols="40">';
	$code .= $row['tip'];
	$code .= '</textarea>';
	return $code;
}

$pconfig = array(
	'title'   => '魂侍提示信息',
	'table'   => 'ghost_tip',
	'links'   => $ghost_links,
	
	'columns' => array(
		array('name' => 'tip', 'text' => '提示信息', 'width' => '180px',
			'render' => "tip_render",
			'editor' => "tip_editor",
		),
	),
);

?>