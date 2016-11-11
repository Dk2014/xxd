<?php 
$extend_columns = array(
	/*'字段' => 配置 */

	'type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
);


function range_type(){
	global $types;
	return $types;
}


?>