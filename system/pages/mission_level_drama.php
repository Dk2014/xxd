<?php

$pconfig = array(
	'title'   => '关卡剧情',
	'table'   => 'mission_level_drama',
	'links'   => array(),
	'columns' => array(
		array('name' => 'name', 'text' => '剧情名称', 'width' => '120px',),
		array('name' => 'quest_id', 'text' => '关联任务', 'width' => '60px',),
		array('name' => 'quest_state', 'text' => '任务状态', 'width' => '60px',),
		array('name' => 'area_x', 'text' => '触发x轴坐标', 'width' => '60px',),
		array('name' => 'area_y', 'text' => '触发y轴坐标', 'width' => '60px',),
		array('name' => 'area_width', 'text' => '触发区域宽度', 'width' => '60px',),
		array('name' => 'area_height', 'text' => '触发区域高度', 'width' => '60px',),
	),
	'where'	  => 'sql_where',
	'insert' => 'sql_insert',
	'js' => 'jsFunction',
);
?>
