<?php

$pconfig = array(
	'title'   => '城镇数据',
	'table'   => 'town',
	'links'   => array(),
	'columns' => array(
		array('name' => 'name', 'text' => '城镇名称', 'width' => '120px',),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '150px',),
		array('name' => 'music', 'text' => '音乐资源标识', 'width' => '150px',),
		array('name' => 'lock', 'text' => '解锁权值', 'width' => '60px',),
		array('name' => 'start_x', 'text' => '出生点x轴坐标', 'width' => '60px',),
		array('name' => 'start_y', 'text' => '出生点y轴坐标', 'width' => '60px',),
		array('name' => 'exit_x', 'text' => '冒险区域出口x轴坐标', 'width' => '60px',),
		array('name' => 'exit_y', 'text' => '冒险区域出口y轴坐标', 'width' => '60px',),
	),
	'opera' => array(
		array('type' => 'link', 'text' => '商人出售物品', 'render' => 'item_blacksmith'),
		),
	'js' => 'js_function',
	'where'	  => 'sql_where',
	'not_delete' => true,
);
?>
