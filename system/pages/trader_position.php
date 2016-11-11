<?php
$pconfig = array(
	'title'   => '随机商人坐标',
	'table'   => 'trader_position',
	'links'   => array(),

	'columns' => array(
		array('name' => 'town_id', 'text' => '城镇ID', 'width' => '80px'),
		array('name' => 'x', 'text' => 'x轴坐标', 'width' => '80px'),
		array('name' => 'y', 'text' => 'y轴坐标', 'width' => '80px'),
		array('name' => 'direction', 'text' => '朝向', 'width' => '80px'),

	),

	'where' 		=> 'sql_where',
	'insert' 		=> 'sql_insert',

	'location'		=> 'location',

);
?>
