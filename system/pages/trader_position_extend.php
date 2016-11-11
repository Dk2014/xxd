<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'town_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
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
);

function render_direction($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_direction($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

$sql = "select * from town";
$rows = db_query($db, $sql);
$allTown = array();
foreach($rows as $row) {
	$allTown[$row['id']] = $row['name'];
}

function render_town_id($column_name, $column_val, $row, $data) {
	global $allTown;
	return $allTown[$column_val];
}

function editor_town_id($column_name, $column_val, $row, $data) {
	global $allTown;
	return html_get_select($column_name,$allTown,$column_val);
}

function sql_where($params) {
	if (!isset($params['m'])){
		return '';
	}
	return "where `trader_id`={$params['m']}";
}

function sql_insert($params) {
	return "`trader_id` = {$params['m']}";
}

function location($params) {
	global $db;
	$sql = "select * from trader where id={$params['m']}";
	$trader = db_query($db, $sql)[0];

	$html = html_get_link("随机商人", "?p=trader", false) . "->";
	$html .= "{$trader['name']}";
	return $html;
}
?>
