<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(/* 0 */'普通', /* 1 */'精英', /* 2 */'Boss'),
	),
	//TODO 选择奖励类型需要让 award_item 的可选项动态地修改
	'award_type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array( /* 0 */'道具', /* 1 */'装备'),
	),

	'award_item' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),


	'daily_num' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),

	'box_dir' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('朝左', '朝右'),
	),

	'flip_horizontal' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('否', '是'),
	),

	'award_box' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('否', '是'),
	),

	'award_pet' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(/* 0 */'否', /* 1 */'是'),
	),
);

function render_award_pet($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_award_pet($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_award_box($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_award_box($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_type($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_type($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_flip_horizontal($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_flip_horizontal($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_daily_num($column_name, $column_val, $row, $data){
	if ($column_val < 0) {
		return "不限";
	}

	return $column_val;
}

function editor_daily_num($column_name, $column_val, $row, $data){
	if ($column_val == 0)
		$column_val = -1;
	return html_get_input($column_name, $column_val);
}

function editor_box_dir($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_box_dir($column_name, $column_val, $row, $data){
	return $data[$column_val];
}


function map_enemy_render($row) {
	$html = html_get_link("怪物组", "?p=mission_enemy&m=".$row['id']."&bt=0", false);
	$html .= " | ";
	$html .= html_get_link("宝箱配置", "?p=mission_level_box&m=".$row['id'], false);
	$html .= " | ";
	$html .= html_get_link("小宝箱", "?p=mission_level_small_box&m=".$row['id'], false);
	$html .= " | ";
	$html .= html_get_link("关卡评星", "?p=level_star&m=" . $row['id'], false);
	$html .= " | ";
	$html .= html_get_link("关卡梦妖", "?p=mission_level_recovery_meng_yao&m=".$row['id'], false);
	$html .= " | ";
	$html .= html_get_link("关卡剧情", "?p=mission_level_drama&m=".$row['id'], false);
	$html .= " | ";
	$html .= html_get_link("影之间隙", "?p=shaded_mission&m=".$row['id'], false);
	return $html;
}

function sql_insert($params) {
	return "`mission_id` = {$params['m']}, `order`='0'";
}

function location($params){
	global $db;
	$html = '';

	$sql = "select `id`, `name`, `town_id` from mission where id = {$params['m']}";
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
	$html .= $mission[0]['name'];
	
	return $html;
}

function sql_where($params) {
	if (!isset($params['m'])){
		return '';
	}

	return " where `mission_id`={$params['m']} order by `lock` asc";
}

$all_items = get_all_item();
function range_award_item(){
	global $all_items;
	$all_items_temp = $all_items;
	$all_items_temp[0] = '无';
	return $all_items_temp;
}
function editor_award_item($column_name, $column_val, $row, $data){
	global $all_items;

	$item_name = '';
	if (isset($all_items[$column_val])) {
		$item_name = $all_items[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item($column_name, $column_val, $row, $data){
	global $all_items;

	return (isset($all_items[$column_val]) ? $all_items[$column_val] : "无");
}

function js_function($params) {
	global $all_items;

	return choose_single_item($all_items, 'award_item');
}

function render_award_type($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_award_type($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

?>
