<?php

$pconfig = array(
	'title'   => '极限关卡-彩虹桥',
	'table'   => 'rainbow_level',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'segment', 'text' => '段数', 'width' => '180px'),
		array('name' => 'award_id', 'text' => '通过魂侍奖励', 'width' => '180px'),
	),
	
	'opera' => array(
		array('type' => 'link', 'text' => '关卡', 'render' => 'leveltable'),
	),

	'sql_where' => 'sql_where',
	'js'	    => 'js_function',
);


function sql_where($params) {	
	return " order by `segment` asc ";
}


function leveltable($row) {
	return html_get_link("配置", "?p=rainbow_level_config&m=".$row['segment'], false)." | ".html_get_link("扫荡宝箱", "?p=rainbow_autofight_award&m=".$row['segment'], false);
}


?>
