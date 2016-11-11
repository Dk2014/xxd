<?php
$extend_columns = array(
	/*'字段' => 配置 */

	'quality' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
);

function range_quality(){
	global $all_qualitys;
	return $all_qualitys;
}

?>
