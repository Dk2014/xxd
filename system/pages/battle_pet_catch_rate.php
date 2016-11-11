<?php
include "battle_pet_links.php";

$pconfig = array(
	'title'   => '区域关卡',
	'table'   => 'battle_pet_catch_rate',
	'links'   => $battle_pet_links,
	'columns' => array(
		array('name' => 'rate', 'text' => '增加概率', 'width' => '50', 'default' => '50'),
		array('name' => 'health_rate', 'text' => '生效血量百分比', 'width' => '50'),
	),
	'opera' => array(),
	'insert' => 'sql_insert',
	'where' => 'sql_where',
);

function sql_insert($params) {
	if(!isset($params['m'])) {
		die("need params `m`");
	}
	return " `enemy_id`={$params['m']} ";
}

function sql_where($params) {
	if(!isset($params['m'])) {
		die("need params `m`");
	}
	return " where  `enemy_id`={$params['m']} ";
}
?>
