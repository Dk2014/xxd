<?php
$extend_columns = array(
/*   '字段' => 配置 */
	'npc_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	'type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(/* 0 */'首次对话', /* 1 */'任务对话'),
	),
	'quest_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	'conversion' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	'award_item_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);
$townId = $_REQUEST['m'];
if(!$townId) {
	die("需要指定城镇ID");
}

$townNpc = array();
$sql = "select * from town_npc where town_id={$townId}";
$townNpcRows = db_query($db, $sql);

foreach($townNpcRows as $row) {
	$townNpc[$row['id']] = $row['name'];
}

function render_npc_id($column_name, $column_val, $row, $data) {
	global $townNpc;
	return $townNpc[$column_val];
}

function editor_npc_id($column_name, $column_val, $row, $data) {
	global $townNpc;
	return html_get_select($column_name,$townNpc,$column_val);
}

$quests = array();
$quests[0] = '无';
$sql = "select * from quest";
$questRows = db_query($db, $sql);
foreach($questRows as $row) {
	global $quests;
	$quests[$row['id']]=$row['name'];
}

function render_quest_id($column_name, $column_val, $row, $data) {
	global $quests;
	return $quests[$column_val];
}

function editor_quest_id($column_name, $column_val, $row, $data) {
	global $quests;
	return html_get_select($column_name,$quests,$column_val);
}


$sql = "select id, name from item where type_id !=10 and type_id!=11";
$allItemRows = db_query($db, $sql);
$allItems = array();
$allItems[0] = '无';
foreach($allItemRows as $row) {
	$allItems[$row['id']] = $row['name'];
}


function render_award_item_id($column_name, $column_val, $row, $data) {
	global $allItems;
	return normal_render($allItems, $row['award_item_id']);
}

function editor_award_item_id($column_name, $column_val, $row, $data){
	global $allItems;
	$item_name = '';
	$award_item_id = $row['award_item_id'];
	if (isset($allItems[$award_item_id])) {
		$item_name = $allItems[$award_item_id];
	}
	return editor_single_item($item_name, $award_item_id, "award_item_id");
}

function editor_conversion($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_conversion($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function render_type($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_type($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}


function location($params){
	global $db;
	$html = '';

	$sql = "select `id`, `name` from town where id = {$params['m']}";
	$town = db_query($db, $sql);
	if (count($town) == 0) {
		return $html;
	}
	$html .= '<a href="?p=town">所有城镇</a> -> '.$town[0]['name'];

	return $html;
}



function sql_where($params) {
	return "where `town_id`={$params['m']} ";
}

function sql_insert($params) {
	return "`town_id` = {$params['m']}";
}

function js_function($params) {
	global $allItems;

	$column_name = 'award_item_id';

	$html = get_items_json($allItems, 'items');

	$html .= <<<EOT

	autoopt_{$column_name} = {
		lookup: items, 
		minChars: 0,
		onSelect: function(s) {
			$(this).siblings('input[class="real_value"]').val(s.id);
		}
	};

	$(".display_{$column_name}").autocomplete(autoopt_{$column_name});

EOT;

	return $html;
}

?>
