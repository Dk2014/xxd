<?php
die("This page is DEPRECATED!");

$pconfig = array(
	'title'   => '魂侍技能威力',
	'table'   => 'ghost_skill_force',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'level', 'text' => '等级', 'width' => '80px'),
		array('name' => 'force', 'text' => '威力', 'width' => '180px'),
	),

	'where' 		=> 'sql_where',
	'insert' 		=> 'sql_insert',

);

function sql_where($params) {
	if(!isset($params['m'])) {
		die('not enough argms');
	}
	return "where ghost_id={$params['m']} order by level";
}

function sql_insert($params) {
	if(!isset($params['m'])) {
		die('not enough argms');
	}
	return " ghost_id={$params['m']} ";
}
?>
