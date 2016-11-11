<?php
$extend_columns = array(
/*   '字段' => 配置 */
	'money_type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('铜币', '元宝', '爱心', '龙币'),
	),
);

$moneyType = array('铜币', '元宝', '爱心', '龙币');
function render_money_type($column_name, $column_val, $row, $data) {
	global $moneyType;
	return $moneyType[$column_val];
}

function editor_money_type($column_name, $column_val, $row, $data) {
	global $moneyType;
	return html_get_select($column_name,$moneyType,$column_val);
}

function location($params) {
	global $db;
	$sql = "select * from trader where id={$params['m']}";
	$trader = db_query($db, $sql)[0];

	$html = html_get_link("随机商人", "?p=trader", false) . "->";
	$html .= "{$trader['name']}";
	return $html;
}



function sql_where($params) {
	return "where trader_id={$params['m']}";
}

function sql_insert($params) {
	return "trader_id={$params['m']}";
}

function trader_grid_config($row) {
	return html_get_link('随机货物列表', "?p=trader_grid_config&m={$row['id']}", true);
}

function trader_grid_config_list($row) {
	$text = "";
	global $db;
	$sql = "select * from trader_grid_config where grid_id={$row['id']}";
	$rows = db_query($db, $sql);
	if(!$rows) {
		return "未配置";
	}
	$probability = 0;

	$text .= "总概率";
	foreach($rows as $row) {
		$probability += intval($row['probability']);
	}
	return $text . "{$probability}";
}

?>
