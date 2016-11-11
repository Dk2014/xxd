<?php

$pconfig = array(
	'title' => '觉醒图谱',
	'table' => 'awaken_graphic',
	'links' => array(),

	'columns' => array(
		array('name' => 'impl_id', 'text' => '实例编号', 'width' => '30px'),
		array('name' => 'attr_id', 'text' => '属性', 'width' => '30px'),
		array('name' => 'dep_impl', 'text' => '依赖实例', 'width' => '30px'),
		array('name' => 'direct', 'text' => '派生方向', 'width' => '30px'),
	),
	'where' => 'sql_where',
	'insert' => 'sql_insert',
);

function sql_where($params) {
	return "where `role_id` = ${params['m']} order by `impl_id`";
}

function sql_insert($params) {
	return "`role_id` = {$params['m']}";
}
