<?php
include "item_equipment_links.php";

$extend_columns = array(
/*   '字段' => 配置 */
	'level' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
);

$allItem = array();
$temp = db_query($db,"select `id`,`name` from `item` where `type_id` in (3,4,5,6)");

foreach ($temp as $value)
{
	$allItem[-1*$value['id']] = $value['name'];
}

function range_level(){
	global $allItem;
	return $allItem;
}

function render_level($column_name, $column_val, $row, $data) {
	global $allItem;
	return $allItem[$column_val];
}

function editor_level($column_name, $column_val, $row, $data) {
	global $allItem;
	return html_get_select($column_name,$allItem,$column_val);
}

function sql_where($params) {
	return " where `level` < 0";
}

?>