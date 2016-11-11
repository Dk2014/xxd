<?php


$extend_columns = array(
/*   '字段' => 配置 */
	'award_type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(/* 0 */'铜钱', /* 1 */'道具', /* 2 */'装备', /* 3 */'经验', /* 4 */'经验倍数', /* 5 */'铜钱倍数', 
				/* 6 */'伙伴技能恢复',  /* 7 */ '魂侍技能恢复', /* 8 */'灵宠状态恢复',
				/* 9 */'增加精气', /* 10 */'百分比恢复血量', /* 11 */'增加魂力'),
	),
	'item_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range'=> array('params' => array()),
	),
);

$all_items = get_all_item();

function range_item_id(){
	global $all_items;
	$all_items_temp = $all_items;
	$all_items_temp[0] = '无';
	return $all_items_temp;
}
function editor_item_id($column_name, $column_val, $row, $data){
	global $all_items;

	$item_name = '';
	if (isset($all_items[$column_val])) {
		$item_name = $all_items[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_item_id($column_name, $column_val, $row, $data){
	global $all_items;

	return (isset($all_items[$column_val]) ? $all_items[$column_val] : "无");
}

function render_award_type($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_award_type($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function sql_insert($params) {
	return "`mission_level_id` = {$params['m']}, `autofight_box` = 1";
}


function sql_where($params) {
	if (!isset($params['m'])){
		return '';
	}

	return "where `mission_level_id`={$params['m']} order by `order` asc";
}

function js_function($params) {
	global $all_items;

	return choose_single_item($all_items, 'item_id');
}

?>
