<?php


function sql_where($params){
	if (!isset($params['m'])){
		return '';
	}
	return "WHERE `mission_level_id` = {$params['m']}";
}

function sql_insert($params) {
	return "`mission_level_id` = {$params['m']}";
}

function add_item_reader($row){
	return html_get_link("添加宝箱物品", "?p=mission_level_small_box_items&m=".$row['id'], false);
}

function location($params){
	global $db;
	$html = '';

	$sql = "select `id`, `mission_id`, `name` from mission_level where id = {$params['m']}";
	$mission_level = db_query($db, $sql);
	if (count($mission_level) == 0) {
		return $html;
	}

	$sql = "select `id`, `name`, `town_id` from mission where id = {$mission_level[0]['mission_id']}";
	$mission = db_query($db, $sql);
	if (count($mission) == 0) {
		return $html;
	}
	
	
	$sql = "select `id`, `name` from town where id = {$mission[0]['town_id']}";
	$town = db_query($db, $sql);
	if (count($town) == 0) {
		return $html;
	}
	
	
	$html .= html_get_link("所有城镇", "?p=town", false).' -> ';
	$html .= html_get_link($town[0]['name'], "?p=mission&m=".$town[0]['id'], false).' -> ';
	$html .= html_get_link($mission[0]['name'], "?p=mission_level&m=".$mission[0]['id'], false) . '-> ';
	$html .= $mission_level[0]['name'];
	
	return $html;
}

function box_dir_render($row){
	global $types;
	return $types[$row['box_dir']];
}

function box_dir_editor($row){
	global $types;

	$code = '';
	if ($row != null) {
		$code .= '<select name="box_dir[]" >';
	foreach ($types as $key => $value){

		if ($key == $row['box_dir']) {
			$code .= '<option value="'.$key.'"  selected="selected">'.$value.'</option>';
		} else {
			$code .= '<option value="'.$key.'" >'.$value.'</option>';
		}
	}
		$code .= '</select>';
	} else {
		$code .= '<select name="box_dir[]" >';
		$code .= '<option value="0" selected="selected">朝左</option>';
		$code .= '<option value="1" >朝右</option>';
		$code .= '</select>';
	}
	
	return $code;
}

// 任务类型
$types[0] = '朝左';
$types[1] = '朝右';

$pconfig = array(
	'title'   => '关卡小宝箱',
	'table'   => 'mission_level_small_box',
	'links'   => array(),
	'where'	  => 'sql_where',
	'insert'  => 'sql_insert',
	'columns' => array(
		array('name' => 'box_x', 'text' => '宝箱x坐标', 'width' => '150px'),
		array('name' => 'box_y', 'text' => '宝箱y坐标', 'width' => '150px'),
		array('name' => 'probability', 'text' => '出现概率%', 'width' => '150px'),
		array('name' => 'box_dir', 'text' => '宝箱朝向', 'width' => '60px',
			'editor' => 'box_dir_editor',
			'render' => 'box_dir_render',
		),
		array('name' => 'items_kind', 'text' => '出现物品有几种', 'width' => '150px'),

	),

	'opera' => array(
		array('type' => 'link', 'text' => '添加宝箱物品', 'render' => 'add_item_reader'),
	),

	'location' => 'location',

);
?>
