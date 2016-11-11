<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'npc_role' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'		=> array('params' => array()),
	),
	'direction' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array("r" => "右",
						"rb" => "右下方",
						"b" => "下",
						"lb" => "左下方",
						"l" => "左",
						"lt" => "左上方",
						"t" => "上",
						"rt" => "右上方",
			),
	),

	'showup_quest' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),

	'disappear_quest' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);

$all_town = array();
$sql = "select id, name from town";
$rows = db_query($db, $sql);
foreach($rows as $row) {
	$all_town[$row['id']] = $row['name'];
}

$all_quest = array();
$all_quest[0] = "一直存在";
$all_quest[-1] = "无效";

foreach($all_town as $town_id=>$town_name) {
	$sql= " select * from quest q inner join (select l.id lid from  (select id,name from mission where town_id = {$town_id}) m  left join mission_level l on m.id = l.mission_id) m 
		on q.town_id = {$town_id}
		or q.mission_level_id = m.lid 
		group by q.id
		order by `town_id`, `order`
		";

	$rows = db_query($db, $sql);
	foreach($rows as $row) {
		$all_quest[$row['id']] = $town_name . $row['name'];
	}
}

function render_showup_quest($column_name, $column_val, $row, $data){
	global $all_quest;
	return normal_render($all_quest, $column_val);
}

function editor_showup_quest($column_name, $column_val, $row, $data){
	global $all_quest;
	$quest_id =-1 ;
	if(isset($row['showup_quest'])) {
		$quest_id =$row['showup_quest'] ;
	}
	$name = '无';
	if (isset($all_quest[$quest_id])) {
		$name = $all_quest[$quest_id];
	}
	return editor_single_item($name, $quest_id, "showup_quest");
}

function render_disappear_quest($column_name, $column_val, $row, $data){
	global $all_quest;
	return normal_render($all_quest, $column_val);
}

function editor_disappear_quest($column_name, $column_val, $row, $data){
	global $all_quest;
	$quest_id =-1 ;
	if(isset($row['disappear_quest'])) {
		$quest_id =$row['disappear_quest'] ;
	}
	$name = '无';
	if (isset($all_quest[$quest_id])) {
		$name = $all_quest[$quest_id];
	}
	return editor_single_item($name,  $quest_id, "disappear_quest");
}

function editor_npc_role($column_name, $column_val, $row, $data){
	global $allNPCs;
	$role_name = '';
	if (isset($allNPCs[$column_val])) {
		$role_name = $allNPCs[$column_val];
	}

	return editor_single_item($role_name, $column_val, $column_name);
}

function render_npc_role($column_name, $column_val, $row, $data){
	global $allNPCs;

	return (isset($allNPCs[$column_val]) ? $allNPCs[$column_val] : "0");
}

function render_direction($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_direction($column_name, $column_val, $row, $data){
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
	$order = ' ORDER BY town_id';
	if (!isset($params['m'])){
		return ''.$order;
	}
	return "where `town_id`={$params['m']}". $order;
}

function sql_insert($params) {
	return "`town_id` = {$params['m']}";
}

$allNPCs = get_all_npc_role();
//$all_quest = get_all_quest();

function jsFunction($params) {
	global $all_quest;
	global $allNPCs;



	$html = '';
	$html .= choose_single_item($allNPCs, 'npc_role');
	$html .= choose_single_item($all_quest, 'showup_quest');
	$html .= choose_single_item($all_quest, 'disappear_quest');

	return $html;
}
?>
