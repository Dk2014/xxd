<?php

$extend_columns = array(
/*   '字段' => 配置 */

	'equ_type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(3=>'武器',4=>'战袍',5=>'靴子',6=>'饰品'),
	),
	'equ_kind' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),	
);

function render_equ_type($column_name, $column_val, $row, $data) {
	return $data[$column_val];
}

function editor_equ_type($column_name, $column_val, $row, $data) {
	return html_get_select($column_name,$data,$column_val);
}

function render_equ_kind($column_name, $column_val, $row, $data) {
	global $equipmentLevelType;
	return $equipmentLevelType[$column_val];
}

function editor_equ_kind($column_name, $column_val, $row, $data) {
	global $equipmentLevelType;
	return html_get_select($column_name,$equipmentLevelType,$column_val);
}

function range_equ_kind(){
	global $equipmentLevelType;
	return $equipmentLevelType;
}

?>