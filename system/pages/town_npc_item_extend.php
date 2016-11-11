<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'town_npc_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),

	'item_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),

	'vip' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('普通','vip特供'),
	),
);

function sql_where($params) {
	return " WHERE `town_npc_id` in (SELECT `id` FROM `town_npc` WHERE `town_id`={$params['m']}) order by `town_npc_id`";
}

function location($params){
	global $currentTown;
	$html = '<a href="?p=town">所有城镇</a> -> '.$currentTown[$params['m']];
	return $html;
}

function render_town_npc_id($column_name, $column_val, $row, $data) {
	global $currentTown, $currentTownNpc;

	$key = $row['town_npc_id'];
	$town_id = $currentTownNpc[$key]['town_id'];

	return $currentTown[$town_id].' '.$currentTownNpc[$key]['name'];
}

function editor_town_npc_id($column_name, $column_val, $row, $data) {
	global $currentTown, $currentTownNpc;
	$name = 'town_npc_id';

	$html = '<select name="'.$name.'[]" >';
	foreach ($currentTownNpc as $key => $value) {
		$town_id = $value['town_id'];
		if (isset($row[$name]) && $row[$name] == $key) {
			$html .= '<option value="'.$key.'" selected="selected" >'.$currentTown[$town_id].' '.$value['name'].'</option>';
		} else {
			$html .= '<option value="'.$key.'">'.$currentTown[$town_id].' '.$value['name'].'</option>';
		}
	}
	$html .= '</select>';
	return $html;
}

function render_item_id($column_name, $column_val, $row, $data) {
	global $allItem;
	return $allItem[$column_val];
}

function editor_item_id($column_name, $column_val, $row, $data) {
	global $allItem;
	return html_get_select($column_name,$allItem,$column_val);
}

function render_vip($column_name, $column_val, $row, $data) {
	return $data[$column_val];
}

function editor_vip($column_name, $column_val, $row, $data) {
	return html_get_select($column_name,$data,$column_val);
}

$sql = "select `id`, `name` from `item`";
$allItemTmp = db_query($db, $sql);
foreach ($allItemTmp as $value)
{
	$allItem[$value['id']] = $value['name'];
}

$town_id = $_GET['m'];
$sql = "select `name` from `town` where `id`=$town_id";
$currentTownTmp = db_query($db, $sql);
foreach ($currentTownTmp as $value)
{
	$currentTown[$town_id] = $value['name'];
	break;
}

$sql = "select `id`, `town_id`, `name` from `town_npc` where `town_id`=$town_id";
$currentTownNpcTmp = db_query($db, $sql);
$currentTownNpc = array();
foreach ($currentTownNpcTmp as $value)
{
	$currentTownNpc[$value['id']] = array(
		'town_id' => $value['town_id'],
		'name'    => $value['name'],
	);
}

?>