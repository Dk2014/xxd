<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'privilege_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);

$privileges = array();
$sql = "select * from `vip_privilege`";
$privilegeRows = db_query($db, $sql);
foreach($privilegeRows as $row) {
	$privileges[$row['id']] = $row['name'];
}




function render_privilege_id($column_name, $column_val, $row, $data) {
	global $privileges;
	return $privileges[$column_val];
}

function editor_privilege_id($column_name, $column_val, $row, $data) {
	global $privileges;
	return html_get_select($column_name,$privileges,$column_val);
}

function sql_where($params) {
	if(!isset($params['m'])) {
		return "";
	}
	return " where `level`='{$params['m']}' ";
}

function sql_insert($params) {
	return " `level`={$params['m']} ";
}

function location($params) {
	$html = "";
	$html .= html_get_link("所有等级", "?p=vip_level", false) . "->";
	$html .= "{$params['m']}级";
	return $html;
}



?>
