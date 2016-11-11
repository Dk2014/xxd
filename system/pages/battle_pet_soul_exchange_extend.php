<?php
die("This page is DEPRECATED!");

$extend_columns = array(
/*   '字段' => 配置 */
	'quality' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array('蓝','紫','金','橙'),
	),
	'star' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(
				1=>'一星',
				2=>'二星',
				3=>'三星',
				4=>'四星',
				5=>'五星',
			),
	),
);

function render_quality($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_quality($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}


function render_star($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_star($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}
?>
