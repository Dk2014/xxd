<?php

$pconfig = array(
	'title'   => '商人出售物品',
	'table'   => 'town_npc_item',
	'links'   => array(),
	'columns' => array(
		array('name' => 'town_npc_id', 'text' => '城镇NPC','width' => '200px',),
		array('name' => 'item_id', 'text' => '出售物品','width' => '200px',),
		array('name' => 'stock', 'text' => '库存', 'width' => '80px',),
		array('name' => 'vip', 'text' => 'vip特供','width' => '100px',),
	),
	'where' => 'sql_where',
	'location' => 'location',
);

?>