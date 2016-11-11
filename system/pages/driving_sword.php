<?php

$pconfig = array(
	'title'   => '云海御剑',
	'table'   => 'driving_sword',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'level', 'text' => '层级', 'width' => '90px'),
		array('name' => 'width', 'text' => '地图宽', 'width' => '90px'),
		array('name' => 'height', 'text' => '地图高', 'width' => '90px'),
		array('name' => 'born_x', 'text' => '出生地坐标x', 'width' => '90px'),
		array('name' => 'born_y', 'text' => '出生地坐标y', 'width' => '90px'),
		array('name' => 'hole_x', 'text' => '传送阵坐标x', 'width' => '90px'),
		array('name' => 'hole_y', 'text' => '传送阵坐标y', 'width' => '90px'),
		array('name' => 'obstacle_count', 'text' => '障碍总数', 'width' => '90px'),
	),
	
	'opera' => array(
		array('type' => 'link', 'text' => '传送点', 'render' => 'teleport_op'),
		array('type' => 'link', 'text' => '探险类', 'render' => 'exploring_op'),
		array('type' => 'link', 'text' => '拜访类', 'render' => 'visiting_op'),
		array('type' => 'link', 'text' => '宝藏类', 'render' => 'treasure_op'),
	),

	'sql_where' => 'sql_where',
);



function teleport_op($row) {
	return html_get_link("传送点", "?p=driving_sword_teleport&m=".$row['level'], false).'|';
}

function exploring_op($row) {
	return html_get_link("探险类", "?p=driving_sword_exploring&m=".$row['level'], false).'|';
}

function visiting_op($row) {
	return html_get_link("拜访类", "?p=driving_sword_visiting&m=".$row['level'], false).'|';
}

function treasure_op($row) {
	return html_get_link("宝藏类", "?p=driving_sword_treasure&m=".$row['level'], false);
}

