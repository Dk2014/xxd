<?php

$extend_columns = array(
	/*'字段' => 配置 */

	'type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'priority' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array("0级","1级"),
		//'range' => array('params' => array()),
	),
	//'expire_time' => array(
	//	'editor' => array('params' => array()),
	//	'render' => array('params' => array()),
	//	'data' => array(0,1),
	//	//'range' => array('params' => array()),
	//),
);

function range_type(){
	global $types;
	return $types;
}

//function range_expire_time(){
//	global $expire_time;
//	return $expire_time;
//}

function render_priority($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_priority($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

?>
