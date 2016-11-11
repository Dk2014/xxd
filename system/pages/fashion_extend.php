<?php
include "item_equipment_links.php";

$extend_columns = array(
/*   '字段' => 配置 */
	'desc' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);


function editor_desc($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_desc($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function fashion_exchange_opera($row) {
	return html_get_link("关联物品", "?p=fashion_exchange&m=".$row['id'], true);
}


?>
