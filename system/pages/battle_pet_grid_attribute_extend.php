<?php
die("This page is DEPRECATED!");

$extend_columns = array(
	/*'字段' => 配置 */

	
	'grid_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(
			1=>'麒麟',
			2=>'朱雀',
			3=>'白虎',
			4=>'玄武',
			5=>'青龙',
			),
	),
);

function render_grid_id($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_grid_id($column_name, $column_val, $row, $data){
	return html_get_select($column_name,$data,$column_val);
}

?>
