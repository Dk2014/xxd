<?php

function sql_where($params) {
	return " order by `times`";
}

$pconfig = array(
	'title'   => '元宝购买资源关卡次数所消耗元宝配置',
	'table'   => 'buy_resource_times_config',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'times', 'text' => '购买次数', 'width' => '100px'),
		array('name' => 'cost', 'text' => '购买所需元宝', 'width' => '100px'),

	),
	'where' 		=> 'sql_where',
);

?>
