<?php

include "skill_links.php";

$pconfig = array(
	'title'   => '绝招训练消耗',
	'table'   => 'skill_training_cost',
	'links'   => $links,

	'columns' => array(
		array('name' => 'level', 'text' => '绝招等级', 'width' => '80px'),
		array('name' => 'cost', 'text' => '消耗铜钱', 'width' => '80px'),
	),

	'where' => 'sql_where',
);

function sql_where($params) {
    return " ORDER BY `level` ASC";
}

