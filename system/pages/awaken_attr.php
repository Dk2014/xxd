<?php

$pconfig = array(
	'title' => '觉醒列表',
	'table' => 'awaken_attr',
	'links' => array(),

	'columns' => array(
		array('name' => 'name', 'text' => '觉醒名称', 'width' => '120px'),
		array('name' => 'is_skill', 'text' => '是否为技能', 'width' => '120px'),
		array('name' => 'skill_id', 'text' => '觉醒技能', 'width' => '120px'),
		array('name' => 'type', 'text' => '觉醒属性', 'width' => '120px'),
		array('name' => 'attr', 'text' => '觉醒属性值', 'width' => '120px'),
		array('name' => 'lights', 'text' => '所需希望之光', 'width' => '120px'),
	),
	'where' => 'sql_where',
);

function sql_where($params) {
	return "order by `id`";
}
