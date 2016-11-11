<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'item_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);

$allType = array();
$sql = "select `id`, `name` from `item` where `type_id` in (select `id` from `item_type` where `name` in ('道具'));";
$allTypeTmp = db_query($db, $sql);

foreach ($allTypeTmp as $value)
{
	$allType[$value['id']] = $value['name'];
}

function render_item_id($column_name, $column_val, $row, $data) {
	global $allType;
	return $allType[$column_val];
}

function editor_item_id($column_name, $column_val, $row, $data) {
	global $allType;
	return html_get_select($column_name,$allType,$column_val);
}

function sql_where($params) {
	return " order by `realm_class`";
}

?>