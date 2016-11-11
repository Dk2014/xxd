<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'need_play' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array("不需要","需要"),
	),
	'type' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array( 1=> "权值",2=>"等级"),
	),
	// 'level' => array(
	// 	'editor' 	=> array('params' => array()),
	// 	'render' 	=> array('params' => array()),
	// 	'data' => array(),
	// ),
);


function sql_where($params) {
	return " ORDER BY `lock` DESC";
}
function render_type($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_type($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_need_play($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_need_play($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

// function render_level($column_name, $column_val, $row, $data){
// 	return isset($column_val)?$column_val:0;
// }

// function editor_level($column_name, $column_val, $row, $data){
// 	return html_get_input($column_name, isset($column_val)?$column_val:0);
// }

?>
