<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'desc' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
	),
	
	'require_open_day' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
	),
	
	'class' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array('0'=>'上香', '1'=>'护送运镖', '2'=>'铜钱捐献', ),
	),
	'award_item1_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'		=> array('params' => array()),
	),
	'award_item2_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range' =>array('params' => array()),
	),
	'award_item3_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range' =>array('params' => array()),
	),
	'award_item4_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range' =>array('params' => array()),
	),

);

$allItem = get_all_item();

function render_require_open_day($column_name, $column_val, $row, $data){
	$html = "每日";
	$arr = array();
	
	if (strlen(trim($column_val)) > 0) {
		$arr = explode(" ", trim($column_val));
	}

	if (count($arr) > 0) {
		$list = array();
		foreach ($arr as $val) {
			switch ($val) {
				case '0':
					$list[] = '周日';
					break;
				case '1':
					$list[] = '周一';
					break;
				case '2':
					$list[] = '周二';
					break;
				case '3':
					$list[] = '周三';
					break;
				case '4':
					$list[] = '周四';
					break;
				case '5':
					$list[] = '周五';
					break;
				case '6':
					$list[] = '周六';
					break;
			}
		}

		$html = implode("、", $list);
	}
	return $html;
}

function editor_require_open_day($column_name, $column_val, $row, $data){
	return html_get_input($column_name, $column_val, 80);
}

function render_class($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_class($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}
/*
function render_level_type($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_level_type($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_level_sub_type($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_level_sub_type($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}
*/
function editor_award_item1_id($column_name, $column_val, $row, $data){
	global $allItem;

	$item_name = '';
	if (isset($allItem[$column_val])) {
		$item_name = $allItem[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item1_id($column_name, $column_val, $row, $data){
	global $allItem;

	return (isset($allItem[$column_val]) ? $allItem[$column_val] : "0");
}
function editor_award_item2_id($column_name, $column_val, $row, $data){
	global $allItem;

	$item_name = '';
	if (isset($allItem[$column_val])) {
		$item_name = $allItem[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item2_id($column_name, $column_val, $row, $data){
	global $allItem;

	return (isset($allItem[$column_val]) ? $allItem[$column_val] : "无");
}

function editor_award_item3_id($column_name, $column_val, $row, $data){
	global $allItem;

	$item_name = '';
	if (isset($allItem[$column_val])) {
		$item_name = $allItem[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item3_id($column_name, $column_val, $row, $data){
	global $allItem;

	return (isset($allItem[$column_val]) ? $allItem[$column_val] : "无");
}
function editor_award_item4_id($column_name, $column_val, $row, $data){
	global $allItem;

	$item_name = '';
	if (isset($allItem[$column_val])) {
		$item_name = $allItem[$column_val];
	}

	return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item4_id($column_name, $column_val, $row, $data){
	global $allItem;

	return (isset($allItem[$column_val]) ? $allItem[$column_val] : "无");
}
function range_award_item1_id(){
	global $allItem;
	$all_item_temp = $allItem;
	$all_item_temp[0] = '无';
	return $all_item_temp;
}
function range_award_item2_id(){
	global $allItem;
	$all_item_temp = $allItem;
	$all_item_temp[0] = '无';
	return $all_item_temp;
}
function range_award_item3_id(){
	global $allItem;
	$all_item_temp = $allItem;
	$all_item_temp[0] = '无';
	return $all_item_temp;
}
function range_award_item4_id(){
	global $allItem;
	$all_item_temp = $allItem;
	$all_item_temp[0] = '无';
	return $all_item_temp;
}

function editor_desc($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_desc($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function sql_where($params) {
	return " order by `order` ASC ";
}


function jsFunction($params) {
	global $allItem;

	$html = '
$("select").change( function(){
		$(this).css("color","#0033FF");
});';

	$html .= choose_single_item($allItem, 'award_item1_id');
	$html .= choose_single_item($allItem, 'award_item2_id');
	$html .= choose_single_item($allItem, 'award_item3_id');
	$html .= choose_single_item($allItem, 'award_item4_id');
	return $html;
}
?>
