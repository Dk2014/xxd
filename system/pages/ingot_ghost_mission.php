<?php

$pconfig = array(
	'title'   => '极暗净土',
	'table'   => 'ingot_ghost_mission',
	'links'   => array(),
	'columns' => array(
		array('name' => 'vip_level', 'text' => '可进入vip等级', 'width' => '120px',),
		array('name' => 'max_egg_num', 'text' => '龙珠数量', 'width' => '150px',),
		array('name' => 'yellow_ghost_rand', 'text' => '金色魂侍概率(万分之)', 'width' => '60px',),
		array('name' => 'purple_ghost_rand', 'text' => '紫色魂侍概率(万分之)', 'width' => '60px',),
		array('name' => 'orange_ghost_rand', 'text' => '橙色魂侍概率(万分之)', 'width' => '60px',),
	),
	'opera' => array(
		array('type' => 'link', 'render' => 'ingot_ghost_mission_opera','width' => '300px'),
	),
);

function ingot_ghost_mission_opera($row) {
	$html = '';
	$html .= '<table border="0" width="400"><tr>';
	// 影界魂侍
	$html .= '<td width="40">';
	$html .= '<a href="?p=ghost&m=-1">影界魂侍</a>';
	$html .= ghost_info("select * from `ghost` where `town_id`= -1 and `id` not in (28,39,40)");
	$html .= '</td>';

	$html .= '</tr></table>';
	return $html;
}

function ghost_info($sql){
	global $db;
	$html = '';
	$ghostInfos = db_query($db, $sql);
	foreach ($ghostInfos as $ghostInfo) {
		$html .= "</br>".$ghostInfo['name'];
	}
	return $html;
}

?>