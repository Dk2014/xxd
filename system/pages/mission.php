<?php


$sql = "select `id`, `name` from `town`;";
$allTown_tmp = db_query($db, $sql);

foreach ($allTown_tmp as $value){
	$allTown[$value['id']] = $value['name'];
}

//城镇
$sql = "select `id`,`name` from `town`";
$allTowns = db_query($db, $sql);
$i = 0;
$linksTown = array();
foreach ($allTowns as $value){
	$allTown[$value['id']] = $value['name'];
	$linksTown[$i]['text'] = $value['name'];
	$linksTown[$i]['id'] = "mission&m=".$value['id'];
	$i++;
}

$pconfig = array(
	'title'   => '区域数据',
	'table'   => 'mission',
	'links'   => $linksTown,
	'columns' => array(
		array('name' => 'name', 'text' => '区域名称', 'width' => '90px'),
		array('name' => 'order','text' => '区域开启顺序', 'width' => '90px'),
		array('name' => 'keys', 'text' => '开启钥匙数', 'width' => '90px'),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '90px'),
	),
	'opera' => array(
		array('type' => 'link', 'text' => '区域怪物', 'render' => 'mission_enemy_render'),
		),
	'where' => 'sql_where',
	'insert'  => 'sql_insert',
	'location' => 'location',
	'js' => 'js_function',
);
?>
