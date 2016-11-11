<?php

$extend_columns = array(
	/*'字段' => 配置 */

	'type_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'quality' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'fragment_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'only_exchange' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
);

function range_type_id(){
	global $allSwordSoulType;
	return $allSwordSoulType;
}

function range_quality(){
	global $qualityConfig;
	return $qualityConfig;
}

function range_fragment_id(){
	global $all_sword_soul_fragments;
	return $all_sword_soul_fragments;
}

function range_only_exchange(){
	global $onlyExchange;
	return $onlyExchange;
}

?>