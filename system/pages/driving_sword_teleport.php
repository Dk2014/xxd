<?php

$pconfig = array(
	'title'   => '云海御剑传送点',
	'table'   => 'driving_sword_teleport',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'event_id', 'text' => '传送点事件id', 'width' => '90px'),
		array('name' => 'dest_event_id', 'text' => '目标传送点id', 'width' => '90px'),
	),
	
	'where' => 'sql_where',
	'insert' => 'sql_insert',
);

