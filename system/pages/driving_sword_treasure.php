<?php

$pconfig = array(
	'title'   => '云海御剑宝藏类',
	'table'   => 'driving_sword_treasure',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'event_id', 'text' => '宝藏山id', 'width' => '90px'),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '90px'),
	),

	'opera' => array(
		array('type' => 'link', 'text' => '宝藏列表', 'render' => 'content_op'),
	),

	'where' => 'sql_where',
	'insert' => 'sql_insert',
);

function content_op($row) {
	return html_get_link("宝藏列表", "?p=driving_sword_treasure_content&m=".$row['id'], false);
}

